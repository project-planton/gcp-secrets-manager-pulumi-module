package pkg

import (
	gcpsecretsmanagerv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/gcp/gcpsecretsmanager/v1"
	"github.com/project-planton/pulumi-module-golang-commons/pkg/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	GcpSecretsManager *gcpsecretsmanagerv1.GcpSecretsManager
	GcpLabels         map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *gcpsecretsmanagerv1.GcpSecretsManagerStackInput) *Locals {
	locals := &Locals{}

	//if the id is empty, use name as id
	if stackInput.Target.Metadata.Id == "" {
		stackInput.Target.Metadata.Id = stackInput.Target.Metadata.Name
	}

	locals.GcpSecretsManager = stackInput.Target

	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceKind: "gcp_secrets_manager",
		gcplabelkeys.ResourceId:   locals.GcpSecretsManager.Metadata.Id,
	}

	if locals.GcpSecretsManager.Spec.EnvironmentInfo != nil {
		locals.GcpLabels[gcplabelkeys.Organization] = locals.GcpSecretsManager.Spec.EnvironmentInfo.OrgId
		locals.GcpLabels[gcplabelkeys.Environment] = locals.GcpSecretsManager.Spec.EnvironmentInfo.EnvId
	}

	return locals
}
