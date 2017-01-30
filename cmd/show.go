package cmd

import (
	"fmt"
	"github.com/bronzdoc/bucket/lib/s3"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show BUCKET_NAME",
	Short: "Show AWS bucket objects",
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle show")
}
