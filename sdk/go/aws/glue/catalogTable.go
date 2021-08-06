// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Glue Catalog Table Resource. You can refer to the [Glue Developer Guide](http://docs.aws.amazon.com/glue/latest/dg/populate-data-catalog.html) for a full explanation of the Glue Data Catalog functionality.
//
// ## Example Usage
// ### Basic Table
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewCatalogTable(ctx, "awsGlueCatalogTable", &glue.CatalogTableArgs{
// 			DatabaseName: pulumi.String("MyCatalogDatabase"),
// 			Name:         pulumi.String("MyCatalogTable"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Parquet Table for Athena
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewCatalogTable(ctx, "awsGlueCatalogTable", &glue.CatalogTableArgs{
// 			DatabaseName: pulumi.String("MyCatalogDatabase"),
// 			Name:         pulumi.String("MyCatalogTable"),
// 			Parameters: pulumi.StringMap{
// 				"EXTERNAL":            pulumi.String("TRUE"),
// 				"parquet.compression": pulumi.String("SNAPPY"),
// 			},
// 			StorageDescriptor: &glue.CatalogTableStorageDescriptorArgs{
// 				Columns: glue.CatalogTableStorageDescriptorColumnArray{
// 					&glue.CatalogTableStorageDescriptorColumnArgs{
// 						Name: pulumi.String("my_string"),
// 						Type: pulumi.String("string"),
// 					},
// 					&glue.CatalogTableStorageDescriptorColumnArgs{
// 						Name: pulumi.String("my_double"),
// 						Type: pulumi.String("double"),
// 					},
// 					&glue.CatalogTableStorageDescriptorColumnArgs{
// 						Comment: pulumi.String(""),
// 						Name:    pulumi.String("my_date"),
// 						Type:    pulumi.String("date"),
// 					},
// 					&glue.CatalogTableStorageDescriptorColumnArgs{
// 						Comment: pulumi.String(""),
// 						Name:    pulumi.String("my_bigint"),
// 						Type:    pulumi.String("bigint"),
// 					},
// 					&glue.CatalogTableStorageDescriptorColumnArgs{
// 						Comment: pulumi.String(""),
// 						Name:    pulumi.String("my_struct"),
// 						Type:    pulumi.String("struct<my_nested_string:string>"),
// 					},
// 				},
// 				InputFormat:  pulumi.String("org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"),
// 				Location:     pulumi.String("s3://my-bucket/event-streams/my-stream"),
// 				OutputFormat: pulumi.String("org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"),
// 				SerDeInfo: &glue.CatalogTableStorageDescriptorSerDeInfoArgs{
// 					Name: pulumi.String("my-stream"),
// 					Parameters: pulumi.StringMap{
// 						"serialization.format": pulumi.String("1"),
// 					},
// 					SerializationLibrary: pulumi.String("org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"),
// 				},
// 			},
// 			TableType: pulumi.String("EXTERNAL_TABLE"),
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
// Glue Tables can be imported with their catalog ID (usually AWS account ID), database name, and table name, e.g.
//
// ```sh
//  $ pulumi import aws:glue/catalogTable:CatalogTable MyTable 123456789012:MyDatabase:MyTable
// ```
type CatalogTable struct {
	pulumi.CustomResourceState

	// The ARN of the Glue Table.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// ID of the Data Catalog in which the table resides.
	CatalogId pulumi.StringOutput `pulumi:"catalogId"`
	// Name of the catalog database that contains the target table.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// Description of the table.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the target table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Owner of the table.
	Owner pulumi.StringPtrOutput `pulumi:"owner"`
	// Map of initialization parameters for the SerDe, in key-value form.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// Configuration block for a maximum of 3 partition indexes. See `partitionIndex` below.
	PartitionIndices CatalogTablePartitionIndexArrayOutput `pulumi:"partitionIndices"`
	// Configuration block of columns by which the table is partitioned. Only primitive types are supported as partition keys. See `partitionKeys` below.
	PartitionKeys CatalogTablePartitionKeyArrayOutput `pulumi:"partitionKeys"`
	// Retention time for this table.
	Retention pulumi.IntPtrOutput `pulumi:"retention"`
	// Configuration block for information about the physical storage of this table. For more information, refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor). See `storageDescriptor` below.
	StorageDescriptor CatalogTableStorageDescriptorPtrOutput `pulumi:"storageDescriptor"`
	// Type of this table (EXTERNAL_TABLE, VIRTUAL_VIEW, etc.). While optional, some Athena DDL queries such as `ALTER TABLE` and `SHOW CREATE TABLE` will fail if this argument is empty.
	TableType pulumi.StringPtrOutput `pulumi:"tableType"`
	// Configuration block of a target table for resource linking. See `targetTable` below.
	TargetTable CatalogTableTargetTablePtrOutput `pulumi:"targetTable"`
	// If the table is a view, the expanded text of the view; otherwise null.
	ViewExpandedText pulumi.StringPtrOutput `pulumi:"viewExpandedText"`
	// If the table is a view, the original text of the view; otherwise null.
	ViewOriginalText pulumi.StringPtrOutput `pulumi:"viewOriginalText"`
}

// NewCatalogTable registers a new resource with the given unique name, arguments, and options.
func NewCatalogTable(ctx *pulumi.Context,
	name string, args *CatalogTableArgs, opts ...pulumi.ResourceOption) (*CatalogTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	var resource CatalogTable
	err := ctx.RegisterResource("aws:glue/catalogTable:CatalogTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCatalogTable gets an existing CatalogTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCatalogTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CatalogTableState, opts ...pulumi.ResourceOption) (*CatalogTable, error) {
	var resource CatalogTable
	err := ctx.ReadResource("aws:glue/catalogTable:CatalogTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CatalogTable resources.
type catalogTableState struct {
	// The ARN of the Glue Table.
	Arn *string `pulumi:"arn"`
	// ID of the Data Catalog in which the table resides.
	CatalogId *string `pulumi:"catalogId"`
	// Name of the catalog database that contains the target table.
	DatabaseName *string `pulumi:"databaseName"`
	// Description of the table.
	Description *string `pulumi:"description"`
	// Name of the target table.
	Name *string `pulumi:"name"`
	// Owner of the table.
	Owner *string `pulumi:"owner"`
	// Map of initialization parameters for the SerDe, in key-value form.
	Parameters map[string]string `pulumi:"parameters"`
	// Configuration block for a maximum of 3 partition indexes. See `partitionIndex` below.
	PartitionIndices []CatalogTablePartitionIndex `pulumi:"partitionIndices"`
	// Configuration block of columns by which the table is partitioned. Only primitive types are supported as partition keys. See `partitionKeys` below.
	PartitionKeys []CatalogTablePartitionKey `pulumi:"partitionKeys"`
	// Retention time for this table.
	Retention *int `pulumi:"retention"`
	// Configuration block for information about the physical storage of this table. For more information, refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor). See `storageDescriptor` below.
	StorageDescriptor *CatalogTableStorageDescriptor `pulumi:"storageDescriptor"`
	// Type of this table (EXTERNAL_TABLE, VIRTUAL_VIEW, etc.). While optional, some Athena DDL queries such as `ALTER TABLE` and `SHOW CREATE TABLE` will fail if this argument is empty.
	TableType *string `pulumi:"tableType"`
	// Configuration block of a target table for resource linking. See `targetTable` below.
	TargetTable *CatalogTableTargetTable `pulumi:"targetTable"`
	// If the table is a view, the expanded text of the view; otherwise null.
	ViewExpandedText *string `pulumi:"viewExpandedText"`
	// If the table is a view, the original text of the view; otherwise null.
	ViewOriginalText *string `pulumi:"viewOriginalText"`
}

type CatalogTableState struct {
	// The ARN of the Glue Table.
	Arn pulumi.StringPtrInput
	// ID of the Data Catalog in which the table resides.
	CatalogId pulumi.StringPtrInput
	// Name of the catalog database that contains the target table.
	DatabaseName pulumi.StringPtrInput
	// Description of the table.
	Description pulumi.StringPtrInput
	// Name of the target table.
	Name pulumi.StringPtrInput
	// Owner of the table.
	Owner pulumi.StringPtrInput
	// Map of initialization parameters for the SerDe, in key-value form.
	Parameters pulumi.StringMapInput
	// Configuration block for a maximum of 3 partition indexes. See `partitionIndex` below.
	PartitionIndices CatalogTablePartitionIndexArrayInput
	// Configuration block of columns by which the table is partitioned. Only primitive types are supported as partition keys. See `partitionKeys` below.
	PartitionKeys CatalogTablePartitionKeyArrayInput
	// Retention time for this table.
	Retention pulumi.IntPtrInput
	// Configuration block for information about the physical storage of this table. For more information, refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor). See `storageDescriptor` below.
	StorageDescriptor CatalogTableStorageDescriptorPtrInput
	// Type of this table (EXTERNAL_TABLE, VIRTUAL_VIEW, etc.). While optional, some Athena DDL queries such as `ALTER TABLE` and `SHOW CREATE TABLE` will fail if this argument is empty.
	TableType pulumi.StringPtrInput
	// Configuration block of a target table for resource linking. See `targetTable` below.
	TargetTable CatalogTableTargetTablePtrInput
	// If the table is a view, the expanded text of the view; otherwise null.
	ViewExpandedText pulumi.StringPtrInput
	// If the table is a view, the original text of the view; otherwise null.
	ViewOriginalText pulumi.StringPtrInput
}

func (CatalogTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*catalogTableState)(nil)).Elem()
}

type catalogTableArgs struct {
	// ID of the Data Catalog in which the table resides.
	CatalogId *string `pulumi:"catalogId"`
	// Name of the catalog database that contains the target table.
	DatabaseName string `pulumi:"databaseName"`
	// Description of the table.
	Description *string `pulumi:"description"`
	// Name of the target table.
	Name *string `pulumi:"name"`
	// Owner of the table.
	Owner *string `pulumi:"owner"`
	// Map of initialization parameters for the SerDe, in key-value form.
	Parameters map[string]string `pulumi:"parameters"`
	// Configuration block for a maximum of 3 partition indexes. See `partitionIndex` below.
	PartitionIndices []CatalogTablePartitionIndex `pulumi:"partitionIndices"`
	// Configuration block of columns by which the table is partitioned. Only primitive types are supported as partition keys. See `partitionKeys` below.
	PartitionKeys []CatalogTablePartitionKey `pulumi:"partitionKeys"`
	// Retention time for this table.
	Retention *int `pulumi:"retention"`
	// Configuration block for information about the physical storage of this table. For more information, refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor). See `storageDescriptor` below.
	StorageDescriptor *CatalogTableStorageDescriptor `pulumi:"storageDescriptor"`
	// Type of this table (EXTERNAL_TABLE, VIRTUAL_VIEW, etc.). While optional, some Athena DDL queries such as `ALTER TABLE` and `SHOW CREATE TABLE` will fail if this argument is empty.
	TableType *string `pulumi:"tableType"`
	// Configuration block of a target table for resource linking. See `targetTable` below.
	TargetTable *CatalogTableTargetTable `pulumi:"targetTable"`
	// If the table is a view, the expanded text of the view; otherwise null.
	ViewExpandedText *string `pulumi:"viewExpandedText"`
	// If the table is a view, the original text of the view; otherwise null.
	ViewOriginalText *string `pulumi:"viewOriginalText"`
}

// The set of arguments for constructing a CatalogTable resource.
type CatalogTableArgs struct {
	// ID of the Data Catalog in which the table resides.
	CatalogId pulumi.StringPtrInput
	// Name of the catalog database that contains the target table.
	DatabaseName pulumi.StringInput
	// Description of the table.
	Description pulumi.StringPtrInput
	// Name of the target table.
	Name pulumi.StringPtrInput
	// Owner of the table.
	Owner pulumi.StringPtrInput
	// Map of initialization parameters for the SerDe, in key-value form.
	Parameters pulumi.StringMapInput
	// Configuration block for a maximum of 3 partition indexes. See `partitionIndex` below.
	PartitionIndices CatalogTablePartitionIndexArrayInput
	// Configuration block of columns by which the table is partitioned. Only primitive types are supported as partition keys. See `partitionKeys` below.
	PartitionKeys CatalogTablePartitionKeyArrayInput
	// Retention time for this table.
	Retention pulumi.IntPtrInput
	// Configuration block for information about the physical storage of this table. For more information, refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor). See `storageDescriptor` below.
	StorageDescriptor CatalogTableStorageDescriptorPtrInput
	// Type of this table (EXTERNAL_TABLE, VIRTUAL_VIEW, etc.). While optional, some Athena DDL queries such as `ALTER TABLE` and `SHOW CREATE TABLE` will fail if this argument is empty.
	TableType pulumi.StringPtrInput
	// Configuration block of a target table for resource linking. See `targetTable` below.
	TargetTable CatalogTableTargetTablePtrInput
	// If the table is a view, the expanded text of the view; otherwise null.
	ViewExpandedText pulumi.StringPtrInput
	// If the table is a view, the original text of the view; otherwise null.
	ViewOriginalText pulumi.StringPtrInput
}

func (CatalogTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*catalogTableArgs)(nil)).Elem()
}

type CatalogTableInput interface {
	pulumi.Input

	ToCatalogTableOutput() CatalogTableOutput
	ToCatalogTableOutputWithContext(ctx context.Context) CatalogTableOutput
}

func (*CatalogTable) ElementType() reflect.Type {
	return reflect.TypeOf((*CatalogTable)(nil))
}

func (i *CatalogTable) ToCatalogTableOutput() CatalogTableOutput {
	return i.ToCatalogTableOutputWithContext(context.Background())
}

func (i *CatalogTable) ToCatalogTableOutputWithContext(ctx context.Context) CatalogTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogTableOutput)
}

func (i *CatalogTable) ToCatalogTablePtrOutput() CatalogTablePtrOutput {
	return i.ToCatalogTablePtrOutputWithContext(context.Background())
}

func (i *CatalogTable) ToCatalogTablePtrOutputWithContext(ctx context.Context) CatalogTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogTablePtrOutput)
}

