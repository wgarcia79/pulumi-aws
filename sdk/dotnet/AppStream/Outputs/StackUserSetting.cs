// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppStream.Outputs
{

    [OutputType]
    public sealed class StackUserSetting
    {
        /// <summary>
        /// Action that is enabled or disabled. Valid values are: `CLIPBOARD_COPY_FROM_LOCAL_DEVICE`,  `CLIPBOARD_COPY_TO_LOCAL_DEVICE`, `FILE_UPLOAD`, `FILE_DOWNLOAD`, `PRINTING_TO_LOCAL_DEVICE`, `DOMAIN_PASSWORD_SIGNIN`, `DOMAIN_SMART_CARD_SIGNIN`.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Indicates whether the action is enabled or disabled. Valid values are: `ENABLED`, `DISABLED`.
        /// </summary>
        public readonly string Permission;

        [OutputConstructor]
        private StackUserSetting(
            string action,

            string permission)
        {
            Action = action;
            Permission = permission;
        }
    }
}
