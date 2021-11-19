# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .certifi import *
from .dns import *
from .get_dns import *
from .get_hosts import *
from .get_time import *
from .get_version import *
from .hosts import *
from .provider import *
from .time import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_proxmoxve.cluster as __cluster
    cluster = __cluster
    import pulumi_proxmoxve.config as __config
    config = __config
    import pulumi_proxmoxve.ct as __ct
    ct = __ct
    import pulumi_proxmoxve.permission as __permission
    permission = __permission
    import pulumi_proxmoxve.storage as __storage
    storage = __storage
    import pulumi_proxmoxve.vm as __vm
    vm = __vm
else:
    cluster = _utilities.lazy_import('pulumi_proxmoxve.cluster')
    config = _utilities.lazy_import('pulumi_proxmoxve.config')
    ct = _utilities.lazy_import('pulumi_proxmoxve.ct')
    permission = _utilities.lazy_import('pulumi_proxmoxve.permission')
    storage = _utilities.lazy_import('pulumi_proxmoxve.storage')
    vm = _utilities.lazy_import('pulumi_proxmoxve.vm')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "proxmoxve",
  "mod": "CT/container",
  "fqn": "pulumi_proxmoxve.ct",
  "classes": {
   "proxmoxve:CT/container:Container": "Container"
  }
 },
 {
  "pkg": "proxmoxve",
  "mod": "Cluster/clusterAlias",
  "fqn": "pulumi_proxmoxve.cluster",
  "classes": {
   "proxmoxve:Cluster/clusterAlias:ClusterAlias": "ClusterAlias"
  }
 },
 {
  "pkg": "proxmoxve",
  "mod": "Cluster/clusterIPSet",
  "fqn": "pulumi_proxmoxve.cluster",
  "classes": {
   "proxmoxve:Cluster/clusterIPSet:ClusterIPSet": "ClusterIPSet"
  }
 },
 {
  "pkg": "proxmoxve",
  "mod": "Permission/group",
  "fqn": "pulumi_proxmoxve.permission",
  "classes": {
   "proxmoxve:Permission/group:Group": "Group"
  }
 },
 {
  "pkg": "proxmoxve",
  "mod": "Permission/pool",
  "fqn": "pulumi_proxmoxve.permission",
  "classes": {
   "proxmoxve:Permission/pool:Pool": "Pool"
  }
 },
 {
  "pkg": "proxmoxve",
  "mod": "Permission/role",
  "fqn": "pulumi_proxmoxve.permission",
  "classes": {
   "proxmoxve:Permission/role:Role": "Role"
  }
 },
 {
  "pkg": "proxmoxve",
  "mod": "Permission/user",
  "fqn": "pulumi_proxmoxve.permission",
  "classes": {
   "proxmoxve:Permission/user:User": "User"
  }
 },
 {
  "pkg": "proxmoxve",
  "mod": "Storage/file",
  "fqn": "pulumi_proxmoxve.storage",
  "classes": {
   "proxmoxve:Storage/file:File": "File"
  }
 },
 {
  "pkg": "proxmoxve",
  "mod": "VM/virtualMachine",
  "fqn": "pulumi_proxmoxve.vm",
  "classes": {
   "proxmoxve:VM/virtualMachine:VirtualMachine": "VirtualMachine"
  }
 },
 {
  "pkg": "proxmoxve",
  "mod": "index/certifi",
  "fqn": "pulumi_proxmoxve",
  "classes": {
   "proxmoxve:index/certifi:Certifi": "Certifi"
  }
 },
 {
  "pkg": "proxmoxve",
  "mod": "index/dNS",
  "fqn": "pulumi_proxmoxve",
  "classes": {
   "proxmoxve:index/dNS:DNS": "DNS"
  }
 },
 {
  "pkg": "proxmoxve",
  "mod": "index/hosts",
  "fqn": "pulumi_proxmoxve",
  "classes": {
   "proxmoxve:index/hosts:Hosts": "Hosts"
  }
 },
 {
  "pkg": "proxmoxve",
  "mod": "index/time",
  "fqn": "pulumi_proxmoxve",
  "classes": {
   "proxmoxve:index/time:Time": "Time"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "proxmoxve",
  "token": "pulumi:providers:proxmoxve",
  "fqn": "pulumi_proxmoxve",
  "class": "Provider"
 }
]
"""
)
