// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a lifecycle configuration for SageMaker Notebook Instances.
//
// ## Import
//
// Models can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:sagemaker/notebookInstanceLifecycleConfiguration:NotebookInstanceLifecycleConfiguration lc foo
// ```
type NotebookInstanceLifecycleConfiguration struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) assigned by AWS to this lifecycle configuration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
	OnCreate pulumi.StringPtrOutput `pulumi:"onCreate"`
	// A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
	OnStart pulumi.StringPtrOutput `pulumi:"onStart"`
}

// NewNotebookInstanceLifecycleConfiguration registers a new resource with the given unique name, arguments, and options.
func NewNotebookInstanceLifecycleConfiguration(ctx *pulumi.Context,
	name string, args *NotebookInstanceLifecycleConfigurationArgs, opts ...pulumi.ResourceOption) (*NotebookInstanceLifecycleConfiguration, error) {
	if args == nil {
		args = &NotebookInstanceLifecycleConfigurationArgs{}
	}

	var resource NotebookInstanceLifecycleConfiguration
	err := ctx.RegisterResource("aws:sagemaker/notebookInstanceLifecycleConfiguration:NotebookInstanceLifecycleConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotebookInstanceLifecycleConfiguration gets an existing NotebookInstanceLifecycleConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotebookInstanceLifecycleConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebookInstanceLifecycleConfigurationState, opts ...pulumi.ResourceOption) (*NotebookInstanceLifecycleConfiguration, error) {
	var resource NotebookInstanceLifecycleConfiguration
	err := ctx.ReadResource("aws:sagemaker/notebookInstanceLifecycleConfiguration:NotebookInstanceLifecycleConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotebookInstanceLifecycleConfiguration resources.
type notebookInstanceLifecycleConfigurationState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this lifecycle configuration.
	Arn *string `pulumi:"arn"`
	// The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
	OnCreate *string `pulumi:"onCreate"`
	// A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
	OnStart *string `pulumi:"onStart"`
}

type NotebookInstanceLifecycleConfigurationState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this lifecycle configuration.
	Arn pulumi.StringPtrInput
	// The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
	OnCreate pulumi.StringPtrInput
	// A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
	OnStart pulumi.StringPtrInput
}

func (NotebookInstanceLifecycleConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookInstanceLifecycleConfigurationState)(nil)).Elem()
}

type notebookInstanceLifecycleConfigurationArgs struct {
	// The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
	OnCreate *string `pulumi:"onCreate"`
	// A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
	OnStart *string `pulumi:"onStart"`
}

// The set of arguments for constructing a NotebookInstanceLifecycleConfiguration resource.
type NotebookInstanceLifecycleConfigurationArgs struct {
	// The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
	OnCreate pulumi.StringPtrInput
	// A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
	OnStart pulumi.StringPtrInput
}

func (NotebookInstanceLifecycleConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookInstanceLifecycleConfigurationArgs)(nil)).Elem()
}

type NotebookInstanceLifecycleConfigurationInput interface {
	pulumi.Input

	ToNotebookInstanceLifecycleConfigurationOutput() NotebookInstanceLifecycleConfigurationOutput
	ToNotebookInstanceLifecycleConfigurationOutputWithContext(ctx context.Context) NotebookInstanceLifecycleConfigurationOutput
}

func (*NotebookInstanceLifecycleConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookInstanceLifecycleConfiguration)(nil))
}

func (i *NotebookInstanceLifecycleConfiguration) ToNotebookInstanceLifecycleConfigurationOutput() NotebookInstanceLifecycleConfigurationOutput {
	return i.ToNotebookInstanceLifecycleConfigurationOutputWithContext(context.Background())
}

func (i *NotebookInstanceLifecycleConfiguration) ToNotebookInstanceLifecycleConfigurationOutputWithContext(ctx context.Context) NotebookInstanceLifecycleConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookInstanceLifecycleConfigurationOutput)
}

func (i *NotebookInstanceLifecycleConfiguration) ToNotebookInstanceLifecycleConfigurationPtrOutput() NotebookInstanceLifecycleConfigurationPtrOutput {
	return i.ToNotebookInstanceLifecycleConfigurationPtrOutputWithContext(context.Background())
}

func (i *NotebookInstanceLifecycleConfiguration) ToNotebookInstanceLifecycleConfigurationPtrOutputWithContext(ctx context.Context) NotebookInstanceLifecycleConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookInstanceLifecycleConfigurationPtrOutput)
}

