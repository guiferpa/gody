# gody

:balloon: **A lightweight struct validator for Go**

[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![GoDoc](https://godoc.org/github.com/guiferpa/gody?status.svg)](https://godoc.org/github.com/guiferpa/gody)
[![Go Report Card](https://goreportcard.com/badge/github.com/guiferpa/gody)](https://goreportcard.com/report/github.com/guiferpa/gody)
[![Build Status](https://cloud.drone.io/api/badges/guiferpa/gody/status.svg)](https://cloud.drone.io/guiferpa/gody)
[![Coverage Status](https://coveralls.io/repos/github/guiferpa/gody/badge.svg?branch=master)](https://coveralls.io/github/guiferpa/gody?branch=master)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/guiferpa/gody?color=purple&label=latest)](https://github.com/guiferpa/gody/releases/latest)

## Quick start

```go
package main

import (
	"log"

	"github.com/guiferpa/gody"
	"github.com/guiferpa/gody/rule"
)

type Body struct {
	Text string `json:"text" validate:"not_empty"`
}

func main() {
	b := Body{}

	if valid, err := gody.Validate(b, nil); err != nil {
		if !valid {
			log.Println("body didn't validate:", err)
		}

		switch err.(type) {
		case *rule.ErrNotEmpty:
			log.Println(err)
		}
	}
}
```

### Kinds of validation

- [Simple](https://github.com/guiferpa/gody/blob/master/example/validate.go#L11-L29)
- [Deep](https://github.com/guiferpa/gody/blob/master/example/validate.go#L84-L115)
- [Custom](https://github.com/guiferpa/gody/blob/master/example/validate.go#L31-L82)