type CatalogTablePtrInput interface {
	pulumi.Input

	ToCatalogTablePtrOutput() CatalogTablePtrOutput
	ToCatalogTablePtrOutputWithContext(ctx context.Context) CatalogTablePtrOutput
}

type catalogTablePtrType CatalogTableArgs

func (*catalogTablePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CatalogTable)(nil))
}

func (i *catalogTablePtrType) ToCatalogTablePtrOutput() CatalogTablePtrOutput {
	return i.ToCatalogTablePtrOutputWithContext(context.Background())
}

func (i *catalogTablePtrType) ToCatalogTablePtrOutputWithContext(ctx context.Context) CatalogTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogTablePtrOutput)
}

// CatalogTableArrayInput is an input type that accepts CatalogTableArray and CatalogTableArrayOutput values.
// You can construct a concrete instance of `CatalogTableArrayInput` via:
//
//          CatalogTableArray{ CatalogTableArgs{...} }
type CatalogTableArrayInput interface {
	pulumi.Input

	ToCatalogTableArrayOutput() CatalogTableArrayOutput
	ToCatalogTableArrayOutputWithContext(context.Context) CatalogTableArrayOutput
}

type CatalogTableArray []CatalogTableInput

func (CatalogTableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CatalogTable)(nil)).Elem()
}

