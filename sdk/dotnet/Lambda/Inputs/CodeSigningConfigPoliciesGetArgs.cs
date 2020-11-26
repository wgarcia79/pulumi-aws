// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Lambda.Inputs
{

    public sealed class CodeSigningConfigPoliciesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Code signing configuration policy for deployment validation failure. If you set the policy to Enforce, Lambda blocks the deployment request if code-signing validation checks fail. If you set the policy to Warn, Lambda allows the deployment and creates a CloudWatch log. Valid values: `Warn`, `Enforce`. Default value: `Warn`.
        /// </summary>
        [Input("untrustedArtifactOnDeployment", required: true)]
        public Input<string> UntrustedArtifactOnDeployment { get; set; } = null!;

        public CodeSigningConfigPoliciesGetArgs()
        {
        }
    }
}
