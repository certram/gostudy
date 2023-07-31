package slice

// golang slice的缩容
func ShrinkSlice[T any](srcSlice []T) []T {
	if len(srcSlice) > 0 && len(srcSlice) < cap(srcSlice)/4 {
		newCap := len(srcSlice)
		newSlice := make([]T, newCap)
		copy(newSlice, srcSlice[:newCap])
		return newSlice
	}
	return srcSlice
}
