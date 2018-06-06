// Adding a logo or heading to the dashboard

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

	exportConfig.Set("chartConfig", "example/resources/multiple.json")
	exportConfig.Set("templateFilePath", "example/resources/template.html")
	exportConfig.Set("dashboardLogo", "example/resources/logo.jpg")
	exportConfig.Set("dashboardHeading", "FusionCharts")
	exportConfig.Set("dashboardSubheading", "The best charting library in the world")

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
