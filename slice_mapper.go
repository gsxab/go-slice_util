/*
 * SPDX-License-Identifier: Apache-2.0
 *
 * Copyright (c) 2023 Gsxab
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package slice_util

import (
	"context"
)

func ignoreError[T any, R any](mapper func(T) (R, error)) func(T) R {
	return func(t T) R {
		r, _ := mapper(t)
		return r
	}
}

func ignoreErrorC[T any, R any](mapper func(context.Context, T) (R, error)) func(context.Context, T) R {
	return func(ctx context.Context, t T) R {
		r, _ := mapper(ctx, t)
		return r
	}
}

func Map[T any, R any](slice []T, mapper func(T) R) []R {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		ret = append(ret, mapper(t))
	}
	return ret
}

func MapC[T any, R any](ctx context.Context, slice []T, mapper func(context.Context, T) R) []R {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		ret = append(ret, mapper(ctx, t))
	}
	return ret
}

func MapE[T any, R any](slice []T, mapper func(T) (R, error)) ([]R, error) {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		r, err := mapper(t)
		if err != nil {
			return nil, err
		}
		ret = append(ret, r)
	}
	return ret, nil
}

func MapIgnoreE[T any, R any](slice []T, mapper func(T) (R, error)) []R {
	return Map(slice, ignoreError(mapper))
}

func MapSkipE[T any, R any](slice []T, mapper func(T) (R, error)) []R {
	return Map(slice, ignoreError(mapper))
}

func MapCE[T any, R any](ctx context.Context, slice []T, mapper func(context.Context, T) (R, error)) ([]R, error) {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		r, err := mapper(ctx, t)
		if err != nil {
			return nil, err
		}
		ret = append(ret, r)
	}
	return ret, nil
}

func MapIgnoreCE[T any, R any](ctx context.Context, slice []T, mapper func(context.Context, T) (R, error)) []R {
	return MapC(ctx, slice, ignoreErrorC(mapper))
}

func MapSkipCE[T any, R any](ctx context.Context, slice []T, mapper func(context.Context, T) (R, error)) []R {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		r, err := mapper(ctx, t)
		if err != nil {
			continue
		}
		ret = append(ret, r)
	}
	return ret
}

func FlatMapSlice[T any, R any](slice []T, mapper func(T) []R) []R {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		ret = append(ret, mapper(t)...)
	}
	return ret
}

func FlatMapSliceC[T any, R any](ctx context.Context, slice []T, mapper func(context.Context, T) []R) []R {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		ret = append(ret, mapper(ctx, t)...)
	}
	return ret
}

func FlatMapSliceE[T any, R any](slice []T, mapper func(T) ([]R, error)) ([]R, error) {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		s, err := mapper(t)
		if err != nil {
			return nil, err
		}
		ret = append(ret, s...)
	}
	return ret, nil
}

func FlatMapSliceIgnoreE[T any, R any](slice []T, mapper func(T) ([]R, error)) []R {
	return FlatMapSlice(slice, ignoreError(mapper))
}

func FlatMapSliceSkipE[T any, R any](slice []T, mapper func(T) ([]R, error)) []R {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		s, err := mapper(t)
		if err != nil {
			continue
		}
		ret = append(ret, s...)
	}
	return ret
}

func FlatMapSliceCE[T any, R any](ctx context.Context, slice []T, mapper func(context.Context, T) ([]R, error)) ([]R, error) {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		s, err := mapper(ctx, t)
		if err != nil {
			return nil, err
		}
		ret = append(ret, s...)
	}
	return ret, nil
}

func FlatMapSliceIgnoreCE[T any, R any](ctx context.Context, slice []T, mapper func(context.Context, T) ([]R, error)) []R {
	return FlatMapSliceC(ctx, slice, ignoreErrorC(mapper))
}

func FlatMapSliceSkipCE[T any, R any](ctx context.Context, slice []T, mapper func(context.Context, T) ([]R, error)) []R {
	ret := make([]R, 0, len(slice))
	for _, t := range slice {
		s, err := mapper(ctx, t)
		if err != nil {
			continue
		}
		ret = append(ret, s...)
	}
	return ret
}

// func FlatMap[T any, R any](slice []T, mapper func(T) container.Iterable[R]) []R {
// 	ret := make([]R, 0, len(slice))
// 	for _, t := range slice {
// 		ret = append(ret, mapper(t)...)
// 	}
// 	return ret
// }
