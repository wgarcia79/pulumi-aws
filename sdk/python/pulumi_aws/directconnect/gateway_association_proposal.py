# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GatewayAssociationProposal(pulumi.CustomResource):
    allowed_prefixes: pulumi.Output[list]
    """
    VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
    """
    dx_gateway_id: pulumi.Output[str]
    """
    Direct Connect Gateway identifier.
    """
    dx_gateway_owner_account_id: pulumi.Output[str]
    """
    AWS Account identifier of the Direct Connect Gateway.
    """
    vpn_gateway_id: pulumi.Output[str]
    """
    Virtual Gateway identifier to associate with the Direct Connect Gateway.
    """
    def __init__(__self__, resource_name, opts=None, allowed_prefixes=None, dx_gateway_id=None, dx_gateway_owner_account_id=None, vpn_gateway_id=None, __name__=None, __opts__=None):
        """
        Manages a Direct Connect Gateway Association Proposal, typically for enabling cross-account associations. For single account associations, see the [`aws_dx_gateway_association` resource](https://www.terraform.io/docs/providers/aws/r/dx_gateway_association.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allowed_prefixes: VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
        :param pulumi.Input[str] dx_gateway_id: Direct Connect Gateway identifier.
        :param pulumi.Input[str] dx_gateway_owner_account_id: AWS Account identifier of the Direct Connect Gateway.
        :param pulumi.Input[str] vpn_gateway_id: Virtual Gateway identifier to associate with the Direct Connect Gateway.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['allowed_prefixes'] = allowed_prefixes

        if dx_gateway_id is None:
            raise TypeError("Missing required property 'dx_gateway_id'")
        __props__['dx_gateway_id'] = dx_gateway_id

        if dx_gateway_owner_account_id is None:
            raise TypeError("Missing required property 'dx_gateway_owner_account_id'")
        __props__['dx_gateway_owner_account_id'] = dx_gateway_owner_account_id

        if vpn_gateway_id is None:
            raise TypeError("Missing required property 'vpn_gateway_id'")
        __props__['vpn_gateway_id'] = vpn_gateway_id

        super(GatewayAssociationProposal, __self__).__init__(
            'aws:directconnect/gatewayAssociationProposal:GatewayAssociationProposal',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

