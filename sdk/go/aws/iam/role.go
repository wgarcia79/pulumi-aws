// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an IAM role.
//
// > **NOTE:** If policies are attached to the role via the `iam.PolicyAttachment` resource and you are modifying the role `name` or `path`, the `forceDetachPolicies` argument must be set to `true` and applied before attempting the operation otherwise you will encounter a `DeleteConflict` error. The `iam.RolePolicyAttachment` resource does not have this requirement.
//
// > **NOTE:** If you use this resource's `managedPolicyArns` argument or `inlinePolicy` configuration blocks, this resource will take over exclusive management of the role's respective policy types (e.g., both policy types if both arguments are used). These arguments are incompatible with other ways of managing a role's policies, such as `iam.PolicyAttachment`, `iam.RolePolicyAttachment`, and `iam.RolePolicy`. If you attempt to manage a role's policies by multiple means, you will get resource cycling and/or errors.
//
// ## Example Usage
// ### Basic Example
//
// ```go
// package main
//
// import (
// 	"encoding/json"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		tmpJSON0, err := json.Marshal(map[string]interface{}{
// 			"Version": "2012-10-17",
// 			"Statement": []map[string]interface{}{
// 				map[string]interface{}{
// 					"Action": "sts:AssumeRole",
// 					"Effect": "Allow",
// 					"Sid":    "",
// 					"Principal": map[string]interface{}{
// 						"Service": "ec2.amazonaws.com",
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json0 := string(tmpJSON0)
// 		_, err := iam.NewRole(ctx, "testRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(json0),
// 			Tags: pulumi.StringMap{
// 				"tag-key": pulumi.String("tag-value"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Example of Using Data Source for Assume Role Policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		instance_assume_role_policy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
// 			Statements: []iam.GetPolicyDocumentStatement{
// 				iam.GetPolicyDocumentStatement{
// 					Actions: []string{
// 						"sts:AssumeRole",
// 					},
// 					Principals: []iam.GetPolicyDocumentStatementPrincipal{
// 						iam.GetPolicyDocumentStatementPrincipal{
// 							Type: "Service",
// 							Identifiers: []string{
// 								"ec2.amazonaws.com",
// 							},
// 						},
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRole(ctx, "instance", &iam.RoleArgs{
// 			Path:             pulumi.String("/system/"),
// 			AssumeRolePolicy: pulumi.String(instance_assume_role_policy.Json),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Example of Removing Inline Policies
//
// This example creates an IAM role with what appears to be empty IAM `inlinePolicy` argument instead of using `inlinePolicy` as a configuration block. The result is that if someone were to add an inline policy out-of-band, on the next apply, the provider will remove that policy.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iam.NewRole(ctx, "example", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.Any(data.Aws_iam_policy_document.Instance_assume_role_policy.Json),
// 			InlinePolicies: iam.RoleInlinePolicyArray{
// 				nil,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Example of Exclusive Managed Policies
//
// This example creates an IAM role and attaches two managed IAM policies. If someone attaches another managed policy out-of-band, on the next apply, the provider will detach that policy. If someone detaches these policies out-of-band, the provider will attach them again.
//
// ```go
// package main
//
// import (
// 	"encoding/json"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		tmpJSON0, err := json.Marshal(map[string]interface{}{
// 			"Version": "2012-10-17",
// 			"Statement": []map[string]interface{}{
// 				map[string]interface{}{
// 					"Action": []string{
// 						"ec2:Describe*",
// 					},
// 					"Effect":   "Allow",
// 					"Resource": "*",
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json0 := string(tmpJSON0)
// 		policyOne, err := iam.NewPolicy(ctx, "policyOne", &iam.PolicyArgs{
// 			Policy: pulumi.String(json0),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		tmpJSON1, err := json.Marshal(map[string]interface{}{
// 			"Version": "2012-10-17",
// 			"Statement": []map[string]interface{}{
// 				map[string]interface{}{
// 					"Action": []string{
// 						"s3:ListAllMyBuckets",
// 						"s3:ListBucket",
// 						"s3:HeadBucket",
// 					},
// 					"Effect":   "Allow",
// 					"Resource": "*",
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json1 := string(tmpJSON1)
// 		policyTwo, err := iam.NewPolicy(ctx, "policyTwo", &iam.PolicyArgs{
// 			Policy: pulumi.String(json1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRole(ctx, "example", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.Any(data.Aws_iam_policy_document.Instance_assume_role_policy.Json),
// 			ManagedPolicyArns: pulumi.StringArray{
// 				policyOne.Arn,
// 				policyTwo.Arn,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Example of Removing Managed Policies
//
// This example creates an IAM role with an empty `managedPolicyArns` argument. If someone attaches a policy out-of-band, on the next apply, the provider will detach that policy.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iam.NewRole(ctx, "example", &iam.RoleArgs{
// 			AssumeRolePolicy:  pulumi.Any(data.Aws_iam_policy_document.Instance_assume_role_policy.Json),
// 			ManagedPolicyArns: []interface{}{},
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
// IAM Roles can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:iam/role:Role developer developer_name
// ```
type Role struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) specifying the role.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Policy that grants an entity permission to assume the role.
	AssumeRolePolicy pulumi.StringOutput `pulumi:"assumeRolePolicy"`
	// Creation date of the IAM role.
	CreateDate pulumi.StringOutput `pulumi:"createDate"`
	// Description of the role.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
	ForceDetachPolicies pulumi.BoolPtrOutput `pulumi:"forceDetachPolicies"`
	// Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. Defined below. If no blocks are configured, the provider will ignore any managing any inline policies in this resource. Configuring one empty block (i.e., `inlinePolicy {}`) will cause the provider to remove _all_ inline policies.
	InlinePolicies RoleInlinePolicyArrayOutput `pulumi:"inlinePolicies"`
	// Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, the provider will ignore policy attachments to this resource. When configured, the provider will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., `managedPolicyArns = []`) will cause the provider to remove _all_ managed policy attachments.
	ManagedPolicyArns pulumi.StringArrayOutput `pulumi:"managedPolicyArns"`
	// Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	MaxSessionDuration pulumi.IntPtrOutput `pulumi:"maxSessionDuration"`
	// Name of the role policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// ARN of the policy that is used to set the permissions boundary for the role.
	PermissionsBoundary pulumi.StringPtrOutput `pulumi:"permissionsBoundary"`
	// Key-value mapping of tags for the IAM role
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Stable and unique string identifying the role.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssumeRolePolicy == nil {
		return nil, errors.New("invalid value for required argument 'AssumeRolePolicy'")
	}
	var resource Role
	err := ctx.RegisterResource("aws:iam/role:Role", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleState, opts ...pulumi.ResourceOption) (*Role, error) {
	var resource Role
	err := ctx.ReadResource("aws:iam/role:Role", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Role resources.
type roleState struct {
	// Amazon Resource Name (ARN) specifying the role.
	Arn *string `pulumi:"arn"`
	// Policy that grants an entity permission to assume the role.
	AssumeRolePolicy *string `pulumi:"assumeRolePolicy"`
	// Creation date of the IAM role.
	CreateDate *string `pulumi:"createDate"`
	// Description of the role.
	Description *string `pulumi:"description"`
	// Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
	ForceDetachPolicies *bool `pulumi:"forceDetachPolicies"`
	// Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. Defined below. If no blocks are configured, the provider will ignore any managing any inline policies in this resource. Configuring one empty block (i.e., `inlinePolicy {}`) will cause the provider to remove _all_ inline policies.
	InlinePolicies []RoleInlinePolicy `pulumi:"inlinePolicies"`
	// Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, the provider will ignore policy attachments to this resource. When configured, the provider will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., `managedPolicyArns = []`) will cause the provider to remove _all_ managed policy attachments.
	ManagedPolicyArns []string `pulumi:"managedPolicyArns"`
	// Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	MaxSessionDuration *int `pulumi:"maxSessionDuration"`
	// Name of the role policy.
	Name *string `pulumi:"name"`
	// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path *string `pulumi:"path"`
	// ARN of the policy that is used to set the permissions boundary for the role.
	PermissionsBoundary *string `pulumi:"permissionsBoundary"`
	// Key-value mapping of tags for the IAM role
	Tags map[string]string `pulumi:"tags"`
	// Stable and unique string identifying the role.
	UniqueId *string `pulumi:"uniqueId"`
}

type RoleState struct {
	// Amazon Resource Name (ARN) specifying the role.
	Arn pulumi.StringPtrInput
	// Policy that grants an entity permission to assume the role.
	AssumeRolePolicy pulumi.StringPtrInput
	// Creation date of the IAM role.
	CreateDate pulumi.StringPtrInput
	// Description of the role.
	Description pulumi.StringPtrInput
	// Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
	ForceDetachPolicies pulumi.BoolPtrInput
	// Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. Defined below. If no blocks are configured, the provider will ignore any managing any inline policies in this resource. Configuring one empty block (i.e., `inlinePolicy {}`) will cause the provider to remove _all_ inline policies.
	InlinePolicies RoleInlinePolicyArrayInput
	// Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, the provider will ignore policy attachments to this resource. When configured, the provider will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., `managedPolicyArns = []`) will cause the provider to remove _all_ managed policy attachments.
	ManagedPolicyArns pulumi.StringArrayInput
	// Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	MaxSessionDuration pulumi.IntPtrInput
	// Name of the role policy.
	Name pulumi.StringPtrInput
	// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path pulumi.StringPtrInput
	// ARN of the policy that is used to set the permissions boundary for the role.
	PermissionsBoundary pulumi.StringPtrInput
	// Key-value mapping of tags for the IAM role
	Tags pulumi.StringMapInput
	// Stable and unique string identifying the role.
	UniqueId pulumi.StringPtrInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	// Policy that grants an entity permission to assume the role.
	AssumeRolePolicy interface{} `pulumi:"assumeRolePolicy"`
	// Description of the role.
	Description *string `pulumi:"description"`
	// Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
	ForceDetachPolicies *bool `pulumi:"forceDetachPolicies"`
	// Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. Defined below. If no blocks are configured, the provider will ignore any managing any inline policies in this resource. Configuring one empty block (i.e., `inlinePolicy {}`) will cause the provider to remove _all_ inline policies.
	InlinePolicies []RoleInlinePolicy `pulumi:"inlinePolicies"`
	// Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, the provider will ignore policy attachments to this resource. When configured, the provider will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., `managedPolicyArns = []`) will cause the provider to remove _all_ managed policy attachments.
	ManagedPolicyArns []string `pulumi:"managedPolicyArns"`
	// Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	MaxSessionDuration *int `pulumi:"maxSessionDuration"`
	// Name of the role policy.
	Name *string `pulumi:"name"`
	// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path *string `pulumi:"path"`
	// ARN of the policy that is used to set the permissions boundary for the role.
	PermissionsBoundary *string `pulumi:"permissionsBoundary"`
	// Key-value mapping of tags for the IAM role
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	// Policy that grants an entity permission to assume the role.
	AssumeRolePolicy pulumi.Input
	// Description of the role.
	Description pulumi.StringPtrInput
	// Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
	ForceDetachPolicies pulumi.BoolPtrInput
	// Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. Defined below. If no blocks are configured, the provider will ignore any managing any inline policies in this resource. Configuring one empty block (i.e., `inlinePolicy {}`) will cause the provider to remove _all_ inline policies.
	InlinePolicies RoleInlinePolicyArrayInput
	// Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, the provider will ignore policy attachments to this resource. When configured, the provider will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., `managedPolicyArns = []`) will cause the provider to remove _all_ managed policy attachments.
	ManagedPolicyArns pulumi.StringArrayInput
	// Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	MaxSessionDuration pulumi.IntPtrInput
	// Name of the role policy.
	Name pulumi.StringPtrInput
	// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path pulumi.StringPtrInput
	// ARN of the policy that is used to set the permissions boundary for the role.
	PermissionsBoundary pulumi.StringPtrInput
	// Key-value mapping of tags for the IAM role
	Tags pulumi.StringMapInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}

type RoleInput interface {
	pulumi.Input

	ToRoleOutput() RoleOutput
	ToRoleOutputWithContext(ctx context.Context) RoleOutput
}

func (*Role) ElementType() reflect.Type {
	return reflect.TypeOf((*Role)(nil))
}

func (i *Role) ToRoleOutput() RoleOutput {
	return i.ToRoleOutputWithContext(context.Background())
}

func (i *Role) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleOutput)
}

func (i *Role) ToRolePtrOutput() RolePtrOutput {
	return i.ToRolePtrOutputWithContext(context.Background())
}

func (i *Role) ToRolePtrOutputWithContext(ctx context.Context) RolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolePtrOutput)
}

