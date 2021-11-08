// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Glue ML Transform resource.
//
// ## Example Usage
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
// 		testCatalogDatabase, err := glue.NewCatalogDatabase(ctx, "testCatalogDatabase", &glue.CatalogDatabaseArgs{
// 			Name: pulumi.String("example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testCatalogTable, err := glue.NewCatalogTable(ctx, "testCatalogTable", &glue.CatalogTableArgs{
// 			Name:             pulumi.String("example"),
// 			DatabaseName:     testCatalogDatabase.Name,
// 			Owner:            pulumi.String("my_owner"),
// 			Retention:        pulumi.Int(1),
// 			TableType:        pulumi.String("VIRTUAL_VIEW"),
// 			ViewExpandedText: pulumi.String("view_expanded_text_1"),
// 			ViewOriginalText: pulumi.String("view_original_text_1"),
// 			StorageDescriptor: &glue.CatalogTableStorageDescriptorArgs{
// 				BucketColumns: pulumi.StringArray{
// 					pulumi.String("bucket_column_1"),
// 				},
// 				Compressed:             pulumi.Bool(false),
// 				InputFormat:            pulumi.String("SequenceFileInputFormat"),
// 				Location:               pulumi.String("my_location"),
// 				NumberOfBuckets:        pulumi.Int(1),
// 				OutputFormat:           pulumi.String("SequenceFileInputFormat"),
// 				StoredAsSubDirectories: pulumi.Bool(false),
// 				Parameters: pulumi.StringMap{
// 					"param1": pulumi.String("param1_val"),
// 				},
// 				Columns: glue.CatalogTableStorageDescriptorColumnArray{
// 					&glue.CatalogTableStorageDescriptorColumnArgs{
// 						Name:    pulumi.String("my_column_1"),
// 						Type:    pulumi.String("int"),
// 						Comment: pulumi.String("my_column1_comment"),
// 					},
// 					&glue.CatalogTableStorageDescriptorColumnArgs{
// 						Name:    pulumi.String("my_column_2"),
// 						Type:    pulumi.String("string"),
// 						Comment: pulumi.String("my_column2_comment"),
// 					},
// 				},
// 				SerDeInfo: &glue.CatalogTableStorageDescriptorSerDeInfoArgs{
// 					Name: pulumi.String("ser_de_name"),
// 					Parameters: pulumi.StringMap{
// 						"param1": pulumi.String("param_val_1"),
// 					},
// 					SerializationLibrary: pulumi.String("org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe"),
// 				},
// 				SortColumns: glue.CatalogTableStorageDescriptorSortColumnArray{
// 					&glue.CatalogTableStorageDescriptorSortColumnArgs{
// 						Column:    pulumi.String("my_column_1"),
// 						SortOrder: pulumi.Int(1),
// 					},
// 				},
// 				SkewedInfo: &glue.CatalogTableStorageDescriptorSkewedInfoArgs{
// 					SkewedColumnNames: pulumi.StringArray{
// 						pulumi.String("my_column_1"),
// 					},
// 					SkewedColumnValueLocationMaps: pulumi.StringMap{
// 						"my_column_1": pulumi.String("my_column_1_val_loc_map"),
// 					},
// 					SkewedColumnValues: pulumi.StringArray{
// 						pulumi.String("skewed_val_1"),
// 					},
// 				},
// 			},
// 			PartitionKeys: glue.CatalogTablePartitionKeyArray{
// 				&glue.CatalogTablePartitionKeyArgs{
// 					Name:    pulumi.String("my_column_1"),
// 					Type:    pulumi.String("int"),
// 					Comment: pulumi.String("my_column_1_comment"),
// 				},
// 				&glue.CatalogTablePartitionKeyArgs{
// 					Name:    pulumi.String("my_column_2"),
// 					Type:    pulumi.String("string"),
// 					Comment: pulumi.String("my_column_2_comment"),
// 				},
// 			},
// 			Parameters: pulumi.StringMap{
// 				"param1": pulumi.String("param1_val"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = glue.NewMLTransform(ctx, "testMLTransform", &glue.MLTransformArgs{
// 			RoleArn: pulumi.Any(aws_iam_role.Test.Arn),
// 			InputRecordTables: glue.MLTransformInputRecordTableArray{
// 				&glue.MLTransformInputRecordTableArgs{
// 					DatabaseName: testCatalogTable.DatabaseName,
// 					TableName:    testCatalogTable.Name,
// 				},
// 			},
// 			Parameters: &glue.MLTransformParametersArgs{
// 				TransformType: pulumi.String("FIND_MATCHES"),
// 				FindMatchesParameters: &glue.MLTransformParametersFindMatchesParametersArgs{
// 					PrimaryKeyColumnName: pulumi.String("my_column_1"),
// 				},
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			aws_iam_role_policy_attachment.Test,
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
// Glue ML Transforms can be imported using `id`, e.g.,
//
// ```sh
//  $ pulumi import aws:glue/mLTransform:MLTransform example tfm-c2cafbe83b1c575f49eaca9939220e2fcd58e2d5
// ```
type MLTransform struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of Glue ML Transform.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the ML Transform.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
	GlueVersion pulumi.StringOutput `pulumi:"glueVersion"`
	// A list of AWS Glue table definitions used by the transform. see Input Record Tables.
	InputRecordTables MLTransformInputRecordTableArrayOutput `pulumi:"inputRecordTables"`
	// The number of labels available for this transform.
	LabelCount pulumi.IntOutput `pulumi:"labelCount"`
	// The number of AWS Glue data processing units (DPUs) that are allocated to task runs for this transform. You can allocate from `2` to `100` DPUs; the default is `10`. `maxCapacity` is a mutually exclusive option with `numberOfWorkers` and `workerType`.
	MaxCapacity pulumi.Float64Output `pulumi:"maxCapacity"`
	// The maximum number of times to retry this ML Transform if it fails.
	MaxRetries pulumi.IntPtrOutput `pulumi:"maxRetries"`
	// The name you assign to this ML Transform. It must be unique in your account.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of workers of a defined `workerType` that are allocated when an ML Transform runs. Required with `workerType`.
	NumberOfWorkers pulumi.IntPtrOutput `pulumi:"numberOfWorkers"`
	// The algorithmic parameters that are specific to the transform type used. Conditionally dependent on the transform type. see Parameters.
	Parameters MLTransformParametersOutput `pulumi:"parameters"`
	// The ARN of the IAM role associated with this ML Transform.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// The object that represents the schema that this transform accepts. see Schema.
	Schemas MLTransformSchemaArrayOutput `pulumi:"schemas"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The ML Transform timeout in minutes. The default is 2880 minutes (48 hours).
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// The type of predefined worker that is allocated when an ML Transform runs. Accepts a value of `Standard`, `G.1X`, or `G.2X`. Required with `numberOfWorkers`.
	WorkerType pulumi.StringPtrOutput `pulumi:"workerType"`
}

// NewMLTransform registers a new resource with the given unique name, arguments, and options.
func NewMLTransform(ctx *pulumi.Context,
	name string, args *MLTransformArgs, opts ...pulumi.ResourceOption) (*MLTransform, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InputRecordTables == nil {
		return nil, errors.New("invalid value for required argument 'InputRecordTables'")
	}
	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource MLTransform
	err := ctx.RegisterResource("aws:glue/mLTransform:MLTransform", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMLTransform gets an existing MLTransform resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMLTransform(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MLTransformState, opts ...pulumi.ResourceOption) (*MLTransform, error) {
	var resource MLTransform
	err := ctx.ReadResource("aws:glue/mLTransform:MLTransform", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MLTransform resources.
type mltransformState struct {
	// Amazon Resource Name (ARN) of Glue ML Transform.
	Arn *string `pulumi:"arn"`
	// Description of the ML Transform.
	Description *string `pulumi:"description"`
	// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
	GlueVersion *string `pulumi:"glueVersion"`
	// A list of AWS Glue table definitions used by the transform. see Input Record Tables.
	InputRecordTables []MLTransformInputRecordTable `pulumi:"inputRecordTables"`
	// The number of labels available for this transform.
	LabelCount *int `pulumi:"labelCount"`
	// The number of AWS Glue data processing units (DPUs) that are allocated to task runs for this transform. You can allocate from `2` to `100` DPUs; the default is `10`. `maxCapacity` is a mutually exclusive option with `numberOfWorkers` and `workerType`.
	MaxCapacity *float64 `pulumi:"maxCapacity"`
	// The maximum number of times to retry this ML Transform if it fails.
	MaxRetries *int `pulumi:"maxRetries"`
	// The name you assign to this ML Transform. It must be unique in your account.
	Name *string `pulumi:"name"`
	// The number of workers of a defined `workerType` that are allocated when an ML Transform runs. Required with `workerType`.
	NumberOfWorkers *int `pulumi:"numberOfWorkers"`
	// The algorithmic parameters that are specific to the transform type used. Conditionally dependent on the transform type. see Parameters.
	Parameters *MLTransformParameters `pulumi:"parameters"`
	// The ARN of the IAM role associated with this ML Transform.
	RoleArn *string `pulumi:"roleArn"`
	// The object that represents the schema that this transform accepts. see Schema.
	Schemas []MLTransformSchema `pulumi:"schemas"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The ML Transform timeout in minutes. The default is 2880 minutes (48 hours).
	Timeout *int `pulumi:"timeout"`
	// The type of predefined worker that is allocated when an ML Transform runs. Accepts a value of `Standard`, `G.1X`, or `G.2X`. Required with `numberOfWorkers`.
	WorkerType *string `pulumi:"workerType"`
}

