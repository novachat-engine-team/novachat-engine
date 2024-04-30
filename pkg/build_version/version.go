/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/11/19 10:36
 * @File : version.go
 */

package build_version

import (
	"fmt"
	"os"
	"runtime"
)

var (
	GitTagsName		   = ""
	GitBranchName  = ""
	GitCommitLog   = ""
	User      	   = ""
	BuildTime      = ""
	BuildGoVersion = ""
)

func StringifySingleLine() string {
	return fmt.Sprintf("GitTagsName=%s. GitBranchName=%s. GitCommitLog=%s. User=%s. BuildTime=%s. GoVersion=%s. runtime=%s/%s.",
		GitTagsName, GitBranchName, GitCommitLog, User, BuildTime, BuildGoVersion, runtime.GOOS, runtime.GOARCH)
}

func StringifyMultiLine() string {
	return fmt.Sprintf("GitTagsName=%s\nGitBranchName=%s\nGitCommitLog=%s\nUser=%s\nBuildTime=%s\nGoVersion=%s\nruntime=%s/%s\n",
		GitTagsName, GitBranchName, GitCommitLog, User, BuildTime, BuildGoVersion, runtime.GOOS, runtime.GOARCH)
}

func init() {
}

func FlagVersion() {
	fmt.Printf(StringifyMultiLine())

	args := os.Args
	if len(args) >= 2 && (args[1] == "--version" || args[1] == "-v") {
		os.Exit(0)
		return
	}
}