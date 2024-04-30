/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package command

import "novachat_engine/pkg/config"

type Application interface {
	Initialize(string) (*config.LogConfig, error)
	RunLoop() error
	Close()
	OnEvent(filename string)
}

var ApplicationIns Application
var Watcher *config.Watcher
