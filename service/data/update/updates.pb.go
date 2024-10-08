// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: updates.proto

package data_update

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	message "novachat_engine/service/data/messages/message"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type UserUpdate struct {
	UpdateKey        string           `protobuf:"bytes,21,opt,name=update_key,json=updateKey,proto3" json:"update_key,omitempty" bson:"_id"`
	Id               int32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"id"`
	UserId           int64            `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" bson:"user_id"`
	PeerId           int64            `protobuf:"varint,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty" bson:"peer_id"`
	PeerType         int32            `protobuf:"varint,4,opt,name=peer_type,json=peerType,proto3" json:"peer_type,omitempty" bson:"peer_type"`
	Pts              int32            `protobuf:"varint,5,opt,name=pts,proto3" json:"pts,omitempty" bson:"pts"`
	PtsCount         int32            `protobuf:"varint,6,opt,name=pts_count,json=ptsCount,proto3" json:"pts_count,omitempty" bson:"pts_count"`
	UpdateType       int32            `protobuf:"varint,7,opt,name=update_type,json=updateType,proto3" json:"update_type,omitempty" bson:"update_type"`
	Date             int32            `protobuf:"varint,8,opt,name=date,proto3" json:"date,omitempty" bson:"date"`
	Version          int32            `protobuf:"varint,9,opt,name=version,proto3" json:"version,omitempty" bson:"version"`
	Time             string           `protobuf:"bytes,10,opt,name=time,proto3" json:"time,omitempty" bson:"time"`
	RandomId         int64            `protobuf:"varint,11,opt,name=random_id,json=randomId,proto3" json:"random_id,omitempty" bson:"random_id"`
	MessageBoxIds    []int32          `protobuf:"varint,12,rep,packed,name=message_box_ids,json=messageBoxIds,proto3" json:"message_box_ids,omitempty" bson:"message_box_ids"`
	MessageAction    string           `protobuf:"bytes,13,opt,name=message_action,json=messageAction,proto3" json:"message_action,omitempty" bson:"message_action"`
	ChatId           int64            `protobuf:"varint,14,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty" bson:"chat_id"`
	MaxId            int32            `protobuf:"varint,15,opt,name=max_id,json=maxId,proto3" json:"max_id,omitempty" bson:"max_id"`
	MessageData      *message.Message `protobuf:"bytes,16,opt,name=message_data,json=messageData,proto3" json:"message_data,omitempty" bson:"message_data"`
	FoldId           int32            `protobuf:"varint,17,opt,name=fold_id,json=foldId,proto3" json:"fold_id,omitempty" bson:"fold_id"`
	StillUnreadCount int32            `protobuf:"varint,18,opt,name=still_unread_count,json=stillUnreadCount,proto3" json:"still_unread_count,omitempty" bson:"still_unread_count"`
	Media            *message.Media   `protobuf:"bytes,19,opt,name=media,proto3" json:"media,omitempty" bson:"media"`
	OwnerUserId      int64            `protobuf:"varint,20,opt,name=owner_user_id,json=ownerUserId,proto3" json:"owner_user_id,omitempty" bson:"owner_user_id"`
}

func (m *UserUpdate) Reset()         { *m = UserUpdate{} }
func (m *UserUpdate) String() string { return proto.CompactTextString(m) }
func (*UserUpdate) ProtoMessage()    {}
func (*UserUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_675fc0bf03cd96fd, []int{0}
}
func (m *UserUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserUpdate.Merge(m, src)
}
func (m *UserUpdate) XXX_Size() int {
	return m.Size()
}
func (m *UserUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_UserUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_UserUpdate proto.InternalMessageInfo

func (m *UserUpdate) GetUpdateKey() string {
	if m != nil {
		return m.UpdateKey
	}
	return ""
}

func (m *UserUpdate) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserUpdate) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserUpdate) GetPeerId() int64 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *UserUpdate) GetPeerType() int32 {
	if m != nil {
		return m.PeerType
	}
	return 0
}

func (m *UserUpdate) GetPts() int32 {
	if m != nil {
		return m.Pts
	}
	return 0
}

func (m *UserUpdate) GetPtsCount() int32 {
	if m != nil {
		return m.PtsCount
	}
	return 0
}

