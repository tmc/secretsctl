// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"google.golang.org/api/iterator"

	"os"

	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var ListSecretVersionsInput secretmanagerpb.ListSecretVersionsRequest

var ListSecretVersionsFromFile string

func init() {
	SecretManagerServiceCmd.AddCommand(ListSecretVersionsCmd)

	ListSecretVersionsCmd.Flags().StringVar(&ListSecretVersionsInput.Parent, "parent", "", "Required. Required. The resource name of the...")

	ListSecretVersionsCmd.Flags().Int32Var(&ListSecretVersionsInput.PageSize, "page_size", 10, "Default is 10. Optional. The maximum number of results to be...")

	ListSecretVersionsCmd.Flags().StringVar(&ListSecretVersionsInput.PageToken, "page_token", "", "Optional. Pagination token, returned earlier via ...")

	ListSecretVersionsCmd.Flags().StringVar(&ListSecretVersionsFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var ListSecretVersionsCmd = &cobra.Command{
	Use:   "list-secret-versions",
	Short: "Lists...",
	Long:  "Lists [SecretVersions][google.cloud.secretmanager.v1.SecretVersion]. This call does not return secret  data.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if ListSecretVersionsFromFile == "" {

			cmd.MarkFlagRequired("parent")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if ListSecretVersionsFromFile != "" {
			in, err = os.Open(ListSecretVersionsFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &ListSecretVersionsInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("SecretManager", "ListSecretVersions", &ListSecretVersionsInput)
		}
		iter := SecretManagerClient.ListSecretVersions(ctx, &ListSecretVersionsInput)

		// populate iterator with a page
		_, err = iter.Next()
		if err != nil && err != iterator.Done {
			return err
		}

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(iter.Response)

		return err
	},
}
