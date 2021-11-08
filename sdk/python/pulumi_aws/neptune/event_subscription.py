# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EventSubscriptionArgs', 'EventSubscription']

@pulumi.input_type
class EventSubscriptionArgs:
    def __init__(__self__, *,
                 sns_topic_arn: pulumi.Input[str],
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_categories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 source_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 source_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a EventSubscription resource.
        :param pulumi.Input[str] sns_topic_arn: The ARN of the SNS topic to send events to.
        :param pulumi.Input[bool] enabled: A boolean flag to enable/disable the subscription. Defaults to true.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_categories: A list of event categories for a `source_type` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
        :param pulumi.Input[str] name: The name of the Neptune event subscription. By default generated by this provider.
        :param pulumi.Input[str] name_prefix: The name of the Neptune event subscription. Conflicts with `name`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] source_ids: A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `source_type` must also be specified.
        :param pulumi.Input[str] source_type: The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "sns_topic_arn", sns_topic_arn)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if event_categories is not None:
            pulumi.set(__self__, "event_categories", event_categories)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)
        if source_ids is not None:
            pulumi.set(__self__, "source_ids", source_ids)
        if source_type is not None:
            pulumi.set(__self__, "source_type", source_type)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="snsTopicArn")
    def sns_topic_arn(self) -> pulumi.Input[str]:
        """
        The ARN of the SNS topic to send events to.
        """
        return pulumi.get(self, "sns_topic_arn")

    @sns_topic_arn.setter
    def sns_topic_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "sns_topic_arn", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean flag to enable/disable the subscription. Defaults to true.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="eventCategories")
    def event_categories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of event categories for a `source_type` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
        """
        return pulumi.get(self, "event_categories")

    @event_categories.setter
    def event_categories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "event_categories", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Neptune event subscription. By default generated by this provider.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Neptune event subscription. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)

    @property
    @pulumi.getter(name="sourceIds")
    def source_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `source_type` must also be specified.
        """
        return pulumi.get(self, "source_ids")

    @source_ids.setter
    def source_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "source_ids", value)

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
        """
        return pulumi.get(self, "source_type")

    @source_type.setter
    def source_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_type", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _EventSubscriptionState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 customer_aws_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_categories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 sns_topic_arn: Optional[pulumi.Input[str]] = None,
                 source_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 source_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering EventSubscription resources.
        :param pulumi.Input[str] arn: The Amazon Resource Name of the Neptune event notification subscription.
        :param pulumi.Input[str] customer_aws_id: The AWS customer account associated with the Neptune event notification subscription.
        :param pulumi.Input[bool] enabled: A boolean flag to enable/disable the subscription. Defaults to true.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_categories: A list of event categories for a `source_type` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
        :param pulumi.Input[str] name: The name of the Neptune event subscription. By default generated by this provider.
        :param pulumi.Input[str] name_prefix: The name of the Neptune event subscription. Conflicts with `name`.
        :param pulumi.Input[str] sns_topic_arn: The ARN of the SNS topic to send events to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] source_ids: A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `source_type` must also be specified.
        :param pulumi.Input[str] source_type: The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if customer_aws_id is not None:
            pulumi.set(__self__, "customer_aws_id", customer_aws_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if event_categories is not None:
            pulumi.set(__self__, "event_categories", event_categories)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)
        if sns_topic_arn is not None:
            pulumi.set(__self__, "sns_topic_arn", sns_topic_arn)
        if source_ids is not None:
            pulumi.set(__self__, "source_ids", source_ids)
        if source_type is not None:
            pulumi.set(__self__, "source_type", source_type)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name of the Neptune event notification subscription.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="customerAwsId")
    def customer_aws_id(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS customer account associated with the Neptune event notification subscription.
        """
        return pulumi.get(self, "customer_aws_id")

    @customer_aws_id.setter
    def customer_aws_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "customer_aws_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean flag to enable/disable the subscription. Defaults to true.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="eventCategories")
    def event_categories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of event categories for a `source_type` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
        """
        return pulumi.get(self, "event_categories")

    @event_categories.setter
    def event_categories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "event_categories", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Neptune event subscription. By default generated by this provider.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Neptune event subscription. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)

    @property
    @pulumi.getter(name="snsTopicArn")
    def sns_topic_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the SNS topic to send events to.
        """
        return pulumi.get(self, "sns_topic_arn")

    @sns_topic_arn.setter
    def sns_topic_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sns_topic_arn", value)

    @property
    @pulumi.getter(name="sourceIds")
    def source_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `source_type` must also be specified.
        """
        return pulumi.get(self, "source_ids")

    @source_ids.setter
    def source_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "source_ids", value)

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
        """
        return pulumi.get(self, "source_type")

    @source_type.setter
    def source_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_type", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class EventSubscription(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_categories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 sns_topic_arn: Optional[pulumi.Input[str]] = None,
                 source_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 source_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        default_cluster = aws.neptune.Cluster("defaultCluster",
            cluster_identifier="neptune-cluster-demo",
            engine="neptune",
            backup_retention_period=5,
            preferred_backup_window="07:00-09:00",
            skip_final_snapshot=True,
            iam_database_authentication_enabled=True,
            apply_immediately=True)
        example = aws.neptune.ClusterInstance("example",
            cluster_identifier=default_cluster.id,
            engine="neptune",
            instance_class="db.r4.large",
            apply_immediately=True)
        default_topic = aws.sns.Topic("defaultTopic")
        default_event_subscription = aws.neptune.EventSubscription("defaultEventSubscription",
            sns_topic_arn=default_topic.arn,
            source_type="db-instance",
            source_ids=[example.id],
            event_categories=[
                "maintenance",
                "availability",
                "creation",
                "backup",
                "restoration",
                "recovery",
                "deletion",
                "failover",
                "failure",
                "notification",
                "configuration change",
                "read replica",
            ],
            tags={
                "env": "test",
            })
        ```

        ## Import

        `aws_neptune_event_subscription` can be imported by using the event subscription name, e.g.,

        ```sh
         $ pulumi import aws:neptune/eventSubscription:EventSubscription example my-event-subscription
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: A boolean flag to enable/disable the subscription. Defaults to true.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_categories: A list of event categories for a `source_type` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
        :param pulumi.Input[str] name: The name of the Neptune event subscription. By default generated by this provider.
        :param pulumi.Input[str] name_prefix: The name of the Neptune event subscription. Conflicts with `name`.
        :param pulumi.Input[str] sns_topic_arn: The ARN of the SNS topic to send events to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] source_ids: A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `source_type` must also be specified.
        :param pulumi.Input[str] source_type: The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EventSubscriptionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        default_cluster = aws.neptune.Cluster("defaultCluster",
            cluster_identifier="neptune-cluster-demo",
            engine="neptune",
            backup_retention_period=5,
            preferred_backup_window="07:00-09:00",
            skip_final_snapshot=True,
            iam_database_authentication_enabled=True,
            apply_immediately=True)
        example = aws.neptune.ClusterInstance("example",
            cluster_identifier=default_cluster.id,
            engine="neptune",
            instance_class="db.r4.large",
            apply_immediately=True)
        default_topic = aws.sns.Topic("defaultTopic")
        default_event_subscription = aws.neptune.EventSubscription("defaultEventSubscription",
            sns_topic_arn=default_topic.arn,
            source_type="db-instance",
            source_ids=[example.id],
            event_categories=[
                "maintenance",
                "availability",
                "creation",
                "backup",
                "restoration",
                "recovery",
                "deletion",
                "failover",
                "failure",
                "notification",
                "configuration change",
                "read replica",
            ],
            tags={
                "env": "test",
            })
        ```

        ## Import

        `aws_neptune_event_subscription` can be imported by using the event subscription name, e.g.,

        ```sh
         $ pulumi import aws:neptune/eventSubscription:EventSubscription example my-event-subscription
        ```

        :param str resource_name: The name of the resource.
        :param EventSubscriptionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EventSubscriptionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_categories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 sns_topic_arn: Optional[pulumi.Input[str]] = None,
                 source_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 source_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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
            __props__ = EventSubscriptionArgs.__new__(EventSubscriptionArgs)

            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["event_categories"] = event_categories
            __props__.__dict__["name"] = name
            __props__.__dict__["name_prefix"] = name_prefix
            if sns_topic_arn is None and not opts.urn:
                raise TypeError("Missing required property 'sns_topic_arn'")
            __props__.__dict__["sns_topic_arn"] = sns_topic_arn
            __props__.__dict__["source_ids"] = source_ids
            __props__.__dict__["source_type"] = source_type
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["customer_aws_id"] = None
            __props__.__dict__["tags_all"] = None
        super(EventSubscription, __self__).__init__(
            'aws:neptune/eventSubscription:EventSubscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            customer_aws_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            event_categories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            sns_topic_arn: Optional[pulumi.Input[str]] = None,
            source_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            source_type: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'EventSubscription':
        """
        Get an existing EventSubscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name of the Neptune event notification subscription.
        :param pulumi.Input[str] customer_aws_id: The AWS customer account associated with the Neptune event notification subscription.
        :param pulumi.Input[bool] enabled: A boolean flag to enable/disable the subscription. Defaults to true.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_categories: A list of event categories for a `source_type` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
        :param pulumi.Input[str] name: The name of the Neptune event subscription. By default generated by this provider.
        :param pulumi.Input[str] name_prefix: The name of the Neptune event subscription. Conflicts with `name`.
        :param pulumi.Input[str] sns_topic_arn: The ARN of the SNS topic to send events to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] source_ids: A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `source_type` must also be specified.
        :param pulumi.Input[str] source_type: The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EventSubscriptionState.__new__(_EventSubscriptionState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["customer_aws_id"] = customer_aws_id
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["event_categories"] = event_categories
        __props__.__dict__["name"] = name
        __props__.__dict__["name_prefix"] = name_prefix
        __props__.__dict__["sns_topic_arn"] = sns_topic_arn
        __props__.__dict__["source_ids"] = source_ids
        __props__.__dict__["source_type"] = source_type
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return EventSubscription(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name of the Neptune event notification subscription.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="customerAwsId")
    def customer_aws_id(self) -> pulumi.Output[str]:
        """
        The AWS customer account associated with the Neptune event notification subscription.
        """
        return pulumi.get(self, "customer_aws_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean flag to enable/disable the subscription. Defaults to true.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="eventCategories")
    def event_categories(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of event categories for a `source_type` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
        """
        return pulumi.get(self, "event_categories")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Neptune event subscription. By default generated by this provider.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Output[str]:
        """
        The name of the Neptune event subscription. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter(name="snsTopicArn")
    def sns_topic_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the SNS topic to send events to.
        """
        return pulumi.get(self, "sns_topic_arn")

    @property
    @pulumi.getter(name="sourceIds")
    def source_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `source_type` must also be specified.
        """
        return pulumi.get(self, "source_ids")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
        """
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

