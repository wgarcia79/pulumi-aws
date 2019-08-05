# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ConfigurationSet(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the configuration set
    """
    def __init__(__self__, resource_name, opts=None, name=None, __name__=None, __opts__=None):
        """
        Provides an SES configuration set resource
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the configuration set

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_configuration_set.html.markdown.
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

        __props__['name'] = name

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="aws:ses/confgurationSet:ConfgurationSet")])
        opts = alias_opts if opts is None else opts.merge(alias_opts)
        super(ConfigurationSet, __self__).__init__(
            'aws:ses/configurationSet:ConfigurationSet',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

