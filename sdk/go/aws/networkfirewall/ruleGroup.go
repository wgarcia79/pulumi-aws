// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networkfirewall

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AWS Network Firewall Rule Group Resource
//
// ## Example Usage
// ### Stateful Inspection for denying access to a domain
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/networkfirewall"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networkfirewall.NewRuleGroup(ctx, "example", &networkfirewall.RuleGroupArgs{
// 			Capacity: pulumi.Int(100),
// 			RuleGroup: &networkfirewall.RuleGroupRuleGroupArgs{
// 				RulesSource: &networkfirewall.RuleGroupRuleGroupRulesSourceArgs{
// 					RulesSourceList: &networkfirewall.RuleGroupRuleGroupRulesSourceRulesSourceListArgs{
// 						GeneratedRulesType: pulumi.String("DENYLIST"),
// 						TargetTypes: pulumi.StringArray{
// 							pulumi.String("HTTP_HOST"),
// 						},
// 						Targets: pulumi.StringArray{
// 							pulumi.String("test.example.com"),
// 						},
// 					},
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Tag1": pulumi.String("Value1"),
// 				"Tag2": pulumi.String("Value2"),
// 			},
// 			Type: pulumi.String("STATEFUL"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Stateful Inspection for permitting packets from a source IP address
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/networkfirewall"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ips := []string{
// 			"1.1.1.1/32",
// 			"1.0.0.1/32",
// 		}
// 		_, err := networkfirewall.NewRuleGroup(ctx, "example", &networkfirewall.RuleGroupArgs{
// 			Capacity:    pulumi.Int(50),
// 			Description: pulumi.String("Permits http traffic from source"),
// 			Type:        pulumi.String("STATEFUL"),
// 			RuleGroup: &networkfirewall.RuleGroupRuleGroupArgs{
// 				RulesSource: &networkfirewall.RuleGroupRuleGroupRulesSourceArgs{
// 					Dynamic: []map[string]interface{}{
// 						map[string]interface{}{
// 							"forEach": ips,
// 							"content": []map[string]interface{}{
// 								map[string]interface{}{
// 									"action": "PASS",
// 									"header": []map[string]interface{}{
// 										map[string]interface{}{
// 											"destination":     "ANY",
// 											"destinationPort": "ANY",
// 											"protocol":        "HTTP",
// 											"direction":       "ANY",
// 											"sourcePort":      "ANY",
// 											"source":          stateful_rule.Value,
// 										},
// 									},
// 									"ruleOption": []map[string]interface{}{
// 										map[string]interface{}{
// 											"keyword": "sid:1",
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("permit HTTP from source"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Stateful Inspection for blocking packets from going to an intended destination
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/networkfirewall"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networkfirewall.NewRuleGroup(ctx, "example", &networkfirewall.RuleGroupArgs{
// 			Capacity: pulumi.Int(100),
// 			RuleGroup: &networkfirewall.RuleGroupRuleGroupArgs{
// 				RulesSource: &networkfirewall.RuleGroupRuleGroupRulesSourceArgs{
// 					StatefulRule: []map[string]interface{}{
// 						map[string]interface{}{
// 							"action": "DROP",
// 							"header": map[string]interface{}{
// 								"destination":     "124.1.1.24/32",
// 								"destinationPort": 53,
// 								"direction":       "ANY",
// 								"protocol":        "TCP",
// 								"source":          "1.2.3.4/32",
// 								"sourcePort":      53,
// 							},
// 							"ruleOption": []map[string]interface{}{
// 								map[string]interface{}{
// 									"keyword": "sid:1",
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Tag1": pulumi.String("Value1"),
// 				"Tag2": pulumi.String("Value2"),
// 			},
// 			Type: pulumi.String("STATEFUL"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Stateful Inspection from rules specifications defined in Suricata flat format
//
// ```go
// package main
//
// import (
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/networkfirewall"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networkfirewall.NewRuleGroup(ctx, "example", &networkfirewall.RuleGroupArgs{
// 			Capacity: pulumi.Int(100),
// 			Type:     pulumi.String("STATEFUL"),
// 			Rules:    readFileOrPanic("example.rules"),
// 			Tags: pulumi.StringMap{
// 				"Tag1": pulumi.String("Value1"),
// 				"Tag2": pulumi.String("Value2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Stateful Inspection from rule group specifications using rule variables and Suricata format rules
//
// ```go
// package main
//
// import (
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/networkfirewall"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networkfirewall.NewRuleGroup(ctx, "example", &networkfirewall.RuleGroupArgs{
// 			Capacity: pulumi.Int(100),
// 			Type:     pulumi.String("STATEFUL"),
// 			RuleGroup: &networkfirewall.RuleGroupRuleGroupArgs{
// 				RuleVariables: &networkfirewall.RuleGroupRuleGroupRuleVariablesArgs{
// 					IpSets: networkfirewall.RuleGroupRuleGroupRuleVariablesIpSetArray{
// 						&networkfirewall.RuleGroupRuleGroupRuleVariablesIpSetArgs{
// 							Key: pulumi.String("WEBSERVERS_HOSTS"),
// 							IpSet: &networkfirewall.RuleGroupRuleGroupRuleVariablesIpSetIpSetArgs{
// 								Definitions: pulumi.StringArray{
// 									pulumi.String("10.0.0.0/16"),
// 									pulumi.String("10.0.1.0/24"),
// 									pulumi.String("192.168.0.0/16"),
// 								},
// 							},
// 						},
// 						&networkfirewall.RuleGroupRuleGroupRuleVariablesIpSetArgs{
// 							Key: pulumi.String("EXTERNAL_HOST"),
// 							IpSet: &networkfirewall.RuleGroupRuleGroupRuleVariablesIpSetIpSetArgs{
// 								Definitions: pulumi.StringArray{
// 									pulumi.String("1.2.3.4/32"),
// 								},
// 							},
// 						},
// 					},
// 					PortSets: networkfirewall.RuleGroupRuleGroupRuleVariablesPortSetArray{
// 						&networkfirewall.RuleGroupRuleGroupRuleVariablesPortSetArgs{
// 							Key: pulumi.String("HTTP_PORTS"),
// 							PortSet: &networkfirewall.RuleGroupRuleGroupRuleVariablesPortSetPortSetArgs{
// 								Definitions: pulumi.StringArray{
// 									pulumi.String("443"),
// 									pulumi.String("80"),
// 								},
// 							},
// 						},
// 					},
// 				},
// 				RulesSource: &networkfirewall.RuleGroupRuleGroupRulesSourceArgs{
// 					RulesString: readFileOrPanic("suricata_rules_file"),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Tag1": pulumi.String("Value1"),
// 				"Tag2": pulumi.String("Value2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Stateless Inspection with a Custom Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/networkfirewall"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networkfirewall.NewRuleGroup(ctx, "example", &networkfirewall.RuleGroupArgs{
// 			Capacity:    pulumi.Int(100),
// 			Description: pulumi.String("Stateless Rate Limiting Rule"),
// 			RuleGroup: &networkfirewall.RuleGroupRuleGroupArgs{
// 				RulesSource: &networkfirewall.RuleGroupRuleGroupRulesSourceArgs{
// 					StatelessRulesAndCustomActions: &networkfirewall.RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsArgs{
// 						CustomAction: []map[string]interface{}{
// 							map[string]interface{}{
// 								"actionDefinition": map[string]interface{}{
// 									"publishMetricAction": map[string]interface{}{
// 										"dimension": []map[string]interface{}{
// 											map[string]interface{}{
// 												"value": "2",
// 											},
// 										},
// 									},
// 								},
// 								"actionName": "ExampleMetricsAction",
// 							},
// 						},
// 						StatelessRule: []map[string]interface{}{
// 							map[string]interface{}{
// 								"priority": 1,
// 								"ruleDefinition": map[string]interface{}{
// 									"actions": []string{
// 										"aws:pass",
// 										"ExampleMetricsAction",
// 									},
// 									"matchAttributes": map[string]interface{}{
// 										"destination": []map[string]interface{}{
// 											map[string]interface{}{
// 												"addressDefinition": "124.1.1.5/32",
// 											},
// 										},
// 										"destinationPort": []map[string]interface{}{
// 											map[string]interface{}{
// 												"fromPort": 443,
// 												"toPort":   443,
// 											},
// 										},
// 										"protocols": []float64{
// 											6,
// 										},
// 										"source": []map[string]interface{}{
// 											map[string]interface{}{
// 												"addressDefinition": "1.2.3.4/32",
// 											},
// 										},
// 										"sourcePort": []map[string]interface{}{
// 											map[string]interface{}{
// 												"fromPort": 443,
// 												"toPort":   443,
// 											},
// 										},
// 										"tcpFlag": []map[string]interface{}{
// 											map[string]interface{}{
// 												"flags": []string{
// 													"SYN",
// 												},
// 												"masks": []string{
// 													"SYN",
// 													"ACK",
// 												},
// 											},
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Tag1": pulumi.String("Value1"),
// 				"Tag2": pulumi.String("Value2"),
// 			},
// 			Type: pulumi.String("STATELESS"),
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
// Network Firewall Rule Groups can be imported using their `ARN`.
//
// ```sh
//  $ pulumi import aws:networkfirewall/ruleGroup:RuleGroup example arn:aws:network-firewall:us-west-1:123456789012:stateful-rulegroup/example
// ```
type RuleGroup struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) that identifies the rule group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
	Capacity pulumi.IntOutput `pulumi:"capacity"`
	// A friendly description of the rule group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A friendly name of the rule group.
	Name pulumi.StringOutput `pulumi:"name"`
	// A configuration block that defines the rule group rules. Required unless `rules` is specified. See Rule Group below for details.
	RuleGroup RuleGroupRuleGroupOutput `pulumi:"ruleGroup"`
	// The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `ruleGroup` is specified.
	Rules pulumi.StringPtrOutput `pulumi:"rules"`
	// A map of key:value pairs to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.
	Type pulumi.StringOutput `pulumi:"type"`
	// A string token used when updating the rule group.
	UpdateToken pulumi.StringOutput `pulumi:"updateToken"`
}

