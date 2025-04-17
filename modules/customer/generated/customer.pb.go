// Note: This is a proto3 file, proto2 is deprecated and should not be used.
// In proto3, all fields are optional by default, required and optional keywords
// are no longer required. Everything is optional by default.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: customer.proto

// Declaring a package name prevents collisions with other packages that use
// methods with the same name.

package generated

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type GetCustomerReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Be explicit with what you want to get, don't just"ID"
	PublicId string `protobuf:"bytes,1,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
	// Include the caller of the request to aid with quickly filtering through
	// structured logs.
	Downstream string `protobuf:"bytes,998,opt,name=downstream,proto3" json:"downstream,omitempty"`
	// It's good practice to include a request ID all requests to help with
	// tracing and debugging.
	RequestId     string `protobuf:"bytes,999,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCustomerReq) Reset() {
	*x = GetCustomerReq{}
	mi := &file_customer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerReq) ProtoMessage() {}

func (x *GetCustomerReq) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerReq.ProtoReflect.Descriptor instead.
func (*GetCustomerReq) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{0}
}

func (x *GetCustomerReq) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *GetCustomerReq) GetDownstream() string {
	if x != nil {
		return x.Downstream
	}
	return ""
}

func (x *GetCustomerReq) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type GetCustomerRes struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Success bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // Optional: populate if success is false
	// The customer object
	Customer *XCustomer `protobuf:"bytes,3,opt,name=customer,proto3" json:"customer,omitempty"`
	// Return the service that handled the request. This is useful when handling
	// switchovers from legacy services and you're still throttling
	Upstream string `protobuf:"bytes,998,opt,name=upstream,proto3" json:"upstream,omitempty"`
	// Return the request ID so we know the response is idempotent
	RequestId     string `protobuf:"bytes,999,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCustomerRes) Reset() {
	*x = GetCustomerRes{}
	mi := &file_customer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerRes) ProtoMessage() {}

func (x *GetCustomerRes) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerRes.ProtoReflect.Descriptor instead.
func (*GetCustomerRes) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{1}
}

func (x *GetCustomerRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetCustomerRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetCustomerRes) GetCustomer() *XCustomer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *GetCustomerRes) GetUpstream() string {
	if x != nil {
		return x.Upstream
	}
	return ""
}

func (x *GetCustomerRes) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type CreateCustomerReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The customer object
	Customer      *XCustomer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	Downstream    string     `protobuf:"bytes,998,opt,name=downstream,proto3" json:"downstream,omitempty"`
	RequestId     string     `protobuf:"bytes,999,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCustomerReq) Reset() {
	*x = CreateCustomerReq{}
	mi := &file_customer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCustomerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerReq) ProtoMessage() {}

func (x *CreateCustomerReq) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerReq.ProtoReflect.Descriptor instead.
func (*CreateCustomerReq) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCustomerReq) GetCustomer() *XCustomer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *CreateCustomerReq) GetDownstream() string {
	if x != nil {
		return x.Downstream
	}
	return ""
}

