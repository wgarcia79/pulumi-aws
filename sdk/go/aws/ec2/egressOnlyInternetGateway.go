// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// [IPv6 only] Creates an egress-only Internet gateway for your VPC.
// An egress-only Internet gateway is used to enable outbound communication
// over IPv6 from instances in your VPC to the Internet, and prevents hosts
// outside of your VPC from initiating an IPv6 connection with your instance.
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
// 		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
// 			CidrBlock:                    pulumi.String("10.1.0.0/16"),
// 			AssignGeneratedIpv6CidrBlock: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewEgressOnlyInternetGateway(ctx, "exampleEgressOnlyInternetGateway", &ec2.EgressOnlyInternetGatewayArgs{
// 			VpcId: exampleVpc.ID(),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("main"),
// 			},
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
// Egress-only Internet gateways can be imported using the `id`, e.g.,
//
// ```sh
//  $ pulumi import aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway example eigw-015e0e244e24dfe8a
// ```
type EgressOnlyInternetGateway struct {
	pulumi.CustomResourceState

	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The VPC ID to create in.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewEgressOnlyInternetGateway registers a new resource with the given unique name, arguments, and options.
func NewEgressOnlyInternetGateway(ctx *pulumi.Context,
	name string, args *EgressOnlyInternetGatewayArgs, opts ...pulumi.ResourceOption) (*EgressOnlyInternetGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource EgressOnlyInternetGateway
	err := ctx.RegisterResource("aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEgressOnlyInternetGateway gets an existing EgressOnlyInternetGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEgressOnlyInternetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EgressOnlyInternetGatewayState, opts ...pulumi.ResourceOption) (*EgressOnlyInternetGateway, error) {
	var resource EgressOnlyInternetGateway
	err := ctx.ReadResource("aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EgressOnlyInternetGateway resources.
type egressOnlyInternetGatewayState struct {
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The VPC ID to create in.
	VpcId *string `pulumi:"vpcId"`
}

type EgressOnlyInternetGatewayState struct {
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The VPC ID to create in.
	VpcId pulumi.StringPtrInput
}

func (EgressOnlyInternetGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*egressOnlyInternetGatewayState)(nil)).Elem()
}

type egressOnlyInternetGatewayArgs struct {
	Tags map[string]string `pulumi:"tags"`
	// The VPC ID to create in.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a EgressOnlyInternetGateway resource.
type EgressOnlyInternetGatewayArgs struct {
	Tags pulumi.StringMapInput
	// The VPC ID to create in.
	VpcId pulumi.StringInput
}

func (EgressOnlyInternetGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*egressOnlyInternetGatewayArgs)(nil)).Elem()
}

type EgressOnlyInternetGatewayInput interface {
	pulumi.Input

	ToEgressOnlyInternetGatewayOutput() EgressOnlyInternetGatewayOutput
	ToEgressOnlyInternetGatewayOutputWithContext(ctx context.Context) EgressOnlyInternetGatewayOutput
}

func (*EgressOnlyInternetGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*EgressOnlyInternetGateway)(nil))
}

func (i *EgressOnlyInternetGateway) ToEgressOnlyInternetGatewayOutput() EgressOnlyInternetGatewayOutput {
	return i.ToEgressOnlyInternetGatewayOutputWithContext(context.Background())
}

func (i *EgressOnlyInternetGateway) ToEgressOnlyInternetGatewayOutputWithContext(ctx context.Context) EgressOnlyInternetGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EgressOnlyInternetGatewayOutput)
}

func (i *EgressOnlyInternetGateway) ToEgressOnlyInternetGatewayPtrOutput() EgressOnlyInternetGatewayPtrOutput {
	return i.ToEgressOnlyInternetGatewayPtrOutputWithContext(context.Background())
}

func (i *EgressOnlyInternetGateway) ToEgressOnlyInternetGatewayPtrOutputWithContext(ctx context.Context) EgressOnlyInternetGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EgressOnlyInternetGatewayPtrOutput)
}

type EgressOnlyInternetGatewayPtrInput interface {
	pulumi.Input

	ToEgressOnlyInternetGatewayPtrOutput() EgressOnlyInternetGatewayPtrOutput
	ToEgressOnlyInternetGatewayPtrOutputWithContext(ctx context.Context) EgressOnlyInternetGatewayPtrOutput
}