// NewRuleGroup registers a new resource with the given unique name, arguments, and options.
func NewRuleGroup(ctx *pulumi.Context,
	name string, args *RuleGroupArgs, opts ...pulumi.ResourceOption) (*RuleGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Capacity == nil {
		return nil, errors.New("invalid value for required argument 'Capacity'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource RuleGroup
	err := ctx.RegisterResource("aws:networkfirewall/ruleGroup:RuleGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleGroup gets an existing RuleGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleGroupState, opts ...pulumi.ResourceOption) (*RuleGroup, error) {
	var resource RuleGroup
	err := ctx.ReadResource("aws:networkfirewall/ruleGroup:RuleGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleGroup resources.
type ruleGroupState struct {
	// The Amazon Resource Name (ARN) that identifies the rule group.
	Arn *string `pulumi:"arn"`
	// The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
	Capacity *int `pulumi:"capacity"`
	// A friendly description of the rule group.
	Description *string `pulumi:"description"`
	// A friendly name of the rule group.
	Name *string `pulumi:"name"`
	// A configuration block that defines the rule group rules. Required unless `rules` is specified. See Rule Group below for details.
	RuleGroup *RuleGroupRuleGroup `pulumi:"ruleGroup"`
	// The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `ruleGroup` is specified.
	Rules *string `pulumi:"rules"`
	// A map of key:value pairs to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.
	Type *string `pulumi:"type"`
	// A string token used when updating the rule group.
	UpdateToken *string `pulumi:"updateToken"`
}

type RuleGroupState struct {
	// The Amazon Resource Name (ARN) that identifies the rule group.
	Arn pulumi.StringPtrInput
	// The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
	Capacity pulumi.IntPtrInput
	// A friendly description of the rule group.
	Description pulumi.StringPtrInput
	// A friendly name of the rule group.
	Name pulumi.StringPtrInput
	// A configuration block that defines the rule group rules. Required unless `rules` is specified. See Rule Group below for details.
	RuleGroup RuleGroupRuleGroupPtrInput
	// The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `ruleGroup` is specified.
	Rules pulumi.StringPtrInput
	// A map of key:value pairs to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.
	Type pulumi.StringPtrInput
	// A string token used when updating the rule group.
	UpdateToken pulumi.StringPtrInput
}

func (RuleGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleGroupState)(nil)).Elem()
}

type ruleGroupArgs struct {
	// The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
	Capacity int `pulumi:"capacity"`
	// A friendly description of the rule group.
	Description *string `pulumi:"description"`
	// A friendly name of the rule group.
	Name *string `pulumi:"name"`
	// A configuration block that defines the rule group rules. Required unless `rules` is specified. See Rule Group below for details.
	RuleGroup *RuleGroupRuleGroup `pulumi:"ruleGroup"`
	// The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `ruleGroup` is specified.
	Rules *string `pulumi:"rules"`
	// A map of key:value pairs to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a RuleGroup resource.
type RuleGroupArgs struct {
	// The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
	Capacity pulumi.IntInput
	// A friendly description of the rule group.
	Description pulumi.StringPtrInput
	// A friendly name of the rule group.
	Name pulumi.StringPtrInput
	// A configuration block that defines the rule group rules. Required unless `rules` is specified. See Rule Group below for details.
	RuleGroup RuleGroupRuleGroupPtrInput
	// The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `ruleGroup` is specified.
	Rules pulumi.StringPtrInput
	// A map of key:value pairs to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.
	Type pulumi.StringInput
}

func (RuleGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleGroupArgs)(nil)).Elem()
}

type RuleGroupInput interface {
	pulumi.Input

	ToRuleGroupOutput() RuleGroupOutput
	ToRuleGroupOutputWithContext(ctx context.Context) RuleGroupOutput
}

func (*RuleGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleGroup)(nil))
}

