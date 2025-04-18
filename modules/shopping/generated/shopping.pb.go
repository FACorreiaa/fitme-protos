// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: shopping.proto

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

type GenerateShoppingListReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateShoppingListReq) Reset() {
	*x = GenerateShoppingListReq{}
	mi := &file_shopping_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateShoppingListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateShoppingListReq) ProtoMessage() {}

func (x *GenerateShoppingListReq) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateShoppingListReq.ProtoReflect.Descriptor instead.
func (*GenerateShoppingListReq) Descriptor() ([]byte, []int) {
	return file_shopping_proto_rawDescGZIP(), []int{0}
}

type GenerateShoppingListRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateShoppingListRes) Reset() {
	*x = GenerateShoppingListRes{}
	mi := &file_shopping_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateShoppingListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateShoppingListRes) ProtoMessage() {}

func (x *GenerateShoppingListRes) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateShoppingListRes.ProtoReflect.Descriptor instead.
func (*GenerateShoppingListRes) Descriptor() ([]byte, []int) {
	return file_shopping_proto_rawDescGZIP(), []int{1}
}

type GetShoppingListReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetShoppingListReq) Reset() {
	*x = GetShoppingListReq{}
	mi := &file_shopping_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetShoppingListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShoppingListReq) ProtoMessage() {}

func (x *GetShoppingListReq) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShoppingListReq.ProtoReflect.Descriptor instead.
func (*GetShoppingListReq) Descriptor() ([]byte, []int) {
	return file_shopping_proto_rawDescGZIP(), []int{2}
}

type GetShoppingListRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetShoppingListRes) Reset() {
	*x = GetShoppingListRes{}
	mi := &file_shopping_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetShoppingListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShoppingListRes) ProtoMessage() {}

func (x *GetShoppingListRes) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShoppingListRes.ProtoReflect.Descriptor instead.
func (*GetShoppingListRes) Descriptor() ([]byte, []int) {
	return file_shopping_proto_rawDescGZIP(), []int{3}
}

type UpdateShoppingListReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateShoppingListReq) Reset() {
	*x = UpdateShoppingListReq{}
	mi := &file_shopping_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateShoppingListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateShoppingListReq) ProtoMessage() {}

func (x *UpdateShoppingListReq) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateShoppingListReq.ProtoReflect.Descriptor instead.
func (*UpdateShoppingListReq) Descriptor() ([]byte, []int) {
	return file_shopping_proto_rawDescGZIP(), []int{4}
}

type UpdateShoppingListRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateShoppingListRes) Reset() {
	*x = UpdateShoppingListRes{}
	mi := &file_shopping_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateShoppingListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateShoppingListRes) ProtoMessage() {}

func (x *UpdateShoppingListRes) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateShoppingListRes.ProtoReflect.Descriptor instead.
func (*UpdateShoppingListRes) Descriptor() ([]byte, []int) {
	return file_shopping_proto_rawDescGZIP(), []int{5}
}

type DeleteShoppingListReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteShoppingListReq) Reset() {
	*x = DeleteShoppingListReq{}
	mi := &file_shopping_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteShoppingListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShoppingListReq) ProtoMessage() {}

func (x *DeleteShoppingListReq) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShoppingListReq.ProtoReflect.Descriptor instead.
func (*DeleteShoppingListReq) Descriptor() ([]byte, []int) {
	return file_shopping_proto_rawDescGZIP(), []int{6}
}

type NilRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NilRes) Reset() {
	*x = NilRes{}
	mi := &file_shopping_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NilRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NilRes) ProtoMessage() {}

func (x *NilRes) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NilRes.ProtoReflect.Descriptor instead.
func (*NilRes) Descriptor() ([]byte, []int) {
	return file_shopping_proto_rawDescGZIP(), []int{7}
}

var File_shopping_proto protoreflect.FileDescriptor

