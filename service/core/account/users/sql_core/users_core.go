/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File :
 * @Desc :
 *
 */

package sql_core

import (
	"database/sql"
	"gorm.io/gorm"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/crypto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/orm"
	"novachat_engine/pkg/util"
	"novachat_engine/service/core/account/users"
	"novachat_engine/service/data/users"
	"strings"
)

type Core struct {
	db *gorm.DB
}

func (c *Core) UserByPhoneNumber(phoneNumber string) (*data_users.Users, error) {
	u := &data_users.Users{}
	if err := c.db.Table(users.TableName).Find(u, "phone = ?", phoneNumber).Error; err != nil {
		log.Errorf("UserByPhoneNumber phoneNumber:%s error:%s", phoneNumber, err.Error())
		return nil, err
	}
	return u, nil
}

func (c *Core) UserByPhoneNumberList(phoneList []string) ([]*data_users.Users, error) {
	if len(phoneList) == 0 {
		return nil, nil
	}

	stringList := make([]sql.NullString, 0, len(phoneList))
	for _, v := range phoneList {
		stringList = append(stringList, sql.NullString{Valid: true, String: v})
	}
	var u []*data_users.Users
	if err := c.db.Table(users.TableName).Find(u, "phone in (?)", stringList).Error; err != nil {
		log.Errorf("UserByPhoneNumberList phoneList:%s error:%s", phoneList, err.Error())
		return nil, err
	}
	return u, nil
}

func (c *Core) UserCreate(phoneNumber string, firstname string, lastname string) (*data_users.Users, error) {
	u := &data_users.Users{
		Phone:      sql.NullString{Valid: true, String: phoneNumber},
		AccessHash: crypto.GenerateAccessHash(),
		FirstName:  firstname,
		LastName:   lastname,
	}
	db := c.db.Table(users.TableName).Create(u)
	if db.Error != nil {
		log.Errorf("UserCreate phoneNumber:%s error:%s", phoneNumber, db.Error.Error())
		return nil, db.Error
	}

	//u.Id = db.RowsAffected
	log.Infof("UserCreate ok value:%v", u)
	return u, nil
}

func (c *Core) UpdateBindPhone(userId int64, number string) (bool, error) {
	db := c.db.Table(users.TableName).Where("id = ?", userId).Updates(map[string]interface{}{
		"phone_number": sql.NullString{Valid: len(number) > 0, String: number},
	})
	if db.Error != nil {
		log.Errorf("UpdateBindPhone userId:%d number:%s error:%s", userId, number, db.Error.Error())
		return false, db.Error
	}
	return db.RowsAffected > 0, nil
}

func (c *Core) DeleteAccount(userId int64, reason string) (bool, error) {
	db := c.db.Table(users.TableName).Where("id = ?", userId).Updates(map[string]interface{}{
		"deleted":            1,
		"restricted":         1,
		"restriction_reason": reason,
	})
	if db.Error != nil {
		log.Errorf("DeleteAccount userId:%d reason:%s error:%s", userId, reason, db.Error.Error())
		return false, db.Error
	}
	return db.RowsAffected > 0, nil
}

func (c *Core) User(id int64) (*data_users.Users, error) {
	u := &data_users.Users{}
	if err := c.db.Table(users.TableName).Find(u, "id = ?", id).Error; err != nil {
		log.Errorf("User id:%d error:%s", id, err.Error())
		return nil, err
	}
	return u, nil
}

func (c *Core) Username(username string) (*data_users.Users, error) {
	u := &data_users.Users{}
	if err := c.db.Table(users.TableName).Find(u, "username = ?", sql.NullString{Valid: true, String: username}).Error; err != nil {
		log.Errorf("Username Username:%s error:%s", username, err.Error())
		return nil, err
	}
	return u, nil
}

func (c *Core) UsernameList(usernameList []string) ([]*data_users.Users, error) {
	if len(usernameList) == 0 {
		return nil, nil
	}

	stringList := make([]string, 0, len(usernameList))
	for _, v := range usernameList {
		stringList = append(stringList, v)
	}
	var u []*data_users.Users
	if err := c.db.Table(users.TableName).Where("phone in ?", stringList).Find(&u).Error; err != nil {
		log.Errorf("UsernameList usernameList:%s error:%s", usernameList, err.Error())
		return nil, err
	}
	return u, nil
}

func (c *Core) UpdateUsername(userId int64, username string) (bool, error) {
	db := c.db.Table(users.TableName).Where("id = ?", userId).Update("username", username)
	if db.Error != nil {
		log.Errorf("UpdateUsername userId:%d username:%s error:%s", userId, username, db.Error.Error())
		if strings.Contains(db.Error.Error(), strings.ToLower("Duplicate")) {
			return false, nil
		}

		return false, db.Error
	}
	return db.RowsAffected > 0, nil
}

func (c *Core) UserIdList(list []int64) ([]*data_users.Users, error) {
	var u []*data_users.Users
	if err := c.db.Table(users.TableName).Where("id in (?)", list).Find(&u).Error; err != nil {
		log.Errorf("UserIdList id:%v error:%s", list, err.Error())
		return nil, err
	}
	return u, nil
}

