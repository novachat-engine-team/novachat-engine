/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/1 18:53
 * @File : partfs.go
 */

package partfs

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"novachat_engine/pkg/cmd/sfs/server/syncfs"
	"novachat_engine/pkg/log"
	"os"
	"strings"
)

func InstallPartFS(rootPath string) {
	ok, err := PathExits(rootPath)
	if err != nil {
		panic(err)
	}

	if ok == false && len(rootPath) > 0 {
		err = os.MkdirAll(FormatPath(rootPath), 0755)
		if err != nil {
			panic(err)
		}
		ok, err = PathExits(rootPath)
		if err != nil {
			panic(err)
		}
	}

	if ok == true && len(rootPath) > 0 {
		_rootPath = FormatPath(rootPath)
	}

	for idx := range _subPaths {
		for i := 0; i < DirNumber; i++ {
			p := fmt.Sprintf("%s/%02X", GetSubPath(SubPathsType(idx)), i)
			if ok, err = PathExits(p); err != nil {
				panic(p)
			}
			if ok == false {
				err = os.MkdirAll(p, 0755)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

type PartFS struct {
	fru    FRU
	syncFS syncfs.SyncFS
}

func NewPartFS(fru FRU, syncFS syncfs.SyncFS) *PartFS {
	return &PartFS{
		fru:    fru,
		syncFS: syncFS,
	}
}

func (m *PartFS) UnClearDir(pathName string) {
	m.fru.UnclearDir(pathName)
}

func (m *PartFS) UnClearFile(fileName string) {
	m.fru.UnclearFile(fileName)
}

func (m *PartFS) WriteFileData(authKeyId int64, fileId int64, filePart int32, bytes []byte) error {
	filePathDir := LocationPartDatDir(authKeyId, fileId)
	exist, err := PathExits(filePathDir)
	if err != nil {
		log.Errorf("WriteFileData fileId:%d filePathDir:%s error:%s", fileId, filePathDir, err.Error())
		return err
	}

	if exist == false {
		err = os.MkdirAll(filePathDir, 0755)
		if err != nil && strings.Contains(err.Error(), "file exists") == false {
			log.Errorf("WriteFileData Mkdir %s error:%s", filePathDir, err.Error())
			return err
		}
	}

	fileName := fmt.Sprintf("%s/%d.part", filePathDir, filePart)
	err = ioutil.WriteFile(fileName, bytes, 0644)
	if err != nil {
		log.Errorf("WriteFileData fileId:%d error:%s", fileId, err.Error())

		m.fru.ClearDir(filePathDir)
		return err
	}

	m.fru.PartDir(filePathDir)
	m.syncFS.SyncFile(fileName)
	return nil
}

func (m *PartFS) WriteFilePartData(authKeyId int64, fileId int64, filePart int32, fileTotalPart int32, bytes []byte) error {
	filePathDir := LocationPartDatDir(authKeyId, fileId)
	exist, err := PathExits(filePathDir)
	if err != nil {
		log.Errorf("WriteFilePartData fileId:%d filePathDir:%s error:%s", fileId, filePathDir, err.Error())
		return err
	}

	if exist == false {
		err = os.MkdirAll(filePathDir, 0755)
		if !errors.Is(err, os.ErrExist) {
			log.Errorf("WriteFilePartData Mkdir %s error:%s", filePathDir, err.Error())
			return err
		}
	}

	fileName := fmt.Sprintf("%s/%d.part", filePathDir, filePart)
	err = ioutil.WriteFile(fileName, bytes, 0644)
	if err != nil {
		log.Errorf("WriteFilePartData fileId:%d error:%s", fileId, err.Error())

		m.fru.ClearDir(filePathDir)
		return err
	}

	m.fru.PartDir(filePathDir)

	if fileTotalPart >= filePart+1 {
		m.syncFS.PartsComplete(filePathDir, fileTotalPart)
	}
	return nil
}

func (m *PartFS) CalcDatPartsSize(datPath string, parts int32) (int64, error) {
	if parts > 0 {
		size := int64(0)
		for idx := int32(0); idx < parts; idx++ {
			fileName := fmt.Sprintf("%s/%d.part", datPath, idx)
			v, err := FileSize(fileName)
			if err != nil {
				log.Errorf("CalcDatPartsSize fileName:%s error:%s", fileName, err.Error())
				return 0, err
			}
			size += v
		}
		return size, nil
	} else {
		fileName := fmt.Sprintf("%s/%d.part", datPath, parts)
		return FileSize(fileName)
	}
}

func (m *PartFS) CalcDatPartsHash(datPath string, parts int32) (int64, string, error) {
	exist, err := PathExits(datPath)
	if err != nil {
		log.Errorf("CalcDatPartsHash fileName:%s PathExits error:%s", datPath, err.Error())
		return 0, "", err
	}

	if exist == false {
		log.Errorf("CalcDatPartsHash fileName:%s PathExits", datPath)
		return 0, "", PathNotExistErr
	}

	if parts < 0 {
		log.Errorf("CalcDatPartsHash fileName:%s PathExits", datPath)
		return 0, "", FileNotExistErr
	}

	var size int64
	if parts == 0 {
		fileName := fmt.Sprintf("%s/%d.part", datPath, parts)
		if size, err = FileSize(fileName); err != nil {
			log.Errorf("CalcDatPartsHash fileName:%s FileSize", datPath)
			return 0, "", err
		}
		m := sha256.New()
		b, err := ioutil.ReadFile(fileName)
		if err != nil {
			err = fmt.Errorf("path not exists: %s", fileName)
			return 0, "", err
		}
		m.Write(b)
		return size, hex.EncodeToString(m.Sum(nil)), nil
	} else {
		m := sha256.New()
		for idx := int32(0); idx < parts; idx++ {
			var s int64
			fileName := fmt.Sprintf("%s/%d.part", datPath, idx)
			if s, err = FileSize(fileName); err != nil {
				log.Errorf("CalcDatPartsHash fileName:%s FileSize", fileName)
				return 0, "", err
			}
			b, err := ioutil.ReadFile(fileName)
			if err != nil {
				log.Errorf("CalcDatPartsHash fileName:%s FileSize", fileName)
				err = fmt.Errorf("path not exists: %s", fileName)
				return 0, "", err
			}
			m.Write(b)
			size += s
		}
		return size, hex.EncodeToString(m.Sum(nil)), nil
	}
}

func (m *PartFS) MergeDatFileParts(datPath string, parts int32, dstPathName string) error {
	if parts > 0 {
		file, err := os.OpenFile(dstPathName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
		if err != nil {
			log.Errorf("CalcDatPartsSize dstPathName:%s error:%s", dstPathName, err.Error())
			return err
		}
		defer file.Close()

		for idx := int32(0); idx < parts; idx++ {
			fileName := fmt.Sprintf("%s/%d.part", datPath, idx)
			b, err := ioutil.ReadFile(fileName)
			if err != nil {
				log.Errorf("CalcDatPartsSize read file :%s error:%s", fileName, err.Error())
				return err
			}

			_, err = file.Write(b)
			if err != nil {
				log.Errorf("CalcDatPartsSize append fileName:%s error:%s", fileName, err.Error())
				return err
			}
		}
		return nil
	} else {
		fileName := fmt.Sprintf("%s/%d.part", datPath, parts)
		return os.Rename(fileName, dstPathName)
	}
}

func (m *PartFS) MergeDatPhotoFileParts(datPath string, parts int32, dstPathName string) error {
	if parts > 0 {
		file, err := os.OpenFile(dstPathName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
		if err != nil {
			log.Errorf("CalcDatPartsSize dstPathName:%s error:%s", dstPathName, err.Error())
			return err
		}
		defer file.Close()

		for idx := int32(0); idx < parts; idx++ {
			fileName := fmt.Sprintf("%s/%d.part", datPath, idx)
			b, err := ioutil.ReadFile(fileName)
			if err != nil {
				log.Errorf("CalcDatPartsSize read file :%s error:%s", fileName, err.Error())
				return err
			}

			_, err = file.Write(b)
			if err != nil {
				log.Errorf("CalcDatPartsSize append fileName:%s error:%s", fileName, err.Error())
				return err
			}
		}
	} else {
		fileName := fmt.Sprintf("%s/%d.part", datPath, parts)
		if err := os.Rename(fileName, dstPathName); err != nil {
			return err
		}
	}

	return nil
}
