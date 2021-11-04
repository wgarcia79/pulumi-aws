// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Sagemaker Device Fleet resource.
//
// ## Example Usage
// ### Basic usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sagemaker"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sagemaker.NewDeviceFleet(ctx, "example", &sagemaker.DeviceFleetArgs{
// 			DeviceFleetName: pulumi.String("example"),
// 			RoleArn:         pulumi.Any(aws_iam_role.Test.Arn),
// 			OutputConfig: &sagemaker.DeviceFleetOutputConfigArgs{
// 				S3OutputLocation: pulumi.String(fmt.Sprintf("%v%v%v", "s3://", aws_s3_bucket.Example.Bucket, "/prefix/")),
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
// Sagemaker Device Fleets can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:sagemaker/deviceFleet:DeviceFleet example my-fleet
// ```
type DeviceFleet struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) assigned by AWS to this Device Fleet.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A description of the fleet.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the Device Fleet (must be unique).
	DeviceFleetName pulumi.StringOutput `pulumi:"deviceFleetName"`
	// Whether to create an AWS IoT Role Alias during device fleet creation. The name of the role alias generated will match this pattern: "SageMakerEdge-{DeviceFleetName}".
	EnableIotRoleAlias pulumi.BoolPtrOutput `pulumi:"enableIotRoleAlias"`
	IotRoleAlias       pulumi.StringOutput  `pulumi:"iotRoleAlias"`
	// Specifies details about the repository. see Output Config details below.
	OutputConfig DeviceFleetOutputConfigOutput `pulumi:"outputConfig"`
	// The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewDeviceFleet registers a new resource with the given unique name, arguments, and options.
func NewDeviceFleet(ctx *pulumi.Context,
	name string, args *DeviceFleetArgs, opts ...pulumi.ResourceOption) (*DeviceFleet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceFleetName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceFleetName'")
	}
	if args.OutputConfig == nil {
		return nil, errors.New("invalid value for required argument 'OutputConfig'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource DeviceFleet
	err := ctx.RegisterResource("aws:sagemaker/deviceFleet:DeviceFleet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeviceFleet gets an existing DeviceFleet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceFleet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceFleetState, opts ...pulumi.ResourceOption) (*DeviceFleet, error) {
	var resource DeviceFleet
	err := ctx.ReadResource("aws:sagemaker/deviceFleet:DeviceFleet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeviceFleet resources.
type deviceFleetState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this Device Fleet.
	Arn *string `pulumi:"arn"`
	// A description of the fleet.
	Description *string `pulumi:"description"`
	// The name of the Device Fleet (must be unique).
	DeviceFleetName *string `pulumi:"deviceFleetName"`
	// Whether to create an AWS IoT Role Alias during device fleet creation. The name of the role alias generated will match this pattern: "SageMakerEdge-{DeviceFleetName}".
	EnableIotRoleAlias *bool   `pulumi:"enableIotRoleAlias"`
	IotRoleAlias       *string `pulumi:"iotRoleAlias"`
	// Specifies details about the repository. see Output Config details below.
	OutputConfig *DeviceFleetOutputConfig `pulumi:"outputConfig"`
	// The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
	RoleArn *string `pulumi:"roleArn"`
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type DeviceFleetState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this Device Fleet.
	Arn pulumi.StringPtrInput
	// A description of the fleet.
	Description pulumi.StringPtrInput
	// The name of the Device Fleet (must be unique).
	DeviceFleetName pulumi.StringPtrInput
	// Whether to create an AWS IoT Role Alias during device fleet creation. The name of the role alias generated will match this pattern: "SageMakerEdge-{DeviceFleetName}".
	EnableIotRoleAlias pulumi.BoolPtrInput
	IotRoleAlias       pulumi.StringPtrInput
	// Specifies details about the repository. see Output Config details below.
	OutputConfig DeviceFleetOutputConfigPtrInput
	// The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
	RoleArn pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapInput
}

func (DeviceFleetState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceFleetState)(nil)).Elem()
}

type deviceFleetArgs struct {
	// A description of the fleet.
	Description *string `pulumi:"description"`
	// The name of the Device Fleet (must be unique).
	DeviceFleetName string `pulumi:"deviceFleetName"`
	// Whether to create an AWS IoT Role Alias during device fleet creation. The name of the role alias generated will match this pattern: "SageMakerEdge-{DeviceFleetName}".
	EnableIotRoleAlias *bool `pulumi:"enableIotRoleAlias"`
	// Specifies details about the repository. see Output Config details below.
	OutputConfig DeviceFleetOutputConfig `pulumi:"outputConfig"`
	// The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
	RoleArn string `pulumi:"roleArn"`
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a DeviceFleet resource.
type DeviceFleetArgs struct {
	// A description of the fleet.
	Description pulumi.StringPtrInput
	// The name of the Device Fleet (must be unique).
	DeviceFleetName pulumi.StringInput
	// Whether to create an AWS IoT Role Alias during device fleet creation. The name of the role alias generated will match this pattern: "SageMakerEdge-{DeviceFleetName}".
	EnableIotRoleAlias pulumi.BoolPtrInput
	// Specifies details about the repository. see Output Config details below.
	OutputConfig DeviceFleetOutputConfigInput
	// The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
	RoleArn pulumi.StringInput
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapInput
}

func (DeviceFleetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceFleetArgs)(nil)).Elem()
}

