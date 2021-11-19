// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package proxmoxve

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterIPSet struct {
	pulumi.CustomResourceState

	// List of IP or Networks
	Cidrs ClusterIPSetCidrArrayOutput `pulumi:"cidrs"`
	// IPSet comment
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// IPSet name
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewClusterIPSet registers a new resource with the given unique name, arguments, and options.
func NewClusterIPSet(ctx *pulumi.Context,
	name string, args *ClusterIPSetArgs, opts ...pulumi.ResourceOption) (*ClusterIPSet, error) {
	if args == nil {
		args = &ClusterIPSetArgs{}
	}

	var resource ClusterIPSet
	err := ctx.RegisterResource("proxmoxve:index/clusterIPSet:ClusterIPSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterIPSet gets an existing ClusterIPSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterIPSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterIPSetState, opts ...pulumi.ResourceOption) (*ClusterIPSet, error) {
	var resource ClusterIPSet
	err := ctx.ReadResource("proxmoxve:index/clusterIPSet:ClusterIPSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterIPSet resources.
type clusterIPSetState struct {
	// List of IP or Networks
	Cidrs []ClusterIPSetCidr `pulumi:"cidrs"`
	// IPSet comment
	Comment *string `pulumi:"comment"`
	// IPSet name
	Name *string `pulumi:"name"`
}

type ClusterIPSetState struct {
	// List of IP or Networks
	Cidrs ClusterIPSetCidrArrayInput
	// IPSet comment
	Comment pulumi.StringPtrInput
	// IPSet name
	Name pulumi.StringPtrInput
}

func (ClusterIPSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterIPSetState)(nil)).Elem()
}

type clusterIPSetArgs struct {
	// List of IP or Networks
	Cidrs []ClusterIPSetCidr `pulumi:"cidrs"`
	// IPSet comment
	Comment *string `pulumi:"comment"`
	// IPSet name
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ClusterIPSet resource.
type ClusterIPSetArgs struct {
	// List of IP or Networks
	Cidrs ClusterIPSetCidrArrayInput
	// IPSet comment
	Comment pulumi.StringPtrInput
	// IPSet name
	Name pulumi.StringPtrInput
}

func (ClusterIPSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterIPSetArgs)(nil)).Elem()
}

type ClusterIPSetInput interface {
	pulumi.Input

	ToClusterIPSetOutput() ClusterIPSetOutput
	ToClusterIPSetOutputWithContext(ctx context.Context) ClusterIPSetOutput
}

func (*ClusterIPSet) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterIPSet)(nil))
}

func (i *ClusterIPSet) ToClusterIPSetOutput() ClusterIPSetOutput {
	return i.ToClusterIPSetOutputWithContext(context.Background())
}

func (i *ClusterIPSet) ToClusterIPSetOutputWithContext(ctx context.Context) ClusterIPSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIPSetOutput)
}

func (i *ClusterIPSet) ToClusterIPSetPtrOutput() ClusterIPSetPtrOutput {
	return i.ToClusterIPSetPtrOutputWithContext(context.Background())
}

func (i *ClusterIPSet) ToClusterIPSetPtrOutputWithContext(ctx context.Context) ClusterIPSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIPSetPtrOutput)
}

type ClusterIPSetPtrInput interface {
	pulumi.Input

	ToClusterIPSetPtrOutput() ClusterIPSetPtrOutput
	ToClusterIPSetPtrOutputWithContext(ctx context.Context) ClusterIPSetPtrOutput
}

type clusterIPSetPtrType ClusterIPSetArgs

func (*clusterIPSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterIPSet)(nil))
}

func (i *clusterIPSetPtrType) ToClusterIPSetPtrOutput() ClusterIPSetPtrOutput {
	return i.ToClusterIPSetPtrOutputWithContext(context.Background())
}

func (i *clusterIPSetPtrType) ToClusterIPSetPtrOutputWithContext(ctx context.Context) ClusterIPSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIPSetPtrOutput)
}

// ClusterIPSetArrayInput is an input type that accepts ClusterIPSetArray and ClusterIPSetArrayOutput values.
// You can construct a concrete instance of `ClusterIPSetArrayInput` via:
//
//          ClusterIPSetArray{ ClusterIPSetArgs{...} }
type ClusterIPSetArrayInput interface {
	pulumi.Input

	ToClusterIPSetArrayOutput() ClusterIPSetArrayOutput
	ToClusterIPSetArrayOutputWithContext(context.Context) ClusterIPSetArrayOutput
}

type ClusterIPSetArray []ClusterIPSetInput

func (ClusterIPSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterIPSet)(nil)).Elem()
}

