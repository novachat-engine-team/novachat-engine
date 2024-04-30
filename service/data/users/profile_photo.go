package data_users

type ProfilePhotoIds struct {
	Default int64   `protobuf:"varint,1,opt,name=d,proto3" json:"d,omitempty"`
	IdList  []int64 `protobuf:"varint,2,rep,packed,name=l,json=idList,proto3" json:"l,omitempty"`
}