type DeviceFleetInput interface {
	pulumi.Input

	ToDeviceFleetOutput() DeviceFleetOutput
	ToDeviceFleetOutputWithContext(ctx context.Context) DeviceFleetOutput
}

func (*DeviceFleet) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceFleet)(nil))
}

func (i *DeviceFleet) ToDeviceFleetOutput() DeviceFleetOutput {
	return i.ToDeviceFleetOutputWithContext(context.Background())
}

func (i *DeviceFleet) ToDeviceFleetOutputWithContext(ctx context.Context) DeviceFleetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceFleetOutput)
}

func (i *DeviceFleet) ToDeviceFleetPtrOutput() DeviceFleetPtrOutput {
	return i.ToDeviceFleetPtrOutputWithContext(context.Background())
}

func (i *DeviceFleet) ToDeviceFleetPtrOutputWithContext(ctx context.Context) DeviceFleetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceFleetPtrOutput)
}

type DeviceFleetPtrInput interface {
	pulumi.Input

	ToDeviceFleetPtrOutput() DeviceFleetPtrOutput
	ToDeviceFleetPtrOutputWithContext(ctx context.Context) DeviceFleetPtrOutput
}

type deviceFleetPtrType DeviceFleetArgs

func (*deviceFleetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceFleet)(nil))
}

func (i *deviceFleetPtrType) ToDeviceFleetPtrOutput() DeviceFleetPtrOutput {
	return i.ToDeviceFleetPtrOutputWithContext(context.Background())
}

func (i *deviceFleetPtrType) ToDeviceFleetPtrOutputWithContext(ctx context.Context) DeviceFleetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceFleetPtrOutput)
}

// DeviceFleetArrayInput is an input type that accepts DeviceFleetArray and DeviceFleetArrayOutput values.
// You can construct a concrete instance of `DeviceFleetArrayInput` via:
//
//          DeviceFleetArray{ DeviceFleetArgs{...} }
type DeviceFleetArrayInput interface {
	pulumi.Input

	ToDeviceFleetArrayOutput() DeviceFleetArrayOutput
	ToDeviceFleetArrayOutputWithContext(context.Context) DeviceFleetArrayOutput
}

type DeviceFleetArray []DeviceFleetInput

func (DeviceFleetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceFleet)(nil)).Elem()
}

func (i DeviceFleetArray) ToDeviceFleetArrayOutput() DeviceFleetArrayOutput {
	return i.ToDeviceFleetArrayOutputWithContext(context.Background())
}

func (i DeviceFleetArray) ToDeviceFleetArrayOutputWithContext(ctx context.Context) DeviceFleetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceFleetArrayOutput)
}

// DeviceFleetMapInput is an input type that accepts DeviceFleetMap and DeviceFleetMapOutput values.
// You can construct a concrete instance of `DeviceFleetMapInput` via:
//
//          DeviceFleetMap{ "key": DeviceFleetArgs{...} }
type DeviceFleetMapInput interface {
	pulumi.Input

	ToDeviceFleetMapOutput() DeviceFleetMapOutput
	ToDeviceFleetMapOutputWithContext(context.Context) DeviceFleetMapOutput
}

