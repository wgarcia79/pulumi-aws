// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an Amazon CloudFront origin access identity.
 *
 * For information about CloudFront distributions, see the
 * [Amazon CloudFront Developer Guide](http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Introduction.html). For more information on generating
 * origin access identities, see
 * [Using an Origin Access Identity to Restrict Access to Your Amazon S3 Content][2].
 *
 * ## Example Usage
 *
 * The following example below creates a CloudFront origin access identity.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.cloudfront.OriginAccessIdentity("example", {
 *     comment: "Some comment",
 * });
 * ```
 * ## Using With CloudFront
 *
 * Normally, when referencing an origin access identity in CloudFront, you need to
 * prefix the ID with the `origin-access-identity/cloudfront/` special path.
 * The `cloudfrontAccessIdentityPath` allows this to be circumvented.
 * The below snippet demonstrates use with the `s3OriginConfig` structure for the
 * `aws.cloudfront.Distribution` resource:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // ... other configuration ...
 * const example = new aws.cloudfront.Distribution("example", {origins: [{
 *     s3OriginConfig: {
 *         originAccessIdentity: aws_cloudfront_origin_access_identity.example.cloudfront_access_identity_path,
 *     },
 * }]});
 * ```
 *
 * ### Updating your bucket policy
 *
 * Note that the AWS API may translate the `s3CanonicalUserId` `CanonicalUser`
 * principal into an `AWS` IAM ARN principal when supplied in an
 * `aws.s3.Bucket` bucket policy, causing spurious diffs. If
 * you see this behaviour, use the `iamArn` instead:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const s3Policy = aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: ["s3:GetObject"],
 *         resources: [`${aws_s3_bucket.example.arn}/*`],
 *         principals: [{
 *             type: "AWS",
 *             identifiers: [aws_cloudfront_origin_access_identity.example.iam_arn],
 *         }],
 *     }],
 * });
 * const example = new aws.s3.BucketPolicy("example", {
 *     bucket: aws_s3_bucket.example.id,
 *     policy: s3Policy.then(s3Policy => s3Policy.json),
 * });
 * ```
 *
 * [1]: http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Introduction.html
 * [2]: http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html
 *
 * ## Import
 *
 * Cloudfront Origin Access Identities can be imported using the `id`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:cloudfront/originAccessIdentity:OriginAccessIdentity origin_access E74FTE3AEXAMPLE
 * ```
 */
export class OriginAccessIdentity extends pulumi.CustomResource {
    /**
     * Get an existing OriginAccessIdentity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OriginAccessIdentityState, opts?: pulumi.CustomResourceOptions): OriginAccessIdentity {
        return new OriginAccessIdentity(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudfront/originAccessIdentity:OriginAccessIdentity';

    /**
     * Returns true if the given object is an instance of OriginAccessIdentity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OriginAccessIdentity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OriginAccessIdentity.__pulumiType;
    }

    /**
     * Internal value used by CloudFront to allow future
     * updates to the origin access identity.
     */
    public /*out*/ readonly callerReference!: pulumi.Output<string>;
    /**
     * A shortcut to the full path for the
     * origin access identity to use in CloudFront, see below.
     */
    public /*out*/ readonly cloudfrontAccessIdentityPath!: pulumi.Output<string>;
    /**
     * An optional comment for the origin access identity.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * The current version of the origin access identity's information.
     * For example: `E2QWRUHAPOMQZL`.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * A pre-generated ARN for use in S3 bucket policies (see below).
     * Example: `arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity
     * E2QWRUHAPOMQZL`.
     */
    public /*out*/ readonly iamArn!: pulumi.Output<string>;
    /**
     * The Amazon S3 canonical user ID for the origin
     * access identity, which you use when giving the origin access identity read
     * permission to an object in Amazon S3.
     */
    public /*out*/ readonly s3CanonicalUserId!: pulumi.Output<string>;

    /**
     * Create a OriginAccessIdentity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OriginAccessIdentityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OriginAccessIdentityArgs | OriginAccessIdentityState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OriginAccessIdentityState | undefined;
            inputs["callerReference"] = state ? state.callerReference : undefined;
            inputs["cloudfrontAccessIdentityPath"] = state ? state.cloudfrontAccessIdentityPath : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["iamArn"] = state ? state.iamArn : undefined;
            inputs["s3CanonicalUserId"] = state ? state.s3CanonicalUserId : undefined;
        } else {
            const args = argsOrState as OriginAccessIdentityArgs | undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["callerReference"] = undefined /*out*/;
            inputs["cloudfrontAccessIdentityPath"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["iamArn"] = undefined /*out*/;
            inputs["s3CanonicalUserId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(OriginAccessIdentity.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OriginAccessIdentity resources.
 */
export interface OriginAccessIdentityState {
    /**
     * Internal value used by CloudFront to allow future
     * updates to the origin access identity.
     */
    callerReference?: pulumi.Input<string>;
    /**
     * A shortcut to the full path for the
     * origin access identity to use in CloudFront, see below.
     */
    cloudfrontAccessIdentityPath?: pulumi.Input<string>;
    /**
     * An optional comment for the origin access identity.
     */
    comment?: pulumi.Input<string>;
    /**
     * The current version of the origin access identity's information.
     * For example: `E2QWRUHAPOMQZL`.
     */
    etag?: pulumi.Input<string>;
    /**
     * A pre-generated ARN for use in S3 bucket policies (see below).
     * Example: `arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity
     * E2QWRUHAPOMQZL`.
     */
    iamArn?: pulumi.Input<string>;
    /**
     * The Amazon S3 canonical user ID for the origin
     * access identity, which you use when giving the origin access identity read
     * permission to an object in Amazon S3.
     */
    s3CanonicalUserId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OriginAccessIdentity resource.
 */
export interface OriginAccessIdentityArgs {
    /**
     * An optional comment for the origin access identity.
     */
    comment?: pulumi.Input<string>;
}
