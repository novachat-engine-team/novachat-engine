/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/12 9:31
 * @File : sfs_handler.go
 */

package server

import (
	"novachat_engine/pkg/cmd/sfs/conf"
	partfs2 "novachat_engine/pkg/cmd/sfs/server/partfs"
)

type SfsHandler struct {
	conf *conf.Config
}

func NewSfsHandler(conf *conf.Config) *SfsHandler {
	return &SfsHandler{
		conf: conf,
	}
}

func (m *SfsHandler) Install() {

	if m.conf.RootPathConfig != nil {
		partfs2.ConfigRootSubPath(m.conf.RootPathConfig.T0, partfs2.SubPathsTypeDat)
		partfs2.ConfigRootSubPath(m.conf.RootPathConfig.T1, partfs2.SubPathsTypePart)
		partfs2.ConfigRootSubPath(m.conf.RootPathConfig.T2, partfs2.SubPathsTypePng)
		partfs2.ConfigRootSubPath(m.conf.RootPathConfig.T3, partfs2.SubPathsTypeJpeg)
		partfs2.ConfigRootSubPath(m.conf.RootPathConfig.T4, partfs2.SubPathsTypeFGif)
		partfs2.ConfigRootSubPath(m.conf.RootPathConfig.T5, partfs2.SubPathsTypePdf)
		partfs2.ConfigRootSubPath(m.conf.RootPathConfig.T6, partfs2.SubPathsTypeMov)
		partfs2.ConfigRootSubPath(m.conf.RootPathConfig.T7, partfs2.SubPathsTypeMp3)
		partfs2.ConfigRootSubPath(m.conf.RootPathConfig.T8, partfs2.SubPathsTypeMp4)
		partfs2.ConfigRootSubPath(m.conf.RootPathConfig.T9, partfs2.SubPathsTypeWebp)
		partfs2.ConfigRootSubPath(m.conf.RootPathConfig.T10, partfs2.SubPathsTypeStickerSet)
	}

	if len(m.conf.RootPath) > 0 {
		partfs2.InstallPartFS(m.conf.RootPath)
	} else if m.conf.RootPathConfig != nil && len(m.conf.RootPathConfig.RootPath) > 0 {
		partfs2.InstallPartFS(m.conf.RootPathConfig.RootPath)
	}
}