func (i CatalogTableArray) ToCatalogTableArrayOutput() CatalogTableArrayOutput {
	return i.ToCatalogTableArrayOutputWithContext(context.Background())
}

func (i CatalogTableArray) ToCatalogTableArrayOutputWithContext(ctx context.Context) CatalogTableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogTableArrayOutput)
}

// CatalogTableMapInput is an input type that accepts CatalogTableMap and CatalogTableMapOutput values.
// You can construct a concrete instance of `CatalogTableMapInput` via:
//
//          CatalogTableMap{ "key": CatalogTableArgs{...} }
type CatalogTableMapInput interface {
	pulumi.Input

	ToCatalogTableMapOutput() CatalogTableMapOutput
	ToCatalogTableMapOutputWithContext(context.Context) CatalogTableMapOutput
}

type CatalogTableMap map[string]CatalogTableInput

func (CatalogTableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CatalogTable)(nil)).Elem()
}

func (i CatalogTableMap) ToCatalogTableMapOutput() CatalogTableMapOutput {
	return i.ToCatalogTableMapOutputWithContext(context.Background())
}

func (i CatalogTableMap) ToCatalogTableMapOutputWithContext(ctx context.Context) CatalogTableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogTableMapOutput)
}

type CatalogTableOutput struct {
	*pulumi.OutputState
}

