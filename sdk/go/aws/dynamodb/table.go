// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dynamodb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DynamoDB table resource
//
// > **Note:** It is recommended to use [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) for `readCapacity` and/or `writeCapacity` if there's `autoscaling policy` attached to the table.
//
// ## Example Usage
//
// The following dynamodb table description models the table and GSI shown
// in the [AWS SDK example documentation](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.html)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/dynamodb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dynamodb.NewTable(ctx, "basic_dynamodb_table", &dynamodb.TableArgs{
// 			Attributes: dynamodb.TableAttributeArray{
// 				&dynamodb.TableAttributeArgs{
// 					Name: pulumi.String("UserId"),
// 					Type: pulumi.String("S"),
// 				},
// 				&dynamodb.TableAttributeArgs{
// 					Name: pulumi.String("GameTitle"),
// 					Type: pulumi.String("S"),
// 				},
// 				&dynamodb.TableAttributeArgs{
// 					Name: pulumi.String("TopScore"),
// 					Type: pulumi.String("N"),
// 				},
// 			},
// 			BillingMode: pulumi.String("PROVISIONED"),
// 			GlobalSecondaryIndexes: dynamodb.TableGlobalSecondaryIndexArray{
// 				&dynamodb.TableGlobalSecondaryIndexArgs{
// 					HashKey: pulumi.String("GameTitle"),
// 					Name:    pulumi.String("GameTitleIndex"),
// 					NonKeyAttributes: pulumi.StringArray{
// 						pulumi.String("UserId"),
// 					},
// 					ProjectionType: pulumi.String("INCLUDE"),
// 					RangeKey:       pulumi.String("TopScore"),
// 					ReadCapacity:   pulumi.Int(10),
// 					WriteCapacity:  pulumi.Int(10),
// 				},
// 			},
// 			HashKey:      pulumi.String("UserId"),
// 			RangeKey:     pulumi.String("GameTitle"),
// 			ReadCapacity: pulumi.Int(20),
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("production"),
// 				"Name":        pulumi.String("dynamodb-table-1"),
// 			},
// 			Ttl: &dynamodb.TableTtlArgs{
// 				AttributeName: pulumi.String("TimeToExist"),
// 				Enabled:       pulumi.Bool(false),
// 			},
// 			WriteCapacity: pulumi.Int(20),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Global Tables
//
// This resource implements support for [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) via `replica` configuration blocks. For working with [DynamoDB Global Tables V1 (version 2017.11.29)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html), see the `dynamodb.GlobalTable` resource.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/dynamodb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dynamodb.NewTable(ctx, "example", &dynamodb.TableArgs{
// 			Attributes: dynamodb.TableAttributeArray{
// 				&dynamodb.TableAttributeArgs{
// 					Name: pulumi.String("TestTableHashKey"),
// 					Type: pulumi.String("S"),
// 				},
// 			},
// 			BillingMode: pulumi.String("PAY_PER_REQUEST"),
// 			HashKey:     pulumi.String("TestTableHashKey"),
// 			Replicas: dynamodb.TableReplicaArray{
// 				&dynamodb.TableReplicaArgs{
// 					RegionName: pulumi.String("us-east-2"),
// 				},
// 				&dynamodb.TableReplicaArgs{
// 					RegionName: pulumi.String("us-west-2"),
// 				},
// 			},
// 			StreamEnabled:  pulumi.Bool(true),
// 			StreamViewType: pulumi.String("NEW_AND_OLD_IMAGES"),
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
// DynamoDB tables can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:dynamodb/table:Table basic-dynamodb-table GameScores
// ```
type Table struct {
	pulumi.CustomResourceState

	// The arn of the table
	Arn pulumi.StringOutput `pulumi:"arn"`
	// List of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes. Each attribute has two properties:
	Attributes TableAttributeArrayOutput `pulumi:"attributes"`
	// Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
	BillingMode pulumi.StringPtrOutput `pulumi:"billingMode"`
	// Describe a GSI for the table;
	// subject to the normal limits on the number of GSIs, projected
	// attributes, etc.
	GlobalSecondaryIndexes TableGlobalSecondaryIndexArrayOutput `pulumi:"globalSecondaryIndexes"`
	// The name of the hash key in the index; must be
	// defined as an attribute in the resource.
	HashKey pulumi.StringOutput `pulumi:"hashKey"`
	// Describe an LSI on the table;
	// these can only be allocated *at creation* so you cannot change this
	// definition after you have created the resource.
	LocalSecondaryIndexes TableLocalSecondaryIndexArrayOutput `pulumi:"localSecondaryIndexes"`
	// The name of the index
	Name pulumi.StringOutput `pulumi:"name"`
	// Point-in-time recovery options.
	PointInTimeRecovery TablePointInTimeRecoveryOutput `pulumi:"pointInTimeRecovery"`
	// The name of the range key; must be defined
	RangeKey pulumi.StringPtrOutput `pulumi:"rangeKey"`
	// The number of read units for this index. Must be set if billingMode is set to PROVISIONED.
	ReadCapacity pulumi.IntPtrOutput `pulumi:"readCapacity"`
	// Configuration block(s) with [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) replication configurations. Detailed below.
	Replicas TableReplicaArrayOutput `pulumi:"replicas"`
	// Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS owned Customer Master Key if this argument isn't specified.
	ServerSideEncryption TableServerSideEncryptionOutput `pulumi:"serverSideEncryption"`
	// The ARN of the Table Stream. Only available when `streamEnabled = true`
	StreamArn pulumi.StringOutput `pulumi:"streamArn"`
	// Indicates whether Streams are to be enabled (true) or disabled (false).
	StreamEnabled pulumi.BoolPtrOutput `pulumi:"streamEnabled"`
	// A timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not
	// a unique identifier for the stream on its own. However, the combination of AWS customer ID,
	// table name and this field is guaranteed to be unique.
	// It can be used for creating CloudWatch Alarms. Only available when `streamEnabled = true`
	StreamLabel pulumi.StringOutput `pulumi:"streamLabel"`
	// When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
	StreamViewType pulumi.StringOutput `pulumi:"streamViewType"`
	// A map of tags to populate on the created table. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Defines ttl, has two properties, and can only be specified once:
	Ttl TableTtlPtrOutput `pulumi:"ttl"`
	// The number of write units for this index. Must be set if billingMode is set to PROVISIONED.
	WriteCapacity pulumi.IntPtrOutput `pulumi:"writeCapacity"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Attributes == nil {
		return nil, errors.New("invalid value for required argument 'Attributes'")
	}
	if args.HashKey == nil {
		return nil, errors.New("invalid value for required argument 'HashKey'")
	}
	var resource Table
	err := ctx.RegisterResource("aws:dynamodb/table:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("aws:dynamodb/table:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
	// The arn of the table
	Arn *string `pulumi:"arn"`
	// List of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes. Each attribute has two properties:
	Attributes []TableAttribute `pulumi:"attributes"`
	// Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
	BillingMode *string `pulumi:"billingMode"`
	// Describe a GSI for the table;
	// subject to the normal limits on the number of GSIs, projected
	// attributes, etc.
	GlobalSecondaryIndexes []TableGlobalSecondaryIndex `pulumi:"globalSecondaryIndexes"`
	// The name of the hash key in the index; must be
	// defined as an attribute in the resource.
	HashKey *string `pulumi:"hashKey"`
	// Describe an LSI on the table;
	// these can only be allocated *at creation* so you cannot change this
	// definition after you have created the resource.
	LocalSecondaryIndexes []TableLocalSecondaryIndex `pulumi:"localSecondaryIndexes"`
	// The name of the index
	Name *string `pulumi:"name"`
	// Point-in-time recovery options.
	PointInTimeRecovery *TablePointInTimeRecovery `pulumi:"pointInTimeRecovery"`
	// The name of the range key; must be defined
	RangeKey *string `pulumi:"rangeKey"`
	// The number of read units for this index. Must be set if billingMode is set to PROVISIONED.
	ReadCapacity *int `pulumi:"readCapacity"`
	// Configuration block(s) with [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) replication configurations. Detailed below.
	Replicas []TableReplica `pulumi:"replicas"`
	// Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS owned Customer Master Key if this argument isn't specified.
	ServerSideEncryption *TableServerSideEncryption `pulumi:"serverSideEncryption"`
	// The ARN of the Table Stream. Only available when `streamEnabled = true`
	StreamArn *string `pulumi:"streamArn"`
	// Indicates whether Streams are to be enabled (true) or disabled (false).
	StreamEnabled *bool `pulumi:"streamEnabled"`
	// A timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not
	// a unique identifier for the stream on its own. However, the combination of AWS customer ID,
	// table name and this field is guaranteed to be unique.
	// It can be used for creating CloudWatch Alarms. Only available when `streamEnabled = true`
	StreamLabel *string `pulumi:"streamLabel"`
	// When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
	StreamViewType *string `pulumi:"streamViewType"`
	// A map of tags to populate on the created table. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Defines ttl, has two properties, and can only be specified once:
	Ttl *TableTtl `pulumi:"ttl"`
	// The number of write units for this index. Must be set if billingMode is set to PROVISIONED.
	WriteCapacity *int `pulumi:"writeCapacity"`
}

type TableState struct {
	// The arn of the table
	Arn pulumi.StringPtrInput
	// List of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes. Each attribute has two properties:
	Attributes TableAttributeArrayInput
	// Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
	BillingMode pulumi.StringPtrInput
	// Describe a GSI for the table;
	// subject to the normal limits on the number of GSIs, projected
	// attributes, etc.
	GlobalSecondaryIndexes TableGlobalSecondaryIndexArrayInput
	// The name of the hash key in the index; must be
	// defined as an attribute in the resource.
	HashKey pulumi.StringPtrInput
	// Describe an LSI on the table;
	// these can only be allocated *at creation* so you cannot change this
	// definition after you have created the resource.
	LocalSecondaryIndexes TableLocalSecondaryIndexArrayInput
	// The name of the index
	Name pulumi.StringPtrInput
	// Point-in-time recovery options.
	PointInTimeRecovery TablePointInTimeRecoveryPtrInput
	// The name of the range key; must be defined
	RangeKey pulumi.StringPtrInput
	// The number of read units for this index. Must be set if billingMode is set to PROVISIONED.
	ReadCapacity pulumi.IntPtrInput
	// Configuration block(s) with [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) replication configurations. Detailed below.
	Replicas TableReplicaArrayInput
	// Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS owned Customer Master Key if this argument isn't specified.
	ServerSideEncryption TableServerSideEncryptionPtrInput
	// The ARN of the Table Stream. Only available when `streamEnabled = true`
	StreamArn pulumi.StringPtrInput
	// Indicates whether Streams are to be enabled (true) or disabled (false).
	StreamEnabled pulumi.BoolPtrInput
	// A timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not
	// a unique identifier for the stream on its own. However, the combination of AWS customer ID,
	// table name and this field is guaranteed to be unique.
	// It can be used for creating CloudWatch Alarms. Only available when `streamEnabled = true`
	StreamLabel pulumi.StringPtrInput
	// When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
	StreamViewType pulumi.StringPtrInput
	// A map of tags to populate on the created table. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Defines ttl, has two properties, and can only be specified once:
	Ttl TableTtlPtrInput
	// The number of write units for this index. Must be set if billingMode is set to PROVISIONED.
	WriteCapacity pulumi.IntPtrInput
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	// List of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes. Each attribute has two properties:
	Attributes []TableAttribute `pulumi:"attributes"`
	// Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
	BillingMode *string `pulumi:"billingMode"`
	// Describe a GSI for the table;
	// subject to the normal limits on the number of GSIs, projected
	// attributes, etc.
	GlobalSecondaryIndexes []TableGlobalSecondaryIndex `pulumi:"globalSecondaryIndexes"`
	// The name of the hash key in the index; must be
	// defined as an attribute in the resource.
	HashKey string `pulumi:"hashKey"`
	// Describe an LSI on the table;
	// these can only be allocated *at creation* so you cannot change this
	// definition after you have created the resource.
	LocalSecondaryIndexes []TableLocalSecondaryIndex `pulumi:"localSecondaryIndexes"`
	// The name of the index
	Name *string `pulumi:"name"`
	// Point-in-time recovery options.
	PointInTimeRecovery *TablePointInTimeRecovery `pulumi:"pointInTimeRecovery"`
	// The name of the range key; must be defined
	RangeKey *string `pulumi:"rangeKey"`
	// The number of read units for this index. Must be set if billingMode is set to PROVISIONED.
	ReadCapacity *int `pulumi:"readCapacity"`
	// Configuration block(s) with [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) replication configurations. Detailed below.
	Replicas []TableReplica `pulumi:"replicas"`
	// Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS owned Customer Master Key if this argument isn't specified.
	ServerSideEncryption *TableServerSideEncryption `pulumi:"serverSideEncryption"`
	// Indicates whether Streams are to be enabled (true) or disabled (false).
	StreamEnabled *bool `pulumi:"streamEnabled"`
	// When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
	StreamViewType *string `pulumi:"streamViewType"`
	// A map of tags to populate on the created table. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Defines ttl, has two properties, and can only be specified once:
	Ttl *TableTtl `pulumi:"ttl"`
	// The number of write units for this index. Must be set if billingMode is set to PROVISIONED.
	WriteCapacity *int `pulumi:"writeCapacity"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	// List of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes. Each attribute has two properties:
	Attributes TableAttributeArrayInput
	// Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
	BillingMode pulumi.StringPtrInput
	// Describe a GSI for the table;
	// subject to the normal limits on the number of GSIs, projected
	// attributes, etc.
	GlobalSecondaryIndexes TableGlobalSecondaryIndexArrayInput
	// The name of the hash key in the index; must be
	// defined as an attribute in the resource.
	HashKey pulumi.StringInput
	// Describe an LSI on the table;
	// these can only be allocated *at creation* so you cannot change this
	// definition after you have created the resource.
	LocalSecondaryIndexes TableLocalSecondaryIndexArrayInput
	// The name of the index
	Name pulumi.StringPtrInput
	// Point-in-time recovery options.
	PointInTimeRecovery TablePointInTimeRecoveryPtrInput
	// The name of the range key; must be defined
	RangeKey pulumi.StringPtrInput
	// The number of read units for this index. Must be set if billingMode is set to PROVISIONED.
	ReadCapacity pulumi.IntPtrInput
	// Configuration block(s) with [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) replication configurations. Detailed below.
	Replicas TableReplicaArrayInput
	// Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS owned Customer Master Key if this argument isn't specified.
	ServerSideEncryption TableServerSideEncryptionPtrInput
	// Indicates whether Streams are to be enabled (true) or disabled (false).
	StreamEnabled pulumi.BoolPtrInput
	// When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
	StreamViewType pulumi.StringPtrInput
	// A map of tags to populate on the created table. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Defines ttl, has two properties, and can only be specified once:
	Ttl TableTtlPtrInput
	// The number of write units for this index. Must be set if billingMode is set to PROVISIONED.
	WriteCapacity pulumi.IntPtrInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}

