/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/5/26 15:19
 * @File : video_call.go
 */

package video_call

type VideoCall struct {}

func NewVideoCall() *VideoCall{
	return &VideoCall{}
}

func (c VideoCall) VideoCallAvailableConf() bool {
	//TODO:(Coder)
	return true
}
