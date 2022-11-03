package utils

import (
	"os"
	"strconv"
	"time"
)

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

func Utint(prev string, next int) int {
	result := 0
	value, err := strconv.Atoi(prev)

	if err != nil || value < 0 {
		result = next
	} else {
		result = value
	}

	if result < 0 {
		result = 0
	}

	return result
}

func ConvertStringToUtc(date string) time.Time {
	result, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return ConvertStringToUtc("0001-01-01 00:00:00 +0000 UTC")
	}

	return result
}

func ConvertStringToInt(number string) int {
	result := 0

	value, err := strconv.Atoi(number)
	if err == nil {
		result = value
	}

	return result
}
