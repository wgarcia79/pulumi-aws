// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.KinesisAnalyticsV2.Outputs
{

    [OutputType]
    public sealed class ApplicationApplicationConfigurationSqlApplicationConfigurationInput
    {
        public readonly ImmutableArray<string> InAppStreamNames;
        public readonly string? InputId;
        /// <summary>
        /// Describes the number of in-application streams to create.
        /// </summary>
        public readonly Outputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputParallelism? InputParallelism;
        /// <summary>
        /// The input processing configuration for the input.
        /// An input processor transforms records as they are received from the stream, before the application's SQL code executes.
        /// </summary>
        public readonly Outputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputProcessingConfiguration? InputProcessingConfiguration;
        /// <summary>
        /// Describes the format of the data in the streaming source, and how each data element maps to corresponding columns in the in-application stream that is being created.
        /// </summary>
        public readonly Outputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchema InputSchema;
        /// <summary>
        /// The point at which the application starts processing records from the streaming source.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfiguration> InputStartingPositionConfigurations;
        /// <summary>
        /// If the streaming source is a Kinesis Data Firehose delivery stream, identifies the delivery stream's ARN.
        /// </summary>
        public readonly Outputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisFirehoseInput? KinesisFirehoseInput;
        /// <summary>
        /// If the streaming source is a Kinesis data stream, identifies the stream's Amazon Resource Name (ARN).
        /// </summary>
        public readonly Outputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInput? KinesisStreamsInput;
        /// <summary>
        /// The name prefix to use when creating an in-application stream.
        /// </summary>
        public readonly string NamePrefix;

        [OutputConstructor]
        private ApplicationApplicationConfigurationSqlApplicationConfigurationInput(
            ImmutableArray<string> inAppStreamNames,

            string? inputId,

            Outputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputParallelism? inputParallelism,

            Outputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputProcessingConfiguration? inputProcessingConfiguration,

            Outputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchema inputSchema,

            ImmutableArray<Outputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfiguration> inputStartingPositionConfigurations,

            Outputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisFirehoseInput? kinesisFirehoseInput,

            Outputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInput? kinesisStreamsInput,

            string namePrefix)
        {
            InAppStreamNames = inAppStreamNames;
            InputId = inputId;
            InputParallelism = inputParallelism;
            InputProcessingConfiguration = inputProcessingConfiguration;
            InputSchema = inputSchema;
            InputStartingPositionConfigurations = inputStartingPositionConfigurations;
            KinesisFirehoseInput = kinesisFirehoseInput;
            KinesisStreamsInput = kinesisStreamsInput;
            NamePrefix = namePrefix;
        }
    }
}
