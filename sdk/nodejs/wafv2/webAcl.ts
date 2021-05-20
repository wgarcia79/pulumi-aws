// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Creates a WAFv2 Web ACL resource.
 *
 * ## Example Usage
 *
 * This resource is based on `aws.wafv2.RuleGroup`, check the documentation of the `aws.wafv2.RuleGroup` resource to see examples of the various available statements.
 * ### Managed Rule
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.wafv2.WebAcl("example", {
 *     defaultAction: {
 *         allow: {},
 *     },
 *     description: "Example of a managed rule.",
 *     rules: [{
 *         name: "rule-1",
 *         overrideAction: {
 *             count: {},
 *         },
 *         priority: 1,
 *         statement: {
 *             managedRuleGroupStatement: {
 *                 excludedRules: [
 *                     {
 *                         name: "SizeRestrictions_QUERYSTRING",
 *                     },
 *                     {
 *                         name: "NoUserAgent_HEADER",
 *                     },
 *                 ],
 *                 name: "AWSManagedRulesCommonRuleSet",
 *                 vendorName: "AWS",
 *             },
 *         },
 *         visibilityConfig: {
 *             cloudwatchMetricsEnabled: false,
 *             metricName: "friendly-rule-metric-name",
 *             sampledRequestsEnabled: false,
 *         },
 *     }],
 *     scope: "REGIONAL",
 *     tags: {
 *         Tag1: "Value1",
 *         Tag2: "Value2",
 *     },
 *     visibilityConfig: {
 *         cloudwatchMetricsEnabled: false,
 *         metricName: "friendly-metric-name",
 *         sampledRequestsEnabled: false,
 *     },
 * });
 * ```
 * ### Rate Based
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.wafv2.WebAcl("example", {
 *     defaultAction: {
 *         block: {},
 *     },
 *     description: "Example of a rate based statement.",
 *     rules: [{
 *         action: {
 *             count: {},
 *         },
 *         name: "rule-1",
 *         priority: 1,
 *         statement: {
 *             rateBasedStatement: {
 *                 aggregateKeyType: "IP",
 *                 limit: 10000,
 *                 scopeDownStatement: {
 *                     geoMatchStatement: {
 *                         countryCodes: [
 *                             "US",
 *                             "NL",
 *                         ],
 *                     },
 *                 },
 *             },
 *         },
 *         visibilityConfig: {
 *             cloudwatchMetricsEnabled: false,
 *             metricName: "friendly-rule-metric-name",
 *             sampledRequestsEnabled: false,
 *         },
 *     }],
 *     scope: "REGIONAL",
 *     tags: {
 *         Tag1: "Value1",
 *         Tag2: "Value2",
 *     },
 *     visibilityConfig: {
 *         cloudwatchMetricsEnabled: false,
 *         metricName: "friendly-metric-name",
 *         sampledRequestsEnabled: false,
 *     },
 * });
 * ```
 * ### Rule Group Reference
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.wafv2.RuleGroup("example", {
 *     capacity: 10,
 *     scope: "REGIONAL",
 *     rules: [
 *         {
 *             name: "rule-1",
 *             priority: 1,
 *             action: {
 *                 count: {},
 *             },
 *             statement: {
 *                 geoMatchStatement: {
 *                     countryCodes: ["NL"],
 *                 },
 *             },
 *             visibilityConfig: {
 *                 cloudwatchMetricsEnabled: false,
 *                 metricName: "friendly-rule-metric-name",
 *                 sampledRequestsEnabled: false,
 *             },
 *         },
 *         {
 *             name: "rule-to-exclude-a",
 *             priority: 10,
 *             action: {
 *                 allow: {},
 *             },
 *             statement: {
 *                 geoMatchStatement: {
 *                     countryCodes: ["US"],
 *                 },
 *             },
 *             visibilityConfig: {
 *                 cloudwatchMetricsEnabled: false,
 *                 metricName: "friendly-rule-metric-name",
 *                 sampledRequestsEnabled: false,
 *             },
 *         },
 *         {
 *             name: "rule-to-exclude-b",
 *             priority: 15,
 *             action: {
 *                 allow: {},
 *             },
 *             statement: {
 *                 geoMatchStatement: {
 *                     countryCodes: ["GB"],
 *                 },
 *             },
 *             visibilityConfig: {
 *                 cloudwatchMetricsEnabled: false,
 *                 metricName: "friendly-rule-metric-name",
 *                 sampledRequestsEnabled: false,
 *             },
 *         },
 *     ],
 *     visibilityConfig: {
 *         cloudwatchMetricsEnabled: false,
 *         metricName: "friendly-metric-name",
 *         sampledRequestsEnabled: false,
 *     },
 * });
 * const test = new aws.wafv2.WebAcl("test", {
 *     scope: "REGIONAL",
 *     defaultAction: {
 *         block: {},
 *     },
 *     rules: [{
 *         name: "rule-1",
 *         priority: 1,
 *         overrideAction: {
 *             count: {},
 *         },
 *         statement: {
 *             ruleGroupReferenceStatement: {
 *                 arn: example.arn,
 *                 excludedRules: [
 *                     {
 *                         name: "rule-to-exclude-b",
 *                     },
 *                     {
 *                         name: "rule-to-exclude-a",
 *                     },
 *                 ],
 *             },
 *         },
 *         visibilityConfig: {
 *             cloudwatchMetricsEnabled: false,
 *             metricName: "friendly-rule-metric-name",
 *             sampledRequestsEnabled: false,
 *         },
 *     }],
 *     tags: {
 *         Tag1: "Value1",
 *         Tag2: "Value2",
 *     },
 *     visibilityConfig: {
 *         cloudwatchMetricsEnabled: false,
 *         metricName: "friendly-metric-name",
 *         sampledRequestsEnabled: false,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * WAFv2 Web ACLs can be imported using `ID/Name/Scope` e.g.
 *
 * ```sh
 *  $ pulumi import aws:wafv2/webAcl:WebAcl example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc/example/REGIONAL
 * ```
 */
export class WebAcl extends pulumi.CustomResource {
    /**
     * Get an existing WebAcl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebAclState, opts?: pulumi.CustomResourceOptions): WebAcl {
        return new WebAcl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:wafv2/webAcl:WebAcl';

    /**
     * Returns true if the given object is an instance of WebAcl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebAcl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebAcl.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the IP Set that this statement references.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The web ACL capacity units (WCUs) currently being used by this web ACL.
     */
    public /*out*/ readonly capacity!: pulumi.Output<number>;
    /**
     * The action to perform if none of the `rules` contained in the WebACL match. See Default Action below for details.
     */
    public readonly defaultAction!: pulumi.Output<outputs.wafv2.WebAclDefaultAction>;
    /**
     * A friendly description of the WebACL.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly lockToken!: pulumi.Output<string>;
    /**
     * A friendly name of the WebACL.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
     */
    public readonly rules!: pulumi.Output<outputs.wafv2.WebAclRule[] | undefined>;
    /**
     * Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
     */
    public readonly scope!: pulumi.Output<string>;
    /**
     * An map of key:value pairs to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
     */
    public readonly visibilityConfig!: pulumi.Output<outputs.wafv2.WebAclVisibilityConfig>;

    /**
     * Create a WebAcl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebAclArgs | WebAclState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebAclState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["capacity"] = state ? state.capacity : undefined;
            inputs["defaultAction"] = state ? state.defaultAction : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["lockToken"] = state ? state.lockToken : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["rules"] = state ? state.rules : undefined;
            inputs["scope"] = state ? state.scope : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["visibilityConfig"] = state ? state.visibilityConfig : undefined;
        } else {
            const args = argsOrState as WebAclArgs | undefined;
            if ((!args || args.defaultAction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultAction'");
            }
            if ((!args || args.scope === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scope'");
            }
            if ((!args || args.visibilityConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'visibilityConfig'");
            }
            inputs["defaultAction"] = args ? args.defaultAction : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["rules"] = args ? args.rules : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["visibilityConfig"] = args ? args.visibilityConfig : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["capacity"] = undefined /*out*/;
            inputs["lockToken"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WebAcl.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebAcl resources.
 */
export interface WebAclState {
    /**
     * The Amazon Resource Name (ARN) of the IP Set that this statement references.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The web ACL capacity units (WCUs) currently being used by this web ACL.
     */
    readonly capacity?: pulumi.Input<number>;
    /**
     * The action to perform if none of the `rules` contained in the WebACL match. See Default Action below for details.
     */
    readonly defaultAction?: pulumi.Input<inputs.wafv2.WebAclDefaultAction>;
    /**
     * A friendly description of the WebACL.
     */
    readonly description?: pulumi.Input<string>;
    readonly lockToken?: pulumi.Input<string>;
    /**
     * A friendly name of the WebACL.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
     */
    readonly rules?: pulumi.Input<pulumi.Input<inputs.wafv2.WebAclRule>[]>;
    /**
     * Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
     */
    readonly scope?: pulumi.Input<string>;
    /**
     * An map of key:value pairs to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    readonly tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
     */
    readonly visibilityConfig?: pulumi.Input<inputs.wafv2.WebAclVisibilityConfig>;
}

/**
 * The set of arguments for constructing a WebAcl resource.
 */
export interface WebAclArgs {
    /**
     * The action to perform if none of the `rules` contained in the WebACL match. See Default Action below for details.
     */
    readonly defaultAction: pulumi.Input<inputs.wafv2.WebAclDefaultAction>;
    /**
     * A friendly description of the WebACL.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A friendly name of the WebACL.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
     */
    readonly rules?: pulumi.Input<pulumi.Input<inputs.wafv2.WebAclRule>[]>;
    /**
     * Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
     */
    readonly scope: pulumi.Input<string>;
    /**
     * An map of key:value pairs to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    readonly tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
     */
    readonly visibilityConfig: pulumi.Input<inputs.wafv2.WebAclVisibilityConfig>;
}
