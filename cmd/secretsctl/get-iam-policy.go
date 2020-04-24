// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	iampb "google.golang.org/genproto/googleapis/iam/v1"

	"os"
)

var GetIamPolicyInput iampb.GetIamPolicyRequest

var GetIamPolicyFromFile string

func init() {
	SecretManagerServiceCmd.AddCommand(GetIamPolicyCmd)

	GetIamPolicyInput.Options = new(iampb.GetPolicyOptions)

	GetIamPolicyCmd.Flags().StringVar(&GetIamPolicyInput.Resource, "resource", "", "Required. REQUIRED: The resource for which the policy is...")

	GetIamPolicyCmd.Flags().Int32Var(&GetIamPolicyInput.Options.RequestedPolicyVersion, "options.requested_policy_version", 0, "Optional. The policy format version to be...")

	GetIamPolicyCmd.Flags().StringVar(&GetIamPolicyFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var GetIamPolicyCmd = &cobra.Command{
	Use:   "get-iam-policy",
	Short: "Gets the access control policy for a secret. ...",
	Long:  "Gets the access control policy for a secret.  Returns empty policy if the secret exists and does not have a policy set.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if GetIamPolicyFromFile == "" {

			cmd.MarkFlagRequired("resource")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if GetIamPolicyFromFile != "" {
			in, err = os.Open(GetIamPolicyFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &GetIamPolicyInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("SecretManager", "GetIamPolicy", &GetIamPolicyInput)
		}
		resp, err := SecretManagerClient.GetIamPolicy(ctx, &GetIamPolicyInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
