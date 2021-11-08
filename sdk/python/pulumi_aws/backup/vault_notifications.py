# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VaultNotificationsArgs', 'VaultNotifications']

@pulumi.input_type
class VaultNotificationsArgs:
    def __init__(__self__, *,
                 backup_vault_events: pulumi.Input[Sequence[pulumi.Input[str]]],
                 backup_vault_name: pulumi.Input[str],
                 sns_topic_arn: pulumi.Input[str]):
        """
        The set of arguments for constructing a VaultNotifications resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] backup_vault_events: An array of events that indicate the status of jobs to back up resources to the backup vault.
        :param pulumi.Input[str] backup_vault_name: Name of the backup vault to add notifications for.
        :param pulumi.Input[str] sns_topic_arn: The Amazon Resource Name (ARN) that specifies the topic for a backup vault’s events
        """
        pulumi.set(__self__, "backup_vault_events", backup_vault_events)
        pulumi.set(__self__, "backup_vault_name", backup_vault_name)
        pulumi.set(__self__, "sns_topic_arn", sns_topic_arn)

    @property
    @pulumi.getter(name="backupVaultEvents")
    def backup_vault_events(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        An array of events that indicate the status of jobs to back up resources to the backup vault.
        """
        return pulumi.get(self, "backup_vault_events")

    @backup_vault_events.setter
    def backup_vault_events(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "backup_vault_events", value)

    @property
    @pulumi.getter(name="backupVaultName")
    def backup_vault_name(self) -> pulumi.Input[str]:
        """
        Name of the backup vault to add notifications for.
        """
        return pulumi.get(self, "backup_vault_name")

    @backup_vault_name.setter
    def backup_vault_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "backup_vault_name", value)

    @property
    @pulumi.getter(name="snsTopicArn")
    def sns_topic_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) that specifies the topic for a backup vault’s events
        """
        return pulumi.get(self, "sns_topic_arn")

    @sns_topic_arn.setter
    def sns_topic_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "sns_topic_arn", value)


@pulumi.input_type
class _VaultNotificationsState:
    def __init__(__self__, *,
                 backup_vault_arn: Optional[pulumi.Input[str]] = None,
                 backup_vault_events: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 backup_vault_name: Optional[pulumi.Input[str]] = None,
                 sns_topic_arn: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VaultNotifications resources.
        :param pulumi.Input[str] backup_vault_arn: The ARN of the vault.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] backup_vault_events: An array of events that indicate the status of jobs to back up resources to the backup vault.
        :param pulumi.Input[str] backup_vault_name: Name of the backup vault to add notifications for.
        :param pulumi.Input[str] sns_topic_arn: The Amazon Resource Name (ARN) that specifies the topic for a backup vault’s events
        """
        if backup_vault_arn is not None:
            pulumi.set(__self__, "backup_vault_arn", backup_vault_arn)
        if backup_vault_events is not None:
            pulumi.set(__self__, "backup_vault_events", backup_vault_events)
        if backup_vault_name is not None:
            pulumi.set(__self__, "backup_vault_name", backup_vault_name)
        if sns_topic_arn is not None:
            pulumi.set(__self__, "sns_topic_arn", sns_topic_arn)

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
    @pulumi.getter(name="backupVaultEvents")
    def backup_vault_events(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of events that indicate the status of jobs to back up resources to the backup vault.
        """
        return pulumi.get(self, "backup_vault_events")

    @backup_vault_events.setter
    def backup_vault_events(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "backup_vault_events", value)

    @property
    @pulumi.getter(name="backupVaultName")
    def backup_vault_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the backup vault to add notifications for.
        """
        return pulumi.get(self, "backup_vault_name")

    @backup_vault_name.setter
    def backup_vault_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_vault_name", value)

    @property
    @pulumi.getter(name="snsTopicArn")
    def sns_topic_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) that specifies the topic for a backup vault’s events
        """
        return pulumi.get(self, "sns_topic_arn")

    @sns_topic_arn.setter
    def sns_topic_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sns_topic_arn", value)


