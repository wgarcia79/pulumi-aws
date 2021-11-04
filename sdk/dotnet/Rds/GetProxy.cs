// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Aws.Rds
{
    public static class GetProxy
    {
        /// <summary>
        /// Use this data source to get information about a DB Proxy.
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
        ///         var proxy = Output.Create(Aws.Rds.GetProxy.InvokeAsync(new Aws.Rds.GetProxyArgs
        ///         {
        ///             Name = "my-test-db-proxy",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProxyResult> InvokeAsync(GetProxyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProxyResult>("aws:rds/getProxy:getProxy", args ?? new GetProxyArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to get information about a DB Proxy.
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
        ///         var proxy = Output.Create(Aws.Rds.GetProxy.InvokeAsync(new Aws.Rds.GetProxyArgs
        ///         {
        ///             Name = "my-test-db-proxy",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProxyResult> Invoke(GetProxyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProxyResult>("aws:rds/getProxy:getProxy", args ?? new GetProxyInvokeArgs(), options.WithVersion());
    }


    public sealed class GetProxyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the DB proxy.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetProxyArgs()
        {
        }
    }

    public sealed class GetProxyInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the DB proxy.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetProxyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProxyResult
    {
        /// <summary>
        /// The ARN of the DB Proxy.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The configuration(s) with authorization mechanisms to connect to the associated instance or cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProxyAuthResult> Auths;
        /// <summary>
        /// Whether the proxy includes detailed information about SQL statements in its logs.
        /// </summary>
        public readonly bool DebugLogging;
        /// <summary>
        /// The endpoint that you can use to connect to the DB proxy.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// The kinds of databases that the proxy can connect to.
        /// </summary>
        public readonly string EngineFamily;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The number of seconds a connection to the proxy can have no activity before the proxy drops the client connection.
        /// </summary>
        public readonly int IdleClientTimeout;
        public readonly string Name;
        /// <summary>
        /// Indicates whether Transport Layer Security (TLS) encryption is required for connections to the proxy.
        /// </summary>
        public readonly bool RequireTls;
        /// <summary>
        /// The Amazon Resource Name (ARN) for the IAM role that the proxy uses to access Amazon Secrets Manager.
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// Provides the VPC ID of the DB proxy.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// Provides a list of VPC security groups that the proxy belongs to.
        /// </summary>
        public readonly ImmutableArray<string> VpcSecurityGroupIds;
        /// <summary>
        /// The EC2 subnet IDs for the proxy.
        /// </summary>
        public readonly ImmutableArray<string> VpcSubnetIds;

        [OutputConstructor]
        private GetProxyResult(
            string arn,

            ImmutableArray<Outputs.GetProxyAuthResult> auths,

            bool debugLogging,

            string endpoint,

            string engineFamily,

            string id,

            int idleClientTimeout,

            string name,

            bool requireTls,

            string roleArn,

            string vpcId,

            ImmutableArray<string> vpcSecurityGroupIds,

            ImmutableArray<string> vpcSubnetIds)
        {
            Arn = arn;
            Auths = auths;
            DebugLogging = debugLogging;
            Endpoint = endpoint;
            EngineFamily = engineFamily;
            Id = id;
            IdleClientTimeout = idleClientTimeout;
            Name = name;
            RequireTls = requireTls;
            RoleArn = roleArn;
            VpcId = vpcId;
            VpcSecurityGroupIds = vpcSecurityGroupIds;
            VpcSubnetIds = vpcSubnetIds;
        }
    }
}
