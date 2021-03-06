package ssmiface

import (
	"github.com/datacratic/aws-sdk-go/service/ssm"
)

type SSMAPI interface {
	CreateAssociation(*ssm.CreateAssociationInput) (*ssm.CreateAssociationOutput, error)

	CreateAssociationBatch(*ssm.CreateAssociationBatchInput) (*ssm.CreateAssociationBatchOutput, error)

	CreateDocument(*ssm.CreateDocumentInput) (*ssm.CreateDocumentOutput, error)

	DeleteAssociation(*ssm.DeleteAssociationInput) (*ssm.DeleteAssociationOutput, error)

	DeleteDocument(*ssm.DeleteDocumentInput) (*ssm.DeleteDocumentOutput, error)

	DescribeAssociation(*ssm.DescribeAssociationInput) (*ssm.DescribeAssociationOutput, error)

	DescribeDocument(*ssm.DescribeDocumentInput) (*ssm.DescribeDocumentOutput, error)

	GetDocument(*ssm.GetDocumentInput) (*ssm.GetDocumentOutput, error)

	ListAssociations(*ssm.ListAssociationsInput) (*ssm.ListAssociationsOutput, error)

	ListDocuments(*ssm.ListDocumentsInput) (*ssm.ListDocumentsOutput, error)

	UpdateAssociationStatus(*ssm.UpdateAssociationStatusInput) (*ssm.UpdateAssociationStatusOutput, error)
}