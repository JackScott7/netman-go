package main

import (
    "net/http"
    "os"
)

var client = &http.Client{}

func main() {
    var auth = os.Getenv("ROUTER_AUTH")

    if auth == "" {
        panic("ROUTER_AUTH environment variable not set")
    }

    data := getTplinkDslData(auth)
    parsed := parseRouterDslResponse(data)

    if err := prettyPrintJSON(parsed); err != nil {
        panic(err)
    }
}
