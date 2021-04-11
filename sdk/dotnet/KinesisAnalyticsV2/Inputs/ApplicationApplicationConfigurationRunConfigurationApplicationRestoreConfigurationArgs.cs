// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.KinesisAnalyticsV2.Inputs
{

    public sealed class ApplicationApplicationConfigurationRunConfigurationApplicationRestoreConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies how the application should be restored. Valid values: `RESTORE_FROM_CUSTOM_SNAPSHOT`, `RESTORE_FROM_LATEST_SNAPSHOT`, `SKIP_RESTORE_FROM_SNAPSHOT`.
        /// </summary>
        [Input("applicationRestoreType")]
        public Input<string>? ApplicationRestoreType { get; set; }

        /// <summary>
        /// The identifier of an existing snapshot of application state to use to restart an application. The application uses this value if `RESTORE_FROM_CUSTOM_SNAPSHOT` is specified for `application_restore_type`.
        /// </summary>
        [Input("snapshotName")]
        public Input<string>? SnapshotName { get; set; }

        public ApplicationApplicationConfigurationRunConfigurationApplicationRestoreConfigurationArgs()
        {
        }
    }
}
