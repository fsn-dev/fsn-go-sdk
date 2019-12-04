# Fusion Go SDK Repository

## first run

1. clone the repository

```
    mkdir -p ${GOPATH:-$HOME/go}/src/github.com/FusionFoundation  
    cd ${GOPATH:-$HOME/go}/src/github.com/FusionFoundation  
    git clone https://github.com/fsn-dev/fsn-go-sdk.git  
    cd fsn-go-sdk  
```

2. build project

```
    # set env GOPROXY if you can't get packages from golang.org
    export GOPROXY=https://goproxy.io

    make fsn-cli (take `fsn-cli` as example here)  
```

3. run project

```
    ./bin/fsn-cli (take `fsn-cli` as example here)  
```

## commitment notes

1. please make a new top level directory for each new project
2. please provide `help information` for each command and sub-commands
3. please run `make fmt` to format codes before committing
4. please run `./scripts/add-license.sh <newfile>` to add lincense for new files

## common directories

* efsn      -- fusion base code

    import from `https://github.com/FUSIONFoundation/efsn`

* fsnapi    -- supply API to build and sign transaction, etc.

* bin       -- binary output directory

* scripts   -- scripts used to manage project

    build.sh       -- build specified projects  
    gofmt.sh       -- format `*.go` files  
    add-license.sh -- add LICENSE content to the file header  

[//]: # (/* vim: set ts=4 sts=4 sw=4 et : */)
