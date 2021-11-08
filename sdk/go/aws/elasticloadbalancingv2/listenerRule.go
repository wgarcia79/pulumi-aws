// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Load Balancer Listener Rule resource.
//
// > **Note:** `alb.ListenerRule` is known as `lb.ListenerRule`. The functionality is identical.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cognito"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
// 		if err != nil {
// 			return err
// 		}
// 		frontEndListener, err := lb.NewListener(ctx, "frontEndListener", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "static", &lb.ListenerRuleArgs{
// 			ListenerArn: frontEndListener.Arn,
// 			Priority:    pulumi.Int(100),
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					Type:           pulumi.String("forward"),
// 					TargetGroupArn: pulumi.Any(aws_lb_target_group.Static.Arn),
// 				},
// 			},
// 			Conditions: lb.ListenerRuleConditionArray{
// 				&lb.ListenerRuleConditionArgs{
// 					PathPattern: &lb.ListenerRuleConditionPathPatternArgs{
// 						Values: pulumi.StringArray{
// 							pulumi.String("/static/*"),
// 						},
// 					},
// 				},
// 				&lb.ListenerRuleConditionArgs{
// 					HostHeader: &lb.ListenerRuleConditionHostHeaderArgs{
// 						Values: pulumi.StringArray{
// 							pulumi.String("example.com"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "hostBasedWeightedRouting", &lb.ListenerRuleArgs{
// 			ListenerArn: frontEndListener.Arn,
// 			Priority:    pulumi.Int(99),
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					Type:           pulumi.String("forward"),
// 					TargetGroupArn: pulumi.Any(aws_lb_target_group.Static.Arn),
// 				},
// 			},
// 			Conditions: lb.ListenerRuleConditionArray{
// 				&lb.ListenerRuleConditionArgs{
// 					HostHeader: &lb.ListenerRuleConditionHostHeaderArgs{
// 						Values: pulumi.StringArray{
// 							pulumi.String("my-service.*.mycompany.io"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "hostBasedRouting", &lb.ListenerRuleArgs{
// 			ListenerArn: frontEndListener.Arn,
// 			Priority:    pulumi.Int(99),
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					Type: pulumi.String("forward"),
// 					Forward: &lb.ListenerRuleActionForwardArgs{
// 						TargetGroups: lb.ListenerRuleActionForwardTargetGroupArray{
// 							&lb.ListenerRuleActionForwardTargetGroupArgs{
// 								Arn:    pulumi.Any(aws_lb_target_group.Main.Arn),
// 								Weight: pulumi.Int(80),
// 							},
// 							&lb.ListenerRuleActionForwardTargetGroupArgs{
// 								Arn:    pulumi.Any(aws_lb_target_group.Canary.Arn),
// 								Weight: pulumi.Int(20),
// 							},
// 						},
// 						Stickiness: &lb.ListenerRuleActionForwardStickinessArgs{
// 							Enabled:  pulumi.Bool(true),
// 							Duration: pulumi.Int(600),
// 						},
// 					},
// 				},
// 			},
// 			Conditions: lb.ListenerRuleConditionArray{
// 				&lb.ListenerRuleConditionArgs{
// 					HostHeader: &lb.ListenerRuleConditionHostHeaderArgs{
// 						Values: pulumi.StringArray{
// 							pulumi.String("my-service.*.mycompany.io"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "redirectHttpToHttps", &lb.ListenerRuleArgs{
// 			ListenerArn: frontEndListener.Arn,
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					Type: pulumi.String("redirect"),
// 					Redirect: &lb.ListenerRuleActionRedirectArgs{
// 						Port:       pulumi.String("443"),
// 						Protocol:   pulumi.String("HTTPS"),
// 						StatusCode: pulumi.String("HTTP_301"),
// 					},
// 				},
// 			},
// 			Conditions: lb.ListenerRuleConditionArray{
// 				&lb.ListenerRuleConditionArgs{
// 					HttpHeader: &lb.ListenerRuleConditionHttpHeaderArgs{
// 						HttpHeaderName: pulumi.String("X-Forwarded-For"),
// 						Values: pulumi.StringArray{
// 							pulumi.String("192.168.1.*"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "healthCheck", &lb.ListenerRuleArgs{
// 			ListenerArn: frontEndListener.Arn,
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					Type: pulumi.String("fixed-response"),
// 					FixedResponse: &lb.ListenerRuleActionFixedResponseArgs{
// 						ContentType: pulumi.String("text/plain"),
// 						MessageBody: pulumi.String("HEALTHY"),
// 						StatusCode:  pulumi.String("200"),
// 					},
// 				},
// 			},
// 			Conditions: lb.ListenerRuleConditionArray{
// 				&lb.ListenerRuleConditionArgs{
// 					QueryStrings: lb.ListenerRuleConditionQueryStringArray{
// 						&lb.ListenerRuleConditionQueryStringArgs{
// 							Key:   pulumi.String("health"),
// 							Value: pulumi.String("check"),
// 						},
// 						&lb.ListenerRuleConditionQueryStringArgs{
// 							Value: pulumi.String("bar"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		pool, err := cognito.NewUserPool(ctx, "pool", nil)
// 		if err != nil {
// 			return err
// 		}
// 		client, err := cognito.NewUserPoolClient(ctx, "client", nil)
// 		if err != nil {
// 			return err
// 		}
// 		domain, err := cognito.NewUserPoolDomain(ctx, "domain", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "admin", &lb.ListenerRuleArgs{
// 			ListenerArn: frontEndListener.Arn,
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					Type: pulumi.String("authenticate-cognito"),
// 					AuthenticateCognito: &lb.ListenerRuleActionAuthenticateCognitoArgs{
// 						UserPoolArn:      pool.Arn,
// 						UserPoolClientId: client.ID(),
// 						UserPoolDomain:   domain.Domain,
// 					},
// 				},
// 				&lb.ListenerRuleActionArgs{
// 					Type:           pulumi.String("forward"),
// 					TargetGroupArn: pulumi.Any(aws_lb_target_group.Static.Arn),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewListenerRule(ctx, "oidc", &lb.ListenerRuleArgs{
// 			ListenerArn: frontEndListener.Arn,
// 			Actions: lb.ListenerRuleActionArray{
// 				&lb.ListenerRuleActionArgs{
// 					Type: pulumi.String("authenticate-oidc"),
// 					AuthenticateOidc: &lb.ListenerRuleActionAuthenticateOidcArgs{
// 						AuthorizationEndpoint: pulumi.String("https://example.com/authorization_endpoint"),
// 						ClientId:              pulumi.String("client_id"),
// 						ClientSecret:          pulumi.String("client_secret"),
// 						Issuer:                pulumi.String("https://example.com"),
// 						TokenEndpoint:         pulumi.String("https://example.com/token_endpoint"),
// 						UserInfoEndpoint:      pulumi.String("https://example.com/user_info_endpoint"),
// 					},
// 				},
// 				&lb.ListenerRuleActionArgs{
// 					Type:           pulumi.String("forward"),
// 					TargetGroupArn: pulumi.Any(aws_lb_target_group.Static.Arn),
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
// Rules can be imported using their ARN, e.g.,
//
// ```sh
//  $ pulumi import aws:elasticloadbalancingv2/listenerRule:ListenerRule front_end arn:aws:elasticloadbalancing:us-west-2:187416307283:listener-rule/app/test/8e4497da625e2d8a/9ab28ade35828f96/67b3d2d36dd7c26b
// ```
//
// Deprecated: aws.elasticloadbalancingv2.ListenerRule has been deprecated in favor of aws.lb.ListenerRule
type ListenerRule struct {
	pulumi.CustomResourceState

	// An Action block. Action blocks are documented below.
	Actions ListenerRuleActionArrayOutput `pulumi:"actions"`
	// The Amazon Resource Name (ARN) of the target group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	Conditions ListenerRuleConditionArrayOutput `pulumi:"conditions"`
	// The ARN of the listener to which to attach the rule.
	ListenerArn pulumi.StringOutput `pulumi:"listenerArn"`
	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewListenerRule registers a new resource with the given unique name, arguments, and options.
func NewListenerRule(ctx *pulumi.Context,
	name string, args *ListenerRuleArgs, opts ...pulumi.ResourceOption) (*ListenerRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.Conditions == nil {
		return nil, errors.New("invalid value for required argument 'Conditions'")
	}
	if args.ListenerArn == nil {
		return nil, errors.New("invalid value for required argument 'ListenerArn'")
	}
	var resource ListenerRule
	err := ctx.RegisterResource("aws:elasticloadbalancingv2/listenerRule:ListenerRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListenerRule gets an existing ListenerRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListenerRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListenerRuleState, opts ...pulumi.ResourceOption) (*ListenerRule, error) {
	var resource ListenerRule
	err := ctx.ReadResource("aws:elasticloadbalancingv2/listenerRule:ListenerRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ListenerRule resources.
type listenerRuleState struct {
	// An Action block. Action blocks are documented below.
	Actions []ListenerRuleAction `pulumi:"actions"`
	// The Amazon Resource Name (ARN) of the target group.
	Arn *string `pulumi:"arn"`
	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	Conditions []ListenerRuleCondition `pulumi:"conditions"`
	// The ARN of the listener to which to attach the rule.
	ListenerArn *string `pulumi:"listenerArn"`
	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority *int `pulumi:"priority"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ListenerRuleState struct {
	// An Action block. Action blocks are documented below.
	Actions ListenerRuleActionArrayInput
	// The Amazon Resource Name (ARN) of the target group.
	Arn pulumi.StringPtrInput
	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	Conditions ListenerRuleConditionArrayInput
	// The ARN of the listener to which to attach the rule.
	ListenerArn pulumi.StringPtrInput
	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority pulumi.IntPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (ListenerRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerRuleState)(nil)).Elem()
}

type listenerRuleArgs struct {
	// An Action block. Action blocks are documented below.
	Actions []ListenerRuleAction `pulumi:"actions"`
	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	Conditions []ListenerRuleCondition `pulumi:"conditions"`
	// The ARN of the listener to which to attach the rule.
	ListenerArn string `pulumi:"listenerArn"`
	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority *int `pulumi:"priority"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ListenerRule resource.
type ListenerRuleArgs struct {
	// An Action block. Action blocks are documented below.
	Actions ListenerRuleActionArrayInput
	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	Conditions ListenerRuleConditionArrayInput
	// The ARN of the listener to which to attach the rule.
	ListenerArn pulumi.StringInput
	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority pulumi.IntPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ListenerRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerRuleArgs)(nil)).Elem()
}

type ListenerRuleInput interface {
	pulumi.Input

	ToListenerRuleOutput() ListenerRuleOutput
	ToListenerRuleOutputWithContext(ctx context.Context) ListenerRuleOutput
}

func (*ListenerRule) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenerRule)(nil))
}

