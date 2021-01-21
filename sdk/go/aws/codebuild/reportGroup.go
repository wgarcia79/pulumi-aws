// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codebuild

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CodeBuild Report Groups Resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/codebuild"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/kms"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
// 			Description:          pulumi.String("my test kms key"),
// 			DeletionWindowInDays: pulumi.Int(7),
// 			Policy:               pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Id\": \"kms-tf-1\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"Enable IAM User Permissions\",\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"AWS\": \"*\"\n", "      },\n", "      \"Action\": \"kms:*\",\n", "      \"Resource\": \"*\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleBucket, err := s3.NewBucket(ctx, "exampleBucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = codebuild.NewReportGroup(ctx, "exampleReportGroup", &codebuild.ReportGroupArgs{
// 			Type: pulumi.String("TEST"),
// 			ExportConfig: &codebuild.ReportGroupExportConfigArgs{
// 				Type: pulumi.String("S3"),
// 				S3Destination: &codebuild.ReportGroupExportConfigS3DestinationArgs{
// 					Bucket:             exampleBucket.ID(),
// 					EncryptionDisabled: pulumi.Bool(false),
// 					EncryptionKey:      exampleKey.Arn,
// 					Packaging:          pulumi.String("NONE"),
// 					Path:               pulumi.String("/some"),
// 				},
// 			},
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
// CodeBuild Report Group can be imported using the CodeBuild Report Group arn, e.g.
//
// ```sh
//  $ pulumi import aws:codebuild/reportGroup:ReportGroup example arn:aws:codebuild:us-west-2:123456789:report-group/report-group-name
// ```
type ReportGroup struct {
	pulumi.CustomResourceState

	// The ARN of Report Group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The date and time this Report Group was created.
	Created pulumi.StringOutput `pulumi:"created"`
	// Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
	ExportConfig ReportGroupExportConfigOutput `pulumi:"exportConfig"`
	// The name of a Report Group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value mapping of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The export configuration type. Valid values are `S3` and `NO_EXPORT`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReportGroup registers a new resource with the given unique name, arguments, and options.
func NewReportGroup(ctx *pulumi.Context,
	name string, args *ReportGroupArgs, opts ...pulumi.ResourceOption) (*ReportGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExportConfig == nil {
		return nil, errors.New("invalid value for required argument 'ExportConfig'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource ReportGroup
	err := ctx.RegisterResource("aws:codebuild/reportGroup:ReportGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReportGroup gets an existing ReportGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReportGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReportGroupState, opts ...pulumi.ResourceOption) (*ReportGroup, error) {
	var resource ReportGroup
	err := ctx.ReadResource("aws:codebuild/reportGroup:ReportGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReportGroup resources.
type reportGroupState struct {
	// The ARN of Report Group.
	Arn *string `pulumi:"arn"`
	// The date and time this Report Group was created.
	Created *string `pulumi:"created"`
	// Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
	ExportConfig *ReportGroupExportConfig `pulumi:"exportConfig"`
	// The name of a Report Group.
	Name *string `pulumi:"name"`
	// Key-value mapping of resource tags
	Tags map[string]string `pulumi:"tags"`
	// The export configuration type. Valid values are `S3` and `NO_EXPORT`.
	Type *string `pulumi:"type"`
}

type ReportGroupState struct {
	// The ARN of Report Group.
	Arn pulumi.StringPtrInput
	// The date and time this Report Group was created.
	Created pulumi.StringPtrInput
	// Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
	ExportConfig ReportGroupExportConfigPtrInput
	// The name of a Report Group.
	Name pulumi.StringPtrInput
	// Key-value mapping of resource tags
	Tags pulumi.StringMapInput
	// The export configuration type. Valid values are `S3` and `NO_EXPORT`.
	Type pulumi.StringPtrInput
}

func (ReportGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*reportGroupState)(nil)).Elem()
}

type reportGroupArgs struct {
	// Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
	ExportConfig ReportGroupExportConfig `pulumi:"exportConfig"`
	// The name of a Report Group.
	Name *string `pulumi:"name"`
	// Key-value mapping of resource tags
	Tags map[string]string `pulumi:"tags"`
	// The export configuration type. Valid values are `S3` and `NO_EXPORT`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a ReportGroup resource.
type ReportGroupArgs struct {
	// Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
	ExportConfig ReportGroupExportConfigInput
	// The name of a Report Group.
	Name pulumi.StringPtrInput
	// Key-value mapping of resource tags
	Tags pulumi.StringMapInput
	// The export configuration type. Valid values are `S3` and `NO_EXPORT`.
	Type pulumi.StringInput
}

func (ReportGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reportGroupArgs)(nil)).Elem()
}

type ReportGroupInput interface {
	pulumi.Input

	ToReportGroupOutput() ReportGroupOutput
	ToReportGroupOutputWithContext(ctx context.Context) ReportGroupOutput
}

func (*ReportGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportGroup)(nil))
}

func (i *ReportGroup) ToReportGroupOutput() ReportGroupOutput {
	return i.ToReportGroupOutputWithContext(context.Background())
}

func (i *ReportGroup) ToReportGroupOutputWithContext(ctx context.Context) ReportGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportGroupOutput)
}

func (i *ReportGroup) ToReportGroupPtrOutput() ReportGroupPtrOutput {
	return i.ToReportGroupPtrOutputWithContext(context.Background())
}

