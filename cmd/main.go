package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "gogit"}
	rootCmd.AddCommand(
		internal.addCmd,
	// addCmd,
	// catFileCmd,
	// checkIgnoreCmd,
	// checkoutCmd,
	// commitCmd,
	// hashObjectCmd,
	// initCmd,
	// logCmd,
	// lsFilesCmd,
	// lsTreeCmd,
	// revParseCmd,
	// rmCmd,
	// showRefCmd,
	// statusCmd,
	// tagCmd,
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
