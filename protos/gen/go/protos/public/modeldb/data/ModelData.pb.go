// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.2
// source: modeldb/data/ModelData.proto

package data

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ModelDataMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimestampMillis int64  `protobuf:"varint,1,opt,name=timestamp_millis,json=timestampMillis,proto3" json:"timestamp_millis,omitempty"`
	ModelId         string `protobuf:"bytes,2,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	Endpoint        string `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *ModelDataMetadata) Reset() {
	*x = ModelDataMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modeldb_data_ModelData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelDataMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelDataMetadata) ProtoMessage() {}

func (x *ModelDataMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_modeldb_data_ModelData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelDataMetadata.ProtoReflect.Descriptor instead.
func (*ModelDataMetadata) Descriptor() ([]byte, []int) {
	return file_modeldb_data_ModelData_proto_rawDescGZIP(), []int{0}
}

func (x *ModelDataMetadata) GetTimestampMillis() int64 {
	if x != nil {
		return x.TimestampMillis
	}
	return 0
}

func (x *ModelDataMetadata) GetModelId() string {
	if x != nil {
		return x.ModelId
	}
	return ""
}

func (x *ModelDataMetadata) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type ModelData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *ModelDataMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Data     string             `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"` // json
}

func (x *ModelData) Reset() {
	*x = ModelData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modeldb_data_ModelData_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelData) ProtoMessage() {}

func (x *ModelData) ProtoReflect() protoreflect.Message {
	mi := &file_modeldb_data_ModelData_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelData.ProtoReflect.Descriptor instead.
func (*ModelData) Descriptor() ([]byte, []int) {
	return file_modeldb_data_ModelData_proto_rawDescGZIP(), []int{1}
}

func (x *ModelData) GetMetadata() *ModelDataMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ModelData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type StoreModelDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelData *ModelData `protobuf:"bytes,1,opt,name=modelData,proto3" json:"modelData,omitempty"`
}

func (x *StoreModelDataRequest) Reset() {
	*x = StoreModelDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modeldb_data_ModelData_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreModelDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreModelDataRequest) ProtoMessage() {}

func (x *StoreModelDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modeldb_data_ModelData_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreModelDataRequest.ProtoReflect.Descriptor instead.
func (*StoreModelDataRequest) Descriptor() ([]byte, []int) {
	return file_modeldb_data_ModelData_proto_rawDescGZIP(), []int{2}
}

func (x *StoreModelDataRequest) GetModelData() *ModelData {
	if x != nil {
		return x.ModelData
	}
	return nil
}

type GetModelDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTimeMillis int64  `protobuf:"varint,1,opt,name=start_time_millis,json=startTimeMillis,proto3" json:"start_time_millis,omitempty"`
	EndTimeMillis   int64  `protobuf:"varint,2,opt,name=end_time_millis,json=endTimeMillis,proto3" json:"end_time_millis,omitempty"`
	ModelId         string `protobuf:"bytes,3,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	Endpoint        string `protobuf:"bytes,4,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *GetModelDataRequest) Reset() {
	*x = GetModelDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modeldb_data_ModelData_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModelDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModelDataRequest) ProtoMessage() {}

