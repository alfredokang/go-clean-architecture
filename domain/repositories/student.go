package repositories

import (
	"awesomeProject/domain/entities"
	"github.com/google/uuid"
)

type StudentUsecaseContract interface {
	Create(name string, age int) (entities.Student, error)
	Delete(id uuid.UUID) error
	List() ([]entities.Student, error)
	SearchByID(id uuid.UUID) (entities.Student, error)
	Update(id uuid.UUID, name string, age int) (entities.Student, error)
}
