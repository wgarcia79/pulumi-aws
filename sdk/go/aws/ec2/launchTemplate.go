// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an EC2 launch template resource. Can be used to create instances or auto scaling groups.
//
// ## Import
//
// Launch Templates can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:ec2/launchTemplate:LaunchTemplate web lt-12345678
// ```
type LaunchTemplate struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the instance profile.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specify volumes to attach to the instance besides the volumes specified by the AMI.
	// See Block Devices below for details.
	BlockDeviceMappings LaunchTemplateBlockDeviceMappingArrayOutput `pulumi:"blockDeviceMappings"`
	// Targeting for EC2 capacity reservations. See Capacity Reservation Specification below for more details.
	CapacityReservationSpecification LaunchTemplateCapacityReservationSpecificationPtrOutput `pulumi:"capacityReservationSpecification"`
	// The CPU options for the instance. See CPU Options below for more details.
	CpuOptions LaunchTemplateCpuOptionsPtrOutput `pulumi:"cpuOptions"`
	// Customize the credit specification of the instance. See Credit
	// Specification below for more details.
	CreditSpecification LaunchTemplateCreditSpecificationPtrOutput `pulumi:"creditSpecification"`
	// Default Version of the launch template.
	DefaultVersion pulumi.IntOutput `pulumi:"defaultVersion"`
	// Description of the launch template.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If `true`, enables [EC2 Instance
	// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
	DisableApiTermination pulumi.BoolPtrOutput `pulumi:"disableApiTermination"`
	// If `true`, the launched EC2 instance will be EBS-optimized.
	EbsOptimized pulumi.StringPtrOutput `pulumi:"ebsOptimized"`
	// The elastic GPU to attach to the instance. See Elastic GPU
	// below for more details.
	ElasticGpuSpecifications LaunchTemplateElasticGpuSpecificationArrayOutput `pulumi:"elasticGpuSpecifications"`
	// Configuration block containing an Elastic Inference Accelerator to attach to the instance. See Elastic Inference Accelerator below for more details.
	ElasticInferenceAccelerator LaunchTemplateElasticInferenceAcceleratorPtrOutput `pulumi:"elasticInferenceAccelerator"`
	// Enable Nitro Enclaves on launched instances. See Enclave Options below for more details.
	EnclaveOptions LaunchTemplateEnclaveOptionsPtrOutput `pulumi:"enclaveOptions"`
	// The hibernation options for the instance. See Hibernation Options below for more details.
	HibernationOptions LaunchTemplateHibernationOptionsPtrOutput `pulumi:"hibernationOptions"`
	// The IAM Instance Profile to launch the instance with. See Instance Profile
	// below for more details.
	IamInstanceProfile LaunchTemplateIamInstanceProfilePtrOutput `pulumi:"iamInstanceProfile"`
	// The AMI from which to launch the instance.
	ImageId pulumi.StringPtrOutput `pulumi:"imageId"`
	// Shutdown behavior for the instance. Can be `stop` or `terminate`.
	// (Default: `stop`).
	InstanceInitiatedShutdownBehavior pulumi.StringPtrOutput `pulumi:"instanceInitiatedShutdownBehavior"`
	// The market (purchasing) option for the instance. See Market Options
	// below for details.
	InstanceMarketOptions LaunchTemplateInstanceMarketOptionsPtrOutput `pulumi:"instanceMarketOptions"`
	// The type of the instance.
	InstanceType pulumi.StringPtrOutput `pulumi:"instanceType"`
	// The kernel ID.
	KernelId pulumi.StringPtrOutput `pulumi:"kernelId"`
	// The key name to use for the instance.
	KeyName pulumi.StringPtrOutput `pulumi:"keyName"`
	// The latest version of the launch template.
	LatestVersion pulumi.IntOutput `pulumi:"latestVersion"`
	// A list of license specifications to associate with. See License Specification below for more details.
	LicenseSpecifications LaunchTemplateLicenseSpecificationArrayOutput `pulumi:"licenseSpecifications"`
	// Customize the metadata options for the instance. See Metadata Options below for more details.
	MetadataOptions LaunchTemplateMetadataOptionsOutput `pulumi:"metadataOptions"`
	// The monitoring option for the instance. See Monitoring below for more details.
	Monitoring LaunchTemplateMonitoringPtrOutput `pulumi:"monitoring"`
	// The name of the launch template. If you leave this blank, this provider will auto-generate a unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// Customize network interfaces to be attached at instance boot time. See Network
	// Interfaces below for more details.
	NetworkInterfaces LaunchTemplateNetworkInterfaceArrayOutput `pulumi:"networkInterfaces"`
	// The placement of the instance. See Placement below for more details.
	Placement LaunchTemplatePlacementPtrOutput `pulumi:"placement"`
	// The ID of the RAM disk.
	RamDiskId pulumi.StringPtrOutput `pulumi:"ramDiskId"`
	// A list of security group names to associate with. If you are creating Instances in a VPC, use
	// `vpcSecurityGroupIds` instead.
	SecurityGroupNames pulumi.StringArrayOutput `pulumi:"securityGroupNames"`
	// The tags to apply to the resources during launch. See Tag Specifications below for more details.
	TagSpecifications LaunchTemplateTagSpecificationArrayOutput `pulumi:"tagSpecifications"`
	// A map of tags to assign to the launch template. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Whether to update Default Version each update. Conflicts with `defaultVersion`.
	UpdateDefaultVersion pulumi.BoolPtrOutput `pulumi:"updateDefaultVersion"`
	// The Base64-encoded user data to provide when launching the instance.
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
	// A list of security group IDs to associate with.
	VpcSecurityGroupIds pulumi.StringArrayOutput `pulumi:"vpcSecurityGroupIds"`
}

