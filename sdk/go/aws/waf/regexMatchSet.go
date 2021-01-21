// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a WAF Regex Match Set Resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/waf"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleRegexPatternSet, err := waf.NewRegexPatternSet(ctx, "exampleRegexPatternSet", &waf.RegexPatternSetArgs{
// 			RegexPatternStrings: pulumi.StringArray{
// 				pulumi.String("one"),
// 				pulumi.String("two"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = waf.NewRegexMatchSet(ctx, "exampleRegexMatchSet", &waf.RegexMatchSetArgs{
// 			RegexMatchTuples: waf.RegexMatchSetRegexMatchTupleArray{
// 				&waf.RegexMatchSetRegexMatchTupleArgs{
// 					FieldToMatch: &waf.RegexMatchSetRegexMatchTupleFieldToMatchArgs{
// 						Data: pulumi.String("User-Agent"),
// 						Type: pulumi.String("HEADER"),
// 					},
// 					RegexPatternSetId:  exampleRegexPatternSet.ID(),
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
// WAF Regex Match Set can be imported using their ID, e.g.
//
// ```sh
//  $ pulumi import aws:waf/regexMatchSet:RegexMatchSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
// ```
type RegexMatchSet struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name or description of the Regex Match Set.
	Name pulumi.StringOutput `pulumi:"name"`
	// The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples RegexMatchSetRegexMatchTupleArrayOutput `pulumi:"regexMatchTuples"`
}

// NewRegexMatchSet registers a new resource with the given unique name, arguments, and options.
func NewRegexMatchSet(ctx *pulumi.Context,
	name string, args *RegexMatchSetArgs, opts ...pulumi.ResourceOption) (*RegexMatchSet, error) {
	if args == nil {
		args = &RegexMatchSetArgs{}
	}

	var resource RegexMatchSet
	err := ctx.RegisterResource("aws:waf/regexMatchSet:RegexMatchSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegexMatchSet gets an existing RegexMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegexMatchSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegexMatchSetState, opts ...pulumi.ResourceOption) (*RegexMatchSet, error) {
	var resource RegexMatchSet
	err := ctx.ReadResource("aws:waf/regexMatchSet:RegexMatchSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegexMatchSet resources.
type regexMatchSetState struct {
	// Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// The name or description of the Regex Match Set.
	Name *string `pulumi:"name"`
	// The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples []RegexMatchSetRegexMatchTuple `pulumi:"regexMatchTuples"`
}

type RegexMatchSetState struct {
	// Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// The name or description of the Regex Match Set.
	Name pulumi.StringPtrInput
	// The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples RegexMatchSetRegexMatchTupleArrayInput
}

func (RegexMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*regexMatchSetState)(nil)).Elem()
}

type regexMatchSetArgs struct {
	// The name or description of the Regex Match Set.
	Name *string `pulumi:"name"`
	// The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples []RegexMatchSetRegexMatchTuple `pulumi:"regexMatchTuples"`
}

// The set of arguments for constructing a RegexMatchSet resource.
type RegexMatchSetArgs struct {
	// The name or description of the Regex Match Set.
	Name pulumi.StringPtrInput
	// The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples RegexMatchSetRegexMatchTupleArrayInput
}

func (RegexMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regexMatchSetArgs)(nil)).Elem()
}

type RegexMatchSetInput interface {
	pulumi.Input

	ToRegexMatchSetOutput() RegexMatchSetOutput
	ToRegexMatchSetOutputWithContext(ctx context.Context) RegexMatchSetOutput
}

func (*RegexMatchSet) ElementType() reflect.Type {
	return reflect.TypeOf((*RegexMatchSet)(nil))
}

func (i *RegexMatchSet) ToRegexMatchSetOutput() RegexMatchSetOutput {
	return i.ToRegexMatchSetOutputWithContext(context.Background())
}

func (i *RegexMatchSet) ToRegexMatchSetOutputWithContext(ctx context.Context) RegexMatchSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegexMatchSetOutput)
}

func (i *RegexMatchSet) ToRegexMatchSetPtrOutput() RegexMatchSetPtrOutput {
	return i.ToRegexMatchSetPtrOutputWithContext(context.Background())
}

func (i *RegexMatchSet) ToRegexMatchSetPtrOutputWithContext(ctx context.Context) RegexMatchSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegexMatchSetPtrOutput)
}

type RegexMatchSetPtrInput interface {
	pulumi.Input

	ToRegexMatchSetPtrOutput() RegexMatchSetPtrOutput
	ToRegexMatchSetPtrOutputWithContext(ctx context.Context) RegexMatchSetPtrOutput
}

type regexMatchSetPtrType RegexMatchSetArgs

func (*regexMatchSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegexMatchSet)(nil))
}

func (i *regexMatchSetPtrType) ToRegexMatchSetPtrOutput() RegexMatchSetPtrOutput {
	return i.ToRegexMatchSetPtrOutputWithContext(context.Background())
}

func (i *regexMatchSetPtrType) ToRegexMatchSetPtrOutputWithContext(ctx context.Context) RegexMatchSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegexMatchSetPtrOutput)
}

