package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/qingtao/version"
)

func main() {
	name := filepath.Base(os.Args[0])
	log.Printf("appName=(%s)\n", name)
	log.Printf("appVersion=(%s)\n", version.Version)
	log.Printf("gitTag=(%s)\n", version.GitTag)
	log.Printf("gitSha1=(%s)\n", version.GitSHA)
}
