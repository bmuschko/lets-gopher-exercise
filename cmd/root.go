package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "letsgopher",
	Short: "Let's Gopher is a project generator for Golang projects",
	Long: "A flexible and customizable project generator for Golang projects.",
	Run: doGenerateCmd,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
