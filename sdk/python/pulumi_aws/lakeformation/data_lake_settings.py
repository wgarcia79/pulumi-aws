# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DataLakeSettingsArgs', 'DataLakeSettings']

@pulumi.input_type
class DataLakeSettingsArgs:
    def __init__(__self__, *,
                 admins: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 create_database_default_permissions: Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeSettingsCreateDatabaseDefaultPermissionArgs']]]] = None,
                 create_table_default_permissions: Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeSettingsCreateTableDefaultPermissionArgs']]]] = None,
                 trusted_resource_owners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a DataLakeSettings resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] admins: Set of ARNs of AWS Lake Formation principals (IAM users or roles).
        :param pulumi.Input[str] catalog_id: Identifier for the Data Catalog. By default, the account ID.
        :param pulumi.Input[Sequence[pulumi.Input['DataLakeSettingsCreateDatabaseDefaultPermissionArgs']]] create_database_default_permissions: Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input['DataLakeSettingsCreateTableDefaultPermissionArgs']]] create_table_default_permissions: Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] trusted_resource_owners: List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
        """
        if admins is not None:
            pulumi.set(__self__, "admins", admins)
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)
        if create_database_default_permissions is not None:
            pulumi.set(__self__, "create_database_default_permissions", create_database_default_permissions)
        if create_table_default_permissions is not None:
            pulumi.set(__self__, "create_table_default_permissions", create_table_default_permissions)
        if trusted_resource_owners is not None:
            pulumi.set(__self__, "trusted_resource_owners", trusted_resource_owners)

    @property
    @pulumi.getter
    def admins(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of ARNs of AWS Lake Formation principals (IAM users or roles).
        """
        return pulumi.get(self, "admins")

    @admins.setter
    def admins(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "admins", value)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier for the Data Catalog. By default, the account ID.
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter(name="createDatabaseDefaultPermissions")
    def create_database_default_permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeSettingsCreateDatabaseDefaultPermissionArgs']]]]:
        """
        Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
        """
        return pulumi.get(self, "create_database_default_permissions")

    @create_database_default_permissions.setter
    def create_database_default_permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeSettingsCreateDatabaseDefaultPermissionArgs']]]]):
        pulumi.set(self, "create_database_default_permissions", value)

    @property
    @pulumi.getter(name="createTableDefaultPermissions")
    def create_table_default_permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeSettingsCreateTableDefaultPermissionArgs']]]]:
        """
        Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
        """
        return pulumi.get(self, "create_table_default_permissions")

    @create_table_default_permissions.setter
    def create_table_default_permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeSettingsCreateTableDefaultPermissionArgs']]]]):
        pulumi.set(self, "create_table_default_permissions", value)

    @property
    @pulumi.getter(name="trustedResourceOwners")
    def trusted_resource_owners(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
        """
        return pulumi.get(self, "trusted_resource_owners")

    @trusted_resource_owners.setter
    def trusted_resource_owners(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "trusted_resource_owners", value)


@pulumi.input_type
class _DataLakeSettingsState:
    def __init__(__self__, *,
                 admins: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 create_database_default_permissions: Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeSettingsCreateDatabaseDefaultPermissionArgs']]]] = None,
                 create_table_default_permissions: Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeSettingsCreateTableDefaultPermissionArgs']]]] = None,
                 trusted_resource_owners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering DataLakeSettings resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] admins: Set of ARNs of AWS Lake Formation principals (IAM users or roles).
        :param pulumi.Input[str] catalog_id: Identifier for the Data Catalog. By default, the account ID.
        :param pulumi.Input[Sequence[pulumi.Input['DataLakeSettingsCreateDatabaseDefaultPermissionArgs']]] create_database_default_permissions: Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input['DataLakeSettingsCreateTableDefaultPermissionArgs']]] create_table_default_permissions: Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] trusted_resource_owners: List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
        """
        if admins is not None:
            pulumi.set(__self__, "admins", admins)
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)
        if create_database_default_permissions is not None:
            pulumi.set(__self__, "create_database_default_permissions", create_database_default_permissions)
        if create_table_default_permissions is not None:
            pulumi.set(__self__, "create_table_default_permissions", create_table_default_permissions)
        if trusted_resource_owners is not None:
            pulumi.set(__self__, "trusted_resource_owners", trusted_resource_owners)

    @property
    @pulumi.getter
    def admins(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of ARNs of AWS Lake Formation principals (IAM users or roles).
        """
        return pulumi.get(self, "admins")

    @admins.setter
    def admins(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "admins", value)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier for the Data Catalog. By default, the account ID.
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter(name="createDatabaseDefaultPermissions")
    def create_database_default_permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeSettingsCreateDatabaseDefaultPermissionArgs']]]]:
        """
        Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
        """
        return pulumi.get(self, "create_database_default_permissions")

    @create_database_default_permissions.setter
    def create_database_default_permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeSettingsCreateDatabaseDefaultPermissionArgs']]]]):
        pulumi.set(self, "create_database_default_permissions", value)

    @property
    @pulumi.getter(name="createTableDefaultPermissions")
    def create_table_default_permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeSettingsCreateTableDefaultPermissionArgs']]]]:
        """
        Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
        """
        return pulumi.get(self, "create_table_default_permissions")

    @create_table_default_permissions.setter
    def create_table_default_permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeSettingsCreateTableDefaultPermissionArgs']]]]):
        pulumi.set(self, "create_table_default_permissions", value)

    @property
    @pulumi.getter(name="trustedResourceOwners")
    def trusted_resource_owners(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
        """
        return pulumi.get(self, "trusted_resource_owners")

    @trusted_resource_owners.setter
    def trusted_resource_owners(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "trusted_resource_owners", value)


class DataLakeSettings(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admins: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 create_database_default_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataLakeSettingsCreateDatabaseDefaultPermissionArgs']]]]] = None,
                 create_table_default_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataLakeSettingsCreateTableDefaultPermissionArgs']]]]] = None,
                 trusted_resource_owners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages Lake Formation principals designated as data lake administrators and lists of principal permission entries for default create database and default create table permissions.

        > **NOTE:** Lake Formation introduces fine-grained access control for data in your data lake. Part of the changes include the `IAMAllowedPrincipals` principal in order to make Lake Formation backwards compatible with existing IAM and Glue permissions. For more information, see [Changing the Default Security Settings for Your Data Lake](https://docs.aws.amazon.com/lake-formation/latest/dg/change-settings.html) and [Upgrading AWS Glue Data Permissions to the AWS Lake Formation Model](https://docs.aws.amazon.com/lake-formation/latest/dg/upgrade-glue-lake-formation.html).

        ## Example Usage
        ### Data Lake Admins

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.lakeformation.DataLakeSettings("example", admins=[
            aws_iam_user["test"]["arn"],
            aws_iam_role["test"]["arn"],
        ])
        ```
        ### Create Default Permissions

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.lakeformation.DataLakeSettings("example",
            admins=[
                aws_iam_user["test"]["arn"],
                aws_iam_role["test"]["arn"],
            ],
            create_database_default_permissions=[aws.lakeformation.DataLakeSettingsCreateDatabaseDefaultPermissionArgs(
                permissions=[
                    "SELECT",
                    "ALTER",
                    "DROP",
                ],
                principal=aws_iam_user["test"]["arn"],
            )],
            create_table_default_permissions=[aws.lakeformation.DataLakeSettingsCreateTableDefaultPermissionArgs(
                permissions=["ALL"],
                principal=aws_iam_role["test"]["arn"],
            )])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] admins: Set of ARNs of AWS Lake Formation principals (IAM users or roles).
        :param pulumi.Input[str] catalog_id: Identifier for the Data Catalog. By default, the account ID.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataLakeSettingsCreateDatabaseDefaultPermissionArgs']]]] create_database_default_permissions: Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataLakeSettingsCreateTableDefaultPermissionArgs']]]] create_table_default_permissions: Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] trusted_resource_owners: List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DataLakeSettingsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Lake Formation principals designated as data lake administrators and lists of principal permission entries for default create database and default create table permissions.

        > **NOTE:** Lake Formation introduces fine-grained access control for data in your data lake. Part of the changes include the `IAMAllowedPrincipals` principal in order to make Lake Formation backwards compatible with existing IAM and Glue permissions. For more information, see [Changing the Default Security Settings for Your Data Lake](https://docs.aws.amazon.com/lake-formation/latest/dg/change-settings.html) and [Upgrading AWS Glue Data Permissions to the AWS Lake Formation Model](https://docs.aws.amazon.com/lake-formation/latest/dg/upgrade-glue-lake-formation.html).

        ## Example Usage
        ### Data Lake Admins

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.lakeformation.DataLakeSettings("example", admins=[
            aws_iam_user["test"]["arn"],
            aws_iam_role["test"]["arn"],
        ])
        ```
        ### Create Default Permissions

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.lakeformation.DataLakeSettings("example",
            admins=[
                aws_iam_user["test"]["arn"],
                aws_iam_role["test"]["arn"],
            ],
            create_database_default_permissions=[aws.lakeformation.DataLakeSettingsCreateDatabaseDefaultPermissionArgs(
                permissions=[
                    "SELECT",
                    "ALTER",
                    "DROP",
                ],
                principal=aws_iam_user["test"]["arn"],
            )],
            create_table_default_permissions=[aws.lakeformation.DataLakeSettingsCreateTableDefaultPermissionArgs(
                permissions=["ALL"],
                principal=aws_iam_role["test"]["arn"],
            )])
        ```

        :param str resource_name: The name of the resource.
        :param DataLakeSettingsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataLakeSettingsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admins: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 create_database_default_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataLakeSettingsCreateDatabaseDefaultPermissionArgs']]]]] = None,
                 create_table_default_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataLakeSettingsCreateTableDefaultPermissionArgs']]]]] = None,
                 trusted_resource_owners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DataLakeSettingsArgs.__new__(DataLakeSettingsArgs)

            __props__.__dict__["admins"] = admins
            __props__.__dict__["catalog_id"] = catalog_id
            __props__.__dict__["create_database_default_permissions"] = create_database_default_permissions
            __props__.__dict__["create_table_default_permissions"] = create_table_default_permissions
            __props__.__dict__["trusted_resource_owners"] = trusted_resource_owners
        super(DataLakeSettings, __self__).__init__(
            'aws:lakeformation/dataLakeSettings:DataLakeSettings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admins: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            catalog_id: Optional[pulumi.Input[str]] = None,
            create_database_default_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataLakeSettingsCreateDatabaseDefaultPermissionArgs']]]]] = None,
            create_table_default_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataLakeSettingsCreateTableDefaultPermissionArgs']]]]] = None,
            trusted_resource_owners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'DataLakeSettings':
        """
        Get an existing DataLakeSettings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] admins: Set of ARNs of AWS Lake Formation principals (IAM users or roles).
        :param pulumi.Input[str] catalog_id: Identifier for the Data Catalog. By default, the account ID.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataLakeSettingsCreateDatabaseDefaultPermissionArgs']]]] create_database_default_permissions: Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataLakeSettingsCreateTableDefaultPermissionArgs']]]] create_table_default_permissions: Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] trusted_resource_owners: List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DataLakeSettingsState.__new__(_DataLakeSettingsState)

        __props__.__dict__["admins"] = admins
        __props__.__dict__["catalog_id"] = catalog_id
        __props__.__dict__["create_database_default_permissions"] = create_database_default_permissions
        __props__.__dict__["create_table_default_permissions"] = create_table_default_permissions
        __props__.__dict__["trusted_resource_owners"] = trusted_resource_owners
        return DataLakeSettings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def admins(self) -> pulumi.Output[Sequence[str]]:
        """
        Set of ARNs of AWS Lake Formation principals (IAM users or roles).
        """
        return pulumi.get(self, "admins")

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Output[Optional[str]]:
        """
        Identifier for the Data Catalog. By default, the account ID.
        """
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter(name="createDatabaseDefaultPermissions")
    def create_database_default_permissions(self) -> pulumi.Output[Sequence['outputs.DataLakeSettingsCreateDatabaseDefaultPermission']]:
        """
        Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
        """
        return pulumi.get(self, "create_database_default_permissions")

    @property
    @pulumi.getter(name="createTableDefaultPermissions")
    def create_table_default_permissions(self) -> pulumi.Output[Sequence['outputs.DataLakeSettingsCreateTableDefaultPermission']]:
        """
        Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
        """
        return pulumi.get(self, "create_table_default_permissions")

    @property
    @pulumi.getter(name="trustedResourceOwners")
    def trusted_resource_owners(self) -> pulumi.Output[Sequence[str]]:
        """
        List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
        """
        return pulumi.get(self, "trusted_resource_owners")

