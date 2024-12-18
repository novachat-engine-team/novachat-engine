// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: document.proto

package data_fs

import (
	encoding_binary "encoding/binary"
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

type Document struct {
	VolumeId      int64         `protobuf:"varint,1,opt,name=volume_id,json=volumeId,proto3" json:"volume_id,omitempty" bson:"_id"`
	LocalId       int32         `protobuf:"varint,2,opt,name=local_id,json=localId,proto3" json:"local_id,omitempty" bson:"local_id"`
	MimeType      string        `protobuf:"bytes,3,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty" bson:"mime_type"`
	Size_         int32         `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty" bson:"size"`
	Thumbs        []*Photo      `protobuf:"bytes,5,rep,name=thumbs,proto3" json:"thumbs,omitempty" bson:"thumbs"`
	Attributes    []*Attributes `protobuf:"bytes,7,rep,name=attributes,proto3" json:"attributes,omitempty" bson:"attributes"`
	Date          int32         `protobuf:"varint,8,opt,name=date,proto3" json:"date,omitempty" bson:"date"`
	FileType      int32         `protobuf:"varint,9,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty" bson:"file_type"`
	FilePath      string        `protobuf:"bytes,10,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty" bson:"file_path"`
	AccessHash    int64         `protobuf:"varint,11,opt,name=access_hash,json=accessHash,proto3" json:"access_hash,omitempty" bson:"access_hash"`
	FileReference []byte        `protobuf:"bytes,12,opt,name=file_reference,json=fileReference,proto3" json:"file_reference,omitempty" bson:"file_reference"`
	Version       int32         `protobuf:"varint,13,opt,name=version,proto3" json:"version,omitempty" bson:"version"`
	VideoStartTs  float64       `protobuf:"fixed64,14,opt,name=video_start_ts,json=videoStartTs,proto3" json:"video_start_ts,omitempty" bson:"video_start_ts"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d2790a4091b3173, []int{0}
}
func (m *Document) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Document.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(m, src)
}
func (m *Document) XXX_Size() int {
	return m.Size()
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetVolumeId() int64 {
	if m != nil {
		return m.VolumeId
	}
	return 0
}

func (m *Document) GetLocalId() int32 {
	if m != nil {
		return m.LocalId
	}
	return 0
}

func (m *Document) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *Document) GetSize_() int32 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *Document) GetThumbs() []*Photo {
	if m != nil {
		return m.Thumbs
	}
	return nil
}

func (m *Document) GetAttributes() []*Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Document) GetDate() int32 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *Document) GetFileType() int32 {
	if m != nil {
		return m.FileType
	}
	return 0
}

func (m *Document) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *Document) GetAccessHash() int64 {
	if m != nil {
		return m.AccessHash
	}
	return 0
}

func (m *Document) GetFileReference() []byte {
	if m != nil {
		return m.FileReference
	}
	return nil
}

func (m *Document) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Document) GetVideoStartTs() float64 {
	if m != nil {
		return m.VideoStartTs
	}
	return 0
}

func init() {
	proto.RegisterType((*Document)(nil), "data_fs.Document")
}

func init() { proto.RegisterFile("document.proto", fileDescriptor_9d2790a4091b3173) }

var fileDescriptor_9d2790a4091b3173 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x53, 0xcf, 0x6e, 0xd3, 0x30,
	0x1c, 0x6e, 0xe8, 0xba, 0xa6, 0x6e, 0x9b, 0x6d, 0x1e, 0x43, 0xd9, 0x0e, 0x75, 0x64, 0x2e, 0x91,
	0x80, 0x4e, 0xc0, 0x01, 0xb1, 0x0b, 0x10, 0x71, 0x60, 0x9c, 0x26, 0xb3, 0x13, 0x97, 0xc8, 0x49,
	0xdc, 0x26, 0x52, 0x53, 0x57, 0xb1, 0x53, 0x69, 0x3c, 0x05, 0x47, 0xae, 0xec, 0x69, 0xe0, 0xb6,
	0x27, 0x88, 0xa0, 0x4f, 0x80, 0xf2, 0x04, 0xc8, 0xce, 0x9f, 0x76, 0xbb, 0xfd, 0xbe, 0xdf, 0xf7,
	0x7d, 0xfe, 0xf9, 0xf7, 0xc5, 0x01, 0x56, 0xc4, 0xc3, 0x3c, 0x65, 0x4b, 0x39, 0x5d, 0x65, 0x5c,
	0x72, 0xd8, 0x8f, 0xa8, 0xa4, 0xfe, 0x4c, 0x9c, 0xbd, 0x98, 0x27, 0x32, 0xce, 0x83, 0x69, 0xc8,
	0xd3, 0xf3, 0x39, 0x9f, 0xf3, 0x73, 0xcd, 0x07, 0xf9, 0x4c, 0x23, 0x0d, 0x74, 0x55, 0xf9, 0xce,
	0x86, 0xab, 0x98, 0xcb, 0x16, 0xac, 0x93, 0x88, 0x35, 0xe0, 0x90, 0x4a, 0x99, 0x25, 0x41, 0x2e,
	0x99, 0xa8, 0x3a, 0xf8, 0x77, 0x0f, 0x98, 0x1f, 0xeb, 0xb1, 0xf0, 0x19, 0x18, 0xac, 0xf9, 0x22,
	0x4f, 0x99, 0x9f, 0x44, 0xb6, 0xe1, 0x18, 0x6e, 0xd7, 0xb3, 0xca, 0x02, 0x81, 0x40, 0xf0, 0xe5,
	0x05, 0xf6, 0x93, 0x08, 0x13, 0xb3, 0x12, 0x5c, 0x46, 0x70, 0x0a, 0xcc, 0x05, 0x0f, 0xe9, 0x42,
	0x69, 0x1f, 0x39, 0x86, 0xdb, 0xf3, 0x8e, 0xcb, 0x02, 0x1d, 0x54, 0xda, 0x86, 0xc1, 0xa4, 0xaf,
	0xcb, 0xcb, 0x08, 0xbe, 0x04, 0x83, 0x34, 0x49, 0x99, 0x2f, 0x6f, 0x56, 0xcc, 0xee, 0x3a, 0x86,
	0x3b, 0xf0, 0x1e, 0x97, 0x05, 0x3a, 0xac, 0x0c, 0x2d, 0x85, 0x89, 0xa9, 0xea, 0xeb, 0x9b, 0x15,
	0x83, 0x4f, 0xc1, 0x9e, 0x48, 0xbe, 0x31, 0x7b, 0x4f, 0x1f, 0x7f, 0x50, 0x16, 0x68, 0x58, 0xa9,
	0x55, 0x17, 0x13, 0x4d, 0xc2, 0xb7, 0x60, 0x5f, 0xc6, 0x79, 0x1a, 0x08, 0xbb, 0xe7, 0x74, 0xdd,
	0xe1, 0x2b, 0x6b, 0x5a, 0xc7, 0x36, 0xbd, 0x52, 0x31, 0x78, 0x47, 0x65, 0x81, 0xc6, 0x95, 0xad,
	0xd2, 0x61, 0x52, 0x1b, 0xe0, 0x67, 0x00, 0xb6, 0x81, 0xd8, 0x7d, 0x6d, 0x3f, 0x6e, 0xed, 0x1f,
	0x5a, 0xca, 0x3b, 0x29, 0x0b, 0x74, 0x54, 0x9d, 0xb1, 0x35, 0x60, 0xb2, 0xe3, 0x56, 0x77, 0x8d,
	0xa8, 0x64, 0xb6, 0xf9, 0xf0, 0xae, 0xaa, 0x8b, 0x89, 0x26, 0x55, 0x06, 0xb3, 0x64, 0x51, 0x67,
	0x30, 0xd0, 0xca, 0x9d, 0x0c, 0x5a, 0x0a, 0x13, 0x53, 0xd5, 0x3a, 0x83, 0xc6, 0xb2, 0xa2, 0x32,
	0xb6, 0xc1, 0xc3, 0xd8, 0x5a, 0xaa, 0xb6, 0x5c, 0x51, 0x19, 0xc3, 0x37, 0x60, 0x48, 0xc3, 0x90,
	0x09, 0xe1, 0xc7, 0x54, 0xc4, 0xf6, 0x50, 0x7f, 0xc8, 0x27, 0x65, 0x81, 0x60, 0xbd, 0xc2, 0x96,
	0x54, 0x3b, 0x68, 0xf4, 0x89, 0x8a, 0x18, 0xbe, 0x07, 0x96, 0x3e, 0x30, 0x63, 0x33, 0x96, 0xb1,
	0x65, 0xc8, 0xec, 0x91, 0x63, 0xb8, 0x23, 0xef, 0xb4, 0x2c, 0xd0, 0xc9, 0xce, 0xc0, 0x96, 0xc7,
	0x64, 0xac, 0x1a, 0xa4, 0xc1, 0xf0, 0x39, 0xe8, 0xaf, 0x59, 0x26, 0x12, 0xbe, 0xb4, 0xc7, 0x7a,
	0x3d, 0x58, 0x16, 0xc8, 0xaa, 0xac, 0x35, 0x81, 0x49, 0x23, 0x81, 0xef, 0x80, 0xa5, 0x5f, 0xa7,
	0x2f, 0x24, 0xcd, 0xa4, 0x2f, 0x85, 0x6d, 0x39, 0x86, 0x6b, 0xec, 0xce, 0xbb, 0xcf, 0x63, 0x32,
	0xd2, 0x8d, 0x2f, 0x0a, 0x5f, 0x8b, 0x8b, 0xd1, 0xdd, 0x4f, 0xd4, 0xf9, 0x7e, 0x8b, 0x3a, 0x3f,
	0x6e, 0x51, 0xc7, 0x3b, 0xfd, 0xf7, 0x77, 0x62, 0xfc, 0xda, 0x4c, 0x8c, 0xbb, 0xcd, 0xc4, 0xf8,
	0xb3, 0x99, 0x18, 0x5f, 0x9b, 0x3f, 0x28, 0xd8, 0xd7, 0xaf, 0xfd, 0xf5, 0xff, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xfa, 0xc8, 0x46, 0x8c, 0x63, 0x03, 0x00, 0x00,
}

func (this *Document) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 17)
	s = append(s, "&data_fs.Document{")
	s = append(s, "VolumeId: "+fmt.Sprintf("%#v", this.VolumeId)+",\n")
	s = append(s, "LocalId: "+fmt.Sprintf("%#v", this.LocalId)+",\n")
	s = append(s, "MimeType: "+fmt.Sprintf("%#v", this.MimeType)+",\n")
	s = append(s, "Size_: "+fmt.Sprintf("%#v", this.Size_)+",\n")
	if this.Thumbs != nil {
		s = append(s, "Thumbs: "+fmt.Sprintf("%#v", this.Thumbs)+",\n")
	}
	if this.Attributes != nil {
		s = append(s, "Attributes: "+fmt.Sprintf("%#v", this.Attributes)+",\n")
	}
	s = append(s, "Date: "+fmt.Sprintf("%#v", this.Date)+",\n")
	s = append(s, "FileType: "+fmt.Sprintf("%#v", this.FileType)+",\n")
	s = append(s, "FilePath: "+fmt.Sprintf("%#v", this.FilePath)+",\n")
	s = append(s, "AccessHash: "+fmt.Sprintf("%#v", this.AccessHash)+",\n")
	s = append(s, "FileReference: "+fmt.Sprintf("%#v", this.FileReference)+",\n")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "VideoStartTs: "+fmt.Sprintf("%#v", this.VideoStartTs)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringDocument(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Document) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Document) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Document) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VideoStartTs != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.VideoStartTs))))
		i--
		dAtA[i] = 0x71
	}
	if m.Version != 0 {
		i = encodeVarintDocument(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x68
	}
	if len(m.FileReference) > 0 {
		i -= len(m.FileReference)
		copy(dAtA[i:], m.FileReference)
		i = encodeVarintDocument(dAtA, i, uint64(len(m.FileReference)))
		i--
		dAtA[i] = 0x62
	}
	if m.AccessHash != 0 {
		i = encodeVarintDocument(dAtA, i, uint64(m.AccessHash))
		i--
		dAtA[i] = 0x58
	}
	if len(m.FilePath) > 0 {
		i -= len(m.FilePath)
		copy(dAtA[i:], m.FilePath)
		i = encodeVarintDocument(dAtA, i, uint64(len(m.FilePath)))
		i--
		dAtA[i] = 0x52
	}
	if m.FileType != 0 {
		i = encodeVarintDocument(dAtA, i, uint64(m.FileType))
		i--
		dAtA[i] = 0x48
	}
	if m.Date != 0 {
		i = encodeVarintDocument(dAtA, i, uint64(m.Date))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDocument(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Thumbs) > 0 {
		for iNdEx := len(m.Thumbs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Thumbs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDocument(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Size_ != 0 {
		i = encodeVarintDocument(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x20
	}
	if len(m.MimeType) > 0 {
		i -= len(m.MimeType)
		copy(dAtA[i:], m.MimeType)
		i = encodeVarintDocument(dAtA, i, uint64(len(m.MimeType)))
		i--
		dAtA[i] = 0x1a
	}
	if m.LocalId != 0 {
		i = encodeVarintDocument(dAtA, i, uint64(m.LocalId))
		i--
		dAtA[i] = 0x10
	}
	if m.VolumeId != 0 {
		i = encodeVarintDocument(dAtA, i, uint64(m.VolumeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDocument(dAtA []byte, offset int, v uint64) int {
	offset -= sovDocument(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Document) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VolumeId != 0 {
		n += 1 + sovDocument(uint64(m.VolumeId))
	}
	if m.LocalId != 0 {
		n += 1 + sovDocument(uint64(m.LocalId))
	}
	l = len(m.MimeType)
	if l > 0 {
		n += 1 + l + sovDocument(uint64(l))
	}
	if m.Size_ != 0 {
		n += 1 + sovDocument(uint64(m.Size_))
	}
	if len(m.Thumbs) > 0 {
		for _, e := range m.Thumbs {
			l = e.Size()
			n += 1 + l + sovDocument(uint64(l))
		}
	}
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovDocument(uint64(l))
		}
	}
	if m.Date != 0 {
		n += 1 + sovDocument(uint64(m.Date))
	}
	if m.FileType != 0 {
		n += 1 + sovDocument(uint64(m.FileType))
	}
	l = len(m.FilePath)
	if l > 0 {
		n += 1 + l + sovDocument(uint64(l))
	}
	if m.AccessHash != 0 {
		n += 1 + sovDocument(uint64(m.AccessHash))
	}
	l = len(m.FileReference)
	if l > 0 {
		n += 1 + l + sovDocument(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovDocument(uint64(m.Version))
	}
	if m.VideoStartTs != 0 {
		n += 9
	}
	return n
}

func sovDocument(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDocument(x uint64) (n int) {
	return sovDocument(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Document) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDocument
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
			return fmt.Errorf("proto: Document: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Document: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VolumeId", wireType)
			}
			m.VolumeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VolumeId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalId", wireType)
			}
			m.LocalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LocalId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MimeType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
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
				return ErrInvalidLengthDocument
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MimeType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Thumbs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
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
				return ErrInvalidLengthDocument
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Thumbs = append(m.Thumbs, &Photo{})
			if err := m.Thumbs[len(m.Thumbs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
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
				return ErrInvalidLengthDocument
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, &Attributes{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			m.Date = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
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
				return fmt.Errorf("proto: wrong wireType = %d for field FileType", wireType)
			}
			m.FileType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FileType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilePath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
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
				return ErrInvalidLengthDocument
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FilePath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessHash", wireType)
			}
			m.AccessHash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccessHash |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileReference", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDocument
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileReference = append(m.FileReference[:0], dAtA[iNdEx:postIndex]...)
			if m.FileReference == nil {
				m.FileReference = []byte{}
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
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
		case 14:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field VideoStartTs", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.VideoStartTs = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipDocument(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDocument
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDocument
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
func skipDocument(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDocument
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
					return 0, ErrIntOverflowDocument
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
					return 0, ErrIntOverflowDocument
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
				return 0, ErrInvalidLengthDocument
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDocument
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDocument
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDocument        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDocument          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDocument = fmt.Errorf("proto: unexpected end of group")
)
