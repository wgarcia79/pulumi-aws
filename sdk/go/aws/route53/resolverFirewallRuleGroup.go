// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Route 53 Resolver DNS Firewall rule group resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := route53.NewResolverFirewallRuleGroup(ctx, "example", nil)
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
//  Route 53 Resolver DNS Firewall rule groups can be imported using the Route 53 Resolver DNS Firewall rule group ID, e.g.
//
// ```sh
//  $ pulumi import aws:route53/resolverFirewallRuleGroup:ResolverFirewallRuleGroup example rslvr-frg-0123456789abcdef
// ```
type ResolverFirewallRuleGroup struct {
	pulumi.CustomResourceState

	// The ARN (Amazon Resource Name) of the rule group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A name that lets you identify the rule group, to manage and use it.
	Name pulumi.StringOutput `pulumi:"name"`
	// The AWS account ID for the account that created the rule group. When a rule group is shared with your account, this is the account that has shared the rule group with you.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// Whether the rule group is shared with other AWS accounts, or was shared with the current account by another AWS account. Sharing is configured through AWS Resource Access Manager (AWS RAM). Valid values: `NOT_SHARED`, `SHARED_BY_ME`, `SHARED_WITH_ME`
	ShareStatus pulumi.StringOutput `pulumi:"shareStatus"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewResolverFirewallRuleGroup registers a new resource with the given unique name, arguments, and options.
func NewResolverFirewallRuleGroup(ctx *pulumi.Context,
	name string, args *ResolverFirewallRuleGroupArgs, opts ...pulumi.ResourceOption) (*ResolverFirewallRuleGroup, error) {
	if args == nil {
		args = &ResolverFirewallRuleGroupArgs{}
	}

	var resource ResolverFirewallRuleGroup
	err := ctx.RegisterResource("aws:route53/resolverFirewallRuleGroup:ResolverFirewallRuleGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverFirewallRuleGroup gets an existing ResolverFirewallRuleGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverFirewallRuleGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverFirewallRuleGroupState, opts ...pulumi.ResourceOption) (*ResolverFirewallRuleGroup, error) {
	var resource ResolverFirewallRuleGroup
	err := ctx.ReadResource("aws:route53/resolverFirewallRuleGroup:ResolverFirewallRuleGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverFirewallRuleGroup resources.
type resolverFirewallRuleGroupState struct {
	// The ARN (Amazon Resource Name) of the rule group.
	Arn *string `pulumi:"arn"`
	// A name that lets you identify the rule group, to manage and use it.
	Name *string `pulumi:"name"`
	// The AWS account ID for the account that created the rule group. When a rule group is shared with your account, this is the account that has shared the rule group with you.
	OwnerId *string `pulumi:"ownerId"`
	// Whether the rule group is shared with other AWS accounts, or was shared with the current account by another AWS account. Sharing is configured through AWS Resource Access Manager (AWS RAM). Valid values: `NOT_SHARED`, `SHARED_BY_ME`, `SHARED_WITH_ME`
	ShareStatus *string `pulumi:"shareStatus"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ResolverFirewallRuleGroupState struct {
	// The ARN (Amazon Resource Name) of the rule group.
	Arn pulumi.StringPtrInput
	// A name that lets you identify the rule group, to manage and use it.
	Name pulumi.StringPtrInput
	// The AWS account ID for the account that created the rule group. When a rule group is shared with your account, this is the account that has shared the rule group with you.
	OwnerId pulumi.StringPtrInput
	// Whether the rule group is shared with other AWS accounts, or was shared with the current account by another AWS account. Sharing is configured through AWS Resource Access Manager (AWS RAM). Valid values: `NOT_SHARED`, `SHARED_BY_ME`, `SHARED_WITH_ME`
	ShareStatus pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (ResolverFirewallRuleGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverFirewallRuleGroupState)(nil)).Elem()
}

type resolverFirewallRuleGroupArgs struct {
	// A name that lets you identify the rule group, to manage and use it.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a ResolverFirewallRuleGroup resource.
type ResolverFirewallRuleGroupArgs struct {
	// A name that lets you identify the rule group, to manage and use it.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (ResolverFirewallRuleGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverFirewallRuleGroupArgs)(nil)).Elem()
}

type ResolverFirewallRuleGroupInput interface {
	pulumi.Input

	ToResolverFirewallRuleGroupOutput() ResolverFirewallRuleGroupOutput
	ToResolverFirewallRuleGroupOutputWithContext(ctx context.Context) ResolverFirewallRuleGroupOutput
}

func (*ResolverFirewallRuleGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ResolverFirewallRuleGroup)(nil))
}

