package util

func ContainString(container []string, check string) bool {
	for _, v := range container {
		if v == check {
			return true
		}
	}
	return false
}
