# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'ClusterNodeArgs',
    'ClusterServerSideEncryptionArgs',
    'ParameterGroupParameterArgs',
]

@pulumi.input_type
class ClusterNodeArgs:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[float]] = None):
        """
        :param pulumi.Input[float] port: The port used by the configuration endpoint
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[float]]:
        """
        The port used by the configuration endpoint
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "port", value)


@pulumi.input_type
class ClusterServerSideEncryptionArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] enabled: Whether to enable encryption at rest. Defaults to `false`.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable encryption at rest. Defaults to `false`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class ParameterGroupParameterArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] name: The name of the parameter.
        :param pulumi.Input[str] value: The value for the parameter.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the parameter.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the parameter.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


