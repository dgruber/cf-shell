#!/usr/bin/env bash
set -e

mkdir -p $GOPATH/src/github.com/dgruber
cp -r cf-shell $GOPATH/src/github.com/dgruber

cd $GOPATH/src/github.com/dgruber/cf-shell

./xcompile.sh builds cf-shell 0.0.0
