// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata/types/types.proto

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Bool struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bool) Reset()         { *m = Bool{} }
func (m *Bool) String() string { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()    {}
func (*Bool) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{1}
}
func (m *Bool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bool.Unmarshal(m, b)
}
func (m *Bool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bool.Marshal(b, m, deterministic)
}
func (dst *Bool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bool.Merge(dst, src)
}
func (m *Bool) XXX_Size() int {
	return xxx_messageInfo_Bool.Size(m)
}
func (m *Bool) XXX_DiscardUnknown() {
	xxx_messageInfo_Bool.DiscardUnknown(m)
}

var xxx_messageInfo_Bool proto.InternalMessageInfo

func (m *Bool) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type BoolList struct {
	ValueList            []bool   `protobuf:"varint,1,rep,packed,name=value_list,json=valueList,proto3" json:"value_list"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoolList) Reset()         { *m = BoolList{} }
func (m *BoolList) String() string { return proto.CompactTextString(m) }
func (*BoolList) ProtoMessage()    {}
func (*BoolList) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{2}
}
func (m *BoolList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolList.Unmarshal(m, b)
}
func (m *BoolList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolList.Marshal(b, m, deterministic)
}
func (dst *BoolList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolList.Merge(dst, src)
}
func (m *BoolList) XXX_Size() int {
	return xxx_messageInfo_BoolList.Size(m)
}
func (m *BoolList) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolList.DiscardUnknown(m)
}

var xxx_messageInfo_BoolList proto.InternalMessageInfo

func (m *BoolList) GetValueList() []bool {
	if m != nil {
		return m.ValueList
	}
	return nil
}

type Int32 struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32) Reset()         { *m = Int32{} }
func (m *Int32) String() string { return proto.CompactTextString(m) }
func (*Int32) ProtoMessage()    {}
func (*Int32) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{3}
}
func (m *Int32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32.Unmarshal(m, b)
}
func (m *Int32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32.Marshal(b, m, deterministic)
}
func (dst *Int32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32.Merge(dst, src)
}
func (m *Int32) XXX_Size() int {
	return xxx_messageInfo_Int32.Size(m)
}
func (m *Int32) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32.DiscardUnknown(m)
}

var xxx_messageInfo_Int32 proto.InternalMessageInfo

func (m *Int32) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Int32List struct {
	ValueList            []int32  `protobuf:"varint,1,rep,packed,name=value_list,json=valueList,proto3" json:"value_list"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32List) Reset()         { *m = Int32List{} }
func (m *Int32List) String() string { return proto.CompactTextString(m) }
func (*Int32List) ProtoMessage()    {}
func (*Int32List) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{4}
}
func (m *Int32List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32List.Unmarshal(m, b)
}
func (m *Int32List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32List.Marshal(b, m, deterministic)
}
func (dst *Int32List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32List.Merge(dst, src)
}
func (m *Int32List) XXX_Size() int {
	return xxx_messageInfo_Int32List.Size(m)
}
func (m *Int32List) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32List.DiscardUnknown(m)
}

var xxx_messageInfo_Int32List proto.InternalMessageInfo

func (m *Int32List) GetValueList() []int32 {
	if m != nil {
		return m.ValueList
	}
	return nil
}

type UInt32 struct {
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UInt32) Reset()         { *m = UInt32{} }
func (m *UInt32) String() string { return proto.CompactTextString(m) }
func (*UInt32) ProtoMessage()    {}
func (*UInt32) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{5}
}
func (m *UInt32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UInt32.Unmarshal(m, b)
}
func (m *UInt32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UInt32.Marshal(b, m, deterministic)
}
func (dst *UInt32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UInt32.Merge(dst, src)
}
func (m *UInt32) XXX_Size() int {
	return xxx_messageInfo_UInt32.Size(m)
}
func (m *UInt32) XXX_DiscardUnknown() {
	xxx_messageInfo_UInt32.DiscardUnknown(m)
}

var xxx_messageInfo_UInt32 proto.InternalMessageInfo