func (i ClusterIPSetArray) ToClusterIPSetArrayOutput() ClusterIPSetArrayOutput {
	return i.ToClusterIPSetArrayOutputWithContext(context.Background())
}

func (i ClusterIPSetArray) ToClusterIPSetArrayOutputWithContext(ctx context.Context) ClusterIPSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIPSetArrayOutput)
}

// ClusterIPSetMapInput is an input type that accepts ClusterIPSetMap and ClusterIPSetMapOutput values.
// You can construct a concrete instance of `ClusterIPSetMapInput` via:
//
//          ClusterIPSetMap{ "key": ClusterIPSetArgs{...} }
type ClusterIPSetMapInput interface {
	pulumi.Input

	ToClusterIPSetMapOutput() ClusterIPSetMapOutput
	ToClusterIPSetMapOutputWithContext(context.Context) ClusterIPSetMapOutput
}

type ClusterIPSetMap map[string]ClusterIPSetInput

func (ClusterIPSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterIPSet)(nil)).Elem()
}

func (i ClusterIPSetMap) ToClusterIPSetMapOutput() ClusterIPSetMapOutput {
	return i.ToClusterIPSetMapOutputWithContext(context.Background())
}

func (i ClusterIPSetMap) ToClusterIPSetMapOutputWithContext(ctx context.Context) ClusterIPSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIPSetMapOutput)
}

type ClusterIPSetOutput struct{ *pulumi.OutputState }

func (ClusterIPSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterIPSet)(nil))
}

func (o ClusterIPSetOutput) ToClusterIPSetOutput() ClusterIPSetOutput {
	return o
}

func (o ClusterIPSetOutput) ToClusterIPSetOutputWithContext(ctx context.Context) ClusterIPSetOutput {
	return o
}

func (o ClusterIPSetOutput) ToClusterIPSetPtrOutput() ClusterIPSetPtrOutput {
	return o.ToClusterIPSetPtrOutputWithContext(context.Background())
}

func (o ClusterIPSetOutput) ToClusterIPSetPtrOutputWithContext(ctx context.Context) ClusterIPSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterIPSet) *ClusterIPSet {
		return &v
	}).(ClusterIPSetPtrOutput)
}

type ClusterIPSetPtrOutput struct{ *pulumi.OutputState }

func (ClusterIPSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterIPSet)(nil))
}

func (o ClusterIPSetPtrOutput) ToClusterIPSetPtrOutput() ClusterIPSetPtrOutput {
	return o
}

func (o ClusterIPSetPtrOutput) ToClusterIPSetPtrOutputWithContext(ctx context.Context) ClusterIPSetPtrOutput {
	return o
}

func (o ClusterIPSetPtrOutput) Elem() ClusterIPSetOutput {
	return o.ApplyT(func(v *ClusterIPSet) ClusterIPSet {
		if v != nil {
			return *v
		}
		var ret ClusterIPSet
		return ret
	}).(ClusterIPSetOutput)
}

type ClusterIPSetArrayOutput struct{ *pulumi.OutputState }

func (ClusterIPSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterIPSet)(nil))
}

func (o ClusterIPSetArrayOutput) ToClusterIPSetArrayOutput() ClusterIPSetArrayOutput {
	return o
}

func (o ClusterIPSetArrayOutput) ToClusterIPSetArrayOutputWithContext(ctx context.Context) ClusterIPSetArrayOutput {
	return o
}

func (o ClusterIPSetArrayOutput) Index(i pulumi.IntInput) ClusterIPSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterIPSet {
		return vs[0].([]ClusterIPSet)[vs[1].(int)]
	}).(ClusterIPSetOutput)
}

type ClusterIPSetMapOutput struct{ *pulumi.OutputState }

func (ClusterIPSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ClusterIPSet)(nil))
}

func (o ClusterIPSetMapOutput) ToClusterIPSetMapOutput() ClusterIPSetMapOutput {
	return o
}

func (o ClusterIPSetMapOutput) ToClusterIPSetMapOutputWithContext(ctx context.Context) ClusterIPSetMapOutput {
	return o
}

func (o ClusterIPSetMapOutput) MapIndex(k pulumi.StringInput) ClusterIPSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ClusterIPSet {
		return vs[0].(map[string]ClusterIPSet)[vs[1].(string)]
	}).(ClusterIPSetOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterIPSetOutput{})
	pulumi.RegisterOutputType(ClusterIPSetPtrOutput{})
	pulumi.RegisterOutputType(ClusterIPSetArrayOutput{})
	pulumi.RegisterOutputType(ClusterIPSetMapOutput{})
}
