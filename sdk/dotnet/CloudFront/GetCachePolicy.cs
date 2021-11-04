// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Aws.CloudFront
{
    public static class GetCachePolicy
    {
        /// <summary>
        /// Use this data source to retrieve information about a CloudFront cache policy.
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
        ///         var example = Output.Create(Aws.CloudFront.GetCachePolicy.InvokeAsync(new Aws.CloudFront.GetCachePolicyArgs
        ///         {
        ///             Name = "example-policy",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCachePolicyResult> InvokeAsync(GetCachePolicyArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCachePolicyResult>("aws:cloudfront/getCachePolicy:getCachePolicy", args ?? new GetCachePolicyArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to retrieve information about a CloudFront cache policy.
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
        ///         var example = Output.Create(Aws.CloudFront.GetCachePolicy.InvokeAsync(new Aws.CloudFront.GetCachePolicyArgs
        ///         {
        ///             Name = "example-policy",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCachePolicyResult> Invoke(GetCachePolicyInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCachePolicyResult>("aws:cloudfront/getCachePolicy:getCachePolicy", args ?? new GetCachePolicyInvokeArgs(), options.WithVersion());
    }


    public sealed class GetCachePolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier for the cache policy.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// A unique name to identify the cache policy.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetCachePolicyArgs()
        {
        }
    }

    public sealed class GetCachePolicyInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier for the cache policy.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// A unique name to identify the cache policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetCachePolicyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCachePolicyResult
    {
        /// <summary>
        /// A comment to describe the cache policy.
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        /// </summary>
        public readonly int DefaultTtl;
        /// <summary>
        /// The current version of the cache policy.
        /// </summary>
        public readonly string Etag;
        public readonly string? Id;
        /// <summary>
        /// The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        /// </summary>
        public readonly int MaxTtl;
        /// <summary>
        /// The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
        /// </summary>
        public readonly int MinTtl;
        public readonly string? Name;
        /// <summary>
        /// The HTTP headers, cookies, and URL query strings to include in the cache key. See Parameters In Cache Key And Forwarded To Origin for more information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCachePolicyParametersInCacheKeyAndForwardedToOriginResult> ParametersInCacheKeyAndForwardedToOrigins;

        [OutputConstructor]
        private GetCachePolicyResult(
            string comment,

            int defaultTtl,

            string etag,

            string? id,

            int maxTtl,

            int minTtl,

            string? name,

            ImmutableArray<Outputs.GetCachePolicyParametersInCacheKeyAndForwardedToOriginResult> parametersInCacheKeyAndForwardedToOrigins)
        {
            Comment = comment;
            DefaultTtl = defaultTtl;
            Etag = etag;
            Id = id;
            MaxTtl = maxTtl;
            MinTtl = minTtl;
            Name = name;
            ParametersInCacheKeyAndForwardedToOrigins = parametersInCacheKeyAndForwardedToOrigins;
        }
    }
}
