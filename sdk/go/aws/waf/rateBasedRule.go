// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a WAF Rate Based Rule Resource
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
// 		ipset, err := waf.NewIpSet(ctx, "ipset", &waf.IpSetArgs{
// 			IpSetDescriptors: waf.IpSetIpSetDescriptorArray{
// 				&waf.IpSetIpSetDescriptorArgs{
// 					Type:  pulumi.String("IPV4"),
// 					Value: pulumi.String("192.0.7.0/24"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = waf.NewRateBasedRule(ctx, "wafrule", &waf.RateBasedRuleArgs{
// 			MetricName: pulumi.String("tfWAFRule"),
// 			RateKey:    pulumi.String("IP"),
// 			RateLimit:  pulumi.Int(100),
// 			Predicates: waf.RateBasedRulePredicateArray{
// 				&waf.RateBasedRulePredicateArgs{
// 					DataId:  ipset.ID(),
// 					Negated: pulumi.Bool(false),
// 					Type:    pulumi.String("IPMatch"),
// 				},
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			ipset,
// 		}))
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
// WAF Rated Based Rule can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import aws:waf/rateBasedRule:RateBasedRule wafrule a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
// ```
type RateBasedRule struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// The name or description of the rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The objects to include in a rule (documented below).
	Predicates RateBasedRulePredicateArrayOutput `pulumi:"predicates"`
	// Valid value is IP.
	RateKey pulumi.StringOutput `pulumi:"rateKey"`
	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit pulumi.IntOutput `pulumi:"rateLimit"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewRateBasedRule registers a new resource with the given unique name, arguments, and options.
func NewRateBasedRule(ctx *pulumi.Context,
	name string, args *RateBasedRuleArgs, opts ...pulumi.ResourceOption) (*RateBasedRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MetricName == nil {
		return nil, errors.New("invalid value for required argument 'MetricName'")
	}
	if args.RateKey == nil {
		return nil, errors.New("invalid value for required argument 'RateKey'")
	}
	if args.RateLimit == nil {
		return nil, errors.New("invalid value for required argument 'RateLimit'")
	}
	var resource RateBasedRule
	err := ctx.RegisterResource("aws:waf/rateBasedRule:RateBasedRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRateBasedRule gets an existing RateBasedRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRateBasedRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RateBasedRuleState, opts ...pulumi.ResourceOption) (*RateBasedRule, error) {
	var resource RateBasedRule
	err := ctx.ReadResource("aws:waf/rateBasedRule:RateBasedRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RateBasedRule resources.
type rateBasedRuleState struct {
	// Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName *string `pulumi:"metricName"`
	// The name or description of the rule.
	Name *string `pulumi:"name"`
	// The objects to include in a rule (documented below).
	Predicates []RateBasedRulePredicate `pulumi:"predicates"`
	// Valid value is IP.
	RateKey *string `pulumi:"rateKey"`
	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit *int `pulumi:"rateLimit"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type RateBasedRuleState struct {
	// Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName pulumi.StringPtrInput
	// The name or description of the rule.
	Name pulumi.StringPtrInput
	// The objects to include in a rule (documented below).
	Predicates RateBasedRulePredicateArrayInput
	// Valid value is IP.
	RateKey pulumi.StringPtrInput
	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit pulumi.IntPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (RateBasedRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*rateBasedRuleState)(nil)).Elem()
}

