# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetTimeResult',
    'AwaitableGetTimeResult',
    'get_time',
    'get_time_output',
]

@pulumi.output_type
class GetTimeResult:
    """
    A collection of values returned by getTime.
    """
    def __init__(__self__, id=None, local_time=None, node_name=None, time_zone=None, utc_time=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if local_time and not isinstance(local_time, str):
            raise TypeError("Expected argument 'local_time' to be a str")
        pulumi.set(__self__, "local_time", local_time)
        if node_name and not isinstance(node_name, str):
            raise TypeError("Expected argument 'node_name' to be a str")
        pulumi.set(__self__, "node_name", node_name)
        if time_zone and not isinstance(time_zone, str):
            raise TypeError("Expected argument 'time_zone' to be a str")
        pulumi.set(__self__, "time_zone", time_zone)
        if utc_time and not isinstance(utc_time, str):
            raise TypeError("Expected argument 'utc_time' to be a str")
        pulumi.set(__self__, "utc_time", utc_time)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="localTime")
    def local_time(self) -> str:
        return pulumi.get(self, "local_time")

    @property
    @pulumi.getter(name="nodeName")
    def node_name(self) -> str:
        return pulumi.get(self, "node_name")

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> str:
        return pulumi.get(self, "time_zone")

    @property
    @pulumi.getter(name="utcTime")
    def utc_time(self) -> str:
        return pulumi.get(self, "utc_time")


class AwaitableGetTimeResult(GetTimeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTimeResult(
            id=self.id,
            local_time=self.local_time,
            node_name=self.node_name,
            time_zone=self.time_zone,
            utc_time=self.utc_time)


def get_time(node_name: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTimeResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['nodeName'] = node_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('proxmoxve:index/getTime:getTime', __args__, opts=opts, typ=GetTimeResult).value

    return AwaitableGetTimeResult(
        id=__ret__.id,
        local_time=__ret__.local_time,
        node_name=__ret__.node_name,
        time_zone=__ret__.time_zone,
        utc_time=__ret__.utc_time)


@_utilities.lift_output_func(get_time)
def get_time_output(node_name: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTimeResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...