package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func DecodeAndValidate(obj any, w *http.ResponseWriter, r *http.Request) error {
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		tag := fld.Tag.Get("json")
		if tag == "-" {
			return ""
		}
		name := strings.Split(tag, ",")[0]
		if name == "" {
			return fld.Name
		}
		return name
	})

	err := json.NewDecoder(r.Body).Decode(obj)

	if err != nil {
		if errors.Is(err, io.EOF) {
			http.Error(*w, "Request body is empty", http.StatusBadRequest)
		}

		http.Error(*w, err.Error(), http.StatusBadRequest)
	}

	err = validate.Struct(obj)

	if err != nil {
		http.Error(*w, "Bad request: "+getJsonErrorFromValidator(err), http.StatusBadRequest)
	}

	return err

}

func getJsonErrorFromValidator(err error) string {

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return "Invalid validation error"
	}

	var jsonError string
	for _, fieldError := range err.(validator.ValidationErrors) {
		jsonError += fmt.Sprintf("Error:Field validation for '%s' failed on the '%s' tag\n",
			fieldError.Field(), fieldError.Tag())
		fmt.Println(jsonError)
	}
	return jsonError

}
