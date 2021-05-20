// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networkfirewall

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AWS Network Firewall Firewall Policy Resource
//
// ## Example Usage
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
// 		_, err := networkfirewall.NewFirewallPolicy(ctx, "example", &networkfirewall.FirewallPolicyArgs{
// 			FirewallPolicy: &networkfirewall.FirewallPolicyFirewallPolicyArgs{
// 				StatelessDefaultActions: pulumi.StringArray{
// 					pulumi.String("aws:pass"),
// 				},
// 				StatelessFragmentDefaultActions: pulumi.StringArray{
// 					pulumi.String("aws:drop"),
// 				},
// 				StatelessRuleGroupReferences: networkfirewall.FirewallPolicyFirewallPolicyStatelessRuleGroupReferenceArray{
// 					&networkfirewall.FirewallPolicyFirewallPolicyStatelessRuleGroupReferenceArgs{
// 						Priority:    pulumi.Int(1),
// 						ResourceArn: pulumi.Any(aws_networkfirewall_rule_group.Example.Arn),
// 					},
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
// ## Policy with a Custom Action for Stateless Inspection
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
// 		_, err := networkfirewall.NewFirewallPolicy(ctx, "test", &networkfirewall.FirewallPolicyArgs{
// 			FirewallPolicy: &networkfirewall.FirewallPolicyFirewallPolicyArgs{
// 				StatelessCustomActions: networkfirewall.FirewallPolicyFirewallPolicyStatelessCustomActionArray{
// 					&networkfirewall.FirewallPolicyFirewallPolicyStatelessCustomActionArgs{
// 						ActionDefinition: &networkfirewall.FirewallPolicyFirewallPolicyStatelessCustomActionActionDefinitionArgs{
// 							PublishMetricAction: &networkfirewall.FirewallPolicyFirewallPolicyStatelessCustomActionActionDefinitionPublishMetricActionArgs{
// 								Dimension: pulumi.StringMapArray{
// 									pulumi.StringMap{
// 										"value": pulumi.String("1"),
// 									},
// 								},
// 							},
// 						},
// 						ActionName: pulumi.String("ExampleCustomAction"),
// 					},
// 				},
// 				StatelessDefaultActions: pulumi.StringArray{
// 					pulumi.String("aws:pass"),
// 					pulumi.String("ExampleCustomAction"),
// 				},
// 				StatelessFragmentDefaultActions: pulumi.StringArray{
// 					pulumi.String("aws:drop"),
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
// Network Firewall Policies can be imported using their `ARN`.
//
// ```sh
//  $ pulumi import aws:networkfirewall/firewallPolicy:FirewallPolicy example arn:aws:network-firewall:us-west-1:123456789012:firewall-policy/example
// ```
type FirewallPolicy struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) that identifies the firewall policy.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A friendly description of the firewall policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A configuration block describing the rule groups and policy actions to use in the firewall policy. See Firewall Policy below for details.
	FirewallPolicy FirewallPolicyFirewallPolicyOutput `pulumi:"firewallPolicy"`
	// A friendly name of the firewall policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// Map of resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// A string token used when updating a firewall policy.
	UpdateToken pulumi.StringOutput `pulumi:"updateToken"`
}

