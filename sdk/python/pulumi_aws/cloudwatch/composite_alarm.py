# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CompositeAlarmArgs', 'CompositeAlarm']

@pulumi.input_type
class CompositeAlarmArgs:
    def __init__(__self__, *,
                 alarm_name: pulumi.Input[str],
                 alarm_rule: pulumi.Input[str],
                 actions_enabled: Optional[pulumi.Input[bool]] = None,
                 alarm_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 alarm_description: Optional[pulumi.Input[str]] = None,
                 insufficient_data_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ok_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a CompositeAlarm resource.
        :param pulumi.Input[str] alarm_name: The name for the composite alarm. This name must be unique within the region.
        :param pulumi.Input[str] alarm_rule: An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see [Creating a Composite Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Composite_Alarm.html). The maximum length is 10240 characters.
        :param pulumi.Input[bool] actions_enabled: Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alarm_actions: The set of actions to execute when this alarm transitions to the `ALARM` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[str] alarm_description: The description for the composite alarm.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] insufficient_data_actions: The set of actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ok_actions: The set of actions to execute when this alarm transitions to an `OK` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to associate with the alarm. Up to 50 tags are allowed.
        """
        pulumi.set(__self__, "alarm_name", alarm_name)
        pulumi.set(__self__, "alarm_rule", alarm_rule)
        if actions_enabled is not None:
            pulumi.set(__self__, "actions_enabled", actions_enabled)
        if alarm_actions is not None:
            pulumi.set(__self__, "alarm_actions", alarm_actions)
        if alarm_description is not None:
            pulumi.set(__self__, "alarm_description", alarm_description)
        if insufficient_data_actions is not None:
            pulumi.set(__self__, "insufficient_data_actions", insufficient_data_actions)
        if ok_actions is not None:
            pulumi.set(__self__, "ok_actions", ok_actions)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="alarmName")
    def alarm_name(self) -> pulumi.Input[str]:
        """
        The name for the composite alarm. This name must be unique within the region.
        """
        return pulumi.get(self, "alarm_name")

    @alarm_name.setter
    def alarm_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "alarm_name", value)

    @property
    @pulumi.getter(name="alarmRule")
    def alarm_rule(self) -> pulumi.Input[str]:
        """
        An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see [Creating a Composite Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Composite_Alarm.html). The maximum length is 10240 characters.
        """
        return pulumi.get(self, "alarm_rule")

    @alarm_rule.setter
    def alarm_rule(self, value: pulumi.Input[str]):
        pulumi.set(self, "alarm_rule", value)

    @property
    @pulumi.getter(name="actionsEnabled")
    def actions_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to `true`.
        """
        return pulumi.get(self, "actions_enabled")

    @actions_enabled.setter
    def actions_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "actions_enabled", value)

    @property
    @pulumi.getter(name="alarmActions")
    def alarm_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The set of actions to execute when this alarm transitions to the `ALARM` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        """
        return pulumi.get(self, "alarm_actions")

    @alarm_actions.setter
    def alarm_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "alarm_actions", value)

    @property
    @pulumi.getter(name="alarmDescription")
    def alarm_description(self) -> Optional[pulumi.Input[str]]:
        """
        The description for the composite alarm.
        """
        return pulumi.get(self, "alarm_description")

    @alarm_description.setter
    def alarm_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alarm_description", value)

    @property
    @pulumi.getter(name="insufficientDataActions")
    def insufficient_data_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The set of actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        """
        return pulumi.get(self, "insufficient_data_actions")

    @insufficient_data_actions.setter
    def insufficient_data_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "insufficient_data_actions", value)

    @property
    @pulumi.getter(name="okActions")
    def ok_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The set of actions to execute when this alarm transitions to an `OK` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        """
        return pulumi.get(self, "ok_actions")

    @ok_actions.setter
    def ok_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ok_actions", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to associate with the alarm. Up to 50 tags are allowed.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _CompositeAlarmState:
    def __init__(__self__, *,
                 actions_enabled: Optional[pulumi.Input[bool]] = None,
                 alarm_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 alarm_description: Optional[pulumi.Input[str]] = None,
                 alarm_name: Optional[pulumi.Input[str]] = None,
                 alarm_rule: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 insufficient_data_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ok_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering CompositeAlarm resources.
        :param pulumi.Input[bool] actions_enabled: Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alarm_actions: The set of actions to execute when this alarm transitions to the `ALARM` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[str] alarm_description: The description for the composite alarm.
        :param pulumi.Input[str] alarm_name: The name for the composite alarm. This name must be unique within the region.
        :param pulumi.Input[str] alarm_rule: An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see [Creating a Composite Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Composite_Alarm.html). The maximum length is 10240 characters.
        :param pulumi.Input[str] arn: The ARN of the composite alarm.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] insufficient_data_actions: The set of actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ok_actions: The set of actions to execute when this alarm transitions to an `OK` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to associate with the alarm. Up to 50 tags are allowed.
        """
        if actions_enabled is not None:
            pulumi.set(__self__, "actions_enabled", actions_enabled)
        if alarm_actions is not None:
            pulumi.set(__self__, "alarm_actions", alarm_actions)
        if alarm_description is not None:
            pulumi.set(__self__, "alarm_description", alarm_description)
        if alarm_name is not None:
            pulumi.set(__self__, "alarm_name", alarm_name)
        if alarm_rule is not None:
            pulumi.set(__self__, "alarm_rule", alarm_rule)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if insufficient_data_actions is not None:
            pulumi.set(__self__, "insufficient_data_actions", insufficient_data_actions)
        if ok_actions is not None:
            pulumi.set(__self__, "ok_actions", ok_actions)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="actionsEnabled")
    def actions_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to `true`.
        """
        return pulumi.get(self, "actions_enabled")

    @actions_enabled.setter
    def actions_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "actions_enabled", value)

    @property
    @pulumi.getter(name="alarmActions")
    def alarm_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The set of actions to execute when this alarm transitions to the `ALARM` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        """
        return pulumi.get(self, "alarm_actions")

    @alarm_actions.setter
    def alarm_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "alarm_actions", value)

    @property
    @pulumi.getter(name="alarmDescription")
    def alarm_description(self) -> Optional[pulumi.Input[str]]:
        """
        The description for the composite alarm.
        """
        return pulumi.get(self, "alarm_description")

    @alarm_description.setter
    def alarm_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alarm_description", value)

    @property
    @pulumi.getter(name="alarmName")
    def alarm_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for the composite alarm. This name must be unique within the region.
        """
        return pulumi.get(self, "alarm_name")

    @alarm_name.setter
    def alarm_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alarm_name", value)

    @property
    @pulumi.getter(name="alarmRule")
    def alarm_rule(self) -> Optional[pulumi.Input[str]]:
        """
        An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see [Creating a Composite Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Composite_Alarm.html). The maximum length is 10240 characters.
        """
        return pulumi.get(self, "alarm_rule")

    @alarm_rule.setter
    def alarm_rule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alarm_rule", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the composite alarm.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="insufficientDataActions")
    def insufficient_data_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The set of actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        """
        return pulumi.get(self, "insufficient_data_actions")

    @insufficient_data_actions.setter
    def insufficient_data_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "insufficient_data_actions", value)

    @property
    @pulumi.getter(name="okActions")
    def ok_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The set of actions to execute when this alarm transitions to an `OK` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        """
        return pulumi.get(self, "ok_actions")

    @ok_actions.setter
    def ok_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ok_actions", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to associate with the alarm. Up to 50 tags are allowed.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class CompositeAlarm(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions_enabled: Optional[pulumi.Input[bool]] = None,
                 alarm_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 alarm_description: Optional[pulumi.Input[str]] = None,
                 alarm_name: Optional[pulumi.Input[str]] = None,
                 alarm_rule: Optional[pulumi.Input[str]] = None,
                 insufficient_data_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ok_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a CloudWatch Composite Alarm resource.

        > **NOTE:** An alarm (composite or metric) cannot be destroyed when there are other composite alarms depending on it. This can lead to a cyclical dependency on update, as the provider will unsuccessfully attempt to destroy alarms before updating the rule. Consider using `depends_on`, references to alarm names, and two-stage updates.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.cloudwatch.CompositeAlarm("example",
            alarm_description="This is a composite alarm!",
            alarm_name="example-composite-alarm",
            alarm_actions=aws_sns_topic["example"]["arn"],
            ok_actions=aws_sns_topic["example"]["arn"],
            alarm_rule=f\"\"\"ALARM({aws_cloudwatch_metric_alarm["alpha"]["alarm_name"]}) OR
        ALARM({aws_cloudwatch_metric_alarm["bravo"]["alarm_name"]})
        \"\"\")
        ```

        ## Import

        Use the `alarm_name` to import a CloudWatch Composite Alarm. For example

        ```sh
         $ pulumi import aws:cloudwatch/compositeAlarm:CompositeAlarm test my-alarm
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] actions_enabled: Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alarm_actions: The set of actions to execute when this alarm transitions to the `ALARM` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[str] alarm_description: The description for the composite alarm.
        :param pulumi.Input[str] alarm_name: The name for the composite alarm. This name must be unique within the region.
        :param pulumi.Input[str] alarm_rule: An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see [Creating a Composite Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Composite_Alarm.html). The maximum length is 10240 characters.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] insufficient_data_actions: The set of actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ok_actions: The set of actions to execute when this alarm transitions to an `OK` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to associate with the alarm. Up to 50 tags are allowed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CompositeAlarmArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a CloudWatch Composite Alarm resource.

        > **NOTE:** An alarm (composite or metric) cannot be destroyed when there are other composite alarms depending on it. This can lead to a cyclical dependency on update, as the provider will unsuccessfully attempt to destroy alarms before updating the rule. Consider using `depends_on`, references to alarm names, and two-stage updates.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.cloudwatch.CompositeAlarm("example",
            alarm_description="This is a composite alarm!",
            alarm_name="example-composite-alarm",
            alarm_actions=aws_sns_topic["example"]["arn"],
            ok_actions=aws_sns_topic["example"]["arn"],
            alarm_rule=f\"\"\"ALARM({aws_cloudwatch_metric_alarm["alpha"]["alarm_name"]}) OR
        ALARM({aws_cloudwatch_metric_alarm["bravo"]["alarm_name"]})
        \"\"\")
        ```

        ## Import

        Use the `alarm_name` to import a CloudWatch Composite Alarm. For example

        ```sh
         $ pulumi import aws:cloudwatch/compositeAlarm:CompositeAlarm test my-alarm
        ```

        :param str resource_name: The name of the resource.
        :param CompositeAlarmArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CompositeAlarmArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions_enabled: Optional[pulumi.Input[bool]] = None,
                 alarm_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 alarm_description: Optional[pulumi.Input[str]] = None,
                 alarm_name: Optional[pulumi.Input[str]] = None,
                 alarm_rule: Optional[pulumi.Input[str]] = None,
                 insufficient_data_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ok_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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
            __props__ = CompositeAlarmArgs.__new__(CompositeAlarmArgs)

            __props__.__dict__["actions_enabled"] = actions_enabled
            __props__.__dict__["alarm_actions"] = alarm_actions
            __props__.__dict__["alarm_description"] = alarm_description
            if alarm_name is None and not opts.urn:
                raise TypeError("Missing required property 'alarm_name'")
            __props__.__dict__["alarm_name"] = alarm_name
            if alarm_rule is None and not opts.urn:
                raise TypeError("Missing required property 'alarm_rule'")
            __props__.__dict__["alarm_rule"] = alarm_rule
            __props__.__dict__["insufficient_data_actions"] = insufficient_data_actions
            __props__.__dict__["ok_actions"] = ok_actions
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
        super(CompositeAlarm, __self__).__init__(
            'aws:cloudwatch/compositeAlarm:CompositeAlarm',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            actions_enabled: Optional[pulumi.Input[bool]] = None,
            alarm_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            alarm_description: Optional[pulumi.Input[str]] = None,
            alarm_name: Optional[pulumi.Input[str]] = None,
            alarm_rule: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            insufficient_data_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ok_actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'CompositeAlarm':
        """
        Get an existing CompositeAlarm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] actions_enabled: Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alarm_actions: The set of actions to execute when this alarm transitions to the `ALARM` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[str] alarm_description: The description for the composite alarm.
        :param pulumi.Input[str] alarm_name: The name for the composite alarm. This name must be unique within the region.
        :param pulumi.Input[str] alarm_rule: An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see [Creating a Composite Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Composite_Alarm.html). The maximum length is 10240 characters.
        :param pulumi.Input[str] arn: The ARN of the composite alarm.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] insufficient_data_actions: The set of actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ok_actions: The set of actions to execute when this alarm transitions to an `OK` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to associate with the alarm. Up to 50 tags are allowed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CompositeAlarmState.__new__(_CompositeAlarmState)

        __props__.__dict__["actions_enabled"] = actions_enabled
        __props__.__dict__["alarm_actions"] = alarm_actions
        __props__.__dict__["alarm_description"] = alarm_description
        __props__.__dict__["alarm_name"] = alarm_name
        __props__.__dict__["alarm_rule"] = alarm_rule
        __props__.__dict__["arn"] = arn
        __props__.__dict__["insufficient_data_actions"] = insufficient_data_actions
        __props__.__dict__["ok_actions"] = ok_actions
        __props__.__dict__["tags"] = tags
        return CompositeAlarm(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="actionsEnabled")
    def actions_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to `true`.
        """
        return pulumi.get(self, "actions_enabled")

    @property
    @pulumi.getter(name="alarmActions")
    def alarm_actions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The set of actions to execute when this alarm transitions to the `ALARM` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        """
        return pulumi.get(self, "alarm_actions")

    @property
    @pulumi.getter(name="alarmDescription")
    def alarm_description(self) -> pulumi.Output[Optional[str]]:
        """
        The description for the composite alarm.
        """
        return pulumi.get(self, "alarm_description")

    @property
    @pulumi.getter(name="alarmName")
    def alarm_name(self) -> pulumi.Output[str]:
        """
        The name for the composite alarm. This name must be unique within the region.
        """
        return pulumi.get(self, "alarm_name")

    @property
    @pulumi.getter(name="alarmRule")
    def alarm_rule(self) -> pulumi.Output[str]:
        """
        An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see [Creating a Composite Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Composite_Alarm.html). The maximum length is 10240 characters.
        """
        return pulumi.get(self, "alarm_rule")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the composite alarm.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="insufficientDataActions")
    def insufficient_data_actions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The set of actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        """
        return pulumi.get(self, "insufficient_data_actions")

    @property
    @pulumi.getter(name="okActions")
    def ok_actions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The set of actions to execute when this alarm transitions to an `OK` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
        """
        return pulumi.get(self, "ok_actions")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to associate with the alarm. Up to 50 tags are allowed.
        """
        return pulumi.get(self, "tags")

