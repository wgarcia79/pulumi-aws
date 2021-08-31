// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling.Inputs
{

    public sealed class TagTagGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Tag name.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Whether to propagate the tags to instances launched by the ASG.
        /// </summary>
        [Input("propagateAtLaunch", required: true)]
        public Input<bool> PropagateAtLaunch { get; set; } = null!;

        /// <summary>
        /// Tag value.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public TagTagGetArgs()
        {
        }
    }
}
