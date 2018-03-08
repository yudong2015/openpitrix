// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pilot.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SendTaskRequest struct {
	TaskId     *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	TaskAction *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=task_action,json=taskAction" json:"task_action,omitempty"`
	Directive  *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=directive" json:"directive,omitempty"`
}

func (m *SendTaskRequest) Reset()                    { *m = SendTaskRequest{} }
func (m *SendTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*SendTaskRequest) ProtoMessage()               {}
func (*SendTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *SendTaskRequest) GetTaskId() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *SendTaskRequest) GetTaskAction() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskAction
	}
	return nil
}

func (m *SendTaskRequest) GetDirective() *google_protobuf.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

type SendTaskResponse struct {
	TaskId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
}

func (m *SendTaskResponse) Reset()                    { *m = SendTaskResponse{} }
func (m *SendTaskResponse) String() string            { return proto.CompactTextString(m) }
func (*SendTaskResponse) ProtoMessage()               {}
func (*SendTaskResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *SendTaskResponse) GetTaskId() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

type GetTaskStatusRequest struct {
	TaskId []string `protobuf:"bytes,1,rep,name=task_id,json=taskId" json:"task_id,omitempty"`
}

func (m *GetTaskStatusRequest) Reset()                    { *m = GetTaskStatusRequest{} }
func (m *GetTaskStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTaskStatusRequest) ProtoMessage()               {}
func (*GetTaskStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *GetTaskStatusRequest) GetTaskId() []string {
	if m != nil {
		return m.TaskId
	}
	return nil
}

type TaskStatus struct {
	TaskId     *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	TaskStatus *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=task_status,json=taskStatus" json:"task_status,omitempty"`
}

func (m *TaskStatus) Reset()                    { *m = TaskStatus{} }
func (m *TaskStatus) String() string            { return proto.CompactTextString(m) }
func (*TaskStatus) ProtoMessage()               {}
func (*TaskStatus) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *TaskStatus) GetTaskId() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *TaskStatus) GetTaskStatus() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskStatus
	}
	return nil
}

