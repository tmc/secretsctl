// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var GetSecretInput secretmanagerpb.GetSecretRequest

var GetSecretFromFile string

func init() {
	SecretManagerServiceCmd.AddCommand(GetSecretCmd)

	GetSecretCmd.Flags().StringVar(&GetSecretInput.Name, "name", "", "Required. Required. The resource name of the...")

	GetSecretCmd.Flags().StringVar(&GetSecretFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var GetSecretCmd = &cobra.Command{
	Use:   "get-secret",
	Short: "Gets metadata for a given...",
	Long:  "Gets metadata for a given [Secret][google.cloud.secretmanager.v1.Secret].",
	PreRun: func(cmd *cobra.Command, args []string) {

		if GetSecretFromFile == "" {

			cmd.MarkFlagRequired("name")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if GetSecretFromFile != "" {
			in, err = os.Open(GetSecretFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &GetSecretInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("SecretManager", "GetSecret", &GetSecretInput)
		}
		resp, err := SecretManagerClient.GetSecret(ctx, &GetSecretInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
