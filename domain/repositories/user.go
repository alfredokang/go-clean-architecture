package repositories

import (
	"awesomeProject/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecaseContract interface {
	Create(user entities.UserDto) (entities.User, error)
	List() ([]entities.User, error)
	FindById(id primitive.ObjectID) (entities.User, error)
	Update(id primitive.ObjectID, user entities.User) (entities.User, error)
	Delete(id primitive.ObjectID) error
}
