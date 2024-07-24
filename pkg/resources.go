package pkg

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcpsecretsmanagersecretset/model"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/secretmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	PlaceholderSecretValue = "placeholder"
)

type ResourceStack struct {
	Input     *model.GcpSecretsManagerSecretSetStackInput
	GcpLabels map[string]string
}

func (s *ResourceStack) Resources(ctx *pulumi.Context) error {
	//create gcp provider using the credentials from the input
	gcpProvider, err := pulumigoogleprovider.Get(ctx, s.Input.GcpCredential)
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	//create a variable with descriptive name for the api-resource
	gcpSecretsManagerSecretSet := s.Input.ApiResource

	//for each secret in the input spec, create a secret on gcp secrets-manager
	for _, secretName := range gcpSecretsManagerSecretSet.Spec.SecretNames {
		if secretName == "" {
			continue
		}

		//construct the id of the secret to make it unique with in the google cloud project
		secretId := fmt.Sprintf("%s-%s", gcpSecretsManagerSecretSet.Metadata.Id, secretName)

		//create the secret resource
		createdSecret, err := secretmanager.NewSecret(ctx, secretName, &secretmanager.SecretArgs{
			Labels:   pulumi.ToStringMap(s.GcpLabels),
			Project:  pulumi.String(gcpSecretsManagerSecretSet.Spec.ProjectId),
			SecretId: pulumi.String(secretId),
			Replication: secretmanager.SecretReplicationArgs{
				Auto: secretmanager.SecretReplicationAutoArgs{},
			},
		}, pulumi.Provider(gcpProvider))
		if err != nil {
			return errors.Wrap(err, "failed to create secret")
		}

		//create secret-version with a placeholder value
		_, err = secretmanager.NewSecretVersion(ctx, secretId, &secretmanager.SecretVersionArgs{
			Enabled:    pulumi.Bool(true),
			Secret:     createdSecret.Name,
			SecretData: pulumi.String(PlaceholderSecretValue),
		},
			pulumi.Parent(createdSecret),
			pulumi.IgnoreChanges([]string{"secretData"}))
		if err != nil {
			return errors.Wrap(err, "failed to create placeholder secret version")
		}

		//export the id of the secret
		ctx.Export(fmt.Sprintf("%s-id", secretName), createdSecret.SecretId)
	}
	return nil
}