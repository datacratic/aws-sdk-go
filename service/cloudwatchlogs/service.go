package cloudwatchlogs

import (
	"github.com/datacratic/aws-sdk-go/aws"
	"github.com/datacratic/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/datacratic/aws-sdk-go/internal/signer/v4"
)

// CloudWatchLogs is a client for Amazon CloudWatch Logs.
type CloudWatchLogs struct {
	*aws.Service
}

// Used for custom service initialization logic
var initService func(*aws.Service)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// New returns a new CloudWatchLogs client.
func New(config *aws.Config) *CloudWatchLogs {
	if config == nil {
		config = &aws.Config{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config),
		ServiceName:  "logs",
		APIVersion:   "2014-03-28",
		JSONVersion:  "1.1",
		TargetPrefix: "Logs_20140328",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &CloudWatchLogs{service}
}

// newRequest creates a new request for a CloudWatchLogs operation and runs any
// custom request initialization.
func (c *CloudWatchLogs) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := aws.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
