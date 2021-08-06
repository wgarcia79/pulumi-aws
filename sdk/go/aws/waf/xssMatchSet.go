// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a WAF XSS Match Set Resource
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
// 		_, err := waf.NewXssMatchSet(ctx, "xssMatchSet", &waf.XssMatchSetArgs{
// 			XssMatchTuples: waf.XssMatchSetXssMatchTupleArray{
// 				&waf.XssMatchSetXssMatchTupleArgs{
// 					FieldToMatch: &waf.XssMatchSetXssMatchTupleFieldToMatchArgs{
// 						Type: pulumi.String("URI"),
// 					},
// 					TextTransformation: pulumi.String("NONE"),
// 				},
// 				&waf.XssMatchSetXssMatchTupleArgs{
// 					FieldToMatch: &waf.XssMatchSetXssMatchTupleFieldToMatchArgs{
// 						Type: pulumi.String("QUERY_STRING"),
// 					},
// 					TextTransformation: pulumi.String("NONE"),
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
// WAF XSS Match Set can be imported using their ID, e.g.
//
// ```sh
//  $ pulumi import aws:waf/xssMatchSet:XssMatchSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
// ```
type XssMatchSet struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name or description of the SizeConstraintSet.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples XssMatchSetXssMatchTupleArrayOutput `pulumi:"xssMatchTuples"`
}

// NewXssMatchSet registers a new resource with the given unique name, arguments, and options.
func NewXssMatchSet(ctx *pulumi.Context,
	name string, args *XssMatchSetArgs, opts ...pulumi.ResourceOption) (*XssMatchSet, error) {
	if args == nil {
		args = &XssMatchSetArgs{}
	}

	var resource XssMatchSet
	err := ctx.RegisterResource("aws:waf/xssMatchSet:XssMatchSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetXssMatchSet gets an existing XssMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetXssMatchSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *XssMatchSetState, opts ...pulumi.ResourceOption) (*XssMatchSet, error) {
	var resource XssMatchSet
	err := ctx.ReadResource("aws:waf/xssMatchSet:XssMatchSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering XssMatchSet resources.
type xssMatchSetState struct {
	// Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// The name or description of the SizeConstraintSet.
	Name *string `pulumi:"name"`
	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples []XssMatchSetXssMatchTuple `pulumi:"xssMatchTuples"`
}

type XssMatchSetState struct {
	// Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// The name or description of the SizeConstraintSet.
	Name pulumi.StringPtrInput
	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples XssMatchSetXssMatchTupleArrayInput
}

func (XssMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*xssMatchSetState)(nil)).Elem()
}

type xssMatchSetArgs struct {
	// The name or description of the SizeConstraintSet.
	Name *string `pulumi:"name"`
	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples []XssMatchSetXssMatchTuple `pulumi:"xssMatchTuples"`
}

// The set of arguments for constructing a XssMatchSet resource.
type XssMatchSetArgs struct {
	// The name or description of the SizeConstraintSet.
	Name pulumi.StringPtrInput
	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples XssMatchSetXssMatchTupleArrayInput
}

func (XssMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*xssMatchSetArgs)(nil)).Elem()
}

type XssMatchSetInput interface {
	pulumi.Input

	ToXssMatchSetOutput() XssMatchSetOutput
	ToXssMatchSetOutputWithContext(ctx context.Context) XssMatchSetOutput
}

func (*XssMatchSet) ElementType() reflect.Type {
	return reflect.TypeOf((*XssMatchSet)(nil))
}

func (i *XssMatchSet) ToXssMatchSetOutput() XssMatchSetOutput {
	return i.ToXssMatchSetOutputWithContext(context.Background())
}

func (i *XssMatchSet) ToXssMatchSetOutputWithContext(ctx context.Context) XssMatchSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(XssMatchSetOutput)
}

func (i *XssMatchSet) ToXssMatchSetPtrOutput() XssMatchSetPtrOutput {
	return i.ToXssMatchSetPtrOutputWithContext(context.Background())
}

func (i *XssMatchSet) ToXssMatchSetPtrOutputWithContext(ctx context.Context) XssMatchSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(XssMatchSetPtrOutput)
}

type XssMatchSetPtrInput interface {
	pulumi.Input

	ToXssMatchSetPtrOutput() XssMatchSetPtrOutput
	ToXssMatchSetPtrOutputWithContext(ctx context.Context) XssMatchSetPtrOutput
}

type xssMatchSetPtrType XssMatchSetArgs

func (*xssMatchSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**XssMatchSet)(nil))
}

func (i *xssMatchSetPtrType) ToXssMatchSetPtrOutput() XssMatchSetPtrOutput {
	return i.ToXssMatchSetPtrOutputWithContext(context.Background())
}

func (i *xssMatchSetPtrType) ToXssMatchSetPtrOutputWithContext(ctx context.Context) XssMatchSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(XssMatchSetPtrOutput)
}

