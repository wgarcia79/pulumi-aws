// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing SES Identity Notification Topics
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ses"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ses.NewIdentityNotificationTopic(ctx, "test", &ses.IdentityNotificationTopicArgs{
// 			TopicArn:               pulumi.Any(aws_sns_topic.Example.Arn),
// 			NotificationType:       pulumi.String("Bounce"),
// 			Identity:               pulumi.Any(aws_ses_domain_identity.Example.Domain),
// 			IncludeOriginalHeaders: pulumi.Bool(true),
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
// Identity Notification Topics can be imported using ID of the record. The ID is made up as IDENTITY|TYPE where IDENTITY is the SES Identity and TYPE is the Notification Type.
//
// ```sh
//  $ pulumi import aws:ses/identityNotificationTopic:IdentityNotificationTopic test 'example.com|Bounce'
// ```
type IdentityNotificationTopic struct {
	pulumi.CustomResourceState

	// The identity for which the Amazon SNS topic will be set. You can specify an identity by using its name or by using its Amazon Resource Name (ARN).
	Identity pulumi.StringOutput `pulumi:"identity"`
	// Whether SES should include original email headers in SNS notifications of this type. *false* by default.
	IncludeOriginalHeaders pulumi.BoolPtrOutput `pulumi:"includeOriginalHeaders"`
	// The type of notifications that will be published to the specified Amazon SNS topic. Valid Values: *Bounce*, *Complaint* or *Delivery*.
	NotificationType pulumi.StringOutput `pulumi:"notificationType"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic. Can be set to "" (an empty string) to disable publishing.
	TopicArn pulumi.StringPtrOutput `pulumi:"topicArn"`
}

// NewIdentityNotificationTopic registers a new resource with the given unique name, arguments, and options.
func NewIdentityNotificationTopic(ctx *pulumi.Context,
	name string, args *IdentityNotificationTopicArgs, opts ...pulumi.ResourceOption) (*IdentityNotificationTopic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
	}
	if args.NotificationType == nil {
		return nil, errors.New("invalid value for required argument 'NotificationType'")
	}
	var resource IdentityNotificationTopic
	err := ctx.RegisterResource("aws:ses/identityNotificationTopic:IdentityNotificationTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityNotificationTopic gets an existing IdentityNotificationTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityNotificationTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityNotificationTopicState, opts ...pulumi.ResourceOption) (*IdentityNotificationTopic, error) {
	var resource IdentityNotificationTopic
	err := ctx.ReadResource("aws:ses/identityNotificationTopic:IdentityNotificationTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityNotificationTopic resources.
type identityNotificationTopicState struct {
	// The identity for which the Amazon SNS topic will be set. You can specify an identity by using its name or by using its Amazon Resource Name (ARN).
	Identity *string `pulumi:"identity"`
	// Whether SES should include original email headers in SNS notifications of this type. *false* by default.
	IncludeOriginalHeaders *bool `pulumi:"includeOriginalHeaders"`
	// The type of notifications that will be published to the specified Amazon SNS topic. Valid Values: *Bounce*, *Complaint* or *Delivery*.
	NotificationType *string `pulumi:"notificationType"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic. Can be set to "" (an empty string) to disable publishing.
	TopicArn *string `pulumi:"topicArn"`
}

type IdentityNotificationTopicState struct {
	// The identity for which the Amazon SNS topic will be set. You can specify an identity by using its name or by using its Amazon Resource Name (ARN).
	Identity pulumi.StringPtrInput
	// Whether SES should include original email headers in SNS notifications of this type. *false* by default.
	IncludeOriginalHeaders pulumi.BoolPtrInput
	// The type of notifications that will be published to the specified Amazon SNS topic. Valid Values: *Bounce*, *Complaint* or *Delivery*.
	NotificationType pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the Amazon SNS topic. Can be set to "" (an empty string) to disable publishing.
	TopicArn pulumi.StringPtrInput
}

func (IdentityNotificationTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityNotificationTopicState)(nil)).Elem()
}

