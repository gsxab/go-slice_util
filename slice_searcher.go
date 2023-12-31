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

import "context"

func LinearSearch[T any](slice []T, pred func(T) bool) (bool, int) {
	for idx, elem := range slice {
		if pred(elem) {
			return true, idx
		}
	}
	return false, 0
}

func LinearSearchC[T any](ctx context.Context, slice []T, pred func(context.Context, T) bool) (bool, int) {
	for idx, elem := range slice {
		if pred(ctx, elem) {
			return true, idx
		}
	}
	return false, 0
}

func LinearSearchE[T any](slice []T, pred func(T) (bool, error)) (bool, int) {
	for idx, elem := range slice {
		if b, err := pred(elem); err != nil && b {
			return true, idx
		}
	}
	return false, 0
}

func LinearSearchCE[T any](ctx context.Context, slice []T, pred func(context.Context, T) (bool, error)) (bool, int) {
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

func LinearSearchPtrC[T any](ctx context.Context, slice []T, pred func(context.Context, *T) bool) (bool, int) {
	for idx, elem := range slice {
		if pred(ctx, &elem) {
			return true, idx
		}
	}
	return false, 0
}

func LinearSearchPtrE[T any](slice []T, pred func(*T) (bool, error)) (bool, int) {
	for idx, elem := range slice {
		if b, err := pred(&elem); err != nil && b {
			return true, idx
		}
	}
	return false, 0
}

func LinearSearchPtrCE[T any](ctx context.Context, slice []T, pred func(context.Context, *T) (bool, error)) (bool, int) {
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
