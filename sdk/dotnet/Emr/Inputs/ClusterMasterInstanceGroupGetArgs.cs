// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Emr.Inputs
{

    public sealed class ClusterMasterInstanceGroupGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bid price for each EC2 Spot instance type as defined by `instance_type`. Expressed in USD. If neither `bid_price` nor `bid_price_as_percentage_of_on_demand_price` is provided, `bid_price_as_percentage_of_on_demand_price` defaults to 100%.
        /// </summary>
        [Input("bidPrice")]
        public Input<string>? BidPrice { get; set; }

        [Input("ebsConfigs")]
        private InputList<Inputs.ClusterMasterInstanceGroupEbsConfigGetArgs>? _ebsConfigs;

        /// <summary>
        /// Configuration block(s) for EBS volumes attached to each instance in the instance group. Detailed below.
        /// </summary>
        public InputList<Inputs.ClusterMasterInstanceGroupEbsConfigGetArgs> EbsConfigs
        {
            get => _ebsConfigs ?? (_ebsConfigs = new InputList<Inputs.ClusterMasterInstanceGroupEbsConfigGetArgs>());
            set => _ebsConfigs = value;
        }

        /// <summary>
        /// The ID of the EMR Cluster
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Target number of instances for the instance group. Must be 1 or 3. Defaults to 1. Launching with multiple master nodes is only supported in EMR version 5.23.0+, and requires this resource's `core_instance_group` to be configured. Public (Internet accessible) instances must be created in VPC subnets that have `map public IP on launch` enabled. Termination protection is automatically enabled when launched with multiple master nodes and this provider must have the `termination_protection = false` configuration applied before destroying this resource.
        /// </summary>
        [Input("instanceCount")]
        public Input<int>? InstanceCount { get; set; }

        /// <summary>
        /// An EC2 instance type, such as m4.xlarge.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// Friendly name given to the instance fleet.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ClusterMasterInstanceGroupGetArgs()
        {
        }
    }
}
