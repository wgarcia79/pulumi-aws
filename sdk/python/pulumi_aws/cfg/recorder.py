# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Recorder(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the recorder. Defaults to `default`. Changing it recreates the resource.
    """
    recording_group: pulumi.Output[dict]
    """
    Recording group - see below.
    """
    role_arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the IAM role.
    used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account.
    See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.
    """
    def __init__(__self__, __name__, __opts__=None, name=None, recording_group=None, role_arn=None):
        """
        Provides an AWS Config Configuration Recorder. Please note that this resource **does not start** the created recorder automatically.
        
        > **Note:** _Starting_ the Configuration Recorder requires a [delivery channel](https://www.terraform.io/docs/providers/aws/r/config_delivery_channel.html) (while delivery channel creation requires Configuration Recorder). This is why [`aws_config_configuration_recorder_status`](https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder_status.html) is a separate resource.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] name: The name of the recorder. Defaults to `default`. Changing it recreates the resource.
        :param pulumi.Input[dict] recording_group: Recording group - see below.
        :param pulumi.Input[str] role_arn: Amazon Resource Name (ARN) of the IAM role.
               used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account.
               See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['name'] = name

        __props__['recording_group'] = recording_group

        if not role_arn:
            raise TypeError('Missing required property role_arn')
        __props__['role_arn'] = role_arn

        super(Recorder, __self__).__init__(
            'aws:cfg/recorder:Recorder',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

