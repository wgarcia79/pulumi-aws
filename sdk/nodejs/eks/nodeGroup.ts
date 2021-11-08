// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an EKS Node Group, which can provision and optionally update an Auto Scaling Group of Kubernetes worker nodes compatible with EKS. Additional documentation about this functionality can be found in the [EKS User Guide](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.eks.NodeGroup("example", {
 *     clusterName: aws_eks_cluster.example.name,
 *     nodeRoleArn: aws_iam_role.example.arn,
 *     subnetIds: aws_subnet.example.map(__item => __item.id),
 *     scalingConfig: {
 *         desiredSize: 1,
 *         maxSize: 1,
 *         minSize: 1,
 *     },
 *     updateConfig: {
 *         maxUnavailable: 2,
 *     },
 * }, {
 *     dependsOn: [
 *         aws_iam_role_policy_attachment["example-AmazonEKSWorkerNodePolicy"],
 *         aws_iam_role_policy_attachment["example-AmazonEKS_CNI_Policy"],
 *         aws_iam_role_policy_attachment["example-AmazonEC2ContainerRegistryReadOnly"],
 *     ],
 * });
 * ```
 * ### Ignoring Changes to Desired Size
 *
 * You can utilize [ignoreChanges](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) create an EKS Node Group with an initial size of running instances, then ignore any changes to that count caused externally (e.g. Application Autoscaling).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // ... other configurations ...
 * const example = new aws.eks.NodeGroup("example", {scalingConfig: {
 *     desiredSize: 2,
 * }});
 * ```
 * ### Example IAM Role for EKS Node Group
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.iam.Role("example", {assumeRolePolicy: JSON.stringify({
 *     Statement: [{
 *         Action: "sts:AssumeRole",
 *         Effect: "Allow",
 *         Principal: {
 *             Service: "ec2.amazonaws.com",
 *         },
 *     }],
 *     Version: "2012-10-17",
 * })});
 * const example_AmazonEKSWorkerNodePolicy = new aws.iam.RolePolicyAttachment("example-AmazonEKSWorkerNodePolicy", {
 *     policyArn: "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
 *     role: example.name,
 * });
 * const example_AmazonEKSCNIPolicy = new aws.iam.RolePolicyAttachment("example-AmazonEKSCNIPolicy", {
 *     policyArn: "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
 *     role: example.name,
 * });
 * const example_AmazonEC2ContainerRegistryReadOnly = new aws.iam.RolePolicyAttachment("example-AmazonEC2ContainerRegistryReadOnly", {
 *     policyArn: "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
 *     role: example.name,
 * });
 * ```
 *
 * ## Import
 *
 * EKS Node Groups can be imported using the `cluster_name` and `node_group_name` separated by a colon (`:`), e.g.,
 *
 * ```sh
 *  $ pulumi import aws:eks/nodeGroup:NodeGroup my_node_group my_cluster:my_node_group
 * ```
 */
export class NodeGroup extends pulumi.CustomResource {
    /**
     * Get an existing NodeGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NodeGroupState, opts?: pulumi.CustomResourceOptions): NodeGroup {
        return new NodeGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:eks/nodeGroup:NodeGroup';

    /**
     * Returns true if the given object is an instance of NodeGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NodeGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NodeGroup.__pulumiType;
    }

    /**
     * Type of Amazon Machine Image (AMI) associated with the EKS Node Group. See the [AWS documentation](https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType) for valid values. This provider will only perform drift detection if a configuration value is provided.
     */
    public readonly amiType!: pulumi.Output<string>;
    /**
     * Amazon Resource Name (ARN) of the EKS Node Group.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Type of capacity associated with the EKS Node Group. Valid values: `ON_DEMAND`, `SPOT`. This provider will only perform drift detection if a configuration value is provided.
     */
    public readonly capacityType!: pulumi.Output<string>;
    /**
     * Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * Disk size in GiB for worker nodes. Defaults to `20`. This provider will only perform drift detection if a configuration value is provided.
     */
    public readonly diskSize!: pulumi.Output<number>;
    /**
     * Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
     */
    public readonly forceUpdateVersion!: pulumi.Output<boolean | undefined>;
    /**
     * Set of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. This provider will only perform drift detection if a configuration value is provided.
     */
    public readonly instanceTypes!: pulumi.Output<string[]>;
    /**
     * Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Configuration block with Launch Template settings. Detailed below.
     */
    public readonly launchTemplate!: pulumi.Output<outputs.eks.NodeGroupLaunchTemplate | undefined>;
    /**
     * Name of the EKS Node Group. If omitted, this provider will assign a random, unique name. Conflicts with `nodeGroupNamePrefix`.
     */
    public readonly nodeGroupName!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `nodeGroupName`.
     */
    public readonly nodeGroupNamePrefix!: pulumi.Output<string>;
    /**
     * Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
     */
    public readonly nodeRoleArn!: pulumi.Output<string>;
    /**
     * AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
     */
    public readonly releaseVersion!: pulumi.Output<string>;
    /**
     * Configuration block with remote access settings. Detailed below.
     */
    public readonly remoteAccess!: pulumi.Output<outputs.eks.NodeGroupRemoteAccess | undefined>;
    /**
     * List of objects containing information about underlying resources.
     */
    public /*out*/ readonly resources!: pulumi.Output<outputs.eks.NodeGroupResource[]>;
    /**
     * Configuration block with scaling settings. Detailed below.
     */
    public readonly scalingConfig!: pulumi.Output<outputs.eks.NodeGroupScalingConfig>;
    /**
     * Status of the EKS Node Group.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
     */
    public readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * Key-value map of resource tags. If configured with a provider defaultTags present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. Detailed below.
     */
    public readonly taints!: pulumi.Output<outputs.eks.NodeGroupTaint[] | undefined>;
    public readonly updateConfig!: pulumi.Output<outputs.eks.NodeGroupUpdateConfig>;
    /**
     * EC2 Launch Template version number. While the API accepts values like `$Default` and `$Latest`, the API will convert the value to the associated version number (e.g. `1`) on read and This provider will show a difference on next plan. Using the `defaultVersion` or `latestVersion` attribute of the `aws.ec2.LaunchTemplate` resource or data source is recommended for this argument.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a NodeGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodeGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodeGroupArgs | NodeGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NodeGroupState | undefined;
            inputs["amiType"] = state ? state.amiType : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["capacityType"] = state ? state.capacityType : undefined;
            inputs["clusterName"] = state ? state.clusterName : undefined;
            inputs["diskSize"] = state ? state.diskSize : undefined;
            inputs["forceUpdateVersion"] = state ? state.forceUpdateVersion : undefined;
            inputs["instanceTypes"] = state ? state.instanceTypes : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["launchTemplate"] = state ? state.launchTemplate : undefined;
            inputs["nodeGroupName"] = state ? state.nodeGroupName : undefined;
            inputs["nodeGroupNamePrefix"] = state ? state.nodeGroupNamePrefix : undefined;
            inputs["nodeRoleArn"] = state ? state.nodeRoleArn : undefined;
            inputs["releaseVersion"] = state ? state.releaseVersion : undefined;
            inputs["remoteAccess"] = state ? state.remoteAccess : undefined;
            inputs["resources"] = state ? state.resources : undefined;
            inputs["scalingConfig"] = state ? state.scalingConfig : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["taints"] = state ? state.taints : undefined;
            inputs["updateConfig"] = state ? state.updateConfig : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as NodeGroupArgs | undefined;
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            if ((!args || args.nodeRoleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeRoleArn'");
            }
            if ((!args || args.scalingConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalingConfig'");
            }
            if ((!args || args.subnetIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetIds'");
            }
            inputs["amiType"] = args ? args.amiType : undefined;
            inputs["capacityType"] = args ? args.capacityType : undefined;
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["diskSize"] = args ? args.diskSize : undefined;
            inputs["forceUpdateVersion"] = args ? args.forceUpdateVersion : undefined;
            inputs["instanceTypes"] = args ? args.instanceTypes : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["launchTemplate"] = args ? args.launchTemplate : undefined;
            inputs["nodeGroupName"] = args ? args.nodeGroupName : undefined;
            inputs["nodeGroupNamePrefix"] = args ? args.nodeGroupNamePrefix : undefined;
            inputs["nodeRoleArn"] = args ? args.nodeRoleArn : undefined;
            inputs["releaseVersion"] = args ? args.releaseVersion : undefined;
            inputs["remoteAccess"] = args ? args.remoteAccess : undefined;
            inputs["scalingConfig"] = args ? args.scalingConfig : undefined;
            inputs["subnetIds"] = args ? args.subnetIds : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["taints"] = args ? args.taints : undefined;
            inputs["updateConfig"] = args ? args.updateConfig : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["resources"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["tagsAll"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NodeGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NodeGroup resources.
 */
export interface NodeGroupState {
    /**
     * Type of Amazon Machine Image (AMI) associated with the EKS Node Group. See the [AWS documentation](https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType) for valid values. This provider will only perform drift detection if a configuration value is provided.
     */
    amiType?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the EKS Node Group.
     */
    arn?: pulumi.Input<string>;
    /**
     * Type of capacity associated with the EKS Node Group. Valid values: `ON_DEMAND`, `SPOT`. This provider will only perform drift detection if a configuration value is provided.
     */
    capacityType?: pulumi.Input<string>;
    /**
     * Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
     */
    clusterName?: pulumi.Input<string>;
    /**
     * Disk size in GiB for worker nodes. Defaults to `20`. This provider will only perform drift detection if a configuration value is provided.
     */
    diskSize?: pulumi.Input<number>;
    /**
     * Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
     */
    forceUpdateVersion?: pulumi.Input<boolean>;
    /**
     * Set of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. This provider will only perform drift detection if a configuration value is provided.
     */
    instanceTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block with Launch Template settings. Detailed below.
     */
    launchTemplate?: pulumi.Input<inputs.eks.NodeGroupLaunchTemplate>;
    /**
     * Name of the EKS Node Group. If omitted, this provider will assign a random, unique name. Conflicts with `nodeGroupNamePrefix`.
     */
    nodeGroupName?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `nodeGroupName`.
     */
    nodeGroupNamePrefix?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
     */
    nodeRoleArn?: pulumi.Input<string>;
    /**
     * AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
     */
    releaseVersion?: pulumi.Input<string>;
    /**
     * Configuration block with remote access settings. Detailed below.
     */
    remoteAccess?: pulumi.Input<inputs.eks.NodeGroupRemoteAccess>;
    /**
     * List of objects containing information about underlying resources.
     */
    resources?: pulumi.Input<pulumi.Input<inputs.eks.NodeGroupResource>[]>;
    /**
     * Configuration block with scaling settings. Detailed below.
     */
    scalingConfig?: pulumi.Input<inputs.eks.NodeGroupScalingConfig>;
    /**
     * Status of the EKS Node Group.
     */
    status?: pulumi.Input<string>;
    /**
     * Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key-value map of resource tags. If configured with a provider defaultTags present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. Detailed below.
     */
    taints?: pulumi.Input<pulumi.Input<inputs.eks.NodeGroupTaint>[]>;
    updateConfig?: pulumi.Input<inputs.eks.NodeGroupUpdateConfig>;
    /**
     * EC2 Launch Template version number. While the API accepts values like `$Default` and `$Latest`, the API will convert the value to the associated version number (e.g. `1`) on read and This provider will show a difference on next plan. Using the `defaultVersion` or `latestVersion` attribute of the `aws.ec2.LaunchTemplate` resource or data source is recommended for this argument.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NodeGroup resource.
 */
export interface NodeGroupArgs {
    /**
     * Type of Amazon Machine Image (AMI) associated with the EKS Node Group. See the [AWS documentation](https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType) for valid values. This provider will only perform drift detection if a configuration value is provided.
     */
    amiType?: pulumi.Input<string>;
    /**
     * Type of capacity associated with the EKS Node Group. Valid values: `ON_DEMAND`, `SPOT`. This provider will only perform drift detection if a configuration value is provided.
     */
    capacityType?: pulumi.Input<string>;
    /**
     * Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
     */
    clusterName: pulumi.Input<string>;
    /**
     * Disk size in GiB for worker nodes. Defaults to `20`. This provider will only perform drift detection if a configuration value is provided.
     */
    diskSize?: pulumi.Input<number>;
    /**
     * Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
     */
    forceUpdateVersion?: pulumi.Input<boolean>;
    /**
     * Set of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. This provider will only perform drift detection if a configuration value is provided.
     */
    instanceTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block with Launch Template settings. Detailed below.
     */
    launchTemplate?: pulumi.Input<inputs.eks.NodeGroupLaunchTemplate>;
    /**
     * Name of the EKS Node Group. If omitted, this provider will assign a random, unique name. Conflicts with `nodeGroupNamePrefix`.
     */
    nodeGroupName?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `nodeGroupName`.
     */
    nodeGroupNamePrefix?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
     */
    nodeRoleArn: pulumi.Input<string>;
    /**
     * AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
     */
    releaseVersion?: pulumi.Input<string>;
    /**
     * Configuration block with remote access settings. Detailed below.
     */
    remoteAccess?: pulumi.Input<inputs.eks.NodeGroupRemoteAccess>;
    /**
     * Configuration block with scaling settings. Detailed below.
     */
    scalingConfig: pulumi.Input<inputs.eks.NodeGroupScalingConfig>;
    /**
     * Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
     */
    subnetIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key-value map of resource tags. If configured with a provider defaultTags present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. Detailed below.
     */
    taints?: pulumi.Input<pulumi.Input<inputs.eks.NodeGroupTaint>[]>;
    updateConfig?: pulumi.Input<inputs.eks.NodeGroupUpdateConfig>;
    /**
     * EC2 Launch Template version number. While the API accepts values like `$Default` and `$Latest`, the API will convert the value to the associated version number (e.g. `1`) on read and This provider will show a difference on next plan. Using the `defaultVersion` or `latestVersion` attribute of the `aws.ec2.LaunchTemplate` resource or data source is recommended for this argument.
     */
    version?: pulumi.Input<string>;
}
