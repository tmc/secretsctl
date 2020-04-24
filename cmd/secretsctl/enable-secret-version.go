// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var EnableSecretVersionInput secretmanagerpb.EnableSecretVersionRequest

var EnableSecretVersionFromFile string

func init() {
	SecretManagerServiceCmd.AddCommand(EnableSecretVersionCmd)

	EnableSecretVersionCmd.Flags().StringVar(&EnableSecretVersionInput.Name, "name", "", "Required. Required. The resource name of the...")

	EnableSecretVersionCmd.Flags().StringVar(&EnableSecretVersionFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var EnableSecretVersionCmd = &cobra.Command{
	Use:   "enable-secret-version",
	Short: "Enables a...",
	Long:  "Enables a [SecretVersion][google.cloud.secretmanager.v1.SecretVersion].   Sets the [state][google.cloud.secretmanager.v1.SecretVersion.state] of the...",
	PreRun: func(cmd *cobra.Command, args []string) {

		if EnableSecretVersionFromFile == "" {

			cmd.MarkFlagRequired("name")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if EnableSecretVersionFromFile != "" {
			in, err = os.Open(EnableSecretVersionFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &EnableSecretVersionInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("SecretManager", "EnableSecretVersion", &EnableSecretVersionInput)
		}
		resp, err := SecretManagerClient.EnableSecretVersion(ctx, &EnableSecretVersionInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
