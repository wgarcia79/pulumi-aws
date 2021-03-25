// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Auto Scaling Attachment resource.
//
// > **NOTE on Auto Scaling Groups and ASG Attachments:** This provider currently provides
// both a standalone `autoscaling.Attachment` resource
// (describing an ASG attached to an ELB or ALB), and an `autoscaling.Group`
// with `loadBalancers` and `targetGroupArns` defined in-line. These two methods are not
// mutually-exclusive. If `autoscaling.Attachment` resources are used, either alone or with inline
// `loadBalancers` or `targetGroupArns`, the `autoscaling.Group` resource must be configured
// to [ignore changes](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to the `loadBalancers` and `targetGroupArns` arguments.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/autoscaling"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := autoscaling.NewAttachment(ctx, "asgAttachmentBar", &autoscaling.AttachmentArgs{
// 			AutoscalingGroupName: pulumi.Any(aws_autoscaling_group.Asg.Id),
// 			Elb:                  pulumi.Any(aws_elb.Bar.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/autoscaling"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := autoscaling.NewAttachment(ctx, "asgAttachmentBar", &autoscaling.AttachmentArgs{
// 			AutoscalingGroupName: pulumi.Any(aws_autoscaling_group.Asg.Id),
// 			AlbTargetGroupArn:    pulumi.Any(aws_alb_target_group.Test.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## With An AutoScaling Group Resource
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/autoscaling"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		asg, err := autoscaling.NewGroup(ctx, "asg", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = autoscaling.NewAttachment(ctx, "asgAttachmentBar", &autoscaling.AttachmentArgs{
// 			AutoscalingGroupName: asg.ID(),
// 			Elb:                  pulumi.Any(aws_elb.Test.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Attachment struct {
	pulumi.CustomResourceState

	// The ARN of an ALB Target Group.
	AlbTargetGroupArn pulumi.StringPtrOutput `pulumi:"albTargetGroupArn"`
	// Name of ASG to associate with the ELB.
	AutoscalingGroupName pulumi.StringOutput `pulumi:"autoscalingGroupName"`
	// The name of the ELB.
	Elb pulumi.StringPtrOutput `pulumi:"elb"`
}

// NewAttachment registers a new resource with the given unique name, arguments, and options.
func NewAttachment(ctx *pulumi.Context,
	name string, args *AttachmentArgs, opts ...pulumi.ResourceOption) (*Attachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoscalingGroupName == nil {
		return nil, errors.New("invalid value for required argument 'AutoscalingGroupName'")
	}
	var resource Attachment
	err := ctx.RegisterResource("aws:autoscaling/attachment:Attachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttachment gets an existing Attachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachmentState, opts ...pulumi.ResourceOption) (*Attachment, error) {
	var resource Attachment
	err := ctx.ReadResource("aws:autoscaling/attachment:Attachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Attachment resources.
type attachmentState struct {
	// The ARN of an ALB Target Group.
	AlbTargetGroupArn *string `pulumi:"albTargetGroupArn"`
	// Name of ASG to associate with the ELB.
	AutoscalingGroupName *string `pulumi:"autoscalingGroupName"`
	// The name of the ELB.
	Elb *string `pulumi:"elb"`
}

type AttachmentState struct {
	// The ARN of an ALB Target Group.
	AlbTargetGroupArn pulumi.StringPtrInput
	// Name of ASG to associate with the ELB.
	AutoscalingGroupName pulumi.StringPtrInput
	// The name of the ELB.
	Elb pulumi.StringPtrInput
}

func (AttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachmentState)(nil)).Elem()
}

type attachmentArgs struct {
	// The ARN of an ALB Target Group.
	AlbTargetGroupArn *string `pulumi:"albTargetGroupArn"`
	// Name of ASG to associate with the ELB.
	AutoscalingGroupName string `pulumi:"autoscalingGroupName"`
	// The name of the ELB.
	Elb *string `pulumi:"elb"`
}

// The set of arguments for constructing a Attachment resource.
type AttachmentArgs struct {
	// The ARN of an ALB Target Group.
	AlbTargetGroupArn pulumi.StringPtrInput
	// Name of ASG to associate with the ELB.
	AutoscalingGroupName pulumi.StringInput
	// The name of the ELB.
	Elb pulumi.StringPtrInput
}

func (AttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachmentArgs)(nil)).Elem()
}

type AttachmentInput interface {
	pulumi.Input

	ToAttachmentOutput() AttachmentOutput
	ToAttachmentOutputWithContext(ctx context.Context) AttachmentOutput
}

func (*Attachment) ElementType() reflect.Type {
	return reflect.TypeOf((*Attachment)(nil))
}

func (i *Attachment) ToAttachmentOutput() AttachmentOutput {
	return i.ToAttachmentOutputWithContext(context.Background())
}

func (i *Attachment) ToAttachmentOutputWithContext(ctx context.Context) AttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachmentOutput)
}

func (i *Attachment) ToAttachmentPtrOutput() AttachmentPtrOutput {
	return i.ToAttachmentPtrOutputWithContext(context.Background())
}

