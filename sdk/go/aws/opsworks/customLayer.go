// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an OpsWorks custom layer resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/opsworks"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := opsworks.NewCustomLayer(ctx, "custlayer", &opsworks.CustomLayerArgs{
// 			ShortName: pulumi.String("awesome"),
// 			StackId:   pulumi.Any(aws_opsworks_stack.Main.Id),
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
// OpsWorks Custom Layers can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:opsworks/customLayer:CustomLayer bar 00000000-0000-0000-0000-000000000000
// ```
type CustomLayer struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name(ARN) of the layer.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrOutput `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrOutput `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing            pulumi.BoolPtrOutput     `pulumi:"autoHealing"`
	CustomConfigureRecipes pulumi.StringArrayOutput `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes    pulumi.StringArrayOutput `pulumi:"customDeployRecipes"`
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn pulumi.StringPtrOutput `pulumi:"customInstanceProfileArn"`
	// Custom JSON attributes to apply to the layer.
	CustomJson pulumi.StringPtrOutput `pulumi:"customJson"`
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.StringArrayOutput `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes     pulumi.StringArrayOutput `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes  pulumi.StringArrayOutput `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes  pulumi.StringArrayOutput `pulumi:"customUndeployRecipes"`
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolPtrOutput `pulumi:"drainElbOnShutdown"`
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes CustomLayerEbsVolumeArrayOutput `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrOutput `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrOutput `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrOutput `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name pulumi.StringOutput `pulumi:"name"`
	// A short, machine-readable name for the layer, which will be used to identify it in the Chef node JSON.
	ShortName pulumi.StringOutput `pulumi:"shortName"`
	// The id of the stack the layer will belong to.
	StackId pulumi.StringOutput `pulumi:"stackId"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayOutput `pulumi:"systemPackages"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrOutput `pulumi:"useEbsOptimizedInstances"`
}

