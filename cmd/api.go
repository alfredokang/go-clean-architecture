package cmd

import (
	"awesomeProject/infra/config"
	"awesomeProject/interfaces/api"
	"context"
	"github.com/spf13/cobra"
)

var Api = &cobra.Command{
	Use:   "api",
	Short: "API manager student register",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ctx := context.TODO()

		err = config.StartConfig()
		FatalError(err)

		db := GetDatabase(ctx)

		err = api.NewService(db).Start()
		FatalError(err)
	},
}
