package FusionExport

import (
    "net"
    "strconv"
    "fmt"
    "bufio"
    "io"
    "strings"
    "encoding/json"
    "errors"
)

type Exporter struct {
    ExportDoneListener func([]OutFileBag, error)
    ExportStateChangeListener func(ExportEvent)
    ExportConfig ExportConfig
    ExportServerHost string
    ExportServerPort int
    tcpClient net.Conn
}

type ExportEvent struct {
    Reporter string `json:"reporter"`
    Weight int `json:"weight"`
    CustomMsg string `json:"customMsg"`
    Uuid string `json:"uuid"`
}

type OutFileBag struct {
    RealName string `json:"realName"`
    TmpPath string `json:"tmpPath"`
}

func (exp *Exporter) Start () error {
    return exp.handleSocketConnection()
}

func (exp *Exporter) Cancel () error {
    return exp.tcpClient.Close()
}

func (exp *Exporter) handleSocketConnection () error {
    var err error

    address := exp.ExportServerHost + ":" + strconv.Itoa(exp.ExportServerPort)
    exp.tcpClient, err = net.Dial("tcp", address)
    if err != nil { return err }

    _, err = fmt.Fprint(exp.tcpClient, exp.getFormattedExportConfigs())
    if err != nil { return err }

    var data string
    for {
        out, err := bufio.NewReader(exp.tcpClient).ReadString('\n')
        if err != nil && err != io.EOF { return err }

        if len(out) < 1 { break }

        data += string(out)
        data = exp.processDataReceived(data)
    }

    return exp.tcpClient.Close()
}

func (exp *Exporter) processDataReceived (data string) string {
    parts := strings.Split(data, UNIQUE_BORDER)

    for _, part := range parts {
        if strings.HasPrefix(part, EXPORT_EVENT) {
            exp.processExportStateChangedData(part)
        } else if strings.HasPrefix(part, EXPORT_DATA) {
            exp.processExportDoneData(part)
        }
    }

    return parts[len(parts) - 1]
}

func (exp *Exporter) processExportStateChangedData (data string) {
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

func (exp *Exporter) processExportDoneData (data string) {
    exportResult := strings.TrimLeft(data, EXPORT_DATA)

    var outFileBagData map[string][]OutFileBag
    err := json.Unmarshal([]byte(exportResult), &outFileBagData)
    warn(err)

    exp.onExportDone(outFileBagData["data"], nil)
}

func (exp *Exporter) checkExportError (data string) error {
    var exportError map[string]string
    json.Unmarshal([]byte(data), &exportError)

    if val, ok := exportError["error"]; ok {
        return errors.New(val)
    }

    return nil
}

func (exp *Exporter) onExportStateChanged (event ExportEvent) {
    exp.ExportStateChangeListener(event)
}

func (exp *Exporter) onExportDone(bag []OutFileBag, err error) {
    exp.ExportDoneListener(bag, err)
}

func (exp *Exporter) getFormattedExportConfigs () string {
    return fmt.Sprintf("%s.%s<=:=>%s", "ExportManager", "export", exp.ExportConfig.GetFormattedConfigs())
}

func warn (err error) {
    if err != nil {
        fmt.Println(err)
    }
}

