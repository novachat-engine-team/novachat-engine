/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package users

import (
	"context"
	"novachat_engine/mtproto"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/cmd/sfs/utils"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/util"
	"novachat_engine/service/account/contact"
	"novachat_engine/service/core/account/account"
	"novachat_engine/service/core/account/users"
	"novachat_engine/service/core/account/users/sql_core"
	data_contact "novachat_engine/service/data/contact"
	data_fs "novachat_engine/service/data/fs"
	data_users "novachat_engine/service/data/users"
	"novachat_engine/service/upload/photo"
)

type Core struct {
	contact     *contact.Core
	accountCore *account.Core
	usersCore   *sql_core.Core
}

func NewUsersCore(accountCore *account.Core, contact *contact.Core, usersCore *sql_core.Core) *Core {
	return &Core{
		accountCore: accountCore,
		contact:     contact,
		usersCore:   usersCore,
	}
}

func (c *Core) GetUserList(userId int64, userIdList []int64, layer int32) ([]*mtproto.User, error) {
	if len(userIdList) == 0 {
		return nil, nil
	}

	contacts, err := c.contact.GetContactByIdList(userId, userIdList)
	if err != nil {
		log.Warnf("GetUserList - GetContactByIdList error:%s", err.Error())
	}

	contactsCacheMap := make(map[int64]*data_contact.Contact, len(contacts))
	for _, v := range contacts {
		contactsCacheMap[v.PeerId] = v
	}

	photoIdList := make([]int64, 0, len(userIdList))
	videoIdList := make([]int64, 0, len(userIdList))
	userPhotoMap := make(map[int64]*data_users.ProfilePhotoData, len(userIdList))
	userInfoList := make([]*data_users.Users, 0, len(userIdList))
	userCacheList, err := c.accountCore.UserList(userIdList)
	if err != nil {
		log.Warnf("GetUserList UserList error:%s", err.Error())
	} else {
		var userInfo *data_users.Users
		for _, v := range userCacheList {
			if len(v) == 0 {
				continue
			}
			userInfo, _ = account.CacheInfo2User(v)
			if userInfo == nil {
				continue
			}
			userInfoList = append(userInfoList, userInfo)
			if len(userInfo.Photos) > 0 {
				if profilePhotoData := users.MakeProfilePhotoData(userInfo.Photos); profilePhotoData.Default != 0 {
					if idx := util.Index(profilePhotoData.IdList, func(idx int) bool {
						return profilePhotoData.IdList[idx].PhotoId == profilePhotoData.Default
					}); idx >= 0 {
						pd := profilePhotoData.IdList[idx]
						userPhotoMap[userInfo.Id] = pd
						if profilePhotoData.IdList[idx].VideoId == 0 {
							if idx = util.Index(photoIdList, func(idx int) bool {
								return photoIdList[idx] == pd.PhotoId
							}); idx < 0 {
								photoIdList = append(photoIdList, pd.PhotoId)
							}
						} else {
							if idx = util.Index(videoIdList, func(idx int) bool {
								return videoIdList[idx] == pd.VideoId
							}); idx < 0 {
								videoIdList = append(videoIdList, pd.VideoId)
							}
						}
					}
				}
			}
		}
	}

	loseUserIdList := make([]int64, 0, len(userIdList))
	for _, v := range userIdList {
		if util.Index(userInfoList, func(idx int) bool {
			return userInfoList[idx].Id == v
		}) < 0 {
			loseUserIdList = append(loseUserIdList, v)
		}
	}
	if len(loseUserIdList) > 0 {
		localUserList, err := c.usersCore.UserIdList(loseUserIdList)
		if err != nil {
			log.Warnf("GetUserList UserIdList error:%s", err.Error())
		}
		for _, v := range localUserList {
			if len(v.Photos) > 0 {
				if profilePhotoData := users.MakeProfilePhotoData(v.Photos); profilePhotoData.Default != 0 {
					if idx := util.Index(profilePhotoData.IdList, func(idx int) bool {
						return profilePhotoData.IdList[idx].PhotoId == profilePhotoData.Default
					}); idx >= 0 {
						pd := profilePhotoData.IdList[idx]
						userPhotoMap[v.Id] = pd
						if profilePhotoData.IdList[idx].VideoId == 0 {
							if idx = util.Index(photoIdList, func(idx int) bool {
								return photoIdList[idx] == pd.PhotoId
							}); idx < 0 {
								photoIdList = append(photoIdList, pd.PhotoId)
							}
						} else {
							if idx = util.Index(videoIdList, func(idx int) bool {
								return videoIdList[idx] == pd.VideoId
							}); idx < 0 {
								videoIdList = append(videoIdList, pd.VideoId)
							}
						}
					}
				}
			}
		}
		userInfoList = append(userInfoList, localUserList...)
	}

	photoMap := make(map[int64]*data_fs.PhotoProfile, len(photoIdList))
	videoMap := make(map[int64]*data_fs.PhotoProfile, len(videoIdList))
	if len(photoIdList) > 0 || len(videoIdList) > 0 {
		sfsIns := sfsService.GetSfsClient("0")
		if sfsIns == nil {
			log.Warnf("GetUserList sfsService = nil")
		} else {
			if len(photoIdList) > 0 {
				photoInfoList, err := sfsIns.ReqGetPhotoList(context.Background(), &sfsService.GetPhotoList{
					PhotoIdList: photoIdList,
				})
				if err != nil {
					log.Infof("GetUserList ReqGetPhotoList error:%s", err.Error())
				} else {
					for _, v := range photoInfoList.Values {
						detail := utils.ToDataPhoto(v).Detail
						if len(detail) > 0 {
							photoMap[v.VolumeId] = &data_fs.PhotoProfile{
								Photo: detail[0],
								Video: nil,
							}
						}
					}
				}
			}
			if len(videoIdList) > 0 {
				videoInfoList, err := sfsIns.ReqGetDocumentList(context.Background(), &sfsService.GetDocumentList{
					FileId: videoIdList,
				})
				if err != nil {
					log.Infof("GetUserList ReqGetDocumentList error:%s", err.Error())
				} else {
					for _, v := range videoInfoList.Values {
						videoMap[v.VolumeId] = &data_fs.PhotoProfile{
							Photo: utils.ToDataPhoto(v.Thumbs[0]).Detail[0],
							Video: utils.DocumentToVideo(v).Detail[0],
						}
					}
				}
			}
		}
	}

	var user *mtproto.User
	var dataPhoto *data_fs.PhotoProfile
	userList := make([]*mtproto.User, 0, len(userInfoList))
	for _, v := range userInfoList {
		vv, _ := contactsCacheMap[v.Id]
		photoV, _ := userPhotoMap[v.Id]
		if photoV != nil {
			if photoV.PhotoId != 0 {
				dataPhoto, _ = photoMap[photoV.PhotoId]
			} else if photoV.VideoId != 0 {
				dataPhoto, _ = videoMap[photoV.VideoId]
			}
		}
		user = UserCore2Users(v)
		if v.Id == userId {
			UserCoreSelfUsers(user)
		} else {
			user = UserCoreContactUser(user, vv != nil && !vv.Deleted, vv != nil && !vv.Deleted && vv.GetContact() > data_contact.MutualTypeMyContact, contactsCacheMap[v.Id])
		}
		userList = append(userList, user)
		user.Photo = photo.PhotoProfileUserProfilePhoto(dataPhoto, layer)
	}
	return userList, nil
}

