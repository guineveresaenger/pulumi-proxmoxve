// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ct

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Container struct {
	pulumi.CustomResourceState

	// The cloning configuration
	Clone ContainerClonePtrOutput `pulumi:"clone"`
	// The console configuration
	Console ContainerConsolePtrOutput `pulumi:"console"`
	// The CPU allocation
	Cpu ContainerCpuPtrOutput `pulumi:"cpu"`
	// The description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The disks
	Disk ContainerDiskPtrOutput `pulumi:"disk"`
	// The initialization configuration
	Initialization ContainerInitializationPtrOutput `pulumi:"initialization"`
	// The memory allocation
	Memory ContainerMemoryPtrOutput `pulumi:"memory"`
	// The network interfaces
	NetworkInterfaces ContainerNetworkInterfaceArrayOutput `pulumi:"networkInterfaces"`
	// The node name
	NodeName pulumi.StringOutput `pulumi:"nodeName"`
	// The operating system configuration
	OperatingSystem ContainerOperatingSystemPtrOutput `pulumi:"operatingSystem"`
	// The ID of the pool to assign the container to
	PoolId pulumi.StringPtrOutput `pulumi:"poolId"`
	// Whether to start the container
	Started pulumi.BoolPtrOutput `pulumi:"started"`
	// Whether to create a template
	Template pulumi.BoolPtrOutput `pulumi:"template"`
	// The VM identifier
	VmId pulumi.IntPtrOutput `pulumi:"vmId"`
}

