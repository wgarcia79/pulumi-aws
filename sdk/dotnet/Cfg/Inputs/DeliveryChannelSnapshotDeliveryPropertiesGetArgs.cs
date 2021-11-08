// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cfg.Inputs
{

    public sealed class DeliveryChannelSnapshotDeliveryPropertiesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// - The frequency with which AWS Config recurringly delivers configuration snapshotsE.g., `One_Hour` or `Three_Hours`. Valid values are listed [here](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigSnapshotDeliveryProperties.html#API_ConfigSnapshotDeliveryProperties_Contents).
        /// </summary>
        [Input("deliveryFrequency")]
        public Input<string>? DeliveryFrequency { get; set; }

        public DeliveryChannelSnapshotDeliveryPropertiesGetArgs()
        {
        }
    }
}
