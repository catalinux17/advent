/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use: "day3",
	Run: func(cmd *cobra.Command, args []string) {
		day3()
	},
}

func init() {
	rootCmd.AddCommand(day3Cmd)

}

func findDuplicateItem(text1, text2 string) rune {
	for _, c1 := range text1 {
		for _, c2 := range text2 {
			if c1 == c2 {
				return c1
			}
		}
	}
	return ' '
}

func itemCost(char int) int {
	if (65 <= char) && (char <= 90) {
		return char - 65 + 27
	}

	if (97 <= char) && (char <= 122) {
		return char - 97 + 1
	}

	return 0
}

func findCommonInGroup(a, b, c string) rune {
	for _, c1 := range a {
		for _, c2 := range b {
			for _, c3 := range c {
				if c1 == c2 && c2 == c3 {
					return c1
				}
			}
		}
	}
	return ' '
}

func day3() {

	file, _ := os.Open("inputs/day3p")
	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		t1 := scanner.Text()

		scanner.Scan()
		t2 := scanner.Text()

		scanner.Scan()
		t3 := scanner.Text()

		dupItem := findCommonInGroup(t1, t2, t3)

		total += itemCost(int(dupItem))

	}

	fmt.Println(total)
}
