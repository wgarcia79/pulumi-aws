// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Route 53 Resolver DNSSEC config resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
// 			CidrBlock:          pulumi.String("10.0.0.0/16"),
// 			EnableDnsSupport:   pulumi.Bool(true),
// 			EnableDnsHostnames: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewResolverDnsSecConfig(ctx, "exampleResolverDnsSecConfig", &route53.ResolverDnsSecConfigArgs{
// 			ResourceId: exampleVpc.ID(),
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
//  Route 53 Resolver DNSSEC configs can be imported using the Route 53 Resolver DNSSEC config ID, e.g.
//
// ```sh
//  $ pulumi import aws:route53/resolverDnsSecConfig:ResolverDnsSecConfig example rdsc-be1866ecc1683e95
// ```
type ResolverDnsSecConfig struct {
	pulumi.CustomResourceState

	// The ARN for a configuration for DNSSEC validation.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The owner account ID of the virtual private cloud (VPC) for a configuration for DNSSEC validation.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The validation status for a DNSSEC configuration. The status can be one of the following: `ENABLING`, `ENABLED`, `DISABLING` and `DISABLED`.
	ValidationStatus pulumi.StringOutput `pulumi:"validationStatus"`
}

// NewResolverDnsSecConfig registers a new resource with the given unique name, arguments, and options.
func NewResolverDnsSecConfig(ctx *pulumi.Context,
	name string, args *ResolverDnsSecConfigArgs, opts ...pulumi.ResourceOption) (*ResolverDnsSecConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	var resource ResolverDnsSecConfig
	err := ctx.RegisterResource("aws:route53/resolverDnsSecConfig:ResolverDnsSecConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverDnsSecConfig gets an existing ResolverDnsSecConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverDnsSecConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverDnsSecConfigState, opts ...pulumi.ResourceOption) (*ResolverDnsSecConfig, error) {
	var resource ResolverDnsSecConfig
	err := ctx.ReadResource("aws:route53/resolverDnsSecConfig:ResolverDnsSecConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverDnsSecConfig resources.
type resolverDnsSecConfigState struct {
	// The ARN for a configuration for DNSSEC validation.
	Arn *string `pulumi:"arn"`
	// The owner account ID of the virtual private cloud (VPC) for a configuration for DNSSEC validation.
	OwnerId *string `pulumi:"ownerId"`
	// The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.
	ResourceId *string `pulumi:"resourceId"`
	// The validation status for a DNSSEC configuration. The status can be one of the following: `ENABLING`, `ENABLED`, `DISABLING` and `DISABLED`.
	ValidationStatus *string `pulumi:"validationStatus"`
}

type ResolverDnsSecConfigState struct {
	// The ARN for a configuration for DNSSEC validation.
	Arn pulumi.StringPtrInput
	// The owner account ID of the virtual private cloud (VPC) for a configuration for DNSSEC validation.
	OwnerId pulumi.StringPtrInput
	// The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.
	ResourceId pulumi.StringPtrInput
	// The validation status for a DNSSEC configuration. The status can be one of the following: `ENABLING`, `ENABLED`, `DISABLING` and `DISABLED`.
	ValidationStatus pulumi.StringPtrInput
}

func (ResolverDnsSecConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverDnsSecConfigState)(nil)).Elem()
}

type resolverDnsSecConfigArgs struct {
	// The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.
	ResourceId string `pulumi:"resourceId"`
}

// The set of arguments for constructing a ResolverDnsSecConfig resource.
type ResolverDnsSecConfigArgs struct {
	// The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.
	ResourceId pulumi.StringInput
}

func (ResolverDnsSecConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverDnsSecConfigArgs)(nil)).Elem()
}

type ResolverDnsSecConfigInput interface {
	pulumi.Input

	ToResolverDnsSecConfigOutput() ResolverDnsSecConfigOutput
	ToResolverDnsSecConfigOutputWithContext(ctx context.Context) ResolverDnsSecConfigOutput
}

func (*ResolverDnsSecConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*ResolverDnsSecConfig)(nil))
}

func (i *ResolverDnsSecConfig) ToResolverDnsSecConfigOutput() ResolverDnsSecConfigOutput {
	return i.ToResolverDnsSecConfigOutputWithContext(context.Background())
}

