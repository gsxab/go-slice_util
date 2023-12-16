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