func (x *GetModelDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modeldb_data_ModelData_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModelDataRequest.ProtoReflect.Descriptor instead.
func (*GetModelDataRequest) Descriptor() ([]byte, []int) {
	return file_modeldb_data_ModelData_proto_rawDescGZIP(), []int{3}
}

func (x *GetModelDataRequest) GetStartTimeMillis() int64 {
	if x != nil {
		return x.StartTimeMillis
	}
	return 0
}

func (x *GetModelDataRequest) GetEndTimeMillis() int64 {
	if x != nil {
		return x.EndTimeMillis
	}
	return 0
}

func (x *GetModelDataRequest) GetModelId() string {
	if x != nil {
		return x.ModelId
	}
	return ""
}

func (x *GetModelDataRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type ModelDataDiffMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTimeMillis int64  `protobuf:"varint,1,opt,name=start_time_millis,json=startTimeMillis,proto3" json:"start_time_millis,omitempty"`
	EndTimeMillis   int64  `protobuf:"varint,2,opt,name=end_time_millis,json=endTimeMillis,proto3" json:"end_time_millis,omitempty"`
	ModelIdA        string `protobuf:"bytes,3,opt,name=model_id_a,json=modelIdA,proto3" json:"model_id_a,omitempty"`
	ModelIdB        string `protobuf:"bytes,4,opt,name=model_id_b,json=modelIdB,proto3" json:"model_id_b,omitempty"`
	Endpoint        string `protobuf:"bytes,5,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *ModelDataDiffMetadata) Reset() {
	*x = ModelDataDiffMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modeldb_data_ModelData_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelDataDiffMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelDataDiffMetadata) ProtoMessage() {}

func (x *ModelDataDiffMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_modeldb_data_ModelData_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelDataDiffMetadata.ProtoReflect.Descriptor instead.
func (*ModelDataDiffMetadata) Descriptor() ([]byte, []int) {
	return file_modeldb_data_ModelData_proto_rawDescGZIP(), []int{4}
}

func (x *ModelDataDiffMetadata) GetStartTimeMillis() int64 {
	if x != nil {
		return x.StartTimeMillis
	}
	return 0
}

func (x *ModelDataDiffMetadata) GetEndTimeMillis() int64 {
	if x != nil {
		return x.EndTimeMillis
	}
	return 0
}

func (x *ModelDataDiffMetadata) GetModelIdA() string {
	if x != nil {
		return x.ModelIdA
	}
	return ""
}

func (x *ModelDataDiffMetadata) GetModelIdB() string {
	if x != nil {
		return x.ModelIdB
	}
	return ""
}

func (x *ModelDataDiffMetadata) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type GetModelDataDiffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTimeMillis int64  `protobuf:"varint,1,opt,name=start_time_millis,json=startTimeMillis,proto3" json:"start_time_millis,omitempty"`
	EndTimeMillis   int64  `protobuf:"varint,2,opt,name=end_time_millis,json=endTimeMillis,proto3" json:"end_time_millis,omitempty"`
	ModelIdA        string `protobuf:"bytes,3,opt,name=model_id_a,json=modelIdA,proto3" json:"model_id_a,omitempty"`
	ModelIdB        string `protobuf:"bytes,4,opt,name=model_id_b,json=modelIdB,proto3" json:"model_id_b,omitempty"`
}

func (x *GetModelDataDiffRequest) Reset() {
	*x = GetModelDataDiffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modeldb_data_ModelData_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModelDataDiffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModelDataDiffRequest) ProtoMessage() {}

func (x *GetModelDataDiffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modeldb_data_ModelData_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModelDataDiffRequest.ProtoReflect.Descriptor instead.
func (*GetModelDataDiffRequest) Descriptor() ([]byte, []int) {
	return file_modeldb_data_ModelData_proto_rawDescGZIP(), []int{5}
}

func (x *GetModelDataDiffRequest) GetStartTimeMillis() int64 {
	if x != nil {
		return x.StartTimeMillis
	}
	return 0
}

func (x *GetModelDataDiffRequest) GetEndTimeMillis() int64 {
	if x != nil {
		return x.EndTimeMillis
	}
	return 0
}

func (x *GetModelDataDiffRequest) GetModelIdA() string {
	if x != nil {
		return x.ModelIdA
	}
	return ""
}

func (x *GetModelDataDiffRequest) GetModelIdB() string {
	if x != nil {
		return x.ModelIdB
	}
	return ""
}

type StoreModelDataRequest_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StoreModelDataRequest_Response) Reset() {
	*x = StoreModelDataRequest_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modeldb_data_ModelData_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreModelDataRequest_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreModelDataRequest_Response) ProtoMessage() {}

func (x *StoreModelDataRequest_Response) ProtoReflect() protoreflect.Message {
	mi := &file_modeldb_data_ModelData_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreModelDataRequest_Response.ProtoReflect.Descriptor instead.
func (*StoreModelDataRequest_Response) Descriptor() ([]byte, []int) {
	return file_modeldb_data_ModelData_proto_rawDescGZIP(), []int{2, 0}
}

type GetModelDataRequest_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"` // json
}

func (x *GetModelDataRequest_Response) Reset() {
	*x = GetModelDataRequest_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modeldb_data_ModelData_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModelDataRequest_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModelDataRequest_Response) ProtoMessage() {}

func (x *GetModelDataRequest_Response) ProtoReflect() protoreflect.Message {
	mi := &file_modeldb_data_ModelData_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModelDataRequest_Response.ProtoReflect.Descriptor instead.
func (*GetModelDataRequest_Response) Descriptor() ([]byte, []int) {
	return file_modeldb_data_ModelData_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetModelDataRequest_Response) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type GetModelDataDiffRequest_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"` // json
}

func (x *GetModelDataDiffRequest_Response) Reset() {
	*x = GetModelDataDiffRequest_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modeldb_data_ModelData_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModelDataDiffRequest_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModelDataDiffRequest_Response) ProtoMessage() {}

func (x *GetModelDataDiffRequest_Response) ProtoReflect() protoreflect.Message {
	mi := &file_modeldb_data_ModelData_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModelDataDiffRequest_Response.ProtoReflect.Descriptor instead.
func (*GetModelDataDiffRequest_Response) Descriptor() ([]byte, []int) {
	return file_modeldb_data_ModelData_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetModelDataDiffRequest_Response) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_modeldb_data_ModelData_proto protoreflect.FileDescriptor

var file_modeldb_data_ModelData_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15,
	0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x62,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d, 0x69, 0x6c,
	0x6c, 0x69, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x65, 0x0a, 0x09, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x44, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x69, 0x2e, 0x76,
	0x65, 0x72, 0x74, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x62, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x63, 0x0a, 0x15, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x09, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x62,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0a, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc0, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x11, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x69, 0x6c,
	0x6c, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x69, 0x6c, 0x6c,
	0x69, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc3, 0x01, 0x0a, 0x15, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x44, 0x69, 0x66, 0x66, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x69, 0x6c, 0x6c,
	0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x12, 0x1c, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x5f, 0x69, 0x64, 0x5f, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x49, 0x64, 0x41, 0x12, 0x1c, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69,
	0x64, 0x5f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x49, 0x64, 0x42, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22,
	0xc9, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61,
	0x44, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x12,
	0x1c, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x5f, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x41, 0x12, 0x1c, 0x0a,
	0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x5f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x42, 0x1a, 0x1e, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xe0, 0x03, 0x0a, 0x10,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x99, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x2c, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x64, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x35, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x64, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c,
	0x1a, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x3a, 0x01, 0x2a, 0x12, 0x8e, 0x01, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x2e,
	0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x62,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x61, 0x69, 0x2e, 0x76,
	0x65, 0x72, 0x74, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x62, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x67, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x9e, 0x01,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x44, 0x69,
	0x66, 0x66, 0x12, 0x2e, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x64, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x44, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x37, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x64, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x44, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x67, 0x65,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x44, 0x69, 0x66, 0x66, 0x42, 0x47,
	0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56,
	0x65, 0x72, 0x74, 0x61, 0x41, 0x49, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x62, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x64, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modeldb_data_ModelData_proto_rawDescOnce sync.Once
	file_modeldb_data_ModelData_proto_rawDescData = file_modeldb_data_ModelData_proto_rawDesc
)

func file_modeldb_data_ModelData_proto_rawDescGZIP() []byte {
	file_modeldb_data_ModelData_proto_rawDescOnce.Do(func() {
		file_modeldb_data_ModelData_proto_rawDescData = protoimpl.X.CompressGZIP(file_modeldb_data_ModelData_proto_rawDescData)
	})
	return file_modeldb_data_ModelData_proto_rawDescData
}

var file_modeldb_data_ModelData_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_modeldb_data_ModelData_proto_goTypes = []interface{}{
	(*ModelDataMetadata)(nil),                // 0: ai.verta.modeldb.data.ModelDataMetadata
	(*ModelData)(nil),                        // 1: ai.verta.modeldb.data.ModelData
	(*StoreModelDataRequest)(nil),            // 2: ai.verta.modeldb.data.StoreModelDataRequest
	(*GetModelDataRequest)(nil),              // 3: ai.verta.modeldb.data.GetModelDataRequest
	(*ModelDataDiffMetadata)(nil),            // 4: ai.verta.modeldb.data.ModelDataDiffMetadata
	(*GetModelDataDiffRequest)(nil),          // 5: ai.verta.modeldb.data.GetModelDataDiffRequest
	(*StoreModelDataRequest_Response)(nil),   // 6: ai.verta.modeldb.data.StoreModelDataRequest.Response
	(*GetModelDataRequest_Response)(nil),     // 7: ai.verta.modeldb.data.GetModelDataRequest.Response
	(*GetModelDataDiffRequest_Response)(nil), // 8: ai.verta.modeldb.data.GetModelDataDiffRequest.Response
}
var file_modeldb_data_ModelData_proto_depIdxs = []int32{
	0, // 0: ai.verta.modeldb.data.ModelData.metadata:type_name -> ai.verta.modeldb.data.ModelDataMetadata
	1, // 1: ai.verta.modeldb.data.StoreModelDataRequest.modelData:type_name -> ai.verta.modeldb.data.ModelData
	2, // 2: ai.verta.modeldb.data.ModelDataService.StoreModelData:input_type -> ai.verta.modeldb.data.StoreModelDataRequest
	3, // 3: ai.verta.modeldb.data.ModelDataService.GetModelData:input_type -> ai.verta.modeldb.data.GetModelDataRequest
	5, // 4: ai.verta.modeldb.data.ModelDataService.GetModelDataDiff:input_type -> ai.verta.modeldb.data.GetModelDataDiffRequest
	6, // 5: ai.verta.modeldb.data.ModelDataService.StoreModelData:output_type -> ai.verta.modeldb.data.StoreModelDataRequest.Response
	7, // 6: ai.verta.modeldb.data.ModelDataService.GetModelData:output_type -> ai.verta.modeldb.data.GetModelDataRequest.Response
	8, // 7: ai.verta.modeldb.data.ModelDataService.GetModelDataDiff:output_type -> ai.verta.modeldb.data.GetModelDataDiffRequest.Response
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_modeldb_data_ModelData_proto_init() }
func file_modeldb_data_ModelData_proto_init() {
	if File_modeldb_data_ModelData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modeldb_data_ModelData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelDataMetadata); i {
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
		file_modeldb_data_ModelData_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelData); i {
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
		file_modeldb_data_ModelData_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreModelDataRequest); i {
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
		file_modeldb_data_ModelData_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModelDataRequest); i {
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
		file_modeldb_data_ModelData_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelDataDiffMetadata); i {
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
		file_modeldb_data_ModelData_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModelDataDiffRequest); i {
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
		file_modeldb_data_ModelData_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreModelDataRequest_Response); i {
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
		file_modeldb_data_ModelData_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModelDataRequest_Response); i {
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
		file_modeldb_data_ModelData_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModelDataDiffRequest_Response); i {
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
			RawDescriptor: file_modeldb_data_ModelData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modeldb_data_ModelData_proto_goTypes,
		DependencyIndexes: file_modeldb_data_ModelData_proto_depIdxs,
		MessageInfos:      file_modeldb_data_ModelData_proto_msgTypes,
	}.Build()
	File_modeldb_data_ModelData_proto = out.File
	file_modeldb_data_ModelData_proto_rawDesc = nil
	file_modeldb_data_ModelData_proto_goTypes = nil
	file_modeldb_data_ModelData_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ModelDataServiceClient is the client API for ModelDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ModelDataServiceClient interface {
	StoreModelData(ctx context.Context, in *StoreModelDataRequest, opts ...grpc.CallOption) (*StoreModelDataRequest_Response, error)
	GetModelData(ctx context.Context, in *GetModelDataRequest, opts ...grpc.CallOption) (*GetModelDataRequest_Response, error)
	GetModelDataDiff(ctx context.Context, in *GetModelDataDiffRequest, opts ...grpc.CallOption) (*GetModelDataDiffRequest_Response, error)
}

type modelDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModelDataServiceClient(cc grpc.ClientConnInterface) ModelDataServiceClient {
	return &modelDataServiceClient{cc}
}

func (c *modelDataServiceClient) StoreModelData(ctx context.Context, in *StoreModelDataRequest, opts ...grpc.CallOption) (*StoreModelDataRequest_Response, error) {
	out := new(StoreModelDataRequest_Response)
	err := c.cc.Invoke(ctx, "/ai.verta.modeldb.data.ModelDataService/StoreModelData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelDataServiceClient) GetModelData(ctx context.Context, in *GetModelDataRequest, opts ...grpc.CallOption) (*GetModelDataRequest_Response, error) {
	out := new(GetModelDataRequest_Response)
	err := c.cc.Invoke(ctx, "/ai.verta.modeldb.data.ModelDataService/GetModelData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelDataServiceClient) GetModelDataDiff(ctx context.Context, in *GetModelDataDiffRequest, opts ...grpc.CallOption) (*GetModelDataDiffRequest_Response, error) {
	out := new(GetModelDataDiffRequest_Response)
	err := c.cc.Invoke(ctx, "/ai.verta.modeldb.data.ModelDataService/GetModelDataDiff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelDataServiceServer is the server API for ModelDataService service.
type ModelDataServiceServer interface {
	StoreModelData(context.Context, *StoreModelDataRequest) (*StoreModelDataRequest_Response, error)
	GetModelData(context.Context, *GetModelDataRequest) (*GetModelDataRequest_Response, error)
	GetModelDataDiff(context.Context, *GetModelDataDiffRequest) (*GetModelDataDiffRequest_Response, error)
}

// UnimplementedModelDataServiceServer can be embedded to have forward compatible implementations.
type UnimplementedModelDataServiceServer struct {
}

func (*UnimplementedModelDataServiceServer) StoreModelData(context.Context, *StoreModelDataRequest) (*StoreModelDataRequest_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreModelData not implemented")
}
func (*UnimplementedModelDataServiceServer) GetModelData(context.Context, *GetModelDataRequest) (*GetModelDataRequest_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelData not implemented")
}
func (*UnimplementedModelDataServiceServer) GetModelDataDiff(context.Context, *GetModelDataDiffRequest) (*GetModelDataDiffRequest_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelDataDiff not implemented")
}

func RegisterModelDataServiceServer(s *grpc.Server, srv ModelDataServiceServer) {
	s.RegisterService(&_ModelDataService_serviceDesc, srv)
}

func _ModelDataService_StoreModelData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreModelDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelDataServiceServer).StoreModelData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.verta.modeldb.data.ModelDataService/StoreModelData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelDataServiceServer).StoreModelData(ctx, req.(*StoreModelDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelDataService_GetModelData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelDataServiceServer).GetModelData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.verta.modeldb.data.ModelDataService/GetModelData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelDataServiceServer).GetModelData(ctx, req.(*GetModelDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelDataService_GetModelDataDiff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelDataDiffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelDataServiceServer).GetModelDataDiff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.verta.modeldb.data.ModelDataService/GetModelDataDiff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelDataServiceServer).GetModelDataDiff(ctx, req.(*GetModelDataDiffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ModelDataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.verta.modeldb.data.ModelDataService",
	HandlerType: (*ModelDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreModelData",
			Handler:    _ModelDataService_StoreModelData_Handler,
		},
		{
			MethodName: "GetModelData",
			Handler:    _ModelDataService_GetModelData_Handler,
		},
		{
			MethodName: "GetModelDataDiff",
			Handler:    _ModelDataService_GetModelDataDiff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modeldb/data/ModelData.proto",
}
