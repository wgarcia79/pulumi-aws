// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package appsync

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AppSync Function.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appsync_function.html.markdown.
type Function struct {
	pulumi.CustomResourceState

	// The ID of the associated AppSync API.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// The ARN of the Function object.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Function DataSource name.
	DataSource pulumi.StringOutput `pulumi:"dataSource"`
	// The Function description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A unique ID representing the Function object.
	FunctionId pulumi.StringOutput `pulumi:"functionId"`
	// The version of the request mapping template. Currently the supported value is `2018-05-29`.
	FunctionVersion pulumi.StringPtrOutput `pulumi:"functionVersion"`
	// The Function name. The function name does not have to be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate pulumi.StringOutput `pulumi:"requestMappingTemplate"`
	// The Function response mapping template.
	ResponseMappingTemplate pulumi.StringOutput `pulumi:"responseMappingTemplate"`
}

// NewFunction registers a new resource with the given unique name, arguments, and options.
func NewFunction(ctx *pulumi.Context,
	name string, args *FunctionArgs, opts ...pulumi.ResourceOption) (*Function, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
	}
	if args == nil || args.DataSource == nil {
		return nil, errors.New("missing required argument 'DataSource'")
	}
	if args == nil || args.RequestMappingTemplate == nil {
		return nil, errors.New("missing required argument 'RequestMappingTemplate'")
	}
	if args == nil || args.ResponseMappingTemplate == nil {
		return nil, errors.New("missing required argument 'ResponseMappingTemplate'")
	}
	if args == nil {
		args = &FunctionArgs{}
	}
	var resource Function
	err := ctx.RegisterResource("aws:appsync/function:Function", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunction gets an existing Function resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionState, opts ...pulumi.ResourceOption) (*Function, error) {
	var resource Function
	err := ctx.ReadResource("aws:appsync/function:Function", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Function resources.
type functionState struct {
	// The ID of the associated AppSync API.
	ApiId *string `pulumi:"apiId"`
	// The ARN of the Function object.
	Arn *string `pulumi:"arn"`
	// The Function DataSource name.
	DataSource *string `pulumi:"dataSource"`
	// The Function description.
	Description *string `pulumi:"description"`
	// A unique ID representing the Function object.
	FunctionId *string `pulumi:"functionId"`
	// The version of the request mapping template. Currently the supported value is `2018-05-29`.
	FunctionVersion *string `pulumi:"functionVersion"`
	// The Function name. The function name does not have to be unique.
	Name *string `pulumi:"name"`
	// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate *string `pulumi:"requestMappingTemplate"`
	// The Function response mapping template.
	ResponseMappingTemplate *string `pulumi:"responseMappingTemplate"`
}

type FunctionState struct {
	// The ID of the associated AppSync API.
	ApiId pulumi.StringPtrInput
	// The ARN of the Function object.
	Arn pulumi.StringPtrInput
	// The Function DataSource name.
	DataSource pulumi.StringPtrInput
	// The Function description.
	Description pulumi.StringPtrInput
	// A unique ID representing the Function object.
	FunctionId pulumi.StringPtrInput
	// The version of the request mapping template. Currently the supported value is `2018-05-29`.
	FunctionVersion pulumi.StringPtrInput
	// The Function name. The function name does not have to be unique.
	Name pulumi.StringPtrInput
	// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate pulumi.StringPtrInput
	// The Function response mapping template.
	ResponseMappingTemplate pulumi.StringPtrInput
}

func (FunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionState)(nil)).Elem()
}

type functionArgs struct {
	// The ID of the associated AppSync API.
	ApiId string `pulumi:"apiId"`
	// The Function DataSource name.
	DataSource string `pulumi:"dataSource"`
	// The Function description.
	Description *string `pulumi:"description"`
	// The version of the request mapping template. Currently the supported value is `2018-05-29`.
	FunctionVersion *string `pulumi:"functionVersion"`
	// The Function name. The function name does not have to be unique.
	Name *string `pulumi:"name"`
	// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate string `pulumi:"requestMappingTemplate"`
	// The Function response mapping template.
	ResponseMappingTemplate string `pulumi:"responseMappingTemplate"`
}

// The set of arguments for constructing a Function resource.
type FunctionArgs struct {
	// The ID of the associated AppSync API.
	ApiId pulumi.StringInput
	// The Function DataSource name.
	DataSource pulumi.StringInput
	// The Function description.
	Description pulumi.StringPtrInput
	// The version of the request mapping template. Currently the supported value is `2018-05-29`.
	FunctionVersion pulumi.StringPtrInput
	// The Function name. The function name does not have to be unique.
	Name pulumi.StringPtrInput
	// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate pulumi.StringInput
	// The Function response mapping template.
	ResponseMappingTemplate pulumi.StringInput
}

func (FunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionArgs)(nil)).Elem()
}

