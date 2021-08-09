// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AWS Backup vault resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/backup"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := backup.NewVault(ctx, "example", &backup.VaultArgs{
// 			KmsKeyArn: pulumi.Any(aws_kms_key.Example.Arn),
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
// Backup vault can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:backup/vault:Vault test-vault TestVault
// ```
type Vault struct {
	pulumi.CustomResourceState

	// The ARN of the vault.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The server-side encryption key that is used to protect your backups.
	KmsKeyArn pulumi.StringOutput `pulumi:"kmsKeyArn"`
	// Name of the backup vault to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of recovery points that are stored in a backup vault.
	RecoveryPoints pulumi.IntOutput `pulumi:"recoveryPoints"`
	// Metadata that you can assign to help organize the resources that you create. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewVault registers a new resource with the given unique name, arguments, and options.
func NewVault(ctx *pulumi.Context,
	name string, args *VaultArgs, opts ...pulumi.ResourceOption) (*Vault, error) {
	if args == nil {
		args = &VaultArgs{}
	}

	var resource Vault
	err := ctx.RegisterResource("aws:backup/vault:Vault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVault gets an existing Vault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VaultState, opts ...pulumi.ResourceOption) (*Vault, error) {
	var resource Vault
	err := ctx.ReadResource("aws:backup/vault:Vault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vault resources.
type vaultState struct {
	// The ARN of the vault.
	Arn *string `pulumi:"arn"`
	// The server-side encryption key that is used to protect your backups.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// Name of the backup vault to create.
	Name *string `pulumi:"name"`
	// The number of recovery points that are stored in a backup vault.
	RecoveryPoints *int `pulumi:"recoveryPoints"`
	// Metadata that you can assign to help organize the resources that you create. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type VaultState struct {
	// The ARN of the vault.
	Arn pulumi.StringPtrInput
	// The server-side encryption key that is used to protect your backups.
	KmsKeyArn pulumi.StringPtrInput
	// Name of the backup vault to create.
	Name pulumi.StringPtrInput
	// The number of recovery points that are stored in a backup vault.
	RecoveryPoints pulumi.IntPtrInput
	// Metadata that you can assign to help organize the resources that you create. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (VaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultState)(nil)).Elem()
}

type vaultArgs struct {
	// The server-side encryption key that is used to protect your backups.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// Name of the backup vault to create.
	Name *string `pulumi:"name"`
	// Metadata that you can assign to help organize the resources that you create. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a Vault resource.
type VaultArgs struct {
	// The server-side encryption key that is used to protect your backups.
	KmsKeyArn pulumi.StringPtrInput
	// Name of the backup vault to create.
	Name pulumi.StringPtrInput
	// Metadata that you can assign to help organize the resources that you create. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (VaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultArgs)(nil)).Elem()
}

type VaultInput interface {
	pulumi.Input

	ToVaultOutput() VaultOutput
	ToVaultOutputWithContext(ctx context.Context) VaultOutput
}

func (*Vault) ElementType() reflect.Type {
	return reflect.TypeOf((*Vault)(nil))
}

func (i *Vault) ToVaultOutput() VaultOutput {
	return i.ToVaultOutputWithContext(context.Background())
}

func (i *Vault) ToVaultOutputWithContext(ctx context.Context) VaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultOutput)
}

func (i *Vault) ToVaultPtrOutput() VaultPtrOutput {
	return i.ToVaultPtrOutputWithContext(context.Background())
}

func (i *Vault) ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPtrOutput)
}

type VaultPtrInput interface {
	pulumi.Input

	ToVaultPtrOutput() VaultPtrOutput
	ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput
}

type vaultPtrType VaultArgs

func (*vaultPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Vault)(nil))
}

func (i *vaultPtrType) ToVaultPtrOutput() VaultPtrOutput {
	return i.ToVaultPtrOutputWithContext(context.Background())
}

func (i *vaultPtrType) ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPtrOutput)
}

