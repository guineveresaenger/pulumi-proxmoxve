// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package proxmoxve

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupClusterAlias(ctx *pulumi.Context, args *LookupClusterAliasArgs, opts ...pulumi.InvokeOption) (*LookupClusterAliasResult, error) {
	var rv LookupClusterAliasResult
	err := ctx.Invoke("proxmoxve:index/getClusterAlias:getClusterAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusterAlias.
type LookupClusterAliasArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by getClusterAlias.
type LookupClusterAliasResult struct {
	Cidr    string `pulumi:"cidr"`
	Comment string `pulumi:"comment"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

func LookupClusterAliasOutput(ctx *pulumi.Context, args LookupClusterAliasOutputArgs, opts ...pulumi.InvokeOption) LookupClusterAliasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterAliasResult, error) {
			args := v.(LookupClusterAliasArgs)
			r, err := LookupClusterAlias(ctx, &args, opts...)
			return *r, err
		}).(LookupClusterAliasResultOutput)
}

// A collection of arguments for invoking getClusterAlias.
type LookupClusterAliasOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupClusterAliasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterAliasArgs)(nil)).Elem()
}

// A collection of values returned by getClusterAlias.
type LookupClusterAliasResultOutput struct{ *pulumi.OutputState }

func (LookupClusterAliasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterAliasResult)(nil)).Elem()
}

func (o LookupClusterAliasResultOutput) ToLookupClusterAliasResultOutput() LookupClusterAliasResultOutput {
	return o
}

func (o LookupClusterAliasResultOutput) ToLookupClusterAliasResultOutputWithContext(ctx context.Context) LookupClusterAliasResultOutput {
	return o
}

func (o LookupClusterAliasResultOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterAliasResult) string { return v.Cidr }).(pulumi.StringOutput)
}

func (o LookupClusterAliasResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterAliasResult) string { return v.Comment }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupClusterAliasResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterAliasResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupClusterAliasResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterAliasResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterAliasResultOutput{})
}