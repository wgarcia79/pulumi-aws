// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Aws.Route53
{
    public static class GetResolverEndpoint
    {
        /// <summary>
        /// `aws.route53.ResolverEndpoint` provides details about a specific Route53 Resolver Endpoint.
        /// 
        /// This data source allows to find a list of IPaddresses associated with a specific Route53 Resolver Endpoint.
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
        ///         var example = Output.Create(Aws.Route53.GetResolverEndpoint.InvokeAsync(new Aws.Route53.GetResolverEndpointArgs
        ///         {
        ///             ResolverEndpointId = "rslvr-in-1abc2345ef678g91h",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.Route53.GetResolverEndpoint.InvokeAsync(new Aws.Route53.GetResolverEndpointArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Aws.Route53.Inputs.GetResolverEndpointFilterArgs
        ///                 {
        ///                     Name = "NAME",
        ///                     Values = 
        ///                     {
        ///                         "MyResolverExampleName",
        ///                     },
        ///                 },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetResolverEndpointResult> InvokeAsync(GetResolverEndpointArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResolverEndpointResult>("aws:route53/getResolverEndpoint:getResolverEndpoint", args ?? new GetResolverEndpointArgs(), options.WithVersion());

        /// <summary>
        /// `aws.route53.ResolverEndpoint` provides details about a specific Route53 Resolver Endpoint.
        /// 
        /// This data source allows to find a list of IPaddresses associated with a specific Route53 Resolver Endpoint.
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
        ///         var example = Output.Create(Aws.Route53.GetResolverEndpoint.InvokeAsync(new Aws.Route53.GetResolverEndpointArgs
        ///         {
        ///             ResolverEndpointId = "rslvr-in-1abc2345ef678g91h",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.Route53.GetResolverEndpoint.InvokeAsync(new Aws.Route53.GetResolverEndpointArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Aws.Route53.Inputs.GetResolverEndpointFilterArgs
        ///                 {
        ///                     Name = "NAME",
        ///                     Values = 
        ///                     {
        ///                         "MyResolverExampleName",
        ///                     },
        ///                 },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetResolverEndpointResult> Invoke(GetResolverEndpointInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetResolverEndpointResult>("aws:route53/getResolverEndpoint:getResolverEndpoint", args ?? new GetResolverEndpointInvokeArgs(), options.WithVersion());
    }


    public sealed class GetResolverEndpointArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetResolverEndpointFilterArgs>? _filters;

        /// <summary>
        /// One or more name/value pairs to use as filters. There are
        /// several valid keys, for a full reference, check out
        /// [Route53resolver Filter value in the AWS API reference][1].
        /// </summary>
        public List<Inputs.GetResolverEndpointFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetResolverEndpointFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The ID of the Route53 Resolver Endpoint.
        /// </summary>
        [Input("resolverEndpointId")]
        public string? ResolverEndpointId { get; set; }

        public GetResolverEndpointArgs()
        {
        }
    }

    public sealed class GetResolverEndpointInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetResolverEndpointFilterInputArgs>? _filters;

        /// <summary>
        /// One or more name/value pairs to use as filters. There are
        /// several valid keys, for a full reference, check out
        /// [Route53resolver Filter value in the AWS API reference][1].
        /// </summary>
        public InputList<Inputs.GetResolverEndpointFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetResolverEndpointFilterInputArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The ID of the Route53 Resolver Endpoint.
        /// </summary>
        [Input("resolverEndpointId")]
        public Input<string>? ResolverEndpointId { get; set; }

        public GetResolverEndpointInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResolverEndpointResult
    {
        public readonly string Arn;
        public readonly string Direction;
        public readonly ImmutableArray<Outputs.GetResolverEndpointFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> IpAddresses;
        public readonly string Name;
        public readonly string? ResolverEndpointId;
        public readonly string Status;
        public readonly string VpcId;

        [OutputConstructor]
        private GetResolverEndpointResult(
            string arn,

            string direction,

            ImmutableArray<Outputs.GetResolverEndpointFilterResult> filters,

            string id,

            ImmutableArray<string> ipAddresses,

            string name,

            string? resolverEndpointId,

            string status,

            string vpcId)
        {
            Arn = arn;
            Direction = direction;
            Filters = filters;
            Id = id;
            IpAddresses = ipAddresses;
            Name = name;
            ResolverEndpointId = resolverEndpointId;
            Status = status;
            VpcId = vpcId;
        }
    }
}
