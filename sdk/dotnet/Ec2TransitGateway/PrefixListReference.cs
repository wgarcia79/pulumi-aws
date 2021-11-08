// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2TransitGateway
{
    /// <summary>
    /// Manages an EC2 Transit Gateway Prefix List Reference.
    /// 
    /// ## Example Usage
    /// ### Attachment Routing
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Ec2TransitGateway.PrefixListReference("example", new Aws.Ec2TransitGateway.PrefixListReferenceArgs
    ///         {
    ///             PrefixListId = aws_ec2_managed_prefix_list.Example.Id,
    ///             TransitGatewayAttachmentId = aws_ec2_transit_gateway_vpc_attachment.Example.Id,
    ///             TransitGatewayRouteTableId = aws_ec2_transit_gateway.Example.Association_default_route_table_id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Blackhole Routing
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Ec2TransitGateway.PrefixListReference("example", new Aws.Ec2TransitGateway.PrefixListReferenceArgs
    ///         {
    ///             Blackhole = true,
    ///             PrefixListId = aws_ec2_managed_prefix_list.Example.Id,
    ///             TransitGatewayRouteTableId = aws_ec2_transit_gateway.Example.Association_default_route_table_id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_ec2_transit_gateway_prefix_list_reference` can be imported by using the EC2 Transit Gateway Route Table identifier and EC2 Prefix List identifier, separated by an underscore (`_`), e.g., console
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2transitgateway/prefixListReference:PrefixListReference example tgw-rtb-12345678_pl-12345678
    /// ```
    /// </summary>
    [AwsResourceType("aws:ec2transitgateway/prefixListReference:PrefixListReference")]
    public partial class PrefixListReference : Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether to drop traffic that matches the Prefix List. Defaults to `false`.
        /// </summary>
        [Output("blackhole")]
        public Output<bool?> Blackhole { get; private set; } = null!;

        /// <summary>
        /// Identifier of EC2 Prefix List.
        /// </summary>
        [Output("prefixListId")]
        public Output<string> PrefixListId { get; private set; } = null!;

        [Output("prefixListOwnerId")]
        public Output<string> PrefixListOwnerId { get; private set; } = null!;

        /// <summary>
        /// Identifier of EC2 Transit Gateway Attachment.
        /// </summary>
        [Output("transitGatewayAttachmentId")]
        public Output<string?> TransitGatewayAttachmentId { get; private set; } = null!;

        /// <summary>
        /// Identifier of EC2 Transit Gateway Route Table.
        /// </summary>
        [Output("transitGatewayRouteTableId")]
        public Output<string> TransitGatewayRouteTableId { get; private set; } = null!;


        /// <summary>
        /// Create a PrefixListReference resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrefixListReference(string name, PrefixListReferenceArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2transitgateway/prefixListReference:PrefixListReference", name, args ?? new PrefixListReferenceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrefixListReference(string name, Input<string> id, PrefixListReferenceState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2transitgateway/prefixListReference:PrefixListReference", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrefixListReference resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrefixListReference Get(string name, Input<string> id, PrefixListReferenceState? state = null, CustomResourceOptions? options = null)
        {
            return new PrefixListReference(name, id, state, options);
        }
    }

    public sealed class PrefixListReferenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether to drop traffic that matches the Prefix List. Defaults to `false`.
        /// </summary>
        [Input("blackhole")]
        public Input<bool>? Blackhole { get; set; }

        /// <summary>
        /// Identifier of EC2 Prefix List.
        /// </summary>
        [Input("prefixListId", required: true)]
        public Input<string> PrefixListId { get; set; } = null!;

        /// <summary>
        /// Identifier of EC2 Transit Gateway Attachment.
        /// </summary>
        [Input("transitGatewayAttachmentId")]
        public Input<string>? TransitGatewayAttachmentId { get; set; }

        /// <summary>
        /// Identifier of EC2 Transit Gateway Route Table.
        /// </summary>
        [Input("transitGatewayRouteTableId", required: true)]
        public Input<string> TransitGatewayRouteTableId { get; set; } = null!;

        public PrefixListReferenceArgs()
        {
        }
    }

    public sealed class PrefixListReferenceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether to drop traffic that matches the Prefix List. Defaults to `false`.
        /// </summary>
        [Input("blackhole")]
        public Input<bool>? Blackhole { get; set; }

        /// <summary>
        /// Identifier of EC2 Prefix List.
        /// </summary>
        [Input("prefixListId")]
        public Input<string>? PrefixListId { get; set; }

        [Input("prefixListOwnerId")]
        public Input<string>? PrefixListOwnerId { get; set; }

        /// <summary>
        /// Identifier of EC2 Transit Gateway Attachment.
        /// </summary>
        [Input("transitGatewayAttachmentId")]
        public Input<string>? TransitGatewayAttachmentId { get; set; }

        /// <summary>
        /// Identifier of EC2 Transit Gateway Route Table.
        /// </summary>
        [Input("transitGatewayRouteTableId")]
        public Input<string>? TransitGatewayRouteTableId { get; set; }

        public PrefixListReferenceState()
        {
        }
    }
}
