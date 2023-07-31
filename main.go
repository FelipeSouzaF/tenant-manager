package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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

var subCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lista todos os tenants",
	Run: func(cmd *cobra.Command, args []string) {
		// This function will be executed when the sub-command is run
		fmt.Printf("Hello, %s!\n", flagName)
	},
}

func init() {
	rootCmd.AddCommand(subCmd)

	subCmd.Flags().StringVarP(&flagName, "name", "n", "nome_do_tenant", "Nome do tenant")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
