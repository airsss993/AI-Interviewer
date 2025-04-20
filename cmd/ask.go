package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask the AI a question",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("AI: Привет! Я твой собеседующий бот. Что хочешь узнать?")
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}
