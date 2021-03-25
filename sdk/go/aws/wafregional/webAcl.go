// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a WAF Regional Web ACL Resource for use with Application Load Balancer.
//
// ## Example Usage
// ### Regular Rule
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/wafregional"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ipset, err := wafregional.NewIpSet(ctx, "ipset", &wafregional.IpSetArgs{
// 			IpSetDescriptors: wafregional.IpSetIpSetDescriptorArray{
// 				&wafregional.IpSetIpSetDescriptorArgs{
// 					Type:  pulumi.String("IPV4"),
// 					Value: pulumi.String("192.0.7.0/24"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		wafrule, err := wafregional.NewRule(ctx, "wafrule", &wafregional.RuleArgs{
// 			MetricName: pulumi.String("tfWAFRule"),
// 			Predicates: wafregional.RulePredicateArray{
// 				&wafregional.RulePredicateArgs{
// 					DataId:  ipset.ID(),
// 					Negated: pulumi.Bool(false),
// 					Type:    pulumi.String("IPMatch"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = wafregional.NewWebAcl(ctx, "wafacl", &wafregional.WebAclArgs{
// 			MetricName: pulumi.String("tfWebACL"),
// 			DefaultAction: &wafregional.WebAclDefaultActionArgs{
// 				Type: pulumi.String("ALLOW"),
// 			},
// 			Rules: wafregional.WebAclRuleArray{
// 				&wafregional.WebAclRuleArgs{
// 					Action: &wafregional.WebAclRuleActionArgs{
// 						Type: pulumi.String("BLOCK"),
// 					},
// 					Priority: pulumi.Int(1),
// 					RuleId:   wafrule.ID(),
// 					Type:     pulumi.String("REGULAR"),
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
// ### Group Rule
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/wafregional"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := wafregional.NewWebAcl(ctx, "example", &wafregional.WebAclArgs{
// 			MetricName: pulumi.String("example"),
// 			DefaultAction: &wafregional.WebAclDefaultActionArgs{
// 				Type: pulumi.String("ALLOW"),
// 			},
// 			Rules: wafregional.WebAclRuleArray{
// 				&wafregional.WebAclRuleArgs{
// 					Priority: pulumi.Int(1),
// 					RuleId:   pulumi.Any(aws_wafregional_rule_group.Example.Id),
// 					Type:     pulumi.String("GROUP"),
// 					OverrideAction: &wafregional.WebAclRuleOverrideActionArgs{
// 						Type: pulumi.String("NONE"),
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
// ### Logging
//
// > *NOTE:* The Kinesis Firehose Delivery Stream name must begin with `aws-waf-logs-`. See the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/logging.html) for more information about enabling WAF logging.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/wafregional"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := wafregional.NewWebAcl(ctx, "example", &wafregional.WebAclArgs{
// 			LoggingConfiguration: &wafregional.WebAclLoggingConfigurationArgs{
// 				LogDestination: pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Arn),
// 				RedactedFields: &wafregional.WebAclLoggingConfigurationRedactedFieldsArgs{
// 					FieldToMatches: wafregional.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArray{
// 						&wafregional.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs{
// 							Type: pulumi.String("URI"),
// 						},
// 						&wafregional.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs{
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
// WAF Regional Web ACL can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import aws:wafregional/webAcl:WebAcl wafacl a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
// ```
type WebAcl struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the WAF Regional WebACL.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction WebAclDefaultActionOutput `pulumi:"defaultAction"`
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration WebAclLoggingConfigurationPtrOutput `pulumi:"loggingConfiguration"`
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// The name or description of the web ACL.
	Name pulumi.StringOutput `pulumi:"name"`
	// Set of configuration blocks containing rules for the web ACL. Detailed below.
	Rules WebAclRuleArrayOutput `pulumi:"rules"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
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
	err := ctx.RegisterResource("aws:wafregional/webAcl:WebAcl", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:wafregional/webAcl:WebAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAcl resources.
type webAclState struct {
	// Amazon Resource Name (ARN) of the WAF Regional WebACL.
	Arn *string `pulumi:"arn"`
	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction *WebAclDefaultAction `pulumi:"defaultAction"`
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration *WebAclLoggingConfiguration `pulumi:"loggingConfiguration"`
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName *string `pulumi:"metricName"`
	// The name or description of the web ACL.
	Name *string `pulumi:"name"`
	// Set of configuration blocks containing rules for the web ACL. Detailed below.
	Rules []WebAclRule `pulumi:"rules"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

type WebAclState struct {
	// Amazon Resource Name (ARN) of the WAF Regional WebACL.
	Arn pulumi.StringPtrInput
	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction WebAclDefaultActionPtrInput
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration WebAclLoggingConfigurationPtrInput
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName pulumi.StringPtrInput
	// The name or description of the web ACL.
	Name pulumi.StringPtrInput
	// Set of configuration blocks containing rules for the web ACL. Detailed below.
	Rules WebAclRuleArrayInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (WebAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclState)(nil)).Elem()
}

type webAclArgs struct {
	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction WebAclDefaultAction `pulumi:"defaultAction"`
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration *WebAclLoggingConfiguration `pulumi:"loggingConfiguration"`
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName string `pulumi:"metricName"`
	// The name or description of the web ACL.
	Name *string `pulumi:"name"`
	// Set of configuration blocks containing rules for the web ACL. Detailed below.
	Rules []WebAclRule `pulumi:"rules"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a WebAcl resource.
type WebAclArgs struct {
	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction WebAclDefaultActionInput
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration WebAclLoggingConfigurationPtrInput
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName pulumi.StringInput
	// The name or description of the web ACL.
	Name pulumi.StringPtrInput
	// Set of configuration blocks containing rules for the web ACL. Detailed below.
	Rules WebAclRuleArrayInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
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