// NewFirewallPolicy registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicy(ctx *pulumi.Context,
	name string, args *FirewallPolicyArgs, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FirewallPolicy == nil {
		return nil, errors.New("invalid value for required argument 'FirewallPolicy'")
	}
	var resource FirewallPolicy
	err := ctx.RegisterResource("aws:networkfirewall/firewallPolicy:FirewallPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicy gets an existing FirewallPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyState, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	var resource FirewallPolicy
	err := ctx.ReadResource("aws:networkfirewall/firewallPolicy:FirewallPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicy resources.
type firewallPolicyState struct {
	// The Amazon Resource Name (ARN) that identifies the firewall policy.
	Arn *string `pulumi:"arn"`
	// A friendly description of the firewall policy.
	Description *string `pulumi:"description"`
	// A configuration block describing the rule groups and policy actions to use in the firewall policy. See Firewall Policy below for details.
	FirewallPolicy *FirewallPolicyFirewallPolicy `pulumi:"firewallPolicy"`
	// A friendly name of the firewall policy.
	Name *string `pulumi:"name"`
	// Map of resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// A string token used when updating a firewall policy.
	UpdateToken *string `pulumi:"updateToken"`
}

type FirewallPolicyState struct {
	// The Amazon Resource Name (ARN) that identifies the firewall policy.
	Arn pulumi.StringPtrInput
	// A friendly description of the firewall policy.
	Description pulumi.StringPtrInput
	// A configuration block describing the rule groups and policy actions to use in the firewall policy. See Firewall Policy below for details.
	FirewallPolicy FirewallPolicyFirewallPolicyPtrInput
	// A friendly name of the firewall policy.
	Name pulumi.StringPtrInput
	// Map of resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// A string token used when updating a firewall policy.
	UpdateToken pulumi.StringPtrInput
}

func (FirewallPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyState)(nil)).Elem()
}

type firewallPolicyArgs struct {
	// A friendly description of the firewall policy.
	Description *string `pulumi:"description"`
	// A configuration block describing the rule groups and policy actions to use in the firewall policy. See Firewall Policy below for details.
	FirewallPolicy FirewallPolicyFirewallPolicy `pulumi:"firewallPolicy"`
	// A friendly name of the firewall policy.
	Name *string `pulumi:"name"`
	// Map of resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a FirewallPolicy resource.
type FirewallPolicyArgs struct {
	// A friendly description of the firewall policy.
	Description pulumi.StringPtrInput
	// A configuration block describing the rule groups and policy actions to use in the firewall policy. See Firewall Policy below for details.
	FirewallPolicy FirewallPolicyFirewallPolicyInput
	// A friendly name of the firewall policy.
	Name pulumi.StringPtrInput
	// Map of resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (FirewallPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyArgs)(nil)).Elem()
}

type FirewallPolicyInput interface {
	pulumi.Input

	ToFirewallPolicyOutput() FirewallPolicyOutput
	ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput
}

func (*FirewallPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicy)(nil))
}

func (i *FirewallPolicy) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return i.ToFirewallPolicyOutputWithContext(context.Background())
}

func (i *FirewallPolicy) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyOutput)
}

func (i *FirewallPolicy) ToFirewallPolicyPtrOutput() FirewallPolicyPtrOutput {
	return i.ToFirewallPolicyPtrOutputWithContext(context.Background())
}

func (i *FirewallPolicy) ToFirewallPolicyPtrOutputWithContext(ctx context.Context) FirewallPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyPtrOutput)
}

type FirewallPolicyPtrInput interface {
	pulumi.Input

	ToFirewallPolicyPtrOutput() FirewallPolicyPtrOutput
	ToFirewallPolicyPtrOutputWithContext(ctx context.Context) FirewallPolicyPtrOutput
}

type firewallPolicyPtrType FirewallPolicyArgs

func (*firewallPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy)(nil))
}

func (i *firewallPolicyPtrType) ToFirewallPolicyPtrOutput() FirewallPolicyPtrOutput {
	return i.ToFirewallPolicyPtrOutputWithContext(context.Background())
}

func (i *firewallPolicyPtrType) ToFirewallPolicyPtrOutputWithContext(ctx context.Context) FirewallPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyPtrOutput)
}

// FirewallPolicyArrayInput is an input type that accepts FirewallPolicyArray and FirewallPolicyArrayOutput values.
// You can construct a concrete instance of `FirewallPolicyArrayInput` via:
//
//          FirewallPolicyArray{ FirewallPolicyArgs{...} }
type FirewallPolicyArrayInput interface {
	pulumi.Input

	ToFirewallPolicyArrayOutput() FirewallPolicyArrayOutput
	ToFirewallPolicyArrayOutputWithContext(context.Context) FirewallPolicyArrayOutput
}

type FirewallPolicyArray []FirewallPolicyInput

