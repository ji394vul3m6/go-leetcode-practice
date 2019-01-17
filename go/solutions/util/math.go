package util

func Min(container []int) int {
	if len(container) == 0 {
		return -1
	}

	ret := container[0]
	for _, v := range container {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func Max(container []int) int {
	if len(container) == 0 {
		return -1
	}

	ret := container[0]
	for _, v := range container {
		if v > ret {
			ret = v
		}
	}
	return ret
}
