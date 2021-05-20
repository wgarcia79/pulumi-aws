// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Glue Schema resource.
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
// 		_, err := glue.NewSchema(ctx, "example", &glue.SchemaArgs{
// 			SchemaName:       pulumi.String("example"),
// 			RegistryArn:      pulumi.Any(aws_glue_registry.Test.Arn),
// 			DataFormat:       pulumi.String("AVRO"),
// 			Compatibility:    pulumi.String("NONE"),
// 			SchemaDefinition: pulumi.String("{\"type\": \"record\", \"name\": \"r1\", \"fields\": [ {\"name\": \"f1\", \"type\": \"int\"}, {\"name\": \"f2\", \"type\": \"string\"} ]}"),
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
// Glue Registries can be imported using `arn`, e.g.
//
// ```sh
//  $ pulumi import aws:glue/schema:Schema example arn:aws:glue:us-west-2:123456789012:schema/example/example
// ```
type Schema struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the schema.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
	Compatibility pulumi.StringOutput `pulumi:"compatibility"`
	// The data format of the schema definition. Currently only `AVRO` is supported.
	DataFormat pulumi.StringOutput `pulumi:"dataFormat"`
	// A description of the schema.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The latest version of the schema associated with the returned schema definition.
	LatestSchemaVersion pulumi.IntOutput `pulumi:"latestSchemaVersion"`
	// The next version of the schema associated with the returned schema definition.
	NextSchemaVersion pulumi.IntOutput `pulumi:"nextSchemaVersion"`
	// The ARN of the Glue Registry to create the schema in.
	RegistryArn pulumi.StringOutput `pulumi:"registryArn"`
	// The name of the Glue Registry.
	RegistryName pulumi.StringOutput `pulumi:"registryName"`
	// The version number of the checkpoint (the last time the compatibility mode was changed).
	SchemaCheckpoint pulumi.IntOutput `pulumi:"schemaCheckpoint"`
	// The schema definition using the `dataFormat` setting for `schemaName`.
	SchemaDefinition pulumi.StringOutput `pulumi:"schemaDefinition"`
	// The Name of the schema.
	SchemaName pulumi.StringOutput `pulumi:"schemaName"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    pulumi.StringMapOutput `pulumi:"tags"`
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewSchema registers a new resource with the given unique name, arguments, and options.
func NewSchema(ctx *pulumi.Context,
	name string, args *SchemaArgs, opts ...pulumi.ResourceOption) (*Schema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Compatibility == nil {
		return nil, errors.New("invalid value for required argument 'Compatibility'")
	}
	if args.DataFormat == nil {
		return nil, errors.New("invalid value for required argument 'DataFormat'")
	}
	if args.SchemaDefinition == nil {
		return nil, errors.New("invalid value for required argument 'SchemaDefinition'")
	}
	if args.SchemaName == nil {
		return nil, errors.New("invalid value for required argument 'SchemaName'")
	}
	var resource Schema
	err := ctx.RegisterResource("aws:glue/schema:Schema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchema gets an existing Schema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemaState, opts ...pulumi.ResourceOption) (*Schema, error) {
	var resource Schema
	err := ctx.ReadResource("aws:glue/schema:Schema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Schema resources.
type schemaState struct {
	// Amazon Resource Name (ARN) of the schema.
	Arn *string `pulumi:"arn"`
	// The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
	Compatibility *string `pulumi:"compatibility"`
	// The data format of the schema definition. Currently only `AVRO` is supported.
	DataFormat *string `pulumi:"dataFormat"`
	// A description of the schema.
	Description *string `pulumi:"description"`
	// The latest version of the schema associated with the returned schema definition.
	LatestSchemaVersion *int `pulumi:"latestSchemaVersion"`
	// The next version of the schema associated with the returned schema definition.
	NextSchemaVersion *int `pulumi:"nextSchemaVersion"`
	// The ARN of the Glue Registry to create the schema in.
	RegistryArn *string `pulumi:"registryArn"`
	// The name of the Glue Registry.
	RegistryName *string `pulumi:"registryName"`
	// The version number of the checkpoint (the last time the compatibility mode was changed).
	SchemaCheckpoint *int `pulumi:"schemaCheckpoint"`
	// The schema definition using the `dataFormat` setting for `schemaName`.
	SchemaDefinition *string `pulumi:"schemaDefinition"`
	// The Name of the schema.
	SchemaName *string `pulumi:"schemaName"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    map[string]string `pulumi:"tags"`
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type SchemaState struct {
	// Amazon Resource Name (ARN) of the schema.
	Arn pulumi.StringPtrInput
	// The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
	Compatibility pulumi.StringPtrInput
	// The data format of the schema definition. Currently only `AVRO` is supported.
	DataFormat pulumi.StringPtrInput
	// A description of the schema.
	Description pulumi.StringPtrInput
	// The latest version of the schema associated with the returned schema definition.
	LatestSchemaVersion pulumi.IntPtrInput
	// The next version of the schema associated with the returned schema definition.
	NextSchemaVersion pulumi.IntPtrInput
	// The ARN of the Glue Registry to create the schema in.
	RegistryArn pulumi.StringPtrInput
	// The name of the Glue Registry.
	RegistryName pulumi.StringPtrInput
	// The version number of the checkpoint (the last time the compatibility mode was changed).
	SchemaCheckpoint pulumi.IntPtrInput
	// The schema definition using the `dataFormat` setting for `schemaName`.
	SchemaDefinition pulumi.StringPtrInput
	// The Name of the schema.
	SchemaName pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    pulumi.StringMapInput
	TagsAll pulumi.StringMapInput
}

