// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a security group rule resource. Represents a single `ingress` or
 * `egress` group rule, which can be added to external Security Groups.
 *
 * > **NOTE on Security Groups and Security Group Rules:** This provider currently
 * provides both a standalone Security Group Rule resource (a single `ingress` or
 * `egress` rule), and a Security Group resource with `ingress` and `egress` rules
 * defined in-line. At this time you cannot use a Security Group with in-line rules
 * in conjunction with any Security Group Rule resources. Doing so will cause
 * a conflict of rule settings and will overwrite rules.
 *
 * > **NOTE:** Setting `protocol = "all"` or `protocol = -1` with `fromPort` and `toPort` will result in the EC2 API creating a security group rule with all ports open. This API behavior cannot be controlled by this provider and may generate warnings in the future.
 *
 * > **NOTE:** Referencing Security Groups across VPC peering has certain restrictions. More information is available in the [VPC Peering User Guide](https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-security-groups.html).
 *
 * ## Example Usage
 *
 * Basic usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ec2.SecurityGroupRule("example", {
 *     type: "ingress",
 *     fromPort: 0,
 *     toPort: 65535,
 *     protocol: "tcp",
 *     cidrBlocks: [aws_vpc.example.cidr_block],
 *     securityGroupId: "sg-123456",
 * });
 * ```
 * ## Usage with prefix list IDs
 *
 * Prefix Lists are either managed by AWS internally, or created by the customer using a
 * Managed Prefix List resource. Prefix Lists provided by
 * AWS are associated with a prefix list name, or service name, that is linked to a specific region.
 *
 * Prefix list IDs are exported on VPC Endpoints, so you can use this format:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // ...
 * const myEndpoint = new aws.ec2.VpcEndpoint("myEndpoint", {});
 * // ...
 * const allowAll = new aws.ec2.SecurityGroupRule("allowAll", {
 *     type: "egress",
 *     toPort: 0,
 *     protocol: "-1",
 *     prefixListIds: [myEndpoint.prefixListId],
 *     fromPort: 0,
 *     securityGroupId: "sg-123456",
 * });
 * ```
 *
 * You can also find a specific Prefix List using the `aws.ec2.getPrefixList` data source.
 *
 * ## Import
 *
 * ### Examples Import an ingress rule in security group `sg-6e616f6d69` for TCP port 8000 with an IPv4 destination CIDR of `10.0.3.0/24`console
 *
 * ```sh
 *  $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule ingress sg-6e616f6d69_ingress_tcp_8000_8000_10.0.3.0/24
 * ```
 *
 *  Import a rule with various IPv4 and IPv6 source CIDR blocksconsole
 *
 * ```sh
 *  $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule ingress sg-4973616163_ingress_tcp_100_121_10.1.0.0/16_2001:db8::/48_10.2.0.0/16_2002:db8::/48
 * ```
 *
 *  Import a rule, applicable to all ports, with a protocol other than TCP/UDP/ICMP/ALL, e.g., Multicast Transport Protocol (MTP), using the IANA protocol number, e.g., 92. console
 *
 * ```sh
 *  $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule ingress sg-6777656e646f6c796e_ingress_92_0_65536_10.0.3.0/24_10.0.4.0/24
 * ```
 *
 *  Import an egress rule with a prefix list ID destinationconsole
 *
 * ```sh
 *  $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule egress sg-62726f6479_egress_tcp_8000_8000_pl-6469726b
 * ```
 *
 *  Import a rule applicable to all protocols and ports with a security group sourceconsole
 *
 * ```sh
 *  $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule ingress_rule sg-7472697374616e_ingress_all_0_65536_sg-6176657279
 * ```
 *
 *  Import a rule that has itself and an IPv6 CIDR block as sourcesconsole
 *
 * ```sh
 *  $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule rule_name sg-656c65616e6f72_ingress_tcp_80_80_self_2001:db8::/48
 * ```
 */
export class SecurityGroupRule extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroupRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityGroupRuleState, opts?: pulumi.CustomResourceOptions): SecurityGroupRule {
        return new SecurityGroupRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/securityGroupRule:SecurityGroupRule';

