// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an IAM policy attached to a group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"encoding/json"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myDevelopers, err := iam.NewGroup(ctx, "myDevelopers", &iam.GroupArgs{
// 			Path: pulumi.String("/users/"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		tmpJSON0, err := json.Marshal(map[string]interface{}{
// 			"Version": "2012-10-17",
// 			"Statement": []map[string]interface{}{
// 				map[string]interface{}{
// 					"Action": []string{
// 						"ec2:Describe*",
// 					},
// 					"Effect":   "Allow",
// 					"Resource": "*",
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json0 := string(tmpJSON0)
// 		_, err = iam.NewGroupPolicy(ctx, "myDeveloperPolicy", &iam.GroupPolicyArgs{
// 			Group:  myDevelopers.Name,
// 			Policy: pulumi.String(json0),
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
// IAM Group Policies can be imported using the `group_name:group_policy_name`, e.g.
//
// ```sh
//  $ pulumi import aws:iam/groupPolicy:GroupPolicy mypolicy group_of_mypolicy_name:mypolicy_name
// ```
type GroupPolicy struct {
	pulumi.CustomResourceState

	// The IAM group to attach to the policy.
	Group pulumi.StringOutput `pulumi:"group"`
	// The name of the policy. If omitted, this provider will
	// assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// The policy document. This is a JSON formatted string.
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewGroupPolicy registers a new resource with the given unique name, arguments, and options.
func NewGroupPolicy(ctx *pulumi.Context,
	name string, args *GroupPolicyArgs, opts ...pulumi.ResourceOption) (*GroupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	var resource GroupPolicy
	err := ctx.RegisterResource("aws:iam/groupPolicy:GroupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupPolicy gets an existing GroupPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupPolicyState, opts ...pulumi.ResourceOption) (*GroupPolicy, error) {
	var resource GroupPolicy
	err := ctx.ReadResource("aws:iam/groupPolicy:GroupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupPolicy resources.
type groupPolicyState struct {
	// The IAM group to attach to the policy.
	Group *string `pulumi:"group"`
	// The name of the policy. If omitted, this provider will
	// assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The policy document. This is a JSON formatted string.
	Policy interface{} `pulumi:"policy"`
}

type GroupPolicyState struct {
	// The IAM group to attach to the policy.
	Group pulumi.StringPtrInput
	// The name of the policy. If omitted, this provider will
	// assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The policy document. This is a JSON formatted string.
	Policy pulumi.Input
}

func (GroupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPolicyState)(nil)).Elem()
}

type groupPolicyArgs struct {
	// The IAM group to attach to the policy.
	Group string `pulumi:"group"`
	// The name of the policy. If omitted, this provider will
	// assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The policy document. This is a JSON formatted string.
	Policy interface{} `pulumi:"policy"`
}

// The set of arguments for constructing a GroupPolicy resource.
type GroupPolicyArgs struct {
	// The IAM group to attach to the policy.
	Group pulumi.StringInput
	// The name of the policy. If omitted, this provider will
	// assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The policy document. This is a JSON formatted string.
	Policy pulumi.Input
}

func (GroupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPolicyArgs)(nil)).Elem()
}

type GroupPolicyInput interface {
	pulumi.Input

	ToGroupPolicyOutput() GroupPolicyOutput
	ToGroupPolicyOutputWithContext(ctx context.Context) GroupPolicyOutput
}

func (*GroupPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupPolicy)(nil))
}

func (i *GroupPolicy) ToGroupPolicyOutput() GroupPolicyOutput {
	return i.ToGroupPolicyOutputWithContext(context.Background())
}

func (i *GroupPolicy) ToGroupPolicyOutputWithContext(ctx context.Context) GroupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPolicyOutput)
}

func (i *GroupPolicy) ToGroupPolicyPtrOutput() GroupPolicyPtrOutput {
	return i.ToGroupPolicyPtrOutputWithContext(context.Background())
}

func (i *GroupPolicy) ToGroupPolicyPtrOutputWithContext(ctx context.Context) GroupPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPolicyPtrOutput)
}

type GroupPolicyPtrInput interface {
	pulumi.Input

	ToGroupPolicyPtrOutput() GroupPolicyPtrOutput
	ToGroupPolicyPtrOutputWithContext(ctx context.Context) GroupPolicyPtrOutput
}

type groupPolicyPtrType GroupPolicyArgs

func (*groupPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupPolicy)(nil))
}

func (i *groupPolicyPtrType) ToGroupPolicyPtrOutput() GroupPolicyPtrOutput {
	return i.ToGroupPolicyPtrOutputWithContext(context.Background())
}

func (i *groupPolicyPtrType) ToGroupPolicyPtrOutputWithContext(ctx context.Context) GroupPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPolicyPtrOutput)
}

