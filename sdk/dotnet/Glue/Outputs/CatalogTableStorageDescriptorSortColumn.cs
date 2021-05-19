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
    public sealed class CatalogTableStorageDescriptorSortColumn
    {
        /// <summary>
        /// Name of the column.
        /// </summary>
        public readonly string Column;
        /// <summary>
        /// Whether the column is sorted in ascending (`1`) or descending order (`0`).
        /// </summary>
        public readonly int SortOrder;

        [OutputConstructor]
        private CatalogTableStorageDescriptorSortColumn(
            string column,

            int sortOrder)
        {
            Column = column;
            SortOrder = sortOrder;
        }
    }
}
