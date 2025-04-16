package utils

func DeleteSlice(a []int, elem int) []int {
	j := 0
	for _, v := range a {
		if v != elem {
			a[j] = v
			j++
		}
	}
	return a[:j]
}

// 数组去重
func RemoveDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// SliceEqual 两个数组是否相同
func SliceEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func RemoveElement(slice []int, index int) []int {
	// 检查索引是否越界
	if index < 0 || index >= len(slice) {
		return nil
	}

	newSlice := make([]int, len(slice)-1)

	copy(newSlice, slice[:index])

	copy(newSlice[index:], slice[index+1:])

	return newSlice
}
