// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Route53 Hosted Zone. For managing Domain Name System Security Extensions (DNSSEC), see the `route53.KeySigningKey` and `route53.HostedZoneDnsSec` resources.
//
// ## Example Usage
// ### Public Zone
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := route53.NewZone(ctx, "primary", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Public Subdomain Zone
//
// For use in subdomains, note that you need to create a
// `route53.Record` of type `NS` as well as the subdomain
// zone.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		main, err := route53.NewZone(ctx, "main", nil)
// 		if err != nil {
// 			return err
// 		}
// 		dev, err := route53.NewZone(ctx, "dev", &route53.ZoneArgs{
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("dev"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewRecord(ctx, "dev_ns", &route53.RecordArgs{
// 			ZoneId:  main.ZoneId,
// 			Name:    pulumi.String("dev.example.com"),
// 			Type:    pulumi.String("NS"),
// 			Ttl:     pulumi.Int(30),
// 			Records: dev.NameServers,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Private Zone
//
// > **NOTE:** This provider provides both exclusive VPC associations defined in-line in this resource via `vpc` configuration blocks and a separate ` Zone VPC Association resource. At this time, you cannot use in-line VPC associations in conjunction with any  `route53.ZoneAssociation`  resources with the same zone ID otherwise it will cause a perpetual difference in plan output. You can optionally use [ `ignoreChanges` ](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to manage additional associations via the  `route53.ZoneAssociation` resource.
//
// > **NOTE:** Private zones require at least one VPC association at all times.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := route53.NewZone(ctx, "private", &route53.ZoneArgs{
// 			Vpcs: route53.ZoneVpcArray{
// 				&route53.ZoneVpcArgs{
// 					VpcId: pulumi.Any(aws_vpc.Example.Id),
// 				},
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
// Route53 Zones can be imported using the `zone id`, e.g.,
//
// ```sh
//  $ pulumi import aws:route53/zone:Zone myzone Z1D633PJN98FT9
// ```
type Zone struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the Hosted Zone.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A comment for the hosted zone. Defaults to 'Managed by Pulumi'.
	Comment pulumi.StringOutput `pulumi:"comment"`
	// The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
	DelegationSetId pulumi.StringPtrOutput `pulumi:"delegationSetId"`
	// Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// This is the name of the hosted zone.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of name servers in associated (or default) delegation set.
	// Find more about delegation sets in [AWS docs](https://docs.aws.amazon.com/Route53/latest/APIReference/actions-on-reusable-delegation-sets.html).
	NameServers pulumi.StringArrayOutput `pulumi:"nameServers"`
	// A mapping of tags to assign to the zone.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegationSetId` argument in this resource and any `route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
	Vpcs ZoneVpcArrayOutput `pulumi:"vpcs"`
	// The Hosted Zone ID. This can be referenced by zone records.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewZone registers a new resource with the given unique name, arguments, and options.
func NewZone(ctx *pulumi.Context,
	name string, args *ZoneArgs, opts ...pulumi.ResourceOption) (*Zone, error) {
	if args == nil {
		args = &ZoneArgs{}
	}

	if args.Comment == nil {
		args.Comment = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource Zone
	err := ctx.RegisterResource("aws:route53/zone:Zone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZone gets an existing Zone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneState, opts ...pulumi.ResourceOption) (*Zone, error) {
	var resource Zone
	err := ctx.ReadResource("aws:route53/zone:Zone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Zone resources.
type zoneState struct {
	// The Amazon Resource Name (ARN) of the Hosted Zone.
	Arn *string `pulumi:"arn"`
	// A comment for the hosted zone. Defaults to 'Managed by Pulumi'.
	Comment *string `pulumi:"comment"`
	// The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
	DelegationSetId *string `pulumi:"delegationSetId"`
	// Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// This is the name of the hosted zone.
	Name *string `pulumi:"name"`
	// A list of name servers in associated (or default) delegation set.
	// Find more about delegation sets in [AWS docs](https://docs.aws.amazon.com/Route53/latest/APIReference/actions-on-reusable-delegation-sets.html).
	NameServers []string `pulumi:"nameServers"`
	// A mapping of tags to assign to the zone.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegationSetId` argument in this resource and any `route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
	Vpcs []ZoneVpc `pulumi:"vpcs"`
	// The Hosted Zone ID. This can be referenced by zone records.
	ZoneId *string `pulumi:"zoneId"`
}

type ZoneState struct {
	// The Amazon Resource Name (ARN) of the Hosted Zone.
	Arn pulumi.StringPtrInput
	// A comment for the hosted zone. Defaults to 'Managed by Pulumi'.
	Comment pulumi.StringPtrInput
	// The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
	DelegationSetId pulumi.StringPtrInput
	// Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
	ForceDestroy pulumi.BoolPtrInput
	// This is the name of the hosted zone.
	Name pulumi.StringPtrInput
	// A list of name servers in associated (or default) delegation set.
	// Find more about delegation sets in [AWS docs](https://docs.aws.amazon.com/Route53/latest/APIReference/actions-on-reusable-delegation-sets.html).
	NameServers pulumi.StringArrayInput
	// A mapping of tags to assign to the zone.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegationSetId` argument in this resource and any `route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
	Vpcs ZoneVpcArrayInput
	// The Hosted Zone ID. This can be referenced by zone records.
	ZoneId pulumi.StringPtrInput
}

func (ZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneState)(nil)).Elem()
}

type zoneArgs struct {
	// A comment for the hosted zone. Defaults to 'Managed by Pulumi'.
	Comment *string `pulumi:"comment"`
	// The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
	DelegationSetId *string `pulumi:"delegationSetId"`
	// Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// This is the name of the hosted zone.
	Name *string `pulumi:"name"`
	// A mapping of tags to assign to the zone.
	Tags map[string]string `pulumi:"tags"`
	// Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegationSetId` argument in this resource and any `route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
	Vpcs []ZoneVpc `pulumi:"vpcs"`
}

// The set of arguments for constructing a Zone resource.
type ZoneArgs struct {
	// A comment for the hosted zone. Defaults to 'Managed by Pulumi'.
	Comment pulumi.StringPtrInput
	// The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
	DelegationSetId pulumi.StringPtrInput
	// Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
	ForceDestroy pulumi.BoolPtrInput
	// This is the name of the hosted zone.
	Name pulumi.StringPtrInput
	// A mapping of tags to assign to the zone.
	Tags pulumi.StringMapInput
	// Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegationSetId` argument in this resource and any `route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
	Vpcs ZoneVpcArrayInput
}

func (ZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneArgs)(nil)).Elem()
}

type ZoneInput interface {
	pulumi.Input

	ToZoneOutput() ZoneOutput
	ToZoneOutputWithContext(ctx context.Context) ZoneOutput
}

func (*Zone) ElementType() reflect.Type {
	return reflect.TypeOf((*Zone)(nil))
}

func (i *Zone) ToZoneOutput() ZoneOutput {
	return i.ToZoneOutputWithContext(context.Background())
}

func (i *Zone) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneOutput)
}

func (i *Zone) ToZonePtrOutput() ZonePtrOutput {
	return i.ToZonePtrOutputWithContext(context.Background())
}

func (i *Zone) ToZonePtrOutputWithContext(ctx context.Context) ZonePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZonePtrOutput)
}

