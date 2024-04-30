package auth

type Device struct {
	TokenType    int32   `protobuf:"varint,1,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty" bson:"token_type"`
	AppSandBox   bool    `protobuf:"varint,2,opt,name=app_sand_box,json=appSandBox,proto3" json:"app_sand_box,omitempty" bson:"app_sand_box"`
	Secret       string  `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty" bson:"secret"`
	NoMuted      bool    `protobuf:"varint,4,opt,name=no_muted,json=noMuted,proto3" json:"no_muted,omitempty" bson:"no_muted"`
	DeviceModel  string  `protobuf:"bytes,5,opt,name=device_model,json=deviceModel,proto3" json:"device_model,omitempty" bson:"device_model"`
	LangPack     string  `protobuf:"bytes,6,opt,name=lang_pack,json=langPack,proto3" json:"lang_pack,omitempty" bson:"lang_pack"`
	OtherUidList []int64 `protobuf:"varint,7,rep,packed,name=other_uid_list,json=otherUidList,proto3" json:"other_uid_list,omitempty" bson:"other_uid_list"`
}
