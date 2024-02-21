// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package qconnectiface provides an interface to enable mocking the Amazon Q Connect service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package qconnectiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/qconnect"
)

// QConnectAPI provides an interface to enable mocking the
// qconnect.QConnect service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Q Connect.
//	func myFunc(svc qconnectiface.QConnectAPI) bool {
//	    // Make svc.CreateAssistant request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := qconnect.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockQConnectClient struct {
//	    qconnectiface.QConnectAPI
//	}
//	func (m *mockQConnectClient) CreateAssistant(input *qconnect.CreateAssistantInput) (*qconnect.CreateAssistantOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockQConnectClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type QConnectAPI interface {
	CreateAssistant(*qconnect.CreateAssistantInput) (*qconnect.CreateAssistantOutput, error)
	CreateAssistantWithContext(aws.Context, *qconnect.CreateAssistantInput, ...request.Option) (*qconnect.CreateAssistantOutput, error)
	CreateAssistantRequest(*qconnect.CreateAssistantInput) (*request.Request, *qconnect.CreateAssistantOutput)

	CreateAssistantAssociation(*qconnect.CreateAssistantAssociationInput) (*qconnect.CreateAssistantAssociationOutput, error)
	CreateAssistantAssociationWithContext(aws.Context, *qconnect.CreateAssistantAssociationInput, ...request.Option) (*qconnect.CreateAssistantAssociationOutput, error)
	CreateAssistantAssociationRequest(*qconnect.CreateAssistantAssociationInput) (*request.Request, *qconnect.CreateAssistantAssociationOutput)

	CreateContent(*qconnect.CreateContentInput) (*qconnect.CreateContentOutput, error)
	CreateContentWithContext(aws.Context, *qconnect.CreateContentInput, ...request.Option) (*qconnect.CreateContentOutput, error)
	CreateContentRequest(*qconnect.CreateContentInput) (*request.Request, *qconnect.CreateContentOutput)

	CreateKnowledgeBase(*qconnect.CreateKnowledgeBaseInput) (*qconnect.CreateKnowledgeBaseOutput, error)
	CreateKnowledgeBaseWithContext(aws.Context, *qconnect.CreateKnowledgeBaseInput, ...request.Option) (*qconnect.CreateKnowledgeBaseOutput, error)
	CreateKnowledgeBaseRequest(*qconnect.CreateKnowledgeBaseInput) (*request.Request, *qconnect.CreateKnowledgeBaseOutput)

	CreateQuickResponse(*qconnect.CreateQuickResponseInput) (*qconnect.CreateQuickResponseOutput, error)
	CreateQuickResponseWithContext(aws.Context, *qconnect.CreateQuickResponseInput, ...request.Option) (*qconnect.CreateQuickResponseOutput, error)
	CreateQuickResponseRequest(*qconnect.CreateQuickResponseInput) (*request.Request, *qconnect.CreateQuickResponseOutput)

	CreateSession(*qconnect.CreateSessionInput) (*qconnect.CreateSessionOutput, error)
	CreateSessionWithContext(aws.Context, *qconnect.CreateSessionInput, ...request.Option) (*qconnect.CreateSessionOutput, error)
	CreateSessionRequest(*qconnect.CreateSessionInput) (*request.Request, *qconnect.CreateSessionOutput)

	DeleteAssistant(*qconnect.DeleteAssistantInput) (*qconnect.DeleteAssistantOutput, error)
	DeleteAssistantWithContext(aws.Context, *qconnect.DeleteAssistantInput, ...request.Option) (*qconnect.DeleteAssistantOutput, error)
	DeleteAssistantRequest(*qconnect.DeleteAssistantInput) (*request.Request, *qconnect.DeleteAssistantOutput)

	DeleteAssistantAssociation(*qconnect.DeleteAssistantAssociationInput) (*qconnect.DeleteAssistantAssociationOutput, error)
	DeleteAssistantAssociationWithContext(aws.Context, *qconnect.DeleteAssistantAssociationInput, ...request.Option) (*qconnect.DeleteAssistantAssociationOutput, error)
	DeleteAssistantAssociationRequest(*qconnect.DeleteAssistantAssociationInput) (*request.Request, *qconnect.DeleteAssistantAssociationOutput)

	DeleteContent(*qconnect.DeleteContentInput) (*qconnect.DeleteContentOutput, error)
	DeleteContentWithContext(aws.Context, *qconnect.DeleteContentInput, ...request.Option) (*qconnect.DeleteContentOutput, error)
	DeleteContentRequest(*qconnect.DeleteContentInput) (*request.Request, *qconnect.DeleteContentOutput)

	DeleteImportJob(*qconnect.DeleteImportJobInput) (*qconnect.DeleteImportJobOutput, error)
	DeleteImportJobWithContext(aws.Context, *qconnect.DeleteImportJobInput, ...request.Option) (*qconnect.DeleteImportJobOutput, error)
	DeleteImportJobRequest(*qconnect.DeleteImportJobInput) (*request.Request, *qconnect.DeleteImportJobOutput)

	DeleteKnowledgeBase(*qconnect.DeleteKnowledgeBaseInput) (*qconnect.DeleteKnowledgeBaseOutput, error)
	DeleteKnowledgeBaseWithContext(aws.Context, *qconnect.DeleteKnowledgeBaseInput, ...request.Option) (*qconnect.DeleteKnowledgeBaseOutput, error)
	DeleteKnowledgeBaseRequest(*qconnect.DeleteKnowledgeBaseInput) (*request.Request, *qconnect.DeleteKnowledgeBaseOutput)

	DeleteQuickResponse(*qconnect.DeleteQuickResponseInput) (*qconnect.DeleteQuickResponseOutput, error)
	DeleteQuickResponseWithContext(aws.Context, *qconnect.DeleteQuickResponseInput, ...request.Option) (*qconnect.DeleteQuickResponseOutput, error)
	DeleteQuickResponseRequest(*qconnect.DeleteQuickResponseInput) (*request.Request, *qconnect.DeleteQuickResponseOutput)

	GetAssistant(*qconnect.GetAssistantInput) (*qconnect.GetAssistantOutput, error)
	GetAssistantWithContext(aws.Context, *qconnect.GetAssistantInput, ...request.Option) (*qconnect.GetAssistantOutput, error)
	GetAssistantRequest(*qconnect.GetAssistantInput) (*request.Request, *qconnect.GetAssistantOutput)

	GetAssistantAssociation(*qconnect.GetAssistantAssociationInput) (*qconnect.GetAssistantAssociationOutput, error)
	GetAssistantAssociationWithContext(aws.Context, *qconnect.GetAssistantAssociationInput, ...request.Option) (*qconnect.GetAssistantAssociationOutput, error)
	GetAssistantAssociationRequest(*qconnect.GetAssistantAssociationInput) (*request.Request, *qconnect.GetAssistantAssociationOutput)

	GetContent(*qconnect.GetContentInput) (*qconnect.GetContentOutput, error)
	GetContentWithContext(aws.Context, *qconnect.GetContentInput, ...request.Option) (*qconnect.GetContentOutput, error)
	GetContentRequest(*qconnect.GetContentInput) (*request.Request, *qconnect.GetContentOutput)

	GetContentSummary(*qconnect.GetContentSummaryInput) (*qconnect.GetContentSummaryOutput, error)
	GetContentSummaryWithContext(aws.Context, *qconnect.GetContentSummaryInput, ...request.Option) (*qconnect.GetContentSummaryOutput, error)
	GetContentSummaryRequest(*qconnect.GetContentSummaryInput) (*request.Request, *qconnect.GetContentSummaryOutput)

	GetImportJob(*qconnect.GetImportJobInput) (*qconnect.GetImportJobOutput, error)
	GetImportJobWithContext(aws.Context, *qconnect.GetImportJobInput, ...request.Option) (*qconnect.GetImportJobOutput, error)
	GetImportJobRequest(*qconnect.GetImportJobInput) (*request.Request, *qconnect.GetImportJobOutput)

	GetKnowledgeBase(*qconnect.GetKnowledgeBaseInput) (*qconnect.GetKnowledgeBaseOutput, error)
	GetKnowledgeBaseWithContext(aws.Context, *qconnect.GetKnowledgeBaseInput, ...request.Option) (*qconnect.GetKnowledgeBaseOutput, error)
	GetKnowledgeBaseRequest(*qconnect.GetKnowledgeBaseInput) (*request.Request, *qconnect.GetKnowledgeBaseOutput)

	GetQuickResponse(*qconnect.GetQuickResponseInput) (*qconnect.GetQuickResponseOutput, error)
	GetQuickResponseWithContext(aws.Context, *qconnect.GetQuickResponseInput, ...request.Option) (*qconnect.GetQuickResponseOutput, error)
	GetQuickResponseRequest(*qconnect.GetQuickResponseInput) (*request.Request, *qconnect.GetQuickResponseOutput)

	GetRecommendations(*qconnect.GetRecommendationsInput) (*qconnect.GetRecommendationsOutput, error)
	GetRecommendationsWithContext(aws.Context, *qconnect.GetRecommendationsInput, ...request.Option) (*qconnect.GetRecommendationsOutput, error)
	GetRecommendationsRequest(*qconnect.GetRecommendationsInput) (*request.Request, *qconnect.GetRecommendationsOutput)

	GetSession(*qconnect.GetSessionInput) (*qconnect.GetSessionOutput, error)
	GetSessionWithContext(aws.Context, *qconnect.GetSessionInput, ...request.Option) (*qconnect.GetSessionOutput, error)
	GetSessionRequest(*qconnect.GetSessionInput) (*request.Request, *qconnect.GetSessionOutput)

	ListAssistantAssociations(*qconnect.ListAssistantAssociationsInput) (*qconnect.ListAssistantAssociationsOutput, error)
	ListAssistantAssociationsWithContext(aws.Context, *qconnect.ListAssistantAssociationsInput, ...request.Option) (*qconnect.ListAssistantAssociationsOutput, error)
	ListAssistantAssociationsRequest(*qconnect.ListAssistantAssociationsInput) (*request.Request, *qconnect.ListAssistantAssociationsOutput)

	ListAssistantAssociationsPages(*qconnect.ListAssistantAssociationsInput, func(*qconnect.ListAssistantAssociationsOutput, bool) bool) error
	ListAssistantAssociationsPagesWithContext(aws.Context, *qconnect.ListAssistantAssociationsInput, func(*qconnect.ListAssistantAssociationsOutput, bool) bool, ...request.Option) error

	ListAssistants(*qconnect.ListAssistantsInput) (*qconnect.ListAssistantsOutput, error)
	ListAssistantsWithContext(aws.Context, *qconnect.ListAssistantsInput, ...request.Option) (*qconnect.ListAssistantsOutput, error)
	ListAssistantsRequest(*qconnect.ListAssistantsInput) (*request.Request, *qconnect.ListAssistantsOutput)

	ListAssistantsPages(*qconnect.ListAssistantsInput, func(*qconnect.ListAssistantsOutput, bool) bool) error
	ListAssistantsPagesWithContext(aws.Context, *qconnect.ListAssistantsInput, func(*qconnect.ListAssistantsOutput, bool) bool, ...request.Option) error

	ListContents(*qconnect.ListContentsInput) (*qconnect.ListContentsOutput, error)
	ListContentsWithContext(aws.Context, *qconnect.ListContentsInput, ...request.Option) (*qconnect.ListContentsOutput, error)
	ListContentsRequest(*qconnect.ListContentsInput) (*request.Request, *qconnect.ListContentsOutput)

	ListContentsPages(*qconnect.ListContentsInput, func(*qconnect.ListContentsOutput, bool) bool) error
	ListContentsPagesWithContext(aws.Context, *qconnect.ListContentsInput, func(*qconnect.ListContentsOutput, bool) bool, ...request.Option) error

	ListImportJobs(*qconnect.ListImportJobsInput) (*qconnect.ListImportJobsOutput, error)
	ListImportJobsWithContext(aws.Context, *qconnect.ListImportJobsInput, ...request.Option) (*qconnect.ListImportJobsOutput, error)
	ListImportJobsRequest(*qconnect.ListImportJobsInput) (*request.Request, *qconnect.ListImportJobsOutput)

	ListImportJobsPages(*qconnect.ListImportJobsInput, func(*qconnect.ListImportJobsOutput, bool) bool) error
	ListImportJobsPagesWithContext(aws.Context, *qconnect.ListImportJobsInput, func(*qconnect.ListImportJobsOutput, bool) bool, ...request.Option) error

	ListKnowledgeBases(*qconnect.ListKnowledgeBasesInput) (*qconnect.ListKnowledgeBasesOutput, error)
	ListKnowledgeBasesWithContext(aws.Context, *qconnect.ListKnowledgeBasesInput, ...request.Option) (*qconnect.ListKnowledgeBasesOutput, error)
	ListKnowledgeBasesRequest(*qconnect.ListKnowledgeBasesInput) (*request.Request, *qconnect.ListKnowledgeBasesOutput)

	ListKnowledgeBasesPages(*qconnect.ListKnowledgeBasesInput, func(*qconnect.ListKnowledgeBasesOutput, bool) bool) error
	ListKnowledgeBasesPagesWithContext(aws.Context, *qconnect.ListKnowledgeBasesInput, func(*qconnect.ListKnowledgeBasesOutput, bool) bool, ...request.Option) error

	ListQuickResponses(*qconnect.ListQuickResponsesInput) (*qconnect.ListQuickResponsesOutput, error)
	ListQuickResponsesWithContext(aws.Context, *qconnect.ListQuickResponsesInput, ...request.Option) (*qconnect.ListQuickResponsesOutput, error)
	ListQuickResponsesRequest(*qconnect.ListQuickResponsesInput) (*request.Request, *qconnect.ListQuickResponsesOutput)

	ListQuickResponsesPages(*qconnect.ListQuickResponsesInput, func(*qconnect.ListQuickResponsesOutput, bool) bool) error
	ListQuickResponsesPagesWithContext(aws.Context, *qconnect.ListQuickResponsesInput, func(*qconnect.ListQuickResponsesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*qconnect.ListTagsForResourceInput) (*qconnect.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *qconnect.ListTagsForResourceInput, ...request.Option) (*qconnect.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*qconnect.ListTagsForResourceInput) (*request.Request, *qconnect.ListTagsForResourceOutput)

	NotifyRecommendationsReceived(*qconnect.NotifyRecommendationsReceivedInput) (*qconnect.NotifyRecommendationsReceivedOutput, error)
	NotifyRecommendationsReceivedWithContext(aws.Context, *qconnect.NotifyRecommendationsReceivedInput, ...request.Option) (*qconnect.NotifyRecommendationsReceivedOutput, error)
	NotifyRecommendationsReceivedRequest(*qconnect.NotifyRecommendationsReceivedInput) (*request.Request, *qconnect.NotifyRecommendationsReceivedOutput)

	QueryAssistant(*qconnect.QueryAssistantInput) (*qconnect.QueryAssistantOutput, error)
	QueryAssistantWithContext(aws.Context, *qconnect.QueryAssistantInput, ...request.Option) (*qconnect.QueryAssistantOutput, error)
	QueryAssistantRequest(*qconnect.QueryAssistantInput) (*request.Request, *qconnect.QueryAssistantOutput)

	QueryAssistantPages(*qconnect.QueryAssistantInput, func(*qconnect.QueryAssistantOutput, bool) bool) error
	QueryAssistantPagesWithContext(aws.Context, *qconnect.QueryAssistantInput, func(*qconnect.QueryAssistantOutput, bool) bool, ...request.Option) error

	RemoveKnowledgeBaseTemplateUri(*qconnect.RemoveKnowledgeBaseTemplateUriInput) (*qconnect.RemoveKnowledgeBaseTemplateUriOutput, error)
	RemoveKnowledgeBaseTemplateUriWithContext(aws.Context, *qconnect.RemoveKnowledgeBaseTemplateUriInput, ...request.Option) (*qconnect.RemoveKnowledgeBaseTemplateUriOutput, error)
	RemoveKnowledgeBaseTemplateUriRequest(*qconnect.RemoveKnowledgeBaseTemplateUriInput) (*request.Request, *qconnect.RemoveKnowledgeBaseTemplateUriOutput)

	SearchContent(*qconnect.SearchContentInput) (*qconnect.SearchContentOutput, error)
	SearchContentWithContext(aws.Context, *qconnect.SearchContentInput, ...request.Option) (*qconnect.SearchContentOutput, error)
	SearchContentRequest(*qconnect.SearchContentInput) (*request.Request, *qconnect.SearchContentOutput)

	SearchContentPages(*qconnect.SearchContentInput, func(*qconnect.SearchContentOutput, bool) bool) error
	SearchContentPagesWithContext(aws.Context, *qconnect.SearchContentInput, func(*qconnect.SearchContentOutput, bool) bool, ...request.Option) error

	SearchQuickResponses(*qconnect.SearchQuickResponsesInput) (*qconnect.SearchQuickResponsesOutput, error)
	SearchQuickResponsesWithContext(aws.Context, *qconnect.SearchQuickResponsesInput, ...request.Option) (*qconnect.SearchQuickResponsesOutput, error)
	SearchQuickResponsesRequest(*qconnect.SearchQuickResponsesInput) (*request.Request, *qconnect.SearchQuickResponsesOutput)

	SearchQuickResponsesPages(*qconnect.SearchQuickResponsesInput, func(*qconnect.SearchQuickResponsesOutput, bool) bool) error
	SearchQuickResponsesPagesWithContext(aws.Context, *qconnect.SearchQuickResponsesInput, func(*qconnect.SearchQuickResponsesOutput, bool) bool, ...request.Option) error

	SearchSessions(*qconnect.SearchSessionsInput) (*qconnect.SearchSessionsOutput, error)
	SearchSessionsWithContext(aws.Context, *qconnect.SearchSessionsInput, ...request.Option) (*qconnect.SearchSessionsOutput, error)
	SearchSessionsRequest(*qconnect.SearchSessionsInput) (*request.Request, *qconnect.SearchSessionsOutput)

	SearchSessionsPages(*qconnect.SearchSessionsInput, func(*qconnect.SearchSessionsOutput, bool) bool) error
	SearchSessionsPagesWithContext(aws.Context, *qconnect.SearchSessionsInput, func(*qconnect.SearchSessionsOutput, bool) bool, ...request.Option) error

	StartContentUpload(*qconnect.StartContentUploadInput) (*qconnect.StartContentUploadOutput, error)
	StartContentUploadWithContext(aws.Context, *qconnect.StartContentUploadInput, ...request.Option) (*qconnect.StartContentUploadOutput, error)
	StartContentUploadRequest(*qconnect.StartContentUploadInput) (*request.Request, *qconnect.StartContentUploadOutput)

	StartImportJob(*qconnect.StartImportJobInput) (*qconnect.StartImportJobOutput, error)
	StartImportJobWithContext(aws.Context, *qconnect.StartImportJobInput, ...request.Option) (*qconnect.StartImportJobOutput, error)
	StartImportJobRequest(*qconnect.StartImportJobInput) (*request.Request, *qconnect.StartImportJobOutput)

	TagResource(*qconnect.TagResourceInput) (*qconnect.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *qconnect.TagResourceInput, ...request.Option) (*qconnect.TagResourceOutput, error)
	TagResourceRequest(*qconnect.TagResourceInput) (*request.Request, *qconnect.TagResourceOutput)

	UntagResource(*qconnect.UntagResourceInput) (*qconnect.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *qconnect.UntagResourceInput, ...request.Option) (*qconnect.UntagResourceOutput, error)
	UntagResourceRequest(*qconnect.UntagResourceInput) (*request.Request, *qconnect.UntagResourceOutput)

	UpdateContent(*qconnect.UpdateContentInput) (*qconnect.UpdateContentOutput, error)
	UpdateContentWithContext(aws.Context, *qconnect.UpdateContentInput, ...request.Option) (*qconnect.UpdateContentOutput, error)
	UpdateContentRequest(*qconnect.UpdateContentInput) (*request.Request, *qconnect.UpdateContentOutput)

	UpdateKnowledgeBaseTemplateUri(*qconnect.UpdateKnowledgeBaseTemplateUriInput) (*qconnect.UpdateKnowledgeBaseTemplateUriOutput, error)
	UpdateKnowledgeBaseTemplateUriWithContext(aws.Context, *qconnect.UpdateKnowledgeBaseTemplateUriInput, ...request.Option) (*qconnect.UpdateKnowledgeBaseTemplateUriOutput, error)
	UpdateKnowledgeBaseTemplateUriRequest(*qconnect.UpdateKnowledgeBaseTemplateUriInput) (*request.Request, *qconnect.UpdateKnowledgeBaseTemplateUriOutput)

	UpdateQuickResponse(*qconnect.UpdateQuickResponseInput) (*qconnect.UpdateQuickResponseOutput, error)
	UpdateQuickResponseWithContext(aws.Context, *qconnect.UpdateQuickResponseInput, ...request.Option) (*qconnect.UpdateQuickResponseOutput, error)
	UpdateQuickResponseRequest(*qconnect.UpdateQuickResponseInput) (*request.Request, *qconnect.UpdateQuickResponseOutput)
}

var _ QConnectAPI = (*qconnect.QConnect)(nil)