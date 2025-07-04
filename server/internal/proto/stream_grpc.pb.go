// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: stream.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Streamer_UploadFile_FullMethodName   = "/stream.Streamer/UploadFile"
	Streamer_DownloadFile_FullMethodName = "/stream.Streamer/DownloadFile"
	Streamer_OpenFile_FullMethodName     = "/stream.Streamer/OpenFile"
	Streamer_HeartBeat_FullMethodName    = "/stream.Streamer/HeartBeat"
	Streamer_Commandline_FullMethodName  = "/stream.Streamer/Commandline"
	Streamer_StreamFrames_FullMethodName = "/stream.Streamer/StreamFrames"
)

// StreamerClient is the client API for Streamer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamerClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[DataChunk, Response], error)
	DownloadFile(ctx context.Context, in *Command, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DataChunk], error)
	OpenFile(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error)
	HeartBeat(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error)
	Commandline(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error)
	StreamFrames(ctx context.Context, in *Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DataChunk], error)
}

type streamerClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamerClient(cc grpc.ClientConnInterface) StreamerClient {
	return &streamerClient{cc}
}

func (c *streamerClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[DataChunk, Response], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Streamer_ServiceDesc.Streams[0], Streamer_UploadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DataChunk, Response]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Streamer_UploadFileClient = grpc.ClientStreamingClient[DataChunk, Response]

func (c *streamerClient) DownloadFile(ctx context.Context, in *Command, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DataChunk], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Streamer_ServiceDesc.Streams[1], Streamer_DownloadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Command, DataChunk]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Streamer_DownloadFileClient = grpc.ServerStreamingClient[DataChunk]

func (c *streamerClient) OpenFile(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Streamer_OpenFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamerClient) HeartBeat(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Streamer_HeartBeat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamerClient) Commandline(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Streamer_Commandline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamerClient) StreamFrames(ctx context.Context, in *Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DataChunk], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Streamer_ServiceDesc.Streams[2], Streamer_StreamFrames_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Empty, DataChunk]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Streamer_StreamFramesClient = grpc.ServerStreamingClient[DataChunk]

// StreamerServer is the server API for Streamer service.
// All implementations must embed UnimplementedStreamerServer
// for forward compatibility.
type StreamerServer interface {
	UploadFile(grpc.ClientStreamingServer[DataChunk, Response]) error
	DownloadFile(*Command, grpc.ServerStreamingServer[DataChunk]) error
	OpenFile(context.Context, *Command) (*Response, error)
	HeartBeat(context.Context, *Command) (*Response, error)
	Commandline(context.Context, *Command) (*Response, error)
	StreamFrames(*Empty, grpc.ServerStreamingServer[DataChunk]) error
	mustEmbedUnimplementedStreamerServer()
}

// UnimplementedStreamerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStreamerServer struct{}

func (UnimplementedStreamerServer) UploadFile(grpc.ClientStreamingServer[DataChunk, Response]) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedStreamerServer) DownloadFile(*Command, grpc.ServerStreamingServer[DataChunk]) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedStreamerServer) OpenFile(context.Context, *Command) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenFile not implemented")
}
func (UnimplementedStreamerServer) HeartBeat(context.Context, *Command) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (UnimplementedStreamerServer) Commandline(context.Context, *Command) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commandline not implemented")
}
func (UnimplementedStreamerServer) StreamFrames(*Empty, grpc.ServerStreamingServer[DataChunk]) error {
	return status.Errorf(codes.Unimplemented, "method StreamFrames not implemented")
}
func (UnimplementedStreamerServer) mustEmbedUnimplementedStreamerServer() {}
func (UnimplementedStreamerServer) testEmbeddedByValue()                  {}

// UnsafeStreamerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamerServer will
// result in compilation errors.
type UnsafeStreamerServer interface {
	mustEmbedUnimplementedStreamerServer()
}

func RegisterStreamerServer(s grpc.ServiceRegistrar, srv StreamerServer) {
	// If the following call pancis, it indicates UnimplementedStreamerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Streamer_ServiceDesc, srv)
}

func _Streamer_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamerServer).UploadFile(&grpc.GenericServerStream[DataChunk, Response]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Streamer_UploadFileServer = grpc.ClientStreamingServer[DataChunk, Response]

func _Streamer_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Command)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamerServer).DownloadFile(m, &grpc.GenericServerStream[Command, DataChunk]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Streamer_DownloadFileServer = grpc.ServerStreamingServer[DataChunk]

func _Streamer_OpenFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamerServer).OpenFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Streamer_OpenFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamerServer).OpenFile(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Streamer_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamerServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Streamer_HeartBeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamerServer).HeartBeat(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Streamer_Commandline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamerServer).Commandline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Streamer_Commandline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamerServer).Commandline(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Streamer_StreamFrames_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamerServer).StreamFrames(m, &grpc.GenericServerStream[Empty, DataChunk]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Streamer_StreamFramesServer = grpc.ServerStreamingServer[DataChunk]

// Streamer_ServiceDesc is the grpc.ServiceDesc for Streamer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Streamer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream.Streamer",
	HandlerType: (*StreamerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenFile",
			Handler:    _Streamer_OpenFile_Handler,
		},
		{
			MethodName: "HeartBeat",
			Handler:    _Streamer_HeartBeat_Handler,
		},
		{
			MethodName: "Commandline",
			Handler:    _Streamer_Commandline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _Streamer_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _Streamer_DownloadFile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamFrames",
			Handler:       _Streamer_StreamFrames_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stream.proto",
}
