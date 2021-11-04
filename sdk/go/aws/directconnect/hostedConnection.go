// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a hosted connection on the specified interconnect or a link aggregation group (LAG) of interconnects. Intended for use by AWS Direct Connect Partners only.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/directconnect"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := directconnect.NewHostedConnection(ctx, "hosted", &directconnect.HostedConnectionArgs{
// 			Bandwidth:      pulumi.String("100Mbps"),
// 			ConnectionId:   pulumi.String("dxcon-ffabc123"),
// 			OwnerAccountId: pulumi.String("123456789012"),
// 			Vlan:           pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type HostedConnection struct {
	pulumi.CustomResourceState

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDevice pulumi.StringOutput `pulumi:"awsDevice"`
	// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps and 10Gbps. Case sensitive.
	Bandwidth pulumi.StringOutput `pulumi:"bandwidth"`
	// The ID of the interconnect or LAG.
	ConnectionId pulumi.StringOutput `pulumi:"connectionId"`
	// Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
	HasLogicalRedundancy pulumi.StringOutput `pulumi:"hasLogicalRedundancy"`
	// Boolean value representing if jumbo frames have been enabled for this connection.
	JumboFrameCapable pulumi.BoolOutput `pulumi:"jumboFrameCapable"`
	// The ID of the LAG.
	LagId pulumi.StringOutput `pulumi:"lagId"`
	// The time of the most recent call to [DescribeLoa](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLoa.html) for this connection.
	LoaIssueTime pulumi.StringOutput `pulumi:"loaIssueTime"`
	// The location of the connection.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the connection.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the AWS account of the customer for the connection.
	OwnerAccountId pulumi.StringOutput `pulumi:"ownerAccountId"`
	// The name of the AWS Direct Connect service provider associated with the connection.
	PartnerName pulumi.StringOutput `pulumi:"partnerName"`
	// The name of the service provider associated with the connection.
	ProviderName pulumi.StringOutput `pulumi:"providerName"`
	// The AWS Region where the connection is located.
	Region pulumi.StringOutput `pulumi:"region"`
	// The state of the connection. Possible values include: ordering, requested, pending, available, down, deleting, deleted, rejected, unknown. See [AllocateHostedConnection](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_AllocateHostedConnection.html) for a description of each connection state.
	State pulumi.StringOutput `pulumi:"state"`
	// The dedicated VLAN provisioned to the hosted connection.
	Vlan pulumi.IntOutput `pulumi:"vlan"`
}

// NewHostedConnection registers a new resource with the given unique name, arguments, and options.
func NewHostedConnection(ctx *pulumi.Context,
	name string, args *HostedConnectionArgs, opts ...pulumi.ResourceOption) (*HostedConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bandwidth == nil {
		return nil, errors.New("invalid value for required argument 'Bandwidth'")
	}
	if args.ConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionId'")
	}
	if args.OwnerAccountId == nil {
		return nil, errors.New("invalid value for required argument 'OwnerAccountId'")
	}
	if args.Vlan == nil {
		return nil, errors.New("invalid value for required argument 'Vlan'")
	}
	var resource HostedConnection
	err := ctx.RegisterResource("aws:directconnect/hostedConnection:HostedConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostedConnection gets an existing HostedConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostedConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostedConnectionState, opts ...pulumi.ResourceOption) (*HostedConnection, error) {
	var resource HostedConnection
	err := ctx.ReadResource("aws:directconnect/hostedConnection:HostedConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostedConnection resources.
type hostedConnectionState struct {
	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDevice *string `pulumi:"awsDevice"`
	// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps and 10Gbps. Case sensitive.
	Bandwidth *string `pulumi:"bandwidth"`
	// The ID of the interconnect or LAG.
	ConnectionId *string `pulumi:"connectionId"`
	// Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
	HasLogicalRedundancy *string `pulumi:"hasLogicalRedundancy"`
	// Boolean value representing if jumbo frames have been enabled for this connection.
	JumboFrameCapable *bool `pulumi:"jumboFrameCapable"`
	// The ID of the LAG.
	LagId *string `pulumi:"lagId"`
	// The time of the most recent call to [DescribeLoa](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLoa.html) for this connection.
	LoaIssueTime *string `pulumi:"loaIssueTime"`
	// The location of the connection.
	Location *string `pulumi:"location"`
	// The name of the connection.
	Name *string `pulumi:"name"`
	// The ID of the AWS account of the customer for the connection.
	OwnerAccountId *string `pulumi:"ownerAccountId"`
	// The name of the AWS Direct Connect service provider associated with the connection.
	PartnerName *string `pulumi:"partnerName"`
	// The name of the service provider associated with the connection.
	ProviderName *string `pulumi:"providerName"`
	// The AWS Region where the connection is located.
	Region *string `pulumi:"region"`
	// The state of the connection. Possible values include: ordering, requested, pending, available, down, deleting, deleted, rejected, unknown. See [AllocateHostedConnection](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_AllocateHostedConnection.html) for a description of each connection state.
	State *string `pulumi:"state"`
	// The dedicated VLAN provisioned to the hosted connection.
	Vlan *int `pulumi:"vlan"`
}

type HostedConnectionState struct {
	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDevice pulumi.StringPtrInput
	// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps and 10Gbps. Case sensitive.
	Bandwidth pulumi.StringPtrInput
	// The ID of the interconnect or LAG.
	ConnectionId pulumi.StringPtrInput
	// Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
	HasLogicalRedundancy pulumi.StringPtrInput
	// Boolean value representing if jumbo frames have been enabled for this connection.
	JumboFrameCapable pulumi.BoolPtrInput
	// The ID of the LAG.
	LagId pulumi.StringPtrInput
	// The time of the most recent call to [DescribeLoa](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLoa.html) for this connection.
	LoaIssueTime pulumi.StringPtrInput
	// The location of the connection.
	Location pulumi.StringPtrInput
	// The name of the connection.
	Name pulumi.StringPtrInput
	// The ID of the AWS account of the customer for the connection.
	OwnerAccountId pulumi.StringPtrInput
	// The name of the AWS Direct Connect service provider associated with the connection.
	PartnerName pulumi.StringPtrInput
	// The name of the service provider associated with the connection.
	ProviderName pulumi.StringPtrInput
	// The AWS Region where the connection is located.
	Region pulumi.StringPtrInput
	// The state of the connection. Possible values include: ordering, requested, pending, available, down, deleting, deleted, rejected, unknown. See [AllocateHostedConnection](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_AllocateHostedConnection.html) for a description of each connection state.
	State pulumi.StringPtrInput
	// The dedicated VLAN provisioned to the hosted connection.
	Vlan pulumi.IntPtrInput
}

func (HostedConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedConnectionState)(nil)).Elem()
}