func (i *ReportGroup) ToReportGroupPtrOutputWithContext(ctx context.Context) ReportGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportGroupPtrOutput)
}

type ReportGroupPtrInput interface {
	pulumi.Input

	ToReportGroupPtrOutput() ReportGroupPtrOutput
	ToReportGroupPtrOutputWithContext(ctx context.Context) ReportGroupPtrOutput
}

type reportGroupPtrType ReportGroupArgs

func (*reportGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportGroup)(nil))
}

func (i *reportGroupPtrType) ToReportGroupPtrOutput() ReportGroupPtrOutput {
	return i.ToReportGroupPtrOutputWithContext(context.Background())
}

func (i *reportGroupPtrType) ToReportGroupPtrOutputWithContext(ctx context.Context) ReportGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportGroupPtrOutput)
}

// ReportGroupArrayInput is an input type that accepts ReportGroupArray and ReportGroupArrayOutput values.
// You can construct a concrete instance of `ReportGroupArrayInput` via:
//
//          ReportGroupArray{ ReportGroupArgs{...} }
type ReportGroupArrayInput interface {
	pulumi.Input

	ToReportGroupArrayOutput() ReportGroupArrayOutput
	ToReportGroupArrayOutputWithContext(context.Context) ReportGroupArrayOutput
}

type ReportGroupArray []ReportGroupInput

func (ReportGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ReportGroup)(nil))
}

func (i ReportGroupArray) ToReportGroupArrayOutput() ReportGroupArrayOutput {
	return i.ToReportGroupArrayOutputWithContext(context.Background())
}

func (i ReportGroupArray) ToReportGroupArrayOutputWithContext(ctx context.Context) ReportGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportGroupArrayOutput)
}

// ReportGroupMapInput is an input type that accepts ReportGroupMap and ReportGroupMapOutput values.
// You can construct a concrete instance of `ReportGroupMapInput` via:
//
//          ReportGroupMap{ "key": ReportGroupArgs{...} }
type ReportGroupMapInput interface {
	pulumi.Input

	ToReportGroupMapOutput() ReportGroupMapOutput
	ToReportGroupMapOutputWithContext(context.Context) ReportGroupMapOutput
}

type ReportGroupMap map[string]ReportGroupInput

func (ReportGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ReportGroup)(nil))
}

func (i ReportGroupMap) ToReportGroupMapOutput() ReportGroupMapOutput {
	return i.ToReportGroupMapOutputWithContext(context.Background())
}

func (i ReportGroupMap) ToReportGroupMapOutputWithContext(ctx context.Context) ReportGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportGroupMapOutput)
}

type ReportGroupOutput struct {
	*pulumi.OutputState
}

func (ReportGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportGroup)(nil))
}

func (o ReportGroupOutput) ToReportGroupOutput() ReportGroupOutput {
	return o
}

func (o ReportGroupOutput) ToReportGroupOutputWithContext(ctx context.Context) ReportGroupOutput {
	return o
}

func (o ReportGroupOutput) ToReportGroupPtrOutput() ReportGroupPtrOutput {
	return o.ToReportGroupPtrOutputWithContext(context.Background())
}

func (o ReportGroupOutput) ToReportGroupPtrOutputWithContext(ctx context.Context) ReportGroupPtrOutput {
	return o.ApplyT(func(v ReportGroup) *ReportGroup {
		return &v
	}).(ReportGroupPtrOutput)
}

type ReportGroupPtrOutput struct {
	*pulumi.OutputState
}

func (ReportGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportGroup)(nil))
}

func (o ReportGroupPtrOutput) ToReportGroupPtrOutput() ReportGroupPtrOutput {
	return o
}

func (o ReportGroupPtrOutput) ToReportGroupPtrOutputWithContext(ctx context.Context) ReportGroupPtrOutput {
	return o
}

type ReportGroupArrayOutput struct{ *pulumi.OutputState }

func (ReportGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportGroup)(nil))
}

func (o ReportGroupArrayOutput) ToReportGroupArrayOutput() ReportGroupArrayOutput {
	return o
}

func (o ReportGroupArrayOutput) ToReportGroupArrayOutputWithContext(ctx context.Context) ReportGroupArrayOutput {
	return o
}

func (o ReportGroupArrayOutput) Index(i pulumi.IntInput) ReportGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportGroup {
		return vs[0].([]ReportGroup)[vs[1].(int)]
	}).(ReportGroupOutput)
}

type ReportGroupMapOutput struct{ *pulumi.OutputState }

func (ReportGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReportGroup)(nil))
}

func (o ReportGroupMapOutput) ToReportGroupMapOutput() ReportGroupMapOutput {
	return o
}

func (o ReportGroupMapOutput) ToReportGroupMapOutputWithContext(ctx context.Context) ReportGroupMapOutput {
	return o
}

func (o ReportGroupMapOutput) MapIndex(k pulumi.StringInput) ReportGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReportGroup {
		return vs[0].(map[string]ReportGroup)[vs[1].(string)]
	}).(ReportGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(ReportGroupOutput{})
	pulumi.RegisterOutputType(ReportGroupPtrOutput{})
	pulumi.RegisterOutputType(ReportGroupArrayOutput{})
	pulumi.RegisterOutputType(ReportGroupMapOutput{})
}
