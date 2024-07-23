package gcpsecretsmanagersecretset

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcpsecretsmanagersecretset/model"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/enums/stackjoboperationtype"
	"github.com/plantoncloud/stack-job-runner-golang-sdk/pkg/automationapi/autoapistackoutput"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

func OutputMapTransformer(stackOutput auto.OutputMap,
	input *model.GcpSecretsManagerSecretSetStackInput) *model.GcpSecretsManagerSecretSetStackOutputs {
	if input.StackJob.Spec.OperationType != stackjoboperationtype.StackJobOperationType_apply || stackOutput == nil {
		return &model.GcpSecretsManagerSecretSetStackOutputs{}
	}
	stackOutputs := &model.GcpSecretsManagerSecretSetStackOutputs{
		SecretIdMap: make(map[string]string),
	}
	for _, s := range input.ApiResource.Spec.SecretNames {
		stackOutputs.SecretIdMap[s] = autoapistackoutput.GetVal(stackOutput, GetSecretIdOutputName(s))
	}
	return stackOutputs
}
