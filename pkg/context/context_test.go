package context

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type dummyKey struct {
	/* explicitly empty */
}

func TestGet(t *testing.T) {
	type testSet[K, V any] struct {
		newContextFunc func() context.Context
		key            K
		shouldExists   bool
		value          V
	}
	testSets := map[string]testSet[dummyKey, string]{
		"get existing value": {
			newContextFunc: func() context.Context {
				return context.WithValue(context.TODO(), dummyKey{}, "some value")
			},
			key:          dummyKey{},
			value:        "some value",
			shouldExists: true,
		},
		"get non-existing value": {
			newContextFunc: context.TODO,
			key:            dummyKey{},
			shouldExists:   false,
		},
	}
	for name, ts := range testSets {
		t.Run(name, func(t *testing.T) {
			ctx := ts.newContextFunc()
			actual, ok := Get[dummyKey, string](ctx, ts.key)
			if !ts.shouldExists {
				require.False(t, ok)
				return
			}
			require.True(t, ok)
			require.Equal(t, ts.value, actual)
		})
	}
}

func TestSet(t *testing.T) {
	testSets := map[string]struct {
		key   interface{}
		value interface{}
	}{
		"set boolean value": {
			key:   dummyKey{},
			value: true,
		},
		"set integer value": {
			key:   dummyKey{},
			value: 1234,
		},
		"set string value": {
			key:   dummyKey{},
			value: "some value",
		},
	}
	for name, ts := range testSets {
		t.Run(name, func(t *testing.T) {
			ctx := Set(context.TODO(), ts.key, ts.value)
			require.Equal(t, ctx.Value(ts.key), ts.value)
		})
	}
}