func (i *ListenerRule) ToListenerRuleOutput() ListenerRuleOutput {
	return i.ToListenerRuleOutputWithContext(context.Background())
}

func (i *ListenerRule) ToListenerRuleOutputWithContext(ctx context.Context) ListenerRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRuleOutput)
}

func (i *ListenerRule) ToListenerRulePtrOutput() ListenerRulePtrOutput {
	return i.ToListenerRulePtrOutputWithContext(context.Background())
}

func (i *ListenerRule) ToListenerRulePtrOutputWithContext(ctx context.Context) ListenerRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRulePtrOutput)
}

type ListenerRulePtrInput interface {
	pulumi.Input

	ToListenerRulePtrOutput() ListenerRulePtrOutput
	ToListenerRulePtrOutputWithContext(ctx context.Context) ListenerRulePtrOutput
}

type listenerRulePtrType ListenerRuleArgs

func (*listenerRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ListenerRule)(nil))
}

func (i *listenerRulePtrType) ToListenerRulePtrOutput() ListenerRulePtrOutput {
	return i.ToListenerRulePtrOutputWithContext(context.Background())
}

func (i *listenerRulePtrType) ToListenerRulePtrOutputWithContext(ctx context.Context) ListenerRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRulePtrOutput)
}

// ListenerRuleArrayInput is an input type that accepts ListenerRuleArray and ListenerRuleArrayOutput values.
// You can construct a concrete instance of `ListenerRuleArrayInput` via:
//
//          ListenerRuleArray{ ListenerRuleArgs{...} }
type ListenerRuleArrayInput interface {
	pulumi.Input

	ToListenerRuleArrayOutput() ListenerRuleArrayOutput
	ToListenerRuleArrayOutputWithContext(context.Context) ListenerRuleArrayOutput
}

