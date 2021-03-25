// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an EC2 Transit Gateway Prefix List Reference.
//
// ## Example Usage
// ### Attachment Routing
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2transitgateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2transitgateway.NewPrefixListReference(ctx, "example", &ec2transitgateway.PrefixListReferenceArgs{
// 			PrefixListId:               pulumi.Any(aws_ec2_managed_prefix_list.Example.Id),
// 			TransitGatewayAttachmentId: pulumi.Any(aws_ec2_transit_gateway_vpc_attachment.Example.Id),
// 			TransitGatewayRouteTableId: pulumi.Any(aws_ec2_transit_gateway.Example.Association_default_route_table_id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Blackhole Routing
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2transitgateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2transitgateway.NewPrefixListReference(ctx, "example", &ec2transitgateway.PrefixListReferenceArgs{
// 			Blackhole:                  pulumi.Bool(true),
// 			PrefixListId:               pulumi.Any(aws_ec2_managed_prefix_list.Example.Id),
// 			TransitGatewayRouteTableId: pulumi.Any(aws_ec2_transit_gateway.Example.Association_default_route_table_id),
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
// `aws_ec2_transit_gateway_prefix_list_reference` can be imported by using the EC2 Transit Gateway Route Table identifier and EC2 Prefix List identifier, separated by an underscore (`_`), e.g. console
//
// ```sh
//  $ pulumi import aws:ec2transitgateway/prefixListReference:PrefixListReference example tgw-rtb-12345678_pl-12345678
// ```
type PrefixListReference struct {
	pulumi.CustomResourceState

	// Indicates whether to drop traffic that matches the Prefix List. Defaults to `false`.
	Blackhole pulumi.BoolPtrOutput `pulumi:"blackhole"`
	// Identifier of EC2 Prefix List.
	PrefixListId      pulumi.StringOutput `pulumi:"prefixListId"`
	PrefixListOwnerId pulumi.StringOutput `pulumi:"prefixListOwnerId"`
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId pulumi.StringPtrOutput `pulumi:"transitGatewayAttachmentId"`
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId pulumi.StringOutput `pulumi:"transitGatewayRouteTableId"`
}

// NewPrefixListReference registers a new resource with the given unique name, arguments, and options.
func NewPrefixListReference(ctx *pulumi.Context,
	name string, args *PrefixListReferenceArgs, opts ...pulumi.ResourceOption) (*PrefixListReference, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrefixListId == nil {
		return nil, errors.New("invalid value for required argument 'PrefixListId'")
	}
	if args.TransitGatewayRouteTableId == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayRouteTableId'")
	}
	var resource PrefixListReference
	err := ctx.RegisterResource("aws:ec2transitgateway/prefixListReference:PrefixListReference", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrefixListReference gets an existing PrefixListReference resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrefixListReference(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrefixListReferenceState, opts ...pulumi.ResourceOption) (*PrefixListReference, error) {
	var resource PrefixListReference
	err := ctx.ReadResource("aws:ec2transitgateway/prefixListReference:PrefixListReference", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrefixListReference resources.
type prefixListReferenceState struct {
	// Indicates whether to drop traffic that matches the Prefix List. Defaults to `false`.
	Blackhole *bool `pulumi:"blackhole"`
	// Identifier of EC2 Prefix List.
	PrefixListId      *string `pulumi:"prefixListId"`
	PrefixListOwnerId *string `pulumi:"prefixListOwnerId"`
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId *string `pulumi:"transitGatewayAttachmentId"`
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId *string `pulumi:"transitGatewayRouteTableId"`
}

type PrefixListReferenceState struct {
	// Indicates whether to drop traffic that matches the Prefix List. Defaults to `false`.
	Blackhole pulumi.BoolPtrInput
	// Identifier of EC2 Prefix List.
	PrefixListId      pulumi.StringPtrInput
	PrefixListOwnerId pulumi.StringPtrInput
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId pulumi.StringPtrInput
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId pulumi.StringPtrInput
}

func (PrefixListReferenceState) ElementType() reflect.Type {
	return reflect.TypeOf((*prefixListReferenceState)(nil)).Elem()
}

type prefixListReferenceArgs struct {
	// Indicates whether to drop traffic that matches the Prefix List. Defaults to `false`.
	Blackhole *bool `pulumi:"blackhole"`
	// Identifier of EC2 Prefix List.
	PrefixListId string `pulumi:"prefixListId"`
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId *string `pulumi:"transitGatewayAttachmentId"`
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId string `pulumi:"transitGatewayRouteTableId"`
}

// The set of arguments for constructing a PrefixListReference resource.
type PrefixListReferenceArgs struct {
	// Indicates whether to drop traffic that matches the Prefix List. Defaults to `false`.
	Blackhole pulumi.BoolPtrInput
	// Identifier of EC2 Prefix List.
	PrefixListId pulumi.StringInput
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId pulumi.StringPtrInput
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId pulumi.StringInput
}

func (PrefixListReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*prefixListReferenceArgs)(nil)).Elem()
}

type PrefixListReferenceInput interface {
	pulumi.Input

	ToPrefixListReferenceOutput() PrefixListReferenceOutput
	ToPrefixListReferenceOutputWithContext(ctx context.Context) PrefixListReferenceOutput
}

func (*PrefixListReference) ElementType() reflect.Type {
	return reflect.TypeOf((*PrefixListReference)(nil))
}

func (i *PrefixListReference) ToPrefixListReferenceOutput() PrefixListReferenceOutput {
	return i.ToPrefixListReferenceOutputWithContext(context.Background())
}

func (i *PrefixListReference) ToPrefixListReferenceOutputWithContext(ctx context.Context) PrefixListReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrefixListReferenceOutput)
}

func (i *PrefixListReference) ToPrefixListReferencePtrOutput() PrefixListReferencePtrOutput {
	return i.ToPrefixListReferencePtrOutputWithContext(context.Background())
}

func (i *PrefixListReference) ToPrefixListReferencePtrOutputWithContext(ctx context.Context) PrefixListReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrefixListReferencePtrOutput)
}

type PrefixListReferencePtrInput interface {
	pulumi.Input

	ToPrefixListReferencePtrOutput() PrefixListReferencePtrOutput
	ToPrefixListReferencePtrOutputWithContext(ctx context.Context) PrefixListReferencePtrOutput
}

type prefixListReferencePtrType PrefixListReferenceArgs

func (*prefixListReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrefixListReference)(nil))
}

