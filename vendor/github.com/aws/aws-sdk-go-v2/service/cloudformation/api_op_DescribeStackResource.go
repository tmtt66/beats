// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for DescribeStackResource action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeStackResourceInput
type DescribeStackResourceInput struct {
	_ struct{} `type:"structure"`

	// The logical name of the resource as specified in the template.
	//
	// Default: There is no default value.
	//
	// LogicalResourceId is a required field
	LogicalResourceId *string `type:"string" required:"true"`

	// The name or the unique stack ID that is associated with the stack, which
	// are not always interchangeable:
	//
	//    * Running stacks: You can specify either the stack's name or its unique
	//    stack ID.
	//
	//    * Deleted stacks: You must specify the unique stack ID.
	//
	// Default: There is no default value.
	//
	// StackName is a required field
	StackName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeStackResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStackResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStackResourceInput"}

	if s.LogicalResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogicalResourceId"))
	}

	if s.StackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output for a DescribeStackResource action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeStackResourceOutput
type DescribeStackResourceOutput struct {
	_ struct{} `type:"structure"`

	// A StackResourceDetail structure containing the description of the specified
	// resource in the specified stack.
	StackResourceDetail *StackResourceDetail `type:"structure"`
}

// String returns the string representation
func (s DescribeStackResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeStackResource = "DescribeStackResource"

// DescribeStackResourceRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns a description of the specified resource in the specified stack.
//
// For deleted stacks, DescribeStackResource returns resource information for
// up to 90 days after the stack has been deleted.
//
//    // Example sending a request using DescribeStackResourceRequest.
//    req := client.DescribeStackResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeStackResource
func (c *Client) DescribeStackResourceRequest(input *DescribeStackResourceInput) DescribeStackResourceRequest {
	op := &aws.Operation{
		Name:       opDescribeStackResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeStackResourceInput{}
	}

	req := c.newRequest(op, input, &DescribeStackResourceOutput{})
	return DescribeStackResourceRequest{Request: req, Input: input, Copy: c.DescribeStackResourceRequest}
}

// DescribeStackResourceRequest is the request type for the
// DescribeStackResource API operation.
type DescribeStackResourceRequest struct {
	*aws.Request
	Input *DescribeStackResourceInput
	Copy  func(*DescribeStackResourceInput) DescribeStackResourceRequest
}

// Send marshals and sends the DescribeStackResource API request.
func (r DescribeStackResourceRequest) Send(ctx context.Context) (*DescribeStackResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStackResourceResponse{
		DescribeStackResourceOutput: r.Request.Data.(*DescribeStackResourceOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStackResourceResponse is the response type for the
// DescribeStackResource API operation.
type DescribeStackResourceResponse struct {
	*DescribeStackResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStackResource request.
func (r *DescribeStackResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
