// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AWS EIP Association as a top level resource, to associate and
// disassociate Elastic IPs from AWS Instances and Network Interfaces.
//
// > **NOTE:** Do not use this resource to associate an EIP to `lb.LoadBalancer` or `ec2.NatGateway` resources. Instead use the `allocationId` available in those resources to allow AWS to manage the association, otherwise you will see `AuthFailure` errors.
//
// > **NOTE:** `ec2.EipAssociation` is useful in scenarios where EIPs are either
// pre-existing or distributed to customers or users and therefore cannot be changed.
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
// 		web, err := ec2.NewInstance(ctx, "web", &ec2.InstanceArgs{
// 			Ami:              pulumi.String("ami-21f78e11"),
// 			AvailabilityZone: pulumi.String("us-west-2a"),
// 			InstanceType:     pulumi.String("t2.micro"),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("HelloWorld"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		example, err := ec2.NewEip(ctx, "example", &ec2.EipArgs{
// 			Vpc: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewEipAssociation(ctx, "eipAssoc", &ec2.EipAssociationArgs{
// 			InstanceId:   web.ID(),
// 			AllocationId: example.ID(),
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
// EIP Assocations can be imported using their association ID.
//
// ```sh
//  $ pulumi import aws:ec2/eipAssociation:EipAssociation test eipassoc-ab12c345
// ```
type EipAssociation struct {
	pulumi.CustomResourceState

	// The allocation ID. This is required for EC2-VPC.
	AllocationId pulumi.StringOutput `pulumi:"allocationId"`
	// Whether to allow an Elastic IP to
	// be re-associated. Defaults to `true` in VPC.
	AllowReassociation pulumi.BoolPtrOutput `pulumi:"allowReassociation"`
	// The ID of the instance. This is required for
	// EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
	// network interface ID, but not both. The operation fails if you specify an
	// instance ID unless exactly one network interface is attached.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The ID of the network interface. If the
	// instance has more than one network interface, you must specify a network
	// interface ID.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The primary or secondary private IP address
	// to associate with the Elastic IP address. If no private IP address is
	// specified, the Elastic IP address is associated with the primary private IP
	// address.
	PrivateIpAddress pulumi.StringOutput `pulumi:"privateIpAddress"`
	// The Elastic IP address. This is required for EC2-Classic.
	PublicIp pulumi.StringOutput `pulumi:"publicIp"`
}

