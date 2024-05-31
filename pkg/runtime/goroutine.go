/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package runtime

import (
	"fmt"
	"novachat_engine/pkg/log"
	"os"
	"runtime/debug"
	"time"
)

var debugIgnoreStdout = false

// GoSafeWithRecover wraps a `go func()` with recover()
func GoSafeWithRecover(handler func([]interface{}), params []interface{}, recoverHandler func(r interface{})) {
	go func(param []interface{}) {
		defer func() {
			if r := recover(); r != nil {
				// TODO: log
				if !debugIgnoreStdout {
					log.Errorf("goroutine panic: %v\n%s\n", r, string(debug.Stack()))
				}
				if recoverHandler != nil {
					go func() {
						defer func() {
							if p := recover(); p != nil {
								if !debugIgnoreStdout {
									log.Errorf("recover goroutine panic:%v\n%s\n", p, string(debug.Stack()))
								}
							}
						}()
						recoverHandler(r)
					}()
				}
			}
		}()
		handler(param)
	}(params)
}

// GoWithRecover wraps a `go func()` with recover()
func GoWithRecover(handler func(), recoverHandler func(r interface{})) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				// TODO: log
				if !debugIgnoreStdout {
					fmt.Fprintf(os.Stderr, "%s goroutine panic: %v\n%s\n",
						time.Now(), r, string(debug.Stack()))
				}
				if recoverHandler != nil {
					go func() {
						defer func() {
							if p := recover(); p != nil {
								if !debugIgnoreStdout {
									fmt.Fprintf(os.Stderr, "recover goroutine panic:%v\n%s\n", p, string(debug.Stack()))
								}
							}
						}()
						recoverHandler(r)
					}()
				}
			}
		}()
		handler()
	}()
}

func GoSafeCall(fn func(), recoverHandler func(r interface{})) {
	defer func() {
		if err := recover(); err != nil {
			if !debugIgnoreStdout {
				fmt.Fprintf(os.Stderr, "%s safe call panic: %v\n%s\n",
					time.Now(), err, string(debug.Stack()))
			}
			if recoverHandler != nil {
				go func() {
					defer func() {
						if p := recover(); p != nil {
							if !debugIgnoreStdout {
								fmt.Fprintf(os.Stderr, "recover goroutine panic:%v\n%s\n", p, string(debug.Stack()))
							}
						}
					}()
					recoverHandler(err)
				}()
			}
		}
	}()

	fn()
}
