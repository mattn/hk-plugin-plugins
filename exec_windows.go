package main

import (
	"os"
	"strings"
)

func executable(file string) bool {
	file = strings.ToLower(file)
	for _, pathext := range strings.Split(os.Getenv("PATHEXT"), ";") {
		if strings.HasSuffix(file, strings.ToLower(pathext)) {
			return true
		}
	}
	return false
}
