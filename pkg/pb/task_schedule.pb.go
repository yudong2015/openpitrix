// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task_schedule.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type ScheduleTask struct {
	TaskId               *wrappers.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Handler              *wrappers.StringValue `protobuf:"bytes,2,opt,name=handler,proto3" json:"handler,omitempty"`
	Action               *wrappers.StringValue `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Conf                 *wrappers.StringValue `protobuf:"bytes,4,opt,name=conf,proto3" json:"conf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ScheduleTask) Reset()         { *m = ScheduleTask{} }
func (m *ScheduleTask) String() string { return proto.CompactTextString(m) }
func (*ScheduleTask) ProtoMessage()    {}
func (*ScheduleTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_838bf5d950b53e21, []int{0}
}

func (m *ScheduleTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduleTask.Unmarshal(m, b)
}
func (m *ScheduleTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduleTask.Marshal(b, m, deterministic)
}
func (m *ScheduleTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleTask.Merge(m, src)
}
func (m *ScheduleTask) XXX_Size() int {
	return xxx_messageInfo_ScheduleTask.Size(m)
}
func (m *ScheduleTask) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleTask.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleTask proto.InternalMessageInfo

func (m *ScheduleTask) GetTaskId() *wrappers.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *ScheduleTask) GetHandler() *wrappers.StringValue {
	if m != nil {
		return m.Handler
	}
	return nil
}

func (m *ScheduleTask) GetAction() *wrappers.StringValue {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *ScheduleTask) GetConf() *wrappers.StringValue {
	if m != nil {
		return m.Conf
	}
	return nil
}

