// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an OpsWorks Java application layer resource.
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
// 		_, err := opsworks.NewJavaAppLayer(ctx, "app", &opsworks.JavaAppLayerArgs{
// 			StackId: pulumi.Any(aws_opsworks_stack.Main.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type JavaAppLayer struct {
	pulumi.CustomResourceState

	// Keyword for the application container to use. Defaults to "tomcat".
	AppServer pulumi.StringPtrOutput `pulumi:"appServer"`
	// Version of the selected application container to use. Defaults to "7".
	AppServerVersion pulumi.StringPtrOutput `pulumi:"appServerVersion"`
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
	EbsVolumes JavaAppLayerEbsVolumeArrayOutput `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrOutput `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrOutput `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrOutput `pulumi:"instanceShutdownTimeout"`
	// Options to set for the JVM.
	JvmOptions pulumi.StringPtrOutput `pulumi:"jvmOptions"`
	// Keyword for the type of JVM to use. Defaults to `openjdk`.
	JvmType pulumi.StringPtrOutput `pulumi:"jvmType"`
	// Version of JVM to use. Defaults to "7".
	JvmVersion pulumi.StringPtrOutput `pulumi:"jvmVersion"`
	// A human-readable name for the layer.
	Name pulumi.StringOutput `pulumi:"name"`
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

// NewJavaAppLayer registers a new resource with the given unique name, arguments, and options.
func NewJavaAppLayer(ctx *pulumi.Context,
	name string, args *JavaAppLayerArgs, opts ...pulumi.ResourceOption) (*JavaAppLayer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	var resource JavaAppLayer
	err := ctx.RegisterResource("aws:opsworks/javaAppLayer:JavaAppLayer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJavaAppLayer gets an existing JavaAppLayer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJavaAppLayer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JavaAppLayerState, opts ...pulumi.ResourceOption) (*JavaAppLayer, error) {
	var resource JavaAppLayer
	err := ctx.ReadResource("aws:opsworks/javaAppLayer:JavaAppLayer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JavaAppLayer resources.
type javaAppLayerState struct {
	// Keyword for the application container to use. Defaults to "tomcat".
	AppServer *string `pulumi:"appServer"`
	// Version of the selected application container to use. Defaults to "7".
	AppServerVersion *string `pulumi:"appServerVersion"`
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
	EbsVolumes []JavaAppLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// Options to set for the JVM.
	JvmOptions *string `pulumi:"jvmOptions"`
	// Keyword for the type of JVM to use. Defaults to `openjdk`.
	JvmType *string `pulumi:"jvmType"`
	// Version of JVM to use. Defaults to "7".
	JvmVersion *string `pulumi:"jvmVersion"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
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

