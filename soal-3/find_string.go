package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const (
	openBracket  = "("
	closeBracket = ")"
)

var rootCmd = &cobra.Command{
	Use: "soal3",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if len(args) == 1 {
			str := findFirstStringInBracket(args[0])
			fmt.Println(str)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func findFirstStringInBracket(str string) string {
	if len(str) < 0 {
		return str
	}

	if strings.HasPrefix(str, openBracket) {
		str = strings.TrimPrefix(str, openBracket)
	}

	if strings.HasSuffix(str, closeBracket) {
		str = strings.TrimSuffix(str, closeBracket)
	}

	return str
}

func main() {
	Execute()
}
