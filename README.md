# go-mgt-client

Golang client for MGT REST service

## Use

To use in a project you can do something similar to the example:

```go
package main

import (
    "fmt"
    mgt "github.com/cpurta/go-mgt-client"
)

func main() {
    var (
        host = "http://127.0.0.1:8082"
        license = "fakeLicense"
        client = mgt.NewClient(http.DefaultClient, host, license)
        promotions []*mgt.PlayerPromotions
        err error
    )

    if promotions, err = client.GetPromotions("fakeUUID"); err != nil {
        fmt.Println("unable to get promotions", err.Error())
        return
    }

    fmt.Println("Successfully got player promotions", promotions)
}
```
