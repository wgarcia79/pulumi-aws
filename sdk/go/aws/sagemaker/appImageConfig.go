// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Sagemaker App Image Config resource.
//
// ## Example Usage
// ### Basic usage
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
// 		_, err := sagemaker.NewAppImageConfig(ctx, "test", &sagemaker.AppImageConfigArgs{
// 			AppImageConfigName: pulumi.String("example"),
// 			KernelGatewayImageConfig: &sagemaker.AppImageConfigKernelGatewayImageConfigArgs{
// 				KernelSpec: &sagemaker.AppImageConfigKernelGatewayImageConfigKernelSpecArgs{
// 					Name: pulumi.String("example"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Default File System Config
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
// 		_, err := sagemaker.NewAppImageConfig(ctx, "test", &sagemaker.AppImageConfigArgs{
// 			AppImageConfigName: pulumi.String("example"),
// 			KernelGatewayImageConfig: &sagemaker.AppImageConfigKernelGatewayImageConfigArgs{
// 				FileSystemConfig: nil,
// 				KernelSpec: &sagemaker.AppImageConfigKernelGatewayImageConfigKernelSpecArgs{
// 					Name: pulumi.String("example"),
// 				},
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
// Sagemaker App Image Configs can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:sagemaker/appImageConfig:AppImageConfig example example
// ```
type AppImageConfig struct {
	pulumi.CustomResourceState

	// The name of the App Image Config.
	AppImageConfigName pulumi.StringOutput `pulumi:"appImageConfigName"`
	// The Amazon Resource Name (ARN) assigned by AWS to this App Image Config.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The configuration for the file system and kernels in a SageMaker image running as a KernelGateway app. See Kernel Gateway Image Config details below.
	KernelGatewayImageConfig AppImageConfigKernelGatewayImageConfigPtrOutput `pulumi:"kernelGatewayImageConfig"`
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewAppImageConfig registers a new resource with the given unique name, arguments, and options.
func NewAppImageConfig(ctx *pulumi.Context,
	name string, args *AppImageConfigArgs, opts ...pulumi.ResourceOption) (*AppImageConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppImageConfigName == nil {
		return nil, errors.New("invalid value for required argument 'AppImageConfigName'")
	}
	var resource AppImageConfig
	err := ctx.RegisterResource("aws:sagemaker/appImageConfig:AppImageConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppImageConfig gets an existing AppImageConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppImageConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppImageConfigState, opts ...pulumi.ResourceOption) (*AppImageConfig, error) {
	var resource AppImageConfig
	err := ctx.ReadResource("aws:sagemaker/appImageConfig:AppImageConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppImageConfig resources.
type appImageConfigState struct {
	// The name of the App Image Config.
	AppImageConfigName *string `pulumi:"appImageConfigName"`
	// The Amazon Resource Name (ARN) assigned by AWS to this App Image Config.
	Arn *string `pulumi:"arn"`
	// The configuration for the file system and kernels in a SageMaker image running as a KernelGateway app. See Kernel Gateway Image Config details below.
	KernelGatewayImageConfig *AppImageConfigKernelGatewayImageConfig `pulumi:"kernelGatewayImageConfig"`
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type AppImageConfigState struct {
	// The name of the App Image Config.
	AppImageConfigName pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) assigned by AWS to this App Image Config.
	Arn pulumi.StringPtrInput
	// The configuration for the file system and kernels in a SageMaker image running as a KernelGateway app. See Kernel Gateway Image Config details below.
	KernelGatewayImageConfig AppImageConfigKernelGatewayImageConfigPtrInput
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapInput
}

func (AppImageConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*appImageConfigState)(nil)).Elem()
}

type appImageConfigArgs struct {
	// The name of the App Image Config.
	AppImageConfigName string `pulumi:"appImageConfigName"`
	// The configuration for the file system and kernels in a SageMaker image running as a KernelGateway app. See Kernel Gateway Image Config details below.
	KernelGatewayImageConfig *AppImageConfigKernelGatewayImageConfig `pulumi:"kernelGatewayImageConfig"`
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a AppImageConfig resource.
type AppImageConfigArgs struct {
	// The name of the App Image Config.
	AppImageConfigName pulumi.StringInput
	// The configuration for the file system and kernels in a SageMaker image running as a KernelGateway app. See Kernel Gateway Image Config details below.
	KernelGatewayImageConfig AppImageConfigKernelGatewayImageConfigPtrInput
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapInput
}

func (AppImageConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appImageConfigArgs)(nil)).Elem()
}

type AppImageConfigInput interface {
	pulumi.Input

	ToAppImageConfigOutput() AppImageConfigOutput
	ToAppImageConfigOutputWithContext(ctx context.Context) AppImageConfigOutput
}

func (*AppImageConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*AppImageConfig)(nil))
}

func (i *AppImageConfig) ToAppImageConfigOutput() AppImageConfigOutput {
	return i.ToAppImageConfigOutputWithContext(context.Background())
}

func (i *AppImageConfig) ToAppImageConfigOutputWithContext(ctx context.Context) AppImageConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppImageConfigOutput)
}

func (i *AppImageConfig) ToAppImageConfigPtrOutput() AppImageConfigPtrOutput {
	return i.ToAppImageConfigPtrOutputWithContext(context.Background())
}

func (i *AppImageConfig) ToAppImageConfigPtrOutputWithContext(ctx context.Context) AppImageConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppImageConfigPtrOutput)
}

type AppImageConfigPtrInput interface {
	pulumi.Input

	ToAppImageConfigPtrOutput() AppImageConfigPtrOutput
	ToAppImageConfigPtrOutputWithContext(ctx context.Context) AppImageConfigPtrOutput
}

