// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Glue Partition Resource.
//
// ## Import
//
// Glue Partitions can be imported with their catalog ID (usually AWS account ID), database name, table name and partition values e.g.,
//
// ```sh
//  $ pulumi import aws:glue/partition:Partition part 123456789012:MyDatabase:MyTable:val1#val2
// ```
type Partition struct {
	pulumi.CustomResourceState

	// ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
	CatalogId pulumi.StringOutput `pulumi:"catalogId"`
	// The time at which the partition was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The last time at which the partition was accessed.
	LastAccessedTime pulumi.StringOutput `pulumi:"lastAccessedTime"`
	// The last time at which column statistics were computed for this partition.
	LastAnalyzedTime pulumi.StringOutput `pulumi:"lastAnalyzedTime"`
	// A map of initialization parameters for the SerDe, in key-value form.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The values that define the partition.
	PartitionValues pulumi.StringArrayOutput `pulumi:"partitionValues"`
	// A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
	StorageDescriptor PartitionStorageDescriptorPtrOutput `pulumi:"storageDescriptor"`
	TableName         pulumi.StringOutput                 `pulumi:"tableName"`
}

// NewPartition registers a new resource with the given unique name, arguments, and options.
func NewPartition(ctx *pulumi.Context,
	name string, args *PartitionArgs, opts ...pulumi.ResourceOption) (*Partition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.PartitionValues == nil {
		return nil, errors.New("invalid value for required argument 'PartitionValues'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	var resource Partition
	err := ctx.RegisterResource("aws:glue/partition:Partition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPartition gets an existing Partition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPartition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartitionState, opts ...pulumi.ResourceOption) (*Partition, error) {
	var resource Partition
	err := ctx.ReadResource("aws:glue/partition:Partition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Partition resources.
type partitionState struct {
	// ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
	CatalogId *string `pulumi:"catalogId"`
	// The time at which the partition was created.
	CreationTime *string `pulumi:"creationTime"`
	// Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
	DatabaseName *string `pulumi:"databaseName"`
	// The last time at which the partition was accessed.
	LastAccessedTime *string `pulumi:"lastAccessedTime"`
	// The last time at which column statistics were computed for this partition.
	LastAnalyzedTime *string `pulumi:"lastAnalyzedTime"`
	// A map of initialization parameters for the SerDe, in key-value form.
	Parameters map[string]string `pulumi:"parameters"`
	// The values that define the partition.
	PartitionValues []string `pulumi:"partitionValues"`
	// A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
	StorageDescriptor *PartitionStorageDescriptor `pulumi:"storageDescriptor"`
	TableName         *string                     `pulumi:"tableName"`
}

type PartitionState struct {
	// ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
	CatalogId pulumi.StringPtrInput
	// The time at which the partition was created.
	CreationTime pulumi.StringPtrInput
	// Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
	DatabaseName pulumi.StringPtrInput
	// The last time at which the partition was accessed.
	LastAccessedTime pulumi.StringPtrInput
	// The last time at which column statistics were computed for this partition.
	LastAnalyzedTime pulumi.StringPtrInput
	// A map of initialization parameters for the SerDe, in key-value form.
	Parameters pulumi.StringMapInput
	// The values that define the partition.
	PartitionValues pulumi.StringArrayInput
	// A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
	StorageDescriptor PartitionStorageDescriptorPtrInput
	TableName         pulumi.StringPtrInput
}

func (PartitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*partitionState)(nil)).Elem()
}

type partitionArgs struct {
	// ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
	CatalogId *string `pulumi:"catalogId"`
	// Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
	DatabaseName string `pulumi:"databaseName"`
	// A map of initialization parameters for the SerDe, in key-value form.
	Parameters map[string]string `pulumi:"parameters"`
	// The values that define the partition.
	PartitionValues []string `pulumi:"partitionValues"`
	// A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
	StorageDescriptor *PartitionStorageDescriptor `pulumi:"storageDescriptor"`
	TableName         string                      `pulumi:"tableName"`
}

// The set of arguments for constructing a Partition resource.
type PartitionArgs struct {
	// ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
	CatalogId pulumi.StringPtrInput
	// Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
	DatabaseName pulumi.StringInput
	// A map of initialization parameters for the SerDe, in key-value form.
	Parameters pulumi.StringMapInput
	// The values that define the partition.
	PartitionValues pulumi.StringArrayInput
	// A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
	StorageDescriptor PartitionStorageDescriptorPtrInput
	TableName         pulumi.StringInput
}

func (PartitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partitionArgs)(nil)).Elem()
}

type PartitionInput interface {
	pulumi.Input

	ToPartitionOutput() PartitionOutput
	ToPartitionOutputWithContext(ctx context.Context) PartitionOutput
}

func (*Partition) ElementType() reflect.Type {
	return reflect.TypeOf((*Partition)(nil))
}

func (i *Partition) ToPartitionOutput() PartitionOutput {
	return i.ToPartitionOutputWithContext(context.Background())
}

func (i *Partition) ToPartitionOutputWithContext(ctx context.Context) PartitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionOutput)
}

func (i *Partition) ToPartitionPtrOutput() PartitionPtrOutput {
	return i.ToPartitionPtrOutputWithContext(context.Background())
}

