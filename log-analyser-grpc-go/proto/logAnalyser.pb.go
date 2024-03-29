// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: logAnalyser.proto

package proto

import (
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

type GetMessageInIntervalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date    string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Time    string `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Delta   string `protobuf:"bytes,3,opt,name=delta,proto3" json:"delta,omitempty"`
	Pattern string `protobuf:"bytes,4,opt,name=pattern,proto3" json:"pattern,omitempty"`
}

func (x *GetMessageInIntervalRequest) Reset() {
	*x = GetMessageInIntervalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logAnalyser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageInIntervalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageInIntervalRequest) ProtoMessage() {}

func (x *GetMessageInIntervalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logAnalyser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageInIntervalRequest.ProtoReflect.Descriptor instead.
func (*GetMessageInIntervalRequest) Descriptor() ([]byte, []int) {
	return file_logAnalyser_proto_rawDescGZIP(), []int{0}
}

func (x *GetMessageInIntervalRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *GetMessageInIntervalRequest) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *GetMessageInIntervalRequest) GetDelta() string {
	if x != nil {
		return x.Delta
	}
	return ""
}

func (x *GetMessageInIntervalRequest) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

type GetMessageInIntervalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Logs []*Log `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *GetMessageInIntervalResponse) Reset() {
	*x = GetMessageInIntervalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logAnalyser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageInIntervalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageInIntervalResponse) ProtoMessage() {}

func (x *GetMessageInIntervalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logAnalyser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageInIntervalResponse.ProtoReflect.Descriptor instead.
func (*GetMessageInIntervalResponse) Descriptor() ([]byte, []int) {
	return file_logAnalyser_proto_rawDescGZIP(), []int{1}
}

func (x *GetMessageInIntervalResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *GetMessageInIntervalResponse) GetLogs() []*Log {
	if x != nil {
		return x.Logs
	}
	return nil
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time        string `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	MessageHash string `protobuf:"bytes,2,opt,name=messageHash,proto3" json:"messageHash,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logAnalyser_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_logAnalyser_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_logAnalyser_proto_rawDescGZIP(), []int{2}
}

func (x *Log) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Log) GetMessageHash() string {
	if x != nil {
		return x.MessageHash
	}
	return ""
}

var File_logAnalyser_proto protoreflect.FileDescriptor

var file_logAnalyser_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6c, 0x6f, 0x67, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6c, 0x6f, 0x67, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x72,
	0x22, 0x75, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x22, 0x58, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6c,
	0x6f, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6c, 0x6f, 0x67, 0x41,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67,
	0x73, 0x22, 0x3b, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x61, 0x73, 0x68, 0x32, 0x83,
	0x01, 0x0a, 0x12, 0x6c, 0x6f, 0x67, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x28, 0x2e,
	0x6c, 0x6f, 0x67, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6c, 0x6f, 0x67, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x0a, 0x1b, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x42, 0x0f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logAnalyser_proto_rawDescOnce sync.Once
	file_logAnalyser_proto_rawDescData = file_logAnalyser_proto_rawDesc
)

func file_logAnalyser_proto_rawDescGZIP() []byte {
	file_logAnalyser_proto_rawDescOnce.Do(func() {
		file_logAnalyser_proto_rawDescData = protoimpl.X.CompressGZIP(file_logAnalyser_proto_rawDescData)
	})
	return file_logAnalyser_proto_rawDescData
}

var file_logAnalyser_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_logAnalyser_proto_goTypes = []interface{}{
	(*GetMessageInIntervalRequest)(nil),  // 0: logAnalyser.GetMessageInIntervalRequest
	(*GetMessageInIntervalResponse)(nil), // 1: logAnalyser.GetMessageInIntervalResponse
	(*Log)(nil),                          // 2: logAnalyser.Log
}
var file_logAnalyser_proto_depIdxs = []int32{
	2, // 0: logAnalyser.GetMessageInIntervalResponse.logs:type_name -> logAnalyser.Log
	0, // 1: logAnalyser.logAnalyserService.GetMessageInInterval:input_type -> logAnalyser.GetMessageInIntervalRequest
	1, // 2: logAnalyser.logAnalyserService.GetMessageInInterval:output_type -> logAnalyser.GetMessageInIntervalResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_logAnalyser_proto_init() }
func file_logAnalyser_proto_init() {
	if File_logAnalyser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logAnalyser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageInIntervalRequest); i {
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
		file_logAnalyser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageInIntervalResponse); i {
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
		file_logAnalyser_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
			RawDescriptor: file_logAnalyser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logAnalyser_proto_goTypes,
		DependencyIndexes: file_logAnalyser_proto_depIdxs,
		MessageInfos:      file_logAnalyser_proto_msgTypes,
	}.Build()
	File_logAnalyser_proto = out.File
	file_logAnalyser_proto_rawDesc = nil
	file_logAnalyser_proto_goTypes = nil
	file_logAnalyser_proto_depIdxs = nil
}
