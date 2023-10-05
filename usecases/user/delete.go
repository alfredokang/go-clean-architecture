package user

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *UserUseCase) Delete(id primitive.ObjectID) error {
	user, err := s.Database.UserRepository.FindById(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("Não foi possível encontrar o estudante!")
		}
		return err
	}

	return s.Database.UserRepository.Delete(user.Id)
}
