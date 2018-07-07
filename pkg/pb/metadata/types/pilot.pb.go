// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata/types/pilot.proto

package pbtypes // import "openpitrix.io/openpitrix/pkg/pb/metadata/types"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PilotConfig struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host"`
	ListenPort           int32    `protobuf:"varint,4,opt,name=listen_port,json=listenPort,proto3" json:"listen_port"`
	LogLevel             string   `protobuf:"bytes,5,opt,name=log_level,json=logLevel,proto3" json:"log_level"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PilotConfig) Reset()         { *m = PilotConfig{} }
func (m *PilotConfig) String() string { return proto.CompactTextString(m) }
func (*PilotConfig) ProtoMessage()    {}
func (*PilotConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_pilot_d962ee71cde2a4ba, []int{0}
}
func (m *PilotConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PilotConfig.Unmarshal(m, b)
}
func (m *PilotConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PilotConfig.Marshal(b, m, deterministic)
}
func (dst *PilotConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PilotConfig.Merge(dst, src)
}
func (m *PilotConfig) XXX_Size() int {
	return xxx_messageInfo_PilotConfig.Size(m)
}
func (m *PilotConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PilotConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PilotConfig proto.InternalMessageInfo

func (m *PilotConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PilotConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *PilotConfig) GetListenPort() int32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

func (m *PilotConfig) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

type PilotEndpoint struct {
	PilotId              string   `protobuf:"bytes,1,opt,name=pilot_id,json=pilotId,proto3" json:"pilot_id"`
	PilotHost            string   `protobuf:"bytes,2,opt,name=pilot_host,json=pilotHost,proto3" json:"pilot_host"`
	PilotPort            int32    `protobuf:"varint,3,opt,name=pilot_port,json=pilotPort,proto3" json:"pilot_port"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PilotEndpoint) Reset()         { *m = PilotEndpoint{} }
func (m *PilotEndpoint) String() string { return proto.CompactTextString(m) }
func (*PilotEndpoint) ProtoMessage()    {}
func (*PilotEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_pilot_d962ee71cde2a4ba, []int{1}
}
func (m *PilotEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PilotEndpoint.Unmarshal(m, b)
}
func (m *PilotEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PilotEndpoint.Marshal(b, m, deterministic)
}
func (dst *PilotEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PilotEndpoint.Merge(dst, src)
}
func (m *PilotEndpoint) XXX_Size() int {
	return xxx_messageInfo_PilotEndpoint.Size(m)
}
func (m *PilotEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_PilotEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_PilotEndpoint proto.InternalMessageInfo

func (m *PilotEndpoint) GetPilotId() string {
	if m != nil {
		return m.PilotId
	}
	return ""
}

func (m *PilotEndpoint) GetPilotHost() string {
	if m != nil {
		return m.PilotHost
	}
	return ""
}

func (m *PilotEndpoint) GetPilotPort() int32 {
	if m != nil {
		return m.PilotPort
	}
	return 0
}

func init() {
	proto.RegisterType((*PilotConfig)(nil), "metadata.types.PilotConfig")
	proto.RegisterType((*PilotEndpoint)(nil), "metadata.types.PilotEndpoint")
}

func init() { proto.RegisterFile("metadata/types/pilot.proto", fileDescriptor_pilot_d962ee71cde2a4ba) }

var fileDescriptor_pilot_d962ee71cde2a4ba = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0x6c, 0xb5, 0x99, 0x62, 0x0f, 0x7b, 0x8a, 0x8a, 0x58, 0x7a, 0xea, 0x29, 0x7b,
	0x10, 0x44, 0xf0, 0xa6, 0x08, 0x0a, 0x1e, 0x4a, 0x8f, 0x5e, 0x42, 0xc2, 0xae, 0xe9, 0xe0, 0xba,
	0x33, 0x24, 0x83, 0xe8, 0xbf, 0x97, 0x9d, 0x60, 0xb5, 0xa7, 0xdd, 0xf7, 0x3e, 0x86, 0xf9, 0x18,
	0x38, 0xff, 0xf0, 0xd2, 0xb8, 0x46, 0x1a, 0x2b, 0xdf, 0xec, 0x07, 0xcb, 0x18, 0x48, 0x2a, 0xee,
	0x49, 0xc8, 0x2c, 0x7e, 0x59, 0xa5, 0x6c, 0x45, 0x30, 0xdf, 0x24, 0xfc, 0x40, 0xf1, 0x0d, 0x3b,
	0xb3, 0x80, 0x1c, 0x5d, 0x99, 0x2d, 0xb3, 0x75, 0xb1, 0xcd, 0xd1, 0x19, 0x03, 0x93, 0x1d, 0x0d,
	0x52, 0xe6, 0xda, 0xe8, 0xdf, 0x5c, 0xc1, 0x3c, 0xe0, 0x20, 0x3e, 0xd6, 0x4c, 0xbd, 0x94, 0x93,
	0x65, 0xb6, 0x9e, 0x6e, 0x61, 0xac, 0x36, 0xd4, 0x8b, 0xb9, 0x80, 0x22, 0x50, 0x57, 0x07, 0xff,
	0xe9, 0x43, 0x39, 0xd5, 0xc9, 0x59, 0xa0, 0xee, 0x25, 0xe5, 0xd5, 0x0e, 0x4e, 0x75, 0xe1, 0x63,
	0x74, 0x4c, 0x18, 0xc5, 0x9c, 0xc1, 0x4c, 0x05, 0xeb, 0xfd, 0xe2, 0x13, 0xcd, 0xcf, 0xce, 0x5c,
	0x02, 0x8c, 0xe8, 0x9f, 0x43, 0xa1, 0xcd, 0x53, 0x12, 0xd9, 0x63, 0xf5, 0x38, 0x52, 0x8f, 0x11,
	0x27, 0x8d, 0xfb, 0xdb, 0xd7, 0x1b, 0x62, 0x1f, 0x19, 0xa5, 0xc7, 0xaf, 0x0a, 0xc9, 0xfe, 0x25,
	0xcb, 0xef, 0x9d, 0xe5, 0xd6, 0x1e, 0x1e, 0xea, 0x8e, 0x5b, 0x7d, 0xdb, 0x63, 0xbd, 0xd5, 0xf5,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x02, 0x35, 0xfc, 0x49, 0x01, 0x00, 0x00,
}
