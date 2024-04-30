#!/bin/bash

GitCommitLog=`git log --pretty=oneline -n 1`
GitCommitLog=${GitCommitLog//\'/\"}
GitBranchName=`git rev-parse --abbrev-ref HEAD`
GitBranchName=${GitBranchName//\'/\"}
GitTagsName=`git describe --tags $(git rev-list --tags --max-count=1)`
WHO=`whoami`
BuildTime=`date +'%Y.%m.%d.%H%M%S'`
BuildGoVersion=`go version`

#    -s -w \

LDFlags=" \
    -X 'novachat_engine/pkg/build_version.GitTagsName=${GitTagsName}' \
    -X 'novachat_engine/pkg/build_version.GitBranchName=${GitBranchName}' \
    -X 'novachat_engine/pkg/build_version.GitCommitLog=${GitCommitLog}' \
    -X 'novachat_engine/pkg/build_version.User=${WHO}' \
    -X 'novachat_engine/pkg/build_version.BuildTime=${BuildTime}' \
    -X 'novachat_engine/pkg/build_version.BuildGoVersion=${BuildGoVersion}' \
"

export LDFlags=$LDFlags