func (i *RuleGroup) ToRuleGroupOutput() RuleGroupOutput {
	return i.ToRuleGroupOutputWithContext(context.Background())
}

func (i *RuleGroup) ToRuleGroupOutputWithContext(ctx context.Context) RuleGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleGroupOutput)
}

func (i *RuleGroup) ToRuleGroupPtrOutput() RuleGroupPtrOutput {
	return i.ToRuleGroupPtrOutputWithContext(context.Background())
}

func (i *RuleGroup) ToRuleGroupPtrOutputWithContext(ctx context.Context) RuleGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleGroupPtrOutput)
}

type RuleGroupPtrInput interface {
	pulumi.Input

	ToRuleGroupPtrOutput() RuleGroupPtrOutput
	ToRuleGroupPtrOutputWithContext(ctx context.Context) RuleGroupPtrOutput
}

type ruleGroupPtrType RuleGroupArgs

func (*ruleGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleGroup)(nil))
}

func (i *ruleGroupPtrType) ToRuleGroupPtrOutput() RuleGroupPtrOutput {
	return i.ToRuleGroupPtrOutputWithContext(context.Background())
}

func (i *ruleGroupPtrType) ToRuleGroupPtrOutputWithContext(ctx context.Context) RuleGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleGroupPtrOutput)
}

