package upload

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	sfsClient "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/cmd/sfs/server/partfs"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/log"
	"novachat_engine/service/app/dc"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/upload/attribute"
	"novachat_engine/service/upload/photo"
)

type Core struct {
}

func NewUploadCore(cnf *config.EtcdClientConfig) *Core {
	sfsClient.InstallSfsClient(*cnf)
	return &Core{}
}

func (c *Core) GetFile(file *sfsClient.GetFile) (*mtproto.Upload_File, error) {
	log.Debugf("UploadCore GetFile file:%v", file)
	fileInfo, err := sfsClient.GetSfsClient(fmt.Sprintf("%d", file.VolumeId)).ReqGetFile(context.Background(), file)
	if err != nil {
		log.Errorf("UploadCore GetFile file:%v error:%s", file, err.Error())
		return nil, err
	}
	if len(fileInfo.Bytes) == 0 {
		log.Warnf("UploadCore GetFile file:%v file is empty", file)
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_UPLOAD)
	}

	//  upload.file#96a18d5 type:storage.FileType mtime:int bytes:bytes = upload.File;
	ret := mtproto.NewTLUploadFile(nil).To_Upload_File()
	switch partfs.SubPathsType(fileInfo.FileType) {
	case partfs.SubPathsTypeDat:
		ret.Type = mtproto.NewTLStorageFilePartial(nil).To_Storage_FileType()
	case partfs.SubPathsTypePart:
		ret.Type = mtproto.NewTLStorageFilePartial(nil).To_Storage_FileType()
	case partfs.SubPathsTypePng:
		ret.Type = mtproto.NewTLStorageFilePng(nil).To_Storage_FileType()
	case partfs.SubPathsTypeJpeg:
		ret.Type = mtproto.NewTLStorageFileJpeg(nil).To_Storage_FileType()
	case partfs.SubPathsTypeFGif:
		ret.Type = mtproto.NewTLStorageFileGif(nil).To_Storage_FileType()
	case partfs.SubPathsTypePdf:
		ret.Type = mtproto.NewTLStorageFilePdf(nil).To_Storage_FileType()
	case partfs.SubPathsTypeMov:
		ret.Type = mtproto.NewTLStorageFileMov(nil).To_Storage_FileType()
	case partfs.SubPathsTypeMp4:
		ret.Type = mtproto.NewTLStorageFileMp4(nil).To_Storage_FileType()
	case partfs.SubPathsTypeWebp:
		ret.Type = mtproto.NewTLStorageFileWebp(nil).To_Storage_FileType()
	case partfs.SubPathsTypeStickerSet:
		ret.Type = mtproto.NewTLStorageFilePartial(nil).To_Storage_FileType()
	default:
		ret.Type = mtproto.NewTLStorageFileUnknown(nil).To_Storage_FileType()
	}

	ret.Bytes = fileInfo.Bytes
	ret.Mtime = fileInfo.Mtime
	return ret, nil
}

func (c *Core) SavePartFileData(req *sfsClient.UploadSavePartFileData) (*mtproto.Bool, error) {
	log.Debugf("SavePartFileData req:%v", req.FileId)

	ret, err := sfsClient.GetSfsClient(fmt.Sprintf("%d", req.AuthKeyId)).
		ReqUploadSavePartFileData(context.Background(), req)
	if err != nil {
		log.Errorf("SavePartFileData req:%v error:%s", req.FileId, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_UPLOAD)
	}

	var ok mtproto.Bool
	types.UnmarshalAny(ret, &ok)
	return &ok, nil
}

func (c *Core) UploadSaveMultiPartFileData(req *sfsClient.UploadSaveMultiPartFileData) (*mtproto.Bool, error) {
	log.Debugf("UploadSaveMultiPartFileData req:%v", req.FileId)

	ret, err := sfsClient.GetSfsClient(fmt.Sprintf("%d", req.AuthKeyId)).
		ReqUploadSaveMultiPartFileData(context.Background(), req)
	if err != nil {
		log.Errorf("UploadSaveMultiPartFileData req:%v error:%s", req.FileId, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_UPLOAD)
	}

	var ok mtproto.Bool
	types.UnmarshalAny(ret, &ok)
	return &ok, nil
}

