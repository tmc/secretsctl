// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var AccessSecretVersionInput secretmanagerpb.AccessSecretVersionRequest

var AccessSecretVersionFromFile string

func init() {
	SecretManagerServiceCmd.AddCommand(AccessSecretVersionCmd)

	AccessSecretVersionCmd.Flags().StringVar(&AccessSecretVersionInput.Name, "name", "", "Required. Required. The resource name of the...")

	AccessSecretVersionCmd.Flags().StringVar(&AccessSecretVersionFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var AccessSecretVersionCmd = &cobra.Command{
	Use:   "access-secret-version",
	Short: "Accesses a...",
	Long:  "Accesses a [SecretVersion][google.cloud.secretmanager.v1.SecretVersion]. This call returns the secret data.   `projects/*/secrets/*/versions/latest`...",
	PreRun: func(cmd *cobra.Command, args []string) {

		if AccessSecretVersionFromFile == "" {

			cmd.MarkFlagRequired("name")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if AccessSecretVersionFromFile != "" {
			in, err = os.Open(AccessSecretVersionFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &AccessSecretVersionInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("SecretManager", "AccessSecretVersion", &AccessSecretVersionInput)
		}
		resp, err := SecretManagerClient.AccessSecretVersion(ctx, &AccessSecretVersionInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