type JavaAppLayerState struct {
	// Keyword for the application container to use. Defaults to "tomcat".
	AppServer pulumi.StringPtrInput
	// Version of the selected application container to use. Defaults to "7".
	AppServerVersion pulumi.StringPtrInput
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
	EbsVolumes JavaAppLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// Options to set for the JVM.
	JvmOptions pulumi.StringPtrInput
	// Keyword for the type of JVM to use. Defaults to `openjdk`.
	JvmType pulumi.StringPtrInput
	// Version of JVM to use. Defaults to "7".
	JvmVersion pulumi.StringPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
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

func (JavaAppLayerState) ElementType() reflect.Type {
	return reflect.TypeOf((*javaAppLayerState)(nil)).Elem()
}

type javaAppLayerArgs struct {
	// Keyword for the application container to use. Defaults to "tomcat".
	AppServer *string `pulumi:"appServer"`
	// Version of the selected application container to use. Defaults to "7".
	AppServerVersion *string `pulumi:"appServerVersion"`
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
	EbsVolumes []JavaAppLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// Options to set for the JVM.
	JvmOptions *string `pulumi:"jvmOptions"`
	// Keyword for the type of JVM to use. Defaults to `openjdk`.
	JvmType *string `pulumi:"jvmType"`
	// Version of JVM to use. Defaults to "7".
	JvmVersion *string `pulumi:"jvmVersion"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
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

// The set of arguments for constructing a JavaAppLayer resource.
type JavaAppLayerArgs struct {
	// Keyword for the application container to use. Defaults to "tomcat".
	AppServer pulumi.StringPtrInput
	// Version of the selected application container to use. Defaults to "7".
	AppServerVersion pulumi.StringPtrInput
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
	EbsVolumes JavaAppLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// Options to set for the JVM.
	JvmOptions pulumi.StringPtrInput
	// Keyword for the type of JVM to use. Defaults to `openjdk`.
	JvmType pulumi.StringPtrInput
	// Version of JVM to use. Defaults to "7".
	JvmVersion pulumi.StringPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
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

func (JavaAppLayerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*javaAppLayerArgs)(nil)).Elem()
}

type JavaAppLayerInput interface {
	pulumi.Input

	ToJavaAppLayerOutput() JavaAppLayerOutput
	ToJavaAppLayerOutputWithContext(ctx context.Context) JavaAppLayerOutput
}

func (*JavaAppLayer) ElementType() reflect.Type {
	return reflect.TypeOf((*JavaAppLayer)(nil))
}

func (i *JavaAppLayer) ToJavaAppLayerOutput() JavaAppLayerOutput {
	return i.ToJavaAppLayerOutputWithContext(context.Background())
}

func (i *JavaAppLayer) ToJavaAppLayerOutputWithContext(ctx context.Context) JavaAppLayerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JavaAppLayerOutput)
}

func (i *JavaAppLayer) ToJavaAppLayerPtrOutput() JavaAppLayerPtrOutput {
	return i.ToJavaAppLayerPtrOutputWithContext(context.Background())
}

func (i *JavaAppLayer) ToJavaAppLayerPtrOutputWithContext(ctx context.Context) JavaAppLayerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JavaAppLayerPtrOutput)
}

type JavaAppLayerPtrInput interface {
	pulumi.Input

	ToJavaAppLayerPtrOutput() JavaAppLayerPtrOutput
	ToJavaAppLayerPtrOutputWithContext(ctx context.Context) JavaAppLayerPtrOutput
}

type javaAppLayerPtrType JavaAppLayerArgs

func (*javaAppLayerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JavaAppLayer)(nil))
}

func (i *javaAppLayerPtrType) ToJavaAppLayerPtrOutput() JavaAppLayerPtrOutput {
	return i.ToJavaAppLayerPtrOutputWithContext(context.Background())
}

func (i *javaAppLayerPtrType) ToJavaAppLayerPtrOutputWithContext(ctx context.Context) JavaAppLayerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JavaAppLayerPtrOutput)
}

// JavaAppLayerArrayInput is an input type that accepts JavaAppLayerArray and JavaAppLayerArrayOutput values.
// You can construct a concrete instance of `JavaAppLayerArrayInput` via:
//
//          JavaAppLayerArray{ JavaAppLayerArgs{...} }
type JavaAppLayerArrayInput interface {
	pulumi.Input

	ToJavaAppLayerArrayOutput() JavaAppLayerArrayOutput
	ToJavaAppLayerArrayOutputWithContext(context.Context) JavaAppLayerArrayOutput
}

type JavaAppLayerArray []JavaAppLayerInput

func (JavaAppLayerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*JavaAppLayer)(nil)).Elem()
}

func (i JavaAppLayerArray) ToJavaAppLayerArrayOutput() JavaAppLayerArrayOutput {
	return i.ToJavaAppLayerArrayOutputWithContext(context.Background())
}

func (i JavaAppLayerArray) ToJavaAppLayerArrayOutputWithContext(ctx context.Context) JavaAppLayerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JavaAppLayerArrayOutput)
}

// JavaAppLayerMapInput is an input type that accepts JavaAppLayerMap and JavaAppLayerMapOutput values.
// You can construct a concrete instance of `JavaAppLayerMapInput` via:
//
//          JavaAppLayerMap{ "key": JavaAppLayerArgs{...} }
type JavaAppLayerMapInput interface {
	pulumi.Input

	ToJavaAppLayerMapOutput() JavaAppLayerMapOutput
	ToJavaAppLayerMapOutputWithContext(context.Context) JavaAppLayerMapOutput
}

type JavaAppLayerMap map[string]JavaAppLayerInput

func (JavaAppLayerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*JavaAppLayer)(nil)).Elem()
}

func (i JavaAppLayerMap) ToJavaAppLayerMapOutput() JavaAppLayerMapOutput {
	return i.ToJavaAppLayerMapOutputWithContext(context.Background())
}

func (i JavaAppLayerMap) ToJavaAppLayerMapOutputWithContext(ctx context.Context) JavaAppLayerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JavaAppLayerMapOutput)
}

type JavaAppLayerOutput struct{ *pulumi.OutputState }

func (JavaAppLayerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JavaAppLayer)(nil))
}

func (o JavaAppLayerOutput) ToJavaAppLayerOutput() JavaAppLayerOutput {
	return o
}

func (o JavaAppLayerOutput) ToJavaAppLayerOutputWithContext(ctx context.Context) JavaAppLayerOutput {
	return o
}

func (o JavaAppLayerOutput) ToJavaAppLayerPtrOutput() JavaAppLayerPtrOutput {
	return o.ToJavaAppLayerPtrOutputWithContext(context.Background())
}

func (o JavaAppLayerOutput) ToJavaAppLayerPtrOutputWithContext(ctx context.Context) JavaAppLayerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JavaAppLayer) *JavaAppLayer {
		return &v
	}).(JavaAppLayerPtrOutput)
}

type JavaAppLayerPtrOutput struct{ *pulumi.OutputState }

func (JavaAppLayerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JavaAppLayer)(nil))
}

func (o JavaAppLayerPtrOutput) ToJavaAppLayerPtrOutput() JavaAppLayerPtrOutput {
	return o
}

func (o JavaAppLayerPtrOutput) ToJavaAppLayerPtrOutputWithContext(ctx context.Context) JavaAppLayerPtrOutput {
	return o
}

func (o JavaAppLayerPtrOutput) Elem() JavaAppLayerOutput {
	return o.ApplyT(func(v *JavaAppLayer) JavaAppLayer {
		if v != nil {
			return *v
		}
		var ret JavaAppLayer
		return ret
	}).(JavaAppLayerOutput)
}

type JavaAppLayerArrayOutput struct{ *pulumi.OutputState }

func (JavaAppLayerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JavaAppLayer)(nil))
}

func (o JavaAppLayerArrayOutput) ToJavaAppLayerArrayOutput() JavaAppLayerArrayOutput {
	return o
}

func (o JavaAppLayerArrayOutput) ToJavaAppLayerArrayOutputWithContext(ctx context.Context) JavaAppLayerArrayOutput {
	return o
}

func (o JavaAppLayerArrayOutput) Index(i pulumi.IntInput) JavaAppLayerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JavaAppLayer {
		return vs[0].([]JavaAppLayer)[vs[1].(int)]
	}).(JavaAppLayerOutput)
}

type JavaAppLayerMapOutput struct{ *pulumi.OutputState }

func (JavaAppLayerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]JavaAppLayer)(nil))
}

func (o JavaAppLayerMapOutput) ToJavaAppLayerMapOutput() JavaAppLayerMapOutput {
	return o
}

func (o JavaAppLayerMapOutput) ToJavaAppLayerMapOutputWithContext(ctx context.Context) JavaAppLayerMapOutput {
	return o
}

func (o JavaAppLayerMapOutput) MapIndex(k pulumi.StringInput) JavaAppLayerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) JavaAppLayer {
		return vs[0].(map[string]JavaAppLayer)[vs[1].(string)]
	}).(JavaAppLayerOutput)
}

func init() {
	pulumi.RegisterOutputType(JavaAppLayerOutput{})
	pulumi.RegisterOutputType(JavaAppLayerPtrOutput{})
	pulumi.RegisterOutputType(JavaAppLayerArrayOutput{})
	pulumi.RegisterOutputType(JavaAppLayerMapOutput{})
}
