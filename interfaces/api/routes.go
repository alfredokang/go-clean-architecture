package api

import (
	infraController "awesomeProject/interfaces/controllers/auth"
	"awesomeProject/interfaces/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func (s *Service) GetRoutes() {
	s.Engine.Use(middleware.Logger)
	s.Engine.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Permite qualquer origem
		AllowedMethods:   []string{"*"}, // Permite qualquer método
		AllowedHeaders:   []string{"*"}, // Permite qualquer cabeçalho
		AllowCredentials: true,          // Permite credenciais
	}))
	s.Engine.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		_, err := w.Write([]byte("Method is not valid"))
		if err != nil {
			return
		}
	})
	// Authentication Routes
	s.Engine.Get("/login", infraController.Heart)

	// Public Students Routes
	s.Engine.Group(func(r chi.Router) {
		r.Get("/students", s.StudentController.List)
		r.Post("/students", s.StudentController.Create)
		r.Put("/students/{id}", s.StudentController.Update)
		r.Delete("/students/{id}", s.StudentController.Delete)
		r.Get("/students/{id}", s.StudentController.Details)
	})

	// Public Users Routes
	s.Engine.Group(func(r chi.Router) {
		r.Get("/users", s.UserController.List)
		r.Get("/users/{id}", s.UserController.Details)
		r.Post("/users", s.UserController.Create)
		r.Put("/users/{id}", s.UserController.Update)
		r.Delete("/users/{id}", s.UserController.Delete)
	})

	// Private Routes, Require Authentication
	s.Engine.Group(func(r chi.Router) {
		r.Use(middlewares.Authenticator)
		//r.Use(jwtauth.Verifier(s.AuthMiddleware.Token))
		//r.Use(jwtauth.Authenticator)
		r.Post("/manage", s.StudentController.List)
	})
}
