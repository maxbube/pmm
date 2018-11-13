// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inventory/services.proto

package inventory

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// MySQLService represents MySQL Service configuration.
type MySQLService struct {
	// Unique Service identifier.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique user-defined Service name.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MySQLService) Reset()         { *m = MySQLService{} }
func (m *MySQLService) String() string { return proto.CompactTextString(m) }
func (*MySQLService) ProtoMessage()    {}
func (*MySQLService) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_1ee5b9b0047f2b0a, []int{0}
}
func (m *MySQLService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MySQLService.Unmarshal(m, b)
}
func (m *MySQLService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MySQLService.Marshal(b, m, deterministic)
}
func (dst *MySQLService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MySQLService.Merge(dst, src)
}
func (m *MySQLService) XXX_Size() int {
	return xxx_messageInfo_MySQLService.Size(m)
}
func (m *MySQLService) XXX_DiscardUnknown() {
	xxx_messageInfo_MySQLService.DiscardUnknown(m)
}

var xxx_messageInfo_MySQLService proto.InternalMessageInfo

func (m *MySQLService) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MySQLService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListServicesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListServicesRequest) Reset()         { *m = ListServicesRequest{} }
func (m *ListServicesRequest) String() string { return proto.CompactTextString(m) }
func (*ListServicesRequest) ProtoMessage()    {}
func (*ListServicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_1ee5b9b0047f2b0a, []int{1}
}
func (m *ListServicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServicesRequest.Unmarshal(m, b)
}
func (m *ListServicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServicesRequest.Marshal(b, m, deterministic)
}
func (dst *ListServicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServicesRequest.Merge(dst, src)
}
func (m *ListServicesRequest) XXX_Size() int {
	return xxx_messageInfo_ListServicesRequest.Size(m)
}
func (m *ListServicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListServicesRequest proto.InternalMessageInfo

type ListServicesResponse struct {
	Mysql                []*MySQLService `protobuf:"bytes,1,rep,name=mysql,proto3" json:"mysql,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListServicesResponse) Reset()         { *m = ListServicesResponse{} }
func (m *ListServicesResponse) String() string { return proto.CompactTextString(m) }
func (*ListServicesResponse) ProtoMessage()    {}
func (*ListServicesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_1ee5b9b0047f2b0a, []int{2}
}
func (m *ListServicesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServicesResponse.Unmarshal(m, b)
}
func (m *ListServicesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServicesResponse.Marshal(b, m, deterministic)
}
func (dst *ListServicesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServicesResponse.Merge(dst, src)
}
func (m *ListServicesResponse) XXX_Size() int {
	return xxx_messageInfo_ListServicesResponse.Size(m)
}
func (m *ListServicesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServicesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListServicesResponse proto.InternalMessageInfo

func (m *ListServicesResponse) GetMysql() []*MySQLService {
	if m != nil {
		return m.Mysql
	}
	return nil
}

type GetServiceRequest struct {
	// Unique Service identifier.
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServiceRequest) Reset()         { *m = GetServiceRequest{} }
func (m *GetServiceRequest) String() string { return proto.CompactTextString(m) }
func (*GetServiceRequest) ProtoMessage()    {}
func (*GetServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_1ee5b9b0047f2b0a, []int{3}
}
func (m *GetServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceRequest.Unmarshal(m, b)
}
func (m *GetServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceRequest.Marshal(b, m, deterministic)
}
func (dst *GetServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceRequest.Merge(dst, src)
}
func (m *GetServiceRequest) XXX_Size() int {
	return xxx_messageInfo_GetServiceRequest.Size(m)
}
func (m *GetServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceRequest proto.InternalMessageInfo

func (m *GetServiceRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetServiceResponse struct {
	// Types that are valid to be assigned to Service:
	//	*GetServiceResponse_Mysql
	Service              isGetServiceResponse_Service `protobuf_oneof:"service"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GetServiceResponse) Reset()         { *m = GetServiceResponse{} }
func (m *GetServiceResponse) String() string { return proto.CompactTextString(m) }
func (*GetServiceResponse) ProtoMessage()    {}
func (*GetServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_1ee5b9b0047f2b0a, []int{4}
}
func (m *GetServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceResponse.Unmarshal(m, b)
}
func (m *GetServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceResponse.Marshal(b, m, deterministic)
}
func (dst *GetServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceResponse.Merge(dst, src)
}
func (m *GetServiceResponse) XXX_Size() int {
	return xxx_messageInfo_GetServiceResponse.Size(m)
}
func (m *GetServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceResponse proto.InternalMessageInfo

type isGetServiceResponse_Service interface {
	isGetServiceResponse_Service()
}

type GetServiceResponse_Mysql struct {
	Mysql *MySQLService `protobuf:"bytes,1,opt,name=mysql,proto3,oneof"`
}

func (*GetServiceResponse_Mysql) isGetServiceResponse_Service() {}

func (m *GetServiceResponse) GetService() isGetServiceResponse_Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *GetServiceResponse) GetMysql() *MySQLService {
	if x, ok := m.GetService().(*GetServiceResponse_Mysql); ok {
		return x.Mysql
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GetServiceResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GetServiceResponse_OneofMarshaler, _GetServiceResponse_OneofUnmarshaler, _GetServiceResponse_OneofSizer, []interface{}{
		(*GetServiceResponse_Mysql)(nil),
	}
}

func _GetServiceResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GetServiceResponse)
	// service
	switch x := m.Service.(type) {
	case *GetServiceResponse_Mysql:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Mysql); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GetServiceResponse.Service has unexpected type %T", x)
	}
	return nil
}

