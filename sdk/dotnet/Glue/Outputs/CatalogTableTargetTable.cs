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
    public sealed class CatalogTableTargetTable
    {
        /// <summary>
        /// ID of the Data Catalog in which the table resides.
        /// </summary>
        public readonly string CatalogId;
        /// <summary>
        /// Name of the catalog database that contains the target table.
        /// </summary>
        public readonly string DatabaseName;
        /// <summary>
        /// Name of the target table.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private CatalogTableTargetTable(
            string catalogId,

            string databaseName,

            string name)
        {
            CatalogId = catalogId;
            DatabaseName = databaseName;
            Name = name;
        }
    }
}
