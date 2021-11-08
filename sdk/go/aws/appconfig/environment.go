// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appconfig

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AppConfig Environment resource for an `appconfig.Application` resource. One or more environments can be defined for an application.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/appconfig"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleApplication, err := appconfig.NewApplication(ctx, "exampleApplication", &appconfig.ApplicationArgs{
// 			Description: pulumi.String("Example AppConfig Application"),
// 			Tags: pulumi.StringMap{
// 				"Type": pulumi.String("AppConfig Application"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appconfig.NewEnvironment(ctx, "exampleEnvironment", &appconfig.EnvironmentArgs{
// 			Description:   pulumi.String("Example AppConfig Environment"),
// 			ApplicationId: exampleApplication.ID(),
// 			Monitors: appconfig.EnvironmentMonitorArray{
// 				&appconfig.EnvironmentMonitorArgs{
// 					AlarmArn:     pulumi.Any(aws_cloudwatch_metric_alarm.Example.Arn),
// 					AlarmRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Type": pulumi.String("AppConfig Environment"),
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
// AppConfig Environments can be imported by using the environment ID and application ID separated by a colon (`:`), e.g.,
//
// ```sh
//  $ pulumi import aws:appconfig/environment:Environment example 71abcde:11xxxxx
// ```
type Environment struct {
	pulumi.CustomResourceState

	// The AppConfig application ID. Must be between 4 and 7 characters in length.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The Amazon Resource Name (ARN) of the AppConfig Environment.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the environment. Can be at most 1024 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The AppConfig environment ID.
	EnvironmentId pulumi.StringOutput `pulumi:"environmentId"`
	// Set of Amazon CloudWatch alarms to monitor during the deployment process. Maximum of 5. See Monitor below for more details.
	Monitors EnvironmentMonitorArrayOutput `pulumi:"monitors"`
	// The name for the environment. Must be between 1 and 64 characters in length.
	Name  pulumi.StringOutput `pulumi:"name"`
	State pulumi.StringOutput `pulumi:"state"`
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	var resource Environment
	err := ctx.RegisterResource("aws:appconfig/environment:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironment gets an existing Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("aws:appconfig/environment:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Environment resources.
type environmentState struct {
	// The AppConfig application ID. Must be between 4 and 7 characters in length.
	ApplicationId *string `pulumi:"applicationId"`
	// The Amazon Resource Name (ARN) of the AppConfig Environment.
	Arn *string `pulumi:"arn"`
	// The description of the environment. Can be at most 1024 characters.
	Description *string `pulumi:"description"`
	// The AppConfig environment ID.
	EnvironmentId *string `pulumi:"environmentId"`
	// Set of Amazon CloudWatch alarms to monitor during the deployment process. Maximum of 5. See Monitor below for more details.
	Monitors []EnvironmentMonitor `pulumi:"monitors"`
	// The name for the environment. Must be between 1 and 64 characters in length.
	Name  *string `pulumi:"name"`
	State *string `pulumi:"state"`
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type EnvironmentState struct {
	// The AppConfig application ID. Must be between 4 and 7 characters in length.
	ApplicationId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the AppConfig Environment.
	Arn pulumi.StringPtrInput
	// The description of the environment. Can be at most 1024 characters.
	Description pulumi.StringPtrInput
	// The AppConfig environment ID.
	EnvironmentId pulumi.StringPtrInput
	// Set of Amazon CloudWatch alarms to monitor during the deployment process. Maximum of 5. See Monitor below for more details.
	Monitors EnvironmentMonitorArrayInput
	// The name for the environment. Must be between 1 and 64 characters in length.
	Name  pulumi.StringPtrInput
	State pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapInput
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	// The AppConfig application ID. Must be between 4 and 7 characters in length.
	ApplicationId string `pulumi:"applicationId"`
	// The description of the environment. Can be at most 1024 characters.
	Description *string `pulumi:"description"`
	// Set of Amazon CloudWatch alarms to monitor during the deployment process. Maximum of 5. See Monitor below for more details.
	Monitors []EnvironmentMonitor `pulumi:"monitors"`
	// The name for the environment. Must be between 1 and 64 characters in length.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	// The AppConfig application ID. Must be between 4 and 7 characters in length.
	ApplicationId pulumi.StringInput
	// The description of the environment. Can be at most 1024 characters.
	Description pulumi.StringPtrInput
	// Set of Amazon CloudWatch alarms to monitor during the deployment process. Maximum of 5. See Monitor below for more details.
	Monitors EnvironmentMonitorArrayInput
	// The name for the environment. Must be between 1 and 64 characters in length.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}

type EnvironmentInput interface {
	pulumi.Input

	ToEnvironmentOutput() EnvironmentOutput
	ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput
}

func (*Environment) ElementType() reflect.Type {
	return reflect.TypeOf((*Environment)(nil))
}

func (i *Environment) ToEnvironmentOutput() EnvironmentOutput {
	return i.ToEnvironmentOutputWithContext(context.Background())
}

func (i *Environment) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentOutput)
}

func (i *Environment) ToEnvironmentPtrOutput() EnvironmentPtrOutput {
	return i.ToEnvironmentPtrOutputWithContext(context.Background())
}

func (i *Environment) ToEnvironmentPtrOutputWithContext(ctx context.Context) EnvironmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentPtrOutput)
}

