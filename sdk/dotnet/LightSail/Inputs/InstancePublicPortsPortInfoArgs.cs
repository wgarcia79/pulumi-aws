// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LightSail.Inputs
{

    public sealed class InstancePublicPortsPortInfoArgs : Pulumi.ResourceArgs
    {
        [Input("cidrs")]
        private InputList<string>? _cidrs;

        /// <summary>
        /// Set of CIDR blocks.
        /// </summary>
        public InputList<string> Cidrs
        {
            get => _cidrs ?? (_cidrs = new InputList<string>());
            set => _cidrs = value;
        }

        /// <summary>
        /// First port in a range of open ports on an instance.
        /// </summary>
        [Input("fromPort", required: true)]
        public Input<int> FromPort { get; set; } = null!;

        /// <summary>
        /// IP protocol name. Valid values are `tcp`, `all`, `udp`, and `icmp`.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// Last port in a range of open ports on an instance.
        /// </summary>
        [Input("toPort", required: true)]
        public Input<int> ToPort { get; set; } = null!;

        public InstancePublicPortsPortInfoArgs()
        {
        }
    }
}
