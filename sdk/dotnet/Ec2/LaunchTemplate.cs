// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Provides an EC2 launch template resource. Can be used to create instances or auto scaling groups.
    /// 
    /// ## Import
    /// 
    /// Launch Templates can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2/launchTemplate:LaunchTemplate web lt-12345678
    /// ```
    /// </summary>
    [AwsResourceType("aws:ec2/launchTemplate:LaunchTemplate")]
    public partial class LaunchTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the instance profile.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Specify volumes to attach to the instance besides the volumes specified by the AMI.
        /// See Block Devices below for details.
        /// </summary>
        [Output("blockDeviceMappings")]
        public Output<ImmutableArray<Outputs.LaunchTemplateBlockDeviceMapping>> BlockDeviceMappings { get; private set; } = null!;

        /// <summary>
        /// Targeting for EC2 capacity reservations. See Capacity Reservation Specification below for more details.
        /// </summary>
        [Output("capacityReservationSpecification")]
        public Output<Outputs.LaunchTemplateCapacityReservationSpecification?> CapacityReservationSpecification { get; private set; } = null!;

        /// <summary>
        /// The CPU options for the instance. See CPU Options below for more details.
        /// </summary>
        [Output("cpuOptions")]
        public Output<Outputs.LaunchTemplateCpuOptions?> CpuOptions { get; private set; } = null!;

        /// <summary>
        /// Customize the credit specification of the instance. See Credit
        /// Specification below for more details.
        /// </summary>
        [Output("creditSpecification")]
        public Output<Outputs.LaunchTemplateCreditSpecification?> CreditSpecification { get; private set; } = null!;

        /// <summary>
        /// Default Version of the launch template.
        /// </summary>
        [Output("defaultVersion")]
        public Output<int> DefaultVersion { get; private set; } = null!;

        /// <summary>
        /// Description of the launch template.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// If `true`, enables [EC2 Instance
        /// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
        /// </summary>
        [Output("disableApiTermination")]
        public Output<bool?> DisableApiTermination { get; private set; } = null!;

        /// <summary>
        /// If `true`, the launched EC2 instance will be EBS-optimized.
        /// </summary>
        [Output("ebsOptimized")]
        public Output<string?> EbsOptimized { get; private set; } = null!;

        /// <summary>
        /// The elastic GPU to attach to the instance. See Elastic GPU
        /// below for more details.
        /// </summary>
        [Output("elasticGpuSpecifications")]
        public Output<ImmutableArray<Outputs.LaunchTemplateElasticGpuSpecification>> ElasticGpuSpecifications { get; private set; } = null!;

        /// <summary>
        /// Configuration block containing an Elastic Inference Accelerator to attach to the instance. See Elastic Inference Accelerator below for more details.
        /// </summary>
        [Output("elasticInferenceAccelerator")]
        public Output<Outputs.LaunchTemplateElasticInferenceAccelerator?> ElasticInferenceAccelerator { get; private set; } = null!;

        /// <summary>
        /// Enable Nitro Enclaves on launched instances. See Enclave Options below for more details.
        /// </summary>
        [Output("enclaveOptions")]
        public Output<Outputs.LaunchTemplateEnclaveOptions?> EnclaveOptions { get; private set; } = null!;

        /// <summary>
        /// The hibernation options for the instance. See Hibernation Options below for more details.
        /// </summary>
        [Output("hibernationOptions")]
        public Output<Outputs.LaunchTemplateHibernationOptions?> HibernationOptions { get; private set; } = null!;

        /// <summary>
        /// The IAM Instance Profile to launch the instance with. See Instance Profile
        /// below for more details.
        /// </summary>
        [Output("iamInstanceProfile")]
        public Output<Outputs.LaunchTemplateIamInstanceProfile?> IamInstanceProfile { get; private set; } = null!;

        /// <summary>
        /// The AMI from which to launch the instance.
        /// </summary>
        [Output("imageId")]
        public Output<string?> ImageId { get; private set; } = null!;

        /// <summary>
        /// Shutdown behavior for the instance. Can be `stop` or `terminate`.
        /// (Default: `stop`).
        /// </summary>
        [Output("instanceInitiatedShutdownBehavior")]
        public Output<string?> InstanceInitiatedShutdownBehavior { get; private set; } = null!;

        /// <summary>
        /// The market (purchasing) option for the instance. See Market Options
        /// below for details.
        /// </summary>
        [Output("instanceMarketOptions")]
        public Output<Outputs.LaunchTemplateInstanceMarketOptions?> InstanceMarketOptions { get; private set; } = null!;

        /// <summary>
        /// The type of the instance.
        /// </summary>
        [Output("instanceType")]
        public Output<string?> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The kernel ID.
        /// </summary>
        [Output("kernelId")]
        public Output<string?> KernelId { get; private set; } = null!;

        /// <summary>
        /// The key name to use for the instance.
        /// </summary>
        [Output("keyName")]
        public Output<string?> KeyName { get; private set; } = null!;

        /// <summary>
        /// The latest version of the launch template.
        /// </summary>
        [Output("latestVersion")]
        public Output<int> LatestVersion { get; private set; } = null!;

        /// <summary>
        /// A list of license specifications to associate with. See License Specification below for more details.
        /// </summary>
        [Output("licenseSpecifications")]
        public Output<ImmutableArray<Outputs.LaunchTemplateLicenseSpecification>> LicenseSpecifications { get; private set; } = null!;

        /// <summary>
        /// Customize the metadata options for the instance. See Metadata Options below for more details.
        /// </summary>
        [Output("metadataOptions")]
        public Output<Outputs.LaunchTemplateMetadataOptions> MetadataOptions { get; private set; } = null!;

        /// <summary>
        /// The monitoring option for the instance. See Monitoring below for more details.
        /// </summary>
        [Output("monitoring")]
        public Output<Outputs.LaunchTemplateMonitoring?> Monitoring { get; private set; } = null!;

        /// <summary>
        /// The name of the launch template. If you leave this blank, this provider will auto-generate a unique name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// Customize network interfaces to be attached at instance boot time. See Network
        /// Interfaces below for more details.
        /// </summary>
        [Output("networkInterfaces")]
        public Output<ImmutableArray<Outputs.LaunchTemplateNetworkInterface>> NetworkInterfaces { get; private set; } = null!;

        /// <summary>
        /// The placement of the instance. See Placement below for more details.
        /// </summary>
        [Output("placement")]
        public Output<Outputs.LaunchTemplatePlacement?> Placement { get; private set; } = null!;

        /// <summary>
        /// The ID of the RAM disk.
        /// </summary>
        [Output("ramDiskId")]
        public Output<string?> RamDiskId { get; private set; } = null!;

        /// <summary>
        /// A list of security group names to associate with. If you are creating Instances in a VPC, use
        /// `vpc_security_group_ids` instead.
        /// </summary>
        [Output("securityGroupNames")]
        public Output<ImmutableArray<string>> SecurityGroupNames { get; private set; } = null!;

        /// <summary>
        /// The tags to apply to the resources during launch. See Tag Specifications below for more details.
        /// </summary>
        [Output("tagSpecifications")]
        public Output<ImmutableArray<Outputs.LaunchTemplateTagSpecification>> TagSpecifications { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the launch template. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Whether to update Default Version each update. Conflicts with `default_version`.
        /// </summary>
        [Output("updateDefaultVersion")]
        public Output<bool?> UpdateDefaultVersion { get; private set; } = null!;

        /// <summary>
        /// The Base64-encoded user data to provide when launching the instance.
        /// </summary>
        [Output("userData")]
        public Output<string?> UserData { get; private set; } = null!;

        /// <summary>
        /// A list of security group IDs to associate with.
        /// </summary>
        [Output("vpcSecurityGroupIds")]
        public Output<ImmutableArray<string>> VpcSecurityGroupIds { get; private set; } = null!;


        /// <summary>
        /// Create a LaunchTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LaunchTemplate(string name, LaunchTemplateArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:ec2/launchTemplate:LaunchTemplate", name, args ?? new LaunchTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LaunchTemplate(string name, Input<string> id, LaunchTemplateState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/launchTemplate:LaunchTemplate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LaunchTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LaunchTemplate Get(string name, Input<string> id, LaunchTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new LaunchTemplate(name, id, state, options);
        }
    }

    public sealed class LaunchTemplateArgs : Pulumi.ResourceArgs
    {
        [Input("blockDeviceMappings")]
        private InputList<Inputs.LaunchTemplateBlockDeviceMappingArgs>? _blockDeviceMappings;

        /// <summary>
        /// Specify volumes to attach to the instance besides the volumes specified by the AMI.
        /// See Block Devices below for details.
        /// </summary>
        public InputList<Inputs.LaunchTemplateBlockDeviceMappingArgs> BlockDeviceMappings
        {
            get => _blockDeviceMappings ?? (_blockDeviceMappings = new InputList<Inputs.LaunchTemplateBlockDeviceMappingArgs>());
            set => _blockDeviceMappings = value;
        }

        /// <summary>
        /// Targeting for EC2 capacity reservations. See Capacity Reservation Specification below for more details.
        /// </summary>
        [Input("capacityReservationSpecification")]
        public Input<Inputs.LaunchTemplateCapacityReservationSpecificationArgs>? CapacityReservationSpecification { get; set; }

        /// <summary>
        /// The CPU options for the instance. See CPU Options below for more details.
        /// </summary>
        [Input("cpuOptions")]
        public Input<Inputs.LaunchTemplateCpuOptionsArgs>? CpuOptions { get; set; }

        /// <summary>
        /// Customize the credit specification of the instance. See Credit
        /// Specification below for more details.
        /// </summary>
        [Input("creditSpecification")]
        public Input<Inputs.LaunchTemplateCreditSpecificationArgs>? CreditSpecification { get; set; }

        /// <summary>
        /// Default Version of the launch template.
        /// </summary>
        [Input("defaultVersion")]
        public Input<int>? DefaultVersion { get; set; }

        /// <summary>
        /// Description of the launch template.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If `true`, enables [EC2 Instance
        /// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
        /// </summary>
        [Input("disableApiTermination")]
        public Input<bool>? DisableApiTermination { get; set; }

        /// <summary>
        /// If `true`, the launched EC2 instance will be EBS-optimized.
        /// </summary>
        [Input("ebsOptimized")]
        public Input<string>? EbsOptimized { get; set; }

        [Input("elasticGpuSpecifications")]
        private InputList<Inputs.LaunchTemplateElasticGpuSpecificationArgs>? _elasticGpuSpecifications;

        /// <summary>
        /// The elastic GPU to attach to the instance. See Elastic GPU
        /// below for more details.
        /// </summary>
        public InputList<Inputs.LaunchTemplateElasticGpuSpecificationArgs> ElasticGpuSpecifications
        {
            get => _elasticGpuSpecifications ?? (_elasticGpuSpecifications = new InputList<Inputs.LaunchTemplateElasticGpuSpecificationArgs>());
            set => _elasticGpuSpecifications = value;
        }

        /// <summary>
        /// Configuration block containing an Elastic Inference Accelerator to attach to the instance. See Elastic Inference Accelerator below for more details.
        /// </summary>
        [Input("elasticInferenceAccelerator")]
        public Input<Inputs.LaunchTemplateElasticInferenceAcceleratorArgs>? ElasticInferenceAccelerator { get; set; }

        /// <summary>
        /// Enable Nitro Enclaves on launched instances. See Enclave Options below for more details.
        /// </summary>
        [Input("enclaveOptions")]
        public Input<Inputs.LaunchTemplateEnclaveOptionsArgs>? EnclaveOptions { get; set; }

        /// <summary>
        /// The hibernation options for the instance. See Hibernation Options below for more details.
        /// </summary>
        [Input("hibernationOptions")]
        public Input<Inputs.LaunchTemplateHibernationOptionsArgs>? HibernationOptions { get; set; }

        /// <summary>
        /// The IAM Instance Profile to launch the instance with. See Instance Profile
        /// below for more details.
        /// </summary>
        [Input("iamInstanceProfile")]
        public Input<Inputs.LaunchTemplateIamInstanceProfileArgs>? IamInstanceProfile { get; set; }

        /// <summary>
        /// The AMI from which to launch the instance.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// Shutdown behavior for the instance. Can be `stop` or `terminate`.
        /// (Default: `stop`).
        /// </summary>
        [Input("instanceInitiatedShutdownBehavior")]
        public Input<string>? InstanceInitiatedShutdownBehavior { get; set; }

        /// <summary>
        /// The market (purchasing) option for the instance. See Market Options
        /// below for details.
        /// </summary>
        [Input("instanceMarketOptions")]
        public Input<Inputs.LaunchTemplateInstanceMarketOptionsArgs>? InstanceMarketOptions { get; set; }

        /// <summary>
        /// The type of the instance.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The kernel ID.
        /// </summary>
        [Input("kernelId")]
        public Input<string>? KernelId { get; set; }

        /// <summary>
        /// The key name to use for the instance.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        [Input("licenseSpecifications")]
        private InputList<Inputs.LaunchTemplateLicenseSpecificationArgs>? _licenseSpecifications;

        /// <summary>
        /// A list of license specifications to associate with. See License Specification below for more details.
        /// </summary>
        public InputList<Inputs.LaunchTemplateLicenseSpecificationArgs> LicenseSpecifications
        {
            get => _licenseSpecifications ?? (_licenseSpecifications = new InputList<Inputs.LaunchTemplateLicenseSpecificationArgs>());
            set => _licenseSpecifications = value;
        }

        /// <summary>
        /// Customize the metadata options for the instance. See Metadata Options below for more details.
        /// </summary>
        [Input("metadataOptions")]
        public Input<Inputs.LaunchTemplateMetadataOptionsArgs>? MetadataOptions { get; set; }

        /// <summary>
        /// The monitoring option for the instance. See Monitoring below for more details.
        /// </summary>
        [Input("monitoring")]
        public Input<Inputs.LaunchTemplateMonitoringArgs>? Monitoring { get; set; }

        /// <summary>
        /// The name of the launch template. If you leave this blank, this provider will auto-generate a unique name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.LaunchTemplateNetworkInterfaceArgs>? _networkInterfaces;

        /// <summary>
        /// Customize network interfaces to be attached at instance boot time. See Network
        /// Interfaces below for more details.
        /// </summary>
        public InputList<Inputs.LaunchTemplateNetworkInterfaceArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.LaunchTemplateNetworkInterfaceArgs>());
            set => _networkInterfaces = value;
        }

        /// <summary>
        /// The placement of the instance. See Placement below for more details.
        /// </summary>
        [Input("placement")]
        public Input<Inputs.LaunchTemplatePlacementArgs>? Placement { get; set; }

        /// <summary>
        /// The ID of the RAM disk.
        /// </summary>
        [Input("ramDiskId")]
        public Input<string>? RamDiskId { get; set; }

        [Input("securityGroupNames")]
        private InputList<string>? _securityGroupNames;

        /// <summary>
        /// A list of security group names to associate with. If you are creating Instances in a VPC, use
        /// `vpc_security_group_ids` instead.
        /// </summary>
        public InputList<string> SecurityGroupNames
        {
            get => _securityGroupNames ?? (_securityGroupNames = new InputList<string>());
            set => _securityGroupNames = value;
        }

        [Input("tagSpecifications")]
        private InputList<Inputs.LaunchTemplateTagSpecificationArgs>? _tagSpecifications;

        /// <summary>
        /// The tags to apply to the resources during launch. See Tag Specifications below for more details.
        /// </summary>
        public InputList<Inputs.LaunchTemplateTagSpecificationArgs> TagSpecifications
        {
            get => _tagSpecifications ?? (_tagSpecifications = new InputList<Inputs.LaunchTemplateTagSpecificationArgs>());
            set => _tagSpecifications = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the launch template. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// Whether to update Default Version each update. Conflicts with `default_version`.
        /// </summary>
        [Input("updateDefaultVersion")]
        public Input<bool>? UpdateDefaultVersion { get; set; }

        /// <summary>
        /// The Base64-encoded user data to provide when launching the instance.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        [Input("vpcSecurityGroupIds")]
        private InputList<string>? _vpcSecurityGroupIds;

        /// <summary>
        /// A list of security group IDs to associate with.
        /// </summary>
        public InputList<string> VpcSecurityGroupIds
        {
            get => _vpcSecurityGroupIds ?? (_vpcSecurityGroupIds = new InputList<string>());
            set => _vpcSecurityGroupIds = value;
        }

        public LaunchTemplateArgs()
        {
        }
    }

    public sealed class LaunchTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the instance profile.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("blockDeviceMappings")]
        private InputList<Inputs.LaunchTemplateBlockDeviceMappingGetArgs>? _blockDeviceMappings;

        /// <summary>
        /// Specify volumes to attach to the instance besides the volumes specified by the AMI.
        /// See Block Devices below for details.
        /// </summary>
        public InputList<Inputs.LaunchTemplateBlockDeviceMappingGetArgs> BlockDeviceMappings
        {
            get => _blockDeviceMappings ?? (_blockDeviceMappings = new InputList<Inputs.LaunchTemplateBlockDeviceMappingGetArgs>());
            set => _blockDeviceMappings = value;
        }

        /// <summary>
        /// Targeting for EC2 capacity reservations. See Capacity Reservation Specification below for more details.
        /// </summary>
        [Input("capacityReservationSpecification")]
        public Input<Inputs.LaunchTemplateCapacityReservationSpecificationGetArgs>? CapacityReservationSpecification { get; set; }

        /// <summary>
        /// The CPU options for the instance. See CPU Options below for more details.
        /// </summary>
        [Input("cpuOptions")]
        public Input<Inputs.LaunchTemplateCpuOptionsGetArgs>? CpuOptions { get; set; }

        /// <summary>
        /// Customize the credit specification of the instance. See Credit
        /// Specification below for more details.
        /// </summary>
        [Input("creditSpecification")]
        public Input<Inputs.LaunchTemplateCreditSpecificationGetArgs>? CreditSpecification { get; set; }

        /// <summary>
        /// Default Version of the launch template.
        /// </summary>
        [Input("defaultVersion")]
        public Input<int>? DefaultVersion { get; set; }

        /// <summary>
        /// Description of the launch template.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If `true`, enables [EC2 Instance
        /// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
        /// </summary>
        [Input("disableApiTermination")]
        public Input<bool>? DisableApiTermination { get; set; }

        /// <summary>
        /// If `true`, the launched EC2 instance will be EBS-optimized.
        /// </summary>
        [Input("ebsOptimized")]
        public Input<string>? EbsOptimized { get; set; }

        [Input("elasticGpuSpecifications")]
        private InputList<Inputs.LaunchTemplateElasticGpuSpecificationGetArgs>? _elasticGpuSpecifications;

        /// <summary>
        /// The elastic GPU to attach to the instance. See Elastic GPU
        /// below for more details.
        /// </summary>
        public InputList<Inputs.LaunchTemplateElasticGpuSpecificationGetArgs> ElasticGpuSpecifications
        {
            get => _elasticGpuSpecifications ?? (_elasticGpuSpecifications = new InputList<Inputs.LaunchTemplateElasticGpuSpecificationGetArgs>());
            set => _elasticGpuSpecifications = value;
        }

        /// <summary>
        /// Configuration block containing an Elastic Inference Accelerator to attach to the instance. See Elastic Inference Accelerator below for more details.
        /// </summary>
        [Input("elasticInferenceAccelerator")]
        public Input<Inputs.LaunchTemplateElasticInferenceAcceleratorGetArgs>? ElasticInferenceAccelerator { get; set; }

        /// <summary>
        /// Enable Nitro Enclaves on launched instances. See Enclave Options below for more details.
        /// </summary>
        [Input("enclaveOptions")]
        public Input<Inputs.LaunchTemplateEnclaveOptionsGetArgs>? EnclaveOptions { get; set; }

        /// <summary>
        /// The hibernation options for the instance. See Hibernation Options below for more details.
        /// </summary>
        [Input("hibernationOptions")]
        public Input<Inputs.LaunchTemplateHibernationOptionsGetArgs>? HibernationOptions { get; set; }

        /// <summary>
        /// The IAM Instance Profile to launch the instance with. See Instance Profile
        /// below for more details.
        /// </summary>
        [Input("iamInstanceProfile")]
        public Input<Inputs.LaunchTemplateIamInstanceProfileGetArgs>? IamInstanceProfile { get; set; }

        /// <summary>
        /// The AMI from which to launch the instance.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// Shutdown behavior for the instance. Can be `stop` or `terminate`.
        /// (Default: `stop`).
        /// </summary>
        [Input("instanceInitiatedShutdownBehavior")]
        public Input<string>? InstanceInitiatedShutdownBehavior { get; set; }

        /// <summary>
        /// The market (purchasing) option for the instance. See Market Options
        /// below for details.
        /// </summary>
        [Input("instanceMarketOptions")]
        public Input<Inputs.LaunchTemplateInstanceMarketOptionsGetArgs>? InstanceMarketOptions { get; set; }

        /// <summary>
        /// The type of the instance.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The kernel ID.
        /// </summary>
        [Input("kernelId")]
        public Input<string>? KernelId { get; set; }

        /// <summary>
        /// The key name to use for the instance.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// The latest version of the launch template.
        /// </summary>
        [Input("latestVersion")]
        public Input<int>? LatestVersion { get; set; }

        [Input("licenseSpecifications")]
        private InputList<Inputs.LaunchTemplateLicenseSpecificationGetArgs>? _licenseSpecifications;

        /// <summary>
        /// A list of license specifications to associate with. See License Specification below for more details.
        /// </summary>
        public InputList<Inputs.LaunchTemplateLicenseSpecificationGetArgs> LicenseSpecifications
        {
            get => _licenseSpecifications ?? (_licenseSpecifications = new InputList<Inputs.LaunchTemplateLicenseSpecificationGetArgs>());
            set => _licenseSpecifications = value;
        }

        /// <summary>
        /// Customize the metadata options for the instance. See Metadata Options below for more details.
        /// </summary>
        [Input("metadataOptions")]
        public Input<Inputs.LaunchTemplateMetadataOptionsGetArgs>? MetadataOptions { get; set; }

        /// <summary>
        /// The monitoring option for the instance. See Monitoring below for more details.
        /// </summary>
        [Input("monitoring")]
        public Input<Inputs.LaunchTemplateMonitoringGetArgs>? Monitoring { get; set; }

        /// <summary>
        /// The name of the launch template. If you leave this blank, this provider will auto-generate a unique name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.LaunchTemplateNetworkInterfaceGetArgs>? _networkInterfaces;

        /// <summary>
        /// Customize network interfaces to be attached at instance boot time. See Network
        /// Interfaces below for more details.
        /// </summary>
        public InputList<Inputs.LaunchTemplateNetworkInterfaceGetArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.LaunchTemplateNetworkInterfaceGetArgs>());
            set => _networkInterfaces = value;
        }

        /// <summary>
        /// The placement of the instance. See Placement below for more details.
        /// </summary>
        [Input("placement")]
        public Input<Inputs.LaunchTemplatePlacementGetArgs>? Placement { get; set; }

        /// <summary>
        /// The ID of the RAM disk.
        /// </summary>
        [Input("ramDiskId")]
        public Input<string>? RamDiskId { get; set; }

        [Input("securityGroupNames")]
        private InputList<string>? _securityGroupNames;

        /// <summary>
        /// A list of security group names to associate with. If you are creating Instances in a VPC, use
        /// `vpc_security_group_ids` instead.
        /// </summary>
        public InputList<string> SecurityGroupNames
        {
            get => _securityGroupNames ?? (_securityGroupNames = new InputList<string>());
            set => _securityGroupNames = value;
        }

        [Input("tagSpecifications")]
        private InputList<Inputs.LaunchTemplateTagSpecificationGetArgs>? _tagSpecifications;

        /// <summary>
        /// The tags to apply to the resources during launch. See Tag Specifications below for more details.
        /// </summary>
        public InputList<Inputs.LaunchTemplateTagSpecificationGetArgs> TagSpecifications
        {
            get => _tagSpecifications ?? (_tagSpecifications = new InputList<Inputs.LaunchTemplateTagSpecificationGetArgs>());
            set => _tagSpecifications = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the launch template. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// Whether to update Default Version each update. Conflicts with `default_version`.
        /// </summary>
        [Input("updateDefaultVersion")]
        public Input<bool>? UpdateDefaultVersion { get; set; }

        /// <summary>
        /// The Base64-encoded user data to provide when launching the instance.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        [Input("vpcSecurityGroupIds")]
        private InputList<string>? _vpcSecurityGroupIds;

        /// <summary>
        /// A list of security group IDs to associate with.
        /// </summary>
        public InputList<string> VpcSecurityGroupIds
        {
            get => _vpcSecurityGroupIds ?? (_vpcSecurityGroupIds = new InputList<string>());
            set => _vpcSecurityGroupIds = value;
        }

        public LaunchTemplateState()
        {
        }
    }
}
