// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kinesis
{
    /// <summary>
    /// Provides a resource to manage a Kinesis Stream Consumer.
    /// 
    /// &gt; **Note:** You can register up to 20 consumers per stream. A given consumer can only be registered with one stream at a time.
    /// 
    /// For more details, see the [Amazon Kinesis Stream Consumer Documentation](https://docs.aws.amazon.com/streams/latest/dev/amazon-kinesis-consumers.html).
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
    ///         var exampleStream = new Aws.Kinesis.Stream("exampleStream", new Aws.Kinesis.StreamArgs
    ///         {
    ///             ShardCount = 1,
    ///         });
    ///         var exampleStreamConsumer = new Aws.Kinesis.StreamConsumer("exampleStreamConsumer", new Aws.Kinesis.StreamConsumerArgs
    ///         {
    ///             StreamArn = exampleStream.Arn,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Kinesis Stream Consumers can be imported using the Amazon Resource Name (ARN) e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:kinesis/streamConsumer:StreamConsumer example arn:aws:kinesis:us-west-2:123456789012:stream/example/consumer/example:1616044553
    /// ```
    /// 
    ///  [1]https://docs.aws.amazon.com/streams/latest/dev/amazon-kinesis-consumers.html
    /// </summary>
    [AwsResourceType("aws:kinesis/streamConsumer:StreamConsumer")]
    public partial class StreamConsumer : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the stream consumer.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Approximate timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of when the stream consumer was created.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// Name of the stream consumer.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the data stream the consumer is registered with.
        /// </summary>
        [Output("streamArn")]
        public Output<string> StreamArn { get; private set; } = null!;


        /// <summary>
        /// Create a StreamConsumer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StreamConsumer(string name, StreamConsumerArgs args, CustomResourceOptions? options = null)
            : base("aws:kinesis/streamConsumer:StreamConsumer", name, args ?? new StreamConsumerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StreamConsumer(string name, Input<string> id, StreamConsumerState? state = null, CustomResourceOptions? options = null)
            : base("aws:kinesis/streamConsumer:StreamConsumer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StreamConsumer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StreamConsumer Get(string name, Input<string> id, StreamConsumerState? state = null, CustomResourceOptions? options = null)
        {
            return new StreamConsumer(name, id, state, options);
        }
    }

    public sealed class StreamConsumerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the stream consumer.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the data stream the consumer is registered with.
        /// </summary>
        [Input("streamArn", required: true)]
        public Input<string> StreamArn { get; set; } = null!;

        public StreamConsumerArgs()
        {
        }
    }

    public sealed class StreamConsumerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the stream consumer.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Approximate timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of when the stream consumer was created.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// Name of the stream consumer.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the data stream the consumer is registered with.
        /// </summary>
        [Input("streamArn")]
        public Input<string>? StreamArn { get; set; }

        public StreamConsumerState()
        {
        }
    }
}
