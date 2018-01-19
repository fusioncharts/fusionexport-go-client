package main

import (
	"fmt"
)

func main() {
	//exportConfig := FusionExport.NewExportConfig()
    //
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


	b := "true"
	getType(b)
}

func getType (d interface{}) {
	fmt.Println(d.(bool))
}
