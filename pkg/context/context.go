package context

import "context"

func Get[K comparable, V any](ctx context.Context, key K) (V, bool) {
	v, ok := ctx.Value(key).(V)
	return v, ok
}

func Set[K comparable, V any](ctx context.Context, key K, val V) context.Context {
	return context.WithValue(ctx, key, val)
}
