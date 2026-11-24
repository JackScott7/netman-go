# netman-go

A small Go CLI for retrieving DSL and LAN information from a **TP-Link Archer D2** router.

## Features

- Display DSL status, modulation, Annex type, sync and attainable rates, SNR margins, line attenuation, and CRC error counts.
- List connected LAN clients, including IP and MAC addresses, hostnames, lease information, and connection status.
- Print the router log.
- Restart the router.

## Requirements

- Go 1.27 or later.
- A TP-Link Archer D2 router reachable at `http://192.168.1.1`.
- The value of the router session's `Authorization` cookie.

## Run

Set the `ROUTER_AUTH` environment variable before running the CLI.

**PowerShell**

```powershell
$env:ROUTER_AUTH = "<authorization-cookie-value>"
go run .
```

**Bash**

```bash
export ROUTER_AUTH="<authorization-cookie-value>"
go run .
```

## Commands

| Command | Description |
| --- | --- |
| `go run .` | Display DSL statistics (default). |
| `go run . --clients` | List connected LAN clients as JSON. |
| `go run . --log` | Print the router log. |
| `go run . --restart` | Display DSL statistics, then restart the router. |

Short flags are also available: `-c`, `-l`, and `-r`.

## Configuration

The router address is currently hard-coded to `192.168.1.1`. The `ROUTER_AUTH` variable should contain the value of the router session's `Authorization` cookie (not the full `Cookie` header).

Keep this value private and do not commit it to the repository.
