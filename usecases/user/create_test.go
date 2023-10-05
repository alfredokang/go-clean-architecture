package user

import (
	"awesomeProject/domain/entities"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func (s *UserUsecaseSuite) TestCreateShouldReturnSuccess() {
	name := "Marcelo"
	mobile := "(11) 96151-1000"
	email := "alfredok@icloud.com"
	role := "admin"
	password := "123456"

	userDto := entities.UserDto{Name: name, Mobile: mobile, Email: email, Role: role, Password: password}

	s.UserRepository.EXPECT().Create(gomock.Any()).Return(nil).Times(1)

	user, err := s.UserUseCase.Create(userDto)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), name, user.Name)
	assert.Equal(s.T(), mobile, user.Mobile)
	assert.Equal(s.T(), email, user.Email)
	assert.Equal(s.T(), role, user.Role)
	assert.Equal(s.T(), password, user.Password)
}

func (s *UserUsecaseSuite) TestCreateShouldReturnErr() {
	name := "Marcelo"
	mobile := "(11) 96151-1000"
	email := "alfredok@icloud.com"
	role := "admin"
	password := "123456"
	errExpect := fmt.Errorf("Error Expect")

	userDto := entities.UserDto{Name: name, Mobile: mobile, Email: email, Role: role, Password: password}

	s.UserRepository.EXPECT().Create(gomock.Any()).Return(errExpect).Times(1)

	user, err := s.UserUseCase.Create(userDto)

	assert.Error(s.T(), err, errExpect.Error())
	assert.Equal(s.T(), name, user.Name)
	assert.Equal(s.T(), mobile, user.Mobile)
	assert.Equal(s.T(), email, user.Email)
	assert.Equal(s.T(), role, user.Role)
	assert.Equal(s.T(), password, user.Password)
}
