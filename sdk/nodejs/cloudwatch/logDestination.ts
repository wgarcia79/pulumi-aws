// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CloudWatch Logs destination resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testDestination = new aws.cloudwatch.LogDestination("testDestination", {
 *     roleArn: aws_iam_role.iam_for_cloudwatch.arn,
 *     targetArn: aws_kinesis_stream.kinesis_for_cloudwatch.arn,
 * });
 * ```
 *
 * ## Import
 *
 * CloudWatch Logs destinations can be imported using the `name`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:cloudwatch/logDestination:LogDestination test_destination test_destination
 * ```
 */
export class LogDestination extends pulumi.CustomResource {
    /**
     * Get an existing LogDestination resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogDestinationState, opts?: pulumi.CustomResourceOptions): LogDestination {
        return new LogDestination(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudwatch/logDestination:LogDestination';

    /**
     * Returns true if the given object is an instance of LogDestination.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogDestination {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogDestination.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) specifying the log destination.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A name for the log destination
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * The ARN of the target Amazon Kinesis stream resource for the destination
     */
    public readonly targetArn!: pulumi.Output<string>;

    /**
     * Create a LogDestination resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LogDestinationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogDestinationArgs | LogDestinationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogDestinationState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["targetArn"] = state ? state.targetArn : undefined;
        } else {
            const args = argsOrState as LogDestinationArgs | undefined;
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            if ((!args || args.targetArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetArn'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["targetArn"] = args ? args.targetArn : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LogDestination.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogDestination resources.
 */
export interface LogDestinationState {
    /**
     * The Amazon Resource Name (ARN) specifying the log destination.
     */
    arn?: pulumi.Input<string>;
    /**
     * A name for the log destination
     */
    name?: pulumi.Input<string>;
    /**
     * The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target
     */
    roleArn?: pulumi.Input<string>;
    /**
     * The ARN of the target Amazon Kinesis stream resource for the destination
     */
    targetArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogDestination resource.
 */
export interface LogDestinationArgs {
    /**
     * A name for the log destination
     */
    name?: pulumi.Input<string>;
    /**
     * The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target
     */
    roleArn: pulumi.Input<string>;
    /**
     * The ARN of the target Amazon Kinesis stream resource for the destination
     */
    targetArn: pulumi.Input<string>;
}
