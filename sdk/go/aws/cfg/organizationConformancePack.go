// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Config Organization Conformance Pack. More information can be found in the [Managing Conformance Packs Across all Accounts in Your Organization](https://docs.aws.amazon.com/config/latest/developerguide/conformance-pack-organization-apis.html) and [AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) documentation. Example conformance pack templates may be found in the [AWS Config Rules Repository](https://github.com/awslabs/aws-config-rules/tree/master/aws-config-conformance-packs).
//
// > **NOTE:** This resource must be created in the Organization master account or a delegated administrator account, and the Organization must have all features enabled. Every Organization account except those configured in the `excludedAccounts` argument must have a Configuration Recorder with proper IAM permissions before the Organization Conformance Pack will successfully create or update. See also the [`cfg.Recorder` resource](https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html).
//
// ## Example Usage
// ### Using Template Body
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cfg"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleOrganization, err := organizations.NewOrganization(ctx, "exampleOrganization", &organizations.OrganizationArgs{
// 			AwsServiceAccessPrincipals: pulumi.StringArray{
// 				pulumi.String("config-multiaccountsetup.amazonaws.com"),
// 			},
// 			FeatureSet: pulumi.String("ALL"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cfg.NewOrganizationConformancePack(ctx, "exampleOrganizationConformancePack", &cfg.OrganizationConformancePackArgs{
// 			InputParameters: cfg.OrganizationConformancePackInputParameterArray{
// 				&cfg.OrganizationConformancePackInputParameterArgs{
// 					ParameterName:  pulumi.String("AccessKeysRotatedParameterMaxAccessKeyAge"),
// 					ParameterValue: pulumi.String("90"),
// 				},
// 			},
// 			TemplateBody: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v", "Parameters:\n", "  AccessKeysRotatedParameterMaxAccessKeyAge:\n", "    Type: String\n", "Resources:\n", "  IAMPasswordPolicy:\n", "    Properties:\n", "      ConfigRuleName: IAMPasswordPolicy\n", "      Source:\n", "        Owner: AWS\n", "        SourceIdentifier: IAM_PASSWORD_POLICY\n", "    Type: AWS::Config::ConfigRule\n")),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			aws_config_configuration_recorder.Example,
// 			exampleOrganization,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Using Template S3 URI
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cfg"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/organizations"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleOrganization, err := organizations.NewOrganization(ctx, "exampleOrganization", &organizations.OrganizationArgs{
// 			AwsServiceAccessPrincipals: pulumi.StringArray{
// 				pulumi.String("config-multiaccountsetup.amazonaws.com"),
// 			},
// 			FeatureSet: pulumi.String("ALL"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleBucket, err := s3.NewBucket(ctx, "exampleBucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleBucketObject, err := s3.NewBucketObject(ctx, "exampleBucketObject", &s3.BucketObjectArgs{
// 			Bucket:  exampleBucket.ID(),
// 			Key:     pulumi.String("example-key"),
// 			Content: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v", "Resources:\n", "  IAMPasswordPolicy:\n", "    Properties:\n", "      ConfigRuleName: IAMPasswordPolicy\n", "      Source:\n", "        Owner: AWS\n", "        SourceIdentifier: IAM_PASSWORD_POLICY\n", "    Type: AWS::Config::ConfigRule\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cfg.NewOrganizationConformancePack(ctx, "exampleOrganizationConformancePack", &cfg.OrganizationConformancePackArgs{
// 			TemplateS3Uri: pulumi.All(exampleBucket.Bucket, exampleBucketObject.Key).ApplyT(func(_args []interface{}) (string, error) {
// 				bucket := _args[0].(string)
// 				key := _args[1].(string)
// 				return fmt.Sprintf("%v%v%v%v", "s3://", bucket, "/", key), nil
// 			}).(pulumi.StringOutput),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			aws_config_configuration_recorder.Example,
// 			exampleOrganization,
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
// Config Organization Conformance Packs can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:cfg/organizationConformancePack:OrganizationConformancePack example example
// ```
type OrganizationConformancePack struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the organization conformance pack.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Amazon S3 bucket where AWS Config stores conformance pack templates. Delivery bucket must begin with `awsconfigconforms` prefix. Maximum length of 63.
	DeliveryS3Bucket pulumi.StringPtrOutput `pulumi:"deliveryS3Bucket"`
	// The prefix for the Amazon S3 bucket. Maximum length of 1024.
	DeliveryS3KeyPrefix pulumi.StringPtrOutput `pulumi:"deliveryS3KeyPrefix"`
	// Set of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack. Maximum of 1000 accounts.
	ExcludedAccounts pulumi.StringArrayOutput `pulumi:"excludedAccounts"`
	// Set of configuration blocks describing input parameters passed to the conformance pack template. Documented below. When configured, the parameters must also be included in the `templateBody` or in the template stored in Amazon S3 if using `templateS3Uri`.
	InputParameters OrganizationConformancePackInputParameterArrayOutput `pulumi:"inputParameters"`
	// The name of the organization conformance pack. Must begin with a letter and contain from 1 to 128 alphanumeric characters and hyphens.
	Name pulumi.StringOutput `pulumi:"name"`
	// A string containing full conformance pack template body. Maximum length of 51200. Drift detection is not possible with this argument.
	TemplateBody pulumi.StringPtrOutput `pulumi:"templateBody"`
	// Location of file, e.g. `s3://bucketname/prefix`, containing the template body. The uri must point to the conformance pack template that is located in an Amazon S3 bucket in the same region as the conformance pack. Maximum length of 1024. Drift detection is not possible with this argument.
	TemplateS3Uri pulumi.StringPtrOutput `pulumi:"templateS3Uri"`
}

// NewOrganizationConformancePack registers a new resource with the given unique name, arguments, and options.
func NewOrganizationConformancePack(ctx *pulumi.Context,
	name string, args *OrganizationConformancePackArgs, opts ...pulumi.ResourceOption) (*OrganizationConformancePack, error) {
	if args == nil {
		args = &OrganizationConformancePackArgs{}
	}

	var resource OrganizationConformancePack
	err := ctx.RegisterResource("aws:cfg/organizationConformancePack:OrganizationConformancePack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationConformancePack gets an existing OrganizationConformancePack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationConformancePack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationConformancePackState, opts ...pulumi.ResourceOption) (*OrganizationConformancePack, error) {
	var resource OrganizationConformancePack
	err := ctx.ReadResource("aws:cfg/organizationConformancePack:OrganizationConformancePack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationConformancePack resources.
type organizationConformancePackState struct {
	// Amazon Resource Name (ARN) of the organization conformance pack.
	Arn *string `pulumi:"arn"`
	// Amazon S3 bucket where AWS Config stores conformance pack templates. Delivery bucket must begin with `awsconfigconforms` prefix. Maximum length of 63.
	DeliveryS3Bucket *string `pulumi:"deliveryS3Bucket"`
	// The prefix for the Amazon S3 bucket. Maximum length of 1024.
	DeliveryS3KeyPrefix *string `pulumi:"deliveryS3KeyPrefix"`
	// Set of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack. Maximum of 1000 accounts.
	ExcludedAccounts []string `pulumi:"excludedAccounts"`
	// Set of configuration blocks describing input parameters passed to the conformance pack template. Documented below. When configured, the parameters must also be included in the `templateBody` or in the template stored in Amazon S3 if using `templateS3Uri`.
	InputParameters []OrganizationConformancePackInputParameter `pulumi:"inputParameters"`
	// The name of the organization conformance pack. Must begin with a letter and contain from 1 to 128 alphanumeric characters and hyphens.
	Name *string `pulumi:"name"`
	// A string containing full conformance pack template body. Maximum length of 51200. Drift detection is not possible with this argument.
	TemplateBody *string `pulumi:"templateBody"`
	// Location of file, e.g. `s3://bucketname/prefix`, containing the template body. The uri must point to the conformance pack template that is located in an Amazon S3 bucket in the same region as the conformance pack. Maximum length of 1024. Drift detection is not possible with this argument.
	TemplateS3Uri *string `pulumi:"templateS3Uri"`
}

type OrganizationConformancePackState struct {
	// Amazon Resource Name (ARN) of the organization conformance pack.
	Arn pulumi.StringPtrInput
	// Amazon S3 bucket where AWS Config stores conformance pack templates. Delivery bucket must begin with `awsconfigconforms` prefix. Maximum length of 63.
	DeliveryS3Bucket pulumi.StringPtrInput
	// The prefix for the Amazon S3 bucket. Maximum length of 1024.
	DeliveryS3KeyPrefix pulumi.StringPtrInput
	// Set of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack. Maximum of 1000 accounts.
	ExcludedAccounts pulumi.StringArrayInput
	// Set of configuration blocks describing input parameters passed to the conformance pack template. Documented below. When configured, the parameters must also be included in the `templateBody` or in the template stored in Amazon S3 if using `templateS3Uri`.
	InputParameters OrganizationConformancePackInputParameterArrayInput
	// The name of the organization conformance pack. Must begin with a letter and contain from 1 to 128 alphanumeric characters and hyphens.
	Name pulumi.StringPtrInput
	// A string containing full conformance pack template body. Maximum length of 51200. Drift detection is not possible with this argument.
	TemplateBody pulumi.StringPtrInput
	// Location of file, e.g. `s3://bucketname/prefix`, containing the template body. The uri must point to the conformance pack template that is located in an Amazon S3 bucket in the same region as the conformance pack. Maximum length of 1024. Drift detection is not possible with this argument.
	TemplateS3Uri pulumi.StringPtrInput
}

func (OrganizationConformancePackState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationConformancePackState)(nil)).Elem()
}

type organizationConformancePackArgs struct {
	// Amazon S3 bucket where AWS Config stores conformance pack templates. Delivery bucket must begin with `awsconfigconforms` prefix. Maximum length of 63.
	DeliveryS3Bucket *string `pulumi:"deliveryS3Bucket"`
	// The prefix for the Amazon S3 bucket. Maximum length of 1024.
	DeliveryS3KeyPrefix *string `pulumi:"deliveryS3KeyPrefix"`
	// Set of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack. Maximum of 1000 accounts.
	ExcludedAccounts []string `pulumi:"excludedAccounts"`
	// Set of configuration blocks describing input parameters passed to the conformance pack template. Documented below. When configured, the parameters must also be included in the `templateBody` or in the template stored in Amazon S3 if using `templateS3Uri`.
	InputParameters []OrganizationConformancePackInputParameter `pulumi:"inputParameters"`
	// The name of the organization conformance pack. Must begin with a letter and contain from 1 to 128 alphanumeric characters and hyphens.
	Name *string `pulumi:"name"`
	// A string containing full conformance pack template body. Maximum length of 51200. Drift detection is not possible with this argument.
	TemplateBody *string `pulumi:"templateBody"`
	// Location of file, e.g. `s3://bucketname/prefix`, containing the template body. The uri must point to the conformance pack template that is located in an Amazon S3 bucket in the same region as the conformance pack. Maximum length of 1024. Drift detection is not possible with this argument.
	TemplateS3Uri *string `pulumi:"templateS3Uri"`
}

// The set of arguments for constructing a OrganizationConformancePack resource.
type OrganizationConformancePackArgs struct {
	// Amazon S3 bucket where AWS Config stores conformance pack templates. Delivery bucket must begin with `awsconfigconforms` prefix. Maximum length of 63.
	DeliveryS3Bucket pulumi.StringPtrInput
	// The prefix for the Amazon S3 bucket. Maximum length of 1024.
	DeliveryS3KeyPrefix pulumi.StringPtrInput
	// Set of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack. Maximum of 1000 accounts.
	ExcludedAccounts pulumi.StringArrayInput
	// Set of configuration blocks describing input parameters passed to the conformance pack template. Documented below. When configured, the parameters must also be included in the `templateBody` or in the template stored in Amazon S3 if using `templateS3Uri`.
	InputParameters OrganizationConformancePackInputParameterArrayInput
	// The name of the organization conformance pack. Must begin with a letter and contain from 1 to 128 alphanumeric characters and hyphens.
	Name pulumi.StringPtrInput
	// A string containing full conformance pack template body. Maximum length of 51200. Drift detection is not possible with this argument.
	TemplateBody pulumi.StringPtrInput
	// Location of file, e.g. `s3://bucketname/prefix`, containing the template body. The uri must point to the conformance pack template that is located in an Amazon S3 bucket in the same region as the conformance pack. Maximum length of 1024. Drift detection is not possible with this argument.
	TemplateS3Uri pulumi.StringPtrInput
}

func (OrganizationConformancePackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationConformancePackArgs)(nil)).Elem()
}

