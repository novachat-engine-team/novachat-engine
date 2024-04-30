/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package command

import (
	"errors"
	"fmt"
	"novachat_engine/pkg/build_version"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/runtime"
	"novachat_engine/pkg/syscall"
	"os"
	"os/signal"
	"sync"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "name",
	Short: "name service",
	Long:  "Demo name service",
}

func cmdRun(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		fmt.Println("invalid params")

		cmd.Help()
		return
	}

	logConf, err := ApplicationIns.Initialize(args[0])
	if err != nil {
		fmt.Println(err.Error())

		cmd.Help()
		return
	}

	var options *log.Options
	if logConf != nil {
		options = log.ConfigOptions(logConf.LogPath, logConf.LogLevel, logConf.RotateMaxSize, logConf.RotateMaxAge, logConf.RotateMaxBackups)
	} else {
		options = log.DefaultOptions()
	}
	options.AttachCobraFlags(RootCmd)
	if err = log.Configure(options); err != nil {
		panic(err)
	}

	Watcher, err = config.NewWatcher(ApplicationIns)
	if err != nil {
		panic(fmt.Errorf("watcher error:%s", err.Error()))
	}
	Watcher.Add(args[0])
	Watcher.Run()

	defer log.Sync()

	wg := sync.WaitGroup{}
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP)

	closeChan := make(chan struct{})
	wg.Add(1)
	runtime.GoWithRecover(func() {
		defer wg.Done()
		close(closeChan)

		<-sigChan
		ApplicationIns.Close()

		log.Debug("shutdown")
	}, nil)
	<-closeChan

	wg.Add(1)
	runtime.GoWithRecover(func() {
		defer wg.Done()

		err = ApplicationIns.RunLoop()
		if err != nil {
			close(sigChan)
		}
	}, func(r interface{}) {
		close(sigChan)
	})

	log.Debugf("start ok!!")
	wg.Wait()
}

func Run(args []string, app Application) error {
	build_version.FlagVersion()
	fmt.Println("")

	if app == nil {
		return errors.New("app is nil, exit")
	}

	RootCmd.Run = cmdRun

	RootCmd.SetArgs(args)

	ApplicationIns = app

	err := RootCmd.Execute()
	if err != nil {
		return err
	}

	return nil
}
