// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to create an EventBridge permission to support cross-account events in the current account default event bus.
//
// > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
//
// ## Example Usage
// ### Account Access
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudwatch.NewEventPermission(ctx, "devAccountAccess", &cloudwatch.EventPermissionArgs{
// 			Principal:   pulumi.String("123456789012"),
// 			StatementId: pulumi.String("DevAccountAccess"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Organization Access
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudwatch.NewEventPermission(ctx, "organizationAccess", &cloudwatch.EventPermissionArgs{
// 			Principal:   pulumi.String("*"),
// 			StatementId: pulumi.String("OrganizationAccess"),
// 			Condition: &cloudwatch.EventPermissionConditionArgs{
// 				Key:   pulumi.String("aws:PrincipalOrgID"),
// 				Type:  pulumi.String("StringEquals"),
// 				Value: pulumi.Any(aws_organizations_organization.Example.Id),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type EventPermission struct {
	pulumi.CustomResourceState

	// The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
	Action pulumi.StringPtrOutput `pulumi:"action"`
	// Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
	Condition EventPermissionConditionPtrOutput `pulumi:"condition"`
	// The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// An identifier string for the external account that you are granting permissions to.
	StatementId pulumi.StringOutput `pulumi:"statementId"`
}

// NewEventPermission registers a new resource with the given unique name, arguments, and options.
func NewEventPermission(ctx *pulumi.Context,
	name string, args *EventPermissionArgs, opts ...pulumi.ResourceOption) (*EventPermission, error) {
	if args == nil || args.Principal == nil {
		return nil, errors.New("missing required argument 'Principal'")
	}
	if args == nil || args.StatementId == nil {
		return nil, errors.New("missing required argument 'StatementId'")
	}
	if args == nil {
		args = &EventPermissionArgs{}
	}
	var resource EventPermission
	err := ctx.RegisterResource("aws:cloudwatch/eventPermission:EventPermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventPermission gets an existing EventPermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventPermissionState, opts ...pulumi.ResourceOption) (*EventPermission, error) {
	var resource EventPermission
	err := ctx.ReadResource("aws:cloudwatch/eventPermission:EventPermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventPermission resources.
type eventPermissionState struct {
	// The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
	Action *string `pulumi:"action"`
	// Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
	Condition *EventPermissionCondition `pulumi:"condition"`
	// The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
	Principal *string `pulumi:"principal"`
	// An identifier string for the external account that you are granting permissions to.
	StatementId *string `pulumi:"statementId"`
}

type EventPermissionState struct {
	// The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
	Action pulumi.StringPtrInput
	// Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
	Condition EventPermissionConditionPtrInput
	// The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
	Principal pulumi.StringPtrInput
	// An identifier string for the external account that you are granting permissions to.
	StatementId pulumi.StringPtrInput
}

func (EventPermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventPermissionState)(nil)).Elem()
}

type eventPermissionArgs struct {
	// The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
	Action *string `pulumi:"action"`
	// Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
	Condition *EventPermissionCondition `pulumi:"condition"`
	// The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
	Principal string `pulumi:"principal"`
	// An identifier string for the external account that you are granting permissions to.
	StatementId string `pulumi:"statementId"`
}

// The set of arguments for constructing a EventPermission resource.
type EventPermissionArgs struct {
	// The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
	Action pulumi.StringPtrInput
	// Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
	Condition EventPermissionConditionPtrInput
	// The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
	Principal pulumi.StringInput
	// An identifier string for the external account that you are granting permissions to.
	StatementId pulumi.StringInput
}

func (EventPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventPermissionArgs)(nil)).Elem()
}
