// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppStream.Outputs
{

    [OutputType]
    public sealed class FleetComputeCapacity
    {
        /// <summary>
        /// Number of currently available instances that can be used to stream sessions.
        /// </summary>
        public readonly int? Available;
        /// <summary>
        /// Desired number of streaming instances.
        /// </summary>
        public readonly int DesiredInstances;
        /// <summary>
        /// Number of instances in use for streaming.
        /// </summary>
        public readonly int? InUse;
        /// <summary>
        /// Total number of simultaneous streaming instances that are running.
        /// </summary>
        public readonly int? Running;

        [OutputConstructor]
        private FleetComputeCapacity(
            int? available,

            int desiredInstances,

            int? inUse,

            int? running)
        {
            Available = available;
            DesiredInstances = desiredInstances;
            InUse = inUse;
            Running = running;
        }
    }
}
