// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package sagemaker

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a SageMaker endpoint configuration resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/sagemaker_endpoint_configuration.html.markdown.
type EndpointConfiguration struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKeyArn pulumi.StringPtrOutput `pulumi:"kmsKeyArn"`
	// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Fields are documented below.
	ProductionVariants EndpointConfigurationProductionVariantArrayOutput `pulumi:"productionVariants"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewEndpointConfiguration registers a new resource with the given unique name, arguments, and options.
func NewEndpointConfiguration(ctx *pulumi.Context,
	name string, args *EndpointConfigurationArgs, opts ...pulumi.ResourceOption) (*EndpointConfiguration, error) {
	if args == nil || args.ProductionVariants == nil {
		return nil, errors.New("missing required argument 'ProductionVariants'")
	}
	if args == nil {
		args = &EndpointConfigurationArgs{}
	}
	var resource EndpointConfiguration
	err := ctx.RegisterResource("aws:sagemaker/endpointConfiguration:EndpointConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointConfiguration gets an existing EndpointConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointConfigurationState, opts ...pulumi.ResourceOption) (*EndpointConfiguration, error) {
	var resource EndpointConfiguration
	err := ctx.ReadResource("aws:sagemaker/endpointConfiguration:EndpointConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointConfiguration resources.
type endpointConfigurationState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
	Arn *string `pulumi:"arn"`
	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Fields are documented below.
	ProductionVariants []EndpointConfigurationProductionVariant `pulumi:"productionVariants"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type EndpointConfigurationState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
	Arn pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKeyArn pulumi.StringPtrInput
	// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Fields are documented below.
	ProductionVariants EndpointConfigurationProductionVariantArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (EndpointConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointConfigurationState)(nil)).Elem()
}

type endpointConfigurationArgs struct {
	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Fields are documented below.
	ProductionVariants []EndpointConfigurationProductionVariant `pulumi:"productionVariants"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a EndpointConfiguration resource.
type EndpointConfigurationArgs struct {
	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKeyArn pulumi.StringPtrInput
	// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Fields are documented below.
	ProductionVariants EndpointConfigurationProductionVariantArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (EndpointConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointConfigurationArgs)(nil)).Elem()
}

