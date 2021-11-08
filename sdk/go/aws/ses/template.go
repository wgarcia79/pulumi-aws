// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a SES template.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ses"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ses.NewTemplate(ctx, "myTemplate", &ses.TemplateArgs{
// 			Html:    pulumi.String("<h1>Hello {{name}},</h1><p>Your favorite animal is {{favoriteanimal}}.</p>"),
// 			Subject: pulumi.String("Greetings, {{name}}!"),
// 			Text: pulumi.String(fmt.Sprintf("%v%v", "Hello {{name}},\n", "Your favorite animal is {{favoriteanimal}}.\n")),
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
// SES templates can be imported using the template name, e.g.,
//
// ```sh
//  $ pulumi import aws:ses/template:Template MyTemplate MyTemplate
// ```
type Template struct {
	pulumi.CustomResourceState

	// The ARN of the SES template
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The HTML body of the email. Must be less than 500KB in size, including both the text and HTML parts.
	Html pulumi.StringPtrOutput `pulumi:"html"`
	// The name of the template. Cannot exceed 64 characters. You will refer to this name when you send email.
	Name pulumi.StringOutput `pulumi:"name"`
	// The subject line of the email.
	Subject pulumi.StringPtrOutput `pulumi:"subject"`
	// The email body that will be visible to recipients whose email clients do not display HTML. Must be less than 500KB in size, including both the text and HTML parts.
	Text pulumi.StringPtrOutput `pulumi:"text"`
}

// NewTemplate registers a new resource with the given unique name, arguments, and options.
func NewTemplate(ctx *pulumi.Context,
	name string, args *TemplateArgs, opts ...pulumi.ResourceOption) (*Template, error) {
	if args == nil {
		args = &TemplateArgs{}
	}

	var resource Template
	err := ctx.RegisterResource("aws:ses/template:Template", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplate gets an existing Template resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateState, opts ...pulumi.ResourceOption) (*Template, error) {
	var resource Template
	err := ctx.ReadResource("aws:ses/template:Template", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Template resources.
type templateState struct {
	// The ARN of the SES template
	Arn *string `pulumi:"arn"`
	// The HTML body of the email. Must be less than 500KB in size, including both the text and HTML parts.
	Html *string `pulumi:"html"`
	// The name of the template. Cannot exceed 64 characters. You will refer to this name when you send email.
	Name *string `pulumi:"name"`
	// The subject line of the email.
	Subject *string `pulumi:"subject"`
	// The email body that will be visible to recipients whose email clients do not display HTML. Must be less than 500KB in size, including both the text and HTML parts.
	Text *string `pulumi:"text"`
}

type TemplateState struct {
	// The ARN of the SES template
	Arn pulumi.StringPtrInput
	// The HTML body of the email. Must be less than 500KB in size, including both the text and HTML parts.
	Html pulumi.StringPtrInput
	// The name of the template. Cannot exceed 64 characters. You will refer to this name when you send email.
	Name pulumi.StringPtrInput
	// The subject line of the email.
	Subject pulumi.StringPtrInput
	// The email body that will be visible to recipients whose email clients do not display HTML. Must be less than 500KB in size, including both the text and HTML parts.
	Text pulumi.StringPtrInput
}

func (TemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateState)(nil)).Elem()
}

type templateArgs struct {
	// The HTML body of the email. Must be less than 500KB in size, including both the text and HTML parts.
	Html *string `pulumi:"html"`
	// The name of the template. Cannot exceed 64 characters. You will refer to this name when you send email.
	Name *string `pulumi:"name"`
	// The subject line of the email.
	Subject *string `pulumi:"subject"`
	// The email body that will be visible to recipients whose email clients do not display HTML. Must be less than 500KB in size, including both the text and HTML parts.
	Text *string `pulumi:"text"`
}

// The set of arguments for constructing a Template resource.
type TemplateArgs struct {
	// The HTML body of the email. Must be less than 500KB in size, including both the text and HTML parts.
	Html pulumi.StringPtrInput
	// The name of the template. Cannot exceed 64 characters. You will refer to this name when you send email.
	Name pulumi.StringPtrInput
	// The subject line of the email.
	Subject pulumi.StringPtrInput
	// The email body that will be visible to recipients whose email clients do not display HTML. Must be less than 500KB in size, including both the text and HTML parts.
	Text pulumi.StringPtrInput
}

func (TemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateArgs)(nil)).Elem()
}

type TemplateInput interface {
	pulumi.Input

	ToTemplateOutput() TemplateOutput
	ToTemplateOutputWithContext(ctx context.Context) TemplateOutput
}

func (*Template) ElementType() reflect.Type {
	return reflect.TypeOf((*Template)(nil))
}

func (i *Template) ToTemplateOutput() TemplateOutput {
	return i.ToTemplateOutputWithContext(context.Background())
}

