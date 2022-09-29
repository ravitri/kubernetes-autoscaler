// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ecrpubliciface provides an interface to enable mocking the Amazon Elastic Container Registry Public service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ecrpubliciface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/ecrpublic"
)

// ECRPublicAPI provides an interface to enable mocking the
// ecrpublic.ECRPublic service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Elastic Container Registry Public.
//	func myFunc(svc ecrpubliciface.ECRPublicAPI) bool {
//	    // Make svc.BatchCheckLayerAvailability request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := ecrpublic.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockECRPublicClient struct {
//	    ecrpubliciface.ECRPublicAPI
//	}
//	func (m *mockECRPublicClient) BatchCheckLayerAvailability(input *ecrpublic.BatchCheckLayerAvailabilityInput) (*ecrpublic.BatchCheckLayerAvailabilityOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockECRPublicClient{}
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
type ECRPublicAPI interface {
	BatchCheckLayerAvailability(*ecrpublic.BatchCheckLayerAvailabilityInput) (*ecrpublic.BatchCheckLayerAvailabilityOutput, error)
	BatchCheckLayerAvailabilityWithContext(aws.Context, *ecrpublic.BatchCheckLayerAvailabilityInput, ...request.Option) (*ecrpublic.BatchCheckLayerAvailabilityOutput, error)
	BatchCheckLayerAvailabilityRequest(*ecrpublic.BatchCheckLayerAvailabilityInput) (*request.Request, *ecrpublic.BatchCheckLayerAvailabilityOutput)

	BatchDeleteImage(*ecrpublic.BatchDeleteImageInput) (*ecrpublic.BatchDeleteImageOutput, error)
	BatchDeleteImageWithContext(aws.Context, *ecrpublic.BatchDeleteImageInput, ...request.Option) (*ecrpublic.BatchDeleteImageOutput, error)
	BatchDeleteImageRequest(*ecrpublic.BatchDeleteImageInput) (*request.Request, *ecrpublic.BatchDeleteImageOutput)

	CompleteLayerUpload(*ecrpublic.CompleteLayerUploadInput) (*ecrpublic.CompleteLayerUploadOutput, error)
	CompleteLayerUploadWithContext(aws.Context, *ecrpublic.CompleteLayerUploadInput, ...request.Option) (*ecrpublic.CompleteLayerUploadOutput, error)
	CompleteLayerUploadRequest(*ecrpublic.CompleteLayerUploadInput) (*request.Request, *ecrpublic.CompleteLayerUploadOutput)

	CreateRepository(*ecrpublic.CreateRepositoryInput) (*ecrpublic.CreateRepositoryOutput, error)
	CreateRepositoryWithContext(aws.Context, *ecrpublic.CreateRepositoryInput, ...request.Option) (*ecrpublic.CreateRepositoryOutput, error)
	CreateRepositoryRequest(*ecrpublic.CreateRepositoryInput) (*request.Request, *ecrpublic.CreateRepositoryOutput)

	DeleteRepository(*ecrpublic.DeleteRepositoryInput) (*ecrpublic.DeleteRepositoryOutput, error)
	DeleteRepositoryWithContext(aws.Context, *ecrpublic.DeleteRepositoryInput, ...request.Option) (*ecrpublic.DeleteRepositoryOutput, error)
	DeleteRepositoryRequest(*ecrpublic.DeleteRepositoryInput) (*request.Request, *ecrpublic.DeleteRepositoryOutput)

	DeleteRepositoryPolicy(*ecrpublic.DeleteRepositoryPolicyInput) (*ecrpublic.DeleteRepositoryPolicyOutput, error)
	DeleteRepositoryPolicyWithContext(aws.Context, *ecrpublic.DeleteRepositoryPolicyInput, ...request.Option) (*ecrpublic.DeleteRepositoryPolicyOutput, error)
	DeleteRepositoryPolicyRequest(*ecrpublic.DeleteRepositoryPolicyInput) (*request.Request, *ecrpublic.DeleteRepositoryPolicyOutput)

	DescribeImageTags(*ecrpublic.DescribeImageTagsInput) (*ecrpublic.DescribeImageTagsOutput, error)
	DescribeImageTagsWithContext(aws.Context, *ecrpublic.DescribeImageTagsInput, ...request.Option) (*ecrpublic.DescribeImageTagsOutput, error)
	DescribeImageTagsRequest(*ecrpublic.DescribeImageTagsInput) (*request.Request, *ecrpublic.DescribeImageTagsOutput)

	DescribeImageTagsPages(*ecrpublic.DescribeImageTagsInput, func(*ecrpublic.DescribeImageTagsOutput, bool) bool) error
	DescribeImageTagsPagesWithContext(aws.Context, *ecrpublic.DescribeImageTagsInput, func(*ecrpublic.DescribeImageTagsOutput, bool) bool, ...request.Option) error

	DescribeImages(*ecrpublic.DescribeImagesInput) (*ecrpublic.DescribeImagesOutput, error)
	DescribeImagesWithContext(aws.Context, *ecrpublic.DescribeImagesInput, ...request.Option) (*ecrpublic.DescribeImagesOutput, error)
	DescribeImagesRequest(*ecrpublic.DescribeImagesInput) (*request.Request, *ecrpublic.DescribeImagesOutput)

	DescribeImagesPages(*ecrpublic.DescribeImagesInput, func(*ecrpublic.DescribeImagesOutput, bool) bool) error
	DescribeImagesPagesWithContext(aws.Context, *ecrpublic.DescribeImagesInput, func(*ecrpublic.DescribeImagesOutput, bool) bool, ...request.Option) error

	DescribeRegistries(*ecrpublic.DescribeRegistriesInput) (*ecrpublic.DescribeRegistriesOutput, error)
	DescribeRegistriesWithContext(aws.Context, *ecrpublic.DescribeRegistriesInput, ...request.Option) (*ecrpublic.DescribeRegistriesOutput, error)
	DescribeRegistriesRequest(*ecrpublic.DescribeRegistriesInput) (*request.Request, *ecrpublic.DescribeRegistriesOutput)

	DescribeRegistriesPages(*ecrpublic.DescribeRegistriesInput, func(*ecrpublic.DescribeRegistriesOutput, bool) bool) error
	DescribeRegistriesPagesWithContext(aws.Context, *ecrpublic.DescribeRegistriesInput, func(*ecrpublic.DescribeRegistriesOutput, bool) bool, ...request.Option) error

	DescribeRepositories(*ecrpublic.DescribeRepositoriesInput) (*ecrpublic.DescribeRepositoriesOutput, error)
	DescribeRepositoriesWithContext(aws.Context, *ecrpublic.DescribeRepositoriesInput, ...request.Option) (*ecrpublic.DescribeRepositoriesOutput, error)
	DescribeRepositoriesRequest(*ecrpublic.DescribeRepositoriesInput) (*request.Request, *ecrpublic.DescribeRepositoriesOutput)

	DescribeRepositoriesPages(*ecrpublic.DescribeRepositoriesInput, func(*ecrpublic.DescribeRepositoriesOutput, bool) bool) error
	DescribeRepositoriesPagesWithContext(aws.Context, *ecrpublic.DescribeRepositoriesInput, func(*ecrpublic.DescribeRepositoriesOutput, bool) bool, ...request.Option) error

	GetAuthorizationToken(*ecrpublic.GetAuthorizationTokenInput) (*ecrpublic.GetAuthorizationTokenOutput, error)
	GetAuthorizationTokenWithContext(aws.Context, *ecrpublic.GetAuthorizationTokenInput, ...request.Option) (*ecrpublic.GetAuthorizationTokenOutput, error)
	GetAuthorizationTokenRequest(*ecrpublic.GetAuthorizationTokenInput) (*request.Request, *ecrpublic.GetAuthorizationTokenOutput)

	GetRegistryCatalogData(*ecrpublic.GetRegistryCatalogDataInput) (*ecrpublic.GetRegistryCatalogDataOutput, error)
	GetRegistryCatalogDataWithContext(aws.Context, *ecrpublic.GetRegistryCatalogDataInput, ...request.Option) (*ecrpublic.GetRegistryCatalogDataOutput, error)
	GetRegistryCatalogDataRequest(*ecrpublic.GetRegistryCatalogDataInput) (*request.Request, *ecrpublic.GetRegistryCatalogDataOutput)

	GetRepositoryCatalogData(*ecrpublic.GetRepositoryCatalogDataInput) (*ecrpublic.GetRepositoryCatalogDataOutput, error)
	GetRepositoryCatalogDataWithContext(aws.Context, *ecrpublic.GetRepositoryCatalogDataInput, ...request.Option) (*ecrpublic.GetRepositoryCatalogDataOutput, error)
	GetRepositoryCatalogDataRequest(*ecrpublic.GetRepositoryCatalogDataInput) (*request.Request, *ecrpublic.GetRepositoryCatalogDataOutput)

	GetRepositoryPolicy(*ecrpublic.GetRepositoryPolicyInput) (*ecrpublic.GetRepositoryPolicyOutput, error)
	GetRepositoryPolicyWithContext(aws.Context, *ecrpublic.GetRepositoryPolicyInput, ...request.Option) (*ecrpublic.GetRepositoryPolicyOutput, error)
	GetRepositoryPolicyRequest(*ecrpublic.GetRepositoryPolicyInput) (*request.Request, *ecrpublic.GetRepositoryPolicyOutput)

	InitiateLayerUpload(*ecrpublic.InitiateLayerUploadInput) (*ecrpublic.InitiateLayerUploadOutput, error)
	InitiateLayerUploadWithContext(aws.Context, *ecrpublic.InitiateLayerUploadInput, ...request.Option) (*ecrpublic.InitiateLayerUploadOutput, error)
	InitiateLayerUploadRequest(*ecrpublic.InitiateLayerUploadInput) (*request.Request, *ecrpublic.InitiateLayerUploadOutput)

	ListTagsForResource(*ecrpublic.ListTagsForResourceInput) (*ecrpublic.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *ecrpublic.ListTagsForResourceInput, ...request.Option) (*ecrpublic.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*ecrpublic.ListTagsForResourceInput) (*request.Request, *ecrpublic.ListTagsForResourceOutput)

	PutImage(*ecrpublic.PutImageInput) (*ecrpublic.PutImageOutput, error)
	PutImageWithContext(aws.Context, *ecrpublic.PutImageInput, ...request.Option) (*ecrpublic.PutImageOutput, error)
	PutImageRequest(*ecrpublic.PutImageInput) (*request.Request, *ecrpublic.PutImageOutput)

	PutRegistryCatalogData(*ecrpublic.PutRegistryCatalogDataInput) (*ecrpublic.PutRegistryCatalogDataOutput, error)
	PutRegistryCatalogDataWithContext(aws.Context, *ecrpublic.PutRegistryCatalogDataInput, ...request.Option) (*ecrpublic.PutRegistryCatalogDataOutput, error)
	PutRegistryCatalogDataRequest(*ecrpublic.PutRegistryCatalogDataInput) (*request.Request, *ecrpublic.PutRegistryCatalogDataOutput)

	PutRepositoryCatalogData(*ecrpublic.PutRepositoryCatalogDataInput) (*ecrpublic.PutRepositoryCatalogDataOutput, error)
	PutRepositoryCatalogDataWithContext(aws.Context, *ecrpublic.PutRepositoryCatalogDataInput, ...request.Option) (*ecrpublic.PutRepositoryCatalogDataOutput, error)
	PutRepositoryCatalogDataRequest(*ecrpublic.PutRepositoryCatalogDataInput) (*request.Request, *ecrpublic.PutRepositoryCatalogDataOutput)

	SetRepositoryPolicy(*ecrpublic.SetRepositoryPolicyInput) (*ecrpublic.SetRepositoryPolicyOutput, error)
	SetRepositoryPolicyWithContext(aws.Context, *ecrpublic.SetRepositoryPolicyInput, ...request.Option) (*ecrpublic.SetRepositoryPolicyOutput, error)
	SetRepositoryPolicyRequest(*ecrpublic.SetRepositoryPolicyInput) (*request.Request, *ecrpublic.SetRepositoryPolicyOutput)

	TagResource(*ecrpublic.TagResourceInput) (*ecrpublic.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *ecrpublic.TagResourceInput, ...request.Option) (*ecrpublic.TagResourceOutput, error)
	TagResourceRequest(*ecrpublic.TagResourceInput) (*request.Request, *ecrpublic.TagResourceOutput)

	UntagResource(*ecrpublic.UntagResourceInput) (*ecrpublic.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *ecrpublic.UntagResourceInput, ...request.Option) (*ecrpublic.UntagResourceOutput, error)
	UntagResourceRequest(*ecrpublic.UntagResourceInput) (*request.Request, *ecrpublic.UntagResourceOutput)

	UploadLayerPart(*ecrpublic.UploadLayerPartInput) (*ecrpublic.UploadLayerPartOutput, error)
	UploadLayerPartWithContext(aws.Context, *ecrpublic.UploadLayerPartInput, ...request.Option) (*ecrpublic.UploadLayerPartOutput, error)
	UploadLayerPartRequest(*ecrpublic.UploadLayerPartInput) (*request.Request, *ecrpublic.UploadLayerPartOutput)
}

var _ ECRPublicAPI = (*ecrpublic.ECRPublic)(nil)