func (i *ResolverFirewallRuleGroup) ToResolverFirewallRuleGroupOutput() ResolverFirewallRuleGroupOutput {
	return i.ToResolverFirewallRuleGroupOutputWithContext(context.Background())
}

func (i *ResolverFirewallRuleGroup) ToResolverFirewallRuleGroupOutputWithContext(ctx context.Context) ResolverFirewallRuleGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverFirewallRuleGroupOutput)
}

func (i *ResolverFirewallRuleGroup) ToResolverFirewallRuleGroupPtrOutput() ResolverFirewallRuleGroupPtrOutput {
	return i.ToResolverFirewallRuleGroupPtrOutputWithContext(context.Background())
}

func (i *ResolverFirewallRuleGroup) ToResolverFirewallRuleGroupPtrOutputWithContext(ctx context.Context) ResolverFirewallRuleGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverFirewallRuleGroupPtrOutput)
}

type ResolverFirewallRuleGroupPtrInput interface {
	pulumi.Input

	ToResolverFirewallRuleGroupPtrOutput() ResolverFirewallRuleGroupPtrOutput
	ToResolverFirewallRuleGroupPtrOutputWithContext(ctx context.Context) ResolverFirewallRuleGroupPtrOutput
}

type resolverFirewallRuleGroupPtrType ResolverFirewallRuleGroupArgs

func (*resolverFirewallRuleGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverFirewallRuleGroup)(nil))
}

func (i *resolverFirewallRuleGroupPtrType) ToResolverFirewallRuleGroupPtrOutput() ResolverFirewallRuleGroupPtrOutput {
	return i.ToResolverFirewallRuleGroupPtrOutputWithContext(context.Background())
}

func (i *resolverFirewallRuleGroupPtrType) ToResolverFirewallRuleGroupPtrOutputWithContext(ctx context.Context) ResolverFirewallRuleGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverFirewallRuleGroupPtrOutput)
}

// ResolverFirewallRuleGroupArrayInput is an input type that accepts ResolverFirewallRuleGroupArray and ResolverFirewallRuleGroupArrayOutput values.
// You can construct a concrete instance of `ResolverFirewallRuleGroupArrayInput` via:
//
//          ResolverFirewallRuleGroupArray{ ResolverFirewallRuleGroupArgs{...} }
type ResolverFirewallRuleGroupArrayInput interface {
	pulumi.Input

	ToResolverFirewallRuleGroupArrayOutput() ResolverFirewallRuleGroupArrayOutput
	ToResolverFirewallRuleGroupArrayOutputWithContext(context.Context) ResolverFirewallRuleGroupArrayOutput
}

type ResolverFirewallRuleGroupArray []ResolverFirewallRuleGroupInput

func (ResolverFirewallRuleGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ResolverFirewallRuleGroup)(nil))
}

func (i ResolverFirewallRuleGroupArray) ToResolverFirewallRuleGroupArrayOutput() ResolverFirewallRuleGroupArrayOutput {
	return i.ToResolverFirewallRuleGroupArrayOutputWithContext(context.Background())
}

func (i ResolverFirewallRuleGroupArray) ToResolverFirewallRuleGroupArrayOutputWithContext(ctx context.Context) ResolverFirewallRuleGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverFirewallRuleGroupArrayOutput)
}

// ResolverFirewallRuleGroupMapInput is an input type that accepts ResolverFirewallRuleGroupMap and ResolverFirewallRuleGroupMapOutput values.
// You can construct a concrete instance of `ResolverFirewallRuleGroupMapInput` via:
//
//          ResolverFirewallRuleGroupMap{ "key": ResolverFirewallRuleGroupArgs{...} }
type ResolverFirewallRuleGroupMapInput interface {
	pulumi.Input

	ToResolverFirewallRuleGroupMapOutput() ResolverFirewallRuleGroupMapOutput
	ToResolverFirewallRuleGroupMapOutputWithContext(context.Context) ResolverFirewallRuleGroupMapOutput
}

type ResolverFirewallRuleGroupMap map[string]ResolverFirewallRuleGroupInput

func (ResolverFirewallRuleGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ResolverFirewallRuleGroup)(nil))
}