// RuleGroupArrayInput is an input type that accepts RuleGroupArray and RuleGroupArrayOutput values.
// You can construct a concrete instance of `RuleGroupArrayInput` via:
//
//          RuleGroupArray{ RuleGroupArgs{...} }
type RuleGroupArrayInput interface {
	pulumi.Input

	ToRuleGroupArrayOutput() RuleGroupArrayOutput
	ToRuleGroupArrayOutputWithContext(context.Context) RuleGroupArrayOutput
}

type RuleGroupArray []RuleGroupInput

func (RuleGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuleGroup)(nil)).Elem()
}

func (i RuleGroupArray) ToRuleGroupArrayOutput() RuleGroupArrayOutput {
	return i.ToRuleGroupArrayOutputWithContext(context.Background())
}

func (i RuleGroupArray) ToRuleGroupArrayOutputWithContext(ctx context.Context) RuleGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleGroupArrayOutput)
}

// RuleGroupMapInput is an input type that accepts RuleGroupMap and RuleGroupMapOutput values.
// You can construct a concrete instance of `RuleGroupMapInput` via:
//
//          RuleGroupMap{ "key": RuleGroupArgs{...} }
type RuleGroupMapInput interface {
	pulumi.Input

	ToRuleGroupMapOutput() RuleGroupMapOutput
	ToRuleGroupMapOutputWithContext(context.Context) RuleGroupMapOutput
}

