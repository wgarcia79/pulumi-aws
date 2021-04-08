# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VpcEndpointConnectionNotificationArgs', 'VpcEndpointConnectionNotification']

@pulumi.input_type
class VpcEndpointConnectionNotificationArgs:
    def __init__(__self__, *,
                 connection_events: pulumi.Input[Sequence[pulumi.Input[str]]],
                 connection_notification_arn: pulumi.Input[str],
                 vpc_endpoint_id: Optional[pulumi.Input[str]] = None,
                 vpc_endpoint_service_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VpcEndpointConnectionNotification resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] connection_events: One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
        :param pulumi.Input[str] connection_notification_arn: The ARN of the SNS topic for the notifications.
        :param pulumi.Input[str] vpc_endpoint_id: The ID of the VPC Endpoint to receive notifications for.
        :param pulumi.Input[str] vpc_endpoint_service_id: The ID of the VPC Endpoint Service to receive notifications for.
        """
        pulumi.set(__self__, "connection_events", connection_events)
        pulumi.set(__self__, "connection_notification_arn", connection_notification_arn)
        if vpc_endpoint_id is not None:
            pulumi.set(__self__, "vpc_endpoint_id", vpc_endpoint_id)
        if vpc_endpoint_service_id is not None:
            pulumi.set(__self__, "vpc_endpoint_service_id", vpc_endpoint_service_id)

    @property
    @pulumi.getter(name="connectionEvents")
    def connection_events(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
        """
        return pulumi.get(self, "connection_events")

    @connection_events.setter
    def connection_events(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "connection_events", value)

    @property
    @pulumi.getter(name="connectionNotificationArn")
    def connection_notification_arn(self) -> pulumi.Input[str]:
        """
        The ARN of the SNS topic for the notifications.
        """
        return pulumi.get(self, "connection_notification_arn")

    @connection_notification_arn.setter
    def connection_notification_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "connection_notification_arn", value)

    @property
    @pulumi.getter(name="vpcEndpointId")
    def vpc_endpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPC Endpoint to receive notifications for.
        """
        return pulumi.get(self, "vpc_endpoint_id")

    @vpc_endpoint_id.setter
    def vpc_endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_endpoint_id", value)

    @property
    @pulumi.getter(name="vpcEndpointServiceId")
    def vpc_endpoint_service_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPC Endpoint Service to receive notifications for.
        """
        return pulumi.get(self, "vpc_endpoint_service_id")

    @vpc_endpoint_service_id.setter
    def vpc_endpoint_service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_endpoint_service_id", value)


@pulumi.input_type
class _VpcEndpointConnectionNotificationState:
    def __init__(__self__, *,
                 connection_events: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 connection_notification_arn: Optional[pulumi.Input[str]] = None,
                 notification_type: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 vpc_endpoint_id: Optional[pulumi.Input[str]] = None,
                 vpc_endpoint_service_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcEndpointConnectionNotification resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] connection_events: One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
        :param pulumi.Input[str] connection_notification_arn: The ARN of the SNS topic for the notifications.
        :param pulumi.Input[str] notification_type: The type of notification.
        :param pulumi.Input[str] state: The state of the notification.
        :param pulumi.Input[str] vpc_endpoint_id: The ID of the VPC Endpoint to receive notifications for.
        :param pulumi.Input[str] vpc_endpoint_service_id: The ID of the VPC Endpoint Service to receive notifications for.
        """
        if connection_events is not None:
            pulumi.set(__self__, "connection_events", connection_events)
        if connection_notification_arn is not None:
            pulumi.set(__self__, "connection_notification_arn", connection_notification_arn)
        if notification_type is not None:
            pulumi.set(__self__, "notification_type", notification_type)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if vpc_endpoint_id is not None:
            pulumi.set(__self__, "vpc_endpoint_id", vpc_endpoint_id)
        if vpc_endpoint_service_id is not None:
            pulumi.set(__self__, "vpc_endpoint_service_id", vpc_endpoint_service_id)

    @property
    @pulumi.getter(name="connectionEvents")
    def connection_events(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
        """
        return pulumi.get(self, "connection_events")

    @connection_events.setter
    def connection_events(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "connection_events", value)

    @property
    @pulumi.getter(name="connectionNotificationArn")
    def connection_notification_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the SNS topic for the notifications.
        """
        return pulumi.get(self, "connection_notification_arn")

    @connection_notification_arn.setter
    def connection_notification_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_notification_arn", value)

    @property
    @pulumi.getter(name="notificationType")
    def notification_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of notification.
        """
        return pulumi.get(self, "notification_type")

    @notification_type.setter
    def notification_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_type", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The state of the notification.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="vpcEndpointId")
    def vpc_endpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPC Endpoint to receive notifications for.
        """
        return pulumi.get(self, "vpc_endpoint_id")

    @vpc_endpoint_id.setter
    def vpc_endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_endpoint_id", value)

    @property
    @pulumi.getter(name="vpcEndpointServiceId")
    def vpc_endpoint_service_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPC Endpoint Service to receive notifications for.
        """
        return pulumi.get(self, "vpc_endpoint_service_id")

    @vpc_endpoint_service_id.setter
    def vpc_endpoint_service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_endpoint_service_id", value)


