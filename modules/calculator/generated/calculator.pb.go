// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: calculator.proto

package generated

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

// Message definitions for UserMacroDistribution and other related entities
type UserMacroDistribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId                          int32  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Age                             uint32 `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Height                          uint32 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	Weight                          uint32 `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	Gender                          string `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
	System                          string `protobuf:"bytes,7,opt,name=system,proto3" json:"system,omitempty"`
	Activity                        string `protobuf:"bytes,8,opt,name=activity,proto3" json:"activity,omitempty"`
	ActivityDescription             string `protobuf:"bytes,9,opt,name=activity_description,json=activityDescription,proto3" json:"activity_description,omitempty"`
	Objective                       string `protobuf:"bytes,10,opt,name=objective,proto3" json:"objective,omitempty"`
	ObjectiveDescription            string `protobuf:"bytes,11,opt,name=objective_description,json=objectiveDescription,proto3" json:"objective_description,omitempty"`
	CaloriesDistribution            string `protobuf:"bytes,12,opt,name=calories_distribution,json=caloriesDistribution,proto3" json:"calories_distribution,omitempty"`
	CaloriesDistributionDescription string `protobuf:"bytes,13,opt,name=calories_distribution_description,json=caloriesDistributionDescription,proto3" json:"calories_distribution_description,omitempty"`
	Protein                         uint32 `protobuf:"varint,14,opt,name=protein,proto3" json:"protein,omitempty"`
	Fats                            uint32 `protobuf:"varint,15,opt,name=fats,proto3" json:"fats,omitempty"`
	Carbs                           uint32 `protobuf:"varint,16,opt,name=carbs,proto3" json:"carbs,omitempty"`
	Bmr                             uint32 `protobuf:"varint,17,opt,name=bmr,proto3" json:"bmr,omitempty"`
	Tdee                            uint32 `protobuf:"varint,18,opt,name=tdee,proto3" json:"tdee,omitempty"`
	Goal                            uint32 `protobuf:"varint,19,opt,name=goal,proto3" json:"goal,omitempty"`
	CreatedAt                       string `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *UserMacroDistribution) Reset() {
	*x = UserMacroDistribution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMacroDistribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMacroDistribution) ProtoMessage() {}

func (x *UserMacroDistribution) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMacroDistribution.ProtoReflect.Descriptor instead.
func (*UserMacroDistribution) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *UserMacroDistribution) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserMacroDistribution) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserMacroDistribution) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserMacroDistribution) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *UserMacroDistribution) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *UserMacroDistribution) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserMacroDistribution) GetSystem() string {
	if x != nil {
		return x.System
	}
	return ""
}

func (x *UserMacroDistribution) GetActivity() string {
	if x != nil {
		return x.Activity
	}
	return ""
}

func (x *UserMacroDistribution) GetActivityDescription() string {
	if x != nil {
		return x.ActivityDescription
	}
	return ""
}

func (x *UserMacroDistribution) GetObjective() string {
	if x != nil {
		return x.Objective
	}
	return ""
}

func (x *UserMacroDistribution) GetObjectiveDescription() string {
	if x != nil {
		return x.ObjectiveDescription
	}
	return ""
}

func (x *UserMacroDistribution) GetCaloriesDistribution() string {
	if x != nil {
		return x.CaloriesDistribution
	}
	return ""
}

func (x *UserMacroDistribution) GetCaloriesDistributionDescription() string {
	if x != nil {
		return x.CaloriesDistributionDescription
	}
	return ""
}

func (x *UserMacroDistribution) GetProtein() uint32 {
	if x != nil {
		return x.Protein
	}
	return 0
}

func (x *UserMacroDistribution) GetFats() uint32 {
	if x != nil {
		return x.Fats
	}
	return 0
}

func (x *UserMacroDistribution) GetCarbs() uint32 {
	if x != nil {
		return x.Carbs
	}
	return 0
}

func (x *UserMacroDistribution) GetBmr() uint32 {
	if x != nil {
		return x.Bmr
	}
	return 0
}

