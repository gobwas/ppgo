package main

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

const script = "src/github.com/gobwas/ppgo/script/ppgo"

func main() {
	var scriptPath string
	for _, p := range strings.Split(os.Getenv("GOPATH"), ":") {
		f := path.Join(p, script)
		if _, err := os.Stat(f); err == nil {
			scriptPath = f
			break
		}
	}
	if scriptPath == "" {
		log.Fatalf("could not find ppgo generation script")
	}
	cmd := exec.Command("/bin/sh", scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