type RolePtrInput interface {
	pulumi.Input

	ToRolePtrOutput() RolePtrOutput
	ToRolePtrOutputWithContext(ctx context.Context) RolePtrOutput
}

type rolePtrType RoleArgs

func (*rolePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil))
}

func (i *rolePtrType) ToRolePtrOutput() RolePtrOutput {
	return i.ToRolePtrOutputWithContext(context.Background())
}

func (i *rolePtrType) ToRolePtrOutputWithContext(ctx context.Context) RolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolePtrOutput)
}

// RoleArrayInput is an input type that accepts RoleArray and RoleArrayOutput values.
// You can construct a concrete instance of `RoleArrayInput` via:
//
//          RoleArray{ RoleArgs{...} }
type RoleArrayInput interface {
	pulumi.Input

	ToRoleArrayOutput() RoleArrayOutput
	ToRoleArrayOutputWithContext(context.Context) RoleArrayOutput
}

type RoleArray []RoleInput

func (RoleArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Role)(nil))
}

func (i RoleArray) ToRoleArrayOutput() RoleArrayOutput {
	return i.ToRoleArrayOutputWithContext(context.Background())
}

func (i RoleArray) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleArrayOutput)
}

// RoleMapInput is an input type that accepts RoleMap and RoleMapOutput values.
// You can construct a concrete instance of `RoleMapInput` via:
//
//          RoleMap{ "key": RoleArgs{...} }
type RoleMapInput interface {
	pulumi.Input

	ToRoleMapOutput() RoleMapOutput
	ToRoleMapOutputWithContext(context.Context) RoleMapOutput
}