func (m *UInt32) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type UInt32List struct {
	ValueList            []uint32 `protobuf:"varint,1,rep,packed,name=value_list,json=valueList,proto3" json:"value_list"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UInt32List) Reset()         { *m = UInt32List{} }
func (m *UInt32List) String() string { return proto.CompactTextString(m) }
func (*UInt32List) ProtoMessage()    {}
func (*UInt32List) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{6}
}
func (m *UInt32List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UInt32List.Unmarshal(m, b)
}
func (m *UInt32List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UInt32List.Marshal(b, m, deterministic)
}
func (dst *UInt32List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UInt32List.Merge(dst, src)
}
func (m *UInt32List) XXX_Size() int {
	return xxx_messageInfo_UInt32List.Size(m)
}
func (m *UInt32List) XXX_DiscardUnknown() {
	xxx_messageInfo_UInt32List.DiscardUnknown(m)
}

var xxx_messageInfo_UInt32List proto.InternalMessageInfo

func (m *UInt32List) GetValueList() []uint32 {
	if m != nil {
		return m.ValueList
	}
	return nil
}

type Int64 struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64) Reset()         { *m = Int64{} }
func (m *Int64) String() string { return proto.CompactTextString(m) }
func (*Int64) ProtoMessage()    {}
func (*Int64) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{7}
}
func (m *Int64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64.Unmarshal(m, b)
}
func (m *Int64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64.Marshal(b, m, deterministic)
}
func (dst *Int64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64.Merge(dst, src)
}
func (m *Int64) XXX_Size() int {
	return xxx_messageInfo_Int64.Size(m)
}
func (m *Int64) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64.DiscardUnknown(m)
}

var xxx_messageInfo_Int64 proto.InternalMessageInfo

func (m *Int64) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Int64List struct {
	ValueList            []int64  `protobuf:"varint,1,rep,packed,name=value_list,json=valueList,proto3" json:"value_list"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64List) Reset()         { *m = Int64List{} }
func (m *Int64List) String() string { return proto.CompactTextString(m) }
func (*Int64List) ProtoMessage()    {}
func (*Int64List) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{8}
}
func (m *Int64List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64List.Unmarshal(m, b)
}
func (m *Int64List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64List.Marshal(b, m, deterministic)
}
func (dst *Int64List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64List.Merge(dst, src)
}
func (m *Int64List) XXX_Size() int {
	return xxx_messageInfo_Int64List.Size(m)
}
func (m *Int64List) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64List.DiscardUnknown(m)
}

var xxx_messageInfo_Int64List proto.InternalMessageInfo

func (m *Int64List) GetValueList() []int64 {
	if m != nil {
		return m.ValueList
	}
	return nil
}

type UInt64 struct {
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UInt64) Reset()         { *m = UInt64{} }
func (m *UInt64) String() string { return proto.CompactTextString(m) }
func (*UInt64) ProtoMessage()    {}
func (*UInt64) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{9}
}
func (m *UInt64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UInt64.Unmarshal(m, b)
}
func (m *UInt64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UInt64.Marshal(b, m, deterministic)
}
func (dst *UInt64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UInt64.Merge(dst, src)
}
func (m *UInt64) XXX_Size() int {
	return xxx_messageInfo_UInt64.Size(m)
}
func (m *UInt64) XXX_DiscardUnknown() {
	xxx_messageInfo_UInt64.DiscardUnknown(m)
}

var xxx_messageInfo_UInt64 proto.InternalMessageInfo

