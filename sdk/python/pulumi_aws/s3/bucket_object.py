# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['BucketObject']


class BucketObject(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[str]] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 bucket_key_enabled: Optional[pulumi.Input[bool]] = None,
                 cache_control: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 content_base64: Optional[pulumi.Input[str]] = None,
                 content_disposition: Optional[pulumi.Input[str]] = None,
                 content_encoding: Optional[pulumi.Input[str]] = None,
                 content_language: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 object_lock_legal_hold_status: Optional[pulumi.Input[str]] = None,
                 object_lock_mode: Optional[pulumi.Input[str]] = None,
                 object_lock_retain_until_date: Optional[pulumi.Input[str]] = None,
                 server_side_encryption: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]] = None,
                 storage_class: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 website_redirect: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a S3 bucket object resource.

        ## Example Usage
        ### Encrypting with KMS Key

        ```python
        import pulumi
        import pulumi_aws as aws

        examplekms = aws.kms.Key("examplekms",
            description="KMS key 1",
            deletion_window_in_days=7)
        examplebucket = aws.s3.Bucket("examplebucket", acl="private")
        examplebucket_object = aws.s3.BucketObject("examplebucketObject",
            key="someobject",
            bucket=examplebucket.id,
            source=pulumi.FileAsset("index.html"),
            kms_key_id=examplekms.arn)
        ```
        ### Server Side Encryption with S3 Default Master Key

        ```python
        import pulumi
        import pulumi_aws as aws

        examplebucket = aws.s3.Bucket("examplebucket", acl="private")
        examplebucket_object = aws.s3.BucketObject("examplebucketObject",
            key="someobject",
            bucket=examplebucket.id,
            source=pulumi.FileAsset("index.html"),
            server_side_encryption="aws:kms")
        ```
        ### Server Side Encryption with AWS-Managed Key

        ```python
        import pulumi
        import pulumi_aws as aws

        examplebucket = aws.s3.Bucket("examplebucket", acl="private")
        examplebucket_object = aws.s3.BucketObject("examplebucketObject",
            key="someobject",
            bucket=examplebucket.id,
            source=pulumi.FileAsset("index.html"),
            server_side_encryption="AES256")
        ```
        ### S3 Object Lock

        ```python
        import pulumi
        import pulumi_aws as aws

        examplebucket = aws.s3.Bucket("examplebucket",
            acl="private",
            versioning=aws.s3.BucketVersioningArgs(
                enabled=True,
            ),
            object_lock_configuration=aws.s3.BucketObjectLockConfigurationArgs(
                object_lock_enabled="Enabled",
            ))
        examplebucket_object = aws.s3.BucketObject("examplebucketObject",
            key="someobject",
            bucket=examplebucket.id,
            source=pulumi.FileAsset("important.txt"),
            object_lock_legal_hold_status="ON",
            object_lock_mode="GOVERNANCE",
            object_lock_retain_until_date="2021-12-31T23:59:60Z",
            force_destroy=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] acl: The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Valid values are `private`, `public-read`, `public-read-write`, `aws-exec-read`, `authenticated-read`, `bucket-owner-read`, and `bucket-owner-full-control`. Defaults to `private`.
        :param pulumi.Input[str] bucket: The name of the bucket to put the file in. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified.
        :param pulumi.Input[bool] bucket_key_enabled: Whether or not to use [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) for SSE-KMS.
        :param pulumi.Input[str] cache_control: Specifies caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
        :param pulumi.Input[str] content: Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
        :param pulumi.Input[str] content_base64: Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
        :param pulumi.Input[str] content_disposition: Specifies presentational information for the object. Read [w3c content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
        :param pulumi.Input[str] content_encoding: Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
        :param pulumi.Input[str] content_language: The language the content is in e.g. en-US or en-GB.
        :param pulumi.Input[str] content_type: A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
        :param pulumi.Input[str] etag: Used to trigger updates. The only meaningful value is `${filemd5("path/to/file")}` (this provider 0.11.12 or later) or `${md5(file("path/to/file"))}` (this provider 0.11.11 or earlier).
               This attribute is not compatible with KMS encryption, `kms_key_id` or `server_side_encryption = "aws:kms"`.
        :param pulumi.Input[bool] force_destroy: Allow the object to be deleted by removing any legal hold on any object version.
               Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
        :param pulumi.Input[str] key: The name of the object once it is in the bucket.
        :param pulumi.Input[str] kms_key_id: Amazon Resource Name (ARN) of the KMS Key to use for object encryption. If the S3 Bucket has server-side encryption enabled, that value will automatically be used. If referencing the
               `kms.Key` resource, use the `arn` attribute. If referencing the `kms.Alias` data source or resource, use the `target_key_arn` attribute. The provider will only perform drift detection if a configuration value
               is provided.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: A map of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
        :param pulumi.Input[str] object_lock_legal_hold_status: The [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
        :param pulumi.Input[str] object_lock_mode: The object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
        :param pulumi.Input[str] object_lock_retain_until_date: The date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
        :param pulumi.Input[str] server_side_encryption: Specifies server-side encryption of the object in S3. Valid values are "`AES256`" and "`aws:kms`".
        :param pulumi.Input[Union[pulumi.Asset, pulumi.Archive]] source: The path to a file that will be read and uploaded as raw bytes for the object content.
        :param pulumi.Input[str] storage_class: Specifies the desired [Storage Class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
               for the object. Can be either "`STANDARD`", "`REDUCED_REDUNDANCY`", "`ONEZONE_IA`", "`INTELLIGENT_TIERING`", "`GLACIER`", "`DEEP_ARCHIVE`", or "`STANDARD_IA`". Defaults to "`STANDARD`".
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the object.
        :param pulumi.Input[str] website_redirect: Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
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

            __props__['acl'] = acl
            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__['bucket'] = bucket
            __props__['bucket_key_enabled'] = bucket_key_enabled
            __props__['cache_control'] = cache_control
            __props__['content'] = content
            __props__['content_base64'] = content_base64
            __props__['content_disposition'] = content_disposition
            __props__['content_encoding'] = content_encoding
            __props__['content_language'] = content_language
            __props__['content_type'] = content_type
            __props__['etag'] = etag
            __props__['force_destroy'] = force_destroy
            __props__['key'] = key
            __props__['kms_key_id'] = kms_key_id
            __props__['metadata'] = metadata
            __props__['object_lock_legal_hold_status'] = object_lock_legal_hold_status
            __props__['object_lock_mode'] = object_lock_mode
            __props__['object_lock_retain_until_date'] = object_lock_retain_until_date
            __props__['server_side_encryption'] = server_side_encryption
            __props__['source'] = source
            __props__['storage_class'] = storage_class
            __props__['tags'] = tags
            __props__['website_redirect'] = website_redirect
            __props__['version_id'] = None
        super(BucketObject, __self__).__init__(
            'aws:s3/bucketObject:BucketObject',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl: Optional[pulumi.Input[str]] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            bucket_key_enabled: Optional[pulumi.Input[bool]] = None,
            cache_control: Optional[pulumi.Input[str]] = None,
            content: Optional[pulumi.Input[str]] = None,
            content_base64: Optional[pulumi.Input[str]] = None,
            content_disposition: Optional[pulumi.Input[str]] = None,
            content_encoding: Optional[pulumi.Input[str]] = None,
            content_language: Optional[pulumi.Input[str]] = None,
            content_type: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            force_destroy: Optional[pulumi.Input[bool]] = None,
            key: Optional[pulumi.Input[str]] = None,
            kms_key_id: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            object_lock_legal_hold_status: Optional[pulumi.Input[str]] = None,
            object_lock_mode: Optional[pulumi.Input[str]] = None,
            object_lock_retain_until_date: Optional[pulumi.Input[str]] = None,
            server_side_encryption: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]] = None,
            storage_class: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            version_id: Optional[pulumi.Input[str]] = None,
            website_redirect: Optional[pulumi.Input[str]] = None) -> 'BucketObject':
        """
        Get an existing BucketObject resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] acl: The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Valid values are `private`, `public-read`, `public-read-write`, `aws-exec-read`, `authenticated-read`, `bucket-owner-read`, and `bucket-owner-full-control`. Defaults to `private`.
        :param pulumi.Input[str] bucket: The name of the bucket to put the file in. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified.
        :param pulumi.Input[bool] bucket_key_enabled: Whether or not to use [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) for SSE-KMS.
        :param pulumi.Input[str] cache_control: Specifies caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
        :param pulumi.Input[str] content: Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
        :param pulumi.Input[str] content_base64: Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
        :param pulumi.Input[str] content_disposition: Specifies presentational information for the object. Read [w3c content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
        :param pulumi.Input[str] content_encoding: Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
        :param pulumi.Input[str] content_language: The language the content is in e.g. en-US or en-GB.
        :param pulumi.Input[str] content_type: A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
        :param pulumi.Input[str] etag: Used to trigger updates. The only meaningful value is `${filemd5("path/to/file")}` (this provider 0.11.12 or later) or `${md5(file("path/to/file"))}` (this provider 0.11.11 or earlier).
               This attribute is not compatible with KMS encryption, `kms_key_id` or `server_side_encryption = "aws:kms"`.
        :param pulumi.Input[bool] force_destroy: Allow the object to be deleted by removing any legal hold on any object version.
               Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
        :param pulumi.Input[str] key: The name of the object once it is in the bucket.
        :param pulumi.Input[str] kms_key_id: Amazon Resource Name (ARN) of the KMS Key to use for object encryption. If the S3 Bucket has server-side encryption enabled, that value will automatically be used. If referencing the
               `kms.Key` resource, use the `arn` attribute. If referencing the `kms.Alias` data source or resource, use the `target_key_arn` attribute. The provider will only perform drift detection if a configuration value
               is provided.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: A map of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
        :param pulumi.Input[str] object_lock_legal_hold_status: The [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
        :param pulumi.Input[str] object_lock_mode: The object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
        :param pulumi.Input[str] object_lock_retain_until_date: The date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
        :param pulumi.Input[str] server_side_encryption: Specifies server-side encryption of the object in S3. Valid values are "`AES256`" and "`aws:kms`".
        :param pulumi.Input[Union[pulumi.Asset, pulumi.Archive]] source: The path to a file that will be read and uploaded as raw bytes for the object content.
        :param pulumi.Input[str] storage_class: Specifies the desired [Storage Class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
               for the object. Can be either "`STANDARD`", "`REDUCED_REDUNDANCY`", "`ONEZONE_IA`", "`INTELLIGENT_TIERING`", "`GLACIER`", "`DEEP_ARCHIVE`", or "`STANDARD_IA`". Defaults to "`STANDARD`".
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the object.
        :param pulumi.Input[str] version_id: A unique version ID value for the object, if bucket versioning
               is enabled.
        :param pulumi.Input[str] website_redirect: Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["acl"] = acl
        __props__["bucket"] = bucket
        __props__["bucket_key_enabled"] = bucket_key_enabled
        __props__["cache_control"] = cache_control
        __props__["content"] = content
        __props__["content_base64"] = content_base64
        __props__["content_disposition"] = content_disposition
        __props__["content_encoding"] = content_encoding
        __props__["content_language"] = content_language
        __props__["content_type"] = content_type
        __props__["etag"] = etag
        __props__["force_destroy"] = force_destroy
        __props__["key"] = key
        __props__["kms_key_id"] = kms_key_id
        __props__["metadata"] = metadata
        __props__["object_lock_legal_hold_status"] = object_lock_legal_hold_status
        __props__["object_lock_mode"] = object_lock_mode
        __props__["object_lock_retain_until_date"] = object_lock_retain_until_date
        __props__["server_side_encryption"] = server_side_encryption
        __props__["source"] = source
        __props__["storage_class"] = storage_class
        __props__["tags"] = tags
        __props__["version_id"] = version_id
        __props__["website_redirect"] = website_redirect
        return BucketObject(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acl(self) -> pulumi.Output[Optional[str]]:
        """
        The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Valid values are `private`, `public-read`, `public-read-write`, `aws-exec-read`, `authenticated-read`, `bucket-owner-read`, and `bucket-owner-full-control`. Defaults to `private`.
        """
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        The name of the bucket to put the file in. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="bucketKeyEnabled")
    def bucket_key_enabled(self) -> pulumi.Output[bool]:
        """
        Whether or not to use [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) for SSE-KMS.
        """
        return pulumi.get(self, "bucket_key_enabled")

    @property
    @pulumi.getter(name="cacheControl")
    def cache_control(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
        """
        return pulumi.get(self, "cache_control")

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[Optional[str]]:
        """
        Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contentBase64")
    def content_base64(self) -> pulumi.Output[Optional[str]]:
        """
        Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
        """
        return pulumi.get(self, "content_base64")

    @property
    @pulumi.getter(name="contentDisposition")
    def content_disposition(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies presentational information for the object. Read [w3c content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
        """
        return pulumi.get(self, "content_disposition")

    @property
    @pulumi.getter(name="contentEncoding")
    def content_encoding(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
        """
        return pulumi.get(self, "content_encoding")

    @property
    @pulumi.getter(name="contentLanguage")
    def content_language(self) -> pulumi.Output[Optional[str]]:
        """
        The language the content is in e.g. en-US or en-GB.
        """
        return pulumi.get(self, "content_language")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> pulumi.Output[str]:
        """
        A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        Used to trigger updates. The only meaningful value is `${filemd5("path/to/file")}` (this provider 0.11.12 or later) or `${md5(file("path/to/file"))}` (this provider 0.11.11 or earlier).
        This attribute is not compatible with KMS encryption, `kms_key_id` or `server_side_encryption = "aws:kms"`.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        Allow the object to be deleted by removing any legal hold on any object version.
        Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
        """
        return pulumi.get(self, "force_destroy")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        The name of the object once it is in the bucket.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the KMS Key to use for object encryption. If the S3 Bucket has server-side encryption enabled, that value will automatically be used. If referencing the
        `kms.Key` resource, use the `arn` attribute. If referencing the `kms.Alias` data source or resource, use the `target_key_arn` attribute. The provider will only perform drift detection if a configuration value
        is provided.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="objectLockLegalHoldStatus")
    def object_lock_legal_hold_status(self) -> pulumi.Output[Optional[str]]:
        """
        The [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
        """
        return pulumi.get(self, "object_lock_legal_hold_status")

    @property
    @pulumi.getter(name="objectLockMode")
    def object_lock_mode(self) -> pulumi.Output[Optional[str]]:
        """
        The object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
        """
        return pulumi.get(self, "object_lock_mode")

    @property
    @pulumi.getter(name="objectLockRetainUntilDate")
    def object_lock_retain_until_date(self) -> pulumi.Output[Optional[str]]:
        """
        The date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
        """
        return pulumi.get(self, "object_lock_retain_until_date")

    @property
    @pulumi.getter(name="serverSideEncryption")
    def server_side_encryption(self) -> pulumi.Output[str]:
        """
        Specifies server-side encryption of the object in S3. Valid values are "`AES256`" and "`aws:kms`".
        """
        return pulumi.get(self, "server_side_encryption")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[Optional[Union[pulumi.Asset, pulumi.Archive]]]:
        """
        The path to a file that will be read and uploaded as raw bytes for the object content.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> pulumi.Output[str]:
        """
        Specifies the desired [Storage Class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
        for the object. Can be either "`STANDARD`", "`REDUCED_REDUNDANCY`", "`ONEZONE_IA`", "`INTELLIGENT_TIERING`", "`GLACIER`", "`DEEP_ARCHIVE`", or "`STANDARD_IA`". Defaults to "`STANDARD`".
        """
        return pulumi.get(self, "storage_class")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the object.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="versionId")
    def version_id(self) -> pulumi.Output[str]:
        """
        A unique version ID value for the object, if bucket versioning
        is enabled.
        """
        return pulumi.get(self, "version_id")

    @property
    @pulumi.getter(name="websiteRedirect")
    def website_redirect(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
        """
        return pulumi.get(self, "website_redirect")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

