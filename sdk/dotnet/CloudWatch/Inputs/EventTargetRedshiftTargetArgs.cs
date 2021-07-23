// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch.Inputs
{

    public sealed class EventTargetRedshiftTargetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// The database user name.
        /// </summary>
        [Input("dbUser")]
        public Input<string>? DbUser { get; set; }

        /// <summary>
        /// The name or ARN of the secret that enables access to the database.
        /// </summary>
        [Input("secretsManagerArn")]
        public Input<string>? SecretsManagerArn { get; set; }

        /// <summary>
        /// The SQL statement text to run.
        /// </summary>
        [Input("sql")]
        public Input<string>? Sql { get; set; }

        /// <summary>
        /// The name of the SQL statement.
        /// </summary>
        [Input("statementName")]
        public Input<string>? StatementName { get; set; }

        /// <summary>
        /// Indicates whether to send an event back to EventBridge after the SQL statement runs.
        /// </summary>
        [Input("withEvent")]
        public Input<bool>? WithEvent { get; set; }

        public EventTargetRedshiftTargetArgs()
        {
        }
    }
}
