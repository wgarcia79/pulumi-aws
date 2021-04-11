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
    public sealed class ReplicationConfigurationReplicationConfigurationRuleDestination
    {
        /// <summary>
        /// A Region to replicate to.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The account ID of the destination registry to replicate to.
        /// </summary>
        public readonly string RegistryId;

        [OutputConstructor]
        private ReplicationConfigurationReplicationConfigurationRuleDestination(
            string region,

            string registryId)
        {
            Region = region;
            RegistryId = registryId;
        }
    }
}
