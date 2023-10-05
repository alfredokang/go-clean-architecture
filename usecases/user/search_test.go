package user

import (
	"awesomeProject/domain/entities"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *UserUsecaseSuite) TestSearchShouldReturnSuccess() {
	id := primitive.NewObjectID()

	s.UserRepository.EXPECT().FindById(id).Return(entities.User{Id: id}, nil)

	user, err := s.UserUseCase.FindById(id)

	assert.NotEmpty(s.T(), user)
	assert.EqualValues(s.T(), user.Id, id)
	assert.Nil(s.T(), err)
}

func (s *UserUsecaseSuite) TestSearchShouldReturnErr() {
	id := primitive.NewObjectID()
	errExpect := fmt.Errorf("Error expected")

	s.UserRepository.EXPECT().FindById(id).Return(entities.User{Id: id}, errExpect)

	user, err := s.UserUseCase.FindById(id)

	assert.NotEmpty(s.T(), user)
	assert.EqualValues(s.T(), user.Id, id)
	assert.Error(s.T(), err, errExpect.Error())
}
