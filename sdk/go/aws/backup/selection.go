// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages selection conditions for AWS Backup plan resources.
//
// ## Example Usage
// ### IAM Role
//
// > For more information about creating and managing IAM Roles for backups and restores, see the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/iam-service-roles.html).
//
// The below example creates an IAM role with the default managed IAM Policy for allowing AWS Backup to create backups.
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/backup"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.Any(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\"sts:AssumeRole\"],\n", "      \"Effect\": \"allow\",\n", "      \"Principal\": {\n", "        \"Service\": [\"backup.amazonaws.com\"]\n", "      }\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicyAttachment(ctx, "exampleRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
// 			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWSBackupServiceRolePolicyForBackup"),
// 			Role:      exampleRole.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = backup.NewSelection(ctx, "exampleSelection", &backup.SelectionArgs{
// 			IamRoleArn: exampleRole.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Selecting Backups By Tag
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/backup"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := backup.NewSelection(ctx, "example", &backup.SelectionArgs{
// 			IamRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
// 			PlanId:     pulumi.Any(aws_backup_plan.Example.Id),
// 			SelectionTags: backup.SelectionSelectionTagArray{
// 				&backup.SelectionSelectionTagArgs{
// 					Type:  pulumi.String("STRINGEQUALS"),
// 					Key:   pulumi.String("foo"),
// 					Value: pulumi.String("bar"),
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
// ### Selecting Backups By Resource
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/backup"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := backup.NewSelection(ctx, "example", &backup.SelectionArgs{
// 			IamRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
// 			PlanId:     pulumi.Any(aws_backup_plan.Example.Id),
// 			Resources: pulumi.StringArray{
// 				pulumi.Any(aws_db_instance.Example.Arn),
// 				pulumi.Any(aws_ebs_volume.Example.Arn),
// 				pulumi.Any(aws_efs_file_system.Example.Arn),
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
// Backup selection can be imported using the role plan_id and id separated by `|`.
//
// ```sh
//  $ pulumi import aws:backup/selection:Selection example plan-id|selection-id
// ```
type Selection struct {
	pulumi.CustomResourceState

	// The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
	IamRoleArn pulumi.StringOutput `pulumi:"iamRoleArn"`
	// The display name of a resource selection document.
	Name pulumi.StringOutput `pulumi:"name"`
	// The backup plan ID to be associated with the selection of resources.
	PlanId pulumi.StringOutput `pulumi:"planId"`
	// An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan..
	Resources pulumi.StringArrayOutput `pulumi:"resources"`
	// Tag-based conditions used to specify a set of resources to assign to a backup plan.
	SelectionTags SelectionSelectionTagArrayOutput `pulumi:"selectionTags"`
}

// NewSelection registers a new resource with the given unique name, arguments, and options.
func NewSelection(ctx *pulumi.Context,
	name string, args *SelectionArgs, opts ...pulumi.ResourceOption) (*Selection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IamRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'IamRoleArn'")
	}
	if args.PlanId == nil {
		return nil, errors.New("invalid value for required argument 'PlanId'")
	}
	var resource Selection
	err := ctx.RegisterResource("aws:backup/selection:Selection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSelection gets an existing Selection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSelection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SelectionState, opts ...pulumi.ResourceOption) (*Selection, error) {
	var resource Selection
	err := ctx.ReadResource("aws:backup/selection:Selection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Selection resources.
type selectionState struct {
	// The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
	IamRoleArn *string `pulumi:"iamRoleArn"`
	// The display name of a resource selection document.
	Name *string `pulumi:"name"`
	// The backup plan ID to be associated with the selection of resources.
	PlanId *string `pulumi:"planId"`
	// An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan..
	Resources []string `pulumi:"resources"`
	// Tag-based conditions used to specify a set of resources to assign to a backup plan.
	SelectionTags []SelectionSelectionTag `pulumi:"selectionTags"`
}

type SelectionState struct {
	// The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
	IamRoleArn pulumi.StringPtrInput
	// The display name of a resource selection document.
	Name pulumi.StringPtrInput
	// The backup plan ID to be associated with the selection of resources.
	PlanId pulumi.StringPtrInput
	// An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan..
	Resources pulumi.StringArrayInput
	// Tag-based conditions used to specify a set of resources to assign to a backup plan.
	SelectionTags SelectionSelectionTagArrayInput
}

func (SelectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*selectionState)(nil)).Elem()
}

type selectionArgs struct {
	// The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
	IamRoleArn string `pulumi:"iamRoleArn"`
	// The display name of a resource selection document.
	Name *string `pulumi:"name"`
	// The backup plan ID to be associated with the selection of resources.
	PlanId string `pulumi:"planId"`
	// An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan..
	Resources []string `pulumi:"resources"`
	// Tag-based conditions used to specify a set of resources to assign to a backup plan.
	SelectionTags []SelectionSelectionTag `pulumi:"selectionTags"`
}

// The set of arguments for constructing a Selection resource.
type SelectionArgs struct {
	// The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
	IamRoleArn pulumi.StringInput
	// The display name of a resource selection document.
	Name pulumi.StringPtrInput
	// The backup plan ID to be associated with the selection of resources.
	PlanId pulumi.StringInput
	// An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan..
	Resources pulumi.StringArrayInput
	// Tag-based conditions used to specify a set of resources to assign to a backup plan.
	SelectionTags SelectionSelectionTagArrayInput
}

func (SelectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*selectionArgs)(nil)).Elem()
}

type SelectionInput interface {
	pulumi.Input

	ToSelectionOutput() SelectionOutput
	ToSelectionOutputWithContext(ctx context.Context) SelectionOutput
}