type identityNotificationTopicArgs struct {
	// The identity for which the Amazon SNS topic will be set. You can specify an identity by using its name or by using its Amazon Resource Name (ARN).
	Identity string `pulumi:"identity"`
	// Whether SES should include original email headers in SNS notifications of this type. *false* by default.
	IncludeOriginalHeaders *bool `pulumi:"includeOriginalHeaders"`
	// The type of notifications that will be published to the specified Amazon SNS topic. Valid Values: *Bounce*, *Complaint* or *Delivery*.
	NotificationType string `pulumi:"notificationType"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic. Can be set to "" (an empty string) to disable publishing.
	TopicArn *string `pulumi:"topicArn"`
}

// The set of arguments for constructing a IdentityNotificationTopic resource.
type IdentityNotificationTopicArgs struct {
	// The identity for which the Amazon SNS topic will be set. You can specify an identity by using its name or by using its Amazon Resource Name (ARN).
	Identity pulumi.StringInput
	// Whether SES should include original email headers in SNS notifications of this type. *false* by default.
	IncludeOriginalHeaders pulumi.BoolPtrInput
	// The type of notifications that will be published to the specified Amazon SNS topic. Valid Values: *Bounce*, *Complaint* or *Delivery*.
	NotificationType pulumi.StringInput
	// The Amazon Resource Name (ARN) of the Amazon SNS topic. Can be set to "" (an empty string) to disable publishing.
	TopicArn pulumi.StringPtrInput
}

func (IdentityNotificationTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityNotificationTopicArgs)(nil)).Elem()
}

type IdentityNotificationTopicInput interface {
	pulumi.Input

	ToIdentityNotificationTopicOutput() IdentityNotificationTopicOutput
	ToIdentityNotificationTopicOutputWithContext(ctx context.Context) IdentityNotificationTopicOutput
}

func (*IdentityNotificationTopic) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityNotificationTopic)(nil))
}

func (i *IdentityNotificationTopic) ToIdentityNotificationTopicOutput() IdentityNotificationTopicOutput {
	return i.ToIdentityNotificationTopicOutputWithContext(context.Background())
}

func (i *IdentityNotificationTopic) ToIdentityNotificationTopicOutputWithContext(ctx context.Context) IdentityNotificationTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityNotificationTopicOutput)
}

func (i *IdentityNotificationTopic) ToIdentityNotificationTopicPtrOutput() IdentityNotificationTopicPtrOutput {
	return i.ToIdentityNotificationTopicPtrOutputWithContext(context.Background())
}

func (i *IdentityNotificationTopic) ToIdentityNotificationTopicPtrOutputWithContext(ctx context.Context) IdentityNotificationTopicPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityNotificationTopicPtrOutput)
}

type IdentityNotificationTopicPtrInput interface {
	pulumi.Input

	ToIdentityNotificationTopicPtrOutput() IdentityNotificationTopicPtrOutput
	ToIdentityNotificationTopicPtrOutputWithContext(ctx context.Context) IdentityNotificationTopicPtrOutput
}

type identityNotificationTopicPtrType IdentityNotificationTopicArgs

func (*identityNotificationTopicPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityNotificationTopic)(nil))
}

func (i *identityNotificationTopicPtrType) ToIdentityNotificationTopicPtrOutput() IdentityNotificationTopicPtrOutput {
	return i.ToIdentityNotificationTopicPtrOutputWithContext(context.Background())
}

func (i *identityNotificationTopicPtrType) ToIdentityNotificationTopicPtrOutputWithContext(ctx context.Context) IdentityNotificationTopicPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityNotificationTopicPtrOutput)
}

// IdentityNotificationTopicArrayInput is an input type that accepts IdentityNotificationTopicArray and IdentityNotificationTopicArrayOutput values.
// You can construct a concrete instance of `IdentityNotificationTopicArrayInput` via:
//
//          IdentityNotificationTopicArray{ IdentityNotificationTopicArgs{...} }
type IdentityNotificationTopicArrayInput interface {
	pulumi.Input

	ToIdentityNotificationTopicArrayOutput() IdentityNotificationTopicArrayOutput
	ToIdentityNotificationTopicArrayOutputWithContext(context.Context) IdentityNotificationTopicArrayOutput
}

type IdentityNotificationTopicArray []IdentityNotificationTopicInput

func (IdentityNotificationTopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdentityNotificationTopic)(nil)).Elem()
}

func (i IdentityNotificationTopicArray) ToIdentityNotificationTopicArrayOutput() IdentityNotificationTopicArrayOutput {
	return i.ToIdentityNotificationTopicArrayOutputWithContext(context.Background())
}

func (i IdentityNotificationTopicArray) ToIdentityNotificationTopicArrayOutputWithContext(ctx context.Context) IdentityNotificationTopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityNotificationTopicArrayOutput)
}

// IdentityNotificationTopicMapInput is an input type that accepts IdentityNotificationTopicMap and IdentityNotificationTopicMapOutput values.
// You can construct a concrete instance of `IdentityNotificationTopicMapInput` via:
//
//          IdentityNotificationTopicMap{ "key": IdentityNotificationTopicArgs{...} }
type IdentityNotificationTopicMapInput interface {
	pulumi.Input

	ToIdentityNotificationTopicMapOutput() IdentityNotificationTopicMapOutput
	ToIdentityNotificationTopicMapOutputWithContext(context.Context) IdentityNotificationTopicMapOutput
}

type IdentityNotificationTopicMap map[string]IdentityNotificationTopicInput

func (IdentityNotificationTopicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdentityNotificationTopic)(nil)).Elem()
}

func (i IdentityNotificationTopicMap) ToIdentityNotificationTopicMapOutput() IdentityNotificationTopicMapOutput {
	return i.ToIdentityNotificationTopicMapOutputWithContext(context.Background())
}

func (i IdentityNotificationTopicMap) ToIdentityNotificationTopicMapOutputWithContext(ctx context.Context) IdentityNotificationTopicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityNotificationTopicMapOutput)
}

type IdentityNotificationTopicOutput struct {
	*pulumi.OutputState
}

func (IdentityNotificationTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityNotificationTopic)(nil))
}

func (o IdentityNotificationTopicOutput) ToIdentityNotificationTopicOutput() IdentityNotificationTopicOutput {
	return o
}

func (o IdentityNotificationTopicOutput) ToIdentityNotificationTopicOutputWithContext(ctx context.Context) IdentityNotificationTopicOutput {
	return o
}

func (o IdentityNotificationTopicOutput) ToIdentityNotificationTopicPtrOutput() IdentityNotificationTopicPtrOutput {
	return o.ToIdentityNotificationTopicPtrOutputWithContext(context.Background())
}

func (o IdentityNotificationTopicOutput) ToIdentityNotificationTopicPtrOutputWithContext(ctx context.Context) IdentityNotificationTopicPtrOutput {
	return o.ApplyT(func(v IdentityNotificationTopic) *IdentityNotificationTopic {
		return &v
	}).(IdentityNotificationTopicPtrOutput)
}

type IdentityNotificationTopicPtrOutput struct {
	*pulumi.OutputState
}

func (IdentityNotificationTopicPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityNotificationTopic)(nil))
}

func (o IdentityNotificationTopicPtrOutput) ToIdentityNotificationTopicPtrOutput() IdentityNotificationTopicPtrOutput {
	return o
}

func (o IdentityNotificationTopicPtrOutput) ToIdentityNotificationTopicPtrOutputWithContext(ctx context.Context) IdentityNotificationTopicPtrOutput {
	return o
}

type IdentityNotificationTopicArrayOutput struct{ *pulumi.OutputState }

func (IdentityNotificationTopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentityNotificationTopic)(nil))
}

func (o IdentityNotificationTopicArrayOutput) ToIdentityNotificationTopicArrayOutput() IdentityNotificationTopicArrayOutput {
	return o
}

func (o IdentityNotificationTopicArrayOutput) ToIdentityNotificationTopicArrayOutputWithContext(ctx context.Context) IdentityNotificationTopicArrayOutput {
	return o
}

func (o IdentityNotificationTopicArrayOutput) Index(i pulumi.IntInput) IdentityNotificationTopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IdentityNotificationTopic {
		return vs[0].([]IdentityNotificationTopic)[vs[1].(int)]
	}).(IdentityNotificationTopicOutput)
}

type IdentityNotificationTopicMapOutput struct{ *pulumi.OutputState }

func (IdentityNotificationTopicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IdentityNotificationTopic)(nil))
}

func (o IdentityNotificationTopicMapOutput) ToIdentityNotificationTopicMapOutput() IdentityNotificationTopicMapOutput {
	return o
}

func (o IdentityNotificationTopicMapOutput) ToIdentityNotificationTopicMapOutputWithContext(ctx context.Context) IdentityNotificationTopicMapOutput {
	return o
}

func (o IdentityNotificationTopicMapOutput) MapIndex(k pulumi.StringInput) IdentityNotificationTopicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IdentityNotificationTopic {
		return vs[0].(map[string]IdentityNotificationTopic)[vs[1].(string)]
	}).(IdentityNotificationTopicOutput)
}

func init() {
	pulumi.RegisterOutputType(IdentityNotificationTopicOutput{})
	pulumi.RegisterOutputType(IdentityNotificationTopicPtrOutput{})
	pulumi.RegisterOutputType(IdentityNotificationTopicArrayOutput{})
	pulumi.RegisterOutputType(IdentityNotificationTopicMapOutput{})
}