class VpcEndpointConnectionNotification(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_events: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 connection_notification_arn: Optional[pulumi.Input[str]] = None,
                 vpc_endpoint_id: Optional[pulumi.Input[str]] = None,
                 vpc_endpoint_service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a VPC Endpoint connection notification resource.
        Connection notifications notify subscribers of VPC Endpoint events.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        topic = aws.sns.Topic("topic", policy=\"\"\"{
            "Version":"2012-10-17",
            "Statement":[{
                "Effect": "Allow",
                "Principal": {
                    "Service": "vpce.amazonaws.com"
                },
                "Action": "SNS:Publish",
                "Resource": "arn:aws:sns:*:*:vpce-notification-topic"
            }]
        }
        \"\"\")
        foo_vpc_endpoint_service = aws.ec2.VpcEndpointService("fooVpcEndpointService",
            acceptance_required=False,
            network_load_balancer_arns=[aws_lb["test"]["arn"]])
        foo_vpc_endpoint_connection_notification = aws.ec2.VpcEndpointConnectionNotification("fooVpcEndpointConnectionNotification",
            vpc_endpoint_service_id=foo_vpc_endpoint_service.id,
            connection_notification_arn=topic.arn,
            connection_events=[
                "Accept",
                "Reject",
            ])
        ```

        ## Import

        VPC Endpoint connection notifications can be imported using the `VPC endpoint connection notification id`, e.g.

        ```sh
         $ pulumi import aws:ec2/vpcEndpointConnectionNotification:VpcEndpointConnectionNotification foo vpce-nfn-09e6ed3b4efba2263
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] connection_events: One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
        :param pulumi.Input[str] connection_notification_arn: The ARN of the SNS topic for the notifications.
        :param pulumi.Input[str] vpc_endpoint_id: The ID of the VPC Endpoint to receive notifications for.
        :param pulumi.Input[str] vpc_endpoint_service_id: The ID of the VPC Endpoint Service to receive notifications for.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcEndpointConnectionNotificationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a VPC Endpoint connection notification resource.
        Connection notifications notify subscribers of VPC Endpoint events.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        topic = aws.sns.Topic("topic", policy=\"\"\"{
            "Version":"2012-10-17",
            "Statement":[{
                "Effect": "Allow",
                "Principal": {
                    "Service": "vpce.amazonaws.com"
                },
                "Action": "SNS:Publish",
                "Resource": "arn:aws:sns:*:*:vpce-notification-topic"
            }]
        }
        \"\"\")
        foo_vpc_endpoint_service = aws.ec2.VpcEndpointService("fooVpcEndpointService",
            acceptance_required=False,
            network_load_balancer_arns=[aws_lb["test"]["arn"]])
        foo_vpc_endpoint_connection_notification = aws.ec2.VpcEndpointConnectionNotification("fooVpcEndpointConnectionNotification",
            vpc_endpoint_service_id=foo_vpc_endpoint_service.id,
            connection_notification_arn=topic.arn,
            connection_events=[
                "Accept",
                "Reject",
            ])
        ```

        ## Import

        VPC Endpoint connection notifications can be imported using the `VPC endpoint connection notification id`, e.g.

        ```sh
         $ pulumi import aws:ec2/vpcEndpointConnectionNotification:VpcEndpointConnectionNotification foo vpce-nfn-09e6ed3b4efba2263
        ```

        :param str resource_name: The name of the resource.
        :param VpcEndpointConnectionNotificationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcEndpointConnectionNotificationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_events: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 connection_notification_arn: Optional[pulumi.Input[str]] = None,
                 vpc_endpoint_id: Optional[pulumi.Input[str]] = None,
                 vpc_endpoint_service_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = VpcEndpointConnectionNotificationArgs.__new__(VpcEndpointConnectionNotificationArgs)

            if connection_events is None and not opts.urn:
                raise TypeError("Missing required property 'connection_events'")
            __props__.__dict__["connection_events"] = connection_events
            if connection_notification_arn is None and not opts.urn:
                raise TypeError("Missing required property 'connection_notification_arn'")
            __props__.__dict__["connection_notification_arn"] = connection_notification_arn
            __props__.__dict__["vpc_endpoint_id"] = vpc_endpoint_id
            __props__.__dict__["vpc_endpoint_service_id"] = vpc_endpoint_service_id
            __props__.__dict__["notification_type"] = None
            __props__.__dict__["state"] = None
        super(VpcEndpointConnectionNotification, __self__).__init__(
            'aws:ec2/vpcEndpointConnectionNotification:VpcEndpointConnectionNotification',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            connection_events: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            connection_notification_arn: Optional[pulumi.Input[str]] = None,
            notification_type: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            vpc_endpoint_id: Optional[pulumi.Input[str]] = None,
            vpc_endpoint_service_id: Optional[pulumi.Input[str]] = None) -> 'VpcEndpointConnectionNotification':
        """
        Get an existing VpcEndpointConnectionNotification resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] connection_events: One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
        :param pulumi.Input[str] connection_notification_arn: The ARN of the SNS topic for the notifications.
        :param pulumi.Input[str] notification_type: The type of notification.
        :param pulumi.Input[str] state: The state of the notification.
        :param pulumi.Input[str] vpc_endpoint_id: The ID of the VPC Endpoint to receive notifications for.
        :param pulumi.Input[str] vpc_endpoint_service_id: The ID of the VPC Endpoint Service to receive notifications for.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcEndpointConnectionNotificationState.__new__(_VpcEndpointConnectionNotificationState)

        __props__.__dict__["connection_events"] = connection_events
        __props__.__dict__["connection_notification_arn"] = connection_notification_arn
        __props__.__dict__["notification_type"] = notification_type
        __props__.__dict__["state"] = state
        __props__.__dict__["vpc_endpoint_id"] = vpc_endpoint_id
        __props__.__dict__["vpc_endpoint_service_id"] = vpc_endpoint_service_id
        return VpcEndpointConnectionNotification(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectionEvents")
    def connection_events(self) -> pulumi.Output[Sequence[str]]:
        """
        One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
        """
        return pulumi.get(self, "connection_events")

    @property
    @pulumi.getter(name="connectionNotificationArn")
    def connection_notification_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the SNS topic for the notifications.
        """
        return pulumi.get(self, "connection_notification_arn")

    @property
    @pulumi.getter(name="notificationType")
    def notification_type(self) -> pulumi.Output[str]:
        """
        The type of notification.
        """
        return pulumi.get(self, "notification_type")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The state of the notification.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="vpcEndpointId")
    def vpc_endpoint_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the VPC Endpoint to receive notifications for.
        """
        return pulumi.get(self, "vpc_endpoint_id")

    @property
    @pulumi.getter(name="vpcEndpointServiceId")
    def vpc_endpoint_service_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the VPC Endpoint Service to receive notifications for.
        """
        return pulumi.get(self, "vpc_endpoint_service_id")

