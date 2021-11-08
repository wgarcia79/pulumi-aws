// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3
{
    /// <summary>
    /// Manages a S3 Bucket Notification Configuration. For additional information, see the [Configuring S3 Event Notifications section in the Amazon S3 Developer Guide](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html).
    /// 
    /// &gt; **NOTE:** S3 Buckets only support a single notification configuration. Declaring multiple `aws.s3.BucketNotification` resources to the same S3 Bucket will cause a perpetual difference in configuration. See the example "Trigger multiple Lambda functions" for an option.
    /// 
    /// ## Example Usage
    /// ### Add notification configuration to SNS Topic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
    ///         {
    ///         });
    ///         var topic = new Aws.Sns.Topic("topic", new Aws.Sns.TopicArgs
    ///         {
    ///             Policy = bucket.Arn.Apply(arn =&gt; @$"{{
    ///     ""Version"":""2012-10-17"",
    ///     ""Statement"":[{{
    ///         ""Effect"": ""Allow"",
    ///         ""Principal"": {{ ""Service"": ""s3.amazonaws.com"" }},
    ///         ""Action"": ""SNS:Publish"",
    ///         ""Resource"": ""arn:aws:sns:*:*:s3-event-notification-topic"",
    ///         ""Condition"":{{
    ///             ""ArnLike"":{{""aws:SourceArn"":""{arn}""}}
    ///         }}
    ///     }}]
    /// }}
    /// "),
    ///         });
    ///         var bucketNotification = new Aws.S3.BucketNotification("bucketNotification", new Aws.S3.BucketNotificationArgs
    ///         {
    ///             Bucket = bucket.Id,
    ///             Topics = 
    ///             {
    ///                 new Aws.S3.Inputs.BucketNotificationTopicArgs
    ///                 {
    ///                     TopicArn = topic.Arn,
    ///                     Events = 
    ///                     {
    ///                         "s3:ObjectCreated:*",
    ///                     },
    ///                     FilterSuffix = ".log",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Add notification configuration to SQS Queue
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
    ///         {
    ///         });
    ///         var queue = new Aws.Sqs.Queue("queue", new Aws.Sqs.QueueArgs
    ///         {
    ///             Policy = bucket.Arn.Apply(arn =&gt; @$"{{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///     {{
    ///       ""Effect"": ""Allow"",
    ///       ""Principal"": ""*"",
    ///       ""Action"": ""sqs:SendMessage"",
    /// 	  ""Resource"": ""arn:aws:sqs:*:*:s3-event-notification-queue"",
    ///       ""Condition"": {{
    ///         ""ArnEquals"": {{ ""aws:SourceArn"": ""{arn}"" }}
    ///       }}
    ///     }}
    ///   ]
    /// }}
    /// "),
    ///         });
    ///         var bucketNotification = new Aws.S3.BucketNotification("bucketNotification", new Aws.S3.BucketNotificationArgs
    ///         {
    ///             Bucket = bucket.Id,
    ///             Queues = 
    ///             {
    ///                 new Aws.S3.Inputs.BucketNotificationQueueArgs
    ///                 {
    ///                     QueueArn = queue.Arn,
    ///                     Events = 
    ///                     {
    ///                         "s3:ObjectCreated:*",
    ///                     },
    ///                     FilterSuffix = ".log",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Add notification configuration to Lambda Function
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var iamForLambda = new Aws.Iam.Role("iamForLambda", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = @"{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///     {
    ///       ""Action"": ""sts:AssumeRole"",
    ///       ""Principal"": {
    ///         ""Service"": ""lambda.amazonaws.com""
    ///       },
    ///       ""Effect"": ""Allow""
    ///     }
    ///   ]
    /// }
    /// ",
    ///         });
    ///         var func = new Aws.Lambda.Function("func", new Aws.Lambda.FunctionArgs
    ///         {
    ///             Code = new FileArchive("your-function.zip"),
    ///             Role = iamForLambda.Arn,
    ///             Handler = "exports.example",
    ///             Runtime = "go1.x",
    ///         });
    ///         var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
    ///         {
    ///         });
    ///         var allowBucket = new Aws.Lambda.Permission("allowBucket", new Aws.Lambda.PermissionArgs
    ///         {
    ///             Action = "lambda:InvokeFunction",
    ///             Function = func.Arn,
    ///             Principal = "s3.amazonaws.com",
    ///             SourceArn = bucket.Arn,
    ///         });
    ///         var bucketNotification = new Aws.S3.BucketNotification("bucketNotification", new Aws.S3.BucketNotificationArgs
    ///         {
    ///             Bucket = bucket.Id,
    ///             LambdaFunctions = 
    ///             {
    ///                 new Aws.S3.Inputs.BucketNotificationLambdaFunctionArgs
    ///                 {
    ///                     LambdaFunctionArn = func.Arn,
    ///                     Events = 
    ///                     {
    ///                         "s3:ObjectCreated:*",
    ///                     },
    ///                     FilterPrefix = "AWSLogs/",
    ///                     FilterSuffix = ".log",
    ///                 },
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 allowBucket,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Trigger multiple Lambda functions
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var iamForLambda = new Aws.Iam.Role("iamForLambda", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = @"{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///     {
    ///       ""Action"": ""sts:AssumeRole"",
    ///       ""Principal"": {
    ///         ""Service"": ""lambda.amazonaws.com""
    ///       },
    ///       ""Effect"": ""Allow""
    ///     }
    ///   ]
    /// }
    /// ",
    ///         });
    ///         var func1 = new Aws.Lambda.Function("func1", new Aws.Lambda.FunctionArgs
    ///         {
    ///             Code = new FileArchive("your-function1.zip"),
    ///             Role = iamForLambda.Arn,
    ///             Handler = "exports.example",
    ///             Runtime = "go1.x",
    ///         });
    ///         var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
    ///         {
    ///         });
    ///         var allowBucket1 = new Aws.Lambda.Permission("allowBucket1", new Aws.Lambda.PermissionArgs
    ///         {
    ///             Action = "lambda:InvokeFunction",
    ///             Function = func1.Arn,
    ///             Principal = "s3.amazonaws.com",
    ///             SourceArn = bucket.Arn,
    ///         });
    ///         var func2 = new Aws.Lambda.Function("func2", new Aws.Lambda.FunctionArgs
    ///         {
    ///             Code = new FileArchive("your-function2.zip"),
    ///             Role = iamForLambda.Arn,
    ///             Handler = "exports.example",
    ///         });
    ///         var allowBucket2 = new Aws.Lambda.Permission("allowBucket2", new Aws.Lambda.PermissionArgs
    ///         {
    ///             Action = "lambda:InvokeFunction",
    ///             Function = func2.Arn,
    ///             Principal = "s3.amazonaws.com",
    ///             SourceArn = bucket.Arn,
    ///         });
    ///         var bucketNotification = new Aws.S3.BucketNotification("bucketNotification", new Aws.S3.BucketNotificationArgs
    ///         {
    ///             Bucket = bucket.Id,
    ///             LambdaFunctions = 
    ///             {
    ///                 new Aws.S3.Inputs.BucketNotificationLambdaFunctionArgs
    ///                 {
    ///                     LambdaFunctionArn = func1.Arn,
    ///                     Events = 
    ///                     {
    ///                         "s3:ObjectCreated:*",
    ///                     },
    ///                     FilterPrefix = "AWSLogs/",
    ///                     FilterSuffix = ".log",
    ///                 },
    ///                 new Aws.S3.Inputs.BucketNotificationLambdaFunctionArgs
    ///                 {
    ///                     LambdaFunctionArn = func2.Arn,
    ///                     Events = 
    ///                     {
    ///                         "s3:ObjectCreated:*",
    ///                     },
    ///                     FilterPrefix = "OtherLogs/",
    ///                     FilterSuffix = ".log",
    ///                 },
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 allowBucket1,
    ///                 allowBucket2,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Add multiple notification configurations to SQS Queue
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
    ///         {
    ///         });
    ///         var queue = new Aws.Sqs.Queue("queue", new Aws.Sqs.QueueArgs
    ///         {
    ///             Policy = bucket.Arn.Apply(arn =&gt; @$"{{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///     {{
    ///       ""Effect"": ""Allow"",
    ///       ""Principal"": ""*"",
    ///       ""Action"": ""sqs:SendMessage"",
    /// 	  ""Resource"": ""arn:aws:sqs:*:*:s3-event-notification-queue"",
    ///       ""Condition"": {{
    ///         ""ArnEquals"": {{ ""aws:SourceArn"": ""{arn}"" }}
    ///       }}
    ///     }}
    ///   ]
    /// }}
    /// "),
    ///         });
    ///         var bucketNotification = new Aws.S3.BucketNotification("bucketNotification", new Aws.S3.BucketNotificationArgs
    ///         {
    ///             Bucket = bucket.Id,
    ///             Queues = 
    ///             {
    ///                 new Aws.S3.Inputs.BucketNotificationQueueArgs
    ///                 {
    ///                     Id = "image-upload-event",
    ///                     QueueArn = queue.Arn,
    ///                     Events = 
    ///                     {
    ///                         "s3:ObjectCreated:*",
    ///                     },
    ///                     FilterPrefix = "images/",
    ///                 },
    ///                 new Aws.S3.Inputs.BucketNotificationQueueArgs
    ///                 {
    ///                     Id = "video-upload-event",
    ///                     QueueArn = queue.Arn,
    ///                     Events = 
    ///                     {
    ///                         "s3:ObjectCreated:*",
    ///                     },
    ///                     FilterPrefix = "videos/",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// S3 bucket notification can be imported using the `bucket`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:s3/bucketNotification:BucketNotification bucket_notification bucket-name
    /// ```
    /// </summary>
    [AwsResourceType("aws:s3/bucketNotification:BucketNotification")]
    public partial class BucketNotification : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the bucket to put notification configuration.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// Used to configure notifications to a Lambda Function (documented below).
        /// </summary>
        [Output("lambdaFunctions")]
        public Output<ImmutableArray<Outputs.BucketNotificationLambdaFunction>> LambdaFunctions { get; private set; } = null!;

        /// <summary>
        /// The notification configuration to SQS Queue (documented below).
        /// </summary>
        [Output("queues")]
        public Output<ImmutableArray<Outputs.BucketNotificationQueue>> Queues { get; private set; } = null!;

        /// <summary>
        /// The notification configuration to SNS Topic (documented below).
        /// </summary>
        [Output("topics")]
        public Output<ImmutableArray<Outputs.BucketNotificationTopic>> Topics { get; private set; } = null!;


        /// <summary>
        /// Create a BucketNotification resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketNotification(string name, BucketNotificationArgs args, CustomResourceOptions? options = null)
            : base("aws:s3/bucketNotification:BucketNotification", name, args ?? new BucketNotificationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BucketNotification(string name, Input<string> id, BucketNotificationState? state = null, CustomResourceOptions? options = null)
            : base("aws:s3/bucketNotification:BucketNotification", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BucketNotification resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketNotification Get(string name, Input<string> id, BucketNotificationState? state = null, CustomResourceOptions? options = null)
        {
            return new BucketNotification(name, id, state, options);
        }
    }

    public sealed class BucketNotificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket to put notification configuration.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("lambdaFunctions")]
        private InputList<Inputs.BucketNotificationLambdaFunctionArgs>? _lambdaFunctions;

        /// <summary>
        /// Used to configure notifications to a Lambda Function (documented below).
        /// </summary>
        public InputList<Inputs.BucketNotificationLambdaFunctionArgs> LambdaFunctions
        {
            get => _lambdaFunctions ?? (_lambdaFunctions = new InputList<Inputs.BucketNotificationLambdaFunctionArgs>());
            set => _lambdaFunctions = value;
        }

        [Input("queues")]
        private InputList<Inputs.BucketNotificationQueueArgs>? _queues;

        /// <summary>
        /// The notification configuration to SQS Queue (documented below).
        /// </summary>
        public InputList<Inputs.BucketNotificationQueueArgs> Queues
        {
            get => _queues ?? (_queues = new InputList<Inputs.BucketNotificationQueueArgs>());
            set => _queues = value;
        }

        [Input("topics")]
        private InputList<Inputs.BucketNotificationTopicArgs>? _topics;

        /// <summary>
        /// The notification configuration to SNS Topic (documented below).
        /// </summary>
        public InputList<Inputs.BucketNotificationTopicArgs> Topics
        {
            get => _topics ?? (_topics = new InputList<Inputs.BucketNotificationTopicArgs>());
            set => _topics = value;
        }

        public BucketNotificationArgs()
        {
        }
    }

    public sealed class BucketNotificationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket to put notification configuration.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("lambdaFunctions")]
        private InputList<Inputs.BucketNotificationLambdaFunctionGetArgs>? _lambdaFunctions;

        /// <summary>
        /// Used to configure notifications to a Lambda Function (documented below).
        /// </summary>
        public InputList<Inputs.BucketNotificationLambdaFunctionGetArgs> LambdaFunctions
        {
            get => _lambdaFunctions ?? (_lambdaFunctions = new InputList<Inputs.BucketNotificationLambdaFunctionGetArgs>());
            set => _lambdaFunctions = value;
        }

        [Input("queues")]
        private InputList<Inputs.BucketNotificationQueueGetArgs>? _queues;

        /// <summary>
        /// The notification configuration to SQS Queue (documented below).
        /// </summary>
        public InputList<Inputs.BucketNotificationQueueGetArgs> Queues
        {
            get => _queues ?? (_queues = new InputList<Inputs.BucketNotificationQueueGetArgs>());
            set => _queues = value;
        }

        [Input("topics")]
        private InputList<Inputs.BucketNotificationTopicGetArgs>? _topics;

        /// <summary>
        /// The notification configuration to SNS Topic (documented below).
        /// </summary>
        public InputList<Inputs.BucketNotificationTopicGetArgs> Topics
        {
            get => _topics ?? (_topics = new InputList<Inputs.BucketNotificationTopicGetArgs>());
            set => _topics = value;
        }

        public BucketNotificationState()
        {
        }
    }
}
