// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an OpsWorks Ruby on Rails application layer resource.
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
// 		_, err := opsworks.NewRailsAppLayer(ctx, "app", &opsworks.RailsAppLayerArgs{
// 			StackId: pulumi.Any(aws_opsworks_stack.Main.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RailsAppLayer struct {
	pulumi.CustomResourceState

	// Keyword for the app server to use. Defaults to "apachePassenger".
	AppServer pulumi.StringPtrOutput `pulumi:"appServer"`
	// The Amazon Resource Name(ARN) of the layer.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrOutput `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrOutput `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing pulumi.BoolPtrOutput `pulumi:"autoHealing"`
	// When OpsWorks is managing Bundler, which version to use. Defaults to "1.5.3".
	BundlerVersion         pulumi.StringPtrOutput   `pulumi:"bundlerVersion"`
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
	EbsVolumes RailsAppLayerEbsVolumeArrayOutput `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrOutput `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrOutput `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrOutput `pulumi:"instanceShutdownTimeout"`
	// Whether OpsWorks should manage bundler. On by default.
	ManageBundler pulumi.BoolPtrOutput `pulumi:"manageBundler"`
	// A human-readable name for the layer.
	Name pulumi.StringOutput `pulumi:"name"`
	// The version of Passenger to use. Defaults to "4.0.46".
	PassengerVersion pulumi.StringPtrOutput `pulumi:"passengerVersion"`
	// The version of Ruby to use. Defaults to "2.0.0".
	RubyVersion pulumi.StringPtrOutput `pulumi:"rubyVersion"`
	// The version of RubyGems to use. Defaults to "2.2.2".
	RubygemsVersion pulumi.StringPtrOutput `pulumi:"rubygemsVersion"`
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

// NewRailsAppLayer registers a new resource with the given unique name, arguments, and options.
func NewRailsAppLayer(ctx *pulumi.Context,
	name string, args *RailsAppLayerArgs, opts ...pulumi.ResourceOption) (*RailsAppLayer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	var resource RailsAppLayer
	err := ctx.RegisterResource("aws:opsworks/railsAppLayer:RailsAppLayer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRailsAppLayer gets an existing RailsAppLayer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRailsAppLayer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RailsAppLayerState, opts ...pulumi.ResourceOption) (*RailsAppLayer, error) {
	var resource RailsAppLayer
	err := ctx.ReadResource("aws:opsworks/railsAppLayer:RailsAppLayer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RailsAppLayer resources.
type railsAppLayerState struct {
	// Keyword for the app server to use. Defaults to "apachePassenger".
	AppServer *string `pulumi:"appServer"`
	// The Amazon Resource Name(ARN) of the layer.
	Arn *string `pulumi:"arn"`
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps *bool `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps *bool `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing *bool `pulumi:"autoHealing"`
	// When OpsWorks is managing Bundler, which version to use. Defaults to "1.5.3".
	BundlerVersion         *string  `pulumi:"bundlerVersion"`
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
	EbsVolumes []RailsAppLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// Whether OpsWorks should manage bundler. On by default.
	ManageBundler *bool `pulumi:"manageBundler"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
	// The version of Passenger to use. Defaults to "4.0.46".
	PassengerVersion *string `pulumi:"passengerVersion"`
	// The version of Ruby to use. Defaults to "2.0.0".
	RubyVersion *string `pulumi:"rubyVersion"`
	// The version of RubyGems to use. Defaults to "2.2.2".
	RubygemsVersion *string `pulumi:"rubygemsVersion"`
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

