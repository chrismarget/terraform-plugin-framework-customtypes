// Copyright (c) Christopher Marget, 2024-2024.
// SPDX-License-Identifier: MIT

package emptytype

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ basetypes.BoolValuable                   = (*Empty)(nil)
	_ basetypes.BoolValuableWithSemanticEquals = (*Empty)(nil)
	_ xattr.ValidateableAttribute              = (*Empty)(nil)
)

type Empty struct {
	basetypes.BoolValue
}

// Type returns an EmptyType.
func (e Empty) Type(_ context.Context) attr.Type {
	return EmptyType{}
}

// Equal returns true if the given value is equivalent.
func (e Empty) Equal(o attr.Value) bool {
	other, ok := o.(Empty)
	if !ok {
		return false
	}

	return e.BoolValue.Equal(other.BoolValue)
}

// BoolSemanticEquals treats false and null as equal
func (e Empty) BoolSemanticEquals(_ context.Context, newValuable basetypes.BoolValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(Empty)
	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected value type was received while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Expected Value Type: "+fmt.Sprintf("%T", e)+"\n"+
				"Got Value Type: "+fmt.Sprintf("%T", newValuable),
		)

		return false, diags
	}

	return e.ValueBool() == newValue.ValueBool(), nil
}

// ValidateAttribute implements attribute value validation. This type only accepts `true`
func (e Empty) ValidateAttribute(_ context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	if e.IsUnknown() || e.IsNull() {
		return
	}

	if !e.ValueBool() {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Attribute must not be `false`",
			fmt.Sprintf("Valid values are `true` and `null`, got `%t`", e.ValueBool()),
		)

		return
	}
}

// NewEmptyNull creates an Empty with a null value. Determine whether the value is null via IsNull method.
func NewEmptyNull() Empty {
	return Empty{
		BoolValue: basetypes.NewBoolNull(),
	}
}

// NewEmptyUnknown creates an Empty with an unknown value. Determine whether the value is unknown via IsUnknown method.
func NewEmptyUnknown() Empty {
	return Empty{
		BoolValue: basetypes.NewBoolUnknown(),
	}
}

// NewEmptyValue creates an Empty with a known value. Access the value via ValueBool method.
func NewEmptyValue(value bool) Empty {
	return Empty{
		BoolValue: basetypes.NewBoolValue(value),
	}
}

// NewEmptyPointerValue creates an Empty with a null value if nil or a known value. Access the value via ValueBoolPointer method.
func NewEmptyPointerValue(value *bool) Empty {
	return Empty{
		BoolValue: basetypes.NewBoolPointerValue(value),
	}
}
