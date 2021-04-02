// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CloudWatch Logs query definition resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.cloudwatch.QueryDefinition("example", {
 *     logGroups: [
 *         "/aws/logGroup1",
 *         "/aws/logGroup2",
 *     ],
 *     query: `fields @timestamp, @message
 * | sort @timestamp desc
 * | limit 25
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * CloudWatch query definitions can be imported using the query definition ARN. The ARN can be found on the "Edit Query" page for the query in the AWS Console.
 *
 * ```sh
 *  $ pulumi import aws:cloudwatch/queryDefinition:QueryDefinition example arn:aws:logs:us-west-2:123456789012:query-definition:269951d7-6f75-496d-9d7b-6b7a5486bdbd
 * ```
 */
export class QueryDefinition extends pulumi.CustomResource {
    /**
     * Get an existing QueryDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QueryDefinitionState, opts?: pulumi.CustomResourceOptions): QueryDefinition {
        return new QueryDefinition(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudwatch/queryDefinition:QueryDefinition';

    /**
     * Returns true if the given object is an instance of QueryDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is QueryDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === QueryDefinition.__pulumiType;
    }

    /**
     * Specific log groups to use with the query.
     */
    public readonly logGroupNames!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the query.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The query definition ID.
     */
    public /*out*/ readonly queryDefinitionId!: pulumi.Output<string>;
    /**
     * The query to save. You can read more about CloudWatch Logs Query Syntax in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
     */
    public readonly queryString!: pulumi.Output<string>;

    /**
     * Create a QueryDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QueryDefinitionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QueryDefinitionArgs | QueryDefinitionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as QueryDefinitionState | undefined;
            inputs["logGroupNames"] = state ? state.logGroupNames : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["queryDefinitionId"] = state ? state.queryDefinitionId : undefined;
            inputs["queryString"] = state ? state.queryString : undefined;
        } else {
            const args = argsOrState as QueryDefinitionArgs | undefined;
            if ((!args || args.queryString === undefined) && !opts.urn) {
                throw new Error("Missing required property 'queryString'");
            }
            inputs["logGroupNames"] = args ? args.logGroupNames : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["queryString"] = args ? args.queryString : undefined;
            inputs["queryDefinitionId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(QueryDefinition.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering QueryDefinition resources.
 */
export interface QueryDefinitionState {
    /**
     * Specific log groups to use with the query.
     */
    readonly logGroupNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the query.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The query definition ID.
     */
    readonly queryDefinitionId?: pulumi.Input<string>;
    /**
     * The query to save. You can read more about CloudWatch Logs Query Syntax in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
     */
    readonly queryString?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a QueryDefinition resource.
 */
export interface QueryDefinitionArgs {
    /**
     * Specific log groups to use with the query.
     */
    readonly logGroupNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the query.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The query to save. You can read more about CloudWatch Logs Query Syntax in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
     */
    readonly queryString: pulumi.Input<string>;
}
