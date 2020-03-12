// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public partial class SecurityGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the security group
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The security group description. Defaults to
        /// "Managed by Pulumi". Cannot be "". __NOTE__: This field maps to the AWS
        /// `GroupDescription` attribute, for which there is no Update API. If you'd like
        /// to classify your security groups in a way that can be updated, use `tags`.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Can be specified multiple times for each
        /// egress rule. Each egress block supports fields documented below.
        /// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        [Output("egress")]
        public Output<ImmutableArray<Outputs.SecurityGroupEgress>> Egress { get; private set; } = null!;

        /// <summary>
        /// Can be specified multiple times for each
        /// ingress rule. Each ingress block supports fields documented below.
        /// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        [Output("ingress")]
        public Output<ImmutableArray<Outputs.SecurityGroupIngress>> Ingress { get; private set; } = null!;

        /// <summary>
        /// The name of the security group. If omitted, this provider will
        /// assign a random, unique name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the specified
        /// prefix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// The owner ID.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// Instruct this provider to revoke all of the
        /// Security Groups attached ingress and egress rules before deleting the rule
        /// itself. This is normally not needed, however certain AWS services such as
        /// Elastic Map Reduce may automatically add required rules to security groups used
        /// with the service, and those rules may contain a cyclic dependency that prevent
        /// the security groups from being destroyed without removing the dependency first.
        /// Default `false`
        /// </summary>
        [Output("revokeRulesOnDelete")]
        public Output<bool?> RevokeRulesOnDelete { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The VPC ID.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityGroup(string name, SecurityGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:ec2/securityGroup:SecurityGroup", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SecurityGroup(string name, Input<string> id, SecurityGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/securityGroup:SecurityGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityGroup Get(string name, Input<string> id, SecurityGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityGroup(name, id, state, options);
        }
    }

    public sealed class SecurityGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The security group description. Defaults to
        /// "Managed by Pulumi". Cannot be "". __NOTE__: This field maps to the AWS
        /// `GroupDescription` attribute, for which there is no Update API. If you'd like
        /// to classify your security groups in a way that can be updated, use `tags`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("egress")]
        private InputList<Inputs.SecurityGroupEgressArgs>? _egress;

        /// <summary>
        /// Can be specified multiple times for each
        /// egress rule. Each egress block supports fields documented below.
        /// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        public InputList<Inputs.SecurityGroupEgressArgs> Egress
        {
            get => _egress ?? (_egress = new InputList<Inputs.SecurityGroupEgressArgs>());
            set => _egress = value;
        }

        [Input("ingress")]
        private InputList<Inputs.SecurityGroupIngressArgs>? _ingress;

        /// <summary>
        /// Can be specified multiple times for each
        /// ingress rule. Each ingress block supports fields documented below.
        /// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        public InputList<Inputs.SecurityGroupIngressArgs> Ingress
        {
            get => _ingress ?? (_ingress = new InputList<Inputs.SecurityGroupIngressArgs>());
            set => _ingress = value;
        }

        /// <summary>
        /// The name of the security group. If omitted, this provider will
        /// assign a random, unique name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified
        /// prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// Instruct this provider to revoke all of the
        /// Security Groups attached ingress and egress rules before deleting the rule
        /// itself. This is normally not needed, however certain AWS services such as
        /// Elastic Map Reduce may automatically add required rules to security groups used
        /// with the service, and those rules may contain a cyclic dependency that prevent
        /// the security groups from being destroyed without removing the dependency first.
        /// Default `false`
        /// </summary>
        [Input("revokeRulesOnDelete")]
        public Input<bool>? RevokeRulesOnDelete { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The VPC ID.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public SecurityGroupArgs()
        {
            Description = "Managed by Pulumi";
        }
    }

    public sealed class SecurityGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the security group
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The security group description. Defaults to
        /// "Managed by Pulumi". Cannot be "". __NOTE__: This field maps to the AWS
        /// `GroupDescription` attribute, for which there is no Update API. If you'd like
        /// to classify your security groups in a way that can be updated, use `tags`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("egress")]
        private InputList<Inputs.SecurityGroupEgressGetArgs>? _egress;

        /// <summary>
        /// Can be specified multiple times for each
        /// egress rule. Each egress block supports fields documented below.
        /// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        public InputList<Inputs.SecurityGroupEgressGetArgs> Egress
        {
            get => _egress ?? (_egress = new InputList<Inputs.SecurityGroupEgressGetArgs>());
            set => _egress = value;
        }

        [Input("ingress")]
        private InputList<Inputs.SecurityGroupIngressGetArgs>? _ingress;

        /// <summary>
        /// Can be specified multiple times for each
        /// ingress rule. Each ingress block supports fields documented below.
        /// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        public InputList<Inputs.SecurityGroupIngressGetArgs> Ingress
        {
            get => _ingress ?? (_ingress = new InputList<Inputs.SecurityGroupIngressGetArgs>());
            set => _ingress = value;
        }

        /// <summary>
        /// The name of the security group. If omitted, this provider will
        /// assign a random, unique name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified
        /// prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The owner ID.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        /// <summary>
        /// Instruct this provider to revoke all of the
        /// Security Groups attached ingress and egress rules before deleting the rule
        /// itself. This is normally not needed, however certain AWS services such as
        /// Elastic Map Reduce may automatically add required rules to security groups used
        /// with the service, and those rules may contain a cyclic dependency that prevent
        /// the security groups from being destroyed without removing the dependency first.
        /// Default `false`
        /// </summary>
        [Input("revokeRulesOnDelete")]
        public Input<bool>? RevokeRulesOnDelete { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The VPC ID.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public SecurityGroupState()
        {
            Description = "Managed by Pulumi";
        }
    }

    namespace Inputs
    {

    public sealed class SecurityGroupEgressArgs : Pulumi.ResourceArgs
    {
        [Input("cidrBlocks")]
        private InputList<string>? _cidrBlocks;
        public InputList<string> CidrBlocks
        {
            get => _cidrBlocks ?? (_cidrBlocks = new InputList<string>());
            set => _cidrBlocks = value;
        }

        /// <summary>
        /// The security group description. Defaults to
        /// "Managed by Pulumi". Cannot be "". __NOTE__: This field maps to the AWS
        /// `GroupDescription` attribute, for which there is no Update API. If you'd like
        /// to classify your security groups in a way that can be updated, use `tags`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("fromPort", required: true)]
        public Input<int> FromPort { get; set; } = null!;

        [Input("ipv6CidrBlocks")]
        private InputList<string>? _ipv6CidrBlocks;
        public InputList<string> Ipv6CidrBlocks
        {
            get => _ipv6CidrBlocks ?? (_ipv6CidrBlocks = new InputList<string>());
            set => _ipv6CidrBlocks = value;
        }

        [Input("prefixListIds")]
        private InputList<string>? _prefixListIds;
        public InputList<string> PrefixListIds
        {
            get => _prefixListIds ?? (_prefixListIds = new InputList<string>());
            set => _prefixListIds = value;
        }

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("self")]
        public Input<bool>? Self { get; set; }

        [Input("toPort", required: true)]
        public Input<int> ToPort { get; set; } = null!;

        public SecurityGroupEgressArgs()
        {
        }
    }

    public sealed class SecurityGroupEgressGetArgs : Pulumi.ResourceArgs
    {
        [Input("cidrBlocks")]
        private InputList<string>? _cidrBlocks;
        public InputList<string> CidrBlocks
        {
            get => _cidrBlocks ?? (_cidrBlocks = new InputList<string>());
            set => _cidrBlocks = value;
        }

        /// <summary>
        /// The security group description. Defaults to
        /// "Managed by Pulumi". Cannot be "". __NOTE__: This field maps to the AWS
        /// `GroupDescription` attribute, for which there is no Update API. If you'd like
        /// to classify your security groups in a way that can be updated, use `tags`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("fromPort", required: true)]
        public Input<int> FromPort { get; set; } = null!;

        [Input("ipv6CidrBlocks")]
        private InputList<string>? _ipv6CidrBlocks;
        public InputList<string> Ipv6CidrBlocks
        {
            get => _ipv6CidrBlocks ?? (_ipv6CidrBlocks = new InputList<string>());
            set => _ipv6CidrBlocks = value;
        }

        [Input("prefixListIds")]
        private InputList<string>? _prefixListIds;
        public InputList<string> PrefixListIds
        {
            get => _prefixListIds ?? (_prefixListIds = new InputList<string>());
            set => _prefixListIds = value;
        }

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("self")]
        public Input<bool>? Self { get; set; }

        [Input("toPort", required: true)]
        public Input<int> ToPort { get; set; } = null!;

        public SecurityGroupEgressGetArgs()
        {
        }
    }

    public sealed class SecurityGroupIngressArgs : Pulumi.ResourceArgs
    {
        [Input("cidrBlocks")]
        private InputList<string>? _cidrBlocks;
        public InputList<string> CidrBlocks
        {
            get => _cidrBlocks ?? (_cidrBlocks = new InputList<string>());
            set => _cidrBlocks = value;
        }

        /// <summary>
        /// The security group description. Defaults to
        /// "Managed by Pulumi". Cannot be "". __NOTE__: This field maps to the AWS
        /// `GroupDescription` attribute, for which there is no Update API. If you'd like
        /// to classify your security groups in a way that can be updated, use `tags`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("fromPort", required: true)]
        public Input<int> FromPort { get; set; } = null!;

        [Input("ipv6CidrBlocks")]
        private InputList<string>? _ipv6CidrBlocks;
        public InputList<string> Ipv6CidrBlocks
        {
            get => _ipv6CidrBlocks ?? (_ipv6CidrBlocks = new InputList<string>());
            set => _ipv6CidrBlocks = value;
        }

        [Input("prefixListIds")]
        private InputList<string>? _prefixListIds;
        public InputList<string> PrefixListIds
        {
            get => _prefixListIds ?? (_prefixListIds = new InputList<string>());
            set => _prefixListIds = value;
        }

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("self")]
        public Input<bool>? Self { get; set; }

        [Input("toPort", required: true)]
        public Input<int> ToPort { get; set; } = null!;

        public SecurityGroupIngressArgs()
        {
        }
    }

    public sealed class SecurityGroupIngressGetArgs : Pulumi.ResourceArgs
    {
        [Input("cidrBlocks")]
        private InputList<string>? _cidrBlocks;
        public InputList<string> CidrBlocks
        {
            get => _cidrBlocks ?? (_cidrBlocks = new InputList<string>());
            set => _cidrBlocks = value;
        }

        /// <summary>
        /// The security group description. Defaults to
        /// "Managed by Pulumi". Cannot be "". __NOTE__: This field maps to the AWS
        /// `GroupDescription` attribute, for which there is no Update API. If you'd like
        /// to classify your security groups in a way that can be updated, use `tags`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("fromPort", required: true)]
        public Input<int> FromPort { get; set; } = null!;

        [Input("ipv6CidrBlocks")]
        private InputList<string>? _ipv6CidrBlocks;
        public InputList<string> Ipv6CidrBlocks
        {
            get => _ipv6CidrBlocks ?? (_ipv6CidrBlocks = new InputList<string>());
            set => _ipv6CidrBlocks = value;
        }

        [Input("prefixListIds")]
        private InputList<string>? _prefixListIds;
        public InputList<string> PrefixListIds
        {
            get => _prefixListIds ?? (_prefixListIds = new InputList<string>());
            set => _prefixListIds = value;
        }

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("self")]
        public Input<bool>? Self { get; set; }

        [Input("toPort", required: true)]
        public Input<int> ToPort { get; set; } = null!;

        public SecurityGroupIngressGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class SecurityGroupEgress
    {
        public readonly ImmutableArray<string> CidrBlocks;
        /// <summary>
        /// The security group description. Defaults to
        /// "Managed by Pulumi". Cannot be "". __NOTE__: This field maps to the AWS
        /// `GroupDescription` attribute, for which there is no Update API. If you'd like
        /// to classify your security groups in a way that can be updated, use `tags`.
        /// </summary>
        public readonly string? Description;
        public readonly int FromPort;
        public readonly ImmutableArray<string> Ipv6CidrBlocks;
        public readonly ImmutableArray<string> PrefixListIds;
        public readonly string Protocol;
        public readonly ImmutableArray<string> SecurityGroups;
        public readonly bool? Self;
        public readonly int ToPort;

        [OutputConstructor]
        private SecurityGroupEgress(
            ImmutableArray<string> cidrBlocks,
            string? description,
            int fromPort,
            ImmutableArray<string> ipv6CidrBlocks,
            ImmutableArray<string> prefixListIds,
            string protocol,
            ImmutableArray<string> securityGroups,
            bool? self,
            int toPort)
        {
            CidrBlocks = cidrBlocks;
            Description = description;
            FromPort = fromPort;
            Ipv6CidrBlocks = ipv6CidrBlocks;
            PrefixListIds = prefixListIds;
            Protocol = protocol;
            SecurityGroups = securityGroups;
            Self = self;
            ToPort = toPort;
        }
    }

    [OutputType]
    public sealed class SecurityGroupIngress
    {
        public readonly ImmutableArray<string> CidrBlocks;
        /// <summary>
        /// The security group description. Defaults to
        /// "Managed by Pulumi". Cannot be "". __NOTE__: This field maps to the AWS
        /// `GroupDescription` attribute, for which there is no Update API. If you'd like
        /// to classify your security groups in a way that can be updated, use `tags`.
        /// </summary>
        public readonly string? Description;
        public readonly int FromPort;
        public readonly ImmutableArray<string> Ipv6CidrBlocks;
        public readonly ImmutableArray<string> PrefixListIds;
        public readonly string Protocol;
        public readonly ImmutableArray<string> SecurityGroups;
        public readonly bool? Self;
        public readonly int ToPort;

        [OutputConstructor]
        private SecurityGroupIngress(
            ImmutableArray<string> cidrBlocks,
            string? description,
            int fromPort,
            ImmutableArray<string> ipv6CidrBlocks,
            ImmutableArray<string> prefixListIds,
            string protocol,
            ImmutableArray<string> securityGroups,
            bool? self,
            int toPort)
        {
            CidrBlocks = cidrBlocks;
            Description = description;
            FromPort = fromPort;
            Ipv6CidrBlocks = ipv6CidrBlocks;
            PrefixListIds = prefixListIds;
            Protocol = protocol;
            SecurityGroups = securityGroups;
            Self = self;
            ToPort = toPort;
        }
    }
    }
}
