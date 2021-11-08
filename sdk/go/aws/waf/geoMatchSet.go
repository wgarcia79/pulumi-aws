// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a WAF Geo Match Set Resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/waf"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := waf.NewGeoMatchSet(ctx, "geoMatchSet", &waf.GeoMatchSetArgs{
// 			GeoMatchConstraints: waf.GeoMatchSetGeoMatchConstraintArray{
// 				&waf.GeoMatchSetGeoMatchConstraintArgs{
// 					Type:  pulumi.String("Country"),
// 					Value: pulumi.String("US"),
// 				},
// 				&waf.GeoMatchSetGeoMatchConstraintArgs{
// 					Type:  pulumi.String("Country"),
// 					Value: pulumi.String("CA"),
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
// WAF Geo Match Set can be imported using their ID, e.g.,
//
// ```sh
//  $ pulumi import aws:waf/geoMatchSet:GeoMatchSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
// ```
type GeoMatchSet struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints GeoMatchSetGeoMatchConstraintArrayOutput `pulumi:"geoMatchConstraints"`
	// The name or description of the GeoMatchSet.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewGeoMatchSet registers a new resource with the given unique name, arguments, and options.
func NewGeoMatchSet(ctx *pulumi.Context,
	name string, args *GeoMatchSetArgs, opts ...pulumi.ResourceOption) (*GeoMatchSet, error) {
	if args == nil {
		args = &GeoMatchSetArgs{}
	}

	var resource GeoMatchSet
	err := ctx.RegisterResource("aws:waf/geoMatchSet:GeoMatchSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGeoMatchSet gets an existing GeoMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGeoMatchSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GeoMatchSetState, opts ...pulumi.ResourceOption) (*GeoMatchSet, error) {
	var resource GeoMatchSet
	err := ctx.ReadResource("aws:waf/geoMatchSet:GeoMatchSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GeoMatchSet resources.
type geoMatchSetState struct {
	// Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints []GeoMatchSetGeoMatchConstraint `pulumi:"geoMatchConstraints"`
	// The name or description of the GeoMatchSet.
	Name *string `pulumi:"name"`
}

type GeoMatchSetState struct {
	// Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints GeoMatchSetGeoMatchConstraintArrayInput
	// The name or description of the GeoMatchSet.
	Name pulumi.StringPtrInput
}

func (GeoMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*geoMatchSetState)(nil)).Elem()
}

type geoMatchSetArgs struct {
	// The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints []GeoMatchSetGeoMatchConstraint `pulumi:"geoMatchConstraints"`
	// The name or description of the GeoMatchSet.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a GeoMatchSet resource.
type GeoMatchSetArgs struct {
	// The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints GeoMatchSetGeoMatchConstraintArrayInput
	// The name or description of the GeoMatchSet.
	Name pulumi.StringPtrInput
}

func (GeoMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*geoMatchSetArgs)(nil)).Elem()
}

type GeoMatchSetInput interface {
	pulumi.Input

	ToGeoMatchSetOutput() GeoMatchSetOutput
	ToGeoMatchSetOutputWithContext(ctx context.Context) GeoMatchSetOutput
}

func (*GeoMatchSet) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoMatchSet)(nil))
}

func (i *GeoMatchSet) ToGeoMatchSetOutput() GeoMatchSetOutput {
	return i.ToGeoMatchSetOutputWithContext(context.Background())
}

func (i *GeoMatchSet) ToGeoMatchSetOutputWithContext(ctx context.Context) GeoMatchSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoMatchSetOutput)
}

func (i *GeoMatchSet) ToGeoMatchSetPtrOutput() GeoMatchSetPtrOutput {
	return i.ToGeoMatchSetPtrOutputWithContext(context.Background())
}

func (i *GeoMatchSet) ToGeoMatchSetPtrOutputWithContext(ctx context.Context) GeoMatchSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoMatchSetPtrOutput)
}

type GeoMatchSetPtrInput interface {
	pulumi.Input

	ToGeoMatchSetPtrOutput() GeoMatchSetPtrOutput
	ToGeoMatchSetPtrOutputWithContext(ctx context.Context) GeoMatchSetPtrOutput
}

type geoMatchSetPtrType GeoMatchSetArgs

func (*geoMatchSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GeoMatchSet)(nil))
}

func (i *geoMatchSetPtrType) ToGeoMatchSetPtrOutput() GeoMatchSetPtrOutput {
	return i.ToGeoMatchSetPtrOutputWithContext(context.Background())
}

func (i *geoMatchSetPtrType) ToGeoMatchSetPtrOutputWithContext(ctx context.Context) GeoMatchSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoMatchSetPtrOutput)
}

// GeoMatchSetArrayInput is an input type that accepts GeoMatchSetArray and GeoMatchSetArrayOutput values.
// You can construct a concrete instance of `GeoMatchSetArrayInput` via:
//
//          GeoMatchSetArray{ GeoMatchSetArgs{...} }
type GeoMatchSetArrayInput interface {
	pulumi.Input

	ToGeoMatchSetArrayOutput() GeoMatchSetArrayOutput
	ToGeoMatchSetArrayOutputWithContext(context.Context) GeoMatchSetArrayOutput
}

