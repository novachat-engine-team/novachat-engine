// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stickerpack.proto

package data_stickerset

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

type StickerPack struct {
	Emoticon  string         `protobuf:"bytes,1,opt,name=emoticon,proto3" json:"emoticon,omitempty" bson:"emoticon"`
	Documents []*fs.Document `protobuf:"bytes,2,rep,name=documents,proto3" json:"documents,omitempty" bson:"documents"`
}

func (m *StickerPack) Reset()         { *m = StickerPack{} }
func (m *StickerPack) String() string { return proto.CompactTextString(m) }
func (*StickerPack) ProtoMessage()    {}
func (*StickerPack) Descriptor() ([]byte, []int) {
	return fileDescriptor_b33a155f76835abf, []int{0}
}
func (m *StickerPack) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StickerPack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StickerPack.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StickerPack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StickerPack.Merge(m, src)
}
func (m *StickerPack) XXX_Size() int {
	return m.Size()
}
func (m *StickerPack) XXX_DiscardUnknown() {
	xxx_messageInfo_StickerPack.DiscardUnknown(m)
}

var xxx_messageInfo_StickerPack proto.InternalMessageInfo

func (m *StickerPack) GetEmoticon() string {
	if m != nil {
		return m.Emoticon
	}
	return ""
}

func (m *StickerPack) GetDocuments() []*fs.Document {
	if m != nil {
		return m.Documents
	}
	return nil
}

func init() {
	proto.RegisterType((*StickerPack)(nil), "data_stickerset.StickerPack")
}

func init() { proto.RegisterFile("stickerpack.proto", fileDescriptor_b33a155f76835abf) }

var fileDescriptor_b33a155f76835abf = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2e, 0xc9, 0x4c,
	0xce, 0x4e, 0x2d, 0x2a, 0x48, 0x4c, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4f,
	0x49, 0x2c, 0x49, 0x8c, 0x87, 0x8a, 0x17, 0xa7, 0x96, 0x48, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94,
	0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0xd5, 0x25, 0x95, 0xa6,
	0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0x2f, 0xc5, 0x97, 0x92, 0x9f, 0x5c, 0x9a, 0x9b, 0x9a,
	0x57, 0x02, 0xe1, 0x2b, 0x4d, 0x66, 0xe4, 0xe2, 0x0e, 0x86, 0x98, 0x16, 0x90, 0x98, 0x9c, 0x2d,
	0xa4, 0xcf, 0xc5, 0x91, 0x9a, 0x9b, 0x5f, 0x92, 0x99, 0x9c, 0x9f, 0x27, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0xe9, 0x24, 0xfc, 0xe9, 0x9e, 0x3c, 0x7f, 0x52, 0x71, 0x7e, 0x9e, 0x95, 0x12, 0x4c, 0x46,
	0x29, 0x08, 0xae, 0x48, 0xc8, 0x95, 0x8b, 0x13, 0x66, 0x64, 0xb1, 0x04, 0x93, 0x02, 0xb3, 0x06,
	0xb7, 0x91, 0xa0, 0x1e, 0xd8, 0x91, 0x69, 0xc5, 0x7a, 0x2e, 0x50, 0x19, 0x27, 0x91, 0x4f, 0xf7,
	0xe4, 0x05, 0x20, 0x86, 0xc0, 0x55, 0x2b, 0x05, 0x21, 0x74, 0x5a, 0xf1, 0x5c, 0x58, 0x28, 0xcf,
	0x30, 0x61, 0x91, 0x3c, 0xc3, 0x8c, 0x45, 0xf2, 0x0c, 0x4e, 0x8a, 0x1f, 0x1e, 0xca, 0x31, 0x9e,
	0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x51, 0xe8, 0xfe, 0x4e,
	0x62, 0x03, 0xbb, 0xdf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x15, 0x32, 0xd7, 0x24, 0x01,
	0x00, 0x00,
}

func (this *StickerPack) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&data_stickerset.StickerPack{")
	s = append(s, "Emoticon: "+fmt.Sprintf("%#v", this.Emoticon)+",\n")
	if this.Documents != nil {
		s = append(s, "Documents: "+fmt.Sprintf("%#v", this.Documents)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringStickerpack(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *StickerPack) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StickerPack) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StickerPack) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Documents) > 0 {
		for iNdEx := len(m.Documents) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Documents[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStickerpack(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Emoticon) > 0 {
		i -= len(m.Emoticon)
		copy(dAtA[i:], m.Emoticon)
		i = encodeVarintStickerpack(dAtA, i, uint64(len(m.Emoticon)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStickerpack(dAtA []byte, offset int, v uint64) int {
	offset -= sovStickerpack(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StickerPack) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Emoticon)
	if l > 0 {
		n += 1 + l + sovStickerpack(uint64(l))
	}
	if len(m.Documents) > 0 {
		for _, e := range m.Documents {
			l = e.Size()
			n += 1 + l + sovStickerpack(uint64(l))
		}
	}
	return n
}

func sovStickerpack(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStickerpack(x uint64) (n int) {
	return sovStickerpack(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StickerPack) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStickerpack
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
			return fmt.Errorf("proto: StickerPack: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StickerPack: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Emoticon", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStickerpack
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
				return ErrInvalidLengthStickerpack
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStickerpack
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Emoticon = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Documents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStickerpack
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
				return ErrInvalidLengthStickerpack
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStickerpack
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Documents = append(m.Documents, &fs.Document{})
			if err := m.Documents[len(m.Documents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStickerpack(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStickerpack
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStickerpack
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
func skipStickerpack(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStickerpack
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
					return 0, ErrIntOverflowStickerpack
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
					return 0, ErrIntOverflowStickerpack
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
				return 0, ErrInvalidLengthStickerpack
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStickerpack
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStickerpack
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStickerpack        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStickerpack          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStickerpack = fmt.Errorf("proto: unexpected end of group")
)