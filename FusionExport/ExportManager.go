package FusionExport

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

type ExportManager struct {
    Host string
    Port string
}

var conn net.Conn

func connect(host, port string) {
    var err error

    address := host + ":" + port
    conn, err = net.Dial("tcp", address)
    check(err)
}

func (em *ExportManager) Export (exportConfig string, exportDone func([]OutFileBag), exportStateChanged func()) {
    connect(em.Host, em.Port)

    data := emitData("ExportManager", "export", exportConfig)

    var outFileBagData map[string][]OutFileBag
    err := json.Unmarshal([]byte(data), &outFileBagData)
    check(err)

    exportDone(outFileBagData["data"])
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
