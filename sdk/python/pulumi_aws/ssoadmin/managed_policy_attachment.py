# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ManagedPolicyAttachmentArgs', 'ManagedPolicyAttachment']

@pulumi.input_type
class ManagedPolicyAttachmentArgs:
    def __init__(__self__, *,
                 instance_arn: pulumi.Input[str],
                 managed_policy_arn: pulumi.Input[str],
                 permission_set_arn: pulumi.Input[str]):
        """
        The set of arguments for constructing a ManagedPolicyAttachment resource.
        :param pulumi.Input[str] instance_arn: The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
        :param pulumi.Input[str] managed_policy_arn: The IAM managed policy Amazon Resource Name (ARN) to be attached to the Permission Set.
        :param pulumi.Input[str] permission_set_arn: The Amazon Resource Name (ARN) of the Permission Set.
        """
        pulumi.set(__self__, "instance_arn", instance_arn)
        pulumi.set(__self__, "managed_policy_arn", managed_policy_arn)
        pulumi.set(__self__, "permission_set_arn", permission_set_arn)

    @property
    @pulumi.getter(name="instanceArn")
    def instance_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
        """
        return pulumi.get(self, "instance_arn")

    @instance_arn.setter
    def instance_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_arn", value)

    @property
    @pulumi.getter(name="managedPolicyArn")
    def managed_policy_arn(self) -> pulumi.Input[str]:
        """
        The IAM managed policy Amazon Resource Name (ARN) to be attached to the Permission Set.
        """
        return pulumi.get(self, "managed_policy_arn")

    @managed_policy_arn.setter
    def managed_policy_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "managed_policy_arn", value)

    @property
    @pulumi.getter(name="permissionSetArn")
    def permission_set_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of the Permission Set.
        """
        return pulumi.get(self, "permission_set_arn")

    @permission_set_arn.setter
    def permission_set_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "permission_set_arn", value)