func (i *ResolverDnsSecConfig) ToResolverDnsSecConfigOutputWithContext(ctx context.Context) ResolverDnsSecConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverDnsSecConfigOutput)
}

func (i *ResolverDnsSecConfig) ToResolverDnsSecConfigPtrOutput() ResolverDnsSecConfigPtrOutput {
	return i.ToResolverDnsSecConfigPtrOutputWithContext(context.Background())
}

func (i *ResolverDnsSecConfig) ToResolverDnsSecConfigPtrOutputWithContext(ctx context.Context) ResolverDnsSecConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverDnsSecConfigPtrOutput)
}

type ResolverDnsSecConfigPtrInput interface {
	pulumi.Input

	ToResolverDnsSecConfigPtrOutput() ResolverDnsSecConfigPtrOutput
	ToResolverDnsSecConfigPtrOutputWithContext(ctx context.Context) ResolverDnsSecConfigPtrOutput
}

type resolverDnsSecConfigPtrType ResolverDnsSecConfigArgs

func (*resolverDnsSecConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverDnsSecConfig)(nil))
}

func (i *resolverDnsSecConfigPtrType) ToResolverDnsSecConfigPtrOutput() ResolverDnsSecConfigPtrOutput {
	return i.ToResolverDnsSecConfigPtrOutputWithContext(context.Background())
}

func (i *resolverDnsSecConfigPtrType) ToResolverDnsSecConfigPtrOutputWithContext(ctx context.Context) ResolverDnsSecConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverDnsSecConfigPtrOutput)
}

// ResolverDnsSecConfigArrayInput is an input type that accepts ResolverDnsSecConfigArray and ResolverDnsSecConfigArrayOutput values.
// You can construct a concrete instance of `ResolverDnsSecConfigArrayInput` via:
//
//          ResolverDnsSecConfigArray{ ResolverDnsSecConfigArgs{...} }
type ResolverDnsSecConfigArrayInput interface {
	pulumi.Input

	ToResolverDnsSecConfigArrayOutput() ResolverDnsSecConfigArrayOutput
	ToResolverDnsSecConfigArrayOutputWithContext(context.Context) ResolverDnsSecConfigArrayOutput
}

type ResolverDnsSecConfigArray []ResolverDnsSecConfigInput

func (ResolverDnsSecConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResolverDnsSecConfig)(nil)).Elem()
}

func (i ResolverDnsSecConfigArray) ToResolverDnsSecConfigArrayOutput() ResolverDnsSecConfigArrayOutput {
	return i.ToResolverDnsSecConfigArrayOutputWithContext(context.Background())
}

func (i ResolverDnsSecConfigArray) ToResolverDnsSecConfigArrayOutputWithContext(ctx context.Context) ResolverDnsSecConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverDnsSecConfigArrayOutput)
}

// ResolverDnsSecConfigMapInput is an input type that accepts ResolverDnsSecConfigMap and ResolverDnsSecConfigMapOutput values.
// You can construct a concrete instance of `ResolverDnsSecConfigMapInput` via:
//
//          ResolverDnsSecConfigMap{ "key": ResolverDnsSecConfigArgs{...} }
type ResolverDnsSecConfigMapInput interface {
	pulumi.Input

	ToResolverDnsSecConfigMapOutput() ResolverDnsSecConfigMapOutput
	ToResolverDnsSecConfigMapOutputWithContext(context.Context) ResolverDnsSecConfigMapOutput
}

type ResolverDnsSecConfigMap map[string]ResolverDnsSecConfigInput

func (ResolverDnsSecConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResolverDnsSecConfig)(nil)).Elem()
}

func (i ResolverDnsSecConfigMap) ToResolverDnsSecConfigMapOutput() ResolverDnsSecConfigMapOutput {
	return i.ToResolverDnsSecConfigMapOutputWithContext(context.Background())
}

func (i ResolverDnsSecConfigMap) ToResolverDnsSecConfigMapOutputWithContext(ctx context.Context) ResolverDnsSecConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverDnsSecConfigMapOutput)
}

type ResolverDnsSecConfigOutput struct{ *pulumi.OutputState }