// GroupPolicyArrayInput is an input type that accepts GroupPolicyArray and GroupPolicyArrayOutput values.
// You can construct a concrete instance of `GroupPolicyArrayInput` via:
//
//          GroupPolicyArray{ GroupPolicyArgs{...} }
type GroupPolicyArrayInput interface {
	pulumi.Input

	ToGroupPolicyArrayOutput() GroupPolicyArrayOutput
	ToGroupPolicyArrayOutputWithContext(context.Context) GroupPolicyArrayOutput
}

type GroupPolicyArray []GroupPolicyInput

func (GroupPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupPolicy)(nil)).Elem()
}

func (i GroupPolicyArray) ToGroupPolicyArrayOutput() GroupPolicyArrayOutput {
	return i.ToGroupPolicyArrayOutputWithContext(context.Background())
}

func (i GroupPolicyArray) ToGroupPolicyArrayOutputWithContext(ctx context.Context) GroupPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPolicyArrayOutput)
}

// GroupPolicyMapInput is an input type that accepts GroupPolicyMap and GroupPolicyMapOutput values.
// You can construct a concrete instance of `GroupPolicyMapInput` via:
//
//          GroupPolicyMap{ "key": GroupPolicyArgs{...} }
type GroupPolicyMapInput interface {
	pulumi.Input

	ToGroupPolicyMapOutput() GroupPolicyMapOutput
	ToGroupPolicyMapOutputWithContext(context.Context) GroupPolicyMapOutput
}

type GroupPolicyMap map[string]GroupPolicyInput

func (GroupPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupPolicy)(nil)).Elem()
}

func (i GroupPolicyMap) ToGroupPolicyMapOutput() GroupPolicyMapOutput {
	return i.ToGroupPolicyMapOutputWithContext(context.Background())
}

func (i GroupPolicyMap) ToGroupPolicyMapOutputWithContext(ctx context.Context) GroupPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPolicyMapOutput)
}

type GroupPolicyOutput struct{ *pulumi.OutputState }

func (GroupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupPolicy)(nil))
}

func (o GroupPolicyOutput) ToGroupPolicyOutput() GroupPolicyOutput {
	return o
}

func (o GroupPolicyOutput) ToGroupPolicyOutputWithContext(ctx context.Context) GroupPolicyOutput {
	return o
}

func (o GroupPolicyOutput) ToGroupPolicyPtrOutput() GroupPolicyPtrOutput {
	return o.ToGroupPolicyPtrOutputWithContext(context.Background())
}

func (o GroupPolicyOutput) ToGroupPolicyPtrOutputWithContext(ctx context.Context) GroupPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GroupPolicy) *GroupPolicy {
		return &v
	}).(GroupPolicyPtrOutput)
}

type GroupPolicyPtrOutput struct{ *pulumi.OutputState }

func (GroupPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupPolicy)(nil))
}

func (o GroupPolicyPtrOutput) ToGroupPolicyPtrOutput() GroupPolicyPtrOutput {
	return o
}

func (o GroupPolicyPtrOutput) ToGroupPolicyPtrOutputWithContext(ctx context.Context) GroupPolicyPtrOutput {
	return o
}

func (o GroupPolicyPtrOutput) Elem() GroupPolicyOutput {
	return o.ApplyT(func(v *GroupPolicy) GroupPolicy {
		if v != nil {
			return *v
		}
		var ret GroupPolicy
		return ret
	}).(GroupPolicyOutput)
}

type GroupPolicyArrayOutput struct{ *pulumi.OutputState }

func (GroupPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupPolicy)(nil))
}

func (o GroupPolicyArrayOutput) ToGroupPolicyArrayOutput() GroupPolicyArrayOutput {
	return o
}

func (o GroupPolicyArrayOutput) ToGroupPolicyArrayOutputWithContext(ctx context.Context) GroupPolicyArrayOutput {
	return o
}

func (o GroupPolicyArrayOutput) Index(i pulumi.IntInput) GroupPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupPolicy {
		return vs[0].([]GroupPolicy)[vs[1].(int)]
	}).(GroupPolicyOutput)
}

type GroupPolicyMapOutput struct{ *pulumi.OutputState }

func (GroupPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GroupPolicy)(nil))
}

func (o GroupPolicyMapOutput) ToGroupPolicyMapOutput() GroupPolicyMapOutput {
	return o
}

func (o GroupPolicyMapOutput) ToGroupPolicyMapOutputWithContext(ctx context.Context) GroupPolicyMapOutput {
	return o
}

func (o GroupPolicyMapOutput) MapIndex(k pulumi.StringInput) GroupPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GroupPolicy {
		return vs[0].(map[string]GroupPolicy)[vs[1].(string)]
	}).(GroupPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupPolicyOutput{})
	pulumi.RegisterOutputType(GroupPolicyPtrOutput{})
	pulumi.RegisterOutputType(GroupPolicyArrayOutput{})
	pulumi.RegisterOutputType(GroupPolicyMapOutput{})
}
