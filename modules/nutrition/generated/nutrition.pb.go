// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: nutrition.proto

package generated

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type AnalyzeMealPlanReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnalyzeMealPlanReq) Reset() {
	*x = AnalyzeMealPlanReq{}
	mi := &file_nutrition_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnalyzeMealPlanReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzeMealPlanReq) ProtoMessage() {}

func (x *AnalyzeMealPlanReq) ProtoReflect() protoreflect.Message {
	mi := &file_nutrition_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzeMealPlanReq.ProtoReflect.Descriptor instead.
func (*AnalyzeMealPlanReq) Descriptor() ([]byte, []int) {
	return file_nutrition_proto_rawDescGZIP(), []int{0}
}

type AnalyzeMealPlanRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnalyzeMealPlanRes) Reset() {
	*x = AnalyzeMealPlanRes{}
	mi := &file_nutrition_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnalyzeMealPlanRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzeMealPlanRes) ProtoMessage() {}

func (x *AnalyzeMealPlanRes) ProtoReflect() protoreflect.Message {
	mi := &file_nutrition_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzeMealPlanRes.ProtoReflect.Descriptor instead.
func (*AnalyzeMealPlanRes) Descriptor() ([]byte, []int) {
	return file_nutrition_proto_rawDescGZIP(), []int{1}
}

type RecommendChangesReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RecommendChangesReq) Reset() {
	*x = RecommendChangesReq{}
	mi := &file_nutrition_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecommendChangesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendChangesReq) ProtoMessage() {}

func (x *RecommendChangesReq) ProtoReflect() protoreflect.Message {
	mi := &file_nutrition_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendChangesReq.ProtoReflect.Descriptor instead.
func (*RecommendChangesReq) Descriptor() ([]byte, []int) {
	return file_nutrition_proto_rawDescGZIP(), []int{2}
}

type RecommendChangesRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RecommendChangesRes) Reset() {
	*x = RecommendChangesRes{}
	mi := &file_nutrition_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecommendChangesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendChangesRes) ProtoMessage() {}

func (x *RecommendChangesRes) ProtoReflect() protoreflect.Message {
	mi := &file_nutrition_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendChangesRes.ProtoReflect.Descriptor instead.
func (*RecommendChangesRes) Descriptor() ([]byte, []int) {
	return file_nutrition_proto_rawDescGZIP(), []int{3}
}

var File_nutrition_proto protoreflect.FileDescriptor

var file_nutrition_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x6e, 0x75, 0x74, 0x72, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x66, 0x69, 0x74, 0x53, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x6e, 0x75, 0x74,
	0x72, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x7a, 0x65, 0x4d, 0x65, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x22, 0x14, 0x0a,
	0x12, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x4d, 0x65, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x32, 0xe2, 0x01, 0x0a, 0x13, 0x4e, 0x75, 0x74, 0x72, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x63, 0x0a, 0x0f, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x7a, 0x65, 0x4d, 0x65, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x27, 0x2e, 0x66,
	0x69, 0x74, 0x53, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x6e, 0x75, 0x74, 0x72, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x4d, 0x65, 0x61, 0x6c, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x66, 0x69, 0x74, 0x53, 0x70, 0x68, 0x65, 0x72,
	0x65, 0x2e, 0x6e, 0x75, 0x74, 0x72, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x7a, 0x65, 0x4d, 0x65, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x66,
	0x0a, 0x10, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x12, 0x28, 0x2e, 0x66, 0x69, 0x74, 0x53, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x6e,
	0x75, 0x74, 0x72, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x66,
	0x69, 0x74, 0x53, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x6e, 0x75, 0x74, 0x72, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nutrition_proto_rawDescOnce sync.Once
	file_nutrition_proto_rawDescData []byte
)

func file_nutrition_proto_rawDescGZIP() []byte {
	file_nutrition_proto_rawDescOnce.Do(func() {
		file_nutrition_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nutrition_proto_rawDesc), len(file_nutrition_proto_rawDesc)))
	})
	return file_nutrition_proto_rawDescData
}

var file_nutrition_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nutrition_proto_goTypes = []any{
	(*AnalyzeMealPlanReq)(nil),  // 0: fitSphere.nutrition.AnalyzeMealPlanReq
	(*AnalyzeMealPlanRes)(nil),  // 1: fitSphere.nutrition.AnalyzeMealPlanRes
	(*RecommendChangesReq)(nil), // 2: fitSphere.nutrition.RecommendChangesReq
	(*RecommendChangesRes)(nil), // 3: fitSphere.nutrition.RecommendChangesRes
}
var file_nutrition_proto_depIdxs = []int32{
	0, // 0: fitSphere.nutrition.NutritionalAnalysis.AnalyzeMealPlan:input_type -> fitSphere.nutrition.AnalyzeMealPlanReq
	2, // 1: fitSphere.nutrition.NutritionalAnalysis.RecommendChanges:input_type -> fitSphere.nutrition.RecommendChangesReq
	1, // 2: fitSphere.nutrition.NutritionalAnalysis.AnalyzeMealPlan:output_type -> fitSphere.nutrition.AnalyzeMealPlanRes
	3, // 3: fitSphere.nutrition.NutritionalAnalysis.RecommendChanges:output_type -> fitSphere.nutrition.RecommendChangesRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nutrition_proto_init() }
func file_nutrition_proto_init() {
	if File_nutrition_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nutrition_proto_rawDesc), len(file_nutrition_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nutrition_proto_goTypes,
		DependencyIndexes: file_nutrition_proto_depIdxs,
		MessageInfos:      file_nutrition_proto_msgTypes,
	}.Build()
	File_nutrition_proto = out.File
	file_nutrition_proto_goTypes = nil
	file_nutrition_proto_depIdxs = nil
}