func (x *CreateCustomerReq) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type CreateCustomerRes struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Success bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// The customer object
	Customer      *XCustomer `protobuf:"bytes,3,opt,name=customer,proto3" json:"customer,omitempty"`
	Upstream      string     `protobuf:"bytes,998,opt,name=upstream,proto3" json:"upstream,omitempty"`
	RequestId     string     `protobuf:"bytes,999,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCustomerRes) Reset() {
	*x = CreateCustomerRes{}
	mi := &file_customer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCustomerRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerRes) ProtoMessage() {}

func (x *CreateCustomerRes) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerRes.ProtoReflect.Descriptor instead.
func (*CreateCustomerRes) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCustomerRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateCustomerRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateCustomerRes) GetCustomer() *XCustomer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *CreateCustomerRes) GetUpstream() string {
	if x != nil {
		return x.Upstream
	}
	return ""
}

func (x *CreateCustomerRes) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type UpdateCustomerReq struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	CustomerId string                 `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// This saves us from having to return the entire customer object and saves
	// on implementation complexity due to not all customer fields being
	// returned (so we can't just dump the entire object into the DB).
	Updates       []*XDiff `protobuf:"bytes,2,rep,name=updates,proto3" json:"updates,omitempty"`
	Downstream    string   `protobuf:"bytes,998,opt,name=downstream,proto3" json:"downstream,omitempty"`
	RequestId     string   `protobuf:"bytes,999,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCustomerReq) Reset() {
	*x = UpdateCustomerReq{}
	mi := &file_customer_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCustomerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCustomerReq) ProtoMessage() {}

func (x *UpdateCustomerReq) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCustomerReq.ProtoReflect.Descriptor instead.
func (*UpdateCustomerReq) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCustomerReq) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *UpdateCustomerReq) GetUpdates() []*XDiff {
	if x != nil {
		return x.Updates
	}
	return nil
}

func (x *UpdateCustomerReq) GetDownstream() string {
	if x != nil {
		return x.Downstream
	}
	return ""
}

func (x *UpdateCustomerReq) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type UpdateCustomerRes struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Success bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// The customer object
	Customer      *XCustomer `protobuf:"bytes,3,opt,name=customer,proto3" json:"customer,omitempty"`
	Upstream      string     `protobuf:"bytes,998,opt,name=upstream,proto3" json:"upstream,omitempty"`
	RequestId     string     `protobuf:"bytes,999,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCustomerRes) Reset() {
	*x = UpdateCustomerRes{}
	mi := &file_customer_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCustomerRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCustomerRes) ProtoMessage() {}

func (x *UpdateCustomerRes) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCustomerRes.ProtoReflect.Descriptor instead.
func (*UpdateCustomerRes) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCustomerRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdateCustomerRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpdateCustomerRes) GetCustomer() *XCustomer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *UpdateCustomerRes) GetUpstream() string {
	if x != nil {
		return x.Upstream
	}
	return ""
}

func (x *UpdateCustomerRes) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type DeleteCustomerReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    string                 `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	HardDelete    bool                   `protobuf:"varint,2,opt,name=hard_delete,json=hardDelete,proto3" json:"hard_delete,omitempty"` // soft delete by default
	Downstream    string                 `protobuf:"bytes,998,opt,name=downstream,proto3" json:"downstream,omitempty"`
	RequestId     string                 `protobuf:"bytes,999,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCustomerReq) Reset() {
	*x = DeleteCustomerReq{}
	mi := &file_customer_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCustomerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCustomerReq) ProtoMessage() {}

func (x *DeleteCustomerReq) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCustomerReq.ProtoReflect.Descriptor instead.
func (*DeleteCustomerReq) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCustomerReq) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *DeleteCustomerReq) GetHardDelete() bool {
	if x != nil {
		return x.HardDelete
	}
	return false
}

func (x *DeleteCustomerReq) GetDownstream() string {
	if x != nil {
		return x.Downstream
	}
	return ""
}

func (x *DeleteCustomerReq) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

// NilRes effectively mirrors google.protobuf.Empty, we use our own type though
// as the google.protobuf.Empty type is not available in all languages that we
// support.
type NilRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NilRes) Reset() {
	*x = NilRes{}
	mi := &file_customer_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NilRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NilRes) ProtoMessage() {}

func (x *NilRes) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[7]
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
	return file_customer_proto_rawDescGZIP(), []int{7}
}

// -- Subtype messages
type XCustomer struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PublicId      string                 `protobuf:"bytes,1,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
	PrivateId     string                 `protobuf:"bytes,2,opt,name=private_id,json=privateId,proto3" json:"private_id,omitempty"`
	Name          *XName                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"` // Optional - string due to intl dial code
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *XCustomer) Reset() {
	*x = XCustomer{}
	mi := &file_customer_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *XCustomer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XCustomer) ProtoMessage() {}

func (x *XCustomer) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XCustomer.ProtoReflect.Descriptor instead.
func (*XCustomer) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{8}
}

func (x *XCustomer) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *XCustomer) GetPrivateId() string {
	if x != nil {
		return x.PrivateId
	}
	return ""
}

