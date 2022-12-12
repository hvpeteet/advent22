package fuck_go

func NonMutateAppend[T any](slice []T, x T) []T {
	ret := make([]T, len(slice)+1)
	copy(ret, slice)
	ret[len(slice)] = x
	return ret
}
