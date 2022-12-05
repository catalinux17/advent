package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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

func isIncluded(p11, p12, p21, p22 int) bool {
	return (p11 <= p21 && p22 <= p12)
}

func divideParts(text string) (h1, h2 int) {
	h1, _ = strconv.Atoi(strings.Split(text, "-")[0])
	h2, _ = strconv.Atoi(strings.Split(text, "-")[1])
	return
}

func day4() {
	file, _ := os.Open("inputs/day4p")
	scanner := bufio.NewScanner(file)

	includedPairs := 0

	for scanner.Scan() {
		text := scanner.Text()

		p1 := strings.Split(text, ",")[0]
		p2 := strings.Split(text, ",")[1]

		p11, p12 := divideParts(p1)
		p21, p22 := divideParts(p2)

		if isIncluded(p11, p12, p21, p22) {
			includedPairs++
			continue
		}

		if isIncluded(p21, p22, p11, p12) {
			includedPairs++
		}

	}

	fmt.Println(includedPairs)
}
