# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetSigningJobResult',
    'AwaitableGetSigningJobResult',
    'get_signing_job',
]

@pulumi.output_type
class GetSigningJobResult:
    """
    A collection of values returned by getSigningJob.
    """
    def __init__(__self__, completed_at=None, created_at=None, id=None, job_id=None, job_invoker=None, job_owner=None, platform_display_name=None, platform_id=None, profile_name=None, profile_version=None, requested_by=None, revocation_records=None, signature_expires_at=None, signed_objects=None, sources=None, status=None, status_reason=None):
        if completed_at and not isinstance(completed_at, str):
            raise TypeError("Expected argument 'completed_at' to be a str")
        pulumi.set(__self__, "completed_at", completed_at)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if job_id and not isinstance(job_id, str):
            raise TypeError("Expected argument 'job_id' to be a str")
        pulumi.set(__self__, "job_id", job_id)
        if job_invoker and not isinstance(job_invoker, str):
            raise TypeError("Expected argument 'job_invoker' to be a str")
        pulumi.set(__self__, "job_invoker", job_invoker)
        if job_owner and not isinstance(job_owner, str):
            raise TypeError("Expected argument 'job_owner' to be a str")
        pulumi.set(__self__, "job_owner", job_owner)
        if platform_display_name and not isinstance(platform_display_name, str):
            raise TypeError("Expected argument 'platform_display_name' to be a str")
        pulumi.set(__self__, "platform_display_name", platform_display_name)
        if platform_id and not isinstance(platform_id, str):
            raise TypeError("Expected argument 'platform_id' to be a str")
        pulumi.set(__self__, "platform_id", platform_id)
        if profile_name and not isinstance(profile_name, str):
            raise TypeError("Expected argument 'profile_name' to be a str")
        pulumi.set(__self__, "profile_name", profile_name)
        if profile_version and not isinstance(profile_version, str):
            raise TypeError("Expected argument 'profile_version' to be a str")
        pulumi.set(__self__, "profile_version", profile_version)
        if requested_by and not isinstance(requested_by, str):
            raise TypeError("Expected argument 'requested_by' to be a str")
        pulumi.set(__self__, "requested_by", requested_by)
        if revocation_records and not isinstance(revocation_records, list):
            raise TypeError("Expected argument 'revocation_records' to be a list")
        pulumi.set(__self__, "revocation_records", revocation_records)
        if signature_expires_at and not isinstance(signature_expires_at, str):
            raise TypeError("Expected argument 'signature_expires_at' to be a str")
        pulumi.set(__self__, "signature_expires_at", signature_expires_at)
        if signed_objects and not isinstance(signed_objects, list):
            raise TypeError("Expected argument 'signed_objects' to be a list")
        pulumi.set(__self__, "signed_objects", signed_objects)
        if sources and not isinstance(sources, list):
            raise TypeError("Expected argument 'sources' to be a list")
        pulumi.set(__self__, "sources", sources)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if status_reason and not isinstance(status_reason, str):
            raise TypeError("Expected argument 'status_reason' to be a str")
        pulumi.set(__self__, "status_reason", status_reason)

    @property
    @pulumi.getter(name="completedAt")
    def completed_at(self) -> str:
        """
        Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the signing job was completed.
        """
        return pulumi.get(self, "completed_at")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the signing job was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> str:
        return pulumi.get(self, "job_id")

    @property
    @pulumi.getter(name="jobInvoker")
    def job_invoker(self) -> str:
        """
        The IAM entity that initiated the signing job.
        """
        return pulumi.get(self, "job_invoker")

    @property
    @pulumi.getter(name="jobOwner")
    def job_owner(self) -> str:
        """
        The AWS account ID of the job owner.
        """
        return pulumi.get(self, "job_owner")

    @property
    @pulumi.getter(name="platformDisplayName")
    def platform_display_name(self) -> str:
        """
        A human-readable name for the signing platform associated with the signing job.
        """
        return pulumi.get(self, "platform_display_name")

    @property
    @pulumi.getter(name="platformId")
    def platform_id(self) -> str:
        """
        The platform to which your signed code image will be distributed.
        """
        return pulumi.get(self, "platform_id")

    @property
    @pulumi.getter(name="profileName")
    def profile_name(self) -> str:
        """
        The name of the profile that initiated the signing operation.
        """
        return pulumi.get(self, "profile_name")

    @property
    @pulumi.getter(name="profileVersion")
    def profile_version(self) -> str:
        """
        The version of the signing profile used to initiate the signing job.
        """
        return pulumi.get(self, "profile_version")

    @property
    @pulumi.getter(name="requestedBy")
    def requested_by(self) -> str:
        """
        The IAM principal that requested the signing job.
        """
        return pulumi.get(self, "requested_by")

    @property
    @pulumi.getter(name="revocationRecords")
    def revocation_records(self) -> Sequence['outputs.GetSigningJobRevocationRecordResult']:
        """
        A revocation record if the signature generated by the signing job has been revoked. Contains a timestamp and the ID of the IAM entity that revoked the signature.
        """
        return pulumi.get(self, "revocation_records")

    @property
    @pulumi.getter(name="signatureExpiresAt")
    def signature_expires_at(self) -> str:
        """
        The time when the signature of a signing job expires.
        """
        return pulumi.get(self, "signature_expires_at")

    @property
    @pulumi.getter(name="signedObjects")
    def signed_objects(self) -> Sequence['outputs.GetSigningJobSignedObjectResult']:
        """
        Name of the S3 bucket where the signed code image is saved by code signing.
        """
        return pulumi.get(self, "signed_objects")

    @property
    @pulumi.getter
    def sources(self) -> Sequence['outputs.GetSigningJobSourceResult']:
        """
        The object that contains the name of your S3 bucket or your raw code.
        """
        return pulumi.get(self, "sources")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the signing job.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusReason")
    def status_reason(self) -> str:
        """
        String value that contains the status reason.
        """
        return pulumi.get(self, "status_reason")


