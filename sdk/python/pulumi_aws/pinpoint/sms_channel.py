# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SmsChannelArgs', 'SmsChannel']

@pulumi.input_type
class SmsChannelArgs:
    def __init__(__self__, *,
                 application_id: pulumi.Input[str],
                 enabled: Optional[pulumi.Input[bool]] = None,
                 sender_id: Optional[pulumi.Input[str]] = None,
                 short_code: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SmsChannel resource.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[bool] enabled: Whether the channel is enabled or disabled. Defaults to `true`.
        :param pulumi.Input[str] sender_id: Sender identifier of your messages.
        :param pulumi.Input[str] short_code: The Short Code registered with the phone provider.
        """
        pulumi.set(__self__, "application_id", application_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if sender_id is not None:
            pulumi.set(__self__, "sender_id", sender_id)
        if short_code is not None:
            pulumi.set(__self__, "short_code", short_code)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Input[str]:
        """
        The application ID.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the channel is enabled or disabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="senderId")
    def sender_id(self) -> Optional[pulumi.Input[str]]:
        """
        Sender identifier of your messages.
        """
        return pulumi.get(self, "sender_id")

    @sender_id.setter
    def sender_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sender_id", value)

    @property
    @pulumi.getter(name="shortCode")
    def short_code(self) -> Optional[pulumi.Input[str]]:
        """
        The Short Code registered with the phone provider.
        """
        return pulumi.get(self, "short_code")

    @short_code.setter
    def short_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "short_code", value)


@pulumi.input_type
class _SmsChannelState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 promotional_messages_per_second: Optional[pulumi.Input[int]] = None,
                 sender_id: Optional[pulumi.Input[str]] = None,
                 short_code: Optional[pulumi.Input[str]] = None,
                 transactional_messages_per_second: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering SmsChannel resources.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[bool] enabled: Whether the channel is enabled or disabled. Defaults to `true`.
        :param pulumi.Input[int] promotional_messages_per_second: Promotional messages per second that can be sent.
        :param pulumi.Input[str] sender_id: Sender identifier of your messages.
        :param pulumi.Input[str] short_code: The Short Code registered with the phone provider.
        :param pulumi.Input[int] transactional_messages_per_second: Transactional messages per second that can be sent.
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if promotional_messages_per_second is not None:
            pulumi.set(__self__, "promotional_messages_per_second", promotional_messages_per_second)
        if sender_id is not None:
            pulumi.set(__self__, "sender_id", sender_id)
        if short_code is not None:
            pulumi.set(__self__, "short_code", short_code)
        if transactional_messages_per_second is not None:
            pulumi.set(__self__, "transactional_messages_per_second", transactional_messages_per_second)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        The application ID.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the channel is enabled or disabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="promotionalMessagesPerSecond")
    def promotional_messages_per_second(self) -> Optional[pulumi.Input[int]]:
        """
        Promotional messages per second that can be sent.
        """
        return pulumi.get(self, "promotional_messages_per_second")

    @promotional_messages_per_second.setter
    def promotional_messages_per_second(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "promotional_messages_per_second", value)

    @property
    @pulumi.getter(name="senderId")
    def sender_id(self) -> Optional[pulumi.Input[str]]:
        """
        Sender identifier of your messages.
        """
        return pulumi.get(self, "sender_id")

    @sender_id.setter
    def sender_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sender_id", value)

    @property
    @pulumi.getter(name="shortCode")
    def short_code(self) -> Optional[pulumi.Input[str]]:
        """
        The Short Code registered with the phone provider.
        """
        return pulumi.get(self, "short_code")

    @short_code.setter
    def short_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "short_code", value)

    @property
    @pulumi.getter(name="transactionalMessagesPerSecond")
    def transactional_messages_per_second(self) -> Optional[pulumi.Input[int]]:
        """
        Transactional messages per second that can be sent.
        """
        return pulumi.get(self, "transactional_messages_per_second")

    @transactional_messages_per_second.setter
    def transactional_messages_per_second(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "transactional_messages_per_second", value)


