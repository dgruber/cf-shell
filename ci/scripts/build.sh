#!/usr/bin/env bash
set -e

go get github.com/onsi/ginkgo/ginkgo

cd cf-shell 
go build