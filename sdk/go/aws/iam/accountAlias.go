// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **Note:** There is only a single account alias per AWS account.
//
// Manages the account alias for the AWS Account.
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
// 		_, err := iam.NewAccountAlias(ctx, "alias", &iam.AccountAliasArgs{
// 			AccountAlias: pulumi.String("my-account-alias"),
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
// The current Account Alias can be imported using the `account_alias`, e.g.,
//
// ```sh
//  $ pulumi import aws:iam/accountAlias:AccountAlias alias my-account-alias
// ```
type AccountAlias struct {
	pulumi.CustomResourceState

	// The account alias
	AccountAlias pulumi.StringOutput `pulumi:"accountAlias"`
}

// NewAccountAlias registers a new resource with the given unique name, arguments, and options.
func NewAccountAlias(ctx *pulumi.Context,
	name string, args *AccountAliasArgs, opts ...pulumi.ResourceOption) (*AccountAlias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountAlias == nil {
		return nil, errors.New("invalid value for required argument 'AccountAlias'")
	}
	var resource AccountAlias
	err := ctx.RegisterResource("aws:iam/accountAlias:AccountAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountAlias gets an existing AccountAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountAliasState, opts ...pulumi.ResourceOption) (*AccountAlias, error) {
	var resource AccountAlias
	err := ctx.ReadResource("aws:iam/accountAlias:AccountAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountAlias resources.
type accountAliasState struct {
	// The account alias
	AccountAlias *string `pulumi:"accountAlias"`
}

type AccountAliasState struct {
	// The account alias
	AccountAlias pulumi.StringPtrInput
}

func (AccountAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountAliasState)(nil)).Elem()
}

type accountAliasArgs struct {
	// The account alias
	AccountAlias string `pulumi:"accountAlias"`
}

// The set of arguments for constructing a AccountAlias resource.
type AccountAliasArgs struct {
	// The account alias
	AccountAlias pulumi.StringInput
}

func (AccountAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountAliasArgs)(nil)).Elem()
}

type AccountAliasInput interface {
	pulumi.Input

	ToAccountAliasOutput() AccountAliasOutput
	ToAccountAliasOutputWithContext(ctx context.Context) AccountAliasOutput
}

func (*AccountAlias) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountAlias)(nil))
}

func (i *AccountAlias) ToAccountAliasOutput() AccountAliasOutput {
	return i.ToAccountAliasOutputWithContext(context.Background())
}

func (i *AccountAlias) ToAccountAliasOutputWithContext(ctx context.Context) AccountAliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountAliasOutput)
}

func (i *AccountAlias) ToAccountAliasPtrOutput() AccountAliasPtrOutput {
	return i.ToAccountAliasPtrOutputWithContext(context.Background())
}

func (i *AccountAlias) ToAccountAliasPtrOutputWithContext(ctx context.Context) AccountAliasPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountAliasPtrOutput)
}

type AccountAliasPtrInput interface {
	pulumi.Input

	ToAccountAliasPtrOutput() AccountAliasPtrOutput
	ToAccountAliasPtrOutputWithContext(ctx context.Context) AccountAliasPtrOutput
}

type accountAliasPtrType AccountAliasArgs

func (*accountAliasPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountAlias)(nil))
}

func (i *accountAliasPtrType) ToAccountAliasPtrOutput() AccountAliasPtrOutput {
	return i.ToAccountAliasPtrOutputWithContext(context.Background())
}

func (i *accountAliasPtrType) ToAccountAliasPtrOutputWithContext(ctx context.Context) AccountAliasPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountAliasPtrOutput)
}

// AccountAliasArrayInput is an input type that accepts AccountAliasArray and AccountAliasArrayOutput values.
// You can construct a concrete instance of `AccountAliasArrayInput` via:
//
//          AccountAliasArray{ AccountAliasArgs{...} }
type AccountAliasArrayInput interface {
	pulumi.Input

	ToAccountAliasArrayOutput() AccountAliasArrayOutput
	ToAccountAliasArrayOutputWithContext(context.Context) AccountAliasArrayOutput
}

type AccountAliasArray []AccountAliasInput

func (AccountAliasArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountAlias)(nil)).Elem()
}

func (i AccountAliasArray) ToAccountAliasArrayOutput() AccountAliasArrayOutput {
	return i.ToAccountAliasArrayOutputWithContext(context.Background())
}

