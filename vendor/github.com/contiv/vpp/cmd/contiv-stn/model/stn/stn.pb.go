// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stn.proto

package stn

/*
Package stn defines STN (Steal the NIC) GRPC service.
*/

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// STNRequest represents a request to steal or release of an interface.
type STNRequest struct {
	// The interface to be stolen / released.
	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	// True if DHCP is enabled on the interface.
	DhcpEnabled          bool     `protobuf:"varint,2,opt,name=dhcp_enabled,json=dhcpEnabled,proto3" json:"dhcp_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *STNRequest) Reset()         { *m = STNRequest{} }
func (m *STNRequest) String() string { return proto.CompactTextString(m) }
func (*STNRequest) ProtoMessage()    {}
func (*STNRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_stn_2aee956be88da58d, []int{0}
}
func (m *STNRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_STNRequest.Unmarshal(m, b)
}
func (m *STNRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_STNRequest.Marshal(b, m, deterministic)
}
func (dst *STNRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_STNRequest.Merge(dst, src)
}
func (m *STNRequest) XXX_Size() int {
	return xxx_messageInfo_STNRequest.Size(m)
}
func (m *STNRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_STNRequest.DiscardUnknown(m)
}

var xxx_messageInfo_STNRequest proto.InternalMessageInfo

func (m *STNRequest) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *STNRequest) GetDhcpEnabled() bool {
	if m != nil {
		return m.DhcpEnabled
	}
	return false
}

// The reply to the STNRequest. Contains the original config of the stolen interface.
type STNReply struct {
	// Result code. 0 = success, non-zero = error.
	Result uint32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	// Error string in case that result != 0.
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// PCI address of the interface.
	PciAddress string `protobuf:"bytes,3,opt,name=pci_address,json=pciAddress,proto3" json:"pci_address,omitempty"`
	// List of IP addresses assigned to the interface.
	IpAddresses []string `protobuf:"bytes,4,rep,name=ip_addresses,json=ipAddresses" json:"ip_addresses,omitempty"`
	// List of routes related to the interface.
	Routes []*STNReply_Route `protobuf:"bytes,5,rep,name=routes" json:"routes,omitempty"`
	// Kernel driver used by this interface before stealing.
	KernelDriver         string   `protobuf:"bytes,6,opt,name=kernel_driver,json=kernelDriver,proto3" json:"kernel_driver,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *STNReply) Reset()         { *m = STNReply{} }
func (m *STNReply) String() string { return proto.CompactTextString(m) }
func (*STNReply) ProtoMessage()    {}
func (*STNReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_stn_2aee956be88da58d, []int{1}
}
func (m *STNReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_STNReply.Unmarshal(m, b)
}
func (m *STNReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_STNReply.Marshal(b, m, deterministic)
}
func (dst *STNReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_STNReply.Merge(dst, src)
}
func (m *STNReply) XXX_Size() int {
	return xxx_messageInfo_STNReply.Size(m)
}
func (m *STNReply) XXX_DiscardUnknown() {
	xxx_messageInfo_STNReply.DiscardUnknown(m)
}

var xxx_messageInfo_STNReply proto.InternalMessageInfo

func (m *STNReply) GetResult() uint32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *STNReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *STNReply) GetPciAddress() string {
	if m != nil {
		return m.PciAddress
	}
	return ""
}

func (m *STNReply) GetIpAddresses() []string {
	if m != nil {
		return m.IpAddresses
	}
	return nil
}

func (m *STNReply) GetRoutes() []*STNReply_Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *STNReply) GetKernelDriver() string {
	if m != nil {
		return m.KernelDriver
	}
	return ""
}

