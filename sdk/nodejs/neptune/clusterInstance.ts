// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A Cluster Instance Resource defines attributes that are specific to a single instance in a Neptune Cluster.
 *
 * You can simply add neptune instances and Neptune manages the replication. You can use the [count](https://www.terraform.io/docs/configuration/meta-arguments/count.html)
 * meta-parameter to make multiple instances and join them all to the same Neptune Cluster, or you may specify different Cluster Instance resources with various `instanceClass` sizes.
 *
 * ## Example Usage
 *
 * The following example will create a neptune cluster with two neptune instances(one writer and one reader).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const _default = new aws.neptune.Cluster("default", {
 *     clusterIdentifier: "neptune-cluster-demo",
 *     engine: "neptune",
 *     backupRetentionPeriod: 5,
 *     preferredBackupWindow: "07:00-09:00",
 *     skipFinalSnapshot: true,
 *     iamDatabaseAuthenticationEnabled: true,
 *     applyImmediately: true,
 * });
 * const example: aws.neptune.ClusterInstance[];
 * for (const range = {value: 0}; range.value < 2; range.value++) {
 *     example.push(new aws.neptune.ClusterInstance(`example-${range.value}`, {
 *         clusterIdentifier: _default.id,
 *         engine: "neptune",
 *         instanceClass: "db.r4.large",
 *         applyImmediately: true,
 *     }));
 * }
 * ```
 *
 * ## Import
 *
 * `aws_neptune_cluster_instance` can be imported by using the instance identifier, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:neptune/clusterInstance:ClusterInstance example my-instance
 * ```
 */
export class ClusterInstance extends pulumi.CustomResource {
    /**
     * Get an existing ClusterInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterInstanceState, opts?: pulumi.CustomResourceOptions): ClusterInstance {
        return new ClusterInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:neptune/clusterInstance:ClusterInstance';

    /**
     * Returns true if the given object is an instance of ClusterInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterInstance.__pulumiType;
    }

    /**
     * The hostname of the instance. See also `endpoint` and `port`.
     */
    public /*out*/ readonly address!: pulumi.Output<string>;
    /**
     * Specifies whether any instance modifications
     * are applied immediately, or during the next maintenance window. Default is`false`.
     */
    public readonly applyImmediately!: pulumi.Output<boolean>;
    /**
     * Amazon Resource Name (ARN) of neptune instance
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Indicates that minor engine upgrades will be applied automatically to the instance during the maintenance window. Default is `true`.
     */
    public readonly autoMinorVersionUpgrade!: pulumi.Output<boolean | undefined>;
    /**
     * The EC2 Availability Zone that the neptune instance is created in.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * The identifier of the `aws.neptune.Cluster` in which to launch this instance.
     */
    public readonly clusterIdentifier!: pulumi.Output<string>;
    /**
     * The region-unique, immutable identifier for the neptune instance.
     */
    public /*out*/ readonly dbiResourceId!: pulumi.Output<string>;
    /**
     * The connection endpoint in `address:port` format.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * The name of the database engine to be used for the neptune instance. Defaults to `neptune`. Valid Values: `neptune`.
     */
    public readonly engine!: pulumi.Output<string | undefined>;
    /**
     * The neptune engine version.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * The identifier for the neptune instance, if omitted, this provider will assign a random, unique identifier.
     */
    public readonly identifier!: pulumi.Output<string>;
    /**
     * Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
     */
    public readonly identifierPrefix!: pulumi.Output<string>;
    /**
     * The instance class to use.
     */
    public readonly instanceClass!: pulumi.Output<string>;
    /**
     * The ARN for the KMS encryption key if one is set to the neptune cluster.
     */
    public /*out*/ readonly kmsKeyArn!: pulumi.Output<string>;
    /**
     * The name of the neptune parameter group to associate with this instance.
     */
    public readonly neptuneParameterGroupName!: pulumi.Output<string | undefined>;
    /**
     * A subnet group to associate with this neptune instance. **NOTE:** This must match the `neptuneSubnetGroupName` of the attached `aws.neptune.Cluster`.
     */
    public readonly neptuneSubnetGroupName!: pulumi.Output<string>;
    /**
     * The port on which the DB accepts connections. Defaults to `8182`.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * The daily time range during which automated backups are created if automated backups are enabled. Eg: "04:00-09:00"
     */
    public readonly preferredBackupWindow!: pulumi.Output<string>;
    /**
     * The window to perform maintenance in.
     * Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
     */
    public readonly preferredMaintenanceWindow!: pulumi.Output<string>;
    /**
     * Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
     */
    public readonly promotionTier!: pulumi.Output<number | undefined>;
    /**
     * Bool to control if instance is publicly accessible. Default is `false`.
     */
    public readonly publiclyAccessible!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether the neptune cluster is encrypted.
     */
    public /*out*/ readonly storageEncrypted!: pulumi.Output<boolean>;
    /**
     * A map of tags to assign to the instance. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
     */
    public /*out*/ readonly writer!: pulumi.Output<boolean>;

    /**
     * Create a ClusterInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterInstanceArgs | ClusterInstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterInstanceState | undefined;
            inputs["address"] = state ? state.address : undefined;
            inputs["applyImmediately"] = state ? state.applyImmediately : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["autoMinorVersionUpgrade"] = state ? state.autoMinorVersionUpgrade : undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["clusterIdentifier"] = state ? state.clusterIdentifier : undefined;
            inputs["dbiResourceId"] = state ? state.dbiResourceId : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["engineVersion"] = state ? state.engineVersion : undefined;
            inputs["identifier"] = state ? state.identifier : undefined;
            inputs["identifierPrefix"] = state ? state.identifierPrefix : undefined;
            inputs["instanceClass"] = state ? state.instanceClass : undefined;
            inputs["kmsKeyArn"] = state ? state.kmsKeyArn : undefined;
            inputs["neptuneParameterGroupName"] = state ? state.neptuneParameterGroupName : undefined;
            inputs["neptuneSubnetGroupName"] = state ? state.neptuneSubnetGroupName : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["preferredBackupWindow"] = state ? state.preferredBackupWindow : undefined;
            inputs["preferredMaintenanceWindow"] = state ? state.preferredMaintenanceWindow : undefined;
            inputs["promotionTier"] = state ? state.promotionTier : undefined;
            inputs["publiclyAccessible"] = state ? state.publiclyAccessible : undefined;
            inputs["storageEncrypted"] = state ? state.storageEncrypted : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["writer"] = state ? state.writer : undefined;
        } else {
            const args = argsOrState as ClusterInstanceArgs | undefined;
            if ((!args || args.clusterIdentifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterIdentifier'");
            }
            if ((!args || args.instanceClass === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceClass'");
            }
            inputs["applyImmediately"] = args ? args.applyImmediately : undefined;
            inputs["autoMinorVersionUpgrade"] = args ? args.autoMinorVersionUpgrade : undefined;
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["clusterIdentifier"] = args ? args.clusterIdentifier : undefined;
            inputs["engine"] = args ? args.engine : undefined;
            inputs["engineVersion"] = args ? args.engineVersion : undefined;
            inputs["identifier"] = args ? args.identifier : undefined;
            inputs["identifierPrefix"] = args ? args.identifierPrefix : undefined;
            inputs["instanceClass"] = args ? args.instanceClass : undefined;
            inputs["neptuneParameterGroupName"] = args ? args.neptuneParameterGroupName : undefined;
            inputs["neptuneSubnetGroupName"] = args ? args.neptuneSubnetGroupName : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["preferredBackupWindow"] = args ? args.preferredBackupWindow : undefined;
            inputs["preferredMaintenanceWindow"] = args ? args.preferredMaintenanceWindow : undefined;
            inputs["promotionTier"] = args ? args.promotionTier : undefined;
            inputs["publiclyAccessible"] = args ? args.publiclyAccessible : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["address"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["dbiResourceId"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
            inputs["kmsKeyArn"] = undefined /*out*/;
            inputs["storageEncrypted"] = undefined /*out*/;
            inputs["tagsAll"] = undefined /*out*/;
            inputs["writer"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ClusterInstance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClusterInstance resources.
 */
export interface ClusterInstanceState {
    /**
     * The hostname of the instance. See also `endpoint` and `port`.
     */
    address?: pulumi.Input<string>;
    /**
     * Specifies whether any instance modifications
     * are applied immediately, or during the next maintenance window. Default is`false`.
     */
    applyImmediately?: pulumi.Input<boolean>;
    /**
     * Amazon Resource Name (ARN) of neptune instance
     */
    arn?: pulumi.Input<string>;
    /**
     * Indicates that minor engine upgrades will be applied automatically to the instance during the maintenance window. Default is `true`.
     */
    autoMinorVersionUpgrade?: pulumi.Input<boolean>;
    /**
     * The EC2 Availability Zone that the neptune instance is created in.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The identifier of the `aws.neptune.Cluster` in which to launch this instance.
     */
    clusterIdentifier?: pulumi.Input<string>;
    /**
     * The region-unique, immutable identifier for the neptune instance.
     */
    dbiResourceId?: pulumi.Input<string>;
    /**
     * The connection endpoint in `address:port` format.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * The name of the database engine to be used for the neptune instance. Defaults to `neptune`. Valid Values: `neptune`.
     */
    engine?: pulumi.Input<string>;
    /**
     * The neptune engine version.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * The identifier for the neptune instance, if omitted, this provider will assign a random, unique identifier.
     */
    identifier?: pulumi.Input<string>;
    /**
     * Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
     */
    identifierPrefix?: pulumi.Input<string>;
    /**
     * The instance class to use.
     */
    instanceClass?: pulumi.Input<string>;
    /**
     * The ARN for the KMS encryption key if one is set to the neptune cluster.
     */
    kmsKeyArn?: pulumi.Input<string>;
    /**
     * The name of the neptune parameter group to associate with this instance.
     */
    neptuneParameterGroupName?: pulumi.Input<string>;
    /**
     * A subnet group to associate with this neptune instance. **NOTE:** This must match the `neptuneSubnetGroupName` of the attached `aws.neptune.Cluster`.
     */
    neptuneSubnetGroupName?: pulumi.Input<string>;
    /**
     * The port on which the DB accepts connections. Defaults to `8182`.
     */
    port?: pulumi.Input<number>;
    /**
     * The daily time range during which automated backups are created if automated backups are enabled. Eg: "04:00-09:00"
     */
    preferredBackupWindow?: pulumi.Input<string>;
    /**
     * The window to perform maintenance in.
     * Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
     */
    preferredMaintenanceWindow?: pulumi.Input<string>;
    /**
     * Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
     */
    promotionTier?: pulumi.Input<number>;
    /**
     * Bool to control if instance is publicly accessible. Default is `false`.
     */
    publiclyAccessible?: pulumi.Input<boolean>;
    /**
     * Specifies whether the neptune cluster is encrypted.
     */
    storageEncrypted?: pulumi.Input<boolean>;
    /**
     * A map of tags to assign to the instance. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
     */
    writer?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ClusterInstance resource.
 */
export interface ClusterInstanceArgs {
    /**
     * Specifies whether any instance modifications
     * are applied immediately, or during the next maintenance window. Default is`false`.
     */
    applyImmediately?: pulumi.Input<boolean>;
    /**
     * Indicates that minor engine upgrades will be applied automatically to the instance during the maintenance window. Default is `true`.
     */
    autoMinorVersionUpgrade?: pulumi.Input<boolean>;
    /**
     * The EC2 Availability Zone that the neptune instance is created in.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The identifier of the `aws.neptune.Cluster` in which to launch this instance.
     */
    clusterIdentifier: pulumi.Input<string>;
    /**
     * The name of the database engine to be used for the neptune instance. Defaults to `neptune`. Valid Values: `neptune`.
     */
    engine?: pulumi.Input<string>;
    /**
     * The neptune engine version.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * The identifier for the neptune instance, if omitted, this provider will assign a random, unique identifier.
     */
    identifier?: pulumi.Input<string>;
    /**
     * Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
     */
    identifierPrefix?: pulumi.Input<string>;
    /**
     * The instance class to use.
     */
    instanceClass: pulumi.Input<string>;
    /**
     * The name of the neptune parameter group to associate with this instance.
     */
    neptuneParameterGroupName?: pulumi.Input<string>;
    /**
     * A subnet group to associate with this neptune instance. **NOTE:** This must match the `neptuneSubnetGroupName` of the attached `aws.neptune.Cluster`.
     */
    neptuneSubnetGroupName?: pulumi.Input<string>;
    /**
     * The port on which the DB accepts connections. Defaults to `8182`.
     */
    port?: pulumi.Input<number>;
    /**
     * The daily time range during which automated backups are created if automated backups are enabled. Eg: "04:00-09:00"
     */
    preferredBackupWindow?: pulumi.Input<string>;
    /**
     * The window to perform maintenance in.
     * Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
     */
    preferredMaintenanceWindow?: pulumi.Input<string>;
    /**
     * Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
     */
    promotionTier?: pulumi.Input<number>;
    /**
     * Bool to control if instance is publicly accessible. Default is `false`.
     */
    publiclyAccessible?: pulumi.Input<boolean>;
    /**
     * A map of tags to assign to the instance. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
