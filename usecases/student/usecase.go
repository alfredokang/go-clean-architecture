package student

import "awesomeProject/infra/database"

type StudentUseCase struct {
	Database *database.Database
}

func NewStudentUseCase(db *database.Database) *StudentUseCase {
	return &StudentUseCase{
		Database: db,
	}
}