func (c *Core) GetUser(userId int64, peerId int64, layer int32) (*mtproto.User, error) {

	var userInfo *data_users.Users
	var err error
	userInfo, err = c.UserData(userId)
	if err != nil {
		log.Warnf("GetUser GetUserData error:%s", err.Error())
		return nil, err
	}

	var videoId int64
	var photoId int64
	if profilePhotoData := users.MakeProfilePhotoData(userInfo.Photos); profilePhotoData.Default != 0 {
		if idx := util.Index(profilePhotoData.IdList, func(idx int) bool {
			return profilePhotoData.IdList[idx].PhotoId == profilePhotoData.Default
		}); idx >= 0 {
			if profilePhotoData.IdList[idx].VideoId == 0 {
				photoId = profilePhotoData.IdList[idx].PhotoId
			} else {
				videoId = profilePhotoData.IdList[idx].VideoId
			}
		}
	}

	var dataPhoto *data_fs.PhotoProfile
	if photoId != 0 || videoId != 0 {
		sfsIns := sfsService.GetSfsClient("0")
		if sfsIns == nil {
			log.Warnf("GetUserList sfsService = nil")
		} else {
			if photoId != 0 {
				photoInfo, err := sfsIns.ReqGetPhoto(context.Background(), &sfsService.GetPhoto{
					PhotoId: photoId,
					LocalId: 0,
				})
				if err != nil {
					log.Infof("GetUserList ReqGetPhoto error:%s", err.Error())
				} else {
					detail := utils.ToDataPhoto(photoInfo).Detail
					if len(detail) > 0 {
						dataPhoto = &data_fs.PhotoProfile{
							Photo: detail[0],
							Video: nil,
						}
					}
				}
			}
			if videoId != 0 {
				videoInfo, err := sfsIns.ReqGetDocument(context.Background(), &sfsService.GetDocument{
					FileId:  videoId,
					LocalId: 0,
				})
				if err != nil {
					log.Infof("GetUserList ReqGetDocument error:%s", err.Error())
				} else {
					dataPhoto = &data_fs.PhotoProfile{
						Photo: utils.ToDataPhoto(videoInfo.Thumbs[0]).Detail[0],
						Video: utils.DocumentToVideo(videoInfo).Detail[0],
					}
				}
			}
		}
	}

	var user *mtproto.User
	user = UserCore2Users(userInfo)
	if peerId == userId {
		UserCoreSelfUsers(user)
	} else {
		contact, err1 := c.contact.GetContactById(userId, peerId)
		if err1 != nil {
			log.Warnf("GetUser - GetContactById error:%s", err1.Error())
		}

		user = UserCoreContactUser(user, contact != nil && !contact.Deleted, contact != nil && !contact.Deleted && contact.GetContact() > data_contact.MutualTypeMyContact, contact)
	}
	user.Photo = photo.PhotoProfileUserProfilePhoto(dataPhoto, layer)
	return user, nil
}

func (c *Core) SearchUserData(query string, offset int32, limit int32) ([]*data_users.Users, error) {
	return c.usersCore.Search(query, offset, limit)
}
