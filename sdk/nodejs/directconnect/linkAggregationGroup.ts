// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Direct Connect LAG. Connections can be added to the LAG via the `aws.directconnect.Connection` and `aws.directconnect.ConnectionAssociation` resources.
 *
 * > *NOTE:* When creating a LAG, Direct Connect requires creating a Connection. This provider will remove this unmanaged connection during resource creation.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const hoge = new aws.directconnect.LinkAggregationGroup("hoge", {
 *     connectionsBandwidth: "1Gbps",
 *     forceDestroy: true,
 *     location: "EqDC2",
 * });
 * ```
 *
 * ## Import
 *
 * Direct Connect LAGs can be imported using the `lag id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:directconnect/linkAggregationGroup:LinkAggregationGroup test_lag dxlag-fgnsp5rq
 * ```
 */
export class LinkAggregationGroup extends pulumi.CustomResource {
    /**
     * Get an existing LinkAggregationGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LinkAggregationGroupState, opts?: pulumi.CustomResourceOptions): LinkAggregationGroup {
        return new LinkAggregationGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:directconnect/linkAggregationGroup:LinkAggregationGroup';

    /**
     * Returns true if the given object is an instance of LinkAggregationGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LinkAggregationGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LinkAggregationGroup.__pulumiType;
    }

    /**
     * The ARN of the LAG.
     * * `jumboFrameCapable` -Indicates whether jumbo frames (9001 MTU) are supported.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
     */
    public readonly connectionsBandwidth!: pulumi.Output<string>;
    /**
     * A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
     */
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether the LAG supports a secondary BGP peer in the same address family (IPv4/IPv6).
     */
    public /*out*/ readonly hasLogicalRedundancy!: pulumi.Output<string>;
    public /*out*/ readonly jumboFrameCapable!: pulumi.Output<boolean>;
    /**
     * The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the LAG.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a LinkAggregationGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkAggregationGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LinkAggregationGroupArgs | LinkAggregationGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LinkAggregationGroupState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["connectionsBandwidth"] = state ? state.connectionsBandwidth : undefined;
            inputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            inputs["hasLogicalRedundancy"] = state ? state.hasLogicalRedundancy : undefined;
            inputs["jumboFrameCapable"] = state ? state.jumboFrameCapable : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as LinkAggregationGroupArgs | undefined;
            if ((!args || args.connectionsBandwidth === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionsBandwidth'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            inputs["connectionsBandwidth"] = args ? args.connectionsBandwidth : undefined;
            inputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["hasLogicalRedundancy"] = undefined /*out*/;
            inputs["jumboFrameCapable"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LinkAggregationGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LinkAggregationGroup resources.
 */
export interface LinkAggregationGroupState {
    /**
     * The ARN of the LAG.
     * * `jumboFrameCapable` -Indicates whether jumbo frames (9001 MTU) are supported.
     */
    arn?: pulumi.Input<string>;
    /**
     * The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
     */
    connectionsBandwidth?: pulumi.Input<string>;
    /**
     * A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * Indicates whether the LAG supports a secondary BGP peer in the same address family (IPv4/IPv6).
     */
    hasLogicalRedundancy?: pulumi.Input<string>;
    jumboFrameCapable?: pulumi.Input<boolean>;
    /**
     * The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the LAG.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a LinkAggregationGroup resource.
 */
export interface LinkAggregationGroupArgs {
    /**
     * The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
     */
    connectionsBandwidth: pulumi.Input<string>;
    /**
     * A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
     */
    location: pulumi.Input<string>;
    /**
     * The name of the LAG.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
