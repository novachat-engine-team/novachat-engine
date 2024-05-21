package data_users

type ProfilePhotoData struct {
	PhotoId int64 `protobuf:"varint,1,opt,name=p,proto3" json:"p,omitempty"`
	VideoId int64 `protobuf:"varint,2,opt,name=v,proto3" json:"v,omitempty"`
}

type ProfilePhotoIds struct {
	Default int64               `protobuf:"varint,1,opt,name=d,proto3" json:"d,omitempty"`
	IdList  []*ProfilePhotoData `protobuf:"varint,2,rep,packed,name=l,json=idList,proto3" json:"l,omitempty"`
}
