# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RecorderStatusArgs', 'RecorderStatus']

@pulumi.input_type
class RecorderStatusArgs:
    def __init__(__self__, *,
                 is_enabled: pulumi.Input[bool],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RecorderStatus resource.
        :param pulumi.Input[bool] is_enabled: Whether the configuration recorder should be enabled or disabled.
        :param pulumi.Input[str] name: The name of the recorder
        """
        pulumi.set(__self__, "is_enabled", is_enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Input[bool]:
        """
        Whether the configuration recorder should be enabled or disabled.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the recorder
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _RecorderStatusState:
    def __init__(__self__, *,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RecorderStatus resources.
        :param pulumi.Input[bool] is_enabled: Whether the configuration recorder should be enabled or disabled.
        :param pulumi.Input[str] name: The name of the recorder
        """
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the configuration recorder should be enabled or disabled.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the recorder
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class RecorderStatus(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages status (recording / stopped) of an AWS Config Configuration Recorder.

        > **Note:** Starting Configuration Recorder requires a `Delivery Channel` to be present. Use of `depends_on` (as shown below) is recommended to avoid race conditions.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        bucket = aws.s3.Bucket("bucket")
        foo_delivery_channel = aws.cfg.DeliveryChannel("fooDeliveryChannel", s3_bucket_name=bucket.bucket)
        foo_recorder_status = aws.cfg.RecorderStatus("fooRecorderStatus", is_enabled=True,
        opts=pulumi.ResourceOptions(depends_on=[foo_delivery_channel]))
        role = aws.iam.Role("role", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": "sts:AssumeRole",
              "Principal": {
                "Service": "config.amazonaws.com"
              },
              "Effect": "Allow",
              "Sid": ""
            }
          ]
        }
        \"\"\")
        role_policy_attachment = aws.iam.RolePolicyAttachment("rolePolicyAttachment",
            role=role.name,
            policy_arn="arn:aws:iam::aws:policy/service-role/AWSConfigRole")
        foo_recorder = aws.cfg.Recorder("fooRecorder", role_arn=role.arn)
        role_policy = aws.iam.RolePolicy("rolePolicy",
            role=role.id,
            policy=pulumi.Output.all(bucket.arn, bucket.arn).apply(lambda bucketArn, bucketArn1: f\"\"\"{{
          "Version": "2012-10-17",
          "Statement": [
            {{
              "Action": [
                "s3:*"
              ],
              "Effect": "Allow",
              "Resource": [
                "{bucket_arn}",
                "{bucket_arn1}/*"
              ]
            }}
          ]
        }}
        \"\"\"))
        ```

        ## Import

        Configuration Recorder Status can be imported using the name of the Configuration Recorder, e.g.

        ```sh
         $ pulumi import aws:cfg/recorderStatus:RecorderStatus foo example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] is_enabled: Whether the configuration recorder should be enabled or disabled.
        :param pulumi.Input[str] name: The name of the recorder
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RecorderStatusArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages status (recording / stopped) of an AWS Config Configuration Recorder.

        > **Note:** Starting Configuration Recorder requires a `Delivery Channel` to be present. Use of `depends_on` (as shown below) is recommended to avoid race conditions.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        bucket = aws.s3.Bucket("bucket")
        foo_delivery_channel = aws.cfg.DeliveryChannel("fooDeliveryChannel", s3_bucket_name=bucket.bucket)
        foo_recorder_status = aws.cfg.RecorderStatus("fooRecorderStatus", is_enabled=True,
        opts=pulumi.ResourceOptions(depends_on=[foo_delivery_channel]))
        role = aws.iam.Role("role", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": "sts:AssumeRole",
              "Principal": {
                "Service": "config.amazonaws.com"
              },
              "Effect": "Allow",
              "Sid": ""
            }
          ]
        }
        \"\"\")
        role_policy_attachment = aws.iam.RolePolicyAttachment("rolePolicyAttachment",
            role=role.name,
            policy_arn="arn:aws:iam::aws:policy/service-role/AWSConfigRole")
        foo_recorder = aws.cfg.Recorder("fooRecorder", role_arn=role.arn)
        role_policy = aws.iam.RolePolicy("rolePolicy",
            role=role.id,
            policy=pulumi.Output.all(bucket.arn, bucket.arn).apply(lambda bucketArn, bucketArn1: f\"\"\"{{
          "Version": "2012-10-17",
          "Statement": [
            {{
              "Action": [
                "s3:*"
              ],
              "Effect": "Allow",
              "Resource": [
                "{bucket_arn}",
                "{bucket_arn1}/*"
              ]
            }}
          ]
        }}
        \"\"\"))
        ```

        ## Import

        Configuration Recorder Status can be imported using the name of the Configuration Recorder, e.g.

        ```sh
         $ pulumi import aws:cfg/recorderStatus:RecorderStatus foo example
        ```

        :param str resource_name: The name of the resource.
        :param RecorderStatusArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RecorderStatusArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
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
            __props__ = RecorderStatusArgs.__new__(RecorderStatusArgs)

            if is_enabled is None and not opts.urn:
                raise TypeError("Missing required property 'is_enabled'")
            __props__.__dict__["is_enabled"] = is_enabled
            __props__.__dict__["name"] = name
        super(RecorderStatus, __self__).__init__(
            'aws:cfg/recorderStatus:RecorderStatus',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            is_enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'RecorderStatus':
        """
        Get an existing RecorderStatus resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] is_enabled: Whether the configuration recorder should be enabled or disabled.
        :param pulumi.Input[str] name: The name of the recorder
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RecorderStatusState.__new__(_RecorderStatusState)

        __props__.__dict__["is_enabled"] = is_enabled
        __props__.__dict__["name"] = name
        return RecorderStatus(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Output[bool]:
        """
        Whether the configuration recorder should be enabled or disabled.
        """
        return pulumi.get(self, "is_enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the recorder
        """
        return pulumi.get(self, "name")

