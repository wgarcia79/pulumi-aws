// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Amazon API Gateway Version 2 deployment.
// More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).
//
// > **Note:** Creating a deployment for an API requires at least one `apigatewayv2.Route` resource associated with that API. To avoid race conditions when all resources are being created together, you need to add implicit resource references via the `triggers` argument or explicit resource references using the [resource `dependsOn` meta-argument](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson).
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/apigatewayv2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apigatewayv2.NewDeployment(ctx, "example", &apigatewayv2.DeploymentArgs{
// 			ApiId:       pulumi.Any(aws_apigatewayv2_route.Example.Api_id),
// 			Description: pulumi.String("Example deployment"),
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
// `aws_apigatewayv2_deployment` can be imported by using the API identifier and deployment identifier, e.g.
//
// ```sh
//  $ pulumi import aws:apigatewayv2/deployment:Deployment example aabbccddee/1122334
// ```
//
//  The `triggers` argument cannot be imported.
type Deployment struct {
	pulumi.CustomResourceState

	// The API identifier.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// Whether the deployment was automatically released.
	AutoDeployed pulumi.BoolOutput `pulumi:"autoDeployed"`
	// The description for the deployment resource. Must be less than or equal to 1024 characters in length.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A map of arbitrary keys and values that, when changed, will trigger a redeployment.
	Triggers pulumi.StringMapOutput `pulumi:"triggers"`
}

// NewDeployment registers a new resource with the given unique name, arguments, and options.
func NewDeployment(ctx *pulumi.Context,
	name string, args *DeploymentArgs, opts ...pulumi.ResourceOption) (*Deployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	var resource Deployment
	err := ctx.RegisterResource("aws:apigatewayv2/deployment:Deployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployment gets an existing Deployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentState, opts ...pulumi.ResourceOption) (*Deployment, error) {
	var resource Deployment
	err := ctx.ReadResource("aws:apigatewayv2/deployment:Deployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Deployment resources.
type deploymentState struct {
	// The API identifier.
	ApiId *string `pulumi:"apiId"`
	// Whether the deployment was automatically released.
	AutoDeployed *bool `pulumi:"autoDeployed"`
	// The description for the deployment resource. Must be less than or equal to 1024 characters in length.
	Description *string `pulumi:"description"`
	// A map of arbitrary keys and values that, when changed, will trigger a redeployment.
	Triggers map[string]string `pulumi:"triggers"`
}

type DeploymentState struct {
	// The API identifier.
	ApiId pulumi.StringPtrInput
	// Whether the deployment was automatically released.
	AutoDeployed pulumi.BoolPtrInput
	// The description for the deployment resource. Must be less than or equal to 1024 characters in length.
	Description pulumi.StringPtrInput
	// A map of arbitrary keys and values that, when changed, will trigger a redeployment.
	Triggers pulumi.StringMapInput
}

func (DeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentState)(nil)).Elem()
}

type deploymentArgs struct {
	// The API identifier.
	ApiId string `pulumi:"apiId"`
	// The description for the deployment resource. Must be less than or equal to 1024 characters in length.
	Description *string `pulumi:"description"`
	// A map of arbitrary keys and values that, when changed, will trigger a redeployment.
	Triggers map[string]string `pulumi:"triggers"`
}

// The set of arguments for constructing a Deployment resource.
type DeploymentArgs struct {
	// The API identifier.
	ApiId pulumi.StringInput
	// The description for the deployment resource. Must be less than or equal to 1024 characters in length.
	Description pulumi.StringPtrInput
	// A map of arbitrary keys and values that, when changed, will trigger a redeployment.
	Triggers pulumi.StringMapInput
}

func (DeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentArgs)(nil)).Elem()
}

type DeploymentInput interface {
	pulumi.Input

	ToDeploymentOutput() DeploymentOutput
	ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput
}

func (*Deployment) ElementType() reflect.Type {
	return reflect.TypeOf((*Deployment)(nil))
}

func (i *Deployment) ToDeploymentOutput() DeploymentOutput {
	return i.ToDeploymentOutputWithContext(context.Background())
}

func (i *Deployment) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentOutput)
}

func (i *Deployment) ToDeploymentPtrOutput() DeploymentPtrOutput {
	return i.ToDeploymentPtrOutputWithContext(context.Background())
}

func (i *Deployment) ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPtrOutput)
}

type DeploymentPtrInput interface {
	pulumi.Input

	ToDeploymentPtrOutput() DeploymentPtrOutput
	ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput
}