// NewLaunchTemplate registers a new resource with the given unique name, arguments, and options.
func NewLaunchTemplate(ctx *pulumi.Context,
	name string, args *LaunchTemplateArgs, opts ...pulumi.ResourceOption) (*LaunchTemplate, error) {
	if args == nil {
		args = &LaunchTemplateArgs{}
	}

	var resource LaunchTemplate
	err := ctx.RegisterResource("aws:ec2/launchTemplate:LaunchTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLaunchTemplate gets an existing LaunchTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLaunchTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LaunchTemplateState, opts ...pulumi.ResourceOption) (*LaunchTemplate, error) {
	var resource LaunchTemplate
	err := ctx.ReadResource("aws:ec2/launchTemplate:LaunchTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LaunchTemplate resources.
type launchTemplateState struct {
	// The Amazon Resource Name (ARN) of the instance profile.
	Arn *string `pulumi:"arn"`
	// Specify volumes to attach to the instance besides the volumes specified by the AMI.
	// See Block Devices below for details.
	BlockDeviceMappings []LaunchTemplateBlockDeviceMapping `pulumi:"blockDeviceMappings"`
	// Targeting for EC2 capacity reservations. See Capacity Reservation Specification below for more details.
	CapacityReservationSpecification *LaunchTemplateCapacityReservationSpecification `pulumi:"capacityReservationSpecification"`
	// The CPU options for the instance. See CPU Options below for more details.
	CpuOptions *LaunchTemplateCpuOptions `pulumi:"cpuOptions"`
	// Customize the credit specification of the instance. See Credit
	// Specification below for more details.
	CreditSpecification *LaunchTemplateCreditSpecification `pulumi:"creditSpecification"`
	// Default Version of the launch template.
	DefaultVersion *int `pulumi:"defaultVersion"`
	// Description of the launch template.
	Description *string `pulumi:"description"`
	// If `true`, enables [EC2 Instance
	// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
	DisableApiTermination *bool `pulumi:"disableApiTermination"`
	// If `true`, the launched EC2 instance will be EBS-optimized.
	EbsOptimized *string `pulumi:"ebsOptimized"`
	// The elastic GPU to attach to the instance. See Elastic GPU
	// below for more details.
	ElasticGpuSpecifications []LaunchTemplateElasticGpuSpecification `pulumi:"elasticGpuSpecifications"`
	// Configuration block containing an Elastic Inference Accelerator to attach to the instance. See Elastic Inference Accelerator below for more details.
	ElasticInferenceAccelerator *LaunchTemplateElasticInferenceAccelerator `pulumi:"elasticInferenceAccelerator"`
	// Enable Nitro Enclaves on launched instances. See Enclave Options below for more details.
	EnclaveOptions *LaunchTemplateEnclaveOptions `pulumi:"enclaveOptions"`
	// The hibernation options for the instance. See Hibernation Options below for more details.
	HibernationOptions *LaunchTemplateHibernationOptions `pulumi:"hibernationOptions"`
	// The IAM Instance Profile to launch the instance with. See Instance Profile
	// below for more details.
	IamInstanceProfile *LaunchTemplateIamInstanceProfile `pulumi:"iamInstanceProfile"`
	// The AMI from which to launch the instance.
	ImageId *string `pulumi:"imageId"`
	// Shutdown behavior for the instance. Can be `stop` or `terminate`.
	// (Default: `stop`).
	InstanceInitiatedShutdownBehavior *string `pulumi:"instanceInitiatedShutdownBehavior"`
	// The market (purchasing) option for the instance. See Market Options
	// below for details.
	InstanceMarketOptions *LaunchTemplateInstanceMarketOptions `pulumi:"instanceMarketOptions"`
	// The type of the instance.
	InstanceType *string `pulumi:"instanceType"`
	// The kernel ID.
	KernelId *string `pulumi:"kernelId"`
	// The key name to use for the instance.
	KeyName *string `pulumi:"keyName"`
	// The latest version of the launch template.
	LatestVersion *int `pulumi:"latestVersion"`
	// A list of license specifications to associate with. See License Specification below for more details.
	LicenseSpecifications []LaunchTemplateLicenseSpecification `pulumi:"licenseSpecifications"`
	// Customize the metadata options for the instance. See Metadata Options below for more details.
	MetadataOptions *LaunchTemplateMetadataOptions `pulumi:"metadataOptions"`
	// The monitoring option for the instance. See Monitoring below for more details.
	Monitoring *LaunchTemplateMonitoring `pulumi:"monitoring"`
	// The name of the launch template. If you leave this blank, this provider will auto-generate a unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Customize network interfaces to be attached at instance boot time. See Network
	// Interfaces below for more details.
	NetworkInterfaces []LaunchTemplateNetworkInterface `pulumi:"networkInterfaces"`
	// The placement of the instance. See Placement below for more details.
	Placement *LaunchTemplatePlacement `pulumi:"placement"`
	// The ID of the RAM disk.
	RamDiskId *string `pulumi:"ramDiskId"`
	// A list of security group names to associate with. If you are creating Instances in a VPC, use
	// `vpcSecurityGroupIds` instead.
	SecurityGroupNames []string `pulumi:"securityGroupNames"`
	// The tags to apply to the resources during launch. See Tag Specifications below for more details.
	TagSpecifications []LaunchTemplateTagSpecification `pulumi:"tagSpecifications"`
	// A map of tags to assign to the launch template. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Whether to update Default Version each update. Conflicts with `defaultVersion`.
	UpdateDefaultVersion *bool `pulumi:"updateDefaultVersion"`
	// The Base64-encoded user data to provide when launching the instance.
	UserData *string `pulumi:"userData"`
	// A list of security group IDs to associate with.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
}

type LaunchTemplateState struct {
	// The Amazon Resource Name (ARN) of the instance profile.
	Arn pulumi.StringPtrInput
	// Specify volumes to attach to the instance besides the volumes specified by the AMI.
	// See Block Devices below for details.
	BlockDeviceMappings LaunchTemplateBlockDeviceMappingArrayInput
	// Targeting for EC2 capacity reservations. See Capacity Reservation Specification below for more details.
	CapacityReservationSpecification LaunchTemplateCapacityReservationSpecificationPtrInput
	// The CPU options for the instance. See CPU Options below for more details.
	CpuOptions LaunchTemplateCpuOptionsPtrInput
	// Customize the credit specification of the instance. See Credit
	// Specification below for more details.
	CreditSpecification LaunchTemplateCreditSpecificationPtrInput
	// Default Version of the launch template.
	DefaultVersion pulumi.IntPtrInput
	// Description of the launch template.
	Description pulumi.StringPtrInput
	// If `true`, enables [EC2 Instance
	// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
	DisableApiTermination pulumi.BoolPtrInput
	// If `true`, the launched EC2 instance will be EBS-optimized.
	EbsOptimized pulumi.StringPtrInput
	// The elastic GPU to attach to the instance. See Elastic GPU
	// below for more details.
	ElasticGpuSpecifications LaunchTemplateElasticGpuSpecificationArrayInput
	// Configuration block containing an Elastic Inference Accelerator to attach to the instance. See Elastic Inference Accelerator below for more details.
	ElasticInferenceAccelerator LaunchTemplateElasticInferenceAcceleratorPtrInput
	// Enable Nitro Enclaves on launched instances. See Enclave Options below for more details.
	EnclaveOptions LaunchTemplateEnclaveOptionsPtrInput
	// The hibernation options for the instance. See Hibernation Options below for more details.
	HibernationOptions LaunchTemplateHibernationOptionsPtrInput
	// The IAM Instance Profile to launch the instance with. See Instance Profile
	// below for more details.
	IamInstanceProfile LaunchTemplateIamInstanceProfilePtrInput
	// The AMI from which to launch the instance.
	ImageId pulumi.StringPtrInput
	// Shutdown behavior for the instance. Can be `stop` or `terminate`.
	// (Default: `stop`).
	InstanceInitiatedShutdownBehavior pulumi.StringPtrInput
	// The market (purchasing) option for the instance. See Market Options
	// below for details.
	InstanceMarketOptions LaunchTemplateInstanceMarketOptionsPtrInput
	// The type of the instance.
	InstanceType pulumi.StringPtrInput
	// The kernel ID.
	KernelId pulumi.StringPtrInput
	// The key name to use for the instance.
	KeyName pulumi.StringPtrInput
	// The latest version of the launch template.
	LatestVersion pulumi.IntPtrInput
	// A list of license specifications to associate with. See License Specification below for more details.
	LicenseSpecifications LaunchTemplateLicenseSpecificationArrayInput
	// Customize the metadata options for the instance. See Metadata Options below for more details.
	MetadataOptions LaunchTemplateMetadataOptionsPtrInput
	// The monitoring option for the instance. See Monitoring below for more details.
	Monitoring LaunchTemplateMonitoringPtrInput
	// The name of the launch template. If you leave this blank, this provider will auto-generate a unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Customize network interfaces to be attached at instance boot time. See Network
	// Interfaces below for more details.
	NetworkInterfaces LaunchTemplateNetworkInterfaceArrayInput
	// The placement of the instance. See Placement below for more details.
	Placement LaunchTemplatePlacementPtrInput
	// The ID of the RAM disk.
	RamDiskId pulumi.StringPtrInput
	// A list of security group names to associate with. If you are creating Instances in a VPC, use
	// `vpcSecurityGroupIds` instead.
	SecurityGroupNames pulumi.StringArrayInput
	// The tags to apply to the resources during launch. See Tag Specifications below for more details.
	TagSpecifications LaunchTemplateTagSpecificationArrayInput
	// A map of tags to assign to the launch template. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Whether to update Default Version each update. Conflicts with `defaultVersion`.
	UpdateDefaultVersion pulumi.BoolPtrInput
	// The Base64-encoded user data to provide when launching the instance.
	UserData pulumi.StringPtrInput
	// A list of security group IDs to associate with.
	VpcSecurityGroupIds pulumi.StringArrayInput
}

func (LaunchTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*launchTemplateState)(nil)).Elem()
}

type launchTemplateArgs struct {
	// Specify volumes to attach to the instance besides the volumes specified by the AMI.
	// See Block Devices below for details.
	BlockDeviceMappings []LaunchTemplateBlockDeviceMapping `pulumi:"blockDeviceMappings"`
	// Targeting for EC2 capacity reservations. See Capacity Reservation Specification below for more details.
	CapacityReservationSpecification *LaunchTemplateCapacityReservationSpecification `pulumi:"capacityReservationSpecification"`
	// The CPU options for the instance. See CPU Options below for more details.
	CpuOptions *LaunchTemplateCpuOptions `pulumi:"cpuOptions"`
	// Customize the credit specification of the instance. See Credit
	// Specification below for more details.
	CreditSpecification *LaunchTemplateCreditSpecification `pulumi:"creditSpecification"`
	// Default Version of the launch template.
	DefaultVersion *int `pulumi:"defaultVersion"`
	// Description of the launch template.
	Description *string `pulumi:"description"`
	// If `true`, enables [EC2 Instance
	// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
	DisableApiTermination *bool `pulumi:"disableApiTermination"`
	// If `true`, the launched EC2 instance will be EBS-optimized.
	EbsOptimized *string `pulumi:"ebsOptimized"`
	// The elastic GPU to attach to the instance. See Elastic GPU
	// below for more details.
	ElasticGpuSpecifications []LaunchTemplateElasticGpuSpecification `pulumi:"elasticGpuSpecifications"`
	// Configuration block containing an Elastic Inference Accelerator to attach to the instance. See Elastic Inference Accelerator below for more details.
	ElasticInferenceAccelerator *LaunchTemplateElasticInferenceAccelerator `pulumi:"elasticInferenceAccelerator"`
	// Enable Nitro Enclaves on launched instances. See Enclave Options below for more details.
	EnclaveOptions *LaunchTemplateEnclaveOptions `pulumi:"enclaveOptions"`
	// The hibernation options for the instance. See Hibernation Options below for more details.
	HibernationOptions *LaunchTemplateHibernationOptions `pulumi:"hibernationOptions"`
	// The IAM Instance Profile to launch the instance with. See Instance Profile
	// below for more details.
	IamInstanceProfile *LaunchTemplateIamInstanceProfile `pulumi:"iamInstanceProfile"`
	// The AMI from which to launch the instance.
	ImageId *string `pulumi:"imageId"`
	// Shutdown behavior for the instance. Can be `stop` or `terminate`.
	// (Default: `stop`).
	InstanceInitiatedShutdownBehavior *string `pulumi:"instanceInitiatedShutdownBehavior"`
	// The market (purchasing) option for the instance. See Market Options
	// below for details.
	InstanceMarketOptions *LaunchTemplateInstanceMarketOptions `pulumi:"instanceMarketOptions"`
	// The type of the instance.
	InstanceType *string `pulumi:"instanceType"`
	// The kernel ID.
	KernelId *string `pulumi:"kernelId"`
	// The key name to use for the instance.
	KeyName *string `pulumi:"keyName"`
	// A list of license specifications to associate with. See License Specification below for more details.
	LicenseSpecifications []LaunchTemplateLicenseSpecification `pulumi:"licenseSpecifications"`
	// Customize the metadata options for the instance. See Metadata Options below for more details.
	MetadataOptions *LaunchTemplateMetadataOptions `pulumi:"metadataOptions"`
	// The monitoring option for the instance. See Monitoring below for more details.
	Monitoring *LaunchTemplateMonitoring `pulumi:"monitoring"`
	// The name of the launch template. If you leave this blank, this provider will auto-generate a unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Customize network interfaces to be attached at instance boot time. See Network
	// Interfaces below for more details.
	NetworkInterfaces []LaunchTemplateNetworkInterface `pulumi:"networkInterfaces"`
	// The placement of the instance. See Placement below for more details.
	Placement *LaunchTemplatePlacement `pulumi:"placement"`
	// The ID of the RAM disk.
	RamDiskId *string `pulumi:"ramDiskId"`
	// A list of security group names to associate with. If you are creating Instances in a VPC, use
	// `vpcSecurityGroupIds` instead.
	SecurityGroupNames []string `pulumi:"securityGroupNames"`
	// The tags to apply to the resources during launch. See Tag Specifications below for more details.
	TagSpecifications []LaunchTemplateTagSpecification `pulumi:"tagSpecifications"`
	// A map of tags to assign to the launch template. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Whether to update Default Version each update. Conflicts with `defaultVersion`.
	UpdateDefaultVersion *bool `pulumi:"updateDefaultVersion"`
	// The Base64-encoded user data to provide when launching the instance.
	UserData *string `pulumi:"userData"`
	// A list of security group IDs to associate with.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
}

// The set of arguments for constructing a LaunchTemplate resource.
type LaunchTemplateArgs struct {
	// Specify volumes to attach to the instance besides the volumes specified by the AMI.
	// See Block Devices below for details.
	BlockDeviceMappings LaunchTemplateBlockDeviceMappingArrayInput
	// Targeting for EC2 capacity reservations. See Capacity Reservation Specification below for more details.
	CapacityReservationSpecification LaunchTemplateCapacityReservationSpecificationPtrInput
	// The CPU options for the instance. See CPU Options below for more details.
	CpuOptions LaunchTemplateCpuOptionsPtrInput
	// Customize the credit specification of the instance. See Credit
	// Specification below for more details.
	CreditSpecification LaunchTemplateCreditSpecificationPtrInput
	// Default Version of the launch template.
	DefaultVersion pulumi.IntPtrInput
	// Description of the launch template.
	Description pulumi.StringPtrInput
	// If `true`, enables [EC2 Instance
	// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
	DisableApiTermination pulumi.BoolPtrInput
	// If `true`, the launched EC2 instance will be EBS-optimized.
	EbsOptimized pulumi.StringPtrInput
	// The elastic GPU to attach to the instance. See Elastic GPU
	// below for more details.
	ElasticGpuSpecifications LaunchTemplateElasticGpuSpecificationArrayInput
	// Configuration block containing an Elastic Inference Accelerator to attach to the instance. See Elastic Inference Accelerator below for more details.
	ElasticInferenceAccelerator LaunchTemplateElasticInferenceAcceleratorPtrInput
	// Enable Nitro Enclaves on launched instances. See Enclave Options below for more details.
	EnclaveOptions LaunchTemplateEnclaveOptionsPtrInput
	// The hibernation options for the instance. See Hibernation Options below for more details.
	HibernationOptions LaunchTemplateHibernationOptionsPtrInput
	// The IAM Instance Profile to launch the instance with. See Instance Profile
	// below for more details.
	IamInstanceProfile LaunchTemplateIamInstanceProfilePtrInput
	// The AMI from which to launch the instance.
	ImageId pulumi.StringPtrInput
	// Shutdown behavior for the instance. Can be `stop` or `terminate`.
	// (Default: `stop`).
	InstanceInitiatedShutdownBehavior pulumi.StringPtrInput
	// The market (purchasing) option for the instance. See Market Options
	// below for details.
	InstanceMarketOptions LaunchTemplateInstanceMarketOptionsPtrInput
	// The type of the instance.
	InstanceType pulumi.StringPtrInput
	// The kernel ID.
	KernelId pulumi.StringPtrInput
	// The key name to use for the instance.
	KeyName pulumi.StringPtrInput
	// A list of license specifications to associate with. See License Specification below for more details.
	LicenseSpecifications LaunchTemplateLicenseSpecificationArrayInput
	// Customize the metadata options for the instance. See Metadata Options below for more details.
	MetadataOptions LaunchTemplateMetadataOptionsPtrInput
	// The monitoring option for the instance. See Monitoring below for more details.
	Monitoring LaunchTemplateMonitoringPtrInput
	// The name of the launch template. If you leave this blank, this provider will auto-generate a unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Customize network interfaces to be attached at instance boot time. See Network
	// Interfaces below for more details.
	NetworkInterfaces LaunchTemplateNetworkInterfaceArrayInput
	// The placement of the instance. See Placement below for more details.
	Placement LaunchTemplatePlacementPtrInput
	// The ID of the RAM disk.
	RamDiskId pulumi.StringPtrInput
	// A list of security group names to associate with. If you are creating Instances in a VPC, use
	// `vpcSecurityGroupIds` instead.
	SecurityGroupNames pulumi.StringArrayInput
	// The tags to apply to the resources during launch. See Tag Specifications below for more details.
	TagSpecifications LaunchTemplateTagSpecificationArrayInput
	// A map of tags to assign to the launch template. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Whether to update Default Version each update. Conflicts with `defaultVersion`.
	UpdateDefaultVersion pulumi.BoolPtrInput
	// The Base64-encoded user data to provide when launching the instance.
	UserData pulumi.StringPtrInput
	// A list of security group IDs to associate with.
	VpcSecurityGroupIds pulumi.StringArrayInput
}

func (LaunchTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*launchTemplateArgs)(nil)).Elem()
}

