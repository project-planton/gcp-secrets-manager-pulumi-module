package pkg

import (
	"fmt"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcpsecretsmanagersecretset/model"
	"github.com/plantoncloud/stack-job-runner-golang-sdk/pkg/automationapi/autoapistackoutput"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

func PulumiOutputsToStackOutputsConverter(stackOutput auto.OutputMap,
	input *model.GcpSecretsManagerSecretSetStackInput) *model.GcpSecretsManagerSecretSetStackOutputs {
	stackOutputs := &model.GcpSecretsManagerSecretSetStackOutputs{
		SecretIdMap: make(map[string]string),
	}
	for _, secretName := range input.ApiResource.Spec.SecretNames {
		stackOutputs.SecretIdMap[secretName] = autoapistackoutput.GetVal(stackOutput,
			fmt.Sprintf("%s-id", secretName))
	}
	return stackOutputs
}
