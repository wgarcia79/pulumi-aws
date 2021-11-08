# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VaultPolicyArgs', 'VaultPolicy']

@pulumi.input_type
class VaultPolicyArgs:
    def __init__(__self__, *,
                 backup_vault_name: pulumi.Input[str],
                 policy: pulumi.Input[str]):
        """
        The set of arguments for constructing a VaultPolicy resource.
        :param pulumi.Input[str] backup_vault_name: Name of the backup vault to add policy for.
        :param pulumi.Input[str] policy: The backup vault access policy document in JSON format.
        """
        pulumi.set(__self__, "backup_vault_name", backup_vault_name)
        pulumi.set(__self__, "policy", policy)

    @property
    @pulumi.getter(name="backupVaultName")
    def backup_vault_name(self) -> pulumi.Input[str]:
        """
        Name of the backup vault to add policy for.
        """
        return pulumi.get(self, "backup_vault_name")

    @backup_vault_name.setter
    def backup_vault_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "backup_vault_name", value)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Input[str]:
        """
        The backup vault access policy document in JSON format.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy", value)


@pulumi.input_type
class _VaultPolicyState:
    def __init__(__self__, *,
                 backup_vault_arn: Optional[pulumi.Input[str]] = None,
                 backup_vault_name: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VaultPolicy resources.
        :param pulumi.Input[str] backup_vault_arn: The ARN of the vault.
        :param pulumi.Input[str] backup_vault_name: Name of the backup vault to add policy for.
        :param pulumi.Input[str] policy: The backup vault access policy document in JSON format.
        """
        if backup_vault_arn is not None:
            pulumi.set(__self__, "backup_vault_arn", backup_vault_arn)
        if backup_vault_name is not None:
            pulumi.set(__self__, "backup_vault_name", backup_vault_name)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)

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
        Name of the backup vault to add policy for.
        """
        return pulumi.get(self, "backup_vault_name")

    @backup_vault_name.setter
    def backup_vault_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_vault_name", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        The backup vault access policy document in JSON format.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)


class VaultPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_vault_name: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an AWS Backup vault policy resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_vault = aws.backup.Vault("exampleVault")
        example_vault_policy = aws.backup.VaultPolicy("exampleVaultPolicy",
            backup_vault_name=example_vault.name,
            policy=example_vault.arn.apply(lambda arn: f\"\"\"{{
          "Version": "2012-10-17",
          "Id": "default",
          "Statement": [
            {{
              "Sid": "default",
              "Effect": "Allow",
              "Principal": {{
                "AWS": "*"
              }},
              "Action": [
        		"backup:DescribeBackupVault",
        		"backup:DeleteBackupVault",
        		"backup:PutBackupVaultAccessPolicy",
        		"backup:DeleteBackupVaultAccessPolicy",
        		"backup:GetBackupVaultAccessPolicy",
        		"backup:StartBackupJob",
        		"backup:GetBackupVaultNotifications",
        		"backup:PutBackupVaultNotifications"
              ],
              "Resource": "{arn}"
            }}
          ]
        }}
        \"\"\"))
        ```

        ## Import

        Backup vault policy can be imported using the `name`, e.g.,

        ```sh
         $ pulumi import aws:backup/vaultPolicy:VaultPolicy test TestVault
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_vault_name: Name of the backup vault to add policy for.
        :param pulumi.Input[str] policy: The backup vault access policy document in JSON format.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VaultPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AWS Backup vault policy resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_vault = aws.backup.Vault("exampleVault")
        example_vault_policy = aws.backup.VaultPolicy("exampleVaultPolicy",
            backup_vault_name=example_vault.name,
            policy=example_vault.arn.apply(lambda arn: f\"\"\"{{
          "Version": "2012-10-17",
          "Id": "default",
          "Statement": [
            {{
              "Sid": "default",
              "Effect": "Allow",
              "Principal": {{
                "AWS": "*"
              }},
              "Action": [
        		"backup:DescribeBackupVault",
        		"backup:DeleteBackupVault",
        		"backup:PutBackupVaultAccessPolicy",
        		"backup:DeleteBackupVaultAccessPolicy",
        		"backup:GetBackupVaultAccessPolicy",
        		"backup:StartBackupJob",
        		"backup:GetBackupVaultNotifications",
        		"backup:PutBackupVaultNotifications"
              ],
              "Resource": "{arn}"
            }}
          ]
        }}
        \"\"\"))
        ```

        ## Import

        Backup vault policy can be imported using the `name`, e.g.,

        ```sh
         $ pulumi import aws:backup/vaultPolicy:VaultPolicy test TestVault
        ```

        :param str resource_name: The name of the resource.
        :param VaultPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VaultPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_vault_name: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
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
            __props__ = VaultPolicyArgs.__new__(VaultPolicyArgs)

            if backup_vault_name is None and not opts.urn:
                raise TypeError("Missing required property 'backup_vault_name'")
            __props__.__dict__["backup_vault_name"] = backup_vault_name
            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__.__dict__["policy"] = policy
            __props__.__dict__["backup_vault_arn"] = None
        super(VaultPolicy, __self__).__init__(
            'aws:backup/vaultPolicy:VaultPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_vault_arn: Optional[pulumi.Input[str]] = None,
            backup_vault_name: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[str]] = None) -> 'VaultPolicy':
        """
        Get an existing VaultPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_vault_arn: The ARN of the vault.
        :param pulumi.Input[str] backup_vault_name: Name of the backup vault to add policy for.
        :param pulumi.Input[str] policy: The backup vault access policy document in JSON format.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VaultPolicyState.__new__(_VaultPolicyState)

        __props__.__dict__["backup_vault_arn"] = backup_vault_arn
        __props__.__dict__["backup_vault_name"] = backup_vault_name
        __props__.__dict__["policy"] = policy
        return VaultPolicy(resource_name, opts=opts, __props__=__props__)

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
        Name of the backup vault to add policy for.
        """
        return pulumi.get(self, "backup_vault_name")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        The backup vault access policy document in JSON format.
        """
        return pulumi.get(self, "policy")

