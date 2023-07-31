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
	Use:   "your_app_name",
	Short: "A brief description of your CLI application",
	Long:  `A longer description of your CLI application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// This function will be executed when your command is run
		fmt.Println("Hello, world!")
	},
}

var subCmd = &cobra.Command{
	Use:   "subcommand",
	Short: "A brief description of your sub-command",
	Long:  `A longer description of your sub-command.`,
	Run: func(cmd *cobra.Command, args []string) {
		// This function will be executed when the sub-command is run
		fmt.Printf("Hello, %s!\n", flagName)
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
	subCmd.Flags().StringVarP(&flagName, "name", "n", "John", "Name to greet")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
