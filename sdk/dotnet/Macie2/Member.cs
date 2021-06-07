// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Macie2
{
    /// <summary>
    /// Provides a resource to manage an [Amazon Macie Member](https://docs.aws.amazon.com/macie/latest/APIReference/members-id.html).
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
    ///         var exampleAccount = new Aws.Macie2.Account("exampleAccount", new Aws.Macie2.AccountArgs
    ///         {
    ///         });
    ///         var exampleMember = new Aws.Macie2.Member("exampleMember", new Aws.Macie2.MemberArgs
    ///         {
    ///             AccountId = "AWS ACCOUNT ID",
    ///             Email = "EMAIL",
    ///             Invite = true,
    ///             InvitationMessage = "Message of the invitation",
    ///             InvitationDisableEmailNotification = "true",
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 exampleAccount,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_macie2_member` can be imported using the account ID of the member account, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:macie2/member:Member example 123456789012
    /// ```
    /// </summary>
    [AwsResourceType("aws:macie2/member:Member")]
    public partial class Member : Pulumi.CustomResource
    {
        /// <summary>
        /// The AWS account ID for the account.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// The AWS account ID for the administrator account.
        /// </summary>
        [Output("administratorAccountId")]
        public Output<string> AdministratorAccountId { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the account.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The email address for the account.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to send an email notification to the root user of each account that the invitation will be sent to. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. To send an email notification to the root user of each account, set this value to `true`.
        /// </summary>
        [Output("invitationDisableEmailNotification")]
        public Output<string?> InvitationDisableEmailNotification { get; private set; } = null!;

        /// <summary>
        /// A custom message to include in the invitation. Amazon Macie adds this message to the standard content that it sends for an invitation.
        /// </summary>
        [Output("invitationMessage")]
        public Output<string?> InvitationMessage { get; private set; } = null!;

        /// <summary>
        /// Send an invitation to a member
        /// </summary>
        [Output("invite")]
        public Output<bool> Invite { get; private set; } = null!;

        /// <summary>
        /// The date and time, in UTC and extended RFC 3339 format, when an Amazon Macie membership invitation was last sent to the account. This value is null if a Macie invitation hasn't been sent to the account.
        /// </summary>
        [Output("invitedAt")]
        public Output<string> InvitedAt { get; private set; } = null!;

        [Output("masterAccountId")]
        public Output<string> MasterAccountId { get; private set; } = null!;

        /// <summary>
        /// The current status of the relationship between the account and the administrator account.
        /// </summary>
        [Output("relationshipStatus")]
        public Output<string> RelationshipStatus { get; private set; } = null!;

        /// <summary>
        /// Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to `ENABLED`. Valid values are `ENABLED` or `PAUSED`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A map of key-value pairs that specifies the tags to associate with the account in Amazon Macie.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The date and time, in UTC and extended RFC 3339 format, of the most recent change to the status of the relationship between the account and the administrator account.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a Member resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Member(string name, MemberArgs args, CustomResourceOptions? options = null)
            : base("aws:macie2/member:Member", name, args ?? new MemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Member(string name, Input<string> id, MemberState? state = null, CustomResourceOptions? options = null)
            : base("aws:macie2/member:Member", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Member resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Member Get(string name, Input<string> id, MemberState? state = null, CustomResourceOptions? options = null)
        {
            return new Member(name, id, state, options);
        }
    }

    public sealed class MemberArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS account ID for the account.
        /// </summary>
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        /// <summary>
        /// The email address for the account.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        /// <summary>
        /// Specifies whether to send an email notification to the root user of each account that the invitation will be sent to. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. To send an email notification to the root user of each account, set this value to `true`.
        /// </summary>
        [Input("invitationDisableEmailNotification")]
        public Input<string>? InvitationDisableEmailNotification { get; set; }

        /// <summary>
        /// A custom message to include in the invitation. Amazon Macie adds this message to the standard content that it sends for an invitation.
        /// </summary>
        [Input("invitationMessage")]
        public Input<string>? InvitationMessage { get; set; }

        /// <summary>
        /// Send an invitation to a member
        /// </summary>
        [Input("invite")]
        public Input<bool>? Invite { get; set; }

        /// <summary>
        /// Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to `ENABLED`. Valid values are `ENABLED` or `PAUSED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of key-value pairs that specifies the tags to associate with the account in Amazon Macie.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public MemberArgs()
        {
        }
    }

    public sealed class MemberState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS account ID for the account.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The AWS account ID for the administrator account.
        /// </summary>
        [Input("administratorAccountId")]
        public Input<string>? AdministratorAccountId { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the account.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The email address for the account.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Specifies whether to send an email notification to the root user of each account that the invitation will be sent to. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. To send an email notification to the root user of each account, set this value to `true`.
        /// </summary>
        [Input("invitationDisableEmailNotification")]
        public Input<string>? InvitationDisableEmailNotification { get; set; }

        /// <summary>
        /// A custom message to include in the invitation. Amazon Macie adds this message to the standard content that it sends for an invitation.
        /// </summary>
        [Input("invitationMessage")]
        public Input<string>? InvitationMessage { get; set; }

        /// <summary>
        /// Send an invitation to a member
        /// </summary>
        [Input("invite")]
        public Input<bool>? Invite { get; set; }

        /// <summary>
        /// The date and time, in UTC and extended RFC 3339 format, when an Amazon Macie membership invitation was last sent to the account. This value is null if a Macie invitation hasn't been sent to the account.
        /// </summary>
        [Input("invitedAt")]
        public Input<string>? InvitedAt { get; set; }

        [Input("masterAccountId")]
        public Input<string>? MasterAccountId { get; set; }

        /// <summary>
        /// The current status of the relationship between the account and the administrator account.
        /// </summary>
        [Input("relationshipStatus")]
        public Input<string>? RelationshipStatus { get; set; }

        /// <summary>
        /// Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to `ENABLED`. Valid values are `ENABLED` or `PAUSED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of key-value pairs that specifies the tags to associate with the account in Amazon Macie.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// The date and time, in UTC and extended RFC 3339 format, of the most recent change to the status of the relationship between the account and the administrator account.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public MemberState()
        {
        }
    }
}