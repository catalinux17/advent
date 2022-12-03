package cmd

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use: "day1",
	Run: func(cmd *cobra.Command, args []string) {
		day1()
	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)
}

func day1() {
	file, _ := os.Open("inputs/day1p")

	scanner := bufio.NewScanner(file)

	backpacks := []int{}
	current_backpack := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			backpacks = append(backpacks, current_backpack)
			current_backpack = 0
		}
		n, _ := strconv.Atoi(scanner.Text())
		current_backpack += n
	}

	sort.Ints(backpacks)
	fmt.Println(backpacks)

	fmt.Println(backpacks[len(backpacks)-1] + backpacks[len(backpacks)-2] + backpacks[len(backpacks)-3])
}