type CreateScheduleTaskRequest struct {
	Handler              *wrappers.StringValue `protobuf:"bytes,2,opt,name=handler,proto3" json:"handler,omitempty"`
	Action               *wrappers.StringValue `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Conf                 *wrappers.StringValue `protobuf:"bytes,4,opt,name=conf,proto3" json:"conf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateScheduleTaskRequest) Reset()         { *m = CreateScheduleTaskRequest{} }
func (m *CreateScheduleTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateScheduleTaskRequest) ProtoMessage()    {}
func (*CreateScheduleTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_838bf5d950b53e21, []int{1}
}

func (m *CreateScheduleTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateScheduleTaskRequest.Unmarshal(m, b)
}
func (m *CreateScheduleTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateScheduleTaskRequest.Marshal(b, m, deterministic)
}
func (m *CreateScheduleTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateScheduleTaskRequest.Merge(m, src)
}
func (m *CreateScheduleTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateScheduleTaskRequest.Size(m)
}
func (m *CreateScheduleTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateScheduleTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateScheduleTaskRequest proto.InternalMessageInfo

func (m *CreateScheduleTaskRequest) GetHandler() *wrappers.StringValue {
	if m != nil {
		return m.Handler
	}
	return nil
}

func (m *CreateScheduleTaskRequest) GetAction() *wrappers.StringValue {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *CreateScheduleTaskRequest) GetConf() *wrappers.StringValue {
	if m != nil {
		return m.Conf
	}
	return nil
}

type CreateScheduleTaskResponse struct {
	TaskId               *wrappers.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateScheduleTaskResponse) Reset()         { *m = CreateScheduleTaskResponse{} }
func (m *CreateScheduleTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateScheduleTaskResponse) ProtoMessage()    {}
func (*CreateScheduleTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_838bf5d950b53e21, []int{2}
}

func (m *CreateScheduleTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateScheduleTaskResponse.Unmarshal(m, b)
}
func (m *CreateScheduleTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateScheduleTaskResponse.Marshal(b, m, deterministic)
}
func (m *CreateScheduleTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateScheduleTaskResponse.Merge(m, src)
}
func (m *CreateScheduleTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateScheduleTaskResponse.Size(m)
}
func (m *CreateScheduleTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateScheduleTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateScheduleTaskResponse proto.InternalMessageInfo

func (m *CreateScheduleTaskResponse) GetTaskId() *wrappers.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func init() {
	proto.RegisterType((*ScheduleTask)(nil), "openpitrix.ScheduleTask")
	proto.RegisterType((*CreateScheduleTaskRequest)(nil), "openpitrix.CreateScheduleTaskRequest")
	proto.RegisterType((*CreateScheduleTaskResponse)(nil), "openpitrix.CreateScheduleTaskResponse")
}

func init() { proto.RegisterFile("task_schedule.proto", fileDescriptor_838bf5d950b53e21) }

var fileDescriptor_838bf5d950b53e21 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x87, 0xdd, 0x5a, 0x5a, 0x18, 0x3d, 0xa5, 0x97, 0x75, 0x29, 0x22, 0x05, 0xc5, 0x83, 0xcd,
	0x4a, 0xfd, 0xf3, 0x00, 0x7a, 0xf2, 0xe0, 0xa5, 0x15, 0x0f, 0x82, 0xc8, 0x74, 0x77, 0x9a, 0x86,
	0x2e, 0x49, 0x4c, 0xb2, 0xd4, 0x77, 0xf3, 0x79, 0x7c, 0x0f, 0xd9, 0x6c, 0x17, 0x05, 0x15, 0x17,
	0x4f, 0x9e, 0x72, 0x98, 0x6f, 0xe6, 0xf7, 0x25, 0x19, 0x18, 0x78, 0x74, 0xab, 0x27, 0x97, 0x2d,
	0x29, 0x2f, 0x0b, 0xe2, 0xc6, 0x6a, 0xaf, 0x19, 0x68, 0x43, 0xca, 0x48, 0x6f, 0xe5, 0x4b, 0xb2,
	0x2f, 0xb4, 0x16, 0x05, 0xa5, 0xa1, 0x32, 0x2f, 0x17, 0xe9, 0xda, 0xa2, 0x31, 0x64, 0x5d, 0xcd,
	0x26, 0xc3, 0x4d, 0x1d, 0x8d, 0x4c, 0x51, 0x29, 0xed, 0xd1, 0x4b, 0xad, 0x9a, 0xea, 0x49, 0x38,
	0xb2, 0xb1, 0x20, 0x35, 0x76, 0x6b, 0x14, 0x82, 0x6c, 0xaa, 0x4d, 0x20, 0xbe, 0xd2, 0xa3, 0xb7,
	0x08, 0x76, 0x67, 0x1b, 0x95, 0x3b, 0x74, 0x2b, 0x76, 0x01, 0xfd, 0xe0, 0x27, 0xf3, 0x38, 0x3a,
	0x88, 0x8e, 0x77, 0x26, 0x43, 0x5e, 0xc7, 0xf1, 0x46, 0x87, 0xcf, 0xbc, 0x95, 0x4a, 0xdc, 0x63,
	0x51, 0xd2, 0xb4, 0x57, 0xc1, 0x37, 0x39, 0xbb, 0x84, 0xfe, 0x12, 0x55, 0x5e, 0x90, 0x8d, 0x3b,
	0x2d, 0xda, 0x1a, 0x98, 0x9d, 0x43, 0x0f, 0xb3, 0x4a, 0x28, 0xde, 0x6e, 0x93, 0x56, 0xb3, 0xec,
	0x14, 0xba, 0x99, 0x56, 0x8b, 0xb8, 0xdb, 0xa2, 0x27, 0x90, 0xa3, 0xd7, 0x08, 0xf6, 0xae, 0x2d,
	0xa1, 0xa7, 0xcf, 0xb7, 0x9d, 0xd2, 0x73, 0x49, 0xce, 0xff, 0x7b, 0xfb, 0x19, 0x24, 0xdf, 0xc9,
	0x3b, 0xa3, 0x95, 0xa3, 0x3f, 0x7e, 0xd9, 0xc4, 0xc3, 0xa0, 0x1a, 0xd3, 0x8c, 0xbc, 0x45, 0x85,
	0x82, 0x2c, 0x7b, 0x04, 0xa8, 0xb3, 0xc2, 0x3a, 0x1c, 0xf2, 0x8f, 0xc5, 0xe4, 0x3f, 0x3e, 0x60,
	0x72, 0xf4, 0x1b, 0x56, 0xab, 0x8e, 0xb6, 0xae, 0xba, 0x0f, 0x1d, 0x33, 0x9f, 0xf7, 0x82, 0xd9,
	0xd9, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0xc6, 0xcc, 0xce, 0x0c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskScheduleManagerClient is the client API for TaskScheduleManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskScheduleManagerClient interface {
	CreateTask(ctx context.Context, in *CreateScheduleTaskRequest, opts ...grpc.CallOption) (*CreateScheduleTaskResponse, error)
}

type taskScheduleManagerClient struct {
	cc *grpc.ClientConn
}

func NewTaskScheduleManagerClient(cc *grpc.ClientConn) TaskScheduleManagerClient {
	return &taskScheduleManagerClient{cc}
}

func (c *taskScheduleManagerClient) CreateTask(ctx context.Context, in *CreateScheduleTaskRequest, opts ...grpc.CallOption) (*CreateScheduleTaskResponse, error) {
	out := new(CreateScheduleTaskResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.TaskScheduleManager/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskScheduleManagerServer is the server API for TaskScheduleManager service.
type TaskScheduleManagerServer interface {
	CreateTask(context.Context, *CreateScheduleTaskRequest) (*CreateScheduleTaskResponse, error)
}

func RegisterTaskScheduleManagerServer(s *grpc.Server, srv TaskScheduleManagerServer) {
	s.RegisterService(&_TaskScheduleManager_serviceDesc, srv)
}

func _TaskScheduleManager_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScheduleTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskScheduleManagerServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.TaskScheduleManager/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskScheduleManagerServer).CreateTask(ctx, req.(*CreateScheduleTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskScheduleManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.TaskScheduleManager",
	HandlerType: (*TaskScheduleManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskScheduleManager_CreateTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task_schedule.proto",
}