func (x *UserMacroDistribution) GetTdee() uint32 {
	if x != nil {
		return x.Tdee
	}
	return 0
}

func (x *UserMacroDistribution) GetGoal() uint32 {
	if x != nil {
		return x.Goal
	}
	return 0
}

func (x *UserMacroDistribution) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

// Request and response messages for the Create operation
type CreateUserMacroRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserMacro *UserMacroDistribution `protobuf:"bytes,1,opt,name=user_macro,json=userMacro,proto3" json:"user_macro,omitempty"`
}

func (x *CreateUserMacroRequest) Reset() {
	*x = CreateUserMacroRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserMacroRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserMacroRequest) ProtoMessage() {}

func (x *CreateUserMacroRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserMacroRequest.ProtoReflect.Descriptor instead.
func (*CreateUserMacroRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserMacroRequest) GetUserMacro() *UserMacroDistribution {
	if x != nil {
		return x.UserMacro
	}
	return nil
}

type CreateUserMacroResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserMacro *UserMacroDistribution `protobuf:"bytes,1,opt,name=user_macro,json=userMacro,proto3" json:"user_macro,omitempty"`
}

func (x *CreateUserMacroResponse) Reset() {
	*x = CreateUserMacroResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserMacroResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserMacroResponse) ProtoMessage() {}

func (x *CreateUserMacroResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserMacroResponse.ProtoReflect.Descriptor instead.
func (*CreateUserMacroResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserMacroResponse) GetUserMacro() *UserMacroDistribution {
	if x != nil {
		return x.UserMacro
	}
	return nil
}

// Request and response messages for the GetAll operation
type GetAllUserMacrosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetAllUserMacrosRequest) Reset() {
	*x = GetAllUserMacrosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserMacrosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserMacrosRequest) ProtoMessage() {}

