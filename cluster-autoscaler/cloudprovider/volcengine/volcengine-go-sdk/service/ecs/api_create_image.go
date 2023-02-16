// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/response"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateImageCommon = "CreateImage"

// CreateImageCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateImageCommon operation. The "output" return
// value will be populated with the CreateImageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateImageCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateImageCommon Send returns without error.
//
// See CreateImageCommon for more information on using the CreateImageCommon
// API call, and error handling.
//
//	// Example sending a request using the CreateImageCommonRequest method.
//	req, resp := client.CreateImageCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) CreateImageCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateImageCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// CreateImageCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation CreateImageCommon for usage and error information.
func (c *ECS) CreateImageCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateImageCommonRequest(input)
	return out, req.Send()
}

// CreateImageCommonWithContext is the same as CreateImageCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateImageCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) CreateImageCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateImageCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateImage = "CreateImage"

// CreateImageRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateImage operation. The "output" return
// value will be populated with the CreateImageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateImageCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateImageCommon Send returns without error.
//
// See CreateImage for more information on using the CreateImage
// API call, and error handling.
//
//	// Example sending a request using the CreateImageRequest method.
//	req, resp := client.CreateImageRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) CreateImageRequest(input *CreateImageInput) (req *request.Request, output *CreateImageOutput) {
	op := &request.Operation{
		Name:       opCreateImage,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateImageInput{}
	}

	output = &CreateImageOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateImage API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation CreateImage for usage and error information.
func (c *ECS) CreateImage(input *CreateImageInput) (*CreateImageOutput, error) {
	req, out := c.CreateImageRequest(input)
	return out, req.Send()
}

// CreateImageWithContext is the same as CreateImage with the addition of
// the ability to pass a context and additional request options.
//
// See CreateImage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) CreateImageWithContext(ctx volcengine.Context, input *CreateImageInput, opts ...request.Option) (*CreateImageOutput, error) {
	req, out := c.CreateImageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateImageInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	ImageName *string `type:"string"`

	InstanceId *string `type:"string"`

	ProjectName *string `type:"string"`
}

// String returns the string representation
func (s CreateImageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateImageInput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *CreateImageInput) SetDescription(v string) *CreateImageInput {
	s.Description = &v
	return s
}

// SetImageName sets the ImageName field's value.
func (s *CreateImageInput) SetImageName(v string) *CreateImageInput {
	s.ImageName = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateImageInput) SetInstanceId(v string) *CreateImageInput {
	s.InstanceId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateImageInput) SetProjectName(v string) *CreateImageInput {
	s.ProjectName = &v
	return s
}

type CreateImageOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ImageId *string `type:"string"`
}

// String returns the string representation
func (s CreateImageOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateImageOutput) GoString() string {
	return s.String()
}

// SetImageId sets the ImageId field's value.
func (s *CreateImageOutput) SetImageId(v string) *CreateImageOutput {
	s.ImageId = &v
	return s
}
