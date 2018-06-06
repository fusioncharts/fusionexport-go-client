# FusionExport Go Client

This is a Go export client for FusionExport. It communicates with FusionExport through the socket protocol and does the export.

## Installation

To install the Go package, simply use Go get:

```
go get github.com/fusioncharts/fusionexport-go-client
```

## Usage

To require this into your project:

```go
import "github.com/fusioncharts/fusionexport-go-client"
```

## Getting Started

Let’s start with a simple chart export. For exporting a single chart, save the chartConfig in a JSON file. The config should be inside an array.

```go
// Exporting a chart

package main

import (
	"fmt"

	"github.com/fusioncharts/fusionexport-go-client"
)

// Called when export is done
func onDone(outFileBag []FusionExport.OutFileBag, err error) {
	check(err)
	FusionExport.SaveExportedFiles(outFileBag)
}

// Called on each export state change
func onStateChange(event FusionExport.ExportEvent) {
	fmt.Println("[" + event.Reporter + "] " + event.CustomMsg)
}

func main() {
	// Instantiate ExportConfig and add the required configurations
	exportConfig := FusionExport.NewExportConfig()

	exportConfig.Set("chartConfig", "example/resources/single.json")

	// Instantiate ExportManager
	exportManager := FusionExport.NewExportManager()
	// Call the Export() method with the export config and the respective callbacks
	exportManager.Export(exportConfig, onDone, onStateChange)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
```

## API Reference

You can find the full reference [here](https://www.fusioncharts.com/dev/exporting-charts/using-fusionexport/sdk-api-reference/golang.html).