#!/usr/bin/env bash
set -e

go get github.com/onsi/ginkgo/ginkgo
go install github.com/onsi/ginkgo/ginkgo

mkdir -p $GOPATH/src/github.com/dgruber
cp -r cf-shell $GOPATH/src/github.com/dgruber

cd $GOPATH/src/github.com/dgruber/cf-shell/cfcli

ginkgo

