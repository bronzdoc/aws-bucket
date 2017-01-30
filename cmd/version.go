package cmd

import (
	"fmt"
	"github.com/bronzdoc/bucket/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Bucket",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Bucket %s\n", version.Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