type LaunchTemplateInput interface {
	pulumi.Input

	ToLaunchTemplateOutput() LaunchTemplateOutput
	ToLaunchTemplateOutputWithContext(ctx context.Context) LaunchTemplateOutput
}

func (*LaunchTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((*LaunchTemplate)(nil))
}

func (i *LaunchTemplate) ToLaunchTemplateOutput() LaunchTemplateOutput {
	return i.ToLaunchTemplateOutputWithContext(context.Background())
}

func (i *LaunchTemplate) ToLaunchTemplateOutputWithContext(ctx context.Context) LaunchTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateOutput)
}

func (i *LaunchTemplate) ToLaunchTemplatePtrOutput() LaunchTemplatePtrOutput {
	return i.ToLaunchTemplatePtrOutputWithContext(context.Background())
}

func (i *LaunchTemplate) ToLaunchTemplatePtrOutputWithContext(ctx context.Context) LaunchTemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplatePtrOutput)
}

type LaunchTemplatePtrInput interface {
	pulumi.Input

	ToLaunchTemplatePtrOutput() LaunchTemplatePtrOutput
	ToLaunchTemplatePtrOutputWithContext(ctx context.Context) LaunchTemplatePtrOutput
}

type launchTemplatePtrType LaunchTemplateArgs

