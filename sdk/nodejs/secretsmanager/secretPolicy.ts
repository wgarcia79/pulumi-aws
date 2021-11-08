// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage AWS Secrets Manager secret policy.
 *
 * ## Example Usage
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleSecret = new aws.secretsmanager.Secret("exampleSecret", {});
 * const exampleSecretPolicy = new aws.secretsmanager.SecretPolicy("exampleSecretPolicy", {
 *     secretArn: exampleSecret.arn,
 *     policy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 * 	{
 * 	  "Sid": "EnableAllPermissions",
 * 	  "Effect": "Allow",
 * 	  "Principal": {
 * 		"AWS": "*"
 * 	  },
 * 	  "Action": "secretsmanager:GetSecretValue",
 * 	  "Resource": "*"
 * 	}
 *   ]
 * }
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * `aws_secretsmanager_secret_policy` can be imported by using the secret Amazon Resource Name (ARN), e.g.,
 *
 * ```sh
 *  $ pulumi import aws:secretsmanager/secretPolicy:SecretPolicy example arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456
 * ```
 */
export class SecretPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SecretPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretPolicyState, opts?: pulumi.CustomResourceOptions): SecretPolicy {
        return new SecretPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:secretsmanager/secretPolicy:SecretPolicy';

    /**
     * Returns true if the given object is an instance of SecretPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretPolicy.__pulumiType;
    }

    /**
     * Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.
     */
    public readonly blockPublicPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * A valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html).
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * Secret ARN.
     */
    public readonly secretArn!: pulumi.Output<string>;

    /**
     * Create a SecretPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretPolicyArgs | SecretPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretPolicyState | undefined;
            inputs["blockPublicPolicy"] = state ? state.blockPublicPolicy : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["secretArn"] = state ? state.secretArn : undefined;
        } else {
            const args = argsOrState as SecretPolicyArgs | undefined;
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            if ((!args || args.secretArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretArn'");
            }
            inputs["blockPublicPolicy"] = args ? args.blockPublicPolicy : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["secretArn"] = args ? args.secretArn : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SecretPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretPolicy resources.
 */
export interface SecretPolicyState {
    /**
     * Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.
     */
    blockPublicPolicy?: pulumi.Input<boolean>;
    /**
     * A valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html).
     */
    policy?: pulumi.Input<string>;
    /**
     * Secret ARN.
     */
    secretArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretPolicy resource.
 */
export interface SecretPolicyArgs {
    /**
     * Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.
     */
    blockPublicPolicy?: pulumi.Input<boolean>;
    /**
     * A valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html).
     */
    policy: pulumi.Input<string>;
    /**
     * Secret ARN.
     */
    secretArn: pulumi.Input<string>;
}
