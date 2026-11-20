package main

import (
    "fmt"
    "net/http"
    "os"
)

var client = &http.Client{}


func main() {
    const addr = "http://192.168.1.1"
    var auth = os.Getenv("ROUTER_AUTH")

    if auth == "" {
        panic("ROUTER_AUTH environment variable not set")
    }

    data := getTplinkDslData(addr, auth)
    parsed := parseRouterDslResponse(data)

    if err := prettyPrintJSON(parsed); err != nil {
        panic(err)
    }
    log := getRouterLog(addr,auth)
    fmt.Println("\nRouter logs:\n", log)
}
