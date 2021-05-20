# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ApplicationSnapshotArgs', 'ApplicationSnapshot']

@pulumi.input_type
class ApplicationSnapshotArgs:
    def __init__(__self__, *,
                 application_name: pulumi.Input[str],
                 snapshot_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a ApplicationSnapshot resource.
        :param pulumi.Input[str] application_name: The name of an existing  Kinesis Analytics v2 Application. Note that the application must be running for a snapshot to be created.
        :param pulumi.Input[str] snapshot_name: The name of the application snapshot.
        """
        pulumi.set(__self__, "application_name", application_name)
        pulumi.set(__self__, "snapshot_name", snapshot_name)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> pulumi.Input[str]:
        """
        The name of an existing  Kinesis Analytics v2 Application. Note that the application must be running for a snapshot to be created.
        """
        return pulumi.get(self, "application_name")

    @application_name.setter
    def application_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_name", value)

    @property
    @pulumi.getter(name="snapshotName")
    def snapshot_name(self) -> pulumi.Input[str]:
        """
        The name of the application snapshot.
        """
        return pulumi.get(self, "snapshot_name")

    @snapshot_name.setter
    def snapshot_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "snapshot_name", value)


@pulumi.input_type
class _ApplicationSnapshotState:
    def __init__(__self__, *,
                 application_name: Optional[pulumi.Input[str]] = None,
                 application_version_id: Optional[pulumi.Input[int]] = None,
                 snapshot_creation_timestamp: Optional[pulumi.Input[str]] = None,
                 snapshot_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplicationSnapshot resources.
        :param pulumi.Input[str] application_name: The name of an existing  Kinesis Analytics v2 Application. Note that the application must be running for a snapshot to be created.
        :param pulumi.Input[int] application_version_id: The current application version ID when the snapshot was created.
        :param pulumi.Input[str] snapshot_creation_timestamp: The timestamp of the application snapshot.
        :param pulumi.Input[str] snapshot_name: The name of the application snapshot.
        """
        if application_name is not None:
            pulumi.set(__self__, "application_name", application_name)
        if application_version_id is not None:
            pulumi.set(__self__, "application_version_id", application_version_id)
        if snapshot_creation_timestamp is not None:
            pulumi.set(__self__, "snapshot_creation_timestamp", snapshot_creation_timestamp)
        if snapshot_name is not None:
            pulumi.set(__self__, "snapshot_name", snapshot_name)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of an existing  Kinesis Analytics v2 Application. Note that the application must be running for a snapshot to be created.
        """
        return pulumi.get(self, "application_name")

    @application_name.setter
    def application_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_name", value)

    @property
    @pulumi.getter(name="applicationVersionId")
    def application_version_id(self) -> Optional[pulumi.Input[int]]:
        """
        The current application version ID when the snapshot was created.
        """
        return pulumi.get(self, "application_version_id")

    @application_version_id.setter
    def application_version_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "application_version_id", value)

    @property
    @pulumi.getter(name="snapshotCreationTimestamp")
    def snapshot_creation_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        The timestamp of the application snapshot.
        """
        return pulumi.get(self, "snapshot_creation_timestamp")

    @snapshot_creation_timestamp.setter
    def snapshot_creation_timestamp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_creation_timestamp", value)

    @property
    @pulumi.getter(name="snapshotName")
    def snapshot_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the application snapshot.
        """
        return pulumi.get(self, "snapshot_name")

    @snapshot_name.setter
    def snapshot_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_name", value)


class ApplicationSnapshot(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_name: Optional[pulumi.Input[str]] = None,
                 snapshot_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Kinesis Analytics v2 Application Snapshot.
        Snapshots are the AWS implementation of [Flink Savepoints](https://ci.apache.org/projects/flink/flink-docs-release-1.11/ops/state/savepoints.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.kinesisanalyticsv2.ApplicationSnapshot("example",
            application_name=aws_kinesisanalyticsv2_application["example"]["name"],
            snapshot_name="example-snapshot")
        ```

        ## Import

        `aws_kinesisanalyticsv2_application` can be imported by using `application_name` together with `snapshot_name`, e.g.

        ```sh
         $ pulumi import aws:kinesisanalyticsv2/applicationSnapshot:ApplicationSnapshot example example-application/example-snapshot
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_name: The name of an existing  Kinesis Analytics v2 Application. Note that the application must be running for a snapshot to be created.
        :param pulumi.Input[str] snapshot_name: The name of the application snapshot.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationSnapshotArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Kinesis Analytics v2 Application Snapshot.
        Snapshots are the AWS implementation of [Flink Savepoints](https://ci.apache.org/projects/flink/flink-docs-release-1.11/ops/state/savepoints.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.kinesisanalyticsv2.ApplicationSnapshot("example",
            application_name=aws_kinesisanalyticsv2_application["example"]["name"],
            snapshot_name="example-snapshot")
        ```

        ## Import

        `aws_kinesisanalyticsv2_application` can be imported by using `application_name` together with `snapshot_name`, e.g.

        ```sh
         $ pulumi import aws:kinesisanalyticsv2/applicationSnapshot:ApplicationSnapshot example example-application/example-snapshot
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationSnapshotArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationSnapshotArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_name: Optional[pulumi.Input[str]] = None,
                 snapshot_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = ApplicationSnapshotArgs.__new__(ApplicationSnapshotArgs)

            if application_name is None and not opts.urn:
                raise TypeError("Missing required property 'application_name'")
            __props__.__dict__["application_name"] = application_name
            if snapshot_name is None and not opts.urn:
                raise TypeError("Missing required property 'snapshot_name'")
            __props__.__dict__["snapshot_name"] = snapshot_name
            __props__.__dict__["application_version_id"] = None
            __props__.__dict__["snapshot_creation_timestamp"] = None
        super(ApplicationSnapshot, __self__).__init__(
            'aws:kinesisanalyticsv2/applicationSnapshot:ApplicationSnapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_name: Optional[pulumi.Input[str]] = None,
            application_version_id: Optional[pulumi.Input[int]] = None,
            snapshot_creation_timestamp: Optional[pulumi.Input[str]] = None,
            snapshot_name: Optional[pulumi.Input[str]] = None) -> 'ApplicationSnapshot':
        """
        Get an existing ApplicationSnapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_name: The name of an existing  Kinesis Analytics v2 Application. Note that the application must be running for a snapshot to be created.
        :param pulumi.Input[int] application_version_id: The current application version ID when the snapshot was created.
        :param pulumi.Input[str] snapshot_creation_timestamp: The timestamp of the application snapshot.
        :param pulumi.Input[str] snapshot_name: The name of the application snapshot.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationSnapshotState.__new__(_ApplicationSnapshotState)

        __props__.__dict__["application_name"] = application_name
        __props__.__dict__["application_version_id"] = application_version_id
        __props__.__dict__["snapshot_creation_timestamp"] = snapshot_creation_timestamp
        __props__.__dict__["snapshot_name"] = snapshot_name
        return ApplicationSnapshot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> pulumi.Output[str]:
        """
        The name of an existing  Kinesis Analytics v2 Application. Note that the application must be running for a snapshot to be created.
        """
        return pulumi.get(self, "application_name")

    @property
    @pulumi.getter(name="applicationVersionId")
    def application_version_id(self) -> pulumi.Output[int]:
        """
        The current application version ID when the snapshot was created.
        """
        return pulumi.get(self, "application_version_id")

    @property
    @pulumi.getter(name="snapshotCreationTimestamp")
    def snapshot_creation_timestamp(self) -> pulumi.Output[str]:
        """
        The timestamp of the application snapshot.
        """
        return pulumi.get(self, "snapshot_creation_timestamp")

    @property
    @pulumi.getter(name="snapshotName")
    def snapshot_name(self) -> pulumi.Output[str]:
        """
        The name of the application snapshot.
        """
        return pulumi.get(self, "snapshot_name")

