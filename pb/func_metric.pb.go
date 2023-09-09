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

	RequestId     string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	AppId         string `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	FunctionId    string `protobuf:"bytes,3,opt,name=function_id,json=functionId,proto3" json:"function_id,omitempty"`
	Namespace     string `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	FunctionName  string `protobuf:"bytes,5,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	Qualifier     string `protobuf:"bytes,6,opt,name=qualifier,proto3" json:"qualifier,omitempty"`
	Runtime       string `protobuf:"bytes,7,opt,name=runtime,proto3" json:"runtime,omitempty"`
	CurQuota      string `protobuf:"bytes,8,opt,name=cur_quota,json=curQuota,proto3" json:"cur_quota,omitempty"` //实时并发
	ContainerId   string `protobuf:"bytes,9,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	StatusCode    string `protobuf:"bytes,10,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	RetCode       string `protobuf:"bytes,11,opt,name=ret_code,json=retCode,proto3" json:"ret_code,omitempty"`
	ColdStart     bool   `protobuf:"varint,12,opt,name=cold_start,json=coldStart,proto3" json:"cold_start,omitempty"`
	WorkerMgrCost string `protobuf:"bytes,13,opt,name=worker_mgr_cost,json=workerMgrCost,proto3" json:"worker_mgr_cost,omitempty"`
	WorkerCost    string `protobuf:"bytes,14,opt,name=worker_cost,json=workerCost,proto3" json:"worker_cost,omitempty"`
	Duration      string `protobuf:"bytes,15,opt,name=duration,proto3" json:"duration,omitempty"`
	EventType     string `protobuf:"bytes,16,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	Region        string `protobuf:"bytes,17,opt,name=region,proto3" json:"region,omitempty"`
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

func (x *InvokeMetricRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *InvokeMetricRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *InvokeMetricRequest) GetFunctionId() string {
	if x != nil {
		return x.FunctionId
	}
	return ""
}

func (x *InvokeMetricRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *InvokeMetricRequest) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *InvokeMetricRequest) GetQualifier() string {
	if x != nil {
		return x.Qualifier
	}
	return ""
}

func (x *InvokeMetricRequest) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *InvokeMetricRequest) GetCurQuota() string {
	if x != nil {
		return x.CurQuota
	}
	return ""
}

func (x *InvokeMetricRequest) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *InvokeMetricRequest) GetStatusCode() string {
	if x != nil {
		return x.StatusCode
	}
	return ""
}

func (x *InvokeMetricRequest) GetRetCode() string {
	if x != nil {
		return x.RetCode
	}
	return ""
}

func (x *InvokeMetricRequest) GetColdStart() bool {
	if x != nil {
		return x.ColdStart
	}
	return false
}

func (x *InvokeMetricRequest) GetWorkerMgrCost() string {
	if x != nil {
		return x.WorkerMgrCost
	}
	return ""
}

func (x *InvokeMetricRequest) GetWorkerCost() string {
	if x != nil {
		return x.WorkerCost
	}
	return ""
}

func (x *InvokeMetricRequest) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *InvokeMetricRequest) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *InvokeMetricRequest) GetRegion() string {
	if x != nil {
		return x.Region
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
		mi := &file_func_metric_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_func_metric_proto_rawDescGZIP(), []int{1}
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
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x9e, 0x04, 0x0a, 0x13, 0x49, 0x6e, 0x76, 0x6f,
	0x6b, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x61,
	0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75,
	0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6f, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x63, 0x6f, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x26, 0x0a, 0x0f,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x6d, 0x67, 0x72, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4d, 0x67, 0x72,
	0x43, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x4c, 0x0a, 0x10, 0x46, 0x75, 0x6e, 0x63, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x38, 0x0a, 0x12, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x49,
	0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
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

var file_func_metric_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_func_metric_proto_goTypes = []interface{}{
	(*InvokeMetricRequest)(nil), // 0: pb.InvokeMetricRequest
	(*Reply)(nil),               // 1: pb.Reply
}
var file_func_metric_proto_depIdxs = []int32{
	0, // 0: pb.FuncMetricReport.InvokeMetricReport:input_type -> pb.InvokeMetricRequest
	1, // 1: pb.FuncMetricReport.InvokeMetricReport:output_type -> pb.Reply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
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
			NumMessages:   2,
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