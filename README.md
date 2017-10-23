# FusionCharts Go Export Client

This is a Go Export Client for FC Export Service. It communicates with Export Service through the socket protocol and does the export.

## Installation

```
go get github.com/fc-export-go-client
```

## Usage

The `FcGoExportManager` exposes two simple interface to communicate with Export Service.

`Connect(host, port string)`

It takes the host and port of Export Service as string and sets up the connection to it.

`Export(exportConfig string) []OutFileBag`

It takes the exportConfig as a JSON string format.

It returns an array of OutFileBag which contains the temporary file and the resolved name of that file as specified in the `output-file` option of the config.
