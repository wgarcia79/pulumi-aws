// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package securityhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **Note:** AWS accounts can only be associated with a single Security Hub master account. Destroying this resource will disassociate the member account from the master account.
//
// Accepts a Security Hub invitation.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/securityhub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleMember, err := securityhub.NewMember(ctx, "exampleMember", &securityhub.MemberArgs{
// 			AccountId: pulumi.String("123456789012"),
// 			Email:     pulumi.String("example@example.com"),
// 			Invite:    pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = securityhub.NewAccount(ctx, "inviteeAccount", nil, pulumi.Provider("aws.invitee"))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = securityhub.NewInviteAccepter(ctx, "inviteeInviteAccepter", &securityhub.InviteAccepterArgs{
// 			MasterId: exampleMember.MasterId,
// 		}, pulumi.Provider("aws.invitee"), pulumi.DependsOn([]pulumi.Resource{
// 			aws_securityhub_account.Accepter,
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
// Security Hub invite acceptance can be imported using the account ID, e.g.
//
// ```sh
//  $ pulumi import aws:securityhub/inviteAccepter:InviteAccepter example 123456789012
// ```
type InviteAccepter struct {
	pulumi.CustomResourceState

	// The ID of the invitation.
	InvitationId pulumi.StringOutput `pulumi:"invitationId"`
	// The account ID of the master Security Hub account whose invitation you're accepting.
	MasterId pulumi.StringOutput `pulumi:"masterId"`
}

// NewInviteAccepter registers a new resource with the given unique name, arguments, and options.
func NewInviteAccepter(ctx *pulumi.Context,
	name string, args *InviteAccepterArgs, opts ...pulumi.ResourceOption) (*InviteAccepter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MasterId == nil {
		return nil, errors.New("invalid value for required argument 'MasterId'")
	}
	var resource InviteAccepter
	err := ctx.RegisterResource("aws:securityhub/inviteAccepter:InviteAccepter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInviteAccepter gets an existing InviteAccepter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInviteAccepter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InviteAccepterState, opts ...pulumi.ResourceOption) (*InviteAccepter, error) {
	var resource InviteAccepter
	err := ctx.ReadResource("aws:securityhub/inviteAccepter:InviteAccepter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InviteAccepter resources.
type inviteAccepterState struct {
	// The ID of the invitation.
	InvitationId *string `pulumi:"invitationId"`
	// The account ID of the master Security Hub account whose invitation you're accepting.
	MasterId *string `pulumi:"masterId"`
}

type InviteAccepterState struct {
	// The ID of the invitation.
	InvitationId pulumi.StringPtrInput
	// The account ID of the master Security Hub account whose invitation you're accepting.
	MasterId pulumi.StringPtrInput
}

func (InviteAccepterState) ElementType() reflect.Type {
	return reflect.TypeOf((*inviteAccepterState)(nil)).Elem()
}

type inviteAccepterArgs struct {
	// The account ID of the master Security Hub account whose invitation you're accepting.
	MasterId string `pulumi:"masterId"`
}

// The set of arguments for constructing a InviteAccepter resource.
type InviteAccepterArgs struct {
	// The account ID of the master Security Hub account whose invitation you're accepting.
	MasterId pulumi.StringInput
}

func (InviteAccepterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inviteAccepterArgs)(nil)).Elem()
}

type InviteAccepterInput interface {
	pulumi.Input

	ToInviteAccepterOutput() InviteAccepterOutput
	ToInviteAccepterOutputWithContext(ctx context.Context) InviteAccepterOutput
}

func (*InviteAccepter) ElementType() reflect.Type {
	return reflect.TypeOf((*InviteAccepter)(nil))
}

func (i *InviteAccepter) ToInviteAccepterOutput() InviteAccepterOutput {
	return i.ToInviteAccepterOutputWithContext(context.Background())
}

func (i *InviteAccepter) ToInviteAccepterOutputWithContext(ctx context.Context) InviteAccepterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InviteAccepterOutput)
}

func (i *InviteAccepter) ToInviteAccepterPtrOutput() InviteAccepterPtrOutput {
	return i.ToInviteAccepterPtrOutputWithContext(context.Background())
}

func (i *InviteAccepter) ToInviteAccepterPtrOutputWithContext(ctx context.Context) InviteAccepterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InviteAccepterPtrOutput)
}

type InviteAccepterPtrInput interface {
	pulumi.Input

	ToInviteAccepterPtrOutput() InviteAccepterPtrOutput
	ToInviteAccepterPtrOutputWithContext(ctx context.Context) InviteAccepterPtrOutput
}

type inviteAccepterPtrType InviteAccepterArgs

func (*inviteAccepterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InviteAccepter)(nil))
}

func (i *inviteAccepterPtrType) ToInviteAccepterPtrOutput() InviteAccepterPtrOutput {
	return i.ToInviteAccepterPtrOutputWithContext(context.Background())
}

func (i *inviteAccepterPtrType) ToInviteAccepterPtrOutputWithContext(ctx context.Context) InviteAccepterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InviteAccepterPtrOutput)
}

