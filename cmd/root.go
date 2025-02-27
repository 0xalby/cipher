package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var file string

var rootCmd = &cobra.Command{
	Use:   "cipher",
	Short: "(En/de)code anything with a variety of ciphers",
	Long:  `(En/de)code anything with a variety of popular ciphers with a mnemonic and constant command syntax`,
}

// This is called by main.main()
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "Input file")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Help for cipher")
	rootCmd.Flags().BoolP("help", "h", false, "Help for cipher")
	rootCmd.SilenceUsage = true
}
