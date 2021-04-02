// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecs.Inputs
{

    public sealed class ServiceOrderedPlacementStrategyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// For the `spread` placement strategy, valid values are `instanceId` (or `host`,
        /// which has the same effect), or any platform or custom attribute that is applied to a container instance.
        /// For the `binpack` type, valid values are `memory` and `cpu`. For the `random` type, this attribute is not
        /// needed. For more information, see [Placement Strategy](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PlacementStrategy.html).
        /// </summary>
        [Input("field")]
        public Input<string>? Field { get; set; }

        /// <summary>
        /// Type of placement strategy. Must be one of: `binpack`, `random`, or `spread`
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ServiceOrderedPlacementStrategyArgs()
        {
        }
    }
}