type NotebookInstanceLifecycleConfigurationPtrInput interface {
	pulumi.Input

	ToNotebookInstanceLifecycleConfigurationPtrOutput() NotebookInstanceLifecycleConfigurationPtrOutput
	ToNotebookInstanceLifecycleConfigurationPtrOutputWithContext(ctx context.Context) NotebookInstanceLifecycleConfigurationPtrOutput
}

type notebookInstanceLifecycleConfigurationPtrType NotebookInstanceLifecycleConfigurationArgs

func (*notebookInstanceLifecycleConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookInstanceLifecycleConfiguration)(nil))
}

func (i *notebookInstanceLifecycleConfigurationPtrType) ToNotebookInstanceLifecycleConfigurationPtrOutput() NotebookInstanceLifecycleConfigurationPtrOutput {
	return i.ToNotebookInstanceLifecycleConfigurationPtrOutputWithContext(context.Background())
}

func (i *notebookInstanceLifecycleConfigurationPtrType) ToNotebookInstanceLifecycleConfigurationPtrOutputWithContext(ctx context.Context) NotebookInstanceLifecycleConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookInstanceLifecycleConfigurationPtrOutput)
}

// NotebookInstanceLifecycleConfigurationArrayInput is an input type that accepts NotebookInstanceLifecycleConfigurationArray and NotebookInstanceLifecycleConfigurationArrayOutput values.
// You can construct a concrete instance of `NotebookInstanceLifecycleConfigurationArrayInput` via:
//
//          NotebookInstanceLifecycleConfigurationArray{ NotebookInstanceLifecycleConfigurationArgs{...} }
type NotebookInstanceLifecycleConfigurationArrayInput interface {
	pulumi.Input

	ToNotebookInstanceLifecycleConfigurationArrayOutput() NotebookInstanceLifecycleConfigurationArrayOutput
	ToNotebookInstanceLifecycleConfigurationArrayOutputWithContext(context.Context) NotebookInstanceLifecycleConfigurationArrayOutput
}

type NotebookInstanceLifecycleConfigurationArray []NotebookInstanceLifecycleConfigurationInput

func (NotebookInstanceLifecycleConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotebookInstanceLifecycleConfiguration)(nil)).Elem()
}

func (i NotebookInstanceLifecycleConfigurationArray) ToNotebookInstanceLifecycleConfigurationArrayOutput() NotebookInstanceLifecycleConfigurationArrayOutput {
	return i.ToNotebookInstanceLifecycleConfigurationArrayOutputWithContext(context.Background())
}

func (i NotebookInstanceLifecycleConfigurationArray) ToNotebookInstanceLifecycleConfigurationArrayOutputWithContext(ctx context.Context) NotebookInstanceLifecycleConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookInstanceLifecycleConfigurationArrayOutput)
}

// NotebookInstanceLifecycleConfigurationMapInput is an input type that accepts NotebookInstanceLifecycleConfigurationMap and NotebookInstanceLifecycleConfigurationMapOutput values.
// You can construct a concrete instance of `NotebookInstanceLifecycleConfigurationMapInput` via:
//
//          NotebookInstanceLifecycleConfigurationMap{ "key": NotebookInstanceLifecycleConfigurationArgs{...} }
type NotebookInstanceLifecycleConfigurationMapInput interface {
	pulumi.Input

	ToNotebookInstanceLifecycleConfigurationMapOutput() NotebookInstanceLifecycleConfigurationMapOutput
	ToNotebookInstanceLifecycleConfigurationMapOutputWithContext(context.Context) NotebookInstanceLifecycleConfigurationMapOutput
}

type NotebookInstanceLifecycleConfigurationMap map[string]NotebookInstanceLifecycleConfigurationInput

func (NotebookInstanceLifecycleConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotebookInstanceLifecycleConfiguration)(nil)).Elem()
}

func (i NotebookInstanceLifecycleConfigurationMap) ToNotebookInstanceLifecycleConfigurationMapOutput() NotebookInstanceLifecycleConfigurationMapOutput {
	return i.ToNotebookInstanceLifecycleConfigurationMapOutputWithContext(context.Background())
}

func (i NotebookInstanceLifecycleConfigurationMap) ToNotebookInstanceLifecycleConfigurationMapOutputWithContext(ctx context.Context) NotebookInstanceLifecycleConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookInstanceLifecycleConfigurationMapOutput)
}

type NotebookInstanceLifecycleConfigurationOutput struct {
	*pulumi.OutputState
}

func (NotebookInstanceLifecycleConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookInstanceLifecycleConfiguration)(nil))
}

func (o NotebookInstanceLifecycleConfigurationOutput) ToNotebookInstanceLifecycleConfigurationOutput() NotebookInstanceLifecycleConfigurationOutput {
	return o
}

func (o NotebookInstanceLifecycleConfigurationOutput) ToNotebookInstanceLifecycleConfigurationOutputWithContext(ctx context.Context) NotebookInstanceLifecycleConfigurationOutput {
	return o
}

func (o NotebookInstanceLifecycleConfigurationOutput) ToNotebookInstanceLifecycleConfigurationPtrOutput() NotebookInstanceLifecycleConfigurationPtrOutput {
	return o.ToNotebookInstanceLifecycleConfigurationPtrOutputWithContext(context.Background())
}

func (o NotebookInstanceLifecycleConfigurationOutput) ToNotebookInstanceLifecycleConfigurationPtrOutputWithContext(ctx context.Context) NotebookInstanceLifecycleConfigurationPtrOutput {
	return o.ApplyT(func(v NotebookInstanceLifecycleConfiguration) *NotebookInstanceLifecycleConfiguration {
		return &v
	}).(NotebookInstanceLifecycleConfigurationPtrOutput)
}

type NotebookInstanceLifecycleConfigurationPtrOutput struct {
	*pulumi.OutputState
}

func (NotebookInstanceLifecycleConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookInstanceLifecycleConfiguration)(nil))
}

func (o NotebookInstanceLifecycleConfigurationPtrOutput) ToNotebookInstanceLifecycleConfigurationPtrOutput() NotebookInstanceLifecycleConfigurationPtrOutput {
	return o
}

func (o NotebookInstanceLifecycleConfigurationPtrOutput) ToNotebookInstanceLifecycleConfigurationPtrOutputWithContext(ctx context.Context) NotebookInstanceLifecycleConfigurationPtrOutput {
	return o
}

type NotebookInstanceLifecycleConfigurationArrayOutput struct{ *pulumi.OutputState }

func (NotebookInstanceLifecycleConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotebookInstanceLifecycleConfiguration)(nil))
}

func (o NotebookInstanceLifecycleConfigurationArrayOutput) ToNotebookInstanceLifecycleConfigurationArrayOutput() NotebookInstanceLifecycleConfigurationArrayOutput {
	return o
}

func (o NotebookInstanceLifecycleConfigurationArrayOutput) ToNotebookInstanceLifecycleConfigurationArrayOutputWithContext(ctx context.Context) NotebookInstanceLifecycleConfigurationArrayOutput {
	return o
}

func (o NotebookInstanceLifecycleConfigurationArrayOutput) Index(i pulumi.IntInput) NotebookInstanceLifecycleConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NotebookInstanceLifecycleConfiguration {
		return vs[0].([]NotebookInstanceLifecycleConfiguration)[vs[1].(int)]
	}).(NotebookInstanceLifecycleConfigurationOutput)
}

type NotebookInstanceLifecycleConfigurationMapOutput struct{ *pulumi.OutputState }

func (NotebookInstanceLifecycleConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NotebookInstanceLifecycleConfiguration)(nil))
}

func (o NotebookInstanceLifecycleConfigurationMapOutput) ToNotebookInstanceLifecycleConfigurationMapOutput() NotebookInstanceLifecycleConfigurationMapOutput {
	return o
}

func (o NotebookInstanceLifecycleConfigurationMapOutput) ToNotebookInstanceLifecycleConfigurationMapOutputWithContext(ctx context.Context) NotebookInstanceLifecycleConfigurationMapOutput {
	return o
}

func (o NotebookInstanceLifecycleConfigurationMapOutput) MapIndex(k pulumi.StringInput) NotebookInstanceLifecycleConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NotebookInstanceLifecycleConfiguration {
		return vs[0].(map[string]NotebookInstanceLifecycleConfiguration)[vs[1].(string)]
	}).(NotebookInstanceLifecycleConfigurationOutput)
}

func init() {
	pulumi.RegisterOutputType(NotebookInstanceLifecycleConfigurationOutput{})
	pulumi.RegisterOutputType(NotebookInstanceLifecycleConfigurationPtrOutput{})
	pulumi.RegisterOutputType(NotebookInstanceLifecycleConfigurationArrayOutput{})
	pulumi.RegisterOutputType(NotebookInstanceLifecycleConfigurationMapOutput{})
}
