// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	field_maskpb "google.golang.org/genproto/protobuf/field_mask"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var UpdateSecretInput secretmanagerpb.UpdateSecretRequest

var UpdateSecretFromFile string

var UpdateSecretInputSecretReplicationReplication string

var UpdateSecretInputSecretReplicationReplicationAutomatic secretmanagerpb.Replication_Automatic

var UpdateSecretInputSecretReplicationReplicationUserManaged secretmanagerpb.Replication_UserManaged

var UpdateSecretInputSecretReplicationReplicationUserManagedReplicas []string

var UpdateSecretInputSecretLabels []string

func init() {
	SecretManagerServiceCmd.AddCommand(UpdateSecretCmd)

	UpdateSecretInput.Secret = new(secretmanagerpb.Secret)

	UpdateSecretInput.Secret.Replication = new(secretmanagerpb.Replication)

	UpdateSecretInputSecretReplicationReplicationAutomatic.Automatic = new(secretmanagerpb.Replication_Automatic)

	UpdateSecretInputSecretReplicationReplicationUserManaged.UserManaged = new(secretmanagerpb.Replication_UserManaged)

	UpdateSecretInput.UpdateMask = new(field_maskpb.FieldMask)

	UpdateSecretCmd.Flags().StringArrayVar(&UpdateSecretInputSecretReplicationReplicationUserManagedReplicas, "secret.replication.replication.user_managed.replicas", []string{}, "Required. Required. The list of Replicas for this...")

	UpdateSecretCmd.Flags().StringArrayVar(&UpdateSecretInputSecretLabels, "secret.labels", []string{}, "The labels assigned to this Secret.   Label keys...")

	UpdateSecretCmd.Flags().StringSliceVar(&UpdateSecretInput.UpdateMask.Paths, "update_mask.paths", []string{}, "The set of field mask paths.")

	UpdateSecretCmd.Flags().StringVar(&UpdateSecretInputSecretReplicationReplication, "secret.replication.replication", "", "Choices: automatic, user_managed")

	UpdateSecretCmd.Flags().StringVar(&UpdateSecretFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var UpdateSecretCmd = &cobra.Command{
	Use:   "update-secret",
	Short: "Updates metadata of an existing...",
	Long:  "Updates metadata of an existing [Secret][google.cloud.secretmanager.v1.Secret].",
	PreRun: func(cmd *cobra.Command, args []string) {

		if UpdateSecretFromFile == "" {

			cmd.MarkFlagRequired("secret.replication.replication")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if UpdateSecretFromFile != "" {
			in, err = os.Open(UpdateSecretFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &UpdateSecretInput)
			if err != nil {
				return err
			}

		} else {

			switch UpdateSecretInputSecretReplicationReplication {

			case "automatic":
				UpdateSecretInput.Secret.Replication.Replication = &UpdateSecretInputSecretReplicationReplicationAutomatic

			case "user_managed":
				UpdateSecretInput.Secret.Replication.Replication = &UpdateSecretInputSecretReplicationReplicationUserManaged

			default:
				return fmt.Errorf("Missing oneof choice for secret.replication.replication")
			}

		}

		// unmarshal JSON strings into slice of structs
		for _, item := range UpdateSecretInputSecretReplicationReplicationUserManagedReplicas {
			tmp := secretmanagerpb.Replication_UserManaged_Replica{}
			err = jsonpb.UnmarshalString(item, &tmp)
			if err != nil {
				return
			}

			UpdateSecretInputSecretReplicationReplicationUserManaged.UserManaged.Replicas = append(UpdateSecretInputSecretReplicationReplicationUserManaged.UserManaged.Replicas, &tmp)
		}

		// unmarshal JSON strings into slice of structs
		for _, item := range UpdateSecretInputSecretLabels {
			tmp := secretmanagerpb.Secret_LabelsEntry{}
			err = jsonpb.UnmarshalString(item, &tmp)
			if err != nil {
				return
			}

			UpdateSecretInput.Secret.Labels = append(UpdateSecretInput.Secret.Labels, &tmp)
		}

		if Verbose {
			printVerboseInput("SecretManager", "UpdateSecret", &UpdateSecretInput)
		}
		resp, err := SecretManagerClient.UpdateSecret(ctx, &UpdateSecretInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
