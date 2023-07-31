package slice

import "github.com/certram/gostudy/e"

func AddElementAt[T any](srcSlice []T, index int, val T) ([]T, error) {
	if srcSlice == nil {
		return nil, e.ErrSrcSliceNil
	}
	if index > len(srcSlice) || index < 0 {
		return srcSlice, e.ErrOutOfRange
	}
	newSlice := make([]T, 0, len(srcSlice)+1)
	for i := 0; i < index; i++ {
		newSlice = append(newSlice, srcSlice[i])
	}
	newSlice = append(newSlice, val)
	for i := index; i < len(srcSlice); i++ {
		newSlice = append(newSlice, srcSlice[i])
	}
	return newSlice, nil
}
