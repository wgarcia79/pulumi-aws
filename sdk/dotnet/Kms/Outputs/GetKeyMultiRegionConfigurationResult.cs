// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kms.Outputs
{

    [OutputType]
    public sealed class GetKeyMultiRegionConfigurationResult
    {
        public readonly string MultiRegionKeyType;
        public readonly ImmutableArray<Outputs.GetKeyMultiRegionConfigurationPrimaryKeyResult> PrimaryKeys;
        public readonly ImmutableArray<Outputs.GetKeyMultiRegionConfigurationReplicaKeyResult> ReplicaKeys;

        [OutputConstructor]
        private GetKeyMultiRegionConfigurationResult(
            string multiRegionKeyType,

            ImmutableArray<Outputs.GetKeyMultiRegionConfigurationPrimaryKeyResult> primaryKeys,

            ImmutableArray<Outputs.GetKeyMultiRegionConfigurationReplicaKeyResult> replicaKeys)
        {
            MultiRegionKeyType = multiRegionKeyType;
            PrimaryKeys = primaryKeys;
            ReplicaKeys = replicaKeys;
        }
    }
}
