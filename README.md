# cf-shell - A Shell for Cloud Foundry Commands with Auto-Completion

cf-shell is a simple plugin for the Cloud Foundry *cf* command line tool which 
supports auto-completion of *cf* arguments. Its core functionality is based on
_go-prompt_.

## Installation

### By Source

    mkdir -p $GOPATH/src/github.com/dgruber
    cd $GOPATH/src/github.com/dgruber
    git clone $GOPATH/src/github.com/dgruber/cf-shell
    cd cf-shell
    go build
    cf install-plugin ./cf-shell

### Pre-Build

For MacOS X:

    cf install-plugin https://github.com/dgruber/cf-shell/blob/master/builds/current/cf-shell.osx?raw=true

## Usage

    cf shell

Typing _quit_ or _exit_ (or ctrl-d) closes the shell.

## Uninstall plugin

    cf uninstall-plugin Shell

## Screenshots 

Arguments

![cli completion](images/example.png?raw=true "arguments")

...and options

![cli completion](images/example_context.png?raw=true "options")


