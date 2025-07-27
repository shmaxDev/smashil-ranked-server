package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func DecodeAndValidate(obj any, w *http.ResponseWriter, r *http.Request) error {
	validate = validator.New()

	err := json.NewDecoder(r.Body).Decode(obj)

	if err != nil {
		if errors.Is(err, io.EOF) {
			http.Error(*w, "Request body is empty", http.StatusBadRequest)
		}

		http.Error(*w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(obj)

	err = validate.Struct(obj)

	if err != nil {
		fmt.Println(err.Error())

		http.Error(*w, "Bad request: "+err.Error(), http.StatusBadRequest)
	}

	return err

}
