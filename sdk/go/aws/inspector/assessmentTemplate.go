// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package inspector

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Inspector assessment template
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/inspector"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := inspector.NewAssessmentTemplate(ctx, "example", &inspector.AssessmentTemplateArgs{
// 			TargetArn: pulumi.Any(aws_inspector_assessment_target.Example.Arn),
// 			Duration:  pulumi.Int(3600),
// 			RulesPackageArns: pulumi.StringArray{
// 				pulumi.String("arn:aws:inspector:us-west-2:758058086616:rulespackage/0-9hgA516p"),
// 				pulumi.String("arn:aws:inspector:us-west-2:758058086616:rulespackage/0-H5hpSawc"),
// 				pulumi.String("arn:aws:inspector:us-west-2:758058086616:rulespackage/0-JJOtZiqQ"),
// 				pulumi.String("arn:aws:inspector:us-west-2:758058086616:rulespackage/0-vg5GGHSD"),
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
// `aws_inspector_assessment_template` can be imported by using the template assessment ARN, e.g.
//
// ```sh
//  $ pulumi import aws:inspector/assessmentTemplate:AssessmentTemplate example arn:aws:inspector:us-west-2:123456789012:target/0-9IaAzhGR/template/0-WEcjR8CH
// ```
type AssessmentTemplate struct {
	pulumi.CustomResourceState

	// The template assessment ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The duration of the inspector run.
	Duration pulumi.IntOutput `pulumi:"duration"`
	// The name of the assessment template.
	Name pulumi.StringOutput `pulumi:"name"`
	// The rules to be used during the run.
	RulesPackageArns pulumi.StringArrayOutput `pulumi:"rulesPackageArns"`
	// Key-value map of tags for the Inspector assessment template.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The assessment target ARN to attach the template to.
	TargetArn pulumi.StringOutput `pulumi:"targetArn"`
}

// NewAssessmentTemplate registers a new resource with the given unique name, arguments, and options.
func NewAssessmentTemplate(ctx *pulumi.Context,
	name string, args *AssessmentTemplateArgs, opts ...pulumi.ResourceOption) (*AssessmentTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Duration == nil {
		return nil, errors.New("invalid value for required argument 'Duration'")
	}
	if args.RulesPackageArns == nil {
		return nil, errors.New("invalid value for required argument 'RulesPackageArns'")
	}
	if args.TargetArn == nil {
		return nil, errors.New("invalid value for required argument 'TargetArn'")
	}
	var resource AssessmentTemplate
	err := ctx.RegisterResource("aws:inspector/assessmentTemplate:AssessmentTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssessmentTemplate gets an existing AssessmentTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssessmentTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssessmentTemplateState, opts ...pulumi.ResourceOption) (*AssessmentTemplate, error) {
	var resource AssessmentTemplate
	err := ctx.ReadResource("aws:inspector/assessmentTemplate:AssessmentTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AssessmentTemplate resources.
type assessmentTemplateState struct {
	// The template assessment ARN.
	Arn *string `pulumi:"arn"`
	// The duration of the inspector run.
	Duration *int `pulumi:"duration"`
	// The name of the assessment template.
	Name *string `pulumi:"name"`
	// The rules to be used during the run.
	RulesPackageArns []string `pulumi:"rulesPackageArns"`
	// Key-value map of tags for the Inspector assessment template.
	Tags map[string]string `pulumi:"tags"`
	// The assessment target ARN to attach the template to.
	TargetArn *string `pulumi:"targetArn"`
}

type AssessmentTemplateState struct {
	// The template assessment ARN.
	Arn pulumi.StringPtrInput
	// The duration of the inspector run.
	Duration pulumi.IntPtrInput
	// The name of the assessment template.
	Name pulumi.StringPtrInput
	// The rules to be used during the run.
	RulesPackageArns pulumi.StringArrayInput
	// Key-value map of tags for the Inspector assessment template.
	Tags pulumi.StringMapInput
	// The assessment target ARN to attach the template to.
	TargetArn pulumi.StringPtrInput
}

func (AssessmentTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentTemplateState)(nil)).Elem()
}

type assessmentTemplateArgs struct {
	// The duration of the inspector run.
	Duration int `pulumi:"duration"`
	// The name of the assessment template.
	Name *string `pulumi:"name"`
	// The rules to be used during the run.
	RulesPackageArns []string `pulumi:"rulesPackageArns"`
	// Key-value map of tags for the Inspector assessment template.
	Tags map[string]string `pulumi:"tags"`
	// The assessment target ARN to attach the template to.
	TargetArn string `pulumi:"targetArn"`
}

// The set of arguments for constructing a AssessmentTemplate resource.
type AssessmentTemplateArgs struct {
	// The duration of the inspector run.
	Duration pulumi.IntInput
	// The name of the assessment template.
	Name pulumi.StringPtrInput
	// The rules to be used during the run.
	RulesPackageArns pulumi.StringArrayInput
	// Key-value map of tags for the Inspector assessment template.
	Tags pulumi.StringMapInput
	// The assessment target ARN to attach the template to.
	TargetArn pulumi.StringInput
}

func (AssessmentTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentTemplateArgs)(nil)).Elem()
}

