// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Direct Connect hosted transit virtual interface resource.
 * This resource represents the allocator's side of the hosted virtual interface.
 * A hosted virtual interface is a virtual interface that is owned by another AWS account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.directconnect.HostedTransitVirtualInterface("example", {
 *     connectionId: aws_dx_connection.example.id,
 *     vlan: 4094,
 *     addressFamily: "ipv4",
 *     bgpAsn: 65352,
 * });
 * ```
 *
 * ## Import
 *
 * Direct Connect hosted transit virtual interfaces can be imported using the `vif id`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:directconnect/hostedTransitVirtualInterface:HostedTransitVirtualInterface test dxvif-33cc44dd
 * ```
 */
export class HostedTransitVirtualInterface extends pulumi.CustomResource {
    /**
     * Get an existing HostedTransitVirtualInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostedTransitVirtualInterfaceState, opts?: pulumi.CustomResourceOptions): HostedTransitVirtualInterface {
        return new HostedTransitVirtualInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:directconnect/hostedTransitVirtualInterface:HostedTransitVirtualInterface';

    /**
     * Returns true if the given object is an instance of HostedTransitVirtualInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HostedTransitVirtualInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HostedTransitVirtualInterface.__pulumiType;
    }

    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    public readonly addressFamily!: pulumi.Output<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
     */
    public readonly amazonAddress!: pulumi.Output<string>;
    public /*out*/ readonly amazonSideAsn!: pulumi.Output<string>;
    /**
     * The ARN of the virtual interface.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The Direct Connect endpoint on which the virtual interface terminates.
     */
    public /*out*/ readonly awsDevice!: pulumi.Output<string>;
    /**
     * The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
     */
    public readonly bgpAsn!: pulumi.Output<number>;
    /**
     * The authentication key for BGP configuration.
     */
    public readonly bgpAuthKey!: pulumi.Output<string>;
    /**
     * The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
     */
    public readonly connectionId!: pulumi.Output<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
     */
    public readonly customerAddress!: pulumi.Output<string>;
    /**
     * Indicates whether jumbo frames (8500 MTU) are supported.
     */
    public /*out*/ readonly jumboFrameCapable!: pulumi.Output<boolean>;
    /**
     * The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual transit interface can be either `1500` or `8500` (jumbo frames). Default is `1500`.
     */
    public readonly mtu!: pulumi.Output<number | undefined>;
    /**
     * The name for the virtual interface.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The AWS account that will own the new virtual interface.
     */
    public readonly ownerAccountId!: pulumi.Output<string>;
    /**
     * The VLAN ID.
     */
    public readonly vlan!: pulumi.Output<number>;

    /**
     * Create a HostedTransitVirtualInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostedTransitVirtualInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostedTransitVirtualInterfaceArgs | HostedTransitVirtualInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HostedTransitVirtualInterfaceState | undefined;
            inputs["addressFamily"] = state ? state.addressFamily : undefined;
            inputs["amazonAddress"] = state ? state.amazonAddress : undefined;
            inputs["amazonSideAsn"] = state ? state.amazonSideAsn : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["awsDevice"] = state ? state.awsDevice : undefined;
            inputs["bgpAsn"] = state ? state.bgpAsn : undefined;
            inputs["bgpAuthKey"] = state ? state.bgpAuthKey : undefined;
            inputs["connectionId"] = state ? state.connectionId : undefined;
            inputs["customerAddress"] = state ? state.customerAddress : undefined;
            inputs["jumboFrameCapable"] = state ? state.jumboFrameCapable : undefined;
            inputs["mtu"] = state ? state.mtu : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["ownerAccountId"] = state ? state.ownerAccountId : undefined;
            inputs["vlan"] = state ? state.vlan : undefined;
        } else {
            const args = argsOrState as HostedTransitVirtualInterfaceArgs | undefined;
            if ((!args || args.addressFamily === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addressFamily'");
            }
            if ((!args || args.bgpAsn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bgpAsn'");
            }
            if ((!args || args.connectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionId'");
            }
            if ((!args || args.ownerAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ownerAccountId'");
            }
            if ((!args || args.vlan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vlan'");
            }
            inputs["addressFamily"] = args ? args.addressFamily : undefined;
            inputs["amazonAddress"] = args ? args.amazonAddress : undefined;
            inputs["bgpAsn"] = args ? args.bgpAsn : undefined;
            inputs["bgpAuthKey"] = args ? args.bgpAuthKey : undefined;
            inputs["connectionId"] = args ? args.connectionId : undefined;
            inputs["customerAddress"] = args ? args.customerAddress : undefined;
            inputs["mtu"] = args ? args.mtu : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["ownerAccountId"] = args ? args.ownerAccountId : undefined;
            inputs["vlan"] = args ? args.vlan : undefined;
            inputs["amazonSideAsn"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["awsDevice"] = undefined /*out*/;
            inputs["jumboFrameCapable"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(HostedTransitVirtualInterface.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HostedTransitVirtualInterface resources.
 */
export interface HostedTransitVirtualInterfaceState {
    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    addressFamily?: pulumi.Input<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
     */
    amazonAddress?: pulumi.Input<string>;
    amazonSideAsn?: pulumi.Input<string>;
    /**
     * The ARN of the virtual interface.
     */
    arn?: pulumi.Input<string>;
    /**
     * The Direct Connect endpoint on which the virtual interface terminates.
     */
    awsDevice?: pulumi.Input<string>;
    /**
     * The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
     */
    bgpAsn?: pulumi.Input<number>;
    /**
     * The authentication key for BGP configuration.
     */
    bgpAuthKey?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
     */
    connectionId?: pulumi.Input<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
     */
    customerAddress?: pulumi.Input<string>;
    /**
     * Indicates whether jumbo frames (8500 MTU) are supported.
     */
    jumboFrameCapable?: pulumi.Input<boolean>;
    /**
     * The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual transit interface can be either `1500` or `8500` (jumbo frames). Default is `1500`.
     */
    mtu?: pulumi.Input<number>;
    /**
     * The name for the virtual interface.
     */
    name?: pulumi.Input<string>;
    /**
     * The AWS account that will own the new virtual interface.
     */
    ownerAccountId?: pulumi.Input<string>;
    /**
     * The VLAN ID.
     */
    vlan?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a HostedTransitVirtualInterface resource.
 */
export interface HostedTransitVirtualInterfaceArgs {
    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    addressFamily: pulumi.Input<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
     */
    amazonAddress?: pulumi.Input<string>;
    /**
     * The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
     */
    bgpAsn: pulumi.Input<number>;
    /**
     * The authentication key for BGP configuration.
     */
    bgpAuthKey?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
     */
    connectionId: pulumi.Input<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
     */
    customerAddress?: pulumi.Input<string>;
    /**
     * The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual transit interface can be either `1500` or `8500` (jumbo frames). Default is `1500`.
     */
    mtu?: pulumi.Input<number>;
    /**
     * The name for the virtual interface.
     */
    name?: pulumi.Input<string>;
    /**
     * The AWS account that will own the new virtual interface.
     */
    ownerAccountId: pulumi.Input<string>;
    /**
     * The VLAN ID.
     */
    vlan: pulumi.Input<number>;
}
