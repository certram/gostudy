package e

import "errors"

var (
	ErrOutOfRange  = errors.New("index111 out of range")
	ErrSrcSliceNil = errors.New("the resoure slice is nil")
)
