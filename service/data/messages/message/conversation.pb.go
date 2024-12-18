// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: conversation.proto

package data_message

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

type PinnedMessage struct {
	Id          int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"id"`
	OwnerUserId int64 `protobuf:"varint,2,opt,name=owner_user_id,json=ownerUserId,proto3" json:"owner_user_id,omitempty" bson:"oui"`
	Full        bool  `protobuf:"varint,3,opt,name=full,proto3" json:"full,omitempty" bson:"m"`
}

func (m *PinnedMessage) Reset()         { *m = PinnedMessage{} }
func (m *PinnedMessage) String() string { return proto.CompactTextString(m) }
func (*PinnedMessage) ProtoMessage()    {}
func (*PinnedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6ae46d001825652, []int{0}
}
func (m *PinnedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PinnedMessage.Unmarshal(m, b)
}
func (m *PinnedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PinnedMessage.Marshal(b, m, deterministic)
}
func (m *PinnedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PinnedMessage.Merge(m, src)
}
func (m *PinnedMessage) XXX_Size() int {
	return xxx_messageInfo_PinnedMessage.Size(m)
}
func (m *PinnedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PinnedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PinnedMessage proto.InternalMessageInfo

func (m *PinnedMessage) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PinnedMessage) GetOwnerUserId() int64 {
	if m != nil {
		return m.OwnerUserId
	}
	return 0
}

func (m *PinnedMessage) GetFull() bool {
	if m != nil {
		return m.Full
	}
	return false
}

type Conversation struct {
	Id          string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	UserId      int64            `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" bson:"user_id"`
	PeerId      int64            `protobuf:"varint,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty" bson:"peer_id"`
	PeerType    int32            `protobuf:"varint,4,opt,name=peer_type,json=peerType,proto3" json:"peer_type,omitempty" bson:"peer_type"`
	Top         int32            `protobuf:"varint,5,opt,name=top,proto3" json:"top,omitempty" bson:"top"`
	Date        int32            `protobuf:"varint,6,opt,name=date,proto3" json:"date,omitempty" bson:"date"`
	UnreadCount int32            `protobuf:"varint,7,opt,name=unread_count,json=unreadCount,proto3" json:"unread_count,omitempty" bson:"unread_count"`
	InboxMaxId  int32            `protobuf:"varint,8,opt,name=inbox_max_id,json=inboxMaxId,proto3" json:"inbox_max_id,omitempty" bson:"inbox_max_id"`
	OutboxMaxId int32            `protobuf:"varint,9,opt,name=outbox_max_id,json=outboxMaxId,proto3" json:"outbox_max_id,omitempty" bson:"outbox_max_id"`
	Draft       string           `protobuf:"bytes,10,opt,name=draft,proto3" json:"draft,omitempty" bson:"draft"`
	Pinned      bool             `protobuf:"varint,11,opt,name=pinned,proto3" json:"pinned,omitempty" bson:"pinned"`
	FolderId    int32            `protobuf:"varint,12,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty" bson:"folder_id"`
	PinIds      []*PinnedMessage `protobuf:"bytes,13,rep,name=pin_ids,json=pinIds,proto3" json:"pin_ids,omitempty" bson:"pin_ids"`
	Pts         int32            `protobuf:"varint,14,opt,name=pts,proto3" json:"pts,omitempty" bson:"pts"`
	ClearMaxId  int32            `protobuf:"varint,15,opt,name=clear_max_id,json=clearMaxId,proto3" json:"clear_max_id,omitempty" bson:"clear_max_id"`
	Deleted     bool             `protobuf:"varint,16,opt,name=deleted,proto3" json:"deleted,omitempty" bson:"deleted"`
}

func (m *Conversation) Reset()         { *m = Conversation{} }
func (m *Conversation) String() string { return proto.CompactTextString(m) }
func (*Conversation) ProtoMessage()    {}
func (*Conversation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6ae46d001825652, []int{1}
}
func (m *Conversation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Conversation.Unmarshal(m, b)
}
func (m *Conversation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Conversation.Marshal(b, m, deterministic)
}
func (m *Conversation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Conversation.Merge(m, src)
}
func (m *Conversation) XXX_Size() int {
	return xxx_messageInfo_Conversation.Size(m)
}
func (m *Conversation) XXX_DiscardUnknown() {
	xxx_messageInfo_Conversation.DiscardUnknown(m)
}

var xxx_messageInfo_Conversation proto.InternalMessageInfo

func (m *Conversation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Conversation) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Conversation) GetPeerId() int64 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *Conversation) GetPeerType() int32 {
	if m != nil {
		return m.PeerType
	}
	return 0
}

