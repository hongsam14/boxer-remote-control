// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: stream.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_stream_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{0}
}

type DataChunk struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DataChunk) Reset() {
	*x = DataChunk{}
	mi := &file_stream_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataChunk) ProtoMessage() {}

func (x *DataChunk) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataChunk.ProtoReflect.Descriptor instead.
func (*DataChunk) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{1}
}

func (x *DataChunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Command struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Commandline   string                 `protobuf:"bytes,1,opt,name=commandline,proto3" json:"commandline,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Command) Reset() {
	*x = Command{}
	mi := &file_stream_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{2}
}

func (x *Command) GetCommandline() string {
	if x != nil {
		return x.Commandline
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReturnCode    int32                  `protobuf:"varint,1,opt,name=return_code,json=returnCode,proto3" json:"return_code,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_stream_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetReturnCode() int32 {
	if x != nil {
		return x.ReturnCode
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_stream_proto protoreflect.FileDescriptor

const file_stream_proto_rawDesc = "" +
	"\n" +
	"\fstream.proto\x12\x06stream\"\a\n" +
	"\x05Empty\"\x1f\n" +
	"\tDataChunk\x12\x12\n" +
	"\x04data\x18\x01 \x01(\fR\x04data\"+\n" +
	"\aCommand\x12 \n" +
	"\vcommandline\x18\x01 \x01(\tR\vcommandline\"E\n" +
	"\bResponse\x12\x1f\n" +
	"\vreturn_code\x18\x01 \x01(\x05R\n" +
	"returnCode\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage2\xba\x02\n" +
	"\bStreamer\x123\n" +
	"\n" +
	"UploadFile\x12\x11.stream.DataChunk\x1a\x10.stream.Response(\x01\x124\n" +
	"\fDownloadFile\x12\x0f.stream.Command\x1a\x11.stream.DataChunk0\x01\x12-\n" +
	"\bOpenFile\x12\x0f.stream.Command\x1a\x10.stream.Response\x12.\n" +
	"\tHeartBeat\x12\x0f.stream.Command\x1a\x10.stream.Response\x120\n" +
	"\vCommandline\x12\x0f.stream.Command\x1a\x10.stream.Response\x122\n" +
	"\fStreamFrames\x12\r.stream.Empty\x1a\x11.stream.DataChunk0\x01B\x04Z\x02./b\x06proto3"

var (
	file_stream_proto_rawDescOnce sync.Once
	file_stream_proto_rawDescData []byte
)

func file_stream_proto_rawDescGZIP() []byte {
	file_stream_proto_rawDescOnce.Do(func() {
		file_stream_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_stream_proto_rawDesc), len(file_stream_proto_rawDesc)))
	})
	return file_stream_proto_rawDescData
}

var file_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_stream_proto_goTypes = []any{
	(*Empty)(nil),     // 0: stream.Empty
	(*DataChunk)(nil), // 1: stream.DataChunk
	(*Command)(nil),   // 2: stream.Command
	(*Response)(nil),  // 3: stream.Response
}
var file_stream_proto_depIdxs = []int32{
	1, // 0: stream.Streamer.UploadFile:input_type -> stream.DataChunk
	2, // 1: stream.Streamer.DownloadFile:input_type -> stream.Command
	2, // 2: stream.Streamer.OpenFile:input_type -> stream.Command
	2, // 3: stream.Streamer.HeartBeat:input_type -> stream.Command
	2, // 4: stream.Streamer.Commandline:input_type -> stream.Command
	0, // 5: stream.Streamer.StreamFrames:input_type -> stream.Empty
	3, // 6: stream.Streamer.UploadFile:output_type -> stream.Response
	1, // 7: stream.Streamer.DownloadFile:output_type -> stream.DataChunk
	3, // 8: stream.Streamer.OpenFile:output_type -> stream.Response
	3, // 9: stream.Streamer.HeartBeat:output_type -> stream.Response
	3, // 10: stream.Streamer.Commandline:output_type -> stream.Response
	1, // 11: stream.Streamer.StreamFrames:output_type -> stream.DataChunk
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stream_proto_init() }
func file_stream_proto_init() {
	if File_stream_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_stream_proto_rawDesc), len(file_stream_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stream_proto_goTypes,
		DependencyIndexes: file_stream_proto_depIdxs,
		MessageInfos:      file_stream_proto_msgTypes,
	}.Build()
	File_stream_proto = out.File
	file_stream_proto_goTypes = nil
	file_stream_proto_depIdxs = nil
}
