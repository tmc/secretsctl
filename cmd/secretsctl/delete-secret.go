// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"github.com/golang/protobuf/jsonpb"

	"os"

	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var DeleteSecretInput secretmanagerpb.DeleteSecretRequest

var DeleteSecretFromFile string

func init() {
	SecretManagerServiceCmd.AddCommand(DeleteSecretCmd)

	DeleteSecretCmd.Flags().StringVar(&DeleteSecretInput.Name, "name", "", "Required. Required. The resource name of the...")

	DeleteSecretCmd.Flags().StringVar(&DeleteSecretFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var DeleteSecretCmd = &cobra.Command{
	Use:   "delete-secret",
	Short: "Deletes a...",
	Long:  "Deletes a [Secret][google.cloud.secretmanager.v1.Secret].",
	PreRun: func(cmd *cobra.Command, args []string) {

		if DeleteSecretFromFile == "" {

			cmd.MarkFlagRequired("name")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if DeleteSecretFromFile != "" {
			in, err = os.Open(DeleteSecretFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &DeleteSecretInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("SecretManager", "DeleteSecret", &DeleteSecretInput)
		}
		err = SecretManagerClient.DeleteSecret(ctx, &DeleteSecretInput)

		return err
	},
}