type egressOnlyInternetGatewayPtrType EgressOnlyInternetGatewayArgs

func (*egressOnlyInternetGatewayPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EgressOnlyInternetGateway)(nil))
}

func (i *egressOnlyInternetGatewayPtrType) ToEgressOnlyInternetGatewayPtrOutput() EgressOnlyInternetGatewayPtrOutput {
	return i.ToEgressOnlyInternetGatewayPtrOutputWithContext(context.Background())
}

func (i *egressOnlyInternetGatewayPtrType) ToEgressOnlyInternetGatewayPtrOutputWithContext(ctx context.Context) EgressOnlyInternetGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EgressOnlyInternetGatewayPtrOutput)
}

// EgressOnlyInternetGatewayArrayInput is an input type that accepts EgressOnlyInternetGatewayArray and EgressOnlyInternetGatewayArrayOutput values.
// You can construct a concrete instance of `EgressOnlyInternetGatewayArrayInput` via:
//
//          EgressOnlyInternetGatewayArray{ EgressOnlyInternetGatewayArgs{...} }
type EgressOnlyInternetGatewayArrayInput interface {
	pulumi.Input

	ToEgressOnlyInternetGatewayArrayOutput() EgressOnlyInternetGatewayArrayOutput
	ToEgressOnlyInternetGatewayArrayOutputWithContext(context.Context) EgressOnlyInternetGatewayArrayOutput
}

type EgressOnlyInternetGatewayArray []EgressOnlyInternetGatewayInput

func (EgressOnlyInternetGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EgressOnlyInternetGateway)(nil)).Elem()
}

func (i EgressOnlyInternetGatewayArray) ToEgressOnlyInternetGatewayArrayOutput() EgressOnlyInternetGatewayArrayOutput {
	return i.ToEgressOnlyInternetGatewayArrayOutputWithContext(context.Background())
}

func (i EgressOnlyInternetGatewayArray) ToEgressOnlyInternetGatewayArrayOutputWithContext(ctx context.Context) EgressOnlyInternetGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EgressOnlyInternetGatewayArrayOutput)
}

// EgressOnlyInternetGatewayMapInput is an input type that accepts EgressOnlyInternetGatewayMap and EgressOnlyInternetGatewayMapOutput values.
// You can construct a concrete instance of `EgressOnlyInternetGatewayMapInput` via:
//
//          EgressOnlyInternetGatewayMap{ "key": EgressOnlyInternetGatewayArgs{...} }
type EgressOnlyInternetGatewayMapInput interface {
	pulumi.Input

	ToEgressOnlyInternetGatewayMapOutput() EgressOnlyInternetGatewayMapOutput
	ToEgressOnlyInternetGatewayMapOutputWithContext(context.Context) EgressOnlyInternetGatewayMapOutput
}

type EgressOnlyInternetGatewayMap map[string]EgressOnlyInternetGatewayInput

func (EgressOnlyInternetGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EgressOnlyInternetGateway)(nil)).Elem()
}

func (i EgressOnlyInternetGatewayMap) ToEgressOnlyInternetGatewayMapOutput() EgressOnlyInternetGatewayMapOutput {
	return i.ToEgressOnlyInternetGatewayMapOutputWithContext(context.Background())
}

func (i EgressOnlyInternetGatewayMap) ToEgressOnlyInternetGatewayMapOutputWithContext(ctx context.Context) EgressOnlyInternetGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EgressOnlyInternetGatewayMapOutput)
}

type EgressOnlyInternetGatewayOutput struct{ *pulumi.OutputState }

func (EgressOnlyInternetGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EgressOnlyInternetGateway)(nil))
}

func (o EgressOnlyInternetGatewayOutput) ToEgressOnlyInternetGatewayOutput() EgressOnlyInternetGatewayOutput {
	return o
}

func (o EgressOnlyInternetGatewayOutput) ToEgressOnlyInternetGatewayOutputWithContext(ctx context.Context) EgressOnlyInternetGatewayOutput {
	return o
}

