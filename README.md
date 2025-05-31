# tochka

[Tochka API](https://enter.tochka.com/doc/v2/) client written in Go.

> **WARNING:**
> This is a work in progress. It is not battle tested. Use at your own risk.

## Get

```sh
go get github.com/sharpvik/tochka
```

## Use

```go
package main

import (
    "github.com/sharpvik/tochka"
    "github.com/sharpvik/tochka/dto"
)

func main() {
    client := tochka.Live("YOUR JWT TOKEN")
    params := dto.CreateInvoiceParams{/* ... */}
    result, err := client.CreateInvoice(params)
    // ...
}
```

## Test

```go
package some_test

import (
    "github.com/sharpvik/tochka"
    "github.com/sharpvik/tochka/dto"
    "github.com/stretchr/testify/assert"
)

func TestTochkaCreateInvoice(t *testing.T) {
    client := tochka.Sandbox()
    params := dto.CreateInvoiceParams{/* ... */}
    _, err := client.CreateInvoice(params)
    assert.NoError(t, err)
}
```
