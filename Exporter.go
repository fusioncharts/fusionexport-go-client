package FusionExport

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"golang.org/x/net/websocket"
)

type Exporter struct {
	ExportDoneListener        func([]OutFileBag, error)
	ExportStateChangeListener func(ExportEvent)
	ExportConfig              ExportConfig
	ExportServerHost          string
	ExportServerPort          int
	wsClient                  *websocket.Conn
}

type ExportEvent struct {
	Reporter  string `json:"reporter"`
	Weight    int    `json:"weight"`
	CustomMsg string `json:"customMsg"`
	Uuid      string `json:"uuid"`
}

type OutFileBag struct {
	RealName string `json:"realName"`
	TmpPath  string `json:"tmpPath"`
}

func (exp *Exporter) Start() error {
	return exp.handleSocketConnection()
}

func (exp *Exporter) Cancel() error {
	return exp.wsClient.Close()
}

func (exp *Exporter) handleSocketConnection() error {
	var err error
	var data string

	server := "ws://" + exp.ExportServerHost + ":" + strconv.Itoa(exp.ExportServerPort)
	origin := "http://localhost/"

	wsConfig, err := websocket.NewConfig(server, origin)
	if err != nil {
		return err
	}

	exp.wsClient, err = websocket.DialConfig(wsConfig)
	if err != nil {
		return err
	}

	payload := exp.getFormattedExportConfigs()

	err = websocket.Message.Send(exp.wsClient, payload)
	if err != nil {
		return err
	}

	for {
		err = websocket.Message.Receive(exp.wsClient, &data)

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		exp.processDataReceived(data)
	}

	return exp.wsClient.Close()
}

func (exp *Exporter) processDataReceived(data string) {
	if strings.HasPrefix(data, EXPORT_EVENT) {
		exp.processExportStateChangedData(data)
	} else if strings.HasPrefix(data, EXPORT_DATA) {
		exp.processExportDoneData(data)
	}
}

func (exp *Exporter) processExportStateChangedData(data string) {
	state := strings.TrimLeft(data, EXPORT_EVENT)
	exportError := exp.checkExportError(state)

	if exportError != nil {
		exp.onExportDone(nil, exportError)
	}

	var exportEvent ExportEvent
	err := json.Unmarshal([]byte(state), &exportEvent)
	warn(err)

	exp.onExportStateChanged(exportEvent)
}

func (exp *Exporter) processExportDoneData(data string) {
	exportResult := strings.TrimLeft(data, EXPORT_DATA)

	var outFileBagData map[string][]OutFileBag
	err := json.Unmarshal([]byte(exportResult), &outFileBagData)
	warn(err)

	exp.onExportDone(outFileBagData["data"], nil)
}

func (exp *Exporter) checkExportError(data string) error {
	var exportError map[string]string
	json.Unmarshal([]byte(data), &exportError)

	if val, ok := exportError["error"]; ok {
		return errors.New(val)
	}

	return nil
}

func (exp *Exporter) onExportStateChanged(event ExportEvent) {
	exp.ExportStateChangeListener(event)
}

func (exp *Exporter) onExportDone(bag []OutFileBag, err error) {
	exp.ExportDoneListener(bag, err)
}

func (exp *Exporter) getFormattedExportConfigs() string {
	return fmt.Sprintf("%s.%s<=:=>%s", "ExportManager", "export", exp.ExportConfig.GetFormattedConfigs())
}

func warn(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
