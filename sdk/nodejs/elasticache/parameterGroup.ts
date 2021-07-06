// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an ElastiCache parameter group resource.
 *
 * > **NOTE:** Attempting to remove the `reserved-memory` parameter when `family` is set to `redis2.6` or `redis2.8` may show a perpetual difference in this provider due to an Elasticache API limitation. Leave that parameter configured with any value to workaround the issue.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const defaultParameterGroup = new aws.elasticache.ParameterGroup("default", {
 *     family: "redis2.8",
 *     parameters: [
 *         {
 *             name: "activerehashing",
 *             value: "yes",
 *         },
 *         {
 *             name: "min-slaves-to-write",
 *             value: "2",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * ElastiCache Parameter Groups can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:elasticache/parameterGroup:ParameterGroup default redis-params
 * ```
 */
export class ParameterGroup extends pulumi.CustomResource {
    /**
     * Get an existing ParameterGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ParameterGroupState, opts?: pulumi.CustomResourceOptions): ParameterGroup {
        return new ParameterGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:elasticache/parameterGroup:ParameterGroup';

    /**
     * Returns true if the given object is an instance of ParameterGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ParameterGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ParameterGroup.__pulumiType;
    }

    /**
     * The AWS ARN associated with the parameter group.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The description of the ElastiCache parameter group. Defaults to "Managed by Pulumi".
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The family of the ElastiCache parameter group.
     */
    public readonly family!: pulumi.Output<string>;
    /**
     * The name of the ElastiCache parameter.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of ElastiCache parameters to apply.
     */
    public readonly parameters!: pulumi.Output<outputs.elasticache.ParameterGroupParameter[] | undefined>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a ParameterGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ParameterGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ParameterGroupArgs | ParameterGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ParameterGroupState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["family"] = state ? state.family : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as ParameterGroupArgs | undefined;
            if ((!args || args.family === undefined) && !opts.urn) {
                throw new Error("Missing required property 'family'");
            }
            inputs["description"] = (args ? args.description : undefined) ?? "Managed by Pulumi";
            inputs["family"] = args ? args.family : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ParameterGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ParameterGroup resources.
 */
export interface ParameterGroupState {
    /**
     * The AWS ARN associated with the parameter group.
     */
    arn?: pulumi.Input<string>;
    /**
     * The description of the ElastiCache parameter group. Defaults to "Managed by Pulumi".
     */
    description?: pulumi.Input<string>;
    /**
     * The family of the ElastiCache parameter group.
     */
    family?: pulumi.Input<string>;
    /**
     * The name of the ElastiCache parameter.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of ElastiCache parameters to apply.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.elasticache.ParameterGroupParameter>[]>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a ParameterGroup resource.
 */
export interface ParameterGroupArgs {
    /**
     * The description of the ElastiCache parameter group. Defaults to "Managed by Pulumi".
     */
    description?: pulumi.Input<string>;
    /**
     * The family of the ElastiCache parameter group.
     */
    family: pulumi.Input<string>;
    /**
     * The name of the ElastiCache parameter.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of ElastiCache parameters to apply.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.elasticache.ParameterGroupParameter>[]>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
