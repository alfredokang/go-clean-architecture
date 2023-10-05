package user

import (
	"awesomeProject/domain/entities"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *UserUsecaseSuite) TestUpdateShouldReturnSuccess() {
	id := primitive.NewObjectID()
	name := "Marcelo"
	mobile := "(11) 96151-1000"
	email := "alfredok@icloud.com"
	role := "admin"

	s.UserRepository.EXPECT().FindById(id).Return(entities.User{Id: id}, nil).Times(1)
	s.UserRepository.EXPECT().Update(gomock.Any()).Return(nil).Times(1)

	userDto := entities.UserDto{
		Name:   name,
		Mobile: mobile,
		Email:  email,
		Role:   role,
	}

	newUser := *entities.NewUser(userDto)

	user, err := s.UserUseCase.Update(id, newUser)

	assert.NotNil(s.T(), user)
	assert.EqualValues(s.T(), id, user.Id)
	assert.EqualValues(s.T(), name, user.Name)
	assert.EqualValues(s.T(), mobile, user.Mobile)
	assert.EqualValues(s.T(), email, user.Email)
	assert.EqualValues(s.T(), role, user.Role)
	assert.Nil(s.T(), err)
}

func (s *UserUsecaseSuite) TestUpdateShouldReturnErrFindById() {
	id := primitive.NewObjectID()
	name := "Marcelo"
	mobile := "(11) 96151-1000"
	email := "alfredok@icloud.com"
	role := "admin"
	errExpect := fmt.Errorf("Error expected")

	s.UserRepository.EXPECT().FindById(id).Return(entities.User{Id: id}, errExpect).Times(1)

	userDto := entities.UserDto{
		Name:   name,
		Mobile: mobile,
		Email:  email,
		Role:   role,
	}

	newUser := *entities.NewUser(userDto)

	user, err := s.UserUseCase.Update(id, newUser)

	assert.NotNil(s.T(), user)
	assert.Error(s.T(), err, errExpect.Error())
}

//func (s *UserUsecaseSuite) TestUpdateShouldReturnErrIdEmpty() {
//	id := primitive.NewObjectID()
//	name := "Marcelo"
//	mobile := "(11) 96151-1000"
//	email := "alfredok@icloud.com"
//	role := "admin"
//
//	s.UserRepository.EXPECT().FindById(id).Return(entities.User{}, nil).Times(1)
//
//	userDto := entities.UserDto{
//		Name:   name,
//		Mobile: mobile,
//		Email:  email,
//		Role:   role,
//	}
//
//	newUser := *entities.NewUser(userDto)
//
//	user, err := s.UserUseCase.Update(id, newUser)
//
//	assert.NotNil(s.T(), user)
//	assert.Error(s.T(), err, "Não foi possível encontrar o usuário!")
//}

func (s *UserUsecaseSuite) TestUpdateShouldReturnErrUpdate() {
	id := primitive.NewObjectID()
	name := "Marcelo"
	mobile := "(11) 96151-1000"
	email := "alfredok@icloud.com"
	role := "admin"
	errExpect := fmt.Errorf("Error expected")

	s.UserRepository.EXPECT().FindById(id).Return(entities.User{Id: id}, nil).Times(1)
	s.UserRepository.EXPECT().Update(gomock.Any()).Return(errExpect).Times(1)

	userDto := entities.UserDto{
		Name:   name,
		Mobile: mobile,
		Email:  email,
		Role:   role,
	}

	newUser := *entities.NewUser(userDto)

	user, err := s.UserUseCase.Update(id, newUser)

	assert.NotNil(s.T(), user)
	assert.EqualValues(s.T(), id, user.Id)
	assert.EqualValues(s.T(), name, user.Name)
	assert.EqualValues(s.T(), mobile, user.Mobile)
	assert.EqualValues(s.T(), email, user.Email)
	assert.EqualValues(s.T(), role, user.Role)
	assert.Error(s.T(), err, errExpect.Error())
}
