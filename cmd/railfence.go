package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	railFenceRails  int
	railFenceDecode bool
)

var railFenceCmd = &cobra.Command{
	Use:   "railfence",
	Short: "Encode and decode using the Rail Fence cipher",
	Long:  `Encode and decode using the Rail Fence cipher with a mnemonic command syntax`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := readInput(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := railFenceCipher(input, railFenceRails, railFenceDecode)
		fmt.Println(result)
	},
	TraverseChildren: true,
	Args:             cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(railFenceCmd)
	railFenceCmd.Flags().BoolVarP(&railFenceDecode, "decode", "d", false, "Decode the input")
	railFenceCmd.Flags().IntVarP(&railFenceRails, "rails", "r", 0, "Number of rails for the Rail Fence cipher(required)")
	railFenceCmd.MarkFlagRequired("rails")
}

// The Rail Fence cipher
func railFenceCipher(input string, rails int, decode bool) string {
	if rails < 2 {
		return "error number of rails must be at least 2."
	}
	if decode {
		return railFenceDecodeFunc(input, rails)
	}
	return railFenceEncryptFunc(input, rails)
}

func railFenceEncryptFunc(input string, rails int) string {
	fence := make([][]rune, rails)
	for i := range fence {
		fence[i] = make([]rune, len(input))
	}
	direction := 1
	row := 0
	for i, char := range input {
		fence[row][i] = char
		row += direction
		if row == 0 || row == rails-1 {
			direction *= -1
		}
	}
	var result strings.Builder
	for _, r := range fence {
		for _, char := range r {
			if char != 0 {
				result.WriteRune(char)
			}
		}
	}
	return result.String()
}

func railFenceDecodeFunc(input string, rails int) string {
	fence := make([][]rune, rails)
	for i := range fence {
		fence[i] = make([]rune, len(input))
	}
	direction := 1
	row := 0
	for i := range input {
		fence[row][i] = '*'
		row += direction
		if row == 0 || row == rails-1 {
			direction *= -1
		}
	}
	index := 0
	for r := 0; r < rails; r++ {
		for c := 0; c < len(input); c++ {
			if fence[r][c] == '*' {
				fence[r][c] = rune(input[index])
				index++
			}
		}
	}
	var result strings.Builder
	row = 0
	direction = 1
	for i := 0; i < len(input); i++ {
		result.WriteRune(fence[row][i])
		row += direction
		if row == 0 || row == rails-1 {
			direction *= -1
		}
	}
	return result.String()
}