func (*launchTemplatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LaunchTemplate)(nil))
}

func (i *launchTemplatePtrType) ToLaunchTemplatePtrOutput() LaunchTemplatePtrOutput {
	return i.ToLaunchTemplatePtrOutputWithContext(context.Background())
}

func (i *launchTemplatePtrType) ToLaunchTemplatePtrOutputWithContext(ctx context.Context) LaunchTemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplatePtrOutput)
}

// LaunchTemplateArrayInput is an input type that accepts LaunchTemplateArray and LaunchTemplateArrayOutput values.
// You can construct a concrete instance of `LaunchTemplateArrayInput` via:
//
//          LaunchTemplateArray{ LaunchTemplateArgs{...} }
type LaunchTemplateArrayInput interface {
	pulumi.Input

	ToLaunchTemplateArrayOutput() LaunchTemplateArrayOutput
	ToLaunchTemplateArrayOutputWithContext(context.Context) LaunchTemplateArrayOutput
}

type LaunchTemplateArray []LaunchTemplateInput

func (LaunchTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LaunchTemplate)(nil))
}

func (i LaunchTemplateArray) ToLaunchTemplateArrayOutput() LaunchTemplateArrayOutput {
	return i.ToLaunchTemplateArrayOutputWithContext(context.Background())
}

