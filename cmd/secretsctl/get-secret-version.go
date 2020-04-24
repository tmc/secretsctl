// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var GetSecretVersionInput secretmanagerpb.GetSecretVersionRequest

var GetSecretVersionFromFile string

func init() {
	SecretManagerServiceCmd.AddCommand(GetSecretVersionCmd)

	GetSecretVersionCmd.Flags().StringVar(&GetSecretVersionInput.Name, "name", "", "Required. Required. The resource name of the...")

	GetSecretVersionCmd.Flags().StringVar(&GetSecretVersionFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var GetSecretVersionCmd = &cobra.Command{
	Use:   "get-secret-version",
	Short: "Gets metadata for a...",
	Long:  "Gets metadata for a [SecretVersion][google.cloud.secretmanager.v1.SecretVersion].   `projects/*/secrets/*/versions/latest` is an alias to the...",
	PreRun: func(cmd *cobra.Command, args []string) {

		if GetSecretVersionFromFile == "" {

			cmd.MarkFlagRequired("name")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if GetSecretVersionFromFile != "" {
			in, err = os.Open(GetSecretVersionFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &GetSecretVersionInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("SecretManager", "GetSecretVersion", &GetSecretVersionInput)
		}
		resp, err := SecretManagerClient.GetSecretVersion(ctx, &GetSecretVersionInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