type ListenerRuleArray []ListenerRuleInput

func (ListenerRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ListenerRule)(nil)).Elem()
}

func (i ListenerRuleArray) ToListenerRuleArrayOutput() ListenerRuleArrayOutput {
	return i.ToListenerRuleArrayOutputWithContext(context.Background())
}

func (i ListenerRuleArray) ToListenerRuleArrayOutputWithContext(ctx context.Context) ListenerRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRuleArrayOutput)
}

// ListenerRuleMapInput is an input type that accepts ListenerRuleMap and ListenerRuleMapOutput values.
// You can construct a concrete instance of `ListenerRuleMapInput` via:
//
//          ListenerRuleMap{ "key": ListenerRuleArgs{...} }
type ListenerRuleMapInput interface {
	pulumi.Input

	ToListenerRuleMapOutput() ListenerRuleMapOutput
	ToListenerRuleMapOutputWithContext(context.Context) ListenerRuleMapOutput
}

type ListenerRuleMap map[string]ListenerRuleInput

func (ListenerRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ListenerRule)(nil)).Elem()
}

func (i ListenerRuleMap) ToListenerRuleMapOutput() ListenerRuleMapOutput {
	return i.ToListenerRuleMapOutputWithContext(context.Background())
}

func (i ListenerRuleMap) ToListenerRuleMapOutputWithContext(ctx context.Context) ListenerRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRuleMapOutput)
}

