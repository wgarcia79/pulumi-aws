// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an Gamelift Build resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/gamelift"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := gamelift.NewBuild(ctx, "test", &gamelift.BuildArgs{
// 			OperatingSystem: pulumi.String("WINDOWS_2012"),
// 			StorageLocation: &gamelift.BuildStorageLocationArgs{
// 				Bucket:  pulumi.Any(aws_s3_bucket.Test.Bucket),
// 				Key:     pulumi.Any(aws_s3_bucket_object.Test.Key),
// 				RoleArn: pulumi.Any(aws_iam_role.Test.Arn),
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			aws_iam_role_policy.Test,
// 		}))
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
// Gamelift Builds cannot be imported at this time.
type Build struct {
	pulumi.CustomResourceState

	// Gamelift Build ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name of the build
	Name pulumi.StringOutput `pulumi:"name"`
	// Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
	OperatingSystem pulumi.StringOutput `pulumi:"operatingSystem"`
	// Information indicating where your game build files are stored. See below.
	StorageLocation BuildStorageLocationOutput `pulumi:"storageLocation"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Version that is associated with this build.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewBuild registers a new resource with the given unique name, arguments, and options.
func NewBuild(ctx *pulumi.Context,
	name string, args *BuildArgs, opts ...pulumi.ResourceOption) (*Build, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OperatingSystem == nil {
		return nil, errors.New("invalid value for required argument 'OperatingSystem'")
	}
	if args.StorageLocation == nil {
		return nil, errors.New("invalid value for required argument 'StorageLocation'")
	}
	var resource Build
	err := ctx.RegisterResource("aws:gamelift/build:Build", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBuild gets an existing Build resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBuild(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildState, opts ...pulumi.ResourceOption) (*Build, error) {
	var resource Build
	err := ctx.ReadResource("aws:gamelift/build:Build", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Build resources.
type buildState struct {
	// Gamelift Build ARN.
	Arn *string `pulumi:"arn"`
	// Name of the build
	Name *string `pulumi:"name"`
	// Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
	OperatingSystem *string `pulumi:"operatingSystem"`
	// Information indicating where your game build files are stored. See below.
	StorageLocation *BuildStorageLocation `pulumi:"storageLocation"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// Version that is associated with this build.
	Version *string `pulumi:"version"`
}

type BuildState struct {
	// Gamelift Build ARN.
	Arn pulumi.StringPtrInput
	// Name of the build
	Name pulumi.StringPtrInput
	// Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
	OperatingSystem pulumi.StringPtrInput
	// Information indicating where your game build files are stored. See below.
	StorageLocation BuildStorageLocationPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// Version that is associated with this build.
	Version pulumi.StringPtrInput
}

func (BuildState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildState)(nil)).Elem()
}

