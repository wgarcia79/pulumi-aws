// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue.Inputs
{

    public sealed class GetScriptDagEdgeInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the node at which the edge starts.
        /// </summary>
        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        /// <summary>
        /// The ID of the node at which the edge ends.
        /// </summary>
        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        /// <summary>
        /// The target of the edge.
        /// </summary>
        [Input("targetParameter")]
        public Input<string>? TargetParameter { get; set; }

        public GetScriptDagEdgeInputArgs()
        {
        }
    }
}
