// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the storage task assignment properties
// Azure REST API version: 2023-05-01.
func LookupStorageTaskAssignment(ctx *pulumi.Context, args *LookupStorageTaskAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupStorageTaskAssignmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStorageTaskAssignmentResult
	err := ctx.Invoke("azure-native:storage:getStorageTaskAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageTaskAssignmentArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the storage task assignment within the specified resource group. Storage task assignment names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	StorageTaskAssignmentName string `pulumi:"storageTaskAssignmentName"`
}

// The storage task assignment.
type LookupStorageTaskAssignmentResult struct {
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Properties of the storage task assignment.
	Properties StorageTaskAssignmentPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupStorageTaskAssignmentOutput(ctx *pulumi.Context, args LookupStorageTaskAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupStorageTaskAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageTaskAssignmentResult, error) {
			args := v.(LookupStorageTaskAssignmentArgs)
			r, err := LookupStorageTaskAssignment(ctx, &args, opts...)
			var s LookupStorageTaskAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageTaskAssignmentResultOutput)
}

type LookupStorageTaskAssignmentOutputArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the storage task assignment within the specified resource group. Storage task assignment names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	StorageTaskAssignmentName pulumi.StringInput `pulumi:"storageTaskAssignmentName"`
}

func (LookupStorageTaskAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageTaskAssignmentArgs)(nil)).Elem()
}

// The storage task assignment.
type LookupStorageTaskAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupStorageTaskAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageTaskAssignmentResult)(nil)).Elem()
}

func (o LookupStorageTaskAssignmentResultOutput) ToLookupStorageTaskAssignmentResultOutput() LookupStorageTaskAssignmentResultOutput {
	return o
}

func (o LookupStorageTaskAssignmentResultOutput) ToLookupStorageTaskAssignmentResultOutputWithContext(ctx context.Context) LookupStorageTaskAssignmentResultOutput {
	return o
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupStorageTaskAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageTaskAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupStorageTaskAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageTaskAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the storage task assignment.
func (o LookupStorageTaskAssignmentResultOutput) Properties() StorageTaskAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupStorageTaskAssignmentResult) StorageTaskAssignmentPropertiesResponse { return v.Properties }).(StorageTaskAssignmentPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupStorageTaskAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStorageTaskAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupStorageTaskAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageTaskAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageTaskAssignmentResultOutput{})
}