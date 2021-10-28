package env

import (
	"os"
	"strconv"
)

func ResolveString(name, def string) (res string) {
	if v, ok := os.LookupEnv(name); ok {
		return v
	}

	return def
}

func ResolveInt(name string, def int) (res int) {
	if v, ok := os.LookupEnv(name); ok {
		if u, err := strconv.ParseInt(v, 10, 64); err == nil {
			return int(u)
		}
	}

	return def
}
