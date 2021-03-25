// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lakeformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages Lake Formation principals designated as data lake administrators and lists of principal permission entries for default create database and default create table permissions.
//
// > **NOTE:** Lake Formation introduces fine-grained access control for data in your data lake. Part of the changes include the `IAMAllowedPrincipals` principal in order to make Lake Formation backwards compatible with existing IAM and Glue permissions. For more information, see [Changing the Default Security Settings for Your Data Lake](https://docs.aws.amazon.com/lake-formation/latest/dg/change-settings.html) and [Upgrading AWS Glue Data Permissions to the AWS Lake Formation Model](https://docs.aws.amazon.com/lake-formation/latest/dg/upgrade-glue-lake-formation.html).
//
// ## Example Usage
type DataLakeSettings struct {
	pulumi.CustomResourceState

	// Set of ARNs of AWS Lake Formation principals (IAM users or roles).
	Admins pulumi.StringArrayOutput `pulumi:"admins"`
	// Identifier for the Data Catalog. By default, the account ID.
	CatalogId pulumi.StringPtrOutput `pulumi:"catalogId"`
	// Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
	CreateDatabaseDefaultPermissions DataLakeSettingsCreateDatabaseDefaultPermissionArrayOutput `pulumi:"createDatabaseDefaultPermissions"`
	// Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
	CreateTableDefaultPermissions DataLakeSettingsCreateTableDefaultPermissionArrayOutput `pulumi:"createTableDefaultPermissions"`
	// List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
	TrustedResourceOwners pulumi.StringArrayOutput `pulumi:"trustedResourceOwners"`
}

// NewDataLakeSettings registers a new resource with the given unique name, arguments, and options.
func NewDataLakeSettings(ctx *pulumi.Context,
	name string, args *DataLakeSettingsArgs, opts ...pulumi.ResourceOption) (*DataLakeSettings, error) {
	if args == nil {
		args = &DataLakeSettingsArgs{}
	}

	var resource DataLakeSettings
	err := ctx.RegisterResource("aws:lakeformation/dataLakeSettings:DataLakeSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataLakeSettings gets an existing DataLakeSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataLakeSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataLakeSettingsState, opts ...pulumi.ResourceOption) (*DataLakeSettings, error) {
	var resource DataLakeSettings
	err := ctx.ReadResource("aws:lakeformation/dataLakeSettings:DataLakeSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataLakeSettings resources.
type dataLakeSettingsState struct {
	// Set of ARNs of AWS Lake Formation principals (IAM users or roles).
	Admins []string `pulumi:"admins"`
	// Identifier for the Data Catalog. By default, the account ID.
	CatalogId *string `pulumi:"catalogId"`
	// Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
	CreateDatabaseDefaultPermissions []DataLakeSettingsCreateDatabaseDefaultPermission `pulumi:"createDatabaseDefaultPermissions"`
	// Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
	CreateTableDefaultPermissions []DataLakeSettingsCreateTableDefaultPermission `pulumi:"createTableDefaultPermissions"`
	// List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
	TrustedResourceOwners []string `pulumi:"trustedResourceOwners"`
}

type DataLakeSettingsState struct {
	// Set of ARNs of AWS Lake Formation principals (IAM users or roles).
	Admins pulumi.StringArrayInput
	// Identifier for the Data Catalog. By default, the account ID.
	CatalogId pulumi.StringPtrInput
	// Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
	CreateDatabaseDefaultPermissions DataLakeSettingsCreateDatabaseDefaultPermissionArrayInput
	// Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
	CreateTableDefaultPermissions DataLakeSettingsCreateTableDefaultPermissionArrayInput
	// List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
	TrustedResourceOwners pulumi.StringArrayInput
}

func (DataLakeSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataLakeSettingsState)(nil)).Elem()
}

type dataLakeSettingsArgs struct {
	// Set of ARNs of AWS Lake Formation principals (IAM users or roles).
	Admins []string `pulumi:"admins"`
	// Identifier for the Data Catalog. By default, the account ID.
	CatalogId *string `pulumi:"catalogId"`
	// Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
	CreateDatabaseDefaultPermissions []DataLakeSettingsCreateDatabaseDefaultPermission `pulumi:"createDatabaseDefaultPermissions"`
	// Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
	CreateTableDefaultPermissions []DataLakeSettingsCreateTableDefaultPermission `pulumi:"createTableDefaultPermissions"`
	// List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
	TrustedResourceOwners []string `pulumi:"trustedResourceOwners"`
}

// The set of arguments for constructing a DataLakeSettings resource.
type DataLakeSettingsArgs struct {
	// Set of ARNs of AWS Lake Formation principals (IAM users or roles).
	Admins pulumi.StringArrayInput
	// Identifier for the Data Catalog. By default, the account ID.
	CatalogId pulumi.StringPtrInput
	// Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
	CreateDatabaseDefaultPermissions DataLakeSettingsCreateDatabaseDefaultPermissionArrayInput
	// Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
	CreateTableDefaultPermissions DataLakeSettingsCreateTableDefaultPermissionArrayInput
	// List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
	TrustedResourceOwners pulumi.StringArrayInput
}

func (DataLakeSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataLakeSettingsArgs)(nil)).Elem()
}

