// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ec2.getSubnetIds` provides a set of ids for a vpcId
//
// This resource can be useful for getting back a set of subnet ids for a vpc.
//
// > **NOTE:** The `ec2.getSubnetIds` data source has been deprecated and will be removed in a future version. Use the `ec2.getSubnets` data source instead.
func GetSubnetIds(ctx *pulumi.Context, args *GetSubnetIdsArgs, opts ...pulumi.InvokeOption) (*GetSubnetIdsResult, error) {
	var rv GetSubnetIdsResult
	err := ctx.Invoke("aws:ec2/getSubnetIds:getSubnetIds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubnetIds.
type GetSubnetIdsArgs struct {
	// Custom filter block as described below.
	Filters []GetSubnetIdsFilter `pulumi:"filters"`
	// Map of tags, each pair of which must exactly match
	// a pair on the desired subnets.
	Tags map[string]string `pulumi:"tags"`
	// VPC ID that you want to filter from.
	VpcId string `pulumi:"vpcId"`
}

// A collection of values returned by getSubnetIds.
type GetSubnetIdsResult struct {
	Filters []GetSubnetIdsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Set of all the subnet ids found. This data source will fail if none are found.
	Ids   []string          `pulumi:"ids"`
	Tags  map[string]string `pulumi:"tags"`
	VpcId string            `pulumi:"vpcId"`
}

func GetSubnetIdsOutput(ctx *pulumi.Context, args GetSubnetIdsOutputArgs, opts ...pulumi.InvokeOption) GetSubnetIdsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSubnetIdsResult, error) {
			args := v.(GetSubnetIdsArgs)
			r, err := GetSubnetIds(ctx, &args, opts...)
			var s GetSubnetIdsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSubnetIdsResultOutput)
}

// A collection of arguments for invoking getSubnetIds.
type GetSubnetIdsOutputArgs struct {
	// Custom filter block as described below.
	Filters GetSubnetIdsFilterArrayInput `pulumi:"filters"`
	// Map of tags, each pair of which must exactly match
	// a pair on the desired subnets.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// VPC ID that you want to filter from.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

func (GetSubnetIdsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetIdsArgs)(nil)).Elem()
}

// A collection of values returned by getSubnetIds.
type GetSubnetIdsResultOutput struct{ *pulumi.OutputState }

func (GetSubnetIdsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetIdsResult)(nil)).Elem()
}

func (o GetSubnetIdsResultOutput) ToGetSubnetIdsResultOutput() GetSubnetIdsResultOutput {
	return o
}

func (o GetSubnetIdsResultOutput) ToGetSubnetIdsResultOutputWithContext(ctx context.Context) GetSubnetIdsResultOutput {
	return o
}

func (o GetSubnetIdsResultOutput) Filters() GetSubnetIdsFilterArrayOutput {
	return o.ApplyT(func(v GetSubnetIdsResult) []GetSubnetIdsFilter { return v.Filters }).(GetSubnetIdsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSubnetIdsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetIdsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Set of all the subnet ids found. This data source will fail if none are found.
func (o GetSubnetIdsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSubnetIdsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetSubnetIdsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetSubnetIdsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetSubnetIdsResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetIdsResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSubnetIdsResultOutput{})
}