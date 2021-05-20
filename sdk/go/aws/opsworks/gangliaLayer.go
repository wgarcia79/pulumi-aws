// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an OpsWorks Ganglia layer resource.
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
// 		_, err := opsworks.NewGangliaLayer(ctx, "monitor", &opsworks.GangliaLayerArgs{
// 			StackId:  pulumi.Any(aws_opsworks_stack.Main.Id),
// 			Password: pulumi.String("foobarbaz"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type GangliaLayer struct {
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
	EbsVolumes GangliaLayerEbsVolumeArrayOutput `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrOutput `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrOutput `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrOutput `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password to use for Ganglia.
	Password pulumi.StringOutput `pulumi:"password"`
	// The id of the stack the layer will belong to.
	StackId pulumi.StringOutput `pulumi:"stackId"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayOutput `pulumi:"systemPackages"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The URL path to use for Ganglia. Defaults to "/ganglia".
	Url pulumi.StringPtrOutput `pulumi:"url"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrOutput `pulumi:"useEbsOptimizedInstances"`
	// The username to use for Ganglia. Defaults to "opsworks".
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewGangliaLayer registers a new resource with the given unique name, arguments, and options.
func NewGangliaLayer(ctx *pulumi.Context,
	name string, args *GangliaLayerArgs, opts ...pulumi.ResourceOption) (*GangliaLayer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	var resource GangliaLayer
	err := ctx.RegisterResource("aws:opsworks/gangliaLayer:GangliaLayer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGangliaLayer gets an existing GangliaLayer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGangliaLayer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GangliaLayerState, opts ...pulumi.ResourceOption) (*GangliaLayer, error) {
	var resource GangliaLayer
	err := ctx.ReadResource("aws:opsworks/gangliaLayer:GangliaLayer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GangliaLayer resources.
type gangliaLayerState struct {
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
	EbsVolumes []GangliaLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
	// The password to use for Ganglia.
	Password *string `pulumi:"password"`
	// The id of the stack the layer will belong to.
	StackId *string `pulumi:"stackId"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []string `pulumi:"systemPackages"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The URL path to use for Ganglia. Defaults to "/ganglia".
	Url *string `pulumi:"url"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances *bool `pulumi:"useEbsOptimizedInstances"`
	// The username to use for Ganglia. Defaults to "opsworks".
	Username *string `pulumi:"username"`
}

type GangliaLayerState struct {
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
	EbsVolumes GangliaLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
	// The password to use for Ganglia.
	Password pulumi.StringPtrInput
	// The id of the stack the layer will belong to.
	StackId pulumi.StringPtrInput
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// The URL path to use for Ganglia. Defaults to "/ganglia".
	Url pulumi.StringPtrInput
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrInput
	// The username to use for Ganglia. Defaults to "opsworks".
	Username pulumi.StringPtrInput
}

func (GangliaLayerState) ElementType() reflect.Type {
	return reflect.TypeOf((*gangliaLayerState)(nil)).Elem()
}