type TableInput interface {
	pulumi.Input

	ToTableOutput() TableOutput
	ToTableOutputWithContext(ctx context.Context) TableOutput
}

func (*Table) ElementType() reflect.Type {
	return reflect.TypeOf((*Table)(nil))
}

func (i *Table) ToTableOutput() TableOutput {
	return i.ToTableOutputWithContext(context.Background())
}

func (i *Table) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableOutput)
}

func (i *Table) ToTablePtrOutput() TablePtrOutput {
	return i.ToTablePtrOutputWithContext(context.Background())
}

func (i *Table) ToTablePtrOutputWithContext(ctx context.Context) TablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TablePtrOutput)
}

type TablePtrInput interface {
	pulumi.Input

	ToTablePtrOutput() TablePtrOutput
	ToTablePtrOutputWithContext(ctx context.Context) TablePtrOutput
}

type tablePtrType TableArgs

func (*tablePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil))
}

func (i *tablePtrType) ToTablePtrOutput() TablePtrOutput {
	return i.ToTablePtrOutputWithContext(context.Background())
}

func (i *tablePtrType) ToTablePtrOutputWithContext(ctx context.Context) TablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TablePtrOutput)
}

// TableArrayInput is an input type that accepts TableArray and TableArrayOutput values.
// You can construct a concrete instance of `TableArrayInput` via:
//
//          TableArray{ TableArgs{...} }
type TableArrayInput interface {
	pulumi.Input

	ToTableArrayOutput() TableArrayOutput
	ToTableArrayOutputWithContext(context.Context) TableArrayOutput
}