func (c *Core) UpdateAccountTTL(userId int64, days int32) (bool, error) {
	db := c.db.Table(users.TableName).Where("id = ?", userId).Update("days_ttl", days)
	if db.Error != nil {
		log.Errorf("UpdateAccountTTL userId:%d days:%d error:%s", userId, days, db.Error.Error())
		return false, db.Error
	}
	return db.RowsAffected > 0, nil
}

func (c *Core) UpdateProfile(userId int64, firstname string, lastname string, about string) (bool, error) {
	db := c.db.Table(users.TableName).Where("id = ?", userId).Updates(map[string]interface{}{
		"first_name": firstname,
		"last_name":  lastname,
		"about":      about,
	})
	if db.Error != nil {
		log.Errorf("UpdateProfile userId:%d fname:`%s` lname:`%s` about:`%s` error:%s", userId, firstname, lastname, about, db.Error.Error())
		return false, db.Error
	}
	return db.RowsAffected > 0, nil
}

func (c *Core) UserPhotos(userId int64) ([]*data_users.ProfilePhotoData, error) {
	u := &data_users.Users{}
	db := c.db.Table(users.TableName).Where("id = ?", userId).Find(u)
	if db.Error != nil {
		log.Errorf("UserPhotos userId:%d error:%s", userId, db.Error.Error())
		return nil, db.Error
	}

	v := users.MakeProfilePhotoData(u.Photos)
	return v.IdList, nil
}

func (c *Core) DeletePhoto(userId int64, photoList []int64) (*data_users.ProfilePhotoData, error) {

	var id *data_users.ProfilePhotoData
	err := c.db.Transaction(func(tx *gorm.DB) error {
		u := &data_users.Users{}
		if err := tx.Table(users.TableName).Where("id = ?", userId).Find(&u).Error; err != nil {
			log.Errorf("DeletePhoto userId:%d error:%s", userId, err.Error())
			return nil
		}
		profileUserPhotoId := users.MakeProfilePhotoData(u.Photos)

		found := false
		l := make([]*data_users.ProfilePhotoData, 0, len(profileUserPhotoId.IdList))
		for _, vv := range profileUserPhotoId.IdList {
			found = false
			for _, v := range photoList {
				if vv.PhotoId == v {
					found = true
					break
				}
			}
			if !found {
				l = append(l, vv)
			}
		}
		if len(l) > 0 {
			found = false
			for _, v := range l {
				if v.PhotoId == profileUserPhotoId.Default {
					found = true
					break
				}
			}
			profileUserPhotoId.IdList = l
			if !found {
				profileUserPhotoId.Default = l[0].PhotoId
				id = l[0]
			}
		} else {
			profileUserPhotoId.Default = 0
		}

		c.db.Table(users.TableName).Where("id = ?", userId).Update("photos", users.ToProfilePhotoIDsJson(profileUserPhotoId))
		return nil
	})

	if err != nil {
		log.Errorf("DeletePhoto userId:%d error:%s", userId, err.Error())
		return id, err
	}
	return id, nil
}

func (c *Core) SetPhotoId(userId int64, photoId int64, videoId int64) (*data_users.ProfilePhotoData, error) {
	dataProfile := data_users.ProfilePhotoData{
		PhotoId: photoId,
		VideoId: videoId,
	}
	err := c.db.Transaction(func(tx *gorm.DB) error {
		u := &data_users.Users{}
		if err := tx.Table(users.TableName).Where("id = ?", userId).Find(&u).Error; err != nil {
			log.Errorf("DeletePhoto userId:%d error:%s", userId, err.Error())
			return nil
		}
		profileUserPhotoId := users.MakeProfilePhotoData(u.Photos)
		if photoId != 0 {
			if idx := util.Index(profileUserPhotoId.IdList, func(idx int) bool {
				return profileUserPhotoId.IdList[idx].PhotoId == photoId
			}); idx < 0 {
				profileUserPhotoId.IdList = append(profileUserPhotoId.IdList, &dataProfile)
			} else {
				dataProfile = *profileUserPhotoId.IdList[idx]
			}
		} else {
			return nil
		}
		if profileUserPhotoId.Default == photoId {
			return nil
		}
		profileUserPhotoId.Default = photoId
		c.db.Table(users.TableName).Where("id = ?", userId).Update("photos", users.ToProfilePhotoIDsJson(profileUserPhotoId))
		return nil
	})

	if err != nil {
		log.Errorf("DeletePhoto userId:%d error:%s", userId, err.Error())
		return nil, err
	}
	return &dataProfile, nil
}

func (c *Core) Search(query string, offset int32, limit int32) ([]*data_users.Users, error) {
	var u []*data_users.Users
	if err := c.db.Table(users.TableName).Find(&u, "id > ? and username like ?", offset, query+"%").Limit(int(limit)).Error; err != nil {
		log.Errorf("Search query:%s error:%s", query, err.Error())
		return nil, err
	}
	return u, nil
}

func NewUsersCore(mysql *config.MySQLConfig) *Core {
	db, err := orm.NewORM(mysql)
	if err != nil {
		panic(err)
	}

	if mysql.Migrate {
		db.AutoMigrate(data_users.Users{})
	}

	return &Core{
		db: db,
	}
}
