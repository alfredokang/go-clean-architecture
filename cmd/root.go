package cmd

import (
	"awesomeProject/infra/database"
	"awesomeProject/infra/database/mongo"
	"awesomeProject/infra/database/mongo/repositories"
	"context"
	"github.com/spf13/cobra"
	"log"
)

func Execute() {
	root := &cobra.Command{
		Short:   "Student Mananeger",
		Version: "1.0.0",
	}

	root.AddCommand(
		Api,
	)

	err := root.Execute()
	if err != nil {
		return
	}
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetDatabase(ctx context.Context) *database.Database {
	client, err := mongo.GetConnection(ctx)
	FatalError(err)

	studentRepository := repositories.NewStudentRepository(ctx, client)
	userRepository := repositories.NewUserRepository(ctx, client)

	return database.NewDatabase(client, studentRepository, userRepository)
}