type GeoMatchSetArray []GeoMatchSetInput

func (GeoMatchSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GeoMatchSet)(nil)).Elem()
}

func (i GeoMatchSetArray) ToGeoMatchSetArrayOutput() GeoMatchSetArrayOutput {
	return i.ToGeoMatchSetArrayOutputWithContext(context.Background())
}

func (i GeoMatchSetArray) ToGeoMatchSetArrayOutputWithContext(ctx context.Context) GeoMatchSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoMatchSetArrayOutput)
}

// GeoMatchSetMapInput is an input type that accepts GeoMatchSetMap and GeoMatchSetMapOutput values.
// You can construct a concrete instance of `GeoMatchSetMapInput` via:
//
//          GeoMatchSetMap{ "key": GeoMatchSetArgs{...} }
type GeoMatchSetMapInput interface {
	pulumi.Input

	ToGeoMatchSetMapOutput() GeoMatchSetMapOutput
	ToGeoMatchSetMapOutputWithContext(context.Context) GeoMatchSetMapOutput
}

type GeoMatchSetMap map[string]GeoMatchSetInput

func (GeoMatchSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GeoMatchSet)(nil)).Elem()
}

func (i GeoMatchSetMap) ToGeoMatchSetMapOutput() GeoMatchSetMapOutput {
	return i.ToGeoMatchSetMapOutputWithContext(context.Background())
}

func (i GeoMatchSetMap) ToGeoMatchSetMapOutputWithContext(ctx context.Context) GeoMatchSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoMatchSetMapOutput)
}

type GeoMatchSetOutput struct{ *pulumi.OutputState }

func (GeoMatchSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoMatchSet)(nil))
}

func (o GeoMatchSetOutput) ToGeoMatchSetOutput() GeoMatchSetOutput {
	return o
}

func (o GeoMatchSetOutput) ToGeoMatchSetOutputWithContext(ctx context.Context) GeoMatchSetOutput {
	return o
}

func (o GeoMatchSetOutput) ToGeoMatchSetPtrOutput() GeoMatchSetPtrOutput {
	return o.ToGeoMatchSetPtrOutputWithContext(context.Background())
}

func (o GeoMatchSetOutput) ToGeoMatchSetPtrOutputWithContext(ctx context.Context) GeoMatchSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GeoMatchSet) *GeoMatchSet {
		return &v
	}).(GeoMatchSetPtrOutput)
}

type GeoMatchSetPtrOutput struct{ *pulumi.OutputState }

func (GeoMatchSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GeoMatchSet)(nil))
}

func (o GeoMatchSetPtrOutput) ToGeoMatchSetPtrOutput() GeoMatchSetPtrOutput {
	return o
}

func (o GeoMatchSetPtrOutput) ToGeoMatchSetPtrOutputWithContext(ctx context.Context) GeoMatchSetPtrOutput {
	return o
}

func (o GeoMatchSetPtrOutput) Elem() GeoMatchSetOutput {
	return o.ApplyT(func(v *GeoMatchSet) GeoMatchSet {
		if v != nil {
			return *v
		}
		var ret GeoMatchSet
		return ret
	}).(GeoMatchSetOutput)
}

type GeoMatchSetArrayOutput struct{ *pulumi.OutputState }

func (GeoMatchSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GeoMatchSet)(nil))
}

func (o GeoMatchSetArrayOutput) ToGeoMatchSetArrayOutput() GeoMatchSetArrayOutput {
	return o
}

func (o GeoMatchSetArrayOutput) ToGeoMatchSetArrayOutputWithContext(ctx context.Context) GeoMatchSetArrayOutput {
	return o
}

func (o GeoMatchSetArrayOutput) Index(i pulumi.IntInput) GeoMatchSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GeoMatchSet {
		return vs[0].([]GeoMatchSet)[vs[1].(int)]
	}).(GeoMatchSetOutput)
}

type GeoMatchSetMapOutput struct{ *pulumi.OutputState }

func (GeoMatchSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GeoMatchSet)(nil))
}

func (o GeoMatchSetMapOutput) ToGeoMatchSetMapOutput() GeoMatchSetMapOutput {
	return o
}

func (o GeoMatchSetMapOutput) ToGeoMatchSetMapOutputWithContext(ctx context.Context) GeoMatchSetMapOutput {
	return o
}

func (o GeoMatchSetMapOutput) MapIndex(k pulumi.StringInput) GeoMatchSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GeoMatchSet {
		return vs[0].(map[string]GeoMatchSet)[vs[1].(string)]
	}).(GeoMatchSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GeoMatchSetInput)(nil)).Elem(), &GeoMatchSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*GeoMatchSetPtrInput)(nil)).Elem(), &GeoMatchSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*GeoMatchSetArrayInput)(nil)).Elem(), GeoMatchSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GeoMatchSetMapInput)(nil)).Elem(), GeoMatchSetMap{})
	pulumi.RegisterOutputType(GeoMatchSetOutput{})
	pulumi.RegisterOutputType(GeoMatchSetPtrOutput{})
	pulumi.RegisterOutputType(GeoMatchSetArrayOutput{})
	pulumi.RegisterOutputType(GeoMatchSetMapOutput{})
}