func (o EgressOnlyInternetGatewayOutput) ToEgressOnlyInternetGatewayPtrOutput() EgressOnlyInternetGatewayPtrOutput {
	return o.ToEgressOnlyInternetGatewayPtrOutputWithContext(context.Background())
}

func (o EgressOnlyInternetGatewayOutput) ToEgressOnlyInternetGatewayPtrOutputWithContext(ctx context.Context) EgressOnlyInternetGatewayPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EgressOnlyInternetGateway) *EgressOnlyInternetGateway {
		return &v
	}).(EgressOnlyInternetGatewayPtrOutput)
}

type EgressOnlyInternetGatewayPtrOutput struct{ *pulumi.OutputState }

func (EgressOnlyInternetGatewayPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EgressOnlyInternetGateway)(nil))
}

func (o EgressOnlyInternetGatewayPtrOutput) ToEgressOnlyInternetGatewayPtrOutput() EgressOnlyInternetGatewayPtrOutput {
	return o
}

func (o EgressOnlyInternetGatewayPtrOutput) ToEgressOnlyInternetGatewayPtrOutputWithContext(ctx context.Context) EgressOnlyInternetGatewayPtrOutput {
	return o
}

func (o EgressOnlyInternetGatewayPtrOutput) Elem() EgressOnlyInternetGatewayOutput {
	return o.ApplyT(func(v *EgressOnlyInternetGateway) EgressOnlyInternetGateway {
		if v != nil {
			return *v
		}
		var ret EgressOnlyInternetGateway
		return ret
	}).(EgressOnlyInternetGatewayOutput)
}

type EgressOnlyInternetGatewayArrayOutput struct{ *pulumi.OutputState }

func (EgressOnlyInternetGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EgressOnlyInternetGateway)(nil))
}

func (o EgressOnlyInternetGatewayArrayOutput) ToEgressOnlyInternetGatewayArrayOutput() EgressOnlyInternetGatewayArrayOutput {
	return o
}

func (o EgressOnlyInternetGatewayArrayOutput) ToEgressOnlyInternetGatewayArrayOutputWithContext(ctx context.Context) EgressOnlyInternetGatewayArrayOutput {
	return o
}

func (o EgressOnlyInternetGatewayArrayOutput) Index(i pulumi.IntInput) EgressOnlyInternetGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EgressOnlyInternetGateway {
		return vs[0].([]EgressOnlyInternetGateway)[vs[1].(int)]
	}).(EgressOnlyInternetGatewayOutput)
}

type EgressOnlyInternetGatewayMapOutput struct{ *pulumi.OutputState }

func (EgressOnlyInternetGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EgressOnlyInternetGateway)(nil))
}

func (o EgressOnlyInternetGatewayMapOutput) ToEgressOnlyInternetGatewayMapOutput() EgressOnlyInternetGatewayMapOutput {
	return o
}

func (o EgressOnlyInternetGatewayMapOutput) ToEgressOnlyInternetGatewayMapOutputWithContext(ctx context.Context) EgressOnlyInternetGatewayMapOutput {
	return o
}

func (o EgressOnlyInternetGatewayMapOutput) MapIndex(k pulumi.StringInput) EgressOnlyInternetGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EgressOnlyInternetGateway {
		return vs[0].(map[string]EgressOnlyInternetGateway)[vs[1].(string)]
	}).(EgressOnlyInternetGatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EgressOnlyInternetGatewayInput)(nil)).Elem(), &EgressOnlyInternetGateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*EgressOnlyInternetGatewayPtrInput)(nil)).Elem(), &EgressOnlyInternetGateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*EgressOnlyInternetGatewayArrayInput)(nil)).Elem(), EgressOnlyInternetGatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EgressOnlyInternetGatewayMapInput)(nil)).Elem(), EgressOnlyInternetGatewayMap{})
	pulumi.RegisterOutputType(EgressOnlyInternetGatewayOutput{})
	pulumi.RegisterOutputType(EgressOnlyInternetGatewayPtrOutput{})
	pulumi.RegisterOutputType(EgressOnlyInternetGatewayArrayOutput{})
	pulumi.RegisterOutputType(EgressOnlyInternetGatewayMapOutput{})
}
