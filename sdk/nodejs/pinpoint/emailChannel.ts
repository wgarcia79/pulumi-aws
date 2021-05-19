// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Pinpoint Email Channel resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const app = new aws.pinpoint.App("app", {});
 * const identity = new aws.ses.DomainIdentity("identity", {domain: "example.com"});
 * const role = new aws.iam.Role("role", {assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "pinpoint.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `});
 * const email = new aws.pinpoint.EmailChannel("email", {
 *     applicationId: app.applicationId,
 *     fromAddress: "user@example.com",
 *     identity: identity.arn,
 *     roleArn: role.arn,
 * });
 * const rolePolicy = new aws.iam.RolePolicy("rolePolicy", {
 *     role: role.id,
 *     policy: `{
 *   "Version": "2012-10-17",
 *   "Statement": {
 *     "Action": [
 *       "mobileanalytics:PutEvents",
 *       "mobileanalytics:PutItems"
 *     ],
 *     "Effect": "Allow",
 *     "Resource": [
 *       "*"
 *     ]
 *   }
 * }
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * Pinpoint Email Channel can be imported using the `application-id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:pinpoint/emailChannel:EmailChannel email application-id
 * ```
 */
export class EmailChannel extends pulumi.CustomResource {
    /**
     * Get an existing EmailChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EmailChannelState, opts?: pulumi.CustomResourceOptions): EmailChannel {
        return new EmailChannel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:pinpoint/emailChannel:EmailChannel';

    /**
     * Returns true if the given object is an instance of EmailChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EmailChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EmailChannel.__pulumiType;
    }

    /**
     * The application ID.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * The ARN of the Amazon SES configuration set that you want to apply to messages that you send through the channel.
     */
    public readonly configurationSet!: pulumi.Output<string | undefined>;
    /**
     * Whether the channel is enabled or disabled. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The email address used to send emails from.
     */
    public readonly fromAddress!: pulumi.Output<string>;
    /**
     * The ARN of an identity verified with SES.
     */
    public readonly identity!: pulumi.Output<string>;
    /**
     * Messages per second that can be sent.
     */
    public /*out*/ readonly messagesPerSecond!: pulumi.Output<number>;
    /**
     * The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;

    /**
     * Create a EmailChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EmailChannelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EmailChannelArgs | EmailChannelState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EmailChannelState | undefined;
            inputs["applicationId"] = state ? state.applicationId : undefined;
            inputs["configurationSet"] = state ? state.configurationSet : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["fromAddress"] = state ? state.fromAddress : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["messagesPerSecond"] = state ? state.messagesPerSecond : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
        } else {
            const args = argsOrState as EmailChannelArgs | undefined;
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            if ((!args || args.fromAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fromAddress'");
            }
            if ((!args || args.identity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identity'");
            }
            inputs["applicationId"] = args ? args.applicationId : undefined;
            inputs["configurationSet"] = args ? args.configurationSet : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["fromAddress"] = args ? args.fromAddress : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["messagesPerSecond"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EmailChannel.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EmailChannel resources.
 */
export interface EmailChannelState {
    /**
     * The application ID.
     */
    readonly applicationId?: pulumi.Input<string>;
    /**
     * The ARN of the Amazon SES configuration set that you want to apply to messages that you send through the channel.
     */
    readonly configurationSet?: pulumi.Input<string>;
    /**
     * Whether the channel is enabled or disabled. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The email address used to send emails from.
     */
    readonly fromAddress?: pulumi.Input<string>;
    /**
     * The ARN of an identity verified with SES.
     */
    readonly identity?: pulumi.Input<string>;
    /**
     * Messages per second that can be sent.
     */
    readonly messagesPerSecond?: pulumi.Input<number>;
    /**
     * The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
     */
    readonly roleArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EmailChannel resource.
 */
export interface EmailChannelArgs {
    /**
     * The application ID.
     */
    readonly applicationId: pulumi.Input<string>;
    /**
     * The ARN of the Amazon SES configuration set that you want to apply to messages that you send through the channel.
     */
    readonly configurationSet?: pulumi.Input<string>;
    /**
     * Whether the channel is enabled or disabled. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The email address used to send emails from.
     */
    readonly fromAddress: pulumi.Input<string>;
    /**
     * The ARN of an identity verified with SES.
     */
    readonly identity: pulumi.Input<string>;
    /**
     * The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
     */
    readonly roleArn?: pulumi.Input<string>;
}