type RoleMap map[string]RoleInput

func (RoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Role)(nil))
}

func (i RoleMap) ToRoleMapOutput() RoleMapOutput {
	return i.ToRoleMapOutputWithContext(context.Background())
}

func (i RoleMap) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleMapOutput)
}

type RoleOutput struct {
	*pulumi.OutputState
}

func (RoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Role)(nil))
}

func (o RoleOutput) ToRoleOutput() RoleOutput {
	return o
}

func (o RoleOutput) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return o
}

func (o RoleOutput) ToRolePtrOutput() RolePtrOutput {
	return o.ToRolePtrOutputWithContext(context.Background())
}

func (o RoleOutput) ToRolePtrOutputWithContext(ctx context.Context) RolePtrOutput {
	return o.ApplyT(func(v Role) *Role {
		return &v
	}).(RolePtrOutput)
}

type RolePtrOutput struct {
	*pulumi.OutputState
}

func (RolePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil))
}

func (o RolePtrOutput) ToRolePtrOutput() RolePtrOutput {
	return o
}

func (o RolePtrOutput) ToRolePtrOutputWithContext(ctx context.Context) RolePtrOutput {
	return o
}

type RoleArrayOutput struct{ *pulumi.OutputState }

func (RoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Role)(nil))
}

func (o RoleArrayOutput) ToRoleArrayOutput() RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) Index(i pulumi.IntInput) RoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Role {
		return vs[0].([]Role)[vs[1].(int)]
	}).(RoleOutput)
}

type RoleMapOutput struct{ *pulumi.OutputState }

func (RoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Role)(nil))
}

func (o RoleMapOutput) ToRoleMapOutput() RoleMapOutput {
	return o
}

func (o RoleMapOutput) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return o
}

func (o RoleMapOutput) MapIndex(k pulumi.StringInput) RoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Role {
		return vs[0].(map[string]Role)[vs[1].(string)]
	}).(RoleOutput)
}

func init() {
	pulumi.RegisterOutputType(RoleOutput{})
	pulumi.RegisterOutputType(RolePtrOutput{})
	pulumi.RegisterOutputType(RoleArrayOutput{})
	pulumi.RegisterOutputType(RoleMapOutput{})
}