type OrganizationConformancePackInput interface {
	pulumi.Input

	ToOrganizationConformancePackOutput() OrganizationConformancePackOutput
	ToOrganizationConformancePackOutputWithContext(ctx context.Context) OrganizationConformancePackOutput
}

func (*OrganizationConformancePack) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationConformancePack)(nil))
}

func (i *OrganizationConformancePack) ToOrganizationConformancePackOutput() OrganizationConformancePackOutput {
	return i.ToOrganizationConformancePackOutputWithContext(context.Background())
}

func (i *OrganizationConformancePack) ToOrganizationConformancePackOutputWithContext(ctx context.Context) OrganizationConformancePackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationConformancePackOutput)
}

func (i *OrganizationConformancePack) ToOrganizationConformancePackPtrOutput() OrganizationConformancePackPtrOutput {
	return i.ToOrganizationConformancePackPtrOutputWithContext(context.Background())
}

func (i *OrganizationConformancePack) ToOrganizationConformancePackPtrOutputWithContext(ctx context.Context) OrganizationConformancePackPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationConformancePackPtrOutput)
}

type OrganizationConformancePackPtrInput interface {
	pulumi.Input

	ToOrganizationConformancePackPtrOutput() OrganizationConformancePackPtrOutput
	ToOrganizationConformancePackPtrOutputWithContext(ctx context.Context) OrganizationConformancePackPtrOutput
}

