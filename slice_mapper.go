package slice_util

import "context"

func Map[T any, R any](slice []T, mapper func(T) R) []R {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		ret = append(ret, mapper(t))
	}
	return ret
}

func MapWithCtx[T any, R any](ctx context.Context, slice []T, mapper func(context.Context, T) R) []R {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		ret = append(ret, mapper(ctx, t))
	}
	return ret
}

func MapWithErr[T any, R any](slice []T, mapper func(T) (R, error)) []R {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		if r, err := mapper(t); err != nil {
			ret = append(ret, r)
		}
	}
	return ret
}

func MapWithCtxErr[T any, R any](ctx context.Context, slice []T, mapper func(context.Context, T) (R, error)) []R {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		if r, err := mapper(ctx, t); err != nil {
			ret = append(ret, r)
		}
	}
	return ret
}

func FlatMap[T any, R any](slice []T, mapper func(T) []R) []R {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		ret = append(ret, mapper(t)...)
	}
	return ret
}

func FlatMapWithCtx[T any, R any](ctx context.Context, slice []T, mapper func(context.Context, T) []R) []R {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		ret = append(ret, mapper(ctx, t)...)
	}
	return ret
}

func FlatMapWithErr[T any, R any](slice []T, mapper func(T) ([]R, error)) []R {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		ret = append(ret, mapper(t)...)
	}
	return ret
}

func FlatMapWithCtx[T any, R any](ctx context.Context, slice []T, mapper func(context.Context, T) []R) []R {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		ret = append(ret, mapper(ctx, t)...)
	}
	return ret
}
