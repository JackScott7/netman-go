package main

import (
	"net/http"
	"os"
)

var client = &http.Client{}

const ROUTERADDRESS = "http://192.168.1.1"

var routerAuth = os.Getenv("ROUTER_AUTH")


func main() {
	if routerAuth == "" {
		panic("ROUTER_AUTH environment variable not set")
	}

	data := getTplinkDslData()
	parsed := parseRouterDslResponse(data)

	if err := prettyPrintJSON(parsed); err != nil {
		panic(err)
	}
}