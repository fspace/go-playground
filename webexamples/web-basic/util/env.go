package util

import "os"

func GetenvWithDefault(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		val = defaultValue
	}

	return val
}
