package authentication

import (
	"feldrise.com/animal-api/config"
	"github.com/go-chi/chi"
)

func New(configuration *config.Config) *Config {
	return &Config{configuration}
}

func (config *Config) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/register", config.Register)
	router.Post("/login", config.Login)
	router.Put("/{id}", config.Update)
	router.Get("/check-email", config.CheckIfEmailExists)
	router.Get("/me", config.Me)

	return router
}