// NewEipAssociation registers a new resource with the given unique name, arguments, and options.
func NewEipAssociation(ctx *pulumi.Context,
	name string, args *EipAssociationArgs, opts ...pulumi.ResourceOption) (*EipAssociation, error) {
	if args == nil {
		args = &EipAssociationArgs{}
	}

	var resource EipAssociation
	err := ctx.RegisterResource("aws:ec2/eipAssociation:EipAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEipAssociation gets an existing EipAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEipAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EipAssociationState, opts ...pulumi.ResourceOption) (*EipAssociation, error) {
	var resource EipAssociation
	err := ctx.ReadResource("aws:ec2/eipAssociation:EipAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EipAssociation resources.
type eipAssociationState struct {
	// The allocation ID. This is required for EC2-VPC.
	AllocationId *string `pulumi:"allocationId"`
	// Whether to allow an Elastic IP to
	// be re-associated. Defaults to `true` in VPC.
	AllowReassociation *bool `pulumi:"allowReassociation"`
	// The ID of the instance. This is required for
	// EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
	// network interface ID, but not both. The operation fails if you specify an
	// instance ID unless exactly one network interface is attached.
	InstanceId *string `pulumi:"instanceId"`
	// The ID of the network interface. If the
	// instance has more than one network interface, you must specify a network
	// interface ID.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The primary or secondary private IP address
	// to associate with the Elastic IP address. If no private IP address is
	// specified, the Elastic IP address is associated with the primary private IP
	// address.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The Elastic IP address. This is required for EC2-Classic.
	PublicIp *string `pulumi:"publicIp"`
}

type EipAssociationState struct {
	// The allocation ID. This is required for EC2-VPC.
	AllocationId pulumi.StringPtrInput
	// Whether to allow an Elastic IP to
	// be re-associated. Defaults to `true` in VPC.
	AllowReassociation pulumi.BoolPtrInput
	// The ID of the instance. This is required for
	// EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
	// network interface ID, but not both. The operation fails if you specify an
	// instance ID unless exactly one network interface is attached.
	InstanceId pulumi.StringPtrInput
	// The ID of the network interface. If the
	// instance has more than one network interface, you must specify a network
	// interface ID.
	NetworkInterfaceId pulumi.StringPtrInput
	// The primary or secondary private IP address
	// to associate with the Elastic IP address. If no private IP address is
	// specified, the Elastic IP address is associated with the primary private IP
	// address.
	PrivateIpAddress pulumi.StringPtrInput
	// The Elastic IP address. This is required for EC2-Classic.
	PublicIp pulumi.StringPtrInput
}

func (EipAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*eipAssociationState)(nil)).Elem()
}

type eipAssociationArgs struct {
	// The allocation ID. This is required for EC2-VPC.
	AllocationId *string `pulumi:"allocationId"`
	// Whether to allow an Elastic IP to
	// be re-associated. Defaults to `true` in VPC.
	AllowReassociation *bool `pulumi:"allowReassociation"`
	// The ID of the instance. This is required for
	// EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
	// network interface ID, but not both. The operation fails if you specify an
	// instance ID unless exactly one network interface is attached.
	InstanceId *string `pulumi:"instanceId"`
	// The ID of the network interface. If the
	// instance has more than one network interface, you must specify a network
	// interface ID.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The primary or secondary private IP address
	// to associate with the Elastic IP address. If no private IP address is
	// specified, the Elastic IP address is associated with the primary private IP
	// address.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The Elastic IP address. This is required for EC2-Classic.
	PublicIp *string `pulumi:"publicIp"`
}

// The set of arguments for constructing a EipAssociation resource.
type EipAssociationArgs struct {
	// The allocation ID. This is required for EC2-VPC.
	AllocationId pulumi.StringPtrInput
	// Whether to allow an Elastic IP to
	// be re-associated. Defaults to `true` in VPC.
	AllowReassociation pulumi.BoolPtrInput
	// The ID of the instance. This is required for
	// EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
	// network interface ID, but not both. The operation fails if you specify an
	// instance ID unless exactly one network interface is attached.
	InstanceId pulumi.StringPtrInput
	// The ID of the network interface. If the
	// instance has more than one network interface, you must specify a network
	// interface ID.
	NetworkInterfaceId pulumi.StringPtrInput
	// The primary or secondary private IP address
	// to associate with the Elastic IP address. If no private IP address is
	// specified, the Elastic IP address is associated with the primary private IP
	// address.
	PrivateIpAddress pulumi.StringPtrInput
	// The Elastic IP address. This is required for EC2-Classic.
	PublicIp pulumi.StringPtrInput
}

func (EipAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eipAssociationArgs)(nil)).Elem()
}

type EipAssociationInput interface {
	pulumi.Input

	ToEipAssociationOutput() EipAssociationOutput
	ToEipAssociationOutputWithContext(ctx context.Context) EipAssociationOutput
}

func (*EipAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*EipAssociation)(nil))
}

func (i *EipAssociation) ToEipAssociationOutput() EipAssociationOutput {
	return i.ToEipAssociationOutputWithContext(context.Background())
}

func (i *EipAssociation) ToEipAssociationOutputWithContext(ctx context.Context) EipAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipAssociationOutput)
}

func (i *EipAssociation) ToEipAssociationPtrOutput() EipAssociationPtrOutput {
	return i.ToEipAssociationPtrOutputWithContext(context.Background())
}

func (i *EipAssociation) ToEipAssociationPtrOutputWithContext(ctx context.Context) EipAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipAssociationPtrOutput)
}

type EipAssociationPtrInput interface {
	pulumi.Input

	ToEipAssociationPtrOutput() EipAssociationPtrOutput
	ToEipAssociationPtrOutputWithContext(ctx context.Context) EipAssociationPtrOutput
}

type eipAssociationPtrType EipAssociationArgs

func (*eipAssociationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EipAssociation)(nil))
}

func (i *eipAssociationPtrType) ToEipAssociationPtrOutput() EipAssociationPtrOutput {
	return i.ToEipAssociationPtrOutputWithContext(context.Background())
}

func (i *eipAssociationPtrType) ToEipAssociationPtrOutputWithContext(ctx context.Context) EipAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipAssociationPtrOutput)
}

