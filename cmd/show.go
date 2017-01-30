package cmd

import (
	"fmt"
	"github.com/bronzdoc/bucket/lib/s3"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show BUCKET_NAME",
	Short: "Show AWS S3 bucket objects",
	Run: func(cmd *cobra.Command, args []string) {
		// Return early if no bucket name specified
		if len(args) < 1 {
			fmt.Println(`"bucket show" requires at least 1 argument(s).
see 'bucket show --help'

Usage: bucket show BUCKET_NAME

show bucket objects`)

			return
		}

		bucketName := args[0]

		s3 := s3.New()
		s3.ShowBucketObjects(bucketName)
	},
}

func init() {
	RootCmd.AddCommand(showCmd)
}
