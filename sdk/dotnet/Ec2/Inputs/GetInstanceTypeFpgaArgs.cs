// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Inputs
{

    public sealed class GetInstanceTypeFpgaInputArgs : Pulumi.ResourceArgs
    {
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        [Input("manufacturer", required: true)]
        public Input<string> Manufacturer { get; set; } = null!;

        /// <summary>
        /// Size of the instance memory, in MiB.
        /// </summary>
        [Input("memorySize", required: true)]
        public Input<int> MemorySize { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetInstanceTypeFpgaInputArgs()
        {
        }
    }
}
