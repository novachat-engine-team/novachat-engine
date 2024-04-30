/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/8/16 17:15
 * @File :
 */

package account

import (
	"novachat_engine/pkg/cache"
)

type Core struct {
	redisClient *cache.RedisClient
}

func NewAccountCore(r *cache.RedisClient) *Core {
	//if mgo.GetMongoDB() == nil {
	//	panic("Need InstallMongodb")
	//}
	return &Core{redisClient: r}
}
