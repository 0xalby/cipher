package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	base64Decode bool
)

var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "(En/de)code using Base64",
	Long:  `(En/de)code input using Base64 with a mnemonic command syntax`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := readInput(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := base64Cipher(input, base64Decode)
		fmt.Print(result)
	},
	TraverseChildren: true,
	Args:             cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.Flags().BoolVarP(&base64Decode, "decode", "d", false, "Decode the input")
}

// Base64 implementation
func base64Cipher(input string, decode bool) string {
	if decode {
		decoded, err := base64.StdEncoding.DecodeString(input)
		if err != nil {
			return "error invalid input"
		}
		return string(decoded)
	}
	return base64.StdEncoding.EncodeToString([]byte(input))
}
