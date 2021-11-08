// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CloudFront Function resource. With CloudFront Functions in Amazon CloudFront, you can write lightweight functions in JavaScript for high-scale, latency-sensitive CDN customizations.
 *
 * See [CloudFront Functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-functions.html)
 *
 * > **NOTE:** You cannot delete a function if it’s associated with a cache behavior. First, update your distributions to remove the function association from all cache behaviors, then delete the function.
 *
 * ## Example Usage
 * ### Basic Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * from "fs";
 *
 * const test = new aws.cloudfront.Function("test", {
 *     runtime: "cloudfront-js-1.0",
 *     comment: "my function",
 *     publish: true,
 *     code: fs.readFileSync(`${path.module}/function.js`),
 * });
 * ```
 *
 * ## Import
 *
 * CloudFront Functions can be imported using the `name`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:cloudfront/function:Function test my_test_function
 * ```
 */
export class Function extends pulumi.CustomResource {
    /**
     * Get an existing Function resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionState, opts?: pulumi.CustomResourceOptions): Function {
        return new Function(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudfront/function:Function';

    /**
     * Returns true if the given object is an instance of Function.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Function {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Function.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) identifying your CloudFront Function.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Source code of the function
     */
    public readonly code!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * ETag hash of the function
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Unique name for your CloudFront Function.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
     */
    public readonly publish!: pulumi.Output<boolean | undefined>;
    /**
     * Identifier of the function's runtime. Currently only `cloudfront-js-1.0` is valid.
     */
    public readonly runtime!: pulumi.Output<string>;
    /**
     * Status of the function. Can be `UNPUBLISHED`, `UNASSOCIATED` or `ASSOCIATED`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Function resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionArgs | FunctionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["code"] = state ? state.code : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["publish"] = state ? state.publish : undefined;
            inputs["runtime"] = state ? state.runtime : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as FunctionArgs | undefined;
            if ((!args || args.code === undefined) && !opts.urn) {
                throw new Error("Missing required property 'code'");
            }
            if ((!args || args.runtime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runtime'");
            }
            inputs["code"] = args ? args.code : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["publish"] = args ? args.publish : undefined;
            inputs["runtime"] = args ? args.runtime : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Function.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Function resources.
 */
export interface FunctionState {
    /**
     * Amazon Resource Name (ARN) identifying your CloudFront Function.
     */
    arn?: pulumi.Input<string>;
    /**
     * Source code of the function
     */
    code?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * ETag hash of the function
     */
    etag?: pulumi.Input<string>;
    /**
     * Unique name for your CloudFront Function.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
     */
    publish?: pulumi.Input<boolean>;
    /**
     * Identifier of the function's runtime. Currently only `cloudfront-js-1.0` is valid.
     */
    runtime?: pulumi.Input<string>;
    /**
     * Status of the function. Can be `UNPUBLISHED`, `UNASSOCIATED` or `ASSOCIATED`.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Function resource.
 */
export interface FunctionArgs {
    /**
     * Source code of the function
     */
    code: pulumi.Input<string>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Unique name for your CloudFront Function.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
     */
    publish?: pulumi.Input<boolean>;
    /**
     * Identifier of the function's runtime. Currently only `cloudfront-js-1.0` is valid.
     */
    runtime: pulumi.Input<string>;
}
