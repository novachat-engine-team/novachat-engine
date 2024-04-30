/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/11/2 14:13
 * @File : watcher.go
 */

package config

import (
	"fmt"
	"github.com/howeyc/fsnotify"
	"log"
	"os"
)

type WatcherEvent interface {
	OnEvent(string)
}

type Watcher struct {
	watcher *fsnotify.Watcher
	_cb WatcherEvent
}

func NewWatcher(onEvent WatcherEvent) (*Watcher, error) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &Watcher{watcher: w, _cb: onEvent}, nil
}

func (m *Watcher) Add(name string) error {
	ok, err := _pathExits(name)
	if err != nil {
		return err
	}
	if ok == false {
		return fmt.Errorf("not found name:%s", name)
	}

	return m.watcher.Watch(name)
}

func (m *Watcher) Close() {
	m.watcher.Close()
}

func (m *Watcher) Run() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()

		for {
			select {
			case event := <-m.watcher.Event:
				if event.IsModify() {
					if m._cb != nil {
						m._cb.OnEvent(event.Name)
					} else {
						log.Printf("file modify name:%s\n", event.Name)
					}
				}
			case err := <-m.watcher.Error:
				log.Println(err.Error())
				return
			}
		}
	}()
}

func _pathExits(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if fileInfo != nil {
		return true, nil
	}

	if os.IsNotExist(err) == true {
		return false, nil
	}

	return false, err
}