// NewCustomLayer registers a new resource with the given unique name, arguments, and options.
func NewCustomLayer(ctx *pulumi.Context,
	name string, args *CustomLayerArgs, opts ...pulumi.ResourceOption) (*CustomLayer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ShortName == nil {
		return nil, errors.New("invalid value for required argument 'ShortName'")
	}
	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	var resource CustomLayer
	err := ctx.RegisterResource("aws:opsworks/customLayer:CustomLayer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomLayer gets an existing CustomLayer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomLayer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomLayerState, opts ...pulumi.ResourceOption) (*CustomLayer, error) {
	var resource CustomLayer
	err := ctx.ReadResource("aws:opsworks/customLayer:CustomLayer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomLayer resources.
type customLayerState struct {
	// The Amazon Resource Name(ARN) of the layer.
	Arn *string `pulumi:"arn"`
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps *bool `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps *bool `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing            *bool    `pulumi:"autoHealing"`
	CustomConfigureRecipes []string `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes    []string `pulumi:"customDeployRecipes"`
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn *string `pulumi:"customInstanceProfileArn"`
	// Custom JSON attributes to apply to the layer.
	CustomJson *string `pulumi:"customJson"`
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds []string `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes     []string `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes  []string `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes  []string `pulumi:"customUndeployRecipes"`
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown *bool `pulumi:"drainElbOnShutdown"`
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes []CustomLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
	// A short, machine-readable name for the layer, which will be used to identify it in the Chef node JSON.
	ShortName *string `pulumi:"shortName"`
	// The id of the stack the layer will belong to.
	StackId *string `pulumi:"stackId"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []string `pulumi:"systemPackages"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances *bool `pulumi:"useEbsOptimizedInstances"`
}

type CustomLayerState struct {
	// The Amazon Resource Name(ARN) of the layer.
	Arn pulumi.StringPtrInput
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrInput
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrInput
	// Whether to enable auto-healing for the layer.
	AutoHealing            pulumi.BoolPtrInput
	CustomConfigureRecipes pulumi.StringArrayInput
	CustomDeployRecipes    pulumi.StringArrayInput
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn pulumi.StringPtrInput
	// Custom JSON attributes to apply to the layer.
	CustomJson pulumi.StringPtrInput
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.StringArrayInput
	CustomSetupRecipes     pulumi.StringArrayInput
	CustomShutdownRecipes  pulumi.StringArrayInput
	CustomUndeployRecipes  pulumi.StringArrayInput
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolPtrInput
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes CustomLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
	// A short, machine-readable name for the layer, which will be used to identify it in the Chef node JSON.
	ShortName pulumi.StringPtrInput
	// The id of the stack the layer will belong to.
	StackId pulumi.StringPtrInput
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrInput
}

func (CustomLayerState) ElementType() reflect.Type {
	return reflect.TypeOf((*customLayerState)(nil)).Elem()
}

type customLayerArgs struct {
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps *bool `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps *bool `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing            *bool    `pulumi:"autoHealing"`
	CustomConfigureRecipes []string `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes    []string `pulumi:"customDeployRecipes"`
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn *string `pulumi:"customInstanceProfileArn"`
	// Custom JSON attributes to apply to the layer.
	CustomJson *string `pulumi:"customJson"`
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds []string `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes     []string `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes  []string `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes  []string `pulumi:"customUndeployRecipes"`
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown *bool `pulumi:"drainElbOnShutdown"`
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes []CustomLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
	// A short, machine-readable name for the layer, which will be used to identify it in the Chef node JSON.
	ShortName string `pulumi:"shortName"`
	// The id of the stack the layer will belong to.
	StackId string `pulumi:"stackId"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []string `pulumi:"systemPackages"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances *bool `pulumi:"useEbsOptimizedInstances"`
}

// The set of arguments for constructing a CustomLayer resource.
type CustomLayerArgs struct {
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrInput
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrInput
	// Whether to enable auto-healing for the layer.
	AutoHealing            pulumi.BoolPtrInput
	CustomConfigureRecipes pulumi.StringArrayInput
	CustomDeployRecipes    pulumi.StringArrayInput
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn pulumi.StringPtrInput
	// Custom JSON attributes to apply to the layer.
	CustomJson pulumi.StringPtrInput
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.StringArrayInput
	CustomSetupRecipes     pulumi.StringArrayInput
	CustomShutdownRecipes  pulumi.StringArrayInput
	CustomUndeployRecipes  pulumi.StringArrayInput
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolPtrInput
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes CustomLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
	// A short, machine-readable name for the layer, which will be used to identify it in the Chef node JSON.
	ShortName pulumi.StringInput
	// The id of the stack the layer will belong to.
	StackId pulumi.StringInput
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrInput
}

func (CustomLayerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customLayerArgs)(nil)).Elem()
}

type CustomLayerInput interface {
	pulumi.Input

	ToCustomLayerOutput() CustomLayerOutput
	ToCustomLayerOutputWithContext(ctx context.Context) CustomLayerOutput
}

func (*CustomLayer) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLayer)(nil))
}

func (i *CustomLayer) ToCustomLayerOutput() CustomLayerOutput {
	return i.ToCustomLayerOutputWithContext(context.Background())
}

func (i *CustomLayer) ToCustomLayerOutputWithContext(ctx context.Context) CustomLayerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLayerOutput)
}

func (i *CustomLayer) ToCustomLayerPtrOutput() CustomLayerPtrOutput {
	return i.ToCustomLayerPtrOutputWithContext(context.Background())
}

func (i *CustomLayer) ToCustomLayerPtrOutputWithContext(ctx context.Context) CustomLayerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLayerPtrOutput)
}

type CustomLayerPtrInput interface {
	pulumi.Input

	ToCustomLayerPtrOutput() CustomLayerPtrOutput
	ToCustomLayerPtrOutputWithContext(ctx context.Context) CustomLayerPtrOutput
}

type customLayerPtrType CustomLayerArgs

func (*customLayerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLayer)(nil))
}

func (i *customLayerPtrType) ToCustomLayerPtrOutput() CustomLayerPtrOutput {
	return i.ToCustomLayerPtrOutputWithContext(context.Background())
}

func (i *customLayerPtrType) ToCustomLayerPtrOutputWithContext(ctx context.Context) CustomLayerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLayerPtrOutput)
}

// CustomLayerArrayInput is an input type that accepts CustomLayerArray and CustomLayerArrayOutput values.
// You can construct a concrete instance of `CustomLayerArrayInput` via:
//
//          CustomLayerArray{ CustomLayerArgs{...} }
type CustomLayerArrayInput interface {
	pulumi.Input

	ToCustomLayerArrayOutput() CustomLayerArrayOutput
	ToCustomLayerArrayOutputWithContext(context.Context) CustomLayerArrayOutput
}

