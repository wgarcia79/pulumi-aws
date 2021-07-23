# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
]

@pulumi.output_type
class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, access_string=None, engine=None, id=None, no_password_required=None, passwords=None, user_id=None, user_name=None):
        if access_string and not isinstance(access_string, str):
            raise TypeError("Expected argument 'access_string' to be a str")
        pulumi.set(__self__, "access_string", access_string)
        if engine and not isinstance(engine, str):
            raise TypeError("Expected argument 'engine' to be a str")
        pulumi.set(__self__, "engine", engine)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if no_password_required and not isinstance(no_password_required, bool):
            raise TypeError("Expected argument 'no_password_required' to be a bool")
        pulumi.set(__self__, "no_password_required", no_password_required)
        if passwords and not isinstance(passwords, list):
            raise TypeError("Expected argument 'passwords' to be a list")
        pulumi.set(__self__, "passwords", passwords)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="accessString")
    def access_string(self) -> Optional[str]:
        """
        A string for what access a user possesses within the associated ElastiCache replication groups or clusters.
        """
        return pulumi.get(self, "access_string")

    @property
    @pulumi.getter
    def engine(self) -> Optional[str]:
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="noPasswordRequired")
    def no_password_required(self) -> Optional[bool]:
        return pulumi.get(self, "no_password_required")

    @property
    @pulumi.getter
    def passwords(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "passwords")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        """
        The identifier for the user.
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[str]:
        """
        The user name of the user.
        """
        return pulumi.get(self, "user_name")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            access_string=self.access_string,
            engine=self.engine,
            id=self.id,
            no_password_required=self.no_password_required,
            passwords=self.passwords,
            user_id=self.user_id,
            user_name=self.user_name)


def get_user(access_string: Optional[str] = None,
             engine: Optional[str] = None,
             no_password_required: Optional[bool] = None,
             passwords: Optional[Sequence[str]] = None,
             user_id: Optional[str] = None,
             user_name: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    Use this data source to get information about an Elasticache User.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    bar = aws.elasticache.get_user(user_id="example")
    ```


    :param str access_string: A string for what access a user possesses within the associated ElastiCache replication groups or clusters.
    :param str user_id: The identifier for the user.
    :param str user_name: The user name of the user.
    """
    __args__ = dict()
    __args__['accessString'] = access_string
    __args__['engine'] = engine
    __args__['noPasswordRequired'] = no_password_required
    __args__['passwords'] = passwords
    __args__['userId'] = user_id
    __args__['userName'] = user_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:elasticache/getUser:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        access_string=__ret__.access_string,
        engine=__ret__.engine,
        id=__ret__.id,
        no_password_required=__ret__.no_password_required,
        passwords=__ret__.passwords,
        user_id=__ret__.user_id,
        user_name=__ret__.user_name)
