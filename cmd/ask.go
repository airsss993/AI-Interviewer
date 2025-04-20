package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask the AI a question",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("AI: Привет! Я твой собеседующий бот. Что хочешь узнать?")
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("Ты: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if strings.ToLower(input) == "выход" {
				break
			}
			//fmt.Printf("%s\n", input)
			fmt.Print("AI: Nice!\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}
