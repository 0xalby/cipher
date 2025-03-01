package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	railFenceDecrypt bool
	railFenceRails   int
)

var railFenceCmd = &cobra.Command{
	Use:   "railfence",
	Short: "Encrypt and decrypt using the Rail Fence cipher",
	Long:  `Encrypt and decrypt using the Rail Fence cipher with a mnemonic command syntax`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := readInput(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := railFenceCipher(input, railFenceRails, railFenceDecrypt)
		fmt.Println(result)
	},
	TraverseChildren: true,
	Args:             cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(railFenceCmd)
	railFenceCmd.Flags().BoolVarP(&railFenceDecrypt, "decrypt", "d", false, "Decrypt the input")
	railFenceCmd.Flags().IntVarP(&railFenceRails, "rails", "r", 0, "Number of rails for the Rail Fence cipher(required)")
	railFenceCmd.MarkFlagRequired("rails")
}

// The Rail Fence cipher
func railFenceCipher(input string, rails int, decrypt bool) string {
	if rails < 2 {
		return "error number of rails must be at least 2."
	}
	if decrypt {
		return railFenceDecryptFunc(input, rails)
	}
	return railFenceEncryptFunc(input, rails)
}

// Rail Fenc encryption
func railFenceEncryptFunc(input string, rails int) string {
	// Creating a fence to represent the rails
	fence := make([][]rune, rails)
	for i := range fence {
		fence[i] = make([]rune, len(input))
	}
	direction := 1
	row := 0
	// Filling the fence
	for i, char := range input {
		fence[row][i] = char
		row += direction
		if row == 0 || row == rails-1 {
			direction *= -1
		}
	}
	// Bulding the encrypted result
	var result strings.Builder
	for _, r := range fence {
		for _, char := range r {
			// Skipping over empty slots
			if char != 0 {
				result.WriteRune(char)
			}
		}
	}
	return result.String()
}

// Rail Fence decryption
func railFenceDecryptFunc(input string, rails int) string {
	// Creating a fence to represent the rails
	fence := make([][]rune, rails)
	for i := range fence {
		fence[i] = make([]rune, len(input))
	}
	direction := 1
	row := 0
	// Marking the positions in the fence to position the characters
	for i := range input {
		// Setting the placeholder
		fence[row][i] = '*'
		row += direction
		// Reversing the direction when hitting the top or the bottom rail
		if row == 0 || row == rails-1 {
			direction *= -1
		}
	}
	// Filling the fence
	index := 0
	for r := 0; r < rails; r++ {
		for c := 0; c < len(input); c++ {
			if fence[r][c] == '*' {
				fence[r][c] = rune(input[index])
				index++
			}
		}
	}
	// Bulding the decrypted result
	var result strings.Builder
	row = 0
	direction = 1
	for i := 0; i < len(input); i++ {
		result.WriteRune(fence[row][i])
		row += direction
		// Reversing the direction when hitting the top or the bottom rail
		if row == 0 || row == rails-1 {
			direction *= -1
		}
	}
	return result.String()
}
