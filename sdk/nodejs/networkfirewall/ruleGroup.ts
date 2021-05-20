// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an AWS Network Firewall Rule Group Resource
 *
 * ## Example Usage
 * ### Stateful Inspection for denying access to a domain
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.networkfirewall.RuleGroup("example", {
 *     capacity: 100,
 *     ruleGroup: {
 *         rulesSource: {
 *             rulesSourceList: {
 *                 generatedRulesType: "DENYLIST",
 *                 targetTypes: ["HTTP_HOST"],
 *                 targets: ["test.example.com"],
 *             },
 *         },
 *     },
 *     tags: {
 *         Tag1: "Value1",
 *         Tag2: "Value2",
 *     },
 *     type: "STATEFUL",
 * });
 * ```
 * ### Stateful Inspection for permitting packets from a source IP address
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const ips = [
 *     "1.1.1.1/32",
 *     "1.0.0.1/32",
 * ];
 * const example = new aws.networkfirewall.RuleGroup("example", {
 *     capacity: 50,
 *     description: "Permits http traffic from source",
 *     type: "STATEFUL",
 *     ruleGroup: {
 *         rulesSource: {
 *             dynamic: [{
 *                 forEach: ips,
 *                 content: [{
 *                     action: "PASS",
 *                     header: [{
 *                         destination: "ANY",
 *                         destinationPort: "ANY",
 *                         protocol: "HTTP",
 *                         direction: "ANY",
 *                         sourcePort: "ANY",
 *                         source: stateful_rule.value,
 *                     }],
 *                     ruleOption: [{
 *                         keyword: "sid:1",
 *                     }],
 *                 }],
 *             }],
 *         },
 *     },
 *     tags: {
 *         Name: "permit HTTP from source",
 *     },
 * });
 * ```
 * ### Stateful Inspection for blocking packets from going to an intended destination
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.networkfirewall.RuleGroup("example", {
 *     capacity: 100,
 *     ruleGroup: {
 *         rulesSource: {
 *             statefulRules: [{
 *                 action: "DROP",
 *                 header: {
 *                     destination: "124.1.1.24/32",
 *                     destinationPort: "53",
 *                     direction: "ANY",
 *                     protocol: "TCP",
 *                     source: "1.2.3.4/32",
 *                     sourcePort: "53",
 *                 },
 *                 ruleOptions: [{
 *                     keyword: "sid:1",
 *                 }],
 *             }],
 *         },
 *     },
 *     tags: {
 *         Tag1: "Value1",
 *         Tag2: "Value2",
 *     },
 *     type: "STATEFUL",
 * });
 * ```
 * ### Stateful Inspection from rules specifications defined in Suricata flat format
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * from "fs";
 *
 * const example = new aws.networkfirewall.RuleGroup("example", {
 *     capacity: 100,
 *     type: "STATEFUL",
 *     rules: fs.readFileSync("example.rules"),
 *     tags: {
 *         Tag1: "Value1",
 *         Tag2: "Value2",
 *     },
 * });
 * ```
 * ### Stateless Inspection with a Custom Action
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.networkfirewall.RuleGroup("example", {
 *     capacity: 100,
 *     description: "Stateless Rate Limiting Rule",
 *     ruleGroup: {
 *         rulesSource: {
 *             statelessRulesAndCustomActions: {
 *                 customActions: [{
 *                     actionDefinition: {
 *                         publishMetricAction: {
 *                             dimensions: [{
 *                                 value: "2",
 *                             }],
 *                         },
 *                     },
 *                     actionName: "ExampleMetricsAction",
 *                 }],
 *                 statelessRules: [{
 *                     priority: 1,
 *                     ruleDefinition: {
 *                         actions: [
 *                             "aws:pass",
 *                             "ExampleMetricsAction",
 *                         ],
 *                         matchAttributes: {
 *                             destinations: [{
 *                                 addressDefinition: "124.1.1.5/32",
 *                             }],
 *                             destinationPorts: [{
 *                                 fromPort: 443,
 *                                 toPort: 443,
 *                             }],
 *                             protocols: [6],
 *                             sources: [{
 *                                 addressDefinition: "1.2.3.4/32",
 *                             }],
 *                             sourcePorts: [{
 *                                 fromPort: 443,
 *                                 toPort: 443,
 *                             }],
 *                             tcpFlags: [{
 *                                 flags: ["SYN"],
 *                                 masks: [
 *                                     "SYN",
 *                                     "ACK",
 *                                 ],
 *                             }],
 *                         },
 *                     },
 *                 }],
 *             },
 *         },
 *     },
 *     tags: {
 *         Tag1: "Value1",
 *         Tag2: "Value2",
 *     },
 *     type: "STATELESS",
 * });
 * ```
 *
 * ## Import
 *
 * Network Firewall Rule Groups can be imported using their `ARN`.
 *
 * ```sh
 *  $ pulumi import aws:networkfirewall/ruleGroup:RuleGroup example arn:aws:network-firewall:us-west-1:123456789012:stateful-rulegroup/example
 * ```
 */
export class RuleGroup extends pulumi.CustomResource {
    /**
     * Get an existing RuleGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleGroupState, opts?: pulumi.CustomResourceOptions): RuleGroup {
        return new RuleGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:networkfirewall/ruleGroup:RuleGroup';

    /**
     * Returns true if the given object is an instance of RuleGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuleGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuleGroup.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) that identifies the rule group.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
     */
    public readonly capacity!: pulumi.Output<number>;
    /**
     * A friendly description of the rule group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A friendly name of the rule group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A configuration block that defines the rule group rules. Required unless `rules` is specified. See Rule Group below for details.
     */
    public readonly ruleGroup!: pulumi.Output<outputs.networkfirewall.RuleGroupRuleGroup>;
    /**
     * The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `ruleGroup` is specified.
     */
    public readonly rules!: pulumi.Output<string | undefined>;
    /**
     * A map of key:value pairs to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * A string token used when updating the rule group.
     */
    public /*out*/ readonly updateToken!: pulumi.Output<string>;

    /**
     * Create a RuleGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleGroupArgs | RuleGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleGroupState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["capacity"] = state ? state.capacity : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["ruleGroup"] = state ? state.ruleGroup : undefined;
            inputs["rules"] = state ? state.rules : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["updateToken"] = state ? state.updateToken : undefined;
        } else {
            const args = argsOrState as RuleGroupArgs | undefined;
            if ((!args || args.capacity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'capacity'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["capacity"] = args ? args.capacity : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["ruleGroup"] = args ? args.ruleGroup : undefined;
            inputs["rules"] = args ? args.rules : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["updateToken"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RuleGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RuleGroup resources.
 */
export interface RuleGroupState {
    /**
     * The Amazon Resource Name (ARN) that identifies the rule group.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
     */
    readonly capacity?: pulumi.Input<number>;
    /**
     * A friendly description of the rule group.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A friendly name of the rule group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A configuration block that defines the rule group rules. Required unless `rules` is specified. See Rule Group below for details.
     */
    readonly ruleGroup?: pulumi.Input<inputs.networkfirewall.RuleGroupRuleGroup>;
    /**
     * The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `ruleGroup` is specified.
     */
    readonly rules?: pulumi.Input<string>;
    /**
     * A map of key:value pairs to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    readonly tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * A string token used when updating the rule group.
     */
    readonly updateToken?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RuleGroup resource.
 */
export interface RuleGroupArgs {
    /**
     * The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
     */
    readonly capacity: pulumi.Input<number>;
    /**
     * A friendly description of the rule group.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A friendly name of the rule group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A configuration block that defines the rule group rules. Required unless `rules` is specified. See Rule Group below for details.
     */
    readonly ruleGroup?: pulumi.Input<inputs.networkfirewall.RuleGroupRuleGroup>;
    /**
     * The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `ruleGroup` is specified.
     */
    readonly rules?: pulumi.Input<string>;
    /**
     * A map of key:value pairs to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    readonly tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.
     */
    readonly type: pulumi.Input<string>;
}
