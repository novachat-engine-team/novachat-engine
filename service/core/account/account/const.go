/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/8/16 17:15
 * @File :
 */

package account

const (
	AccountCacheUserNamePrefix         = "bf:account:username"
	AccountCachePhoneNumberPrefix      = "bf:account:phonenumber"
	AccountCacheUsersUserIdPrefix      = "string:account:userid:"
	AccountCacheUsersPhoneNumberPrefix = "string:account:phonenumber:"
	AccountCacheUsersUsernamePrefix    = "string:account:username:"
)

const (
	BFReserve   = "BF.RESERVE"
	BFAdd       = "BF.ADD"
	BFMAdd      = "BF.MADD"
	BFInsert    = "BF.INSERT"
	BFExists    = "BF.EXISTS"
	BFMExists   = "BF.MEXISTS"
	BFScanDump  = "BF.SCANDUMP"
	BFLoadChunk = "BF.LOADCHUNK"
	BFInfo      = "BF.INFO"
	BFDebug     = "BF.DEBUG"
	SET         = "SET"
	MSET        = "MSET"
	GET         = "GET"
	MGET        = "MGET"
)
