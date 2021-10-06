// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing QuickSight Data Source
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/quicksight"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := quicksight.NewDataSource(ctx, "_default", &quicksight.DataSourceArgs{
// 			DataSourceId: pulumi.String("example-id"),
// 			Parameters: &quicksight.DataSourceParametersArgs{
// 				S3: &quicksight.DataSourceParametersS3Args{
// 					ManifestFileLocation: &quicksight.DataSourceParametersS3ManifestFileLocationArgs{
// 						Bucket: pulumi.String("my-bucket"),
// 						Key:    pulumi.String("path/to/manifest.json"),
// 					},
// 				},
// 			},
// 			Type: pulumi.String("S3"),
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
// A QuickSight data source can be imported using the AWS account ID, and data source ID name separated by a slash (`/`) e.g.
//
// ```sh
//  $ pulumi import aws:quicksight/dataSource:DataSource example 123456789123/my-data-source-id
// ```
type DataSource struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the data source
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID for the AWS account that the data source is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId pulumi.StringOutput `pulumi:"awsAccountId"`
	// The credentials Amazon QuickSight uses to connect to your underlying source. Currently, only credentials based on user name and password are supported. See Credentials below for more details.
	Credentials DataSourceCredentialsPtrOutput `pulumi:"credentials"`
	// An identifier for the data source.
	DataSourceId pulumi.StringOutput `pulumi:"dataSourceId"`
	// A name for the data source, maximum of 128 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parameters used to connect to this data source (exactly one).
	Parameters DataSourceParametersOutput `pulumi:"parameters"`
	// A set of resource permissions on the data source. Maximum of 64 items. See Permission below for more details.
	Permissions DataSourcePermissionArrayOutput `pulumi:"permissions"`
	// Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source. See SSL Properties below for more details.
	SslProperties DataSourceSslPropertiesPtrOutput `pulumi:"sslProperties"`
	// Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The type of the data source. See the [AWS Documentation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CreateDataSource.html#QS-CreateDataSource-request-Type) for the complete list of valid values.
	Type pulumi.StringOutput `pulumi:"type"`
	// Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source. See VPC Connection Properties below for more details.
	VpcConnectionProperties DataSourceVpcConnectionPropertiesPtrOutput `pulumi:"vpcConnectionProperties"`
}

