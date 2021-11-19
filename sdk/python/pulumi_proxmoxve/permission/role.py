# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RoleArgs', 'Role']

@pulumi.input_type
class RoleArgs:
    def __init__(__self__, *,
                 privileges: pulumi.Input[Sequence[pulumi.Input[str]]],
                 role_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a Role resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] privileges: The role privileges
        :param pulumi.Input[str] role_id: The role id
        """
        pulumi.set(__self__, "privileges", privileges)
        pulumi.set(__self__, "role_id", role_id)

    @property
    @pulumi.getter
    def privileges(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The role privileges
        """
        return pulumi.get(self, "privileges")

    @privileges.setter
    def privileges(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "privileges", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Input[str]:
        """
        The role id
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_id", value)


@pulumi.input_type
class _RoleState:
    def __init__(__self__, *,
                 privileges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Role resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] privileges: The role privileges
        :param pulumi.Input[str] role_id: The role id
        """
        if privileges is not None:
            pulumi.set(__self__, "privileges", privileges)
        if role_id is not None:
            pulumi.set(__self__, "role_id", role_id)

    @property
    @pulumi.getter
    def privileges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The role privileges
        """
        return pulumi.get(self, "privileges")

    @privileges.setter
    def privileges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "privileges", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[str]]:
        """
        The role id
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_id", value)


class Role(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 privileges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Role resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] privileges: The role privileges
        :param pulumi.Input[str] role_id: The role id
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Role resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 privileges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = RoleArgs.__new__(RoleArgs)

            if privileges is None and not opts.urn:
                raise TypeError("Missing required property 'privileges'")
            __props__.__dict__["privileges"] = privileges
            if role_id is None and not opts.urn:
                raise TypeError("Missing required property 'role_id'")
            __props__.__dict__["role_id"] = role_id
        super(Role, __self__).__init__(
            'proxmoxve:Permission/role:Role',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            privileges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            role_id: Optional[pulumi.Input[str]] = None) -> 'Role':
        """
        Get an existing Role resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] privileges: The role privileges
        :param pulumi.Input[str] role_id: The role id
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RoleState.__new__(_RoleState)

        __props__.__dict__["privileges"] = privileges
        __props__.__dict__["role_id"] = role_id
        return Role(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def privileges(self) -> pulumi.Output[Sequence[str]]:
        """
        The role privileges
        """
        return pulumi.get(self, "privileges")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Output[str]:
        """
        The role id
        """
        return pulumi.get(self, "role_id")