type hostedConnectionArgs struct {
	// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps and 10Gbps. Case sensitive.
	Bandwidth string `pulumi:"bandwidth"`
	// The ID of the interconnect or LAG.
	ConnectionId string `pulumi:"connectionId"`
	// The name of the connection.
	Name *string `pulumi:"name"`
	// The ID of the AWS account of the customer for the connection.
	OwnerAccountId string `pulumi:"ownerAccountId"`
	// The dedicated VLAN provisioned to the hosted connection.
	Vlan int `pulumi:"vlan"`
}

// The set of arguments for constructing a HostedConnection resource.
type HostedConnectionArgs struct {
	// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps and 10Gbps. Case sensitive.
	Bandwidth pulumi.StringInput
	// The ID of the interconnect or LAG.
	ConnectionId pulumi.StringInput
	// The name of the connection.
	Name pulumi.StringPtrInput
	// The ID of the AWS account of the customer for the connection.
	OwnerAccountId pulumi.StringInput
	// The dedicated VLAN provisioned to the hosted connection.
	Vlan pulumi.IntInput
}

func (HostedConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedConnectionArgs)(nil)).Elem()
}

type HostedConnectionInput interface {
	pulumi.Input

	ToHostedConnectionOutput() HostedConnectionOutput
	ToHostedConnectionOutputWithContext(ctx context.Context) HostedConnectionOutput
}

func (*HostedConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*HostedConnection)(nil))
}

func (i *HostedConnection) ToHostedConnectionOutput() HostedConnectionOutput {
	return i.ToHostedConnectionOutputWithContext(context.Background())
}

func (i *HostedConnection) ToHostedConnectionOutputWithContext(ctx context.Context) HostedConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedConnectionOutput)
}

func (i *HostedConnection) ToHostedConnectionPtrOutput() HostedConnectionPtrOutput {
	return i.ToHostedConnectionPtrOutputWithContext(context.Background())
}

func (i *HostedConnection) ToHostedConnectionPtrOutputWithContext(ctx context.Context) HostedConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedConnectionPtrOutput)
}

type HostedConnectionPtrInput interface {
	pulumi.Input

	ToHostedConnectionPtrOutput() HostedConnectionPtrOutput
	ToHostedConnectionPtrOutputWithContext(ctx context.Context) HostedConnectionPtrOutput
}

type hostedConnectionPtrType HostedConnectionArgs

func (*hostedConnectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HostedConnection)(nil))
}

func (i *hostedConnectionPtrType) ToHostedConnectionPtrOutput() HostedConnectionPtrOutput {
	return i.ToHostedConnectionPtrOutputWithContext(context.Background())
}

func (i *hostedConnectionPtrType) ToHostedConnectionPtrOutputWithContext(ctx context.Context) HostedConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedConnectionPtrOutput)
}

// HostedConnectionArrayInput is an input type that accepts HostedConnectionArray and HostedConnectionArrayOutput values.
// You can construct a concrete instance of `HostedConnectionArrayInput` via:
//
//          HostedConnectionArray{ HostedConnectionArgs{...} }
type HostedConnectionArrayInput interface {
	pulumi.Input

	ToHostedConnectionArrayOutput() HostedConnectionArrayOutput
	ToHostedConnectionArrayOutputWithContext(context.Context) HostedConnectionArrayOutput
}