type rateBasedRuleArgs struct {
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName string `pulumi:"metricName"`
	// The name or description of the rule.
	Name *string `pulumi:"name"`
	// The objects to include in a rule (documented below).
	Predicates []RateBasedRulePredicate `pulumi:"predicates"`
	// Valid value is IP.
	RateKey string `pulumi:"rateKey"`
	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit int `pulumi:"rateLimit"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a RateBasedRule resource.
type RateBasedRuleArgs struct {
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName pulumi.StringInput
	// The name or description of the rule.
	Name pulumi.StringPtrInput
	// The objects to include in a rule (documented below).
	Predicates RateBasedRulePredicateArrayInput
	// Valid value is IP.
	RateKey pulumi.StringInput
	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit pulumi.IntInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (RateBasedRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rateBasedRuleArgs)(nil)).Elem()
}

type RateBasedRuleInput interface {
	pulumi.Input

	ToRateBasedRuleOutput() RateBasedRuleOutput
	ToRateBasedRuleOutputWithContext(ctx context.Context) RateBasedRuleOutput
}

func (*RateBasedRule) ElementType() reflect.Type {
	return reflect.TypeOf((*RateBasedRule)(nil))
}

func (i *RateBasedRule) ToRateBasedRuleOutput() RateBasedRuleOutput {
	return i.ToRateBasedRuleOutputWithContext(context.Background())
}

func (i *RateBasedRule) ToRateBasedRuleOutputWithContext(ctx context.Context) RateBasedRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateBasedRuleOutput)
}

func (i *RateBasedRule) ToRateBasedRulePtrOutput() RateBasedRulePtrOutput {
	return i.ToRateBasedRulePtrOutputWithContext(context.Background())
}

func (i *RateBasedRule) ToRateBasedRulePtrOutputWithContext(ctx context.Context) RateBasedRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateBasedRulePtrOutput)
}

type RateBasedRulePtrInput interface {
	pulumi.Input

	ToRateBasedRulePtrOutput() RateBasedRulePtrOutput
	ToRateBasedRulePtrOutputWithContext(ctx context.Context) RateBasedRulePtrOutput
}

type rateBasedRulePtrType RateBasedRuleArgs

func (*rateBasedRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RateBasedRule)(nil))
}

func (i *rateBasedRulePtrType) ToRateBasedRulePtrOutput() RateBasedRulePtrOutput {
	return i.ToRateBasedRulePtrOutputWithContext(context.Background())
}

func (i *rateBasedRulePtrType) ToRateBasedRulePtrOutputWithContext(ctx context.Context) RateBasedRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateBasedRulePtrOutput)
}

// RateBasedRuleArrayInput is an input type that accepts RateBasedRuleArray and RateBasedRuleArrayOutput values.
// You can construct a concrete instance of `RateBasedRuleArrayInput` via:
//
//          RateBasedRuleArray{ RateBasedRuleArgs{...} }
type RateBasedRuleArrayInput interface {
	pulumi.Input

	ToRateBasedRuleArrayOutput() RateBasedRuleArrayOutput
	ToRateBasedRuleArrayOutputWithContext(context.Context) RateBasedRuleArrayOutput
}

type RateBasedRuleArray []RateBasedRuleInput

func (RateBasedRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RateBasedRule)(nil))
}

func (i RateBasedRuleArray) ToRateBasedRuleArrayOutput() RateBasedRuleArrayOutput {
	return i.ToRateBasedRuleArrayOutputWithContext(context.Background())
}

func (i RateBasedRuleArray) ToRateBasedRuleArrayOutputWithContext(ctx context.Context) RateBasedRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateBasedRuleArrayOutput)
}

// RateBasedRuleMapInput is an input type that accepts RateBasedRuleMap and RateBasedRuleMapOutput values.
// You can construct a concrete instance of `RateBasedRuleMapInput` via:
//
//          RateBasedRuleMap{ "key": RateBasedRuleArgs{...} }
type RateBasedRuleMapInput interface {
	pulumi.Input

	ToRateBasedRuleMapOutput() RateBasedRuleMapOutput
	ToRateBasedRuleMapOutputWithContext(context.Context) RateBasedRuleMapOutput
}

type RateBasedRuleMap map[string]RateBasedRuleInput

func (RateBasedRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RateBasedRule)(nil))
}

func (i RateBasedRuleMap) ToRateBasedRuleMapOutput() RateBasedRuleMapOutput {
	return i.ToRateBasedRuleMapOutputWithContext(context.Background())
}

func (i RateBasedRuleMap) ToRateBasedRuleMapOutputWithContext(ctx context.Context) RateBasedRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateBasedRuleMapOutput)
}

type RateBasedRuleOutput struct {
	*pulumi.OutputState
}

func (RateBasedRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RateBasedRule)(nil))
}

func (o RateBasedRuleOutput) ToRateBasedRuleOutput() RateBasedRuleOutput {
	return o
}

func (o RateBasedRuleOutput) ToRateBasedRuleOutputWithContext(ctx context.Context) RateBasedRuleOutput {
	return o
}

func (o RateBasedRuleOutput) ToRateBasedRulePtrOutput() RateBasedRulePtrOutput {
	return o.ToRateBasedRulePtrOutputWithContext(context.Background())
}

func (o RateBasedRuleOutput) ToRateBasedRulePtrOutputWithContext(ctx context.Context) RateBasedRulePtrOutput {
	return o.ApplyT(func(v RateBasedRule) *RateBasedRule {
		return &v
	}).(RateBasedRulePtrOutput)
}

type RateBasedRulePtrOutput struct {
	*pulumi.OutputState
}

func (RateBasedRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RateBasedRule)(nil))
}

func (o RateBasedRulePtrOutput) ToRateBasedRulePtrOutput() RateBasedRulePtrOutput {
	return o
}

func (o RateBasedRulePtrOutput) ToRateBasedRulePtrOutputWithContext(ctx context.Context) RateBasedRulePtrOutput {
	return o
}

type RateBasedRuleArrayOutput struct{ *pulumi.OutputState }

func (RateBasedRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RateBasedRule)(nil))
}

func (o RateBasedRuleArrayOutput) ToRateBasedRuleArrayOutput() RateBasedRuleArrayOutput {
	return o
}

func (o RateBasedRuleArrayOutput) ToRateBasedRuleArrayOutputWithContext(ctx context.Context) RateBasedRuleArrayOutput {
	return o
}

func (o RateBasedRuleArrayOutput) Index(i pulumi.IntInput) RateBasedRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RateBasedRule {
		return vs[0].([]RateBasedRule)[vs[1].(int)]
	}).(RateBasedRuleOutput)
}

type RateBasedRuleMapOutput struct{ *pulumi.OutputState }

func (RateBasedRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RateBasedRule)(nil))
}

func (o RateBasedRuleMapOutput) ToRateBasedRuleMapOutput() RateBasedRuleMapOutput {
	return o
}

func (o RateBasedRuleMapOutput) ToRateBasedRuleMapOutputWithContext(ctx context.Context) RateBasedRuleMapOutput {
	return o
}

func (o RateBasedRuleMapOutput) MapIndex(k pulumi.StringInput) RateBasedRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RateBasedRule {
		return vs[0].(map[string]RateBasedRule)[vs[1].(string)]
	}).(RateBasedRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(RateBasedRuleOutput{})
	pulumi.RegisterOutputType(RateBasedRulePtrOutput{})
	pulumi.RegisterOutputType(RateBasedRuleArrayOutput{})
	pulumi.RegisterOutputType(RateBasedRuleMapOutput{})
}
