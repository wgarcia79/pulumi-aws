// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kinesis.Outputs
{

    [OutputType]
    public sealed class FirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJsonSerDe
    {
        /// <summary>
        /// When set to true, which is the default, Kinesis Data Firehose converts JSON keys to lowercase before deserializing them.
        /// </summary>
        public readonly bool? CaseInsensitive;
        /// <summary>
        /// A map of column names to JSON keys that aren't identical to the column names. This is useful when the JSON contains keys that are Hive keywords. For example, timestamp is a Hive keyword. If you have a JSON key named timestamp, set this parameter to `{ ts = "timestamp" }` to map this key to a column named ts.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? ColumnToJsonKeyMappings;
        /// <summary>
        /// When set to `true`, specifies that the names of the keys include dots and that you want Kinesis Data Firehose to replace them with underscores. This is useful because Apache Hive does not allow dots in column names. For example, if the JSON contains a key whose name is "a.b", you can define the column name to be "a_b" when using this option. Defaults to `false`.
        /// </summary>
        public readonly bool? ConvertDotsInJsonKeysToUnderscores;

        [OutputConstructor]
        private FirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJsonSerDe(
            bool? caseInsensitive,

            ImmutableDictionary<string, string>? columnToJsonKeyMappings,

            bool? convertDotsInJsonKeysToUnderscores)
        {
            CaseInsensitive = caseInsensitive;
            ColumnToJsonKeyMappings = columnToJsonKeyMappings;
            ConvertDotsInJsonKeysToUnderscores = convertDotsInJsonKeysToUnderscores;
        }
    }
}