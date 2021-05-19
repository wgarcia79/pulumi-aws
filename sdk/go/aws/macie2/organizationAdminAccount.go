// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package macie2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage an [Amazon Macie Organization Admin Account](https://docs.aws.amazon.com/macie/latest/APIReference/admin.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/macie2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := macie2.NewAccount(ctx, "example", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = macie2.NewOrganizationAdminAccount(ctx, "test", &macie2.OrganizationAdminAccountArgs{
// 			AdminAccountId: pulumi.String("ID OF THE ADMIN ACCOUNT"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			aws_macie2_account.Test,
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
// `aws_macie2_organization_admin_account` can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import aws:macie2/organizationAdminAccount:OrganizationAdminAccount example abcd1
// ```
type OrganizationAdminAccount struct {
	pulumi.CustomResourceState

	// The AWS account ID for the account to designate as the delegated Amazon Macie administrator account for the organization.
	AdminAccountId pulumi.StringOutput `pulumi:"adminAccountId"`
}

// NewOrganizationAdminAccount registers a new resource with the given unique name, arguments, and options.
func NewOrganizationAdminAccount(ctx *pulumi.Context,
	name string, args *OrganizationAdminAccountArgs, opts ...pulumi.ResourceOption) (*OrganizationAdminAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminAccountId == nil {
		return nil, errors.New("invalid value for required argument 'AdminAccountId'")
	}
	var resource OrganizationAdminAccount
	err := ctx.RegisterResource("aws:macie2/organizationAdminAccount:OrganizationAdminAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationAdminAccount gets an existing OrganizationAdminAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationAdminAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationAdminAccountState, opts ...pulumi.ResourceOption) (*OrganizationAdminAccount, error) {
	var resource OrganizationAdminAccount
	err := ctx.ReadResource("aws:macie2/organizationAdminAccount:OrganizationAdminAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationAdminAccount resources.
type organizationAdminAccountState struct {
	// The AWS account ID for the account to designate as the delegated Amazon Macie administrator account for the organization.
	AdminAccountId *string `pulumi:"adminAccountId"`
}

type OrganizationAdminAccountState struct {
	// The AWS account ID for the account to designate as the delegated Amazon Macie administrator account for the organization.
	AdminAccountId pulumi.StringPtrInput
}

func (OrganizationAdminAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationAdminAccountState)(nil)).Elem()
}

type organizationAdminAccountArgs struct {
	// The AWS account ID for the account to designate as the delegated Amazon Macie administrator account for the organization.
	AdminAccountId string `pulumi:"adminAccountId"`
}

// The set of arguments for constructing a OrganizationAdminAccount resource.
type OrganizationAdminAccountArgs struct {
	// The AWS account ID for the account to designate as the delegated Amazon Macie administrator account for the organization.
	AdminAccountId pulumi.StringInput
}

func (OrganizationAdminAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationAdminAccountArgs)(nil)).Elem()
}

type OrganizationAdminAccountInput interface {
	pulumi.Input

	ToOrganizationAdminAccountOutput() OrganizationAdminAccountOutput
	ToOrganizationAdminAccountOutputWithContext(ctx context.Context) OrganizationAdminAccountOutput
}

func (*OrganizationAdminAccount) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationAdminAccount)(nil))
}

func (i *OrganizationAdminAccount) ToOrganizationAdminAccountOutput() OrganizationAdminAccountOutput {
	return i.ToOrganizationAdminAccountOutputWithContext(context.Background())
}

func (i *OrganizationAdminAccount) ToOrganizationAdminAccountOutputWithContext(ctx context.Context) OrganizationAdminAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationAdminAccountOutput)
}

func (i *OrganizationAdminAccount) ToOrganizationAdminAccountPtrOutput() OrganizationAdminAccountPtrOutput {
	return i.ToOrganizationAdminAccountPtrOutputWithContext(context.Background())
}

func (i *OrganizationAdminAccount) ToOrganizationAdminAccountPtrOutputWithContext(ctx context.Context) OrganizationAdminAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationAdminAccountPtrOutput)
}

type OrganizationAdminAccountPtrInput interface {
	pulumi.Input

	ToOrganizationAdminAccountPtrOutput() OrganizationAdminAccountPtrOutput
	ToOrganizationAdminAccountPtrOutputWithContext(ctx context.Context) OrganizationAdminAccountPtrOutput
}

type organizationAdminAccountPtrType OrganizationAdminAccountArgs

func (*organizationAdminAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationAdminAccount)(nil))
}

func (i *organizationAdminAccountPtrType) ToOrganizationAdminAccountPtrOutput() OrganizationAdminAccountPtrOutput {
	return i.ToOrganizationAdminAccountPtrOutputWithContext(context.Background())
}

func (i *organizationAdminAccountPtrType) ToOrganizationAdminAccountPtrOutputWithContext(ctx context.Context) OrganizationAdminAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationAdminAccountPtrOutput)
}

// OrganizationAdminAccountArrayInput is an input type that accepts OrganizationAdminAccountArray and OrganizationAdminAccountArrayOutput values.
// You can construct a concrete instance of `OrganizationAdminAccountArrayInput` via:
//
//          OrganizationAdminAccountArray{ OrganizationAdminAccountArgs{...} }
type OrganizationAdminAccountArrayInput interface {
	pulumi.Input

	ToOrganizationAdminAccountArrayOutput() OrganizationAdminAccountArrayOutput
	ToOrganizationAdminAccountArrayOutputWithContext(context.Context) OrganizationAdminAccountArrayOutput
}