type AssessmentTemplateInput interface {
	pulumi.Input

	ToAssessmentTemplateOutput() AssessmentTemplateOutput
	ToAssessmentTemplateOutputWithContext(ctx context.Context) AssessmentTemplateOutput
}

func (*AssessmentTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentTemplate)(nil))
}

func (i *AssessmentTemplate) ToAssessmentTemplateOutput() AssessmentTemplateOutput {
	return i.ToAssessmentTemplateOutputWithContext(context.Background())
}

func (i *AssessmentTemplate) ToAssessmentTemplateOutputWithContext(ctx context.Context) AssessmentTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentTemplateOutput)
}

func (i *AssessmentTemplate) ToAssessmentTemplatePtrOutput() AssessmentTemplatePtrOutput {
	return i.ToAssessmentTemplatePtrOutputWithContext(context.Background())
}

func (i *AssessmentTemplate) ToAssessmentTemplatePtrOutputWithContext(ctx context.Context) AssessmentTemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentTemplatePtrOutput)
}

type AssessmentTemplatePtrInput interface {
	pulumi.Input

	ToAssessmentTemplatePtrOutput() AssessmentTemplatePtrOutput
	ToAssessmentTemplatePtrOutputWithContext(ctx context.Context) AssessmentTemplatePtrOutput
}

type assessmentTemplatePtrType AssessmentTemplateArgs

func (*assessmentTemplatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentTemplate)(nil))
}

func (i *assessmentTemplatePtrType) ToAssessmentTemplatePtrOutput() AssessmentTemplatePtrOutput {
	return i.ToAssessmentTemplatePtrOutputWithContext(context.Background())
}

func (i *assessmentTemplatePtrType) ToAssessmentTemplatePtrOutputWithContext(ctx context.Context) AssessmentTemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentTemplatePtrOutput)
}

// AssessmentTemplateArrayInput is an input type that accepts AssessmentTemplateArray and AssessmentTemplateArrayOutput values.
// You can construct a concrete instance of `AssessmentTemplateArrayInput` via:
//
//          AssessmentTemplateArray{ AssessmentTemplateArgs{...} }
type AssessmentTemplateArrayInput interface {
	pulumi.Input

	ToAssessmentTemplateArrayOutput() AssessmentTemplateArrayOutput
	ToAssessmentTemplateArrayOutputWithContext(context.Context) AssessmentTemplateArrayOutput
}

type AssessmentTemplateArray []AssessmentTemplateInput

func (AssessmentTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AssessmentTemplate)(nil))
}

func (i AssessmentTemplateArray) ToAssessmentTemplateArrayOutput() AssessmentTemplateArrayOutput {
	return i.ToAssessmentTemplateArrayOutputWithContext(context.Background())
}

func (i AssessmentTemplateArray) ToAssessmentTemplateArrayOutputWithContext(ctx context.Context) AssessmentTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentTemplateArrayOutput)
}