func (SchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaState)(nil)).Elem()
}

type schemaArgs struct {
	// The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
	Compatibility string `pulumi:"compatibility"`
	// The data format of the schema definition. Currently only `AVRO` is supported.
	DataFormat string `pulumi:"dataFormat"`
	// A description of the schema.
	Description *string `pulumi:"description"`
	// The ARN of the Glue Registry to create the schema in.
	RegistryArn *string `pulumi:"registryArn"`
	// The schema definition using the `dataFormat` setting for `schemaName`.
	SchemaDefinition string `pulumi:"schemaDefinition"`
	// The Name of the schema.
	SchemaName string `pulumi:"schemaName"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    map[string]string `pulumi:"tags"`
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a Schema resource.
type SchemaArgs struct {
	// The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
	Compatibility pulumi.StringInput
	// The data format of the schema definition. Currently only `AVRO` is supported.
	DataFormat pulumi.StringInput
	// A description of the schema.
	Description pulumi.StringPtrInput
	// The ARN of the Glue Registry to create the schema in.
	RegistryArn pulumi.StringPtrInput
	// The schema definition using the `dataFormat` setting for `schemaName`.
	SchemaDefinition pulumi.StringInput
	// The Name of the schema.
	SchemaName pulumi.StringInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    pulumi.StringMapInput
	TagsAll pulumi.StringMapInput
}

func (SchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaArgs)(nil)).Elem()
}

type SchemaInput interface {
	pulumi.Input

	ToSchemaOutput() SchemaOutput
	ToSchemaOutputWithContext(ctx context.Context) SchemaOutput
}

func (*Schema) ElementType() reflect.Type {
	return reflect.TypeOf((*Schema)(nil))
}

func (i *Schema) ToSchemaOutput() SchemaOutput {
	return i.ToSchemaOutputWithContext(context.Background())
}

func (i *Schema) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaOutput)
}

func (i *Schema) ToSchemaPtrOutput() SchemaPtrOutput {
	return i.ToSchemaPtrOutputWithContext(context.Background())
}

func (i *Schema) ToSchemaPtrOutputWithContext(ctx context.Context) SchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaPtrOutput)
}

type SchemaPtrInput interface {
	pulumi.Input

	ToSchemaPtrOutput() SchemaPtrOutput
	ToSchemaPtrOutputWithContext(ctx context.Context) SchemaPtrOutput
}

type schemaPtrType SchemaArgs

func (*schemaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil))
}

func (i *schemaPtrType) ToSchemaPtrOutput() SchemaPtrOutput {
	return i.ToSchemaPtrOutputWithContext(context.Background())
}

func (i *schemaPtrType) ToSchemaPtrOutputWithContext(ctx context.Context) SchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaPtrOutput)
}

// SchemaArrayInput is an input type that accepts SchemaArray and SchemaArrayOutput values.
// You can construct a concrete instance of `SchemaArrayInput` via:
//
//          SchemaArray{ SchemaArgs{...} }
type SchemaArrayInput interface {
	pulumi.Input

	ToSchemaArrayOutput() SchemaArrayOutput
	ToSchemaArrayOutputWithContext(context.Context) SchemaArrayOutput
}

type SchemaArray []SchemaInput

func (SchemaArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Schema)(nil))
}

func (i SchemaArray) ToSchemaArrayOutput() SchemaArrayOutput {
	return i.ToSchemaArrayOutputWithContext(context.Background())
}

func (i SchemaArray) ToSchemaArrayOutputWithContext(ctx context.Context) SchemaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaArrayOutput)
}

// SchemaMapInput is an input type that accepts SchemaMap and SchemaMapOutput values.
// You can construct a concrete instance of `SchemaMapInput` via:
//
//          SchemaMap{ "key": SchemaArgs{...} }
type SchemaMapInput interface {
	pulumi.Input

	ToSchemaMapOutput() SchemaMapOutput
	ToSchemaMapOutputWithContext(context.Context) SchemaMapOutput
}

type SchemaMap map[string]SchemaInput

func (SchemaMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Schema)(nil))
}

func (i SchemaMap) ToSchemaMapOutput() SchemaMapOutput {
	return i.ToSchemaMapOutputWithContext(context.Background())
}

func (i SchemaMap) ToSchemaMapOutputWithContext(ctx context.Context) SchemaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaMapOutput)
}

type SchemaOutput struct {
	*pulumi.OutputState
}

func (SchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Schema)(nil))
}

func (o SchemaOutput) ToSchemaOutput() SchemaOutput {
	return o
}

func (o SchemaOutput) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return o
}

func (o SchemaOutput) ToSchemaPtrOutput() SchemaPtrOutput {
	return o.ToSchemaPtrOutputWithContext(context.Background())
}

func (o SchemaOutput) ToSchemaPtrOutputWithContext(ctx context.Context) SchemaPtrOutput {
	return o.ApplyT(func(v Schema) *Schema {
		return &v
	}).(SchemaPtrOutput)
}

type SchemaPtrOutput struct {
	*pulumi.OutputState
}

func (SchemaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil))
}

func (o SchemaPtrOutput) ToSchemaPtrOutput() SchemaPtrOutput {
	return o
}

func (o SchemaPtrOutput) ToSchemaPtrOutputWithContext(ctx context.Context) SchemaPtrOutput {
	return o
}

type SchemaArrayOutput struct{ *pulumi.OutputState }

func (SchemaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Schema)(nil))
}

func (o SchemaArrayOutput) ToSchemaArrayOutput() SchemaArrayOutput {
	return o
}

func (o SchemaArrayOutput) ToSchemaArrayOutputWithContext(ctx context.Context) SchemaArrayOutput {
	return o
}

func (o SchemaArrayOutput) Index(i pulumi.IntInput) SchemaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Schema {
		return vs[0].([]Schema)[vs[1].(int)]
	}).(SchemaOutput)
}

type SchemaMapOutput struct{ *pulumi.OutputState }

func (SchemaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Schema)(nil))
}

func (o SchemaMapOutput) ToSchemaMapOutput() SchemaMapOutput {
	return o
}

func (o SchemaMapOutput) ToSchemaMapOutputWithContext(ctx context.Context) SchemaMapOutput {
	return o
}

func (o SchemaMapOutput) MapIndex(k pulumi.StringInput) SchemaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Schema {
		return vs[0].(map[string]Schema)[vs[1].(string)]
	}).(SchemaOutput)
}

func init() {
	pulumi.RegisterOutputType(SchemaOutput{})
	pulumi.RegisterOutputType(SchemaPtrOutput{})
	pulumi.RegisterOutputType(SchemaArrayOutput{})
	pulumi.RegisterOutputType(SchemaMapOutput{})
}
