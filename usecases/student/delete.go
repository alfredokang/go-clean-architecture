package student

import (
	"awesomeProject/domain/entities/shared"
	"errors"
	"github.com/google/uuid"
)

func (s *StudentUseCase) Delete(id uuid.UUID) error {
	student, err := s.Database.StudentRepository.FindById(id)
	if err != nil {
		return err
	}

	if student.ID == shared.GetUuidEmpty() {
		return errors.New("Não foi possível encontrar o estudante")
	}

	return s.Database.StudentRepository.Delete(id)
}
