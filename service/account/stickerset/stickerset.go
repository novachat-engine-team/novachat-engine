/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package stickerset

import (
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/cache"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/util"
	"novachat_engine/service/core/stickerset"
	data_stickerset "novachat_engine/service/data/stickerset"
	"novachat_engine/service/upload/photo"
	"sort"
)

type Core struct {
	core        *stickerset.Core
	redisClient *cache.RedisClient
	favedLimit  int32
	recentMax   int32
}

func NewStickerSetCore(core *stickerset.Core, redisConf config.RedisConfig, favedLimit int32, recentMax int32) *Core {
	cache.InstallRedis(redisConf)
	if favedLimit <= 0 {
		favedLimit = defaultFavedLimit
	}
	if recentMax <= 0 {
		recentMax = defaultRecentLimit
	}

	return &Core{
		core:        core,
		favedLimit:  favedLimit,
		recentMax:   recentMax,
		redisClient: cache.GetRedisClient(),
	}
}

func (s *Core) GetAllStickers() ([]*data_stickerset.StickerSet, error) {
	stickerSetIdList, err := s.cache2AllStickersId()
	if err != nil {
		log.Warnf("GetAllStickers SMEMBERS error:%s", err.Error())
	}

	var stickerSetList []*data_stickerset.StickerSet
	if len(stickerSetIdList) == 0 {
		stickerSetList, err = s.core.GetAllStickers()
		if err != nil {
			log.Errorf("GetAllStickers from db error:%s", err.Error())
			return nil, err
		}

		if len(stickerSetList) > 0 {
			err = s.allStickersId2Cache(stickerSetList)
			if err != nil {
				log.Warnf("GetAllStickers allStickersId2Cache :%s", err.Error())
			}
			err = s.stickers2Cache(stickerSetList)
			if err != nil {
				log.Warnf("GetAllStickers stickers2Cache :%s", err.Error())
			}
		}
	} else {
		stickerSetList, err = s.GetStickers(stickerSetIdList)
		if err != nil {
			log.Errorf("GetAllStickers cache2stickers :%s", err.Error())
			return nil, err
		}
	}
	return stickerSetList, nil
}

func (s *Core) GetStickers(idList []int64) ([]*data_stickerset.StickerSet, error) {
	stickerSetList, err := s.cache2stickers(idList)
	if err != nil {
		log.Warnf("GetStickers cache2stickers :%s", err.Error())
	}

	lessIdList := make([]int64, 0, len(idList))
	for _, v := range idList {
		if util.Index(&stickerSetList, func(idx int) bool {
			return stickerSetList[idx].Id == v
		}) < 0 {
			lessIdList = append(lessIdList, v)
		}
	}

	if len(lessIdList) > 0 {
		lessStickerList, err2 := s.core.GetStickers(lessIdList)
		if err2 != nil {
			log.Errorf("GetStickers GetStickers error:%s", err.Error())
			return nil, err2
		}
		stickerSetList = append(stickerSetList, lessStickerList...)
		err = s.stickers2Cache(lessStickerList)
		if err != nil {
			log.Warnf("GetStickers stickers2Cache :%s", err.Error())
		}
	}
	return stickerSetList, nil
}

func (s *Core) GetInstalled(userId int64) ([]*data_stickerset.StickerInstall, error) {
	l, err := s.cache2Installed(userId)
	if err != nil {
		log.Warnf("GetInstalled cache2Installed :%s", err.Error())
	}
	if len(l) == 0 {
		l, err = s.core.Installed(userId)
		if err != nil {
			log.Errorf("GetInstalled Installed :%s", err.Error())
			return nil, err
		}
		sort.SliceStable(l, func(i, j int) bool {
			return l[i].Date < l[j].Date
		})

		err = s.installed2Cache(userId, l)
		if err != nil {
			log.Warnf("GetInstalled installed2Cache :%s", err.Error())
		}
	}
	return l, nil
}

