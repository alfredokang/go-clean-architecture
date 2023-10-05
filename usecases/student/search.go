package student

import (
	"awesomeProject/domain/entities"
	"github.com/google/uuid"
)

func (s *StudentUseCase) SearchByID(id uuid.UUID) (entities.Student, error) {
	return s.Database.StudentRepository.FindById(id)
}