// AssessmentTemplateMapInput is an input type that accepts AssessmentTemplateMap and AssessmentTemplateMapOutput values.
// You can construct a concrete instance of `AssessmentTemplateMapInput` via:
//
//          AssessmentTemplateMap{ "key": AssessmentTemplateArgs{...} }
type AssessmentTemplateMapInput interface {
	pulumi.Input

	ToAssessmentTemplateMapOutput() AssessmentTemplateMapOutput
	ToAssessmentTemplateMapOutputWithContext(context.Context) AssessmentTemplateMapOutput
}

type AssessmentTemplateMap map[string]AssessmentTemplateInput

func (AssessmentTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AssessmentTemplate)(nil))
}

func (i AssessmentTemplateMap) ToAssessmentTemplateMapOutput() AssessmentTemplateMapOutput {
	return i.ToAssessmentTemplateMapOutputWithContext(context.Background())
}

func (i AssessmentTemplateMap) ToAssessmentTemplateMapOutputWithContext(ctx context.Context) AssessmentTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentTemplateMapOutput)
}

type AssessmentTemplateOutput struct {
	*pulumi.OutputState
}

func (AssessmentTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentTemplate)(nil))
}

func (o AssessmentTemplateOutput) ToAssessmentTemplateOutput() AssessmentTemplateOutput {
	return o
}

func (o AssessmentTemplateOutput) ToAssessmentTemplateOutputWithContext(ctx context.Context) AssessmentTemplateOutput {
	return o
}

func (o AssessmentTemplateOutput) ToAssessmentTemplatePtrOutput() AssessmentTemplatePtrOutput {
	return o.ToAssessmentTemplatePtrOutputWithContext(context.Background())
}

func (o AssessmentTemplateOutput) ToAssessmentTemplatePtrOutputWithContext(ctx context.Context) AssessmentTemplatePtrOutput {
	return o.ApplyT(func(v AssessmentTemplate) *AssessmentTemplate {
		return &v
	}).(AssessmentTemplatePtrOutput)
}

type AssessmentTemplatePtrOutput struct {
	*pulumi.OutputState
}

func (AssessmentTemplatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentTemplate)(nil))
}

func (o AssessmentTemplatePtrOutput) ToAssessmentTemplatePtrOutput() AssessmentTemplatePtrOutput {
	return o
}

func (o AssessmentTemplatePtrOutput) ToAssessmentTemplatePtrOutputWithContext(ctx context.Context) AssessmentTemplatePtrOutput {
	return o
}

type AssessmentTemplateArrayOutput struct{ *pulumi.OutputState }

func (AssessmentTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AssessmentTemplate)(nil))
}

func (o AssessmentTemplateArrayOutput) ToAssessmentTemplateArrayOutput() AssessmentTemplateArrayOutput {
	return o
}

func (o AssessmentTemplateArrayOutput) ToAssessmentTemplateArrayOutputWithContext(ctx context.Context) AssessmentTemplateArrayOutput {
	return o
}

func (o AssessmentTemplateArrayOutput) Index(i pulumi.IntInput) AssessmentTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AssessmentTemplate {
		return vs[0].([]AssessmentTemplate)[vs[1].(int)]
	}).(AssessmentTemplateOutput)
}

type AssessmentTemplateMapOutput struct{ *pulumi.OutputState }

func (AssessmentTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AssessmentTemplate)(nil))
}

func (o AssessmentTemplateMapOutput) ToAssessmentTemplateMapOutput() AssessmentTemplateMapOutput {
	return o
}

func (o AssessmentTemplateMapOutput) ToAssessmentTemplateMapOutputWithContext(ctx context.Context) AssessmentTemplateMapOutput {
	return o
}

func (o AssessmentTemplateMapOutput) MapIndex(k pulumi.StringInput) AssessmentTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AssessmentTemplate {
		return vs[0].(map[string]AssessmentTemplate)[vs[1].(string)]
	}).(AssessmentTemplateOutput)
}

func init() {
	pulumi.RegisterOutputType(AssessmentTemplateOutput{})
	pulumi.RegisterOutputType(AssessmentTemplatePtrOutput{})
	pulumi.RegisterOutputType(AssessmentTemplateArrayOutput{})
	pulumi.RegisterOutputType(AssessmentTemplateMapOutput{})
}