type MLTransformState struct {
	// Amazon Resource Name (ARN) of Glue ML Transform.
	Arn pulumi.StringPtrInput
	// Description of the ML Transform.
	Description pulumi.StringPtrInput
	// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
	GlueVersion pulumi.StringPtrInput
	// A list of AWS Glue table definitions used by the transform. see Input Record Tables.
	InputRecordTables MLTransformInputRecordTableArrayInput
	// The number of labels available for this transform.
	LabelCount pulumi.IntPtrInput
	// The number of AWS Glue data processing units (DPUs) that are allocated to task runs for this transform. You can allocate from `2` to `100` DPUs; the default is `10`. `maxCapacity` is a mutually exclusive option with `numberOfWorkers` and `workerType`.
	MaxCapacity pulumi.Float64PtrInput
	// The maximum number of times to retry this ML Transform if it fails.
	MaxRetries pulumi.IntPtrInput
	// The name you assign to this ML Transform. It must be unique in your account.
	Name pulumi.StringPtrInput
	// The number of workers of a defined `workerType` that are allocated when an ML Transform runs. Required with `workerType`.
	NumberOfWorkers pulumi.IntPtrInput
	// The algorithmic parameters that are specific to the transform type used. Conditionally dependent on the transform type. see Parameters.
	Parameters MLTransformParametersPtrInput
	// The ARN of the IAM role associated with this ML Transform.
	RoleArn pulumi.StringPtrInput
	// The object that represents the schema that this transform accepts. see Schema.
	Schemas MLTransformSchemaArrayInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The ML Transform timeout in minutes. The default is 2880 minutes (48 hours).
	Timeout pulumi.IntPtrInput
	// The type of predefined worker that is allocated when an ML Transform runs. Accepts a value of `Standard`, `G.1X`, or `G.2X`. Required with `numberOfWorkers`.
	WorkerType pulumi.StringPtrInput
}