class VaultNotifications(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_vault_events: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 backup_vault_name: Optional[pulumi.Input[str]] = None,
                 sns_topic_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an AWS Backup vault notifications resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test_topic = aws.sns.Topic("testTopic")
        test_policy_document = test_topic.arn.apply(lambda arn: aws.iam.get_policy_document(policy_id="__default_policy_ID",
            statements=[aws.iam.GetPolicyDocumentStatementArgs(
                actions=["SNS:Publish"],
                effect="Allow",
                principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                    type="Service",
                    identifiers=["backup.amazonaws.com"],
                )],
                resources=[arn],
                sid="__default_statement_ID",
            )]))
        test_topic_policy = aws.sns.TopicPolicy("testTopicPolicy",
            arn=test_topic.arn,
            policy=test_policy_document.json)
        test_vault_notifications = aws.backup.VaultNotifications("testVaultNotifications",
            backup_vault_name="example_backup_vault",
            sns_topic_arn=test_topic.arn,
            backup_vault_events=[
                "BACKUP_JOB_STARTED",
                "RESTORE_JOB_COMPLETED",
            ])
        ```

        ## Import

        Backup vault notifications can be imported using the `name`, e.g.,

        ```sh
         $ pulumi import aws:backup/vaultNotifications:VaultNotifications test TestVault
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] backup_vault_events: An array of events that indicate the status of jobs to back up resources to the backup vault.
        :param pulumi.Input[str] backup_vault_name: Name of the backup vault to add notifications for.
        :param pulumi.Input[str] sns_topic_arn: The Amazon Resource Name (ARN) that specifies the topic for a backup vault’s events
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VaultNotificationsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AWS Backup vault notifications resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test_topic = aws.sns.Topic("testTopic")
        test_policy_document = test_topic.arn.apply(lambda arn: aws.iam.get_policy_document(policy_id="__default_policy_ID",
            statements=[aws.iam.GetPolicyDocumentStatementArgs(
                actions=["SNS:Publish"],
                effect="Allow",
                principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                    type="Service",
                    identifiers=["backup.amazonaws.com"],
                )],
                resources=[arn],
                sid="__default_statement_ID",
            )]))
        test_topic_policy = aws.sns.TopicPolicy("testTopicPolicy",
            arn=test_topic.arn,
            policy=test_policy_document.json)
        test_vault_notifications = aws.backup.VaultNotifications("testVaultNotifications",
            backup_vault_name="example_backup_vault",
            sns_topic_arn=test_topic.arn,
            backup_vault_events=[
                "BACKUP_JOB_STARTED",
                "RESTORE_JOB_COMPLETED",
            ])
        ```

        ## Import

        Backup vault notifications can be imported using the `name`, e.g.,

        ```sh
         $ pulumi import aws:backup/vaultNotifications:VaultNotifications test TestVault
        ```

        :param str resource_name: The name of the resource.
        :param VaultNotificationsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VaultNotificationsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_vault_events: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 backup_vault_name: Optional[pulumi.Input[str]] = None,
                 sns_topic_arn: Optional[pulumi.Input[str]] = None,
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
            __props__ = VaultNotificationsArgs.__new__(VaultNotificationsArgs)

            if backup_vault_events is None and not opts.urn:
                raise TypeError("Missing required property 'backup_vault_events'")
            __props__.__dict__["backup_vault_events"] = backup_vault_events
            if backup_vault_name is None and not opts.urn:
                raise TypeError("Missing required property 'backup_vault_name'")
            __props__.__dict__["backup_vault_name"] = backup_vault_name
            if sns_topic_arn is None and not opts.urn:
                raise TypeError("Missing required property 'sns_topic_arn'")
            __props__.__dict__["sns_topic_arn"] = sns_topic_arn
            __props__.__dict__["backup_vault_arn"] = None
        super(VaultNotifications, __self__).__init__(
            'aws:backup/vaultNotifications:VaultNotifications',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_vault_arn: Optional[pulumi.Input[str]] = None,
            backup_vault_events: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            backup_vault_name: Optional[pulumi.Input[str]] = None,
            sns_topic_arn: Optional[pulumi.Input[str]] = None) -> 'VaultNotifications':
        """
        Get an existing VaultNotifications resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_vault_arn: The ARN of the vault.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] backup_vault_events: An array of events that indicate the status of jobs to back up resources to the backup vault.
        :param pulumi.Input[str] backup_vault_name: Name of the backup vault to add notifications for.
        :param pulumi.Input[str] sns_topic_arn: The Amazon Resource Name (ARN) that specifies the topic for a backup vault’s events
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VaultNotificationsState.__new__(_VaultNotificationsState)

        __props__.__dict__["backup_vault_arn"] = backup_vault_arn
        __props__.__dict__["backup_vault_events"] = backup_vault_events
        __props__.__dict__["backup_vault_name"] = backup_vault_name
        __props__.__dict__["sns_topic_arn"] = sns_topic_arn
        return VaultNotifications(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupVaultArn")
    def backup_vault_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the vault.
        """
        return pulumi.get(self, "backup_vault_arn")

    @property
    @pulumi.getter(name="backupVaultEvents")
    def backup_vault_events(self) -> pulumi.Output[Sequence[str]]:
        """
        An array of events that indicate the status of jobs to back up resources to the backup vault.
        """
        return pulumi.get(self, "backup_vault_events")

    @property
    @pulumi.getter(name="backupVaultName")
    def backup_vault_name(self) -> pulumi.Output[str]:
        """
        Name of the backup vault to add notifications for.
        """
        return pulumi.get(self, "backup_vault_name")

    @property
    @pulumi.getter(name="snsTopicArn")
    def sns_topic_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) that specifies the topic for a backup vault’s events
        """
        return pulumi.get(self, "sns_topic_arn")

