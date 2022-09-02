package utils

import "os"

func ArrayContains(strings []string, key string) bool {
	for _, s := range strings {
		if s == key {
			return true
		}
	}

	return false
}

func GetEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return value
}
