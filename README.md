### mycli - Golang CLI application - Search for IP and name server

A simple CLI (Command line interface) application written in Go (Golang) to search
for IP and name servers of an URL.

### Dependencies

You will need to [install](https://go.dev/doc/install) Golang in your enviroment.

An external package is required:
``
$ go get github.com/urfave/cli
``
### How to run
Searching for name servers:

```
$ go run main.go server --host amazon.com
==============================
SEARCH FOR IP AND NAME SERVER
==============================
ns1.p31.dynect.net.
ns2.p31.dynect.net.
ns3.p31.dynect.net.
ns4.p31.dynect.net.
pdns1.ultradns.net.
pdns6.ultradns.co.uk.
```

Searching for IPs:

```
$ go run main.go ip --host amazon.com
==============================
SEARCH FOR IP AND NAME SERVER
==============================
176.32.103.205
54.239.28.85
205.251.242.103
```