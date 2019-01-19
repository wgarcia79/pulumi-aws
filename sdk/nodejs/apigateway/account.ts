// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a settings of an API Gateway Account. Settings is applied region-wide per `provider` block.
 * 
 * -> **Note:** As there is no API method for deleting account settings or resetting it to defaults, destroying this resource will keep your account settings intact
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const aws_iam_role_cloudwatch = new aws.iam.Role("cloudwatch", {
 *     assumeRolePolicy: "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Sid\": \"\",\n      \"Effect\": \"Allow\",\n      \"Principal\": {\n        \"Service\": \"apigateway.amazonaws.com\"\n      },\n      \"Action\": \"sts:AssumeRole\"\n    }\n  ]\n}\n",
 *     name: "api_gateway_cloudwatch_global",
 * });
 * const aws_api_gateway_account_demo = new aws.apigateway.Account("demo", {
 *     cloudwatchRoleArn: aws_iam_role_cloudwatch.arn,
 * });
 * const aws_iam_role_policy_cloudwatch = new aws.iam.RolePolicy("cloudwatch", {
 *     name: "default",
 *     policy: "{\n    \"Version\": \"2012-10-17\",\n    \"Statement\": [\n        {\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"logs:CreateLogGroup\",\n                \"logs:CreateLogStream\",\n                \"logs:DescribeLogGroups\",\n                \"logs:DescribeLogStreams\",\n                \"logs:PutLogEvents\",\n                \"logs:GetLogEvents\",\n                \"logs:FilterLogEvents\"\n            ],\n            \"Resource\": \"*\"\n        }\n    ]\n}\n",
 *     role: aws_iam_role_cloudwatch.id,
 * });
 * ```
 */
export class Account extends pulumi.CustomResource {
    /**
     * Get an existing Account resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountState, opts?: pulumi.CustomResourceOptions): Account {
        return new Account(name, <any>state, { ...opts, id: id });
    }

    /**
     * The ARN of an IAM role for CloudWatch (to allow logging & monitoring).
     * See more [in AWS Docs](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-stage-settings.html#how-to-stage-settings-console).
     * Logging & monitoring can be enabled/disabled and otherwise tuned on the API Gateway Stage level.
     */
    public readonly cloudwatchRoleArn: pulumi.Output<string | undefined>;
    /**
     * Account-Level throttle settings. See exported fields below.
     */
    public /*out*/ readonly throttleSettings: pulumi.Output<{ burstLimit: number, rateLimit: number }>;

    /**
     * Create a Account resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountArgs | AccountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: AccountState = argsOrState as AccountState | undefined;
            inputs["cloudwatchRoleArn"] = state ? state.cloudwatchRoleArn : undefined;
            inputs["throttleSettings"] = state ? state.throttleSettings : undefined;
        } else {
            const args = argsOrState as AccountArgs | undefined;
            inputs["cloudwatchRoleArn"] = args ? args.cloudwatchRoleArn : undefined;
            inputs["throttleSettings"] = undefined /*out*/;
        }
        super("aws:apigateway/account:Account", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Account resources.
 */
export interface AccountState {
    /**
     * The ARN of an IAM role for CloudWatch (to allow logging & monitoring).
     * See more [in AWS Docs](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-stage-settings.html#how-to-stage-settings-console).
     * Logging & monitoring can be enabled/disabled and otherwise tuned on the API Gateway Stage level.
     */
    readonly cloudwatchRoleArn?: pulumi.Input<string>;
    /**
     * Account-Level throttle settings. See exported fields below.
     */
    readonly throttleSettings?: pulumi.Input<{ burstLimit?: pulumi.Input<number>, rateLimit?: pulumi.Input<number> }>;
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    /**
     * The ARN of an IAM role for CloudWatch (to allow logging & monitoring).
     * See more [in AWS Docs](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-stage-settings.html#how-to-stage-settings-console).
     * Logging & monitoring can be enabled/disabled and otherwise tuned on the API Gateway Stage level.
     */
    readonly cloudwatchRoleArn?: pulumi.Input<string>;
}
