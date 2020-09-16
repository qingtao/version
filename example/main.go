package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/qingtao/version"
)

func main() {
	name := filepath.Base(os.Args[0])
	log.Printf("app_name=(%s)\n", name)
	log.Printf("app_version=(%s)\n", version.Version)
	log.Printf("git_branch=(%s)\n", version.GitBranch)
	log.Printf("git_tag=(%s)\n", version.GitTag)
	log.Printf("git_sha=(%s)\n", version.GitSHA)
	log.Printf("%s_%s\n", name, version.Version)
	log.Printf("%s_%s_%s\n", version.GitBranch, version.GitTag, version.GitSHA)
}
