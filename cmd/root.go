package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ai-interviewer",
		Short: "AI Interviewer CLI tool",
		//Run: func(cmd *cobra.Command, args []string) {
		//	fmt.Println("Вы вызвали команду info123")
		//},
	}
)

func Execute() error {
	return rootCmd.Execute()
}
