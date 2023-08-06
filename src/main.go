package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/FelipeSouzaF/tenant-manager/db"
	"github.com/FelipeSouzaF/tenant-manager/env"
)

var (
	flagName string
)

var rootCmd = &cobra.Command{
	Use:   "tenant",
	Short: "Executa operações nos tenants",
	Run: func(cmd *cobra.Command, args []string) {
		// This function will be executed when your command is run
		fmt.Println("CLI tenant")
	},
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lista todos os tenants",
	Run: func(cmd *cobra.Command, args []string) {
		// This function will be executed when the sub-command is run
		db.GetConnection()

		fmt.Printf("Lista, %s!\n", flagName)
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Cria um tenant",
	Run: func(cmd *cobra.Command, args []string) {
		// This function will be executed when the sub-command is run
		fmt.Printf("Cria, %s!\n", flagName)
	},
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove um tenant",
	Run: func(cmd *cobra.Command, args []string) {
		// This function will be executed when the sub-command is run
		fmt.Printf("Remove, %s!\n", flagName)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(rmCmd)

	lsCmd.Flags().StringVarP(&flagName, "name", "n", "nome_do_tenant", "Nome do tenant")
	addCmd.Flags().StringVarP(&flagName, "name", "n", "nome_do_tenant", "Nome do tenant")
	rmCmd.Flags().StringVarP(&flagName, "name", "n", "nome_do_tenant", "Nome do tenant")
}

func main() {
	environment := "dev"

	err := env.InitConfig(environment)
	if err != nil {
		panic(err)
	}

	appConfig := env.GetConfig()

	fmt.Printf("Running %s on port %d\n", appConfig.AppName, appConfig.Port)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
