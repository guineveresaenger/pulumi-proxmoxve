// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package permission

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRoles(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetRolesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetRolesResult
	err := ctx.Invoke("proxmoxve:Permission/getRoles:getRoles", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getRoles.
type GetRolesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string     `pulumi:"id"`
	Privileges [][]string `pulumi:"privileges"`
	RoleIds    []string   `pulumi:"roleIds"`
	Specials   []bool     `pulumi:"specials"`
}
