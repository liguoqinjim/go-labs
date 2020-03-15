#!/usr/bin/env bash

set -e
echo "" > coverage.txt

# 去除vendor文件夹
for d in $(go list ./... | grep -v vendor); do
    go test -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done

# 不在测试覆盖率里面去除lab002
#for d in $(go list ./... | grep -v vendor); do
#    go test -race -coverprofile=profile.out -covermode=atomic $d
#    if [ -f profile.out ]; then
#        cat profile.out >> coverage.txt
#        rm profile.out
#    fi
#done

# 去除lab002
#for d in $(go list ./... | grep -v vendor | grep -v lab002); do
#    go test -race -coverprofile=profile.out -covermode=atomic $d
#    if [ -f profile.out ]; then
#        cat profile.out >> coverage.txt
#        rm profile.out
#    fi
#done