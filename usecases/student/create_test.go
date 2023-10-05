package student

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func (s *StudentUsecaseSuite) TestCreateShouldReturnSuccess() {
	name := "Test"
	age := 10

	s.StudentRepository.EXPECT().Create(gomock.Any()).Return(nil).Times(1)

	student, err := s.StudentUseCase.Create(name, age)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), name, student.Name)
	assert.Equal(s.T(), age, student.Age)
}

func (s *StudentUsecaseSuite) TestCreateShouldReturnError() {
	name := "Test"
	age := 10
	errExpect := fmt.Errorf("Error expect")

	s.StudentRepository.EXPECT().Create(gomock.Any()).Return(errExpect).Times(1)

	student, err := s.StudentUseCase.Create(name, age)

	assert.Error(s.T(), err, errExpect.Error())
	assert.Equal(s.T(), name, student.Name)
	assert.Equal(s.T(), age, student.Age)
}