func (m *UInt64) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type UInt64List struct {
	ValueList            []uint64 `protobuf:"varint,1,rep,packed,name=value_list,json=valueList,proto3" json:"value_list"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UInt64List) Reset()         { *m = UInt64List{} }
func (m *UInt64List) String() string { return proto.CompactTextString(m) }
func (*UInt64List) ProtoMessage()    {}
func (*UInt64List) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{10}
}
func (m *UInt64List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UInt64List.Unmarshal(m, b)
}
func (m *UInt64List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UInt64List.Marshal(b, m, deterministic)
}
func (dst *UInt64List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UInt64List.Merge(dst, src)
}
func (m *UInt64List) XXX_Size() int {
	return xxx_messageInfo_UInt64List.Size(m)
}
func (m *UInt64List) XXX_DiscardUnknown() {
	xxx_messageInfo_UInt64List.DiscardUnknown(m)
}

var xxx_messageInfo_UInt64List proto.InternalMessageInfo

func (m *UInt64List) GetValueList() []uint64 {
	if m != nil {
		return m.ValueList
	}
	return nil
}

type Float32 struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Float32) Reset()         { *m = Float32{} }
func (m *Float32) String() string { return proto.CompactTextString(m) }
func (*Float32) ProtoMessage()    {}
func (*Float32) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{11}
}
func (m *Float32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Float32.Unmarshal(m, b)
}
func (m *Float32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Float32.Marshal(b, m, deterministic)
}
func (dst *Float32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Float32.Merge(dst, src)
}
func (m *Float32) XXX_Size() int {
	return xxx_messageInfo_Float32.Size(m)
}
func (m *Float32) XXX_DiscardUnknown() {
	xxx_messageInfo_Float32.DiscardUnknown(m)
}

var xxx_messageInfo_Float32 proto.InternalMessageInfo

func (m *Float32) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Float32List struct {
	ValueList            []float32 `protobuf:"fixed32,1,rep,packed,name=value_list,json=valueList,proto3" json:"value_list"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Float32List) Reset()         { *m = Float32List{} }
func (m *Float32List) String() string { return proto.CompactTextString(m) }
func (*Float32List) ProtoMessage()    {}
func (*Float32List) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{12}
}
func (m *Float32List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Float32List.Unmarshal(m, b)
}
func (m *Float32List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Float32List.Marshal(b, m, deterministic)
}
func (dst *Float32List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Float32List.Merge(dst, src)
}
func (m *Float32List) XXX_Size() int {
	return xxx_messageInfo_Float32List.Size(m)
}
func (m *Float32List) XXX_DiscardUnknown() {
	xxx_messageInfo_Float32List.DiscardUnknown(m)
}

var xxx_messageInfo_Float32List proto.InternalMessageInfo

func (m *Float32List) GetValueList() []float32 {
	if m != nil {
		return m.ValueList
	}
	return nil
}

type Float64 struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Float64) Reset()         { *m = Float64{} }
func (m *Float64) String() string { return proto.CompactTextString(m) }
func (*Float64) ProtoMessage()    {}
func (*Float64) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{13}
}
func (m *Float64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Float64.Unmarshal(m, b)
}
func (m *Float64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Float64.Marshal(b, m, deterministic)
}
func (dst *Float64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Float64.Merge(dst, src)
}
func (m *Float64) XXX_Size() int {
	return xxx_messageInfo_Float64.Size(m)
}
func (m *Float64) XXX_DiscardUnknown() {
	xxx_messageInfo_Float64.DiscardUnknown(m)
}

var xxx_messageInfo_Float64 proto.InternalMessageInfo

func (m *Float64) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Float64List struct {
	ValueList            []float64 `protobuf:"fixed64,1,rep,packed,name=value_list,json=valueList,proto3" json:"value_list"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Float64List) Reset()         { *m = Float64List{} }
func (m *Float64List) String() string { return proto.CompactTextString(m) }
func (*Float64List) ProtoMessage()    {}
func (*Float64List) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{14}
}
func (m *Float64List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Float64List.Unmarshal(m, b)
}
func (m *Float64List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Float64List.Marshal(b, m, deterministic)
}
func (dst *Float64List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Float64List.Merge(dst, src)
}
func (m *Float64List) XXX_Size() int {
	return xxx_messageInfo_Float64List.Size(m)
}
func (m *Float64List) XXX_DiscardUnknown() {
	xxx_messageInfo_Float64List.DiscardUnknown(m)
}

var xxx_messageInfo_Float64List proto.InternalMessageInfo

func (m *Float64List) GetValueList() []float64 {
	if m != nil {
		return m.ValueList
	}
	return nil
}

type String struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{15}
}
func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (dst *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(dst, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type StringList struct {
	ValueList            []string `protobuf:"bytes,1,rep,name=value_list,json=valueList,proto3" json:"value_list"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringList) Reset()         { *m = StringList{} }