// NewContainer registers a new resource with the given unique name, arguments, and options.
func NewContainer(ctx *pulumi.Context,
	name string, args *ContainerArgs, opts ...pulumi.ResourceOption) (*Container, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NodeName == nil {
		return nil, errors.New("invalid value for required argument 'NodeName'")
	}
	var resource Container
	err := ctx.RegisterResource("proxmoxve:CT/container:Container", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainer gets an existing Container resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerState, opts ...pulumi.ResourceOption) (*Container, error) {
	var resource Container
	err := ctx.ReadResource("proxmoxve:CT/container:Container", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Container resources.
type containerState struct {
	// The cloning configuration
	Clone *ContainerClone `pulumi:"clone"`
	// The console configuration
	Console *ContainerConsole `pulumi:"console"`
	// The CPU allocation
	Cpu *ContainerCpu `pulumi:"cpu"`
	// The description
	Description *string `pulumi:"description"`
	// The disks
	Disk *ContainerDisk `pulumi:"disk"`
	// The initialization configuration
	Initialization *ContainerInitialization `pulumi:"initialization"`
	// The memory allocation
	Memory *ContainerMemory `pulumi:"memory"`
	// The network interfaces
	NetworkInterfaces []ContainerNetworkInterface `pulumi:"networkInterfaces"`
	// The node name
	NodeName *string `pulumi:"nodeName"`
	// The operating system configuration
	OperatingSystem *ContainerOperatingSystem `pulumi:"operatingSystem"`
	// The ID of the pool to assign the container to
	PoolId *string `pulumi:"poolId"`
	// Whether to start the container
	Started *bool `pulumi:"started"`
	// Whether to create a template
	Template *bool `pulumi:"template"`
	// The VM identifier
	VmId *int `pulumi:"vmId"`
}

type ContainerState struct {
	// The cloning configuration
	Clone ContainerClonePtrInput
	// The console configuration
	Console ContainerConsolePtrInput
	// The CPU allocation
	Cpu ContainerCpuPtrInput
	// The description
	Description pulumi.StringPtrInput
	// The disks
	Disk ContainerDiskPtrInput
	// The initialization configuration
	Initialization ContainerInitializationPtrInput
	// The memory allocation
	Memory ContainerMemoryPtrInput
	// The network interfaces
	NetworkInterfaces ContainerNetworkInterfaceArrayInput
	// The node name
	NodeName pulumi.StringPtrInput
	// The operating system configuration
	OperatingSystem ContainerOperatingSystemPtrInput
	// The ID of the pool to assign the container to
	PoolId pulumi.StringPtrInput
	// Whether to start the container
	Started pulumi.BoolPtrInput
	// Whether to create a template
	Template pulumi.BoolPtrInput
	// The VM identifier
	VmId pulumi.IntPtrInput
}

func (ContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerState)(nil)).Elem()
}

type containerArgs struct {
	// The cloning configuration
	Clone *ContainerClone `pulumi:"clone"`
	// The console configuration
	Console *ContainerConsole `pulumi:"console"`
	// The CPU allocation
	Cpu *ContainerCpu `pulumi:"cpu"`
	// The description
	Description *string `pulumi:"description"`
	// The disks
	Disk *ContainerDisk `pulumi:"disk"`
	// The initialization configuration
	Initialization *ContainerInitialization `pulumi:"initialization"`
	// The memory allocation
	Memory *ContainerMemory `pulumi:"memory"`
	// The network interfaces
	NetworkInterfaces []ContainerNetworkInterface `pulumi:"networkInterfaces"`
	// The node name
	NodeName string `pulumi:"nodeName"`
	// The operating system configuration
	OperatingSystem *ContainerOperatingSystem `pulumi:"operatingSystem"`
	// The ID of the pool to assign the container to
	PoolId *string `pulumi:"poolId"`
	// Whether to start the container
	Started *bool `pulumi:"started"`
	// Whether to create a template
	Template *bool `pulumi:"template"`
	// The VM identifier
	VmId *int `pulumi:"vmId"`
}

// The set of arguments for constructing a Container resource.
type ContainerArgs struct {
	// The cloning configuration
	Clone ContainerClonePtrInput
	// The console configuration
	Console ContainerConsolePtrInput
	// The CPU allocation
	Cpu ContainerCpuPtrInput
	// The description
	Description pulumi.StringPtrInput
	// The disks
	Disk ContainerDiskPtrInput
	// The initialization configuration
	Initialization ContainerInitializationPtrInput
	// The memory allocation
	Memory ContainerMemoryPtrInput
	// The network interfaces
	NetworkInterfaces ContainerNetworkInterfaceArrayInput
	// The node name
	NodeName pulumi.StringInput
	// The operating system configuration
	OperatingSystem ContainerOperatingSystemPtrInput
	// The ID of the pool to assign the container to
	PoolId pulumi.StringPtrInput
	// Whether to start the container
	Started pulumi.BoolPtrInput
	// Whether to create a template
	Template pulumi.BoolPtrInput
	// The VM identifier
	VmId pulumi.IntPtrInput
}

func (ContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerArgs)(nil)).Elem()
}

type ContainerInput interface {
	pulumi.Input

	ToContainerOutput() ContainerOutput
	ToContainerOutputWithContext(ctx context.Context) ContainerOutput
}

func (*Container) ElementType() reflect.Type {
	return reflect.TypeOf((*Container)(nil))
}

func (i *Container) ToContainerOutput() ContainerOutput {
	return i.ToContainerOutputWithContext(context.Background())
}

func (i *Container) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerOutput)
}

func (i *Container) ToContainerPtrOutput() ContainerPtrOutput {
	return i.ToContainerPtrOutputWithContext(context.Background())
}

func (i *Container) ToContainerPtrOutputWithContext(ctx context.Context) ContainerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPtrOutput)
}

type ContainerPtrInput interface {
	pulumi.Input

	ToContainerPtrOutput() ContainerPtrOutput
	ToContainerPtrOutputWithContext(ctx context.Context) ContainerPtrOutput
}

type containerPtrType ContainerArgs

func (*containerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Container)(nil))
}

func (i *containerPtrType) ToContainerPtrOutput() ContainerPtrOutput {
	return i.ToContainerPtrOutputWithContext(context.Background())
}

func (i *containerPtrType) ToContainerPtrOutputWithContext(ctx context.Context) ContainerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPtrOutput)
}