func (m *Conversation) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *Conversation) GetDate() int32 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *Conversation) GetUnreadCount() int32 {
	if m != nil {
		return m.UnreadCount
	}
	return 0
}

func (m *Conversation) GetInboxMaxId() int32 {
	if m != nil {
		return m.InboxMaxId
	}
	return 0
}

func (m *Conversation) GetOutboxMaxId() int32 {
	if m != nil {
		return m.OutboxMaxId
	}
	return 0
}

func (m *Conversation) GetDraft() string {
	if m != nil {
		return m.Draft
	}
	return ""
}

func (m *Conversation) GetPinned() bool {
	if m != nil {
		return m.Pinned
	}
	return false
}

func (m *Conversation) GetFolderId() int32 {
	if m != nil {
		return m.FolderId
	}
	return 0
}

func (m *Conversation) GetPinIds() []*PinnedMessage {
	if m != nil {
		return m.PinIds
	}
	return nil
}

func (m *Conversation) GetPts() int32 {
	if m != nil {
		return m.Pts
	}
	return 0
}

func (m *Conversation) GetClearMaxId() int32 {
	if m != nil {
		return m.ClearMaxId
	}
	return 0
}

func (m *Conversation) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func init() {
	proto.RegisterType((*PinnedMessage)(nil), "data_message.PinnedMessage")
	proto.RegisterType((*Conversation)(nil), "data_message.Conversation")
}

func init() { proto.RegisterFile("conversation.proto", fileDescriptor_d6ae46d001825652) }

