// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53
{
    /// <summary>
    /// Provides a Route 53 Resolver query logging configuration resource.
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
    ///         var example = new Aws.Route53.ResolverQueryLogConfig("example", new Aws.Route53.ResolverQueryLogConfigArgs
    ///         {
    ///             DestinationArn = aws_s3_bucket.Example.Arn,
    ///             Tags = 
    ///             {
    ///                 { "Environment", "Prod" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    ///  Route 53 Resolver query logging configurations can be imported using the Route 53 Resolver query logging configuration ID, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:route53/resolverQueryLogConfig:ResolverQueryLogConfig example rqlc-92edc3b1838248bf
    /// ```
    /// </summary>
    [AwsResourceType("aws:route53/resolverQueryLogConfig:ResolverQueryLogConfig")]
    public partial class ResolverQueryLogConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN (Amazon Resource Name) of the Route 53 Resolver query logging configuration.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The ARN of the resource that you want Route 53 Resolver to send query logs.
        /// You can send query logs to an S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
        /// </summary>
        [Output("destinationArn")]
        public Output<string> DestinationArn { get; private set; } = null!;

        /// <summary>
        /// The name of the Route 53 Resolver query logging configuration.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The AWS account ID of the account that created the query logging configuration.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// An indication of whether the query logging configuration is shared with other AWS accounts, or was shared with the current account by another AWS account.
        /// Sharing is configured through AWS Resource Access Manager (AWS RAM).
        /// Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
        /// </summary>
        [Output("shareStatus")]
        public Output<string> ShareStatus { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a ResolverQueryLogConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResolverQueryLogConfig(string name, ResolverQueryLogConfigArgs args, CustomResourceOptions? options = null)
            : base("aws:route53/resolverQueryLogConfig:ResolverQueryLogConfig", name, args ?? new ResolverQueryLogConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResolverQueryLogConfig(string name, Input<string> id, ResolverQueryLogConfigState? state = null, CustomResourceOptions? options = null)
            : base("aws:route53/resolverQueryLogConfig:ResolverQueryLogConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResolverQueryLogConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResolverQueryLogConfig Get(string name, Input<string> id, ResolverQueryLogConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new ResolverQueryLogConfig(name, id, state, options);
        }
    }

    public sealed class ResolverQueryLogConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the resource that you want Route 53 Resolver to send query logs.
        /// You can send query logs to an S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
        /// </summary>
        [Input("destinationArn", required: true)]
        public Input<string> DestinationArn { get; set; } = null!;

        /// <summary>
        /// The name of the Route 53 Resolver query logging configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ResolverQueryLogConfigArgs()
        {
        }
    }

    public sealed class ResolverQueryLogConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN (Amazon Resource Name) of the Route 53 Resolver query logging configuration.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The ARN of the resource that you want Route 53 Resolver to send query logs.
        /// You can send query logs to an S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
        /// </summary>
        [Input("destinationArn")]
        public Input<string>? DestinationArn { get; set; }

        /// <summary>
        /// The name of the Route 53 Resolver query logging configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The AWS account ID of the account that created the query logging configuration.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        /// <summary>
        /// An indication of whether the query logging configuration is shared with other AWS accounts, or was shared with the current account by another AWS account.
        /// Sharing is configured through AWS Resource Access Manager (AWS RAM).
        /// Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
        /// </summary>
        [Input("shareStatus")]
        public Input<string>? ShareStatus { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public ResolverQueryLogConfigState()
        {
        }
    }
}
