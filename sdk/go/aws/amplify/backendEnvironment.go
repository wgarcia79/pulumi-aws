// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package amplify

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Amplify Backend Environment resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/amplify"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleApp, err := amplify.NewApp(ctx, "exampleApp", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = amplify.NewBackendEnvironment(ctx, "exampleBackendEnvironment", &amplify.BackendEnvironmentArgs{
// 			AppId:               exampleApp.ID(),
// 			EnvironmentName:     pulumi.String("example"),
// 			DeploymentArtifacts: pulumi.String("app-example-deployment"),
// 			StackName:           pulumi.String("amplify-app-example"),
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
// Amplify backend environment can be imported using `app_id` and `environment_name`, e.g.
//
// ```sh
//  $ pulumi import aws:amplify/backendEnvironment:BackendEnvironment example d2ypk4k47z8u6/example
// ```
type BackendEnvironment struct {
	pulumi.CustomResourceState

	// The unique ID for an Amplify app.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The Amazon Resource Name (ARN) for a backend environment that is part of an Amplify app.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of deployment artifacts.
	DeploymentArtifacts pulumi.StringOutput `pulumi:"deploymentArtifacts"`
	// The name for the backend environment.
	EnvironmentName pulumi.StringOutput `pulumi:"environmentName"`
	// The AWS CloudFormation stack name of a backend environment.
	StackName pulumi.StringOutput `pulumi:"stackName"`
}

// NewBackendEnvironment registers a new resource with the given unique name, arguments, and options.
func NewBackendEnvironment(ctx *pulumi.Context,
	name string, args *BackendEnvironmentArgs, opts ...pulumi.ResourceOption) (*BackendEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	var resource BackendEnvironment
	err := ctx.RegisterResource("aws:amplify/backendEnvironment:BackendEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendEnvironment gets an existing BackendEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendEnvironmentState, opts ...pulumi.ResourceOption) (*BackendEnvironment, error) {
	var resource BackendEnvironment
	err := ctx.ReadResource("aws:amplify/backendEnvironment:BackendEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendEnvironment resources.
type backendEnvironmentState struct {
	// The unique ID for an Amplify app.
	AppId *string `pulumi:"appId"`
	// The Amazon Resource Name (ARN) for a backend environment that is part of an Amplify app.
	Arn *string `pulumi:"arn"`
	// The name of deployment artifacts.
	DeploymentArtifacts *string `pulumi:"deploymentArtifacts"`
	// The name for the backend environment.
	EnvironmentName *string `pulumi:"environmentName"`
	// The AWS CloudFormation stack name of a backend environment.
	StackName *string `pulumi:"stackName"`
}

type BackendEnvironmentState struct {
	// The unique ID for an Amplify app.
	AppId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for a backend environment that is part of an Amplify app.
	Arn pulumi.StringPtrInput
	// The name of deployment artifacts.
	DeploymentArtifacts pulumi.StringPtrInput
	// The name for the backend environment.
	EnvironmentName pulumi.StringPtrInput
	// The AWS CloudFormation stack name of a backend environment.
	StackName pulumi.StringPtrInput
}

func (BackendEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendEnvironmentState)(nil)).Elem()
}

type backendEnvironmentArgs struct {
	// The unique ID for an Amplify app.
	AppId string `pulumi:"appId"`
	// The name of deployment artifacts.
	DeploymentArtifacts *string `pulumi:"deploymentArtifacts"`
	// The name for the backend environment.
	EnvironmentName string `pulumi:"environmentName"`
	// The AWS CloudFormation stack name of a backend environment.
	StackName *string `pulumi:"stackName"`
}

// The set of arguments for constructing a BackendEnvironment resource.
type BackendEnvironmentArgs struct {
	// The unique ID for an Amplify app.
	AppId pulumi.StringInput
	// The name of deployment artifacts.
	DeploymentArtifacts pulumi.StringPtrInput
	// The name for the backend environment.
	EnvironmentName pulumi.StringInput
	// The AWS CloudFormation stack name of a backend environment.
	StackName pulumi.StringPtrInput
}

func (BackendEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendEnvironmentArgs)(nil)).Elem()
}

