// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Aws.Msk
{
    public static class GetCluster
    {
        /// <summary>
        /// Get information on an Amazon MSK Cluster.
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
        ///         var example = Output.Create(Aws.Msk.GetCluster.InvokeAsync(new Aws.Msk.GetClusterArgs
        ///         {
        ///             ClusterName = "example",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("aws:msk/getCluster:getCluster", args ?? new GetClusterArgs(), options.WithVersion());

        /// <summary>
        /// Get information on an Amazon MSK Cluster.
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
        ///         var example = Output.Create(Aws.Msk.GetCluster.InvokeAsync(new Aws.Msk.GetClusterArgs
        ///         {
        ///             ClusterName = "example",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetClusterResult>("aws:msk/getCluster:getCluster", args ?? new GetClusterInvokeArgs(), options.WithVersion());
    }


    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Map of key-value pairs assigned to the cluster.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetClusterArgs()
        {
        }
    }

    public sealed class GetClusterInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of key-value pairs assigned to the cluster.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetClusterInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the MSK cluster.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Comma separated list of one or more hostname:port pairs of kafka brokers suitable to bootstrap connectivity to the kafka cluster. Contains a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `PLAINTEXT` or `TLS_PLAINTEXT`. The resource sorts values alphabetically. AWS may not always return all endpoints so this value is not guaranteed to be stable across applies.
        /// </summary>
        public readonly string BootstrapBrokers;
        /// <summary>
        /// One or more DNS names (or IP addresses) and SASL IAM port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.iam` is set to `true`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
        /// </summary>
        public readonly string BootstrapBrokersSaslIam;
        /// <summary>
        /// One or more DNS names (or IP addresses) and SASL SCRAM port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.scram` is set to `true`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
        /// </summary>
        public readonly string BootstrapBrokersSaslScram;
        /// <summary>
        /// One or more DNS names (or IP addresses) and TLS port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
        /// </summary>
        public readonly string BootstrapBrokersTls;
        public readonly string ClusterName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Apache Kafka version.
        /// </summary>
        public readonly string KafkaVersion;
        /// <summary>
        /// Number of broker nodes in the cluster.
        /// </summary>
        public readonly int NumberOfBrokerNodes;
        /// <summary>
        /// Map of key-value pairs assigned to the cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster. The returned values are sorted alphbetically. The AWS API may not return all endpoints, so this value is not guaranteed to be stable across applies.
        /// </summary>
        public readonly string ZookeeperConnectString;

        [OutputConstructor]
        private GetClusterResult(
            string arn,

            string bootstrapBrokers,

            string bootstrapBrokersSaslIam,

            string bootstrapBrokersSaslScram,

            string bootstrapBrokersTls,

            string clusterName,

            string id,

            string kafkaVersion,

            int numberOfBrokerNodes,

            ImmutableDictionary<string, string> tags,

            string zookeeperConnectString)
        {
            Arn = arn;
            BootstrapBrokers = bootstrapBrokers;
            BootstrapBrokersSaslIam = bootstrapBrokersSaslIam;
            BootstrapBrokersSaslScram = bootstrapBrokersSaslScram;
            BootstrapBrokersTls = bootstrapBrokersTls;
            ClusterName = clusterName;
            Id = id;
            KafkaVersion = kafkaVersion;
            NumberOfBrokerNodes = numberOfBrokerNodes;
            Tags = tags;
            ZookeeperConnectString = zookeeperConnectString;
        }
    }
}
