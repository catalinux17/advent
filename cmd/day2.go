package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// day2Cmd represents the day2 command
var day2Cmd = &cobra.Command{
	Use: "day2",
	Run: func(cmd *cobra.Command, args []string) {
		day2()
	},
}

func init() {
	rootCmd.AddCommand(day2Cmd)
}

// declareWinner returns
// 0, if it's a draw
// 1, if player 1 wins
// 2, if player 2 wins
func declareWinner(p1, p2 int) int {
	switch p1 {
	case 0:
		switch p2 {
		case 0:
			return 0
		case 1:
			return 2
		case 2:
			return 1
		}
	case 1:
		switch p2 {
		case 0:
			return 1
		case 1:
			return 0
		case 2:
			return 2
		}
	case 2:
		switch p2 {
		case 0:
			return 2
		case 1:
			return 1
		case 2:
			return 0
		}
	}

	return -1
}

// 0 = ROCK
// 1 = PAPER
// 2 = SCISSORS
func translate(shape string) int {
	switch shape {
	case "A":
		fallthrough
	case "X":
		return 0
	case "B":
		fallthrough
	case "Y":
		return 1
	case "C":
		fallthrough
	case "Z":
		return 2
	default:
		return -1
	}
}

func day2() {

	file, _ := os.Open("inputs/day2p")

	scanner := bufio.NewScanner(file)

	totalScore := 0

	for scanner.Scan() {
		text := scanner.Text()

		p1 := translate(text[0:1])
		p2 := translate(text[2:3])

		score := 0

		switch p2 {
		case 1:
			score += 3
		case 2:
			score += 6
		}

		aux := 0
		switch p2 {
		case 0:
			aux = 1
		case 1:
			aux = 0
		}

		switch {
		case declareWinner(p1, 0) == aux:
			score += 1
		case declareWinner(p1, 1) == aux:
			score += 2
		case declareWinner(p1, 2) == aux:
			score += 3
		}

		totalScore += score

	}

	fmt.Println(totalScore)

}