type ListenerRuleOutput struct{ *pulumi.OutputState }

func (ListenerRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenerRule)(nil))
}

func (o ListenerRuleOutput) ToListenerRuleOutput() ListenerRuleOutput {
	return o
}

func (o ListenerRuleOutput) ToListenerRuleOutputWithContext(ctx context.Context) ListenerRuleOutput {
	return o
}

func (o ListenerRuleOutput) ToListenerRulePtrOutput() ListenerRulePtrOutput {
	return o.ToListenerRulePtrOutputWithContext(context.Background())
}

func (o ListenerRuleOutput) ToListenerRulePtrOutputWithContext(ctx context.Context) ListenerRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ListenerRule) *ListenerRule {
		return &v
	}).(ListenerRulePtrOutput)
}

type ListenerRulePtrOutput struct{ *pulumi.OutputState }

func (ListenerRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ListenerRule)(nil))
}

func (o ListenerRulePtrOutput) ToListenerRulePtrOutput() ListenerRulePtrOutput {
	return o
}

func (o ListenerRulePtrOutput) ToListenerRulePtrOutputWithContext(ctx context.Context) ListenerRulePtrOutput {
	return o
}

func (o ListenerRulePtrOutput) Elem() ListenerRuleOutput {
	return o.ApplyT(func(v *ListenerRule) ListenerRule {
		if v != nil {
			return *v
		}
		var ret ListenerRule
		return ret
	}).(ListenerRuleOutput)
}

type ListenerRuleArrayOutput struct{ *pulumi.OutputState }

func (ListenerRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ListenerRule)(nil))
}

func (o ListenerRuleArrayOutput) ToListenerRuleArrayOutput() ListenerRuleArrayOutput {
	return o
}

func (o ListenerRuleArrayOutput) ToListenerRuleArrayOutputWithContext(ctx context.Context) ListenerRuleArrayOutput {
	return o
}

func (o ListenerRuleArrayOutput) Index(i pulumi.IntInput) ListenerRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ListenerRule {
		return vs[0].([]ListenerRule)[vs[1].(int)]
	}).(ListenerRuleOutput)
}

type ListenerRuleMapOutput struct{ *pulumi.OutputState }

func (ListenerRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ListenerRule)(nil))
}

func (o ListenerRuleMapOutput) ToListenerRuleMapOutput() ListenerRuleMapOutput {
	return o
}

func (o ListenerRuleMapOutput) ToListenerRuleMapOutputWithContext(ctx context.Context) ListenerRuleMapOutput {
	return o
}

func (o ListenerRuleMapOutput) MapIndex(k pulumi.StringInput) ListenerRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ListenerRule {
		return vs[0].(map[string]ListenerRule)[vs[1].(string)]
	}).(ListenerRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerRuleInput)(nil)).Elem(), &ListenerRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerRulePtrInput)(nil)).Elem(), &ListenerRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerRuleArrayInput)(nil)).Elem(), ListenerRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerRuleMapInput)(nil)).Elem(), ListenerRuleMap{})
	pulumi.RegisterOutputType(ListenerRuleOutput{})
	pulumi.RegisterOutputType(ListenerRulePtrOutput{})
	pulumi.RegisterOutputType(ListenerRuleArrayOutput{})
	pulumi.RegisterOutputType(ListenerRuleMapOutput{})
}
