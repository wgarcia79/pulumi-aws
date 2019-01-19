# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class OriginAccessIdentity(pulumi.CustomResource):
    caller_reference: pulumi.Output[str]
    """
    Internal value used by CloudFront to allow future
    updates to the origin access identity.
    """
    cloudfront_access_identity_path: pulumi.Output[str]
    """
    A shortcut to the full path for the
    origin access identity to use in CloudFront, see below.
    """
    comment: pulumi.Output[str]
    """
    An optional comment for the origin access identity.
    """
    etag: pulumi.Output[str]
    """
    The current version of the origin access identity's information.
    For example: `E2QWRUHAPOMQZL`.
    """
    iam_arn: pulumi.Output[str]
    """
    A pre-generated ARN for use in S3 bucket policies (see below).
    Example: `arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity
    E2QWRUHAPOMQZL`.
    """
    s3_canonical_user_id: pulumi.Output[str]
    """
    The Amazon S3 canonical user ID for the origin
    access identity, which you use when giving the origin access identity read
    permission to an object in Amazon S3.
    """
    def __init__(__self__, __name__, __opts__=None, comment=None):
        """
        Creates an Amazon CloudFront origin access identity.
        
        For information about CloudFront distributions, see the
        [Amazon CloudFront Developer Guide][1]. For more information on generating
        origin access identities, see
        [Using an Origin Access Identity to Restrict Access to Your Amazon S3 Content][2].
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] comment: An optional comment for the origin access identity.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['comment'] = comment

        __props__['caller_reference'] = None
        __props__['cloudfront_access_identity_path'] = None
        __props__['etag'] = None
        __props__['iam_arn'] = None
        __props__['s3_canonical_user_id'] = None

        super(OriginAccessIdentity, __self__).__init__(
            'aws:cloudfront/originAccessIdentity:OriginAccessIdentity',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

