package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	pluginMode := os.Getenv("HKPLUGINMODE")
	description := "Plugin that provides a list of plugins"
	switch pluginMode {
	case "info":
		fmt.Printf("plugins 0.1: List plugins\n\n%s", description)
		return
	}

	hkpath := os.Getenv("HKPATH")
	if hkpath == "" {
		fmt.Fprintf(os.Stderr, "$HKPATH is not set")
		return
	}
	files, err := filepath.Glob(filepath.Join(hkpath, "*"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "$HKPATH is not set")
		return
	}
	os.Setenv("HKPLUGINMODE", "info")
	for _, file := range files {
		fi, err := os.Stat(file)
		if err != nil || fi.IsDir() {
			continue
		}
		if executable(file) {
			b, err := exec.Command(file).Output()
			if err != nil {
				continue
			}
			_, name := filepath.Split(file)
			if ext := filepath.Ext(file); ext != "" {
				name = name[0:len(name)-len(ext)]
			}
			fmt.Println(strings.Replace(string(b), "\n", "\n\t", -1) + "\n")
		}
	}
}
