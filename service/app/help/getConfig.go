/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/15 11:52
 * @File : getConfig.go
 */

package help

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	"os"
	"strings"
	"sync"
)

const (
	ExpiresTimeout = 3600
)

type DcOption struct {
	Ipv6                 bool     `protobuf:"varint,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	MediaOnly            bool     `protobuf:"varint,4,opt,name=media_only,json=mediaOnly,proto3" json:"media_only,omitempty"`
	TcpoOnly             bool     `protobuf:"varint,5,opt,name=tcpo_only,json=tcpoOnly,proto3" json:"tcpo_only,omitempty"`
	Cdn                  bool     `protobuf:"varint,6,opt,name=cdn,proto3" json:"cdn,omitempty"`
	Static               bool     `protobuf:"varint,7,opt,name=static,proto3" json:"static,omitempty"`
	Id                   int32    `protobuf:"varint,8,opt,name=id,proto3" json:"id,omitempty"`
	IpAddress            string   `protobuf:"bytes,9,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Port                 int32    `protobuf:"varint,10,opt,name=port,proto3" json:"port,omitempty"`
	Secret               []byte   `protobuf:"bytes,11,opt,name=secret,proto3" json:"secret,omitempty"`
	Platform			 string   `protobuf:"bytes,11,opt,name=platform,proto3" json:"platform,omitempty"`
}

var (
	locker sync.RWMutex
	testMode = false
	dcOptions []*DcOption
	feature []*mtproto.DisabledFeature
	conf mtproto.Config
)

var fileNames = []string {
	"config.json",
	"dc_options.json",
	"test_mode.json",
	"disabled_feature.json",
}

const (
	FileIndexConfig int32 = 0
	FileIndexDcOptions int32 = 1
	FileIndexTestMode int32 = 2
	FileIndexDisabledFeature int32 = 3
)

func GetFileNameList() []string {
	return fileNames
}

func _parseConfig(fileName string, obj interface{}) error {
	configFile, err := os.Open(fmt.Sprintf("%s", fileName))
	if err != nil {
		log.Errorf("not found %s: error:%s", fileName, err.Error())
		return err
	}

	defer configFile.Close()
	fileData, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Errorf("readAll %s: error:%s", fileName, err.Error())
		return err
	}
	err = json.Unmarshal(fileData, obj)
	if err != nil {
		log.Errorf("json.Unmarshal %s: error:%s", fileName, err.Error())
		return err
	}

	return nil
}

func ParseConfig(pathName string) error {
	var err error
	 _parseConfig(fmt.Sprintf("%s/%s", pathName, fileNames[FileIndexConfig]), &conf)

	if dcOptions == nil {
		dcOptions = make([]*DcOption, 0, 0)
	}

	_parseConfig(fmt.Sprintf("%s/%s", pathName, fileNames[FileIndexDcOptions]), &dcOptions)

	type _test_mode struct {
		TestMode bool `json:"test_mode""`
	}

	test_mode := _test_mode{}
	if err = _parseConfig(fmt.Sprintf("%s/%s", pathName, fileNames[FileIndexTestMode]), &test_mode); err == nil {
		testMode = test_mode.TestMode
	}

	if feature == nil {
		feature = make([]*mtproto.DisabledFeature, 0, 0)
	}
	_parseConfig(fmt.Sprintf("%s/%s", pathName, fileNames[FileIndexDisabledFeature]), &feature)

	locker.Lock()
	defer locker.Unlock()

	conf.ClassName = mtproto.ClassConfig

	conf.TestMode = mtproto.ToMTBool(testMode)

	//conf.DcOptions = []*mtproto.DcOption{}
	//if len(dcOptions) == 0 {
	//	err = fmt.Errorf("ParseConfig error: not found dcOptions")
	//	return err
	//} else {
	//	for idx := range dcOptions {
	//		dcOptions[idx].Cmd = mtproto.Cmd_DcOption
	//		dcOptions[idx].ClassName = mtproto.ClassDcOption
	//
	//		conf.DcOptions = append(conf.DcOptions, dcOptions[idx])
	//	}
	//}

	if len(feature) > 0 {
		conf.DisabledFeatures = make([]*mtproto.DisabledFeature, 0, len(feature))
		for idx := range feature {
			feature[idx].Cmd = mtproto.Cmd_DisabledFeature
			feature[idx].ClassName = mtproto.ClassDisabledFeature

			conf.DisabledFeatures = append(conf.DisabledFeatures, feature[idx])
		}
	}
	return nil
}

