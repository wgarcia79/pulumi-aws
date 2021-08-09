// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an SES receipt rule set resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ses"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ses.NewReceiptRuleSet(ctx, "main", &ses.ReceiptRuleSetArgs{
// 			RuleSetName: pulumi.String("primary-rules"),
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
// SES receipt rule sets can be imported using the rule set name.
//
// ```sh
//  $ pulumi import aws:ses/receiptRuleSet:ReceiptRuleSet my_rule_set my_rule_set_name
// ```
type ReceiptRuleSet struct {
	pulumi.CustomResourceState

	// SES receipt rule set ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name of the rule set.
	RuleSetName pulumi.StringOutput `pulumi:"ruleSetName"`
}

// NewReceiptRuleSet registers a new resource with the given unique name, arguments, and options.
func NewReceiptRuleSet(ctx *pulumi.Context,
	name string, args *ReceiptRuleSetArgs, opts ...pulumi.ResourceOption) (*ReceiptRuleSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RuleSetName == nil {
		return nil, errors.New("invalid value for required argument 'RuleSetName'")
	}
	var resource ReceiptRuleSet
	err := ctx.RegisterResource("aws:ses/receiptRuleSet:ReceiptRuleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReceiptRuleSet gets an existing ReceiptRuleSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReceiptRuleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReceiptRuleSetState, opts ...pulumi.ResourceOption) (*ReceiptRuleSet, error) {
	var resource ReceiptRuleSet
	err := ctx.ReadResource("aws:ses/receiptRuleSet:ReceiptRuleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReceiptRuleSet resources.
type receiptRuleSetState struct {
	// SES receipt rule set ARN.
	Arn *string `pulumi:"arn"`
	// Name of the rule set.
	RuleSetName *string `pulumi:"ruleSetName"`
}

type ReceiptRuleSetState struct {
	// SES receipt rule set ARN.
	Arn pulumi.StringPtrInput
	// Name of the rule set.
	RuleSetName pulumi.StringPtrInput
}

func (ReceiptRuleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*receiptRuleSetState)(nil)).Elem()
}

type receiptRuleSetArgs struct {
	// Name of the rule set.
	RuleSetName string `pulumi:"ruleSetName"`
}

// The set of arguments for constructing a ReceiptRuleSet resource.
type ReceiptRuleSetArgs struct {
	// Name of the rule set.
	RuleSetName pulumi.StringInput
}

func (ReceiptRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*receiptRuleSetArgs)(nil)).Elem()
}

type ReceiptRuleSetInput interface {
	pulumi.Input

	ToReceiptRuleSetOutput() ReceiptRuleSetOutput
	ToReceiptRuleSetOutputWithContext(ctx context.Context) ReceiptRuleSetOutput
}

func (*ReceiptRuleSet) ElementType() reflect.Type {
	return reflect.TypeOf((*ReceiptRuleSet)(nil))
}

func (i *ReceiptRuleSet) ToReceiptRuleSetOutput() ReceiptRuleSetOutput {
	return i.ToReceiptRuleSetOutputWithContext(context.Background())
}

func (i *ReceiptRuleSet) ToReceiptRuleSetOutputWithContext(ctx context.Context) ReceiptRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReceiptRuleSetOutput)
}

func (i *ReceiptRuleSet) ToReceiptRuleSetPtrOutput() ReceiptRuleSetPtrOutput {
	return i.ToReceiptRuleSetPtrOutputWithContext(context.Background())
}

func (i *ReceiptRuleSet) ToReceiptRuleSetPtrOutputWithContext(ctx context.Context) ReceiptRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReceiptRuleSetPtrOutput)
}

type ReceiptRuleSetPtrInput interface {
	pulumi.Input

	ToReceiptRuleSetPtrOutput() ReceiptRuleSetPtrOutput
	ToReceiptRuleSetPtrOutputWithContext(ctx context.Context) ReceiptRuleSetPtrOutput
}

type receiptRuleSetPtrType ReceiptRuleSetArgs

func (*receiptRuleSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReceiptRuleSet)(nil))
}

func (i *receiptRuleSetPtrType) ToReceiptRuleSetPtrOutput() ReceiptRuleSetPtrOutput {
	return i.ToReceiptRuleSetPtrOutputWithContext(context.Background())
}

func (i *receiptRuleSetPtrType) ToReceiptRuleSetPtrOutputWithContext(ctx context.Context) ReceiptRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReceiptRuleSetPtrOutput)
}

// ReceiptRuleSetArrayInput is an input type that accepts ReceiptRuleSetArray and ReceiptRuleSetArrayOutput values.
// You can construct a concrete instance of `ReceiptRuleSetArrayInput` via:
//
//          ReceiptRuleSetArray{ ReceiptRuleSetArgs{...} }
type ReceiptRuleSetArrayInput interface {
	pulumi.Input

	ToReceiptRuleSetArrayOutput() ReceiptRuleSetArrayOutput
	ToReceiptRuleSetArrayOutputWithContext(context.Context) ReceiptRuleSetArrayOutput
}

type ReceiptRuleSetArray []ReceiptRuleSetInput

func (ReceiptRuleSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReceiptRuleSet)(nil)).Elem()
}

