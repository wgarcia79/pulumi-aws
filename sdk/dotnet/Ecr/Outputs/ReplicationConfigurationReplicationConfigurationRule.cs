// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecr.Outputs
{

    [OutputType]
    public sealed class ReplicationConfigurationReplicationConfigurationRule
    {
        /// <summary>
        /// the details of a replication destination. See Destination.
        /// </summary>
        public readonly ImmutableArray<Outputs.ReplicationConfigurationReplicationConfigurationRuleDestination> Destinations;

        [OutputConstructor]
        private ReplicationConfigurationReplicationConfigurationRule(ImmutableArray<Outputs.ReplicationConfigurationReplicationConfigurationRuleDestination> destinations)
        {
            Destinations = destinations;
        }
    }
}