class SmsChannel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 sender_id: Optional[pulumi.Input[str]] = None,
                 short_code: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Pinpoint SMS Channel resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        app = aws.pinpoint.App("app")
        sms = aws.pinpoint.SmsChannel("sms", application_id=app.application_id)
        ```

        ## Import

        Pinpoint SMS Channel can be imported using the `application-id`, e.g.

        ```sh
         $ pulumi import aws:pinpoint/smsChannel:SmsChannel sms application-id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[bool] enabled: Whether the channel is enabled or disabled. Defaults to `true`.
        :param pulumi.Input[str] sender_id: Sender identifier of your messages.
        :param pulumi.Input[str] short_code: The Short Code registered with the phone provider.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SmsChannelArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Pinpoint SMS Channel resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        app = aws.pinpoint.App("app")
        sms = aws.pinpoint.SmsChannel("sms", application_id=app.application_id)
        ```

        ## Import

        Pinpoint SMS Channel can be imported using the `application-id`, e.g.

        ```sh
         $ pulumi import aws:pinpoint/smsChannel:SmsChannel sms application-id
        ```

        :param str resource_name: The name of the resource.
        :param SmsChannelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SmsChannelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 sender_id: Optional[pulumi.Input[str]] = None,
                 short_code: Optional[pulumi.Input[str]] = None,
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
            __props__ = SmsChannelArgs.__new__(SmsChannelArgs)

            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__.__dict__["application_id"] = application_id
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["sender_id"] = sender_id
            __props__.__dict__["short_code"] = short_code
            __props__.__dict__["promotional_messages_per_second"] = None
            __props__.__dict__["transactional_messages_per_second"] = None
        super(SmsChannel, __self__).__init__(
            'aws:pinpoint/smsChannel:SmsChannel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            promotional_messages_per_second: Optional[pulumi.Input[int]] = None,
            sender_id: Optional[pulumi.Input[str]] = None,
            short_code: Optional[pulumi.Input[str]] = None,
            transactional_messages_per_second: Optional[pulumi.Input[int]] = None) -> 'SmsChannel':
        """
        Get an existing SmsChannel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[bool] enabled: Whether the channel is enabled or disabled. Defaults to `true`.
        :param pulumi.Input[int] promotional_messages_per_second: Promotional messages per second that can be sent.
        :param pulumi.Input[str] sender_id: Sender identifier of your messages.
        :param pulumi.Input[str] short_code: The Short Code registered with the phone provider.
        :param pulumi.Input[int] transactional_messages_per_second: Transactional messages per second that can be sent.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SmsChannelState.__new__(_SmsChannelState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["promotional_messages_per_second"] = promotional_messages_per_second
        __props__.__dict__["sender_id"] = sender_id
        __props__.__dict__["short_code"] = short_code
        __props__.__dict__["transactional_messages_per_second"] = transactional_messages_per_second
        return SmsChannel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        The application ID.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the channel is enabled or disabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="promotionalMessagesPerSecond")
    def promotional_messages_per_second(self) -> pulumi.Output[int]:
        """
        Promotional messages per second that can be sent.
        """
        return pulumi.get(self, "promotional_messages_per_second")

    @property
    @pulumi.getter(name="senderId")
    def sender_id(self) -> pulumi.Output[Optional[str]]:
        """
        Sender identifier of your messages.
        """
        return pulumi.get(self, "sender_id")

    @property
    @pulumi.getter(name="shortCode")
    def short_code(self) -> pulumi.Output[Optional[str]]:
        """
        The Short Code registered with the phone provider.
        """
        return pulumi.get(self, "short_code")

    @property
    @pulumi.getter(name="transactionalMessagesPerSecond")
    def transactional_messages_per_second(self) -> pulumi.Output[int]:
        """
        Transactional messages per second that can be sent.
        """
        return pulumi.get(self, "transactional_messages_per_second")

