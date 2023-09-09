// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: func_metric.proto

package pb

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

type InvokeMetricRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomFields map[string]string `protobuf:"bytes,1,rep,name=custom_fields,json=customFields,proto3" json:"custom_fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *InvokeMetricRequest) Reset() {
	*x = InvokeMetricRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_func_metric_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeMetricRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeMetricRequest) ProtoMessage() {}

func (x *InvokeMetricRequest) ProtoReflect() protoreflect.Message {
	mi := &file_func_metric_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeMetricRequest.ProtoReflect.Descriptor instead.
func (*InvokeMetricRequest) Descriptor() ([]byte, []int) {
	return file_func_metric_proto_rawDescGZIP(), []int{0}
}

func (x *InvokeMetricRequest) GetCustomFields() map[string]string {
	if x != nil {
		return x.CustomFields
	}
	return nil
}

type LocalMetricRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecondOverMetric string `protobuf:"bytes,1,opt,name=second_over_metric,json=secondOverMetric,proto3" json:"second_over_metric,omitempty"`
	LocalIp          string `protobuf:"bytes,2,opt,name=local_ip,json=localIp,proto3" json:"local_ip,omitempty"`
}

func (x *LocalMetricRequest) Reset() {
	*x = LocalMetricRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_func_metric_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalMetricRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalMetricRequest) ProtoMessage() {}

func (x *LocalMetricRequest) ProtoReflect() protoreflect.Message {
	mi := &file_func_metric_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalMetricRequest.ProtoReflect.Descriptor instead.
func (*LocalMetricRequest) Descriptor() ([]byte, []int) {
	return file_func_metric_proto_rawDescGZIP(), []int{1}
}

func (x *LocalMetricRequest) GetSecondOverMetric() string {
	if x != nil {
		return x.SecondOverMetric
	}
	return ""
}

func (x *LocalMetricRequest) GetLocalIp() string {
	if x != nil {
		return x.LocalIp
	}
	return ""
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_func_metric_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_func_metric_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_func_metric_proto_rawDescGZIP(), []int{2}
}

func (x *Reply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Reply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_func_metric_proto protoreflect.FileDescriptor

var file_func_metric_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x75, 0x6e, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xa6, 0x01, 0x0a, 0x13, 0x49, 0x6e, 0x76, 0x6f,
	0x6b, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x4e, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x76, 0x6f,
	0x6b, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a,
	0x3f, 0x0a, 0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x5d, 0x0a, 0x12, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x5f, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4f, 0x76, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x69, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x49, 0x70, 0x22,
	0x35, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x84, 0x01, 0x0a, 0x10, 0x46, 0x75, 0x6e, 0x63, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x38, 0x0a, 0x12, 0x49,
	0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_func_metric_proto_rawDescOnce sync.Once
	file_func_metric_proto_rawDescData = file_func_metric_proto_rawDesc
)

func file_func_metric_proto_rawDescGZIP() []byte {
	file_func_metric_proto_rawDescOnce.Do(func() {
		file_func_metric_proto_rawDescData = protoimpl.X.CompressGZIP(file_func_metric_proto_rawDescData)
	})
	return file_func_metric_proto_rawDescData
}

var file_func_metric_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_func_metric_proto_goTypes = []interface{}{
	(*InvokeMetricRequest)(nil), // 0: pb.InvokeMetricRequest
	(*LocalMetricRequest)(nil),  // 1: pb.LocalMetricRequest
	(*Reply)(nil),               // 2: pb.Reply
	nil,                         // 3: pb.InvokeMetricRequest.CustomFieldsEntry
}
var file_func_metric_proto_depIdxs = []int32{
	3, // 0: pb.InvokeMetricRequest.custom_fields:type_name -> pb.InvokeMetricRequest.CustomFieldsEntry
	0, // 1: pb.FuncMetricReport.InvokeMetricReport:input_type -> pb.InvokeMetricRequest
	1, // 2: pb.FuncMetricReport.ReportLocalMetric:input_type -> pb.LocalMetricRequest
	2, // 3: pb.FuncMetricReport.InvokeMetricReport:output_type -> pb.Reply
	2, // 4: pb.FuncMetricReport.ReportLocalMetric:output_type -> pb.Reply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_func_metric_proto_init() }
func file_func_metric_proto_init() {
	if File_func_metric_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_func_metric_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeMetricRequest); i {
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
		file_func_metric_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalMetricRequest); i {
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
		file_func_metric_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
			RawDescriptor: file_func_metric_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_func_metric_proto_goTypes,
		DependencyIndexes: file_func_metric_proto_depIdxs,
		MessageInfos:      file_func_metric_proto_msgTypes,
	}.Build()
	File_func_metric_proto = out.File
	file_func_metric_proto_rawDesc = nil
	file_func_metric_proto_goTypes = nil
	file_func_metric_proto_depIdxs = nil
}
