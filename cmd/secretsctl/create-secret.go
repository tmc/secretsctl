// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	durationpb "github.com/golang/protobuf/ptypes/duration"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	"strings"

	timestamppb "github.com/golang/protobuf/ptypes/timestamp"
)

var CreateSecretInput secretmanagerpb.CreateSecretRequest

var CreateSecretFromFile string

var CreateSecretInputSecretExpiration string

var CreateSecretInputSecretExpirationExpireTime secretmanagerpb.Secret_ExpireTime

var CreateSecretInputSecretExpirationTtl secretmanagerpb.Secret_Ttl

var CreateSecretInputSecretReplicationReplication string

var CreateSecretInputSecretReplicationReplicationAutomatic secretmanagerpb.Replication_Automatic_

var CreateSecretInputSecretReplicationReplicationUserManaged secretmanagerpb.Replication_UserManaged_

var CreateSecretInputSecretReplicationReplicationUserManagedReplicas []string

var CreateSecretInputSecretLabels []string

func init() {
	SecretManagerServiceCmd.AddCommand(CreateSecretCmd)

	CreateSecretInput.Secret = new(secretmanagerpb.Secret)

	CreateSecretInput.Secret.Replication = new(secretmanagerpb.Replication)

	CreateSecretInputSecretReplicationReplicationAutomatic.Automatic = new(secretmanagerpb.Replication_Automatic)

	CreateSecretInputSecretReplicationReplicationAutomatic.Automatic.CustomerManagedEncryption = new(secretmanagerpb.CustomerManagedEncryption)

	CreateSecretInputSecretReplicationReplicationUserManaged.UserManaged = new(secretmanagerpb.Replication_UserManaged)

	CreateSecretInputSecretExpirationExpireTime.ExpireTime = new(timestamppb.Timestamp)

	CreateSecretInputSecretExpirationTtl.Ttl = new(durationpb.Duration)

	CreateSecretCmd.Flags().StringVar(&CreateSecretInput.Parent, "parent", "", "Required. Required. The resource name of the project to...")

	CreateSecretCmd.Flags().StringVar(&CreateSecretInput.SecretId, "secret_id", "", "Required. Required. This must be unique within the project....")

	CreateSecretCmd.Flags().StringVar(&CreateSecretInputSecretReplicationReplicationAutomatic.Automatic.CustomerManagedEncryption.KmsKeyName, "secret.replication.replication.automatic.customer_managed_encryption.kms_key_name", "", "Required. Required. The resource name of the Cloud KMS...")

	CreateSecretCmd.Flags().StringArrayVar(&CreateSecretInputSecretReplicationReplicationUserManagedReplicas, "secret.replication.replication.user_managed.replicas", []string{}, "Required. Required. The list of Replicas for this...")

	CreateSecretCmd.Flags().StringArrayVar(&CreateSecretInputSecretLabels, "secret.labels", []string{}, "key=value pairs. The labels assigned to this Secret.   Label keys...")

	CreateSecretCmd.Flags().Int64Var(&CreateSecretInputSecretExpirationExpireTime.ExpireTime.Seconds, "secret.expiration.expire_time.seconds", 0, "Represents seconds of UTC time since Unix epoch ...")

	CreateSecretCmd.Flags().Int32Var(&CreateSecretInputSecretExpirationExpireTime.ExpireTime.Nanos, "secret.expiration.expire_time.nanos", 0, "Non-negative fractions of a second at nanosecond...")

	CreateSecretCmd.Flags().Int64Var(&CreateSecretInputSecretExpirationTtl.Ttl.Seconds, "secret.expiration.ttl.seconds", 0, "Signed seconds of the span of time. Must be from...")

	CreateSecretCmd.Flags().Int32Var(&CreateSecretInputSecretExpirationTtl.Ttl.Nanos, "secret.expiration.ttl.nanos", 0, "Signed fractions of a second at nanosecond...")

	CreateSecretCmd.Flags().StringVar(&CreateSecretInputSecretExpiration, "secret.expiration", "", "Choices: expire_time, ttl")

	CreateSecretCmd.Flags().StringVar(&CreateSecretInputSecretReplicationReplication, "secret.replication.replication", "", "Choices: automatic, user_managed")

	CreateSecretCmd.Flags().StringVar(&CreateSecretFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var CreateSecretCmd = &cobra.Command{
	Use:   "create-secret",
	Short: "Creates a new...",
	Long:  "Creates a new [Secret][google.cloud.secretmanager.v1.Secret] containing no [SecretVersions][google.cloud.secretmanager.v1.SecretVersion].",
	PreRun: func(cmd *cobra.Command, args []string) {

		if CreateSecretFromFile == "" {

			cmd.MarkFlagRequired("parent")

			cmd.MarkFlagRequired("secret_id")

			cmd.MarkFlagRequired("secret.expiration")

			cmd.MarkFlagRequired("secret.replication.replication")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if CreateSecretFromFile != "" {
			in, err = os.Open(CreateSecretFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &CreateSecretInput)
			if err != nil {
				return err
			}

		} else {

			switch CreateSecretInputSecretExpiration {

			case "expire_time":
				CreateSecretInput.Secret.Expiration = &CreateSecretInputSecretExpirationExpireTime

			case "ttl":
				CreateSecretInput.Secret.Expiration = &CreateSecretInputSecretExpirationTtl

			default:
				return fmt.Errorf("Missing oneof choice for secret.expiration")
			}

			switch CreateSecretInputSecretReplicationReplication {

			case "automatic":
				CreateSecretInput.Secret.Replication.Replication = &CreateSecretInputSecretReplicationReplicationAutomatic

			case "user_managed":
				CreateSecretInput.Secret.Replication.Replication = &CreateSecretInputSecretReplicationReplicationUserManaged

			default:
				return fmt.Errorf("Missing oneof choice for secret.replication.replication")
			}

		}

		// unmarshal JSON strings into slice of structs
		for _, item := range CreateSecretInputSecretReplicationReplicationUserManagedReplicas {
			tmp := secretmanagerpb.Replication_UserManaged_Replica{}
			err = jsonpb.UnmarshalString(item, &tmp)
			if err != nil {
				return
			}

			CreateSecretInputSecretReplicationReplicationUserManaged.UserManaged.Replicas = append(CreateSecretInputSecretReplicationReplicationUserManaged.UserManaged.Replicas, &tmp)
		}

		if len(CreateSecretInputSecretLabels) > 0 {
			CreateSecretInput.Secret.Labels = make(map[string]string)
		}
		for _, item := range CreateSecretInputSecretLabels {
			split := strings.Split(item, "=")
			if len(split) < 2 {
				err = fmt.Errorf("Invalid map item: %q", item)
				return
			}

			CreateSecretInput.Secret.Labels[split[0]] = split[1]
		}

		if Verbose {
			printVerboseInput("SecretManager", "CreateSecret", &CreateSecretInput)
		}
		resp, err := SecretManagerClient.CreateSecret(ctx, &CreateSecretInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
