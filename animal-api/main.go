package main

import (
	"log"
	"net/http"

	_ "feldrise.com/animal-api/docs"

	"feldrise.com/animal-api/config"
	"feldrise.com/animal-api/pkg/authentication"
	"feldrise.com/animal-api/pkg/cat"
	"feldrise.com/animal-api/pkg/visit"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rs/cors"

	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Recoverer,
		authentication.Middelware(configuration),
	)

	// We initialize CORS
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	}).Handler)

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/authentication", authentication.New(configuration).Routes())
		r.Mount("/cat", cat.New(configuration).Routes())
		r.Mount("/{catid}/visits", visit.New(configuration).Routes())
	})

	return router
}

// @title Veterinary API
// @version 1.0
// @description This is the veterinary API
// @contact contact@feldrise.com
// @host
// @BasePath /api/v1
func main() {
	// We initialize the projet
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error: ", err)
	}

	// We initialize the routes
	router := Routes(configuration)

	// Swagger configs
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	// We show all the routes in the logs
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	// We serve the api
	log.Printf("connect to http://localhost:%s/swagger/index.html for documentation", configuration.Constants.Port)
	log.Fatal(http.ListenAndServe(":"+configuration.Constants.Port, router))
}
