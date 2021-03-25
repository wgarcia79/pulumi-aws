// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an EC2 host resource. This allows hosts to be created, updated,
// and deleted.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		test, err := ec2.NewDedicatedHost(ctx, "test", &ec2.DedicatedHostArgs{
// 			AutoPlacement:    pulumi.String("on"),
// 			AvailabilityZone: pulumi.String("us-west-1a"),
// 			HostRecovery:     pulumi.String("on"),
// 			InstanceType:     pulumi.String("c5.18xlarge"),
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
// hosts can be imported using the `host_id`, e.g.
//
// ```sh
//  $ pulumi import aws:ec2/dedicatedHost:DedicatedHost host_id h-0385a99d0e4b20cbb
// ```
type DedicatedHost struct {
	pulumi.CustomResourceState

	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
	AutoPlacement pulumi.StringOutput `pulumi:"autoPlacement"`
	// The AZ to start the host in.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.
	HostRecovery pulumi.StringOutput `pulumi:"hostRecovery"`
	// Specifies the instance family for which to configure your Dedicated Host. Mutually exclusive with `instanceType`.
	InstanceFamily pulumi.StringPtrOutput `pulumi:"instanceFamily"`
	// Specifies the instance type for which to configure your Dedicated Host. When you specify the instance type, that is the only instance type that you can launch onto that host. Mutually exclusive with `instanceFamily`.
	InstanceType pulumi.StringPtrOutput `pulumi:"instanceType"`
	Tags         pulumi.StringMapOutput `pulumi:"tags"`
}

// NewDedicatedHost registers a new resource with the given unique name, arguments, and options.
func NewDedicatedHost(ctx *pulumi.Context,
	name string, args *DedicatedHostArgs, opts ...pulumi.ResourceOption) (*DedicatedHost, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	var resource DedicatedHost
	err := ctx.RegisterResource("aws:ec2/dedicatedHost:DedicatedHost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedHost gets an existing DedicatedHost resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedHostState, opts ...pulumi.ResourceOption) (*DedicatedHost, error) {
	var resource DedicatedHost
	err := ctx.ReadResource("aws:ec2/dedicatedHost:DedicatedHost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedHost resources.
type dedicatedHostState struct {
	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
	AutoPlacement *string `pulumi:"autoPlacement"`
	// The AZ to start the host in.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.
	HostRecovery *string `pulumi:"hostRecovery"`
	// Specifies the instance family for which to configure your Dedicated Host. Mutually exclusive with `instanceType`.
	InstanceFamily *string `pulumi:"instanceFamily"`
	// Specifies the instance type for which to configure your Dedicated Host. When you specify the instance type, that is the only instance type that you can launch onto that host. Mutually exclusive with `instanceFamily`.
	InstanceType *string           `pulumi:"instanceType"`
	Tags         map[string]string `pulumi:"tags"`
}

type DedicatedHostState struct {
	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
	AutoPlacement pulumi.StringPtrInput
	// The AZ to start the host in.
	AvailabilityZone pulumi.StringPtrInput
	// Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.
	HostRecovery pulumi.StringPtrInput
	// Specifies the instance family for which to configure your Dedicated Host. Mutually exclusive with `instanceType`.
	InstanceFamily pulumi.StringPtrInput
	// Specifies the instance type for which to configure your Dedicated Host. When you specify the instance type, that is the only instance type that you can launch onto that host. Mutually exclusive with `instanceFamily`.
	InstanceType pulumi.StringPtrInput
	Tags         pulumi.StringMapInput
}

func (DedicatedHostState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostState)(nil)).Elem()
}

type dedicatedHostArgs struct {
	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
	AutoPlacement *string `pulumi:"autoPlacement"`
	// The AZ to start the host in.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.
	HostRecovery *string `pulumi:"hostRecovery"`
	// Specifies the instance family for which to configure your Dedicated Host. Mutually exclusive with `instanceType`.
	InstanceFamily *string `pulumi:"instanceFamily"`
	// Specifies the instance type for which to configure your Dedicated Host. When you specify the instance type, that is the only instance type that you can launch onto that host. Mutually exclusive with `instanceFamily`.
	InstanceType *string           `pulumi:"instanceType"`
	Tags         map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DedicatedHost resource.
type DedicatedHostArgs struct {
	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
	AutoPlacement pulumi.StringPtrInput
	// The AZ to start the host in.
	AvailabilityZone pulumi.StringInput
	// Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.
	HostRecovery pulumi.StringPtrInput
	// Specifies the instance family for which to configure your Dedicated Host. Mutually exclusive with `instanceType`.
	InstanceFamily pulumi.StringPtrInput
	// Specifies the instance type for which to configure your Dedicated Host. When you specify the instance type, that is the only instance type that you can launch onto that host. Mutually exclusive with `instanceFamily`.
	InstanceType pulumi.StringPtrInput
	Tags         pulumi.StringMapInput
}

func (DedicatedHostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostArgs)(nil)).Elem()
}

type DedicatedHostInput interface {
	pulumi.Input

	ToDedicatedHostOutput() DedicatedHostOutput
	ToDedicatedHostOutputWithContext(ctx context.Context) DedicatedHostOutput
}

func (*DedicatedHost) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedHost)(nil))
}

func (i *DedicatedHost) ToDedicatedHostOutput() DedicatedHostOutput {
	return i.ToDedicatedHostOutputWithContext(context.Background())
}

func (i *DedicatedHost) ToDedicatedHostOutputWithContext(ctx context.Context) DedicatedHostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostOutput)
}

func (i *DedicatedHost) ToDedicatedHostPtrOutput() DedicatedHostPtrOutput {
	return i.ToDedicatedHostPtrOutputWithContext(context.Background())
}