func (i *Partition) ToPartitionPtrOutputWithContext(ctx context.Context) PartitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionPtrOutput)
}

type PartitionPtrInput interface {
	pulumi.Input

	ToPartitionPtrOutput() PartitionPtrOutput
	ToPartitionPtrOutputWithContext(ctx context.Context) PartitionPtrOutput
}

type partitionPtrType PartitionArgs

func (*partitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Partition)(nil))
}

func (i *partitionPtrType) ToPartitionPtrOutput() PartitionPtrOutput {
	return i.ToPartitionPtrOutputWithContext(context.Background())
}

func (i *partitionPtrType) ToPartitionPtrOutputWithContext(ctx context.Context) PartitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionPtrOutput)
}

// PartitionArrayInput is an input type that accepts PartitionArray and PartitionArrayOutput values.
// You can construct a concrete instance of `PartitionArrayInput` via:
//
//          PartitionArray{ PartitionArgs{...} }
type PartitionArrayInput interface {
	pulumi.Input

	ToPartitionArrayOutput() PartitionArrayOutput
	ToPartitionArrayOutputWithContext(context.Context) PartitionArrayOutput
}

type PartitionArray []PartitionInput

func (PartitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Partition)(nil)).Elem()
}

func (i PartitionArray) ToPartitionArrayOutput() PartitionArrayOutput {
	return i.ToPartitionArrayOutputWithContext(context.Background())
}

func (i PartitionArray) ToPartitionArrayOutputWithContext(ctx context.Context) PartitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionArrayOutput)
}

// PartitionMapInput is an input type that accepts PartitionMap and PartitionMapOutput values.
// You can construct a concrete instance of `PartitionMapInput` via:
//
//          PartitionMap{ "key": PartitionArgs{...} }
type PartitionMapInput interface {
	pulumi.Input

	ToPartitionMapOutput() PartitionMapOutput
	ToPartitionMapOutputWithContext(context.Context) PartitionMapOutput
}

type PartitionMap map[string]PartitionInput

func (PartitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Partition)(nil)).Elem()
}

func (i PartitionMap) ToPartitionMapOutput() PartitionMapOutput {
	return i.ToPartitionMapOutputWithContext(context.Background())
}

func (i PartitionMap) ToPartitionMapOutputWithContext(ctx context.Context) PartitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionMapOutput)
}

type PartitionOutput struct{ *pulumi.OutputState }

func (PartitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Partition)(nil))
}

func (o PartitionOutput) ToPartitionOutput() PartitionOutput {
	return o
}

func (o PartitionOutput) ToPartitionOutputWithContext(ctx context.Context) PartitionOutput {
	return o
}

func (o PartitionOutput) ToPartitionPtrOutput() PartitionPtrOutput {
	return o.ToPartitionPtrOutputWithContext(context.Background())
}

func (o PartitionOutput) ToPartitionPtrOutputWithContext(ctx context.Context) PartitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Partition) *Partition {
		return &v
	}).(PartitionPtrOutput)
}

type PartitionPtrOutput struct{ *pulumi.OutputState }

func (PartitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Partition)(nil))
}

func (o PartitionPtrOutput) ToPartitionPtrOutput() PartitionPtrOutput {
	return o
}

func (o PartitionPtrOutput) ToPartitionPtrOutputWithContext(ctx context.Context) PartitionPtrOutput {
	return o
}

func (o PartitionPtrOutput) Elem() PartitionOutput {
	return o.ApplyT(func(v *Partition) Partition {
		if v != nil {
			return *v
		}
		var ret Partition
		return ret
	}).(PartitionOutput)
}

type PartitionArrayOutput struct{ *pulumi.OutputState }

func (PartitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Partition)(nil))
}

func (o PartitionArrayOutput) ToPartitionArrayOutput() PartitionArrayOutput {
	return o
}

func (o PartitionArrayOutput) ToPartitionArrayOutputWithContext(ctx context.Context) PartitionArrayOutput {
	return o
}

func (o PartitionArrayOutput) Index(i pulumi.IntInput) PartitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Partition {
		return vs[0].([]Partition)[vs[1].(int)]
	}).(PartitionOutput)
}

type PartitionMapOutput struct{ *pulumi.OutputState }

func (PartitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Partition)(nil))
}

func (o PartitionMapOutput) ToPartitionMapOutput() PartitionMapOutput {
	return o
}

func (o PartitionMapOutput) ToPartitionMapOutputWithContext(ctx context.Context) PartitionMapOutput {
	return o
}

func (o PartitionMapOutput) MapIndex(k pulumi.StringInput) PartitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Partition {
		return vs[0].(map[string]Partition)[vs[1].(string)]
	}).(PartitionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PartitionInput)(nil)).Elem(), &Partition{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartitionPtrInput)(nil)).Elem(), &Partition{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartitionArrayInput)(nil)).Elem(), PartitionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartitionMapInput)(nil)).Elem(), PartitionMap{})
	pulumi.RegisterOutputType(PartitionOutput{})
	pulumi.RegisterOutputType(PartitionPtrOutput{})
	pulumi.RegisterOutputType(PartitionArrayOutput{})
	pulumi.RegisterOutputType(PartitionMapOutput{})
}
