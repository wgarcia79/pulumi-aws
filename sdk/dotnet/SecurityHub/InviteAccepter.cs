// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecurityHub
{
    /// <summary>
    /// &gt; **Note:** AWS accounts can only be associated with a single Security Hub master account. Destroying this resource will disassociate the member account from the master account.
    /// 
    /// Accepts a Security Hub invitation.
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
    ///         var exampleAccount = new Aws.SecurityHub.Account("exampleAccount", new Aws.SecurityHub.AccountArgs
    ///         {
    ///         });
    ///         var exampleMember = new Aws.SecurityHub.Member("exampleMember", new Aws.SecurityHub.MemberArgs
    ///         {
    ///             AccountId = "123456789012",
    ///             Email = "example@example.com",
    ///             Invite = true,
    ///         });
    ///         var inviteeAccount = new Aws.SecurityHub.Account("inviteeAccount", new Aws.SecurityHub.AccountArgs
    ///         {
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = "aws.invitee",
    ///         });
    ///         var inviteeInviteAccepter = new Aws.SecurityHub.InviteAccepter("inviteeInviteAccepter", new Aws.SecurityHub.InviteAccepterArgs
    ///         {
    ///             MasterId = exampleMember.MasterId,
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = "aws.invitee",
    ///             DependsOn = 
    ///             {
    ///                 inviteeAccount,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Security Hub invite acceptance can be imported using the account ID, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:securityhub/inviteAccepter:InviteAccepter example 123456789012
    /// ```
    /// </summary>
    [AwsResourceType("aws:securityhub/inviteAccepter:InviteAccepter")]
    public partial class InviteAccepter : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the invitation.
        /// </summary>
        [Output("invitationId")]
        public Output<string> InvitationId { get; private set; } = null!;

        /// <summary>
        /// The account ID of the master Security Hub account whose invitation you're accepting.
        /// </summary>
        [Output("masterId")]
        public Output<string> MasterId { get; private set; } = null!;


        /// <summary>
        /// Create a InviteAccepter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InviteAccepter(string name, InviteAccepterArgs args, CustomResourceOptions? options = null)
            : base("aws:securityhub/inviteAccepter:InviteAccepter", name, args ?? new InviteAccepterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InviteAccepter(string name, Input<string> id, InviteAccepterState? state = null, CustomResourceOptions? options = null)
            : base("aws:securityhub/inviteAccepter:InviteAccepter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InviteAccepter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InviteAccepter Get(string name, Input<string> id, InviteAccepterState? state = null, CustomResourceOptions? options = null)
        {
            return new InviteAccepter(name, id, state, options);
        }
    }

    public sealed class InviteAccepterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account ID of the master Security Hub account whose invitation you're accepting.
        /// </summary>
        [Input("masterId", required: true)]
        public Input<string> MasterId { get; set; } = null!;

        public InviteAccepterArgs()
        {
        }
    }

    public sealed class InviteAccepterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the invitation.
        /// </summary>
        [Input("invitationId")]
        public Input<string>? InvitationId { get; set; }

        /// <summary>
        /// The account ID of the master Security Hub account whose invitation you're accepting.
        /// </summary>
        [Input("masterId")]
        public Input<string>? MasterId { get; set; }

        public InviteAccepterState()
        {
        }
    }
}
