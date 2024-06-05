// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/names"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
)

// @FrameworkResource("aws_ec2_capacity_block_reservation")
// @Tags(identifierAttribute="id")
// @Testing(tagsTest=false)
func newResourceCapacityBlockReservation(context.Context) (resource.ResourceWithConfigure, error) {
	r := &resourceCapacityBlockReservation{}

	return r, nil
}

type resourceCapacityBlockReservation struct {
	framework.ResourceWithConfigure
	framework.WithNoOpUpdate[resourceCapacityBlockReservationData]
	framework.WithNoOpDelete
}

// Metadata should return the full name of the resource, such as
// examplecloud_thing.
func (r *resourceCapacityBlockReservation) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = "aws_ec2_capacity_block_reservation"
}

// Schema returns the schema for this resource.
func (r *resourceCapacityBlockReservation) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	s := schema.Schema{
		Attributes: map[string]schema.Attribute{
			names.AttrARN: schema.StringAttribute{
				Computed: true,
			},
			names.AttrAvailabilityZone: schema.StringAttribute{
				Computed: true,
			},
			"capacity_block_offering_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ebs_optimized": schema.BoolAttribute{
				Computed: true,
			},
			"end_date": schema.StringAttribute{
				Computed: true,
			},
			"end_date_type": schema.StringAttribute{
				Computed: true,
			},
			"ephemeral_storage": schema.BoolAttribute{
				Computed: true,
			},
			names.AttrID: // TODO framework.IDAttribute()
			schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			names.AttrInstanceCount: schema.Int64Attribute{
				Computed: true,
			},
			"instance_match_criteria": schema.StringAttribute{
				Computed: true,
			},
			"instance_platform": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				// TODO Validate,
			},
			names.AttrInstanceType: schema.StringAttribute{
				Computed: true,
			},
			"outpost_arn": schema.StringAttribute{
				Computed: true,
			},
			names.AttrOwnerID: schema.StringAttribute{
				Computed: true,
			},
			"placement_group_arn": schema.StringAttribute{
				Computed: true,
			},
			"start_date": schema.StringAttribute{
				Computed: true,
			},
			names.AttrTags:    tftags.TagsAttribute(),
			names.AttrTagsAll: tftags.TagsAttributeComputedOnly(),
			"tenancy": schema.StringAttribute{
				Computed: true,
			},
		},
	}

	response.Schema = s
}

const (
	ResNameCapacityBlockReservation = "Capacity Block Reservation"
)

// Create is called when the provider must create a new resource.
// Config and planned state values should be read from the CreateRequest and new state values set on the CreateResponse.
func (r *resourceCapacityBlockReservation) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var data resourceCapacityBlockReservationData

	response.Diagnostics.Append(request.Plan.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	data.ID = types.StringValue("TODO")

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

// Read is called when the provider must read resource values in order to update state.
// Planned state values should be read from the ReadRequest and new state values set on the ReadResponse.
func (r *resourceCapacityBlockReservation) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var data resourceCapacityBlockReservationData

	response.Diagnostics.Append(request.State.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

// Update is called to update the state of the resource.
// Config, planned state, and prior state values should be read from the UpdateRequest and new state values set on the UpdateResponse.
func (r *resourceCapacityBlockReservation) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var old, new resourceCapacityBlockReservationData

	response.Diagnostics.Append(request.State.Get(ctx, &old)...)

	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(request.Plan.Get(ctx, &new)...)

	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &new)...)
}

// Delete is called when the provider must delete the resource.
// Config values may be read from the DeleteRequest.
//
// If execution completes without error, the framework will automatically call DeleteResponse.State.RemoveResource(),
// so it can be omitted from provider logic.
func (r *resourceCapacityBlockReservation) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var data resourceCapacityBlockReservationData

	response.Diagnostics.Append(request.State.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "deleting TODO", map[string]interface{}{
		names.AttrID: data.ID.ValueString(),
	})
}

// ImportState is called when the provider must import the state of a resource instance.
// This method must return enough state so the Read method can properly refresh the full resource.
//
// If setting an attribute with the import identifier, it is recommended to use the ImportStatePassthroughID() call in this method.
func (r *resourceCapacityBlockReservation) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root(names.AttrID), request, response)
}

// ModifyPlan is called when the provider has an opportunity to modify
// the plan: once during the plan phase when Terraform is determining
// the diff that should be shown to the user for approval, and once
// during the apply phase with any unknown values from configuration
// filled in with their final values.
//
// The planned new state is represented by
// ModifyPlanResponse.Plan. It must meet the following
// constraints:
// 1. Any non-Computed attribute set in config must preserve the exact
// config value or return the corresponding attribute value from the
// prior state (ModifyPlanRequest.State).
// 2. Any attribute with a known value must not have its value changed
// in subsequent calls to ModifyPlan or Create/Read/Update.
// 3. Any attribute with an unknown value may either remain unknown
// or take on any value of the expected type.
//
// Any errors will prevent further resource-level plan modifications.
func (r *resourceCapacityBlockReservation) ModifyPlan(ctx context.Context, request resource.ModifyPlanRequest, response *resource.ModifyPlanResponse) {
	r.SetTagsAll(ctx, request, response)
}

type resourceCapacityBlockReservationData struct {
	ARN                     types.String `tfsdk:"arn"`
	AvailabilityZone        types.String `tfsdk:"availability_zone"`
	CapacityBlockOfferingID types.String `tfsdk:"capacity_block_offering_id"`
	EbsOptimized            types.Bool   `tfsdk:"ebs_optimized"`
	EndDate                 types.String `tfsdk:"end_date"`
	EndDateType             types.String `tfsdk:"end_date_type"`
	EphemeralStorage        types.Bool   `tfsdk:"ephemeral_storage"`
	ID                      types.String `tfsdk:"id"`
	InstanceCount           types.Int64  `tfsdk:"instance_count"`
	InstanceMatchCriteria   types.String `tfsdk:"instance_match_criteria"`
	InstancePlatform        types.String `tfsdk:"instance_platform"`
	InstanceType            types.String `tfsdk:"instance_type"`
	OutpostARN              types.String `tfsdk:"outpost_arn"`
	OwnerID                 types.String `tfsdk:"owner_id"`
	PlacementGroupARN       types.String `tfsdk:"placement_group_arn"`
	StartDate               types.String `tfsdk:"start_date"`
	Tags                    types.Map    `tfsdk:"tags"`
	TagsAll                 types.Map    `tfsdk:"tags_all"`
	Tenancy                 types.String `tfsdk:"tenancy"`
}
