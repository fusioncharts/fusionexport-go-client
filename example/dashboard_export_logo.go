// Adding a logo or heading to the dashboard

package main

import (
    "io/ioutil"
    "../FusionExport"
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

func onDone (outFileBag []FusionExport.OutFileBag, err error) {
    check(err)
    saveFiles(outFileBag)
}

func onStateChange (event FusionExport.ExportEvent) {
    fmt.Println("[" + event.Reporter + "] " + event.CustomMsg)
}

func main() {
    exportConfig := FusionExport.NewExportConfig()

    chartConfig, err := ioutil.ReadFile("multiple.json")
    check(err)
    exportConfig.Set("chartConfig", string(chartConfig))

    templateFilePath, err := filepath.Abs("template.html")
    check(err)
    exportConfig.Set("templateFilePath", templateFilePath)

    logoFilePath, err := filepath.Abs("logo.jpg")
    check(err)
    exportConfig.Set("dashboardLogo", logoFilePath);

    exportConfig.Set("dashboardHeading", "FusionCharts");
    exportConfig.Set("dashboardSubheading", "The best charting library in the world");

    exportManager := FusionExport.NewExportManager()

    exportManager.Export(exportConfig, onDone, onStateChange)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
