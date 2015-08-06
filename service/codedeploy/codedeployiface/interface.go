package codedeployiface

import (
	"github.com/datacratic/aws-sdk-go/service/codedeploy"
)

type CodeDeployAPI interface {
	BatchGetApplications(*codedeploy.BatchGetApplicationsInput) (*codedeploy.BatchGetApplicationsOutput, error)

	BatchGetDeployments(*codedeploy.BatchGetDeploymentsInput) (*codedeploy.BatchGetDeploymentsOutput, error)

	CreateApplication(*codedeploy.CreateApplicationInput) (*codedeploy.CreateApplicationOutput, error)

	CreateDeployment(*codedeploy.CreateDeploymentInput) (*codedeploy.CreateDeploymentOutput, error)

	CreateDeploymentConfig(*codedeploy.CreateDeploymentConfigInput) (*codedeploy.CreateDeploymentConfigOutput, error)

	CreateDeploymentGroup(*codedeploy.CreateDeploymentGroupInput) (*codedeploy.CreateDeploymentGroupOutput, error)

	DeleteApplication(*codedeploy.DeleteApplicationInput) (*codedeploy.DeleteApplicationOutput, error)

	DeleteDeploymentConfig(*codedeploy.DeleteDeploymentConfigInput) (*codedeploy.DeleteDeploymentConfigOutput, error)

	DeleteDeploymentGroup(*codedeploy.DeleteDeploymentGroupInput) (*codedeploy.DeleteDeploymentGroupOutput, error)

	GetApplication(*codedeploy.GetApplicationInput) (*codedeploy.GetApplicationOutput, error)

	GetApplicationRevision(*codedeploy.GetApplicationRevisionInput) (*codedeploy.GetApplicationRevisionOutput, error)

	GetDeployment(*codedeploy.GetDeploymentInput) (*codedeploy.GetDeploymentOutput, error)

	GetDeploymentConfig(*codedeploy.GetDeploymentConfigInput) (*codedeploy.GetDeploymentConfigOutput, error)

	GetDeploymentGroup(*codedeploy.GetDeploymentGroupInput) (*codedeploy.GetDeploymentGroupOutput, error)

	GetDeploymentInstance(*codedeploy.GetDeploymentInstanceInput) (*codedeploy.GetDeploymentInstanceOutput, error)

	ListApplicationRevisions(*codedeploy.ListApplicationRevisionsInput) (*codedeploy.ListApplicationRevisionsOutput, error)

	ListApplications(*codedeploy.ListApplicationsInput) (*codedeploy.ListApplicationsOutput, error)

	ListDeploymentConfigs(*codedeploy.ListDeploymentConfigsInput) (*codedeploy.ListDeploymentConfigsOutput, error)

	ListDeploymentGroups(*codedeploy.ListDeploymentGroupsInput) (*codedeploy.ListDeploymentGroupsOutput, error)

	ListDeploymentInstances(*codedeploy.ListDeploymentInstancesInput) (*codedeploy.ListDeploymentInstancesOutput, error)

	ListDeployments(*codedeploy.ListDeploymentsInput) (*codedeploy.ListDeploymentsOutput, error)

	RegisterApplicationRevision(*codedeploy.RegisterApplicationRevisionInput) (*codedeploy.RegisterApplicationRevisionOutput, error)

	StopDeployment(*codedeploy.StopDeploymentInput) (*codedeploy.StopDeploymentOutput, error)

	UpdateApplication(*codedeploy.UpdateApplicationInput) (*codedeploy.UpdateApplicationOutput, error)

	UpdateDeploymentGroup(*codedeploy.UpdateDeploymentGroupInput) (*codedeploy.UpdateDeploymentGroupOutput, error)
}