    /**
     * Returns true if the given object is an instance of SecurityGroupRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityGroupRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityGroupRule.__pulumiType;
    }

    /**
     * List of CIDR blocks. Cannot be specified with `sourceSecurityGroupId`.
     */
    public readonly cidrBlocks!: pulumi.Output<string[] | undefined>;
    /**
     * Description of the rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The start port (or ICMP type number if protocol is "icmp" or "icmpv6").
     */
    public readonly fromPort!: pulumi.Output<number>;
    /**
     * List of IPv6 CIDR blocks.
     */
    public readonly ipv6CidrBlocks!: pulumi.Output<string[] | undefined>;
    /**
     * List of Prefix List IDs.
     */
    public readonly prefixListIds!: pulumi.Output<string[] | undefined>;
    /**
     * The protocol. If not icmp, icmpv6, tcp, udp, or all use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The security group to apply this rule to.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * If true, the security group itself will be added as
     * a source to this ingress rule. Cannot be specified with `sourceSecurityGroupId`.
     */
    public readonly self!: pulumi.Output<boolean | undefined>;
    /**
     * The security group id to allow access to/from,
     * depending on the `type`. Cannot be specified with `cidrBlocks` and `self`.
     */
    public readonly sourceSecurityGroupId!: pulumi.Output<string>;
    /**
     * The end port (or ICMP code if protocol is "icmp").
     */
    public readonly toPort!: pulumi.Output<number>;
    /**
     * The type of rule being created. Valid options are `ingress` (inbound)
     * or `egress` (outbound).
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a SecurityGroupRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityGroupRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityGroupRuleArgs | SecurityGroupRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityGroupRuleState | undefined;
            inputs["cidrBlocks"] = state ? state.cidrBlocks : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["fromPort"] = state ? state.fromPort : undefined;
            inputs["ipv6CidrBlocks"] = state ? state.ipv6CidrBlocks : undefined;
            inputs["prefixListIds"] = state ? state.prefixListIds : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            inputs["self"] = state ? state.self : undefined;
            inputs["sourceSecurityGroupId"] = state ? state.sourceSecurityGroupId : undefined;
            inputs["toPort"] = state ? state.toPort : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as SecurityGroupRuleArgs | undefined;
            if ((!args || args.fromPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fromPort'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.securityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            if ((!args || args.toPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'toPort'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["cidrBlocks"] = args ? args.cidrBlocks : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["fromPort"] = args ? args.fromPort : undefined;
            inputs["ipv6CidrBlocks"] = args ? args.ipv6CidrBlocks : undefined;
            inputs["prefixListIds"] = args ? args.prefixListIds : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            inputs["self"] = args ? args.self : undefined;
            inputs["sourceSecurityGroupId"] = args ? args.sourceSecurityGroupId : undefined;
            inputs["toPort"] = args ? args.toPort : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SecurityGroupRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroupRule resources.
 */
export interface SecurityGroupRuleState {
    /**
     * List of CIDR blocks. Cannot be specified with `sourceSecurityGroupId`.
     */
    readonly cidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description of the rule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The start port (or ICMP type number if protocol is "icmp" or "icmpv6").
     */
    readonly fromPort?: pulumi.Input<number>;
    /**
     * List of IPv6 CIDR blocks.
     */
    readonly ipv6CidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of Prefix List IDs.
     */
    readonly prefixListIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The protocol. If not icmp, icmpv6, tcp, udp, or all use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
     */
    readonly protocol?: pulumi.Input<string | enums.ec2.ProtocolType>;
    /**
     * The security group to apply this rule to.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    /**
     * If true, the security group itself will be added as
     * a source to this ingress rule. Cannot be specified with `sourceSecurityGroupId`.
     */
    readonly self?: pulumi.Input<boolean>;
    /**
     * The security group id to allow access to/from,
     * depending on the `type`. Cannot be specified with `cidrBlocks` and `self`.
     */
    readonly sourceSecurityGroupId?: pulumi.Input<string>;
    /**
     * The end port (or ICMP code if protocol is "icmp").
     */
    readonly toPort?: pulumi.Input<number>;
    /**
     * The type of rule being created. Valid options are `ingress` (inbound)
     * or `egress` (outbound).
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityGroupRule resource.
 */
export interface SecurityGroupRuleArgs {
    /**
     * List of CIDR blocks. Cannot be specified with `sourceSecurityGroupId`.
     */
    readonly cidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description of the rule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The start port (or ICMP type number if protocol is "icmp" or "icmpv6").
     */
    readonly fromPort: pulumi.Input<number>;
    /**
     * List of IPv6 CIDR blocks.
     */
    readonly ipv6CidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of Prefix List IDs.
     */
    readonly prefixListIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The protocol. If not icmp, icmpv6, tcp, udp, or all use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
     */
    readonly protocol: pulumi.Input<string | enums.ec2.ProtocolType>;
    /**
     * The security group to apply this rule to.
     */
    readonly securityGroupId: pulumi.Input<string>;
    /**
     * If true, the security group itself will be added as
     * a source to this ingress rule. Cannot be specified with `sourceSecurityGroupId`.
     */
    readonly self?: pulumi.Input<boolean>;
    /**
     * The security group id to allow access to/from,
     * depending on the `type`. Cannot be specified with `cidrBlocks` and `self`.
     */
    readonly sourceSecurityGroupId?: pulumi.Input<string>;
    /**
     * The end port (or ICMP code if protocol is "icmp").
     */
    readonly toPort: pulumi.Input<number>;
    /**
     * The type of rule being created. Valid options are `ingress` (inbound)
     * or `egress` (outbound).
     */
    readonly type: pulumi.Input<string>;
}