type TableArray []TableInput

func (TableArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Table)(nil))
}

func (i TableArray) ToTableArrayOutput() TableArrayOutput {
	return i.ToTableArrayOutputWithContext(context.Background())
}

func (i TableArray) ToTableArrayOutputWithContext(ctx context.Context) TableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableArrayOutput)
}

// TableMapInput is an input type that accepts TableMap and TableMapOutput values.
// You can construct a concrete instance of `TableMapInput` via:
//
//          TableMap{ "key": TableArgs{...} }
type TableMapInput interface {
	pulumi.Input

	ToTableMapOutput() TableMapOutput
	ToTableMapOutputWithContext(context.Context) TableMapOutput
}

type TableMap map[string]TableInput

func (TableMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Table)(nil))
}

func (i TableMap) ToTableMapOutput() TableMapOutput {
	return i.ToTableMapOutputWithContext(context.Background())
}

func (i TableMap) ToTableMapOutputWithContext(ctx context.Context) TableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableMapOutput)
}

type TableOutput struct {
	*pulumi.OutputState
}

func (TableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Table)(nil))
}

func (o TableOutput) ToTableOutput() TableOutput {
	return o
}

func (o TableOutput) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return o
}

func (o TableOutput) ToTablePtrOutput() TablePtrOutput {
	return o.ToTablePtrOutputWithContext(context.Background())
}