func (MLTransformState) ElementType() reflect.Type {
	return reflect.TypeOf((*mltransformState)(nil)).Elem()
}

type mltransformArgs struct {
	// Description of the ML Transform.
	Description *string `pulumi:"description"`
	// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
	GlueVersion *string `pulumi:"glueVersion"`
	// A list of AWS Glue table definitions used by the transform. see Input Record Tables.
	InputRecordTables []MLTransformInputRecordTable `pulumi:"inputRecordTables"`
	// The number of AWS Glue data processing units (DPUs) that are allocated to task runs for this transform. You can allocate from `2` to `100` DPUs; the default is `10`. `maxCapacity` is a mutually exclusive option with `numberOfWorkers` and `workerType`.
	MaxCapacity *float64 `pulumi:"maxCapacity"`
	// The maximum number of times to retry this ML Transform if it fails.
	MaxRetries *int `pulumi:"maxRetries"`
	// The name you assign to this ML Transform. It must be unique in your account.
	Name *string `pulumi:"name"`
	// The number of workers of a defined `workerType` that are allocated when an ML Transform runs. Required with `workerType`.
	NumberOfWorkers *int `pulumi:"numberOfWorkers"`
	// The algorithmic parameters that are specific to the transform type used. Conditionally dependent on the transform type. see Parameters.
	Parameters MLTransformParameters `pulumi:"parameters"`
	// The ARN of the IAM role associated with this ML Transform.
	RoleArn string `pulumi:"roleArn"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The ML Transform timeout in minutes. The default is 2880 minutes (48 hours).
	Timeout *int `pulumi:"timeout"`
	// The type of predefined worker that is allocated when an ML Transform runs. Accepts a value of `Standard`, `G.1X`, or `G.2X`. Required with `numberOfWorkers`.
	WorkerType *string `pulumi:"workerType"`
}

// The set of arguments for constructing a MLTransform resource.
type MLTransformArgs struct {
	// Description of the ML Transform.
	Description pulumi.StringPtrInput
	// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
	GlueVersion pulumi.StringPtrInput
	// A list of AWS Glue table definitions used by the transform. see Input Record Tables.
	InputRecordTables MLTransformInputRecordTableArrayInput
	// The number of AWS Glue data processing units (DPUs) that are allocated to task runs for this transform. You can allocate from `2` to `100` DPUs; the default is `10`. `maxCapacity` is a mutually exclusive option with `numberOfWorkers` and `workerType`.
	MaxCapacity pulumi.Float64PtrInput
	// The maximum number of times to retry this ML Transform if it fails.
	MaxRetries pulumi.IntPtrInput
	// The name you assign to this ML Transform. It must be unique in your account.
	Name pulumi.StringPtrInput
	// The number of workers of a defined `workerType` that are allocated when an ML Transform runs. Required with `workerType`.
	NumberOfWorkers pulumi.IntPtrInput
	// The algorithmic parameters that are specific to the transform type used. Conditionally dependent on the transform type. see Parameters.
	Parameters MLTransformParametersInput
	// The ARN of the IAM role associated with this ML Transform.
	RoleArn pulumi.StringInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The ML Transform timeout in minutes. The default is 2880 minutes (48 hours).
	Timeout pulumi.IntPtrInput
	// The type of predefined worker that is allocated when an ML Transform runs. Accepts a value of `Standard`, `G.1X`, or `G.2X`. Required with `numberOfWorkers`.
	WorkerType pulumi.StringPtrInput
}

func (MLTransformArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mltransformArgs)(nil)).Elem()
}

type MLTransformInput interface {
	pulumi.Input

	ToMLTransformOutput() MLTransformOutput
	ToMLTransformOutputWithContext(ctx context.Context) MLTransformOutput
}

func (*MLTransform) ElementType() reflect.Type {
	return reflect.TypeOf((*MLTransform)(nil))
}

func (i *MLTransform) ToMLTransformOutput() MLTransformOutput {
	return i.ToMLTransformOutputWithContext(context.Background())
}

func (i *MLTransform) ToMLTransformOutputWithContext(ctx context.Context) MLTransformOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MLTransformOutput)
}

func (i *MLTransform) ToMLTransformPtrOutput() MLTransformPtrOutput {
	return i.ToMLTransformPtrOutputWithContext(context.Background())
}

func (i *MLTransform) ToMLTransformPtrOutputWithContext(ctx context.Context) MLTransformPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MLTransformPtrOutput)
}

type MLTransformPtrInput interface {
	pulumi.Input

	ToMLTransformPtrOutput() MLTransformPtrOutput
	ToMLTransformPtrOutputWithContext(ctx context.Context) MLTransformPtrOutput
}

type mltransformPtrType MLTransformArgs

func (*mltransformPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MLTransform)(nil))
}

func (i *mltransformPtrType) ToMLTransformPtrOutput() MLTransformPtrOutput {
	return i.ToMLTransformPtrOutputWithContext(context.Background())
}

func (i *mltransformPtrType) ToMLTransformPtrOutputWithContext(ctx context.Context) MLTransformPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MLTransformPtrOutput)
}

// MLTransformArrayInput is an input type that accepts MLTransformArray and MLTransformArrayOutput values.
// You can construct a concrete instance of `MLTransformArrayInput` via:
//
//          MLTransformArray{ MLTransformArgs{...} }
type MLTransformArrayInput interface {
	pulumi.Input

	ToMLTransformArrayOutput() MLTransformArrayOutput
	ToMLTransformArrayOutputWithContext(context.Context) MLTransformArrayOutput
}

type MLTransformArray []MLTransformInput

func (MLTransformArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MLTransform)(nil)).Elem()
}

func (i MLTransformArray) ToMLTransformArrayOutput() MLTransformArrayOutput {
	return i.ToMLTransformArrayOutputWithContext(context.Background())
}

func (i MLTransformArray) ToMLTransformArrayOutputWithContext(ctx context.Context) MLTransformArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MLTransformArrayOutput)
}

// MLTransformMapInput is an input type that accepts MLTransformMap and MLTransformMapOutput values.
// You can construct a concrete instance of `MLTransformMapInput` via:
//
//          MLTransformMap{ "key": MLTransformArgs{...} }
type MLTransformMapInput interface {
	pulumi.Input

	ToMLTransformMapOutput() MLTransformMapOutput
	ToMLTransformMapOutputWithContext(context.Context) MLTransformMapOutput
}

type MLTransformMap map[string]MLTransformInput

func (MLTransformMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MLTransform)(nil)).Elem()
}

func (i MLTransformMap) ToMLTransformMapOutput() MLTransformMapOutput {
	return i.ToMLTransformMapOutputWithContext(context.Background())
}

func (i MLTransformMap) ToMLTransformMapOutputWithContext(ctx context.Context) MLTransformMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MLTransformMapOutput)
}

type MLTransformOutput struct{ *pulumi.OutputState }

func (MLTransformOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MLTransform)(nil))
}

func (o MLTransformOutput) ToMLTransformOutput() MLTransformOutput {
	return o
}

func (o MLTransformOutput) ToMLTransformOutputWithContext(ctx context.Context) MLTransformOutput {
	return o
}

func (o MLTransformOutput) ToMLTransformPtrOutput() MLTransformPtrOutput {
	return o.ToMLTransformPtrOutputWithContext(context.Background())
}

func (o MLTransformOutput) ToMLTransformPtrOutputWithContext(ctx context.Context) MLTransformPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MLTransform) *MLTransform {
		return &v
	}).(MLTransformPtrOutput)
}

type MLTransformPtrOutput struct{ *pulumi.OutputState }

func (MLTransformPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MLTransform)(nil))
}

func (o MLTransformPtrOutput) ToMLTransformPtrOutput() MLTransformPtrOutput {
	return o
}

func (o MLTransformPtrOutput) ToMLTransformPtrOutputWithContext(ctx context.Context) MLTransformPtrOutput {
	return o
}

func (o MLTransformPtrOutput) Elem() MLTransformOutput {
	return o.ApplyT(func(v *MLTransform) MLTransform {
		if v != nil {
			return *v
		}
		var ret MLTransform
		return ret
	}).(MLTransformOutput)
}

type MLTransformArrayOutput struct{ *pulumi.OutputState }

func (MLTransformArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MLTransform)(nil))
}

func (o MLTransformArrayOutput) ToMLTransformArrayOutput() MLTransformArrayOutput {
	return o
}

func (o MLTransformArrayOutput) ToMLTransformArrayOutputWithContext(ctx context.Context) MLTransformArrayOutput {
	return o
}

func (o MLTransformArrayOutput) Index(i pulumi.IntInput) MLTransformOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MLTransform {
		return vs[0].([]MLTransform)[vs[1].(int)]
	}).(MLTransformOutput)
}

type MLTransformMapOutput struct{ *pulumi.OutputState }

func (MLTransformMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MLTransform)(nil))
}

func (o MLTransformMapOutput) ToMLTransformMapOutput() MLTransformMapOutput {
	return o
}

func (o MLTransformMapOutput) ToMLTransformMapOutputWithContext(ctx context.Context) MLTransformMapOutput {
	return o
}

func (o MLTransformMapOutput) MapIndex(k pulumi.StringInput) MLTransformOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MLTransform {
		return vs[0].(map[string]MLTransform)[vs[1].(string)]
	}).(MLTransformOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MLTransformInput)(nil)).Elem(), &MLTransform{})
	pulumi.RegisterInputType(reflect.TypeOf((*MLTransformPtrInput)(nil)).Elem(), &MLTransform{})
	pulumi.RegisterInputType(reflect.TypeOf((*MLTransformArrayInput)(nil)).Elem(), MLTransformArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MLTransformMapInput)(nil)).Elem(), MLTransformMap{})
	pulumi.RegisterOutputType(MLTransformOutput{})
	pulumi.RegisterOutputType(MLTransformPtrOutput{})
	pulumi.RegisterOutputType(MLTransformArrayOutput{})
	pulumi.RegisterOutputType(MLTransformMapOutput{})
}