type DeviceFleetMap map[string]DeviceFleetInput

func (DeviceFleetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceFleet)(nil)).Elem()
}

func (i DeviceFleetMap) ToDeviceFleetMapOutput() DeviceFleetMapOutput {
	return i.ToDeviceFleetMapOutputWithContext(context.Background())
}

func (i DeviceFleetMap) ToDeviceFleetMapOutputWithContext(ctx context.Context) DeviceFleetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceFleetMapOutput)
}

type DeviceFleetOutput struct{ *pulumi.OutputState }

func (DeviceFleetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceFleet)(nil))
}

func (o DeviceFleetOutput) ToDeviceFleetOutput() DeviceFleetOutput {
	return o
}

func (o DeviceFleetOutput) ToDeviceFleetOutputWithContext(ctx context.Context) DeviceFleetOutput {
	return o
}

func (o DeviceFleetOutput) ToDeviceFleetPtrOutput() DeviceFleetPtrOutput {
	return o.ToDeviceFleetPtrOutputWithContext(context.Background())
}

func (o DeviceFleetOutput) ToDeviceFleetPtrOutputWithContext(ctx context.Context) DeviceFleetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeviceFleet) *DeviceFleet {
		return &v
	}).(DeviceFleetPtrOutput)
}

type DeviceFleetPtrOutput struct{ *pulumi.OutputState }

func (DeviceFleetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceFleet)(nil))
}

func (o DeviceFleetPtrOutput) ToDeviceFleetPtrOutput() DeviceFleetPtrOutput {
	return o
}

func (o DeviceFleetPtrOutput) ToDeviceFleetPtrOutputWithContext(ctx context.Context) DeviceFleetPtrOutput {
	return o
}

func (o DeviceFleetPtrOutput) Elem() DeviceFleetOutput {
	return o.ApplyT(func(v *DeviceFleet) DeviceFleet {
		if v != nil {
			return *v
		}
		var ret DeviceFleet
		return ret
	}).(DeviceFleetOutput)
}

type DeviceFleetArrayOutput struct{ *pulumi.OutputState }

func (DeviceFleetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeviceFleet)(nil))
}

func (o DeviceFleetArrayOutput) ToDeviceFleetArrayOutput() DeviceFleetArrayOutput {
	return o
}

func (o DeviceFleetArrayOutput) ToDeviceFleetArrayOutputWithContext(ctx context.Context) DeviceFleetArrayOutput {
	return o
}

func (o DeviceFleetArrayOutput) Index(i pulumi.IntInput) DeviceFleetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeviceFleet {
		return vs[0].([]DeviceFleet)[vs[1].(int)]
	}).(DeviceFleetOutput)
}

type DeviceFleetMapOutput struct{ *pulumi.OutputState }

func (DeviceFleetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DeviceFleet)(nil))
}

func (o DeviceFleetMapOutput) ToDeviceFleetMapOutput() DeviceFleetMapOutput {
	return o
}

func (o DeviceFleetMapOutput) ToDeviceFleetMapOutputWithContext(ctx context.Context) DeviceFleetMapOutput {
	return o
}

func (o DeviceFleetMapOutput) MapIndex(k pulumi.StringInput) DeviceFleetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DeviceFleet {
		return vs[0].(map[string]DeviceFleet)[vs[1].(string)]
	}).(DeviceFleetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceFleetInput)(nil)).Elem(), &DeviceFleet{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceFleetPtrInput)(nil)).Elem(), &DeviceFleet{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceFleetArrayInput)(nil)).Elem(), DeviceFleetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceFleetMapInput)(nil)).Elem(), DeviceFleetMap{})
	pulumi.RegisterOutputType(DeviceFleetOutput{})
	pulumi.RegisterOutputType(DeviceFleetPtrOutput{})
	pulumi.RegisterOutputType(DeviceFleetArrayOutput{})
	pulumi.RegisterOutputType(DeviceFleetMapOutput{})
}
