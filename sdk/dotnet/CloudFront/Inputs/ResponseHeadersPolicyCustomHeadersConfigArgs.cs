// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront.Inputs
{

    public sealed class ResponseHeadersPolicyCustomHeadersConfigArgs : Pulumi.ResourceArgs
    {
        [Input("items")]
        private InputList<Inputs.ResponseHeadersPolicyCustomHeadersConfigItemArgs>? _items;
        public InputList<Inputs.ResponseHeadersPolicyCustomHeadersConfigItemArgs> Items
        {
            get => _items ?? (_items = new InputList<Inputs.ResponseHeadersPolicyCustomHeadersConfigItemArgs>());
            set => _items = value;
        }

        public ResponseHeadersPolicyCustomHeadersConfigArgs()
        {
        }
    }
}
