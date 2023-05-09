package sliceutil

// RemoveString remove string from slice if function return true.
func RemoveString(slice []string, remove func(item string) bool) []string {
	for i := 0; i < len(slice); i++ {
		if remove(slice[i]) {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}
	return slice
}

// FindString return true if target in slice, return false if not.
func FindString(slice []string, target string) bool {
	for _, str := range slice {
		if str == target {
			return true
		}
	}
	return false
}

// FindInt return true if target in slice, return false if not.
func FindInt(slice []int, target int) bool {
	for _, str := range slice {
		if str == target {
			return true
		}
	}
	return false
}

//func contains(s []int, target int) bool {
//	sort.Ints(s)
//	i := sort.SearchInts(s, target)
//	return i < len(s) && s[i] == target
//}

// FindUint return true if target in slice, return false if not.
func FindUint(slice []uint, target uint) bool {
	for _, str := range slice {
		if str == target {
			return true
		}
	}
	return false
}

// Find return true if target in slice, return false if not.
func Find[T comparable](slice []T, target T) bool {
	for _, val := range slice {
		if val == target {
			return true
		}
	}
	return false
}
