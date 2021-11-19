// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package proxmoxve

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTime(ctx *pulumi.Context, args *LookupTimeArgs, opts ...pulumi.InvokeOption) (*LookupTimeResult, error) {
	var rv LookupTimeResult
	err := ctx.Invoke("proxmoxve:index/getTime:getTime", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTime.
type LookupTimeArgs struct {
	NodeName string `pulumi:"nodeName"`
}

// A collection of values returned by getTime.
type LookupTimeResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	LocalTime string `pulumi:"localTime"`
	NodeName  string `pulumi:"nodeName"`
	TimeZone  string `pulumi:"timeZone"`
	UtcTime   string `pulumi:"utcTime"`
}

func LookupTimeOutput(ctx *pulumi.Context, args LookupTimeOutputArgs, opts ...pulumi.InvokeOption) LookupTimeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTimeResult, error) {
			args := v.(LookupTimeArgs)
			r, err := LookupTime(ctx, &args, opts...)
			return *r, err
		}).(LookupTimeResultOutput)
}

// A collection of arguments for invoking getTime.
type LookupTimeOutputArgs struct {
	NodeName pulumi.StringInput `pulumi:"nodeName"`
}

func (LookupTimeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTimeArgs)(nil)).Elem()
}

// A collection of values returned by getTime.
type LookupTimeResultOutput struct{ *pulumi.OutputState }

func (LookupTimeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTimeResult)(nil)).Elem()
}

func (o LookupTimeResultOutput) ToLookupTimeResultOutput() LookupTimeResultOutput {
	return o
}

func (o LookupTimeResultOutput) ToLookupTimeResultOutputWithContext(ctx context.Context) LookupTimeResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTimeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTimeResultOutput) LocalTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeResult) string { return v.LocalTime }).(pulumi.StringOutput)
}

func (o LookupTimeResultOutput) NodeName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeResult) string { return v.NodeName }).(pulumi.StringOutput)
}

func (o LookupTimeResultOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeResult) string { return v.TimeZone }).(pulumi.StringOutput)
}

func (o LookupTimeResultOutput) UtcTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeResult) string { return v.UtcTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTimeResultOutput{})
}