type BackendEnvironmentInput interface {
	pulumi.Input

	ToBackendEnvironmentOutput() BackendEnvironmentOutput
	ToBackendEnvironmentOutputWithContext(ctx context.Context) BackendEnvironmentOutput
}

func (*BackendEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendEnvironment)(nil))
}

func (i *BackendEnvironment) ToBackendEnvironmentOutput() BackendEnvironmentOutput {
	return i.ToBackendEnvironmentOutputWithContext(context.Background())
}

func (i *BackendEnvironment) ToBackendEnvironmentOutputWithContext(ctx context.Context) BackendEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendEnvironmentOutput)
}

func (i *BackendEnvironment) ToBackendEnvironmentPtrOutput() BackendEnvironmentPtrOutput {
	return i.ToBackendEnvironmentPtrOutputWithContext(context.Background())
}

func (i *BackendEnvironment) ToBackendEnvironmentPtrOutputWithContext(ctx context.Context) BackendEnvironmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendEnvironmentPtrOutput)
}

type BackendEnvironmentPtrInput interface {
	pulumi.Input

	ToBackendEnvironmentPtrOutput() BackendEnvironmentPtrOutput
	ToBackendEnvironmentPtrOutputWithContext(ctx context.Context) BackendEnvironmentPtrOutput
}

type backendEnvironmentPtrType BackendEnvironmentArgs

func (*backendEnvironmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendEnvironment)(nil))
}

func (i *backendEnvironmentPtrType) ToBackendEnvironmentPtrOutput() BackendEnvironmentPtrOutput {
	return i.ToBackendEnvironmentPtrOutputWithContext(context.Background())
}

func (i *backendEnvironmentPtrType) ToBackendEnvironmentPtrOutputWithContext(ctx context.Context) BackendEnvironmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendEnvironmentPtrOutput)
}

// BackendEnvironmentArrayInput is an input type that accepts BackendEnvironmentArray and BackendEnvironmentArrayOutput values.
// You can construct a concrete instance of `BackendEnvironmentArrayInput` via:
//
//          BackendEnvironmentArray{ BackendEnvironmentArgs{...} }
type BackendEnvironmentArrayInput interface {
	pulumi.Input

	ToBackendEnvironmentArrayOutput() BackendEnvironmentArrayOutput
	ToBackendEnvironmentArrayOutputWithContext(context.Context) BackendEnvironmentArrayOutput
}

type BackendEnvironmentArray []BackendEnvironmentInput

func (BackendEnvironmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendEnvironment)(nil)).Elem()
}

func (i BackendEnvironmentArray) ToBackendEnvironmentArrayOutput() BackendEnvironmentArrayOutput {
	return i.ToBackendEnvironmentArrayOutputWithContext(context.Background())
}

func (i BackendEnvironmentArray) ToBackendEnvironmentArrayOutputWithContext(ctx context.Context) BackendEnvironmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendEnvironmentArrayOutput)
}

// BackendEnvironmentMapInput is an input type that accepts BackendEnvironmentMap and BackendEnvironmentMapOutput values.
// You can construct a concrete instance of `BackendEnvironmentMapInput` via:
//
//          BackendEnvironmentMap{ "key": BackendEnvironmentArgs{...} }
type BackendEnvironmentMapInput interface {
	pulumi.Input

	ToBackendEnvironmentMapOutput() BackendEnvironmentMapOutput
	ToBackendEnvironmentMapOutputWithContext(context.Context) BackendEnvironmentMapOutput
}

type BackendEnvironmentMap map[string]BackendEnvironmentInput

func (BackendEnvironmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendEnvironment)(nil)).Elem()
}

func (i BackendEnvironmentMap) ToBackendEnvironmentMapOutput() BackendEnvironmentMapOutput {
	return i.ToBackendEnvironmentMapOutputWithContext(context.Background())
}

