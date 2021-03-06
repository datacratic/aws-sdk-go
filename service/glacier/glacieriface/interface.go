package glacieriface

import (
	"github.com/datacratic/aws-sdk-go/service/glacier"
)

type GlacierAPI interface {
	AbortMultipartUpload(*glacier.AbortMultipartUploadInput) (*glacier.AbortMultipartUploadOutput, error)

	CompleteMultipartUpload(*glacier.CompleteMultipartUploadInput) (*glacier.ArchiveCreationOutput, error)

	CreateVault(*glacier.CreateVaultInput) (*glacier.CreateVaultOutput, error)

	DeleteArchive(*glacier.DeleteArchiveInput) (*glacier.DeleteArchiveOutput, error)

	DeleteVault(*glacier.DeleteVaultInput) (*glacier.DeleteVaultOutput, error)

	DeleteVaultNotifications(*glacier.DeleteVaultNotificationsInput) (*glacier.DeleteVaultNotificationsOutput, error)

	DescribeJob(*glacier.DescribeJobInput) (*glacier.GlacierJobDescription, error)

	DescribeVault(*glacier.DescribeVaultInput) (*glacier.DescribeVaultOutput, error)

	GetDataRetrievalPolicy(*glacier.GetDataRetrievalPolicyInput) (*glacier.GetDataRetrievalPolicyOutput, error)

	GetJobOutput(*glacier.GetJobOutputInput) (*glacier.GetJobOutputOutput, error)

	GetVaultNotifications(*glacier.GetVaultNotificationsInput) (*glacier.GetVaultNotificationsOutput, error)

	InitiateJob(*glacier.InitiateJobInput) (*glacier.InitiateJobOutput, error)

	InitiateMultipartUpload(*glacier.InitiateMultipartUploadInput) (*glacier.InitiateMultipartUploadOutput, error)

	ListJobs(*glacier.ListJobsInput) (*glacier.ListJobsOutput, error)

	ListMultipartUploads(*glacier.ListMultipartUploadsInput) (*glacier.ListMultipartUploadsOutput, error)

	ListParts(*glacier.ListPartsInput) (*glacier.ListPartsOutput, error)

	ListVaults(*glacier.ListVaultsInput) (*glacier.ListVaultsOutput, error)

	SetDataRetrievalPolicy(*glacier.SetDataRetrievalPolicyInput) (*glacier.SetDataRetrievalPolicyOutput, error)

	SetVaultNotifications(*glacier.SetVaultNotificationsInput) (*glacier.SetVaultNotificationsOutput, error)

	UploadArchive(*glacier.UploadArchiveInput) (*glacier.ArchiveCreationOutput, error)

	UploadMultipartPart(*glacier.UploadMultipartPartInput) (*glacier.UploadMultipartPartOutput, error)
}