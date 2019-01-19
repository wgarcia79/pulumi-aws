// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to lookup current AWS partition in which Terraform is working
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const aws_partition_current = pulumi.output(aws.getPartition({}));
 * const aws_iam_policy_document_s3_policy = pulumi.output(aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: ["s3:ListBucket"],
 *         resources: [aws_partition_current.apply(__arg0 => `arn:${__arg0.partition}:s3:::my-bucket`)],
 *         sid: "1",
 *     }],
 * }));
 * ```
 */
export function getPartition(opts?: pulumi.InvokeOptions): Promise<GetPartitionResult> {
    return pulumi.runtime.invoke("aws:index/getPartition:getPartition", {
    }, opts);
}

/**
 * A collection of values returned by getPartition.
 */
export interface GetPartitionResult {
    readonly partition: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
