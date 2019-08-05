# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Attachment(pulumi.CustomResource):
    elb: pulumi.Output[str]
    """
    The name of the ELB.
    """
    instance: pulumi.Output[str]
    """
    Instance ID to place in the ELB pool.
    """
    def __init__(__self__, resource_name, opts=None, elb=None, instance=None, __name__=None, __opts__=None):
        """
        Attaches an EC2 instance to an Elastic Load Balancer (ELB). For attaching resources with Application Load Balancer (ALB) or Network Load Balancer (NLB), see the [`aws_lb_target_group_attachment` resource](https://www.terraform.io/docs/providers/aws/r/lb_target_group_attachment.html).
        
        > **NOTE on ELB Instances and ELB Attachments:** This provider currently provides
        both a standalone ELB Attachment resource (describing an instance attached to
        an ELB), and an Elastic Load Balancer resource with
        `instances` defined in-line. At this time you cannot use an ELB with in-line
        instances in conjunction with an ELB Attachment resource. Doing so will cause a
        conflict and will overwrite attachments.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] elb: The name of the ELB.
        :param pulumi.Input[str] instance: Instance ID to place in the ELB pool.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/elb_attachment_legacy.html.markdown.
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

        if elb is None:
            raise TypeError("Missing required property 'elb'")
        __props__['elb'] = elb

        if instance is None:
            raise TypeError("Missing required property 'instance'")
        __props__['instance'] = instance

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Attachment, __self__).__init__(
            'aws:elasticloadbalancing/attachment:Attachment',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

