package model

import (
	"errors"
	"net/http"
	"time"
)

type Visit struct {
	Date time.Time `json:"date"`
	Cat  *Cat      `json:"cat"`
} // @name Visit

type VisitCreatePayload struct {
	Date *time.Time `json:"date" validate:"required" example:"2021-01-01T00:00:00Z"`
} // @name VisitCreatePayload

func (v VisitCreatePayload) Bind(r *http.Request) error {
	if v.Date == nil {
		return errors.New("missing date property")
	}

	return nil
}
