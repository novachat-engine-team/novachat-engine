/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/1 18:57
 * @File : fs_util.go
 */

package partfs

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	"io/ioutil"
	"novachat_engine/pkg/log"
	"os"
	"path"
	"strings"
)

var PathNotExistErr = errors.New("path not exists")
var FileNotExistErr = errors.New("file not exists")

var _rootPath = "./sfs_dir"
var DirNumber = 0xFF

const (
	KPhotoSizeSmallNormal = "0"
	KPhotoSizeSmallType   = "s"
	KPhotoSizeMediumType  = "m"
	KPhotoSizeXLargeType  = "x"
	KPhotoSizeYLargeType  = "y"

	KPhotoSizeWLargeType = "w"
	KPhotoSizeALargeType = "a"
	KPhotoSizeBLargeType = "b"
	KPhotoSizeCLargeType = "c"
	KPhotoSizeDLargeType = "d"

	KPhotoSizePathSizeType = "j" // webp thumb
)

// s: box 100x100
// m: box 320x320
// x: box 800x800
// y: box 1280x1280
// w: box 2560x2560 // if loading this fix HistoryPhoto::updateFrom
// a: crop 160x160
// b: crop 320x320
// c: crop 640x640
// d: crop 1280x1280

const (
	KPhotoSizeOriginalSize = 0
	KPhotoSizeSmallSize    = 100
	KPhotoSizeMediumSize   = 320
	KPhotoSizeXLargeSize   = 800
	KPhotoSizeYLargeSize   = 1280
)

type PhotoSizeType struct {
	Type string
	Size int
}

var PhotoSizeList = []PhotoSizeType{
	{KPhotoSizeWLargeType, KPhotoSizeOriginalSize},
	{KPhotoSizeSmallType, KPhotoSizeSmallSize},
	//{KPhotoSizeMediumType,KPhotoSizeMediumSize},
	//{KPhotoSizeXLargeType,KPhotoSizeXLargeSize},
	//{KPhotoSizeYLargeType,KPhotoSizeYLargeSize},
}

// storage_FileType <--
//  + TL_storage_fileUnknown
//  + TL_storage_filePartial
//  + TL_storage_fileJpeg
//  + TL_storage_fileGif
//  + TL_storage_filePng
//  + TL_storage_filePdf
//  + TL_storage_fileMp3
//  + TL_storage_fileMov
//  + TL_storage_fileMp4
//  + TL_storage_fileWebp
var _subPaths = []string{"dat", "part", "png", "jpeg", "gif", "pdf", "mp3", "mov", "mp4", "webp", "sticker_set"}

type SubPathsType int32

func (m SubPathsType) ToInt32() int32 {
	return int32(m)
}

const (
	SubPathsTypeDat        SubPathsType = 0
	SubPathsTypePart       SubPathsType = 1
	SubPathsTypePng        SubPathsType = 2
	SubPathsTypeJpeg       SubPathsType = 3
	SubPathsTypeFGif       SubPathsType = 4
	SubPathsTypePdf        SubPathsType = 5
	SubPathsTypeMov        SubPathsType = 6
	SubPathsTypeMp3        SubPathsType = 7
	SubPathsTypeMp4        SubPathsType = 8
	SubPathsTypeWebp       SubPathsType = 9
	SubPathsTypeStickerSet SubPathsType = 10
)

var _rootSubPaths = make([]string, len(_subPaths))

func ConfigRootSubPath(path string, pathsType SubPathsType) error {
	ok, err := PathExits(path)
	if err != nil {
		panic(err)
	}
	if ok == false && len(path) > 0 {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			log.Errorf("ConfigRootSubPath path:%s error:%s", path, err.Error())
			return err
		}

		ok, err = PathExits(path)
		if err != nil {
			panic(err)
		}
	}
	if ok == true && len(path) > 0 {
		path = FormatPath(path)
		_rootSubPaths[pathsType] = path
		return nil
	}

	return fmt.Errorf("path error:%s", path)
}

func getFileExtName(filePath string) string {
	var ext = path.Ext(filePath)
	return strings.ToLower(ext)
}

func GetImageFormat(fileName string) (int, string, error) {
	formats := map[string]imaging.Format{
		".jpg":  imaging.JPEG,
		".jpeg": imaging.JPEG,
		".png":  imaging.PNG,
		".tif":  imaging.TIFF,
		".tiff": imaging.TIFF,
		".bmp":  imaging.BMP,
		".gif":  imaging.GIF,
		".webp": imaging.GIF,
	}

	ext := getFileExtName(fileName)
	f, ok := formats[ext]
	if !ok {
		return -1, "", imaging.ErrUnsupportedFormat
	}

	return int(f), ext, nil
}

func GetSubPath(s SubPathsType) string {
	rootPath := _rootSubPaths[s]
	if len(rootPath) == 0 {
		rootPath = _rootPath
	}
	return fmt.Sprintf("%s/%s", rootPath, _subPaths[s])
}

func FormatPath(path string) string {
	path = strings.Replace(path, "\\", "/", -1)
	if strings.HasSuffix(path, "/") == true {
		return path[:len(path)-1]
	}
	return path
}

func PathExits(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if fileInfo != nil {
		return true, nil
	}

	if os.IsNotExist(err) == true {
		return false, nil
	}

	return false, err
}

func FileSize(filePathName string) (int64, error) {
	fileInfo, err := os.Stat(filePathName)
	if err != nil {
		return 0, err
	}

	if fileInfo.IsDir() {
		return 0, fmt.Errorf("dir error:%s", filePathName)
	}

	return fileInfo.Size(), nil
}

