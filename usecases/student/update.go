package student

import (
	"awesomeProject/domain/entities"
	"awesomeProject/domain/entities/shared"
	"errors"
	"github.com/google/uuid"
)

func (s *StudentUseCase) Update(id uuid.UUID, name string, age int) (entities.Student, error) {
	student, err := s.Database.StudentRepository.FindById(id)
	if err != nil {
		return student, err
	}

	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Não foi possível encontrar o estudante")
	}

	student.Name = name
	student.Age = age

	err = s.Database.StudentRepository.Update(&student)

	return student, err
}