func _GetServiceResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GetServiceResponse)
	switch tag {
	case 1: // service.mysql
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MySQLService)
		err := b.DecodeMessage(msg)
		m.Service = &GetServiceResponse_Mysql{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GetServiceResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GetServiceResponse)
	// service
	switch x := m.Service.(type) {
	case *GetServiceResponse_Mysql:
		s := proto.Size(x.Mysql)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type AddMySQLServiceRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMySQLServiceRequest) Reset()         { *m = AddMySQLServiceRequest{} }
func (m *AddMySQLServiceRequest) String() string { return proto.CompactTextString(m) }
func (*AddMySQLServiceRequest) ProtoMessage()    {}
func (*AddMySQLServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_1ee5b9b0047f2b0a, []int{5}
}
func (m *AddMySQLServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMySQLServiceRequest.Unmarshal(m, b)
}
func (m *AddMySQLServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMySQLServiceRequest.Marshal(b, m, deterministic)
}
func (dst *AddMySQLServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMySQLServiceRequest.Merge(dst, src)
}
func (m *AddMySQLServiceRequest) XXX_Size() int {
	return xxx_messageInfo_AddMySQLServiceRequest.Size(m)
}
func (m *AddMySQLServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMySQLServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddMySQLServiceRequest proto.InternalMessageInfo

type AddMySQLServiceResponse struct {
	Mysql                *MySQLService `protobuf:"bytes,1,opt,name=mysql,proto3" json:"mysql,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AddMySQLServiceResponse) Reset()         { *m = AddMySQLServiceResponse{} }
func (m *AddMySQLServiceResponse) String() string { return proto.CompactTextString(m) }
func (*AddMySQLServiceResponse) ProtoMessage()    {}
func (*AddMySQLServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_1ee5b9b0047f2b0a, []int{6}
}
func (m *AddMySQLServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMySQLServiceResponse.Unmarshal(m, b)
}
func (m *AddMySQLServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMySQLServiceResponse.Marshal(b, m, deterministic)
}
func (dst *AddMySQLServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMySQLServiceResponse.Merge(dst, src)
}
func (m *AddMySQLServiceResponse) XXX_Size() int {
	return xxx_messageInfo_AddMySQLServiceResponse.Size(m)
}
func (m *AddMySQLServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMySQLServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddMySQLServiceResponse proto.InternalMessageInfo

func (m *AddMySQLServiceResponse) GetMysql() *MySQLService {
	if m != nil {
		return m.Mysql
	}
	return nil
}

type RemoveServiceRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveServiceRequest) Reset()         { *m = RemoveServiceRequest{} }
func (m *RemoveServiceRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveServiceRequest) ProtoMessage()    {}
func (*RemoveServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_1ee5b9b0047f2b0a, []int{7}
}
func (m *RemoveServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveServiceRequest.Unmarshal(m, b)
}
func (m *RemoveServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveServiceRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveServiceRequest.Merge(dst, src)
}
func (m *RemoveServiceRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveServiceRequest.Size(m)
}
func (m *RemoveServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveServiceRequest proto.InternalMessageInfo

func (m *RemoveServiceRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RemoveServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveServiceResponse) Reset()         { *m = RemoveServiceResponse{} }
func (m *RemoveServiceResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveServiceResponse) ProtoMessage()    {}
func (*RemoveServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_1ee5b9b0047f2b0a, []int{8}
}
func (m *RemoveServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveServiceResponse.Unmarshal(m, b)
}
func (m *RemoveServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveServiceResponse.Marshal(b, m, deterministic)
}
func (dst *RemoveServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveServiceResponse.Merge(dst, src)
}
func (m *RemoveServiceResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveServiceResponse.Size(m)
}
func (m *RemoveServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveServiceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MySQLService)(nil), "inventory.MySQLService")
	proto.RegisterType((*ListServicesRequest)(nil), "inventory.ListServicesRequest")
	proto.RegisterType((*ListServicesResponse)(nil), "inventory.ListServicesResponse")
	proto.RegisterType((*GetServiceRequest)(nil), "inventory.GetServiceRequest")
	proto.RegisterType((*GetServiceResponse)(nil), "inventory.GetServiceResponse")
	proto.RegisterType((*AddMySQLServiceRequest)(nil), "inventory.AddMySQLServiceRequest")
	proto.RegisterType((*AddMySQLServiceResponse)(nil), "inventory.AddMySQLServiceResponse")
	proto.RegisterType((*RemoveServiceRequest)(nil), "inventory.RemoveServiceRequest")
	proto.RegisterType((*RemoveServiceResponse)(nil), "inventory.RemoveServiceResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServicesClient is the client API for Services service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServicesClient interface {
	// ListServices returns a list of all Services.
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
	// GetService returns a single Service by ID.
	GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error)
	// AddMySQLService adds MySQL Service.
	AddMySQLService(ctx context.Context, in *AddMySQLServiceRequest, opts ...grpc.CallOption) (*AddMySQLServiceResponse, error)
	// RemoveService removes Service without any Agents.
	RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*RemoveServiceResponse, error)
}

type servicesClient struct {
	cc *grpc.ClientConn
}

func NewServicesClient(cc *grpc.ClientConn) ServicesClient {
	return &servicesClient{cc}
}

func (c *servicesClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/ListServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error) {
	out := new(GetServiceResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/GetService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) AddMySQLService(ctx context.Context, in *AddMySQLServiceRequest, opts ...grpc.CallOption) (*AddMySQLServiceResponse, error) {
	out := new(AddMySQLServiceResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/AddMySQLService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*RemoveServiceResponse, error) {
	out := new(RemoveServiceResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/RemoveService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicesServer is the server API for Services service.
type ServicesServer interface {
	// ListServices returns a list of all Services.
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
	// GetService returns a single Service by ID.
	GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error)
	// AddMySQLService adds MySQL Service.
	AddMySQLService(context.Context, *AddMySQLServiceRequest) (*AddMySQLServiceResponse, error)
	// RemoveService removes Service without any Agents.
	RemoveService(context.Context, *RemoveServiceRequest) (*RemoveServiceResponse, error)
}

func RegisterServicesServer(s *grpc.Server, srv ServicesServer) {
	s.RegisterService(&_Services_serviceDesc, srv)
}

func _Services_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/ListServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/GetService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).GetService(ctx, req.(*GetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_AddMySQLService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMySQLServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).AddMySQLService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/AddMySQLService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).AddMySQLService(ctx, req.(*AddMySQLServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_RemoveService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).RemoveService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/RemoveService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).RemoveService(ctx, req.(*RemoveServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Services_serviceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.Services",
	HandlerType: (*ServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListServices",
			Handler:    _Services_ListServices_Handler,
		},
		{
			MethodName: "GetService",
			Handler:    _Services_GetService_Handler,
		},
		{
			MethodName: "AddMySQLService",
			Handler:    _Services_AddMySQLService_Handler,
		},
		{
			MethodName: "RemoveService",
			Handler:    _Services_RemoveService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory/services.proto",
}

func init() { proto.RegisterFile("inventory/services.proto", fileDescriptor_services_1ee5b9b0047f2b0a) }

var fileDescriptor_services_1ee5b9b0047f2b0a = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0xbb, 0x2e, 0xfd, 0xc3, 0x14, 0x5a, 0x75, 0x0b, 0xc5, 0xb2, 0x68, 0x6b, 0x96, 0x0a,
	0xb9, 0xb4, 0xf5, 0xb6, 0xf4, 0xc6, 0xad, 0x95, 0xaa, 0x70, 0x20, 0x52, 0x62, 0x9e, 0xc0, 0x89,
	0x57, 0xc8, 0x12, 0x78, 0x81, 0x75, 0x1c, 0x71, 0x8a, 0x94, 0x1c, 0x73, 0xcc, 0x25, 0xef, 0x95,
	0x57, 0xc8, 0x83, 0x44, 0xd8, 0x0b, 0x5e, 0x1b, 0x4c, 0x72, 0xb3, 0x3c, 0xdf, 0xcc, 0xef, 0x9b,
	0xf9, 0x6c, 0xd0, 0xfd, 0x20, 0x62, 0x41, 0xc8, 0x17, 0x4b, 0x2a, 0xd8, 0x22, 0xf2, 0x4f, 0x99,
	0xb0, 0x67, 0x0b, 0x1e, 0x72, 0x5c, 0xde, 0x54, 0x8c, 0xe6, 0x98, 0xf3, 0xf1, 0x84, 0x51, 0x77,
	0xe6, 0x53, 0x37, 0x08, 0x78, 0xe8, 0x86, 0x3e, 0x0f, 0xa4, 0x90, 0xf4, 0xa0, 0x72, 0xb8, 0x1c,
	0x1d, 0x0f, 0x47, 0x49, 0x3f, 0x7e, 0x0b, 0x9a, 0xef, 0xe9, 0xc8, 0x44, 0x56, 0xd5, 0xd1, 0x7c,
	0x0f, 0x63, 0x28, 0x05, 0xee, 0x94, 0xe9, 0x9a, 0x89, 0xac, 0xb2, 0x13, 0x3f, 0x93, 0x3a, 0x7c,
	0x18, 0xfa, 0x22, 0x94, 0x2d, 0xc2, 0x61, 0xf3, 0x33, 0x26, 0x42, 0xf2, 0x1f, 0x6a, 0xd9, 0xd7,
	0x62, 0xc6, 0x03, 0xc1, 0xf0, 0x4f, 0x78, 0x31, 0x5d, 0x8a, 0xf9, 0x44, 0x47, 0xe6, 0x73, 0xeb,
	0x4d, 0xaf, 0x61, 0x6f, 0xbc, 0xd9, 0x2a, 0xda, 0x49, 0x54, 0xa4, 0x0d, 0xef, 0x0f, 0xd8, 0x7a,
	0x8a, 0x9c, 0x9d, 0xb7, 0x45, 0x8e, 0x00, 0xab, 0x22, 0x49, 0xa2, 0x29, 0x09, 0xed, 0x21, 0x0d,
	0x9e, 0x49, 0xd6, 0xbf, 0x32, 0xbc, 0x92, 0x87, 0x23, 0x3a, 0x7c, 0xfc, 0xeb, 0x79, 0x19, 0x43,
	0x72, 0xaf, 0x01, 0x34, 0xb6, 0x2a, 0xdb, 0xab, 0xa1, 0x27, 0xac, 0xd6, 0x81, 0x9a, 0xc3, 0xa6,
	0x3c, 0x62, 0x8f, 0x6c, 0xd7, 0x80, 0x7a, 0x4e, 0x97, 0xf0, 0x7a, 0xb7, 0x25, 0x78, 0xbd, 0xbe,
	0x2f, 0xbe, 0x80, 0x8a, 0x7a, 0x6f, 0xfc, 0x59, 0xa1, 0xef, 0xc8, 0xc7, 0xf8, 0x52, 0x58, 0x4f,
	0xa6, 0x13, 0xfb, 0xf2, 0xee, 0xfe, 0x46, 0xb3, 0x48, 0x9b, 0x46, 0xbf, 0x68, 0xfa, 0x69, 0xad,
	0x75, 0x54, 0x6d, 0xea, 0xa3, 0x2e, 0x3e, 0x07, 0x48, 0x43, 0xc0, 0x4d, 0x65, 0xfc, 0x56, 0x80,
	0xc6, 0xa7, 0x82, 0xaa, 0x44, 0xff, 0x88, 0xd1, 0x1d, 0xd2, 0x2a, 0x40, 0xa7, 0x2d, 0x2b, 0xf0,
	0x35, 0x82, 0x77, 0xb9, 0x48, 0x70, 0x4b, 0x01, 0xec, 0x0e, 0xd2, 0x20, 0xfb, 0x24, 0xd2, 0xc8,
	0xef, 0xd8, 0xc8, 0x77, 0xd2, 0x29, 0x30, 0x92, 0xeb, 0x5b, 0xb9, 0xb9, 0x42, 0x50, 0xcd, 0xc4,
	0x85, 0xd5, 0x4b, 0xef, 0x0a, 0xdc, 0x30, 0x8b, 0x05, 0xd2, 0x07, 0x8d, 0x7d, 0x7c, 0x23, 0x5f,
	0x0b, 0x7c, 0x64, 0xba, 0xfa, 0xa8, 0x7b, 0xf2, 0x32, 0xfe, 0x9f, 0xff, 0x3c, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xcb, 0xb7, 0xeb, 0x06, 0x14, 0x04, 0x00, 0x00,
}