func (CatalogTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CatalogTable)(nil))
}

func (o CatalogTableOutput) ToCatalogTableOutput() CatalogTableOutput {
	return o
}

func (o CatalogTableOutput) ToCatalogTableOutputWithContext(ctx context.Context) CatalogTableOutput {
	return o
}

func (o CatalogTableOutput) ToCatalogTablePtrOutput() CatalogTablePtrOutput {
	return o.ToCatalogTablePtrOutputWithContext(context.Background())
}

func (o CatalogTableOutput) ToCatalogTablePtrOutputWithContext(ctx context.Context) CatalogTablePtrOutput {
	return o.ApplyT(func(v CatalogTable) *CatalogTable {
		return &v
	}).(CatalogTablePtrOutput)
}

type CatalogTablePtrOutput struct {
	*pulumi.OutputState
}

func (CatalogTablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CatalogTable)(nil))
}

func (o CatalogTablePtrOutput) ToCatalogTablePtrOutput() CatalogTablePtrOutput {
	return o
}

func (o CatalogTablePtrOutput) ToCatalogTablePtrOutputWithContext(ctx context.Context) CatalogTablePtrOutput {
	return o
}

type CatalogTableArrayOutput struct{ *pulumi.OutputState }

func (CatalogTableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CatalogTable)(nil))
}