// RegexMatchSetArrayInput is an input type that accepts RegexMatchSetArray and RegexMatchSetArrayOutput values.
// You can construct a concrete instance of `RegexMatchSetArrayInput` via:
//
//          RegexMatchSetArray{ RegexMatchSetArgs{...} }
type RegexMatchSetArrayInput interface {
	pulumi.Input

	ToRegexMatchSetArrayOutput() RegexMatchSetArrayOutput
	ToRegexMatchSetArrayOutputWithContext(context.Context) RegexMatchSetArrayOutput
}

type RegexMatchSetArray []RegexMatchSetInput

func (RegexMatchSetArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RegexMatchSet)(nil))
}

func (i RegexMatchSetArray) ToRegexMatchSetArrayOutput() RegexMatchSetArrayOutput {
	return i.ToRegexMatchSetArrayOutputWithContext(context.Background())
}

func (i RegexMatchSetArray) ToRegexMatchSetArrayOutputWithContext(ctx context.Context) RegexMatchSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegexMatchSetArrayOutput)
}

// RegexMatchSetMapInput is an input type that accepts RegexMatchSetMap and RegexMatchSetMapOutput values.
// You can construct a concrete instance of `RegexMatchSetMapInput` via:
//
//          RegexMatchSetMap{ "key": RegexMatchSetArgs{...} }
type RegexMatchSetMapInput interface {
	pulumi.Input

	ToRegexMatchSetMapOutput() RegexMatchSetMapOutput
	ToRegexMatchSetMapOutputWithContext(context.Context) RegexMatchSetMapOutput
}

type RegexMatchSetMap map[string]RegexMatchSetInput

func (RegexMatchSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RegexMatchSet)(nil))
}

func (i RegexMatchSetMap) ToRegexMatchSetMapOutput() RegexMatchSetMapOutput {
	return i.ToRegexMatchSetMapOutputWithContext(context.Background())
}

func (i RegexMatchSetMap) ToRegexMatchSetMapOutputWithContext(ctx context.Context) RegexMatchSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegexMatchSetMapOutput)
}

type RegexMatchSetOutput struct {
	*pulumi.OutputState
}

func (RegexMatchSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegexMatchSet)(nil))
}

func (o RegexMatchSetOutput) ToRegexMatchSetOutput() RegexMatchSetOutput {
	return o
}

func (o RegexMatchSetOutput) ToRegexMatchSetOutputWithContext(ctx context.Context) RegexMatchSetOutput {
	return o
}

func (o RegexMatchSetOutput) ToRegexMatchSetPtrOutput() RegexMatchSetPtrOutput {
	return o.ToRegexMatchSetPtrOutputWithContext(context.Background())
}

func (o RegexMatchSetOutput) ToRegexMatchSetPtrOutputWithContext(ctx context.Context) RegexMatchSetPtrOutput {
	return o.ApplyT(func(v RegexMatchSet) *RegexMatchSet {
		return &v
	}).(RegexMatchSetPtrOutput)
}

type RegexMatchSetPtrOutput struct {
	*pulumi.OutputState
}

func (RegexMatchSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegexMatchSet)(nil))
}

func (o RegexMatchSetPtrOutput) ToRegexMatchSetPtrOutput() RegexMatchSetPtrOutput {
	return o
}

func (o RegexMatchSetPtrOutput) ToRegexMatchSetPtrOutputWithContext(ctx context.Context) RegexMatchSetPtrOutput {
	return o
}

type RegexMatchSetArrayOutput struct{ *pulumi.OutputState }

func (RegexMatchSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegexMatchSet)(nil))
}

func (o RegexMatchSetArrayOutput) ToRegexMatchSetArrayOutput() RegexMatchSetArrayOutput {
	return o
}

func (o RegexMatchSetArrayOutput) ToRegexMatchSetArrayOutputWithContext(ctx context.Context) RegexMatchSetArrayOutput {
	return o
}

func (o RegexMatchSetArrayOutput) Index(i pulumi.IntInput) RegexMatchSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RegexMatchSet {
		return vs[0].([]RegexMatchSet)[vs[1].(int)]
	}).(RegexMatchSetOutput)
}

type RegexMatchSetMapOutput struct{ *pulumi.OutputState }

func (RegexMatchSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RegexMatchSet)(nil))
}

func (o RegexMatchSetMapOutput) ToRegexMatchSetMapOutput() RegexMatchSetMapOutput {
	return o
}

func (o RegexMatchSetMapOutput) ToRegexMatchSetMapOutputWithContext(ctx context.Context) RegexMatchSetMapOutput {
	return o
}

func (o RegexMatchSetMapOutput) MapIndex(k pulumi.StringInput) RegexMatchSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RegexMatchSet {
		return vs[0].(map[string]RegexMatchSet)[vs[1].(string)]
	}).(RegexMatchSetOutput)
}

func init() {
	pulumi.RegisterOutputType(RegexMatchSetOutput{})
	pulumi.RegisterOutputType(RegexMatchSetPtrOutput{})
	pulumi.RegisterOutputType(RegexMatchSetArrayOutput{})
	pulumi.RegisterOutputType(RegexMatchSetMapOutput{})
}
