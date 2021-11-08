// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DynamoDB.Outputs
{

    [OutputType]
    public sealed class GlobalTableReplica
    {
        /// <summary>
        /// AWS region name of replica DynamoDB TableE.g., `us-east-1`
        /// </summary>
        public readonly string RegionName;

        [OutputConstructor]
        private GlobalTableReplica(string regionName)
        {
            RegionName = regionName;
        }
    }
}