func (m *StringList) String() string { return proto.CompactTextString(m) }
func (*StringList) ProtoMessage()    {}
func (*StringList) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{16}
}
func (m *StringList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringList.Unmarshal(m, b)
}
func (m *StringList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringList.Marshal(b, m, deterministic)
}
func (dst *StringList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringList.Merge(dst, src)
}
func (m *StringList) XXX_Size() int {
	return xxx_messageInfo_StringList.Size(m)
}
func (m *StringList) XXX_DiscardUnknown() {
	xxx_messageInfo_StringList.DiscardUnknown(m)
}

var xxx_messageInfo_StringList proto.InternalMessageInfo

func (m *StringList) GetValueList() []string {
	if m != nil {
		return m.ValueList
	}
	return nil
}

type StringMap struct {
	ValueMap             map[string]string `protobuf:"bytes,1,rep,name=value_map,json=valueMap,proto3" json:"value_map" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StringMap) Reset()         { *m = StringMap{} }
func (m *StringMap) String() string { return proto.CompactTextString(m) }
func (*StringMap) ProtoMessage()    {}
func (*StringMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{17}
}
func (m *StringMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMap.Unmarshal(m, b)
}
func (m *StringMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMap.Marshal(b, m, deterministic)
}
func (dst *StringMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMap.Merge(dst, src)
}
func (m *StringMap) XXX_Size() int {
	return xxx_messageInfo_StringMap.Size(m)
}
func (m *StringMap) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMap.DiscardUnknown(m)
}

var xxx_messageInfo_StringMap proto.InternalMessageInfo

func (m *StringMap) GetValueMap() map[string]string {
	if m != nil {
		return m.ValueMap
	}
	return nil
}

type StringListMap struct {
	ValueListMap         map[string]*StringList `protobuf:"bytes,1,rep,name=value_list_map,json=valueListMap,proto3" json:"value_list_map" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *StringListMap) Reset()         { *m = StringListMap{} }
func (m *StringListMap) String() string { return proto.CompactTextString(m) }
func (*StringListMap) ProtoMessage()    {}
func (*StringListMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{18}
}
func (m *StringListMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringListMap.Unmarshal(m, b)
}
func (m *StringListMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringListMap.Marshal(b, m, deterministic)
}
func (dst *StringListMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringListMap.Merge(dst, src)
}
func (m *StringListMap) XXX_Size() int {
	return xxx_messageInfo_StringListMap.Size(m)
}
func (m *StringListMap) XXX_DiscardUnknown() {
	xxx_messageInfo_StringListMap.DiscardUnknown(m)
}

var xxx_messageInfo_StringListMap proto.InternalMessageInfo

func (m *StringListMap) GetValueListMap() map[string]*StringList {
	if m != nil {
		return m.ValueListMap
	}
	return nil
}

type Bytes struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bytes) Reset()         { *m = Bytes{} }
func (m *Bytes) String() string { return proto.CompactTextString(m) }
func (*Bytes) ProtoMessage()    {}
func (*Bytes) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{19}
}
func (m *Bytes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bytes.Unmarshal(m, b)
}
func (m *Bytes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bytes.Marshal(b, m, deterministic)
}
func (dst *Bytes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bytes.Merge(dst, src)
}
func (m *Bytes) XXX_Size() int {
	return xxx_messageInfo_Bytes.Size(m)
}
func (m *Bytes) XXX_DiscardUnknown() {
	xxx_messageInfo_Bytes.DiscardUnknown(m)
}

var xxx_messageInfo_Bytes proto.InternalMessageInfo

