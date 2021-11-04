// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Aws.Kinesis
{
    public static class GetFirehoseDeliveryStream
    {
        /// <summary>
        /// Use this data source to get information about a Kinesis Firehose Delivery Stream for use in other resources.
        /// 
        /// For more details, see the [Amazon Kinesis Firehose Documentation](https://aws.amazon.com/documentation/firehose/).
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
        ///         var stream = Output.Create(Aws.Kinesis.GetFirehoseDeliveryStream.InvokeAsync(new Aws.Kinesis.GetFirehoseDeliveryStreamArgs
        ///         {
        ///             Name = "stream-name",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFirehoseDeliveryStreamResult> InvokeAsync(GetFirehoseDeliveryStreamArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFirehoseDeliveryStreamResult>("aws:kinesis/getFirehoseDeliveryStream:getFirehoseDeliveryStream", args ?? new GetFirehoseDeliveryStreamArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to get information about a Kinesis Firehose Delivery Stream for use in other resources.
        /// 
        /// For more details, see the [Amazon Kinesis Firehose Documentation](https://aws.amazon.com/documentation/firehose/).
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
        ///         var stream = Output.Create(Aws.Kinesis.GetFirehoseDeliveryStream.InvokeAsync(new Aws.Kinesis.GetFirehoseDeliveryStreamArgs
        ///         {
        ///             Name = "stream-name",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFirehoseDeliveryStreamResult> Invoke(GetFirehoseDeliveryStreamInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFirehoseDeliveryStreamResult>("aws:kinesis/getFirehoseDeliveryStream:getFirehoseDeliveryStream", args ?? new GetFirehoseDeliveryStreamInvokeArgs(), options.WithVersion());
    }


    public sealed class GetFirehoseDeliveryStreamArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Kinesis Stream.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetFirehoseDeliveryStreamArgs()
        {
        }
    }

    public sealed class GetFirehoseDeliveryStreamInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Kinesis Stream.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetFirehoseDeliveryStreamInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFirehoseDeliveryStreamResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Kinesis Stream (same as id).
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetFirehoseDeliveryStreamResult(
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