func (o CatalogTableArrayOutput) ToCatalogTableArrayOutput() CatalogTableArrayOutput {
	return o
}

func (o CatalogTableArrayOutput) ToCatalogTableArrayOutputWithContext(ctx context.Context) CatalogTableArrayOutput {
	return o
}

func (o CatalogTableArrayOutput) Index(i pulumi.IntInput) CatalogTableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CatalogTable {
		return vs[0].([]CatalogTable)[vs[1].(int)]
	}).(CatalogTableOutput)
}

type CatalogTableMapOutput struct{ *pulumi.OutputState }

func (CatalogTableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CatalogTable)(nil))
}

func (o CatalogTableMapOutput) ToCatalogTableMapOutput() CatalogTableMapOutput {
	return o
}

func (o CatalogTableMapOutput) ToCatalogTableMapOutputWithContext(ctx context.Context) CatalogTableMapOutput {
	return o
}

func (o CatalogTableMapOutput) MapIndex(k pulumi.StringInput) CatalogTableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CatalogTable {
		return vs[0].(map[string]CatalogTable)[vs[1].(string)]
	}).(CatalogTableOutput)
}

func init() {
	pulumi.RegisterOutputType(CatalogTableOutput{})
	pulumi.RegisterOutputType(CatalogTablePtrOutput{})
	pulumi.RegisterOutputType(CatalogTableArrayOutput{})
	pulumi.RegisterOutputType(CatalogTableMapOutput{})
}
