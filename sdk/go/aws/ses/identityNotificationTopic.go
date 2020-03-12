// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ses

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Resource for managing SES Identity Notification Topics
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_identity_notification_topic.markdown.
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
	if args == nil || args.Identity == nil {
		return nil, errors.New("missing required argument 'Identity'")
	}
	if args == nil || args.NotificationType == nil {
		return nil, errors.New("missing required argument 'NotificationType'")
	}
	if args == nil {
		args = &IdentityNotificationTopicArgs{}
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

