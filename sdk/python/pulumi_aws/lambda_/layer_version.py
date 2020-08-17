# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['LayerVersion']


class LayerVersion(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 code: Optional[pulumi.Input[pulumi.Archive]] = None,
                 compatible_runtimes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 layer_name: Optional[pulumi.Input[str]] = None,
                 license_info: Optional[pulumi.Input[str]] = None,
                 s3_bucket: Optional[pulumi.Input[str]] = None,
                 s3_key: Optional[pulumi.Input[str]] = None,
                 s3_object_version: Optional[pulumi.Input[str]] = None,
                 source_code_hash: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Lambda Layer Version resource. Lambda Layers allow you to reuse shared bits of code across multiple lambda functions.

        For information about Lambda Layers and how to use them, see [AWS Lambda Layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        lambda_layer = aws.lambda_.LayerVersion("lambdaLayer",
            compatible_runtimes=["nodejs8.10"],
            code=pulumi.FileArchive("lambda_layer_payload.zip"),
            layer_name="lambda_layer_name")
        ```
        ## Specifying the Deployment Package

        AWS Lambda Layers expect source code to be provided as a deployment package whose structure varies depending on which `compatible_runtimes` this layer specifies.
        See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) for the valid values of `compatible_runtimes`.

        Once you have created your deployment package you can specify it either directly as a local file (using the `filename` argument) or
        indirectly via Amazon S3 (using the `s3_bucket`, `s3_key` and `s3_object_version` arguments). When providing the deployment
        package via S3 it may be useful to use the `s3.BucketObject` resource to upload it.

        For larger deployment packages it is recommended by Amazon to upload via S3, since the S3 API has better support for uploading
        large files efficiently.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.Archive] code: The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
        :param pulumi.Input[List[pulumi.Input[str]]] compatible_runtimes: A list of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 5 runtimes can be specified.
        :param pulumi.Input[str] description: Description of what your Lambda Layer does.
        :param pulumi.Input[str] layer_name: A unique name for your Lambda Layer
        :param pulumi.Input[str] license_info: License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
        :param pulumi.Input[str] s3_bucket: The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
        :param pulumi.Input[str] s3_key: The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
        :param pulumi.Input[str] s3_object_version: The object version containing the function's deployment package. Conflicts with `filename`.
        :param pulumi.Input[str] source_code_hash: Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3_key`. The usual way to set this is `${filebase64sha256("file.zip")}` (this provider 0.11.12 or later) or `${base64sha256(file("file.zip"))}` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda layer source archive.
        """
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
            __props__ = dict()

            __props__['code'] = code
            __props__['compatible_runtimes'] = compatible_runtimes
            __props__['description'] = description
            if layer_name is None:
                raise TypeError("Missing required property 'layer_name'")
            __props__['layer_name'] = layer_name
            __props__['license_info'] = license_info
            __props__['s3_bucket'] = s3_bucket
            __props__['s3_key'] = s3_key
            __props__['s3_object_version'] = s3_object_version
            __props__['source_code_hash'] = source_code_hash
            __props__['arn'] = None
            __props__['created_date'] = None
            __props__['layer_arn'] = None
            __props__['source_code_size'] = None
            __props__['version'] = None
        super(LayerVersion, __self__).__init__(
            'aws:lambda/layerVersion:LayerVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            code: Optional[pulumi.Input[pulumi.Archive]] = None,
            compatible_runtimes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            created_date: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            layer_arn: Optional[pulumi.Input[str]] = None,
            layer_name: Optional[pulumi.Input[str]] = None,
            license_info: Optional[pulumi.Input[str]] = None,
            s3_bucket: Optional[pulumi.Input[str]] = None,
            s3_key: Optional[pulumi.Input[str]] = None,
            s3_object_version: Optional[pulumi.Input[str]] = None,
            source_code_hash: Optional[pulumi.Input[str]] = None,
            source_code_size: Optional[pulumi.Input[float]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'LayerVersion':
        """
        Get an existing LayerVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the Lambda Layer with version.
        :param pulumi.Input[pulumi.Archive] code: The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
        :param pulumi.Input[List[pulumi.Input[str]]] compatible_runtimes: A list of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 5 runtimes can be specified.
        :param pulumi.Input[str] created_date: The date this resource was created.
        :param pulumi.Input[str] description: Description of what your Lambda Layer does.
        :param pulumi.Input[str] layer_arn: The Amazon Resource Name (ARN) of the Lambda Layer without version.
        :param pulumi.Input[str] layer_name: A unique name for your Lambda Layer
        :param pulumi.Input[str] license_info: License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
        :param pulumi.Input[str] s3_bucket: The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
        :param pulumi.Input[str] s3_key: The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
        :param pulumi.Input[str] s3_object_version: The object version containing the function's deployment package. Conflicts with `filename`.
        :param pulumi.Input[str] source_code_hash: Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3_key`. The usual way to set this is `${filebase64sha256("file.zip")}` (this provider 0.11.12 or later) or `${base64sha256(file("file.zip"))}` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda layer source archive.
        :param pulumi.Input[float] source_code_size: The size in bytes of the function .zip file.
        :param pulumi.Input[str] version: This Lamba Layer version.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["code"] = code
        __props__["compatible_runtimes"] = compatible_runtimes
        __props__["created_date"] = created_date
        __props__["description"] = description
        __props__["layer_arn"] = layer_arn
        __props__["layer_name"] = layer_name
        __props__["license_info"] = license_info
        __props__["s3_bucket"] = s3_bucket
        __props__["s3_key"] = s3_key
        __props__["s3_object_version"] = s3_object_version
        __props__["source_code_hash"] = source_code_hash
        __props__["source_code_size"] = source_code_size
        __props__["version"] = version
        return LayerVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The Amazon Resource Name (ARN) of the Lambda Layer with version.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def code(self) -> Optional[pulumi.Archive]:
        """
        The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter(name="compatibleRuntimes")
    def compatible_runtimes(self) -> Optional[List[str]]:
        """
        A list of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 5 runtimes can be specified.
        """
        return pulumi.get(self, "compatible_runtimes")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> str:
        """
        The date this resource was created.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of what your Lambda Layer does.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="layerArn")
    def layer_arn(self) -> str:
        """
        The Amazon Resource Name (ARN) of the Lambda Layer without version.
        """
        return pulumi.get(self, "layer_arn")

    @property
    @pulumi.getter(name="layerName")
    def layer_name(self) -> str:
        """
        A unique name for your Lambda Layer
        """
        return pulumi.get(self, "layer_name")

    @property
    @pulumi.getter(name="licenseInfo")
    def license_info(self) -> Optional[str]:
        """
        License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
        """
        return pulumi.get(self, "license_info")

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> Optional[str]:
        """
        The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
        """
        return pulumi.get(self, "s3_bucket")

    @property
    @pulumi.getter(name="s3Key")
    def s3_key(self) -> Optional[str]:
        """
        The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
        """
        return pulumi.get(self, "s3_key")

    @property
    @pulumi.getter(name="s3ObjectVersion")
    def s3_object_version(self) -> Optional[str]:
        """
        The object version containing the function's deployment package. Conflicts with `filename`.
        """
        return pulumi.get(self, "s3_object_version")

    @property
    @pulumi.getter(name="sourceCodeHash")
    def source_code_hash(self) -> str:
        """
        Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3_key`. The usual way to set this is `${filebase64sha256("file.zip")}` (this provider 0.11.12 or later) or `${base64sha256(file("file.zip"))}` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda layer source archive.
        """
        return pulumi.get(self, "source_code_hash")

    @property
    @pulumi.getter(name="sourceCodeSize")
    def source_code_size(self) -> float:
        """
        The size in bytes of the function .zip file.
        """
        return pulumi.get(self, "source_code_size")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        This Lamba Layer version.
        """
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