type RailsAppLayerState struct {
	// Keyword for the app server to use. Defaults to "apachePassenger".
	AppServer pulumi.StringPtrInput
	// The Amazon Resource Name(ARN) of the layer.
	Arn pulumi.StringPtrInput
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrInput
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrInput
	// Whether to enable auto-healing for the layer.
	AutoHealing pulumi.BoolPtrInput
	// When OpsWorks is managing Bundler, which version to use. Defaults to "1.5.3".
	BundlerVersion         pulumi.StringPtrInput
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
	EbsVolumes RailsAppLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// Whether OpsWorks should manage bundler. On by default.
	ManageBundler pulumi.BoolPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
	// The version of Passenger to use. Defaults to "4.0.46".
	PassengerVersion pulumi.StringPtrInput
	// The version of Ruby to use. Defaults to "2.0.0".
	RubyVersion pulumi.StringPtrInput
	// The version of RubyGems to use. Defaults to "2.2.2".
	RubygemsVersion pulumi.StringPtrInput
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

func (RailsAppLayerState) ElementType() reflect.Type {
	return reflect.TypeOf((*railsAppLayerState)(nil)).Elem()
}

type railsAppLayerArgs struct {
	// Keyword for the app server to use. Defaults to "apachePassenger".
	AppServer *string `pulumi:"appServer"`
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps *bool `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps *bool `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing *bool `pulumi:"autoHealing"`
	// When OpsWorks is managing Bundler, which version to use. Defaults to "1.5.3".
	BundlerVersion         *string  `pulumi:"bundlerVersion"`
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
	EbsVolumes []RailsAppLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// Whether OpsWorks should manage bundler. On by default.
	ManageBundler *bool `pulumi:"manageBundler"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
	// The version of Passenger to use. Defaults to "4.0.46".
	PassengerVersion *string `pulumi:"passengerVersion"`
	// The version of Ruby to use. Defaults to "2.0.0".
	RubyVersion *string `pulumi:"rubyVersion"`
	// The version of RubyGems to use. Defaults to "2.2.2".
	RubygemsVersion *string `pulumi:"rubygemsVersion"`
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

// The set of arguments for constructing a RailsAppLayer resource.
type RailsAppLayerArgs struct {
	// Keyword for the app server to use. Defaults to "apachePassenger".
	AppServer pulumi.StringPtrInput
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrInput
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrInput
	// Whether to enable auto-healing for the layer.
	AutoHealing pulumi.BoolPtrInput
	// When OpsWorks is managing Bundler, which version to use. Defaults to "1.5.3".
	BundlerVersion         pulumi.StringPtrInput
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
	EbsVolumes RailsAppLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// Whether OpsWorks should manage bundler. On by default.
	ManageBundler pulumi.BoolPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
	// The version of Passenger to use. Defaults to "4.0.46".
	PassengerVersion pulumi.StringPtrInput
	// The version of Ruby to use. Defaults to "2.0.0".
	RubyVersion pulumi.StringPtrInput
	// The version of RubyGems to use. Defaults to "2.2.2".
	RubygemsVersion pulumi.StringPtrInput
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

func (RailsAppLayerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*railsAppLayerArgs)(nil)).Elem()
}

type RailsAppLayerInput interface {
	pulumi.Input

	ToRailsAppLayerOutput() RailsAppLayerOutput
	ToRailsAppLayerOutputWithContext(ctx context.Context) RailsAppLayerOutput
}

func (*RailsAppLayer) ElementType() reflect.Type {
	return reflect.TypeOf((*RailsAppLayer)(nil))
}

func (i *RailsAppLayer) ToRailsAppLayerOutput() RailsAppLayerOutput {
	return i.ToRailsAppLayerOutputWithContext(context.Background())
}

func (i *RailsAppLayer) ToRailsAppLayerOutputWithContext(ctx context.Context) RailsAppLayerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RailsAppLayerOutput)
}

func (i *RailsAppLayer) ToRailsAppLayerPtrOutput() RailsAppLayerPtrOutput {
	return i.ToRailsAppLayerPtrOutputWithContext(context.Background())
}

func (i *RailsAppLayer) ToRailsAppLayerPtrOutputWithContext(ctx context.Context) RailsAppLayerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RailsAppLayerPtrOutput)
}

type RailsAppLayerPtrInput interface {
	pulumi.Input

	ToRailsAppLayerPtrOutput() RailsAppLayerPtrOutput
	ToRailsAppLayerPtrOutputWithContext(ctx context.Context) RailsAppLayerPtrOutput
}

type railsAppLayerPtrType RailsAppLayerArgs

func (*railsAppLayerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RailsAppLayer)(nil))
}

func (i *railsAppLayerPtrType) ToRailsAppLayerPtrOutput() RailsAppLayerPtrOutput {
	return i.ToRailsAppLayerPtrOutputWithContext(context.Background())
}

func (i *railsAppLayerPtrType) ToRailsAppLayerPtrOutputWithContext(ctx context.Context) RailsAppLayerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RailsAppLayerPtrOutput)
}

// RailsAppLayerArrayInput is an input type that accepts RailsAppLayerArray and RailsAppLayerArrayOutput values.
// You can construct a concrete instance of `RailsAppLayerArrayInput` via:
//
//          RailsAppLayerArray{ RailsAppLayerArgs{...} }
type RailsAppLayerArrayInput interface {
	pulumi.Input

	ToRailsAppLayerArrayOutput() RailsAppLayerArrayOutput
	ToRailsAppLayerArrayOutputWithContext(context.Context) RailsAppLayerArrayOutput
}

type RailsAppLayerArray []RailsAppLayerInput

func (RailsAppLayerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RailsAppLayer)(nil)).Elem()
}

func (i RailsAppLayerArray) ToRailsAppLayerArrayOutput() RailsAppLayerArrayOutput {
	return i.ToRailsAppLayerArrayOutputWithContext(context.Background())
}

func (i RailsAppLayerArray) ToRailsAppLayerArrayOutputWithContext(ctx context.Context) RailsAppLayerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RailsAppLayerArrayOutput)
}

// RailsAppLayerMapInput is an input type that accepts RailsAppLayerMap and RailsAppLayerMapOutput values.
// You can construct a concrete instance of `RailsAppLayerMapInput` via:
//
//          RailsAppLayerMap{ "key": RailsAppLayerArgs{...} }
type RailsAppLayerMapInput interface {
	pulumi.Input

	ToRailsAppLayerMapOutput() RailsAppLayerMapOutput
	ToRailsAppLayerMapOutputWithContext(context.Context) RailsAppLayerMapOutput
}

type RailsAppLayerMap map[string]RailsAppLayerInput

func (RailsAppLayerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RailsAppLayer)(nil)).Elem()
}

func (i RailsAppLayerMap) ToRailsAppLayerMapOutput() RailsAppLayerMapOutput {
	return i.ToRailsAppLayerMapOutputWithContext(context.Background())
}

func (i RailsAppLayerMap) ToRailsAppLayerMapOutputWithContext(ctx context.Context) RailsAppLayerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RailsAppLayerMapOutput)
}

type RailsAppLayerOutput struct{ *pulumi.OutputState }

func (RailsAppLayerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RailsAppLayer)(nil))
}

func (o RailsAppLayerOutput) ToRailsAppLayerOutput() RailsAppLayerOutput {
	return o
}

func (o RailsAppLayerOutput) ToRailsAppLayerOutputWithContext(ctx context.Context) RailsAppLayerOutput {
	return o
}

func (o RailsAppLayerOutput) ToRailsAppLayerPtrOutput() RailsAppLayerPtrOutput {
	return o.ToRailsAppLayerPtrOutputWithContext(context.Background())
}

func (o RailsAppLayerOutput) ToRailsAppLayerPtrOutputWithContext(ctx context.Context) RailsAppLayerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RailsAppLayer) *RailsAppLayer {
		return &v
	}).(RailsAppLayerPtrOutput)
}

type RailsAppLayerPtrOutput struct{ *pulumi.OutputState }

func (RailsAppLayerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RailsAppLayer)(nil))
}

func (o RailsAppLayerPtrOutput) ToRailsAppLayerPtrOutput() RailsAppLayerPtrOutput {
	return o
}

func (o RailsAppLayerPtrOutput) ToRailsAppLayerPtrOutputWithContext(ctx context.Context) RailsAppLayerPtrOutput {
	return o
}

func (o RailsAppLayerPtrOutput) Elem() RailsAppLayerOutput {
	return o.ApplyT(func(v *RailsAppLayer) RailsAppLayer {
		if v != nil {
			return *v
		}
		var ret RailsAppLayer
		return ret
	}).(RailsAppLayerOutput)
}

type RailsAppLayerArrayOutput struct{ *pulumi.OutputState }

func (RailsAppLayerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RailsAppLayer)(nil))
}

func (o RailsAppLayerArrayOutput) ToRailsAppLayerArrayOutput() RailsAppLayerArrayOutput {
	return o
}

func (o RailsAppLayerArrayOutput) ToRailsAppLayerArrayOutputWithContext(ctx context.Context) RailsAppLayerArrayOutput {
	return o
}

func (o RailsAppLayerArrayOutput) Index(i pulumi.IntInput) RailsAppLayerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RailsAppLayer {
		return vs[0].([]RailsAppLayer)[vs[1].(int)]
	}).(RailsAppLayerOutput)
}

type RailsAppLayerMapOutput struct{ *pulumi.OutputState }

func (RailsAppLayerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RailsAppLayer)(nil))
}

func (o RailsAppLayerMapOutput) ToRailsAppLayerMapOutput() RailsAppLayerMapOutput {
	return o
}

func (o RailsAppLayerMapOutput) ToRailsAppLayerMapOutputWithContext(ctx context.Context) RailsAppLayerMapOutput {
	return o
}

func (o RailsAppLayerMapOutput) MapIndex(k pulumi.StringInput) RailsAppLayerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RailsAppLayer {
		return vs[0].(map[string]RailsAppLayer)[vs[1].(string)]
	}).(RailsAppLayerOutput)
}

func init() {
	pulumi.RegisterOutputType(RailsAppLayerOutput{})
	pulumi.RegisterOutputType(RailsAppLayerPtrOutput{})
	pulumi.RegisterOutputType(RailsAppLayerArrayOutput{})
	pulumi.RegisterOutputType(RailsAppLayerMapOutput{})
}
