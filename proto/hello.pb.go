// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: proto/hello.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HelloRequestWithoutTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"Name" validate:"required"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *HelloRequestWithoutTime) Reset() {
	*x = HelloRequestWithoutTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequestWithoutTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequestWithoutTime) ProtoMessage() {}

func (x *HelloRequestWithoutTime) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequestWithoutTime.ProtoReflect.Descriptor instead.
func (*HelloRequestWithoutTime) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequestWithoutTime) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloResponseWithoutTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"greeting" validate:"required"
	Greeting string `protobuf:"bytes,1,opt,name=Greeting,proto3" json:"Greeting,omitempty"`
}

func (x *HelloResponseWithoutTime) Reset() {
	*x = HelloResponseWithoutTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponseWithoutTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponseWithoutTime) ProtoMessage() {}

func (x *HelloResponseWithoutTime) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponseWithoutTime.ProtoReflect.Descriptor instead.
func (*HelloResponseWithoutTime) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponseWithoutTime) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

type HelloRequestWithTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"Name" validate:"required"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// @inject_tag: json:"start" validate:"required"
	Start *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
}

func (x *HelloRequestWithTime) Reset() {
	*x = HelloRequestWithTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequestWithTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequestWithTime) ProtoMessage() {}

func (x *HelloRequestWithTime) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequestWithTime.ProtoReflect.Descriptor instead.
func (*HelloRequestWithTime) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{2}
}

func (x *HelloRequestWithTime) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloRequestWithTime) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

type HelloResponseWithTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"greeting" validate:"required"
	Greeting string `protobuf:"bytes,1,opt,name=Greeting,proto3" json:"Greeting,omitempty"`
	// @inject_tag: json:"end" validate:"required"
	End *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *HelloResponseWithTime) Reset() {
	*x = HelloResponseWithTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponseWithTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponseWithTime) ProtoMessage() {}

func (x *HelloResponseWithTime) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponseWithTime.ProtoReflect.Descriptor instead.
func (*HelloResponseWithTime) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{3}
}

func (x *HelloResponseWithTime) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

func (x *HelloResponseWithTime) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

var File_proto_hello_proto protoreflect.FileDescriptor

var file_proto_hello_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x17,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68,
	0x6f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x18, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x57, 0x69, 0x74, 0x68,
	0x6f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x22, 0x5c, 0x0a, 0x14, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x22, 0x61, 0x0a, 0x15, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x57, 0x69, 0x74, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x03, 0x65, 0x6e, 0x64, 0x32, 0xb2, 0x01, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x58,
	0x0a, 0x13, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x6f,
	0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x20, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x6f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x10, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x69, 0x74, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x2e, 0x53,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x1d, 0x2e, 0x53, 0x6f, 0x63,
	0x69, 0x61, 0x6c, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x57, 0x69, 0x74, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hello_proto_rawDescOnce sync.Once
	file_proto_hello_proto_rawDescData = file_proto_hello_proto_rawDesc
)

func file_proto_hello_proto_rawDescGZIP() []byte {
	file_proto_hello_proto_rawDescOnce.Do(func() {
		file_proto_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hello_proto_rawDescData)
	})
	return file_proto_hello_proto_rawDescData
}

var file_proto_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_hello_proto_goTypes = []interface{}{
	(*HelloRequestWithoutTime)(nil),  // 0: Social.HelloRequestWithoutTime
	(*HelloResponseWithoutTime)(nil), // 1: Social.HelloResponseWithoutTime
	(*HelloRequestWithTime)(nil),     // 2: Social.HelloRequestWithTime
	(*HelloResponseWithTime)(nil),    // 3: Social.HelloResponseWithTime
	(*timestamppb.Timestamp)(nil),    // 4: google.protobuf.Timestamp
}
var file_proto_hello_proto_depIdxs = []int32{
	4, // 0: Social.HelloRequestWithTime.start:type_name -> google.protobuf.Timestamp
	4, // 1: Social.HelloResponseWithTime.end:type_name -> google.protobuf.Timestamp
	0, // 2: Social.Hello.SayHelloWithoutTime:input_type -> Social.HelloRequestWithoutTime
	2, // 3: Social.Hello.SayHelloWithTime:input_type -> Social.HelloRequestWithTime
	1, // 4: Social.Hello.SayHelloWithoutTime:output_type -> Social.HelloResponseWithoutTime
	3, // 5: Social.Hello.SayHelloWithTime:output_type -> Social.HelloResponseWithTime
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_hello_proto_init() }
func file_proto_hello_proto_init() {
	if File_proto_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequestWithoutTime); i {
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
		file_proto_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponseWithoutTime); i {
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
		file_proto_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequestWithTime); i {
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
		file_proto_hello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponseWithTime); i {
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
			RawDescriptor: file_proto_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hello_proto_goTypes,
		DependencyIndexes: file_proto_hello_proto_depIdxs,
		MessageInfos:      file_proto_hello_proto_msgTypes,
	}.Build()
	File_proto_hello_proto = out.File
	file_proto_hello_proto_rawDesc = nil
	file_proto_hello_proto_goTypes = nil
	file_proto_hello_proto_depIdxs = nil
}
