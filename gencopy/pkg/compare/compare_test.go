package compare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func stringPtr(v string) *string { return &v }

//TODO(a-hilaly) add unit tests for bool, int, float
/*
func boolPtr(v bool) *bool          { return &v }
func int32Ptr(v int32) *int32       { return &v }
func int64Ptr(v int64) *int64       { return &v }
func float32Ptr(v float32) *float32 { return &v }
func float64Ptr(v float64) *float64 { return &v }
*/

func TestEqualStringStrict(t *testing.T) {
	testCases := []struct {
		desc      string
		fieldPath string
		v1        *string
		v2        *string
		expect    bool
	}{
		{
			desc:      "nil ptr strings",
			fieldPath: "Spec.Name",
			v1:        nil,
			v2:        nil,
			expect:    true,
		},
		{
			desc:      "empty string value",
			fieldPath: "Spec.Name",
			v1:        stringPtr(""),
			v2:        stringPtr(""),
			expect:    true,
		},
		{
			desc:      "equal string values",
			fieldPath: "Spec.Name",
			v1:        stringPtr("value"),
			v2:        stringPtr("value"),
			expect:    true,
		},
		{
			desc:      "nil and empty string",
			fieldPath: "Spec.Name",
			v1:        stringPtr(""),
			v2:        nil,
			expect:    false,
		},
		{
			desc:      "reverse nil and empty string",
			fieldPath: "Spec.Name",
			v1:        nil,
			v2:        stringPtr(""),
			expect:    false,
		},
		{
			desc:      "non equal string values",
			fieldPath: "Spec.Name",
			v1:        stringPtr("value 1"),
			v2:        stringPtr("value 2"),
			expect:    false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			d := &comparator{}
			equal := d.StringEqualStrict(tC.fieldPath, tC.v1, tC.v2)
			assert.Equal(t, tC.expect, equal)
			if !equal {
				assert.Equal(t, []string{tC.fieldPath}, d.diffs)
			}
		})
	}
}

func TestEqualStringSoft(t *testing.T) {
	testCases := []struct {
		desc      string
		fieldPath string
		v1        *string
		v2        *string
		expect    bool
	}{
		{
			desc:      "nil ptr strings",
			fieldPath: "Spec.Name",
			v1:        nil,
			v2:        nil,
			expect:    true,
		},
		{
			desc:      "empty string value",
			fieldPath: "Spec.Name",
			v1:        stringPtr(""),
			v2:        stringPtr(""),
			expect:    true,
		},
		{
			desc:      "equal string values",
			fieldPath: "Spec.Name",
			v1:        stringPtr("value"),
			v2:        stringPtr("value"),
			expect:    true,
		},
		{
			desc:      "nil and empty string",
			fieldPath: "Spec.Name",
			v1:        stringPtr(""),
			v2:        nil,
			expect:    true,
		},
		{
			desc:      "reverse nil and empty string",
			fieldPath: "Spec.Name",
			v1:        nil,
			v2:        stringPtr(""),
			expect:    true,
		},
		{
			desc:      "non equal string values",
			fieldPath: "Spec.Name",
			v1:        stringPtr("value 1"),
			v2:        stringPtr("value 2"),
			expect:    false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			d := &comparator{}
			equal := d.StringEqualSoft(tC.fieldPath, tC.v1, tC.v2)
			assert.Equal(t, tC.expect, equal)
			if !equal {
				assert.Equal(t, []string{tC.fieldPath}, d.diffs)
			}
		})
	}
}
