package slice

import "github.com/certram/gostudy/e"

// 泛型实现slice缩容
// 要求一:能够实现删除操作就行
func DeleteElementAtV1(srcSlice []int, index int) ([]int, error) {
	if index >= len(srcSlice) || index < 0 {
		return nil, e.ErrOutOfRange
	}
	res := make([]int, 0, len(srcSlice)-1)
	//前半段
	for i := 0; i < index; i++ {
		res = append(res, srcSlice[i])
	}
	for i := index + 1; i < len(srcSlice); i++ {
		res = append(res, srcSlice[i])
	}
	return res, nil
}

// 要求二：考虑使用比较高性能的实现。
func DeleteElementAtV2(srcSlice []int, index int) ([]int, error) {

	if index >= len(srcSlice) || index < 0 {
		return nil, e.ErrOutOfRange
	}
	res := append(srcSlice[:index], srcSlice[index+1:]...)
	return res, nil
}

// 要求三：改造为泛型方法
func DeleteElementAtV3[T any](srcSlice []T, index int) ([]T, error) {
	if index < 0 || index >= len(srcSlice) {
		return nil, e.ErrOutOfRange
	}

	newSlice := append(srcSlice[:index], srcSlice[index+1:]...)
	return newSlice, nil
}

//要求四：支持缩容，并旦设计缩容机制。
