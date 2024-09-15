package pkg

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcpsecretsmanager"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/enums/apiresourcekind"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	GcpSecretsManager *gcpsecretsmanager.GcpSecretsManager
	GcpLabels         map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *gcpsecretsmanager.GcpSecretsManagerStackInput) *Locals {
	locals := &Locals{}
	locals.GcpSecretsManager = stackInput.Target

	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.Organization: locals.GcpSecretsManager.Spec.EnvironmentInfo.OrgId,
		gcplabelkeys.Environment:  locals.GcpSecretsManager.Spec.EnvironmentInfo.EnvId,
		gcplabelkeys.ResourceKind: apiresourcekind.ApiResourceKind_gcp_secrets_manager.String(),
		gcplabelkeys.ResourceId:   locals.GcpSecretsManager.Metadata.Id,
	}
	return locals
}
