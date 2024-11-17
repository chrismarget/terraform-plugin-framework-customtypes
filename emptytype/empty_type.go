// Copyright (c) Christopher Marget, 2024-2024.
// SPDX-License-Identifier: MIT

package emptytype

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ basetypes.BoolTypable = (*EmptyType)(nil)

type EmptyType struct {
	basetypes.BoolType
}

// String returns a human readable string of the type name.
func (e EmptyType) String() string {
	return "empty.EmptyType"
}

// ValueType returns the Value type.
func (e EmptyType) ValueType(_ context.Context) attr.Value {
	return Empty{}
}

// Equal returns true if the given type is equivalent.
func (e EmptyType) Equal(o attr.Type) bool {
	other, ok := o.(EmptyType)

	if !ok {
		return false
	}

	return e.BoolType.Equal(other.BoolType)
}

// ValueFromString returns a BoolValuable type given a BoolValue.
func (e EmptyType) ValueFromString(_ context.Context, in basetypes.BoolValue) (basetypes.BoolValuable, diag.Diagnostics) {
	return Empty{
		BoolValue: in,
	}, nil
}

// ValueFromTerraform returns a Value given a tftypes.Value.  This is meant to convert the tftypes.Value into a more convenient Go type
// for the provider to consume the data with.
func (e EmptyType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := e.BoolType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	boolValue, ok := attrValue.(basetypes.BoolValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	boolValuable, diags := e.ValueFromString(ctx, boolValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting BoolValue to BoolValuable: %v", diags)
	}

	return boolValuable, nil
}
