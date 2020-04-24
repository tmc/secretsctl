// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var DestroySecretVersionInput secretmanagerpb.DestroySecretVersionRequest

var DestroySecretVersionFromFile string

func init() {
	SecretManagerServiceCmd.AddCommand(DestroySecretVersionCmd)

	DestroySecretVersionCmd.Flags().StringVar(&DestroySecretVersionInput.Name, "name", "", "Required. Required. The resource name of the...")

	DestroySecretVersionCmd.Flags().StringVar(&DestroySecretVersionFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var DestroySecretVersionCmd = &cobra.Command{
	Use:   "destroy-secret-version",
	Short: "Destroys a...",
	Long:  "Destroys a [SecretVersion][google.cloud.secretmanager.v1.SecretVersion].   Sets the [state][google.cloud.secretmanager.v1.SecretVersion.state] of the...",
	PreRun: func(cmd *cobra.Command, args []string) {

		if DestroySecretVersionFromFile == "" {

			cmd.MarkFlagRequired("name")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if DestroySecretVersionFromFile != "" {
			in, err = os.Open(DestroySecretVersionFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &DestroySecretVersionInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("SecretManager", "DestroySecretVersion", &DestroySecretVersionInput)
		}
		resp, err := SecretManagerClient.DestroySecretVersion(ctx, &DestroySecretVersionInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