type OrganizationAdminAccountArray []OrganizationAdminAccountInput

func (OrganizationAdminAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*OrganizationAdminAccount)(nil))
}

func (i OrganizationAdminAccountArray) ToOrganizationAdminAccountArrayOutput() OrganizationAdminAccountArrayOutput {
	return i.ToOrganizationAdminAccountArrayOutputWithContext(context.Background())
}

func (i OrganizationAdminAccountArray) ToOrganizationAdminAccountArrayOutputWithContext(ctx context.Context) OrganizationAdminAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationAdminAccountArrayOutput)
}

// OrganizationAdminAccountMapInput is an input type that accepts OrganizationAdminAccountMap and OrganizationAdminAccountMapOutput values.
// You can construct a concrete instance of `OrganizationAdminAccountMapInput` via:
//
//          OrganizationAdminAccountMap{ "key": OrganizationAdminAccountArgs{...} }
type OrganizationAdminAccountMapInput interface {
	pulumi.Input

	ToOrganizationAdminAccountMapOutput() OrganizationAdminAccountMapOutput
	ToOrganizationAdminAccountMapOutputWithContext(context.Context) OrganizationAdminAccountMapOutput
}

type OrganizationAdminAccountMap map[string]OrganizationAdminAccountInput

func (OrganizationAdminAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*OrganizationAdminAccount)(nil))
}

func (i OrganizationAdminAccountMap) ToOrganizationAdminAccountMapOutput() OrganizationAdminAccountMapOutput {
	return i.ToOrganizationAdminAccountMapOutputWithContext(context.Background())
}

func (i OrganizationAdminAccountMap) ToOrganizationAdminAccountMapOutputWithContext(ctx context.Context) OrganizationAdminAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationAdminAccountMapOutput)
}

type OrganizationAdminAccountOutput struct {
	*pulumi.OutputState
}

func (OrganizationAdminAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationAdminAccount)(nil))
}

func (o OrganizationAdminAccountOutput) ToOrganizationAdminAccountOutput() OrganizationAdminAccountOutput {
	return o
}

func (o OrganizationAdminAccountOutput) ToOrganizationAdminAccountOutputWithContext(ctx context.Context) OrganizationAdminAccountOutput {
	return o
}

func (o OrganizationAdminAccountOutput) ToOrganizationAdminAccountPtrOutput() OrganizationAdminAccountPtrOutput {
	return o.ToOrganizationAdminAccountPtrOutputWithContext(context.Background())
}

func (o OrganizationAdminAccountOutput) ToOrganizationAdminAccountPtrOutputWithContext(ctx context.Context) OrganizationAdminAccountPtrOutput {
	return o.ApplyT(func(v OrganizationAdminAccount) *OrganizationAdminAccount {
		return &v
	}).(OrganizationAdminAccountPtrOutput)
}

type OrganizationAdminAccountPtrOutput struct {
	*pulumi.OutputState
}

func (OrganizationAdminAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationAdminAccount)(nil))
}

func (o OrganizationAdminAccountPtrOutput) ToOrganizationAdminAccountPtrOutput() OrganizationAdminAccountPtrOutput {
	return o
}

func (o OrganizationAdminAccountPtrOutput) ToOrganizationAdminAccountPtrOutputWithContext(ctx context.Context) OrganizationAdminAccountPtrOutput {
	return o
}

type OrganizationAdminAccountArrayOutput struct{ *pulumi.OutputState }

func (OrganizationAdminAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OrganizationAdminAccount)(nil))
}

func (o OrganizationAdminAccountArrayOutput) ToOrganizationAdminAccountArrayOutput() OrganizationAdminAccountArrayOutput {
	return o
}

func (o OrganizationAdminAccountArrayOutput) ToOrganizationAdminAccountArrayOutputWithContext(ctx context.Context) OrganizationAdminAccountArrayOutput {
	return o
}

func (o OrganizationAdminAccountArrayOutput) Index(i pulumi.IntInput) OrganizationAdminAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OrganizationAdminAccount {
		return vs[0].([]OrganizationAdminAccount)[vs[1].(int)]
	}).(OrganizationAdminAccountOutput)
}

type OrganizationAdminAccountMapOutput struct{ *pulumi.OutputState }

func (OrganizationAdminAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OrganizationAdminAccount)(nil))
}

func (o OrganizationAdminAccountMapOutput) ToOrganizationAdminAccountMapOutput() OrganizationAdminAccountMapOutput {
	return o
}

func (o OrganizationAdminAccountMapOutput) ToOrganizationAdminAccountMapOutputWithContext(ctx context.Context) OrganizationAdminAccountMapOutput {
	return o
}

func (o OrganizationAdminAccountMapOutput) MapIndex(k pulumi.StringInput) OrganizationAdminAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OrganizationAdminAccount {
		return vs[0].(map[string]OrganizationAdminAccount)[vs[1].(string)]
	}).(OrganizationAdminAccountOutput)
}

func init() {
	pulumi.RegisterOutputType(OrganizationAdminAccountOutput{})
	pulumi.RegisterOutputType(OrganizationAdminAccountPtrOutput{})
	pulumi.RegisterOutputType(OrganizationAdminAccountArrayOutput{})
	pulumi.RegisterOutputType(OrganizationAdminAccountMapOutput{})
}
