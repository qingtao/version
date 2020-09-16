#!/bin/bash

ORG='github.com/qingtao/version'
GITTAG=`git describe --tags --dirty --always`
GITSHA=`git rev-parse HEAD`

GOFLAGS="$GOFLAGS -X ${ORG}/version.GitSHA=${GITSHA} \
    -X ${ORG}/version.GitTag=${GITTAG} \
"

echo $GOFLAGS