type ZonePtrInput interface {
	pulumi.Input

	ToZonePtrOutput() ZonePtrOutput
	ToZonePtrOutputWithContext(ctx context.Context) ZonePtrOutput
}

type zonePtrType ZoneArgs

func (*zonePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil))
}

func (i *zonePtrType) ToZonePtrOutput() ZonePtrOutput {
	return i.ToZonePtrOutputWithContext(context.Background())
}

func (i *zonePtrType) ToZonePtrOutputWithContext(ctx context.Context) ZonePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZonePtrOutput)
}

// ZoneArrayInput is an input type that accepts ZoneArray and ZoneArrayOutput values.
// You can construct a concrete instance of `ZoneArrayInput` via:
//
//          ZoneArray{ ZoneArgs{...} }
type ZoneArrayInput interface {
	pulumi.Input

	ToZoneArrayOutput() ZoneArrayOutput
	ToZoneArrayOutputWithContext(context.Context) ZoneArrayOutput
}

type ZoneArray []ZoneInput

func (ZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Zone)(nil)).Elem()
}

func (i ZoneArray) ToZoneArrayOutput() ZoneArrayOutput {
	return i.ToZoneArrayOutputWithContext(context.Background())
}

func (i ZoneArray) ToZoneArrayOutputWithContext(ctx context.Context) ZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneArrayOutput)
}

