package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/guiferpa/gody"
)

type ErrIsAdult struct{}

func (err *ErrIsAdult) Error() string {
	return "The client isn't a adult then it isn't allowed buy"
}

type IsAdultRule struct{}

func (r *IsAdultRule) Name() string {
	return "is_adult"
}

func (r *IsAdultRule) Validate(_, value, _ string) (bool, error) {
	if value == "" {
		return true, &ErrIsAdult{}
	}

	adultAge := 21
	clientAge, err := strconv.Atoi(value)
	if err != nil {
		return false, err
	}

	if clientAge < adultAge {
		return true, &ErrIsAdult{}
	}

	return true, nil
}

type User struct {
	Name string `validate:"min_bound=5"`
	Age  int16  `validate:"min=10 is_adult"`
}

type Product struct {
	Name        string `validate:"not_empty"`
	Description string `validate:"not_empty"`
	Price       int
}

type Cart struct {
	Owner    User      `validate:"required"`
	Products []Product `validate:"required"`
}

func HTTPServerAPI() {
	validator := gody.NewValidator()

	rules := []gody.Rule{
		&IsAdultRule{},
	}
	if err := validator.AddRules(rules); err != nil {
		log.Fatalln(err)
	}

	port := ":8092"
	log.Printf("Example for REST API is running on port %v, now send a POST to /carts and set the claimed Cart struct as request body", port)
	http.ListenAndServe(port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/carts" || r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Path or method is wrong: path: %v, method: %v\n", r.URL.Path, r.Method)
			return
		}

		var body Cart
		err := json.NewDecoder(r.Body).Decode(&body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}

		if validated, err := validator.Validate(body); err != nil {
			if !validated {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Validation for body wasn't processed because of error: %v\n", err)
				return
			}

			w.WriteHeader(http.StatusUnprocessableEntity)

			if _, ok := err.(*ErrIsAdult); ok {
				fmt.Fprintf(w, "The client called %v isn't a adult\n", body.Owner.Name)
				return
			}

			fmt.Fprintln(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "The client %v created your cart!\n", body.Owner.Name)
		return
	}))
}
