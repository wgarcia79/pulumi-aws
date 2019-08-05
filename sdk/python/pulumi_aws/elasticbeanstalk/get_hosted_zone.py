# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetHostedZoneResult:
    """
    A collection of values returned by getHostedZone.
    """
    def __init__(__self__, region=None, id=None):
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        """
        The region of the hosted zone.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return self

    __iter__ = __await__

def get_hosted_zone(region=None,opts=None):
    """
    Use this data source to get the ID of an [elastic beanstalk hosted zone](http://docs.aws.amazon.com/general/latest/gr/rande.html#elasticbeanstalk_region).

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elastic_beanstalk_hosted_zone.html.markdown.
    """
    __args__ = dict()

    __args__['region'] = region
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:elasticbeanstalk/getHostedZone:getHostedZone', __args__, opts=opts).value

    return GetHostedZoneResult(
        region=__ret__.get('region'),
        id=__ret__.get('id'))
