/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/5 15:03
 * @File : multifile_util.go
 */

package partfs

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func SaveAnimatedFile(pathsType SubPathsType, filePath string, frameNumber int32, bytes []byte) error {
	path := GetSubPath(pathsType)
	exist, err := PathExits(filePath)
	if err != nil {
		return err
	}

	if exist == false {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	fileName := fmt.Sprintf("%s/%d.dat", path, frameNumber)
	return ioutil.WriteFile(fileName, bytes, 0644)
}

func GetAnimatedFile(pathsType SubPathsType, filePath string, frameNumber int32) ([]byte, error) {
	path := GetSubPath(pathsType)
	exist, err := PathExits(filePath)
	if err != nil {
		return nil, err
	}

	if exist == false {
		return nil, PathNotExistErr
	}

	fileName := fmt.Sprintf("%s/%d.dat", path, frameNumber)
	return ioutil.ReadFile(fileName)
}

func GetAnimatedFileHash(pathsType SubPathsType, filePath string, offset int32, frameCount int32) (int32, string, error) {
	path := GetSubPath(pathsType)
	exist, err := PathExits(filePath)
	if err != nil {
		return 0, "", err
	}

	if offset < 0 {
		offset = 0
	}

	if exist == false {
		return 0, "", PathNotExistErr
	}

	size := int64(0)
	if frameCount == 0 {
		fileName := fmt.Sprintf("%s/%d.dat", path, frameCount)
		if size, err = FileSize(fileName); err != nil {
			return 0, "", err
		}
		m := sha256.New()

		f, err := os.Open(fileName)
		if err != nil {
			err = fmt.Errorf("path not exists: %s", fileName)
			return 0, "", err
		}
		defer f.Close()
		if offset > 0 {
			f.Seek(int64(offset), 0)
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			err = fmt.Errorf("path not exists: %s", fileName)
			return 0, "", err
		}
		m.Write(b)
		return int32(size), hex.EncodeToString(m.Sum(nil)), nil
	} else {
		m := sha256.New()
		sizeCount := int64(0)
		for idx:=int32(offset);idx<frameCount;idx++ {
			fileName := fmt.Sprintf("%s/%d.dat", path, idx)
			if size, err = FileSize(fileName); err != nil {
				return 0, "", err
			}
			b, err := ioutil.ReadFile(fileName)
			if err != nil {
				err = fmt.Errorf("path not exists: %s", fileName)
				return 0, "", err
			}
			sizeCount += size
			m.Write(b)
		}
		return int32(sizeCount), hex.EncodeToString(m.Sum(nil)), nil
	}
}

func LocationStickerSetDir(filePath string) string {
	hash := sha256.New()
	hashStr := hex.EncodeToString(hash.Sum([]byte(filePath)))
	return fmt.Sprintf("%s/%02X/%s", GetSubPath(SubPathsTypeStickerSet), hashStr[:2], hashStr)
}

func StickerSetSave(filePath string, frameNumber int32, bytes []byte) error {
	return SaveAnimatedFile(SubPathsTypeStickerSet, filePath, frameNumber, bytes)
}

func StickerSetGet(filePath string, frameNumber int32) ([]byte, error) {
	return GetAnimatedFile(SubPathsTypeStickerSet, filePath, frameNumber)
}

func StickerSetHash(filePath string, offset int32, frameCount int32) (int32, string, error) {
	return GetAnimatedFileHash(SubPathsTypeStickerSet, filePath, offset, frameCount)
}