// Code generated. DO NOT EDIT.

package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/grpc"

	gapic "github.com/tmc/secretsctl/secrets"
)

var SecretManagerConfig *viper.Viper
var SecretManagerClient *gapic.SecretManagerClient
var SecretManagerSubCommands []string = []string{
	"list-secrets",
	"create-secret",
	"add-secret-version",
	"get-secret",
	"update-secret",
	"delete-secret",
	"list-secret-versions",
	"get-secret-version",
	"access-secret-version",
	"disable-secret-version",
	"enable-secret-version",
	"destroy-secret-version",
	"set-iam-policy",
	"get-iam-policy",
	"test-iam-permissions",
}

func init() {
	rootCmd.AddCommand(SecretManagerServiceCmd)

	SecretManagerConfig = viper.New()
	SecretManagerConfig.SetEnvPrefix("SECRETS_SECRETMANAGER")
	SecretManagerConfig.AutomaticEnv()

	SecretManagerServiceCmd.PersistentFlags().Bool("insecure", false, "Make insecure client connection. Or use SECRETS_SECRETMANAGER_INSECURE. Must be used with \"address\" option")
	SecretManagerConfig.BindPFlag("insecure", SecretManagerServiceCmd.PersistentFlags().Lookup("insecure"))
	SecretManagerConfig.BindEnv("insecure")

	SecretManagerServiceCmd.PersistentFlags().String("address", "", "Set API address used by client. Or use SECRETS_SECRETMANAGER_ADDRESS.")
	SecretManagerConfig.BindPFlag("address", SecretManagerServiceCmd.PersistentFlags().Lookup("address"))
	SecretManagerConfig.BindEnv("address")

	SecretManagerServiceCmd.PersistentFlags().String("token", "", "Set Bearer token used by the client. Or use SECRETS_SECRETMANAGER_TOKEN.")
	SecretManagerConfig.BindPFlag("token", SecretManagerServiceCmd.PersistentFlags().Lookup("token"))
	SecretManagerConfig.BindEnv("token")

	SecretManagerServiceCmd.PersistentFlags().String("api_key", "", "Set API Key used by the client. Or use SECRETS_SECRETMANAGER_API_KEY.")
	SecretManagerConfig.BindPFlag("api_key", SecretManagerServiceCmd.PersistentFlags().Lookup("api_key"))
	SecretManagerConfig.BindEnv("api_key")
}

var SecretManagerServiceCmd = &cobra.Command{
	Use:       "secretmanager",
	Short:     "Secret Manager Service   Manages secrets and...",
	Long:      "Secret Manager Service   Manages secrets and operations using those secrets. Implements a REST  model with the following objects:   *...",
	ValidArgs: SecretManagerSubCommands,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		var opts []option.ClientOption

		address := SecretManagerConfig.GetString("address")
		if address != "" {
			opts = append(opts, option.WithEndpoint(address))
		}

		if SecretManagerConfig.GetBool("insecure") {
			if address == "" {
				return fmt.Errorf("Missing address to use with insecure connection")
			}

			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				return err
			}
			opts = append(opts, option.WithGRPCConn(conn))
		}

		if token := SecretManagerConfig.GetString("token"); token != "" {
			opts = append(opts, option.WithTokenSource(oauth2.StaticTokenSource(
				&oauth2.Token{
					AccessToken: token,
					TokenType:   "Bearer",
				})))
		}

		if key := SecretManagerConfig.GetString("api_key"); key != "" {
			opts = append(opts, option.WithAPIKey(key))
		}

		SecretManagerClient, err = gapic.NewSecretManagerClient(ctx, opts...)
		return
	},
}