func (m *Bytes) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type BytesList struct {
	ValueList            [][]byte `protobuf:"bytes,1,rep,name=value_list,json=valueList,proto3" json:"value_list"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BytesList) Reset()         { *m = BytesList{} }
func (m *BytesList) String() string { return proto.CompactTextString(m) }
func (*BytesList) ProtoMessage()    {}
func (*BytesList) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{20}
}
func (m *BytesList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesList.Unmarshal(m, b)
}
func (m *BytesList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesList.Marshal(b, m, deterministic)
}
func (dst *BytesList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesList.Merge(dst, src)
}
func (m *BytesList) XXX_Size() int {
	return xxx_messageInfo_BytesList.Size(m)
}
func (m *BytesList) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesList.DiscardUnknown(m)
}

var xxx_messageInfo_BytesList proto.InternalMessageInfo

func (m *BytesList) GetValueList() [][]byte {
	if m != nil {
		return m.ValueList
	}
	return nil
}

type BytesMap struct {
	ValueMap             map[string][]byte `protobuf:"bytes,1,rep,name=value_map,json=valueMap,proto3" json:"value_map" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BytesMap) Reset()         { *m = BytesMap{} }
func (m *BytesMap) String() string { return proto.CompactTextString(m) }
func (*BytesMap) ProtoMessage()    {}
func (*BytesMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{21}
}
func (m *BytesMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesMap.Unmarshal(m, b)
}
func (m *BytesMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesMap.Marshal(b, m, deterministic)
}
func (dst *BytesMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesMap.Merge(dst, src)
}
func (m *BytesMap) XXX_Size() int {
	return xxx_messageInfo_BytesMap.Size(m)
}
func (m *BytesMap) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesMap.DiscardUnknown(m)
}

var xxx_messageInfo_BytesMap proto.InternalMessageInfo

func (m *BytesMap) GetValueMap() map[string][]byte {
	if m != nil {
		return m.ValueMap
	}
	return nil
}

