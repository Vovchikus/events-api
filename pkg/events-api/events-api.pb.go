// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: api/events-api/v1/events-api.proto

package events_api_v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []string `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *CreateEventRequest) Reset() {
	*x = CreateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_events_api_v1_events_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventRequest) ProtoMessage() {}

func (x *CreateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_events_api_v1_events_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventRequest.ProtoReflect.Descriptor instead.
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return file_api_events_api_v1_events_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEventRequest) GetEvents() []string {
	if x != nil {
		return x.Events
	}
	return nil
}

type CreateEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CreateEventResponse) Reset() {
	*x = CreateEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_events_api_v1_events_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventResponse) ProtoMessage() {}

func (x *CreateEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_events_api_v1_events_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventResponse.ProtoReflect.Descriptor instead.
func (*CreateEventResponse) Descriptor() ([]byte, []int) {
	return file_api_events_api_v1_events_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEventResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type Health struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uptime               int64   `protobuf:"varint,1,opt,name=Uptime,proto3" json:"Uptime,omitempty"`
	AllocatedMemory      float32 `protobuf:"fixed32,2,opt,name=AllocatedMemory,proto3" json:"AllocatedMemory,omitempty"`
	TotalAllocatedMemory float32 `protobuf:"fixed32,3,opt,name=TotalAllocatedMemory,proto3" json:"TotalAllocatedMemory,omitempty"`
	Goroutines           int32   `protobuf:"varint,4,opt,name=Goroutines,proto3" json:"Goroutines,omitempty"`
	GCCycles             uint32  `protobuf:"varint,5,opt,name=GCCycles,proto3" json:"GCCycles,omitempty"`
	NumberOfCPUs         int32   `protobuf:"varint,6,opt,name=NumberOfCPUs,proto3" json:"NumberOfCPUs,omitempty"`
	HeapSys              float32 `protobuf:"fixed32,7,opt,name=HeapSys,proto3" json:"HeapSys,omitempty"`
	HeapAllocated        float32 `protobuf:"fixed32,8,opt,name=HeapAllocated,proto3" json:"HeapAllocated,omitempty"`
	ObjectsInUse         uint64  `protobuf:"varint,9,opt,name=ObjectsInUse,proto3" json:"ObjectsInUse,omitempty"`
	OSMemoryObtained     float32 `protobuf:"fixed32,10,opt,name=OSMemoryObtained,proto3" json:"OSMemoryObtained,omitempty"`
}

func (x *Health) Reset() {
	*x = Health{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_events_api_v1_events_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Health) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Health) ProtoMessage() {}

func (x *Health) ProtoReflect() protoreflect.Message {
	mi := &file_api_events_api_v1_events_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Health.ProtoReflect.Descriptor instead.
func (*Health) Descriptor() ([]byte, []int) {
	return file_api_events_api_v1_events_api_proto_rawDescGZIP(), []int{2}
}

func (x *Health) GetUptime() int64 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

func (x *Health) GetAllocatedMemory() float32 {
	if x != nil {
		return x.AllocatedMemory
	}
	return 0
}

func (x *Health) GetTotalAllocatedMemory() float32 {
	if x != nil {
		return x.TotalAllocatedMemory
	}
	return 0
}

func (x *Health) GetGoroutines() int32 {
	if x != nil {
		return x.Goroutines
	}
	return 0
}

func (x *Health) GetGCCycles() uint32 {
	if x != nil {
		return x.GCCycles
	}
	return 0
}

func (x *Health) GetNumberOfCPUs() int32 {
	if x != nil {
		return x.NumberOfCPUs
	}
	return 0
}

func (x *Health) GetHeapSys() float32 {
	if x != nil {
		return x.HeapSys
	}
	return 0
}

func (x *Health) GetHeapAllocated() float32 {
	if x != nil {
		return x.HeapAllocated
	}
	return 0
}

func (x *Health) GetObjectsInUse() uint64 {
	if x != nil {
		return x.ObjectsInUse
	}
	return 0
}

func (x *Health) GetOSMemoryObtained() float32 {
	if x != nil {
		return x.OSMemoryObtained
	}
	return 0
}

var File_api_events_api_v1_events_api_proto protoreflect.FileDescriptor

var file_api_events_api_v1_events_api_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x42, 0x0c, 0x92, 0x41, 0x09, 0xd2, 0x01, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x2d, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xee, 0x02, 0x0a, 0x06, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x14, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x14, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x47, 0x6f,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x47, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x43,
	0x43, 0x79, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x47, 0x43,
	0x43, 0x79, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x4f, 0x66, 0x43, 0x50, 0x55, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x43, 0x50, 0x55, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x65,
	0x61, 0x70, 0x53, 0x79, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x48, 0x65, 0x61,
	0x70, 0x53, 0x79, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x70, 0x41, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x48, 0x65, 0x61,
	0x70, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x4f, 0x53, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4f, 0x62, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x4f, 0x53, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x4f, 0x62, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x32, 0x5e, 0x0a, 0x0d, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0b, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x09, 0x12, 0x07, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x32, 0x7f, 0x0a, 0x0c, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x0f, 0x50, 0x6f,
	0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x3e, 0x5a, 0x3c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x6f, 0x76, 0x63, 0x68, 0x69,
	0x6b, 0x75, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_events_api_v1_events_api_proto_rawDescOnce sync.Once
	file_api_events_api_v1_events_api_proto_rawDescData = file_api_events_api_v1_events_api_proto_rawDesc
)

func file_api_events_api_v1_events_api_proto_rawDescGZIP() []byte {
	file_api_events_api_v1_events_api_proto_rawDescOnce.Do(func() {
		file_api_events_api_v1_events_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_events_api_v1_events_api_proto_rawDescData)
	})
	return file_api_events_api_v1_events_api_proto_rawDescData
}

var file_api_events_api_v1_events_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_events_api_v1_events_api_proto_goTypes = []interface{}{
	(*CreateEventRequest)(nil),  // 0: events.api.v1.CreateEventRequest
	(*CreateEventResponse)(nil), // 1: events.api.v1.CreateEventResponse
	(*Health)(nil),              // 2: events.api.v1.Health
	(*emptypb.Empty)(nil),       // 3: google.protobuf.Empty
}
var file_api_events_api_v1_events_api_proto_depIdxs = []int32{
	3, // 0: events.api.v1.HealthService.CheckHealth:input_type -> google.protobuf.Empty
	0, // 1: events.api.v1.EventService.PostCreateEvent:input_type -> events.api.v1.CreateEventRequest
	2, // 2: events.api.v1.HealthService.CheckHealth:output_type -> events.api.v1.Health
	1, // 3: events.api.v1.EventService.PostCreateEvent:output_type -> events.api.v1.CreateEventResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_events_api_v1_events_api_proto_init() }
func file_api_events_api_v1_events_api_proto_init() {
	if File_api_events_api_v1_events_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_events_api_v1_events_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventRequest); i {
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
		file_api_events_api_v1_events_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventResponse); i {
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
		file_api_events_api_v1_events_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Health); i {
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
			RawDescriptor: file_api_events_api_v1_events_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_events_api_v1_events_api_proto_goTypes,
		DependencyIndexes: file_api_events_api_v1_events_api_proto_depIdxs,
		MessageInfos:      file_api_events_api_v1_events_api_proto_msgTypes,
	}.Build()
	File_api_events_api_v1_events_api_proto = out.File
	file_api_events_api_v1_events_api_proto_rawDesc = nil
	file_api_events_api_v1_events_api_proto_goTypes = nil
	file_api_events_api_v1_events_api_proto_depIdxs = nil
}