type deploymentPtrType DeploymentArgs

func (*deploymentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil))
}

func (i *deploymentPtrType) ToDeploymentPtrOutput() DeploymentPtrOutput {
	return i.ToDeploymentPtrOutputWithContext(context.Background())
}

func (i *deploymentPtrType) ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPtrOutput)
}

// DeploymentArrayInput is an input type that accepts DeploymentArray and DeploymentArrayOutput values.
// You can construct a concrete instance of `DeploymentArrayInput` via:
//
//          DeploymentArray{ DeploymentArgs{...} }
type DeploymentArrayInput interface {
	pulumi.Input

	ToDeploymentArrayOutput() DeploymentArrayOutput
	ToDeploymentArrayOutputWithContext(context.Context) DeploymentArrayOutput
}

type DeploymentArray []DeploymentInput

func (DeploymentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Deployment)(nil)).Elem()
}

func (i DeploymentArray) ToDeploymentArrayOutput() DeploymentArrayOutput {
	return i.ToDeploymentArrayOutputWithContext(context.Background())
}

func (i DeploymentArray) ToDeploymentArrayOutputWithContext(ctx context.Context) DeploymentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentArrayOutput)
}

// DeploymentMapInput is an input type that accepts DeploymentMap and DeploymentMapOutput values.
// You can construct a concrete instance of `DeploymentMapInput` via:
//
//          DeploymentMap{ "key": DeploymentArgs{...} }
type DeploymentMapInput interface {
	pulumi.Input

	ToDeploymentMapOutput() DeploymentMapOutput
	ToDeploymentMapOutputWithContext(context.Context) DeploymentMapOutput
}

type DeploymentMap map[string]DeploymentInput

func (DeploymentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Deployment)(nil)).Elem()
}

func (i DeploymentMap) ToDeploymentMapOutput() DeploymentMapOutput {
	return i.ToDeploymentMapOutputWithContext(context.Background())
}

func (i DeploymentMap) ToDeploymentMapOutputWithContext(ctx context.Context) DeploymentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentMapOutput)
}

type DeploymentOutput struct{ *pulumi.OutputState }

func (DeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Deployment)(nil))
}

func (o DeploymentOutput) ToDeploymentOutput() DeploymentOutput {
	return o
}

func (o DeploymentOutput) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return o
}

func (o DeploymentOutput) ToDeploymentPtrOutput() DeploymentPtrOutput {
	return o.ToDeploymentPtrOutputWithContext(context.Background())
}

func (o DeploymentOutput) ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Deployment) *Deployment {
		return &v
	}).(DeploymentPtrOutput)
}

type DeploymentPtrOutput struct{ *pulumi.OutputState }

func (DeploymentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil))
}

func (o DeploymentPtrOutput) ToDeploymentPtrOutput() DeploymentPtrOutput {
	return o
}

func (o DeploymentPtrOutput) ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput {
	return o
}

func (o DeploymentPtrOutput) Elem() DeploymentOutput {
	return o.ApplyT(func(v *Deployment) Deployment {
		if v != nil {
			return *v
		}
		var ret Deployment
		return ret
	}).(DeploymentOutput)
}

type DeploymentArrayOutput struct{ *pulumi.OutputState }

func (DeploymentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Deployment)(nil))
}

func (o DeploymentArrayOutput) ToDeploymentArrayOutput() DeploymentArrayOutput {
	return o
}

func (o DeploymentArrayOutput) ToDeploymentArrayOutputWithContext(ctx context.Context) DeploymentArrayOutput {
	return o
}

func (o DeploymentArrayOutput) Index(i pulumi.IntInput) DeploymentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Deployment {
		return vs[0].([]Deployment)[vs[1].(int)]
	}).(DeploymentOutput)
}

type DeploymentMapOutput struct{ *pulumi.OutputState }

func (DeploymentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Deployment)(nil))
}

func (o DeploymentMapOutput) ToDeploymentMapOutput() DeploymentMapOutput {
	return o
}

func (o DeploymentMapOutput) ToDeploymentMapOutputWithContext(ctx context.Context) DeploymentMapOutput {
	return o
}

func (o DeploymentMapOutput) MapIndex(k pulumi.StringInput) DeploymentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Deployment {
		return vs[0].(map[string]Deployment)[vs[1].(string)]
	}).(DeploymentOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentOutput{})
	pulumi.RegisterOutputType(DeploymentPtrOutput{})
	pulumi.RegisterOutputType(DeploymentArrayOutput{})
	pulumi.RegisterOutputType(DeploymentMapOutput{})
}
