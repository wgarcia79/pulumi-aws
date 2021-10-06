// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.RedShift.Outputs
{

    [OutputType]
    public sealed class ScheduledActionTargetActionResizeCluster
    {
        /// <summary>
        /// A boolean value indicating whether the resize operation is using the classic resize process. Default: `false`.
        /// </summary>
        public readonly bool? Classic;
        /// <summary>
        /// The identifier of the cluster to be resumed.
        /// </summary>
        public readonly string ClusterIdentifier;
        /// <summary>
        /// The new cluster type for the specified cluster.
        /// </summary>
        public readonly string? ClusterType;
        /// <summary>
        /// The new node type for the nodes you are adding.
        /// </summary>
        public readonly string? NodeType;
        /// <summary>
        /// The new number of nodes for the cluster.
        /// </summary>
        public readonly int? NumberOfNodes;

        [OutputConstructor]
        private ScheduledActionTargetActionResizeCluster(
            bool? classic,

            string clusterIdentifier,

            string? clusterType,

            string? nodeType,

            int? numberOfNodes)
        {
            Classic = classic;
            ClusterIdentifier = clusterIdentifier;
            ClusterType = clusterType;
            NodeType = nodeType;
            NumberOfNodes = numberOfNodes;
        }
    }
}