func (i LaunchTemplateArray) ToLaunchTemplateArrayOutputWithContext(ctx context.Context) LaunchTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateArrayOutput)
}

// LaunchTemplateMapInput is an input type that accepts LaunchTemplateMap and LaunchTemplateMapOutput values.
// You can construct a concrete instance of `LaunchTemplateMapInput` via:
//
//          LaunchTemplateMap{ "key": LaunchTemplateArgs{...} }
type LaunchTemplateMapInput interface {
	pulumi.Input

	ToLaunchTemplateMapOutput() LaunchTemplateMapOutput
	ToLaunchTemplateMapOutputWithContext(context.Context) LaunchTemplateMapOutput
}

type LaunchTemplateMap map[string]LaunchTemplateInput

func (LaunchTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LaunchTemplate)(nil))
}

func (i LaunchTemplateMap) ToLaunchTemplateMapOutput() LaunchTemplateMapOutput {
	return i.ToLaunchTemplateMapOutputWithContext(context.Background())
}

func (i LaunchTemplateMap) ToLaunchTemplateMapOutputWithContext(ctx context.Context) LaunchTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateMapOutput)
}

type LaunchTemplateOutput struct {
	*pulumi.OutputState
}

func (LaunchTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LaunchTemplate)(nil))
}

