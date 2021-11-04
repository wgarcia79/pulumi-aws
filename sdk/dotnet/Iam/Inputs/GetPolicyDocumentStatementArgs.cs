// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iam.Inputs
{

    public sealed class GetPolicyDocumentStatementInputArgs : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<string>? _actions;

        /// <summary>
        /// List of actions that this statement either allows or denies. For example, `["ec2:RunInstances", "s3:*"]`.
        /// </summary>
        public InputList<string> Actions
        {
            get => _actions ?? (_actions = new InputList<string>());
            set => _actions = value;
        }

        [Input("conditions")]
        private InputList<Inputs.GetPolicyDocumentStatementConditionInputArgs>? _conditions;

        /// <summary>
        /// Configuration block for a condition. Detailed below.
        /// </summary>
        public InputList<Inputs.GetPolicyDocumentStatementConditionInputArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.GetPolicyDocumentStatementConditionInputArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// Whether this statement allows or denies the given actions. Valid values are `Allow` and `Deny`. Defaults to `Allow`.
        /// </summary>
        [Input("effect")]
        public Input<string>? Effect { get; set; }

        [Input("notActions")]
        private InputList<string>? _notActions;

        /// <summary>
        /// List of actions that this statement does *not* apply to. Use to apply a policy statement to all actions *except* those listed.
        /// </summary>
        public InputList<string> NotActions
        {
            get => _notActions ?? (_notActions = new InputList<string>());
            set => _notActions = value;
        }

        [Input("notPrincipals")]
        private InputList<Inputs.GetPolicyDocumentStatementNotPrincipalInputArgs>? _notPrincipals;

        /// <summary>
        /// Like `principals` except these are principals that the statement does *not* apply to.
        /// </summary>
        public InputList<Inputs.GetPolicyDocumentStatementNotPrincipalInputArgs> NotPrincipals
        {
            get => _notPrincipals ?? (_notPrincipals = new InputList<Inputs.GetPolicyDocumentStatementNotPrincipalInputArgs>());
            set => _notPrincipals = value;
        }

        [Input("notResources")]
        private InputList<string>? _notResources;

        /// <summary>
        /// List of resource ARNs that this statement does *not* apply to. Use to apply a policy statement to all resources *except* those listed.
        /// </summary>
        public InputList<string> NotResources
        {
            get => _notResources ?? (_notResources = new InputList<string>());
            set => _notResources = value;
        }

        [Input("principals")]
        private InputList<Inputs.GetPolicyDocumentStatementPrincipalInputArgs>? _principals;

        /// <summary>
        /// Configuration block for principals. Detailed below.
        /// </summary>
        public InputList<Inputs.GetPolicyDocumentStatementPrincipalInputArgs> Principals
        {
            get => _principals ?? (_principals = new InputList<Inputs.GetPolicyDocumentStatementPrincipalInputArgs>());
            set => _principals = value;
        }

        [Input("resources")]
        private InputList<string>? _resources;

        /// <summary>
        /// List of resource ARNs that this statement applies to. This is required by AWS if used for an IAM policy.
        /// </summary>
        public InputList<string> Resources
        {
            get => _resources ?? (_resources = new InputList<string>());
            set => _resources = value;
        }

        /// <summary>
        /// Sid (statement ID) is an identifier for a policy statement.
        /// </summary>
        [Input("sid")]
        public Input<string>? Sid { get; set; }

        public GetPolicyDocumentStatementInputArgs()
        {
        }
    }
}