func (i AccountAliasArray) ToAccountAliasArrayOutputWithContext(ctx context.Context) AccountAliasArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountAliasArrayOutput)
}

// AccountAliasMapInput is an input type that accepts AccountAliasMap and AccountAliasMapOutput values.
// You can construct a concrete instance of `AccountAliasMapInput` via:
//
//          AccountAliasMap{ "key": AccountAliasArgs{...} }
type AccountAliasMapInput interface {
	pulumi.Input

	ToAccountAliasMapOutput() AccountAliasMapOutput
	ToAccountAliasMapOutputWithContext(context.Context) AccountAliasMapOutput
}

type AccountAliasMap map[string]AccountAliasInput

func (AccountAliasMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountAlias)(nil)).Elem()
}

func (i AccountAliasMap) ToAccountAliasMapOutput() AccountAliasMapOutput {
	return i.ToAccountAliasMapOutputWithContext(context.Background())
}

func (i AccountAliasMap) ToAccountAliasMapOutputWithContext(ctx context.Context) AccountAliasMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountAliasMapOutput)
}

type AccountAliasOutput struct{ *pulumi.OutputState }

func (AccountAliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountAlias)(nil))
}

func (o AccountAliasOutput) ToAccountAliasOutput() AccountAliasOutput {
	return o
}

func (o AccountAliasOutput) ToAccountAliasOutputWithContext(ctx context.Context) AccountAliasOutput {
	return o
}

func (o AccountAliasOutput) ToAccountAliasPtrOutput() AccountAliasPtrOutput {
	return o.ToAccountAliasPtrOutputWithContext(context.Background())
}

func (o AccountAliasOutput) ToAccountAliasPtrOutputWithContext(ctx context.Context) AccountAliasPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccountAlias) *AccountAlias {
		return &v
	}).(AccountAliasPtrOutput)
}

type AccountAliasPtrOutput struct{ *pulumi.OutputState }

func (AccountAliasPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountAlias)(nil))
}

func (o AccountAliasPtrOutput) ToAccountAliasPtrOutput() AccountAliasPtrOutput {
	return o
}

func (o AccountAliasPtrOutput) ToAccountAliasPtrOutputWithContext(ctx context.Context) AccountAliasPtrOutput {
	return o
}

func (o AccountAliasPtrOutput) Elem() AccountAliasOutput {
	return o.ApplyT(func(v *AccountAlias) AccountAlias {
		if v != nil {
			return *v
		}
		var ret AccountAlias
		return ret
	}).(AccountAliasOutput)
}

type AccountAliasArrayOutput struct{ *pulumi.OutputState }

func (AccountAliasArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccountAlias)(nil))
}

func (o AccountAliasArrayOutput) ToAccountAliasArrayOutput() AccountAliasArrayOutput {
	return o
}

func (o AccountAliasArrayOutput) ToAccountAliasArrayOutputWithContext(ctx context.Context) AccountAliasArrayOutput {
	return o
}

func (o AccountAliasArrayOutput) Index(i pulumi.IntInput) AccountAliasOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccountAlias {
		return vs[0].([]AccountAlias)[vs[1].(int)]
	}).(AccountAliasOutput)
}

type AccountAliasMapOutput struct{ *pulumi.OutputState }

func (AccountAliasMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AccountAlias)(nil))
}

func (o AccountAliasMapOutput) ToAccountAliasMapOutput() AccountAliasMapOutput {
	return o
}

func (o AccountAliasMapOutput) ToAccountAliasMapOutputWithContext(ctx context.Context) AccountAliasMapOutput {
	return o
}

func (o AccountAliasMapOutput) MapIndex(k pulumi.StringInput) AccountAliasOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AccountAlias {
		return vs[0].(map[string]AccountAlias)[vs[1].(string)]
	}).(AccountAliasOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountAliasInput)(nil)).Elem(), &AccountAlias{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountAliasPtrInput)(nil)).Elem(), &AccountAlias{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountAliasArrayInput)(nil)).Elem(), AccountAliasArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountAliasMapInput)(nil)).Elem(), AccountAliasMap{})
	pulumi.RegisterOutputType(AccountAliasOutput{})
	pulumi.RegisterOutputType(AccountAliasPtrOutput{})
	pulumi.RegisterOutputType(AccountAliasArrayOutput{})
	pulumi.RegisterOutputType(AccountAliasMapOutput{})
}
