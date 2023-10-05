package student

import (
	"awesomeProject/domain/entities"
)

func (s *StudentUseCase) List() ([]entities.Student, error) {
	return s.Database.StudentRepository.List()
}
