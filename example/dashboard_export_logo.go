// Adding a logo or heading to the dashboard

package main

import (
    "io/ioutil"
    "../FusionExport" // import the sdk
    "path/filepath"
    "fmt"
)

func saveFiles(fileBag []FusionExport.OutFileBag) {
    for _, file := range fileBag {
        fmt.Println(file.RealName)
        fileData, err := ioutil.ReadFile(file.TmpPath)
        check(err)
        err = ioutil.WriteFile(file.RealName, fileData, 0644)
        check(err)
    }
}

// Called when export is done
func onDone (outFileBag []FusionExport.OutFileBag, err error) {
    check(err)
    saveFiles(outFileBag)
}

// Called on each export state change
func onStateChange (event FusionExport.ExportEvent) {
    fmt.Println("[" + event.Reporter + "] " + event.CustomMsg)
}

func main() {
    // Instantiate ExportConfig and add the required configurations
    exportConfig := FusionExport.NewExportConfig()

    chartConfig, err := ioutil.ReadFile("resources/multiple.json")
    check(err)
    exportConfig.Set("chartConfig", string(chartConfig))

    templateFilePath, err := filepath.Abs("resources/template.html")
    check(err)
    exportConfig.Set("templateFilePath", templateFilePath)

    logoFilePath, err := filepath.Abs("resources/logo.jpg")
    check(err)
    exportConfig.Set("dashboardLogo", logoFilePath);

    exportConfig.Set("dashboardHeading", "FusionCharts");
    exportConfig.Set("dashboardSubheading", "The best charting library in the world");

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
