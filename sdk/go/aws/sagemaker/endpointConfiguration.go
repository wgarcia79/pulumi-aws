// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a SageMaker endpoint configuration resource.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sagemaker"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sagemaker.NewEndpointConfiguration(ctx, "ec", &sagemaker.EndpointConfigurationArgs{
// 			ProductionVariants: sagemaker.EndpointConfigurationProductionVariantArray{
// 				&sagemaker.EndpointConfigurationProductionVariantArgs{
// 					VariantName:          pulumi.String("variant-1"),
// 					ModelName:            pulumi.Any(aws_sagemaker_model.M.Name),
// 					InitialInstanceCount: pulumi.Int(1),
// 					InstanceType:         pulumi.String("ml.t2.medium"),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("foo"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Endpoint configurations can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:sagemaker/endpointConfiguration:EndpointConfiguration test_endpoint_config endpoint-config-foo
// ```
type EndpointConfiguration struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies configuration for how an endpoint performs asynchronous inference.
	AsyncInferenceConfig EndpointConfigurationAsyncInferenceConfigPtrOutput `pulumi:"asyncInferenceConfig"`
	// Specifies the parameters to capture input/output of Sagemaker models endpoints. Fields are documented below.
	DataCaptureConfig EndpointConfigurationDataCaptureConfigPtrOutput `pulumi:"dataCaptureConfig"`
	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKeyArn pulumi.StringPtrOutput `pulumi:"kmsKeyArn"`
	// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Fields are documented below.
	ProductionVariants EndpointConfigurationProductionVariantArrayOutput `pulumi:"productionVariants"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewEndpointConfiguration registers a new resource with the given unique name, arguments, and options.
func NewEndpointConfiguration(ctx *pulumi.Context,
	name string, args *EndpointConfigurationArgs, opts ...pulumi.ResourceOption) (*EndpointConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductionVariants == nil {
		return nil, errors.New("invalid value for required argument 'ProductionVariants'")
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
	// Specifies configuration for how an endpoint performs asynchronous inference.
	AsyncInferenceConfig *EndpointConfigurationAsyncInferenceConfig `pulumi:"asyncInferenceConfig"`
	// Specifies the parameters to capture input/output of Sagemaker models endpoints. Fields are documented below.
	DataCaptureConfig *EndpointConfigurationDataCaptureConfig `pulumi:"dataCaptureConfig"`
	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Fields are documented below.
	ProductionVariants []EndpointConfigurationProductionVariant `pulumi:"productionVariants"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type EndpointConfigurationState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
	Arn pulumi.StringPtrInput
	// Specifies configuration for how an endpoint performs asynchronous inference.
	AsyncInferenceConfig EndpointConfigurationAsyncInferenceConfigPtrInput
	// Specifies the parameters to capture input/output of Sagemaker models endpoints. Fields are documented below.
	DataCaptureConfig EndpointConfigurationDataCaptureConfigPtrInput
	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKeyArn pulumi.StringPtrInput
	// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Fields are documented below.
	ProductionVariants EndpointConfigurationProductionVariantArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (EndpointConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointConfigurationState)(nil)).Elem()
}

