package student

import (
	"awesomeProject/domain/entities"
)

func (s *StudentUseCase) Create(name string, age int) (entities.Student, error) {
	student := entities.NewStudent(name, age)

	err := s.Database.StudentRepository.Create(student)

	return *student, err
}