func (i *prefixListReferencePtrType) ToPrefixListReferencePtrOutput() PrefixListReferencePtrOutput {
	return i.ToPrefixListReferencePtrOutputWithContext(context.Background())
}

func (i *prefixListReferencePtrType) ToPrefixListReferencePtrOutputWithContext(ctx context.Context) PrefixListReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrefixListReferencePtrOutput)
}

// PrefixListReferenceArrayInput is an input type that accepts PrefixListReferenceArray and PrefixListReferenceArrayOutput values.
// You can construct a concrete instance of `PrefixListReferenceArrayInput` via:
//
//          PrefixListReferenceArray{ PrefixListReferenceArgs{...} }
type PrefixListReferenceArrayInput interface {
	pulumi.Input

	ToPrefixListReferenceArrayOutput() PrefixListReferenceArrayOutput
	ToPrefixListReferenceArrayOutputWithContext(context.Context) PrefixListReferenceArrayOutput
}

type PrefixListReferenceArray []PrefixListReferenceInput

func (PrefixListReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*PrefixListReference)(nil))
}

func (i PrefixListReferenceArray) ToPrefixListReferenceArrayOutput() PrefixListReferenceArrayOutput {
	return i.ToPrefixListReferenceArrayOutputWithContext(context.Background())
}

func (i PrefixListReferenceArray) ToPrefixListReferenceArrayOutputWithContext(ctx context.Context) PrefixListReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrefixListReferenceArrayOutput)
}

