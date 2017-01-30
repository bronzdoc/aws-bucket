package cmd

import (
	"github.com/bronzdoc/bucket/lib/s3"
	"github.com/spf13/cobra"
)

var sortBy string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List AWS S3 buckets",
	Run: func(cmd *cobra.Command, args []string) {

		s3 := s3.New()
		s3.ListBuckets()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(
		&sortBy,
		"sort",
		"s",
		"",
		`
	- dateCreated
	- dateModified
	- alphaAsc
	- alphaDesc

	Example: bucket list --sort dateCreated
		`,
	)
}
