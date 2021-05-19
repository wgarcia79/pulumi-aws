// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppRunner.Outputs
{

    [OutputType]
    public sealed class ServiceInstanceConfiguration
    {
        /// <summary>
        /// The number of CPU units reserved for each instance of your App Runner service represented as a String. Defaults to `1024`. Valid values: `1024|2048|(1|2) vCPU`.
        /// </summary>
        public readonly string? Cpu;
        /// <summary>
        /// The Amazon Resource Name (ARN) of an IAM role that provides permissions to your App Runner service. These are permissions that your code needs when it calls any AWS APIs.
        /// </summary>
        public readonly string InstanceRoleArn;
        /// <summary>
        /// The amount of memory, in MB or GB, reserved for each instance of your App Runner service. Defaults to `2048`. Valid values: `2048|3072|4096|(2|3|4) GB`.
        /// </summary>
        public readonly string? Memory;

        [OutputConstructor]
        private ServiceInstanceConfiguration(
            string? cpu,

            string instanceRoleArn,

            string? memory)
        {
            Cpu = cpu;
            InstanceRoleArn = instanceRoleArn;
            Memory = memory;
        }
    }
}