func (o LaunchTemplateOutput) ToLaunchTemplateOutput() LaunchTemplateOutput {
	return o
}

func (o LaunchTemplateOutput) ToLaunchTemplateOutputWithContext(ctx context.Context) LaunchTemplateOutput {
	return o
}

func (o LaunchTemplateOutput) ToLaunchTemplatePtrOutput() LaunchTemplatePtrOutput {
	return o.ToLaunchTemplatePtrOutputWithContext(context.Background())
}

func (o LaunchTemplateOutput) ToLaunchTemplatePtrOutputWithContext(ctx context.Context) LaunchTemplatePtrOutput {
	return o.ApplyT(func(v LaunchTemplate) *LaunchTemplate {
		return &v
	}).(LaunchTemplatePtrOutput)
}

type LaunchTemplatePtrOutput struct {
	*pulumi.OutputState
}

func (LaunchTemplatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LaunchTemplate)(nil))
}

func (o LaunchTemplatePtrOutput) ToLaunchTemplatePtrOutput() LaunchTemplatePtrOutput {
	return o
}

func (o LaunchTemplatePtrOutput) ToLaunchTemplatePtrOutputWithContext(ctx context.Context) LaunchTemplatePtrOutput {
	return o
}

type LaunchTemplateArrayOutput struct{ *pulumi.OutputState }