type endpointConfigurationArgs struct {
	// Specifies configuration for how an endpoint performs asynchronous inference.
	AsyncInferenceConfig *EndpointConfigurationAsyncInferenceConfig `pulumi:"asyncInferenceConfig"`
	// Specifies the parameters to capture input/output of Sagemaker models endpoints. Fields are documented below.
	DataCaptureConfig *EndpointConfigurationDataCaptureConfig `pulumi:"dataCaptureConfig"`
	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Fields are documented below.
	ProductionVariants []EndpointConfigurationProductionVariant `pulumi:"productionVariants"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EndpointConfiguration resource.
type EndpointConfigurationArgs struct {
	// Specifies configuration for how an endpoint performs asynchronous inference.
	AsyncInferenceConfig EndpointConfigurationAsyncInferenceConfigPtrInput
	// Specifies the parameters to capture input/output of Sagemaker models endpoints. Fields are documented below.
	DataCaptureConfig EndpointConfigurationDataCaptureConfigPtrInput
	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKeyArn pulumi.StringPtrInput
	// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Fields are documented below.
	ProductionVariants EndpointConfigurationProductionVariantArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (EndpointConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointConfigurationArgs)(nil)).Elem()
}

type EndpointConfigurationInput interface {
	pulumi.Input

	ToEndpointConfigurationOutput() EndpointConfigurationOutput
	ToEndpointConfigurationOutputWithContext(ctx context.Context) EndpointConfigurationOutput
}

func (*EndpointConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointConfiguration)(nil))
}

func (i *EndpointConfiguration) ToEndpointConfigurationOutput() EndpointConfigurationOutput {
	return i.ToEndpointConfigurationOutputWithContext(context.Background())
}

func (i *EndpointConfiguration) ToEndpointConfigurationOutputWithContext(ctx context.Context) EndpointConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointConfigurationOutput)
}

func (i *EndpointConfiguration) ToEndpointConfigurationPtrOutput() EndpointConfigurationPtrOutput {
	return i.ToEndpointConfigurationPtrOutputWithContext(context.Background())
}

func (i *EndpointConfiguration) ToEndpointConfigurationPtrOutputWithContext(ctx context.Context) EndpointConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointConfigurationPtrOutput)
}

type EndpointConfigurationPtrInput interface {
	pulumi.Input

	ToEndpointConfigurationPtrOutput() EndpointConfigurationPtrOutput
	ToEndpointConfigurationPtrOutputWithContext(ctx context.Context) EndpointConfigurationPtrOutput
}

type endpointConfigurationPtrType EndpointConfigurationArgs

func (*endpointConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointConfiguration)(nil))
}

func (i *endpointConfigurationPtrType) ToEndpointConfigurationPtrOutput() EndpointConfigurationPtrOutput {
	return i.ToEndpointConfigurationPtrOutputWithContext(context.Background())
}

func (i *endpointConfigurationPtrType) ToEndpointConfigurationPtrOutputWithContext(ctx context.Context) EndpointConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointConfigurationPtrOutput)
}

// EndpointConfigurationArrayInput is an input type that accepts EndpointConfigurationArray and EndpointConfigurationArrayOutput values.
// You can construct a concrete instance of `EndpointConfigurationArrayInput` via:
//
//          EndpointConfigurationArray{ EndpointConfigurationArgs{...} }
type EndpointConfigurationArrayInput interface {
	pulumi.Input

	ToEndpointConfigurationArrayOutput() EndpointConfigurationArrayOutput
	ToEndpointConfigurationArrayOutputWithContext(context.Context) EndpointConfigurationArrayOutput
}

type EndpointConfigurationArray []EndpointConfigurationInput

func (EndpointConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointConfiguration)(nil)).Elem()
}

func (i EndpointConfigurationArray) ToEndpointConfigurationArrayOutput() EndpointConfigurationArrayOutput {
	return i.ToEndpointConfigurationArrayOutputWithContext(context.Background())
}

func (i EndpointConfigurationArray) ToEndpointConfigurationArrayOutputWithContext(ctx context.Context) EndpointConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointConfigurationArrayOutput)
}

// EndpointConfigurationMapInput is an input type that accepts EndpointConfigurationMap and EndpointConfigurationMapOutput values.
// You can construct a concrete instance of `EndpointConfigurationMapInput` via:
//
//          EndpointConfigurationMap{ "key": EndpointConfigurationArgs{...} }
type EndpointConfigurationMapInput interface {
	pulumi.Input

	ToEndpointConfigurationMapOutput() EndpointConfigurationMapOutput
	ToEndpointConfigurationMapOutputWithContext(context.Context) EndpointConfigurationMapOutput
}

type EndpointConfigurationMap map[string]EndpointConfigurationInput

func (EndpointConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointConfiguration)(nil)).Elem()
}

func (i EndpointConfigurationMap) ToEndpointConfigurationMapOutput() EndpointConfigurationMapOutput {
	return i.ToEndpointConfigurationMapOutputWithContext(context.Background())
}

func (i EndpointConfigurationMap) ToEndpointConfigurationMapOutputWithContext(ctx context.Context) EndpointConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointConfigurationMapOutput)
}

type EndpointConfigurationOutput struct{ *pulumi.OutputState }

func (EndpointConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointConfiguration)(nil))
}

func (o EndpointConfigurationOutput) ToEndpointConfigurationOutput() EndpointConfigurationOutput {
	return o
}

func (o EndpointConfigurationOutput) ToEndpointConfigurationOutputWithContext(ctx context.Context) EndpointConfigurationOutput {
	return o
}

func (o EndpointConfigurationOutput) ToEndpointConfigurationPtrOutput() EndpointConfigurationPtrOutput {
	return o.ToEndpointConfigurationPtrOutputWithContext(context.Background())
}

func (o EndpointConfigurationOutput) ToEndpointConfigurationPtrOutputWithContext(ctx context.Context) EndpointConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointConfiguration) *EndpointConfiguration {
		return &v
	}).(EndpointConfigurationPtrOutput)
}

type EndpointConfigurationPtrOutput struct{ *pulumi.OutputState }

func (EndpointConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointConfiguration)(nil))
}

func (o EndpointConfigurationPtrOutput) ToEndpointConfigurationPtrOutput() EndpointConfigurationPtrOutput {
	return o
}

func (o EndpointConfigurationPtrOutput) ToEndpointConfigurationPtrOutputWithContext(ctx context.Context) EndpointConfigurationPtrOutput {
	return o
}

func (o EndpointConfigurationPtrOutput) Elem() EndpointConfigurationOutput {
	return o.ApplyT(func(v *EndpointConfiguration) EndpointConfiguration {
		if v != nil {
			return *v
		}
		var ret EndpointConfiguration
		return ret
	}).(EndpointConfigurationOutput)
}

type EndpointConfigurationArrayOutput struct{ *pulumi.OutputState }

func (EndpointConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointConfiguration)(nil))
}

func (o EndpointConfigurationArrayOutput) ToEndpointConfigurationArrayOutput() EndpointConfigurationArrayOutput {
	return o
}

func (o EndpointConfigurationArrayOutput) ToEndpointConfigurationArrayOutputWithContext(ctx context.Context) EndpointConfigurationArrayOutput {
	return o
}

func (o EndpointConfigurationArrayOutput) Index(i pulumi.IntInput) EndpointConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointConfiguration {
		return vs[0].([]EndpointConfiguration)[vs[1].(int)]
	}).(EndpointConfigurationOutput)
}

type EndpointConfigurationMapOutput struct{ *pulumi.OutputState }

func (EndpointConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EndpointConfiguration)(nil))
}

func (o EndpointConfigurationMapOutput) ToEndpointConfigurationMapOutput() EndpointConfigurationMapOutput {
	return o
}

func (o EndpointConfigurationMapOutput) ToEndpointConfigurationMapOutputWithContext(ctx context.Context) EndpointConfigurationMapOutput {
	return o
}

func (o EndpointConfigurationMapOutput) MapIndex(k pulumi.StringInput) EndpointConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EndpointConfiguration {
		return vs[0].(map[string]EndpointConfiguration)[vs[1].(string)]
	}).(EndpointConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointConfigurationInput)(nil)).Elem(), &EndpointConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointConfigurationPtrInput)(nil)).Elem(), &EndpointConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointConfigurationArrayInput)(nil)).Elem(), EndpointConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointConfigurationMapInput)(nil)).Elem(), EndpointConfigurationMap{})
	pulumi.RegisterOutputType(EndpointConfigurationOutput{})
	pulumi.RegisterOutputType(EndpointConfigurationPtrOutput{})
	pulumi.RegisterOutputType(EndpointConfigurationArrayOutput{})
	pulumi.RegisterOutputType(EndpointConfigurationMapOutput{})
}
