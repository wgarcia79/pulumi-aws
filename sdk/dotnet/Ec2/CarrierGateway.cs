// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Manages an EC2 Carrier Gateway. See the AWS [documentation](https://docs.aws.amazon.com/vpc/latest/userguide/Carrier_Gateway.html) for more information.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Ec2.CarrierGateway("example", new Aws.Ec2.CarrierGatewayArgs
    ///         {
    ///             VpcId = aws_vpc.Example.Id,
    ///             Tags = 
    ///             {
    ///                 { "Name", "example-carrier-gateway" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_ec2_carrier_gateway` can be imported using the carrier gateway's ID, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2/carrierGateway:CarrierGateway example cgw-12345
    /// ```
    /// </summary>
    [AwsResourceType("aws:ec2/carrierGateway:CarrierGateway")]
    public partial class CarrierGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the carrier gateway.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The AWS account ID of the owner of the carrier gateway.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC to associate with the carrier gateway.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a CarrierGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CarrierGateway(string name, CarrierGatewayArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/carrierGateway:CarrierGateway", name, args ?? new CarrierGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CarrierGateway(string name, Input<string> id, CarrierGatewayState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/carrierGateway:CarrierGateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CarrierGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CarrierGateway Get(string name, Input<string> id, CarrierGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new CarrierGateway(name, id, state, options);
        }
    }

    public sealed class CarrierGatewayArgs : Pulumi.ResourceArgs
    {
        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the VPC to associate with the carrier gateway.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public CarrierGatewayArgs()
        {
        }
    }

    public sealed class CarrierGatewayState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the carrier gateway.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The AWS account ID of the owner of the carrier gateway.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// The ID of the VPC to associate with the carrier gateway.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public CarrierGatewayState()
        {
        }
    }
}