func (c *Core) UploadDocument(uploadDocument *sfsClient.UploadedDocument, layer int32) (*mtproto.Document, error) {
	log.Debugf("UploadDocument req:%v", uploadDocument.FileId)

	documentInfo, err := sfsClient.GetSfsClient(fmt.Sprintf("%d", uploadDocument.FileId)).ReqUploadedDocument(
		context.Background(),
		uploadDocument)
	if err != nil {
		log.Errorf("UploadDocument req:%v error:%s", uploadDocument.FileId, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_UPLOAD)
	}

	attributes := make([]*mtproto.DocumentAttribute, 0, len(documentInfo.Attributes))
	for _, v := range documentInfo.Attributes {
		attributes = append(attributes, attribute.Attributes2MtDocumentAttribute(v))
	}

	thumbs := make([]*mtproto.PhotoSize, 0, len(documentInfo.Thumbs))
	for _, v := range documentInfo.Thumbs {
		thumbs = append(thumbs, photo.PhotoInfo2PhotoSize(v, layer))
		if strippedSize := photo.PhotoInfo2PhotoStrippedSize(v); strippedSize != nil {
			thumbs = append(thumbs, strippedSize)
		}
	}
	var thumb *mtproto.PhotoSize
	if len(thumbs) > 0 {
		thumb = thumbs[0]
	}

	// TODO:(VideoThumbs)
	document := mtproto.NewTLDocument(&mtproto.Document{
		Id:            documentInfo.VolumeId,
		AccessHash:    documentInfo.AccessHash,
		Date:          documentInfo.Date,
		MimeType:      documentInfo.MimeType,
		Size_:         documentInfo.Size_,
		Thumb:         thumb,
		DcId:          dc.DefaultDc,
		Version:       0,
		Attributes:    attributes,
		FileReference: nil,
		Thumbs:        thumbs,
		VideoThumbs:   nil,
	}).To_Document()
	return document, nil
}

func (c *Core) UploadPhoto(uploadPhoto *sfsClient.UploadPhoto, layer int32) (*mtproto.Photo, error) {
	log.Debugf("UploadPhoto req:%v", uploadPhoto.FileId)

	photoInfo, err := sfsClient.GetSfsClient(fmt.Sprintf("%d", uploadPhoto.FileId)).ReqUploadPhoto(
		context.Background(),
		uploadPhoto)
	if err != nil {
		log.Errorf("UploadPhoto req:%v error:%s", uploadPhoto.FileId, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_UPLOAD)
	}

	return photo.PhotoInfo2Photo(photoInfo, layer), nil
}

func (c *Core) GetDocument(getDocument *sfsClient.GetDocument, layer int32) (*mtproto.Document, error) {
	log.Debugf("GetDocument req:%v", getDocument)

	documentInfo, err := sfsClient.GetSfsClient(fmt.Sprintf("%d", getDocument.FileId)).ReqGetDocument(context.Background(),
		getDocument)
	if err != nil {
		log.Errorf("GetDocument req:%v error:%s", getDocument.FileId, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_UPLOAD)
	}

	attributes := make([]*mtproto.DocumentAttribute, 0, len(documentInfo.Attributes))
	for _, v := range documentInfo.Attributes {
		attributes = append(attributes, attribute.Attributes2MtDocumentAttribute(v))
	}
	thumbs := make([]*mtproto.PhotoSize, 0, len(documentInfo.Thumbs))
	for _, v := range documentInfo.Thumbs {
		thumbs = append(thumbs, photo.PhotoInfo2PhotoSize(v, layer))
		if strippedSize := photo.PhotoInfo2PhotoStrippedSize(v); strippedSize != nil {
			thumbs = append(thumbs, strippedSize)
		}
	}
	var thumb *mtproto.PhotoSize
	if len(thumbs) > 0 {
		thumb = thumbs[0]
	}

	// TODO:(VideoThumbs)
	document := mtproto.NewTLDocument(&mtproto.Document{
		Id:            documentInfo.VolumeId,
		AccessHash:    documentInfo.AccessHash,
		Date:          documentInfo.Date,
		MimeType:      documentInfo.MimeType,
		Size_:         documentInfo.Size_,
		Thumb:         thumb,
		DcId:          dc.DefaultDc,
		Version:       0,
		Attributes:    attributes,
		FileReference: nil,
		Thumbs:        thumbs,
		VideoThumbs:   nil,
	}).To_Document()
	return document, nil
}

func (c *Core) GetPhoto(getPhoto *sfsClient.GetPhoto, layer int32) (*mtproto.Photo, error) {
	log.Debugf("GetPhoto req:%v", getPhoto)

	photoInfo, err := sfsClient.GetSfsClient(fmt.Sprintf("%d", getPhoto.PhotoId)).ReqGetPhoto(context.Background(),
		getPhoto)
	if err != nil {
		log.Errorf("GetPhoto req:%v error:%s", getPhoto.PhotoId, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_UPLOAD)
	}

	return photo.PhotoInfo2Photo(photoInfo, layer), nil
}
