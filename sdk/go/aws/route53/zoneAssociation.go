// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Route53 Hosted Zone VPC association. VPC associations can only be made on private zones. See the `route53.VpcAssociationAuthorization` resource for setting up cross-account associations.
//
// > **NOTE:** Unless explicit association ordering is required (e.g. a separate cross-account association authorization), usage of this resource is not recommended. Use the `vpc` configuration blocks available within the `route53.Zone` resource instead.
//
// > **NOTE:** This provider provides both this standalone Zone VPC Association resource and exclusive VPC associations defined in-line in the `route53.Zone` resource via `vpc` configuration blocks. At this time, you cannot use those in-line VPC associations in conjunction with this resource and the same zone ID otherwise it will cause a perpetual difference in plan output. You can optionally use [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) in the `route53.Zone` resource to manage additional associations via this resource.
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
// 		primary, err := ec2.NewVpc(ctx, "primary", &ec2.VpcArgs{
// 			CidrBlock:          pulumi.String("10.6.0.0/16"),
// 			EnableDnsHostnames: pulumi.Bool(true),
// 			EnableDnsSupport:   pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		secondaryVpc, err := ec2.NewVpc(ctx, "secondaryVpc", &ec2.VpcArgs{
// 			CidrBlock:          pulumi.String("10.7.0.0/16"),
// 			EnableDnsHostnames: pulumi.Bool(true),
// 			EnableDnsSupport:   pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		example, err := route53.NewZone(ctx, "example", &route53.ZoneArgs{
// 			Vpcs: route53.ZoneVpcArray{
// 				&route53.ZoneVpcArgs{
// 					VpcId: primary.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewZoneAssociation(ctx, "secondaryZoneAssociation", &route53.ZoneAssociationArgs{
// 			ZoneId: example.ZoneId,
// 			VpcId:  secondaryVpc.ID(),
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
// Route 53 Hosted Zone Associations can be imported via the Hosted Zone ID and VPC ID, separated by a colon (`:`), e.g.
//
// ```sh
//  $ pulumi import aws:route53/zoneAssociation:ZoneAssociation example Z123456ABCDEFG:vpc-12345678
// ```
//
//  If the VPC is in a different region than the provider region configuration, the VPC Region can be added to the end. e.g.
//
// ```sh
//  $ pulumi import aws:route53/zoneAssociation:ZoneAssociation example Z123456ABCDEFG:vpc-12345678:us-east-2
// ```
type ZoneAssociation struct {
	pulumi.CustomResourceState

	// The account ID of the account that created the hosted zone.
	OwningAccount pulumi.StringOutput `pulumi:"owningAccount"`
	// The VPC to associate with the private hosted zone.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion pulumi.StringOutput `pulumi:"vpcRegion"`
	// The private hosted zone to associate.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewZoneAssociation registers a new resource with the given unique name, arguments, and options.
func NewZoneAssociation(ctx *pulumi.Context,
	name string, args *ZoneAssociationArgs, opts ...pulumi.ResourceOption) (*ZoneAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	var resource ZoneAssociation
	err := ctx.RegisterResource("aws:route53/zoneAssociation:ZoneAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZoneAssociation gets an existing ZoneAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZoneAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneAssociationState, opts ...pulumi.ResourceOption) (*ZoneAssociation, error) {
	var resource ZoneAssociation
	err := ctx.ReadResource("aws:route53/zoneAssociation:ZoneAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZoneAssociation resources.
type zoneAssociationState struct {
	// The account ID of the account that created the hosted zone.
	OwningAccount *string `pulumi:"owningAccount"`
	// The VPC to associate with the private hosted zone.
	VpcId *string `pulumi:"vpcId"`
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion *string `pulumi:"vpcRegion"`
	// The private hosted zone to associate.
	ZoneId *string `pulumi:"zoneId"`
}

type ZoneAssociationState struct {
	// The account ID of the account that created the hosted zone.
	OwningAccount pulumi.StringPtrInput
	// The VPC to associate with the private hosted zone.
	VpcId pulumi.StringPtrInput
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion pulumi.StringPtrInput
	// The private hosted zone to associate.
	ZoneId pulumi.StringPtrInput
}

func (ZoneAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneAssociationState)(nil)).Elem()
}

type zoneAssociationArgs struct {
	// The VPC to associate with the private hosted zone.
	VpcId string `pulumi:"vpcId"`
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion *string `pulumi:"vpcRegion"`
	// The private hosted zone to associate.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a ZoneAssociation resource.
type ZoneAssociationArgs struct {
	// The VPC to associate with the private hosted zone.
	VpcId pulumi.StringInput
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion pulumi.StringPtrInput
	// The private hosted zone to associate.
	ZoneId pulumi.StringInput
}

func (ZoneAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneAssociationArgs)(nil)).Elem()
}

type ZoneAssociationInput interface {
	pulumi.Input

	ToZoneAssociationOutput() ZoneAssociationOutput
	ToZoneAssociationOutputWithContext(ctx context.Context) ZoneAssociationOutput
}

func (*ZoneAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneAssociation)(nil))
}

func (i *ZoneAssociation) ToZoneAssociationOutput() ZoneAssociationOutput {
	return i.ToZoneAssociationOutputWithContext(context.Background())
}

func (i *ZoneAssociation) ToZoneAssociationOutputWithContext(ctx context.Context) ZoneAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneAssociationOutput)
}

func (i *ZoneAssociation) ToZoneAssociationPtrOutput() ZoneAssociationPtrOutput {
	return i.ToZoneAssociationPtrOutputWithContext(context.Background())
}

func (i *ZoneAssociation) ToZoneAssociationPtrOutputWithContext(ctx context.Context) ZoneAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneAssociationPtrOutput)
}

