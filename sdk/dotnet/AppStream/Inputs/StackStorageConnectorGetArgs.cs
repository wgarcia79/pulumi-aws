// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppStream.Inputs
{

    public sealed class StackStorageConnectorGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of storage connector. Valid values are: `HOMEFOLDERS`, `GOOGLE_DRIVE`, `ONE_DRIVE`.
        /// </summary>
        [Input("connectorType", required: true)]
        public Input<string> ConnectorType { get; set; } = null!;

        [Input("domains")]
        private InputList<string>? _domains;

        /// <summary>
        /// Names of the domains for the account.
        /// </summary>
        public InputList<string> Domains
        {
            get => _domains ?? (_domains = new InputList<string>());
            set => _domains = value;
        }

        /// <summary>
        /// ARN of the storage connector.
        /// </summary>
        [Input("resourceIdentifier")]
        public Input<string>? ResourceIdentifier { get; set; }

        public StackStorageConnectorGetArgs()
        {
        }
    }
}
