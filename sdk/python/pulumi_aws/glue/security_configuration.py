# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class SecurityConfiguration(pulumi.CustomResource):
    encryption_configuration: pulumi.Output[dict]
    """
    Configuration block containing encryption configuration. Detailed below.
    """
    name: pulumi.Output[str]
    """
    Name of the security configuration.
    """
    def __init__(__self__, __name__, __opts__=None, encryption_configuration=None, name=None):
        """
        Manages a Glue Security Configuration.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[dict] encryption_configuration: Configuration block containing encryption configuration. Detailed below.
        :param pulumi.Input[str] name: Name of the security configuration.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not encryption_configuration:
            raise TypeError('Missing required property encryption_configuration')
        __props__['encryption_configuration'] = encryption_configuration

        __props__['name'] = name

        super(SecurityConfiguration, __self__).__init__(
            'aws:glue/securityConfiguration:SecurityConfiguration',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

