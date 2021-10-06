// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppStream.Inputs
{

    public sealed class ImageBuilderDomainJoinInfoGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fully qualified name of the directory (for example, corp.example.com).
        /// </summary>
        [Input("directoryName")]
        public Input<string>? DirectoryName { get; set; }

        /// <summary>
        /// Distinguished name of the organizational unit for computer accounts.
        /// </summary>
        [Input("organizationalUnitDistinguishedName")]
        public Input<string>? OrganizationalUnitDistinguishedName { get; set; }

        public ImageBuilderDomainJoinInfoGetArgs()
        {
        }
    }
}
