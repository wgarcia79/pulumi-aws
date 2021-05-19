// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppRunner.Inputs
{

    public sealed class ServiceSourceConfigurationCodeRepositorySourceCodeVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of version identifier. For a git-based repository, branches represent versions. Valid values: `BRANCH`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// A source code version. For a git-based repository, a branch name maps to a specific version. App Runner uses the most recent commit to the branch.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ServiceSourceConfigurationCodeRepositorySourceCodeVersionArgs()
        {
        }
    }
}