func (LaunchTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LaunchTemplate)(nil))
}

func (o LaunchTemplateArrayOutput) ToLaunchTemplateArrayOutput() LaunchTemplateArrayOutput {
	return o
}

func (o LaunchTemplateArrayOutput) ToLaunchTemplateArrayOutputWithContext(ctx context.Context) LaunchTemplateArrayOutput {
	return o
}

func (o LaunchTemplateArrayOutput) Index(i pulumi.IntInput) LaunchTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LaunchTemplate {
		return vs[0].([]LaunchTemplate)[vs[1].(int)]
	}).(LaunchTemplateOutput)
}

type LaunchTemplateMapOutput struct{ *pulumi.OutputState }

func (LaunchTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LaunchTemplate)(nil))
}

func (o LaunchTemplateMapOutput) ToLaunchTemplateMapOutput() LaunchTemplateMapOutput {
	return o
}

func (o LaunchTemplateMapOutput) ToLaunchTemplateMapOutputWithContext(ctx context.Context) LaunchTemplateMapOutput {
	return o
}

func (o LaunchTemplateMapOutput) MapIndex(k pulumi.StringInput) LaunchTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LaunchTemplate {
		return vs[0].(map[string]LaunchTemplate)[vs[1].(string)]
	}).(LaunchTemplateOutput)
}

func init() {
	pulumi.RegisterOutputType(LaunchTemplateOutput{})
	pulumi.RegisterOutputType(LaunchTemplatePtrOutput{})
	pulumi.RegisterOutputType(LaunchTemplateArrayOutput{})
	pulumi.RegisterOutputType(LaunchTemplateMapOutput{})
}
