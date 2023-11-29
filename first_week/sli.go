package main

// SliDelete 切片的删除操作
func sliDelete[T int | string | float64](idx int, vals []T) []T {
	// 边界检查
	if idx < 0 || idx >= len(vals) {
		panic("index out of range")
	}
	copy(vals[idx:], vals[idx+1:])
	// 缩短切片长度
	return vals[:len(vals)-1]

}
