package cmd

import (
	"aoc-2024/day04"

	"github.com/spf13/cobra"
)

var cmd4 = &cobra.Command{
	Use: "4",
	Run: func(cmd *cobra.Command, args []string) {
		err := day04.Solve(inputFileName, debug)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	solveCmd.AddCommand(cmd4)
}
