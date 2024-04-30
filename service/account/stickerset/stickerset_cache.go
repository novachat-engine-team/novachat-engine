package stickerset

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
	"novachat_engine/pkg/log"
	data_stickerset "novachat_engine/service/data/stickerset"
)

const (
	defaultFavedLimit   int32  = 20
	defaultRecentLimit  int32  = 50
	KeyAllStickerSet    string = "set:allSticker"
	KeyStickerSet       string = "string:stickerSet:"
	KeyStickerInstalled string = "string:stickerInstall:"
	KeyStickerFaved     string = "string:stickerFaved:"
	KeyStickerRecent    string = "string:stickerRecent:"
)

func makeAllStickerKey() string {
	return KeyAllStickerSet
}

func makeStickerSetKey(id int64) string {
	return fmt.Sprintf("%s%d", KeyStickerSet, id)
}

func makeStickerFavedKey(id int64) string {
	return fmt.Sprintf("%s%d", KeyStickerFaved, id)
}

func makeStickerRecentKey(id int64) string {
	return fmt.Sprintf("%s%d", KeyStickerRecent, id)
}

func makeStickerInstalledKey(id int64) string {
	return fmt.Sprintf("%s%d", KeyStickerInstalled, id)
}

func (s *Core) allStickersId2Cache(stickerSetList []*data_stickerset.StickerSet) error {
	stickerSetIdListValue := make([]interface{}, 0, len(stickerSetList)+1)
	stickerSetIdListValue = append(stickerSetIdListValue, makeAllStickerKey())
	for _, v := range stickerSetList {
		stickerSetIdListValue = append(stickerSetIdListValue, v.Id)
	}

	saddValue, err1 := redis.Int(s.redisClient.Do("SADD", stickerSetIdListValue...))
	if err1 != nil {
		log.Warnf("stickers2Cache SADD :%s", err1.Error())
	} else {
		log.Debugf("stickers2Cache SADD value:%d", saddValue)
	}
	return err1
}

func (s *Core) cache2AllStickersId() ([]int64, error) {
	stickerSetIdList, err := redis.Int64s(s.redisClient.Do("SMEMBERS", makeAllStickerKey()))
	if err != nil {
		log.Warnf("GetAllStickers SMEMBERS error:%s", err.Error())
	}
	return stickerSetIdList, nil
}

func (s *Core) stickers2Cache(stickerSetList []*data_stickerset.StickerSet) error {
	stickerSetIdListValue := make([]interface{}, 0, len(stickerSetList)*2)
	for _, v := range stickerSetList {
		stickerSetIdListValue = append(stickerSetIdListValue, makeStickerSetKey(v.Id))
		ss, _ := jsoniter.MarshalToString(v)
		stickerSetIdListValue = append(stickerSetIdListValue, ss)
	}
	msetValue, err1 := redis.String(s.redisClient.Do("MSET", stickerSetIdListValue...))
	if err1 != nil {
		log.Warnf("stickers2Cache MSET :%s", err1.Error())
	} else {
		log.Debugf("stickers2Cache MSET value:%v", msetValue)
	}
	return err1
}

func (s *Core) cache2stickers(idList []int64) ([]*data_stickerset.StickerSet, error) {
	stickerSetIdListValue := make([]interface{}, 0, len(idList))
	for _, v := range idList {
		stickerSetIdListValue = append(stickerSetIdListValue, makeStickerSetKey(v))
	}

	mgetList, err := redis.Strings(s.redisClient.Do("MGET", stickerSetIdListValue...))
	if err != nil {
		log.Warnf("cache2stickers MGET :%s", err.Error())
		return nil, err
	}
	stickerSetList := make([]*data_stickerset.StickerSet, 0, len(mgetList))
	for _, v := range mgetList {
		var ss data_stickerset.StickerSet
		err = jsoniter.UnmarshalFromString(v, &ss)
		if err == nil {
			stickerSetList = append(stickerSetList, &ss)
		}
	}
	return stickerSetList, nil
}

func (s *Core) cache2Installed(userId int64) ([]*data_stickerset.StickerInstall, error) {
	v, err := redis.String(s.redisClient.Do("GET", makeStickerInstalledKey(userId)))
	if err != nil {
		log.Warnf("cache2Installed GET :%s", err.Error())
		return nil, err
	}
	var l []*data_stickerset.StickerInstall
	err = jsoniter.UnmarshalFromString(v, &l)
	return l, nil
}

func (s *Core) installed2Cache(userId int64, l []*data_stickerset.StickerInstall) error {
	v, err := jsoniter.MarshalToString(l)
	rsp, err := redis.String(s.redisClient.Do("SET", makeStickerInstalledKey(userId), v))
	if err != nil {
		log.Warnf("installed2Cache SET :%s", err.Error())
		return err
	}
	_ = rsp
	return nil
}

func (s *Core) faved2Cache(userId int64, faved []int64) error {
	values := make([]interface{}, 0, len(faved))
	values = append(values, makeStickerFavedKey(userId))
	for _, v := range faved {
		values = append(values, v)
	}

	ok, err := redis.String(s.redisClient.Do("SET", values...))
	if err != nil {
		log.Warnf("faved2Cache SET :%s", err.Error())
	} else {
		log.Debugf("faved2Cache SET value:%d", ok)
	}
	return err
}

func (s *Core) recent2Cache(userId int64, recent []int64) error {
	values := make([]interface{}, 0, len(recent))
	values = append(values, makeStickerRecentKey(userId))
	for _, v := range recent {
		values = append(values, v)
	}

	ok, err := redis.String(s.redisClient.Do("SET", values...))
	if err != nil {
		log.Warnf("faved2Cache SET :%s", err.Error())
	} else {
		log.Debugf("faved2Cache SET value:%d", ok)
	}
	return err
}

func (s *Core) cache2Faved(userId int64) ([]int64, error) {
	ok, err := redis.Int64s(s.redisClient.Do("GET", makeStickerFavedKey(userId)))
	if err != nil {
		log.Warnf("cache2faved GET :%s", err.Error())
	} else {
		log.Debugf("cache2faved GET value:%v", ok)
	}
	return ok, err
}

func (s *Core) cache2Recent(userId int64) ([]int64, error) {
	ok, err := redis.Int64s(s.redisClient.Do("GET", makeStickerRecentKey(userId)))
	if err != nil {
		log.Warnf("cache2Recent GET :%s", err.Error())
	} else {
		log.Debugf("cache2Recent GET value:%v", ok)
	}
	return ok, err
}
