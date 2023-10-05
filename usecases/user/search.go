package user

import (
	"awesomeProject/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *UserUseCase) FindById(id primitive.ObjectID) (entities.User, error) {
	return s.Database.UserRepository.FindById(id)
}
