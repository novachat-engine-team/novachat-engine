// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message.proto

package data_message

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Message struct {
	Did             int64     `protobuf:"varint,1,opt,name=did,proto3" json:"did,omitempty" bson:"_id"`
	UserId          int64     `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" bson:"user_id"`
	PeerId          int64     `protobuf:"varint,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty" bson:"peer_id"`
	IsChannel       bool      `protobuf:"varint,4,opt,name=is_channel,json=isChannel,proto3" json:"is_channel,omitempty" bson:"is_channel"`
	Id              int32     `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty" bson:"id"`
	Out             bool      `protobuf:"varint,6,opt,name=out,proto3" json:"out,omitempty" bson:"out"`
	GroupId         int64     `protobuf:"varint,7,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty" bson:"group_id"`
	Pinned          bool      `protobuf:"varint,8,opt,name=pinned,proto3" json:"pinned,omitempty" bson:"pinned"`
	Media           *Media    `protobuf:"bytes,9,opt,name=media,proto3" json:"media,omitempty" bson:"media"`
	Action          *Action   `protobuf:"bytes,10,opt,name=action,proto3" json:"action,omitempty" bson:"action"`
	Fwd             *Fwd      `protobuf:"bytes,11,opt,name=fwd,proto3" json:"fwd,omitempty" bson:"fwd"`
	Entities        []*Entity `protobuf:"bytes,12,rep,name=entities,proto3" json:"entities,omitempty" bson:"entities"`
	Message         string    `protobuf:"bytes,13,opt,name=message,proto3" json:"message,omitempty" bson:"message"`
	Deleted         bool      `protobuf:"varint,14,opt,name=deleted,proto3" json:"deleted,omitempty" bson:"deleted"`
	ReplyTo         int32     `protobuf:"varint,15,opt,name=reply_to,json=replyTo,proto3" json:"reply_to,omitempty" bson:"reply_to"`
	Type            int32     `protobuf:"varint,16,opt,name=type,proto3" json:"type,omitempty" bson:"type"`
	Date            int32     `protobuf:"varint,17,opt,name=date,proto3" json:"date,omitempty" bson:"date"`
	EditDate        int32     `protobuf:"varint,18,opt,name=edit_date,json=editDate,proto3" json:"edit_date,omitempty" bson:"edit_date"`
	PostAuthor      string    `protobuf:"bytes,19,opt,name=post_author,json=postAuthor,proto3" json:"post_author,omitempty" bson:"post_author"`
	GlobalMessageId int64     `protobuf:"varint,20,opt,name=global_message_id,json=globalMessageId,proto3" json:"global_message_id,omitempty" bson:"global_message_id"`
	IgnoreUserId    []int64   `protobuf:"varint,21,rep,packed,name=ignore_user_id,json=ignoreUserId,proto3" json:"ignore_user_id,omitempty" bson:"ignore_user_id"`
	FromUserId      int64     `protobuf:"varint,22,opt,name=from_user_id,json=fromUserId,proto3" json:"from_user_id,omitempty" bson:"from_user_id"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetDid() int64 {
	if m != nil {
		return m.Did
	}
	return 0
}

func (m *Message) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Message) GetPeerId() int64 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *Message) GetIsChannel() bool {
	if m != nil {
		return m.IsChannel
	}
	return false
}

func (m *Message) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Message) GetOut() bool {
	if m != nil {
		return m.Out
	}
	return false
}

func (m *Message) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *Message) GetPinned() bool {
	if m != nil {
		return m.Pinned
	}
	return false
}

func (m *Message) GetMedia() *Media {
	if m != nil {
		return m.Media
	}
	return nil
}

func (m *Message) GetAction() *Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *Message) GetFwd() *Fwd {
	if m != nil {
		return m.Fwd
	}
	return nil
}

func (m *Message) GetEntities() []*Entity {
	if m != nil {
		return m.Entities
	}
	return nil
}

func (m *Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Message) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *Message) GetReplyTo() int32 {
	if m != nil {
		return m.ReplyTo
	}
	return 0
}

func (m *Message) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Message) GetDate() int32 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *Message) GetEditDate() int32 {
	if m != nil {
		return m.EditDate
	}
	return 0
}

func (m *Message) GetPostAuthor() string {
	if m != nil {
		return m.PostAuthor
	}
	return ""
}

func (m *Message) GetGlobalMessageId() int64 {
	if m != nil {
		return m.GlobalMessageId
	}
	return 0
}

func (m *Message) GetIgnoreUserId() []int64 {
	if m != nil {
		return m.IgnoreUserId
	}
	return nil
}

func (m *Message) GetFromUserId() int64 {
	if m != nil {
		return m.FromUserId
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "data_message.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 698 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x94, 0xcd, 0x6e, 0x2b, 0x35,
	0x14, 0xc7, 0x3b, 0x4d, 0x9b, 0x0f, 0xe7, 0xab, 0x71, 0xda, 0x62, 0x2a, 0x88, 0x47, 0x66, 0x13,
	0x04, 0xa4, 0x82, 0x22, 0x21, 0xca, 0xa2, 0x6a, 0xa0, 0x88, 0x2c, 0xba, 0xb1, 0x60, 0xc3, 0x66,
	0x34, 0xa9, 0x9d, 0xd4, 0x52, 0x32, 0x8e, 0x32, 0x8e, 0xa2, 0xbe, 0x05, 0x4b, 0xb6, 0xf4, 0x69,
	0x58, 0xf6, 0x09, 0x46, 0xd0, 0xcd, 0xdd, 0x5e, 0xcd, 0x13, 0x5c, 0xf9, 0xd8, 0xd3, 0x4e, 0x7a,
	0x77, 0xe7, 0xf8, 0xff, 0xfb, 0x8f, 0xed, 0x73, 0x8e, 0x07, 0xb5, 0x97, 0x32, 0x4d, 0xe3, 0xb9,
	0x1c, 0xad, 0xd6, 0xda, 0x68, 0xdc, 0x12, 0xb1, 0x89, 0x23, 0xbf, 0x76, 0xf6, 0xcd, 0x5c, 0x99,
	0xfb, 0xcd, 0x74, 0x74, 0xa7, 0x97, 0xe7, 0x73, 0x3d, 0xd7, 0xe7, 0x00, 0x4d, 0x37, 0x33, 0xc8,
	0x20, 0x81, 0xc8, 0x99, 0xcf, 0x9a, 0x4b, 0x29, 0x54, 0xec, 0x93, 0xc6, 0x6c, 0x2b, 0x7c, 0xd8,
	0x8a, 0xef, 0x8c, 0xd2, 0x49, 0x91, 0xc9, 0xc4, 0x28, 0xf3, 0xe0, 0x32, 0xf6, 0xae, 0x8e, 0x6a,
	0xb7, 0x6e, 0x3b, 0x1c, 0xa2, 0x8a, 0x50, 0x82, 0x04, 0x61, 0x30, 0xac, 0x8c, 0x3b, 0x79, 0x46,
	0xd1, 0x34, 0xd5, 0xc9, 0x25, 0x8b, 0x94, 0x60, 0xdc, 0x4a, 0xf8, 0x2b, 0x54, 0xdb, 0xa4, 0x72,
	0x1d, 0x29, 0x41, 0xf6, 0x81, 0xc2, 0x79, 0x46, 0x3b, 0x8e, 0xf2, 0x02, 0xe3, 0x55, 0x1b, 0x4d,
	0x00, 0x5e, 0x49, 0x07, 0x57, 0xde, 0xc2, 0x5e, 0x60, 0xbc, 0x6a, 0xa3, 0x89, 0xc0, 0xdf, 0x23,
	0xa4, 0xd2, 0xe8, 0xee, 0x3e, 0x4e, 0x12, 0xb9, 0x20, 0x07, 0x61, 0x30, 0xac, 0x8f, 0x4f, 0xf2,
	0x8c, 0xf6, 0x1c, 0xff, 0xaa, 0x31, 0xde, 0x50, 0xe9, 0xcf, 0x2e, 0xc6, 0x9f, 0xa3, 0x7d, 0x25,
	0xc8, 0x61, 0x18, 0x0c, 0x0f, 0xc7, 0xed, 0x3c, 0xa3, 0x0d, 0x4f, 0x0b, 0xc6, 0xf7, 0x95, 0xb0,
	0x17, 0xd2, 0x1b, 0x43, 0xaa, 0xf0, 0xb5, 0xd2, 0x85, 0xf4, 0xc6, 0x30, 0x6e, 0x25, 0x3c, 0x42,
	0xf5, 0xf9, 0x5a, 0x6f, 0x56, 0xf6, 0x90, 0x35, 0x38, 0x64, 0x3f, 0xcf, 0x68, 0xd7, 0x61, 0x85,
	0xc2, 0x78, 0x0d, 0xc2, 0x89, 0xc0, 0x5f, 0xa2, 0xea, 0x4a, 0x25, 0x89, 0x14, 0xa4, 0x0e, 0x1f,
	0xed, 0xe5, 0x19, 0x6d, 0xfb, 0x2b, 0xc1, 0xba, 0xbd, 0x11, 0x04, 0xf8, 0x27, 0x74, 0x08, 0xfd,
	0x20, 0x8d, 0x30, 0x18, 0x36, 0xbf, 0xeb, 0x8f, 0xca, 0xad, 0x1d, 0xdd, 0x5a, 0x69, 0x7c, 0x94,
	0x67, 0xb4, 0xe5, 0xec, 0xc0, 0x32, 0xee, 0x3c, 0xf8, 0x0a, 0x55, 0x5d, 0xd3, 0x08, 0x02, 0xf7,
	0xf1, 0xae, 0xfb, 0x1a, 0xb4, 0xf2, 0xee, 0x8e, 0x66, 0xdc, 0xdb, 0xf0, 0x05, 0xaa, 0xcc, 0xb6,
	0x82, 0x34, 0xc1, 0xdd, 0xdb, 0x75, 0xff, 0xba, 0x15, 0xe5, 0x6a, 0xcc, 0xb6, 0xb6, 0xbd, 0xb3,
	0xad, 0xc0, 0x37, 0xa8, 0x0e, 0xc3, 0xa1, 0x64, 0x4a, 0x5a, 0x61, 0xe5, 0xe3, 0x7d, 0x6f, 0x60,
	0x74, 0xca, 0x35, 0x2a, 0x78, 0xc6, 0x5f, 0xac, 0xf8, 0x6b, 0x54, 0xf3, 0x06, 0xd2, 0x0e, 0x83,
	0x61, 0xa3, 0xdc, 0x78, 0x2f, 0x30, 0x5e, 0x20, 0x96, 0x16, 0x72, 0x21, 0x8d, 0x14, 0xa4, 0x03,
	0x35, 0x2d, 0xd1, 0x5e, 0x60, 0xbc, 0x40, 0x6c, 0xc3, 0xd6, 0x72, 0xb5, 0x78, 0x88, 0x8c, 0x26,
	0x5d, 0xe8, 0x7b, 0xe9, 0x30, 0x85, 0xc2, 0x78, 0x0d, 0xc2, 0xdf, 0x35, 0xfe, 0x02, 0x1d, 0x98,
	0x87, 0x95, 0x24, 0x47, 0xc0, 0x76, 0xf3, 0x8c, 0x36, 0x1d, 0x6b, 0x57, 0x19, 0x07, 0xd1, 0x42,
	0x22, 0x36, 0x92, 0xf4, 0xde, 0x42, 0x76, 0x95, 0x71, 0x10, 0xf1, 0xb7, 0xa8, 0x21, 0x85, 0x32,
	0x11, 0x90, 0x18, 0xc8, 0xe3, 0x3c, 0xa3, 0x47, 0xbe, 0x0e, 0x85, 0x64, 0x0b, 0x21, 0x94, 0xf9,
	0xc5, 0x5a, 0x7e, 0x40, 0xcd, 0x95, 0x4e, 0x4d, 0x14, 0x6f, 0xcc, 0xbd, 0x5e, 0x93, 0x3e, 0x14,
	0xe3, 0x34, 0xcf, 0x28, 0xf6, 0x23, 0xf3, 0x2a, 0x32, 0x8e, 0x6c, 0x76, 0x0d, 0x09, 0xfe, 0x0d,
	0xf5, 0xe6, 0x0b, 0x3d, 0x8d, 0x17, 0x45, 0xe5, 0xed, 0x7c, 0x1e, 0xc3, 0x7c, 0x7e, 0x96, 0x67,
	0x94, 0xf8, 0xf9, 0x7c, 0x8b, 0x30, 0xde, 0x75, 0x6b, 0xfe, 0x45, 0x4f, 0x04, 0xbe, 0x42, 0x1d,
	0x35, 0x4f, 0xf4, 0x5a, 0x46, 0xc5, 0xc3, 0x3d, 0x09, 0x2b, 0xc3, 0xca, 0xf8, 0xd3, 0x3c, 0xa3,
	0x27, 0xfe, 0xb5, 0xec, 0xe8, 0x8c, 0xb7, 0xdc, 0xc2, 0x1f, 0xee, 0x15, 0xff, 0x88, 0x5a, 0xb3,
	0xb5, 0x5e, 0xbe, 0xd8, 0x4f, 0xe1, 0x14, 0x9f, 0xe4, 0x19, 0xed, 0xfb, 0xf1, 0x29, 0xa9, 0x8c,
	0x23, 0x9b, 0x3a, 0xeb, 0x65, 0xeb, 0xe9, 0x1f, 0xba, 0xf7, 0xd7, 0x23, 0xdd, 0xfb, 0xfb, 0x91,
	0xee, 0x8d, 0x07, 0xef, 0xff, 0x1f, 0x04, 0xff, 0x3e, 0x0f, 0x82, 0xa7, 0xe7, 0x41, 0xf0, 0xdf,
	0xf3, 0x20, 0xf8, 0x73, 0xe7, 0x67, 0x37, 0xad, 0xc2, 0x0f, 0xe9, 0xe2, 0x43, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x95, 0x0c, 0x81, 0x03, 0x12, 0x05, 0x00, 0x00,
}

func (this *Message) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 26)
	s = append(s, "&data_message.Message{")
	s = append(s, "Did: "+fmt.Sprintf("%#v", this.Did)+",\n")
	s = append(s, "UserId: "+fmt.Sprintf("%#v", this.UserId)+",\n")
	s = append(s, "PeerId: "+fmt.Sprintf("%#v", this.PeerId)+",\n")
	s = append(s, "IsChannel: "+fmt.Sprintf("%#v", this.IsChannel)+",\n")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "Out: "+fmt.Sprintf("%#v", this.Out)+",\n")
	s = append(s, "GroupId: "+fmt.Sprintf("%#v", this.GroupId)+",\n")
	s = append(s, "Pinned: "+fmt.Sprintf("%#v", this.Pinned)+",\n")
	if this.Media != nil {
		s = append(s, "Media: "+fmt.Sprintf("%#v", this.Media)+",\n")
	}
	if this.Action != nil {
		s = append(s, "Action: "+fmt.Sprintf("%#v", this.Action)+",\n")
	}
	if this.Fwd != nil {
		s = append(s, "Fwd: "+fmt.Sprintf("%#v", this.Fwd)+",\n")
	}
	if this.Entities != nil {
		s = append(s, "Entities: "+fmt.Sprintf("%#v", this.Entities)+",\n")
	}
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "Deleted: "+fmt.Sprintf("%#v", this.Deleted)+",\n")
	s = append(s, "ReplyTo: "+fmt.Sprintf("%#v", this.ReplyTo)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Date: "+fmt.Sprintf("%#v", this.Date)+",\n")
	s = append(s, "EditDate: "+fmt.Sprintf("%#v", this.EditDate)+",\n")
	s = append(s, "PostAuthor: "+fmt.Sprintf("%#v", this.PostAuthor)+",\n")
	s = append(s, "GlobalMessageId: "+fmt.Sprintf("%#v", this.GlobalMessageId)+",\n")
	s = append(s, "IgnoreUserId: "+fmt.Sprintf("%#v", this.IgnoreUserId)+",\n")
	s = append(s, "FromUserId: "+fmt.Sprintf("%#v", this.FromUserId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMessage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FromUserId != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.FromUserId))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb0
	}
	if len(m.IgnoreUserId) > 0 {
		dAtA2 := make([]byte, len(m.IgnoreUserId)*10)
		var j1 int
		for _, num1 := range m.IgnoreUserId {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintMessage(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xaa
	}
	if m.GlobalMessageId != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.GlobalMessageId))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa0
	}
	if len(m.PostAuthor) > 0 {
		i -= len(m.PostAuthor)
		copy(dAtA[i:], m.PostAuthor)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.PostAuthor)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x9a
	}
	if m.EditDate != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.EditDate))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x90
	}
	if m.Date != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Date))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.Type != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.ReplyTo != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.ReplyTo))
		i--
		dAtA[i] = 0x78
	}
	if m.Deleted {
		i--
		if m.Deleted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x70
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.Entities) > 0 {
		for iNdEx := len(m.Entities) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entities[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if m.Fwd != nil {
		{
			size, err := m.Fwd.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if m.Action != nil {
		{
			size, err := m.Action.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.Media != nil {
		{
			size, err := m.Media.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.Pinned {
		i--
		if m.Pinned {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.GroupId != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.GroupId))
		i--
		dAtA[i] = 0x38
	}
	if m.Out {
		i--
		if m.Out {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.Id != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x28
	}
	if m.IsChannel {
		i--
		if m.IsChannel {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.PeerId != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.PeerId))
		i--
		dAtA[i] = 0x18
	}
	if m.UserId != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.UserId))
		i--
		dAtA[i] = 0x10
	}
	if m.Did != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Did))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Did != 0 {
		n += 1 + sovMessage(uint64(m.Did))
	}
	if m.UserId != 0 {
		n += 1 + sovMessage(uint64(m.UserId))
	}
	if m.PeerId != 0 {
		n += 1 + sovMessage(uint64(m.PeerId))
	}
	if m.IsChannel {
		n += 2
	}
	if m.Id != 0 {
		n += 1 + sovMessage(uint64(m.Id))
	}
	if m.Out {
		n += 2
	}
	if m.GroupId != 0 {
		n += 1 + sovMessage(uint64(m.GroupId))
	}
	if m.Pinned {
		n += 2
	}
	if m.Media != nil {
		l = m.Media.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.Action != nil {
		l = m.Action.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.Fwd != nil {
		l = m.Fwd.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if len(m.Entities) > 0 {
		for _, e := range m.Entities {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.Deleted {
		n += 2
	}
	if m.ReplyTo != 0 {
		n += 1 + sovMessage(uint64(m.ReplyTo))
	}
	if m.Type != 0 {
		n += 2 + sovMessage(uint64(m.Type))
	}
	if m.Date != 0 {
		n += 2 + sovMessage(uint64(m.Date))
	}
	if m.EditDate != 0 {
		n += 2 + sovMessage(uint64(m.EditDate))
	}
	l = len(m.PostAuthor)
	if l > 0 {
		n += 2 + l + sovMessage(uint64(l))
	}
	if m.GlobalMessageId != 0 {
		n += 2 + sovMessage(uint64(m.GlobalMessageId))
	}
	if len(m.IgnoreUserId) > 0 {
		l = 0
		for _, e := range m.IgnoreUserId {
			l += sovMessage(uint64(e))
		}
		n += 2 + sovMessage(uint64(l)) + l
	}
	if m.FromUserId != 0 {
		n += 2 + sovMessage(uint64(m.FromUserId))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			m.Did = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Did |= int64(b&0x7F) << shift
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
					return ErrIntOverflowMessage
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
					return ErrIntOverflowMessage
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
				return fmt.Errorf("proto: wrong wireType = %d for field IsChannel", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsChannel = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Out", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Out = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			m.GroupId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pinned", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Pinned = bool(v != 0)
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Media", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Media == nil {
				m.Media = &Media{}
			}
			if err := m.Media.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Action == nil {
				m.Action = &Action{}
			}
			if err := m.Action.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fwd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fwd == nil {
				m.Fwd = &Fwd{}
			}
			if err := m.Fwd.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entities", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entities = append(m.Entities, &Entity{})
			if err := m.Entities[len(m.Entities)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deleted = bool(v != 0)
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplyTo", wireType)
			}
			m.ReplyTo = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReplyTo |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			m.Date = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
		case 18:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EditDate", wireType)
			}
			m.EditDate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EditDate |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostAuthor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostAuthor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 20:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalMessageId", wireType)
			}
			m.GlobalMessageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GlobalMessageId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 21:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMessage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.IgnoreUserId = append(m.IgnoreUserId, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMessage
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
					return ErrInvalidLengthMessage
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthMessage
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
				if elementCount != 0 && len(m.IgnoreUserId) == 0 {
					m.IgnoreUserId = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMessage
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.IgnoreUserId = append(m.IgnoreUserId, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field IgnoreUserId", wireType)
			}
		case 22:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromUserId", wireType)
			}
			m.FromUserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FromUserId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)