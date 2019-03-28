// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tmcs_rpc.proto

package tmcs_msg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RPCCode int32

const (
	RPCCode_Ping           RPCCode = 0
	RPCCode_Pong           RPCCode = 1
	RPCCode_LinkClicked    RPCCode = 2000
	RPCCode_ContactRequest RPCCode = 2001
	RPCCode_ContactOnline  RPCCode = 2002
	RPCCode_ContactOffline RPCCode = 2003
)

var RPCCode_name = map[int32]string{
	0:    "Ping",
	1:    "Pong",
	2000: "LinkClicked",
	2001: "ContactRequest",
	2002: "ContactOnline",
	2003: "ContactOffline",
}

var RPCCode_value = map[string]int32{
	"Ping":           0,
	"Pong":           1,
	"LinkClicked":    2000,
	"ContactRequest": 2001,
	"ContactOnline":  2002,
	"ContactOffline": 2003,
}

func (x RPCCode) String() string {
	return proto.EnumName(RPCCode_name, int32(x))
}

func (RPCCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_89cbeccdfc94c509, []int{0}
}

type RPC struct {
	RpcId   int32   `protobuf:"varint,1,opt,name=rpcId,proto3" json:"rpcId,omitempty"`
	Code    RPCCode `protobuf:"varint,2,opt,name=code,proto3,enum=tmcs_msg.RPCCode" json:"code,omitempty"`
	Message string  `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*RPC_TextData
	Data                 isRPC_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RPC) Reset()         { *m = RPC{} }
func (m *RPC) String() string { return proto.CompactTextString(m) }
func (*RPC) ProtoMessage()    {}
func (*RPC) Descriptor() ([]byte, []int) {
	return fileDescriptor_89cbeccdfc94c509, []int{0}
}

func (m *RPC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPC.Unmarshal(m, b)
}
func (m *RPC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPC.Marshal(b, m, deterministic)
}
func (m *RPC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPC.Merge(m, src)
}
func (m *RPC) XXX_Size() int {
	return xxx_messageInfo_RPC.Size(m)
}
func (m *RPC) XXX_DiscardUnknown() {
	xxx_messageInfo_RPC.DiscardUnknown(m)
}

var xxx_messageInfo_RPC proto.InternalMessageInfo

func (m *RPC) GetRpcId() int32 {
	if m != nil {
		return m.RpcId
	}
	return 0
}

func (m *RPC) GetCode() RPCCode {
	if m != nil {
		return m.Code
	}
	return RPCCode_Ping
}

func (m *RPC) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type isRPC_Data interface {
	isRPC_Data()
}

type RPC_TextData struct {
	TextData string `protobuf:"bytes,4,opt,name=textData,proto3,oneof"`
}

func (*RPC_TextData) isRPC_Data() {}

func (m *RPC) GetData() isRPC_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *RPC) GetTextData() string {
	if x, ok := m.GetData().(*RPC_TextData); ok {
		return x.TextData
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RPC) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RPC_TextData)(nil),
	}
}

func init() {
	proto.RegisterEnum("tmcs_msg.RPCCode", RPCCode_name, RPCCode_value)
	proto.RegisterType((*RPC)(nil), "tmcs_msg.RPC")
}

func init() { proto.RegisterFile("tmcs_rpc.proto", fileDescriptor_89cbeccdfc94c509) }

var fileDescriptor_89cbeccdfc94c509 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0xc9, 0x4d, 0x2e,
	0x8e, 0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0xf3, 0x73, 0x8b,
	0xd3, 0x95, 0x9a, 0x18, 0xb9, 0x98, 0x83, 0x02, 0x9c, 0x85, 0x44, 0xb8, 0x58, 0x81, 0xd2, 0x9e,
	0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x10, 0x8e, 0x90, 0x2a, 0x17, 0x4b, 0x72, 0x7e,
	0x4a, 0xaa, 0x04, 0x13, 0x50, 0x90, 0xcf, 0x48, 0x50, 0x0f, 0xa6, 0x4d, 0x0f, 0xa8, 0xc5, 0x19,
	0x28, 0x11, 0x04, 0x96, 0x16, 0x92, 0xe0, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x95,
	0x60, 0x06, 0xaa, 0xe4, 0x0c, 0x82, 0x71, 0x85, 0x64, 0xb8, 0x38, 0x4a, 0x52, 0x2b, 0x4a, 0x5c,
	0x12, 0x4b, 0x12, 0x25, 0x58, 0x40, 0x52, 0x1e, 0x0c, 0x41, 0x70, 0x11, 0x27, 0x36, 0x2e, 0x96,
	0x14, 0x20, 0xad, 0x95, 0xcb, 0xc5, 0x0e, 0x35, 0x50, 0x88, 0x83, 0x8b, 0x25, 0x20, 0x33, 0x2f,
	0x5d, 0x80, 0x01, 0xcc, 0xca, 0x07, 0xb2, 0x18, 0x85, 0x04, 0xb8, 0xb8, 0x7d, 0x32, 0xf3, 0xb2,
	0x9d, 0x73, 0x32, 0x93, 0xb3, 0x53, 0x53, 0x04, 0x2e, 0xf0, 0x0b, 0x09, 0x73, 0xf1, 0x39, 0xe7,
	0xe7, 0x95, 0x24, 0x26, 0x97, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x5c, 0xe4, 0x17,
	0x12, 0xe2, 0xe2, 0x85, 0x0a, 0xfa, 0xe7, 0xe5, 0x64, 0xe6, 0xa5, 0x0a, 0x5c, 0x42, 0x56, 0xe8,
	0x9f, 0x96, 0x06, 0x16, 0xbc, 0xcc, 0x9f, 0xc4, 0x06, 0x0e, 0x04, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x5a, 0x5b, 0x02, 0x38, 0x16, 0x01, 0x00, 0x00,
}