@pulumi.input_type
class _ManagedPolicyAttachmentState:
    def __init__(__self__, *,
                 instance_arn: Optional[pulumi.Input[str]] = None,
                 managed_policy_arn: Optional[pulumi.Input[str]] = None,
                 managed_policy_name: Optional[pulumi.Input[str]] = None,
                 permission_set_arn: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ManagedPolicyAttachment resources.
        :param pulumi.Input[str] instance_arn: The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
        :param pulumi.Input[str] managed_policy_arn: The IAM managed policy Amazon Resource Name (ARN) to be attached to the Permission Set.
        :param pulumi.Input[str] managed_policy_name: The name of the IAM Managed Policy.
        :param pulumi.Input[str] permission_set_arn: The Amazon Resource Name (ARN) of the Permission Set.
        """
        if instance_arn is not None:
            pulumi.set(__self__, "instance_arn", instance_arn)
        if managed_policy_arn is not None:
            pulumi.set(__self__, "managed_policy_arn", managed_policy_arn)
        if managed_policy_name is not None:
            pulumi.set(__self__, "managed_policy_name", managed_policy_name)
        if permission_set_arn is not None:
            pulumi.set(__self__, "permission_set_arn", permission_set_arn)

    @property
    @pulumi.getter(name="instanceArn")
    def instance_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
        """
        return pulumi.get(self, "instance_arn")

    @instance_arn.setter
    def instance_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_arn", value)

    @property
    @pulumi.getter(name="managedPolicyArn")
    def managed_policy_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The IAM managed policy Amazon Resource Name (ARN) to be attached to the Permission Set.
        """
        return pulumi.get(self, "managed_policy_arn")

    @managed_policy_arn.setter
    def managed_policy_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "managed_policy_arn", value)

    @property
    @pulumi.getter(name="managedPolicyName")
    def managed_policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IAM Managed Policy.
        """
        return pulumi.get(self, "managed_policy_name")

    @managed_policy_name.setter
    def managed_policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "managed_policy_name", value)

    @property
    @pulumi.getter(name="permissionSetArn")
    def permission_set_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the Permission Set.
        """
        return pulumi.get(self, "permission_set_arn")

    @permission_set_arn.setter
    def permission_set_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "permission_set_arn", value)


class ManagedPolicyAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_arn: Optional[pulumi.Input[str]] = None,
                 managed_policy_arn: Optional[pulumi.Input[str]] = None,
                 permission_set_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an IAM managed policy for a Single Sign-On (SSO) Permission Set resource

        > **NOTE:** Creating this resource will automatically [Provision the Permission Set](https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ProvisionPermissionSet.html) to apply the corresponding updates to all assigned accounts.

        ## Import

        SSO Managed Policy Attachments can be imported using the `managed_policy_arn`, `permission_set_arn`, and `instance_arn` separated by a comma (`,`) e.g.

        ```sh
         $ pulumi import aws:ssoadmin/managedPolicyAttachment:ManagedPolicyAttachment example arn:aws:iam::aws:policy/AlexaForBusinessDeviceSetup,arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_arn: The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
        :param pulumi.Input[str] managed_policy_arn: The IAM managed policy Amazon Resource Name (ARN) to be attached to the Permission Set.
        :param pulumi.Input[str] permission_set_arn: The Amazon Resource Name (ARN) of the Permission Set.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ManagedPolicyAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an IAM managed policy for a Single Sign-On (SSO) Permission Set resource

        > **NOTE:** Creating this resource will automatically [Provision the Permission Set](https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ProvisionPermissionSet.html) to apply the corresponding updates to all assigned accounts.

        ## Import

        SSO Managed Policy Attachments can be imported using the `managed_policy_arn`, `permission_set_arn`, and `instance_arn` separated by a comma (`,`) e.g.

        ```sh
         $ pulumi import aws:ssoadmin/managedPolicyAttachment:ManagedPolicyAttachment example arn:aws:iam::aws:policy/AlexaForBusinessDeviceSetup,arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
        ```

        :param str resource_name: The name of the resource.
        :param ManagedPolicyAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ManagedPolicyAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_arn: Optional[pulumi.Input[str]] = None,
                 managed_policy_arn: Optional[pulumi.Input[str]] = None,
                 permission_set_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ManagedPolicyAttachmentArgs.__new__(ManagedPolicyAttachmentArgs)

            if instance_arn is None and not opts.urn:
                raise TypeError("Missing required property 'instance_arn'")
            __props__.__dict__["instance_arn"] = instance_arn
            if managed_policy_arn is None and not opts.urn:
                raise TypeError("Missing required property 'managed_policy_arn'")
            __props__.__dict__["managed_policy_arn"] = managed_policy_arn
            if permission_set_arn is None and not opts.urn:
                raise TypeError("Missing required property 'permission_set_arn'")
            __props__.__dict__["permission_set_arn"] = permission_set_arn
            __props__.__dict__["managed_policy_name"] = None
        super(ManagedPolicyAttachment, __self__).__init__(
            'aws:ssoadmin/managedPolicyAttachment:ManagedPolicyAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_arn: Optional[pulumi.Input[str]] = None,
            managed_policy_arn: Optional[pulumi.Input[str]] = None,
            managed_policy_name: Optional[pulumi.Input[str]] = None,
            permission_set_arn: Optional[pulumi.Input[str]] = None) -> 'ManagedPolicyAttachment':
        """
        Get an existing ManagedPolicyAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_arn: The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
        :param pulumi.Input[str] managed_policy_arn: The IAM managed policy Amazon Resource Name (ARN) to be attached to the Permission Set.
        :param pulumi.Input[str] managed_policy_name: The name of the IAM Managed Policy.
        :param pulumi.Input[str] permission_set_arn: The Amazon Resource Name (ARN) of the Permission Set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ManagedPolicyAttachmentState.__new__(_ManagedPolicyAttachmentState)

        __props__.__dict__["instance_arn"] = instance_arn
        __props__.__dict__["managed_policy_arn"] = managed_policy_arn
        __props__.__dict__["managed_policy_name"] = managed_policy_name
        __props__.__dict__["permission_set_arn"] = permission_set_arn
        return ManagedPolicyAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceArn")
    def instance_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
        """
        return pulumi.get(self, "instance_arn")

    @property
    @pulumi.getter(name="managedPolicyArn")
    def managed_policy_arn(self) -> pulumi.Output[str]:
        """
        The IAM managed policy Amazon Resource Name (ARN) to be attached to the Permission Set.
        """
        return pulumi.get(self, "managed_policy_arn")

    @property
    @pulumi.getter(name="managedPolicyName")
    def managed_policy_name(self) -> pulumi.Output[str]:
        """
        The name of the IAM Managed Policy.
        """
        return pulumi.get(self, "managed_policy_name")

    @property
    @pulumi.getter(name="permissionSetArn")
    def permission_set_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the Permission Set.
        """
        return pulumi.get(self, "permission_set_arn")