type BytesListMap struct {
	ValueListMap         map[string]*BytesList `protobuf:"bytes,1,rep,name=value_list_map,json=valueListMap,proto3" json:"value_list_map" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BytesListMap) Reset()         { *m = BytesListMap{} }
func (m *BytesListMap) String() string { return proto.CompactTextString(m) }
func (*BytesListMap) ProtoMessage()    {}
func (*BytesListMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0dcb13a33242d03f, []int{22}
}
func (m *BytesListMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesListMap.Unmarshal(m, b)
}
func (m *BytesListMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesListMap.Marshal(b, m, deterministic)
}
func (dst *BytesListMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesListMap.Merge(dst, src)
}
func (m *BytesListMap) XXX_Size() int {
	return xxx_messageInfo_BytesListMap.Size(m)
}
func (m *BytesListMap) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesListMap.DiscardUnknown(m)
}

var xxx_messageInfo_BytesListMap proto.InternalMessageInfo

func (m *BytesListMap) GetValueListMap() map[string]*BytesList {
	if m != nil {
		return m.ValueListMap
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "metadata.types.Empty")
	proto.RegisterType((*Bool)(nil), "metadata.types.Bool")
	proto.RegisterType((*BoolList)(nil), "metadata.types.BoolList")
	proto.RegisterType((*Int32)(nil), "metadata.types.Int32")
	proto.RegisterType((*Int32List)(nil), "metadata.types.Int32List")
	proto.RegisterType((*UInt32)(nil), "metadata.types.UInt32")
	proto.RegisterType((*UInt32List)(nil), "metadata.types.UInt32List")
	proto.RegisterType((*Int64)(nil), "metadata.types.Int64")
	proto.RegisterType((*Int64List)(nil), "metadata.types.Int64List")
	proto.RegisterType((*UInt64)(nil), "metadata.types.UInt64")
	proto.RegisterType((*UInt64List)(nil), "metadata.types.UInt64List")
	proto.RegisterType((*Float32)(nil), "metadata.types.Float32")
	proto.RegisterType((*Float32List)(nil), "metadata.types.Float32List")
	proto.RegisterType((*Float64)(nil), "metadata.types.Float64")
	proto.RegisterType((*Float64List)(nil), "metadata.types.Float64List")
	proto.RegisterType((*String)(nil), "metadata.types.String")
	proto.RegisterType((*StringList)(nil), "metadata.types.StringList")
	proto.RegisterType((*StringMap)(nil), "metadata.types.StringMap")
	proto.RegisterMapType((map[string]string)(nil), "metadata.types.StringMap.ValueMapEntry")
	proto.RegisterType((*StringListMap)(nil), "metadata.types.StringListMap")
	proto.RegisterMapType((map[string]*StringList)(nil), "metadata.types.StringListMap.ValueListMapEntry")
	proto.RegisterType((*Bytes)(nil), "metadata.types.Bytes")
	proto.RegisterType((*BytesList)(nil), "metadata.types.BytesList")
	proto.RegisterType((*BytesMap)(nil), "metadata.types.BytesMap")
	proto.RegisterMapType((map[string][]byte)(nil), "metadata.types.BytesMap.ValueMapEntry")
	proto.RegisterType((*BytesListMap)(nil), "metadata.types.BytesListMap")
	proto.RegisterMapType((map[string]*BytesList)(nil), "metadata.types.BytesListMap.ValueListMapEntry")
}

func init() { proto.RegisterFile("metadata/types/types.proto", fileDescriptor_types_0dcb13a33242d03f) }

var fileDescriptor_types_0dcb13a33242d03f = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x8f, 0x93, 0x40,
	0x14, 0xc7, 0x33, 0x50, 0x76, 0xe1, 0x15, 0x36, 0x4a, 0x3c, 0xac, 0xc4, 0xea, 0x86, 0x83, 0xae,
	0x3f, 0x02, 0xa6, 0xbb, 0x69, 0x36, 0xee, 0xad, 0xba, 0x26, 0x26, 0xf6, 0x52, 0xad, 0x87, 0x7a,
	0x68, 0xa6, 0x91, 0x34, 0xa4, 0x14, 0x26, 0x30, 0x36, 0xf2, 0x3f, 0x18, 0xff, 0x26, 0x13, 0xff,
	0x31, 0xc3, 0x00, 0x65, 0x06, 0x18, 0xd3, 0x83, 0x97, 0xb6, 0xf3, 0xde, 0xeb, 0x87, 0xcf, 0xe3,
	0x4b, 0x00, 0x67, 0x17, 0x50, 0xfc, 0x0d, 0x53, 0xec, 0xd3, 0x9c, 0x04, 0x59, 0xf9, 0xe9, 0x91,
	0x34, 0xa1, 0x89, 0x7d, 0x56, 0xf7, 0x3c, 0x56, 0x75, 0x4f, 0x41, 0xbb, 0xdb, 0x11, 0x9a, 0xbb,
	0x8f, 0x60, 0x30, 0x4d, 0x92, 0xc8, 0x7e, 0x00, 0xda, 0x1e, 0x47, 0xdf, 0x83, 0x73, 0x74, 0x81,
	0x2e, 0xf5, 0x79, 0x79, 0x70, 0x9f, 0x83, 0x5e, 0x74, 0x3f, 0x86, 0x19, 0xb5, 0x47, 0x00, 0xac,
	0xb8, 0x8a, 0xc2, 0x8c, 0x9e, 0xa3, 0x0b, 0xf5, 0x52, 0x9f, 0x1b, 0xac, 0x52, 0xb4, 0xdd, 0x11,
	0x68, 0x1f, 0x62, 0x7a, 0x35, 0x16, 0x49, 0x5a, 0x4d, 0x7a, 0x01, 0x06, 0x6b, 0x4b, 0x50, 0x1a,
	0x8f, 0x7a, 0x0c, 0x27, 0x8b, 0x1e, 0x96, 0x55, 0xb3, 0x5e, 0x02, 0x2c, 0xfe, 0x05, 0xb3, 0xba,
	0x5e, 0x93, 0x6b, 0x91, 0xa5, 0x8a, 0x5e, 0x93, 0x6b, 0x09, 0x4a, 0xed, 0xf1, 0x6a, 0xb3, 0x06,
	0x2d, 0x2f, 0x29, 0x6c, 0xc0, 0xc3, 0x9e, 0xc0, 0xe9, 0xfb, 0x28, 0xc1, 0x9d, 0x2d, 0x95, 0x9a,
	0xf6, 0x0a, 0x86, 0xd5, 0x80, 0x04, 0xa7, 0xf4, 0xe1, 0xda, 0x72, 0xa8, 0x8d, 0x93, 0xda, 0xa1,
	0xd6, 0xaa, 0x9f, 0x68, 0x1a, 0xc6, 0x1b, 0x91, 0x66, 0x70, 0xab, 0x96, 0x7d, 0x09, 0xcc, 0xe0,
	0x61, 0xbf, 0x10, 0x18, 0xe5, 0xf4, 0x0c, 0x13, 0xfb, 0x1d, 0x94, 0xad, 0xd5, 0x0e, 0x13, 0x36,
	0x3b, 0x1c, 0x3f, 0xf3, 0xc4, 0xc7, 0xd3, 0x3b, 0x4c, 0x7b, 0x5f, 0x8a, 0xd1, 0x19, 0x26, 0x77,
	0x31, 0x4d, 0xf3, 0xb9, 0xbe, 0xaf, 0x8e, 0xce, 0x2d, 0x58, 0x42, 0xcb, 0xbe, 0x07, 0xea, 0x36,
	0xc8, 0x2b, 0xcb, 0xe2, 0x67, 0x63, 0xae, 0x70, 0xe6, 0x6f, 0x94, 0x1b, 0xe4, 0xfe, 0x41, 0x60,
	0x35, 0xfa, 0x85, 0xd4, 0x02, 0xce, 0x9a, 0x0d, 0x38, 0x33, 0xbf, 0xdf, 0xac, 0xfa, 0x5b, 0x69,
	0x57, 0x1d, 0x4a, 0x43, 0x73, 0xcf, 0x95, 0x9c, 0xaf, 0x70, 0xbf, 0x33, 0xd2, 0x63, 0xfa, 0x9a,
	0x37, 0x1d, 0x8e, 0x1d, 0xf9, 0x45, 0xf9, 0x2d, 0x46, 0xa0, 0x4d, 0x73, 0x1a, 0x64, 0x62, 0x44,
	0x26, 0xf7, 0x64, 0xb3, 0xb6, 0x24, 0x21, 0x93, 0x4f, 0xe8, 0x27, 0x02, 0x9d, 0x0d, 0x17, 0xf7,
	0xe2, 0x6d, 0x37, 0xa0, 0xa7, 0x6d, 0xa3, 0x7a, 0xf8, 0xbf, 0xe5, 0x63, 0xf2, 0x9b, 0xfd, 0x46,
	0x60, 0x1e, 0xdc, 0x0b, 0xa5, 0xcf, 0x92, 0x78, 0xbc, 0x5e, 0xaf, 0x63, 0xd3, 0x59, 0x1e, 0x97,
	0x8e, 0x2f, 0xa6, 0xf3, 0x50, 0x7a, 0x4d, 0x6e, 0x85, 0xe9, 0xcd, 0x72, 0x92, 0x90, 0x20, 0x26,
	0x21, 0x4d, 0xc3, 0x1f, 0x5e, 0x98, 0xf8, 0xcd, 0xc9, 0x27, 0xdb, 0x8d, 0x4f, 0xd6, 0xbe, 0xf8,
	0xba, 0xbe, 0x25, 0x6b, 0xf6, 0xbd, 0x3e, 0x61, 0x6f, 0xec, 0xab, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x96, 0x58, 0xe0, 0xb4, 0xcf, 0x05, 0x00, 0x00,
}
