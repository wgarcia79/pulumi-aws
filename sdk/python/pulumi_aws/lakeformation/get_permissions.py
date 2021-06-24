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

__all__ = [
    'GetPermissionsResult',
    'AwaitableGetPermissionsResult',
    'get_permissions',
]

@pulumi.output_type
class GetPermissionsResult:
    """
    A collection of values returned by getPermissions.
    """
    def __init__(__self__, catalog_id=None, catalog_resource=None, data_location=None, database=None, id=None, permissions=None, permissions_with_grant_options=None, principal=None, table=None, table_with_columns=None):
        if catalog_id and not isinstance(catalog_id, str):
            raise TypeError("Expected argument 'catalog_id' to be a str")
        pulumi.set(__self__, "catalog_id", catalog_id)
        if catalog_resource and not isinstance(catalog_resource, bool):
            raise TypeError("Expected argument 'catalog_resource' to be a bool")
        pulumi.set(__self__, "catalog_resource", catalog_resource)
        if data_location and not isinstance(data_location, dict):
            raise TypeError("Expected argument 'data_location' to be a dict")
        pulumi.set(__self__, "data_location", data_location)
        if database and not isinstance(database, dict):
            raise TypeError("Expected argument 'database' to be a dict")
        pulumi.set(__self__, "database", database)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if permissions and not isinstance(permissions, list):
            raise TypeError("Expected argument 'permissions' to be a list")
        pulumi.set(__self__, "permissions", permissions)
        if permissions_with_grant_options and not isinstance(permissions_with_grant_options, list):
            raise TypeError("Expected argument 'permissions_with_grant_options' to be a list")
        pulumi.set(__self__, "permissions_with_grant_options", permissions_with_grant_options)
        if principal and not isinstance(principal, str):
            raise TypeError("Expected argument 'principal' to be a str")
        pulumi.set(__self__, "principal", principal)
        if table and not isinstance(table, dict):
            raise TypeError("Expected argument 'table' to be a dict")
        pulumi.set(__self__, "table", table)
        if table_with_columns and not isinstance(table_with_columns, dict):
            raise TypeError("Expected argument 'table_with_columns' to be a dict")
        pulumi.set(__self__, "table_with_columns", table_with_columns)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[str]:
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter(name="catalogResource")
    def catalog_resource(self) -> Optional[bool]:
        return pulumi.get(self, "catalog_resource")

    @property
    @pulumi.getter(name="dataLocation")
    def data_location(self) -> 'outputs.GetPermissionsDataLocationResult':
        return pulumi.get(self, "data_location")

    @property
    @pulumi.getter
    def database(self) -> 'outputs.GetPermissionsDatabaseResult':
        return pulumi.get(self, "database")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def permissions(self) -> Sequence[str]:
        """
        List of permissions granted to the principal. For details on permissions, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="permissionsWithGrantOptions")
    def permissions_with_grant_options(self) -> Sequence[str]:
        """
        Subset of `permissions` which the principal can pass.
        """
        return pulumi.get(self, "permissions_with_grant_options")

    @property
    @pulumi.getter
    def principal(self) -> str:
        return pulumi.get(self, "principal")

    @property
    @pulumi.getter
    def table(self) -> 'outputs.GetPermissionsTableResult':
        return pulumi.get(self, "table")

    @property
    @pulumi.getter(name="tableWithColumns")
    def table_with_columns(self) -> 'outputs.GetPermissionsTableWithColumnsResult':
        return pulumi.get(self, "table_with_columns")


class AwaitableGetPermissionsResult(GetPermissionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPermissionsResult(
            catalog_id=self.catalog_id,
            catalog_resource=self.catalog_resource,
            data_location=self.data_location,
            database=self.database,
            id=self.id,
            permissions=self.permissions,
            permissions_with_grant_options=self.permissions_with_grant_options,
            principal=self.principal,
            table=self.table,
            table_with_columns=self.table_with_columns)


def get_permissions(catalog_id: Optional[str] = None,
                    catalog_resource: Optional[bool] = None,
                    data_location: Optional[pulumi.InputType['GetPermissionsDataLocationArgs']] = None,
                    database: Optional[pulumi.InputType['GetPermissionsDatabaseArgs']] = None,
                    principal: Optional[str] = None,
                    table: Optional[pulumi.InputType['GetPermissionsTableArgs']] = None,
                    table_with_columns: Optional[pulumi.InputType['GetPermissionsTableWithColumnsArgs']] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPermissionsResult:
    """
    Get permissions for a principal to access metadata in the Data Catalog and data organized in underlying data storage such as Amazon S3. Permissions are granted to a principal, in a Data Catalog, relative to a Lake Formation resource, which includes the Data Catalog, databases, and tables. For more information, see [Security and Access Control to Metadata and Data in Lake Formation](https://docs.aws.amazon.com/lake-formation/latest/dg/security-data-access.html).

    > **NOTE:** This data source deals with explicitly granted permissions. Lake Formation grants implicit permissions to data lake administrators, database creators, and table creators. For more information, see [Implicit Lake Formation Permissions](https://docs.aws.amazon.com/lake-formation/latest/dg/implicit-permissions.html).

    ## Example Usage
    ### Permissions For A Lake Formation S3 Resource

    ```python
    import pulumi
    import pulumi_aws as aws

    test = aws.lakeformation.get_permissions(principal=aws_iam_role["workflow_role"]["arn"],
        data_location=aws.lakeformation.GetPermissionsDataLocationArgs(
            arn=aws_lakeformation_resource["test"]["arn"],
        ))
    ```
    ### Permissions For A Glue Catalog Database

    ```python
    import pulumi
    import pulumi_aws as aws

    test = aws.lakeformation.get_permissions(principal=aws_iam_role["workflow_role"]["arn"],
        database=aws.lakeformation.GetPermissionsDatabaseArgs(
            name=aws_glue_catalog_database["test"]["name"],
            catalog_id="110376042874",
        ))
    ```


    :param str catalog_id: Identifier for the Data Catalog. By default, it is the account ID of the caller.
    :param bool catalog_resource: Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
    :param pulumi.InputType['GetPermissionsDataLocationArgs'] data_location: Configuration block for a data location resource. Detailed below.
    :param pulumi.InputType['GetPermissionsDatabaseArgs'] database: Configuration block for a database resource. Detailed below.
    :param str principal: Principal to be granted the permissions on the resource. Supported principals are IAM users or IAM roles.
    :param pulumi.InputType['GetPermissionsTableArgs'] table: Configuration block for a table resource. Detailed below.
    :param pulumi.InputType['GetPermissionsTableWithColumnsArgs'] table_with_columns: Configuration block for a table with columns resource. Detailed below.
    """
    __args__ = dict()
    __args__['catalogId'] = catalog_id
    __args__['catalogResource'] = catalog_resource
    __args__['dataLocation'] = data_location
    __args__['database'] = database
    __args__['principal'] = principal
    __args__['table'] = table
    __args__['tableWithColumns'] = table_with_columns
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:lakeformation/getPermissions:getPermissions', __args__, opts=opts, typ=GetPermissionsResult).value

    return AwaitableGetPermissionsResult(
        catalog_id=__ret__.catalog_id,
        catalog_resource=__ret__.catalog_resource,
        data_location=__ret__.data_location,
        database=__ret__.database,
        id=__ret__.id,
        permissions=__ret__.permissions,
        permissions_with_grant_options=__ret__.permissions_with_grant_options,
        principal=__ret__.principal,
        table=__ret__.table,
        table_with_columns=__ret__.table_with_columns)