type organizationConformancePackPtrType OrganizationConformancePackArgs

func (*organizationConformancePackPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationConformancePack)(nil))
}

func (i *organizationConformancePackPtrType) ToOrganizationConformancePackPtrOutput() OrganizationConformancePackPtrOutput {
	return i.ToOrganizationConformancePackPtrOutputWithContext(context.Background())
}

func (i *organizationConformancePackPtrType) ToOrganizationConformancePackPtrOutputWithContext(ctx context.Context) OrganizationConformancePackPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationConformancePackPtrOutput)
}

// OrganizationConformancePackArrayInput is an input type that accepts OrganizationConformancePackArray and OrganizationConformancePackArrayOutput values.
// You can construct a concrete instance of `OrganizationConformancePackArrayInput` via:
//
//          OrganizationConformancePackArray{ OrganizationConformancePackArgs{...} }
type OrganizationConformancePackArrayInput interface {
	pulumi.Input

	ToOrganizationConformancePackArrayOutput() OrganizationConformancePackArrayOutput
	ToOrganizationConformancePackArrayOutputWithContext(context.Context) OrganizationConformancePackArrayOutput
}

type OrganizationConformancePackArray []OrganizationConformancePackInput

func (OrganizationConformancePackArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationConformancePack)(nil)).Elem()
}

