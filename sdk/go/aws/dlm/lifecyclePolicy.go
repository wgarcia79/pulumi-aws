// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package dlm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a [Data Lifecycle Manager (DLM) lifecycle policy](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshot-lifecycle.html) for managing snapshots.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dlm_lifecycle_policy.markdown.
type LifecyclePolicy struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the DLM Lifecycle Policy.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A description for the DLM lifecycle policy.
	Description pulumi.StringOutput `pulumi:"description"`
	// The ARN of an IAM role that is able to be assumed by the DLM service.
	ExecutionRoleArn pulumi.StringOutput `pulumi:"executionRoleArn"`
	// See the `policyDetails` configuration block. Max of 1.
	PolicyDetails LifecyclePolicyPolicyDetailsOutput `pulumi:"policyDetails"`
	// Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// Key-value mapping of resource tags.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewLifecyclePolicy registers a new resource with the given unique name, arguments, and options.
func NewLifecyclePolicy(ctx *pulumi.Context,
	name string, args *LifecyclePolicyArgs, opts ...pulumi.ResourceOption) (*LifecyclePolicy, error) {
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	if args == nil || args.ExecutionRoleArn == nil {
		return nil, errors.New("missing required argument 'ExecutionRoleArn'")
	}
	if args == nil || args.PolicyDetails == nil {
		return nil, errors.New("missing required argument 'PolicyDetails'")
	}
	if args == nil {
		args = &LifecyclePolicyArgs{}
	}
	var resource LifecyclePolicy
	err := ctx.RegisterResource("aws:dlm/lifecyclePolicy:LifecyclePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLifecyclePolicy gets an existing LifecyclePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLifecyclePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LifecyclePolicyState, opts ...pulumi.ResourceOption) (*LifecyclePolicy, error) {
	var resource LifecyclePolicy
	err := ctx.ReadResource("aws:dlm/lifecyclePolicy:LifecyclePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LifecyclePolicy resources.
type lifecyclePolicyState struct {
	// Amazon Resource Name (ARN) of the DLM Lifecycle Policy.
	Arn *string `pulumi:"arn"`
	// A description for the DLM lifecycle policy.
	Description *string `pulumi:"description"`
	// The ARN of an IAM role that is able to be assumed by the DLM service.
	ExecutionRoleArn *string `pulumi:"executionRoleArn"`
	// See the `policyDetails` configuration block. Max of 1.
	PolicyDetails *LifecyclePolicyPolicyDetails `pulumi:"policyDetails"`
	// Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
	State *string `pulumi:"state"`
	// Key-value mapping of resource tags.
	Tags map[string]interface{} `pulumi:"tags"`
}

type LifecyclePolicyState struct {
	// Amazon Resource Name (ARN) of the DLM Lifecycle Policy.
	Arn pulumi.StringPtrInput
	// A description for the DLM lifecycle policy.
	Description pulumi.StringPtrInput
	// The ARN of an IAM role that is able to be assumed by the DLM service.
	ExecutionRoleArn pulumi.StringPtrInput
	// See the `policyDetails` configuration block. Max of 1.
	PolicyDetails LifecyclePolicyPolicyDetailsPtrInput
	// Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
	State pulumi.StringPtrInput
	// Key-value mapping of resource tags.
	Tags pulumi.MapInput
}

func (LifecyclePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*lifecyclePolicyState)(nil)).Elem()
}

type lifecyclePolicyArgs struct {
	// A description for the DLM lifecycle policy.
	Description string `pulumi:"description"`
	// The ARN of an IAM role that is able to be assumed by the DLM service.
	ExecutionRoleArn string `pulumi:"executionRoleArn"`
	// See the `policyDetails` configuration block. Max of 1.
	PolicyDetails LifecyclePolicyPolicyDetails `pulumi:"policyDetails"`
	// Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
	State *string `pulumi:"state"`
	// Key-value mapping of resource tags.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a LifecyclePolicy resource.
type LifecyclePolicyArgs struct {
	// A description for the DLM lifecycle policy.
	Description pulumi.StringInput
	// The ARN of an IAM role that is able to be assumed by the DLM service.
	ExecutionRoleArn pulumi.StringInput
	// See the `policyDetails` configuration block. Max of 1.
	PolicyDetails LifecyclePolicyPolicyDetailsInput
	// Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
	State pulumi.StringPtrInput
	// Key-value mapping of resource tags.
	Tags pulumi.MapInput
}

func (LifecyclePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lifecyclePolicyArgs)(nil)).Elem()
}

