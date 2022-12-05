package cmd

import (
	"bufio"
	"os"

	"github.com/spf13/cobra"
)

var day4Cmd = &cobra.Command{
	Use: "day4",
	Run: func(cmd *cobra.Command, args []string) {
		day4()
	},
}

func init() {
	rootCmd.AddCommand(day4Cmd)
}

func day4() {
	file, _ := os.Open("inputs/day4")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

	}
}
