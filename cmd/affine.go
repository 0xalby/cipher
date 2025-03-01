package cmd

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var (
	affineDecrypt bool
	affineA       int
	affineB       int
)

var affineCmd = &cobra.Command{
	Use:   "affine",
	Short: "Encrypt and decrypt using the Affine cipher",
	Long:  `Encrypt and decrypt text using the Affine cipher with a mnemonic command syntax`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := readInput(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		result, err := affineCipher(input, affineA, affineB, affineDecrypt)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(affineCmd)
	affineCmd.Flags().BoolVarP(&affineDecrypt, "decrypt", "d", false, "Decrypt the input instead of encrypting")
	affineCmd.Flags().IntVarP(&affineA, "a", "a", 1, "Key 'a' for the Affine cipher(must be coprime with 26 and defaults to 1)")
	affineCmd.Flags().IntVarP(&affineB, "b", "b", 0, "Key 'b' for the Affine cipher")
	affineCmd.MarkFlagRequired("a")
	affineCmd.MarkFlagRequired("b")
}

// The Affine cipher
func affineCipher(input string, a, b int, decrypt bool) (string, error) {
	// Ensure 'a' is coprime with 26
	if gcd(a, 26) != 1 {
		return "", errors.New("key 'a' must be coprime with 26")
	}
	return strings.Map(func(r rune) rune {
		// Skipping non alphabet characters
		if !unicode.IsLetter(r) {
			return r
		}
		base := 'a'
		if unicode.IsUpper(r) {
			base = 'A'
		}
		// Getting the character position in the alphabet
		x := int(r - base)
		var result int
		if decrypt {
			// Decryption formula is D(x) = a^-1 * (x - b) % 26
			aInv := modInverse(a, 26)
			result = (aInv * (x - b + 26)) % 26
		} else {
			// Encryption formula is E(x) = (a * x + b) % 26
			result = (a*x + b) % 26
		}
		// Handling negative results
		if result < 0 {
			result += 26
		}
		return base + rune(result)
	}, input), nil
}

// Computes the greatest common divisor of two numbers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Computes the modular multiplicative inverse of a modulo m
func modInverse(a, m int) int {
	a = a % m
	for x := 1; x < m; x++ {
		if (a*x)%m == 1 {
			return x
		}
	}
	return 1
}
