// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualEnvironment struct {
	Endpoint *string `pulumi:"endpoint"`
	Insecure *bool   `pulumi:"insecure"`
	Otp      *string `pulumi:"otp"`
	Password *string `pulumi:"password"`
	Username *string `pulumi:"username"`
}

// VirtualEnvironmentInput is an input type that accepts VirtualEnvironmentArgs and VirtualEnvironmentOutput values.
// You can construct a concrete instance of `VirtualEnvironmentInput` via:
//
//          VirtualEnvironmentArgs{...}
type VirtualEnvironmentInput interface {
	pulumi.Input

	ToVirtualEnvironmentOutput() VirtualEnvironmentOutput
	ToVirtualEnvironmentOutputWithContext(context.Context) VirtualEnvironmentOutput
}

type VirtualEnvironmentArgs struct {
	Endpoint pulumi.StringPtrInput `pulumi:"endpoint"`
	Insecure pulumi.BoolPtrInput   `pulumi:"insecure"`
	Otp      pulumi.StringPtrInput `pulumi:"otp"`
	Password pulumi.StringPtrInput `pulumi:"password"`
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (VirtualEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualEnvironment)(nil)).Elem()
}

func (i VirtualEnvironmentArgs) ToVirtualEnvironmentOutput() VirtualEnvironmentOutput {
	return i.ToVirtualEnvironmentOutputWithContext(context.Background())
}

func (i VirtualEnvironmentArgs) ToVirtualEnvironmentOutputWithContext(ctx context.Context) VirtualEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualEnvironmentOutput)
}

type VirtualEnvironmentOutput struct{ *pulumi.OutputState }

func (VirtualEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualEnvironment)(nil)).Elem()
}

func (o VirtualEnvironmentOutput) ToVirtualEnvironmentOutput() VirtualEnvironmentOutput {
	return o
}

func (o VirtualEnvironmentOutput) ToVirtualEnvironmentOutputWithContext(ctx context.Context) VirtualEnvironmentOutput {
	return o
}

func (o VirtualEnvironmentOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualEnvironment) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o VirtualEnvironmentOutput) Insecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualEnvironment) *bool { return v.Insecure }).(pulumi.BoolPtrOutput)
}

func (o VirtualEnvironmentOutput) Otp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualEnvironment) *string { return v.Otp }).(pulumi.StringPtrOutput)
}

func (o VirtualEnvironmentOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualEnvironment) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o VirtualEnvironmentOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualEnvironment) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualEnvironmentInput)(nil)).Elem(), VirtualEnvironmentArgs{})
	pulumi.RegisterOutputType(VirtualEnvironmentOutput{})
}
