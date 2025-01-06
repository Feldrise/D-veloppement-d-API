package visit

import (
	"encoding/json"
	"net/http"
	"strconv"

	"feldrise.com/animal-api/database/dbmodel"
	"feldrise.com/animal-api/helper"
	"feldrise.com/animal-api/pkg/authentication"
	"feldrise.com/animal-api/pkg/errors"
	"feldrise.com/animal-api/pkg/model"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Get godoc
// @Summary Get a visit
// @Description Get a visit by ID and for a specific cat
// @Tags visits
// @Param catid path int true "Cat ID"
// @Param id path int true "Visit ID"
// @Success 200 {object} Visit "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /cats/{catid}/visits/{id} [get]
func (config *Config) Get(w http.ResponseWriter, r *http.Request) {
	catID := chi.URLParam(r, "catid")
	visitID := chi.URLParam(r, "id")
	catIDUint, err := strconv.ParseUint(catID, 10, 64)
	visitIDUint, err := strconv.ParseUint(visitID, 10, 64)

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	dbVisit, err := config.VisitsRepository.FindByID(uint(visitIDUint), &dbmodel.VisitsFieldsToInclude{
		Cat: true,
	})

	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	if dbVisit == nil {
		render.Render(w, r, errors.ErrNotFound())
		return
	}

	if dbVisit.CatID != uint(catIDUint) {
		render.Render(w, r, errors.ErrUnauthorized("unauthorized"))
		return
	}

	render.JSON(w, r, dbVisit.ToModel())
}

// GetAll godoc
// @Summary Get all visits
// @Description Get all visits for a specific cat
// @Tags visits
// @Param catid path int true "Cat ID"
// @Success 200 {array} Visit "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /cats/{catid}/visits [get]
func (config *Config) GetAll(w http.ResponseWriter, r *http.Request) {
	catID := chi.URLParam(r, "catid")
	catIDUint, err := strconv.ParseUint(catID, 10, 64)

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	dbVisits, err := config.VisitsRepository.FindAll(&dbmodel.VisitsFilter{
		CatID: uint(catIDUint),
	}, &dbmodel.VisitsFieldsToInclude{
		Cat: true,
	})

	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	visits := make([]model.Visit, 0, len(dbVisits))

	for _, dbVisit := range dbVisits {
		visits = append(visits, *dbVisit.ToModel())
	}

	render.JSON(w, r, visits)
}

// Create godoc
// @Summary Create a visit
// @Description Create a visit for a specific cat
// @Tags visits
// @Param catid path int true "Cat ID"
// @Accept json
// @Produce json
// @Param request body Visit true "Visit info"
// @Success 200 {object} Visit "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /cats/{catid}/visits [post]
func (config *Config) Create(w http.ResponseWriter, r *http.Request) {
	loggedUser := authentication.ForContext(r.Context())

	if loggedUser == nil {
		render.Render(w, r, errors.ErrUnauthorized("unauthorized"))
		return
	}

	catID := chi.URLParam(r, "catid")
	catIDUint, err := strconv.ParseUint(catID, 10, 64)

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	var data model.VisitCreatePayload

	err = render.Bind(r, &data)
	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	dbVisit, err := config.VisitsRepository.Create(&dbmodel.Visit{
		Date:  *data.Date,
		CatID: uint(catIDUint),
	})

	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	render.JSON(w, r, dbVisit.ToModel())
}

// Update godoc
// @Summary Update a visit
// @Description Update a visit for a specific cat
// @Tags visits
// @Param catid path int true "Cat ID"
// @Param id path int true "Visit ID"
// @Accept json
// @Produce json
// @Param request body map[string]interface{} true "Visit info"
// @Success 200 {object} Visit "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /cats/{catid}/visits/{id} [put]
func (config *Config) Update(w http.ResponseWriter, r *http.Request) {
	loggedUser := authentication.ForContext(r.Context())

	if loggedUser == nil {
		render.Render(w, r, errors.ErrUnauthorized("unauthorized"))
		return
	}

	catID := chi.URLParam(r, "catid")
	visitID := chi.URLParam(r, "id")
	catIDUint, err := strconv.ParseUint(catID, 10, 64)
	visitIDUint, err := strconv.ParseUint(visitID, 10, 64)

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	dbVisit, err := config.VisitsRepository.FindByID(uint(visitIDUint), nil)

	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	if dbVisit == nil {
		render.Render(w, r, errors.ErrNotFound())
		return
	}

	if dbVisit.CatID != uint(catIDUint) {
		render.Render(w, r, errors.ErrUnauthorized("unauthorized"))
		return
	}

	var data map[string]interface{}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	helper.ApplyChanges(data, dbVisit)

	dbVisit, err = config.VisitsRepository.Update(dbVisit)

	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	render.JSON(w, r, dbVisit.ToModel())
}
