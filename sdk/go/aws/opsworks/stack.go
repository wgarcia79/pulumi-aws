// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package opsworks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an OpsWorks stack resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/opsworks_stack.html.markdown.
type Stack struct {
	pulumi.CustomResourceState

	// If set to `"LATEST"`, OpsWorks will automatically install the latest version.
	AgentVersion pulumi.StringOutput `pulumi:"agentVersion"`
	Arn pulumi.StringOutput `pulumi:"arn"`
	// If `manageBerkshelf` is enabled, the version of Berkshelf to use.
	BerkshelfVersion pulumi.StringPtrOutput `pulumi:"berkshelfVersion"`
	// Color to paint next to the stack's resources in the OpsWorks console.
	Color pulumi.StringPtrOutput `pulumi:"color"`
	// Name of the configuration manager to use. Defaults to "Chef".
	ConfigurationManagerName pulumi.StringPtrOutput `pulumi:"configurationManagerName"`
	// Version of the configuration manager to use. Defaults to "11.4".
	ConfigurationManagerVersion pulumi.StringPtrOutput `pulumi:"configurationManagerVersion"`
	// When `useCustomCookbooks` is set, provide this sub-object as
	// described below.
	CustomCookbooksSources StackCustomCookbooksSourceArrayOutput `pulumi:"customCookbooksSources"`
	// Custom JSON attributes to apply to the entire stack.
	CustomJson pulumi.StringPtrOutput `pulumi:"customJson"`
	// Name of the availability zone where instances will be created
	// by default. This is required unless you set `vpcId`.
	DefaultAvailabilityZone pulumi.StringOutput `pulumi:"defaultAvailabilityZone"`
	// The ARN of an IAM Instance Profile that created instances
	// will have by default.
	DefaultInstanceProfileArn pulumi.StringOutput `pulumi:"defaultInstanceProfileArn"`
	// Name of OS that will be installed on instances by default.
	DefaultOs pulumi.StringPtrOutput `pulumi:"defaultOs"`
	// Name of the type of root device instances will have by default.
	DefaultRootDeviceType pulumi.StringPtrOutput `pulumi:"defaultRootDeviceType"`
	// Name of the SSH keypair that instances will have by default.
	DefaultSshKeyName pulumi.StringPtrOutput `pulumi:"defaultSshKeyName"`
	// Id of the subnet in which instances will be created by default. Mandatory
	// if `vpcId` is set, and forbidden if it isn't.
	DefaultSubnetId pulumi.StringOutput `pulumi:"defaultSubnetId"`
	// Keyword representing the naming scheme that will be used for instance hostnames
	// within this stack.
	HostnameTheme pulumi.StringPtrOutput `pulumi:"hostnameTheme"`
	// Boolean value controlling whether Opsworks will run Berkshelf for this stack.
	ManageBerkshelf pulumi.BoolPtrOutput `pulumi:"manageBerkshelf"`
	// The name of the stack.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the region where the stack will exist.
	Region pulumi.StringOutput `pulumi:"region"`
	// The ARN of an IAM role that the OpsWorks service will act as.
	ServiceRoleArn pulumi.StringOutput `pulumi:"serviceRoleArn"`
	StackEndpoint pulumi.StringOutput `pulumi:"stackEndpoint"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Boolean value controlling whether the custom cookbook settings are
	// enabled.
	UseCustomCookbooks pulumi.BoolPtrOutput `pulumi:"useCustomCookbooks"`
	// Boolean value controlling whether the standard OpsWorks
	// security groups apply to created instances.
	UseOpsworksSecurityGroups pulumi.BoolPtrOutput `pulumi:"useOpsworksSecurityGroups"`
	// The id of the VPC that this stack belongs to.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewStack registers a new resource with the given unique name, arguments, and options.
func NewStack(ctx *pulumi.Context,
	name string, args *StackArgs, opts ...pulumi.ResourceOption) (*Stack, error) {
	if args == nil || args.DefaultInstanceProfileArn == nil {
		return nil, errors.New("missing required argument 'DefaultInstanceProfileArn'")
	}
	if args == nil || args.Region == nil {
		return nil, errors.New("missing required argument 'Region'")
	}
	if args == nil || args.ServiceRoleArn == nil {
		return nil, errors.New("missing required argument 'ServiceRoleArn'")
	}
	if args == nil {
		args = &StackArgs{}
	}
	var resource Stack
	err := ctx.RegisterResource("aws:opsworks/stack:Stack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStack gets an existing Stack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StackState, opts ...pulumi.ResourceOption) (*Stack, error) {
	var resource Stack
	err := ctx.ReadResource("aws:opsworks/stack:Stack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stack resources.
type stackState struct {
	// If set to `"LATEST"`, OpsWorks will automatically install the latest version.
	AgentVersion *string `pulumi:"agentVersion"`
	Arn *string `pulumi:"arn"`
	// If `manageBerkshelf` is enabled, the version of Berkshelf to use.
	BerkshelfVersion *string `pulumi:"berkshelfVersion"`
	// Color to paint next to the stack's resources in the OpsWorks console.
	Color *string `pulumi:"color"`
	// Name of the configuration manager to use. Defaults to "Chef".
	ConfigurationManagerName *string `pulumi:"configurationManagerName"`
	// Version of the configuration manager to use. Defaults to "11.4".
	ConfigurationManagerVersion *string `pulumi:"configurationManagerVersion"`
	// When `useCustomCookbooks` is set, provide this sub-object as
	// described below.
	CustomCookbooksSources []StackCustomCookbooksSource `pulumi:"customCookbooksSources"`
	// Custom JSON attributes to apply to the entire stack.
	CustomJson *string `pulumi:"customJson"`
	// Name of the availability zone where instances will be created
	// by default. This is required unless you set `vpcId`.
	DefaultAvailabilityZone *string `pulumi:"defaultAvailabilityZone"`
	// The ARN of an IAM Instance Profile that created instances
	// will have by default.
	DefaultInstanceProfileArn *string `pulumi:"defaultInstanceProfileArn"`
	// Name of OS that will be installed on instances by default.
	DefaultOs *string `pulumi:"defaultOs"`
	// Name of the type of root device instances will have by default.
	DefaultRootDeviceType *string `pulumi:"defaultRootDeviceType"`
	// Name of the SSH keypair that instances will have by default.
	DefaultSshKeyName *string `pulumi:"defaultSshKeyName"`
	// Id of the subnet in which instances will be created by default. Mandatory
	// if `vpcId` is set, and forbidden if it isn't.
	DefaultSubnetId *string `pulumi:"defaultSubnetId"`
	// Keyword representing the naming scheme that will be used for instance hostnames
	// within this stack.
	HostnameTheme *string `pulumi:"hostnameTheme"`
	// Boolean value controlling whether Opsworks will run Berkshelf for this stack.
	ManageBerkshelf *bool `pulumi:"manageBerkshelf"`
	// The name of the stack.
	Name *string `pulumi:"name"`
	// The name of the region where the stack will exist.
	Region *string `pulumi:"region"`
	// The ARN of an IAM role that the OpsWorks service will act as.
	ServiceRoleArn *string `pulumi:"serviceRoleArn"`
	StackEndpoint *string `pulumi:"stackEndpoint"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// Boolean value controlling whether the custom cookbook settings are
	// enabled.
	UseCustomCookbooks *bool `pulumi:"useCustomCookbooks"`
	// Boolean value controlling whether the standard OpsWorks
	// security groups apply to created instances.
	UseOpsworksSecurityGroups *bool `pulumi:"useOpsworksSecurityGroups"`
	// The id of the VPC that this stack belongs to.
	VpcId *string `pulumi:"vpcId"`
}

