package student

import (
	"awesomeProject/infra/database"
	mock "awesomeProject/infra/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

type StudentUsecaseSuite struct {
	suite.Suite
	ctrl              *gomock.Controller
	StudentRepository *mock.MockStudentRepository
	UserRepository    *mock.MockUserRepository
	StudentUseCase    *StudentUseCase
}

func TestStudentUseCaseStart(t *testing.T) {
	suite.Run(t, new(StudentUsecaseSuite))
}

func (s *StudentUsecaseSuite) StudentUsecaseSuiteDown() {
	defer s.ctrl.Finish()
}

func (s *StudentUsecaseSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	conn, _ := mongo.NewClient()
	s.StudentRepository = mock.NewMockStudentRepository(s.ctrl)
	db := database.NewDatabase(conn, s.StudentRepository, s.UserRepository)

	s.StudentUseCase = NewStudentUseCase(db)
}