class AwaitableGetSigningJobResult(GetSigningJobResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSigningJobResult(
            completed_at=self.completed_at,
            created_at=self.created_at,
            id=self.id,
            job_id=self.job_id,
            job_invoker=self.job_invoker,
            job_owner=self.job_owner,
            platform_display_name=self.platform_display_name,
            platform_id=self.platform_id,
            profile_name=self.profile_name,
            profile_version=self.profile_version,
            requested_by=self.requested_by,
            revocation_records=self.revocation_records,
            signature_expires_at=self.signature_expires_at,
            signed_objects=self.signed_objects,
            sources=self.sources,
            status=self.status,
            status_reason=self.status_reason)


def get_signing_job(job_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSigningJobResult:
    """
    Provides information about a Signer Signing Job.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    build_signing_job = aws.signer.get_signing_job(job_id="9ed7e5c3-b8d4-4da0-8459-44e0b068f7ee")
    ```


    :param str job_id: The ID of the signing job on output.
    """
    __args__ = dict()
    __args__['jobId'] = job_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:signer/getSigningJob:getSigningJob', __args__, opts=opts, typ=GetSigningJobResult).value

    return AwaitableGetSigningJobResult(
        completed_at=__ret__.completed_at,
        created_at=__ret__.created_at,
        id=__ret__.id,
        job_id=__ret__.job_id,
        job_invoker=__ret__.job_invoker,
        job_owner=__ret__.job_owner,
        platform_display_name=__ret__.platform_display_name,
        platform_id=__ret__.platform_id,
        profile_name=__ret__.profile_name,
        profile_version=__ret__.profile_version,
        requested_by=__ret__.requested_by,
        revocation_records=__ret__.revocation_records,
        signature_expires_at=__ret__.signature_expires_at,
        signed_objects=__ret__.signed_objects,
        sources=__ret__.sources,
        status=__ret__.status,
        status_reason=__ret__.status_reason)
