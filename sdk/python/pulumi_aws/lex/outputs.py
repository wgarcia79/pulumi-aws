# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'SlotTypeEnumerationValue',
    'GetSlotTypeEnumerationValueResult',
]

@pulumi.output_type
class SlotTypeEnumerationValue(dict):
    def __init__(__self__, *,
                 value: str,
                 synonyms: Optional[Sequence[str]] = None):
        """
        :param str value: The value of the slot type.
        :param Sequence[str] synonyms: Additional values related to the slot type value.
        """
        pulumi.set(__self__, "value", value)
        if synonyms is not None:
            pulumi.set(__self__, "synonyms", synonyms)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of the slot type.
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter
    def synonyms(self) -> Optional[Sequence[str]]:
        """
        Additional values related to the slot type value.
        """
        return pulumi.get(self, "synonyms")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetSlotTypeEnumerationValueResult(dict):
    def __init__(__self__, *,
                 synonyms: Sequence[str],
                 value: str):
        pulumi.set(__self__, "synonyms", synonyms)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def synonyms(self) -> Sequence[str]:
        return pulumi.get(self, "synonyms")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