type DataLakeSettingsInput interface {
	pulumi.Input

	ToDataLakeSettingsOutput() DataLakeSettingsOutput
	ToDataLakeSettingsOutputWithContext(ctx context.Context) DataLakeSettingsOutput
}

func (*DataLakeSettings) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeSettings)(nil))
}

func (i *DataLakeSettings) ToDataLakeSettingsOutput() DataLakeSettingsOutput {
	return i.ToDataLakeSettingsOutputWithContext(context.Background())
}

func (i *DataLakeSettings) ToDataLakeSettingsOutputWithContext(ctx context.Context) DataLakeSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeSettingsOutput)
}

func (i *DataLakeSettings) ToDataLakeSettingsPtrOutput() DataLakeSettingsPtrOutput {
	return i.ToDataLakeSettingsPtrOutputWithContext(context.Background())
}

func (i *DataLakeSettings) ToDataLakeSettingsPtrOutputWithContext(ctx context.Context) DataLakeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeSettingsPtrOutput)
}

type DataLakeSettingsPtrInput interface {
	pulumi.Input

	ToDataLakeSettingsPtrOutput() DataLakeSettingsPtrOutput
	ToDataLakeSettingsPtrOutputWithContext(ctx context.Context) DataLakeSettingsPtrOutput
}

type dataLakeSettingsPtrType DataLakeSettingsArgs

func (*dataLakeSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLakeSettings)(nil))
}

func (i *dataLakeSettingsPtrType) ToDataLakeSettingsPtrOutput() DataLakeSettingsPtrOutput {
	return i.ToDataLakeSettingsPtrOutputWithContext(context.Background())
}

func (i *dataLakeSettingsPtrType) ToDataLakeSettingsPtrOutputWithContext(ctx context.Context) DataLakeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeSettingsPtrOutput)
}

// DataLakeSettingsArrayInput is an input type that accepts DataLakeSettingsArray and DataLakeSettingsArrayOutput values.
// You can construct a concrete instance of `DataLakeSettingsArrayInput` via:
//
//          DataLakeSettingsArray{ DataLakeSettingsArgs{...} }
type DataLakeSettingsArrayInput interface {
	pulumi.Input

	ToDataLakeSettingsArrayOutput() DataLakeSettingsArrayOutput
	ToDataLakeSettingsArrayOutputWithContext(context.Context) DataLakeSettingsArrayOutput
}

type DataLakeSettingsArray []DataLakeSettingsInput

func (DataLakeSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DataLakeSettings)(nil))
}

func (i DataLakeSettingsArray) ToDataLakeSettingsArrayOutput() DataLakeSettingsArrayOutput {
	return i.ToDataLakeSettingsArrayOutputWithContext(context.Background())
}

func (i DataLakeSettingsArray) ToDataLakeSettingsArrayOutputWithContext(ctx context.Context) DataLakeSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeSettingsArrayOutput)
}

