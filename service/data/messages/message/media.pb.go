// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: media.proto

package data_message

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	fs "novachat_engine/service/data/fs"
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

type Media struct {
	Type               int32        `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty" bson:"type"`
	PhoneNumber        string       `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty" bson:"phone_number"`
	FirstName          string       `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty" bson:"first_name"`
	LastName           string       `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty" bson:"last_name"`
	UserId             int64        `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" bson:"user_id"`
	Vcard              string       `protobuf:"bytes,6,opt,name=vcard,proto3" json:"vcard,omitempty" bson:"vcard"`
	Document           *fs.Document `protobuf:"bytes,7,opt,name=document,proto3" json:"document,omitempty" bson:"document"`
	Caption            string       `protobuf:"bytes,8,opt,name=caption,proto3" json:"caption,omitempty" bson:"caption"`
	TtlSeconds         int32        `protobuf:"varint,9,opt,name=ttl_seconds,json=ttlSeconds,proto3" json:"ttl_seconds,omitempty" bson:"ttl_seconds"`
	Photo              *fs.Photo    `protobuf:"bytes,10,opt,name=photo,proto3" json:"photo,omitempty" bson:"photo"`
	GeoPoint           *fs.GeoPoint `protobuf:"bytes,11,opt,name=geo_point,json=geoPoint,proto3" json:"geo_point,omitempty" bson:"geo_point"`
	Period             int32        `protobuf:"varint,12,opt,name=period,proto3" json:"period,omitempty" bson:"period"`
	Heading            int32        `protobuf:"varint,13,opt,name=heading,proto3" json:"heading,omitempty" bson:"heading"`
	NotificationRadius int32        `protobuf:"varint,14,opt,name=notification_radius,json=notificationRadius,proto3" json:"notification_radius,omitempty" bson:"notification_radius"`
}

func (m *Media) Reset()         { *m = Media{} }
func (m *Media) String() string { return proto.CompactTextString(m) }
func (*Media) ProtoMessage()    {}
func (*Media) Descriptor() ([]byte, []int) {
	return fileDescriptor_07eb54b56db72a97, []int{0}
}
func (m *Media) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Media) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Media.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Media) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Media.Merge(m, src)
}
func (m *Media) XXX_Size() int {
	return m.Size()
}
func (m *Media) XXX_DiscardUnknown() {
	xxx_messageInfo_Media.DiscardUnknown(m)
}

var xxx_messageInfo_Media proto.InternalMessageInfo

func (m *Media) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Media) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Media) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Media) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Media) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Media) GetVcard() string {
	if m != nil {
		return m.Vcard
	}
	return ""
}

func (m *Media) GetDocument() *fs.Document {
	if m != nil {
		return m.Document
	}
	return nil
}

func (m *Media) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func (m *Media) GetTtlSeconds() int32 {
	if m != nil {
		return m.TtlSeconds
	}
	return 0
}

func (m *Media) GetPhoto() *fs.Photo {
	if m != nil {
		return m.Photo
	}
	return nil
}

func (m *Media) GetGeoPoint() *fs.GeoPoint {
	if m != nil {
		return m.GeoPoint
	}
	return nil
}

func (m *Media) GetPeriod() int32 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *Media) GetHeading() int32 {
	if m != nil {
		return m.Heading
	}
	return 0
}

func (m *Media) GetNotificationRadius() int32 {
	if m != nil {
		return m.NotificationRadius
	}
	return 0
}

func init() {
	proto.RegisterType((*Media)(nil), "data_message.Media")
}

func init() { proto.RegisterFile("media.proto", fileDescriptor_07eb54b56db72a97) }