func (i OrganizationConformancePackArray) ToOrganizationConformancePackArrayOutput() OrganizationConformancePackArrayOutput {
	return i.ToOrganizationConformancePackArrayOutputWithContext(context.Background())
}

func (i OrganizationConformancePackArray) ToOrganizationConformancePackArrayOutputWithContext(ctx context.Context) OrganizationConformancePackArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationConformancePackArrayOutput)
}

// OrganizationConformancePackMapInput is an input type that accepts OrganizationConformancePackMap and OrganizationConformancePackMapOutput values.
// You can construct a concrete instance of `OrganizationConformancePackMapInput` via:
//
//          OrganizationConformancePackMap{ "key": OrganizationConformancePackArgs{...} }
type OrganizationConformancePackMapInput interface {
	pulumi.Input

	ToOrganizationConformancePackMapOutput() OrganizationConformancePackMapOutput
	ToOrganizationConformancePackMapOutputWithContext(context.Context) OrganizationConformancePackMapOutput
}

type OrganizationConformancePackMap map[string]OrganizationConformancePackInput

func (OrganizationConformancePackMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationConformancePack)(nil)).Elem()
}

func (i OrganizationConformancePackMap) ToOrganizationConformancePackMapOutput() OrganizationConformancePackMapOutput {
	return i.ToOrganizationConformancePackMapOutputWithContext(context.Background())
}

