package FusionExport

type ExportManager struct {
    Host string
    Port int
}

func NewExportManager () ExportManager {
    em := ExportManager{
        Host: DEFAULT_HOST,
        Port: DEFAULT_PORT,
    }
    return em
}

func (em *ExportManager) setConnectionConfig (host string, port int) {
    em.Host = host
    em.Port = port
}

func (em *ExportManager) Export (exportConfig ExportConfig, exportDone func([]OutFileBag, error), exportStateChanged func(ExportEvent)) (Exporter, error) {
    exp := Exporter{
        ExportConfig: exportConfig,
        ExportDoneListener: exportDone,
        ExportStateChangeListener: exportStateChanged,
        ExportServerHost: em.Host,
        ExportServerPort: em.Port,
    }

    err := exp.Start()

    return exp, err
}
