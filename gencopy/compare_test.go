package gencopy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func stringPtr(s string) *string {
	return &s
}

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
			d := Comparator{}
			equal := d.stringEqualStrict(tC.fieldPath, tC.v1, tC.v2)
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
			d := Comparator{}
			equal := d.stringEqualSoft(tC.fieldPath, tC.v1, tC.v2)
			assert.Equal(t, tC.expect, equal)
			if !equal {
				assert.Equal(t, []string{tC.fieldPath}, d.diffs)
			}
		})
	}
}