func (x *XCustomer) GetName() *XName {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *XCustomer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *XCustomer) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type XName struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Title should be an enum but we'll keep it simple for now
	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	First    string `protobuf:"bytes,2,opt,name=first,proto3" json:"first,omitempty"`
	Middle   string `protobuf:"bytes,3,opt,name=middle,proto3" json:"middle,omitempty"`
	Last     string `protobuf:"bytes,4,opt,name=last,proto3" json:"last,omitempty"`
	Suffix   string `protobuf:"bytes,5,opt,name=suffix,proto3" json:"suffix,omitempty"`
	Nickname string `protobuf:"bytes,6,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// title + first + middle + last + suffix
	Full string `protobuf:"bytes,7,opt,name=full,proto3" json:"full,omitempty"`
	// first + last
	Friendly      string `protobuf:"bytes,8,opt,name=friendly,proto3" json:"friendly,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *XName) Reset() {
	*x = XName{}
	mi := &file_customer_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *XName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XName) ProtoMessage() {}

func (x *XName) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XName.ProtoReflect.Descriptor instead.
func (*XName) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{9}
}

func (x *XName) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *XName) GetFirst() string {
	if x != nil {
		return x.First
	}
	return ""
}

func (x *XName) GetMiddle() string {
	if x != nil {
		return x.Middle
	}
	return ""
}

func (x *XName) GetLast() string {
	if x != nil {
		return x.Last
	}
	return ""
}

func (x *XName) GetSuffix() string {
	if x != nil {
		return x.Suffix
	}
	return ""
}

func (x *XName) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *XName) GetFull() string {
	if x != nil {
		return x.Full
	}
	return ""
}

func (x *XName) GetFriendly() string {
	if x != nil {
		return x.Friendly
	}
	return ""
}

