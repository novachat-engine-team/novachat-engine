// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc.proto

package rpc

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	metadata "novachat_engine/pkg/rpc/metadata"
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

type RpcData struct {
	Data *types.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *RpcData) Reset()         { *m = RpcData{} }
func (m *RpcData) String() string { return proto.CompactTextString(m) }
func (*RpcData) ProtoMessage()    {}
func (*RpcData) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}
func (m *RpcData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RpcData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RpcData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RpcData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcData.Merge(m, src)
}
func (m *RpcData) XXX_Size() int {
	return m.Size()
}
func (m *RpcData) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcData.DiscardUnknown(m)
}

var xxx_messageInfo_RpcData proto.InternalMessageInfo

func (m *RpcData) GetData() *types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

type RpcStreamData struct {
	Data *types.Any            `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Md   *metadata.RpcMetaData `protobuf:"bytes,2,opt,name=md,proto3" json:"md,omitempty"`
}

func (m *RpcStreamData) Reset()         { *m = RpcStreamData{} }
func (m *RpcStreamData) String() string { return proto.CompactTextString(m) }
func (*RpcStreamData) ProtoMessage()    {}
func (*RpcStreamData) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}
func (m *RpcStreamData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RpcStreamData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RpcStreamData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RpcStreamData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcStreamData.Merge(m, src)
}
func (m *RpcStreamData) XXX_Size() int {
	return m.Size()
}
func (m *RpcStreamData) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcStreamData.DiscardUnknown(m)
}

var xxx_messageInfo_RpcStreamData proto.InternalMessageInfo

func (m *RpcStreamData) GetData() *types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *RpcStreamData) GetMd() *metadata.RpcMetaData {
	if m != nil {
		return m.Md
	}
	return nil
}

func init() {
	proto.RegisterType((*RpcData)(nil), "rpc.RpcData")
	proto.RegisterType((*RpcStreamData)(nil), "rpc.RpcStreamData")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2a, 0x48, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x96, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xcb, 0x25, 0x95,
	0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0x23, 0x25, 0x99, 0x9e, 0x9f, 0x9f, 0x9e, 0x93,
	0x8a, 0x50, 0x95, 0x98, 0x57, 0x09, 0x95, 0x52, 0x28, 0xc8, 0x4e, 0xd7, 0x2f, 0x2a, 0x48, 0xd6,
	0xcf, 0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0x04, 0x33, 0xe2, 0x41, 0x2c, 0x88, 0x0a, 0x25,
	0x47, 0x2e, 0xf6, 0xa0, 0x82, 0x64, 0x97, 0xc4, 0x92, 0x44, 0x21, 0x0d, 0x2e, 0x16, 0x90, 0x84,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x88, 0x1e, 0xc4, 0x58, 0x3d, 0x98, 0xb1, 0x7a, 0x8e,
	0x79, 0x95, 0x41, 0x60, 0x15, 0x56, 0x3c, 0x17, 0x16, 0xca, 0x33, 0x4c, 0x58, 0x24, 0xcf, 0x30,
	0x63, 0x91, 0x3c, 0x83, 0x52, 0x01, 0x17, 0x6f, 0x50, 0x41, 0x72, 0x70, 0x49, 0x51, 0x6a, 0x62,
	0x2e, 0x69, 0x06, 0x09, 0xa9, 0x73, 0x31, 0xe5, 0xa6, 0x48, 0x30, 0x81, 0xd5, 0x89, 0xeb, 0xc1,
	0x1c, 0xa9, 0x57, 0x54, 0x90, 0x1c, 0x0f, 0x77, 0x68, 0x10, 0x53, 0x6e, 0x0a, 0xaa, 0x8d, 0x46,
	0x8e, 0x5c, 0x5c, 0x20, 0x1b, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x8c, 0xb9, 0xb8, 0x83,
	0x52, 0x93, 0x53, 0x33, 0xcb, 0x52, 0xc1, 0xb6, 0xf3, 0x80, 0xb4, 0xeb, 0x41, 0x3d, 0x25, 0x85,
	0xd5, 0x76, 0x25, 0x06, 0x23, 0x7f, 0x2e, 0x01, 0xb8, 0xa3, 0x61, 0x06, 0x59, 0xa3, 0x1a, 0x24,
	0x04, 0x33, 0x08, 0xe1, 0x35, 0x29, 0x2c, 0x62, 0x4a, 0x0c, 0x1a, 0x8c, 0x06, 0x8c, 0x4e, 0xa2,
	0x1f, 0x1e, 0xca, 0x31, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72,
	0x8c, 0x51, 0xa0, 0xb8, 0x4c, 0x62, 0x03, 0xdb, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa2,
	0x37, 0xe7, 0x3d, 0xe4, 0x01, 0x00, 0x00,
}

func (this *RpcData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&rpc.RpcData{")
	if this.Data != nil {
		s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *RpcStreamData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&rpc.RpcStreamData{")
	if this.Data != nil {
		s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	}
	if this.Md != nil {
		s = append(s, "Md: "+fmt.Sprintf("%#v", this.Md)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringRpc(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RpcServiceClient is the client API for RpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RpcServiceClient interface {
	ReceiveData(ctx context.Context, in *RpcData, opts ...grpc.CallOption) (*types.Any, error)
}

type rpcServiceClient struct {
	cc *grpc.ClientConn
}

func NewRpcServiceClient(cc *grpc.ClientConn) RpcServiceClient {
	return &rpcServiceClient{cc}
}

func (c *rpcServiceClient) ReceiveData(ctx context.Context, in *RpcData, opts ...grpc.CallOption) (*types.Any, error) {
	out := new(types.Any)
	err := c.cc.Invoke(ctx, "/rpc.RpcService/ReceiveData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServiceServer is the server API for RpcService service.
type RpcServiceServer interface {
	ReceiveData(context.Context, *RpcData) (*types.Any, error)
}

// UnimplementedRpcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRpcServiceServer struct {
}

func (*UnimplementedRpcServiceServer) ReceiveData(ctx context.Context, req *RpcData) (*types.Any, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveData not implemented")
}

func RegisterRpcServiceServer(s *grpc.Server, srv RpcServiceServer) {
	s.RegisterService(&_RpcService_serviceDesc, srv)
}

func _RpcService_ReceiveData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpcData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).ReceiveData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.RpcService/ReceiveData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).ReceiveData(ctx, req.(*RpcData))
	}
	return interceptor(ctx, in, info, handler)
}

var _RpcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.RpcService",
	HandlerType: (*RpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveData",
			Handler:    _RpcService_ReceiveData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

var RpcService_serviceDesc = map[string][]interface{}{
	"RpcData": {
		"/rpc.RpcService/ReceiveData",
		func() interface{} { return new(types.Any) },
		"ReceiveData",
	},
}

// RpcStreamServiceClient is the client API for RpcStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RpcStreamServiceClient interface {
	ReceiveData(ctx context.Context, opts ...grpc.CallOption) (RpcStreamService_ReceiveDataClient, error)
}

type rpcStreamServiceClient struct {
	cc *grpc.ClientConn
}

func NewRpcStreamServiceClient(cc *grpc.ClientConn) RpcStreamServiceClient {
	return &rpcStreamServiceClient{cc}
}

func (c *rpcStreamServiceClient) ReceiveData(ctx context.Context, opts ...grpc.CallOption) (RpcStreamService_ReceiveDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RpcStreamService_serviceDesc.Streams[0], "/rpc.RpcStreamService/ReceiveData", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcStreamServiceReceiveDataClient{stream}
	return x, nil
}

type RpcStreamService_ReceiveDataClient interface {
	Send(*RpcStreamData) error
	Recv() (*RpcStreamData, error)
	grpc.ClientStream
}

type rpcStreamServiceReceiveDataClient struct {
	grpc.ClientStream
}

func (x *rpcStreamServiceReceiveDataClient) Send(m *RpcStreamData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rpcStreamServiceReceiveDataClient) Recv() (*RpcStreamData, error) {
	m := new(RpcStreamData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RpcStreamServiceServer is the server API for RpcStreamService service.
type RpcStreamServiceServer interface {
	ReceiveData(RpcStreamService_ReceiveDataServer) error
}

// UnimplementedRpcStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRpcStreamServiceServer struct {
}

func (*UnimplementedRpcStreamServiceServer) ReceiveData(srv RpcStreamService_ReceiveDataServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveData not implemented")
}

func RegisterRpcStreamServiceServer(s *grpc.Server, srv RpcStreamServiceServer) {
	s.RegisterService(&_RpcStreamService_serviceDesc, srv)
}

func _RpcStreamService_ReceiveData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RpcStreamServiceServer).ReceiveData(&rpcStreamServiceReceiveDataServer{stream})
}

type RpcStreamService_ReceiveDataServer interface {
	Send(*RpcStreamData) error
	Recv() (*RpcStreamData, error)
	grpc.ServerStream
}

type rpcStreamServiceReceiveDataServer struct {
	grpc.ServerStream
}

func (x *rpcStreamServiceReceiveDataServer) Send(m *RpcStreamData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rpcStreamServiceReceiveDataServer) Recv() (*RpcStreamData, error) {
	m := new(RpcStreamData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RpcStreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.RpcStreamService",
	HandlerType: (*RpcStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveData",
			Handler:       _RpcStreamService_ReceiveData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rpc.proto",
}

var RpcStreamService_serviceDesc = map[string][]interface{}{
	"RpcStreamData": {
		"/rpc.RpcStreamService/ReceiveData",
		func() interface{} { return new(RpcStreamData) },
		"ReceiveData",
	},
}

func (m *RpcData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RpcData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RpcData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RpcStreamData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RpcStreamData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RpcStreamData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Md != nil {
		{
			size, err := m.Md.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRpc(dAtA []byte, offset int, v uint64) int {
	offset -= sovRpc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RpcData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *RpcStreamData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Md != nil {
		l = m.Md.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func sovRpc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRpc(x uint64) (n int) {
	return sovRpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RpcData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: RpcData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RpcData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRpc
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
func (m *RpcStreamData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: RpcStreamData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RpcStreamData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Md", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Md == nil {
				m.Md = &metadata.RpcMetaData{}
			}
			if err := m.Md.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRpc
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
func skipRpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRpc
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
					return 0, ErrIntOverflowRpc
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
					return 0, ErrIntOverflowRpc
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
				return 0, ErrInvalidLengthRpc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRpc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRpc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRpc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRpc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRpc = fmt.Errorf("proto: unexpected end of group")
)