func (o TableOutput) ToTablePtrOutputWithContext(ctx context.Context) TablePtrOutput {
	return o.ApplyT(func(v Table) *Table {
		return &v
	}).(TablePtrOutput)
}

type TablePtrOutput struct {
	*pulumi.OutputState
}

func (TablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil))
}

func (o TablePtrOutput) ToTablePtrOutput() TablePtrOutput {
	return o
}

func (o TablePtrOutput) ToTablePtrOutputWithContext(ctx context.Context) TablePtrOutput {
	return o
}

type TableArrayOutput struct{ *pulumi.OutputState }

func (TableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Table)(nil))
}

func (o TableArrayOutput) ToTableArrayOutput() TableArrayOutput {
	return o
}

func (o TableArrayOutput) ToTableArrayOutputWithContext(ctx context.Context) TableArrayOutput {
	return o
}

func (o TableArrayOutput) Index(i pulumi.IntInput) TableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Table {
		return vs[0].([]Table)[vs[1].(int)]
	}).(TableOutput)
}

type TableMapOutput struct{ *pulumi.OutputState }

func (TableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Table)(nil))
}

func (o TableMapOutput) ToTableMapOutput() TableMapOutput {
	return o
}

func (o TableMapOutput) ToTableMapOutputWithContext(ctx context.Context) TableMapOutput {
	return o
}

func (o TableMapOutput) MapIndex(k pulumi.StringInput) TableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Table {
		return vs[0].(map[string]Table)[vs[1].(string)]
	}).(TableOutput)
}

func init() {
	pulumi.RegisterOutputType(TableOutput{})
	pulumi.RegisterOutputType(TablePtrOutput{})
	pulumi.RegisterOutputType(TableArrayOutput{})
	pulumi.RegisterOutputType(TableMapOutput{})
}
