package gcpsecretsmanagersecretset

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcpsecretsmanagersecretset/model"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/english/enums/englishword"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/gcp/pulumigoogleprovider"
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
	gcpProvider, err := pulumigoogleprovider.Get(ctx, s.Input.GcpCredential)
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	gcpSecretsManagerSecretSet := s.Input.ApiResource

	for _, secretName := range gcpSecretsManagerSecretSet.Spec.SecretNames {
		if secretName == "" {
			continue
		}

		secretId := fmt.Sprintf("%s-%s", gcpSecretsManagerSecretSet.Metadata.Id, secretName)

		addedSecret, err := secretmanager.NewSecret(ctx, secretName, &secretmanager.SecretArgs{
			Labels:   pulumi.ToStringMap(s.GcpLabels),
			Project:  pulumi.String(gcpSecretsManagerSecretSet.Spec.ProjectId),
			SecretId: pulumi.String(secretId),
			Replication: secretmanager.SecretReplicationArgs{
				Auto: secretmanager.SecretReplicationAutoArgs{},
			},
		}, pulumi.Provider(gcpProvider))
		if err != nil {
			return errors.Wrap(err, "failed to add secret")
		}
		_, err = secretmanager.NewSecretVersion(ctx, secretId, &secretmanager.SecretVersionArgs{
			Enabled:    pulumi.Bool(true),
			Secret:     addedSecret.Name,
			SecretData: pulumi.String(PlaceholderSecretValue),
		},
			pulumi.Parent(addedSecret),
			pulumi.IgnoreChanges([]string{"secretData"}))
		if err != nil {
			return errors.Wrap(err, "failed to add placeholder secret version")
		}
		ctx.Export(GetSecretIdOutputName(secretName), addedSecret.SecretId)
	}
	return nil
}

func GetSecretIdOutputName(secretName string) string {
	return pulumigoogleprovider.PulumiOutputName(secretmanager.Secret{}, secretName,
		englishword.EnglishWord_id.String())
}
