// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var AddSecretVersionInput secretmanagerpb.AddSecretVersionRequest

var AddSecretVersionFromFile string

func init() {
	SecretManagerServiceCmd.AddCommand(AddSecretVersionCmd)

	AddSecretVersionInput.Payload = new(secretmanagerpb.SecretPayload)

	AddSecretVersionCmd.Flags().StringVar(&AddSecretVersionInput.Parent, "parent", "", "Required. Required. The resource name of the...")

	AddSecretVersionCmd.Flags().BytesHexVar(&AddSecretVersionInput.Payload.Data, "payload.data", []byte{}, "The secret data. Must be no larger than 64KiB.")

	AddSecretVersionCmd.Flags().StringVar(&AddSecretVersionFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var AddSecretVersionCmd = &cobra.Command{
	Use:   "add-secret-version",
	Short: "Creates a new...",
	Long:  "Creates a new [SecretVersion][google.cloud.secretmanager.v1.SecretVersion] containing secret data and attaches  it to an existing...",
	PreRun: func(cmd *cobra.Command, args []string) {

		if AddSecretVersionFromFile == "" {

			cmd.MarkFlagRequired("parent")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if AddSecretVersionFromFile != "" {
			in, err = os.Open(AddSecretVersionFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &AddSecretVersionInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("SecretManager", "AddSecretVersion", &AddSecretVersionInput)
		}
		resp, err := SecretManagerClient.AddSecretVersion(ctx, &AddSecretVersionInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
