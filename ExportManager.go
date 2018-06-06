package FusionExport

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ExportManager struct {
	Host string
	Port int
}

func NewExportManager() ExportManager {
	em := ExportManager{
		Host: DEFAULT_HOST,
		Port: DEFAULT_PORT,
	}
	return em
}

func (em *ExportManager) SetConnectionConfig(host string, port int) {
	em.Host = host
	em.Port = port
}

func (em *ExportManager) Export(exportConfig ExportConfig, exportDone func([]OutFileBag, error), exportStateChanged func(ExportEvent)) (Exporter, error) {
	exp := Exporter{
		ExportConfig:              exportConfig,
		ExportDoneListener:        exportDone,
		ExportStateChangeListener: exportStateChanged,
		ExportServerHost:          em.Host,
		ExportServerPort:          em.Port,
	}

	err := exp.Start()

	return exp, err
}

func SaveExportedFiles(fileBag []OutFileBag) error {
	for _, file := range fileBag {
		fileData, err := base64.StdEncoding.DecodeString(file.FileContent)
		if err != nil {
			return err
		}

		dir := filepath.Dir(file.RealName)
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(file.RealName, fileData, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetExportedFileNames(fileBag []OutFileBag) []string {
	var fns []string

	for _, file := range fileBag {
		fns = append(fns, file.RealName)
	}

	return fns
}
