/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/26 17:23
 * @File : const_profile.go
 */

package constants

const (
	EnvModelProd = "prod"
	EnvModelTest = "test"

	OfficialAPPID = 1000000
	// id = 42777 429000
	//return user != null && (user.support || user.id == 777000 ||
	//user.id == 333000 || user.id == 4240000 || user.id == 4244000 ||
	//user.id == 4245000 || user.id == 4246000 || user.id == 410000 ||
	//user.id == 420000 || user.id == 431000 || user.id == 431415000 ||
	//user.id == 434000 || user.id == 4243000 || user.id == 439000 ||
	//user.id == 449000 || user.id == 450000 || user.id == 452000 ||
	//user.id == 454000 || user.id == 4254000 || user.id == 455000 ||
	//user.id == 460000 || user.id == 470000 || user.id == 479000 ||
	//user.id == 796000 || user.id == 482000 || user.id == 490000 ||
	//user.id == 496000 || user.id == 497000 || user.id == 498000 ||
	//user.id == 4298000);
	OfficialSupportUserID = int64(777000)
	BotFatherId           = int64(93372553)
	MessageLongLimit      = 4 * 1024 //getConfig message_length_max
	UserAboutLimit        = 500

	ChatMaxLimit     = 200 // migrate channel
	JoinChatMaxLimit = 100 //
)

type AccountStatus int8

const (
	AccountStatusNormal     AccountStatus = 0
	AccountStatusForbid     AccountStatus = 1
	AccountStatusRestricted AccountStatus = 2
)