// While the value may not be a string, for diff patching, we wrap everything in
// a string.
type XDiff struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	OldValue      string                 `protobuf:"bytes,2,opt,name=old_value,json=oldValue,proto3" json:"old_value,omitempty"`
	NewValue      string                 `protobuf:"bytes,3,opt,name=new_value,json=newValue,proto3" json:"new_value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *XDiff) Reset() {
	*x = XDiff{}
	mi := &file_customer_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *XDiff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XDiff) ProtoMessage() {}

func (x *XDiff) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XDiff.ProtoReflect.Descriptor instead.
func (*XDiff) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{10}
}

func (x *XDiff) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *XDiff) GetOldValue() string {
	if x != nil {
		return x.OldValue
	}
	return ""
}

func (x *XDiff) GetNewValue() string {
	if x != nil {
		return x.NewValue
	}
	return ""
}

var File_customer_proto protoreflect.FileDescriptor

const file_customer_proto_rawDesc = "" +
	"\n" +
	"\x0ecustomer.proto\x12\x12fitSphere.customer\"n\n" +
	"\x0eGetCustomerReq\x12\x1b\n" +
	"\tpublic_id\x18\x01 \x01(\tR\bpublicId\x12\x1f\n" +
	"\n" +
	"downstream\x18\xe6\a \x01(\tR\n" +
	"downstream\x12\x1e\n" +
	"\n" +
	"request_id\x18\xe7\a \x01(\tR\trequestId\"\xbc\x01\n" +
	"\x0eGetCustomerRes\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x129\n" +
	"\bcustomer\x18\x03 \x01(\v2\x1d.fitSphere.customer.XCustomerR\bcustomer\x12\x1b\n" +
	"\bupstream\x18\xe6\a \x01(\tR\bupstream\x12\x1e\n" +
	"\n" +
	"request_id\x18\xe7\a \x01(\tR\trequestId\"\x8f\x01\n" +
	"\x11CreateCustomerReq\x129\n" +
	"\bcustomer\x18\x01 \x01(\v2\x1d.fitSphere.customer.XCustomerR\bcustomer\x12\x1f\n" +
	"\n" +
	"downstream\x18\xe6\a \x01(\tR\n" +
	"downstream\x12\x1e\n" +
	"\n" +
	"request_id\x18\xe7\a \x01(\tR\trequestId\"\xbf\x01\n" +
	"\x11CreateCustomerRes\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x129\n" +
	"\bcustomer\x18\x03 \x01(\v2\x1d.fitSphere.customer.XCustomerR\bcustomer\x12\x1b\n" +
	"\bupstream\x18\xe6\a \x01(\tR\bupstream\x12\x1e\n" +
	"\n" +
	"request_id\x18\xe7\a \x01(\tR\trequestId\"\xaa\x01\n" +
	"\x11UpdateCustomerReq\x12\x1f\n" +
	"\vcustomer_id\x18\x01 \x01(\tR\n" +
	"customerId\x123\n" +
	"\aupdates\x18\x02 \x03(\v2\x19.fitSphere.customer.XDiffR\aupdates\x12\x1f\n" +
	"\n" +
	"downstream\x18\xe6\a \x01(\tR\n" +
	"downstream\x12\x1e\n" +
	"\n" +
	"request_id\x18\xe7\a \x01(\tR\trequestId\"\xbf\x01\n" +
	"\x11UpdateCustomerRes\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x129\n" +
	"\bcustomer\x18\x03 \x01(\v2\x1d.fitSphere.customer.XCustomerR\bcustomer\x12\x1b\n" +
	"\bupstream\x18\xe6\a \x01(\tR\bupstream\x12\x1e\n" +
	"\n" +
	"request_id\x18\xe7\a \x01(\tR\trequestId\"\x96\x01\n" +
	"\x11DeleteCustomerReq\x12\x1f\n" +
	"\vcustomer_id\x18\x01 \x01(\tR\n" +
	"customerId\x12\x1f\n" +
	"\vhard_delete\x18\x02 \x01(\bR\n" +
	"hardDelete\x12\x1f\n" +
	"\n" +
	"downstream\x18\xe6\a \x01(\tR\n" +
	"downstream\x12\x1e\n" +
	"\n" +
	"request_id\x18\xe7\a \x01(\tR\trequestId\"\b\n" +
	"\x06NilRes\"\xa2\x01\n" +
	"\tXCustomer\x12\x1b\n" +
	"\tpublic_id\x18\x01 \x01(\tR\bpublicId\x12\x1d\n" +
	"\n" +
	"private_id\x18\x02 \x01(\tR\tprivateId\x12-\n" +
	"\x04name\x18\x03 \x01(\v2\x19.fitSphere.customer.XNameR\x04name\x12\x14\n" +
	"\x05email\x18\x04 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x05 \x01(\tR\x05phone\"\xc3\x01\n" +
	"\x05XName\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x14\n" +
	"\x05first\x18\x02 \x01(\tR\x05first\x12\x16\n" +
	"\x06middle\x18\x03 \x01(\tR\x06middle\x12\x12\n" +
	"\x04last\x18\x04 \x01(\tR\x04last\x12\x16\n" +
	"\x06suffix\x18\x05 \x01(\tR\x06suffix\x12\x1a\n" +
	"\bnickname\x18\x06 \x01(\tR\bnickname\x12\x12\n" +
	"\x04full\x18\a \x01(\tR\x04full\x12\x1a\n" +
	"\bfriendly\x18\b \x01(\tR\bfriendly\"W\n" +
	"\x05XDiff\x12\x14\n" +
	"\x05field\x18\x01 \x01(\tR\x05field\x12\x1b\n" +
	"\told_value\x18\x02 \x01(\tR\boldValue\x12\x1b\n" +
	"\tnew_value\x18\x03 \x01(\tR\bnewValue2\xf6\x02\n" +
	"\bCustomer\x12U\n" +
	"\vGetCustomer\x12\".fitSphere.customer.GetCustomerReq\x1a\".fitSphere.customer.GetCustomerRes\x12^\n" +
	"\x0eCreateCustomer\x12%.fitSphere.customer.CreateCustomerReq\x1a%.fitSphere.customer.CreateCustomerRes\x12^\n" +
	"\x0eUpdateCustomer\x12%.fitSphere.customer.UpdateCustomerReq\x1a%.fitSphere.customer.UpdateCustomerRes\x12S\n" +
	"\x0eDeleteCustomer\x12%.fitSphere.customer.DeleteCustomerReq\x1a\x1a.fitSphere.customer.NilResb\x06proto3"

var (
	file_customer_proto_rawDescOnce sync.Once
	file_customer_proto_rawDescData []byte
)

func file_customer_proto_rawDescGZIP() []byte {
	file_customer_proto_rawDescOnce.Do(func() {
		file_customer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_customer_proto_rawDesc), len(file_customer_proto_rawDesc)))
	})
	return file_customer_proto_rawDescData
}

var file_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_customer_proto_goTypes = []any{
	(*GetCustomerReq)(nil),    // 0: fitSphere.customer.GetCustomerReq
	(*GetCustomerRes)(nil),    // 1: fitSphere.customer.GetCustomerRes
	(*CreateCustomerReq)(nil), // 2: fitSphere.customer.CreateCustomerReq
	(*CreateCustomerRes)(nil), // 3: fitSphere.customer.CreateCustomerRes
	(*UpdateCustomerReq)(nil), // 4: fitSphere.customer.UpdateCustomerReq
	(*UpdateCustomerRes)(nil), // 5: fitSphere.customer.UpdateCustomerRes
	(*DeleteCustomerReq)(nil), // 6: fitSphere.customer.DeleteCustomerReq
	(*NilRes)(nil),            // 7: fitSphere.customer.NilRes
	(*XCustomer)(nil),         // 8: fitSphere.customer.XCustomer
	(*XName)(nil),             // 9: fitSphere.customer.XName
	(*XDiff)(nil),             // 10: fitSphere.customer.XDiff
}
var file_customer_proto_depIdxs = []int32{
	8,  // 0: fitSphere.customer.GetCustomerRes.customer:type_name -> fitSphere.customer.XCustomer
	8,  // 1: fitSphere.customer.CreateCustomerReq.customer:type_name -> fitSphere.customer.XCustomer
	8,  // 2: fitSphere.customer.CreateCustomerRes.customer:type_name -> fitSphere.customer.XCustomer
	10, // 3: fitSphere.customer.UpdateCustomerReq.updates:type_name -> fitSphere.customer.XDiff
	8,  // 4: fitSphere.customer.UpdateCustomerRes.customer:type_name -> fitSphere.customer.XCustomer
	9,  // 5: fitSphere.customer.XCustomer.name:type_name -> fitSphere.customer.XName
	0,  // 6: fitSphere.customer.Customer.GetCustomer:input_type -> fitSphere.customer.GetCustomerReq
	2,  // 7: fitSphere.customer.Customer.CreateCustomer:input_type -> fitSphere.customer.CreateCustomerReq
	4,  // 8: fitSphere.customer.Customer.UpdateCustomer:input_type -> fitSphere.customer.UpdateCustomerReq
	6,  // 9: fitSphere.customer.Customer.DeleteCustomer:input_type -> fitSphere.customer.DeleteCustomerReq
	1,  // 10: fitSphere.customer.Customer.GetCustomer:output_type -> fitSphere.customer.GetCustomerRes
	3,  // 11: fitSphere.customer.Customer.CreateCustomer:output_type -> fitSphere.customer.CreateCustomerRes
	5,  // 12: fitSphere.customer.Customer.UpdateCustomer:output_type -> fitSphere.customer.UpdateCustomerRes
	7,  // 13: fitSphere.customer.Customer.DeleteCustomer:output_type -> fitSphere.customer.NilRes
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_customer_proto_init() }
func file_customer_proto_init() {
	if File_customer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_customer_proto_rawDesc), len(file_customer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customer_proto_goTypes,
		DependencyIndexes: file_customer_proto_depIdxs,
		MessageInfos:      file_customer_proto_msgTypes,
	}.Build()
	File_customer_proto = out.File
	file_customer_proto_goTypes = nil
	file_customer_proto_depIdxs = nil
}