func (i ResolverFirewallRuleGroupMap) ToResolverFirewallRuleGroupMapOutput() ResolverFirewallRuleGroupMapOutput {
	return i.ToResolverFirewallRuleGroupMapOutputWithContext(context.Background())
}

func (i ResolverFirewallRuleGroupMap) ToResolverFirewallRuleGroupMapOutputWithContext(ctx context.Context) ResolverFirewallRuleGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverFirewallRuleGroupMapOutput)
}

type ResolverFirewallRuleGroupOutput struct {
	*pulumi.OutputState
}

func (ResolverFirewallRuleGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResolverFirewallRuleGroup)(nil))
}

func (o ResolverFirewallRuleGroupOutput) ToResolverFirewallRuleGroupOutput() ResolverFirewallRuleGroupOutput {
	return o
}

func (o ResolverFirewallRuleGroupOutput) ToResolverFirewallRuleGroupOutputWithContext(ctx context.Context) ResolverFirewallRuleGroupOutput {
	return o
}

func (o ResolverFirewallRuleGroupOutput) ToResolverFirewallRuleGroupPtrOutput() ResolverFirewallRuleGroupPtrOutput {
	return o.ToResolverFirewallRuleGroupPtrOutputWithContext(context.Background())
}

func (o ResolverFirewallRuleGroupOutput) ToResolverFirewallRuleGroupPtrOutputWithContext(ctx context.Context) ResolverFirewallRuleGroupPtrOutput {
	return o.ApplyT(func(v ResolverFirewallRuleGroup) *ResolverFirewallRuleGroup {
		return &v
	}).(ResolverFirewallRuleGroupPtrOutput)
}

type ResolverFirewallRuleGroupPtrOutput struct {
	*pulumi.OutputState
}

func (ResolverFirewallRuleGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverFirewallRuleGroup)(nil))
}

func (o ResolverFirewallRuleGroupPtrOutput) ToResolverFirewallRuleGroupPtrOutput() ResolverFirewallRuleGroupPtrOutput {
	return o
}

func (o ResolverFirewallRuleGroupPtrOutput) ToResolverFirewallRuleGroupPtrOutputWithContext(ctx context.Context) ResolverFirewallRuleGroupPtrOutput {
	return o
}

type ResolverFirewallRuleGroupArrayOutput struct{ *pulumi.OutputState }

func (ResolverFirewallRuleGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResolverFirewallRuleGroup)(nil))
}

func (o ResolverFirewallRuleGroupArrayOutput) ToResolverFirewallRuleGroupArrayOutput() ResolverFirewallRuleGroupArrayOutput {
	return o
}

func (o ResolverFirewallRuleGroupArrayOutput) ToResolverFirewallRuleGroupArrayOutputWithContext(ctx context.Context) ResolverFirewallRuleGroupArrayOutput {
	return o
}

func (o ResolverFirewallRuleGroupArrayOutput) Index(i pulumi.IntInput) ResolverFirewallRuleGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResolverFirewallRuleGroup {
		return vs[0].([]ResolverFirewallRuleGroup)[vs[1].(int)]
	}).(ResolverFirewallRuleGroupOutput)
}

type ResolverFirewallRuleGroupMapOutput struct{ *pulumi.OutputState }

func (ResolverFirewallRuleGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResolverFirewallRuleGroup)(nil))
}

func (o ResolverFirewallRuleGroupMapOutput) ToResolverFirewallRuleGroupMapOutput() ResolverFirewallRuleGroupMapOutput {
	return o
}

func (o ResolverFirewallRuleGroupMapOutput) ToResolverFirewallRuleGroupMapOutputWithContext(ctx context.Context) ResolverFirewallRuleGroupMapOutput {
	return o
}

func (o ResolverFirewallRuleGroupMapOutput) MapIndex(k pulumi.StringInput) ResolverFirewallRuleGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResolverFirewallRuleGroup {
		return vs[0].(map[string]ResolverFirewallRuleGroup)[vs[1].(string)]
	}).(ResolverFirewallRuleGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(ResolverFirewallRuleGroupOutput{})
	pulumi.RegisterOutputType(ResolverFirewallRuleGroupPtrOutput{})
	pulumi.RegisterOutputType(ResolverFirewallRuleGroupArrayOutput{})
	pulumi.RegisterOutputType(ResolverFirewallRuleGroupMapOutput{})
}