func (i ReceiptRuleSetArray) ToReceiptRuleSetArrayOutput() ReceiptRuleSetArrayOutput {
	return i.ToReceiptRuleSetArrayOutputWithContext(context.Background())
}

func (i ReceiptRuleSetArray) ToReceiptRuleSetArrayOutputWithContext(ctx context.Context) ReceiptRuleSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReceiptRuleSetArrayOutput)
}

// ReceiptRuleSetMapInput is an input type that accepts ReceiptRuleSetMap and ReceiptRuleSetMapOutput values.
// You can construct a concrete instance of `ReceiptRuleSetMapInput` via:
//
//          ReceiptRuleSetMap{ "key": ReceiptRuleSetArgs{...} }
type ReceiptRuleSetMapInput interface {
	pulumi.Input

	ToReceiptRuleSetMapOutput() ReceiptRuleSetMapOutput
	ToReceiptRuleSetMapOutputWithContext(context.Context) ReceiptRuleSetMapOutput
}

type ReceiptRuleSetMap map[string]ReceiptRuleSetInput

func (ReceiptRuleSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReceiptRuleSet)(nil)).Elem()
}

func (i ReceiptRuleSetMap) ToReceiptRuleSetMapOutput() ReceiptRuleSetMapOutput {
	return i.ToReceiptRuleSetMapOutputWithContext(context.Background())
}

func (i ReceiptRuleSetMap) ToReceiptRuleSetMapOutputWithContext(ctx context.Context) ReceiptRuleSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReceiptRuleSetMapOutput)
}

type ReceiptRuleSetOutput struct{ *pulumi.OutputState }

func (ReceiptRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReceiptRuleSet)(nil))
}

func (o ReceiptRuleSetOutput) ToReceiptRuleSetOutput() ReceiptRuleSetOutput {
	return o
}

func (o ReceiptRuleSetOutput) ToReceiptRuleSetOutputWithContext(ctx context.Context) ReceiptRuleSetOutput {
	return o
}

func (o ReceiptRuleSetOutput) ToReceiptRuleSetPtrOutput() ReceiptRuleSetPtrOutput {
	return o.ToReceiptRuleSetPtrOutputWithContext(context.Background())
}

func (o ReceiptRuleSetOutput) ToReceiptRuleSetPtrOutputWithContext(ctx context.Context) ReceiptRuleSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReceiptRuleSet) *ReceiptRuleSet {
		return &v
	}).(ReceiptRuleSetPtrOutput)
}

type ReceiptRuleSetPtrOutput struct{ *pulumi.OutputState }

func (ReceiptRuleSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReceiptRuleSet)(nil))
}

func (o ReceiptRuleSetPtrOutput) ToReceiptRuleSetPtrOutput() ReceiptRuleSetPtrOutput {
	return o
}

func (o ReceiptRuleSetPtrOutput) ToReceiptRuleSetPtrOutputWithContext(ctx context.Context) ReceiptRuleSetPtrOutput {
	return o
}

func (o ReceiptRuleSetPtrOutput) Elem() ReceiptRuleSetOutput {
	return o.ApplyT(func(v *ReceiptRuleSet) ReceiptRuleSet {
		if v != nil {
			return *v
		}
		var ret ReceiptRuleSet
		return ret
	}).(ReceiptRuleSetOutput)
}

type ReceiptRuleSetArrayOutput struct{ *pulumi.OutputState }

func (ReceiptRuleSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReceiptRuleSet)(nil))
}

func (o ReceiptRuleSetArrayOutput) ToReceiptRuleSetArrayOutput() ReceiptRuleSetArrayOutput {
	return o
}

func (o ReceiptRuleSetArrayOutput) ToReceiptRuleSetArrayOutputWithContext(ctx context.Context) ReceiptRuleSetArrayOutput {
	return o
}

func (o ReceiptRuleSetArrayOutput) Index(i pulumi.IntInput) ReceiptRuleSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReceiptRuleSet {
		return vs[0].([]ReceiptRuleSet)[vs[1].(int)]
	}).(ReceiptRuleSetOutput)
}

type ReceiptRuleSetMapOutput struct{ *pulumi.OutputState }

func (ReceiptRuleSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReceiptRuleSet)(nil))
}

func (o ReceiptRuleSetMapOutput) ToReceiptRuleSetMapOutput() ReceiptRuleSetMapOutput {
	return o
}

func (o ReceiptRuleSetMapOutput) ToReceiptRuleSetMapOutputWithContext(ctx context.Context) ReceiptRuleSetMapOutput {
	return o
}

func (o ReceiptRuleSetMapOutput) MapIndex(k pulumi.StringInput) ReceiptRuleSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReceiptRuleSet {
		return vs[0].(map[string]ReceiptRuleSet)[vs[1].(string)]
	}).(ReceiptRuleSetOutput)
}

func init() {
	pulumi.RegisterOutputType(ReceiptRuleSetOutput{})
	pulumi.RegisterOutputType(ReceiptRuleSetPtrOutput{})
	pulumi.RegisterOutputType(ReceiptRuleSetArrayOutput{})
	pulumi.RegisterOutputType(ReceiptRuleSetMapOutput{})
}
