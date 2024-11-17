// Copyright (c) Christopher Marget, 2024-2024.
// SPDX-License-Identifier: MIT

package emptytype_test

import (
	"context"
	"testing"

	"github.com/chrismarget/terraform-framework-types/emptytype"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

func TestEmptyValidateAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		addressValue  emptytype.Empty
		expectedDiags diag.Diagnostics
	}{
		"empty-struct": {
			addressValue: emptytype.Empty{},
		},
		"null": {
			addressValue: emptytype.NewEmptyNull(),
		},
		"unknown": {
			addressValue: emptytype.NewEmptyUnknown(),
		},
		"true": {
			addressValue: emptytype.NewEmptyValue(true),
		},
		"false": {
			addressValue: emptytype.NewEmptyValue(false),
			expectedDiags: diag.Diagnostics{
				diag.NewAttributeErrorDiagnostic(
					path.Root("test"),
					"Attribute must not be `false`",
					"Valid values are `true` and `null`, got `false`",
				),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			resp := xattr.ValidateAttributeResponse{}

			testCase.addressValue.ValidateAttribute(
				context.Background(),
				xattr.ValidateAttributeRequest{Path: path.Root("test")},
				&resp,
			)

			if diff := cmp.Diff(resp.Diagnostics, testCase.expectedDiags); diff != "" {
				t.Errorf("Unexpected diagnostics (-got, +expected): %s", diff)
			}
		})
	}
}

func TestEmptyValue(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		emptyValue    emptytype.Empty
		expectedBool  bool
		expectedDiags diag.Diagnostics
	}{
		"null ": {
			emptyValue:    emptytype.NewEmptyNull(),
			expectedDiags: diag.Diagnostics{},
		},
		"unknown ": {
			emptyValue:    emptytype.NewEmptyUnknown(),
			expectedDiags: diag.Diagnostics{},
		},
		"true": {
			emptyValue:   emptytype.NewEmptyValue(true),
			expectedBool: true,
		},
	}
	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			boolVal := testCase.emptyValue.ValueBool()

			if boolVal != testCase.expectedBool {
				t.Errorf("Unexpected difference in bool value, got: %t, expected: %t", boolVal, testCase.expectedBool)
			}
		})
	}
}
