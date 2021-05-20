// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a WAF Web ACL Resource
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
// 		wafrule, err := waf.NewRule(ctx, "wafrule", &waf.RuleArgs{
// 			MetricName: pulumi.String("tfWAFRule"),
// 			Predicates: waf.RulePredicateArray{
// 				&waf.RulePredicateArgs{
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
// 		_, err = waf.NewWebAcl(ctx, "wafAcl", &waf.WebAclArgs{
// 			MetricName: pulumi.String("tfWebACL"),
// 			DefaultAction: &waf.WebAclDefaultActionArgs{
// 				Type: pulumi.String("ALLOW"),
// 			},
// 			Rules: waf.WebAclRuleArray{
// 				&waf.WebAclRuleArgs{
// 					Action: &waf.WebAclRuleActionArgs{
// 						Type: pulumi.String("BLOCK"),
// 					},
// 					Priority: pulumi.Int(1),
// 					RuleId:   wafrule.ID(),
// 					Type:     pulumi.String("REGULAR"),
// 				},
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			ipset,
// 			wafrule,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Logging
//
// > *NOTE:* The Kinesis Firehose Delivery Stream name must begin with `aws-waf-logs-` and be located in `us-east-1` region. See the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/logging.html) for more information about enabling WAF logging.
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
// 		_, err := waf.NewWebAcl(ctx, "example", &waf.WebAclArgs{
// 			LoggingConfiguration: &waf.WebAclLoggingConfigurationArgs{
// 				LogDestination: pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Arn),
// 				RedactedFields: &waf.WebAclLoggingConfigurationRedactedFieldsArgs{
// 					FieldToMatches: waf.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArray{
// 						&waf.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs{
// 							Type: pulumi.String("URI"),
// 						},
// 						&waf.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs{
// 							Data: pulumi.String("referer"),
// 							Type: pulumi.String("HEADER"),
// 						},
// 					},
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
// WAF Web ACL can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:waf/webAcl:WebAcl main 0c8e583e-18f3-4c13-9e2a-67c4805d2f94
// ```
type WebAcl struct {
	pulumi.CustomResourceState

	// The ARN of the WAF WebACL.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Configuration block with action that you want AWS WAF to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL. Detailed below.
	DefaultAction WebAclDefaultActionOutput `pulumi:"defaultAction"`
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration WebAclLoggingConfigurationPtrOutput `pulumi:"loggingConfiguration"`
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// The name or description of the web ACL.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration blocks containing rules to associate with the web ACL and the settings for each rule. Detailed below.
	Rules WebAclRuleArrayOutput `pulumi:"rules"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewWebAcl registers a new resource with the given unique name, arguments, and options.
func NewWebAcl(ctx *pulumi.Context,
	name string, args *WebAclArgs, opts ...pulumi.ResourceOption) (*WebAcl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultAction == nil {
		return nil, errors.New("invalid value for required argument 'DefaultAction'")
	}
	if args.MetricName == nil {
		return nil, errors.New("invalid value for required argument 'MetricName'")
	}
	var resource WebAcl
	err := ctx.RegisterResource("aws:waf/webAcl:WebAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAcl gets an existing WebAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAclState, opts ...pulumi.ResourceOption) (*WebAcl, error) {
	var resource WebAcl
	err := ctx.ReadResource("aws:waf/webAcl:WebAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAcl resources.
type webAclState struct {
	// The ARN of the WAF WebACL.
	Arn *string `pulumi:"arn"`
	// Configuration block with action that you want AWS WAF to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL. Detailed below.
	DefaultAction *WebAclDefaultAction `pulumi:"defaultAction"`
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration *WebAclLoggingConfiguration `pulumi:"loggingConfiguration"`
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName *string `pulumi:"metricName"`
	// The name or description of the web ACL.
	Name *string `pulumi:"name"`
	// Configuration blocks containing rules to associate with the web ACL and the settings for each rule. Detailed below.
	Rules []WebAclRule `pulumi:"rules"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type WebAclState struct {
	// The ARN of the WAF WebACL.
	Arn pulumi.StringPtrInput
	// Configuration block with action that you want AWS WAF to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL. Detailed below.
	DefaultAction WebAclDefaultActionPtrInput
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration WebAclLoggingConfigurationPtrInput
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName pulumi.StringPtrInput
	// The name or description of the web ACL.
	Name pulumi.StringPtrInput
	// Configuration blocks containing rules to associate with the web ACL and the settings for each rule. Detailed below.
	Rules WebAclRuleArrayInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (WebAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclState)(nil)).Elem()
}