type StackState struct {
	// If set to `"LATEST"`, OpsWorks will automatically install the latest version.
	AgentVersion pulumi.StringPtrInput
	Arn pulumi.StringPtrInput
	// If `manageBerkshelf` is enabled, the version of Berkshelf to use.
	BerkshelfVersion pulumi.StringPtrInput
	// Color to paint next to the stack's resources in the OpsWorks console.
	Color pulumi.StringPtrInput
	// Name of the configuration manager to use. Defaults to "Chef".
	ConfigurationManagerName pulumi.StringPtrInput
	// Version of the configuration manager to use. Defaults to "11.4".
	ConfigurationManagerVersion pulumi.StringPtrInput
	// When `useCustomCookbooks` is set, provide this sub-object as
	// described below.
	CustomCookbooksSources StackCustomCookbooksSourceArrayInput
	// Custom JSON attributes to apply to the entire stack.
	CustomJson pulumi.StringPtrInput
	// Name of the availability zone where instances will be created
	// by default. This is required unless you set `vpcId`.
	DefaultAvailabilityZone pulumi.StringPtrInput
	// The ARN of an IAM Instance Profile that created instances
	// will have by default.
	DefaultInstanceProfileArn pulumi.StringPtrInput
	// Name of OS that will be installed on instances by default.
	DefaultOs pulumi.StringPtrInput
	// Name of the type of root device instances will have by default.
	DefaultRootDeviceType pulumi.StringPtrInput
	// Name of the SSH keypair that instances will have by default.
	DefaultSshKeyName pulumi.StringPtrInput
	// Id of the subnet in which instances will be created by default. Mandatory
	// if `vpcId` is set, and forbidden if it isn't.
	DefaultSubnetId pulumi.StringPtrInput
	// Keyword representing the naming scheme that will be used for instance hostnames
	// within this stack.
	HostnameTheme pulumi.StringPtrInput
	// Boolean value controlling whether Opsworks will run Berkshelf for this stack.
	ManageBerkshelf pulumi.BoolPtrInput
	// The name of the stack.
	Name pulumi.StringPtrInput
	// The name of the region where the stack will exist.
	Region pulumi.StringPtrInput
	// The ARN of an IAM role that the OpsWorks service will act as.
	ServiceRoleArn pulumi.StringPtrInput
	StackEndpoint pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// Boolean value controlling whether the custom cookbook settings are
	// enabled.
	UseCustomCookbooks pulumi.BoolPtrInput
	// Boolean value controlling whether the standard OpsWorks
	// security groups apply to created instances.
	UseOpsworksSecurityGroups pulumi.BoolPtrInput
	// The id of the VPC that this stack belongs to.
	VpcId pulumi.StringPtrInput
}

