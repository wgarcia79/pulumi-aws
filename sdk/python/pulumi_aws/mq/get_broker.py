# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetBrokerResult:
    """
    A collection of values returned by getBroker.
    """
    def __init__(__self__, arn=None, auto_minor_version_upgrade=None, broker_id=None, broker_name=None, configuration=None, deployment_mode=None, engine_type=None, engine_version=None, host_instance_type=None, instances=None, logs=None, maintenance_window_start_time=None, publicly_accessible=None, security_groups=None, subnet_ids=None, tags=None, users=None, id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        if auto_minor_version_upgrade and not isinstance(auto_minor_version_upgrade, bool):
            raise TypeError("Expected argument 'auto_minor_version_upgrade' to be a bool")
        __self__.auto_minor_version_upgrade = auto_minor_version_upgrade
        if broker_id and not isinstance(broker_id, str):
            raise TypeError("Expected argument 'broker_id' to be a str")
        __self__.broker_id = broker_id
        if broker_name and not isinstance(broker_name, str):
            raise TypeError("Expected argument 'broker_name' to be a str")
        __self__.broker_name = broker_name
        if configuration and not isinstance(configuration, dict):
            raise TypeError("Expected argument 'configuration' to be a dict")
        __self__.configuration = configuration
        if deployment_mode and not isinstance(deployment_mode, str):
            raise TypeError("Expected argument 'deployment_mode' to be a str")
        __self__.deployment_mode = deployment_mode
        if engine_type and not isinstance(engine_type, str):
            raise TypeError("Expected argument 'engine_type' to be a str")
        __self__.engine_type = engine_type
        if engine_version and not isinstance(engine_version, str):
            raise TypeError("Expected argument 'engine_version' to be a str")
        __self__.engine_version = engine_version
        if host_instance_type and not isinstance(host_instance_type, str):
            raise TypeError("Expected argument 'host_instance_type' to be a str")
        __self__.host_instance_type = host_instance_type
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        __self__.instances = instances
        if logs and not isinstance(logs, dict):
            raise TypeError("Expected argument 'logs' to be a dict")
        __self__.logs = logs
        if maintenance_window_start_time and not isinstance(maintenance_window_start_time, dict):
            raise TypeError("Expected argument 'maintenance_window_start_time' to be a dict")
        __self__.maintenance_window_start_time = maintenance_window_start_time
        if publicly_accessible and not isinstance(publicly_accessible, bool):
            raise TypeError("Expected argument 'publicly_accessible' to be a bool")
        __self__.publicly_accessible = publicly_accessible
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        __self__.security_groups = security_groups
        if subnet_ids and not isinstance(subnet_ids, list):
            raise TypeError("Expected argument 'subnet_ids' to be a list")
        __self__.subnet_ids = subnet_ids
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        __self__.users = users
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return self

    __iter__ = __await__

def get_broker(broker_id=None,broker_name=None,logs=None,tags=None,opts=None):
    """
    Provides information about a MQ Broker.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/mq_broker.html.markdown.
    """
    __args__ = dict()

    __args__['brokerId'] = broker_id
    __args__['brokerName'] = broker_name
    __args__['logs'] = logs
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:mq/getBroker:getBroker', __args__, opts=opts).value

    return GetBrokerResult(
        arn=__ret__.get('arn'),
        auto_minor_version_upgrade=__ret__.get('autoMinorVersionUpgrade'),
        broker_id=__ret__.get('brokerId'),
        broker_name=__ret__.get('brokerName'),
        configuration=__ret__.get('configuration'),
        deployment_mode=__ret__.get('deploymentMode'),
        engine_type=__ret__.get('engineType'),
        engine_version=__ret__.get('engineVersion'),
        host_instance_type=__ret__.get('hostInstanceType'),
        instances=__ret__.get('instances'),
        logs=__ret__.get('logs'),
        maintenance_window_start_time=__ret__.get('maintenanceWindowStartTime'),
        publicly_accessible=__ret__.get('publiclyAccessible'),
        security_groups=__ret__.get('securityGroups'),
        subnet_ids=__ret__.get('subnetIds'),
        tags=__ret__.get('tags'),
        users=__ret__.get('users'),
        id=__ret__.get('id'))