func (i BackendEnvironmentMap) ToBackendEnvironmentMapOutputWithContext(ctx context.Context) BackendEnvironmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendEnvironmentMapOutput)
}

type BackendEnvironmentOutput struct{ *pulumi.OutputState }

func (BackendEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendEnvironment)(nil))
}

func (o BackendEnvironmentOutput) ToBackendEnvironmentOutput() BackendEnvironmentOutput {
	return o
}

func (o BackendEnvironmentOutput) ToBackendEnvironmentOutputWithContext(ctx context.Context) BackendEnvironmentOutput {
	return o
}

func (o BackendEnvironmentOutput) ToBackendEnvironmentPtrOutput() BackendEnvironmentPtrOutput {
	return o.ToBackendEnvironmentPtrOutputWithContext(context.Background())
}

func (o BackendEnvironmentOutput) ToBackendEnvironmentPtrOutputWithContext(ctx context.Context) BackendEnvironmentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendEnvironment) *BackendEnvironment {
		return &v
	}).(BackendEnvironmentPtrOutput)
}

type BackendEnvironmentPtrOutput struct{ *pulumi.OutputState }

func (BackendEnvironmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendEnvironment)(nil))
}

func (o BackendEnvironmentPtrOutput) ToBackendEnvironmentPtrOutput() BackendEnvironmentPtrOutput {
	return o
}

func (o BackendEnvironmentPtrOutput) ToBackendEnvironmentPtrOutputWithContext(ctx context.Context) BackendEnvironmentPtrOutput {
	return o
}

func (o BackendEnvironmentPtrOutput) Elem() BackendEnvironmentOutput {
	return o.ApplyT(func(v *BackendEnvironment) BackendEnvironment {
		if v != nil {
			return *v
		}
		var ret BackendEnvironment
		return ret
	}).(BackendEnvironmentOutput)
}

type BackendEnvironmentArrayOutput struct{ *pulumi.OutputState }

func (BackendEnvironmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendEnvironment)(nil))
}

func (o BackendEnvironmentArrayOutput) ToBackendEnvironmentArrayOutput() BackendEnvironmentArrayOutput {
	return o
}

func (o BackendEnvironmentArrayOutput) ToBackendEnvironmentArrayOutputWithContext(ctx context.Context) BackendEnvironmentArrayOutput {
	return o
}

func (o BackendEnvironmentArrayOutput) Index(i pulumi.IntInput) BackendEnvironmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackendEnvironment {
		return vs[0].([]BackendEnvironment)[vs[1].(int)]
	}).(BackendEnvironmentOutput)
}

type BackendEnvironmentMapOutput struct{ *pulumi.OutputState }

func (BackendEnvironmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BackendEnvironment)(nil))
}

func (o BackendEnvironmentMapOutput) ToBackendEnvironmentMapOutput() BackendEnvironmentMapOutput {
	return o
}

func (o BackendEnvironmentMapOutput) ToBackendEnvironmentMapOutputWithContext(ctx context.Context) BackendEnvironmentMapOutput {
	return o
}

func (o BackendEnvironmentMapOutput) MapIndex(k pulumi.StringInput) BackendEnvironmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BackendEnvironment {
		return vs[0].(map[string]BackendEnvironment)[vs[1].(string)]
	}).(BackendEnvironmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackendEnvironmentInput)(nil)).Elem(), &BackendEnvironment{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendEnvironmentPtrInput)(nil)).Elem(), &BackendEnvironment{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendEnvironmentArrayInput)(nil)).Elem(), BackendEnvironmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendEnvironmentMapInput)(nil)).Elem(), BackendEnvironmentMap{})
	pulumi.RegisterOutputType(BackendEnvironmentOutput{})
	pulumi.RegisterOutputType(BackendEnvironmentPtrOutput{})
	pulumi.RegisterOutputType(BackendEnvironmentArrayOutput{})
	pulumi.RegisterOutputType(BackendEnvironmentMapOutput{})
}
