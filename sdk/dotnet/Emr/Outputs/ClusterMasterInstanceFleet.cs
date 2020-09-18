// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Emr.Outputs
{

    [OutputType]
    public sealed class ClusterMasterInstanceFleet
    {
        /// <summary>
        /// The ID of the EMR Cluster
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Configuration block for instance fleet
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterMasterInstanceFleetInstanceTypeConfig> InstanceTypeConfigs;
        /// <summary>
        /// Configuration block for launch specification
        /// </summary>
        public readonly Outputs.ClusterMasterInstanceFleetLaunchSpecifications? LaunchSpecifications;
        /// <summary>
        /// Friendly name given to the instance fleet.
        /// </summary>
        public readonly string? Name;
        public readonly int? ProvisionedOnDemandCapacity;
        public readonly int? ProvisionedSpotCapacity;
        /// <summary>
        /// The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
        /// </summary>
        public readonly int? TargetOnDemandCapacity;
        /// <summary>
        /// The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
        /// </summary>
        public readonly int? TargetSpotCapacity;

        [OutputConstructor]
        private ClusterMasterInstanceFleet(
            string? id,

            ImmutableArray<Outputs.ClusterMasterInstanceFleetInstanceTypeConfig> instanceTypeConfigs,

            Outputs.ClusterMasterInstanceFleetLaunchSpecifications? launchSpecifications,

            string? name,

            int? provisionedOnDemandCapacity,

            int? provisionedSpotCapacity,

            int? targetOnDemandCapacity,

            int? targetSpotCapacity)
        {
            Id = id;
            InstanceTypeConfigs = instanceTypeConfigs;
            LaunchSpecifications = launchSpecifications;
            Name = name;
            ProvisionedOnDemandCapacity = provisionedOnDemandCapacity;
            ProvisionedSpotCapacity = provisionedSpotCapacity;
            TargetOnDemandCapacity = targetOnDemandCapacity;
            TargetSpotCapacity = targetSpotCapacity;
        }
    }
}
