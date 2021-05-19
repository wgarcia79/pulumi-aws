// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue.Inputs
{

    public sealed class CatalogTableStorageDescriptorSerDeInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the target table.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// Map of initialization parameters for the SerDe, in key-value form.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// Usually the class that implements the SerDe. An example is `org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe`.
        /// </summary>
        [Input("serializationLibrary")]
        public Input<string>? SerializationLibrary { get; set; }

        public CatalogTableStorageDescriptorSerDeInfoArgs()
        {
        }
    }
}