var fileDescriptor_07eb54b56db72a97 = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xcd, 0x8e, 0xd3, 0x3e,
	0x14, 0xc5, 0x9b, 0xff, 0x4c, 0xbf, 0x9c, 0x4e, 0x67, 0xea, 0xfe, 0x01, 0xab, 0x8b, 0xb8, 0x32,
	0x12, 0x2a, 0x02, 0x3a, 0xe2, 0x43, 0x20, 0x75, 0x19, 0x8d, 0x84, 0x58, 0x30, 0x8c, 0xcc, 0x8e,
	0x4d, 0xe4, 0x26, 0x6e, 0x1a, 0xa9, 0x89, 0xa3, 0xc4, 0x41, 0xe2, 0x09, 0xd8, 0xb2, 0x64, 0xcb,
	0x3c, 0x0d, 0xcb, 0x79, 0x82, 0x08, 0xfa, 0x04, 0x28, 0x4f, 0x80, 0x72, 0x93, 0x74, 0x02, 0x62,
	0xe7, 0x73, 0x7f, 0xe7, 0xd4, 0x37, 0xf7, 0xd6, 0xc8, 0x0c, 0xa5, 0x17, 0x88, 0x65, 0x9c, 0x28,
	0xad, 0xf0, 0xc8, 0x13, 0x5a, 0x38, 0xa1, 0x4c, 0x53, 0xe1, 0xcb, 0xd9, 0x13, 0x3f, 0xd0, 0xdb,
	0x6c, 0xbd, 0x74, 0x55, 0x78, 0xee, 0x2b, 0x5f, 0x9d, 0x83, 0x69, 0x9d, 0x6d, 0x40, 0x81, 0x80,
	0x53, 0x15, 0x9e, 0x8d, 0x3d, 0xe5, 0x66, 0xa1, 0x8c, 0x74, 0xad, 0xcd, 0x78, 0xab, 0xf4, 0x01,
	0xfa, 0x52, 0xc5, 0x2a, 0x68, 0x20, 0xfb, 0xdc, 0x43, 0xdd, 0xb7, 0xe5, 0xcd, 0xf8, 0x3e, 0x3a,
	0xd6, 0x9f, 0x62, 0x49, 0x8c, 0xb9, 0xb1, 0xe8, 0xda, 0xa7, 0x45, 0x4e, 0xcd, 0x75, 0xaa, 0xa2,
	0x15, 0x2b, 0xab, 0x8c, 0x03, 0xc4, 0x2b, 0x34, 0x8a, 0xb7, 0x2a, 0x92, 0x4e, 0x94, 0x85, 0x6b,
	0x99, 0x90, 0xff, 0xe6, 0xc6, 0x62, 0x68, 0xdf, 0x2b, 0x72, 0x3a, 0xad, 0xcc, 0x6d, 0xca, 0xb8,
	0x09, 0xf2, 0x12, 0x14, 0x7e, 0x81, 0xd0, 0x26, 0x48, 0x52, 0xed, 0x44, 0x22, 0x94, 0xe4, 0x08,
	0x92, 0x77, 0x8a, 0x9c, 0x4e, 0xaa, 0xe4, 0x2d, 0x63, 0x7c, 0x08, 0xe2, 0x52, 0x84, 0x12, 0x3f,
	0x45, 0xc3, 0x9d, 0x68, 0x42, 0xc7, 0x10, 0xfa, 0xbf, 0xc8, 0xe9, 0x59, 0x15, 0x3a, 0x20, 0xc6,
	0x07, 0xe5, 0x19, 0x22, 0x8f, 0x50, 0x3f, 0x4b, 0x65, 0xe2, 0x04, 0x1e, 0xe9, 0xce, 0x8d, 0xc5,
	0x91, 0x8d, 0x8b, 0x9c, 0x8e, 0xab, 0x40, 0x0d, 0x18, 0xef, 0x95, 0xa7, 0x37, 0x1e, 0x7e, 0x80,
	0xba, 0x1f, 0x5d, 0x91, 0x78, 0xa4, 0x07, 0xbf, 0x7d, 0x56, 0xe4, 0x74, 0x54, 0x59, 0xa1, 0xcc,
	0x78, 0x85, 0xb1, 0x8d, 0x06, 0xcd, 0x5c, 0x49, 0x7f, 0x6e, 0x2c, 0xcc, 0x67, 0x93, 0x25, 0x6c,
	0x69, 0x93, 0x2e, 0x2f, 0x6a, 0x60, 0x4f, 0x8b, 0x9c, 0x9e, 0x56, 0xe9, 0xc6, 0xcc, 0xf8, 0x21,
	0x87, 0x1f, 0xa3, 0xbe, 0x2b, 0x62, 0x1d, 0xa8, 0x88, 0x0c, 0xe0, 0xb6, 0x56, 0x63, 0x35, 0x60,
	0xbc, 0xb1, 0xe0, 0x57, 0xc8, 0xd4, 0x7a, 0xe7, 0xa4, 0xd2, 0x55, 0x91, 0x97, 0x92, 0x21, 0xec,
	0xe5, 0x6e, 0x91, 0x53, 0x5c, 0xef, 0xe5, 0x16, 0x32, 0x8e, 0xb4, 0xde, 0xbd, 0xaf, 0x04, 0x7e,
	0x89, 0xba, 0xb0, 0x72, 0x82, 0xa0, 0xcf, 0xf1, 0xa1, 0xcf, 0xab, 0xb2, 0xda, 0xfe, 0x44, 0xb0,
	0x31, 0x5e, 0xd9, 0xf1, 0x05, 0x1a, 0xfa, 0x52, 0x39, 0xf0, 0xf7, 0x20, 0xe6, 0x5f, 0xdf, 0xf8,
	0x5a, 0xaa, 0xab, 0x12, 0xb4, 0xa7, 0x7f, 0x70, 0x33, 0x3e, 0xf0, 0x6b, 0x8e, 0x1f, 0xa2, 0x5e,
	0x2c, 0x93, 0x40, 0x79, 0x64, 0x04, 0x1d, 0x4f, 0x8a, 0x9c, 0x9e, 0xd4, 0xd7, 0x41, 0x9d, 0xf1,
	0xda, 0x50, 0xce, 0x63, 0x2b, 0x85, 0x17, 0x44, 0x3e, 0x39, 0x01, 0x6f, 0x6b, 0x1e, 0x35, 0x60,
	0xbc, 0xb1, 0xe0, 0x77, 0x68, 0x1a, 0x29, 0x1d, 0x6c, 0x02, 0x57, 0x94, 0xf3, 0x71, 0x12, 0xe1,
	0x05, 0x59, 0x4a, 0xc6, 0x90, 0xb4, 0x8a, 0x9c, 0xce, 0xaa, 0xe4, 0x3f, 0x4c, 0x8c, 0xe3, 0x76,
	0x95, 0x43, 0x71, 0x35, 0xba, 0xf9, 0x46, 0x3b, 0x5f, 0xae, 0x69, 0xe7, 0xeb, 0x35, 0xed, 0xd8,
	0xd6, 0xaf, 0x9f, 0x96, 0xf1, 0x7d, 0x6f, 0x19, 0x37, 0x7b, 0xcb, 0xf8, 0xb1, 0xb7, 0x8c, 0x0f,
	0x7f, 0xbc, 0xc2, 0x75, 0x0f, 0x1e, 0xcc, 0xf3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x47, 0xff,
	0x21, 0x4c, 0xa9, 0x03, 0x00, 0x00,
}