func GetConfig() *mtproto.Config {
	locker.RLock()
	defer locker.RUnlock()

	return mtproto.NewTLConfig(&conf).To_Config()
}

func GetConfigWithDC(platform constants.DevicePlatformType) *mtproto.Config {
	locker.RLock()
	defer locker.RUnlock()

	c := conf
	c.DcOptions = make([]*mtproto.DcOption, 0, len(dcOptions))
	for _, v := range dcOptions {
		dc := mtproto.NewTLDcOption(&mtproto.DcOption{
			Ipv6:                 v.Ipv6,
			MediaOnly:            v.MediaOnly,
			TcpoOnly:             v.TcpoOnly,
			Cdn:                  v.Cdn,
			Static:               v.Static,
			Id:                   v.Id,
			IpAddress:            v.IpAddress,
			Port:                 v.Port,
			Secret:               v.Secret,
		}).To_DcOption()
		if platform == constants.DevicePlatformTypeIOS {
			if v.Platform == platform.ToString() || len(v.Platform) == 0 {
				c.DcOptions = append(c.DcOptions, dc)
			}
		} else if platform == constants.DevicePlatformTypeAndroid {
			if v.Platform == platform.ToString() || len(v.Platform) == 0 {
				c.DcOptions = append(c.DcOptions, dc)
			}
		} else {
			c.DcOptions = append(c.DcOptions, dc)
		}
	}
	return mtproto.NewTLConfig(&c).To_Config()
}

func GetConfigWithDCReviews(platform constants.DevicePlatformType, review bool) *mtproto.Config {
	locker.RLock()
	defer locker.RUnlock()

	c := conf
	c.DcOptions = make([]*mtproto.DcOption, 0, len(dcOptions))
	for _, v := range dcOptions {
		dc := mtproto.NewTLDcOption(&mtproto.DcOption{
			Ipv6:                 v.Ipv6,
			MediaOnly:            v.MediaOnly,
			TcpoOnly:             v.TcpoOnly,
			Cdn:                  v.Cdn,
			Static:               v.Static,
			Id:                   v.Id,
			IpAddress:            v.IpAddress,
			Port:                 v.Port,
			Secret:               v.Secret,
		}).To_DcOption()

		if review == false {
			if platform == constants.DevicePlatformTypeIOS {
				if v.Platform == platform.ToString() || len(v.Platform) == 0 {
					c.DcOptions = append(c.DcOptions, dc)
				}
			} else if platform == constants.DevicePlatformTypeAndroid {
				if v.Platform == platform.ToString() || len(v.Platform) == 0 {
					c.DcOptions = append(c.DcOptions, dc)
				}
			} else {
				c.DcOptions = append(c.DcOptions, dc)
			}
		} else {
			if platform == constants.DevicePlatformTypeIOS {
				if strings.Compare(v.Platform, platform.ToString() + constants.DCOptionAppStoreReviews) == 0 {
					c.DcOptions = append(c.DcOptions, dc)
				}
			} else if platform == constants.DevicePlatformTypeAndroid {
				if strings.Compare(v.Platform, platform.ToString() + constants.DCOptionAppStoreReviews) == 0 {
					c.DcOptions = append(c.DcOptions, dc)
				}
			}
		}
	}
	return mtproto.NewTLConfig(&c).To_Config()
}