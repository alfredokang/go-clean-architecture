package student

import (
	"awesomeProject/domain/entities"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func (s *StudentUsecaseSuite) TestListShoudReturnSuccess() {
	s.StudentRepository.EXPECT().List().Return([]entities.Student{entities.Student{}}, nil)

	list, err := s.StudentUseCase.List()

	assert.NotNil(s.T(), list)
	assert.EqualValues(s.T(), 1, len(list))
	assert.Nil(s.T(), err)
}

func (s *StudentUsecaseSuite) TestListShoudReturnErr() {
	errExpect := fmt.Errorf("error expected")

	s.StudentRepository.EXPECT().List().Return([]entities.Student{entities.Student{}}, errExpect)

	list, err := s.StudentUseCase.List()

	assert.NotNil(s.T(), list)
	assert.EqualValues(s.T(), 1, len(list))
	assert.Error(s.T(), err, errExpect.Error())
}