type gangliaLayerArgs struct {
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
	EbsVolumes []GangliaLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
	// The password to use for Ganglia.
	Password string `pulumi:"password"`
	// The id of the stack the layer will belong to.
	StackId string `pulumi:"stackId"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []string `pulumi:"systemPackages"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The URL path to use for Ganglia. Defaults to "/ganglia".
	Url *string `pulumi:"url"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances *bool `pulumi:"useEbsOptimizedInstances"`
	// The username to use for Ganglia. Defaults to "opsworks".
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a GangliaLayer resource.
type GangliaLayerArgs struct {
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
	EbsVolumes GangliaLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
	// The password to use for Ganglia.
	Password pulumi.StringInput
	// The id of the stack the layer will belong to.
	StackId pulumi.StringInput
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// The URL path to use for Ganglia. Defaults to "/ganglia".
	Url pulumi.StringPtrInput
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrInput
	// The username to use for Ganglia. Defaults to "opsworks".
	Username pulumi.StringPtrInput
}

func (GangliaLayerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gangliaLayerArgs)(nil)).Elem()
}

type GangliaLayerInput interface {
	pulumi.Input

	ToGangliaLayerOutput() GangliaLayerOutput
	ToGangliaLayerOutputWithContext(ctx context.Context) GangliaLayerOutput
}

func (*GangliaLayer) ElementType() reflect.Type {
	return reflect.TypeOf((*GangliaLayer)(nil))
}

func (i *GangliaLayer) ToGangliaLayerOutput() GangliaLayerOutput {
	return i.ToGangliaLayerOutputWithContext(context.Background())
}

func (i *GangliaLayer) ToGangliaLayerOutputWithContext(ctx context.Context) GangliaLayerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GangliaLayerOutput)
}

func (i *GangliaLayer) ToGangliaLayerPtrOutput() GangliaLayerPtrOutput {
	return i.ToGangliaLayerPtrOutputWithContext(context.Background())
}

func (i *GangliaLayer) ToGangliaLayerPtrOutputWithContext(ctx context.Context) GangliaLayerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GangliaLayerPtrOutput)
}

type GangliaLayerPtrInput interface {
	pulumi.Input

	ToGangliaLayerPtrOutput() GangliaLayerPtrOutput
	ToGangliaLayerPtrOutputWithContext(ctx context.Context) GangliaLayerPtrOutput
}

type gangliaLayerPtrType GangliaLayerArgs

func (*gangliaLayerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GangliaLayer)(nil))
}

func (i *gangliaLayerPtrType) ToGangliaLayerPtrOutput() GangliaLayerPtrOutput {
	return i.ToGangliaLayerPtrOutputWithContext(context.Background())
}

func (i *gangliaLayerPtrType) ToGangliaLayerPtrOutputWithContext(ctx context.Context) GangliaLayerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GangliaLayerPtrOutput)
}

// GangliaLayerArrayInput is an input type that accepts GangliaLayerArray and GangliaLayerArrayOutput values.
// You can construct a concrete instance of `GangliaLayerArrayInput` via:
//
//          GangliaLayerArray{ GangliaLayerArgs{...} }
type GangliaLayerArrayInput interface {
	pulumi.Input

	ToGangliaLayerArrayOutput() GangliaLayerArrayOutput
	ToGangliaLayerArrayOutputWithContext(context.Context) GangliaLayerArrayOutput
}

type GangliaLayerArray []GangliaLayerInput

func (GangliaLayerArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*GangliaLayer)(nil))
}

func (i GangliaLayerArray) ToGangliaLayerArrayOutput() GangliaLayerArrayOutput {
	return i.ToGangliaLayerArrayOutputWithContext(context.Background())
}

func (i GangliaLayerArray) ToGangliaLayerArrayOutputWithContext(ctx context.Context) GangliaLayerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GangliaLayerArrayOutput)
}

// GangliaLayerMapInput is an input type that accepts GangliaLayerMap and GangliaLayerMapOutput values.
// You can construct a concrete instance of `GangliaLayerMapInput` via:
//
//          GangliaLayerMap{ "key": GangliaLayerArgs{...} }
type GangliaLayerMapInput interface {
	pulumi.Input

	ToGangliaLayerMapOutput() GangliaLayerMapOutput
	ToGangliaLayerMapOutputWithContext(context.Context) GangliaLayerMapOutput
}

type GangliaLayerMap map[string]GangliaLayerInput

func (GangliaLayerMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*GangliaLayer)(nil))
}

func (i GangliaLayerMap) ToGangliaLayerMapOutput() GangliaLayerMapOutput {
	return i.ToGangliaLayerMapOutputWithContext(context.Background())
}

func (i GangliaLayerMap) ToGangliaLayerMapOutputWithContext(ctx context.Context) GangliaLayerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GangliaLayerMapOutput)
}

type GangliaLayerOutput struct {
	*pulumi.OutputState
}

func (GangliaLayerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GangliaLayer)(nil))
}

func (o GangliaLayerOutput) ToGangliaLayerOutput() GangliaLayerOutput {
	return o
}

func (o GangliaLayerOutput) ToGangliaLayerOutputWithContext(ctx context.Context) GangliaLayerOutput {
	return o
}

func (o GangliaLayerOutput) ToGangliaLayerPtrOutput() GangliaLayerPtrOutput {
	return o.ToGangliaLayerPtrOutputWithContext(context.Background())
}

func (o GangliaLayerOutput) ToGangliaLayerPtrOutputWithContext(ctx context.Context) GangliaLayerPtrOutput {
	return o.ApplyT(func(v GangliaLayer) *GangliaLayer {
		return &v
	}).(GangliaLayerPtrOutput)
}

type GangliaLayerPtrOutput struct {
	*pulumi.OutputState
}

func (GangliaLayerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GangliaLayer)(nil))
}

func (o GangliaLayerPtrOutput) ToGangliaLayerPtrOutput() GangliaLayerPtrOutput {
	return o
}

func (o GangliaLayerPtrOutput) ToGangliaLayerPtrOutputWithContext(ctx context.Context) GangliaLayerPtrOutput {
	return o
}

type GangliaLayerArrayOutput struct{ *pulumi.OutputState }

func (GangliaLayerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GangliaLayer)(nil))
}

func (o GangliaLayerArrayOutput) ToGangliaLayerArrayOutput() GangliaLayerArrayOutput {
	return o
}

func (o GangliaLayerArrayOutput) ToGangliaLayerArrayOutputWithContext(ctx context.Context) GangliaLayerArrayOutput {
	return o
}

func (o GangliaLayerArrayOutput) Index(i pulumi.IntInput) GangliaLayerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GangliaLayer {
		return vs[0].([]GangliaLayer)[vs[1].(int)]
	}).(GangliaLayerOutput)
}

type GangliaLayerMapOutput struct{ *pulumi.OutputState }

func (GangliaLayerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GangliaLayer)(nil))
}

func (o GangliaLayerMapOutput) ToGangliaLayerMapOutput() GangliaLayerMapOutput {
	return o
}

func (o GangliaLayerMapOutput) ToGangliaLayerMapOutputWithContext(ctx context.Context) GangliaLayerMapOutput {
	return o
}

func (o GangliaLayerMapOutput) MapIndex(k pulumi.StringInput) GangliaLayerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GangliaLayer {
		return vs[0].(map[string]GangliaLayer)[vs[1].(string)]
	}).(GangliaLayerOutput)
}

func init() {
	pulumi.RegisterOutputType(GangliaLayerOutput{})
	pulumi.RegisterOutputType(GangliaLayerPtrOutput{})
	pulumi.RegisterOutputType(GangliaLayerArrayOutput{})
	pulumi.RegisterOutputType(GangliaLayerMapOutput{})
}
