// Async capture

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

    chartConfig, err := ioutil.ReadFile("single.json")
    check(err)
    exportConfig.Set("chartConfig", string(chartConfig))

    callbackFilePath, err := filepath.Abs("expand_scroll.js")
    check(err)
    exportConfig.Set("callbackFilePath", callbackFilePath)

    exportConfig.Set("asyncCapture", "true")

    exportManager := FusionExport.NewExportManager()

    exportManager.Export(exportConfig, onDone, onStateChange)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
