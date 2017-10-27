#!/usr/bin/env bash
set -e

git config --global user.email "cf-shell@danielgruber.com"
git config --global user.name "cf-shell-release-process"

version=`cat version/version`
gosrc="$GOPATH/src/github.com/dgruber"

echo "building version $version"

mkdir -p $gosrc
cp -r cf-shell $gosrc
pushd $gosrc/cf-shell
./xcompile.sh builds cf-shell $version 
popd

git clone cf-shell cf-shell-new-version
pushd cf-shell-new-version/builds
cp -r $gosrc/cf-shell/builds/$version $version
cp -r $gosrc/cf-shell/builds/current current
git add $version/*
git add current/*
git commit -m "[ci skip] new version build $version"

