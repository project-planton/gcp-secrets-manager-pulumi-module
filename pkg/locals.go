package pkg

import (
	gcpsecretsmanagerv1 "buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/provider/gcp/gcpsecretsmanager/v1"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	GcpSecretsManager *gcpsecretsmanagerv1.GcpSecretsManager
	GcpLabels         map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *gcpsecretsmanagerv1.GcpSecretsManagerStackInput) *Locals {
	locals := &Locals{}
	locals.GcpSecretsManager = stackInput.Target

	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.Organization: locals.GcpSecretsManager.Spec.EnvironmentInfo.OrgId,
		gcplabelkeys.Environment:  locals.GcpSecretsManager.Spec.EnvironmentInfo.EnvId,
		gcplabelkeys.ResourceKind: "gcp_secrets_manager",
		gcplabelkeys.ResourceId:   locals.GcpSecretsManager.Metadata.Id,
	}
	return locals
}