func (m *UserUpdate) GetUpdateType() int32 {
	if m != nil {
		return m.UpdateType
	}
	return 0
}

func (m *UserUpdate) GetDate() int32 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *UserUpdate) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *UserUpdate) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *UserUpdate) GetRandomId() int64 {
	if m != nil {
		return m.RandomId
	}
	return 0
}

func (m *UserUpdate) GetMessageBoxIds() []int32 {
	if m != nil {
		return m.MessageBoxIds
	}
	return nil
}

func (m *UserUpdate) GetMessageAction() string {
	if m != nil {
		return m.MessageAction
	}
	return ""
}

func (m *UserUpdate) GetChatId() int64 {
	if m != nil {
		return m.ChatId
	}
	return 0
}

func (m *UserUpdate) GetMaxId() int32 {
	if m != nil {
		return m.MaxId
	}
	return 0
}

func (m *UserUpdate) GetMessageData() *message.Message {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *UserUpdate) GetFoldId() int32 {
	if m != nil {
		return m.FoldId
	}
	return 0
}

func (m *UserUpdate) GetStillUnreadCount() int32 {
	if m != nil {
		return m.StillUnreadCount
	}
	return 0
}

func (m *UserUpdate) GetMedia() *message.Media {
	if m != nil {
		return m.Media
	}
	return nil
}

func (m *UserUpdate) GetOwnerUserId() int64 {
	if m != nil {
		return m.OwnerUserId
	}
	return 0
}

func init() {
	proto.RegisterType((*UserUpdate)(nil), "data_update.UserUpdate")
}

func init() { proto.RegisterFile("updates.proto", fileDescriptor_675fc0bf03cd96fd) }

