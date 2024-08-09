package outputs

import (
	"fmt"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcpsecretset/model"
	"github.com/plantoncloud/stack-job-runner-golang-sdk/pkg/automationapi/autoapistackoutput"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

func PulumiOutputsToStackOutputsConverter(stackOutput auto.OutputMap,
	input *model.GcpSecretSetStackInput) *model.GcpSecretSetStackOutputs {
	stackOutputs := &model.GcpSecretSetStackOutputs{
		SecretIdMap: make(map[string]string),
	}
	for _, secretName := range input.ApiResource.Spec.SecretNames {
		stackOutputs.SecretIdMap[secretName] = autoapistackoutput.GetVal(stackOutput,
			fmt.Sprintf("%s-id", secretName))
	}
	return stackOutputs
}
