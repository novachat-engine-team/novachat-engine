package rpc

import (
	"context"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/cmd/sfs/utils"
	"novachat_engine/pkg/log"
	data_sfs "novachat_engine/service/data/fs"
)

func (m *SfsServiceImpl) ReqUploadPhoto(ctx context.Context, req *sfsService.UploadPhoto) (*sfsService.PhotoInfo, error) {
	log.Debugf("ReqUploadPhoto req:%v", req)
	id := m.node.Generate().Int64()
	photo, err := m.uploadPhoto(id, req)
	if err != nil {
		log.Errorf("ReqUploadPhoto uploadPhoto req:%v error:%s", req, err.Error())
		return nil, err
	}

	return utils.ToPhoto(photo), nil
}

func datPhotoHandler(datPath, dstPathName string, parts int32, info *data_sfs.Photo, m *SfsServiceImpl) ([]*data_sfs.Photo, error) {
	err := m.partCacheFS.MergeDatPhotoFileParts(datPath, parts, dstPathName) // move file
	if err != nil {
		m.partCacheFS.UnClearDir(datPath)
		log.Errorf("datPhotoHandler MergeDatFileParts datPath:%s dstPathName:%s info:%v error:%s", datPath, dstPathName, info, err.Error())
		return nil, err
	}

	photoSizes, err := makeThumbs(dstPathName, info, m.node)
	if err != nil {
		log.Warnf("datPhotoHandler makeThumbs dstPathName:%s info:%v error:%s", dstPathName, info, err.Error())
	}

	log.Infof("datPhotoHandler datHandler datPath:%s dstPathName:%s info:%v completed!!!!!!!!!!!!!!!!!", datPath, dstPathName, info)
	return photoSizes, err
}
