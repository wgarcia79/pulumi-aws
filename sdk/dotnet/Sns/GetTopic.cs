// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Aws.Sns
{
    public static class GetTopic
    {
        /// <summary>
        /// Use this data source to get the ARN of a topic in AWS Simple Notification
        /// Service (SNS). By using this data source, you can reference SNS topics
        /// without having to hard code the ARNs as input.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.Sns.GetTopic.InvokeAsync(new Aws.Sns.GetTopicArgs
        ///         {
        ///             Name = "an_example_topic",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTopicResult> InvokeAsync(GetTopicArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTopicResult>("aws:sns/getTopic:getTopic", args ?? new GetTopicArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to get the ARN of a topic in AWS Simple Notification
        /// Service (SNS). By using this data source, you can reference SNS topics
        /// without having to hard code the ARNs as input.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.Sns.GetTopic.InvokeAsync(new Aws.Sns.GetTopicArgs
        ///         {
        ///             Name = "an_example_topic",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTopicResult> Invoke(GetTopicInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTopicResult>("aws:sns/getTopic:getTopic", args ?? new GetTopicInvokeArgs(), options.WithVersion());
    }


    public sealed class GetTopicArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The friendly name of the topic to match.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetTopicArgs()
        {
        }
    }

    public sealed class GetTopicInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The friendly name of the topic to match.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetTopicInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTopicResult
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the found topic, suitable for referencing in other resources that support SNS topics.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetTopicResult(
            string arn,

            string id,

            string name)
        {
            Arn = arn;
            Id = id;
            Name = name;
        }
    }
}