func (s *Core) Install(userId int64, stickerSetId int64, date int32, archived bool) (*data_stickerset.StickerSet, error) {
	log.Debugf("Install userId:%d stickerSetId：%d date:%d archived:%v", userId, stickerSetId, date, archived)

	stickerSetList, err := s.GetStickers([]int64{stickerSetId})
	if err != nil {
		log.Errorf("Install GetStickers :%s", err.Error())
		return nil, err
	}
	if len(stickerSetList) == 0 {
		err = fmt.Errorf("install GetStickers not found:%d", stickerSetId)
		log.Error(err.Error())
		return nil, err
	}

	l, err := s.cache2Installed(userId)
	if err != nil {
		log.Warnf("Install cache2Installed :%s", err.Error())
	}

	var install *data_stickerset.StickerInstall
	if idx := util.Index(&l, func(idx int) bool {
		return l[idx].StickerSetId == stickerSetId
	}); idx < 0 {
		install = &data_stickerset.StickerInstall{
			UserId:       userId,
			StickerSetId: stickerSetId,
			Installed:    false,
			Date:         date,
		}
		l = append(l, install)
	} else {
		if !l[idx].Installed {
			l[idx].Installed = false
			l[idx].Date = date
			install = l[idx]
		}
	}

	if install != nil {
		err = s.core.Install(userId, stickerSetId, true, date)
		if err != nil {
			log.Errorf("Install :%s", err.Error())
			return nil, err
		}
		install.Installed = true
		install.Archived = archived
		err = s.installed2Cache(userId, l)
		if err != nil {
			log.Warnf("Install installed2Cache :%s", err.Error())
		}
	}
	return stickerSetList[0], nil
}

func (s *Core) Uninstall(userId int64, stickerSetId int64) error {
	l, err := s.cache2Installed(userId)
	if err != nil {
		log.Warnf("Uninstall cache2Installed :%s", err.Error())
	}

	var install *data_stickerset.StickerInstall
	if idx := util.Index(&l, func(idx int) bool {
		return l[idx].StickerSetId == stickerSetId
	}); idx >= 0 {
		if !l[idx].Installed {
			l[idx].Installed = false
			install = l[idx]
		}
	}

	if install != nil {
		err = s.core.Install(userId, stickerSetId, false, 0)
		if err != nil {
			log.Errorf("Uninstall :%s", err.Error())
			return err
		}
		install.Installed = false
		err = s.installed2Cache(userId, l)
		if err != nil {
			log.Warnf("Uninstall installed2Cache :%s", err.Error())
		}
	}
	return nil
}

func (s *Core) Faved(userId int64, fileId int64, fave bool) error {
	favedList, err := s.cache2Faved(userId)
	if err != nil {
		log.Warnf("Faved cache2Faved error:%s", err.Error())

		favedList, err = s.core.Faved(userId)
		if err != nil {
			log.Errorf("Faved error:%s", err.Error())
			return err
		}
	}
	update := false
	if idx := util.IndexInt64s(&favedList, fileId); idx >= 0 {
		if !fave {
			copy(favedList[idx:], favedList[idx+1:])
			favedList = favedList[:len(favedList)-1]
			update = true
		}
	} else {
		if fave {
			if int32(len(favedList)) > s.favedLimit {
				copy(favedList, favedList[1:len(favedList)-1])
				favedList[len(favedList)-1] = fileId
			} else {
				favedList = append(favedList, fileId)
			}
			update = true
		}
	}
	if update {
		err = s.core.UpdateFaved(userId, favedList)
		if err != nil {
			log.Errorf("Faved UpdateFaved error:%s", err.Error())
			return err
		}

		err = s.faved2Cache(userId, favedList)
		if err != nil {
			log.Warnf("Faved faved2Cache error:%s", err.Error())
		}
	}
	return nil
}

