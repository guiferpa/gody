# gody

[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![Last commit](https://img.shields.io/github/last-commit/guiferpa/gody)](https://img.shields.io/github/last-commit/guiferpa/gody)
[![GoDoc](https://godoc.org/github.com/guiferpa/gody?status.svg)](https://godoc.org/github.com/guiferpa/gody)
[![Go Report Card](https://goreportcard.com/badge/github.com/guiferpa/gody)](https://goreportcard.com/report/github.com/guiferpa/gody)
![Pipeline workflow](https://github.com/guiferpa/gody/actions/workflows/pipeline.yml/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/guiferpa/gody/badge.svg?branch=main)](https://coveralls.io/github/guiferpa/gody?branch=main)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/guiferpa/gody?color=purple&label=latest)](https://github.com/guiferpa/gody/releases/latest)

### [Go versions supported](https://github.com/guiferpa/gody/blob/main/.github/workflows/pipeline.yml#L12)

### Installation
```bash
go get github.com/guiferpa/gody/v2
```

### Usage

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    gody "github.com/guiferpa/gody/v2"
    "github.com/guiferpa/gody/v2/rule"
) 

type RequestBody struct {
    Name string `json:"name" validate:"not_empty"`
    Age  int    `json:"age" validate:"min=21"`
}

func HTTPHandler(v *gody.Validator) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var body RequestBody
        if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
	    // ...
    	}
	defer r.Body.Close()

	if isValidated, err := v.Validate(body); err != nil {
	    // ...                                                                                
	}
    })
}

func main() {
    validator := gody.NewValidator()

    validator.AddRules(rule.NotEmpty, rule.Min)

    port := ":3000"
    http.ListenAndServe(port, HTTPHandler(validator))
}
```

### Others ways for validation

There are others ways to valid a struct, take a look on functions below:

- **RawValidate** - It's a function that make validate with no rule, it's necessary put the struct for validation, some rule(s) and tag name.

```go
gody.RawValidate(interface{}, string, []gody.Rule) (bool, error)
```

- **Validate** - It's a function that make validate with no rule, it's necessary put the struct for validation and some rule(s).
```go
gody.Validate(interface{}, []gody.Rule) (bool, error)
```

- **RawDefaultValidate** - It's a function that already have [built-in rules](https://github.com/guiferpa/gody/blob/72ce1caecc5fdacf40ee282716ec1b5abe6f7adf/validate.go#L15-L23) configured, it's necessary put the struct for validation, tag name and optional custom rule(s).
```go
gody.RawDefaultValidate(interface{}, string, []gody.Rule) (bool, error)
```

- **DefaultValidate** - It's a function that already have [built-in rules](https://github.com/guiferpa/gody/blob/72ce1caecc5fdacf40ee282716ec1b5abe6f7adf/validate.go#L15-L23) configured, it's necessary put the struct for validation and optional custom rule(s).
```go
gody.DefaultValidate(interface{}, []gody.Rule) (bool, error)
```

### Dynamic Enum Validation (No Duplication)

You can avoid duplicating enum values in struct tags by using dynamic parameters:

```go
const (
    StatusCreated = "__CREATED__"
    StatusPending = "__PENDING__"
    StatusDoing   = "__DOING__"
    StatusDone    = "__DONE__"
)

type Task struct {
    Name   string `json:"name"`
    Status string `json:"status" validate:"enum={status}"`
}

validator := gody.NewValidator()
validator.AddRuleParameters(map[string]string{
    "status": fmt.Sprintf("%s,%s,%s,%s", StatusCreated, StatusPending, StatusDoing, StatusDone),
})
validator.AddRules(rule.Enum)

// Now you can validate Task structs without duplicating enum values in the tag!
```

### Contribution policies

1. At this time the only policy is don't create a Pull Request directly, it's necessary some discussions for some implementation then open an issue before to dicussion anything about the project.
