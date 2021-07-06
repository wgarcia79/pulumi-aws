// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch
{
    /// <summary>
    /// Provides an EventBridge Target resource.
    /// 
    /// &gt; **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var console = new Aws.CloudWatch.EventRule("console", new Aws.CloudWatch.EventRuleArgs
    ///         {
    ///             Description = "Capture all EC2 scaling events",
    ///             EventPattern = @"{
    ///   ""source"": [
    ///     ""aws.autoscaling""
    ///   ],
    ///   ""detail-type"": [
    ///     ""EC2 Instance Launch Successful"",
    ///     ""EC2 Instance Terminate Successful"",
    ///     ""EC2 Instance Launch Unsuccessful"",
    ///     ""EC2 Instance Terminate Unsuccessful""
    ///   ]
    /// }
    /// ",
    ///         });
    ///         var testStream = new Aws.Kinesis.Stream("testStream", new Aws.Kinesis.StreamArgs
    ///         {
    ///             ShardCount = 1,
    ///         });
    ///         var yada = new Aws.CloudWatch.EventTarget("yada", new Aws.CloudWatch.EventTargetArgs
    ///         {
    ///             Rule = console.Name,
    ///             Arn = testStream.Arn,
    ///             RunCommandTargets = 
    ///             {
    ///                 new Aws.CloudWatch.Inputs.EventTargetRunCommandTargetArgs
    ///                 {
    ///                     Key = "tag:Name",
    ///                     Values = 
    ///                     {
    ///                         "FooBar",
    ///                     },
    ///                 },
    ///                 new Aws.CloudWatch.Inputs.EventTargetRunCommandTargetArgs
    ///                 {
    ///                     Key = "InstanceIds",
    ///                     Values = 
    ///                     {
    ///                         "i-162058cd308bffec2",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Example RunCommand Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var stopInstancesEventRule = new Aws.CloudWatch.EventRule("stopInstancesEventRule", new Aws.CloudWatch.EventRuleArgs
    ///         {
    ///             Description = "Stop instances nightly",
    ///             ScheduleExpression = "cron(0 0 * * ? *)",
    ///         });
    ///         var stopInstancesEventTarget = new Aws.CloudWatch.EventTarget("stopInstancesEventTarget", new Aws.CloudWatch.EventTargetArgs
    ///         {
    ///             Arn = $"arn:aws:ssm:{@var.Aws_region}::document/AWS-RunShellScript",
    ///             Input = "{\"commands\":[\"halt\"]}",
    ///             Rule = stopInstancesEventRule.Name,
    ///             RoleArn = aws_iam_role.Ssm_lifecycle.Arn,
    ///             RunCommandTargets = 
    ///             {
    ///                 new Aws.CloudWatch.Inputs.EventTargetRunCommandTargetArgs
    ///                 {
    ///                     Key = "tag:Terminate",
    ///                     Values = 
    ///                     {
    ///                         "midnight",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Example API Gateway target
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleEventRule = new Aws.CloudWatch.EventRule("exampleEventRule", new Aws.CloudWatch.EventRuleArgs
    ///         {
    ///         });
    ///         // ...
    ///         var exampleDeployment = new Aws.ApiGateway.Deployment("exampleDeployment", new Aws.ApiGateway.DeploymentArgs
    ///         {
    ///             RestApi = aws_api_gateway_rest_api.Example.Id,
    ///         });
    ///         // ...
    ///         var exampleStage = new Aws.ApiGateway.Stage("exampleStage", new Aws.ApiGateway.StageArgs
    ///         {
    ///             RestApi = aws_api_gateway_rest_api.Example.Id,
    ///             Deployment = exampleDeployment.Id,
    ///         });
    ///         // ...
    ///         var exampleEventTarget = new Aws.CloudWatch.EventTarget("exampleEventTarget", new Aws.CloudWatch.EventTargetArgs
    ///         {
    ///             Arn = exampleStage.ExecutionArn.Apply(executionArn =&gt; $"{executionArn}/GET"),
    ///             Rule = exampleEventRule.Id,
    ///             HttpTarget = new Aws.CloudWatch.Inputs.EventTargetHttpTargetArgs
    ///             {
    ///                 QueryStringParameters = 
    ///                 {
    ///                     { "Body", "$.detail.body" },
    ///                 },
    ///                 HeaderParameters = 
    ///                 {
    ///                     { "Env", "Test" },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Example Input Transformer Usage - JSON Object
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleEventRule = new Aws.CloudWatch.EventRule("exampleEventRule", new Aws.CloudWatch.EventRuleArgs
    ///         {
    ///         });
    ///         // ...
    ///         var exampleEventTarget = new Aws.CloudWatch.EventTarget("exampleEventTarget", new Aws.CloudWatch.EventTargetArgs
    ///         {
    ///             Arn = aws_lambda_function.Example.Arn,
    ///             Rule = exampleEventRule.Id,
    ///             InputTransformer = new Aws.CloudWatch.Inputs.EventTargetInputTransformerArgs
    ///             {
    ///                 InputPaths = 
    ///                 {
    ///                     { "instance", "$.detail.instance" },
    ///                     { "status", "$.detail.status" },
    ///                 },
    ///                 InputTemplate = @"{
    ///   ""instance_id"": &lt;instance&gt;,
    ///   ""instance_status"": &lt;status&gt;
    /// }
    /// ",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Example Input Transformer Usage - Simple String
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleEventRule = new Aws.CloudWatch.EventRule("exampleEventRule", new Aws.CloudWatch.EventRuleArgs
    ///         {
    ///         });
    ///         // ...
    ///         var exampleEventTarget = new Aws.CloudWatch.EventTarget("exampleEventTarget", new Aws.CloudWatch.EventTargetArgs
    ///         {
    ///             Arn = aws_lambda_function.Example.Arn,
    ///             Rule = exampleEventRule.Id,
    ///             InputTransformer = new Aws.CloudWatch.Inputs.EventTargetInputTransformerArgs
    ///             {
    ///                 InputPaths = 
    ///                 {
    ///                     { "instance", "$.detail.instance" },
    ///                     { "status", "$.detail.status" },
    ///                 },
    ///                 InputTemplate = "\"&lt;instance&gt; is in state &lt;status&gt;\"",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// EventBridge Targets can be imported using `event_bus_name/rule-name/target-id` (if you omit `event_bus_name`, the `default` event bus will be used).
    /// 
    /// ```sh
    ///  $ pulumi import aws:cloudwatch/eventTarget:EventTarget test-event-target rule-name/target-id
    /// ```
    /// </summary>
    [AwsResourceType("aws:cloudwatch/eventTarget:EventTarget")]
    public partial class EventTarget : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the target.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Output("batchTarget")]
        public Output<Outputs.EventTargetBatchTarget?> BatchTarget { get; private set; } = null!;

        /// <summary>
        /// Parameters used when you are providing a dead letter config. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Output("deadLetterConfig")]
        public Output<Outputs.EventTargetDeadLetterConfig?> DeadLetterConfig { get; private set; } = null!;

        /// <summary>
        /// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Output("ecsTarget")]
        public Output<Outputs.EventTargetEcsTarget?> EcsTarget { get; private set; } = null!;

        /// <summary>
        /// The event bus to associate with the rule. If you omit this, the `default` event bus is used.
        /// </summary>
        [Output("eventBusName")]
        public Output<string?> EventBusName { get; private set; } = null!;

        /// <summary>
        /// Parameters used when you are using the rule to invoke an API Gateway REST endpoint. Documented below. A maximum of 1 is allowed.
        /// </summary>
        [Output("httpTarget")]
        public Output<Outputs.EventTargetHttpTarget?> HttpTarget { get; private set; } = null!;

        /// <summary>
        /// Valid JSON text passed to the target. Conflicts with `input_path` and `input_transformer`.
        /// </summary>
        [Output("input")]
        public Output<string?> Input { get; private set; } = null!;

        /// <summary>
        /// The value of the [JSONPath](http://goessner.net/articles/JsonPath/) that is used for extracting part of the matched event when passing it to the target. Conflicts with `input` and `input_transformer`.
        /// </summary>
        [Output("inputPath")]
        public Output<string?> InputPath { get; private set; } = null!;

        /// <summary>
        /// Parameters used when you are providing a custom input to a target based on certain event data. Conflicts with `input` and `input_path`.
        /// </summary>
        [Output("inputTransformer")]
        public Output<Outputs.EventTargetInputTransformer?> InputTransformer { get; private set; } = null!;

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Output("kinesisTarget")]
        public Output<Outputs.EventTargetKinesisTarget?> KinesisTarget { get; private set; } = null!;

        /// <summary>
        /// Parameters used when you are providing retry policies. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Output("retryPolicy")]
        public Output<Outputs.EventTargetRetryPolicy?> RetryPolicy { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecs_target` is used or target in `arn` is EC2 instance, Kinesis data stream or Step Functions state machine.
        /// </summary>
        [Output("roleArn")]
        public Output<string?> RoleArn { get; private set; } = null!;

        /// <summary>
        /// The name of the rule you want to add targets to.
        /// </summary>
        [Output("rule")]
        public Output<string> Rule { get; private set; } = null!;

        /// <summary>
        /// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
        /// </summary>
        [Output("runCommandTargets")]
        public Output<ImmutableArray<Outputs.EventTargetRunCommandTarget>> RunCommandTargets { get; private set; } = null!;

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Output("sqsTarget")]
        public Output<Outputs.EventTargetSqsTarget?> SqsTarget { get; private set; } = null!;

        /// <summary>
        /// The unique target assignment ID.  If missing, will generate a random, unique id.
        /// </summary>
        [Output("targetId")]
        public Output<string> TargetId { get; private set; } = null!;


        /// <summary>
        /// Create a EventTarget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventTarget(string name, EventTargetArgs args, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/eventTarget:EventTarget", name, args ?? new EventTargetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventTarget(string name, Input<string> id, EventTargetState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/eventTarget:EventTarget", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EventTarget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventTarget Get(string name, Input<string> id, EventTargetState? state = null, CustomResourceOptions? options = null)
        {
            return new EventTarget(name, id, state, options);
        }
    }

    public sealed class EventTargetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the target.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("batchTarget")]
        public Input<Inputs.EventTargetBatchTargetArgs>? BatchTarget { get; set; }

        /// <summary>
        /// Parameters used when you are providing a dead letter config. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("deadLetterConfig")]
        public Input<Inputs.EventTargetDeadLetterConfigArgs>? DeadLetterConfig { get; set; }

        /// <summary>
        /// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("ecsTarget")]
        public Input<Inputs.EventTargetEcsTargetArgs>? EcsTarget { get; set; }

        /// <summary>
        /// The event bus to associate with the rule. If you omit this, the `default` event bus is used.
        /// </summary>
        [Input("eventBusName")]
        public Input<string>? EventBusName { get; set; }

        /// <summary>
        /// Parameters used when you are using the rule to invoke an API Gateway REST endpoint. Documented below. A maximum of 1 is allowed.
        /// </summary>
        [Input("httpTarget")]
        public Input<Inputs.EventTargetHttpTargetArgs>? HttpTarget { get; set; }

        /// <summary>
        /// Valid JSON text passed to the target. Conflicts with `input_path` and `input_transformer`.
        /// </summary>
        [Input("input")]
        public Input<string>? Input { get; set; }

        /// <summary>
        /// The value of the [JSONPath](http://goessner.net/articles/JsonPath/) that is used for extracting part of the matched event when passing it to the target. Conflicts with `input` and `input_transformer`.
        /// </summary>
        [Input("inputPath")]
        public Input<string>? InputPath { get; set; }

        /// <summary>
        /// Parameters used when you are providing a custom input to a target based on certain event data. Conflicts with `input` and `input_path`.
        /// </summary>
        [Input("inputTransformer")]
        public Input<Inputs.EventTargetInputTransformerArgs>? InputTransformer { get; set; }

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("kinesisTarget")]
        public Input<Inputs.EventTargetKinesisTargetArgs>? KinesisTarget { get; set; }

        /// <summary>
        /// Parameters used when you are providing retry policies. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("retryPolicy")]
        public Input<Inputs.EventTargetRetryPolicyArgs>? RetryPolicy { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecs_target` is used or target in `arn` is EC2 instance, Kinesis data stream or Step Functions state machine.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// The name of the rule you want to add targets to.
        /// </summary>
        [Input("rule", required: true)]
        public Input<string> Rule { get; set; } = null!;

        [Input("runCommandTargets")]
        private InputList<Inputs.EventTargetRunCommandTargetArgs>? _runCommandTargets;

        /// <summary>
        /// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
        /// </summary>
        public InputList<Inputs.EventTargetRunCommandTargetArgs> RunCommandTargets
        {
            get => _runCommandTargets ?? (_runCommandTargets = new InputList<Inputs.EventTargetRunCommandTargetArgs>());
            set => _runCommandTargets = value;
        }

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("sqsTarget")]
        public Input<Inputs.EventTargetSqsTargetArgs>? SqsTarget { get; set; }

        /// <summary>
        /// The unique target assignment ID.  If missing, will generate a random, unique id.
        /// </summary>
        [Input("targetId")]
        public Input<string>? TargetId { get; set; }

        public EventTargetArgs()
        {
        }
    }

    public sealed class EventTargetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the target.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("batchTarget")]
        public Input<Inputs.EventTargetBatchTargetGetArgs>? BatchTarget { get; set; }

        /// <summary>
        /// Parameters used when you are providing a dead letter config. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("deadLetterConfig")]
        public Input<Inputs.EventTargetDeadLetterConfigGetArgs>? DeadLetterConfig { get; set; }

        /// <summary>
        /// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("ecsTarget")]
        public Input<Inputs.EventTargetEcsTargetGetArgs>? EcsTarget { get; set; }

        /// <summary>
        /// The event bus to associate with the rule. If you omit this, the `default` event bus is used.
        /// </summary>
        [Input("eventBusName")]
        public Input<string>? EventBusName { get; set; }

        /// <summary>
        /// Parameters used when you are using the rule to invoke an API Gateway REST endpoint. Documented below. A maximum of 1 is allowed.
        /// </summary>
        [Input("httpTarget")]
        public Input<Inputs.EventTargetHttpTargetGetArgs>? HttpTarget { get; set; }

        /// <summary>
        /// Valid JSON text passed to the target. Conflicts with `input_path` and `input_transformer`.
        /// </summary>
        [Input("input")]
        public Input<string>? Input { get; set; }

        /// <summary>
        /// The value of the [JSONPath](http://goessner.net/articles/JsonPath/) that is used for extracting part of the matched event when passing it to the target. Conflicts with `input` and `input_transformer`.
        /// </summary>
        [Input("inputPath")]
        public Input<string>? InputPath { get; set; }

        /// <summary>
        /// Parameters used when you are providing a custom input to a target based on certain event data. Conflicts with `input` and `input_path`.
        /// </summary>
        [Input("inputTransformer")]
        public Input<Inputs.EventTargetInputTransformerGetArgs>? InputTransformer { get; set; }

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("kinesisTarget")]
        public Input<Inputs.EventTargetKinesisTargetGetArgs>? KinesisTarget { get; set; }

        /// <summary>
        /// Parameters used when you are providing retry policies. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("retryPolicy")]
        public Input<Inputs.EventTargetRetryPolicyGetArgs>? RetryPolicy { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecs_target` is used or target in `arn` is EC2 instance, Kinesis data stream or Step Functions state machine.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// The name of the rule you want to add targets to.
        /// </summary>
        [Input("rule")]
        public Input<string>? Rule { get; set; }

        [Input("runCommandTargets")]
        private InputList<Inputs.EventTargetRunCommandTargetGetArgs>? _runCommandTargets;

        /// <summary>
        /// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
        /// </summary>
        public InputList<Inputs.EventTargetRunCommandTargetGetArgs> RunCommandTargets
        {
            get => _runCommandTargets ?? (_runCommandTargets = new InputList<Inputs.EventTargetRunCommandTargetGetArgs>());
            set => _runCommandTargets = value;
        }

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("sqsTarget")]
        public Input<Inputs.EventTargetSqsTargetGetArgs>? SqsTarget { get; set; }

        /// <summary>
        /// The unique target assignment ID.  If missing, will generate a random, unique id.
        /// </summary>
        [Input("targetId")]
        public Input<string>? TargetId { get; set; }

        public EventTargetState()
        {
        }
    }
}
