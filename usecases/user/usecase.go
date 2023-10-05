package user

import (
	"awesomeProject/infra/database"
)

type UserUseCase struct {
	Database *database.Database
}

func NewUserUseCase(db *database.Database) *UserUseCase {
	return &UserUseCase{
		Database: db,
	}
}