type ZoneAssociationPtrInput interface {
	pulumi.Input

	ToZoneAssociationPtrOutput() ZoneAssociationPtrOutput
	ToZoneAssociationPtrOutputWithContext(ctx context.Context) ZoneAssociationPtrOutput
}

type zoneAssociationPtrType ZoneAssociationArgs

func (*zoneAssociationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneAssociation)(nil))
}

func (i *zoneAssociationPtrType) ToZoneAssociationPtrOutput() ZoneAssociationPtrOutput {
	return i.ToZoneAssociationPtrOutputWithContext(context.Background())
}

func (i *zoneAssociationPtrType) ToZoneAssociationPtrOutputWithContext(ctx context.Context) ZoneAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneAssociationPtrOutput)
}

// ZoneAssociationArrayInput is an input type that accepts ZoneAssociationArray and ZoneAssociationArrayOutput values.
// You can construct a concrete instance of `ZoneAssociationArrayInput` via:
//
//          ZoneAssociationArray{ ZoneAssociationArgs{...} }
type ZoneAssociationArrayInput interface {
	pulumi.Input

	ToZoneAssociationArrayOutput() ZoneAssociationArrayOutput
	ToZoneAssociationArrayOutputWithContext(context.Context) ZoneAssociationArrayOutput
}

type ZoneAssociationArray []ZoneAssociationInput

func (ZoneAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ZoneAssociation)(nil))
}

func (i ZoneAssociationArray) ToZoneAssociationArrayOutput() ZoneAssociationArrayOutput {
	return i.ToZoneAssociationArrayOutputWithContext(context.Background())
}

func (i ZoneAssociationArray) ToZoneAssociationArrayOutputWithContext(ctx context.Context) ZoneAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneAssociationArrayOutput)
}

// ZoneAssociationMapInput is an input type that accepts ZoneAssociationMap and ZoneAssociationMapOutput values.
// You can construct a concrete instance of `ZoneAssociationMapInput` via:
//
//          ZoneAssociationMap{ "key": ZoneAssociationArgs{...} }
type ZoneAssociationMapInput interface {
	pulumi.Input

	ToZoneAssociationMapOutput() ZoneAssociationMapOutput
	ToZoneAssociationMapOutputWithContext(context.Context) ZoneAssociationMapOutput
}

type ZoneAssociationMap map[string]ZoneAssociationInput

func (ZoneAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ZoneAssociation)(nil))
}

func (i ZoneAssociationMap) ToZoneAssociationMapOutput() ZoneAssociationMapOutput {
	return i.ToZoneAssociationMapOutputWithContext(context.Background())
}

func (i ZoneAssociationMap) ToZoneAssociationMapOutputWithContext(ctx context.Context) ZoneAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneAssociationMapOutput)
}

type ZoneAssociationOutput struct {
	*pulumi.OutputState
}

func (ZoneAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneAssociation)(nil))
}

func (o ZoneAssociationOutput) ToZoneAssociationOutput() ZoneAssociationOutput {
	return o
}

func (o ZoneAssociationOutput) ToZoneAssociationOutputWithContext(ctx context.Context) ZoneAssociationOutput {
	return o
}

func (o ZoneAssociationOutput) ToZoneAssociationPtrOutput() ZoneAssociationPtrOutput {
	return o.ToZoneAssociationPtrOutputWithContext(context.Background())
}

func (o ZoneAssociationOutput) ToZoneAssociationPtrOutputWithContext(ctx context.Context) ZoneAssociationPtrOutput {
	return o.ApplyT(func(v ZoneAssociation) *ZoneAssociation {
		return &v
	}).(ZoneAssociationPtrOutput)
}

type ZoneAssociationPtrOutput struct {
	*pulumi.OutputState
}

func (ZoneAssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneAssociation)(nil))
}

func (o ZoneAssociationPtrOutput) ToZoneAssociationPtrOutput() ZoneAssociationPtrOutput {
	return o
}

func (o ZoneAssociationPtrOutput) ToZoneAssociationPtrOutputWithContext(ctx context.Context) ZoneAssociationPtrOutput {
	return o
}

type ZoneAssociationArrayOutput struct{ *pulumi.OutputState }

func (ZoneAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ZoneAssociation)(nil))
}

func (o ZoneAssociationArrayOutput) ToZoneAssociationArrayOutput() ZoneAssociationArrayOutput {
	return o
}

func (o ZoneAssociationArrayOutput) ToZoneAssociationArrayOutputWithContext(ctx context.Context) ZoneAssociationArrayOutput {
	return o
}

func (o ZoneAssociationArrayOutput) Index(i pulumi.IntInput) ZoneAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ZoneAssociation {
		return vs[0].([]ZoneAssociation)[vs[1].(int)]
	}).(ZoneAssociationOutput)
}

type ZoneAssociationMapOutput struct{ *pulumi.OutputState }

func (ZoneAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ZoneAssociation)(nil))
}

func (o ZoneAssociationMapOutput) ToZoneAssociationMapOutput() ZoneAssociationMapOutput {
	return o
}

func (o ZoneAssociationMapOutput) ToZoneAssociationMapOutputWithContext(ctx context.Context) ZoneAssociationMapOutput {
	return o
}

func (o ZoneAssociationMapOutput) MapIndex(k pulumi.StringInput) ZoneAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ZoneAssociation {
		return vs[0].(map[string]ZoneAssociation)[vs[1].(string)]
	}).(ZoneAssociationOutput)
}

func init() {
	pulumi.RegisterOutputType(ZoneAssociationOutput{})
	pulumi.RegisterOutputType(ZoneAssociationPtrOutput{})
	pulumi.RegisterOutputType(ZoneAssociationArrayOutput{})
	pulumi.RegisterOutputType(ZoneAssociationMapOutput{})
}
