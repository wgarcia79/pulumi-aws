// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Quicksight.Inputs
{

    public sealed class DataSourceParametersAthenaGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The work-group to which to connect.
        /// </summary>
        [Input("workGroup")]
        public Input<string>? WorkGroup { get; set; }

        public DataSourceParametersAthenaGetArgs()
        {
        }
    }
}