func (*Selection) ElementType() reflect.Type {
	return reflect.TypeOf((*Selection)(nil))
}

func (i *Selection) ToSelectionOutput() SelectionOutput {
	return i.ToSelectionOutputWithContext(context.Background())
}

func (i *Selection) ToSelectionOutputWithContext(ctx context.Context) SelectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectionOutput)
}

func (i *Selection) ToSelectionPtrOutput() SelectionPtrOutput {
	return i.ToSelectionPtrOutputWithContext(context.Background())
}

func (i *Selection) ToSelectionPtrOutputWithContext(ctx context.Context) SelectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectionPtrOutput)
}

type SelectionPtrInput interface {
	pulumi.Input

	ToSelectionPtrOutput() SelectionPtrOutput
	ToSelectionPtrOutputWithContext(ctx context.Context) SelectionPtrOutput
}

type selectionPtrType SelectionArgs

func (*selectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Selection)(nil))
}

func (i *selectionPtrType) ToSelectionPtrOutput() SelectionPtrOutput {
	return i.ToSelectionPtrOutputWithContext(context.Background())
}

func (i *selectionPtrType) ToSelectionPtrOutputWithContext(ctx context.Context) SelectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectionPtrOutput)
}

// SelectionArrayInput is an input type that accepts SelectionArray and SelectionArrayOutput values.
// You can construct a concrete instance of `SelectionArrayInput` via:
//
//          SelectionArray{ SelectionArgs{...} }
type SelectionArrayInput interface {
	pulumi.Input

	ToSelectionArrayOutput() SelectionArrayOutput
	ToSelectionArrayOutputWithContext(context.Context) SelectionArrayOutput
}

type SelectionArray []SelectionInput

func (SelectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Selection)(nil)).Elem()
}

func (i SelectionArray) ToSelectionArrayOutput() SelectionArrayOutput {
	return i.ToSelectionArrayOutputWithContext(context.Background())
}

func (i SelectionArray) ToSelectionArrayOutputWithContext(ctx context.Context) SelectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectionArrayOutput)
}

// SelectionMapInput is an input type that accepts SelectionMap and SelectionMapOutput values.
// You can construct a concrete instance of `SelectionMapInput` via:
//
//          SelectionMap{ "key": SelectionArgs{...} }
type SelectionMapInput interface {
	pulumi.Input

	ToSelectionMapOutput() SelectionMapOutput
	ToSelectionMapOutputWithContext(context.Context) SelectionMapOutput
}

type SelectionMap map[string]SelectionInput

func (SelectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Selection)(nil)).Elem()
}

func (i SelectionMap) ToSelectionMapOutput() SelectionMapOutput {
	return i.ToSelectionMapOutputWithContext(context.Background())
}

func (i SelectionMap) ToSelectionMapOutputWithContext(ctx context.Context) SelectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectionMapOutput)
}

type SelectionOutput struct{ *pulumi.OutputState }

func (SelectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Selection)(nil))
}

func (o SelectionOutput) ToSelectionOutput() SelectionOutput {
	return o
}

func (o SelectionOutput) ToSelectionOutputWithContext(ctx context.Context) SelectionOutput {
	return o
}

func (o SelectionOutput) ToSelectionPtrOutput() SelectionPtrOutput {
	return o.ToSelectionPtrOutputWithContext(context.Background())
}

func (o SelectionOutput) ToSelectionPtrOutputWithContext(ctx context.Context) SelectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Selection) *Selection {
		return &v
	}).(SelectionPtrOutput)
}

type SelectionPtrOutput struct{ *pulumi.OutputState }

func (SelectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Selection)(nil))
}

func (o SelectionPtrOutput) ToSelectionPtrOutput() SelectionPtrOutput {
	return o
}

func (o SelectionPtrOutput) ToSelectionPtrOutputWithContext(ctx context.Context) SelectionPtrOutput {
	return o
}

func (o SelectionPtrOutput) Elem() SelectionOutput {
	return o.ApplyT(func(v *Selection) Selection {
		if v != nil {
			return *v
		}
		var ret Selection
		return ret
	}).(SelectionOutput)
}

type SelectionArrayOutput struct{ *pulumi.OutputState }

func (SelectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Selection)(nil))
}

func (o SelectionArrayOutput) ToSelectionArrayOutput() SelectionArrayOutput {
	return o
}

func (o SelectionArrayOutput) ToSelectionArrayOutputWithContext(ctx context.Context) SelectionArrayOutput {
	return o
}

func (o SelectionArrayOutput) Index(i pulumi.IntInput) SelectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Selection {
		return vs[0].([]Selection)[vs[1].(int)]
	}).(SelectionOutput)
}

type SelectionMapOutput struct{ *pulumi.OutputState }

func (SelectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Selection)(nil))
}

func (o SelectionMapOutput) ToSelectionMapOutput() SelectionMapOutput {
	return o
}

func (o SelectionMapOutput) ToSelectionMapOutputWithContext(ctx context.Context) SelectionMapOutput {
	return o
}

func (o SelectionMapOutput) MapIndex(k pulumi.StringInput) SelectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Selection {
		return vs[0].(map[string]Selection)[vs[1].(string)]
	}).(SelectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SelectionInput)(nil)).Elem(), &Selection{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelectionPtrInput)(nil)).Elem(), &Selection{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelectionArrayInput)(nil)).Elem(), SelectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelectionMapInput)(nil)).Elem(), SelectionMap{})
	pulumi.RegisterOutputType(SelectionOutput{})
	pulumi.RegisterOutputType(SelectionPtrOutput{})
	pulumi.RegisterOutputType(SelectionArrayOutput{})
	pulumi.RegisterOutputType(SelectionMapOutput{})
}
