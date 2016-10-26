package main

import (
	"exec"
	"log"
	"os"
	"path"
	"strings"
)

const script = "src/github.com/gobwas/ppgo/script/ppgo"

func main() {
	var scriptPath string
	for _, p := range strings.Split(os.Getenv("GOPATH"), ":") {
		f := path.Join(p, script)
		if s, err := os.Stat(f); err == nil {
			scriptPath = f
			break
		}
	}
	if scriptPath == "" {
		log.Fatalf("could not find ppgo generation script")
	}
	cmd := exec.Command("/bin/sh", scriptPath)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
