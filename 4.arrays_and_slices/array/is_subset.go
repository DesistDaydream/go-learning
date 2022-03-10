package main

// 判断一个数组是否是另一个数组的子集
func isSubset(subset, superset []string) bool {
	set := make(map[string]int)
	for _, value := range superset {
		set[value] += 1
	}

	for _, value := range subset {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}