func (ResolverDnsSecConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResolverDnsSecConfig)(nil))
}

func (o ResolverDnsSecConfigOutput) ToResolverDnsSecConfigOutput() ResolverDnsSecConfigOutput {
	return o
}

func (o ResolverDnsSecConfigOutput) ToResolverDnsSecConfigOutputWithContext(ctx context.Context) ResolverDnsSecConfigOutput {
	return o
}

func (o ResolverDnsSecConfigOutput) ToResolverDnsSecConfigPtrOutput() ResolverDnsSecConfigPtrOutput {
	return o.ToResolverDnsSecConfigPtrOutputWithContext(context.Background())
}

func (o ResolverDnsSecConfigOutput) ToResolverDnsSecConfigPtrOutputWithContext(ctx context.Context) ResolverDnsSecConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResolverDnsSecConfig) *ResolverDnsSecConfig {
		return &v
	}).(ResolverDnsSecConfigPtrOutput)
}

type ResolverDnsSecConfigPtrOutput struct{ *pulumi.OutputState }

func (ResolverDnsSecConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverDnsSecConfig)(nil))
}

func (o ResolverDnsSecConfigPtrOutput) ToResolverDnsSecConfigPtrOutput() ResolverDnsSecConfigPtrOutput {
	return o
}

func (o ResolverDnsSecConfigPtrOutput) ToResolverDnsSecConfigPtrOutputWithContext(ctx context.Context) ResolverDnsSecConfigPtrOutput {
	return o
}

func (o ResolverDnsSecConfigPtrOutput) Elem() ResolverDnsSecConfigOutput {
	return o.ApplyT(func(v *ResolverDnsSecConfig) ResolverDnsSecConfig {
		if v != nil {
			return *v
		}
		var ret ResolverDnsSecConfig
		return ret
	}).(ResolverDnsSecConfigOutput)
}

type ResolverDnsSecConfigArrayOutput struct{ *pulumi.OutputState }

func (ResolverDnsSecConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResolverDnsSecConfig)(nil))
}

func (o ResolverDnsSecConfigArrayOutput) ToResolverDnsSecConfigArrayOutput() ResolverDnsSecConfigArrayOutput {
	return o
}

func (o ResolverDnsSecConfigArrayOutput) ToResolverDnsSecConfigArrayOutputWithContext(ctx context.Context) ResolverDnsSecConfigArrayOutput {
	return o
}

func (o ResolverDnsSecConfigArrayOutput) Index(i pulumi.IntInput) ResolverDnsSecConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResolverDnsSecConfig {
		return vs[0].([]ResolverDnsSecConfig)[vs[1].(int)]
	}).(ResolverDnsSecConfigOutput)
}

type ResolverDnsSecConfigMapOutput struct{ *pulumi.OutputState }

func (ResolverDnsSecConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResolverDnsSecConfig)(nil))
}

func (o ResolverDnsSecConfigMapOutput) ToResolverDnsSecConfigMapOutput() ResolverDnsSecConfigMapOutput {
	return o
}

func (o ResolverDnsSecConfigMapOutput) ToResolverDnsSecConfigMapOutputWithContext(ctx context.Context) ResolverDnsSecConfigMapOutput {
	return o
}

func (o ResolverDnsSecConfigMapOutput) MapIndex(k pulumi.StringInput) ResolverDnsSecConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResolverDnsSecConfig {
		return vs[0].(map[string]ResolverDnsSecConfig)[vs[1].(string)]
	}).(ResolverDnsSecConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverDnsSecConfigInput)(nil)).Elem(), &ResolverDnsSecConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverDnsSecConfigPtrInput)(nil)).Elem(), &ResolverDnsSecConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverDnsSecConfigArrayInput)(nil)).Elem(), ResolverDnsSecConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverDnsSecConfigMapInput)(nil)).Elem(), ResolverDnsSecConfigMap{})
	pulumi.RegisterOutputType(ResolverDnsSecConfigOutput{})
	pulumi.RegisterOutputType(ResolverDnsSecConfigPtrOutput{})
	pulumi.RegisterOutputType(ResolverDnsSecConfigArrayOutput{})
	pulumi.RegisterOutputType(ResolverDnsSecConfigMapOutput{})
}