// NewDataSource registers a new resource with the given unique name, arguments, and options.
func NewDataSource(ctx *pulumi.Context,
	name string, args *DataSourceArgs, opts ...pulumi.ResourceOption) (*DataSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataSourceId == nil {
		return nil, errors.New("invalid value for required argument 'DataSourceId'")
	}
	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource DataSource
	err := ctx.RegisterResource("aws:quicksight/dataSource:DataSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataSource gets an existing DataSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSourceState, opts ...pulumi.ResourceOption) (*DataSource, error) {
	var resource DataSource
	err := ctx.ReadResource("aws:quicksight/dataSource:DataSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataSource resources.
type dataSourceState struct {
	// Amazon Resource Name (ARN) of the data source
	Arn *string `pulumi:"arn"`
	// The ID for the AWS account that the data source is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// The credentials Amazon QuickSight uses to connect to your underlying source. Currently, only credentials based on user name and password are supported. See Credentials below for more details.
	Credentials *DataSourceCredentials `pulumi:"credentials"`
	// An identifier for the data source.
	DataSourceId *string `pulumi:"dataSourceId"`
	// A name for the data source, maximum of 128 characters.
	Name *string `pulumi:"name"`
	// The parameters used to connect to this data source (exactly one).
	Parameters *DataSourceParameters `pulumi:"parameters"`
	// A set of resource permissions on the data source. Maximum of 64 items. See Permission below for more details.
	Permissions []DataSourcePermission `pulumi:"permissions"`
	// Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source. See SSL Properties below for more details.
	SslProperties *DataSourceSslProperties `pulumi:"sslProperties"`
	// Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The type of the data source. See the [AWS Documentation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CreateDataSource.html#QS-CreateDataSource-request-Type) for the complete list of valid values.
	Type *string `pulumi:"type"`
	// Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source. See VPC Connection Properties below for more details.
	VpcConnectionProperties *DataSourceVpcConnectionProperties `pulumi:"vpcConnectionProperties"`
}

type DataSourceState struct {
	// Amazon Resource Name (ARN) of the data source
	Arn pulumi.StringPtrInput
	// The ID for the AWS account that the data source is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId pulumi.StringPtrInput
	// The credentials Amazon QuickSight uses to connect to your underlying source. Currently, only credentials based on user name and password are supported. See Credentials below for more details.
	Credentials DataSourceCredentialsPtrInput
	// An identifier for the data source.
	DataSourceId pulumi.StringPtrInput
	// A name for the data source, maximum of 128 characters.
	Name pulumi.StringPtrInput
	// The parameters used to connect to this data source (exactly one).
	Parameters DataSourceParametersPtrInput
	// A set of resource permissions on the data source. Maximum of 64 items. See Permission below for more details.
	Permissions DataSourcePermissionArrayInput
	// Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source. See SSL Properties below for more details.
	SslProperties DataSourceSslPropertiesPtrInput
	// Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapInput
	// The type of the data source. See the [AWS Documentation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CreateDataSource.html#QS-CreateDataSource-request-Type) for the complete list of valid values.
	Type pulumi.StringPtrInput
	// Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source. See VPC Connection Properties below for more details.
	VpcConnectionProperties DataSourceVpcConnectionPropertiesPtrInput
}

func (DataSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceState)(nil)).Elem()
}

type dataSourceArgs struct {
	// The ID for the AWS account that the data source is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// The credentials Amazon QuickSight uses to connect to your underlying source. Currently, only credentials based on user name and password are supported. See Credentials below for more details.
	Credentials *DataSourceCredentials `pulumi:"credentials"`
	// An identifier for the data source.
	DataSourceId string `pulumi:"dataSourceId"`
	// A name for the data source, maximum of 128 characters.
	Name *string `pulumi:"name"`
	// The parameters used to connect to this data source (exactly one).
	Parameters DataSourceParameters `pulumi:"parameters"`
	// A set of resource permissions on the data source. Maximum of 64 items. See Permission below for more details.
	Permissions []DataSourcePermission `pulumi:"permissions"`
	// Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source. See SSL Properties below for more details.
	SslProperties *DataSourceSslProperties `pulumi:"sslProperties"`
	// Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The type of the data source. See the [AWS Documentation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CreateDataSource.html#QS-CreateDataSource-request-Type) for the complete list of valid values.
	Type string `pulumi:"type"`
	// Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source. See VPC Connection Properties below for more details.
	VpcConnectionProperties *DataSourceVpcConnectionProperties `pulumi:"vpcConnectionProperties"`
}

// The set of arguments for constructing a DataSource resource.
type DataSourceArgs struct {
	// The ID for the AWS account that the data source is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId pulumi.StringPtrInput
	// The credentials Amazon QuickSight uses to connect to your underlying source. Currently, only credentials based on user name and password are supported. See Credentials below for more details.
	Credentials DataSourceCredentialsPtrInput
	// An identifier for the data source.
	DataSourceId pulumi.StringInput
	// A name for the data source, maximum of 128 characters.
	Name pulumi.StringPtrInput
	// The parameters used to connect to this data source (exactly one).
	Parameters DataSourceParametersInput
	// A set of resource permissions on the data source. Maximum of 64 items. See Permission below for more details.
	Permissions DataSourcePermissionArrayInput
	// Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source. See SSL Properties below for more details.
	SslProperties DataSourceSslPropertiesPtrInput
	// Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapInput
	// The type of the data source. See the [AWS Documentation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CreateDataSource.html#QS-CreateDataSource-request-Type) for the complete list of valid values.
	Type pulumi.StringInput
	// Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source. See VPC Connection Properties below for more details.
	VpcConnectionProperties DataSourceVpcConnectionPropertiesPtrInput
}

func (DataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceArgs)(nil)).Elem()
}

type DataSourceInput interface {
	pulumi.Input

	ToDataSourceOutput() DataSourceOutput
	ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput
}

func (*DataSource) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSource)(nil))
}

func (i *DataSource) ToDataSourceOutput() DataSourceOutput {
	return i.ToDataSourceOutputWithContext(context.Background())
}

func (i *DataSource) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceOutput)
}

func (i *DataSource) ToDataSourcePtrOutput() DataSourcePtrOutput {
	return i.ToDataSourcePtrOutputWithContext(context.Background())
}

func (i *DataSource) ToDataSourcePtrOutputWithContext(ctx context.Context) DataSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourcePtrOutput)
}

type DataSourcePtrInput interface {
	pulumi.Input

	ToDataSourcePtrOutput() DataSourcePtrOutput
	ToDataSourcePtrOutputWithContext(ctx context.Context) DataSourcePtrOutput
}

type dataSourcePtrType DataSourceArgs

func (*dataSourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSource)(nil))
}

func (i *dataSourcePtrType) ToDataSourcePtrOutput() DataSourcePtrOutput {
	return i.ToDataSourcePtrOutputWithContext(context.Background())
}

