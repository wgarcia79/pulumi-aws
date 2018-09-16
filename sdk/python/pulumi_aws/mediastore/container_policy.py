# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class ContainerPolicy(pulumi.CustomResource):
    """
    Provides a MediaStore Container Policy.
    """
    def __init__(__self__, __name__, __opts__=None, container_name=None, policy=None):
        """Create a ContainerPolicy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not container_name:
            raise TypeError('Missing required property container_name')
        elif not isinstance(container_name, basestring):
            raise TypeError('Expected property container_name to be a basestring')
        __self__.container_name = container_name
        """
        The name of the container.
        """
        __props__['containerName'] = container_name

        if not policy:
            raise TypeError('Missing required property policy')
        elif not isinstance(policy, basestring):
            raise TypeError('Expected property policy to be a basestring')
        __self__.policy = policy
        """
        The contents of the policy.
        """
        __props__['policy'] = policy

        super(ContainerPolicy, __self__).__init__(
            'aws:mediastore/containerPolicy:ContainerPolicy',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'containerName' in outs:
            self.container_name = outs['containerName']
        if 'policy' in outs:
            self.policy = outs['policy']