const file_shopping_proto_rawDesc = "" +
	"\n" +
	"\x0eshopping.proto\x12\x12fitSphere.shopping\x1a\x1fgoogle/protobuf/timestamp.proto\"\x19\n" +
	"\x17GenerateShoppingListReq\"\x19\n" +
	"\x17GenerateShoppingListRes\"\x14\n" +
	"\x12GetShoppingListReq\"\x14\n" +
	"\x12GetShoppingListRes\"\x17\n" +
	"\x15UpdateShoppingListReq\"\x17\n" +
	"\x15UpdateShoppingListRes\"\x17\n" +
	"\x15DeleteShoppingListReq\"\b\n" +
	"\x06NilRes2\xa8\x03\n" +
	"\bShopping\x12p\n" +
	"\x14GenerateShoppingList\x12+.fitSphere.shopping.GenerateShoppingListReq\x1a+.fitSphere.shopping.GenerateShoppingListRes\x12a\n" +
	"\x0fGetShoppingList\x12&.fitSphere.shopping.GetShoppingListReq\x1a&.fitSphere.shopping.GetShoppingListRes\x12j\n" +
	"\x12UpdateShoppingList\x12).fitSphere.shopping.UpdateShoppingListReq\x1a).fitSphere.shopping.UpdateShoppingListRes\x12[\n" +
	"\x12DeleteShoppingList\x12).fitSphere.shopping.DeleteShoppingListReq\x1a\x1a.fitSphere.shopping.NilResb\x06proto3"

var (
	file_shopping_proto_rawDescOnce sync.Once
	file_shopping_proto_rawDescData []byte
)

func file_shopping_proto_rawDescGZIP() []byte {
	file_shopping_proto_rawDescOnce.Do(func() {
		file_shopping_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_shopping_proto_rawDesc), len(file_shopping_proto_rawDesc)))
	})
	return file_shopping_proto_rawDescData
}

var file_shopping_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_shopping_proto_goTypes = []any{
	(*GenerateShoppingListReq)(nil), // 0: fitSphere.shopping.GenerateShoppingListReq
	(*GenerateShoppingListRes)(nil), // 1: fitSphere.shopping.GenerateShoppingListRes
	(*GetShoppingListReq)(nil),      // 2: fitSphere.shopping.GetShoppingListReq
	(*GetShoppingListRes)(nil),      // 3: fitSphere.shopping.GetShoppingListRes
	(*UpdateShoppingListReq)(nil),   // 4: fitSphere.shopping.UpdateShoppingListReq
	(*UpdateShoppingListRes)(nil),   // 5: fitSphere.shopping.UpdateShoppingListRes
	(*DeleteShoppingListReq)(nil),   // 6: fitSphere.shopping.DeleteShoppingListReq
	(*NilRes)(nil),                  // 7: fitSphere.shopping.NilRes
}
var file_shopping_proto_depIdxs = []int32{
	0, // 0: fitSphere.shopping.Shopping.GenerateShoppingList:input_type -> fitSphere.shopping.GenerateShoppingListReq
	2, // 1: fitSphere.shopping.Shopping.GetShoppingList:input_type -> fitSphere.shopping.GetShoppingListReq
	4, // 2: fitSphere.shopping.Shopping.UpdateShoppingList:input_type -> fitSphere.shopping.UpdateShoppingListReq
	6, // 3: fitSphere.shopping.Shopping.DeleteShoppingList:input_type -> fitSphere.shopping.DeleteShoppingListReq
	1, // 4: fitSphere.shopping.Shopping.GenerateShoppingList:output_type -> fitSphere.shopping.GenerateShoppingListRes
	3, // 5: fitSphere.shopping.Shopping.GetShoppingList:output_type -> fitSphere.shopping.GetShoppingListRes
	5, // 6: fitSphere.shopping.Shopping.UpdateShoppingList:output_type -> fitSphere.shopping.UpdateShoppingListRes
	7, // 7: fitSphere.shopping.Shopping.DeleteShoppingList:output_type -> fitSphere.shopping.NilRes
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shopping_proto_init() }
func file_shopping_proto_init() {
	if File_shopping_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_shopping_proto_rawDesc), len(file_shopping_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shopping_proto_goTypes,
		DependencyIndexes: file_shopping_proto_depIdxs,
		MessageInfos:      file_shopping_proto_msgTypes,
	}.Build()
	File_shopping_proto = out.File
	file_shopping_proto_goTypes = nil
	file_shopping_proto_depIdxs = nil
}