// PrefixListReferenceMapInput is an input type that accepts PrefixListReferenceMap and PrefixListReferenceMapOutput values.
// You can construct a concrete instance of `PrefixListReferenceMapInput` via:
//
//          PrefixListReferenceMap{ "key": PrefixListReferenceArgs{...} }
type PrefixListReferenceMapInput interface {
	pulumi.Input

	ToPrefixListReferenceMapOutput() PrefixListReferenceMapOutput
	ToPrefixListReferenceMapOutputWithContext(context.Context) PrefixListReferenceMapOutput
}

type PrefixListReferenceMap map[string]PrefixListReferenceInput

func (PrefixListReferenceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*PrefixListReference)(nil))
}

func (i PrefixListReferenceMap) ToPrefixListReferenceMapOutput() PrefixListReferenceMapOutput {
	return i.ToPrefixListReferenceMapOutputWithContext(context.Background())
}

func (i PrefixListReferenceMap) ToPrefixListReferenceMapOutputWithContext(ctx context.Context) PrefixListReferenceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrefixListReferenceMapOutput)
}

type PrefixListReferenceOutput struct {
	*pulumi.OutputState
}

func (PrefixListReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrefixListReference)(nil))
}

func (o PrefixListReferenceOutput) ToPrefixListReferenceOutput() PrefixListReferenceOutput {
	return o
}

func (o PrefixListReferenceOutput) ToPrefixListReferenceOutputWithContext(ctx context.Context) PrefixListReferenceOutput {
	return o
}

func (o PrefixListReferenceOutput) ToPrefixListReferencePtrOutput() PrefixListReferencePtrOutput {
	return o.ToPrefixListReferencePtrOutputWithContext(context.Background())
}

func (o PrefixListReferenceOutput) ToPrefixListReferencePtrOutputWithContext(ctx context.Context) PrefixListReferencePtrOutput {
	return o.ApplyT(func(v PrefixListReference) *PrefixListReference {
		return &v
	}).(PrefixListReferencePtrOutput)
}

type PrefixListReferencePtrOutput struct {
	*pulumi.OutputState
}

func (PrefixListReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrefixListReference)(nil))
}

func (o PrefixListReferencePtrOutput) ToPrefixListReferencePtrOutput() PrefixListReferencePtrOutput {
	return o
}

func (o PrefixListReferencePtrOutput) ToPrefixListReferencePtrOutputWithContext(ctx context.Context) PrefixListReferencePtrOutput {
	return o
}

type PrefixListReferenceArrayOutput struct{ *pulumi.OutputState }

func (PrefixListReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrefixListReference)(nil))
}

func (o PrefixListReferenceArrayOutput) ToPrefixListReferenceArrayOutput() PrefixListReferenceArrayOutput {
	return o
}

func (o PrefixListReferenceArrayOutput) ToPrefixListReferenceArrayOutputWithContext(ctx context.Context) PrefixListReferenceArrayOutput {
	return o
}

func (o PrefixListReferenceArrayOutput) Index(i pulumi.IntInput) PrefixListReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrefixListReference {
		return vs[0].([]PrefixListReference)[vs[1].(int)]
	}).(PrefixListReferenceOutput)
}

type PrefixListReferenceMapOutput struct{ *pulumi.OutputState }

func (PrefixListReferenceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PrefixListReference)(nil))
}

func (o PrefixListReferenceMapOutput) ToPrefixListReferenceMapOutput() PrefixListReferenceMapOutput {
	return o
}

func (o PrefixListReferenceMapOutput) ToPrefixListReferenceMapOutputWithContext(ctx context.Context) PrefixListReferenceMapOutput {
	return o
}

func (o PrefixListReferenceMapOutput) MapIndex(k pulumi.StringInput) PrefixListReferenceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PrefixListReference {
		return vs[0].(map[string]PrefixListReference)[vs[1].(string)]
	}).(PrefixListReferenceOutput)
}

func init() {
	pulumi.RegisterOutputType(PrefixListReferenceOutput{})
	pulumi.RegisterOutputType(PrefixListReferencePtrOutput{})
	pulumi.RegisterOutputType(PrefixListReferenceArrayOutput{})
	pulumi.RegisterOutputType(PrefixListReferenceMapOutput{})
}
