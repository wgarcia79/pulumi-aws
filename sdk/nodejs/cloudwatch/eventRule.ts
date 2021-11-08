// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an EventBridge Rule resource.
 *
 * > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const console = new aws.cloudwatch.EventRule("console", {
 *     description: "Capture each AWS Console Sign In",
 *     eventPattern: `{
 *   "detail-type": [
 *     "AWS Console Sign In via CloudTrail"
 *   ]
 * }
 * `,
 * });
 * const awsLogins = new aws.sns.Topic("awsLogins", {});
 * const sns = new aws.cloudwatch.EventTarget("sns", {
 *     rule: console.name,
 *     arn: awsLogins.arn,
 * });
 * const snsTopicPolicy = awsLogins.arn.apply(arn => aws.iam.getPolicyDocument({
 *     statements: [{
 *         effect: "Allow",
 *         actions: ["SNS:Publish"],
 *         principals: [{
 *             type: "Service",
 *             identifiers: ["events.amazonaws.com"],
 *         }],
 *         resources: [arn],
 *     }],
 * }));
 * const _default = new aws.sns.TopicPolicy("default", {
 *     arn: awsLogins.arn,
 *     policy: snsTopicPolicy.apply(snsTopicPolicy => snsTopicPolicy.json),
 * });
 * ```
 *
 * ## Import
 *
 * EventBridge Rules can be imported using the `event_bus_name/rule_name` (if you omit `event_bus_name`, the `default` event bus will be used), e.g.,
 *
 * ```sh
 *  $ pulumi import aws:cloudwatch/eventRule:EventRule console example-event-bus/capture-console-sign-in
 * ```
 */
export class EventRule extends pulumi.CustomResource {
    /**
     * Get an existing EventRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventRuleState, opts?: pulumi.CustomResourceOptions): EventRule {
        return new EventRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudwatch/eventRule:EventRule';

    /**
     * Returns true if the given object is an instance of EventRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventRule.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the rule.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The description of the rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The event bus to associate with this rule. If you omit this, the `default` event bus is used.
     */
    public readonly eventBusName!: pulumi.Output<string | undefined>;
    /**
     * The event pattern described a JSON object. At least one of `scheduleExpression` or `eventPattern` is required. See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details.
     */
    public readonly eventPattern!: pulumi.Output<string | undefined>;
    /**
     * Whether the rule should be enabled (defaults to `true`).
     */
    public readonly isEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;
    /**
     * The scheduling expression. For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `scheduleExpression` or `eventPattern` is required. Can only be used on the default event bus. For more information, refer to the AWS documentation [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
     */
    public readonly scheduleExpression!: pulumi.Output<string | undefined>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a EventRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EventRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventRuleArgs | EventRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventRuleState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["eventBusName"] = state ? state.eventBusName : undefined;
            inputs["eventPattern"] = state ? state.eventPattern : undefined;
            inputs["isEnabled"] = state ? state.isEnabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["scheduleExpression"] = state ? state.scheduleExpression : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as EventRuleArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["eventBusName"] = args ? args.eventBusName : undefined;
            inputs["eventPattern"] = args ? args.eventPattern : undefined;
            inputs["isEnabled"] = args ? args.isEnabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["scheduleExpression"] = args ? args.scheduleExpression : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["tagsAll"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EventRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventRule resources.
 */
export interface EventRuleState {
    /**
     * The Amazon Resource Name (ARN) of the rule.
     */
    arn?: pulumi.Input<string>;
    /**
     * The description of the rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The event bus to associate with this rule. If you omit this, the `default` event bus is used.
     */
    eventBusName?: pulumi.Input<string>;
    /**
     * The event pattern described a JSON object. At least one of `scheduleExpression` or `eventPattern` is required. See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details.
     */
    eventPattern?: pulumi.Input<string>;
    /**
     * Whether the rule should be enabled (defaults to `true`).
     */
    isEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * The scheduling expression. For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `scheduleExpression` or `eventPattern` is required. Can only be used on the default event bus. For more information, refer to the AWS documentation [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
     */
    scheduleExpression?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a EventRule resource.
 */
export interface EventRuleArgs {
    /**
     * The description of the rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The event bus to associate with this rule. If you omit this, the `default` event bus is used.
     */
    eventBusName?: pulumi.Input<string>;
    /**
     * The event pattern described a JSON object. At least one of `scheduleExpression` or `eventPattern` is required. See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details.
     */
    eventPattern?: pulumi.Input<string>;
    /**
     * Whether the rule should be enabled (defaults to `true`).
     */
    isEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * The scheduling expression. For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `scheduleExpression` or `eventPattern` is required. Can only be used on the default event bus. For more information, refer to the AWS documentation [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
     */
    scheduleExpression?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