func (i *Template) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateOutput)
}

func (i *Template) ToTemplatePtrOutput() TemplatePtrOutput {
	return i.ToTemplatePtrOutputWithContext(context.Background())
}

func (i *Template) ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplatePtrOutput)
}

type TemplatePtrInput interface {
	pulumi.Input

	ToTemplatePtrOutput() TemplatePtrOutput
	ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput
}

type templatePtrType TemplateArgs

func (*templatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil))
}

func (i *templatePtrType) ToTemplatePtrOutput() TemplatePtrOutput {
	return i.ToTemplatePtrOutputWithContext(context.Background())
}

func (i *templatePtrType) ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplatePtrOutput)
}

// TemplateArrayInput is an input type that accepts TemplateArray and TemplateArrayOutput values.
// You can construct a concrete instance of `TemplateArrayInput` via:
//
//          TemplateArray{ TemplateArgs{...} }
type TemplateArrayInput interface {
	pulumi.Input

	ToTemplateArrayOutput() TemplateArrayOutput
	ToTemplateArrayOutputWithContext(context.Context) TemplateArrayOutput
}

type TemplateArray []TemplateInput

func (TemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Template)(nil)).Elem()
}

func (i TemplateArray) ToTemplateArrayOutput() TemplateArrayOutput {
	return i.ToTemplateArrayOutputWithContext(context.Background())
}

func (i TemplateArray) ToTemplateArrayOutputWithContext(ctx context.Context) TemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateArrayOutput)
}

// TemplateMapInput is an input type that accepts TemplateMap and TemplateMapOutput values.
// You can construct a concrete instance of `TemplateMapInput` via:
//
//          TemplateMap{ "key": TemplateArgs{...} }
type TemplateMapInput interface {
	pulumi.Input

	ToTemplateMapOutput() TemplateMapOutput
	ToTemplateMapOutputWithContext(context.Context) TemplateMapOutput
}

type TemplateMap map[string]TemplateInput

func (TemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Template)(nil)).Elem()
}

func (i TemplateMap) ToTemplateMapOutput() TemplateMapOutput {
	return i.ToTemplateMapOutputWithContext(context.Background())
}

func (i TemplateMap) ToTemplateMapOutputWithContext(ctx context.Context) TemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateMapOutput)
}

type TemplateOutput struct{ *pulumi.OutputState }

func (TemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Template)(nil))
}

func (o TemplateOutput) ToTemplateOutput() TemplateOutput {
	return o
}

func (o TemplateOutput) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return o
}

func (o TemplateOutput) ToTemplatePtrOutput() TemplatePtrOutput {
	return o.ToTemplatePtrOutputWithContext(context.Background())
}

func (o TemplateOutput) ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Template) *Template {
		return &v
	}).(TemplatePtrOutput)
}

type TemplatePtrOutput struct{ *pulumi.OutputState }

func (TemplatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil))
}

func (o TemplatePtrOutput) ToTemplatePtrOutput() TemplatePtrOutput {
	return o
}

func (o TemplatePtrOutput) ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput {
	return o
}

func (o TemplatePtrOutput) Elem() TemplateOutput {
	return o.ApplyT(func(v *Template) Template {
		if v != nil {
			return *v
		}
		var ret Template
		return ret
	}).(TemplateOutput)
}

type TemplateArrayOutput struct{ *pulumi.OutputState }

func (TemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Template)(nil))
}

func (o TemplateArrayOutput) ToTemplateArrayOutput() TemplateArrayOutput {
	return o
}

func (o TemplateArrayOutput) ToTemplateArrayOutputWithContext(ctx context.Context) TemplateArrayOutput {
	return o
}

func (o TemplateArrayOutput) Index(i pulumi.IntInput) TemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Template {
		return vs[0].([]Template)[vs[1].(int)]
	}).(TemplateOutput)
}

type TemplateMapOutput struct{ *pulumi.OutputState }

func (TemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Template)(nil))
}

func (o TemplateMapOutput) ToTemplateMapOutput() TemplateMapOutput {
	return o
}

func (o TemplateMapOutput) ToTemplateMapOutputWithContext(ctx context.Context) TemplateMapOutput {
	return o
}

func (o TemplateMapOutput) MapIndex(k pulumi.StringInput) TemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Template {
		return vs[0].(map[string]Template)[vs[1].(string)]
	}).(TemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateInput)(nil)).Elem(), &Template{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplatePtrInput)(nil)).Elem(), &Template{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateArrayInput)(nil)).Elem(), TemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateMapInput)(nil)).Elem(), TemplateMap{})
	pulumi.RegisterOutputType(TemplateOutput{})
	pulumi.RegisterOutputType(TemplatePtrOutput{})
	pulumi.RegisterOutputType(TemplateArrayOutput{})
	pulumi.RegisterOutputType(TemplateMapOutput{})
}