func (StackState) ElementType() reflect.Type {
	return reflect.TypeOf((*stackState)(nil)).Elem()
}

type stackArgs struct {
	// If set to `"LATEST"`, OpsWorks will automatically install the latest version.
	AgentVersion *string `pulumi:"agentVersion"`
	// If `manageBerkshelf` is enabled, the version of Berkshelf to use.
	BerkshelfVersion *string `pulumi:"berkshelfVersion"`
	// Color to paint next to the stack's resources in the OpsWorks console.
	Color *string `pulumi:"color"`
	// Name of the configuration manager to use. Defaults to "Chef".
	ConfigurationManagerName *string `pulumi:"configurationManagerName"`
	// Version of the configuration manager to use. Defaults to "11.4".
	ConfigurationManagerVersion *string `pulumi:"configurationManagerVersion"`
	// When `useCustomCookbooks` is set, provide this sub-object as
	// described below.
	CustomCookbooksSources []StackCustomCookbooksSource `pulumi:"customCookbooksSources"`
	// Custom JSON attributes to apply to the entire stack.
	CustomJson *string `pulumi:"customJson"`
	// Name of the availability zone where instances will be created
	// by default. This is required unless you set `vpcId`.
	DefaultAvailabilityZone *string `pulumi:"defaultAvailabilityZone"`
	// The ARN of an IAM Instance Profile that created instances
	// will have by default.
	DefaultInstanceProfileArn string `pulumi:"defaultInstanceProfileArn"`
	// Name of OS that will be installed on instances by default.
	DefaultOs *string `pulumi:"defaultOs"`
	// Name of the type of root device instances will have by default.
	DefaultRootDeviceType *string `pulumi:"defaultRootDeviceType"`
	// Name of the SSH keypair that instances will have by default.
	DefaultSshKeyName *string `pulumi:"defaultSshKeyName"`
	// Id of the subnet in which instances will be created by default. Mandatory
	// if `vpcId` is set, and forbidden if it isn't.
	DefaultSubnetId *string `pulumi:"defaultSubnetId"`
	// Keyword representing the naming scheme that will be used for instance hostnames
	// within this stack.
	HostnameTheme *string `pulumi:"hostnameTheme"`
	// Boolean value controlling whether Opsworks will run Berkshelf for this stack.
	ManageBerkshelf *bool `pulumi:"manageBerkshelf"`
	// The name of the stack.
	Name *string `pulumi:"name"`
	// The name of the region where the stack will exist.
	Region string `pulumi:"region"`
	// The ARN of an IAM role that the OpsWorks service will act as.
	ServiceRoleArn string `pulumi:"serviceRoleArn"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// Boolean value controlling whether the custom cookbook settings are
	// enabled.
	UseCustomCookbooks *bool `pulumi:"useCustomCookbooks"`
	// Boolean value controlling whether the standard OpsWorks
	// security groups apply to created instances.
	UseOpsworksSecurityGroups *bool `pulumi:"useOpsworksSecurityGroups"`
	// The id of the VPC that this stack belongs to.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Stack resource.
type StackArgs struct {
	// If set to `"LATEST"`, OpsWorks will automatically install the latest version.
	AgentVersion pulumi.StringPtrInput
	// If `manageBerkshelf` is enabled, the version of Berkshelf to use.
	BerkshelfVersion pulumi.StringPtrInput
	// Color to paint next to the stack's resources in the OpsWorks console.
	Color pulumi.StringPtrInput
	// Name of the configuration manager to use. Defaults to "Chef".
	ConfigurationManagerName pulumi.StringPtrInput
	// Version of the configuration manager to use. Defaults to "11.4".
	ConfigurationManagerVersion pulumi.StringPtrInput
	// When `useCustomCookbooks` is set, provide this sub-object as
	// described below.
	CustomCookbooksSources StackCustomCookbooksSourceArrayInput
	// Custom JSON attributes to apply to the entire stack.
	CustomJson pulumi.StringPtrInput
	// Name of the availability zone where instances will be created
	// by default. This is required unless you set `vpcId`.
	DefaultAvailabilityZone pulumi.StringPtrInput
	// The ARN of an IAM Instance Profile that created instances
	// will have by default.
	DefaultInstanceProfileArn pulumi.StringInput
	// Name of OS that will be installed on instances by default.
	DefaultOs pulumi.StringPtrInput
	// Name of the type of root device instances will have by default.
	DefaultRootDeviceType pulumi.StringPtrInput
	// Name of the SSH keypair that instances will have by default.
	DefaultSshKeyName pulumi.StringPtrInput
	// Id of the subnet in which instances will be created by default. Mandatory
	// if `vpcId` is set, and forbidden if it isn't.
	DefaultSubnetId pulumi.StringPtrInput
	// Keyword representing the naming scheme that will be used for instance hostnames
	// within this stack.
	HostnameTheme pulumi.StringPtrInput
	// Boolean value controlling whether Opsworks will run Berkshelf for this stack.
	ManageBerkshelf pulumi.BoolPtrInput
	// The name of the stack.
	Name pulumi.StringPtrInput
	// The name of the region where the stack will exist.
	Region pulumi.StringInput
	// The ARN of an IAM role that the OpsWorks service will act as.
	ServiceRoleArn pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// Boolean value controlling whether the custom cookbook settings are
	// enabled.
	UseCustomCookbooks pulumi.BoolPtrInput
	// Boolean value controlling whether the standard OpsWorks
	// security groups apply to created instances.
	UseOpsworksSecurityGroups pulumi.BoolPtrInput
	// The id of the VPC that this stack belongs to.
	VpcId pulumi.StringPtrInput
}

func (StackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stackArgs)(nil)).Elem()
}

