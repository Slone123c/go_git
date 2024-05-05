package main

import (
	"fmt"
	"github.com/Slone123c/go_git/internal"
	"github.com/spf13/cobra"
	"os"
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