// VaultArrayInput is an input type that accepts VaultArray and VaultArrayOutput values.
// You can construct a concrete instance of `VaultArrayInput` via:
//
//          VaultArray{ VaultArgs{...} }
type VaultArrayInput interface {
	pulumi.Input

	ToVaultArrayOutput() VaultArrayOutput
	ToVaultArrayOutputWithContext(context.Context) VaultArrayOutput
}

type VaultArray []VaultInput

func (VaultArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vault)(nil)).Elem()
}

func (i VaultArray) ToVaultArrayOutput() VaultArrayOutput {
	return i.ToVaultArrayOutputWithContext(context.Background())
}

func (i VaultArray) ToVaultArrayOutputWithContext(ctx context.Context) VaultArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultArrayOutput)
}

// VaultMapInput is an input type that accepts VaultMap and VaultMapOutput values.
// You can construct a concrete instance of `VaultMapInput` via:
//
//          VaultMap{ "key": VaultArgs{...} }
type VaultMapInput interface {
	pulumi.Input

	ToVaultMapOutput() VaultMapOutput
	ToVaultMapOutputWithContext(context.Context) VaultMapOutput
}

type VaultMap map[string]VaultInput

func (VaultMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vault)(nil)).Elem()
}

func (i VaultMap) ToVaultMapOutput() VaultMapOutput {
	return i.ToVaultMapOutputWithContext(context.Background())
}

func (i VaultMap) ToVaultMapOutputWithContext(ctx context.Context) VaultMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultMapOutput)
}

type VaultOutput struct{ *pulumi.OutputState }

func (VaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Vault)(nil))
}

func (o VaultOutput) ToVaultOutput() VaultOutput {
	return o
}

func (o VaultOutput) ToVaultOutputWithContext(ctx context.Context) VaultOutput {
	return o
}

func (o VaultOutput) ToVaultPtrOutput() VaultPtrOutput {
	return o.ToVaultPtrOutputWithContext(context.Background())
}

func (o VaultOutput) ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Vault) *Vault {
		return &v
	}).(VaultPtrOutput)
}

type VaultPtrOutput struct{ *pulumi.OutputState }

func (VaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vault)(nil))
}

func (o VaultPtrOutput) ToVaultPtrOutput() VaultPtrOutput {
	return o
}

func (o VaultPtrOutput) ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput {
	return o
}

func (o VaultPtrOutput) Elem() VaultOutput {
	return o.ApplyT(func(v *Vault) Vault {
		if v != nil {
			return *v
		}
		var ret Vault
		return ret
	}).(VaultOutput)
}

type VaultArrayOutput struct{ *pulumi.OutputState }

func (VaultArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Vault)(nil))
}

func (o VaultArrayOutput) ToVaultArrayOutput() VaultArrayOutput {
	return o
}

func (o VaultArrayOutput) ToVaultArrayOutputWithContext(ctx context.Context) VaultArrayOutput {
	return o
}

func (o VaultArrayOutput) Index(i pulumi.IntInput) VaultOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Vault {
		return vs[0].([]Vault)[vs[1].(int)]
	}).(VaultOutput)
}

type VaultMapOutput struct{ *pulumi.OutputState }

func (VaultMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Vault)(nil))
}

func (o VaultMapOutput) ToVaultMapOutput() VaultMapOutput {
	return o
}

func (o VaultMapOutput) ToVaultMapOutputWithContext(ctx context.Context) VaultMapOutput {
	return o
}

func (o VaultMapOutput) MapIndex(k pulumi.StringInput) VaultOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Vault {
		return vs[0].(map[string]Vault)[vs[1].(string)]
	}).(VaultOutput)
}

func init() {
	pulumi.RegisterOutputType(VaultOutput{})
	pulumi.RegisterOutputType(VaultPtrOutput{})
	pulumi.RegisterOutputType(VaultArrayOutput{})
	pulumi.RegisterOutputType(VaultMapOutput{})
}
