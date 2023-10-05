package entities

import (
	"awesomeProject/domain/entities/shared"
	"github.com/google/uuid"
)

type Student struct {
	ID   uuid.UUID `json:"id" bson:"_id"`
	Name string    `json:"name" bson:"name"`
	Age  int       `json:"age" bson:"age"`
}

func NewStudent(name string, age int) *Student {
	return &Student{
		ID:   shared.GetUuid(),
		Name: name,
		Age:  age,
	}
}

type StudentRepository interface {
	Create(student *Student) error
	List() ([]Student, error)
	FindById(id uuid.UUID) (Student, error)
	Update(student *Student) error
	Delete(id uuid.UUID) error
}
