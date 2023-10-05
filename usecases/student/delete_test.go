package student

import (
	"awesomeProject/domain/entities"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func (s *StudentUsecaseSuite) TestDeleteShoudReturnSuccess() {
	id := uuid.New()

	s.StudentRepository.EXPECT().FindById(id).Return(entities.Student{ID: id}, nil).Times(1)
	s.StudentRepository.EXPECT().Delete(id).Return(nil).Times(1)

	err := s.StudentUseCase.Delete(id)

	assert.Nil(s.T(), err)
}

func (s *StudentUsecaseSuite) TestDeleteShoudReturnErrFindByID() {
	id := uuid.New()
	errExpect := fmt.Errorf("error expected")

	s.StudentRepository.EXPECT().FindById(id).Return(entities.Student{ID: id}, errExpect).Times(1)

	err := s.StudentUseCase.Delete(id)

	assert.Error(s.T(), err, errExpect.Error())
}

func (s *StudentUsecaseSuite) TestDeleteShoudReturnErrDelete() {
	id := uuid.New()
	errExpect := fmt.Errorf("error expected")

	s.StudentRepository.EXPECT().FindById(id).Return(entities.Student{ID: id}, nil).Times(1)
	s.StudentRepository.EXPECT().Delete(id).Return(errExpect).Times(1)

	err := s.StudentUseCase.Delete(id)

	assert.Error(s.T(), err, errExpect.Error())
}

func (s *StudentUsecaseSuite) TestDeleteShoudReturnErrIDEmpty() {
	id := uuid.New()

	s.StudentRepository.EXPECT().FindById(id).Return(entities.Student{}, nil).Times(1)

	err := s.StudentUseCase.Delete(id)

	assert.Error(s.T(), err, "Não foi possivel encontrar o estudante")
}
