package main

import "AI_Interviewer/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		return
	}
}
