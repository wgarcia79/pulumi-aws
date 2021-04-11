// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Mwaa.Inputs
{

    public sealed class EnvironmentLoggingConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) Log configuration options for processing DAGs. See Module logging configuration for more information. Disabled by default.
        /// </summary>
        [Input("dagProcessingLogs")]
        public Input<Inputs.EnvironmentLoggingConfigurationDagProcessingLogsGetArgs>? DagProcessingLogs { get; set; }

        /// <summary>
        /// Log configuration options for the schedulers. See Module logging configuration for more information. Disabled by default.
        /// </summary>
        [Input("schedulerLogs")]
        public Input<Inputs.EnvironmentLoggingConfigurationSchedulerLogsGetArgs>? SchedulerLogs { get; set; }

        /// <summary>
        /// Log configuration options for DAG tasks. See Module logging configuration for more information. Enabled by default with `INFO` log level.
        /// </summary>
        [Input("taskLogs")]
        public Input<Inputs.EnvironmentLoggingConfigurationTaskLogsGetArgs>? TaskLogs { get; set; }

        /// <summary>
        /// Log configuration options for the webservers. See Module logging configuration for more information. Disabled by default.
        /// </summary>
        [Input("webserverLogs")]
        public Input<Inputs.EnvironmentLoggingConfigurationWebserverLogsGetArgs>? WebserverLogs { get; set; }

        /// <summary>
        /// Log configuration options for the workers. See Module logging configuration for more information. Disabled by default.
        /// </summary>
        [Input("workerLogs")]
        public Input<Inputs.EnvironmentLoggingConfigurationWorkerLogsGetArgs>? WorkerLogs { get; set; }

        public EnvironmentLoggingConfigurationGetArgs()
        {
        }
    }
}
