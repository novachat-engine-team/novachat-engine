package data_contact

type MutualType int32

const (
	MutualTypeDefault   MutualType = 0
	MutualTypeMyContact MutualType = 1
	MutualTypeMutual    MutualType = 2
)

type Contact struct {
	UserId    int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" bson:"user_id"`
	PeerId    int64  `protobuf:"varint,2,opt,name=peer_id,json=peer_id,proto3" json:"peer_id,omitempty" bson:"peer_id"`
	Contact   int32  `protobuf:"varint,3,opt,name=contact,proto3" json:"contact,omitempty" bson:"contact"`
	ClientId  int64  `protobuf:"varint,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" bson:"client_id"`
	FirstName string `protobuf:"bytes,5,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty" bson:"first_name"`
	LastName  string `protobuf:"bytes,6,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty" bson:"last_name"`
	NickName  string `protobuf:"bytes,7,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty" bson:"nick_name"`
	FoldId    int32  `protobuf:"varint,8,opt,name=fold_id,json=foldId,proto3" json:"fold_id,omitempty" bson:"fold_id"`
	Date      int32  `protobuf:"varint,9,opt,name=date,proto3" json:"date,omitempty" bson:"date"`
	Block     bool   `protobuf:"varint,10,opt,name=block,proto3" json:"block,omitempty" bson:"block"`
	Deleted   bool   `protobuf:"varint,11,opt,name=deleted,proto3" json:"deleted,omitempty" bson:"deleted"`
	Id        string `protobuf:"bytes,12,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	Phone     string `protobuf:"bytes,13,opt,name=phone,proto3" json:"phone,omitempty" bson:"phone"`
}

func (c *Contact) GetContact() MutualType {
	return MutualType(c.Contact)
}
