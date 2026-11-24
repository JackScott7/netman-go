package main

import (
    "fmt"
    "net/http"
    "os"

    "github.com/akamensky/argparse"
)

var client = &http.Client{}

func main() {
    auth := os.Getenv("ROUTER_AUTH")
    if auth == "" {
        panic("ROUTER_AUTH environment variable not set")
    }

    parser := argparse.NewParser(
        "netman",
        "A small CLI tool that lets you retrieve info about your 'TP-Link Archer D2' router",
    )

    restartFlag := parser.Flag(
        "r",
        "restart",
        &argparse.Options{Required: false, Help: "Restart the router"},
    )
    clientsFlag := parser.Flag(
        "c",
        "clients",
        &argparse.Options{Required: false, Help: "Returns Connect clients to the router"},
    )
    logFlag := parser.Flag(
        "l",
        "log",
        &argparse.Options{Required: false, Help: "Output logs of the router"},
    )

    err := parser.Parse(os.Args)
    if err != nil {
        fmt.Print(parser.Usage(err))
        os.Exit(1)
    }

    if *clientsFlag {
        data := getConnectedRouterClients(auth)
        parsed := parseRouterClients(data)
        if err := prettyPrintJSON(parsed); err != nil {
            panic(err)
        }
    } else if *logFlag {
        data := getRouterLog(auth)
        fmt.Println(data)
    } else if *restartFlag {
        fmt.Println("Not implemented yet")
    } else {
        data := getTplinkDslData(auth)
        parsed := parseRouterDslResponse(data)
        if err := prettyPrintJSON(parsed); err != nil {
            panic(err)
        }
    }
}