func (s *Core) Recent(userId int64, fileId int64, save bool) error {
	//TODO:(Coderxw) thinking find stickerSetId and fileId to save
	recentList, err := s.cache2Recent(userId)
	if err != nil {
		log.Warnf("Recent cache2Recent error:%s", err.Error())

		recentList, err = s.core.Recent(userId)
		if err != nil {
			log.Errorf("Recent error:%s", err.Error())
			return err
		}
	}
	update := false
	if idx := util.IndexInt64s(&recentList, fileId); idx >= 0 {
		if !save {
			copy(recentList[idx:], recentList[idx+1:])
			recentList = recentList[:len(recentList)-1]
			update = true
		}
	} else {
		if save {
			if int32(len(recentList)) > s.favedLimit {
				copy(recentList, recentList[1:len(recentList)-1])
				recentList[len(recentList)-1] = fileId
			} else {
				recentList = append(recentList, fileId)
			}
			update = true
		}
	}
	if update {
		err = s.core.UpdateRecent(userId, recentList)
		if err != nil {
			log.Errorf("Recent UpdateRecent error:%s", err.Error())
			return err
		}

		err = s.recent2Cache(userId, recentList)
		if err != nil {
			log.Warnf("Recent recent2Cache error:%s", err.Error())
		}
	}
	return nil
}

func (s *Core) ClearRecent(userId int64) error {
	err := s.core.UpdateRecent(userId, []int64{})
	if err != nil {
		log.Errorf("ClearRecent UpdateRecent error:%s", err.Error())
		return err
	}

	err = s.recent2Cache(userId, []int64{})
	if err != nil {
		log.Warnf("ClearRecent recent2Cache error:%s", err.Error())
	}
	return nil
}

func (s *Core) AllStickersByUser(userId int64, layer int32) ([]*mtproto.StickerSet, error) {
	stickerSetList, err := s.GetAllStickers()
	if err != nil {
		log.Errorf("AllStickersByUser GetAllStickers error:%s", err.Error())
		return nil, err
	}

	var installStickerIdList []*data_stickerset.StickerInstall
	if userId != 0 {
		installStickerIdList, err = s.GetInstalled(userId)
		if err != nil {
			log.Errorf("AllStickersByUser GetInstalled error:%s", err.Error())
			return nil, err
		}
	}
	resp := make([]*mtproto.StickerSet, 0, len(stickerSetList))
	for _, v := range stickerSetList {
		installed := false
		installedDate := int32(0)
		if idx := util.Index(&installStickerIdList, func(idx int) bool {
			return installStickerIdList[idx].StickerSetId == v.Id
		}); idx >= 0 {
			installed = installStickerIdList[idx].Installed
			installedDate = installStickerIdList[idx].Date
		} else {
			installed = false
		}
		photoSizes := make([]*mtproto.PhotoSize, 0, len(v.Thumbs))
		for _, thumbs := range v.Thumbs {
			for _, thumb := range thumbs.Thumbs {
				photoSizes = append(photoSizes, photo.PhotoData2PhotoSize(thumb, layer))
			}
		}
		var photoSize *mtproto.PhotoSize
		if len(photoSizes) > 0 {
			photoSize = photoSizes[0]
		}
		resp = append(resp, mtproto.NewTLStickerSet(&mtproto.StickerSet{
			Installed:     installed,
			Archived:      v.Animated,
			Official:      v.Official,
			Masks:         v.Masks,
			Id:            v.Id,
			AccessHash:    v.AccessHash,
			Title:         v.Title,
			ShortName:     v.ShortName,
			Count:         v.Count,
			Hash:          v.Hash,
			InstalledDate: installedDate,
			Thumb:         photoSize,
			ThumbDcId:     v.ThumbDcId,
			Animated:      v.Animated,
			Thumbs:        photoSizes,
		}).To_StickerSet())
	}

	return resp, nil
}

