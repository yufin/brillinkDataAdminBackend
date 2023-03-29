package common

func DeleteSliceElms[T int64](sl []T, elms ...T) []T {
	// 先将元素转为 set。
	m := make(map[T]struct{})
	for _, v := range elms {
		m[v] = struct{}{}
	}
	// 过滤掉指定元素。
	res := make([]T, 0, len(sl))
	for _, v := range sl {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
	}
	return res
}