var fileDescriptor_675fc0bf03cd96fd = []byte{
	// 681 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x54, 0xcd, 0x6e, 0x13, 0x3b,
	0x14, 0xee, 0x34, 0x4d, 0x9a, 0x78, 0x92, 0x34, 0x75, 0x7f, 0xae, 0x5b, 0xa9, 0x71, 0xe4, 0xbb,
	0x89, 0x74, 0x6f, 0x5b, 0xdd, 0xcb, 0x02, 0xa9, 0xb0, 0x80, 0x81, 0x4d, 0x54, 0xb1, 0xc0, 0x22,
	0x1b, 0x36, 0x23, 0x27, 0x76, 0xd3, 0x11, 0x4d, 0x26, 0x8a, 0x1d, 0x68, 0xde, 0x82, 0x25, 0x5b,
	0xfa, 0x34, 0x2c, 0x2b, 0xb1, 0x1f, 0x41, 0x9f, 0x00, 0xcd, 0x13, 0x20, 0x1f, 0x7b, 0x68, 0x9a,
	0xae, 0xe6, 0x9c, 0xf3, 0x7d, 0xe7, 0x77, 0xce, 0x31, 0x6a, 0xcc, 0xa7, 0x52, 0x18, 0xa5, 0x4f,
	0xa6, 0xb3, 0xd4, 0xa4, 0x38, 0x94, 0xc2, 0x88, 0xd8, 0xd9, 0x0e, 0x8f, 0x47, 0x89, 0xb9, 0x9c,
	0x0f, 0x4e, 0x86, 0xe9, 0xf8, 0x74, 0x94, 0x8e, 0xd2, 0x53, 0xe0, 0x0c, 0xe6, 0x17, 0xa0, 0x81,
	0x02, 0x92, 0xf3, 0x3d, 0x0c, 0xc7, 0x4a, 0x26, 0xc2, 0x2b, 0x8d, 0xb1, 0xd2, 0x5a, 0x8c, 0x94,
	0x53, 0xd9, 0xf7, 0x2a, 0x42, 0x7d, 0xad, 0x66, 0x7d, 0x88, 0x8c, 0x8f, 0x11, 0x72, 0x39, 0xe2,
	0x0f, 0x6a, 0x41, 0xf6, 0x3a, 0x41, 0xb7, 0x16, 0x35, 0xf3, 0x8c, 0xa2, 0x81, 0x4e, 0x27, 0x67,
	0x2c, 0x4e, 0x24, 0xe3, 0x35, 0xc7, 0x38, 0x57, 0x0b, 0x7c, 0x84, 0xd6, 0x13, 0x49, 0x82, 0x4e,
	0xd0, 0x2d, 0x47, 0x8d, 0x3c, 0xa3, 0x35, 0x47, 0xb3, 0xac, 0xf5, 0x44, 0xe2, 0x7f, 0xd0, 0xe6,
	0x5c, 0xab, 0x59, 0x9c, 0x48, 0xb2, 0xde, 0x09, 0xba, 0xa5, 0x08, 0xe7, 0x19, 0x6d, 0x3a, 0x8e,
	0x07, 0x18, 0xaf, 0x58, 0xa9, 0x07, 0xe4, 0xa9, 0x72, 0xe4, 0xd2, 0x2a, 0xd9, 0x03, 0x8c, 0x57,
	0xac, 0xd4, 0x93, 0xf8, 0x3f, 0x54, 0x03, 0x9b, 0x59, 0x4c, 0x15, 0xd9, 0x80, 0xfc, 0xbb, 0x79,
	0x46, 0x5b, 0x4b, 0x74, 0x0b, 0x31, 0x5e, 0xb5, 0xf2, 0xbb, 0xc5, 0x54, 0xe1, 0x0e, 0x2a, 0x4d,
	0x8d, 0x26, 0x65, 0x20, 0x2f, 0xf5, 0x34, 0x35, 0x9a, 0x71, 0x0b, 0x41, 0x50, 0xa3, 0xe3, 0x61,
	0x3a, 0x9f, 0x18, 0x52, 0x79, 0x14, 0xb4, 0x80, 0x6c, 0x50, 0xa3, 0x5f, 0x59, 0x11, 0x3f, 0x45,
	0xa1, 0x9f, 0x17, 0x54, 0xb2, 0x09, 0x4e, 0xfb, 0x79, 0x46, 0xb1, 0xef, 0xf2, 0x1e, 0x64, 0xdc,
	0x8f, 0x16, 0xaa, 0xf9, 0x1b, 0x6d, 0x58, 0x99, 0x54, 0xc1, 0x63, 0x2b, 0xcf, 0x68, 0xe8, 0x3c,
	0xac, 0x95, 0x71, 0x00, 0xf1, 0xbf, 0x68, 0xf3, 0xa3, 0x9a, 0xe9, 0x24, 0x9d, 0x90, 0x1a, 0xf0,
	0x96, 0x46, 0xe2, 0x01, 0xc6, 0x0b, 0x8a, 0x0d, 0x69, 0x92, 0xb1, 0x22, 0x08, 0xfe, 0xda, 0x52,
	0x48, 0x6b, 0x65, 0x1c, 0x40, 0xdb, 0xe3, 0x4c, 0x4c, 0x64, 0x3a, 0xb6, 0x73, 0x0e, 0x61, 0xce,
	0x4b, 0x3d, 0xfe, 0x81, 0x18, 0xaf, 0x3a, 0xb9, 0x27, 0x71, 0x84, 0xb6, 0xfc, 0xce, 0xc4, 0x83,
	0xf4, 0x3a, 0x4e, 0xa4, 0x26, 0xf5, 0x4e, 0xa9, 0x5b, 0x8e, 0x0e, 0xf3, 0x8c, 0xee, 0x3b, 0xc7,
	0x15, 0x02, 0xe3, 0xc5, 0x9a, 0x45, 0xe9, 0x75, 0x4f, 0x6a, 0xfc, 0x02, 0x35, 0x0b, 0x8a, 0x18,
	0x1a, 0xdb, 0x50, 0x03, 0xaa, 0x3c, 0xc8, 0x33, 0xba, 0xf7, 0x30, 0x84, 0xc3, 0xef, 0x23, 0xbc,
	0x04, 0xdd, 0xae, 0xc7, 0xf0, 0x52, 0x18, 0x5b, 0x76, 0x73, 0x75, 0x3d, 0x3c, 0xc0, 0x78, 0xc5,
	0x4a, 0x3d, 0x89, 0xbb, 0xa8, 0x32, 0x16, 0xb6, 0x12, 0xb2, 0x05, 0x73, 0xdb, 0xce, 0x33, 0xda,
	0xf0, 0x69, 0xc0, 0xce, 0x78, 0x79, 0x2c, 0xae, 0x7b, 0x12, 0xbf, 0x45, 0xf5, 0x22, 0xb1, 0xbd,
	0x30, 0xd2, 0xea, 0x04, 0xdd, 0xf0, 0xff, 0xbd, 0x13, 0x38, 0xb7, 0xe2, 0x54, 0xde, 0xf8, 0x5e,
	0xfe, 0xca, 0x33, 0xba, 0xf3, 0xb0, 0x5a, 0xcb, 0x63, 0x3c, 0xf4, 0xea, 0x6b, 0x61, 0x84, 0xad,
	0xf4, 0x22, 0xbd, 0x92, 0x36, 0xfb, 0xf6, 0xea, 0x5f, 0xf3, 0x00, 0xe3, 0x15, 0x2b, 0xf5, 0x24,
	0x3e, 0x47, 0x58, 0x9b, 0xe4, 0xea, 0x2a, 0x9e, 0x4f, 0x66, 0x4a, 0x48, 0xbf, 0x7c, 0x18, 0xfc,
	0x8e, 0xf2, 0x8c, 0x1e, 0x38, 0xbf, 0xc7, 0x1c, 0xc6, 0x5b, 0x60, 0xec, 0x83, 0xcd, 0x6d, 0xe3,
	0x33, 0x54, 0x86, 0x53, 0x27, 0x3b, 0xd0, 0xc5, 0xce, 0x6a, 0x17, 0x32, 0x11, 0x51, 0x2b, 0xcf,
	0x68, 0xbd, 0xe8, 0x41, 0x26, 0xc2, 0x4e, 0xc2, 0x7e, 0xf1, 0x73, 0xd4, 0x48, 0x3f, 0x4d, 0xd4,
	0x2c, 0x2e, 0x4e, 0x76, 0x17, 0xc6, 0x4c, 0xf2, 0x8c, 0xee, 0x3a, 0xfe, 0x03, 0x98, 0xf1, 0x10,
	0xf4, 0x3e, 0x5c, 0xef, 0x59, 0xfd, 0xf6, 0x2b, 0x5d, 0xfb, 0x7c, 0x43, 0xd7, 0xbe, 0xdc, 0xd0,
	0xb5, 0xe8, 0xe8, 0xd7, 0xcf, 0x76, 0xf0, 0xed, 0xae, 0x1d, 0xdc, 0xde, 0xb5, 0x83, 0x1f, 0x77,
	0xed, 0xe0, 0xfd, 0xf2, 0xfb, 0x35, 0xa8, 0xc0, 0xdb, 0xf3, 0xe4, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x03, 0x0e, 0xbc, 0x27, 0xe4, 0x04, 0x00, 0x00,
}