type webAclArgs struct {
	// Configuration block with action that you want AWS WAF to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL. Detailed below.
	DefaultAction WebAclDefaultAction `pulumi:"defaultAction"`
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration *WebAclLoggingConfiguration `pulumi:"loggingConfiguration"`
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName string `pulumi:"metricName"`
	// The name or description of the web ACL.
	Name *string `pulumi:"name"`
	// Configuration blocks containing rules to associate with the web ACL and the settings for each rule. Detailed below.
	Rules []WebAclRule `pulumi:"rules"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a WebAcl resource.
type WebAclArgs struct {
	// Configuration block with action that you want AWS WAF to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL. Detailed below.
	DefaultAction WebAclDefaultActionInput
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration WebAclLoggingConfigurationPtrInput
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName pulumi.StringInput
	// The name or description of the web ACL.
	Name pulumi.StringPtrInput
	// Configuration blocks containing rules to associate with the web ACL and the settings for each rule. Detailed below.
	Rules WebAclRuleArrayInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (WebAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclArgs)(nil)).Elem()
}

type WebAclInput interface {
	pulumi.Input

	ToWebAclOutput() WebAclOutput
	ToWebAclOutputWithContext(ctx context.Context) WebAclOutput
}

func (*WebAcl) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAcl)(nil))
}

func (i *WebAcl) ToWebAclOutput() WebAclOutput {
	return i.ToWebAclOutputWithContext(context.Background())
}

func (i *WebAcl) ToWebAclOutputWithContext(ctx context.Context) WebAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclOutput)
}

func (i *WebAcl) ToWebAclPtrOutput() WebAclPtrOutput {
	return i.ToWebAclPtrOutputWithContext(context.Background())
}

func (i *WebAcl) ToWebAclPtrOutputWithContext(ctx context.Context) WebAclPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclPtrOutput)
}

type WebAclPtrInput interface {
	pulumi.Input

	ToWebAclPtrOutput() WebAclPtrOutput
	ToWebAclPtrOutputWithContext(ctx context.Context) WebAclPtrOutput
}

type webAclPtrType WebAclArgs

func (*webAclPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAcl)(nil))
}

func (i *webAclPtrType) ToWebAclPtrOutput() WebAclPtrOutput {
	return i.ToWebAclPtrOutputWithContext(context.Background())
}

func (i *webAclPtrType) ToWebAclPtrOutputWithContext(ctx context.Context) WebAclPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclPtrOutput)
}

// WebAclArrayInput is an input type that accepts WebAclArray and WebAclArrayOutput values.
// You can construct a concrete instance of `WebAclArrayInput` via:
//
//          WebAclArray{ WebAclArgs{...} }
type WebAclArrayInput interface {
	pulumi.Input

	ToWebAclArrayOutput() WebAclArrayOutput
	ToWebAclArrayOutputWithContext(context.Context) WebAclArrayOutput
}

type WebAclArray []WebAclInput

func (WebAclArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WebAcl)(nil))
}

func (i WebAclArray) ToWebAclArrayOutput() WebAclArrayOutput {
	return i.ToWebAclArrayOutputWithContext(context.Background())
}

func (i WebAclArray) ToWebAclArrayOutputWithContext(ctx context.Context) WebAclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclArrayOutput)
}

// WebAclMapInput is an input type that accepts WebAclMap and WebAclMapOutput values.
// You can construct a concrete instance of `WebAclMapInput` via:
//
//          WebAclMap{ "key": WebAclArgs{...} }
type WebAclMapInput interface {
	pulumi.Input

	ToWebAclMapOutput() WebAclMapOutput
	ToWebAclMapOutputWithContext(context.Context) WebAclMapOutput
}

type WebAclMap map[string]WebAclInput

func (WebAclMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WebAcl)(nil))
}

func (i WebAclMap) ToWebAclMapOutput() WebAclMapOutput {
	return i.ToWebAclMapOutputWithContext(context.Background())
}

func (i WebAclMap) ToWebAclMapOutputWithContext(ctx context.Context) WebAclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclMapOutput)
}

type WebAclOutput struct {
	*pulumi.OutputState
}

func (WebAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAcl)(nil))
}

func (o WebAclOutput) ToWebAclOutput() WebAclOutput {
	return o
}

func (o WebAclOutput) ToWebAclOutputWithContext(ctx context.Context) WebAclOutput {
	return o
}

func (o WebAclOutput) ToWebAclPtrOutput() WebAclPtrOutput {
	return o.ToWebAclPtrOutputWithContext(context.Background())
}

func (o WebAclOutput) ToWebAclPtrOutputWithContext(ctx context.Context) WebAclPtrOutput {
	return o.ApplyT(func(v WebAcl) *WebAcl {
		return &v
	}).(WebAclPtrOutput)
}

type WebAclPtrOutput struct {
	*pulumi.OutputState
}

func (WebAclPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAcl)(nil))
}

func (o WebAclPtrOutput) ToWebAclPtrOutput() WebAclPtrOutput {
	return o
}

func (o WebAclPtrOutput) ToWebAclPtrOutputWithContext(ctx context.Context) WebAclPtrOutput {
	return o
}

type WebAclArrayOutput struct{ *pulumi.OutputState }

func (WebAclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebAcl)(nil))
}

func (o WebAclArrayOutput) ToWebAclArrayOutput() WebAclArrayOutput {
	return o
}

func (o WebAclArrayOutput) ToWebAclArrayOutputWithContext(ctx context.Context) WebAclArrayOutput {
	return o
}

func (o WebAclArrayOutput) Index(i pulumi.IntInput) WebAclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebAcl {
		return vs[0].([]WebAcl)[vs[1].(int)]
	}).(WebAclOutput)
}

type WebAclMapOutput struct{ *pulumi.OutputState }

func (WebAclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WebAcl)(nil))
}

func (o WebAclMapOutput) ToWebAclMapOutput() WebAclMapOutput {
	return o
}

func (o WebAclMapOutput) ToWebAclMapOutputWithContext(ctx context.Context) WebAclMapOutput {
	return o
}

func (o WebAclMapOutput) MapIndex(k pulumi.StringInput) WebAclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WebAcl {
		return vs[0].(map[string]WebAcl)[vs[1].(string)]
	}).(WebAclOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAclOutput{})
	pulumi.RegisterOutputType(WebAclPtrOutput{})
	pulumi.RegisterOutputType(WebAclArrayOutput{})
	pulumi.RegisterOutputType(WebAclMapOutput{})
}