// ContainerArrayInput is an input type that accepts ContainerArray and ContainerArrayOutput values.
// You can construct a concrete instance of `ContainerArrayInput` via:
//
//          ContainerArray{ ContainerArgs{...} }
type ContainerArrayInput interface {
	pulumi.Input

	ToContainerArrayOutput() ContainerArrayOutput
	ToContainerArrayOutputWithContext(context.Context) ContainerArrayOutput
}

type ContainerArray []ContainerInput

func (ContainerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Container)(nil)).Elem()
}

func (i ContainerArray) ToContainerArrayOutput() ContainerArrayOutput {
	return i.ToContainerArrayOutputWithContext(context.Background())
}

func (i ContainerArray) ToContainerArrayOutputWithContext(ctx context.Context) ContainerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerArrayOutput)
}

// ContainerMapInput is an input type that accepts ContainerMap and ContainerMapOutput values.
// You can construct a concrete instance of `ContainerMapInput` via:
//
//          ContainerMap{ "key": ContainerArgs{...} }
type ContainerMapInput interface {
	pulumi.Input

	ToContainerMapOutput() ContainerMapOutput
	ToContainerMapOutputWithContext(context.Context) ContainerMapOutput
}

type ContainerMap map[string]ContainerInput

func (ContainerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Container)(nil)).Elem()
}

func (i ContainerMap) ToContainerMapOutput() ContainerMapOutput {
	return i.ToContainerMapOutputWithContext(context.Background())
}

func (i ContainerMap) ToContainerMapOutputWithContext(ctx context.Context) ContainerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerMapOutput)
}

type ContainerOutput struct{ *pulumi.OutputState }

func (ContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Container)(nil))
}

func (o ContainerOutput) ToContainerOutput() ContainerOutput {
	return o
}

func (o ContainerOutput) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return o
}

func (o ContainerOutput) ToContainerPtrOutput() ContainerPtrOutput {
	return o.ToContainerPtrOutputWithContext(context.Background())
}

func (o ContainerOutput) ToContainerPtrOutputWithContext(ctx context.Context) ContainerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Container) *Container {
		return &v
	}).(ContainerPtrOutput)
}

type ContainerPtrOutput struct{ *pulumi.OutputState }

func (ContainerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Container)(nil))
}

func (o ContainerPtrOutput) ToContainerPtrOutput() ContainerPtrOutput {
	return o
}

func (o ContainerPtrOutput) ToContainerPtrOutputWithContext(ctx context.Context) ContainerPtrOutput {
	return o
}

func (o ContainerPtrOutput) Elem() ContainerOutput {
	return o.ApplyT(func(v *Container) Container {
		if v != nil {
			return *v
		}
		var ret Container
		return ret
	}).(ContainerOutput)
}

type ContainerArrayOutput struct{ *pulumi.OutputState }

func (ContainerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Container)(nil))
}

func (o ContainerArrayOutput) ToContainerArrayOutput() ContainerArrayOutput {
	return o
}

func (o ContainerArrayOutput) ToContainerArrayOutputWithContext(ctx context.Context) ContainerArrayOutput {
	return o
}

func (o ContainerArrayOutput) Index(i pulumi.IntInput) ContainerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Container {
		return vs[0].([]Container)[vs[1].(int)]
	}).(ContainerOutput)
}

type ContainerMapOutput struct{ *pulumi.OutputState }

func (ContainerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Container)(nil))
}

func (o ContainerMapOutput) ToContainerMapOutput() ContainerMapOutput {
	return o
}

func (o ContainerMapOutput) ToContainerMapOutputWithContext(ctx context.Context) ContainerMapOutput {
	return o
}

func (o ContainerMapOutput) MapIndex(k pulumi.StringInput) ContainerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Container {
		return vs[0].(map[string]Container)[vs[1].(string)]
	}).(ContainerOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerOutput{})
	pulumi.RegisterOutputType(ContainerPtrOutput{})
	pulumi.RegisterOutputType(ContainerArrayOutput{})
	pulumi.RegisterOutputType(ContainerMapOutput{})
}