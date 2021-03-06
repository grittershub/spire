// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nodeattestor.proto

package nodeattestor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/spiffe/spire/proto/common"
import plugin "github.com/spiffe/spire/proto/common/plugin"

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

// ConfigureRequest from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type ConfigureRequest = plugin.ConfigureRequest

// ConfigureResponse from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type ConfigureResponse = plugin.ConfigureResponse

// GetPluginInfoRequest from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type GetPluginInfoRequest = plugin.GetPluginInfoRequest

// GetPluginInfoResponse from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type GetPluginInfoResponse = plugin.GetPluginInfoResponse

// PluginInfoRequest from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type PluginInfoRequest = plugin.PluginInfoRequest

// PluginInfoReply from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type PluginInfoReply = plugin.PluginInfoReply

// StopRequest from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type StopRequest = plugin.StopRequest

// StopReply from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type StopReply = plugin.StopReply

// Empty from public import github.com/spiffe/spire/proto/common/common.proto
type Empty = common.Empty

// AttestedData from public import github.com/spiffe/spire/proto/common/common.proto
type AttestedData = common.AttestedData

// Selector from public import github.com/spiffe/spire/proto/common/common.proto
type Selector = common.Selector

// Selectors from public import github.com/spiffe/spire/proto/common/common.proto
type Selectors = common.Selectors

// RegistrationEntry from public import github.com/spiffe/spire/proto/common/common.proto
type RegistrationEntry = common.RegistrationEntry

// RegistrationEntries from public import github.com/spiffe/spire/proto/common/common.proto
type RegistrationEntries = common.RegistrationEntries