func (FirewallPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FirewallPolicy)(nil))
}

func (i FirewallPolicyArray) ToFirewallPolicyArrayOutput() FirewallPolicyArrayOutput {
	return i.ToFirewallPolicyArrayOutputWithContext(context.Background())
}

func (i FirewallPolicyArray) ToFirewallPolicyArrayOutputWithContext(ctx context.Context) FirewallPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyArrayOutput)
}

// FirewallPolicyMapInput is an input type that accepts FirewallPolicyMap and FirewallPolicyMapOutput values.
// You can construct a concrete instance of `FirewallPolicyMapInput` via:
//
//          FirewallPolicyMap{ "key": FirewallPolicyArgs{...} }
type FirewallPolicyMapInput interface {
	pulumi.Input

	ToFirewallPolicyMapOutput() FirewallPolicyMapOutput
	ToFirewallPolicyMapOutputWithContext(context.Context) FirewallPolicyMapOutput
}

type FirewallPolicyMap map[string]FirewallPolicyInput

func (FirewallPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FirewallPolicy)(nil))
}

func (i FirewallPolicyMap) ToFirewallPolicyMapOutput() FirewallPolicyMapOutput {
	return i.ToFirewallPolicyMapOutputWithContext(context.Background())
}

func (i FirewallPolicyMap) ToFirewallPolicyMapOutputWithContext(ctx context.Context) FirewallPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyMapOutput)
}

type FirewallPolicyOutput struct {
	*pulumi.OutputState
}

func (FirewallPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicy)(nil))
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return o
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return o
}

func (o FirewallPolicyOutput) ToFirewallPolicyPtrOutput() FirewallPolicyPtrOutput {
	return o.ToFirewallPolicyPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyOutput) ToFirewallPolicyPtrOutputWithContext(ctx context.Context) FirewallPolicyPtrOutput {
	return o.ApplyT(func(v FirewallPolicy) *FirewallPolicy {
		return &v
	}).(FirewallPolicyPtrOutput)
}

type FirewallPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (FirewallPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy)(nil))
}

func (o FirewallPolicyPtrOutput) ToFirewallPolicyPtrOutput() FirewallPolicyPtrOutput {
	return o
}

func (o FirewallPolicyPtrOutput) ToFirewallPolicyPtrOutputWithContext(ctx context.Context) FirewallPolicyPtrOutput {
	return o
}

type FirewallPolicyArrayOutput struct{ *pulumi.OutputState }

func (FirewallPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallPolicy)(nil))
}

func (o FirewallPolicyArrayOutput) ToFirewallPolicyArrayOutput() FirewallPolicyArrayOutput {
	return o
}

func (o FirewallPolicyArrayOutput) ToFirewallPolicyArrayOutputWithContext(ctx context.Context) FirewallPolicyArrayOutput {
	return o
}

func (o FirewallPolicyArrayOutput) Index(i pulumi.IntInput) FirewallPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallPolicy {
		return vs[0].([]FirewallPolicy)[vs[1].(int)]
	}).(FirewallPolicyOutput)
}

type FirewallPolicyMapOutput struct{ *pulumi.OutputState }

func (FirewallPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FirewallPolicy)(nil))
}

func (o FirewallPolicyMapOutput) ToFirewallPolicyMapOutput() FirewallPolicyMapOutput {
	return o
}

func (o FirewallPolicyMapOutput) ToFirewallPolicyMapOutputWithContext(ctx context.Context) FirewallPolicyMapOutput {
	return o
}

func (o FirewallPolicyMapOutput) MapIndex(k pulumi.StringInput) FirewallPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FirewallPolicy {
		return vs[0].(map[string]FirewallPolicy)[vs[1].(string)]
	}).(FirewallPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyOutput{})
	pulumi.RegisterOutputType(FirewallPolicyPtrOutput{})
	pulumi.RegisterOutputType(FirewallPolicyArrayOutput{})
	pulumi.RegisterOutputType(FirewallPolicyMapOutput{})
}
