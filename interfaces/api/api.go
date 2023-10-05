package api

import (
	"awesomeProject/infra/config"
	"awesomeProject/infra/database"
	"awesomeProject/interfaces/controllers/students"
	"awesomeProject/interfaces/controllers/users"
	"awesomeProject/interfaces/middlewares"
	studentUsecases "awesomeProject/usecases/student"
	userUsecases "awesomeProject/usecases/user"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Service struct {
	Engine            *chi.Mux
	Database          *database.Database
	StudentController *students.StudentController
	UserController    *users.UserController
	AuthMiddleware    middlewares.AuthMiddleware
}

func NewService(db *database.Database) *Service {
	return &Service{
		Engine:   chi.NewRouter(),
		Database: db,
	}
}

func (s *Service) GetControllers() {
	studentUseCase := studentUsecases.NewStudentUseCase(s.Database)
	userUseCase := userUsecases.NewUserUseCase(s.Database)
	s.StudentController = students.NewStudentController(studentUseCase)
	s.UserController = users.NewUserController(userUseCase)
}

func (s *Service) Start() error {
	s.GetControllers()
	s.GetRoutes()
	middlewares.NewAuthMiddleware()
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Env.ApiPort), s.Engine)
}
