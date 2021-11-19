// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package proxmoxve

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetClusterAliases(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetClusterAliasesResult, error) {
	var rv GetClusterAliasesResult
	err := ctx.Invoke("proxmoxve:index/getClusterAliases:getClusterAliases", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getClusterAliases.
type GetClusterAliasesResult struct {
	AliasIds []string `pulumi:"aliasIds"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