var fileDescriptor_d6ae46d001825652 = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x53, 0xcf, 0x8e, 0xd3, 0x3e,
	0x10, 0xde, 0xfe, 0x4b, 0x5b, 0x27, 0xe9, 0xf6, 0xe7, 0xdf, 0x4a, 0x58, 0x20, 0x70, 0x64, 0x24,
	0x54, 0xb4, 0xd0, 0x15, 0xcb, 0x89, 0x8a, 0x53, 0xf7, 0x42, 0x0f, 0x2b, 0x21, 0x0b, 0x2e, 0x5c,
	0xa2, 0xb4, 0x76, 0x4b, 0xa4, 0x36, 0x8e, 0x12, 0x07, 0xba, 0x6f, 0x81, 0x38, 0x71, 0x84, 0x7d,
	0x22, 0x9e, 0x20, 0x0f, 0x91, 0x27, 0x40, 0x1e, 0xa7, 0xab, 0xb4, 0xb7, 0x99, 0xf9, 0xbe, 0x49,
	0x66, 0xbe, 0xcf, 0x83, 0xf0, 0x4a, 0x25, 0xdf, 0x64, 0x96, 0x47, 0x3a, 0x56, 0xc9, 0x34, 0xcd,
	0x94, 0x56, 0xd8, 0x13, 0x91, 0x8e, 0xc2, 0x9d, 0xcc, 0xf3, 0x68, 0x23, 0x1f, 0xbf, 0xde, 0xc4,
	0xfa, 0x6b, 0xb1, 0x9c, 0xae, 0xd4, 0xee, 0x6a, 0xa3, 0x36, 0xea, 0x0a, 0x48, 0xcb, 0x62, 0x0d,
	0x19, 0x24, 0x10, 0xd9, 0x66, 0xf6, 0xb3, 0x85, 0xfc, 0x8f, 0x71, 0x92, 0x48, 0x71, 0x6b, 0x3f,
	0x80, 0x9f, 0xa2, 0x76, 0x2c, 0x48, 0x2b, 0x68, 0x4d, 0x7a, 0x73, 0xbf, 0x2a, 0xe9, 0x70, 0x99,
	0xab, 0x64, 0xc6, 0x62, 0xc1, 0x78, 0x3b, 0x16, 0xf8, 0x1a, 0xf9, 0xea, 0x7b, 0x22, 0xb3, 0xb0,
	0xc8, 0x65, 0x16, 0xc6, 0x82, 0xb4, 0x83, 0xd6, 0xa4, 0x33, 0x1f, 0x55, 0x25, 0x45, 0x96, 0xa9,
	0x8a, 0x98, 0x71, 0x17, 0x48, 0x9f, 0x73, 0x99, 0x2d, 0x04, 0x0e, 0x50, 0x77, 0x5d, 0x6c, 0xb7,
	0xa4, 0x13, 0xb4, 0x26, 0x83, 0xb9, 0x57, 0x95, 0x74, 0x60, 0xa9, 0x3b, 0xc6, 0x01, 0x99, 0x79,
	0x7f, 0xff, 0xd0, 0xb3, 0x1f, 0xf7, 0xf4, 0xec, 0xd7, 0x3d, 0x3d, 0x63, 0xbf, 0x1d, 0xe4, 0xdd,
	0x34, 0x16, 0xc5, 0xcf, 0x1e, 0x66, 0x1a, 0x36, 0xff, 0x14, 0x1e, 0x86, 0xba, 0x44, 0xfd, 0xe3,
	0x71, 0x70, 0x55, 0xd2, 0x91, 0x25, 0xd5, 0x00, 0xe3, 0x4e, 0x61, 0xa7, 0xb9, 0x44, 0xfd, 0x54,
	0x5a, 0x72, 0xe7, 0x94, 0x5c, 0x03, 0x8c, 0x3b, 0x26, 0x5a, 0x08, 0xfc, 0x06, 0x0d, 0xa1, 0xa6,
	0xef, 0x52, 0x49, 0xba, 0x20, 0xca, 0x45, 0x55, 0xd2, 0x71, 0x83, 0x6e, 0x20, 0xc6, 0x07, 0x26,
	0xfe, 0x74, 0x97, 0x4a, 0x1c, 0xa0, 0x8e, 0x56, 0x29, 0xe9, 0x01, 0xb9, 0x31, 0xad, 0x56, 0x29,
	0xe3, 0x06, 0xc2, 0xcf, 0x51, 0x57, 0x44, 0x5a, 0x12, 0x07, 0x28, 0xe7, 0x55, 0x49, 0x5d, 0x4b,
	0x31, 0x55, 0xc6, 0x01, 0xc4, 0x33, 0xe4, 0x15, 0x49, 0x26, 0x23, 0x11, 0xae, 0x54, 0x91, 0x68,
	0xd2, 0x07, 0xf2, 0xa3, 0xaa, 0xa4, 0xff, 0xd7, 0x8b, 0x35, 0x50, 0xc6, 0x5d, 0x9b, 0xde, 0x98,
	0x0c, 0xbf, 0x43, 0x5e, 0x9c, 0x2c, 0xd5, 0x3e, 0xdc, 0x45, 0x7b, 0xb3, 0xe7, 0xe0, 0xb4, 0xb7,
	0x89, 0x32, 0x8e, 0x20, 0xbd, 0x8d, 0xf6, 0x0b, 0x81, 0xdf, 0x23, 0x5f, 0x15, 0xba, 0xd1, 0x3b,
	0x84, 0x5e, 0x52, 0x95, 0xf4, 0xe2, 0xe0, 0xaf, 0x6e, 0x36, 0xbb, 0x36, 0xb7, 0xdd, 0x2f, 0x50,
	0x4f, 0x64, 0xd1, 0x5a, 0x13, 0x04, 0x5e, 0x8d, 0xab, 0x92, 0x7a, 0xf5, 0x6a, 0xa6, 0xcc, 0xb8,
	0x85, 0xf1, 0x4b, 0xe4, 0xa4, 0xf0, 0xea, 0x88, 0x0b, 0x6f, 0xe2, 0xbf, 0xaa, 0xa4, 0x7e, 0xad,
	0x29, 0xd4, 0x8d, 0x03, 0x10, 0x18, 0x07, 0xd6, 0x6a, 0x2b, 0xac, 0x61, 0xde, 0xa9, 0x03, 0x0f,
	0x10, 0xe3, 0x03, 0x1b, 0x2f, 0x04, 0xfe, 0x80, 0xfa, 0x69, 0x9c, 0x84, 0xb1, 0xc8, 0x89, 0x1f,
	0x74, 0x26, 0xee, 0xf5, 0x93, 0x69, 0xf3, 0x46, 0xa6, 0x47, 0x0f, 0xfe, 0xc8, 0x7e, 0xdb, 0x65,
	0x7f, 0xbe, 0x10, 0xb9, 0xf1, 0x32, 0xd5, 0x39, 0x19, 0x9d, 0x7a, 0x99, 0xea, 0x9c, 0x71, 0x03,
	0x19, 0xa9, 0x57, 0x5b, 0x19, 0x65, 0x07, 0xb9, 0xce, 0x4f, 0xa5, 0x6e, 0xa2, 0x8c, 0x23, 0x48,
	0xad, 0x58, 0xaf, 0x50, 0x5f, 0xc8, 0xad, 0xd4, 0x52, 0x90, 0x31, 0xa8, 0xd0, 0x98, 0xa4, 0x06,
	0x18, 0x3f, 0x50, 0x8e, 0x4f, 0x64, 0x3e, 0xfa, 0x72, 0x74, 0xf6, 0x4b, 0x07, 0xce, 0xf9, 0xed,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x50, 0x87, 0xa8, 0xc5, 0x21, 0x04, 0x00, 0x00,
}