type HostedConnectionArray []HostedConnectionInput

func (HostedConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostedConnection)(nil)).Elem()
}

func (i HostedConnectionArray) ToHostedConnectionArrayOutput() HostedConnectionArrayOutput {
	return i.ToHostedConnectionArrayOutputWithContext(context.Background())
}

func (i HostedConnectionArray) ToHostedConnectionArrayOutputWithContext(ctx context.Context) HostedConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedConnectionArrayOutput)
}

// HostedConnectionMapInput is an input type that accepts HostedConnectionMap and HostedConnectionMapOutput values.
// You can construct a concrete instance of `HostedConnectionMapInput` via:
//
//          HostedConnectionMap{ "key": HostedConnectionArgs{...} }
type HostedConnectionMapInput interface {
	pulumi.Input

	ToHostedConnectionMapOutput() HostedConnectionMapOutput
	ToHostedConnectionMapOutputWithContext(context.Context) HostedConnectionMapOutput
}

type HostedConnectionMap map[string]HostedConnectionInput

func (HostedConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostedConnection)(nil)).Elem()
}

func (i HostedConnectionMap) ToHostedConnectionMapOutput() HostedConnectionMapOutput {
	return i.ToHostedConnectionMapOutputWithContext(context.Background())
}

func (i HostedConnectionMap) ToHostedConnectionMapOutputWithContext(ctx context.Context) HostedConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedConnectionMapOutput)
}

type HostedConnectionOutput struct{ *pulumi.OutputState }

func (HostedConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostedConnection)(nil))
}

func (o HostedConnectionOutput) ToHostedConnectionOutput() HostedConnectionOutput {
	return o
}

func (o HostedConnectionOutput) ToHostedConnectionOutputWithContext(ctx context.Context) HostedConnectionOutput {
	return o
}

func (o HostedConnectionOutput) ToHostedConnectionPtrOutput() HostedConnectionPtrOutput {
	return o.ToHostedConnectionPtrOutputWithContext(context.Background())
}

func (o HostedConnectionOutput) ToHostedConnectionPtrOutputWithContext(ctx context.Context) HostedConnectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostedConnection) *HostedConnection {
		return &v
	}).(HostedConnectionPtrOutput)
}

type HostedConnectionPtrOutput struct{ *pulumi.OutputState }

func (HostedConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostedConnection)(nil))
}

func (o HostedConnectionPtrOutput) ToHostedConnectionPtrOutput() HostedConnectionPtrOutput {
	return o
}

func (o HostedConnectionPtrOutput) ToHostedConnectionPtrOutputWithContext(ctx context.Context) HostedConnectionPtrOutput {
	return o
}

func (o HostedConnectionPtrOutput) Elem() HostedConnectionOutput {
	return o.ApplyT(func(v *HostedConnection) HostedConnection {
		if v != nil {
			return *v
		}
		var ret HostedConnection
		return ret
	}).(HostedConnectionOutput)
}

type HostedConnectionArrayOutput struct{ *pulumi.OutputState }

func (HostedConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostedConnection)(nil))
}

func (o HostedConnectionArrayOutput) ToHostedConnectionArrayOutput() HostedConnectionArrayOutput {
	return o
}

func (o HostedConnectionArrayOutput) ToHostedConnectionArrayOutputWithContext(ctx context.Context) HostedConnectionArrayOutput {
	return o
}

func (o HostedConnectionArrayOutput) Index(i pulumi.IntInput) HostedConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostedConnection {
		return vs[0].([]HostedConnection)[vs[1].(int)]
	}).(HostedConnectionOutput)
}

type HostedConnectionMapOutput struct{ *pulumi.OutputState }

func (HostedConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]HostedConnection)(nil))
}

func (o HostedConnectionMapOutput) ToHostedConnectionMapOutput() HostedConnectionMapOutput {
	return o
}

func (o HostedConnectionMapOutput) ToHostedConnectionMapOutputWithContext(ctx context.Context) HostedConnectionMapOutput {
	return o
}

func (o HostedConnectionMapOutput) MapIndex(k pulumi.StringInput) HostedConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) HostedConnection {
		return vs[0].(map[string]HostedConnection)[vs[1].(string)]
	}).(HostedConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostedConnectionInput)(nil)).Elem(), &HostedConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostedConnectionPtrInput)(nil)).Elem(), &HostedConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostedConnectionArrayInput)(nil)).Elem(), HostedConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostedConnectionMapInput)(nil)).Elem(), HostedConnectionMap{})
	pulumi.RegisterOutputType(HostedConnectionOutput{})
	pulumi.RegisterOutputType(HostedConnectionPtrOutput{})
	pulumi.RegisterOutputType(HostedConnectionArrayOutput{})
	pulumi.RegisterOutputType(HostedConnectionMapOutput{})
}