func (i *DedicatedHost) ToDedicatedHostPtrOutputWithContext(ctx context.Context) DedicatedHostPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostPtrOutput)
}

type DedicatedHostPtrInput interface {
	pulumi.Input

	ToDedicatedHostPtrOutput() DedicatedHostPtrOutput
	ToDedicatedHostPtrOutputWithContext(ctx context.Context) DedicatedHostPtrOutput
}

type dedicatedHostPtrType DedicatedHostArgs

func (*dedicatedHostPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHost)(nil))
}

func (i *dedicatedHostPtrType) ToDedicatedHostPtrOutput() DedicatedHostPtrOutput {
	return i.ToDedicatedHostPtrOutputWithContext(context.Background())
}

func (i *dedicatedHostPtrType) ToDedicatedHostPtrOutputWithContext(ctx context.Context) DedicatedHostPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostPtrOutput)
}

// DedicatedHostArrayInput is an input type that accepts DedicatedHostArray and DedicatedHostArrayOutput values.
// You can construct a concrete instance of `DedicatedHostArrayInput` via:
//
//          DedicatedHostArray{ DedicatedHostArgs{...} }
type DedicatedHostArrayInput interface {
	pulumi.Input

	ToDedicatedHostArrayOutput() DedicatedHostArrayOutput
	ToDedicatedHostArrayOutputWithContext(context.Context) DedicatedHostArrayOutput
}

type DedicatedHostArray []DedicatedHostInput

func (DedicatedHostArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DedicatedHost)(nil))
}

func (i DedicatedHostArray) ToDedicatedHostArrayOutput() DedicatedHostArrayOutput {
	return i.ToDedicatedHostArrayOutputWithContext(context.Background())
}

func (i DedicatedHostArray) ToDedicatedHostArrayOutputWithContext(ctx context.Context) DedicatedHostArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostArrayOutput)
}

// DedicatedHostMapInput is an input type that accepts DedicatedHostMap and DedicatedHostMapOutput values.
// You can construct a concrete instance of `DedicatedHostMapInput` via:
//
//          DedicatedHostMap{ "key": DedicatedHostArgs{...} }
type DedicatedHostMapInput interface {
	pulumi.Input

	ToDedicatedHostMapOutput() DedicatedHostMapOutput
	ToDedicatedHostMapOutputWithContext(context.Context) DedicatedHostMapOutput
}

type DedicatedHostMap map[string]DedicatedHostInput

func (DedicatedHostMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DedicatedHost)(nil))
}

func (i DedicatedHostMap) ToDedicatedHostMapOutput() DedicatedHostMapOutput {
	return i.ToDedicatedHostMapOutputWithContext(context.Background())
}

func (i DedicatedHostMap) ToDedicatedHostMapOutputWithContext(ctx context.Context) DedicatedHostMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostMapOutput)
}

type DedicatedHostOutput struct {
	*pulumi.OutputState
}

func (DedicatedHostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedHost)(nil))
}

func (o DedicatedHostOutput) ToDedicatedHostOutput() DedicatedHostOutput {
	return o
}

func (o DedicatedHostOutput) ToDedicatedHostOutputWithContext(ctx context.Context) DedicatedHostOutput {
	return o
}

func (o DedicatedHostOutput) ToDedicatedHostPtrOutput() DedicatedHostPtrOutput {
	return o.ToDedicatedHostPtrOutputWithContext(context.Background())
}

func (o DedicatedHostOutput) ToDedicatedHostPtrOutputWithContext(ctx context.Context) DedicatedHostPtrOutput {
	return o.ApplyT(func(v DedicatedHost) *DedicatedHost {
		return &v
	}).(DedicatedHostPtrOutput)
}

type DedicatedHostPtrOutput struct {
	*pulumi.OutputState
}

func (DedicatedHostPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHost)(nil))
}

func (o DedicatedHostPtrOutput) ToDedicatedHostPtrOutput() DedicatedHostPtrOutput {
	return o
}

func (o DedicatedHostPtrOutput) ToDedicatedHostPtrOutputWithContext(ctx context.Context) DedicatedHostPtrOutput {
	return o
}

type DedicatedHostArrayOutput struct{ *pulumi.OutputState }

func (DedicatedHostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DedicatedHost)(nil))
}

func (o DedicatedHostArrayOutput) ToDedicatedHostArrayOutput() DedicatedHostArrayOutput {
	return o
}

func (o DedicatedHostArrayOutput) ToDedicatedHostArrayOutputWithContext(ctx context.Context) DedicatedHostArrayOutput {
	return o
}

func (o DedicatedHostArrayOutput) Index(i pulumi.IntInput) DedicatedHostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DedicatedHost {
		return vs[0].([]DedicatedHost)[vs[1].(int)]
	}).(DedicatedHostOutput)
}

type DedicatedHostMapOutput struct{ *pulumi.OutputState }

func (DedicatedHostMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DedicatedHost)(nil))
}

func (o DedicatedHostMapOutput) ToDedicatedHostMapOutput() DedicatedHostMapOutput {
	return o
}

func (o DedicatedHostMapOutput) ToDedicatedHostMapOutputWithContext(ctx context.Context) DedicatedHostMapOutput {
	return o
}

func (o DedicatedHostMapOutput) MapIndex(k pulumi.StringInput) DedicatedHostOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DedicatedHost {
		return vs[0].(map[string]DedicatedHost)[vs[1].(string)]
	}).(DedicatedHostOutput)
}

func init() {
	pulumi.RegisterOutputType(DedicatedHostOutput{})
	pulumi.RegisterOutputType(DedicatedHostPtrOutput{})
	pulumi.RegisterOutputType(DedicatedHostArrayOutput{})
	pulumi.RegisterOutputType(DedicatedHostMapOutput{})
}
