#!/bin/bash

ORG='github.com/qingtao'
GITTAG=`git describe --tags --dirty --always`
GITSHA=`git rev-parse HEAD`

GOFLAGS="$GOFLAGS -X ${ORG}/version.GitSHA=${GITSHA} \
    -X ${ORG}/version.GitTag=${GITTAG} \
"

if [ -f "VERSION" ]; then
    GOFLAGS="$GOFLAGS -X ${ORG}/version.Version=`cat VERSION`"
fi

go build -v -ldflags="$GOFLAGS"