func (i *Attachment) ToAttachmentPtrOutputWithContext(ctx context.Context) AttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachmentPtrOutput)
}

type AttachmentPtrInput interface {
	pulumi.Input

	ToAttachmentPtrOutput() AttachmentPtrOutput
	ToAttachmentPtrOutputWithContext(ctx context.Context) AttachmentPtrOutput
}

type attachmentPtrType AttachmentArgs

func (*attachmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Attachment)(nil))
}

func (i *attachmentPtrType) ToAttachmentPtrOutput() AttachmentPtrOutput {
	return i.ToAttachmentPtrOutputWithContext(context.Background())
}

func (i *attachmentPtrType) ToAttachmentPtrOutputWithContext(ctx context.Context) AttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachmentPtrOutput)
}

// AttachmentArrayInput is an input type that accepts AttachmentArray and AttachmentArrayOutput values.
// You can construct a concrete instance of `AttachmentArrayInput` via:
//
//          AttachmentArray{ AttachmentArgs{...} }
type AttachmentArrayInput interface {
	pulumi.Input

	ToAttachmentArrayOutput() AttachmentArrayOutput
	ToAttachmentArrayOutputWithContext(context.Context) AttachmentArrayOutput
}

type AttachmentArray []AttachmentInput

func (AttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Attachment)(nil))
}

func (i AttachmentArray) ToAttachmentArrayOutput() AttachmentArrayOutput {
	return i.ToAttachmentArrayOutputWithContext(context.Background())
}

func (i AttachmentArray) ToAttachmentArrayOutputWithContext(ctx context.Context) AttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachmentArrayOutput)
}

// AttachmentMapInput is an input type that accepts AttachmentMap and AttachmentMapOutput values.
// You can construct a concrete instance of `AttachmentMapInput` via:
//
//          AttachmentMap{ "key": AttachmentArgs{...} }
type AttachmentMapInput interface {
	pulumi.Input

	ToAttachmentMapOutput() AttachmentMapOutput
	ToAttachmentMapOutputWithContext(context.Context) AttachmentMapOutput
}

type AttachmentMap map[string]AttachmentInput

func (AttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Attachment)(nil))
}

func (i AttachmentMap) ToAttachmentMapOutput() AttachmentMapOutput {
	return i.ToAttachmentMapOutputWithContext(context.Background())
}

func (i AttachmentMap) ToAttachmentMapOutputWithContext(ctx context.Context) AttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachmentMapOutput)
}

type AttachmentOutput struct {
	*pulumi.OutputState
}

func (AttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Attachment)(nil))
}

func (o AttachmentOutput) ToAttachmentOutput() AttachmentOutput {
	return o
}

func (o AttachmentOutput) ToAttachmentOutputWithContext(ctx context.Context) AttachmentOutput {
	return o
}

func (o AttachmentOutput) ToAttachmentPtrOutput() AttachmentPtrOutput {
	return o.ToAttachmentPtrOutputWithContext(context.Background())
}

func (o AttachmentOutput) ToAttachmentPtrOutputWithContext(ctx context.Context) AttachmentPtrOutput {
	return o.ApplyT(func(v Attachment) *Attachment {
		return &v
	}).(AttachmentPtrOutput)
}

type AttachmentPtrOutput struct {
	*pulumi.OutputState
}

func (AttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Attachment)(nil))
}

func (o AttachmentPtrOutput) ToAttachmentPtrOutput() AttachmentPtrOutput {
	return o
}

func (o AttachmentPtrOutput) ToAttachmentPtrOutputWithContext(ctx context.Context) AttachmentPtrOutput {
	return o
}

type AttachmentArrayOutput struct{ *pulumi.OutputState }

func (AttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Attachment)(nil))
}

func (o AttachmentArrayOutput) ToAttachmentArrayOutput() AttachmentArrayOutput {
	return o
}

func (o AttachmentArrayOutput) ToAttachmentArrayOutputWithContext(ctx context.Context) AttachmentArrayOutput {
	return o
}

func (o AttachmentArrayOutput) Index(i pulumi.IntInput) AttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Attachment {
		return vs[0].([]Attachment)[vs[1].(int)]
	}).(AttachmentOutput)
}

type AttachmentMapOutput struct{ *pulumi.OutputState }

func (AttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Attachment)(nil))
}

func (o AttachmentMapOutput) ToAttachmentMapOutput() AttachmentMapOutput {
	return o
}

func (o AttachmentMapOutput) ToAttachmentMapOutputWithContext(ctx context.Context) AttachmentMapOutput {
	return o
}

func (o AttachmentMapOutput) MapIndex(k pulumi.StringInput) AttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Attachment {
		return vs[0].(map[string]Attachment)[vs[1].(string)]
	}).(AttachmentOutput)
}

func init() {
	pulumi.RegisterOutputType(AttachmentOutput{})
	pulumi.RegisterOutputType(AttachmentPtrOutput{})
	pulumi.RegisterOutputType(AttachmentArrayOutput{})
	pulumi.RegisterOutputType(AttachmentMapOutput{})
}
