// Changing the export type

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
	exportConfig.Set("type", "pdf")

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
