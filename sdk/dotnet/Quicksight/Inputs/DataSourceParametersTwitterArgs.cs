// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Quicksight.Inputs
{

    public sealed class DataSourceParametersTwitterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of rows to query.
        /// </summary>
        [Input("maxRows", required: true)]
        public Input<int> MaxRows { get; set; } = null!;

        /// <summary>
        /// The Twitter query to retrieve the data.
        /// </summary>
        [Input("query", required: true)]
        public Input<string> Query { get; set; } = null!;

        public DataSourceParametersTwitterArgs()
        {
        }
    }
}
