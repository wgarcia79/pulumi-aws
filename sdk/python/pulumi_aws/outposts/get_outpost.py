# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetOutpostResult',
    'AwaitableGetOutpostResult',
    'get_outpost',
]

@pulumi.output_type
class GetOutpostResult:
    """
    A collection of values returned by getOutpost.
    """
    def __init__(__self__, arn=None, availability_zone=None, availability_zone_id=None, description=None, id=None, name=None, owner_id=None, site_id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if availability_zone_id and not isinstance(availability_zone_id, str):
            raise TypeError("Expected argument 'availability_zone_id' to be a str")
        pulumi.set(__self__, "availability_zone_id", availability_zone_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        pulumi.set(__self__, "owner_id", owner_id)
        if site_id and not isinstance(site_id, str):
            raise TypeError("Expected argument 'site_id' to be a str")
        pulumi.set(__self__, "site_id", site_id)

    @property
    @pulumi.getter
    def arn(self) -> str:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        Availability Zone name.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="availabilityZoneId")
    def availability_zone_id(self) -> str:
        """
        Availability Zone identifier.
        """
        return pulumi.get(self, "availability_zone_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> str:
        """
        AWS Account identifier of the Outpost owner.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="siteId")
    def site_id(self) -> str:
        """
        Site identifier.
        """
        return pulumi.get(self, "site_id")


class AwaitableGetOutpostResult(GetOutpostResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOutpostResult(
            arn=self.arn,
            availability_zone=self.availability_zone,
            availability_zone_id=self.availability_zone_id,
            description=self.description,
            id=self.id,
            name=self.name,
            owner_id=self.owner_id,
            site_id=self.site_id)


def get_outpost(arn: Optional[str] = None,
                id: Optional[str] = None,
                name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOutpostResult:
    """
    Provides details about an Outposts Outpost.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.outposts.get_outpost(name="example")
    ```


    :param str arn: Amazon Resource Name (ARN).
    :param str id: Identifier of the Outpost.
    :param str name: Name of the Outpost.
    """
    __args__ = dict()
    __args__['arn'] = arn
    __args__['id'] = id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:outposts/getOutpost:getOutpost', __args__, opts=opts, typ=GetOutpostResult).value

    return AwaitableGetOutpostResult(
        arn=__ret__.arn,
        availability_zone=__ret__.availability_zone,
        availability_zone_id=__ret__.availability_zone_id,
        description=__ret__.description,
        id=__ret__.id,
        name=__ret__.name,
        owner_id=__ret__.owner_id,
        site_id=__ret__.site_id)
