// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an EC2 Transit Gateway Route Table propagation.
//
// ## Example Usage
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
// 		_, err := ec2transitgateway.NewRouteTablePropagation(ctx, "example", &ec2transitgateway.RouteTablePropagationArgs{
// 			TransitGatewayAttachmentId: pulumi.Any(aws_ec2_transit_gateway_vpc_attachment.Example.Id),
// 			TransitGatewayRouteTableId: pulumi.Any(aws_ec2_transit_gateway_route_table.Example.Id),
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
// `aws_ec2_transit_gateway_route_table_propagation` can be imported by using the EC2 Transit Gateway Route Table identifier, an underscore, and the EC2 Transit Gateway Attachment identifier, e.g.
//
// ```sh
//  $ pulumi import aws:ec2transitgateway/routeTablePropagation:RouteTablePropagation example tgw-rtb-12345678_tgw-attach-87654321
// ```
type RouteTablePropagation struct {
	pulumi.CustomResourceState

	// Identifier of the resource
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// Type of the resource
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId pulumi.StringOutput `pulumi:"transitGatewayAttachmentId"`
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId pulumi.StringOutput `pulumi:"transitGatewayRouteTableId"`
}

// NewRouteTablePropagation registers a new resource with the given unique name, arguments, and options.
func NewRouteTablePropagation(ctx *pulumi.Context,
	name string, args *RouteTablePropagationArgs, opts ...pulumi.ResourceOption) (*RouteTablePropagation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TransitGatewayAttachmentId == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayAttachmentId'")
	}
	if args.TransitGatewayRouteTableId == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayRouteTableId'")
	}
	var resource RouteTablePropagation
	err := ctx.RegisterResource("aws:ec2transitgateway/routeTablePropagation:RouteTablePropagation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteTablePropagation gets an existing RouteTablePropagation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteTablePropagation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteTablePropagationState, opts ...pulumi.ResourceOption) (*RouteTablePropagation, error) {
	var resource RouteTablePropagation
	err := ctx.ReadResource("aws:ec2transitgateway/routeTablePropagation:RouteTablePropagation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteTablePropagation resources.
type routeTablePropagationState struct {
	// Identifier of the resource
	ResourceId *string `pulumi:"resourceId"`
	// Type of the resource
	ResourceType *string `pulumi:"resourceType"`
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId *string `pulumi:"transitGatewayAttachmentId"`
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId *string `pulumi:"transitGatewayRouteTableId"`
}

type RouteTablePropagationState struct {
	// Identifier of the resource
	ResourceId pulumi.StringPtrInput
	// Type of the resource
	ResourceType pulumi.StringPtrInput
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId pulumi.StringPtrInput
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId pulumi.StringPtrInput
}

func (RouteTablePropagationState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTablePropagationState)(nil)).Elem()
}

type routeTablePropagationArgs struct {
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId string `pulumi:"transitGatewayAttachmentId"`
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId string `pulumi:"transitGatewayRouteTableId"`
}

// The set of arguments for constructing a RouteTablePropagation resource.
type RouteTablePropagationArgs struct {
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId pulumi.StringInput
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId pulumi.StringInput
}

func (RouteTablePropagationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTablePropagationArgs)(nil)).Elem()
}

type RouteTablePropagationInput interface {
	pulumi.Input

	ToRouteTablePropagationOutput() RouteTablePropagationOutput
	ToRouteTablePropagationOutputWithContext(ctx context.Context) RouteTablePropagationOutput
}

func (*RouteTablePropagation) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTablePropagation)(nil))
}

func (i *RouteTablePropagation) ToRouteTablePropagationOutput() RouteTablePropagationOutput {
	return i.ToRouteTablePropagationOutputWithContext(context.Background())
}

func (i *RouteTablePropagation) ToRouteTablePropagationOutputWithContext(ctx context.Context) RouteTablePropagationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTablePropagationOutput)
}

func (i *RouteTablePropagation) ToRouteTablePropagationPtrOutput() RouteTablePropagationPtrOutput {
	return i.ToRouteTablePropagationPtrOutputWithContext(context.Background())
}

func (i *RouteTablePropagation) ToRouteTablePropagationPtrOutputWithContext(ctx context.Context) RouteTablePropagationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTablePropagationPtrOutput)
}

type RouteTablePropagationPtrInput interface {
	pulumi.Input

	ToRouteTablePropagationPtrOutput() RouteTablePropagationPtrOutput
	ToRouteTablePropagationPtrOutputWithContext(ctx context.Context) RouteTablePropagationPtrOutput
}

type routeTablePropagationPtrType RouteTablePropagationArgs

func (*routeTablePropagationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTablePropagation)(nil))
}

func (i *routeTablePropagationPtrType) ToRouteTablePropagationPtrOutput() RouteTablePropagationPtrOutput {
	return i.ToRouteTablePropagationPtrOutputWithContext(context.Background())
}

func (i *routeTablePropagationPtrType) ToRouteTablePropagationPtrOutputWithContext(ctx context.Context) RouteTablePropagationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTablePropagationPtrOutput)
}

