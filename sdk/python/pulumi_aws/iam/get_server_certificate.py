# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetServerCertificateResult',
    'AwaitableGetServerCertificateResult',
    'get_server_certificate',
]



@pulumi.output_type
class GetServerCertificateResult:
    """
    A collection of values returned by getServerCertificate.
    """
    def __init__(__self__, arn=None, certificate_body=None, certificate_chain=None, expiration_date=None, id=None, latest=None, name=None, name_prefix=None, path=None, path_prefix=None, upload_date=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if certificate_body and not isinstance(certificate_body, str):
            raise TypeError("Expected argument 'certificate_body' to be a str")
        pulumi.set(__self__, "certificate_body", certificate_body)
        if certificate_chain and not isinstance(certificate_chain, str):
            raise TypeError("Expected argument 'certificate_chain' to be a str")
        pulumi.set(__self__, "certificate_chain", certificate_chain)
        if expiration_date and not isinstance(expiration_date, str):
            raise TypeError("Expected argument 'expiration_date' to be a str")
        pulumi.set(__self__, "expiration_date", expiration_date)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if latest and not isinstance(latest, bool):
            raise TypeError("Expected argument 'latest' to be a bool")
        pulumi.set(__self__, "latest", latest)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_prefix and not isinstance(name_prefix, str):
            raise TypeError("Expected argument 'name_prefix' to be a str")
        pulumi.set(__self__, "name_prefix", name_prefix)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if path_prefix and not isinstance(path_prefix, str):
            raise TypeError("Expected argument 'path_prefix' to be a str")
        pulumi.set(__self__, "path_prefix", path_prefix)
        if upload_date and not isinstance(upload_date, str):
            raise TypeError("Expected argument 'upload_date' to be a str")
        pulumi.set(__self__, "upload_date", upload_date)

    @property
    @pulumi.getter
    def arn(self) -> str:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="certificateBody")
    def certificate_body(self) -> str:
        return pulumi.get(self, "certificate_body")

    @property
    @pulumi.getter(name="certificateChain")
    def certificate_chain(self) -> str:
        return pulumi.get(self, "certificate_chain")

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> str:
        return pulumi.get(self, "expiration_date")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def latest(self) -> Optional[bool]:
        return pulumi.get(self, "latest")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[str]:
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter
    def path(self) -> str:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="pathPrefix")
    def path_prefix(self) -> Optional[str]:
        return pulumi.get(self, "path_prefix")

    @property
    @pulumi.getter(name="uploadDate")
    def upload_date(self) -> str:
        return pulumi.get(self, "upload_date")



class AwaitableGetServerCertificateResult(GetServerCertificateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerCertificateResult(
            arn=self.arn,
            certificate_body=self.certificate_body,
            certificate_chain=self.certificate_chain,
            expiration_date=self.expiration_date,
            id=self.id,
            latest=self.latest,
            name=self.name,
            name_prefix=self.name_prefix,
            path=self.path,
            path_prefix=self.path_prefix,
            upload_date=self.upload_date)


def get_server_certificate(latest: Optional[bool] = None,
                           name: Optional[str] = None,
                           name_prefix: Optional[str] = None,
                           path_prefix: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServerCertificateResult:
    """
    Use this data source to lookup information about IAM Server Certificates.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    my_domain = aws.iam.get_server_certificate(latest=True,
        name_prefix="my-domain.org")
    elb = aws.elb.LoadBalancer("elb", listeners=[{
        "instance_port": 8000,
        "instanceProtocol": "https",
        "lb_port": 443,
        "lbProtocol": "https",
        "sslCertificateId": my_domain.arn,
    }])
    ```
    ## Import

    The import function will read in certificate body, certificate chain (if it exists), id, name, path, and arn.
    It will not retrieve the private key which is not available through the AWS API.


    :param bool latest: sort results by expiration date. returns the certificate with expiration date in furthest in the future.
    :param str name: exact name of the cert to lookup
    :param str name_prefix: prefix of cert to filter by
    :param str path_prefix: prefix of path to filter by
    """
    __args__ = dict()
    __args__['latest'] = latest
    __args__['name'] = name
    __args__['namePrefix'] = name_prefix
    __args__['pathPrefix'] = path_prefix
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:iam/getServerCertificate:getServerCertificate', __args__, opts=opts, typ=GetServerCertificateResult).value

    return AwaitableGetServerCertificateResult(
        arn=__ret__.arn,
        certificate_body=__ret__.certificate_body,
        certificate_chain=__ret__.certificate_chain,
        expiration_date=__ret__.expiration_date,
        id=__ret__.id,
        latest=__ret__.latest,
        name=__ret__.name,
        name_prefix=__ret__.name_prefix,
        path=__ret__.path,
        path_prefix=__ret__.path_prefix,
        upload_date=__ret__.upload_date)
