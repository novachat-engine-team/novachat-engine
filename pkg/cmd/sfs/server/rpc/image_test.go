/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/9/13 10:15
 * @File : image_test.go
 */

package rpc

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/disintegration/imaging"
	"io/ioutil"
	"math/rand"
	"novachat_engine/pkg/cmd/sfs/server/partfs"
	"novachat_engine/service/data/fs"
	"os"
	"testing"
	"time"
)

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

func local(id int64, t partfs.SubPathsType) string {
	return fmt.Sprintf("./image/%d.jpg", id)
}

func FileHash(fileId int64, pathsType partfs.SubPathsType, offset int32) (int32, string, error) {
	fileName := local(fileId, pathsType)
	size, err := FileSize(fileName)
	if err != nil {
		//log.Errorf("FileHash fileName:%s FileSize error:%s", fileName, err.Error())
		panic(err)
		return 0, "", err
	}

	if offset >= int32(size) {
		offset = 0
	}

	m := sha256.New()

	f, err := os.Open(fileName)
	if err != nil {
		//log.Errorf("FileHash fileName:%s Open error:%s", fileName, err.Error())
		panic(err)
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

func TestImage(t *testing.T) {
	fileFileName := "E:\\test\\novachat_engine\\server\\sfs\\server\\rpc\\image\\1455136455218171904.jpg"
	img, err := imaging.Open(fileFileName)

	if err != nil {
		panic(err)
	}
	photo := &data_fs.Photo{}

	format, _, err := partfs.GetImageFormat(photo.Filename)
	if format == -1 {
		t.Fatal(err)
	}

	photo.Detail = make([]*data_fs.PhotoDetail, 0, len(partfs.PhotoSizeList))
	info := make([]*data_fs.Photo, 0, len(partfs.PhotoSizeList)+1)
	thumbFileName := partfs.LocationDir(photo.VolumeId, partfs.SubPathsType(photo.FileType))

	for idx, s := range partfs.PhotoSizeList {
		if (img.Bounds().Size().X < s.Size || img.Bounds().Size().Y < s.Size) && idx != 0 {
			break
		}

		if idx == 0 {
			thumbFileName = fileFileName
			photo.Detail = append(photo.Detail, &data_fs.PhotoDetail{
				VolumeId: photo.VolumeId,
				LocalId:  photo.LocalId,
				FilePath: photo.FilePath,
				PhotoSize: &data_fs.PhotoSize{
					Width:  int32(img.Bounds().Size().X),
					Height: int32(img.Bounds().Size().Y),
					Size_:  photo.Size_,
					Type:   s.Type,
				},
			})
			info = append(info, photo)
		} else {
			detail := &data_fs.PhotoDetail{
				VolumeId: 0,
				LocalId:  0,
				FilePath: "",
				PhotoSize: &data_fs.PhotoSize{
					Width:  int32(s.Size),
					Height: int32(s.Size),
					Size_:  0,
					Type:   s.Type,
				},
			}
			id := rand.Int63()
			thumbFileName = partfs.LocationDir(id, partfs.SubPathsType(photo.FileType))
			subPhoto := &data_fs.Photo{
				VolumeId: id,
				LocalId:  0,
				FilePath: thumbFileName,
				Filename: "",
				Md5Sum:   "",
				Date:     int32(time.Now().Unix()),
				Detail:   []*data_fs.PhotoDetail{detail},
				FileType: photo.FileType,
				Size_:    0,
			}

			dst := imaging.Fill(img, s.Size, s.Size, imaging.Center, imaging.Lanczos)
			err = saveImage(dst, thumbFileName, imaging.Format(format), imaging.JPEGQuality(100))
			if err != nil {
				continue
			}

			info = append(info, subPhoto)

			size, md5sum, _ := partfs.FileHash(id, partfs.SubPathsType(photo.FileType), 0)
			detail.VolumeId = id
			detail.PhotoSize.Size_ = size
			detail.FilePath = thumbFileName
			subPhoto.Size_ = size
			subPhoto.Md5Sum = md5sum
			photo.Detail = append(photo.Detail, detail)
		}
	}
	t.Logf("%+v", info)
}
