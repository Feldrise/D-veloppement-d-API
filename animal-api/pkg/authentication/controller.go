package authentication

import (
	"encoding/json"
	"net/http"
	"strconv"

	"feldrise.com/animal-api/database/dbmodel"
	"feldrise.com/animal-api/helper"
	"feldrise.com/animal-api/pkg/errors"
	"feldrise.com/animal-api/pkg/model"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @ID register
// @Tags autentication
// @Param request body RegisterPostPayload true "user's info"
// @Success 201 {string} string "created"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /authentication/register [post]
func (config *Config) Register(w http.ResponseWriter, r *http.Request) {
	data := &model.RegisterPostPayload{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	existingUser, err := config.UserRepository.FindByEmail(*data.Email, false)

	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	if existingUser != nil {
		render.Render(w, r, errors.ErrUnauthorized("email already exists"))
		return
	}

	hashedPassword, err := hashPassword(*data.Password)
	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	user := &dbmodel.User{
		Email:        *data.Email,
		PasswordHash: hashedPassword,
		UserProfile: &dbmodel.UserProfile{
			FirstName: data.FirstName,
			LastName:  data.LastName,
		},
	}

	user, err = config.UserRepository.Create(user)

	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, user.ToModel(true, true))

}

// Login godoc
// @Summary Login a user
// @Description Login a user
// @ID login
// @Tags autentication
// @Param request body LoginPostPayload true "user's info"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /authentication/login [post]
func (config *Config) Login(w http.ResponseWriter, r *http.Request) {
	data := &model.LoginPostPayload{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	user, err := config.UserRepository.FindByEmail(*data.Email, true)

	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	if user == nil {
		render.Render(w, r, errors.ErrUnauthorized("invalid email or password"))
		return
	}

	err = checkPassword(user.PasswordHash, *data.Password)

	if err != nil {
		render.Render(w, r, errors.ErrUnauthorized("invalid email or password"))
		return
	}

	tokenString, err := GenerateToken(config.Constants.JWTSecret, user.ID, user.RoleID)
	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	type UserWithToken struct {
		model.User
		Token string `json:"token"`
	}

	render.JSON(w, r, &UserWithToken{
		User:  *user.ToModel(true, true),
		Token: tokenString,
	})
}

// Update godoc
// @Summary Update the current user
// @Description Update the current user
// @ID update
// @Tags autentication
// @Param request body map[string]interface{} true "user's info"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /authentication/{id} [put]
func (config *Config) Update(w http.ResponseWriter, r *http.Request) {
	loggedUser := ForContext(r.Context())

	if loggedUser == nil {
		render.Render(w, r, errors.ErrUnauthorized("unauthorized"))
		return
	}

	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	id := chi.URLParam(r, "id")

	// Convert id to uint
	idUint, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	if loggedUser.ID != uint(idUint) {
		render.Render(w, r, errors.ErrUnauthorized("unauthorized"))
		return
	}

	helper.ApplyChanges(data, loggedUser)

	user, err := config.UserRepository.Update(loggedUser)

	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	render.JSON(w, r, user.ToModel(true, true))
}

// Me godoc
// @Summary Get the current user
// @Description Get the current user
// @ID me
// @Tags autentication
// @Success 200 {string} string "ok"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /authentication/me [get]
func (config *Config) Me(w http.ResponseWriter, r *http.Request) {
	ctxUser := ForContext(r.Context())

	if ctxUser == nil {
		render.Render(w, r, errors.ErrUnauthorized("invalid user"))
		return
	}

	render.JSON(w, r, ctxUser.ToModel(true, true))
}

// CheckIfEmailExists godoc
// @Summary Check if the email exists
// @Description Check if the email exists
// @ID check-if-email-exists
// @Tags autentication
// @Param email query string true "email"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /authentication/check-email [get]
func (config *Config) CheckIfEmailExists(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	if email == "" {
		render.Render(w, r, errors.ErrInvalidRequest(nil))
		return
	}

	user, err := config.UserRepository.FindByEmail(email, false)

	if err != nil {
		render.Render(w, r, errors.ErrServerError(err))
		return
	}

	if user != nil {
		// Here if it exists we return a 401 to be able to display an error message
		render.Render(w, r, errors.ErrUnauthorized("email already exists"))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, "not exists")
}

// Private

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

func checkPassword(initialPassword string, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(initialPassword), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