// RouteTablePropagationArrayInput is an input type that accepts RouteTablePropagationArray and RouteTablePropagationArrayOutput values.
// You can construct a concrete instance of `RouteTablePropagationArrayInput` via:
//
//          RouteTablePropagationArray{ RouteTablePropagationArgs{...} }
type RouteTablePropagationArrayInput interface {
	pulumi.Input

	ToRouteTablePropagationArrayOutput() RouteTablePropagationArrayOutput
	ToRouteTablePropagationArrayOutputWithContext(context.Context) RouteTablePropagationArrayOutput
}

type RouteTablePropagationArray []RouteTablePropagationInput

func (RouteTablePropagationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteTablePropagation)(nil)).Elem()
}

func (i RouteTablePropagationArray) ToRouteTablePropagationArrayOutput() RouteTablePropagationArrayOutput {
	return i.ToRouteTablePropagationArrayOutputWithContext(context.Background())
}

func (i RouteTablePropagationArray) ToRouteTablePropagationArrayOutputWithContext(ctx context.Context) RouteTablePropagationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTablePropagationArrayOutput)
}

// RouteTablePropagationMapInput is an input type that accepts RouteTablePropagationMap and RouteTablePropagationMapOutput values.
// You can construct a concrete instance of `RouteTablePropagationMapInput` via:
//
//          RouteTablePropagationMap{ "key": RouteTablePropagationArgs{...} }
type RouteTablePropagationMapInput interface {
	pulumi.Input

	ToRouteTablePropagationMapOutput() RouteTablePropagationMapOutput
	ToRouteTablePropagationMapOutputWithContext(context.Context) RouteTablePropagationMapOutput
}

type RouteTablePropagationMap map[string]RouteTablePropagationInput

func (RouteTablePropagationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteTablePropagation)(nil)).Elem()
}

func (i RouteTablePropagationMap) ToRouteTablePropagationMapOutput() RouteTablePropagationMapOutput {
	return i.ToRouteTablePropagationMapOutputWithContext(context.Background())
}

func (i RouteTablePropagationMap) ToRouteTablePropagationMapOutputWithContext(ctx context.Context) RouteTablePropagationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTablePropagationMapOutput)
}

type RouteTablePropagationOutput struct {
	*pulumi.OutputState
}

func (RouteTablePropagationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTablePropagation)(nil))
}

func (o RouteTablePropagationOutput) ToRouteTablePropagationOutput() RouteTablePropagationOutput {
	return o
}

func (o RouteTablePropagationOutput) ToRouteTablePropagationOutputWithContext(ctx context.Context) RouteTablePropagationOutput {
	return o
}

func (o RouteTablePropagationOutput) ToRouteTablePropagationPtrOutput() RouteTablePropagationPtrOutput {
	return o.ToRouteTablePropagationPtrOutputWithContext(context.Background())
}

func (o RouteTablePropagationOutput) ToRouteTablePropagationPtrOutputWithContext(ctx context.Context) RouteTablePropagationPtrOutput {
	return o.ApplyT(func(v RouteTablePropagation) *RouteTablePropagation {
		return &v
	}).(RouteTablePropagationPtrOutput)
}

type RouteTablePropagationPtrOutput struct {
	*pulumi.OutputState
}

func (RouteTablePropagationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTablePropagation)(nil))
}

func (o RouteTablePropagationPtrOutput) ToRouteTablePropagationPtrOutput() RouteTablePropagationPtrOutput {
	return o
}

func (o RouteTablePropagationPtrOutput) ToRouteTablePropagationPtrOutputWithContext(ctx context.Context) RouteTablePropagationPtrOutput {
	return o
}

type RouteTablePropagationArrayOutput struct{ *pulumi.OutputState }

func (RouteTablePropagationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteTablePropagation)(nil))
}

func (o RouteTablePropagationArrayOutput) ToRouteTablePropagationArrayOutput() RouteTablePropagationArrayOutput {
	return o
}

func (o RouteTablePropagationArrayOutput) ToRouteTablePropagationArrayOutputWithContext(ctx context.Context) RouteTablePropagationArrayOutput {
	return o
}

func (o RouteTablePropagationArrayOutput) Index(i pulumi.IntInput) RouteTablePropagationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RouteTablePropagation {
		return vs[0].([]RouteTablePropagation)[vs[1].(int)]
	}).(RouteTablePropagationOutput)
}

type RouteTablePropagationMapOutput struct{ *pulumi.OutputState }

func (RouteTablePropagationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RouteTablePropagation)(nil))
}

func (o RouteTablePropagationMapOutput) ToRouteTablePropagationMapOutput() RouteTablePropagationMapOutput {
	return o
}

func (o RouteTablePropagationMapOutput) ToRouteTablePropagationMapOutputWithContext(ctx context.Context) RouteTablePropagationMapOutput {
	return o
}

func (o RouteTablePropagationMapOutput) MapIndex(k pulumi.StringInput) RouteTablePropagationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RouteTablePropagation {
		return vs[0].(map[string]RouteTablePropagation)[vs[1].(string)]
	}).(RouteTablePropagationOutput)
}

func init() {
	pulumi.RegisterOutputType(RouteTablePropagationOutput{})
	pulumi.RegisterOutputType(RouteTablePropagationPtrOutput{})
	pulumi.RegisterOutputType(RouteTablePropagationArrayOutput{})
	pulumi.RegisterOutputType(RouteTablePropagationMapOutput{})
}
