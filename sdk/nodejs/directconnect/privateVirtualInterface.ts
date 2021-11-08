// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Direct Connect private virtual interface resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const foo = new aws.directconnect.PrivateVirtualInterface("foo", {
 *     addressFamily: "ipv4",
 *     bgpAsn: 65352,
 *     connectionId: "dxcon-zzzzzzzz",
 *     vlan: 4094,
 * });
 * ```
 *
 * ## Import
 *
 * Direct Connect private virtual interfaces can be imported using the `vif id`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:directconnect/privateVirtualInterface:PrivateVirtualInterface test dxvif-33cc44dd
 * ```
 */
export class PrivateVirtualInterface extends pulumi.CustomResource {
    /**
     * Get an existing PrivateVirtualInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivateVirtualInterfaceState, opts?: pulumi.CustomResourceOptions): PrivateVirtualInterface {
        return new PrivateVirtualInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:directconnect/privateVirtualInterface:PrivateVirtualInterface';

    /**
     * Returns true if the given object is an instance of PrivateVirtualInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateVirtualInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateVirtualInterface.__pulumiType;
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
     * The ID of the Direct Connect gateway to which to connect the virtual interface.
     */
    public readonly dxGatewayId!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether jumbo frames (9001 MTU) are supported.
     */
    public /*out*/ readonly jumboFrameCapable!: pulumi.Output<boolean>;
    /**
     * The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
     * The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
     */
    public readonly mtu!: pulumi.Output<number | undefined>;
    /**
     * The name for the virtual interface.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The VLAN ID.
     */
    public readonly vlan!: pulumi.Output<number>;
    /**
     * The ID of the virtual private gateway to which to connect the virtual interface.
     */
    public readonly vpnGatewayId!: pulumi.Output<string | undefined>;

    /**
     * Create a PrivateVirtualInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateVirtualInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivateVirtualInterfaceArgs | PrivateVirtualInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrivateVirtualInterfaceState | undefined;
            inputs["addressFamily"] = state ? state.addressFamily : undefined;
            inputs["amazonAddress"] = state ? state.amazonAddress : undefined;
            inputs["amazonSideAsn"] = state ? state.amazonSideAsn : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["awsDevice"] = state ? state.awsDevice : undefined;
            inputs["bgpAsn"] = state ? state.bgpAsn : undefined;
            inputs["bgpAuthKey"] = state ? state.bgpAuthKey : undefined;
            inputs["connectionId"] = state ? state.connectionId : undefined;
            inputs["customerAddress"] = state ? state.customerAddress : undefined;
            inputs["dxGatewayId"] = state ? state.dxGatewayId : undefined;
            inputs["jumboFrameCapable"] = state ? state.jumboFrameCapable : undefined;
            inputs["mtu"] = state ? state.mtu : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["vlan"] = state ? state.vlan : undefined;
            inputs["vpnGatewayId"] = state ? state.vpnGatewayId : undefined;
        } else {
            const args = argsOrState as PrivateVirtualInterfaceArgs | undefined;
            if ((!args || args.addressFamily === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addressFamily'");
            }
            if ((!args || args.bgpAsn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bgpAsn'");
            }
            if ((!args || args.connectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionId'");
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
            inputs["dxGatewayId"] = args ? args.dxGatewayId : undefined;
            inputs["mtu"] = args ? args.mtu : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vlan"] = args ? args.vlan : undefined;
            inputs["vpnGatewayId"] = args ? args.vpnGatewayId : undefined;
            inputs["amazonSideAsn"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["awsDevice"] = undefined /*out*/;
            inputs["jumboFrameCapable"] = undefined /*out*/;
            inputs["tagsAll"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(PrivateVirtualInterface.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivateVirtualInterface resources.
 */
export interface PrivateVirtualInterfaceState {
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
     * The ID of the Direct Connect gateway to which to connect the virtual interface.
     */
    dxGatewayId?: pulumi.Input<string>;
    /**
     * Indicates whether jumbo frames (9001 MTU) are supported.
     */
    jumboFrameCapable?: pulumi.Input<boolean>;
    /**
     * The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
     * The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
     */
    mtu?: pulumi.Input<number>;
    /**
     * The name for the virtual interface.
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
    /**
     * The VLAN ID.
     */
    vlan?: pulumi.Input<number>;
    /**
     * The ID of the virtual private gateway to which to connect the virtual interface.
     */
    vpnGatewayId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PrivateVirtualInterface resource.
 */
export interface PrivateVirtualInterfaceArgs {
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
     * The ID of the Direct Connect gateway to which to connect the virtual interface.
     */
    dxGatewayId?: pulumi.Input<string>;
    /**
     * The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
     * The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
     */
    mtu?: pulumi.Input<number>;
    /**
     * The name for the virtual interface.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The VLAN ID.
     */
    vlan: pulumi.Input<number>;
    /**
     * The ID of the virtual private gateway to which to connect the virtual interface.
     */
    vpnGatewayId?: pulumi.Input<string>;
}
