/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/11/1 18:26
 * @File : make_thumbs.go
 */

package rpc

import (
	"github.com/bwmarrin/snowflake"
	"github.com/disintegration/imaging"
	"image"
	"io"
	"novachat_engine/pkg/cmd/sfs/server/partfs"
	"novachat_engine/pkg/crypto"
	"novachat_engine/pkg/log"
	data_fs "novachat_engine/service/data/fs"
	"os"
	"time"
)

func makeThumbs(photoFileName string, photo *data_fs.Photo, node *snowflake.Node) ([]*data_fs.Photo, error) {
	img, err := imaging.Open(photoFileName)
	if err != nil {
		log.Errorf("makeThumbs fileName:%s error:%s", photoFileName, err.Error())
		return nil, err
	}

	format, _, err1 := partfs.GetImageFormat(photo.Filename)
	if format == -1 {
		log.Warnf("makeThumbs fileName:%s error:%s", photo.Filename, err1.Error())
		return nil, err
	}

	photo.Detail = make([]*data_fs.PhotoDetail, 0, len(partfs.PhotoSizeList))
	info := make([]*data_fs.Photo, 0, len(partfs.PhotoSizeList)+1)
	thumbFileName := partfs.LocationDir(photo.VolumeId, partfs.SubPathsType(photo.FileType))

	for idx, s := range partfs.PhotoSizeList {
		if (img.Bounds().Size().X < s.Size || img.Bounds().Size().Y < s.Size) && idx != 0 {
			break
		}

		if idx == 0 {
			thumbFileName = photoFileName
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
			id := node.Generate().Int64()
			thumbFileName = partfs.LocationDir(id, partfs.SubPathsType(photo.FileType))
			subPhoto := &data_fs.Photo{
				VolumeId:   id,
				LocalId:    0,
				FilePath:   thumbFileName,
				Filename:   "",
				Md5Sum:     "",
				Date:       int32(time.Now().Unix()),
				Detail:     []*data_fs.PhotoDetail{detail},
				FileType:   photo.FileType,
				Size_:      0,
				AccessHash: crypto.GenerateAccessHash(),
			}

			dst := imaging.Fill(img, s.Size, s.Size, imaging.Center, imaging.Lanczos)
			err = saveImage(dst, thumbFileName, imaging.Format(format), imaging.JPEGQuality(100))
			if err != nil {
				log.Warnf("makeThumbs thumbFileName:%s error:%s", thumbFileName, err.Error())
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
	return info, nil
}

func saveImage(img image.Image, filename string, f imaging.Format, opts ...imaging.EncodeOption) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return imaging.Encode(file, img, f, opts...)
}

func fileMove(fileName string, descFileName string) error {
	return os.Rename(fileName, descFileName)
}

// fileCopy copies file from source to target path.
func fileCopy(src, dest string) error {
	// Gather file information to set back later.
	si, err := os.Lstat(src)
	if err != nil {
		return err
	}

	// Handle symbolic link.
	if si.Mode()&os.ModeSymlink != 0 {
		target, err := os.Readlink(src)
		if err != nil {
			return err
		}
		// NOTE: os.Chmod and os.Chtimes don't recoganize symbolic link,
		// which will lead "no such file or directory" error.
		return os.Symlink(target, dest)
	}

	sr, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sr.Close()

	dw, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer dw.Close()

	if _, err = io.Copy(dw, sr); err != nil {
		return err
	}

	// Set back file information.
	if err = os.Chtimes(dest, si.ModTime(), si.ModTime()); err != nil {
		return err
	}
	return os.Chmod(dest, si.Mode())
}
