// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides information about a CloudFront Function.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const config = new pulumi.Config();
 * const functionName = config.require("functionName");
 * const existing = aws.cloudtrail.getFunction({
 *     name: functionName,
 * });
 * ```
 */
export function getFunction(args: GetFunctionArgs, opts?: pulumi.InvokeOptions): Promise<GetFunctionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:cloudtrail/getFunction:getFunction", {
        "name": args.name,
        "stage": args.stage,
    }, opts);
}

/**
 * A collection of arguments for invoking getFunction.
 */
export interface GetFunctionArgs {
    /**
     * Name of the CloudFront function.
     */
    name: string;
    /**
     * The function’s stage, either `DEVELOPMENT` or `LIVE`.
     */
    stage: string;
}

/**
 * A collection of values returned by getFunction.
 */
export interface GetFunctionResult {
    /**
     * Amazon Resource Name (ARN) identifying your CloudFront Function.
     */
    readonly arn: string;
    /**
     * Source code of the function
     */
    readonly code: string;
    /**
     * Comment.
     */
    readonly comment: string;
    /**
     * ETag hash of the function
     */
    readonly etag: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * When this resource was last modified.
     */
    readonly lastModifiedTime: string;
    readonly name: string;
    /**
     * Identifier of the function's runtime.
     */
    readonly runtime: string;
    readonly stage: string;
    /**
     * Status of the function. Can be `UNPUBLISHED`, `UNASSOCIATED` or `ASSOCIATED`.
     */
    readonly status: string;
}
