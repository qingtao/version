#!/bin/bash

ORG='github.com/qingtao'
GITTAG=`git describe --tags --dirty --always`
GITSHA=`git rev-parse HEAD`
GITBRANCH=`git rev-parse --abbrev-ref HEAD`

GOFLAGS="$GOFLAGS -X ${ORG}/version.GitSHA=${GITSHA} \
    -X ${ORG}/version.GitTag=${GITTAG} \
    -X ${ORG}/version.GitBranch=${GITBRANCH} \
"

if [ -f "VERSION" ]; then
    GOFLAGS="$GOFLAGS -X ${ORG}/version.Version=`cat VERSION`"
fi

go build -v -ldflags="$GOFLAGS"

