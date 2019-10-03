Add default header to the http client

## Installation
`go get github.com/Napas/go-default-headers-round-tripper`

of 

`dep ensure -v -add github.com/Napas/go-default-headers-round-tripper`

## Usage
```go
package main

import (

headers "github.com/Napas/go-default-headers-round-tripper"
"net/http"
)

func main() {
    c := http.DefaultClient
    c.Transport = headers.NewDefaultHeaders(c.Transport, map[string]string{
        "Header": "Header value",
    })
}
```