// EipAssociationArrayInput is an input type that accepts EipAssociationArray and EipAssociationArrayOutput values.
// You can construct a concrete instance of `EipAssociationArrayInput` via:
//
//          EipAssociationArray{ EipAssociationArgs{...} }
type EipAssociationArrayInput interface {
	pulumi.Input

	ToEipAssociationArrayOutput() EipAssociationArrayOutput
	ToEipAssociationArrayOutputWithContext(context.Context) EipAssociationArrayOutput
}

type EipAssociationArray []EipAssociationInput

func (EipAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EipAssociation)(nil)).Elem()
}

func (i EipAssociationArray) ToEipAssociationArrayOutput() EipAssociationArrayOutput {
	return i.ToEipAssociationArrayOutputWithContext(context.Background())
}

func (i EipAssociationArray) ToEipAssociationArrayOutputWithContext(ctx context.Context) EipAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipAssociationArrayOutput)
}

// EipAssociationMapInput is an input type that accepts EipAssociationMap and EipAssociationMapOutput values.
// You can construct a concrete instance of `EipAssociationMapInput` via:
//
//          EipAssociationMap{ "key": EipAssociationArgs{...} }
type EipAssociationMapInput interface {
	pulumi.Input

	ToEipAssociationMapOutput() EipAssociationMapOutput
	ToEipAssociationMapOutputWithContext(context.Context) EipAssociationMapOutput
}

type EipAssociationMap map[string]EipAssociationInput

func (EipAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EipAssociation)(nil)).Elem()
}

func (i EipAssociationMap) ToEipAssociationMapOutput() EipAssociationMapOutput {
	return i.ToEipAssociationMapOutputWithContext(context.Background())
}

func (i EipAssociationMap) ToEipAssociationMapOutputWithContext(ctx context.Context) EipAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipAssociationMapOutput)
}

type EipAssociationOutput struct {
	*pulumi.OutputState
}

func (EipAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EipAssociation)(nil))
}

func (o EipAssociationOutput) ToEipAssociationOutput() EipAssociationOutput {
	return o
}

func (o EipAssociationOutput) ToEipAssociationOutputWithContext(ctx context.Context) EipAssociationOutput {
	return o
}

func (o EipAssociationOutput) ToEipAssociationPtrOutput() EipAssociationPtrOutput {
	return o.ToEipAssociationPtrOutputWithContext(context.Background())
}

func (o EipAssociationOutput) ToEipAssociationPtrOutputWithContext(ctx context.Context) EipAssociationPtrOutput {
	return o.ApplyT(func(v EipAssociation) *EipAssociation {
		return &v
	}).(EipAssociationPtrOutput)
}

type EipAssociationPtrOutput struct {
	*pulumi.OutputState
}

func (EipAssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EipAssociation)(nil))
}

func (o EipAssociationPtrOutput) ToEipAssociationPtrOutput() EipAssociationPtrOutput {
	return o
}

func (o EipAssociationPtrOutput) ToEipAssociationPtrOutputWithContext(ctx context.Context) EipAssociationPtrOutput {
	return o
}

type EipAssociationArrayOutput struct{ *pulumi.OutputState }

func (EipAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EipAssociation)(nil))
}

func (o EipAssociationArrayOutput) ToEipAssociationArrayOutput() EipAssociationArrayOutput {
	return o
}

func (o EipAssociationArrayOutput) ToEipAssociationArrayOutputWithContext(ctx context.Context) EipAssociationArrayOutput {
	return o
}

func (o EipAssociationArrayOutput) Index(i pulumi.IntInput) EipAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EipAssociation {
		return vs[0].([]EipAssociation)[vs[1].(int)]
	}).(EipAssociationOutput)
}

type EipAssociationMapOutput struct{ *pulumi.OutputState }

func (EipAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EipAssociation)(nil))
}

func (o EipAssociationMapOutput) ToEipAssociationMapOutput() EipAssociationMapOutput {
	return o
}

func (o EipAssociationMapOutput) ToEipAssociationMapOutputWithContext(ctx context.Context) EipAssociationMapOutput {
	return o
}

func (o EipAssociationMapOutput) MapIndex(k pulumi.StringInput) EipAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EipAssociation {
		return vs[0].(map[string]EipAssociation)[vs[1].(string)]
	}).(EipAssociationOutput)
}

func init() {
	pulumi.RegisterOutputType(EipAssociationOutput{})
	pulumi.RegisterOutputType(EipAssociationPtrOutput{})
	pulumi.RegisterOutputType(EipAssociationArrayOutput{})
	pulumi.RegisterOutputType(EipAssociationMapOutput{})
}