// ZoneMapInput is an input type that accepts ZoneMap and ZoneMapOutput values.
// You can construct a concrete instance of `ZoneMapInput` via:
//
//          ZoneMap{ "key": ZoneArgs{...} }
type ZoneMapInput interface {
	pulumi.Input

	ToZoneMapOutput() ZoneMapOutput
	ToZoneMapOutputWithContext(context.Context) ZoneMapOutput
}

type ZoneMap map[string]ZoneInput

func (ZoneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Zone)(nil)).Elem()
}

func (i ZoneMap) ToZoneMapOutput() ZoneMapOutput {
	return i.ToZoneMapOutputWithContext(context.Background())
}

func (i ZoneMap) ToZoneMapOutputWithContext(ctx context.Context) ZoneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneMapOutput)
}

type ZoneOutput struct{ *pulumi.OutputState }

func (ZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Zone)(nil))
}

func (o ZoneOutput) ToZoneOutput() ZoneOutput {
	return o
}

func (o ZoneOutput) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return o
}

func (o ZoneOutput) ToZonePtrOutput() ZonePtrOutput {
	return o.ToZonePtrOutputWithContext(context.Background())
}

func (o ZoneOutput) ToZonePtrOutputWithContext(ctx context.Context) ZonePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Zone) *Zone {
		return &v
	}).(ZonePtrOutput)
}

type ZonePtrOutput struct{ *pulumi.OutputState }

func (ZonePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil))
}

func (o ZonePtrOutput) ToZonePtrOutput() ZonePtrOutput {
	return o
}

func (o ZonePtrOutput) ToZonePtrOutputWithContext(ctx context.Context) ZonePtrOutput {
	return o
}

func (o ZonePtrOutput) Elem() ZoneOutput {
	return o.ApplyT(func(v *Zone) Zone {
		if v != nil {
			return *v
		}
		var ret Zone
		return ret
	}).(ZoneOutput)
}

type ZoneArrayOutput struct{ *pulumi.OutputState }

func (ZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Zone)(nil))
}

func (o ZoneArrayOutput) ToZoneArrayOutput() ZoneArrayOutput {
	return o
}

func (o ZoneArrayOutput) ToZoneArrayOutputWithContext(ctx context.Context) ZoneArrayOutput {
	return o
}

func (o ZoneArrayOutput) Index(i pulumi.IntInput) ZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Zone {
		return vs[0].([]Zone)[vs[1].(int)]
	}).(ZoneOutput)
}

type ZoneMapOutput struct{ *pulumi.OutputState }

func (ZoneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Zone)(nil))
}

func (o ZoneMapOutput) ToZoneMapOutput() ZoneMapOutput {
	return o
}

func (o ZoneMapOutput) ToZoneMapOutputWithContext(ctx context.Context) ZoneMapOutput {
	return o
}

func (o ZoneMapOutput) MapIndex(k pulumi.StringInput) ZoneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Zone {
		return vs[0].(map[string]Zone)[vs[1].(string)]
	}).(ZoneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneInput)(nil)).Elem(), &Zone{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZonePtrInput)(nil)).Elem(), &Zone{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneArrayInput)(nil)).Elem(), ZoneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneMapInput)(nil)).Elem(), ZoneMap{})
	pulumi.RegisterOutputType(ZoneOutput{})
	pulumi.RegisterOutputType(ZonePtrOutput{})
	pulumi.RegisterOutputType(ZoneArrayOutput{})
	pulumi.RegisterOutputType(ZoneMapOutput{})
}