func (s *Core) GetStickersByEmoticon(emoticon string) ([]*data_stickerset.StickerSet, error) {
	stickerList, err := s.GetAllStickers()
	if err != nil {
		log.Errorf("GetStickersByEmoticon error:%s", err.Error())
		return nil, err
	}

	foundList := make([]*data_stickerset.StickerSet, 0, len(stickerList))
	for _, sticker := range stickerList {
		if idx := util.IndexStrings(&sticker.Emoticons, emoticon); idx >= 0 {
			foundList = append(foundList, sticker)
			continue
		}
	}

	log.Debugf("GetStickersByEmoticon emotion:%s found:%d", emoticon, len(foundList))
	return foundList, nil
}

func (s *Core) GetStickersByShortName(name string) (*data_stickerset.StickerSet, error) {
	stickerList, err := s.GetAllStickers()
	if err != nil {
		log.Errorf("GetStickersByEmoticon error:%s", err.Error())
		return nil, err
	}

	var found *data_stickerset.StickerSet
	for _, sticker := range stickerList {
		if sticker.ShortName == name {
			found = sticker
			break
		}
	}

	log.Debugf("GetStickersByShortName name:%s found:%v", name, found)
	return found, nil
}

func (s *Core) GetRecent(userId int64, attached bool) ([]*data_stickerset.StickerSet, error) {
	//TODO:(Coderxw)
	// find all sticker document need to do....
	recentList, err := s.cache2Recent(userId)
	if err != nil {
		log.Warnf("GetRecent cache2Recent :%s", err.Error())

		recentList, err = s.core.Recent(userId)
		if err != nil {
			log.Errorf("GetRecent Recent error:%s", err.Error())
			return nil, err
		}

		if err = s.recent2Cache(userId, recentList); err != nil {
			log.Warnf("GetRecent recent2Cache :%s", err.Error())
		}
	}
	if len(recentList) == 0 {
		return nil, nil
	}

	allStickerList, err := s.GetAllStickers()
	if err != nil {
		log.Errorf("GetRecent GetAllStickers error:%s", err.Error())
		return nil, err
	}
	stickerSetList := make([]*data_stickerset.StickerSet, 0, len(recentList))
	for _, v := range allStickerList {
		for _, pack := range v.Packs {
			if idx := util.Index(&pack.Documents, func(idx int) bool {
				return util.IndexInt64s(&recentList, pack.Documents[idx].VolumeId) >= 0
			}); idx >= 0 {
				stickerSetList = append(stickerSetList, v)
				break
			}
		}
	}

	log.Debugf("GetRecent userId:%d value:%d", userId, len(stickerSetList))
	return stickerSetList, nil
}

func (s *Core) GetFaved(userId int64) ([]*data_stickerset.StickerSet, error) {
	//TODO:(Coderxw)
	// find all sticker document need to do....
	favedList, err := s.cache2Faved(userId)
	if err != nil {
		log.Warnf("GetFaved cache2Faved :%s", err.Error())

		favedList, err = s.core.Faved(userId)
		if err != nil {
			log.Errorf("GetFaved Faved error:%s", err.Error())
			return nil, err
		}

		if err = s.faved2Cache(userId, favedList); err != nil {
			log.Warnf("GetFaved faved2Cache :%s", err.Error())
		}
	}
	if len(favedList) == 0 {
		return nil, nil
	}

	allStickerList, err := s.GetAllStickers()
	if err != nil {
		log.Errorf("GetFaved GetAllStickers error:%s", err.Error())
		return nil, err
	}
	stickerSetList := make([]*data_stickerset.StickerSet, 0, len(favedList))
	for _, v := range allStickerList {
		for _, pack := range v.Packs {
			if idx := util.Index(&pack.Documents, func(idx int) bool {
				return util.IndexInt64s(&favedList, pack.Documents[idx].VolumeId) >= 0
			}); idx >= 0 {
				stickerSetList = append(stickerSetList, v)
				break
			}
		}
	}

	log.Debugf("GetFaved userId:%d value:%d", userId, len(stickerSetList))
	return stickerSetList, nil
}