func (x *GetAllUserMacrosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserMacrosRequest.ProtoReflect.Descriptor instead.
func (*GetAllUserMacrosRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllUserMacrosRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetAllUserMacrosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserMacros []*UserMacroDistribution `protobuf:"bytes,1,rep,name=user_macros,json=userMacros,proto3" json:"user_macros,omitempty"`
}

func (x *GetAllUserMacrosResponse) Reset() {
	*x = GetAllUserMacrosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserMacrosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserMacrosResponse) ProtoMessage() {}

func (x *GetAllUserMacrosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserMacrosResponse.ProtoReflect.Descriptor instead.
func (*GetAllUserMacrosResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllUserMacrosResponse) GetUserMacros() []*UserMacroDistribution {
	if x != nil {
		return x.UserMacros
	}
	return nil
}

// Request and response messages for the Get operation
type GetUserMacroRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlanId string `protobuf:"bytes,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
}

func (x *GetUserMacroRequest) Reset() {
	*x = GetUserMacroRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserMacroRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserMacroRequest) ProtoMessage() {}

func (x *GetUserMacroRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserMacroRequest.ProtoReflect.Descriptor instead.
func (*GetUserMacroRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserMacroRequest) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

type GetUserMacroResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserMacro *UserMacroDistribution `protobuf:"bytes,1,opt,name=user_macro,json=userMacro,proto3" json:"user_macro,omitempty"`
}

func (x *GetUserMacroResponse) Reset() {
	*x = GetUserMacroResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserMacroResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserMacroResponse) ProtoMessage() {}

func (x *GetUserMacroResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserMacroResponse.ProtoReflect.Descriptor instead.
func (*GetUserMacroResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserMacroResponse) GetUserMacro() *UserMacroDistribution {
	if x != nil {
		return x.UserMacro
	}
	return nil
}

var File_calculator_proto protoreflect.FileDescriptor

var file_calculator_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0xf2,
	0x04, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12,
	0x31, 0x0a, 0x14, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x33, 0x0a, 0x15, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x15, 0x63, 0x61, 0x6c, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x61, 0x6c, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x21, 0x63, 0x61,
	0x6c, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1f, 0x63, 0x61, 0x6c, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69,
	0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x61, 0x74, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x66, 0x61, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x62, 0x73, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x61, 0x72, 0x62, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x6d,
	0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x62, 0x6d, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x64, 0x65, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x64, 0x65, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x67, 0x6f, 0x61, 0x6c, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x67, 0x6f, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x5a, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x22,
	0x5b, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x63,
	0x72, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4d, 0x61, 0x63, 0x72, 0x6f, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x22, 0x32, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x5e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61,
	0x63, 0x72, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63, 0x72, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x73,
	0x22, 0x2e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x72, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64,
	0x22, 0x58, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x72, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6d, 0x61, 0x63, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61,
	0x63, 0x72, 0x6f, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x32, 0x9f, 0x02, 0x0a, 0x11, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5a, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61,
	0x63, 0x72, 0x6f, 0x12, 0x22, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x72, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x61, 0x63, 0x72, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x23,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x72, 0x6f,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x12, 0x1f, 0x2e, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61,
	0x63, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x61, 0x63, 0x72, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculator_proto_rawDescOnce sync.Once
	file_calculator_proto_rawDescData = file_calculator_proto_rawDesc
)

func file_calculator_proto_rawDescGZIP() []byte {
	file_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_proto_rawDescData)
	})
	return file_calculator_proto_rawDescData
}

var file_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_calculator_proto_goTypes = []any{
	(*UserMacroDistribution)(nil),    // 0: calculator.UserMacroDistribution
	(*CreateUserMacroRequest)(nil),   // 1: calculator.CreateUserMacroRequest
	(*CreateUserMacroResponse)(nil),  // 2: calculator.CreateUserMacroResponse
	(*GetAllUserMacrosRequest)(nil),  // 3: calculator.GetAllUserMacrosRequest
	(*GetAllUserMacrosResponse)(nil), // 4: calculator.GetAllUserMacrosResponse
	(*GetUserMacroRequest)(nil),      // 5: calculator.GetUserMacroRequest
	(*GetUserMacroResponse)(nil),     // 6: calculator.GetUserMacroResponse
}
var file_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.CreateUserMacroRequest.user_macro:type_name -> calculator.UserMacroDistribution
	0, // 1: calculator.CreateUserMacroResponse.user_macro:type_name -> calculator.UserMacroDistribution
	0, // 2: calculator.GetAllUserMacrosResponse.user_macros:type_name -> calculator.UserMacroDistribution
	0, // 3: calculator.GetUserMacroResponse.user_macro:type_name -> calculator.UserMacroDistribution
	1, // 4: calculator.CalculatorService.CreateUserMacro:input_type -> calculator.CreateUserMacroRequest
	3, // 5: calculator.CalculatorService.GetUsersMacros:input_type -> calculator.GetAllUserMacrosRequest
	5, // 6: calculator.CalculatorService.GetUserMacro:input_type -> calculator.GetUserMacroRequest
	2, // 7: calculator.CalculatorService.CreateUserMacro:output_type -> calculator.CreateUserMacroResponse
	4, // 8: calculator.CalculatorService.GetUsersMacros:output_type -> calculator.GetAllUserMacrosResponse
	6, // 9: calculator.CalculatorService.GetUserMacro:output_type -> calculator.GetUserMacroResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_calculator_proto_init() }
func file_calculator_proto_init() {
	if File_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UserMacroDistribution); i {
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
		file_calculator_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateUserMacroRequest); i {
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
		file_calculator_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateUserMacroResponse); i {
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
		file_calculator_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllUserMacrosRequest); i {
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
		file_calculator_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllUserMacrosResponse); i {
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
		file_calculator_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserMacroRequest); i {
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
		file_calculator_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserMacroResponse); i {
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
			RawDescriptor: file_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_proto_msgTypes,
	}.Build()
	File_calculator_proto = out.File
	file_calculator_proto_rawDesc = nil
	file_calculator_proto_goTypes = nil
	file_calculator_proto_depIdxs = nil
}