// A route related to the interface.
type STNReply_Route struct {
	// Destination subnet prefix.
	DestinationSubnet string `protobuf:"bytes,1,opt,name=destination_subnet,json=destinationSubnet,proto3" json:"destination_subnet,omitempty"`
	// Next hop IP address.
	NextHopIp            string   `protobuf:"bytes,2,opt,name=next_hop_ip,json=nextHopIp,proto3" json:"next_hop_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *STNReply_Route) Reset()         { *m = STNReply_Route{} }
func (m *STNReply_Route) String() string { return proto.CompactTextString(m) }
func (*STNReply_Route) ProtoMessage()    {}
func (*STNReply_Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_stn_2aee956be88da58d, []int{1, 0}
}
func (m *STNReply_Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_STNReply_Route.Unmarshal(m, b)
}
func (m *STNReply_Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_STNReply_Route.Marshal(b, m, deterministic)
}
func (dst *STNReply_Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_STNReply_Route.Merge(dst, src)
}
func (m *STNReply_Route) XXX_Size() int {
	return xxx_messageInfo_STNReply_Route.Size(m)
}
func (m *STNReply_Route) XXX_DiscardUnknown() {
	xxx_messageInfo_STNReply_Route.DiscardUnknown(m)
}

var xxx_messageInfo_STNReply_Route proto.InternalMessageInfo

func (m *STNReply_Route) GetDestinationSubnet() string {
	if m != nil {
		return m.DestinationSubnet
	}
	return ""
}

func (m *STNReply_Route) GetNextHopIp() string {
	if m != nil {
		return m.NextHopIp
	}
	return ""
}

func init() {
	proto.RegisterType((*STNRequest)(nil), "stn.STNRequest")
	proto.RegisterType((*STNReply)(nil), "stn.STNReply")
	proto.RegisterType((*STNReply_Route)(nil), "stn.STNReply.Route")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for STN service

type STNClient interface {
	// The request to steal (unconfigure) an interface identified by its name.
	StealInterface(ctx context.Context, in *STNRequest, opts ...grpc.CallOption) (*STNReply, error)
	// The request to revert config of a previously "stolen" (unconfigured) interface.
	ReleaseInterface(ctx context.Context, in *STNRequest, opts ...grpc.CallOption) (*STNReply, error)
	// Request to return information about the stolen interface.
	StolenInterfaceInfo(ctx context.Context, in *STNRequest, opts ...grpc.CallOption) (*STNReply, error)
}

type sTNClient struct {
	cc *grpc.ClientConn
}

func NewSTNClient(cc *grpc.ClientConn) STNClient {
	return &sTNClient{cc}
}

func (c *sTNClient) StealInterface(ctx context.Context, in *STNRequest, opts ...grpc.CallOption) (*STNReply, error) {
	out := new(STNReply)
	err := c.cc.Invoke(ctx, "/stn.STN/StealInterface", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sTNClient) ReleaseInterface(ctx context.Context, in *STNRequest, opts ...grpc.CallOption) (*STNReply, error) {
	out := new(STNReply)
	err := c.cc.Invoke(ctx, "/stn.STN/ReleaseInterface", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sTNClient) StolenInterfaceInfo(ctx context.Context, in *STNRequest, opts ...grpc.CallOption) (*STNReply, error) {
	out := new(STNReply)
	err := c.cc.Invoke(ctx, "/stn.STN/StolenInterfaceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for STN service

type STNServer interface {
	// The request to steal (unconfigure) an interface identified by its name.
	StealInterface(context.Context, *STNRequest) (*STNReply, error)
	// The request to revert config of a previously "stolen" (unconfigured) interface.
	ReleaseInterface(context.Context, *STNRequest) (*STNReply, error)
	// Request to return information about the stolen interface.
	StolenInterfaceInfo(context.Context, *STNRequest) (*STNReply, error)
}

func RegisterSTNServer(s *grpc.Server, srv STNServer) {
	s.RegisterService(&_STN_serviceDesc, srv)
}

func _STN_StealInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(STNRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(STNServer).StealInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stn.STN/StealInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(STNServer).StealInterface(ctx, req.(*STNRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _STN_ReleaseInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(STNRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(STNServer).ReleaseInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stn.STN/ReleaseInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(STNServer).ReleaseInterface(ctx, req.(*STNRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _STN_StolenInterfaceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(STNRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(STNServer).StolenInterfaceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stn.STN/StolenInterfaceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(STNServer).StolenInterfaceInfo(ctx, req.(*STNRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _STN_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stn.STN",
	HandlerType: (*STNServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StealInterface",
			Handler:    _STN_StealInterface_Handler,
		},
		{
			MethodName: "ReleaseInterface",
			Handler:    _STN_ReleaseInterface_Handler,
		},
		{
			MethodName: "StolenInterfaceInfo",
			Handler:    _STN_StolenInterfaceInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stn.proto",
}

func init() { proto.RegisterFile("stn.proto", fileDescriptor_stn_2aee956be88da58d) }

var fileDescriptor_stn_2aee956be88da58d = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x8f, 0xd3, 0x30,
	0x10, 0x85, 0x49, 0x43, 0xa3, 0x66, 0xd2, 0x14, 0x70, 0x11, 0x8a, 0x7a, 0x80, 0x50, 0x84, 0x14,
	0x09, 0xd1, 0x43, 0x41, 0xe2, 0x5c, 0x09, 0x24, 0x7a, 0xe9, 0xc1, 0xa9, 0x7a, 0x8d, 0xdc, 0x64,
	0xaa, 0x5a, 0xa4, 0xb6, 0xb1, 0x1d, 0x04, 0x7f, 0x08, 0xed, 0xcf, 0x5c, 0xc5, 0x49, 0xba, 0xbb,
	0xa7, 0xd5, 0x1e, 0xe7, 0x7b, 0x79, 0x33, 0xf3, 0x26, 0x86, 0xd0, 0x58, 0xb1, 0x52, 0x5a, 0x5a,
	0x49, 0x7c, 0x63, 0xc5, 0xf2, 0x00, 0x90, 0xef, 0x77, 0x14, 0x7f, 0x37, 0x68, 0x2c, 0xf9, 0x08,
	0x33, 0x2e, 0x2c, 0xea, 0x13, 0x2b, 0xb1, 0x10, 0xec, 0x82, 0x89, 0x97, 0x7a, 0x59, 0x48, 0xe3,
	0x2b, 0xdd, 0xb1, 0x0b, 0x92, 0xf7, 0x30, 0xad, 0xce, 0xa5, 0x2a, 0x50, 0xb0, 0x63, 0x8d, 0x55,
	0x32, 0x4a, 0xbd, 0x6c, 0x42, 0xa3, 0x96, 0xfd, 0xe8, 0xd0, 0xf2, 0xff, 0x08, 0x26, 0xae, 0xb1,
	0xaa, 0xff, 0x91, 0x37, 0x10, 0x68, 0x34, 0x4d, 0x6d, 0x5d, 0xbb, 0x98, 0xf6, 0x15, 0x79, 0x0d,
	0x63, 0xd4, 0x5a, 0x6a, 0xd7, 0x20, 0xa4, 0x5d, 0x41, 0xde, 0x41, 0xa4, 0x4a, 0x5e, 0xb0, 0xaa,
	0xd2, 0x68, 0x4c, 0xe2, 0x3b, 0x0d, 0x54, 0xc9, 0x37, 0x1d, 0x69, 0xc7, 0x73, 0x35, 0xe8, 0x68,
	0x92, 0xe7, 0xa9, 0x9f, 0x85, 0x34, 0xe2, 0x6a, 0x33, 0x20, 0xf2, 0x09, 0x02, 0x2d, 0x1b, 0x8b,
	0x26, 0x19, 0xa7, 0x7e, 0x16, 0xad, 0xe7, 0xab, 0x36, 0xf7, 0xb0, 0xd0, 0x8a, 0xb6, 0x1a, 0xed,
	0x3f, 0x21, 0x1f, 0x20, 0xfe, 0x85, 0x5a, 0x60, 0x5d, 0x54, 0x9a, 0xff, 0x41, 0x9d, 0x04, 0x6e,
	0xe4, 0xb4, 0x83, 0xdf, 0x1d, 0x5b, 0x1c, 0x60, 0xec, 0x5c, 0xe4, 0x33, 0x90, 0x0a, 0x8d, 0xe5,
	0x82, 0x59, 0x2e, 0x45, 0x61, 0x9a, 0xa3, 0x40, 0xdb, 0xdf, 0xe9, 0xd5, 0x3d, 0x25, 0x77, 0x02,
	0x79, 0x0b, 0x91, 0xc0, 0xbf, 0xb6, 0x38, 0x4b, 0x55, 0x70, 0xd5, 0x27, 0x0d, 0x5b, 0xf4, 0x53,
	0xaa, 0xad, 0x5a, 0xdf, 0x78, 0xe0, 0xe7, 0xfb, 0x1d, 0x59, 0xc3, 0x2c, 0xb7, 0xc8, 0xea, 0xed,
	0x70, 0x69, 0xf2, 0xe2, 0x6e, 0x67, 0xf7, 0x77, 0x16, 0xf1, 0x83, 0x10, 0xcb, 0x67, 0xe4, 0x2b,
	0xbc, 0xa4, 0x58, 0x23, 0x33, 0xf8, 0x14, 0xd7, 0x37, 0x98, 0xe7, 0x56, 0xd6, 0x28, 0xae, 0xa6,
	0xad, 0x38, 0xc9, 0xc7, 0x8d, 0xc7, 0xc0, 0xbd, 0x9b, 0x2f, 0xb7, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x5c, 0xa7, 0xb1, 0x9b, 0x44, 0x02, 0x00, 0x00,
}