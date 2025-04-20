package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ai-interviewer",
		Short: "AI Interviewer CLI tool",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
