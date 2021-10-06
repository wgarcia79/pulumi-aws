# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetParametersByPathResult',
    'AwaitableGetParametersByPathResult',
    'get_parameters_by_path',
    'get_parameters_by_path_output',
]

@pulumi.output_type
class GetParametersByPathResult:
    """
    A collection of values returned by getParametersByPath.
    """
    def __init__(__self__, arns=None, id=None, names=None, path=None, types=None, values=None, with_decryption=None):
        if arns and not isinstance(arns, list):
            raise TypeError("Expected argument 'arns' to be a list")
        pulumi.set(__self__, "arns", arns)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if types and not isinstance(types, list):
            raise TypeError("Expected argument 'types' to be a list")
        pulumi.set(__self__, "types", types)
        if values and not isinstance(values, list):
            raise TypeError("Expected argument 'values' to be a list")
        pulumi.set(__self__, "values", values)
        if with_decryption and not isinstance(with_decryption, bool):
            raise TypeError("Expected argument 'with_decryption' to be a bool")
        pulumi.set(__self__, "with_decryption", with_decryption)

    @property
    @pulumi.getter
    def arns(self) -> Sequence[str]:
        return pulumi.get(self, "arns")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter
    def path(self) -> str:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def types(self) -> Sequence[str]:
        return pulumi.get(self, "types")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        return pulumi.get(self, "values")

    @property
    @pulumi.getter(name="withDecryption")
    def with_decryption(self) -> Optional[bool]:
        return pulumi.get(self, "with_decryption")


class AwaitableGetParametersByPathResult(GetParametersByPathResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetParametersByPathResult(
            arns=self.arns,
            id=self.id,
            names=self.names,
            path=self.path,
            types=self.types,
            values=self.values,
            with_decryption=self.with_decryption)


def get_parameters_by_path(path: Optional[str] = None,
                           with_decryption: Optional[bool] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetParametersByPathResult:
    """
    Use this data source to access information about an existing resource.

    :param str path: The prefix path of the parameter.
    :param bool with_decryption: Whether to return decrypted `SecureString` value. Defaults to `true`.
    """
    __args__ = dict()
    __args__['path'] = path
    __args__['withDecryption'] = with_decryption
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ssm/getParametersByPath:getParametersByPath', __args__, opts=opts, typ=GetParametersByPathResult).value

    return AwaitableGetParametersByPathResult(
        arns=__ret__.arns,
        id=__ret__.id,
        names=__ret__.names,
        path=__ret__.path,
        types=__ret__.types,
        values=__ret__.values,
        with_decryption=__ret__.with_decryption)


@_utilities.lift_output_func(get_parameters_by_path)
def get_parameters_by_path_output(path: Optional[pulumi.Input[str]] = None,
                                  with_decryption: Optional[pulumi.Input[Optional[bool]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetParametersByPathResult]:
    """
    Use this data source to access information about an existing resource.

    :param str path: The prefix path of the parameter.
    :param bool with_decryption: Whether to return decrypted `SecureString` value. Defaults to `true`.
    """
    ...
