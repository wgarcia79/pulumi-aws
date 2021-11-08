# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ControlPanelArgs', 'ControlPanel']

@pulumi.input_type
class ControlPanelArgs:
    def __init__(__self__, *,
                 cluster_arn: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ControlPanel resource.
        :param pulumi.Input[str] cluster_arn: ARN of the cluster in which this control panel will reside.
        :param pulumi.Input[str] name: Name describing the control panel.
        """
        pulumi.set(__self__, "cluster_arn", cluster_arn)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="clusterArn")
    def cluster_arn(self) -> pulumi.Input[str]:
        """
        ARN of the cluster in which this control panel will reside.
        """
        return pulumi.get(self, "cluster_arn")

    @cluster_arn.setter
    def cluster_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_arn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name describing the control panel.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ControlPanelState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 cluster_arn: Optional[pulumi.Input[str]] = None,
                 default_control_panel: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 routing_control_count: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ControlPanel resources.
        :param pulumi.Input[str] arn: ARN of the control panel.
        :param pulumi.Input[str] cluster_arn: ARN of the cluster in which this control panel will reside.
        :param pulumi.Input[bool] default_control_panel: Whether a control panel is default.
        :param pulumi.Input[str] name: Name describing the control panel.
        :param pulumi.Input[int] routing_control_count: Number routing controls in a control panel.
        :param pulumi.Input[str] status: Status of control panel: `PENDING` when it is being created/updated, `PENDING_DELETION` when it is being deleted, and `DEPLOYED` otherwise.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if cluster_arn is not None:
            pulumi.set(__self__, "cluster_arn", cluster_arn)
        if default_control_panel is not None:
            pulumi.set(__self__, "default_control_panel", default_control_panel)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if routing_control_count is not None:
            pulumi.set(__self__, "routing_control_count", routing_control_count)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the control panel.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="clusterArn")
    def cluster_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the cluster in which this control panel will reside.
        """
        return pulumi.get(self, "cluster_arn")

    @cluster_arn.setter
    def cluster_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_arn", value)

    @property
    @pulumi.getter(name="defaultControlPanel")
    def default_control_panel(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether a control panel is default.
        """
        return pulumi.get(self, "default_control_panel")

    @default_control_panel.setter
    def default_control_panel(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default_control_panel", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name describing the control panel.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="routingControlCount")
    def routing_control_count(self) -> Optional[pulumi.Input[int]]:
        """
        Number routing controls in a control panel.
        """
        return pulumi.get(self, "routing_control_count")

    @routing_control_count.setter
    def routing_control_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "routing_control_count", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of control panel: `PENDING` when it is being created/updated, `PENDING_DELETION` when it is being deleted, and `DEPLOYED` otherwise.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class ControlPanel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an AWS Route 53 Recovery Control Config Control Panel.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.route53recoverycontrol.ControlPanel("example", cluster_arn="arn:aws:route53-recovery-control::123456789012:cluster/8d47920e-d789-437d-803a-2dcc4b204393")
        ```

        ## Import

        Route53 Recovery Control Config Control Panel can be imported via the control panel arn, e.g.,

        ```sh
         $ pulumi import aws:route53recoverycontrol/controlPanel:ControlPanel mypanel arn:aws:route53-recovery-control::313517334327:controlpanel/1bfba17df8684f5dab0467b71424f7e8
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_arn: ARN of the cluster in which this control panel will reside.
        :param pulumi.Input[str] name: Name describing the control panel.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ControlPanelArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AWS Route 53 Recovery Control Config Control Panel.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.route53recoverycontrol.ControlPanel("example", cluster_arn="arn:aws:route53-recovery-control::123456789012:cluster/8d47920e-d789-437d-803a-2dcc4b204393")
        ```

        ## Import

        Route53 Recovery Control Config Control Panel can be imported via the control panel arn, e.g.,

        ```sh
         $ pulumi import aws:route53recoverycontrol/controlPanel:ControlPanel mypanel arn:aws:route53-recovery-control::313517334327:controlpanel/1bfba17df8684f5dab0467b71424f7e8
        ```

        :param str resource_name: The name of the resource.
        :param ControlPanelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ControlPanelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ControlPanelArgs.__new__(ControlPanelArgs)

            if cluster_arn is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_arn'")
            __props__.__dict__["cluster_arn"] = cluster_arn
            __props__.__dict__["name"] = name
            __props__.__dict__["arn"] = None
            __props__.__dict__["default_control_panel"] = None
            __props__.__dict__["routing_control_count"] = None
            __props__.__dict__["status"] = None
        super(ControlPanel, __self__).__init__(
            'aws:route53recoverycontrol/controlPanel:ControlPanel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            cluster_arn: Optional[pulumi.Input[str]] = None,
            default_control_panel: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            routing_control_count: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'ControlPanel':
        """
        Get an existing ControlPanel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: ARN of the control panel.
        :param pulumi.Input[str] cluster_arn: ARN of the cluster in which this control panel will reside.
        :param pulumi.Input[bool] default_control_panel: Whether a control panel is default.
        :param pulumi.Input[str] name: Name describing the control panel.
        :param pulumi.Input[int] routing_control_count: Number routing controls in a control panel.
        :param pulumi.Input[str] status: Status of control panel: `PENDING` when it is being created/updated, `PENDING_DELETION` when it is being deleted, and `DEPLOYED` otherwise.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ControlPanelState.__new__(_ControlPanelState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["cluster_arn"] = cluster_arn
        __props__.__dict__["default_control_panel"] = default_control_panel
        __props__.__dict__["name"] = name
        __props__.__dict__["routing_control_count"] = routing_control_count
        __props__.__dict__["status"] = status
        return ControlPanel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN of the control panel.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="clusterArn")
    def cluster_arn(self) -> pulumi.Output[str]:
        """
        ARN of the cluster in which this control panel will reside.
        """
        return pulumi.get(self, "cluster_arn")

    @property
    @pulumi.getter(name="defaultControlPanel")
    def default_control_panel(self) -> pulumi.Output[bool]:
        """
        Whether a control panel is default.
        """
        return pulumi.get(self, "default_control_panel")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name describing the control panel.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="routingControlCount")
    def routing_control_count(self) -> pulumi.Output[int]:
        """
        Number routing controls in a control panel.
        """
        return pulumi.get(self, "routing_control_count")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of control panel: `PENDING` when it is being created/updated, `PENDING_DELETION` when it is being deleted, and `DEPLOYED` otherwise.
        """
        return pulumi.get(self, "status")

