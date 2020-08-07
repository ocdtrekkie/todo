#!/bin/bash
set -euo pipefail

export GOPATH=$HOME/go

[ -d $GOPATH/src/github.com/prologic/ ] || mkdir -p $GOPATH/src/github.com/prologic/
[ -L $GOPATH/src/github.com/prologic/todo ] || \
	ln -s /opt/app $GOPATH/src/github.com/prologic/todo

cd /opt/app
go get -v -d ./...
make build
exit 0
