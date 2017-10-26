package main

import (
    "io/ioutil"
    "./FcGoExportManager"
    "path/filepath"
)

func saveFiles(fileBag []FcGoExportManager.OutFileBag) {
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

    FcGoExportManager.Connect("127.0.0.1", "1337")
    outFileBag := FcGoExportManager.Export(exportConfig)

    saveFiles(outFileBag)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
