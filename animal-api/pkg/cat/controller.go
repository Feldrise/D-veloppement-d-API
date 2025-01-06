package cat

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
// @Summary Get a cat
// @Description Get a cat by its id
// @Tags cats
// @Param id path int true "Cat ID"
// @Success 200 {object} Cat "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /cats/{id} [get]
func (config *Config) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idUint, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	dbCat, err := config.CatsRepository.FindByID(uint(idUint), nil)

	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	if dbCat == nil {
		render.Render(w, r, errors.ErrNotFound())
		return
	}

	render.JSON(w, r, dbCat.ToModel())
}

// GetAll godoc
// @Summary Get all cats
// @Description Get all cats
// @Tags cats
// @Success 200 {array} Cat "ok"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /cats [get]
func (config *Config) GetAll(w http.ResponseWriter, r *http.Request) {
	dbCats, err := config.CatsRepository.FindAll(nil)

	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	cats := make([]model.Cat, 0, len(dbCats))

	for _, dbCat := range dbCats {
		cats = append(cats, *dbCat.ToModel())
	}

	render.JSON(w, r, cats)
}

// Create godoc
// @Summary Create a cat
// @Description Create a cat
// @Tags cats
// @Param cat body CatCreatePayload true "Cat data"
// @Success 200 {object} Cat "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /cats [post]
func (config *Config) Create(w http.ResponseWriter, r *http.Request) {
	loggedUser := authentication.ForContext(r.Context())

	if loggedUser == nil {
		render.Render(w, r, errors.ErrUnauthorized("unauthorized"))
		return
	}

	data := &model.CatCreatePayload{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	dbCat := &dbmodel.Cat{
		Name: data.Name,
	}

	dbCat, err := config.CatsRepository.Create(dbCat)

	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, dbCat.ToModel())
}

// Update godoc
// @Summary Update a cat
// @Description Update a cat
// @Tags cats
// @Param id path int true "Cat ID"
// @Param cat body map[string]interface{} true "Cat data"
// @Success 200 {object} Cat "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /cats/{id} [put]
func (config *Config) Update(w http.ResponseWriter, r *http.Request) {
	loggedUser := authentication.ForContext(r.Context())

	if loggedUser == nil {
		render.Render(w, r, errors.ErrUnauthorized("unauthorized"))
		return
	}

	id := chi.URLParam(r, "id")
	idUint, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	dbCat, err := config.CatsRepository.FindByID(uint(idUint), nil)

	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	if dbCat == nil {
		render.Render(w, r, errors.ErrNotFound())
		return
	}

	var data map[string]interface{}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	helper.ApplyChanges(data, dbCat)

	dbCat, err = config.CatsRepository.Update(dbCat)

	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	render.JSON(w, r, dbCat.ToModel())
}
