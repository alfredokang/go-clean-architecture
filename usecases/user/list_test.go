package user

import (
	"awesomeProject/domain/entities"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func (s *UserUsecaseSuite) TestListShouldReturnSuccess() {
	s.UserRepository.EXPECT().List().Return([]entities.User{{}}, nil)

	list, err := s.UserUseCase.List()

	assert.NotNil(s.T(), list)
	assert.EqualValues(s.T(), 1, len(list))
	assert.Nil(s.T(), err)
}

func (s *UserUsecaseSuite) TestListShouldReturnErr() {
	errExpect := fmt.Errorf("Error expected")

	s.UserRepository.EXPECT().List().Return([]entities.User{{}}, errExpect)

	list, err := s.UserUseCase.List()

	assert.NotNil(s.T(), list)
	assert.EqualValues(s.T(), 1, len(list))
	assert.Error(s.T(), err, errExpect.Error())
}