// InviteAccepterArrayInput is an input type that accepts InviteAccepterArray and InviteAccepterArrayOutput values.
// You can construct a concrete instance of `InviteAccepterArrayInput` via:
//
//          InviteAccepterArray{ InviteAccepterArgs{...} }
type InviteAccepterArrayInput interface {
	pulumi.Input

	ToInviteAccepterArrayOutput() InviteAccepterArrayOutput
	ToInviteAccepterArrayOutputWithContext(context.Context) InviteAccepterArrayOutput
}

type InviteAccepterArray []InviteAccepterInput

func (InviteAccepterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*InviteAccepter)(nil))
}

func (i InviteAccepterArray) ToInviteAccepterArrayOutput() InviteAccepterArrayOutput {
	return i.ToInviteAccepterArrayOutputWithContext(context.Background())
}

func (i InviteAccepterArray) ToInviteAccepterArrayOutputWithContext(ctx context.Context) InviteAccepterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InviteAccepterArrayOutput)
}

// InviteAccepterMapInput is an input type that accepts InviteAccepterMap and InviteAccepterMapOutput values.
// You can construct a concrete instance of `InviteAccepterMapInput` via:
//
//          InviteAccepterMap{ "key": InviteAccepterArgs{...} }
type InviteAccepterMapInput interface {
	pulumi.Input

	ToInviteAccepterMapOutput() InviteAccepterMapOutput
	ToInviteAccepterMapOutputWithContext(context.Context) InviteAccepterMapOutput
}

type InviteAccepterMap map[string]InviteAccepterInput

func (InviteAccepterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*InviteAccepter)(nil))
}

func (i InviteAccepterMap) ToInviteAccepterMapOutput() InviteAccepterMapOutput {
	return i.ToInviteAccepterMapOutputWithContext(context.Background())
}

func (i InviteAccepterMap) ToInviteAccepterMapOutputWithContext(ctx context.Context) InviteAccepterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InviteAccepterMapOutput)
}

type InviteAccepterOutput struct {
	*pulumi.OutputState
}

func (InviteAccepterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InviteAccepter)(nil))
}

func (o InviteAccepterOutput) ToInviteAccepterOutput() InviteAccepterOutput {
	return o
}

func (o InviteAccepterOutput) ToInviteAccepterOutputWithContext(ctx context.Context) InviteAccepterOutput {
	return o
}

func (o InviteAccepterOutput) ToInviteAccepterPtrOutput() InviteAccepterPtrOutput {
	return o.ToInviteAccepterPtrOutputWithContext(context.Background())
}

func (o InviteAccepterOutput) ToInviteAccepterPtrOutputWithContext(ctx context.Context) InviteAccepterPtrOutput {
	return o.ApplyT(func(v InviteAccepter) *InviteAccepter {
		return &v
	}).(InviteAccepterPtrOutput)
}

type InviteAccepterPtrOutput struct {
	*pulumi.OutputState
}

func (InviteAccepterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InviteAccepter)(nil))
}

func (o InviteAccepterPtrOutput) ToInviteAccepterPtrOutput() InviteAccepterPtrOutput {
	return o
}

func (o InviteAccepterPtrOutput) ToInviteAccepterPtrOutputWithContext(ctx context.Context) InviteAccepterPtrOutput {
	return o
}

type InviteAccepterArrayOutput struct{ *pulumi.OutputState }

func (InviteAccepterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InviteAccepter)(nil))
}

func (o InviteAccepterArrayOutput) ToInviteAccepterArrayOutput() InviteAccepterArrayOutput {
	return o
}

func (o InviteAccepterArrayOutput) ToInviteAccepterArrayOutputWithContext(ctx context.Context) InviteAccepterArrayOutput {
	return o
}

func (o InviteAccepterArrayOutput) Index(i pulumi.IntInput) InviteAccepterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InviteAccepter {
		return vs[0].([]InviteAccepter)[vs[1].(int)]
	}).(InviteAccepterOutput)
}

type InviteAccepterMapOutput struct{ *pulumi.OutputState }

func (InviteAccepterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InviteAccepter)(nil))
}

func (o InviteAccepterMapOutput) ToInviteAccepterMapOutput() InviteAccepterMapOutput {
	return o
}

func (o InviteAccepterMapOutput) ToInviteAccepterMapOutputWithContext(ctx context.Context) InviteAccepterMapOutput {
	return o
}

func (o InviteAccepterMapOutput) MapIndex(k pulumi.StringInput) InviteAccepterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InviteAccepter {
		return vs[0].(map[string]InviteAccepter)[vs[1].(string)]
	}).(InviteAccepterOutput)
}

func init() {
	pulumi.RegisterOutputType(InviteAccepterOutput{})
	pulumi.RegisterOutputType(InviteAccepterPtrOutput{})
	pulumi.RegisterOutputType(InviteAccepterArrayOutput{})
	pulumi.RegisterOutputType(InviteAccepterMapOutput{})
}
