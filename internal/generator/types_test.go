package generator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_newStructAccessor(t *testing.T) {
	type input struct {
		parentFields []string
		field        string
	}
	testSets := map[string]struct {
		input    input
		expected string
	}{
		"empty parent fields": {
			input: input{
				parentFields: nil,
				field:        "req",
			},
			expected: "req",
		},
		"single parent field": {
			input: input{
				parentFields: []string{"req"},
				field:        "data",
			},
			expected: "req.data",
		},
		"multiple parent field": {
			input: input{
				parentFields: []string{"req", "data"},
				field:        "tag",
			},
			expected: "req.data.tag",
		},
	}
	for name, ts := range testSets {
		t.Run(name, func(t *testing.T) {
			actual := newStructAccessor(ts.input.parentFields, ts.input.field)
			require.Equal(t, ts.expected, actual)
		})
	}
}

func Test_unexport(t *testing.T) {
	testSets := map[string]struct {
		input    string
		expected string
	}{
		"exported": {
			input:    "TestClient",
			expected: "testClient",
		},
		"unexported": {
			input:    "testClient",
			expected: "testClient",
		},
	}
	for name, ts := range testSets {
		t.Run(name, func(t *testing.T) {
			actual := unexport(ts.input)
			require.Equal(t, ts.expected, actual)
		})
	}
}
