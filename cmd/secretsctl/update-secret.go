// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	durationpb "github.com/golang/protobuf/ptypes/duration"

	field_maskpb "google.golang.org/genproto/protobuf/field_mask"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	"strings"

	timestamppb "github.com/golang/protobuf/ptypes/timestamp"
)

var UpdateSecretInput secretmanagerpb.UpdateSecretRequest

var UpdateSecretFromFile string

var UpdateSecretInputSecretExpiration string

var UpdateSecretInputSecretExpirationExpireTime timestamppb.Secret_ExpireTime

var UpdateSecretInputSecretExpirationTtl durationpb.Secret_Ttl

var UpdateSecretInputSecretReplicationReplication string

var UpdateSecretInputSecretReplicationReplicationAutomatic secretmanagerpb.Replication_Automatic_

var UpdateSecretInputSecretReplicationReplicationUserManaged secretmanagerpb.Replication_UserManaged_

var UpdateSecretInputSecretReplicationReplicationUserManagedReplicas []string

var UpdateSecretInputSecretLabels []string

func init() {
	SecretManagerServiceCmd.AddCommand(UpdateSecretCmd)

	UpdateSecretInput.Secret = new(secretmanagerpb.Secret)

	UpdateSecretInput.Secret.Replication = new(secretmanagerpb.Replication)

	UpdateSecretInputSecretReplicationReplicationAutomatic.Automatic = new(secretmanagerpb.Replication_Automatic)

	UpdateSecretInputSecretReplicationReplicationAutomatic.Automatic.CustomerManagedEncryption = new(secretmanagerpb.CustomerManagedEncryption)

	UpdateSecretInputSecretReplicationReplicationUserManaged.UserManaged = new(secretmanagerpb.Replication_UserManaged)

	UpdateSecretInputSecretExpirationExpireTime.ExpireTime = new(timestamppb.Timestamp)

	UpdateSecretInputSecretExpirationTtl.Ttl = new(durationpb.Duration)

	UpdateSecretInput.UpdateMask = new(field_maskpb.FieldMask)

	UpdateSecretCmd.Flags().StringVar(&UpdateSecretInputSecretReplicationReplicationAutomatic.Automatic.CustomerManagedEncryption.KmsKeyName, "secret.replication.replication.automatic.customer_managed_encryption.kms_key_name", "", "Required. Required. The resource name of the Cloud KMS...")

	UpdateSecretCmd.Flags().StringArrayVar(&UpdateSecretInputSecretReplicationReplicationUserManagedReplicas, "secret.replication.replication.user_managed.replicas", []string{}, "Required. Required. The list of Replicas for this...")

	UpdateSecretCmd.Flags().StringArrayVar(&UpdateSecretInputSecretLabels, "secret.labels", []string{}, "key=value pairs. The labels assigned to this Secret.   Label keys...")

	UpdateSecretCmd.Flags().Int64Var(&UpdateSecretInputSecretExpirationExpireTime.ExpireTime.Seconds, "secret.expiration.expire_time.seconds", 0, "Represents seconds of UTC time since Unix epoch ...")

	UpdateSecretCmd.Flags().Int32Var(&UpdateSecretInputSecretExpirationExpireTime.ExpireTime.Nanos, "secret.expiration.expire_time.nanos", 0, "Non-negative fractions of a second at nanosecond...")

	UpdateSecretCmd.Flags().Int64Var(&UpdateSecretInputSecretExpirationTtl.Ttl.Seconds, "secret.expiration.ttl.seconds", 0, "Signed seconds of the span of time. Must be from...")

	UpdateSecretCmd.Flags().Int32Var(&UpdateSecretInputSecretExpirationTtl.Ttl.Nanos, "secret.expiration.ttl.nanos", 0, "Signed fractions of a second at nanosecond...")

	UpdateSecretCmd.Flags().StringSliceVar(&UpdateSecretInput.UpdateMask.Paths, "update_mask.paths", []string{}, "The set of field mask paths.")

	UpdateSecretCmd.Flags().StringVar(&UpdateSecretInputSecretExpiration, "secret.expiration", "", "Choices: expire_time, ttl")

	UpdateSecretCmd.Flags().StringVar(&UpdateSecretInputSecretReplicationReplication, "secret.replication.replication", "", "Choices: automatic, user_managed")

	UpdateSecretCmd.Flags().StringVar(&UpdateSecretFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var UpdateSecretCmd = &cobra.Command{
	Use:   "update-secret",
	Short: "Updates metadata of an existing...",
	Long:  "Updates metadata of an existing [Secret][google.cloud.secretmanager.v1.Secret].",
	PreRun: func(cmd *cobra.Command, args []string) {

		if UpdateSecretFromFile == "" {

			cmd.MarkFlagRequired("secret.expiration")

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

			switch UpdateSecretInputSecretExpiration {

			case "expire_time":
				UpdateSecretInput.Secret.Expiration = &UpdateSecretInputSecretExpirationExpireTime

			case "ttl":
				UpdateSecretInput.Secret.Expiration = &UpdateSecretInputSecretExpirationTtl

			default:
				return fmt.Errorf("Missing oneof choice for secret.expiration")
			}

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

		for _, item := range UpdateSecretInputSecretLabels {
			split := strings.Split(item, "=")
			if len(split) < 2 {
				err = fmt.Errorf("Invalid map item: %q", item)
				return
			}

			UpdateSecretInput.Secret.Labels[split[0]] = split[1]
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
