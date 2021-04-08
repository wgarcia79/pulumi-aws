# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'CanaryRunConfig',
    'CanarySchedule',
    'CanaryTimeline',
    'CanaryVpcConfig',
]

@pulumi.output_type
class CanaryRunConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "activeTracing":
            suggest = "active_tracing"
        elif key == "memoryInMb":
            suggest = "memory_in_mb"
        elif key == "timeoutInSeconds":
            suggest = "timeout_in_seconds"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CanaryRunConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CanaryRunConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CanaryRunConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 active_tracing: Optional[bool] = None,
                 memory_in_mb: Optional[int] = None,
                 timeout_in_seconds: Optional[int] = None):
        """
        :param bool active_tracing: Whether this canary is to use active AWS X-Ray tracing when it runs. You can enable active tracing only for canaries that use version syn-nodejs-2.0 or later for their canary runtime.
        :param int memory_in_mb: Maximum amount of memory available to the canary while it is running, in MB. The value you specify must be a multiple of 64.
        :param int timeout_in_seconds: Number of seconds the canary is allowed to run before it must stop. If you omit this field, the frequency of the canary is used, up to a maximum of 840 (14 minutes).
        """
        if active_tracing is not None:
            pulumi.set(__self__, "active_tracing", active_tracing)
        if memory_in_mb is not None:
            pulumi.set(__self__, "memory_in_mb", memory_in_mb)
        if timeout_in_seconds is not None:
            pulumi.set(__self__, "timeout_in_seconds", timeout_in_seconds)

    @property
    @pulumi.getter(name="activeTracing")
    def active_tracing(self) -> Optional[bool]:
        """
        Whether this canary is to use active AWS X-Ray tracing when it runs. You can enable active tracing only for canaries that use version syn-nodejs-2.0 or later for their canary runtime.
        """
        return pulumi.get(self, "active_tracing")

    @property
    @pulumi.getter(name="memoryInMb")
    def memory_in_mb(self) -> Optional[int]:
        """
        Maximum amount of memory available to the canary while it is running, in MB. The value you specify must be a multiple of 64.
        """
        return pulumi.get(self, "memory_in_mb")

    @property
    @pulumi.getter(name="timeoutInSeconds")
    def timeout_in_seconds(self) -> Optional[int]:
        """
        Number of seconds the canary is allowed to run before it must stop. If you omit this field, the frequency of the canary is used, up to a maximum of 840 (14 minutes).
        """
        return pulumi.get(self, "timeout_in_seconds")


@pulumi.output_type
class CanarySchedule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "durationInSeconds":
            suggest = "duration_in_seconds"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CanarySchedule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CanarySchedule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CanarySchedule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 expression: str,
                 duration_in_seconds: Optional[int] = None):
        """
        :param str expression: Rate expression that defines how often the canary is to run. The syntax is rate(number unit). unit can be minute, minutes, or hour.
        :param int duration_in_seconds: Duration in seconds, for the canary to continue making regular runs according to the schedule in the Expression value.
        """
        pulumi.set(__self__, "expression", expression)
        if duration_in_seconds is not None:
            pulumi.set(__self__, "duration_in_seconds", duration_in_seconds)

    @property
    @pulumi.getter
    def expression(self) -> str:
        """
        Rate expression that defines how often the canary is to run. The syntax is rate(number unit). unit can be minute, minutes, or hour.
        """
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter(name="durationInSeconds")
    def duration_in_seconds(self) -> Optional[int]:
        """
        Duration in seconds, for the canary to continue making regular runs according to the schedule in the Expression value.
        """
        return pulumi.get(self, "duration_in_seconds")


@pulumi.output_type
class CanaryTimeline(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "lastModified":
            suggest = "last_modified"
        elif key == "lastStarted":
            suggest = "last_started"
        elif key == "lastStopped":
            suggest = "last_stopped"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CanaryTimeline. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CanaryTimeline.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CanaryTimeline.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created: Optional[str] = None,
                 last_modified: Optional[str] = None,
                 last_started: Optional[str] = None,
                 last_stopped: Optional[str] = None):
        """
        :param str created: Date and time the canary was created.
        :param str last_modified: Date and time the canary was most recently modified.
        :param str last_started: Date and time that the canary's most recent run started.
        :param str last_stopped: Date and time that the canary's most recent run ended.
        """
        if created is not None:
            pulumi.set(__self__, "created", created)
        if last_modified is not None:
            pulumi.set(__self__, "last_modified", last_modified)
        if last_started is not None:
            pulumi.set(__self__, "last_started", last_started)
        if last_stopped is not None:
            pulumi.set(__self__, "last_stopped", last_stopped)

    @property
    @pulumi.getter
    def created(self) -> Optional[str]:
        """
        Date and time the canary was created.
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> Optional[str]:
        """
        Date and time the canary was most recently modified.
        """
        return pulumi.get(self, "last_modified")

    @property
    @pulumi.getter(name="lastStarted")
    def last_started(self) -> Optional[str]:
        """
        Date and time that the canary's most recent run started.
        """
        return pulumi.get(self, "last_started")

    @property
    @pulumi.getter(name="lastStopped")
    def last_stopped(self) -> Optional[str]:
        """
        Date and time that the canary's most recent run ended.
        """
        return pulumi.get(self, "last_stopped")


@pulumi.output_type
class CanaryVpcConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "securityGroupIds":
            suggest = "security_group_ids"
        elif key == "subnetIds":
            suggest = "subnet_ids"
        elif key == "vpcId":
            suggest = "vpc_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CanaryVpcConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CanaryVpcConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CanaryVpcConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 security_group_ids: Optional[Sequence[str]] = None,
                 subnet_ids: Optional[Sequence[str]] = None,
                 vpc_id: Optional[str] = None):
        """
        :param Sequence[str] security_group_ids: IDs of the security groups for this canary.
        :param Sequence[str] subnet_ids: IDs of the subnets where this canary is to run.
        :param str vpc_id: ID of the VPC where this canary is to run.
        """
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[Sequence[str]]:
        """
        IDs of the security groups for this canary.
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[Sequence[str]]:
        """
        IDs of the subnets where this canary is to run.
        """
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        ID of the VPC where this canary is to run.
        """
        return pulumi.get(self, "vpc_id")