// DataLakeSettingsMapInput is an input type that accepts DataLakeSettingsMap and DataLakeSettingsMapOutput values.
// You can construct a concrete instance of `DataLakeSettingsMapInput` via:
//
//          DataLakeSettingsMap{ "key": DataLakeSettingsArgs{...} }
type DataLakeSettingsMapInput interface {
	pulumi.Input

	ToDataLakeSettingsMapOutput() DataLakeSettingsMapOutput
	ToDataLakeSettingsMapOutputWithContext(context.Context) DataLakeSettingsMapOutput
}

type DataLakeSettingsMap map[string]DataLakeSettingsInput

func (DataLakeSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DataLakeSettings)(nil))
}

func (i DataLakeSettingsMap) ToDataLakeSettingsMapOutput() DataLakeSettingsMapOutput {
	return i.ToDataLakeSettingsMapOutputWithContext(context.Background())
}

func (i DataLakeSettingsMap) ToDataLakeSettingsMapOutputWithContext(ctx context.Context) DataLakeSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeSettingsMapOutput)
}

type DataLakeSettingsOutput struct {
	*pulumi.OutputState
}

func (DataLakeSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeSettings)(nil))
}

func (o DataLakeSettingsOutput) ToDataLakeSettingsOutput() DataLakeSettingsOutput {
	return o
}

func (o DataLakeSettingsOutput) ToDataLakeSettingsOutputWithContext(ctx context.Context) DataLakeSettingsOutput {
	return o
}

func (o DataLakeSettingsOutput) ToDataLakeSettingsPtrOutput() DataLakeSettingsPtrOutput {
	return o.ToDataLakeSettingsPtrOutputWithContext(context.Background())
}

func (o DataLakeSettingsOutput) ToDataLakeSettingsPtrOutputWithContext(ctx context.Context) DataLakeSettingsPtrOutput {
	return o.ApplyT(func(v DataLakeSettings) *DataLakeSettings {
		return &v
	}).(DataLakeSettingsPtrOutput)
}

type DataLakeSettingsPtrOutput struct {
	*pulumi.OutputState
}

func (DataLakeSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLakeSettings)(nil))
}

func (o DataLakeSettingsPtrOutput) ToDataLakeSettingsPtrOutput() DataLakeSettingsPtrOutput {
	return o
}

func (o DataLakeSettingsPtrOutput) ToDataLakeSettingsPtrOutputWithContext(ctx context.Context) DataLakeSettingsPtrOutput {
	return o
}

type DataLakeSettingsArrayOutput struct{ *pulumi.OutputState }

func (DataLakeSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataLakeSettings)(nil))
}

func (o DataLakeSettingsArrayOutput) ToDataLakeSettingsArrayOutput() DataLakeSettingsArrayOutput {
	return o
}

func (o DataLakeSettingsArrayOutput) ToDataLakeSettingsArrayOutputWithContext(ctx context.Context) DataLakeSettingsArrayOutput {
	return o
}

func (o DataLakeSettingsArrayOutput) Index(i pulumi.IntInput) DataLakeSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataLakeSettings {
		return vs[0].([]DataLakeSettings)[vs[1].(int)]
	}).(DataLakeSettingsOutput)
}

type DataLakeSettingsMapOutput struct{ *pulumi.OutputState }

func (DataLakeSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DataLakeSettings)(nil))
}

func (o DataLakeSettingsMapOutput) ToDataLakeSettingsMapOutput() DataLakeSettingsMapOutput {
	return o
}

func (o DataLakeSettingsMapOutput) ToDataLakeSettingsMapOutputWithContext(ctx context.Context) DataLakeSettingsMapOutput {
	return o
}

func (o DataLakeSettingsMapOutput) MapIndex(k pulumi.StringInput) DataLakeSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DataLakeSettings {
		return vs[0].(map[string]DataLakeSettings)[vs[1].(string)]
	}).(DataLakeSettingsOutput)
}

func init() {
	pulumi.RegisterOutputType(DataLakeSettingsOutput{})
	pulumi.RegisterOutputType(DataLakeSettingsPtrOutput{})
	pulumi.RegisterOutputType(DataLakeSettingsArrayOutput{})
	pulumi.RegisterOutputType(DataLakeSettingsMapOutput{})
}
