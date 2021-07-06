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
    /// Provides a CloudWatch Metric Stream resource.
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
    ///         // https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html
    ///         var metricStreamToFirehoseRole = new Aws.Iam.Role("metricStreamToFirehoseRole", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = @"{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///     {
    ///       ""Action"": ""sts:AssumeRole"",
    ///       ""Principal"": {
    ///         ""Service"": ""streams.metrics.cloudwatch.amazonaws.com""
    ///       },
    ///       ""Effect"": ""Allow"",
    ///       ""Sid"": """"
    ///     }
    ///   ]
    /// }
    /// ",
    ///         });
    ///         var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
    ///         {
    ///             Acl = "private",
    ///         });
    ///         var firehoseToS3Role = new Aws.Iam.Role("firehoseToS3Role", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = @"{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///     {
    ///       ""Action"": ""sts:AssumeRole"",
    ///       ""Principal"": {
    ///         ""Service"": ""firehose.amazonaws.com""
    ///       },
    ///       ""Effect"": ""Allow"",
    ///       ""Sid"": """"
    ///     }
    ///   ]
    /// }
    /// ",
    ///         });
    ///         var s3Stream = new Aws.Kinesis.FirehoseDeliveryStream("s3Stream", new Aws.Kinesis.FirehoseDeliveryStreamArgs
    ///         {
    ///             Destination = "s3",
    ///             S3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamS3ConfigurationArgs
    ///             {
    ///                 RoleArn = firehoseToS3Role.Arn,
    ///                 BucketArn = bucket.Arn,
    ///             },
    ///         });
    ///         var main = new Aws.CloudWatch.MetricStream("main", new Aws.CloudWatch.MetricStreamArgs
    ///         {
    ///             RoleArn = metricStreamToFirehoseRole.Arn,
    ///             FirehoseArn = s3Stream.Arn,
    ///             OutputFormat = "json",
    ///             IncludeFilters = 
    ///             {
    ///                 new Aws.CloudWatch.Inputs.MetricStreamIncludeFilterArgs
    ///                 {
    ///                     Namespace = "AWS/EC2",
    ///                 },
    ///                 new Aws.CloudWatch.Inputs.MetricStreamIncludeFilterArgs
    ///                 {
    ///                     Namespace = "AWS/EBS",
    ///                 },
    ///             },
    ///         });
    ///         // https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html
    ///         var metricStreamToFirehoseRolePolicy = new Aws.Iam.RolePolicy("metricStreamToFirehoseRolePolicy", new Aws.Iam.RolePolicyArgs
    ///         {
    ///             Role = metricStreamToFirehoseRole.Id,
    ///             Policy = s3Stream.Arn.Apply(arn =&gt; @$"{{
    ///     ""Version"": ""2012-10-17"",
    ///     ""Statement"": [
    ///         {{
    ///             ""Effect"": ""Allow"",
    ///             ""Action"": [
    ///                 ""firehose:PutRecord"",
    ///                 ""firehose:PutRecordBatch""
    ///             ],
    ///             ""Resource"": ""{arn}""
    ///         }}
    ///     ]
    /// }}
    /// "),
    ///         });
    ///         var firehoseToS3RolePolicy = new Aws.Iam.RolePolicy("firehoseToS3RolePolicy", new Aws.Iam.RolePolicyArgs
    ///         {
    ///             Role = firehoseToS3Role.Id,
    ///             Policy = Output.Tuple(bucket.Arn, bucket.Arn).Apply(values =&gt;
    ///             {
    ///                 var bucketArn = values.Item1;
    ///                 var bucketArn1 = values.Item2;
    ///                 return @$"{{
    ///     ""Version"": ""2012-10-17"",
    ///     ""Statement"": [
    ///         {{
    ///             ""Effect"": ""Allow"",
    ///             ""Action"": [
    ///                 ""s3:AbortMultipartUpload"",
    ///                 ""s3:GetBucketLocation"",
    ///                 ""s3:GetObject"",
    ///                 ""s3:ListBucket"",
    ///                 ""s3:ListBucketMultipartUploads"",
    ///                 ""s3:PutObject""
    ///             ],
    ///             ""Resource"": [
    ///                 ""{bucketArn}"",
    ///                 ""{bucketArn1}/*""
    ///             ]
    ///         }}
    ///     ]
    /// }}
    /// ";
    ///             }),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CloudWatch metric streams can be imported using the `name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:cloudwatch/metricStream:MetricStream sample &lt;name&gt;
    /// ```
    /// </summary>
    [AwsResourceType("aws:cloudwatch/metricStream:MetricStream")]
    public partial class MetricStream : Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the metric stream.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the metric stream was created.
        /// </summary>
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        /// <summary>
        /// List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces that you specify here. Conflicts with `include_filter`.
        /// </summary>
        [Output("excludeFilters")]
        public Output<ImmutableArray<Outputs.MetricStreamExcludeFilter>> ExcludeFilters { get; private set; } = null!;

        /// <summary>
        /// ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
        /// </summary>
        [Output("firehoseArn")]
        public Output<string> FirehoseArn { get; private set; } = null!;

        /// <summary>
        /// List of inclusive metric filters. If you specify this parameter, the stream sends only the metrics from the metric namespaces that you specify here. Conflicts with `exclude_filter`.
        /// </summary>
        [Output("includeFilters")]
        public Output<ImmutableArray<Outputs.MetricStreamIncludeFilter>> IncludeFilters { get; private set; } = null!;

        /// <summary>
        /// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the metric stream was last updated.
        /// </summary>
        [Output("lastUpdateDate")]
        public Output<string> LastUpdateDate { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// Output format for the stream. Possible values are `json` and `opentelemetry0.7`. For more information about output formats, see [Metric streams output formats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html).
        /// </summary>
        [Output("outputFormat")]
        public Output<string> OutputFormat { get; private set; } = null!;

        /// <summary>
        /// ARN of the IAM role that this metric stream will use to access Amazon Kinesis Firehose resources. For more information about role permissions, see [Trust between CloudWatch and Kinesis Data Firehose](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html).
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// State of the metric stream. Possible values are `running` and `stopped`.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a MetricStream resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MetricStream(string name, MetricStreamArgs args, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/metricStream:MetricStream", name, args ?? new MetricStreamArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MetricStream(string name, Input<string> id, MetricStreamState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/metricStream:MetricStream", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MetricStream resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MetricStream Get(string name, Input<string> id, MetricStreamState? state = null, CustomResourceOptions? options = null)
        {
            return new MetricStream(name, id, state, options);
        }
    }

    public sealed class MetricStreamArgs : Pulumi.ResourceArgs
    {
        [Input("excludeFilters")]
        private InputList<Inputs.MetricStreamExcludeFilterArgs>? _excludeFilters;

        /// <summary>
        /// List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces that you specify here. Conflicts with `include_filter`.
        /// </summary>
        public InputList<Inputs.MetricStreamExcludeFilterArgs> ExcludeFilters
        {
            get => _excludeFilters ?? (_excludeFilters = new InputList<Inputs.MetricStreamExcludeFilterArgs>());
            set => _excludeFilters = value;
        }

        /// <summary>
        /// ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
        /// </summary>
        [Input("firehoseArn", required: true)]
        public Input<string> FirehoseArn { get; set; } = null!;

        [Input("includeFilters")]
        private InputList<Inputs.MetricStreamIncludeFilterArgs>? _includeFilters;

        /// <summary>
        /// List of inclusive metric filters. If you specify this parameter, the stream sends only the metrics from the metric namespaces that you specify here. Conflicts with `exclude_filter`.
        /// </summary>
        public InputList<Inputs.MetricStreamIncludeFilterArgs> IncludeFilters
        {
            get => _includeFilters ?? (_includeFilters = new InputList<Inputs.MetricStreamIncludeFilterArgs>());
            set => _includeFilters = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// Output format for the stream. Possible values are `json` and `opentelemetry0.7`. For more information about output formats, see [Metric streams output formats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html).
        /// </summary>
        [Input("outputFormat", required: true)]
        public Input<string> OutputFormat { get; set; } = null!;

        /// <summary>
        /// ARN of the IAM role that this metric stream will use to access Amazon Kinesis Firehose resources. For more information about role permissions, see [Trust between CloudWatch and Kinesis Data Firehose](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html).
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public MetricStreamArgs()
        {
        }
    }

    public sealed class MetricStreamState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the metric stream.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the metric stream was created.
        /// </summary>
        [Input("creationDate")]
        public Input<string>? CreationDate { get; set; }

        [Input("excludeFilters")]
        private InputList<Inputs.MetricStreamExcludeFilterGetArgs>? _excludeFilters;

        /// <summary>
        /// List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces that you specify here. Conflicts with `include_filter`.
        /// </summary>
        public InputList<Inputs.MetricStreamExcludeFilterGetArgs> ExcludeFilters
        {
            get => _excludeFilters ?? (_excludeFilters = new InputList<Inputs.MetricStreamExcludeFilterGetArgs>());
            set => _excludeFilters = value;
        }

        /// <summary>
        /// ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
        /// </summary>
        [Input("firehoseArn")]
        public Input<string>? FirehoseArn { get; set; }

        [Input("includeFilters")]
        private InputList<Inputs.MetricStreamIncludeFilterGetArgs>? _includeFilters;

        /// <summary>
        /// List of inclusive metric filters. If you specify this parameter, the stream sends only the metrics from the metric namespaces that you specify here. Conflicts with `exclude_filter`.
        /// </summary>
        public InputList<Inputs.MetricStreamIncludeFilterGetArgs> IncludeFilters
        {
            get => _includeFilters ?? (_includeFilters = new InputList<Inputs.MetricStreamIncludeFilterGetArgs>());
            set => _includeFilters = value;
        }

        /// <summary>
        /// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the metric stream was last updated.
        /// </summary>
        [Input("lastUpdateDate")]
        public Input<string>? LastUpdateDate { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// Output format for the stream. Possible values are `json` and `opentelemetry0.7`. For more information about output formats, see [Metric streams output formats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html).
        /// </summary>
        [Input("outputFormat")]
        public Input<string>? OutputFormat { get; set; }

        /// <summary>
        /// ARN of the IAM role that this metric stream will use to access Amazon Kinesis Firehose resources. For more information about role permissions, see [Trust between CloudWatch and Kinesis Data Firehose](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html).
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// State of the metric stream. Possible values are `running` and `stopped`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public MetricStreamState()
        {
        }
    }
}
