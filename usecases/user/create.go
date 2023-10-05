package user

import "awesomeProject/domain/entities"

func (s *UserUseCase) Create(u entities.UserDto) (entities.User, error) {
	user := entities.NewUser(u)

	err := s.Database.UserRepository.Create(user)

	return *user, err
}
