package main

import (
	"fmt"
	"github.com/fusioncharts/fusionexport-go-client"
)

func onDone(outFileBag []FusionExport.OutFileBag, err error) {
	check(err)
	FusionExport.SaveExportedFiles(outFileBag)
}

func onStateChange(event FusionExport.ExportEvent) {
	fmt.Println("[" + event.Reporter + "] " + event.CustomMsg)
}

func main() {
	//svg, _ := filepath.Abs("../resources/vector.svg")
	//exportConfig.Set("inputSVG", svg)
	//
	//exp := FusionExport.Exporter{
	//	ExportConfig:     exportConfig,
	//	ExportServerHost: "0.0.0.0",
	//	ExportServerPort: 1337,
	//}
	//
	//exp.Start()

	//p, err := os.Getwd()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(p)
	//
	//tb := FusionExport.TemplateBundler{
	//	Template: "example/resources/template.html",
	//}
	//
	//tb.Process()

	//file, err := os.Open("example/resources/template.html")
	//check(err)
	//
	//hp := FusionExport.HTMLParser{Data: file}
	//nodes, err := hp.GetElemetsByTagName("img")
	//check(err)
	//
	//for _, node := range nodes {
	//	val := node.GetAttribute("href")
	//	fmt.Println(val)
	//}

	//tb := FusionExport.TemplateBundler{
	//	Template:  "example/resources/template.html",
	//	Resources: "example/resources/resource.json",
	//}
	//
	//err := tb.Process()
	//check(err)
	//
	//resFilePath := tb.GetResourcesZip()
	//
	//copy(resFilePath, "/Users/jimutdhali/Desktop/res.zip")

	exportConfig := FusionExport.NewExportConfig()

	err := exportConfig.Set("chartConfig", "example/resources/multiple.json")
	check(err)

	err = exportConfig.Set("templateFilePath", "example/resources/template.html")
	check(err)

	exportManager := FusionExport.NewExportManager()
	_, err = exportManager.Export(exportConfig, onDone, onStateChange)
	check(err)
}

//func copy(src string, dst string) {
//	data, err := ioutil.ReadFile(src)
//	check(err)
//
//	file, err := os.Create(dst)
//	check(err)
//	defer file.Close()
//
//	err = ioutil.WriteFile(dst, data, 0644)
//	check(err)
//}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
