package slice_util

import "context"

func LinearSearch[T any](slice []T, pred func(T) bool) (bool, int) {
	for idx, elem := range slice {
		if pred(elem) {
			return true, idx
		}
	}
	return false, 0
}

func LinearSearchWithCtx[T any](ctx context.Context, slice []T, pred func(context.Context, T) bool) (bool, int) {
	for idx, elem := range slice {
		if pred(ctx, elem) {
			return true, idx
		}
	}
	return false, 0
}

func LinearSearchWithErr[T any](slice []T, pred func(T) (bool, error)) (bool, int) {
	for idx, elem := range slice {
		if b, err := pred(elem); err != nil && b {
			return true, idx
		}
	}
	return false, 0
}

func LinearSearchWithCtxErr[T any](ctx context.Context, slice []T, pred func(context.Context, T) (bool, error)) (bool, int) {
	for idx, elem := range slice {
		if b, err := pred(ctx, elem); err != nil && b {
			return true, idx
		}
	}
	return false, 0
}

func LinearSearchPtr[T any](slice []T, pred func(*T) bool) (bool, int) {
	for idx, elem := range slice {
		if pred(&elem) {
			return true, idx
		}
	}
	return false, 0
}

func LinearSearchPtrWithCtx[T any](ctx context.Context, slice []T, pred func(context.Context, *T) bool) (bool, int) {
	for idx, elem := range slice {
		if pred(ctx, &elem) {
			return true, idx
		}
	}
	return false, 0
}

func LinearSearchPtrWithErr[T any](slice []T, pred func(*T) (bool, error)) (bool, int) {
	for idx, elem := range slice {
		if b, err := pred(&elem); err != nil && b {
			return true, idx
		}
	}
	return false, 0
}

func LinearSearchPtrWithCtxErr[T any](ctx context.Context, slice []T, pred func(context.Context, *T) (bool, error)) (bool, int) {
	for idx, elem := range slice {
		if b, err := pred(ctx, &elem); err != nil && b {
			return true, idx
		}
	}
	return false, 0
}

func LinearSearchValue[T comparable](slice []T, element T) (bool, int) {
	return LinearSearch(slice, func(t T) bool {
		return t == element
	})
}
