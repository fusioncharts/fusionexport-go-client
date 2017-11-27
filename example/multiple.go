// Manipulating the output filename

package main

import (
    "io/ioutil"
    "../FusionExport"
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
    exportConfig.Set("exportFile", "go-export-<%= number(5) %>")

    exportManager := FusionExport.NewExportManager()

    exportManager.Export(exportConfig, onDone, onStateChange)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
