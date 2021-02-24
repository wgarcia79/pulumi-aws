// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iam.Inputs
{

    public sealed class RoleInlinePolicyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the role policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Policy document as a JSON formatted string.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        public RoleInlinePolicyGetArgs()
        {
        }
    }
}