type CustomLayerArray []CustomLayerInput

func (CustomLayerArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*CustomLayer)(nil))
}

func (i CustomLayerArray) ToCustomLayerArrayOutput() CustomLayerArrayOutput {
	return i.ToCustomLayerArrayOutputWithContext(context.Background())
}

func (i CustomLayerArray) ToCustomLayerArrayOutputWithContext(ctx context.Context) CustomLayerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLayerArrayOutput)
}

// CustomLayerMapInput is an input type that accepts CustomLayerMap and CustomLayerMapOutput values.
// You can construct a concrete instance of `CustomLayerMapInput` via:
//
//          CustomLayerMap{ "key": CustomLayerArgs{...} }
type CustomLayerMapInput interface {
	pulumi.Input

	ToCustomLayerMapOutput() CustomLayerMapOutput
	ToCustomLayerMapOutputWithContext(context.Context) CustomLayerMapOutput
}

type CustomLayerMap map[string]CustomLayerInput

func (CustomLayerMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*CustomLayer)(nil))
}

func (i CustomLayerMap) ToCustomLayerMapOutput() CustomLayerMapOutput {
	return i.ToCustomLayerMapOutputWithContext(context.Background())
}

func (i CustomLayerMap) ToCustomLayerMapOutputWithContext(ctx context.Context) CustomLayerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLayerMapOutput)
}

type CustomLayerOutput struct {
	*pulumi.OutputState
}

func (CustomLayerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLayer)(nil))
}

func (o CustomLayerOutput) ToCustomLayerOutput() CustomLayerOutput {
	return o
}

func (o CustomLayerOutput) ToCustomLayerOutputWithContext(ctx context.Context) CustomLayerOutput {
	return o
}

func (o CustomLayerOutput) ToCustomLayerPtrOutput() CustomLayerPtrOutput {
	return o.ToCustomLayerPtrOutputWithContext(context.Background())
}

func (o CustomLayerOutput) ToCustomLayerPtrOutputWithContext(ctx context.Context) CustomLayerPtrOutput {
	return o.ApplyT(func(v CustomLayer) *CustomLayer {
		return &v
	}).(CustomLayerPtrOutput)
}

type CustomLayerPtrOutput struct {
	*pulumi.OutputState
}

func (CustomLayerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLayer)(nil))
}

func (o CustomLayerPtrOutput) ToCustomLayerPtrOutput() CustomLayerPtrOutput {
	return o
}

func (o CustomLayerPtrOutput) ToCustomLayerPtrOutputWithContext(ctx context.Context) CustomLayerPtrOutput {
	return o
}

type CustomLayerArrayOutput struct{ *pulumi.OutputState }

func (CustomLayerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomLayer)(nil))
}

func (o CustomLayerArrayOutput) ToCustomLayerArrayOutput() CustomLayerArrayOutput {
	return o
}

func (o CustomLayerArrayOutput) ToCustomLayerArrayOutputWithContext(ctx context.Context) CustomLayerArrayOutput {
	return o
}

func (o CustomLayerArrayOutput) Index(i pulumi.IntInput) CustomLayerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomLayer {
		return vs[0].([]CustomLayer)[vs[1].(int)]
	}).(CustomLayerOutput)
}

type CustomLayerMapOutput struct{ *pulumi.OutputState }

func (CustomLayerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomLayer)(nil))
}

func (o CustomLayerMapOutput) ToCustomLayerMapOutput() CustomLayerMapOutput {
	return o
}

func (o CustomLayerMapOutput) ToCustomLayerMapOutputWithContext(ctx context.Context) CustomLayerMapOutput {
	return o
}

func (o CustomLayerMapOutput) MapIndex(k pulumi.StringInput) CustomLayerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CustomLayer {
		return vs[0].(map[string]CustomLayer)[vs[1].(string)]
	}).(CustomLayerOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomLayerOutput{})
	pulumi.RegisterOutputType(CustomLayerPtrOutput{})
	pulumi.RegisterOutputType(CustomLayerArrayOutput{})
	pulumi.RegisterOutputType(CustomLayerMapOutput{})
}