func (this *UserUpdate) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 25)
	s = append(s, "&data_update.UserUpdate{")
	s = append(s, "UpdateKey: "+fmt.Sprintf("%#v", this.UpdateKey)+",\n")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "UserId: "+fmt.Sprintf("%#v", this.UserId)+",\n")
	s = append(s, "PeerId: "+fmt.Sprintf("%#v", this.PeerId)+",\n")
	s = append(s, "PeerType: "+fmt.Sprintf("%#v", this.PeerType)+",\n")
	s = append(s, "Pts: "+fmt.Sprintf("%#v", this.Pts)+",\n")
	s = append(s, "PtsCount: "+fmt.Sprintf("%#v", this.PtsCount)+",\n")
	s = append(s, "UpdateType: "+fmt.Sprintf("%#v", this.UpdateType)+",\n")
	s = append(s, "Date: "+fmt.Sprintf("%#v", this.Date)+",\n")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "Time: "+fmt.Sprintf("%#v", this.Time)+",\n")
	s = append(s, "RandomId: "+fmt.Sprintf("%#v", this.RandomId)+",\n")
	s = append(s, "MessageBoxIds: "+fmt.Sprintf("%#v", this.MessageBoxIds)+",\n")
	s = append(s, "MessageAction: "+fmt.Sprintf("%#v", this.MessageAction)+",\n")
	s = append(s, "ChatId: "+fmt.Sprintf("%#v", this.ChatId)+",\n")
	s = append(s, "MaxId: "+fmt.Sprintf("%#v", this.MaxId)+",\n")
	if this.MessageData != nil {
		s = append(s, "MessageData: "+fmt.Sprintf("%#v", this.MessageData)+",\n")
	}
	s = append(s, "FoldId: "+fmt.Sprintf("%#v", this.FoldId)+",\n")
	s = append(s, "StillUnreadCount: "+fmt.Sprintf("%#v", this.StillUnreadCount)+",\n")
	if this.Media != nil {
		s = append(s, "Media: "+fmt.Sprintf("%#v", this.Media)+",\n")
	}
	s = append(s, "OwnerUserId: "+fmt.Sprintf("%#v", this.OwnerUserId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringUpdates(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *UserUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UpdateKey) > 0 {
		i -= len(m.UpdateKey)
		copy(dAtA[i:], m.UpdateKey)
		i = encodeVarintUpdates(dAtA, i, uint64(len(m.UpdateKey)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xaa
	}
	if m.OwnerUserId != 0 {
		i = encodeVarintUpdates(dAtA, i, uint64(m.OwnerUserId))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa0
	}
	if m.Media != nil {
		{
			size, err := m.Media.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUpdates(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x9a
	}
	if m.StillUnreadCount != 0 {
		i = encodeVarintUpdates(dAtA, i, uint64(m.StillUnreadCount))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x90
	}
	if m.FoldId != 0 {
		i = encodeVarintUpdates(dAtA, i, uint64(m.FoldId))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.MessageData != nil {
		{
			size, err := m.MessageData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUpdates(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if m.MaxId != 0 {
		i = encodeVarintUpdates(dAtA, i, uint64(m.MaxId))
		i--
		dAtA[i] = 0x78
	}
	if m.ChatId != 0 {
		i = encodeVarintUpdates(dAtA, i, uint64(m.ChatId))
		i--
		dAtA[i] = 0x70
	}
	if len(m.MessageAction) > 0 {
		i -= len(m.MessageAction)
		copy(dAtA[i:], m.MessageAction)
		i = encodeVarintUpdates(dAtA, i, uint64(len(m.MessageAction)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.MessageBoxIds) > 0 {
		dAtA4 := make([]byte, len(m.MessageBoxIds)*10)
		var j3 int
		for _, num1 := range m.MessageBoxIds {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintUpdates(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x62
	}
	if m.RandomId != 0 {
		i = encodeVarintUpdates(dAtA, i, uint64(m.RandomId))
		i--
		dAtA[i] = 0x58
	}
	if len(m.Time) > 0 {
		i -= len(m.Time)
		copy(dAtA[i:], m.Time)
		i = encodeVarintUpdates(dAtA, i, uint64(len(m.Time)))
		i--
		dAtA[i] = 0x52
	}
	if m.Version != 0 {
		i = encodeVarintUpdates(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x48
	}
	if m.Date != 0 {
		i = encodeVarintUpdates(dAtA, i, uint64(m.Date))
		i--
		dAtA[i] = 0x40
	}
	if m.UpdateType != 0 {
		i = encodeVarintUpdates(dAtA, i, uint64(m.UpdateType))
		i--
		dAtA[i] = 0x38
	}
	if m.PtsCount != 0 {
		i = encodeVarintUpdates(dAtA, i, uint64(m.PtsCount))
		i--
		dAtA[i] = 0x30
	}
	if m.Pts != 0 {
		i = encodeVarintUpdates(dAtA, i, uint64(m.Pts))
		i--
		dAtA[i] = 0x28
	}
	if m.PeerType != 0 {
		i = encodeVarintUpdates(dAtA, i, uint64(m.PeerType))
		i--
		dAtA[i] = 0x20
	}
	if m.PeerId != 0 {
		i = encodeVarintUpdates(dAtA, i, uint64(m.PeerId))
		i--
		dAtA[i] = 0x18
	}
	if m.UserId != 0 {
		i = encodeVarintUpdates(dAtA, i, uint64(m.UserId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintUpdates(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintUpdates(dAtA []byte, offset int, v uint64) int {
	offset -= sovUpdates(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUpdates(uint64(m.Id))
	}
	if m.UserId != 0 {
		n += 1 + sovUpdates(uint64(m.UserId))
	}
	if m.PeerId != 0 {
		n += 1 + sovUpdates(uint64(m.PeerId))
	}
	if m.PeerType != 0 {
		n += 1 + sovUpdates(uint64(m.PeerType))
	}
	if m.Pts != 0 {
		n += 1 + sovUpdates(uint64(m.Pts))
	}
	if m.PtsCount != 0 {
		n += 1 + sovUpdates(uint64(m.PtsCount))
	}
	if m.UpdateType != 0 {
		n += 1 + sovUpdates(uint64(m.UpdateType))
	}
	if m.Date != 0 {
		n += 1 + sovUpdates(uint64(m.Date))
	}
	if m.Version != 0 {
		n += 1 + sovUpdates(uint64(m.Version))
	}
	l = len(m.Time)
	if l > 0 {
		n += 1 + l + sovUpdates(uint64(l))
	}
	if m.RandomId != 0 {
		n += 1 + sovUpdates(uint64(m.RandomId))
	}
	if len(m.MessageBoxIds) > 0 {
		l = 0
		for _, e := range m.MessageBoxIds {
			l += sovUpdates(uint64(e))
		}
		n += 1 + sovUpdates(uint64(l)) + l
	}
	l = len(m.MessageAction)
	if l > 0 {
		n += 1 + l + sovUpdates(uint64(l))
	}
	if m.ChatId != 0 {
		n += 1 + sovUpdates(uint64(m.ChatId))
	}
	if m.MaxId != 0 {
		n += 1 + sovUpdates(uint64(m.MaxId))
	}
	if m.MessageData != nil {
		l = m.MessageData.Size()
		n += 2 + l + sovUpdates(uint64(l))
	}
	if m.FoldId != 0 {
		n += 2 + sovUpdates(uint64(m.FoldId))
	}
	if m.StillUnreadCount != 0 {
		n += 2 + sovUpdates(uint64(m.StillUnreadCount))
	}
	if m.Media != nil {
		l = m.Media.Size()
		n += 2 + l + sovUpdates(uint64(l))
	}
	if m.OwnerUserId != 0 {
		n += 2 + sovUpdates(uint64(m.OwnerUserId))
	}
	l = len(m.UpdateKey)
	if l > 0 {
		n += 2 + l + sovUpdates(uint64(l))
	}
	return n
}

func sovUpdates(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUpdates(x uint64) (n int) {
	return sovUpdates(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpdates
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			m.PeerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerType", wireType)
			}
			m.PeerType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeerType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pts", wireType)
			}
			m.Pts = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pts |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PtsCount", wireType)
			}
			m.PtsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PtsCount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateType", wireType)
			}
			m.UpdateType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			m.Date = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Date |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUpdates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RandomId", wireType)
			}
			m.RandomId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RandomId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowUpdates
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.MessageBoxIds = append(m.MessageBoxIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowUpdates
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthUpdates
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthUpdates
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.MessageBoxIds) == 0 {
					m.MessageBoxIds = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowUpdates
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.MessageBoxIds = append(m.MessageBoxIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageBoxIds", wireType)
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageAction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUpdates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageAction = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChatId", wireType)
			}
			m.ChatId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChatId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxId", wireType)
			}
			m.MaxId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUpdates
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUpdates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MessageData == nil {
				m.MessageData = &message.Message{}
			}
			if err := m.MessageData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FoldId", wireType)
			}
			m.FoldId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FoldId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 18:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StillUnreadCount", wireType)
			}
			m.StillUnreadCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StillUnreadCount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Media", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUpdates
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUpdates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Media == nil {
				m.Media = &message.Media{}
			}
			if err := m.Media.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 20:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerUserId", wireType)
			}
			m.OwnerUserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OwnerUserId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUpdates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdateKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpdates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUpdates
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUpdates
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUpdates(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUpdates
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUpdates
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthUpdates
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUpdates
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUpdates
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUpdates        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUpdates          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUpdates = fmt.Errorf("proto: unexpected end of group")
)
