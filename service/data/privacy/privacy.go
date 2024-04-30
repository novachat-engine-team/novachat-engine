/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File :
 * @Desc :
 *
 */

package data_privacy

type PrivacyOption struct {
	PrivacyKey   int32   `protobuf:"varint,1,opt,name=privacy_key,json=privacyKey,proto3" json:"privacy_key,omitempty"`
	AllowList    []int64 `protobuf:"varint,2,rep,packed,name=allow_list,json=allowList,proto3" json:"allow_list,omitempty"`
	DisallowList []int64 `protobuf:"varint,3,rep,packed,name=disallow_list,json=disallowList,proto3" json:"disallow_list,omitempty"`
}

type Privacy struct {
	UserId      int64            `protobuf:"varint,1,opt,name=user_id,json=user_id,proto3" json:"user_id,omitempty" bson:"_id,omitempty"`
	Global      []*PrivacyOption `protobuf:"varint,2,opt,name=global,json=global,proto3" json:"global,omitempty" bson:"global,omitempty"`
	UserStatus  []*PrivacyOption `protobuf:"varint,3,opt,name=user_status,json=user_status,proto3" json:"user_status,omitempty" bson:"user_status,omitempty"`
	ChatInvite  []*PrivacyOption `protobuf:"varint,4,opt,name=chat_invite,json=chat_invite,proto3" json:"chat_invite,omitempty" bson:"chat_invite,omitempty"`
	PhoneCall   []*PrivacyOption `protobuf:"varint,5,opt,name=phone_call,json=phone_call,proto3" json:"phone_call,omitempty" bson:"phone_call,omitempty"`
	PhoneP2P    []*PrivacyOption `protobuf:"varint,6,opt,name=phone_p2p,json=phone_p2p,proto3" json:"phone_p2p,omitempty" bson:"phone_p2p,omitempty"`
	Forwards    []*PrivacyOption `protobuf:"varint,7,opt,name=forwards,json=forwards,proto3" json:"forwards,omitempty" bson:"forwards,omitempty"`
	Profile     []*PrivacyOption `protobuf:"varint,8,opt,name=profile,json=profile,proto3" json:"profile,omitempty" bson:"profile,omitempty"`
	PhoneNumber []*PrivacyOption `protobuf:"varint,9,opt,name=phone_number,json=phone_number,proto3" json:"phone_number,omitempty" bson:"phone_number,omitempty"`
	AddByPhone  []*PrivacyOption `protobuf:"varint,10,opt,name=add_by_phone,json=add_by_phone,proto3" json:"add_by_phone,omitempty" bson:"add_by_phone,omitempty"`
}
