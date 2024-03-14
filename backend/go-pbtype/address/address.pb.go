// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: address/address.proto

package address

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Address struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Province             string   `protobuf:"bytes,7,opt,name=province,proto3" json:"province,omitempty"`
	DetailAddress        string   `protobuf:"bytes,8,opt,name=detail_address,json=detailAddress,proto3" json:"detail_address,omitempty"`
	Type                 string   `protobuf:"bytes,9,opt,name=type,proto3" json:"type,omitempty"`
	CreatedAt            string   `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_980ab9fb446d0e03, []int{0}
}
func (m *Address) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Address.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return m.Size()
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Address) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *Address) GetDetailAddress() string {
	if m != nil {
		return m.DetailAddress
	}
	return ""
}

func (m *Address) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Address) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Address) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (*Address) XXX_MessageName() string {
	return "address.Address"
}

type OneAddressRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OneAddressRequest) Reset()         { *m = OneAddressRequest{} }
func (m *OneAddressRequest) String() string { return proto.CompactTextString(m) }
func (*OneAddressRequest) ProtoMessage()    {}
func (*OneAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_980ab9fb446d0e03, []int{1}
}
func (m *OneAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OneAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OneAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OneAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OneAddressRequest.Merge(m, src)
}
func (m *OneAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *OneAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OneAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OneAddressRequest proto.InternalMessageInfo

func (m *OneAddressRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (*OneAddressRequest) XXX_MessageName() string {
	return "address.OneAddressRequest"
}
func init() {
	proto.RegisterType((*Address)(nil), "address.Address")
	golang_proto.RegisterType((*Address)(nil), "address.Address")
	proto.RegisterType((*OneAddressRequest)(nil), "address.OneAddressRequest")
	golang_proto.RegisterType((*OneAddressRequest)(nil), "address.OneAddressRequest")
}

func init() { proto.RegisterFile("address/address.proto", fileDescriptor_980ab9fb446d0e03) }
func init() { golang_proto.RegisterFile("address/address.proto", fileDescriptor_980ab9fb446d0e03) }

var fileDescriptor_980ab9fb446d0e03 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6b, 0xdb, 0x30,
	0x18, 0xc6, 0x51, 0xb6, 0xe5, 0x8f, 0x46, 0xc2, 0xa6, 0x2d, 0x60, 0xcc, 0x66, 0x82, 0xc7, 0x20,
	0x0c, 0x12, 0xc1, 0x76, 0xda, 0x6e, 0xde, 0x60, 0xdb, 0x2d, 0x90, 0xd2, 0x4b, 0x2f, 0x41, 0xb6,
	0xde, 0x38, 0x6a, 0x5c, 0xc9, 0x75, 0xe4, 0x40, 0x53, 0x7a, 0xe9, 0x57, 0xe8, 0xf7, 0xe8, 0x67,
	0xe8, 0x31, 0xc7, 0x42, 0xbf, 0x40, 0x49, 0xda, 0xef, 0x51, 0x22, 0xcb, 0xa1, 0x34, 0x14, 0x7a,
	0xb2, 0xde, 0xe7, 0x79, 0xfc, 0xe3, 0xb1, 0xf5, 0xe2, 0x36, 0xe3, 0x3c, 0x83, 0xd9, 0x8c, 0xda,
	0x67, 0x3f, 0xcd, 0x94, 0x56, 0xa4, 0x66, 0x47, 0xb7, 0x17, 0x0b, 0x3d, 0xc9, 0xc3, 0x7e, 0xa4,
	0x8e, 0x68, 0xac, 0x62, 0x45, 0x8d, 0x1f, 0xe6, 0x63, 0x33, 0x99, 0xc1, 0x9c, 0x8a, 0xf7, 0xdc,
	0x4f, 0xb1, 0x52, 0x71, 0x02, 0x94, 0xa5, 0x82, 0x32, 0x29, 0x95, 0x66, 0x5a, 0x28, 0x69, 0xa9,
	0xfe, 0x25, 0xc2, 0xb5, 0xa0, 0x00, 0x93, 0x16, 0xae, 0x08, 0xee, 0xa0, 0x0e, 0xea, 0xbe, 0x19,
	0x56, 0x04, 0x27, 0x2e, 0xae, 0xa7, 0x99, 0x9a, 0x0b, 0x19, 0x81, 0x53, 0xeb, 0xa0, 0x6e, 0x63,
	0xb8, 0x9d, 0xc9, 0x57, 0xdc, 0xe2, 0xa0, 0x99, 0x48, 0x46, 0xb6, 0x96, 0x53, 0x37, 0x89, 0x66,
	0xa1, 0x96, 0x48, 0x82, 0x5f, 0xeb, 0x93, 0x14, 0x9c, 0x86, 0x31, 0xcd, 0x99, 0x7c, 0xc6, 0x38,
	0xca, 0x80, 0x69, 0xe0, 0x23, 0xa6, 0x9d, 0x8f, 0xc6, 0x69, 0x58, 0x25, 0xd0, 0x1b, 0x3b, 0x4f,
	0x79, 0x69, 0xb7, 0x0b, 0xdb, 0x2a, 0x81, 0xf6, 0xbf, 0xe0, 0xf7, 0x03, 0x09, 0x96, 0x3f, 0x84,
	0xe3, 0x1c, 0x66, 0xfa, 0x69, 0xf3, 0xef, 0xf7, 0x08, 0xb7, 0x6c, 0x64, 0x0f, 0xb2, 0xb9, 0x88,
	0x80, 0xfc, 0xc5, 0xd5, 0x7d, 0x03, 0x21, 0xef, 0xfa, 0xe5, 0x8f, 0xb5, 0x11, 0x77, 0x47, 0xf1,
	0x9d, 0xf3, 0x9b, 0xbb, 0x8b, 0x0a, 0x71, 0x9b, 0xe5, 0x25, 0xd0, 0x53, 0xc1, 0xcf, 0x7e, 0xa1,
	0x6f, 0x24, 0xc0, 0xd5, 0x3f, 0xa6, 0xeb, 0x8b, 0x38, 0x1f, 0x0c, 0xa7, 0xe9, 0xd7, 0x4b, 0xce,
	0x06, 0xf1, 0x1f, 0xbf, 0xfa, 0x07, 0x9a, 0xb8, 0xdb, 0xf4, 0xce, 0x07, 0x3d, 0x4f, 0x22, 0x6f,
	0xe9, 0x44, 0x69, 0x48, 0x4c, 0x9f, 0xdf, 0x83, 0xe5, 0xca, 0x43, 0xd7, 0x2b, 0x0f, 0xdd, 0xae,
	0x3c, 0x74, 0xb5, 0xf6, 0xd0, 0x72, 0xed, 0xa1, 0x83, 0x9f, 0x8f, 0x16, 0x44, 0xc2, 0x54, 0x2c,
	0x16, 0x74, 0x2c, 0x24, 0x4b, 0x7a, 0x69, 0xa6, 0x0e, 0x21, 0xd2, 0x34, 0x64, 0xd1, 0x14, 0x24,
	0xa7, 0xb1, 0xea, 0xa5, 0xe1, 0xe6, 0x4e, 0xca, 0x76, 0x61, 0xd5, 0x6c, 0xc5, 0x8f, 0x87, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x42, 0xd6, 0x2b, 0x06, 0x84, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddressServiceClient is the client API for AddressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddressServiceClient interface {
	Update(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error)
	Create(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error)
	Get(ctx context.Context, in *OneAddressRequest, opts ...grpc.CallOption) (*Address, error)
}

type addressServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddressServiceClient(cc *grpc.ClientConn) AddressServiceClient {
	return &addressServiceClient{cc}
}

func (c *addressServiceClient) Update(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/address.AddressService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) Create(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/address.AddressService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) Get(ctx context.Context, in *OneAddressRequest, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/address.AddressService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressServiceServer is the server API for AddressService service.
type AddressServiceServer interface {
	Update(context.Context, *Address) (*Address, error)
	Create(context.Context, *Address) (*Address, error)
	Get(context.Context, *OneAddressRequest) (*Address, error)
}

// UnimplementedAddressServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAddressServiceServer struct {
}

func (*UnimplementedAddressServiceServer) Update(ctx context.Context, req *Address) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedAddressServiceServer) Create(ctx context.Context, req *Address) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAddressServiceServer) Get(ctx context.Context, req *OneAddressRequest) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterAddressServiceServer(s *grpc.Server, srv AddressServiceServer) {
	s.RegisterService(&_AddressService_serviceDesc, srv)
}

func _AddressService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Address)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/address.AddressService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).Update(ctx, req.(*Address))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Address)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/address.AddressService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).Create(ctx, req.(*Address))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OneAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/address.AddressService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).Get(ctx, req.(*OneAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddressService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "address.AddressService",
	HandlerType: (*AddressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _AddressService_Update_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AddressService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AddressService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "address/address.proto",
}

func (m *Address) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Address) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Address) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.UpdatedAt) > 0 {
		i -= len(m.UpdatedAt)
		copy(dAtA[i:], m.UpdatedAt)
		i = encodeVarintAddress(dAtA, i, uint64(len(m.UpdatedAt)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xaa
	}
	if len(m.CreatedAt) > 0 {
		i -= len(m.CreatedAt)
		copy(dAtA[i:], m.CreatedAt)
		i = encodeVarintAddress(dAtA, i, uint64(len(m.CreatedAt)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa2
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintAddress(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.DetailAddress) > 0 {
		i -= len(m.DetailAddress)
		copy(dAtA[i:], m.DetailAddress)
		i = encodeVarintAddress(dAtA, i, uint64(len(m.DetailAddress)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Province) > 0 {
		i -= len(m.Province)
		copy(dAtA[i:], m.Province)
		i = encodeVarintAddress(dAtA, i, uint64(len(m.Province)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Id != 0 {
		i = encodeVarintAddress(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OneAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OneAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OneAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Id != 0 {
		i = encodeVarintAddress(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAddress(dAtA []byte, offset int, v uint64) int {
	offset -= sovAddress(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Address) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAddress(uint64(m.Id))
	}
	l = len(m.Province)
	if l > 0 {
		n += 1 + l + sovAddress(uint64(l))
	}
	l = len(m.DetailAddress)
	if l > 0 {
		n += 1 + l + sovAddress(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovAddress(uint64(l))
	}
	l = len(m.CreatedAt)
	if l > 0 {
		n += 2 + l + sovAddress(uint64(l))
	}
	l = len(m.UpdatedAt)
	if l > 0 {
		n += 2 + l + sovAddress(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *OneAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAddress(uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAddress(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAddress(x uint64) (n int) {
	return sovAddress(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Address) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAddress
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
			return fmt.Errorf("proto: Address: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Address: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
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
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Province", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
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
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Province = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DetailAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
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
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DetailAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
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
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
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
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
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
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAddress
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OneAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAddress
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
			return fmt.Errorf("proto: OneAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OneAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
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
		default:
			iNdEx = preIndex
			skippy, err := skipAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAddress
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAddress(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAddress
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
					return 0, ErrIntOverflowAddress
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
					return 0, ErrIntOverflowAddress
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
				return 0, ErrInvalidLengthAddress
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAddress
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAddress
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAddress        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAddress          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAddress = fmt.Errorf("proto: unexpected end of group")
)