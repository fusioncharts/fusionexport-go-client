package main

import (
    "io/ioutil"
    "./FusionExport"
    "path/filepath"
    "net/http"
)

func saveFiles(fileBag []ExportFusion.OutFileBag) {
    for _, file := range fileBag {
        fileData, err := ioutil.ReadFile(file.TmpPath)
        check(err)
        err = ioutil.WriteFile(file.RealName, fileData, 0644)
        check(err)
    }
}

func main() {
    chartConfig, _ := ioutil.ReadFile("chart_config.json")

    templateFilePath, err := filepath.Abs("template.html")
    check(err)

    exportConfig := `{
      "chartConfig":` + string(chartConfig) + `,
      "templateFilePath":"` + templateFilePath + `"
    }`

    ExportFusion.Connect("127.0.0.1", "1337")
    outFileBag := ExportFusion.Export(exportConfig)

    saveFiles(outFileBag)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
