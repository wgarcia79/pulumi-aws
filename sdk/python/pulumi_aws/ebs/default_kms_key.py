# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DefaultKmsKeyArgs', 'DefaultKmsKey']

@pulumi.input_type
class DefaultKmsKeyArgs:
    def __init__(__self__, *,
                 key_arn: pulumi.Input[str]):
        """
        The set of arguments for constructing a DefaultKmsKey resource.
        :param pulumi.Input[str] key_arn: The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
        """
        pulumi.set(__self__, "key_arn", key_arn)

    @property
    @pulumi.getter(name="keyArn")
    def key_arn(self) -> pulumi.Input[str]:
        """
        The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
        """
        return pulumi.get(self, "key_arn")

    @key_arn.setter
    def key_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_arn", value)


@pulumi.input_type
class _DefaultKmsKeyState:
    def __init__(__self__, *,
                 key_arn: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DefaultKmsKey resources.
        :param pulumi.Input[str] key_arn: The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
        """
        if key_arn is not None:
            pulumi.set(__self__, "key_arn", key_arn)

    @property
    @pulumi.getter(name="keyArn")
    def key_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
        """
        return pulumi.get(self, "key_arn")

    @key_arn.setter
    def key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_arn", value)


class DefaultKmsKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage the default customer master key (CMK) that your AWS account uses to encrypt EBS volumes.

        Your AWS account has an AWS-managed default CMK that is used for encrypting an EBS volume when no CMK is specified in the API call that creates the volume.
        By using the `ebs.DefaultKmsKey` resource, you can specify a customer-managed CMK to use in place of the AWS-managed default CMK.

        > **NOTE:** Creating an `ebs.DefaultKmsKey` resource does not enable default EBS encryption. Use the `ebs.EncryptionByDefault` to enable default EBS encryption.

        > **NOTE:** Destroying this resource will reset the default CMK to the account's AWS-managed default CMK for EBS.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ebs.DefaultKmsKey("example", key_arn=aws_kms_key["example"]["arn"])
        ```

        ## Import

        The EBS default KMS CMK can be imported with the KMS key ARN, e.g., console

        ```sh
         $ pulumi import aws:ebs/defaultKmsKey:DefaultKmsKey example arn:aws:kms:us-east-1:123456789012:key/abcd-1234
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_arn: The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DefaultKmsKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage the default customer master key (CMK) that your AWS account uses to encrypt EBS volumes.

        Your AWS account has an AWS-managed default CMK that is used for encrypting an EBS volume when no CMK is specified in the API call that creates the volume.
        By using the `ebs.DefaultKmsKey` resource, you can specify a customer-managed CMK to use in place of the AWS-managed default CMK.

        > **NOTE:** Creating an `ebs.DefaultKmsKey` resource does not enable default EBS encryption. Use the `ebs.EncryptionByDefault` to enable default EBS encryption.

        > **NOTE:** Destroying this resource will reset the default CMK to the account's AWS-managed default CMK for EBS.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ebs.DefaultKmsKey("example", key_arn=aws_kms_key["example"]["arn"])
        ```

        ## Import

        The EBS default KMS CMK can be imported with the KMS key ARN, e.g., console

        ```sh
         $ pulumi import aws:ebs/defaultKmsKey:DefaultKmsKey example arn:aws:kms:us-east-1:123456789012:key/abcd-1234
        ```

        :param str resource_name: The name of the resource.
        :param DefaultKmsKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DefaultKmsKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_arn: Optional[pulumi.Input[str]] = None,
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
            __props__ = DefaultKmsKeyArgs.__new__(DefaultKmsKeyArgs)

            if key_arn is None and not opts.urn:
                raise TypeError("Missing required property 'key_arn'")
            __props__.__dict__["key_arn"] = key_arn
        super(DefaultKmsKey, __self__).__init__(
            'aws:ebs/defaultKmsKey:DefaultKmsKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key_arn: Optional[pulumi.Input[str]] = None) -> 'DefaultKmsKey':
        """
        Get an existing DefaultKmsKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_arn: The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DefaultKmsKeyState.__new__(_DefaultKmsKeyState)

        __props__.__dict__["key_arn"] = key_arn
        return DefaultKmsKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="keyArn")
    def key_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
        """
        return pulumi.get(self, "key_arn")