type buildArgs struct {
	// Name of the build
	Name *string `pulumi:"name"`
	// Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
	OperatingSystem string `pulumi:"operatingSystem"`
	// Information indicating where your game build files are stored. See below.
	StorageLocation BuildStorageLocation `pulumi:"storageLocation"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// Version that is associated with this build.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Build resource.
type BuildArgs struct {
	// Name of the build
	Name pulumi.StringPtrInput
	// Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
	OperatingSystem pulumi.StringInput
	// Information indicating where your game build files are stored. See below.
	StorageLocation BuildStorageLocationInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// Version that is associated with this build.
	Version pulumi.StringPtrInput
}

func (BuildArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildArgs)(nil)).Elem()
}

type BuildInput interface {
	pulumi.Input

	ToBuildOutput() BuildOutput
	ToBuildOutputWithContext(ctx context.Context) BuildOutput
}

func (*Build) ElementType() reflect.Type {
	return reflect.TypeOf((*Build)(nil))
}

func (i *Build) ToBuildOutput() BuildOutput {
	return i.ToBuildOutputWithContext(context.Background())
}

func (i *Build) ToBuildOutputWithContext(ctx context.Context) BuildOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildOutput)
}

func (i *Build) ToBuildPtrOutput() BuildPtrOutput {
	return i.ToBuildPtrOutputWithContext(context.Background())
}

func (i *Build) ToBuildPtrOutputWithContext(ctx context.Context) BuildPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildPtrOutput)
}

type BuildPtrInput interface {
	pulumi.Input

	ToBuildPtrOutput() BuildPtrOutput
	ToBuildPtrOutputWithContext(ctx context.Context) BuildPtrOutput
}

type buildPtrType BuildArgs

func (*buildPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Build)(nil))
}

func (i *buildPtrType) ToBuildPtrOutput() BuildPtrOutput {
	return i.ToBuildPtrOutputWithContext(context.Background())
}

func (i *buildPtrType) ToBuildPtrOutputWithContext(ctx context.Context) BuildPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildPtrOutput)
}

// BuildArrayInput is an input type that accepts BuildArray and BuildArrayOutput values.
// You can construct a concrete instance of `BuildArrayInput` via:
//
//          BuildArray{ BuildArgs{...} }
type BuildArrayInput interface {
	pulumi.Input

	ToBuildArrayOutput() BuildArrayOutput
	ToBuildArrayOutputWithContext(context.Context) BuildArrayOutput
}

type BuildArray []BuildInput

func (BuildArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Build)(nil))
}

func (i BuildArray) ToBuildArrayOutput() BuildArrayOutput {
	return i.ToBuildArrayOutputWithContext(context.Background())
}

func (i BuildArray) ToBuildArrayOutputWithContext(ctx context.Context) BuildArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildArrayOutput)
}

// BuildMapInput is an input type that accepts BuildMap and BuildMapOutput values.
// You can construct a concrete instance of `BuildMapInput` via:
//
//          BuildMap{ "key": BuildArgs{...} }
type BuildMapInput interface {
	pulumi.Input

	ToBuildMapOutput() BuildMapOutput
	ToBuildMapOutputWithContext(context.Context) BuildMapOutput
}

type BuildMap map[string]BuildInput

func (BuildMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Build)(nil))
}

func (i BuildMap) ToBuildMapOutput() BuildMapOutput {
	return i.ToBuildMapOutputWithContext(context.Background())
}

func (i BuildMap) ToBuildMapOutputWithContext(ctx context.Context) BuildMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildMapOutput)
}

type BuildOutput struct {
	*pulumi.OutputState
}

func (BuildOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Build)(nil))
}

func (o BuildOutput) ToBuildOutput() BuildOutput {
	return o
}

func (o BuildOutput) ToBuildOutputWithContext(ctx context.Context) BuildOutput {
	return o
}

func (o BuildOutput) ToBuildPtrOutput() BuildPtrOutput {
	return o.ToBuildPtrOutputWithContext(context.Background())
}

func (o BuildOutput) ToBuildPtrOutputWithContext(ctx context.Context) BuildPtrOutput {
	return o.ApplyT(func(v Build) *Build {
		return &v
	}).(BuildPtrOutput)
}

type BuildPtrOutput struct {
	*pulumi.OutputState
}

func (BuildPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Build)(nil))
}

func (o BuildPtrOutput) ToBuildPtrOutput() BuildPtrOutput {
	return o
}

func (o BuildPtrOutput) ToBuildPtrOutputWithContext(ctx context.Context) BuildPtrOutput {
	return o
}

type BuildArrayOutput struct{ *pulumi.OutputState }

func (BuildArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Build)(nil))
}

func (o BuildArrayOutput) ToBuildArrayOutput() BuildArrayOutput {
	return o
}

func (o BuildArrayOutput) ToBuildArrayOutputWithContext(ctx context.Context) BuildArrayOutput {
	return o
}

func (o BuildArrayOutput) Index(i pulumi.IntInput) BuildOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Build {
		return vs[0].([]Build)[vs[1].(int)]
	}).(BuildOutput)
}

type BuildMapOutput struct{ *pulumi.OutputState }

func (BuildMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Build)(nil))
}

func (o BuildMapOutput) ToBuildMapOutput() BuildMapOutput {
	return o
}

func (o BuildMapOutput) ToBuildMapOutputWithContext(ctx context.Context) BuildMapOutput {
	return o
}

func (o BuildMapOutput) MapIndex(k pulumi.StringInput) BuildOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Build {
		return vs[0].(map[string]Build)[vs[1].(string)]
	}).(BuildOutput)
}

func init() {
	pulumi.RegisterOutputType(BuildOutput{})
	pulumi.RegisterOutputType(BuildPtrOutput{})
	pulumi.RegisterOutputType(BuildArrayOutput{})
	pulumi.RegisterOutputType(BuildMapOutput{})
}