// XssMatchSetArrayInput is an input type that accepts XssMatchSetArray and XssMatchSetArrayOutput values.
// You can construct a concrete instance of `XssMatchSetArrayInput` via:
//
//          XssMatchSetArray{ XssMatchSetArgs{...} }
type XssMatchSetArrayInput interface {
	pulumi.Input

	ToXssMatchSetArrayOutput() XssMatchSetArrayOutput
	ToXssMatchSetArrayOutputWithContext(context.Context) XssMatchSetArrayOutput
}

type XssMatchSetArray []XssMatchSetInput

func (XssMatchSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*XssMatchSet)(nil)).Elem()
}

func (i XssMatchSetArray) ToXssMatchSetArrayOutput() XssMatchSetArrayOutput {
	return i.ToXssMatchSetArrayOutputWithContext(context.Background())
}

func (i XssMatchSetArray) ToXssMatchSetArrayOutputWithContext(ctx context.Context) XssMatchSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(XssMatchSetArrayOutput)
}

// XssMatchSetMapInput is an input type that accepts XssMatchSetMap and XssMatchSetMapOutput values.
// You can construct a concrete instance of `XssMatchSetMapInput` via:
//
//          XssMatchSetMap{ "key": XssMatchSetArgs{...} }
type XssMatchSetMapInput interface {
	pulumi.Input

	ToXssMatchSetMapOutput() XssMatchSetMapOutput
	ToXssMatchSetMapOutputWithContext(context.Context) XssMatchSetMapOutput
}

type XssMatchSetMap map[string]XssMatchSetInput

func (XssMatchSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*XssMatchSet)(nil)).Elem()
}

func (i XssMatchSetMap) ToXssMatchSetMapOutput() XssMatchSetMapOutput {
	return i.ToXssMatchSetMapOutputWithContext(context.Background())
}

func (i XssMatchSetMap) ToXssMatchSetMapOutputWithContext(ctx context.Context) XssMatchSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(XssMatchSetMapOutput)
}

type XssMatchSetOutput struct {
	*pulumi.OutputState
}

func (XssMatchSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*XssMatchSet)(nil))
}

func (o XssMatchSetOutput) ToXssMatchSetOutput() XssMatchSetOutput {
	return o
}

func (o XssMatchSetOutput) ToXssMatchSetOutputWithContext(ctx context.Context) XssMatchSetOutput {
	return o
}

func (o XssMatchSetOutput) ToXssMatchSetPtrOutput() XssMatchSetPtrOutput {
	return o.ToXssMatchSetPtrOutputWithContext(context.Background())
}

func (o XssMatchSetOutput) ToXssMatchSetPtrOutputWithContext(ctx context.Context) XssMatchSetPtrOutput {
	return o.ApplyT(func(v XssMatchSet) *XssMatchSet {
		return &v
	}).(XssMatchSetPtrOutput)
}

type XssMatchSetPtrOutput struct {
	*pulumi.OutputState
}

func (XssMatchSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**XssMatchSet)(nil))
}

func (o XssMatchSetPtrOutput) ToXssMatchSetPtrOutput() XssMatchSetPtrOutput {
	return o
}

func (o XssMatchSetPtrOutput) ToXssMatchSetPtrOutputWithContext(ctx context.Context) XssMatchSetPtrOutput {
	return o
}

type XssMatchSetArrayOutput struct{ *pulumi.OutputState }

func (XssMatchSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]XssMatchSet)(nil))
}

func (o XssMatchSetArrayOutput) ToXssMatchSetArrayOutput() XssMatchSetArrayOutput {
	return o
}

func (o XssMatchSetArrayOutput) ToXssMatchSetArrayOutputWithContext(ctx context.Context) XssMatchSetArrayOutput {
	return o
}

func (o XssMatchSetArrayOutput) Index(i pulumi.IntInput) XssMatchSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) XssMatchSet {
		return vs[0].([]XssMatchSet)[vs[1].(int)]
	}).(XssMatchSetOutput)
}

type XssMatchSetMapOutput struct{ *pulumi.OutputState }

func (XssMatchSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]XssMatchSet)(nil))
}

func (o XssMatchSetMapOutput) ToXssMatchSetMapOutput() XssMatchSetMapOutput {
	return o
}

func (o XssMatchSetMapOutput) ToXssMatchSetMapOutputWithContext(ctx context.Context) XssMatchSetMapOutput {
	return o
}

func (o XssMatchSetMapOutput) MapIndex(k pulumi.StringInput) XssMatchSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) XssMatchSet {
		return vs[0].(map[string]XssMatchSet)[vs[1].(string)]
	}).(XssMatchSetOutput)
}

func init() {
	pulumi.RegisterOutputType(XssMatchSetOutput{})
	pulumi.RegisterOutputType(XssMatchSetPtrOutput{})
	pulumi.RegisterOutputType(XssMatchSetArrayOutput{})
	pulumi.RegisterOutputType(XssMatchSetMapOutput{})
}
