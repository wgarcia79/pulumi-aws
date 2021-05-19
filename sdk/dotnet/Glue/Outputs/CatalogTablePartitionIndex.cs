// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue.Outputs
{

    [OutputType]
    public sealed class CatalogTablePartitionIndex
    {
        /// <summary>
        /// Name of the partition index.
        /// </summary>
        public readonly string IndexName;
        public readonly string? IndexStatus;
        /// <summary>
        /// Keys for the partition index.
        /// </summary>
        public readonly ImmutableArray<string> Keys;

        [OutputConstructor]
        private CatalogTablePartitionIndex(
            string indexName,

            string? indexStatus,

            ImmutableArray<string> keys)
        {
            IndexName = indexName;
            IndexStatus = indexStatus;
            Keys = keys;
        }
    }
}
