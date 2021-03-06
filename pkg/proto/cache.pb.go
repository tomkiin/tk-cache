// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: cache.proto

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

type PingCacheReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingCacheReq) Reset() {
	*x = PingCacheReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingCacheReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingCacheReq) ProtoMessage() {}

func (x *PingCacheReq) ProtoReflect() protoreflect.Message {
	mi := &file_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingCacheReq.ProtoReflect.Descriptor instead.
func (*PingCacheReq) Descriptor() ([]byte, []int) {
	return file_cache_proto_rawDescGZIP(), []int{0}
}

type PingCacheRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *PingCacheRes) Reset() {
	*x = PingCacheRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingCacheRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingCacheRes) ProtoMessage() {}

func (x *PingCacheRes) ProtoReflect() protoreflect.Message {
	mi := &file_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingCacheRes.ProtoReflect.Descriptor instead.
func (*PingCacheRes) Descriptor() ([]byte, []int) {
	return file_cache_proto_rawDescGZIP(), []int{1}
}

func (x *PingCacheRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type GetCacheReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetCacheReq) Reset() {
	*x = GetCacheReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCacheReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCacheReq) ProtoMessage() {}

func (x *GetCacheReq) ProtoReflect() protoreflect.Message {
	mi := &file_cache_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCacheReq.ProtoReflect.Descriptor instead.
func (*GetCacheReq) Descriptor() ([]byte, []int) {
	return file_cache_proto_rawDescGZIP(), []int{2}
}

func (x *GetCacheReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetCacheRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Ok    bool   `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *GetCacheRes) Reset() {
	*x = GetCacheRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCacheRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCacheRes) ProtoMessage() {}

func (x *GetCacheRes) ProtoReflect() protoreflect.Message {
	mi := &file_cache_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCacheRes.ProtoReflect.Descriptor instead.
func (*GetCacheRes) Descriptor() ([]byte, []int) {
	return file_cache_proto_rawDescGZIP(), []int{3}
}

func (x *GetCacheRes) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *GetCacheRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type SetCacheReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetCacheReq) Reset() {
	*x = SetCacheReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetCacheReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCacheReq) ProtoMessage() {}

func (x *SetCacheReq) ProtoReflect() protoreflect.Message {
	mi := &file_cache_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetCacheReq.ProtoReflect.Descriptor instead.
func (*SetCacheReq) Descriptor() ([]byte, []int) {
	return file_cache_proto_rawDescGZIP(), []int{4}
}

func (x *SetCacheReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetCacheReq) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type SetCacheRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetCacheRes) Reset() {
	*x = SetCacheRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetCacheRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCacheRes) ProtoMessage() {}

func (x *SetCacheRes) ProtoReflect() protoreflect.Message {
	mi := &file_cache_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetCacheRes.ProtoReflect.Descriptor instead.
func (*SetCacheRes) Descriptor() ([]byte, []int) {
	return file_cache_proto_rawDescGZIP(), []int{5}
}

var File_cache_proto protoreflect.FileDescriptor

var file_cache_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x52, 0x65, 0x71, 0x22, 0x1e, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x02, 0x6f, 0x6b, 0x22, 0x1f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x33, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x35, 0x0a, 0x0b, 0x53, 0x65,
	0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x73,
	0x32, 0x9d, 0x01, 0x0a, 0x05, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2f,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cache_proto_rawDescOnce sync.Once
	file_cache_proto_rawDescData = file_cache_proto_rawDesc
)

func file_cache_proto_rawDescGZIP() []byte {
	file_cache_proto_rawDescOnce.Do(func() {
		file_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_cache_proto_rawDescData)
	})
	return file_cache_proto_rawDescData
}

var file_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cache_proto_goTypes = []interface{}{
	(*PingCacheReq)(nil), // 0: proto.PingCacheReq
	(*PingCacheRes)(nil), // 1: proto.PingCacheRes
	(*GetCacheReq)(nil),  // 2: proto.GetCacheReq
	(*GetCacheRes)(nil),  // 3: proto.GetCacheRes
	(*SetCacheReq)(nil),  // 4: proto.SetCacheReq
	(*SetCacheRes)(nil),  // 5: proto.SetCacheRes
}
var file_cache_proto_depIdxs = []int32{
	0, // 0: proto.Cache.Ping:input_type -> proto.PingCacheReq
	2, // 1: proto.Cache.Get:input_type -> proto.GetCacheReq
	4, // 2: proto.Cache.Set:input_type -> proto.SetCacheReq
	1, // 3: proto.Cache.Ping:output_type -> proto.PingCacheRes
	3, // 4: proto.Cache.Get:output_type -> proto.GetCacheRes
	5, // 5: proto.Cache.Set:output_type -> proto.SetCacheRes
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cache_proto_init() }
func file_cache_proto_init() {
	if File_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingCacheReq); i {
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
		file_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingCacheRes); i {
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
		file_cache_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCacheReq); i {
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
		file_cache_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCacheRes); i {
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
		file_cache_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetCacheReq); i {
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
		file_cache_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetCacheRes); i {
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
			RawDescriptor: file_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cache_proto_goTypes,
		DependencyIndexes: file_cache_proto_depIdxs,
		MessageInfos:      file_cache_proto_msgTypes,
	}.Build()
	File_cache_proto = out.File
	file_cache_proto_rawDesc = nil
	file_cache_proto_goTypes = nil
	file_cache_proto_depIdxs = nil
}