func (this *Media) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 18)
	s = append(s, "&data_message.Media{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "PhoneNumber: "+fmt.Sprintf("%#v", this.PhoneNumber)+",\n")
	s = append(s, "FirstName: "+fmt.Sprintf("%#v", this.FirstName)+",\n")
	s = append(s, "LastName: "+fmt.Sprintf("%#v", this.LastName)+",\n")
	s = append(s, "UserId: "+fmt.Sprintf("%#v", this.UserId)+",\n")
	s = append(s, "Vcard: "+fmt.Sprintf("%#v", this.Vcard)+",\n")
	if this.Document != nil {
		s = append(s, "Document: "+fmt.Sprintf("%#v", this.Document)+",\n")
	}
	s = append(s, "Caption: "+fmt.Sprintf("%#v", this.Caption)+",\n")
	s = append(s, "TtlSeconds: "+fmt.Sprintf("%#v", this.TtlSeconds)+",\n")
	if this.Photo != nil {
		s = append(s, "Photo: "+fmt.Sprintf("%#v", this.Photo)+",\n")
	}
	if this.GeoPoint != nil {
		s = append(s, "GeoPoint: "+fmt.Sprintf("%#v", this.GeoPoint)+",\n")
	}
	s = append(s, "Period: "+fmt.Sprintf("%#v", this.Period)+",\n")
	s = append(s, "Heading: "+fmt.Sprintf("%#v", this.Heading)+",\n")
	s = append(s, "NotificationRadius: "+fmt.Sprintf("%#v", this.NotificationRadius)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMedia(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Media) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Media) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Media) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NotificationRadius != 0 {
		i = encodeVarintMedia(dAtA, i, uint64(m.NotificationRadius))
		i--
		dAtA[i] = 0x70
	}
	if m.Heading != 0 {
		i = encodeVarintMedia(dAtA, i, uint64(m.Heading))
		i--
		dAtA[i] = 0x68
	}
	if m.Period != 0 {
		i = encodeVarintMedia(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x60
	}
	if m.GeoPoint != nil {
		{
			size, err := m.GeoPoint.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMedia(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if m.Photo != nil {
		{
			size, err := m.Photo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMedia(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.TtlSeconds != 0 {
		i = encodeVarintMedia(dAtA, i, uint64(m.TtlSeconds))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Caption) > 0 {
		i -= len(m.Caption)
		copy(dAtA[i:], m.Caption)
		i = encodeVarintMedia(dAtA, i, uint64(len(m.Caption)))
		i--
		dAtA[i] = 0x42
	}
	if m.Document != nil {
		{
			size, err := m.Document.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMedia(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Vcard) > 0 {
		i -= len(m.Vcard)
		copy(dAtA[i:], m.Vcard)
		i = encodeVarintMedia(dAtA, i, uint64(len(m.Vcard)))
		i--
		dAtA[i] = 0x32
	}
	if m.UserId != 0 {
		i = encodeVarintMedia(dAtA, i, uint64(m.UserId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.LastName) > 0 {
		i -= len(m.LastName)
		copy(dAtA[i:], m.LastName)
		i = encodeVarintMedia(dAtA, i, uint64(len(m.LastName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.FirstName) > 0 {
		i -= len(m.FirstName)
		copy(dAtA[i:], m.FirstName)
		i = encodeVarintMedia(dAtA, i, uint64(len(m.FirstName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PhoneNumber) > 0 {
		i -= len(m.PhoneNumber)
		copy(dAtA[i:], m.PhoneNumber)
		i = encodeVarintMedia(dAtA, i, uint64(len(m.PhoneNumber)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintMedia(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMedia(dAtA []byte, offset int, v uint64) int {
	offset -= sovMedia(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Media) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovMedia(uint64(m.Type))
	}
	l = len(m.PhoneNumber)
	if l > 0 {
		n += 1 + l + sovMedia(uint64(l))
	}
	l = len(m.FirstName)
	if l > 0 {
		n += 1 + l + sovMedia(uint64(l))
	}
	l = len(m.LastName)
	if l > 0 {
		n += 1 + l + sovMedia(uint64(l))
	}
	if m.UserId != 0 {
		n += 1 + sovMedia(uint64(m.UserId))
	}
	l = len(m.Vcard)
	if l > 0 {
		n += 1 + l + sovMedia(uint64(l))
	}
	if m.Document != nil {
		l = m.Document.Size()
		n += 1 + l + sovMedia(uint64(l))
	}
	l = len(m.Caption)
	if l > 0 {
		n += 1 + l + sovMedia(uint64(l))
	}
	if m.TtlSeconds != 0 {
		n += 1 + sovMedia(uint64(m.TtlSeconds))
	}
	if m.Photo != nil {
		l = m.Photo.Size()
		n += 1 + l + sovMedia(uint64(l))
	}
	if m.GeoPoint != nil {
		l = m.GeoPoint.Size()
		n += 1 + l + sovMedia(uint64(l))
	}
	if m.Period != 0 {
		n += 1 + sovMedia(uint64(m.Period))
	}
	if m.Heading != 0 {
		n += 1 + sovMedia(uint64(m.Heading))
	}
	if m.NotificationRadius != 0 {
		n += 1 + sovMedia(uint64(m.NotificationRadius))
	}
	return n
}

func sovMedia(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMedia(x uint64) (n int) {
	return sovMedia(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Media) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMedia
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
			return fmt.Errorf("proto: Media: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Media: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PhoneNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
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
				return ErrInvalidLengthMedia
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMedia
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PhoneNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
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
				return ErrInvalidLengthMedia
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMedia
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
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
				return ErrInvalidLengthMedia
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMedia
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
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
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vcard", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
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
				return ErrInvalidLengthMedia
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMedia
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vcard = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Document", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
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
				return ErrInvalidLengthMedia
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMedia
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Document == nil {
				m.Document = &fs.Document{}
			}
			if err := m.Document.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Caption", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
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
				return ErrInvalidLengthMedia
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMedia
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Caption = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TtlSeconds", wireType)
			}
			m.TtlSeconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TtlSeconds |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Photo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
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
				return ErrInvalidLengthMedia
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMedia
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Photo == nil {
				m.Photo = &fs.Photo{}
			}
			if err := m.Photo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GeoPoint", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
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
				return ErrInvalidLengthMedia
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMedia
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GeoPoint == nil {
				m.GeoPoint = &fs.GeoPoint{}
			}
			if err := m.GeoPoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Heading", wireType)
			}
			m.Heading = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Heading |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotificationRadius", wireType)
			}
			m.NotificationRadius = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMedia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NotificationRadius |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMedia(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMedia
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMedia
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
func skipMedia(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMedia
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
					return 0, ErrIntOverflowMedia
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
					return 0, ErrIntOverflowMedia
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
				return 0, ErrInvalidLengthMedia
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMedia
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMedia
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMedia        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMedia          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMedia = fmt.Errorf("proto: unexpected end of group")
)
