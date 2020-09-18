// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeBuild.Outputs
{

    [OutputType]
    public sealed class ReportGroupExportConfigS3Destination
    {
        /// <summary>
        /// The name of the S3 bucket where the raw data of a report are exported.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// A boolean value that specifies if the results of a report are encrypted.
        /// **Note: the API does not currently allow setting encryption as disabled**
        /// </summary>
        public readonly bool? EncryptionDisabled;
        /// <summary>
        /// The encryption key for the report's encrypted raw data. The KMS key ARN.
        /// </summary>
        public readonly string EncryptionKey;
        /// <summary>
        /// The type of build output artifact to create. Valid values are: `NONE` (default) and `ZIP`.
        /// </summary>
        public readonly string? Packaging;
        /// <summary>
        /// The path to the exported report's raw data results.
        /// </summary>
        public readonly string? Path;

        [OutputConstructor]
        private ReportGroupExportConfigS3Destination(
            string bucket,

            bool? encryptionDisabled,

            string encryptionKey,

            string? packaging,

            string? path)
        {
            Bucket = bucket;
            EncryptionDisabled = encryptionDisabled;
            EncryptionKey = encryptionKey;
            Packaging = packaging;
            Path = path;
        }
    }
}
