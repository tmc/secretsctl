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

var ListSecretsInput secretmanagerpb.ListSecretsRequest

var ListSecretsFromFile string

func init() {
	SecretManagerServiceCmd.AddCommand(ListSecretsCmd)

	ListSecretsCmd.Flags().StringVar(&ListSecretsInput.Parent, "parent", "", "Required. Required. The resource name of the project...")

	ListSecretsCmd.Flags().Int32Var(&ListSecretsInput.PageSize, "page_size", 10, "Default is 10. Optional. The maximum number of results to be...")

	ListSecretsCmd.Flags().StringVar(&ListSecretsInput.PageToken, "page_token", "", "Optional. Pagination token, returned earlier via ...")

	ListSecretsCmd.Flags().StringVar(&ListSecretsFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var ListSecretsCmd = &cobra.Command{
	Use:   "list-secrets",
	Short: "Lists...",
	Long:  "Lists [Secrets][google.cloud.secretmanager.v1.Secret].",
	PreRun: func(cmd *cobra.Command, args []string) {

		if ListSecretsFromFile == "" {

			cmd.MarkFlagRequired("parent")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if ListSecretsFromFile != "" {
			in, err = os.Open(ListSecretsFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &ListSecretsInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("SecretManager", "ListSecrets", &ListSecretsInput)
		}
		iter := SecretManagerClient.ListSecrets(ctx, &ListSecretsInput)

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