type RuleGroupMap map[string]RuleGroupInput

func (RuleGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuleGroup)(nil)).Elem()
}

func (i RuleGroupMap) ToRuleGroupMapOutput() RuleGroupMapOutput {
	return i.ToRuleGroupMapOutputWithContext(context.Background())
}

func (i RuleGroupMap) ToRuleGroupMapOutputWithContext(ctx context.Context) RuleGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleGroupMapOutput)
}

type RuleGroupOutput struct{ *pulumi.OutputState }

func (RuleGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleGroup)(nil))
}

func (o RuleGroupOutput) ToRuleGroupOutput() RuleGroupOutput {
	return o
}

func (o RuleGroupOutput) ToRuleGroupOutputWithContext(ctx context.Context) RuleGroupOutput {
	return o
}

func (o RuleGroupOutput) ToRuleGroupPtrOutput() RuleGroupPtrOutput {
	return o.ToRuleGroupPtrOutputWithContext(context.Background())
}

func (o RuleGroupOutput) ToRuleGroupPtrOutputWithContext(ctx context.Context) RuleGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RuleGroup) *RuleGroup {
		return &v
	}).(RuleGroupPtrOutput)
}

type RuleGroupPtrOutput struct{ *pulumi.OutputState }

func (RuleGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleGroup)(nil))
}

func (o RuleGroupPtrOutput) ToRuleGroupPtrOutput() RuleGroupPtrOutput {
	return o
}

func (o RuleGroupPtrOutput) ToRuleGroupPtrOutputWithContext(ctx context.Context) RuleGroupPtrOutput {
	return o
}

func (o RuleGroupPtrOutput) Elem() RuleGroupOutput {
	return o.ApplyT(func(v *RuleGroup) RuleGroup {
		if v != nil {
			return *v
		}
		var ret RuleGroup
		return ret
	}).(RuleGroupOutput)
}

type RuleGroupArrayOutput struct{ *pulumi.OutputState }

func (RuleGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleGroup)(nil))
}

func (o RuleGroupArrayOutput) ToRuleGroupArrayOutput() RuleGroupArrayOutput {
	return o
}

func (o RuleGroupArrayOutput) ToRuleGroupArrayOutputWithContext(ctx context.Context) RuleGroupArrayOutput {
	return o
}

func (o RuleGroupArrayOutput) Index(i pulumi.IntInput) RuleGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuleGroup {
		return vs[0].([]RuleGroup)[vs[1].(int)]
	}).(RuleGroupOutput)
}

type RuleGroupMapOutput struct{ *pulumi.OutputState }

func (RuleGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RuleGroup)(nil))
}

func (o RuleGroupMapOutput) ToRuleGroupMapOutput() RuleGroupMapOutput {
	return o
}

func (o RuleGroupMapOutput) ToRuleGroupMapOutputWithContext(ctx context.Context) RuleGroupMapOutput {
	return o
}

func (o RuleGroupMapOutput) MapIndex(k pulumi.StringInput) RuleGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RuleGroup {
		return vs[0].(map[string]RuleGroup)[vs[1].(string)]
	}).(RuleGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleGroupInput)(nil)).Elem(), &RuleGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleGroupPtrInput)(nil)).Elem(), &RuleGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleGroupArrayInput)(nil)).Elem(), RuleGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleGroupMapInput)(nil)).Elem(), RuleGroupMap{})
	pulumi.RegisterOutputType(RuleGroupOutput{})
	pulumi.RegisterOutputType(RuleGroupPtrOutput{})
	pulumi.RegisterOutputType(RuleGroupArrayOutput{})
	pulumi.RegisterOutputType(RuleGroupMapOutput{})
}
