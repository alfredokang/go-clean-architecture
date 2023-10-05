package user

import "awesomeProject/domain/entities"

func (s *UserUseCase) List() ([]entities.User, error) {
	return s.Database.UserRepository.List()
}