type EnvironmentPtrInput interface {
	pulumi.Input

	ToEnvironmentPtrOutput() EnvironmentPtrOutput
	ToEnvironmentPtrOutputWithContext(ctx context.Context) EnvironmentPtrOutput
}

type environmentPtrType EnvironmentArgs

func (*environmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil))
}

func (i *environmentPtrType) ToEnvironmentPtrOutput() EnvironmentPtrOutput {
	return i.ToEnvironmentPtrOutputWithContext(context.Background())
}

func (i *environmentPtrType) ToEnvironmentPtrOutputWithContext(ctx context.Context) EnvironmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentPtrOutput)
}

// EnvironmentArrayInput is an input type that accepts EnvironmentArray and EnvironmentArrayOutput values.
// You can construct a concrete instance of `EnvironmentArrayInput` via:
//
//          EnvironmentArray{ EnvironmentArgs{...} }
type EnvironmentArrayInput interface {
	pulumi.Input

	ToEnvironmentArrayOutput() EnvironmentArrayOutput
	ToEnvironmentArrayOutputWithContext(context.Context) EnvironmentArrayOutput
}

type EnvironmentArray []EnvironmentInput

func (EnvironmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Environment)(nil)).Elem()
}

func (i EnvironmentArray) ToEnvironmentArrayOutput() EnvironmentArrayOutput {
	return i.ToEnvironmentArrayOutputWithContext(context.Background())
}

func (i EnvironmentArray) ToEnvironmentArrayOutputWithContext(ctx context.Context) EnvironmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentArrayOutput)
}

// EnvironmentMapInput is an input type that accepts EnvironmentMap and EnvironmentMapOutput values.
// You can construct a concrete instance of `EnvironmentMapInput` via:
//
//          EnvironmentMap{ "key": EnvironmentArgs{...} }
type EnvironmentMapInput interface {
	pulumi.Input

	ToEnvironmentMapOutput() EnvironmentMapOutput
	ToEnvironmentMapOutputWithContext(context.Context) EnvironmentMapOutput
}

type EnvironmentMap map[string]EnvironmentInput

func (EnvironmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Environment)(nil)).Elem()
}

func (i EnvironmentMap) ToEnvironmentMapOutput() EnvironmentMapOutput {
	return i.ToEnvironmentMapOutputWithContext(context.Background())
}

func (i EnvironmentMap) ToEnvironmentMapOutputWithContext(ctx context.Context) EnvironmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentMapOutput)
}

type EnvironmentOutput struct{ *pulumi.OutputState }

func (EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Environment)(nil))
}

func (o EnvironmentOutput) ToEnvironmentOutput() EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ToEnvironmentPtrOutput() EnvironmentPtrOutput {
	return o.ToEnvironmentPtrOutputWithContext(context.Background())
}

func (o EnvironmentOutput) ToEnvironmentPtrOutputWithContext(ctx context.Context) EnvironmentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Environment) *Environment {
		return &v
	}).(EnvironmentPtrOutput)
}

type EnvironmentPtrOutput struct{ *pulumi.OutputState }

func (EnvironmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil))
}

func (o EnvironmentPtrOutput) ToEnvironmentPtrOutput() EnvironmentPtrOutput {
	return o
}

func (o EnvironmentPtrOutput) ToEnvironmentPtrOutputWithContext(ctx context.Context) EnvironmentPtrOutput {
	return o
}

func (o EnvironmentPtrOutput) Elem() EnvironmentOutput {
	return o.ApplyT(func(v *Environment) Environment {
		if v != nil {
			return *v
		}
		var ret Environment
		return ret
	}).(EnvironmentOutput)
}

type EnvironmentArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Environment)(nil))
}

func (o EnvironmentArrayOutput) ToEnvironmentArrayOutput() EnvironmentArrayOutput {
	return o
}

func (o EnvironmentArrayOutput) ToEnvironmentArrayOutputWithContext(ctx context.Context) EnvironmentArrayOutput {
	return o
}

func (o EnvironmentArrayOutput) Index(i pulumi.IntInput) EnvironmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Environment {
		return vs[0].([]Environment)[vs[1].(int)]
	}).(EnvironmentOutput)
}

type EnvironmentMapOutput struct{ *pulumi.OutputState }

func (EnvironmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Environment)(nil))
}

func (o EnvironmentMapOutput) ToEnvironmentMapOutput() EnvironmentMapOutput {
	return o
}

func (o EnvironmentMapOutput) ToEnvironmentMapOutputWithContext(ctx context.Context) EnvironmentMapOutput {
	return o
}

func (o EnvironmentMapOutput) MapIndex(k pulumi.StringInput) EnvironmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Environment {
		return vs[0].(map[string]Environment)[vs[1].(string)]
	}).(EnvironmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentInput)(nil)).Elem(), &Environment{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentPtrInput)(nil)).Elem(), &Environment{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentArrayInput)(nil)).Elem(), EnvironmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentMapInput)(nil)).Elem(), EnvironmentMap{})
	pulumi.RegisterOutputType(EnvironmentOutput{})
	pulumi.RegisterOutputType(EnvironmentPtrOutput{})
	pulumi.RegisterOutputType(EnvironmentArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentMapOutput{})
}
