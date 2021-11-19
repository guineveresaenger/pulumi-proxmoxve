// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ProxmoxVE.CT.Outputs
{

    [OutputType]
    public sealed class ContainerNetworkInterface
    {
        public readonly string? Bridge;
        public readonly bool? Enabled;
        public readonly string? MacAddress;
        public readonly string Name;
        public readonly double? RateLimit;
        public readonly int? VlanId;

        [OutputConstructor]
        private ContainerNetworkInterface(
            string? bridge,

            bool? enabled,

            string? macAddress,

            string name,

            double? rateLimit,

            int? vlanId)
        {
            Bridge = bridge;
            Enabled = enabled;
            MacAddress = macAddress;
            Name = name;
            RateLimit = rateLimit;
            VlanId = vlanId;
        }
    }
}