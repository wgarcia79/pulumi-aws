# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VaultLockConfigurationArgs', 'VaultLockConfiguration']

@pulumi.input_type
class VaultLockConfigurationArgs:
    def __init__(__self__, *,
                 backup_vault_name: pulumi.Input[str],
                 changeable_for_days: Optional[pulumi.Input[int]] = None,
                 max_retention_days: Optional[pulumi.Input[int]] = None,
                 min_retention_days: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a VaultLockConfiguration resource.
        :param pulumi.Input[str] backup_vault_name: Name of the backup vault to add a lock configuration for.
        :param pulumi.Input[int] changeable_for_days: The number of days before the lock date.
        :param pulumi.Input[int] max_retention_days: The maximum retention period that the vault retains its recovery points.
        :param pulumi.Input[int] min_retention_days: The minimum retention period that the vault retains its recovery points.
        """
        pulumi.set(__self__, "backup_vault_name", backup_vault_name)
        if changeable_for_days is not None:
            pulumi.set(__self__, "changeable_for_days", changeable_for_days)
        if max_retention_days is not None:
            pulumi.set(__self__, "max_retention_days", max_retention_days)
        if min_retention_days is not None:
            pulumi.set(__self__, "min_retention_days", min_retention_days)

    @property
    @pulumi.getter(name="backupVaultName")
    def backup_vault_name(self) -> pulumi.Input[str]:
        """
        Name of the backup vault to add a lock configuration for.
        """
        return pulumi.get(self, "backup_vault_name")

    @backup_vault_name.setter
    def backup_vault_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "backup_vault_name", value)

    @property
    @pulumi.getter(name="changeableForDays")
    def changeable_for_days(self) -> Optional[pulumi.Input[int]]:
        """
        The number of days before the lock date.
        """
        return pulumi.get(self, "changeable_for_days")

    @changeable_for_days.setter
    def changeable_for_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "changeable_for_days", value)

    @property
    @pulumi.getter(name="maxRetentionDays")
    def max_retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum retention period that the vault retains its recovery points.
        """
        return pulumi.get(self, "max_retention_days")

    @max_retention_days.setter
    def max_retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_retention_days", value)

    @property
    @pulumi.getter(name="minRetentionDays")
    def min_retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum retention period that the vault retains its recovery points.
        """
        return pulumi.get(self, "min_retention_days")

    @min_retention_days.setter
    def min_retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_retention_days", value)


@pulumi.input_type
class _VaultLockConfigurationState:
    def __init__(__self__, *,
                 backup_vault_arn: Optional[pulumi.Input[str]] = None,
                 backup_vault_name: Optional[pulumi.Input[str]] = None,
                 changeable_for_days: Optional[pulumi.Input[int]] = None,
                 max_retention_days: Optional[pulumi.Input[int]] = None,
                 min_retention_days: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering VaultLockConfiguration resources.
        :param pulumi.Input[str] backup_vault_arn: The ARN of the vault.
        :param pulumi.Input[str] backup_vault_name: Name of the backup vault to add a lock configuration for.
        :param pulumi.Input[int] changeable_for_days: The number of days before the lock date.
        :param pulumi.Input[int] max_retention_days: The maximum retention period that the vault retains its recovery points.
        :param pulumi.Input[int] min_retention_days: The minimum retention period that the vault retains its recovery points.
        """
        if backup_vault_arn is not None:
            pulumi.set(__self__, "backup_vault_arn", backup_vault_arn)
        if backup_vault_name is not None:
            pulumi.set(__self__, "backup_vault_name", backup_vault_name)
        if changeable_for_days is not None:
            pulumi.set(__self__, "changeable_for_days", changeable_for_days)
        if max_retention_days is not None:
            pulumi.set(__self__, "max_retention_days", max_retention_days)
        if min_retention_days is not None:
            pulumi.set(__self__, "min_retention_days", min_retention_days)

    @property
    @pulumi.getter(name="backupVaultArn")
    def backup_vault_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the vault.
        """
        return pulumi.get(self, "backup_vault_arn")

    @backup_vault_arn.setter
    def backup_vault_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_vault_arn", value)

    @property
    @pulumi.getter(name="backupVaultName")
    def backup_vault_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the backup vault to add a lock configuration for.
        """
        return pulumi.get(self, "backup_vault_name")

    @backup_vault_name.setter
    def backup_vault_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_vault_name", value)

    @property
    @pulumi.getter(name="changeableForDays")
    def changeable_for_days(self) -> Optional[pulumi.Input[int]]:
        """
        The number of days before the lock date.
        """
        return pulumi.get(self, "changeable_for_days")

    @changeable_for_days.setter
    def changeable_for_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "changeable_for_days", value)

    @property
    @pulumi.getter(name="maxRetentionDays")
    def max_retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum retention period that the vault retains its recovery points.
        """
        return pulumi.get(self, "max_retention_days")

    @max_retention_days.setter
    def max_retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_retention_days", value)

    @property
    @pulumi.getter(name="minRetentionDays")
    def min_retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum retention period that the vault retains its recovery points.
        """
        return pulumi.get(self, "min_retention_days")

    @min_retention_days.setter
    def min_retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_retention_days", value)


class VaultLockConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_vault_name: Optional[pulumi.Input[str]] = None,
                 changeable_for_days: Optional[pulumi.Input[int]] = None,
                 max_retention_days: Optional[pulumi.Input[int]] = None,
                 min_retention_days: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides an AWS Backup vault lock configuration resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.backup.VaultLockConfiguration("test",
            backup_vault_name="example_backup_vault",
            changeable_for_days=3,
            max_retention_days=1200,
            min_retention_days=7)
        ```

        ## Import

        Backup vault lock configuration can be imported using the `name`, e.g.,

        ```sh
         $ pulumi import aws:backup/vaultLockConfiguration:VaultLockConfiguration test TestVault
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_vault_name: Name of the backup vault to add a lock configuration for.
        :param pulumi.Input[int] changeable_for_days: The number of days before the lock date.
        :param pulumi.Input[int] max_retention_days: The maximum retention period that the vault retains its recovery points.
        :param pulumi.Input[int] min_retention_days: The minimum retention period that the vault retains its recovery points.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VaultLockConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AWS Backup vault lock configuration resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.backup.VaultLockConfiguration("test",
            backup_vault_name="example_backup_vault",
            changeable_for_days=3,
            max_retention_days=1200,
            min_retention_days=7)
        ```

        ## Import

        Backup vault lock configuration can be imported using the `name`, e.g.,

        ```sh
         $ pulumi import aws:backup/vaultLockConfiguration:VaultLockConfiguration test TestVault
        ```

        :param str resource_name: The name of the resource.
        :param VaultLockConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VaultLockConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_vault_name: Optional[pulumi.Input[str]] = None,
                 changeable_for_days: Optional[pulumi.Input[int]] = None,
                 max_retention_days: Optional[pulumi.Input[int]] = None,
                 min_retention_days: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VaultLockConfigurationArgs.__new__(VaultLockConfigurationArgs)

            if backup_vault_name is None and not opts.urn:
                raise TypeError("Missing required property 'backup_vault_name'")
            __props__.__dict__["backup_vault_name"] = backup_vault_name
            __props__.__dict__["changeable_for_days"] = changeable_for_days
            __props__.__dict__["max_retention_days"] = max_retention_days
            __props__.__dict__["min_retention_days"] = min_retention_days
            __props__.__dict__["backup_vault_arn"] = None
        super(VaultLockConfiguration, __self__).__init__(
            'aws:backup/vaultLockConfiguration:VaultLockConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_vault_arn: Optional[pulumi.Input[str]] = None,
            backup_vault_name: Optional[pulumi.Input[str]] = None,
            changeable_for_days: Optional[pulumi.Input[int]] = None,
            max_retention_days: Optional[pulumi.Input[int]] = None,
            min_retention_days: Optional[pulumi.Input[int]] = None) -> 'VaultLockConfiguration':
        """
        Get an existing VaultLockConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_vault_arn: The ARN of the vault.
        :param pulumi.Input[str] backup_vault_name: Name of the backup vault to add a lock configuration for.
        :param pulumi.Input[int] changeable_for_days: The number of days before the lock date.
        :param pulumi.Input[int] max_retention_days: The maximum retention period that the vault retains its recovery points.
        :param pulumi.Input[int] min_retention_days: The minimum retention period that the vault retains its recovery points.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VaultLockConfigurationState.__new__(_VaultLockConfigurationState)

        __props__.__dict__["backup_vault_arn"] = backup_vault_arn
        __props__.__dict__["backup_vault_name"] = backup_vault_name
        __props__.__dict__["changeable_for_days"] = changeable_for_days
        __props__.__dict__["max_retention_days"] = max_retention_days
        __props__.__dict__["min_retention_days"] = min_retention_days
        return VaultLockConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupVaultArn")
    def backup_vault_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the vault.
        """
        return pulumi.get(self, "backup_vault_arn")

    @property
    @pulumi.getter(name="backupVaultName")
    def backup_vault_name(self) -> pulumi.Output[str]:
        """
        Name of the backup vault to add a lock configuration for.
        """
        return pulumi.get(self, "backup_vault_name")

    @property
    @pulumi.getter(name="changeableForDays")
    def changeable_for_days(self) -> pulumi.Output[Optional[int]]:
        """
        The number of days before the lock date.
        """
        return pulumi.get(self, "changeable_for_days")

    @property
    @pulumi.getter(name="maxRetentionDays")
    def max_retention_days(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum retention period that the vault retains its recovery points.
        """
        return pulumi.get(self, "max_retention_days")

    @property
    @pulumi.getter(name="minRetentionDays")
    def min_retention_days(self) -> pulumi.Output[Optional[int]]:
        """
        The minimum retention period that the vault retains its recovery points.
        """
        return pulumi.get(self, "min_retention_days")

