# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetClusterAliasesResult',
    'AwaitableGetClusterAliasesResult',
    'get_cluster_aliases',
]

@pulumi.output_type
class GetClusterAliasesResult:
    """
    A collection of values returned by getClusterAliases.
    """
    def __init__(__self__, alias_ids=None, id=None):
        if alias_ids and not isinstance(alias_ids, list):
            raise TypeError("Expected argument 'alias_ids' to be a list")
        pulumi.set(__self__, "alias_ids", alias_ids)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="aliasIds")
    def alias_ids(self) -> Sequence[str]:
        return pulumi.get(self, "alias_ids")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetClusterAliasesResult(GetClusterAliasesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterAliasesResult(
            alias_ids=self.alias_ids,
            id=self.id)


def get_cluster_aliases(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterAliasesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('proxmoxve:Cluster/getClusterAliases:getClusterAliases', __args__, opts=opts, typ=GetClusterAliasesResult).value

    return AwaitableGetClusterAliasesResult(
        alias_ids=__ret__.alias_ids,
        id=__ret__.id)
