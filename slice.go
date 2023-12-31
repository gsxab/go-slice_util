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

func Prepend[T any](slice []T, element T) []T {
	slice = append(slice, element)
	copy(slice[1:], slice)
	slice[0] = element
	return slice
}

func Insert[T any](slice []T, index int, element T) []T {
	slice = append(slice, element)
	copy(slice[index+1:], slice[index:])
	slice[index] = element
	return slice
}
