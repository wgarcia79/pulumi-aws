# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'PlanAdvancedBackupSetting',
    'PlanRule',
    'PlanRuleCopyAction',
    'PlanRuleCopyActionLifecycle',
    'PlanRuleLifecycle',
    'SelectionSelectionTag',
]

@pulumi.output_type
class PlanAdvancedBackupSetting(dict):
    def __init__(__self__, *,
                 backup_options: Mapping[str, str],
                 resource_type: str):
        """
        :param Mapping[str, str] backup_options: Specifies the backup option for a selected resource. This option is only available for Windows VSS backup jobs. Set to `{ WindowsVSS = "enabled" }` to enable Windows VSS backup option and create a VSS Windows backup.
        :param str resource_type: The type of AWS resource to be backed up. For VSS Windows backups, the only supported resource type is Amazon EC2. Valid values: `EC2`.
        """
        pulumi.set(__self__, "backup_options", backup_options)
        pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter(name="backupOptions")
    def backup_options(self) -> Mapping[str, str]:
        """
        Specifies the backup option for a selected resource. This option is only available for Windows VSS backup jobs. Set to `{ WindowsVSS = "enabled" }` to enable Windows VSS backup option and create a VSS Windows backup.
        """
        return pulumi.get(self, "backup_options")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        """
        The type of AWS resource to be backed up. For VSS Windows backups, the only supported resource type is Amazon EC2. Valid values: `EC2`.
        """
        return pulumi.get(self, "resource_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PlanRule(dict):
    def __init__(__self__, *,
                 rule_name: str,
                 target_vault_name: str,
                 completion_window: Optional[int] = None,
                 copy_actions: Optional[Sequence['outputs.PlanRuleCopyAction']] = None,
                 enable_continuous_backup: Optional[bool] = None,
                 lifecycle: Optional['outputs.PlanRuleLifecycle'] = None,
                 recovery_point_tags: Optional[Mapping[str, str]] = None,
                 schedule: Optional[str] = None,
                 start_window: Optional[int] = None):
        """
        :param str rule_name: An display name for a backup rule.
        :param str target_vault_name: The name of a logical container where backups are stored.
        :param int completion_window: The amount of time AWS Backup attempts a backup before canceling the job and returning an error.
        :param Sequence['PlanRuleCopyActionArgs'] copy_actions: Configuration block(s) with copy operation settings. Detailed below.
        :param bool enable_continuous_backup: Enable continuous backups for supported resources.
        :param 'PlanRuleLifecycleArgs' lifecycle: The lifecycle defines when a protected resource is copied over to a backup vault and when it expires.  Fields documented above.
        :param Mapping[str, str] recovery_point_tags: Metadata that you can assign to help organize the resources that you create.
        :param str schedule: A CRON expression specifying when AWS Backup initiates a backup job.
        :param int start_window: The amount of time in minutes before beginning a backup.
        """
        pulumi.set(__self__, "rule_name", rule_name)
        pulumi.set(__self__, "target_vault_name", target_vault_name)
        if completion_window is not None:
            pulumi.set(__self__, "completion_window", completion_window)
        if copy_actions is not None:
            pulumi.set(__self__, "copy_actions", copy_actions)
        if enable_continuous_backup is not None:
            pulumi.set(__self__, "enable_continuous_backup", enable_continuous_backup)
        if lifecycle is not None:
            pulumi.set(__self__, "lifecycle", lifecycle)
        if recovery_point_tags is not None:
            pulumi.set(__self__, "recovery_point_tags", recovery_point_tags)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if start_window is not None:
            pulumi.set(__self__, "start_window", start_window)

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> str:
        """
        An display name for a backup rule.
        """
        return pulumi.get(self, "rule_name")

    @property
    @pulumi.getter(name="targetVaultName")
    def target_vault_name(self) -> str:
        """
        The name of a logical container where backups are stored.
        """
        return pulumi.get(self, "target_vault_name")

    @property
    @pulumi.getter(name="completionWindow")
    def completion_window(self) -> Optional[int]:
        """
        The amount of time AWS Backup attempts a backup before canceling the job and returning an error.
        """
        return pulumi.get(self, "completion_window")

    @property
    @pulumi.getter(name="copyActions")
    def copy_actions(self) -> Optional[Sequence['outputs.PlanRuleCopyAction']]:
        """
        Configuration block(s) with copy operation settings. Detailed below.
        """
        return pulumi.get(self, "copy_actions")

    @property
    @pulumi.getter(name="enableContinuousBackup")
    def enable_continuous_backup(self) -> Optional[bool]:
        """
        Enable continuous backups for supported resources.
        """
        return pulumi.get(self, "enable_continuous_backup")

    @property
    @pulumi.getter
    def lifecycle(self) -> Optional['outputs.PlanRuleLifecycle']:
        """
        The lifecycle defines when a protected resource is copied over to a backup vault and when it expires.  Fields documented above.
        """
        return pulumi.get(self, "lifecycle")

    @property
    @pulumi.getter(name="recoveryPointTags")
    def recovery_point_tags(self) -> Optional[Mapping[str, str]]:
        """
        Metadata that you can assign to help organize the resources that you create.
        """
        return pulumi.get(self, "recovery_point_tags")

    @property
    @pulumi.getter
    def schedule(self) -> Optional[str]:
        """
        A CRON expression specifying when AWS Backup initiates a backup job.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="startWindow")
    def start_window(self) -> Optional[int]:
        """
        The amount of time in minutes before beginning a backup.
        """
        return pulumi.get(self, "start_window")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PlanRuleCopyAction(dict):
    def __init__(__self__, *,
                 destination_vault_arn: str,
                 lifecycle: Optional['outputs.PlanRuleCopyActionLifecycle'] = None):
        """
        :param str destination_vault_arn: An Amazon Resource Name (ARN) that uniquely identifies the destination backup vault for the copied backup.
        :param 'PlanRuleCopyActionLifecycleArgs' lifecycle: The lifecycle defines when a protected resource is copied over to a backup vault and when it expires.  Fields documented above.
        """
        pulumi.set(__self__, "destination_vault_arn", destination_vault_arn)
        if lifecycle is not None:
            pulumi.set(__self__, "lifecycle", lifecycle)

    @property
    @pulumi.getter(name="destinationVaultArn")
    def destination_vault_arn(self) -> str:
        """
        An Amazon Resource Name (ARN) that uniquely identifies the destination backup vault for the copied backup.
        """
        return pulumi.get(self, "destination_vault_arn")

    @property
    @pulumi.getter
    def lifecycle(self) -> Optional['outputs.PlanRuleCopyActionLifecycle']:
        """
        The lifecycle defines when a protected resource is copied over to a backup vault and when it expires.  Fields documented above.
        """
        return pulumi.get(self, "lifecycle")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PlanRuleCopyActionLifecycle(dict):
    def __init__(__self__, *,
                 cold_storage_after: Optional[int] = None,
                 delete_after: Optional[int] = None):
        """
        :param int cold_storage_after: Specifies the number of days after creation that a recovery point is moved to cold storage.
        :param int delete_after: Specifies the number of days after creation that a recovery point is deleted. Must be 90 days greater than `cold_storage_after`.
        """
        if cold_storage_after is not None:
            pulumi.set(__self__, "cold_storage_after", cold_storage_after)
        if delete_after is not None:
            pulumi.set(__self__, "delete_after", delete_after)

    @property
    @pulumi.getter(name="coldStorageAfter")
    def cold_storage_after(self) -> Optional[int]:
        """
        Specifies the number of days after creation that a recovery point is moved to cold storage.
        """
        return pulumi.get(self, "cold_storage_after")

    @property
    @pulumi.getter(name="deleteAfter")
    def delete_after(self) -> Optional[int]:
        """
        Specifies the number of days after creation that a recovery point is deleted. Must be 90 days greater than `cold_storage_after`.
        """
        return pulumi.get(self, "delete_after")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PlanRuleLifecycle(dict):
    def __init__(__self__, *,
                 cold_storage_after: Optional[int] = None,
                 delete_after: Optional[int] = None):
        """
        :param int cold_storage_after: Specifies the number of days after creation that a recovery point is moved to cold storage.
        :param int delete_after: Specifies the number of days after creation that a recovery point is deleted. Must be 90 days greater than `cold_storage_after`.
        """
        if cold_storage_after is not None:
            pulumi.set(__self__, "cold_storage_after", cold_storage_after)
        if delete_after is not None:
            pulumi.set(__self__, "delete_after", delete_after)

    @property
    @pulumi.getter(name="coldStorageAfter")
    def cold_storage_after(self) -> Optional[int]:
        """
        Specifies the number of days after creation that a recovery point is moved to cold storage.
        """
        return pulumi.get(self, "cold_storage_after")

    @property
    @pulumi.getter(name="deleteAfter")
    def delete_after(self) -> Optional[int]:
        """
        Specifies the number of days after creation that a recovery point is deleted. Must be 90 days greater than `cold_storage_after`.
        """
        return pulumi.get(self, "delete_after")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SelectionSelectionTag(dict):
    def __init__(__self__, *,
                 key: str,
                 type: str,
                 value: str):
        """
        :param str key: The key in a key-value pair.
        :param str type: An operation, such as `StringEquals`, that is applied to a key-value pair used to filter resources in a selection.
        :param str value: The value in a key-value pair.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The key in a key-value pair.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        An operation, such as `StringEquals`, that is applied to a key-value pair used to filter resources in a selection.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value in a key-value pair.
        """
        return pulumi.get(self, "value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


