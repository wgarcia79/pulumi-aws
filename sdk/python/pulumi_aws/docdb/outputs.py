# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ClusterParameterGroupParameter',
]

@pulumi.output_type
class ClusterParameterGroupParameter(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "applyMethod":
            suggest = "apply_method"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterParameterGroupParameter. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterParameterGroupParameter.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterParameterGroupParameter.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 value: str,
                 apply_method: Optional[str] = None):
        """
        :param str name: The name of the documentDB parameter.
        :param str value: The value of the documentDB parameter.
        :param str apply_method: Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)
        if apply_method is not None:
            pulumi.set(__self__, "apply_method", apply_method)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the documentDB parameter.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of the documentDB parameter.
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="applyMethod")
    def apply_method(self) -> Optional[str]:
        """
        Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
        """
        return pulumi.get(self, "apply_method")


