// +build !windows

package main

import (
	"os"
	"strings"
)

func executable(file string) bool {
	fi, err := os.Stat(file)
	if err == nil && strings.Contains(fi.Mode(), "x") {
		return true
	}
	return false
}