func (i *dataSourcePtrType) ToDataSourcePtrOutputWithContext(ctx context.Context) DataSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourcePtrOutput)
}

// DataSourceArrayInput is an input type that accepts DataSourceArray and DataSourceArrayOutput values.
// You can construct a concrete instance of `DataSourceArrayInput` via:
//
//          DataSourceArray{ DataSourceArgs{...} }
type DataSourceArrayInput interface {
	pulumi.Input

	ToDataSourceArrayOutput() DataSourceArrayOutput
	ToDataSourceArrayOutputWithContext(context.Context) DataSourceArrayOutput
}

type DataSourceArray []DataSourceInput

func (DataSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataSource)(nil)).Elem()
}

func (i DataSourceArray) ToDataSourceArrayOutput() DataSourceArrayOutput {
	return i.ToDataSourceArrayOutputWithContext(context.Background())
}

func (i DataSourceArray) ToDataSourceArrayOutputWithContext(ctx context.Context) DataSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceArrayOutput)
}

// DataSourceMapInput is an input type that accepts DataSourceMap and DataSourceMapOutput values.
// You can construct a concrete instance of `DataSourceMapInput` via:
//
//          DataSourceMap{ "key": DataSourceArgs{...} }
type DataSourceMapInput interface {
	pulumi.Input

	ToDataSourceMapOutput() DataSourceMapOutput
	ToDataSourceMapOutputWithContext(context.Context) DataSourceMapOutput
}

type DataSourceMap map[string]DataSourceInput

func (DataSourceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataSource)(nil)).Elem()
}

func (i DataSourceMap) ToDataSourceMapOutput() DataSourceMapOutput {
	return i.ToDataSourceMapOutputWithContext(context.Background())
}

func (i DataSourceMap) ToDataSourceMapOutputWithContext(ctx context.Context) DataSourceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceMapOutput)
}

type DataSourceOutput struct{ *pulumi.OutputState }

func (DataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSource)(nil))
}

func (o DataSourceOutput) ToDataSourceOutput() DataSourceOutput {
	return o
}

func (o DataSourceOutput) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return o
}

func (o DataSourceOutput) ToDataSourcePtrOutput() DataSourcePtrOutput {
	return o.ToDataSourcePtrOutputWithContext(context.Background())
}

func (o DataSourceOutput) ToDataSourcePtrOutputWithContext(ctx context.Context) DataSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataSource) *DataSource {
		return &v
	}).(DataSourcePtrOutput)
}

type DataSourcePtrOutput struct{ *pulumi.OutputState }

func (DataSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSource)(nil))
}

func (o DataSourcePtrOutput) ToDataSourcePtrOutput() DataSourcePtrOutput {
	return o
}

func (o DataSourcePtrOutput) ToDataSourcePtrOutputWithContext(ctx context.Context) DataSourcePtrOutput {
	return o
}

func (o DataSourcePtrOutput) Elem() DataSourceOutput {
	return o.ApplyT(func(v *DataSource) DataSource {
		if v != nil {
			return *v
		}
		var ret DataSource
		return ret
	}).(DataSourceOutput)
}

type DataSourceArrayOutput struct{ *pulumi.OutputState }

func (DataSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataSource)(nil))
}

func (o DataSourceArrayOutput) ToDataSourceArrayOutput() DataSourceArrayOutput {
	return o
}

func (o DataSourceArrayOutput) ToDataSourceArrayOutputWithContext(ctx context.Context) DataSourceArrayOutput {
	return o
}

func (o DataSourceArrayOutput) Index(i pulumi.IntInput) DataSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataSource {
		return vs[0].([]DataSource)[vs[1].(int)]
	}).(DataSourceOutput)
}

type DataSourceMapOutput struct{ *pulumi.OutputState }

func (DataSourceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DataSource)(nil))
}

func (o DataSourceMapOutput) ToDataSourceMapOutput() DataSourceMapOutput {
	return o
}

func (o DataSourceMapOutput) ToDataSourceMapOutputWithContext(ctx context.Context) DataSourceMapOutput {
	return o
}

func (o DataSourceMapOutput) MapIndex(k pulumi.StringInput) DataSourceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DataSource {
		return vs[0].(map[string]DataSource)[vs[1].(string)]
	}).(DataSourceOutput)
}

func init() {
	pulumi.RegisterOutputType(DataSourceOutput{})
	pulumi.RegisterOutputType(DataSourcePtrOutput{})
	pulumi.RegisterOutputType(DataSourceArrayOutput{})
	pulumi.RegisterOutputType(DataSourceMapOutput{})
}
