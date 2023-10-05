package user

import (
	"awesomeProject/infra/database"
	mock "awesomeProject/infra/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

type UserUsecaseSuite struct {
	suite.Suite
	ctrl              *gomock.Controller
	StudentRepository *mock.MockStudentRepository
	UserRepository    *mock.MockUserRepository
	UserUseCase       *UserUseCase
}

func TestUserUsecaseStart(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}

func (s *UserUsecaseSuite) UserUsecaseSuiteDown() {
	defer s.ctrl.Finish()
}

func (s *UserUsecaseSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	conn, _ := mongo.NewClient()
	s.UserRepository = mock.NewMockUserRepository(s.ctrl)
	db := database.NewDatabase(conn, s.StudentRepository, s.UserRepository)

	s.UserUseCase = NewUserUseCase(db)
}
