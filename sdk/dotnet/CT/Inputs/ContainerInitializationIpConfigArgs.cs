// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ProxmoxVE.CT.Inputs
{

    public sealed class ContainerInitializationIpConfigArgs : Pulumi.ResourceArgs
    {
        [Input("ipv4")]
        public Input<Inputs.ContainerInitializationIpConfigIpv4Args>? Ipv4 { get; set; }

        [Input("ipv6")]
        public Input<Inputs.ContainerInitializationIpConfigIpv6Args>? Ipv6 { get; set; }

        public ContainerInitializationIpConfigArgs()
        {
        }
    }
}
