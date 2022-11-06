package hw09structvalidator

func containsString(slice []string, str string) bool {
	for _, v := range slice {
		if str == v {
			return true
		}
	}
	return false
}

func containsInt(slice []int, i int) bool {
	for _, v := range slice {
		if i == v {
			return true
		}
	}
	return false
}
