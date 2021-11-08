// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Access Analyzer Analyzer. More information can be found in the [Access Analyzer User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/what-is-access-analyzer.html).
 *
 * ## Example Usage
 * ### Account Analyzer
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.accessanalyzer.Analyzer("example", {
 *     analyzerName: "example",
 * });
 * ```
 * ### Organization Analyzer
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleOrganization = new aws.organizations.Organization("exampleOrganization", {awsServiceAccessPrincipals: ["access-analyzer.amazonaws.com"]});
 * const exampleAnalyzer = new aws.accessanalyzer.Analyzer("exampleAnalyzer", {
 *     analyzerName: "example",
 *     type: "ORGANIZATION",
 * }, {
 *     dependsOn: [exampleOrganization],
 * });
 * ```
 *
 * ## Import
 *
 * Access Analyzer Analyzers can be imported using the `analyzer_name`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:accessanalyzer/analyzer:Analyzer example example
 * ```
 */
export class Analyzer extends pulumi.CustomResource {
    /**
     * Get an existing Analyzer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AnalyzerState, opts?: pulumi.CustomResourceOptions): Analyzer {
        return new Analyzer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:accessanalyzer/analyzer:Analyzer';

    /**
     * Returns true if the given object is an instance of Analyzer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Analyzer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Analyzer.__pulumiType;
    }

    /**
     * Name of the Analyzer.
     */
    public readonly analyzerName!: pulumi.Output<string>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Type of Analyzer. Valid values are `ACCOUNT` or `ORGANIZATION`. Defaults to `ACCOUNT`.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a Analyzer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AnalyzerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AnalyzerArgs | AnalyzerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AnalyzerState | undefined;
            inputs["analyzerName"] = state ? state.analyzerName : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as AnalyzerArgs | undefined;
            if ((!args || args.analyzerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'analyzerName'");
            }
            inputs["analyzerName"] = args ? args.analyzerName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["tagsAll"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Analyzer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Analyzer resources.
 */
export interface AnalyzerState {
    /**
     * Name of the Analyzer.
     */
    analyzerName?: pulumi.Input<string>;
    arn?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Type of Analyzer. Valid values are `ACCOUNT` or `ORGANIZATION`. Defaults to `ACCOUNT`.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Analyzer resource.
 */
export interface AnalyzerArgs {
    /**
     * Name of the Analyzer.
     */
    analyzerName: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Type of Analyzer. Valid values are `ACCOUNT` or `ORGANIZATION`. Defaults to `ACCOUNT`.
     */
    type?: pulumi.Input<string>;
}
