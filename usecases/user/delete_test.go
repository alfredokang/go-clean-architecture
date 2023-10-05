package user

import (
	"awesomeProject/domain/entities"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *UserUsecaseSuite) TestDeleteShouldReturnSuccess() {
	id := primitive.NewObjectID()

	s.UserRepository.EXPECT().FindById(id).Return(entities.User{Id: id}, nil).Times(1)
	s.UserRepository.EXPECT().Delete(id).Return(nil).Times(1)

	err := s.UserUseCase.Delete(id)

	assert.Nil(s.T(), err)
}

func (s *UserUsecaseSuite) TestDeleteShouldReturnErrFindById() {
	id := primitive.NewObjectID()
	errExpect := fmt.Errorf("Error expected")

	s.UserRepository.EXPECT().FindById(id).Return(entities.User{Id: id}, errExpect).Times(1)

	err := s.UserUseCase.Delete(id)

	assert.Error(s.T(), err, errExpect.Error())
}

func (s *UserUsecaseSuite) TestDeleteShouldReturnErrDelete() {
	id := primitive.NewObjectID()
	errExpect := fmt.Errorf("Error expected")

	s.UserRepository.EXPECT().FindById(id).Return(entities.User{Id: id}, nil).Times(1)
	s.UserRepository.EXPECT().Delete(id).Return(errExpect).Times(1)

	err := s.UserUseCase.Delete(id)

	assert.Error(s.T(), err, errExpect.Error())
}

func (s *UserUsecaseSuite) TestDeleteShouldReturnErrIdEmpty() {
	id := primitive.NewObjectID()
	errExpect := fmt.Errorf("Error expected")

	s.UserRepository.EXPECT().FindById(id).Return(entities.User{}, errExpect).Times(1)

	err := s.UserUseCase.Delete(id)

	assert.Error(s.T(), err, "Não foi possível encontrar o estudante!")
}
