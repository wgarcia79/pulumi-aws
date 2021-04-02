// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFormation.Inputs
{

    public sealed class StackSetAutoDeploymentGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not auto-deployment is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Whether or not to retain stacks when the account is removed.
        /// </summary>
        [Input("retainStacksOnAccountRemoval")]
        public Input<bool>? RetainStacksOnAccountRemoval { get; set; }

        public StackSetAutoDeploymentGetArgs()
        {
        }
    }
}