// * Represents a request to attest a node.
type AttestRequest struct {
	// * A type which contains attestation data for specific platform.
	AttestedData *common.AttestedData `protobuf:"bytes,1,opt,name=attestedData" json:"attestedData,omitempty"`
	// * Is true if the Base SPIFFE ID is present in the Attested Node table.
	AttestedBefore       bool     `protobuf:"varint,2,opt,name=attestedBefore" json:"attestedBefore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttestRequest) Reset()         { *m = AttestRequest{} }
func (m *AttestRequest) String() string { return proto.CompactTextString(m) }
func (*AttestRequest) ProtoMessage()    {}
func (*AttestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodeattestor_ea7a50c0f43d681c, []int{0}
}
func (m *AttestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestRequest.Unmarshal(m, b)
}
func (m *AttestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestRequest.Marshal(b, m, deterministic)
}
func (dst *AttestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestRequest.Merge(dst, src)
}
func (m *AttestRequest) XXX_Size() int {
	return xxx_messageInfo_AttestRequest.Size(m)
}
func (m *AttestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttestRequest proto.InternalMessageInfo

func (m *AttestRequest) GetAttestedData() *common.AttestedData {
	if m != nil {
		return m.AttestedData
	}
	return nil
}

func (m *AttestRequest) GetAttestedBefore() bool {
	if m != nil {
		return m.AttestedBefore
	}
	return false
}

// * Represents a response when attesting a node.
type AttestResponse struct {
	// * True/False
	Valid bool `protobuf:"varint,1,opt,name=valid" json:"valid,omitempty"`
	// * Used for the Server to validate the SPIFFE Id in the Certificate signing request.
	BaseSPIFFEID         string   `protobuf:"bytes,2,opt,name=baseSPIFFEID" json:"baseSPIFFEID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttestResponse) Reset()         { *m = AttestResponse{} }
func (m *AttestResponse) String() string { return proto.CompactTextString(m) }
func (*AttestResponse) ProtoMessage()    {}
func (*AttestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodeattestor_ea7a50c0f43d681c, []int{1}
}
func (m *AttestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestResponse.Unmarshal(m, b)
}
func (m *AttestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestResponse.Marshal(b, m, deterministic)
}
func (dst *AttestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestResponse.Merge(dst, src)
}
func (m *AttestResponse) XXX_Size() int {
	return xxx_messageInfo_AttestResponse.Size(m)
}
func (m *AttestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AttestResponse proto.InternalMessageInfo

func (m *AttestResponse) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *AttestResponse) GetBaseSPIFFEID() string {
	if m != nil {
		return m.BaseSPIFFEID
	}
	return ""
}

func init() {
	proto.RegisterType((*AttestRequest)(nil), "spire.agent.nodeattestor.AttestRequest")
	proto.RegisterType((*AttestResponse)(nil), "spire.agent.nodeattestor.AttestResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeAttestor service

type NodeAttestorClient interface {
	// * Attesta a node.
	Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error)
	// * Responsible for configuration of the plugin.
	Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error)
	// * Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error)
}

type nodeAttestorClient struct {
	cc *grpc.ClientConn
}

func NewNodeAttestorClient(cc *grpc.ClientConn) NodeAttestorClient {
	return &nodeAttestorClient{cc}
}

func (c *nodeAttestorClient) Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error) {
	out := new(AttestResponse)
	err := grpc.Invoke(ctx, "/spire.agent.nodeattestor.NodeAttestor/Attest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAttestorClient) Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error) {
	out := new(plugin.ConfigureResponse)
	err := grpc.Invoke(ctx, "/spire.agent.nodeattestor.NodeAttestor/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAttestorClient) GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error) {
	out := new(plugin.GetPluginInfoResponse)
	err := grpc.Invoke(ctx, "/spire.agent.nodeattestor.NodeAttestor/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeAttestor service

type NodeAttestorServer interface {
	// * Attesta a node.
	Attest(context.Context, *AttestRequest) (*AttestResponse, error)
	// * Responsible for configuration of the plugin.
	Configure(context.Context, *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error)
	// * Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(context.Context, *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error)
}

func RegisterNodeAttestorServer(s *grpc.Server, srv NodeAttestorServer) {
	s.RegisterService(&_NodeAttestor_serviceDesc, srv)
}

func _NodeAttestor_Attest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).Attest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.nodeattestor.NodeAttestor/Attest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).Attest(ctx, req.(*AttestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAttestor_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.nodeattestor.NodeAttestor/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).Configure(ctx, req.(*plugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAttestor_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.nodeattestor.NodeAttestor/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, req.(*plugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeAttestor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.agent.nodeattestor.NodeAttestor",
	HandlerType: (*NodeAttestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Attest",
			Handler:    _NodeAttestor_Attest_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _NodeAttestor_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _NodeAttestor_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodeattestor.proto",
}

func init() { proto.RegisterFile("nodeattestor.proto", fileDescriptor_nodeattestor_ea7a50c0f43d681c) }

var fileDescriptor_nodeattestor_ea7a50c0f43d681c = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0x4e, 0xf2, 0x40,
	0x10, 0xff, 0x4a, 0xf2, 0x11, 0x58, 0x0b, 0x87, 0x8d, 0x07, 0xd2, 0x13, 0x21, 0x11, 0xd1, 0xc3,
	0x36, 0xe2, 0xc5, 0x93, 0x09, 0x88, 0x18, 0x3c, 0x98, 0xa6, 0xde, 0xf0, 0xb4, 0xd0, 0xd9, 0xba,
	0x09, 0xec, 0xd4, 0xee, 0x56, 0x1f, 0xcb, 0x57, 0x34, 0xd9, 0x6d, 0x4d, 0x6b, 0x34, 0x70, 0x9a,
	0xce, 0xcc, 0xef, 0xcf, 0xcc, 0x74, 0x09, 0x55, 0x98, 0x00, 0x37, 0x06, 0xb4, 0xc1, 0x9c, 0x65,
	0x39, 0x1a, 0xa4, 0x03, 0x9d, 0xc9, 0x1c, 0x18, 0x4f, 0x41, 0x19, 0x56, 0xef, 0x07, 0x37, 0xa9,
	0x34, 0xaf, 0xc5, 0x86, 0x6d, 0x71, 0x1f, 0xea, 0x4c, 0x0a, 0x01, 0xa1, 0xc5, 0x86, 0x96, 0x18,
	0x6e, 0x71, 0xbf, 0x47, 0x15, 0x66, 0xbb, 0x22, 0x95, 0x55, 0x70, 0x9a, 0xc1, 0xd5, 0x51, 0x4c,
	0x17, 0x1c, 0x65, 0xf4, 0x41, 0x7a, 0x33, 0x6b, 0x1c, 0xc3, 0x5b, 0x01, 0xda, 0xd0, 0x5b, 0xe2,
	0xbb, 0x49, 0x20, 0x59, 0x70, 0xc3, 0x07, 0xde, 0xd0, 0x9b, 0x9c, 0x4c, 0x03, 0xe6, 0xc6, 0x2d,
	0xb9, 0xb3, 0x1a, 0x22, 0x6e, 0xe0, 0xe9, 0x98, 0xf4, 0xab, 0x7c, 0x0e, 0x02, 0x73, 0x18, 0xb4,
	0x86, 0xde, 0xa4, 0x13, 0xff, 0xa8, 0x8e, 0x1e, 0x49, 0xbf, 0x32, 0xd6, 0x19, 0x2a, 0x0d, 0xf4,
	0x94, 0xfc, 0x7f, 0xe7, 0x3b, 0x99, 0x58, 0xcb, 0x4e, 0xec, 0x12, 0x3a, 0x22, 0xfe, 0x86, 0x6b,
	0x78, 0x8e, 0x56, 0xcb, 0xe5, 0xfd, 0x6a, 0x61, 0xd5, 0xba, 0x71, 0xa3, 0x36, 0xfd, 0x6c, 0x11,
	0xff, 0x09, 0x13, 0x98, 0x95, 0x27, 0xa4, 0x2f, 0xa4, 0xed, 0xbe, 0xe9, 0x39, 0xfb, 0xeb, 0xce,
	0xac, 0xb1, 0x77, 0x30, 0x39, 0x0c, 0x2c, 0xe7, 0x5c, 0x93, 0xee, 0x1d, 0x2a, 0x21, 0xd3, 0x22,
	0x07, 0x7a, 0xd6, 0x3c, 0x4c, 0xf9, 0x3b, 0xbe, 0xfb, 0x95, 0xfa, 0xf8, 0x10, 0xac, 0xd4, 0x16,
	0xa4, 0xf7, 0x00, 0x26, 0xb2, 0xed, 0x95, 0x12, 0x48, 0x2f, 0x7e, 0x25, 0x36, 0x30, 0x95, 0xc7,
	0xe5, 0x31, 0x50, 0xe7, 0x33, 0xef, 0xaf, 0xfd, 0xfa, 0x8a, 0xd1, 0xbf, 0xc8, 0xdb, 0xb4, 0xed,
	0x8b, 0xb8, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x7e, 0xc2, 0xb8, 0xae, 0x02, 0x00, 0x00,
}
