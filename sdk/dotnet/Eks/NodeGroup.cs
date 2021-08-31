// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Eks
{
    /// <summary>
    /// Manages an EKS Node Group, which can provision and optionally update an Auto Scaling Group of Kubernetes worker nodes compatible with EKS. Additional documentation about this functionality can be found in the [EKS User Guide](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Eks.NodeGroup("example", new Aws.Eks.NodeGroupArgs
    ///         {
    ///             ClusterName = aws_eks_cluster.Example.Name,
    ///             NodeRoleArn = aws_iam_role.Example.Arn,
    ///             SubnetIds = aws_subnet.Example.Select(__item =&gt; __item.Id).ToList(),
    ///             ScalingConfig = new Aws.Eks.Inputs.NodeGroupScalingConfigArgs
    ///             {
    ///                 DesiredSize = 1,
    ///                 MaxSize = 1,
    ///                 MinSize = 1,
    ///             },
    ///             UpdateConfig = new Aws.Eks.Inputs.NodeGroupUpdateConfigArgs
    ///             {
    ///                 MaxUnavailable = 2,
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 aws_iam_role_policy_attachment.Example_AmazonEKSWorkerNodePolicy,
    ///                 aws_iam_role_policy_attachment.Example_AmazonEKS_CNI_Policy,
    ///                 aws_iam_role_policy_attachment.Example_AmazonEC2ContainerRegistryReadOnly,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Ignoring Changes to Desired Size
    /// 
    /// You can utilize [ignoreChanges](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) create an EKS Node Group with an initial size of running instances, then ignore any changes to that count caused externally (e.g. Application Autoscaling).
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // ... other configurations ...
    ///         var example = new Aws.Eks.NodeGroup("example", new Aws.Eks.NodeGroupArgs
    ///         {
    ///             ScalingConfig = new Aws.Eks.Inputs.NodeGroupScalingConfigArgs
    ///             {
    ///                 DesiredSize = 2,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Example IAM Role for EKS Node Group
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Iam.Role("example", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 { "Statement", new[]
    ///                     {
    ///                         new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             { "Action", "sts:AssumeRole" },
    ///                             { "Effect", "Allow" },
    ///                             { "Principal", new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 { "Service", "ec2.amazonaws.com" },
    ///                             } },
    ///                         },
    ///                     }
    ///                  },
    ///                 { "Version", "2012-10-17" },
    ///             }),
    ///         });
    ///         var example_AmazonEKSWorkerNodePolicy = new Aws.Iam.RolePolicyAttachment("example-AmazonEKSWorkerNodePolicy", new Aws.Iam.RolePolicyAttachmentArgs
    ///         {
    ///             PolicyArn = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
    ///             Role = example.Name,
    ///         });
    ///         var example_AmazonEKSCNIPolicy = new Aws.Iam.RolePolicyAttachment("example-AmazonEKSCNIPolicy", new Aws.Iam.RolePolicyAttachmentArgs
    ///         {
    ///             PolicyArn = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
    ///             Role = example.Name,
    ///         });
    ///         var example_AmazonEC2ContainerRegistryReadOnly = new Aws.Iam.RolePolicyAttachment("example-AmazonEC2ContainerRegistryReadOnly", new Aws.Iam.RolePolicyAttachmentArgs
    ///         {
    ///             PolicyArn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
    ///             Role = example.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// EKS Node Groups can be imported using the `cluster_name` and `node_group_name` separated by a colon (`:`), e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:eks/nodeGroup:NodeGroup my_node_group my_cluster:my_node_group
    /// ```
    /// </summary>
    [AwsResourceType("aws:eks/nodeGroup:NodeGroup")]
    public partial class NodeGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Type of Amazon Machine Image (AMI) associated with the EKS Node Group. Defaults to `AL2_x86_64`. Valid values: `AL2_x86_64`, `AL2_x86_64_GPU`, `AL2_ARM_64`, `CUSTOM`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Output("amiType")]
        public Output<string> AmiType { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the EKS Node Group.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Type of capacity associated with the EKS Node Group. Valid values: `ON_DEMAND`, `SPOT`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Output("capacityType")]
        public Output<string> CapacityType { get; private set; } = null!;

        /// <summary>
        /// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// Disk size in GiB for worker nodes. Defaults to `20`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Output("diskSize")]
        public Output<int> DiskSize { get; private set; } = null!;

        /// <summary>
        /// Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
        /// </summary>
        [Output("forceUpdateVersion")]
        public Output<bool?> ForceUpdateVersion { get; private set; } = null!;

        /// <summary>
        /// Set of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Output("instanceTypes")]
        public Output<ImmutableArray<string>> InstanceTypes { get; private set; } = null!;

        /// <summary>
        /// Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Configuration block with Launch Template settings. Detailed below.
        /// </summary>
        [Output("launchTemplate")]
        public Output<Outputs.NodeGroupLaunchTemplate?> LaunchTemplate { get; private set; } = null!;

        /// <summary>
        /// Name of the EKS Node Group. If omitted, this provider will assign a random, unique name. Conflicts with `node_group_name_prefix`.
        /// </summary>
        [Output("nodeGroupName")]
        public Output<string> NodeGroupName { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `node_group_name`.
        /// </summary>
        [Output("nodeGroupNamePrefix")]
        public Output<string> NodeGroupNamePrefix { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
        /// </summary>
        [Output("nodeRoleArn")]
        public Output<string> NodeRoleArn { get; private set; } = null!;

        /// <summary>
        /// AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
        /// </summary>
        [Output("releaseVersion")]
        public Output<string> ReleaseVersion { get; private set; } = null!;

        /// <summary>
        /// Configuration block with remote access settings. Detailed below.
        /// </summary>
        [Output("remoteAccess")]
        public Output<Outputs.NodeGroupRemoteAccess?> RemoteAccess { get; private set; } = null!;

        /// <summary>
        /// List of objects containing information about underlying resources.
        /// </summary>
        [Output("resources")]
        public Output<ImmutableArray<Outputs.NodeGroupResource>> Resources { get; private set; } = null!;

        /// <summary>
        /// Configuration block with scaling settings. Detailed below.
        /// </summary>
        [Output("scalingConfig")]
        public Output<Outputs.NodeGroupScalingConfig> ScalingConfig { get; private set; } = null!;

        /// <summary>
        /// Status of the EKS Node Group.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider defaultTags present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. Detailed below.
        /// </summary>
        [Output("taints")]
        public Output<ImmutableArray<Outputs.NodeGroupTaint>> Taints { get; private set; } = null!;

        [Output("updateConfig")]
        public Output<Outputs.NodeGroupUpdateConfig> UpdateConfig { get; private set; } = null!;

        /// <summary>
        /// EC2 Launch Template version number. While the API accepts values like `$Default` and `$Latest`, the API will convert the value to the associated version number (e.g. `1`) on read and This provider will show a difference on next plan. Using the `default_version` or `latest_version` attribute of the `aws.ec2.LaunchTemplate` resource or data source is recommended for this argument.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a NodeGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NodeGroup(string name, NodeGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:eks/nodeGroup:NodeGroup", name, args ?? new NodeGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NodeGroup(string name, Input<string> id, NodeGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:eks/nodeGroup:NodeGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NodeGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NodeGroup Get(string name, Input<string> id, NodeGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new NodeGroup(name, id, state, options);
        }
    }

    public sealed class NodeGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of Amazon Machine Image (AMI) associated with the EKS Node Group. Defaults to `AL2_x86_64`. Valid values: `AL2_x86_64`, `AL2_x86_64_GPU`, `AL2_ARM_64`, `CUSTOM`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Input("amiType")]
        public Input<string>? AmiType { get; set; }

        /// <summary>
        /// Type of capacity associated with the EKS Node Group. Valid values: `ON_DEMAND`, `SPOT`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Input("capacityType")]
        public Input<string>? CapacityType { get; set; }

        /// <summary>
        /// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// Disk size in GiB for worker nodes. Defaults to `20`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
        /// </summary>
        [Input("forceUpdateVersion")]
        public Input<bool>? ForceUpdateVersion { get; set; }

        [Input("instanceTypes")]
        private InputList<string>? _instanceTypes;

        /// <summary>
        /// Set of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        public InputList<string> InstanceTypes
        {
            get => _instanceTypes ?? (_instanceTypes = new InputList<string>());
            set => _instanceTypes = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Configuration block with Launch Template settings. Detailed below.
        /// </summary>
        [Input("launchTemplate")]
        public Input<Inputs.NodeGroupLaunchTemplateArgs>? LaunchTemplate { get; set; }

        /// <summary>
        /// Name of the EKS Node Group. If omitted, this provider will assign a random, unique name. Conflicts with `node_group_name_prefix`.
        /// </summary>
        [Input("nodeGroupName")]
        public Input<string>? NodeGroupName { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `node_group_name`.
        /// </summary>
        [Input("nodeGroupNamePrefix")]
        public Input<string>? NodeGroupNamePrefix { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
        /// </summary>
        [Input("nodeRoleArn", required: true)]
        public Input<string> NodeRoleArn { get; set; } = null!;

        /// <summary>
        /// AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
        /// </summary>
        [Input("releaseVersion")]
        public Input<string>? ReleaseVersion { get; set; }

        /// <summary>
        /// Configuration block with remote access settings. Detailed below.
        /// </summary>
        [Input("remoteAccess")]
        public Input<Inputs.NodeGroupRemoteAccessArgs>? RemoteAccess { get; set; }

        /// <summary>
        /// Configuration block with scaling settings. Detailed below.
        /// </summary>
        [Input("scalingConfig", required: true)]
        public Input<Inputs.NodeGroupScalingConfigArgs> ScalingConfig { get; set; } = null!;

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider defaultTags present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("taints")]
        private InputList<Inputs.NodeGroupTaintArgs>? _taints;

        /// <summary>
        /// The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. Detailed below.
        /// </summary>
        public InputList<Inputs.NodeGroupTaintArgs> Taints
        {
            get => _taints ?? (_taints = new InputList<Inputs.NodeGroupTaintArgs>());
            set => _taints = value;
        }

        [Input("updateConfig")]
        public Input<Inputs.NodeGroupUpdateConfigArgs>? UpdateConfig { get; set; }

        /// <summary>
        /// EC2 Launch Template version number. While the API accepts values like `$Default` and `$Latest`, the API will convert the value to the associated version number (e.g. `1`) on read and This provider will show a difference on next plan. Using the `default_version` or `latest_version` attribute of the `aws.ec2.LaunchTemplate` resource or data source is recommended for this argument.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public NodeGroupArgs()
        {
        }
    }

    public sealed class NodeGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of Amazon Machine Image (AMI) associated with the EKS Node Group. Defaults to `AL2_x86_64`. Valid values: `AL2_x86_64`, `AL2_x86_64_GPU`, `AL2_ARM_64`, `CUSTOM`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Input("amiType")]
        public Input<string>? AmiType { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the EKS Node Group.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Type of capacity associated with the EKS Node Group. Valid values: `ON_DEMAND`, `SPOT`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Input("capacityType")]
        public Input<string>? CapacityType { get; set; }

        /// <summary>
        /// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// Disk size in GiB for worker nodes. Defaults to `20`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
        /// </summary>
        [Input("forceUpdateVersion")]
        public Input<bool>? ForceUpdateVersion { get; set; }

        [Input("instanceTypes")]
        private InputList<string>? _instanceTypes;

        /// <summary>
        /// Set of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        public InputList<string> InstanceTypes
        {
            get => _instanceTypes ?? (_instanceTypes = new InputList<string>());
            set => _instanceTypes = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Configuration block with Launch Template settings. Detailed below.
        /// </summary>
        [Input("launchTemplate")]
        public Input<Inputs.NodeGroupLaunchTemplateGetArgs>? LaunchTemplate { get; set; }

        /// <summary>
        /// Name of the EKS Node Group. If omitted, this provider will assign a random, unique name. Conflicts with `node_group_name_prefix`.
        /// </summary>
        [Input("nodeGroupName")]
        public Input<string>? NodeGroupName { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `node_group_name`.
        /// </summary>
        [Input("nodeGroupNamePrefix")]
        public Input<string>? NodeGroupNamePrefix { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
        /// </summary>
        [Input("nodeRoleArn")]
        public Input<string>? NodeRoleArn { get; set; }

        /// <summary>
        /// AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
        /// </summary>
        [Input("releaseVersion")]
        public Input<string>? ReleaseVersion { get; set; }

        /// <summary>
        /// Configuration block with remote access settings. Detailed below.
        /// </summary>
        [Input("remoteAccess")]
        public Input<Inputs.NodeGroupRemoteAccessGetArgs>? RemoteAccess { get; set; }

        [Input("resources")]
        private InputList<Inputs.NodeGroupResourceGetArgs>? _resources;

        /// <summary>
        /// List of objects containing information about underlying resources.
        /// </summary>
        public InputList<Inputs.NodeGroupResourceGetArgs> Resources
        {
            get => _resources ?? (_resources = new InputList<Inputs.NodeGroupResourceGetArgs>());
            set => _resources = value;
        }

        /// <summary>
        /// Configuration block with scaling settings. Detailed below.
        /// </summary>
        [Input("scalingConfig")]
        public Input<Inputs.NodeGroupScalingConfigGetArgs>? ScalingConfig { get; set; }

        /// <summary>
        /// Status of the EKS Node Group.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider defaultTags present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        [Input("taints")]
        private InputList<Inputs.NodeGroupTaintGetArgs>? _taints;

        /// <summary>
        /// The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. Detailed below.
        /// </summary>
        public InputList<Inputs.NodeGroupTaintGetArgs> Taints
        {
            get => _taints ?? (_taints = new InputList<Inputs.NodeGroupTaintGetArgs>());
            set => _taints = value;
        }

        [Input("updateConfig")]
        public Input<Inputs.NodeGroupUpdateConfigGetArgs>? UpdateConfig { get; set; }

        /// <summary>
        /// EC2 Launch Template version number. While the API accepts values like `$Default` and `$Latest`, the API will convert the value to the associated version number (e.g. `1`) on read and This provider will show a difference on next plan. Using the `default_version` or `latest_version` attribute of the `aws.ec2.LaunchTemplate` resource or data source is recommended for this argument.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public NodeGroupState()
        {
        }
    }
}
