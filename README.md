# gody

:balloon: **A lightweight struct validator for Go**

[![GoDoc](https://godoc.org/github.com/guiferpa/gody?status.svg)](https://godoc.org/github.com/guiferpa/gody)
[![Go Report Card](https://goreportcard.com/badge/github.com/guiferpa/gody)](https://goreportcard.com/report/github.com/guiferpa/gody)
[![Build Status](https://cloud.drone.io/api/badges/guiferpa/gody/status.svg)](https://cloud.drone.io/guiferpa/gody)

## Quick start

```go
package main

import (
	"log"

	"github.com/guiferpa/gody"
	"github.com/guiferpa/gody/rule"
)

type Body struct {
	Text string `json:"text" validate:"required=true"`
}

func main() {
	b := Body{}

	if valid, err := gody.Validate(b, nil); err != nil {
		if !valid {
			log.Println("body didn't validate", err)
			return
		}

		switch err.(type) {
		case *rule.ErrRequired:
			log.Println("required error", err)
			break
		}
	}
}
```

### Deep validation

```go
package main

import (
    "log"

    "github.com/guiferpa/gody"
    "github.com/guiferpa/gody/rule"
)

type Price struct {
	Currency string `json:"currency" validate:"enum=BRL,EUR,USD"`
	Value    int    `json:"value" validate:"min=10"`
}

type ItemProduct struct {
	Amount int `json:"amount" validate:"min=1"`

	// validate tag's necessary for validation works if not setted it'll be ignored
	Price Price `json:"price" validate:"required=true"`
}

func main() {
	ip := ItemProduct{Amount: 10, Price: Price{"BYN", 10000}}

	if valid, err := gody.Validate(ip, nil); err != nil {
		if !valid {
			log.Println("product from cart didn't validate because of", err)
			return
		}

		switch err.(type) {
		case *rule.ErrRequired:
			log.Println("required error:", err)
			break

		case *rule.ErrEnum:
			log.Println("enum error:", err)
			break
		}
	}
}
```

> You can access more [examples](https://github.com/guiferpa/gody/blob/master/example/validate.go)
