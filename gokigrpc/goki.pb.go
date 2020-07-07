// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.1
// source: goki.proto

package gokigrpc

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GokiGRPCRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Args    []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *GokiGRPCRequest) Reset() {
	*x = GokiGRPCRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goki_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GokiGRPCRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GokiGRPCRequest) ProtoMessage() {}

func (x *GokiGRPCRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goki_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GokiGRPCRequest.ProtoReflect.Descriptor instead.
func (*GokiGRPCRequest) Descriptor() ([]byte, []int) {
	return file_goki_proto_rawDescGZIP(), []int{0}
}

func (x *GokiGRPCRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *GokiGRPCRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

type GokiGRPCResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    string               `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *GokiGRPCResponse) Reset() {
	*x = GokiGRPCResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goki_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GokiGRPCResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GokiGRPCResponse) ProtoMessage() {}

func (x *GokiGRPCResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goki_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GokiGRPCResponse.ProtoReflect.Descriptor instead.
func (*GokiGRPCResponse) Descriptor() ([]byte, []int) {
	return file_goki_proto_rawDescGZIP(), []int{1}
}

func (x *GokiGRPCResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *GokiGRPCResponse) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type PingPong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *PingPong) Reset() {
	*x = PingPong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goki_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingPong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingPong) ProtoMessage() {}

func (x *PingPong) ProtoReflect() protoreflect.Message {
	mi := &file_goki_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingPong.ProtoReflect.Descriptor instead.
func (*PingPong) Descriptor() ([]byte, []int) {
	return file_goki_proto_rawDescGZIP(), []int{2}
}

func (x *PingPong) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_goki_proto protoreflect.FileDescriptor

var file_goki_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x6f, 0x6b, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x6f,
	0x6b, 0x69, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x0f, 0x47, 0x6f, 0x6b, 0x69, 0x47,
	0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x65, 0x0a, 0x10, 0x47, 0x6f, 0x6b, 0x69,
	0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x1e, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x32,
	0x83, 0x01, 0x0a, 0x08, 0x47, 0x6f, 0x6b, 0x69, 0x47, 0x52, 0x50, 0x43, 0x12, 0x45, 0x0a, 0x0a,
	0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x19, 0x2e, 0x67, 0x6f, 0x6b,
	0x69, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x6f, 0x6b, 0x69, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x6b, 0x69, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x47, 0x6f, 0x6b, 0x69, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x67, 0x6f,
	0x6b, 0x69, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x1a,
	0x12, 0x2e, 0x67, 0x6f, 0x6b, 0x69, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x50,
	0x6f, 0x6e, 0x67, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x67, 0x6f, 0x6b, 0x69, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goki_proto_rawDescOnce sync.Once
	file_goki_proto_rawDescData = file_goki_proto_rawDesc
)

func file_goki_proto_rawDescGZIP() []byte {
	file_goki_proto_rawDescOnce.Do(func() {
		file_goki_proto_rawDescData = protoimpl.X.CompressGZIP(file_goki_proto_rawDescData)
	})
	return file_goki_proto_rawDescData
}

var file_goki_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_goki_proto_goTypes = []interface{}{
	(*GokiGRPCRequest)(nil),     // 0: gokigrpc.GokiGRPCRequest
	(*GokiGRPCResponse)(nil),    // 1: gokigrpc.GokiGRPCResponse
	(*PingPong)(nil),            // 2: gokigrpc.PingPong
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_goki_proto_depIdxs = []int32{
	3, // 0: gokigrpc.GokiGRPCResponse.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: gokigrpc.GokiGRPC.RunCommand:input_type -> gokigrpc.GokiGRPCRequest
	2, // 2: gokigrpc.GokiGRPC.Ping:input_type -> gokigrpc.PingPong
	1, // 3: gokigrpc.GokiGRPC.RunCommand:output_type -> gokigrpc.GokiGRPCResponse
	2, // 4: gokigrpc.GokiGRPC.Ping:output_type -> gokigrpc.PingPong
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_goki_proto_init() }
func file_goki_proto_init() {
	if File_goki_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goki_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GokiGRPCRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_goki_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GokiGRPCResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_goki_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingPong); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_goki_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goki_proto_goTypes,
		DependencyIndexes: file_goki_proto_depIdxs,
		MessageInfos:      file_goki_proto_msgTypes,
	}.Build()
	File_goki_proto = out.File
	file_goki_proto_rawDesc = nil
	file_goki_proto_goTypes = nil
	file_goki_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GokiGRPCClient is the client API for GokiGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GokiGRPCClient interface {
	RunCommand(ctx context.Context, in *GokiGRPCRequest, opts ...grpc.CallOption) (*GokiGRPCResponse, error)
	Ping(ctx context.Context, in *PingPong, opts ...grpc.CallOption) (*PingPong, error)
}

type gokiGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewGokiGRPCClient(cc grpc.ClientConnInterface) GokiGRPCClient {
	return &gokiGRPCClient{cc}
}

func (c *gokiGRPCClient) RunCommand(ctx context.Context, in *GokiGRPCRequest, opts ...grpc.CallOption) (*GokiGRPCResponse, error) {
	out := new(GokiGRPCResponse)
	err := c.cc.Invoke(ctx, "/gokigrpc.GokiGRPC/RunCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gokiGRPCClient) Ping(ctx context.Context, in *PingPong, opts ...grpc.CallOption) (*PingPong, error) {
	out := new(PingPong)
	err := c.cc.Invoke(ctx, "/gokigrpc.GokiGRPC/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GokiGRPCServer is the server API for GokiGRPC service.
type GokiGRPCServer interface {
	RunCommand(context.Context, *GokiGRPCRequest) (*GokiGRPCResponse, error)
	Ping(context.Context, *PingPong) (*PingPong, error)
}

// UnimplementedGokiGRPCServer can be embedded to have forward compatible implementations.
type UnimplementedGokiGRPCServer struct {
}

func (*UnimplementedGokiGRPCServer) RunCommand(context.Context, *GokiGRPCRequest) (*GokiGRPCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCommand not implemented")
}
func (*UnimplementedGokiGRPCServer) Ping(context.Context, *PingPong) (*PingPong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterGokiGRPCServer(s *grpc.Server, srv GokiGRPCServer) {
	s.RegisterService(&_GokiGRPC_serviceDesc, srv)
}

func _GokiGRPC_RunCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GokiGRPCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GokiGRPCServer).RunCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gokigrpc.GokiGRPC/RunCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GokiGRPCServer).RunCommand(ctx, req.(*GokiGRPCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GokiGRPC_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingPong)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GokiGRPCServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gokigrpc.GokiGRPC/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GokiGRPCServer).Ping(ctx, req.(*PingPong))
	}
	return interceptor(ctx, in, info, handler)
}

var _GokiGRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gokigrpc.GokiGRPC",
	HandlerType: (*GokiGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunCommand",
			Handler:    _GokiGRPC_RunCommand_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _GokiGRPC_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goki.proto",
}
