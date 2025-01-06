package visit

import (
	"feldrise.com/animal-api/config"
	"github.com/go-chi/chi/v5"
)

func New(configuration *config.Config) *Config {
	return &Config{configuration}
}

func (config *Config) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", config.GetAll)
	router.Post("/", config.Create)
	router.Get("/{id}", config.Get)
	router.Put("/{id}", config.Update)

	return router
}
