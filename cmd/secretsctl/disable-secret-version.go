// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var DisableSecretVersionInput secretmanagerpb.DisableSecretVersionRequest

var DisableSecretVersionFromFile string

func init() {
	SecretManagerServiceCmd.AddCommand(DisableSecretVersionCmd)

	DisableSecretVersionCmd.Flags().StringVar(&DisableSecretVersionInput.Name, "name", "", "Required. Required. The resource name of the...")

	DisableSecretVersionCmd.Flags().StringVar(&DisableSecretVersionFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var DisableSecretVersionCmd = &cobra.Command{
	Use:   "disable-secret-version",
	Short: "Disables a...",
	Long:  "Disables a [SecretVersion][google.cloud.secretmanager.v1.SecretVersion].   Sets the [state][google.cloud.secretmanager.v1.SecretVersion.state] of the...",
	PreRun: func(cmd *cobra.Command, args []string) {

		if DisableSecretVersionFromFile == "" {

			cmd.MarkFlagRequired("name")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if DisableSecretVersionFromFile != "" {
			in, err = os.Open(DisableSecretVersionFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &DisableSecretVersionInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("SecretManager", "DisableSecretVersion", &DisableSecretVersionInput)
		}
		resp, err := SecretManagerClient.DisableSecretVersion(ctx, &DisableSecretVersionInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