type GetTaskStatusResponse struct {
	TotalCount    uint32        `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	TaskStatusSet []*TaskStatus `protobuf:"bytes,2,rep,name=task_status_set,json=taskStatusSet" json:"task_status_set,omitempty"`
}

func (m *GetTaskStatusResponse) Reset()                    { *m = GetTaskStatusResponse{} }
func (m *GetTaskStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*GetTaskStatusResponse) ProtoMessage()               {}
func (*GetTaskStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *GetTaskStatusResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *GetTaskStatusResponse) GetTaskStatusSet() []*TaskStatus {
	if m != nil {
		return m.TaskStatusSet
	}
	return nil
}

func init() {
	proto.RegisterType((*SendTaskRequest)(nil), "openpitrix.SendTaskRequest")
	proto.RegisterType((*SendTaskResponse)(nil), "openpitrix.SendTaskResponse")
	proto.RegisterType((*GetTaskStatusRequest)(nil), "openpitrix.GetTaskStatusRequest")
	proto.RegisterType((*TaskStatus)(nil), "openpitrix.TaskStatus")
	proto.RegisterType((*GetTaskStatusResponse)(nil), "openpitrix.GetTaskStatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PilotManager service

type PilotManagerClient interface {
	SendTask(ctx context.Context, in *SendTaskRequest, opts ...grpc.CallOption) (*SendTaskResponse, error)
	GetTaskStatus(ctx context.Context, in *GetTaskStatusRequest, opts ...grpc.CallOption) (*GetTaskStatusResponse, error)
}

type pilotManagerClient struct {
	cc *grpc.ClientConn
}

func NewPilotManagerClient(cc *grpc.ClientConn) PilotManagerClient {
	return &pilotManagerClient{cc}
}

func (c *pilotManagerClient) SendTask(ctx context.Context, in *SendTaskRequest, opts ...grpc.CallOption) (*SendTaskResponse, error) {
	out := new(SendTaskResponse)
	err := grpc.Invoke(ctx, "/openpitrix.PilotManager/SendTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotManagerClient) GetTaskStatus(ctx context.Context, in *GetTaskStatusRequest, opts ...grpc.CallOption) (*GetTaskStatusResponse, error) {
	out := new(GetTaskStatusResponse)
	err := grpc.Invoke(ctx, "/openpitrix.PilotManager/GetTaskStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PilotManager service

type PilotManagerServer interface {
	SendTask(context.Context, *SendTaskRequest) (*SendTaskResponse, error)
	GetTaskStatus(context.Context, *GetTaskStatusRequest) (*GetTaskStatusResponse, error)
}

func RegisterPilotManagerServer(s *grpc.Server, srv PilotManagerServer) {
	s.RegisterService(&_PilotManager_serviceDesc, srv)
}

func _PilotManager_SendTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotManagerServer).SendTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.PilotManager/SendTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotManagerServer).SendTask(ctx, req.(*SendTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotManager_GetTaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotManagerServer).GetTaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.PilotManager/GetTaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotManagerServer).GetTaskStatus(ctx, req.(*GetTaskStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PilotManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.PilotManager",
	HandlerType: (*PilotManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTask",
			Handler:    _PilotManager_SendTask_Handler,
		},
		{
			MethodName: "GetTaskStatus",
			Handler:    _PilotManager_GetTaskStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pilot.proto",
}

func init() { proto.RegisterFile("pilot.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x53, 0x51, 0xc8, 0xb8, 0x51, 0xca, 0x8a, 0x42, 0x14, 0x22, 0x30, 0x46, 0x42, 0x51,
	0xa1, 0xde, 0x36, 0x88, 0x4b, 0x24, 0x90, 0x02, 0x07, 0xe8, 0x01, 0x11, 0x25, 0x88, 0x03, 0x97,
	0x68, 0xe3, 0x0c, 0x66, 0xd3, 0x68, 0x77, 0xeb, 0x9d, 0xb4, 0x39, 0x70, 0xea, 0x27, 0x84, 0x4f,
	0xe2, 0x13, 0xf8, 0x05, 0x3e, 0x82, 0x23, 0xf2, 0x26, 0xc1, 0xa6, 0x8d, 0x50, 0x04, 0x27, 0x4b,
	0x33, 0x6f, 0xde, 0x7b, 0xf3, 0x3c, 0x0b, 0xbe, 0x91, 0x13, 0x4d, 0x91, 0x49, 0x35, 0x69, 0x06,
	0xda, 0xa0, 0x32, 0x92, 0x52, 0x39, 0xab, 0xdf, 0x4b, 0xb4, 0x4e, 0x26, 0xc8, 0x5d, 0x67, 0x38,
	0xfd, 0xc4, 0xcf, 0x53, 0x61, 0x0c, 0xa6, 0x76, 0x81, 0xad, 0x37, 0x96, 0x7d, 0x61, 0x24, 0x17,
	0x4a, 0x69, 0x12, 0x24, 0xb5, 0x5a, 0x75, 0x9f, 0xb8, 0x4f, 0x7c, 0x90, 0xa0, 0x3a, 0xb0, 0xe7,
	0x22, 0x49, 0x30, 0xe5, 0xda, 0x38, 0xc4, 0x55, 0x74, 0xf8, 0xcd, 0x83, 0x6a, 0x1f, 0xd5, 0xe8,
	0xbd, 0xb0, 0x27, 0x3d, 0x3c, 0x9d, 0xa2, 0x25, 0xf6, 0x0c, 0xae, 0x93, 0xb0, 0x27, 0x03, 0x39,
	0xaa, 0x79, 0x81, 0xd7, 0xf4, 0x5b, 0x8d, 0x68, 0xa1, 0x18, 0xad, 0x1c, 0x45, 0x7d, 0x4a, 0xa5,
	0x4a, 0x3e, 0x88, 0xc9, 0x14, 0x7b, 0xdb, 0x19, 0xf8, 0x78, 0xc4, 0x9e, 0x83, 0xef, 0xc6, 0x44,
	0x9c, 0x09, 0xd4, 0x4a, 0x1b, 0x8c, 0x42, 0x36, 0xd0, 0x71, 0x78, 0xd6, 0x86, 0xf2, 0x48, 0xa6,
	0x18, 0x93, 0x3c, 0xc3, 0xda, 0xd6, 0x06, 0xc3, 0x39, 0x3c, 0x3c, 0x86, 0xdd, 0x7c, 0x09, 0x6b,
	0xb4, 0xb2, 0xf8, 0x8f, 0x5b, 0x84, 0x1c, 0x6e, 0xbd, 0x46, 0xca, 0x98, 0xfa, 0x24, 0x68, 0x6a,
	0x57, 0xa1, 0xdc, 0x29, 0xd2, 0x6d, 0x35, 0xcb, 0xbf, 0x07, 0x2e, 0x3c, 0x80, 0x1c, 0xfe, 0xbf,
	0xe1, 0x59, 0xc7, 0xb2, 0x79, 0x78, 0x0b, 0xd5, 0x70, 0x06, 0x7b, 0x97, 0x5c, 0x2f, 0x53, 0xb8,
	0x0f, 0x3e, 0x69, 0x12, 0x93, 0x41, 0xac, 0xa7, 0x8a, 0x9c, 0xa5, 0x4a, 0x0f, 0x5c, 0xe9, 0x55,
	0x56, 0x61, 0x2f, 0xa0, 0x5a, 0x10, 0x1e, 0x58, 0xa4, 0x5a, 0x29, 0xd8, 0x6a, 0xfa, 0xad, 0xdb,
	0x51, 0x7e, 0x92, 0x51, 0x81, 0xb9, 0x92, 0xcb, 0xf6, 0x91, 0x5a, 0x3f, 0x3d, 0xd8, 0xe9, 0x66,
	0x87, 0xfc, 0x56, 0x28, 0x91, 0x60, 0xca, 0xc6, 0x70, 0x63, 0xf5, 0x2f, 0xd8, 0xdd, 0x22, 0xc7,
	0xa5, 0x33, 0xab, 0x37, 0xd6, 0x37, 0x17, 0xc6, 0xc3, 0x87, 0xf3, 0x8e, 0xcf, 0xca, 0x16, 0xd5,
	0x28, 0xc8, 0x04, 0x2f, 0xbe, 0xff, 0xf8, 0x5a, 0xaa, 0x86, 0xc0, 0xcf, 0x8e, 0xb8, 0x7b, 0x37,
	0xb6, 0xed, 0xed, 0xb3, 0x2f, 0x50, 0xf9, 0x63, 0x6d, 0x16, 0x14, 0x39, 0xd7, 0xfd, 0xc7, 0xfa,
	0x83, 0xbf, 0x20, 0x96, 0xd2, 0x8f, 0xe6, 0x9d, 0x9b, 0xac, 0x9a, 0x20, 0x39, 0xe5, 0x60, 0x91,
	0x8c, 0x33, 0xb0, 0xc3, 0x0a, 0x06, 0x5e, 0xce, 0xe6, 0x9d, 0x53, 0xf6, 0x06, 0xd8, 0x3b, 0x83,
	0xaa, 0xeb, 0x08, 0x83, 0x6e, 0xaa, 0xc7, 0x18, 0x53, 0xf8, 0x78, 0x5d, 0x95, 0xed, 0x7d, 0x26,
	0x32, 0xb6, 0xcd, 0x79, 0xc1, 0x82, 0xd4, 0xad, 0x6b, 0x87, 0xd1, 0x61, 0x74, 0xb4, 0xef, 0x79,
	0xad, 0x5d, 0x61, 0xcc, 0x44, 0xc6, 0xee, 0x85, 0xf2, 0xb1, 0xd5, 0xaa, 0x7d, 0xa5, 0xf2, 0xb1,
	0x64, 0x86, 0xc3, 0x6d, 0x77, 0x10, 0x4f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x89, 0xc0,
	0xb0, 0x43, 0x04, 0x00, 0x00,
}
