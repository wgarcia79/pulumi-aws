# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetVpnAttachmentResult:
    """
    A collection of values returned by getVpnAttachment.
    """
    def __init__(__self__, id=None, tags=None, transit_gateway_id=None, vpn_connection_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        Key-value tags for the EC2 Transit Gateway VPN Attachment
        """
        if transit_gateway_id and not isinstance(transit_gateway_id, str):
            raise TypeError("Expected argument 'transit_gateway_id' to be a str")
        __self__.transit_gateway_id = transit_gateway_id
        if vpn_connection_id and not isinstance(vpn_connection_id, str):
            raise TypeError("Expected argument 'vpn_connection_id' to be a str")
        __self__.vpn_connection_id = vpn_connection_id
class AwaitableGetVpnAttachmentResult(GetVpnAttachmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpnAttachmentResult(
            id=self.id,
            tags=self.tags,
            transit_gateway_id=self.transit_gateway_id,
            vpn_connection_id=self.vpn_connection_id)

def get_vpn_attachment(tags=None,transit_gateway_id=None,vpn_connection_id=None,opts=None):
    """
    Get information on an EC2 Transit Gateway VPN Attachment.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ec2_transit_gateway_vpn_attachment.html.markdown.


    :param str transit_gateway_id: Identifier of the EC2 Transit Gateway.
    :param str vpn_connection_id: Identifier of the EC2 VPN Connection.
    """
    __args__ = dict()


    __args__['tags'] = tags
    __args__['transitGatewayId'] = transit_gateway_id
    __args__['vpnConnectionId'] = vpn_connection_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2transitgateway/getVpnAttachment:getVpnAttachment', __args__, opts=opts).value

    return AwaitableGetVpnAttachmentResult(
        id=__ret__.get('id'),
        tags=__ret__.get('tags'),
        transit_gateway_id=__ret__.get('transitGatewayId'),
        vpn_connection_id=__ret__.get('vpnConnectionId'))
