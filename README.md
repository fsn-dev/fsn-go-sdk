# Fusion Go SDK Repository

## first run

1. clone the repository

    mkdir -p $GOPATH/src/github.com/FusionFoundation  
    cd $GOPATH/src/github.com/FusionFoundation  
    git clone https://github.com/fsn-dev/fsn-go-sdk.git  
    cd fsn-go-sdk  

2. add vendor packages

    ./scripts/gomod.sh

3. build project

    ./scripts/build.sh fsn-cli
    or use `go build` manually

4. run project

    ./scripts/run.sh fsn-cli
    or ./bin/fsn-cli

## commitment notes

1. please make a new top level directory for each new project
2. please provide help information for each command and sub-commands
3. please run './scripts/gofmt.sh' to format codes before committing
4. please run './scripts/add-license.sh `<newfile>`' to add lincense for new files

## common directories

* efsn		-- fusion base code

    import from `https://github.com/FUSIONFoundation/efsn`

* scripts 	-- scripts used to manage project

    build.sh - build specified projects

    run.sh - run specified project

    gofmt.sh - format `*.go` files

    gomod.sh - import vendor modules

    add-license.sh - add LICENSE content to the file header

* vendor	-- outside modules (use `go mod` to manage)

[//]: # (/* vim: set ts=4 sts=4 sw=4 et : */)
