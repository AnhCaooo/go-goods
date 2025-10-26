package helpers

// Remove duplicated items
func RemoveDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		var zero T
		if item == zero {
			continue
		}

		if !allKeys[item] {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
