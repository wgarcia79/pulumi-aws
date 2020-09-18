// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Emr.Inputs
{

    public sealed class ClusterMasterInstanceFleetInstanceTypeConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bid price for each EC2 Spot instance type as defined by `instance_type`. Expressed in USD. If neither `bid_price` nor `bid_price_as_percentage_of_on_demand_price` is provided, `bid_price_as_percentage_of_on_demand_price` defaults to 100%.
        /// </summary>
        [Input("bidPrice")]
        public Input<string>? BidPrice { get; set; }

        /// <summary>
        /// The bid price, as a percentage of On-Demand price, for each EC2 Spot instance as defined by `instance_type`. Expressed as a number (for example, 20 specifies 20%). If neither `bid_price` nor `bid_price_as_percentage_of_on_demand_price` is provided, `bid_price_as_percentage_of_on_demand_price` defaults to 100%.
        /// </summary>
        [Input("bidPriceAsPercentageOfOnDemandPrice")]
        public Input<double>? BidPriceAsPercentageOfOnDemandPrice { get; set; }

        [Input("configurations")]
        private InputList<Inputs.ClusterMasterInstanceFleetInstanceTypeConfigConfigurationGetArgs>? _configurations;

        /// <summary>
        /// A configuration classification that applies when provisioning cluster instances, which can include configurations for applications and software that run on the cluster. List of `configuration` blocks.
        /// </summary>
        public InputList<Inputs.ClusterMasterInstanceFleetInstanceTypeConfigConfigurationGetArgs> Configurations
        {
            get => _configurations ?? (_configurations = new InputList<Inputs.ClusterMasterInstanceFleetInstanceTypeConfigConfigurationGetArgs>());
            set => _configurations = value;
        }

        [Input("ebsConfigs")]
        private InputList<Inputs.ClusterMasterInstanceFleetInstanceTypeConfigEbsConfigGetArgs>? _ebsConfigs;

        /// <summary>
        /// Configuration block(s) for EBS volumes attached to each instance in the instance group. Detailed below.
        /// </summary>
        public InputList<Inputs.ClusterMasterInstanceFleetInstanceTypeConfigEbsConfigGetArgs> EbsConfigs
        {
            get => _ebsConfigs ?? (_ebsConfigs = new InputList<Inputs.ClusterMasterInstanceFleetInstanceTypeConfigEbsConfigGetArgs>());
            set => _ebsConfigs = value;
        }

        /// <summary>
        /// An EC2 instance type, such as m4.xlarge.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// The number of units that a provisioned instance of this type provides toward fulfilling the target capacities defined in `aws.emr.InstanceFleet`.
        /// </summary>
        [Input("weightedCapacity")]
        public Input<int>? WeightedCapacity { get; set; }

        public ClusterMasterInstanceFleetInstanceTypeConfigGetArgs()
        {
        }
    }
}
