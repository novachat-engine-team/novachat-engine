/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/9 17:52
 * @File : config.go
 */

package conf

import (
	"novachat_engine/pkg/config"
)

//SubPathsTypeDat  	SubPathsType = 0
//SubPathsTypePart  	SubPathsType = 1
//SubPathsTypePng 	SubPathsType = 2
//SubPathsTypeJpeg  	SubPathsType = 3
//SubPathsTypeFGif  	SubPathsType = 4
//SubPathsTypePdf  	SubPathsType = 5
//SubPathsTypeMov  	SubPathsType = 6
//SubPathsTypeMp3  	SubPathsType = 7
//SubPathsTypeMp4  	SubPathsType = 8
//SubPathsTypeWebp  	SubPathsType = 9
//SubPathsTypeStickerSet	SubPathsType = 10
type RootPathConfig struct {
	RootPath string `toml:"rootPath"`
	T0       string `toml:"0"`
	T1       string `toml:"1"`
	T2       string `toml:"2"`
	T3       string `toml:"3"`
	T4       string `toml:"4"`
	T5       string `toml:"5"`
	T6       string `toml:"6"`
	T7       string `toml:"7"`
	T8       string `toml:"8"`
	T9       string `toml:"9"`
	T10      string `toml:"10"`
}

type Config struct {
	Server         config.ServerConfig
	Log            config.LogConfig
	Pprof          config.PprofConfig
	RpcServer      config.RpcServer
	RpcDiscovery   config.EtcdServerConfig
	Redis          config.RedisConfig
	MongoDB        config.MongodbConfig
	RootPath       string
	RootPathConfig *RootPathConfig
}
