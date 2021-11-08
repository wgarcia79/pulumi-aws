# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['StreamConsumerArgs', 'StreamConsumer']

@pulumi.input_type
class StreamConsumerArgs:
    def __init__(__self__, *,
                 stream_arn: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a StreamConsumer resource.
        :param pulumi.Input[str] stream_arn: Amazon Resource Name (ARN) of the data stream the consumer is registered with.
        :param pulumi.Input[str] name: Name of the stream consumer.
        """
        pulumi.set(__self__, "stream_arn", stream_arn)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="streamArn")
    def stream_arn(self) -> pulumi.Input[str]:
        """
        Amazon Resource Name (ARN) of the data stream the consumer is registered with.
        """
        return pulumi.get(self, "stream_arn")

    @stream_arn.setter
    def stream_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "stream_arn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the stream consumer.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _StreamConsumerState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 stream_arn: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StreamConsumer resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the stream consumer.
        :param pulumi.Input[str] creation_timestamp: Approximate timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of when the stream consumer was created.
        :param pulumi.Input[str] name: Name of the stream consumer.
        :param pulumi.Input[str] stream_arn: Amazon Resource Name (ARN) of the data stream the consumer is registered with.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if creation_timestamp is not None:
            pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if stream_arn is not None:
            pulumi.set(__self__, "stream_arn", stream_arn)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the stream consumer.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        Approximate timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of when the stream consumer was created.
        """
        return pulumi.get(self, "creation_timestamp")

    @creation_timestamp.setter
    def creation_timestamp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_timestamp", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the stream consumer.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="streamArn")
    def stream_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the data stream the consumer is registered with.
        """
        return pulumi.get(self, "stream_arn")

    @stream_arn.setter
    def stream_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stream_arn", value)


class StreamConsumer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 stream_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage a Kinesis Stream Consumer.

        > **Note:** You can register up to 20 consumers per stream. A given consumer can only be registered with one stream at a time.

        For more details, see the [Amazon Kinesis Stream Consumer Documentation](https://docs.aws.amazon.com/streams/latest/dev/amazon-kinesis-consumers.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_stream = aws.kinesis.Stream("exampleStream", shard_count=1)
        example_stream_consumer = aws.kinesis.StreamConsumer("exampleStreamConsumer", stream_arn=example_stream.arn)
        ```

        ## Import

        Kinesis Stream Consumers can be imported using the Amazon Resource Name (ARN) e.g.,

        ```sh
         $ pulumi import aws:kinesis/streamConsumer:StreamConsumer example arn:aws:kinesis:us-west-2:123456789012:stream/example/consumer/example:1616044553
        ```

         [1]https://docs.aws.amazon.com/streams/latest/dev/amazon-kinesis-consumers.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the stream consumer.
        :param pulumi.Input[str] stream_arn: Amazon Resource Name (ARN) of the data stream the consumer is registered with.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StreamConsumerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage a Kinesis Stream Consumer.

        > **Note:** You can register up to 20 consumers per stream. A given consumer can only be registered with one stream at a time.

        For more details, see the [Amazon Kinesis Stream Consumer Documentation](https://docs.aws.amazon.com/streams/latest/dev/amazon-kinesis-consumers.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_stream = aws.kinesis.Stream("exampleStream", shard_count=1)
        example_stream_consumer = aws.kinesis.StreamConsumer("exampleStreamConsumer", stream_arn=example_stream.arn)
        ```

        ## Import

        Kinesis Stream Consumers can be imported using the Amazon Resource Name (ARN) e.g.,

        ```sh
         $ pulumi import aws:kinesis/streamConsumer:StreamConsumer example arn:aws:kinesis:us-west-2:123456789012:stream/example/consumer/example:1616044553
        ```

         [1]https://docs.aws.amazon.com/streams/latest/dev/amazon-kinesis-consumers.html

        :param str resource_name: The name of the resource.
        :param StreamConsumerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StreamConsumerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 stream_arn: Optional[pulumi.Input[str]] = None,
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
            __props__ = StreamConsumerArgs.__new__(StreamConsumerArgs)

            __props__.__dict__["name"] = name
            if stream_arn is None and not opts.urn:
                raise TypeError("Missing required property 'stream_arn'")
            __props__.__dict__["stream_arn"] = stream_arn
            __props__.__dict__["arn"] = None
            __props__.__dict__["creation_timestamp"] = None
        super(StreamConsumer, __self__).__init__(
            'aws:kinesis/streamConsumer:StreamConsumer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            creation_timestamp: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            stream_arn: Optional[pulumi.Input[str]] = None) -> 'StreamConsumer':
        """
        Get an existing StreamConsumer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the stream consumer.
        :param pulumi.Input[str] creation_timestamp: Approximate timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of when the stream consumer was created.
        :param pulumi.Input[str] name: Name of the stream consumer.
        :param pulumi.Input[str] stream_arn: Amazon Resource Name (ARN) of the data stream the consumer is registered with.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StreamConsumerState.__new__(_StreamConsumerState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["creation_timestamp"] = creation_timestamp
        __props__.__dict__["name"] = name
        __props__.__dict__["stream_arn"] = stream_arn
        return StreamConsumer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the stream consumer.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        Approximate timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of when the stream consumer was created.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the stream consumer.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="streamArn")
    def stream_arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the data stream the consumer is registered with.
        """
        return pulumi.get(self, "stream_arn")