func LocationPartDatDir(authKeyId int64, fileId int64) string {
	return fmt.Sprintf("%s/%02X/%d/%d.parts", GetSubPath(SubPathsTypePart), uint64(fileId)%uint64(DirNumber), authKeyId, fileId)
}

func LocationPartDir(fileId int64, subPathType SubPathsType) string {
	return fmt.Sprintf("%s/%02X/%d.parts", GetSubPath(subPathType), uint64(fileId)%uint64(DirNumber), fileId)
}

func LocationDir(fileId int64, subPathType SubPathsType) string {
	return fmt.Sprintf("%s/%02X/%d.dat", GetSubPath(subPathType), uint64(fileId)%uint64(DirNumber), fileId)
}

func LocationHashDir(hash string, subPathType SubPathsType) string {
	fileId := "00"
	if len(hash) >= 2 {
		fileId = strings.ToUpper(hash[:2])
	}
	return fmt.Sprintf("%s/%s/%s.dat", GetSubPath(subPathType), fileId, hash)
}

func PartsHash(fileId int64, partsTotal int32) (int64, string, error) {
	path := LocationPartDir(fileId, SubPathsTypeDat)
	exist, err := PathExits(path)
	if err != nil {
		log.Errorf("PartsHash fileName:%s PathExits error:%s", path, err.Error())
		return 0, "", err
	}

	if exist == false {
		log.Errorf("PartsHash fileName:%s PathExits", path)
		return 0, "", PathNotExistErr
	}

	if partsTotal < 0 {
		log.Errorf("PartsHash fileName:%s PathExits", path)
		return 0, "", FileNotExistErr
	}

	var size int64
	if partsTotal == 0 {
		fileName := fmt.Sprintf("%s/%d.part", path, partsTotal)
		if size, err = FileSize(fileName); err != nil {
			log.Errorf("PartsHash fileName:%s FileSize", path)
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
		for idx := int32(0); idx < partsTotal; idx++ {
			var s int64
			fileName := fmt.Sprintf("%s/%d.part", path, idx)
			if s, err = FileSize(fileName); err != nil {
				log.Errorf("PartsHash fileName:%s FileSize", fileName)
				return 0, "", err
			}
			b, err := ioutil.ReadFile(fileName)
			if err != nil {
				log.Errorf("PartsHash fileName:%s FileSize", fileName)
				err = fmt.Errorf("path not exists: %s", fileName)
				return 0, "", err
			}
			m.Write(b)
			size += s
		}
		return size, hex.EncodeToString(m.Sum(nil)), nil
	}
}

func FileHash(fileId int64, pathsType SubPathsType, offset int32) (int32, string, error) {
	fileName := LocationDir(fileId, pathsType)
	size, err := FileSize(fileName)
	if err != nil {
		log.Errorf("FileHash fileName:%s FileSize error:%s", fileName, err.Error())
		return 0, "", err
	}

	if offset >= int32(size) {
		offset = 0
	}

	m := sha256.New()

	f, err := os.Open(fileName)
	if err != nil {
		log.Errorf("FileHash fileName:%s Open error:%s", fileName, err.Error())
		return 0, "", err
	}
	defer f.Close()

	if offset >= 0 {
		f.Seek(int64(offset), 0)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		err = fmt.Errorf("path not exists: %s", fileName)
		return 0, "", err
	}
	m.Write(b)
	return int32(size) - offset, hex.EncodeToString(m.Sum(nil)), nil
}

func HashFileHash(hash string, pathsType SubPathsType, offset int32) (int32, string, error) {
	fileName := LocationHashDir(hash, pathsType)
	size, err := FileSize(fileName)
	if err != nil {
		log.Errorf("HashFileHash fileName:%s FileSize error:%s", fileName, err.Error())
		return 0, "", err
	}

	if offset >= int32(size) {
		offset = 0
	}

	m := sha256.New()

	f, err := os.Open(fileName)
	if err != nil {
		log.Errorf("HashFileHash fileName:%s Open error:%s", fileName, err.Error())
		return 0, "", err
	}
	defer f.Close()

	if offset >= 0 {
		f.Seek(int64(offset), 0)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		err = fmt.Errorf("path not exists: %s", fileName)
		return 0, "", err
	}
	m.Write(b)
	return int32(size) - offset, hex.EncodeToString(m.Sum(nil)), nil
}

func GetFileData(filePathName string, offset int32, limit int32) ([]byte, error) {
	f, err := os.Open(filePathName)
	if err != nil {
		log.Errorf("GetFileData fileName:%s Open error:%s", filePathName, err.Error())
		return nil, err
	}
	defer f.Close()

	if true {
		fileSize, _ := FileSize(filePathName)
		if int64(offset) > fileSize {
			limit = 0
		} else if int64(offset+limit) > fileSize {
			limit = int32(fileSize - int64(offset))
		}
		bytes := make([]byte, limit)
		n, err := f.ReadAt(bytes, int64(offset))
		if err != nil {
			log.Errorf("read file %s error:%s ", filePathName, err)
			return nil, err
		}
		log.Debugf("DEBUG filePathName:%s offset:%d limit:%d n:%d  len()=%d", filePathName, offset, limit, n, len(bytes))
		return bytes, nil
	} else {
		if offset >= 0 {
			f.Seek(int64(offset), 0)
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			err = fmt.Errorf("path not exists: %s", filePathName)
			log.Errorf("GetFileData fileName:%s ReadAll error:%s", filePathName, err.Error())
			return nil, err
		}

		if limit > 0 {
			if int32(len(b)) > limit {
				return b[:limit], nil
			} else {
				return b, nil
			}
		}
		return b, nil
	}
}
