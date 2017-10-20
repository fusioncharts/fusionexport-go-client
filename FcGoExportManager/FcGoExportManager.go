package FcGoExportManager

import (
    "net"
    "fmt"
    "bufio"
    "encoding/json"
    "io"
)

type OutFileBag struct {
    RealName string `json:"realName"`
    TmpPath string `json:"tmpPath"`
}

var conn net.Conn

func Connect(host, port string) {
    var err error

    address := host + ":" + port
    conn, err = net.Dial("tcp", address)
    check(err)
}

func Export (exportConfig string) []OutFileBag {
    data := emitData("ExportManager", "export", exportConfig)

    var outFileBagData map[string][]OutFileBag
    err := json.Unmarshal([]byte(data), &outFileBagData)
    check(err)

    return outFileBagData["data"]
}

func emitData (target, method, body string) string {
    payload := target + "." + method + "<=:=>" + body
    fmt.Fprint(conn, payload)

    out, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil && err != io.EOF {
        check(err)
    }

    return string(out)
}

func check(err error) {
    if err != nil {
        fmt.Print(err)
        panic(err)
    }
}