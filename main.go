package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// 1. Define the Root Command
	var rootCmd = &cobra.Command{
		Use:   "gopherctl",
		Short: "GopherCTL is a simple demo CLI",
		Long:  `A comprehensive example of building a CLI in Go using the Cobra library.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to GopherCTL! Use --help to see available commands.")
		},
	}

	// 2. Define a Sub-command: 'greet'
	var name string
	var greetCmd = &cobra.Command{
		Use:   "greet",
		Short: "Greets the user",
		Run: func(cmd *cobra.Command, args []string) {
			if name != "" {
				fmt.Printf("Hello, %s! Welcome to the Go open-source world.\n", name)
			} else {
				fmt.Println("Hello there, Gopher! Try using the --name flag.")
			}
		},
	}

	// 3. Add Flags to the sub-command
	// This creates a flag: --name (or -n)
	greetCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the person to greet")

	// 4. Add the sub-command to the root
	rootCmd.AddCommand(greetCmd)

	// 5. Execute the CLI
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
