package utils

func ArrayContains(strings []string, key string) bool {
	for _, s := range strings {
		if s == key {
			return true
		}
	}

	return false
}
