# Fusion Go SDK Repository

## first run

1. clone the repository

    mkdir -p $GOPATH/src/github.com/FusionFoundation  
    cd $GOPATH/src/github.com/FusionFoundation  
    git clone https://github.com/fsn-dev/fsn-go-sdk.git  
    cd fsn-go-sdk  

2. add vendor packages

    make vendor  
    or, make vendor_with_proxy (set goproxy if you can't get packages from golang.org)

3. build project

    make fsn-cli (take `fsn-cli` as example here)  

4. run project

    ./bin/fsn-cli (take `fsn-cli` as example here)  

## commitment notes

1. please make a new top level directory for each new project
2. please provide `help information` for each command and sub-commands
3. please run `make fmt` to format codes before committing
4. please run `./scripts/add-license.sh <newfile>` to add lincense for new files

## common directories

* efsn      -- fusion base code

    import from `https://github.com/FUSIONFoundation/efsn`

* vendor    -- outside modules (use `go mod` to manage)

* fsnapi    -- supply API to build and sign transaction, etc.

* bin       -- binary output directory

* scripts   -- scripts used to manage project

    build.sh - build specified projects

    run.sh - run specified project

    gofmt.sh - format `*.go` files

    gomod.sh - import vendor modules

    add-license.sh - add LICENSE content to the file header

[//]: # (/* vim: set ts=4 sts=4 sw=4 et : */)
