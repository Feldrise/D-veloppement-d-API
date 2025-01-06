package model

import (
	"errors"
	"net/http"
)

type Cat struct {
	Name string `json:"name"`
} // @name Cat

type CatCreatePayload struct {
	Name string `json:"name"`
} // @name CatCreatePayload

func (c CatCreatePayload) Bind(r *http.Request) error {
	if c.Name == "" {
		return errors.New("missing name property")
	}

	return nil
}
