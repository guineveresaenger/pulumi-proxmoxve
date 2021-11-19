// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ProxmoxVE.Storage.Inputs
{

    public sealed class FileSourceFileGetArgs : Pulumi.ResourceArgs
    {
        [Input("changed")]
        public Input<bool>? Changed { get; set; }

        [Input("checksum")]
        public Input<string>? Checksum { get; set; }

        [Input("fileName")]
        public Input<string>? FileName { get; set; }

        [Input("insecure")]
        public Input<bool>? Insecure { get; set; }

        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public FileSourceFileGetArgs()
        {
        }
    }
}
