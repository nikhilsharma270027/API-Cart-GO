package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {
	// get jSON payload
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload) // wo body se aaya usko decode payload ke type se karna hai
	// check if the user exist
	// if it doesnt we create the new user
}

func WriteJSon(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSon(w, status, map[string]string{"error": err.Error()})
}
