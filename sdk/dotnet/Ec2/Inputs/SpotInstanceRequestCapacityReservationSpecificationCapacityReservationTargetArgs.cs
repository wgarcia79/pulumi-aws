// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Inputs
{

    public sealed class SpotInstanceRequestCapacityReservationSpecificationCapacityReservationTargetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Capacity Reservation in which to run the instance.
        /// </summary>
        [Input("capacityReservationId")]
        public Input<string>? CapacityReservationId { get; set; }

        public SpotInstanceRequestCapacityReservationSpecificationCapacityReservationTargetArgs()
        {
        }
    }
}