type appImageConfigPtrType AppImageConfigArgs

func (*appImageConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppImageConfig)(nil))
}

func (i *appImageConfigPtrType) ToAppImageConfigPtrOutput() AppImageConfigPtrOutput {
	return i.ToAppImageConfigPtrOutputWithContext(context.Background())
}

func (i *appImageConfigPtrType) ToAppImageConfigPtrOutputWithContext(ctx context.Context) AppImageConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppImageConfigPtrOutput)
}

// AppImageConfigArrayInput is an input type that accepts AppImageConfigArray and AppImageConfigArrayOutput values.
// You can construct a concrete instance of `AppImageConfigArrayInput` via:
//
//          AppImageConfigArray{ AppImageConfigArgs{...} }
type AppImageConfigArrayInput interface {
	pulumi.Input

	ToAppImageConfigArrayOutput() AppImageConfigArrayOutput
	ToAppImageConfigArrayOutputWithContext(context.Context) AppImageConfigArrayOutput
}

type AppImageConfigArray []AppImageConfigInput

func (AppImageConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppImageConfig)(nil)).Elem()
}

func (i AppImageConfigArray) ToAppImageConfigArrayOutput() AppImageConfigArrayOutput {
	return i.ToAppImageConfigArrayOutputWithContext(context.Background())
}

func (i AppImageConfigArray) ToAppImageConfigArrayOutputWithContext(ctx context.Context) AppImageConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppImageConfigArrayOutput)
}

// AppImageConfigMapInput is an input type that accepts AppImageConfigMap and AppImageConfigMapOutput values.
// You can construct a concrete instance of `AppImageConfigMapInput` via:
//
//          AppImageConfigMap{ "key": AppImageConfigArgs{...} }
type AppImageConfigMapInput interface {
	pulumi.Input

	ToAppImageConfigMapOutput() AppImageConfigMapOutput
	ToAppImageConfigMapOutputWithContext(context.Context) AppImageConfigMapOutput
}

type AppImageConfigMap map[string]AppImageConfigInput

func (AppImageConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppImageConfig)(nil)).Elem()
}

func (i AppImageConfigMap) ToAppImageConfigMapOutput() AppImageConfigMapOutput {
	return i.ToAppImageConfigMapOutputWithContext(context.Background())
}

func (i AppImageConfigMap) ToAppImageConfigMapOutputWithContext(ctx context.Context) AppImageConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppImageConfigMapOutput)
}

type AppImageConfigOutput struct{ *pulumi.OutputState }

func (AppImageConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppImageConfig)(nil))
}

func (o AppImageConfigOutput) ToAppImageConfigOutput() AppImageConfigOutput {
	return o
}

func (o AppImageConfigOutput) ToAppImageConfigOutputWithContext(ctx context.Context) AppImageConfigOutput {
	return o
}

func (o AppImageConfigOutput) ToAppImageConfigPtrOutput() AppImageConfigPtrOutput {
	return o.ToAppImageConfigPtrOutputWithContext(context.Background())
}

func (o AppImageConfigOutput) ToAppImageConfigPtrOutputWithContext(ctx context.Context) AppImageConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppImageConfig) *AppImageConfig {
		return &v
	}).(AppImageConfigPtrOutput)
}

type AppImageConfigPtrOutput struct{ *pulumi.OutputState }

func (AppImageConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppImageConfig)(nil))
}

func (o AppImageConfigPtrOutput) ToAppImageConfigPtrOutput() AppImageConfigPtrOutput {
	return o
}

func (o AppImageConfigPtrOutput) ToAppImageConfigPtrOutputWithContext(ctx context.Context) AppImageConfigPtrOutput {
	return o
}

func (o AppImageConfigPtrOutput) Elem() AppImageConfigOutput {
	return o.ApplyT(func(v *AppImageConfig) AppImageConfig {
		if v != nil {
			return *v
		}
		var ret AppImageConfig
		return ret
	}).(AppImageConfigOutput)
}

type AppImageConfigArrayOutput struct{ *pulumi.OutputState }

func (AppImageConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppImageConfig)(nil))
}

func (o AppImageConfigArrayOutput) ToAppImageConfigArrayOutput() AppImageConfigArrayOutput {
	return o
}

func (o AppImageConfigArrayOutput) ToAppImageConfigArrayOutputWithContext(ctx context.Context) AppImageConfigArrayOutput {
	return o
}

func (o AppImageConfigArrayOutput) Index(i pulumi.IntInput) AppImageConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppImageConfig {
		return vs[0].([]AppImageConfig)[vs[1].(int)]
	}).(AppImageConfigOutput)
}

type AppImageConfigMapOutput struct{ *pulumi.OutputState }

func (AppImageConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppImageConfig)(nil))
}

func (o AppImageConfigMapOutput) ToAppImageConfigMapOutput() AppImageConfigMapOutput {
	return o
}

func (o AppImageConfigMapOutput) ToAppImageConfigMapOutputWithContext(ctx context.Context) AppImageConfigMapOutput {
	return o
}

func (o AppImageConfigMapOutput) MapIndex(k pulumi.StringInput) AppImageConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppImageConfig {
		return vs[0].(map[string]AppImageConfig)[vs[1].(string)]
	}).(AppImageConfigOutput)
}

func init() {
	pulumi.RegisterOutputType(AppImageConfigOutput{})
	pulumi.RegisterOutputType(AppImageConfigPtrOutput{})
	pulumi.RegisterOutputType(AppImageConfigArrayOutput{})
	pulumi.RegisterOutputType(AppImageConfigMapOutput{})
}