func (i OrganizationConformancePackMap) ToOrganizationConformancePackMapOutputWithContext(ctx context.Context) OrganizationConformancePackMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationConformancePackMapOutput)
}

type OrganizationConformancePackOutput struct{ *pulumi.OutputState }

func (OrganizationConformancePackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationConformancePack)(nil))
}

func (o OrganizationConformancePackOutput) ToOrganizationConformancePackOutput() OrganizationConformancePackOutput {
	return o
}

func (o OrganizationConformancePackOutput) ToOrganizationConformancePackOutputWithContext(ctx context.Context) OrganizationConformancePackOutput {
	return o
}

func (o OrganizationConformancePackOutput) ToOrganizationConformancePackPtrOutput() OrganizationConformancePackPtrOutput {
	return o.ToOrganizationConformancePackPtrOutputWithContext(context.Background())
}

func (o OrganizationConformancePackOutput) ToOrganizationConformancePackPtrOutputWithContext(ctx context.Context) OrganizationConformancePackPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OrganizationConformancePack) *OrganizationConformancePack {
		return &v
	}).(OrganizationConformancePackPtrOutput)
}

type OrganizationConformancePackPtrOutput struct{ *pulumi.OutputState }

func (OrganizationConformancePackPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationConformancePack)(nil))
}

func (o OrganizationConformancePackPtrOutput) ToOrganizationConformancePackPtrOutput() OrganizationConformancePackPtrOutput {
	return o
}

func (o OrganizationConformancePackPtrOutput) ToOrganizationConformancePackPtrOutputWithContext(ctx context.Context) OrganizationConformancePackPtrOutput {
	return o
}

func (o OrganizationConformancePackPtrOutput) Elem() OrganizationConformancePackOutput {
	return o.ApplyT(func(v *OrganizationConformancePack) OrganizationConformancePack {
		if v != nil {
			return *v
		}
		var ret OrganizationConformancePack
		return ret
	}).(OrganizationConformancePackOutput)
}

type OrganizationConformancePackArrayOutput struct{ *pulumi.OutputState }

func (OrganizationConformancePackArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OrganizationConformancePack)(nil))
}

func (o OrganizationConformancePackArrayOutput) ToOrganizationConformancePackArrayOutput() OrganizationConformancePackArrayOutput {
	return o
}

func (o OrganizationConformancePackArrayOutput) ToOrganizationConformancePackArrayOutputWithContext(ctx context.Context) OrganizationConformancePackArrayOutput {
	return o
}

func (o OrganizationConformancePackArrayOutput) Index(i pulumi.IntInput) OrganizationConformancePackOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OrganizationConformancePack {
		return vs[0].([]OrganizationConformancePack)[vs[1].(int)]
	}).(OrganizationConformancePackOutput)
}

type OrganizationConformancePackMapOutput struct{ *pulumi.OutputState }

func (OrganizationConformancePackMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OrganizationConformancePack)(nil))
}

func (o OrganizationConformancePackMapOutput) ToOrganizationConformancePackMapOutput() OrganizationConformancePackMapOutput {
	return o
}

func (o OrganizationConformancePackMapOutput) ToOrganizationConformancePackMapOutputWithContext(ctx context.Context) OrganizationConformancePackMapOutput {
	return o
}

func (o OrganizationConformancePackMapOutput) MapIndex(k pulumi.StringInput) OrganizationConformancePackOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OrganizationConformancePack {
		return vs[0].(map[string]OrganizationConformancePack)[vs[1].(string)]
	}).(OrganizationConformancePackOutput)
}

func init() {
	pulumi.RegisterOutputType(OrganizationConformancePackOutput{})
	pulumi.RegisterOutputType(OrganizationConformancePackPtrOutput{})
	pulumi.RegisterOutputType(OrganizationConformancePackArrayOutput{})
	pulumi.RegisterOutputType(OrganizationConformancePackMapOutput{})
}
