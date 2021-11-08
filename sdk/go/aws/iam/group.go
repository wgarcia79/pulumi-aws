// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an IAM group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iam.NewGroup(ctx, "developers", &iam.GroupArgs{
// 			Path: pulumi.String("/users/"),
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
// IAM Groups can be imported using the `name`, e.g.,
//
// ```sh
//  $ pulumi import aws:iam/group:Group developers developers
// ```
type Group struct {
	pulumi.CustomResourceState

	// The ARN assigned by AWS for this group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The group's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. Group names are not distinguished by case. For example, you cannot create groups named both "ADMINS" and "admins".
	Name pulumi.StringOutput `pulumi:"name"`
	// Path in which to create the group.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The [unique ID][1] assigned by AWS.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		args = &GroupArgs{}
	}

	var resource Group
	err := ctx.RegisterResource("aws:iam/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("aws:iam/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The ARN assigned by AWS for this group.
	Arn *string `pulumi:"arn"`
	// The group's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. Group names are not distinguished by case. For example, you cannot create groups named both "ADMINS" and "admins".
	Name *string `pulumi:"name"`
	// Path in which to create the group.
	Path *string `pulumi:"path"`
	// The [unique ID][1] assigned by AWS.
	UniqueId *string `pulumi:"uniqueId"`
}

type GroupState struct {
	// The ARN assigned by AWS for this group.
	Arn pulumi.StringPtrInput
	// The group's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. Group names are not distinguished by case. For example, you cannot create groups named both "ADMINS" and "admins".
	Name pulumi.StringPtrInput
	// Path in which to create the group.
	Path pulumi.StringPtrInput
	// The [unique ID][1] assigned by AWS.
	UniqueId pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// The group's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. Group names are not distinguished by case. For example, you cannot create groups named both "ADMINS" and "admins".
	Name *string `pulumi:"name"`
	// Path in which to create the group.
	Path *string `pulumi:"path"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The group's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. Group names are not distinguished by case. For example, you cannot create groups named both "ADMINS" and "admins".
	Name pulumi.StringPtrInput
	// Path in which to create the group.
	Path pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil))
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

func (i *Group) ToGroupPtrOutput() GroupPtrOutput {
	return i.ToGroupPtrOutputWithContext(context.Background())
}

func (i *Group) ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPtrOutput)
}

type GroupPtrInput interface {
	pulumi.Input

	ToGroupPtrOutput() GroupPtrOutput
	ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput
}

type groupPtrType GroupArgs

func (*groupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil))
}

func (i *groupPtrType) ToGroupPtrOutput() GroupPtrOutput {
	return i.ToGroupPtrOutputWithContext(context.Background())
}

func (i *groupPtrType) ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPtrOutput)
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//          GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//          GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil))
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func (o GroupOutput) ToGroupPtrOutput() GroupPtrOutput {
	return o.ToGroupPtrOutputWithContext(context.Background())
}

func (o GroupOutput) ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Group) *Group {
		return &v
	}).(GroupPtrOutput)
}

type GroupPtrOutput struct{ *pulumi.OutputState }

func (GroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil))
}

func (o GroupPtrOutput) ToGroupPtrOutput() GroupPtrOutput {
	return o
}

func (o GroupPtrOutput) ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput {
	return o
}

func (o GroupPtrOutput) Elem() GroupOutput {
	return o.ApplyT(func(v *Group) Group {
		if v != nil {
			return *v
		}
		var ret Group
		return ret
	}).(GroupOutput)
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Group)(nil))
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Group {
		return vs[0].([]Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Group)(nil))
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Group {
		return vs[0].(map[string]Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupPtrInput)(nil)).Elem(), &Group{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupArrayInput)(nil)).Elem(), GroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMapInput)(nil)).Elem(), GroupMap{})
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupPtrOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}
