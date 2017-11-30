# FusionCharts Go Export Client

This is a Go Export Client for FC Export Service. It communicates with Export Service through the socket protocol and does the export.

## Installation

```
go get github.com/fusioncharts/fc-export-go-client
```

## Usage

Everything is stored in the `FusionExport` package.

You can use the `ExportConfig` class to build the export config for each export.

Build a simple export config

```go
exportConfig := FusionExport.NewExportConfig()
exportConfig.Set("chartConfig", chartConfig)
```

Use the `ExportManager` class to export multiple charts.

```go
exportManager := FusionExport.NewExportManager()
exportManager.Export(exportConfig, onDone, onStateChange)
```

The format of the `Export` function is

```go
func (em *ExportManager) Export (exportConfig ExportConfig, exportDone func([]OutFileBag, error), exportStateChanged func(ExportEvent)) (Exporter, error)
```

`exportDone` callback gets an array of OutFileBag which contains the temporary file and the resolved name of that file as specified in the `output-file` option of the config.

`exportStateChanged` callback gets an ExportEvent which contains the state of the export on each progress event.
