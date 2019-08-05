# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetDefaultKmsKeyResult:
    """
    A collection of values returned by getDefaultKmsKey.
    """
    def __init__(__self__, key_arn=None, id=None):
        if key_arn and not isinstance(key_arn, str):
            raise TypeError("Expected argument 'key_arn' to be a str")
        __self__.key_arn = key_arn
        """
        Amazon Resource Name (ARN) of the default KMS key uses to encrypt an EBS volume in this region when no key is specified in an API call that creates the volume and encryption by default is enabled.
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

def get_default_kms_key(opts=None):
    """
    Use this data source to get the default EBS encryption KMS key in the current region.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ebs_default_kms_key.html.markdown.
    """
    __args__ = dict()

    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ebs/getDefaultKmsKey:getDefaultKmsKey', __args__, opts=opts).value

    return GetDefaultKmsKeyResult(
        key_arn=__ret__.get('keyArn'),
        id=__ret__.get('id'))
