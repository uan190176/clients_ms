// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: normalize_address.proto

package api

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

type NormalizeAddressMS_Bool int32

const (
	NormalizeAddressMS_Bool_IS_ANY   NormalizeAddressMS_Bool = 0
	NormalizeAddressMS_Bool_IS_FALSE NormalizeAddressMS_Bool = 1
	NormalizeAddressMS_Bool_IS_TRUE  NormalizeAddressMS_Bool = 2
)

// Enum value maps for NormalizeAddressMS_Bool.
var (
	NormalizeAddressMS_Bool_name = map[int32]string{
		0: "IS_ANY",
		1: "IS_FALSE",
		2: "IS_TRUE",
	}
	NormalizeAddressMS_Bool_value = map[string]int32{
		"IS_ANY":   0,
		"IS_FALSE": 1,
		"IS_TRUE":  2,
	}
)

func (x NormalizeAddressMS_Bool) Enum() *NormalizeAddressMS_Bool {
	p := new(NormalizeAddressMS_Bool)
	*p = x
	return p
}

func (x NormalizeAddressMS_Bool) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NormalizeAddressMS_Bool) Descriptor() protoreflect.EnumDescriptor {
	return file_normalize_address_proto_enumTypes[0].Descriptor()
}

func (NormalizeAddressMS_Bool) Type() protoreflect.EnumType {
	return &file_normalize_address_proto_enumTypes[0]
}

func (x NormalizeAddressMS_Bool) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NormalizeAddressMS_Bool.Descriptor instead.
func (NormalizeAddressMS_Bool) EnumDescriptor() ([]byte, []int) {
	return file_normalize_address_proto_rawDescGZIP(), []int{0}
}

// ***********************************************************************
// Requests
// ***********************************************************************
type RequestNormalizeAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken     string `protobuf:"bytes,1,opt,name=AuthToken,proto3" json:"AuthToken,omitempty"`
	AuthorId      uint64 `protobuf:"fixed64,2,opt,name=AuthorId,proto3" json:"AuthorId,omitempty"`
	Address       string `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	AddressTypeId uint64 `protobuf:"fixed64,4,opt,name=AddressTypeId,proto3" json:"AddressTypeId,omitempty"`
}

func (x *RequestNormalizeAddress) Reset() {
	*x = RequestNormalizeAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_normalize_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestNormalizeAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestNormalizeAddress) ProtoMessage() {}

func (x *RequestNormalizeAddress) ProtoReflect() protoreflect.Message {
	mi := &file_normalize_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestNormalizeAddress.ProtoReflect.Descriptor instead.
func (*RequestNormalizeAddress) Descriptor() ([]byte, []int) {
	return file_normalize_address_proto_rawDescGZIP(), []int{0}
}

func (x *RequestNormalizeAddress) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *RequestNormalizeAddress) GetAuthorId() uint64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *RequestNormalizeAddress) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RequestNormalizeAddress) GetAddressTypeId() uint64 {
	if x != nil {
		return x.AddressTypeId
	}
	return 0
}

// =============== //
//
//	Status        //
//
// =============== //
type NormalizeAddressMS_Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *NormalizeAddressMS_Status) Reset() {
	*x = NormalizeAddressMS_Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_normalize_address_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NormalizeAddressMS_Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormalizeAddressMS_Status) ProtoMessage() {}

func (x *NormalizeAddressMS_Status) ProtoReflect() protoreflect.Message {
	mi := &file_normalize_address_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormalizeAddressMS_Status.ProtoReflect.Descriptor instead.
func (*NormalizeAddressMS_Status) Descriptor() ([]byte, []int) {
	return file_normalize_address_proto_rawDescGZIP(), []int{1}
}

func (x *NormalizeAddressMS_Status) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *NormalizeAddressMS_Status) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// ================== //
// Normalized address //
// ================== //
type NormalizeCodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qc          string `protobuf:"bytes,1,opt,name=Qc,proto3" json:"Qc,omitempty"`
	QcComplete  string `protobuf:"bytes,2,opt,name=QcComplete,proto3" json:"QcComplete,omitempty"`
	QcGeo       string `protobuf:"bytes,3,opt,name=QcGeo,proto3" json:"QcGeo,omitempty"`
	QcHouse     string `protobuf:"bytes,4,opt,name=QcHouse,proto3" json:"QcHouse,omitempty"`
	PostQuality string `protobuf:"bytes,5,opt,name=PostQuality,proto3" json:"PostQuality,omitempty"`
	PostValid   string `protobuf:"bytes,6,opt,name=PostValid,proto3" json:"PostValid,omitempty"`
}

func (x *NormalizeCodes) Reset() {
	*x = NormalizeCodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_normalize_address_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NormalizeCodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormalizeCodes) ProtoMessage() {}

func (x *NormalizeCodes) ProtoReflect() protoreflect.Message {
	mi := &file_normalize_address_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormalizeCodes.ProtoReflect.Descriptor instead.
func (*NormalizeCodes) Descriptor() ([]byte, []int) {
	return file_normalize_address_proto_rawDescGZIP(), []int{2}
}

func (x *NormalizeCodes) GetQc() string {
	if x != nil {
		return x.Qc
	}
	return ""
}

func (x *NormalizeCodes) GetQcComplete() string {
	if x != nil {
		return x.QcComplete
	}
	return ""
}

func (x *NormalizeCodes) GetQcGeo() string {
	if x != nil {
		return x.QcGeo
	}
	return ""
}

func (x *NormalizeCodes) GetQcHouse() string {
	if x != nil {
		return x.QcHouse
	}
	return ""
}

func (x *NormalizeCodes) GetPostQuality() string {
	if x != nil {
		return x.PostQuality
	}
	return ""
}

func (x *NormalizeCodes) GetPostValid() string {
	if x != nil {
		return x.PostValid
	}
	return ""
}

type NormalizedAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressOriginal   string          `protobuf:"bytes,1,opt,name=AddressOriginal,proto3" json:"AddressOriginal,omitempty"`
	AddressNormalized string          `protobuf:"bytes,2,opt,name=AddressNormalized,proto3" json:"AddressNormalized,omitempty"`
	AddressJson       string          `protobuf:"bytes,3,opt,name=AddressJson,proto3" json:"AddressJson,omitempty"`
	PostalCode        string          `protobuf:"bytes,4,opt,name=PostalCode,proto3" json:"PostalCode,omitempty"`
	Timezone          string          `protobuf:"bytes,5,opt,name=Timezone,proto3" json:"Timezone,omitempty"`
	IsQualityByDadata bool            `protobuf:"varint,6,opt,name=IsQualityByDadata,proto3" json:"IsQualityByDadata,omitempty"`
	IsValidByPost     bool            `protobuf:"varint,7,opt,name=IsValidByPost,proto3" json:"IsValidByPost,omitempty"`
	NormalizeCodes    *NormalizeCodes `protobuf:"bytes,8,opt,name=NormalizeCodes,proto3" json:"NormalizeCodes,omitempty"`
}

func (x *NormalizedAddress) Reset() {
	*x = NormalizedAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_normalize_address_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NormalizedAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormalizedAddress) ProtoMessage() {}

func (x *NormalizedAddress) ProtoReflect() protoreflect.Message {
	mi := &file_normalize_address_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormalizedAddress.ProtoReflect.Descriptor instead.
func (*NormalizedAddress) Descriptor() ([]byte, []int) {
	return file_normalize_address_proto_rawDescGZIP(), []int{3}
}

func (x *NormalizedAddress) GetAddressOriginal() string {
	if x != nil {
		return x.AddressOriginal
	}
	return ""
}

func (x *NormalizedAddress) GetAddressNormalized() string {
	if x != nil {
		return x.AddressNormalized
	}
	return ""
}

func (x *NormalizedAddress) GetAddressJson() string {
	if x != nil {
		return x.AddressJson
	}
	return ""
}

func (x *NormalizedAddress) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *NormalizedAddress) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *NormalizedAddress) GetIsQualityByDadata() bool {
	if x != nil {
		return x.IsQualityByDadata
	}
	return false
}

func (x *NormalizedAddress) GetIsValidByPost() bool {
	if x != nil {
		return x.IsValidByPost
	}
	return false
}

func (x *NormalizedAddress) GetNormalizeCodes() *NormalizeCodes {
	if x != nil {
		return x.NormalizeCodes
	}
	return nil
}

type ResponseNormalizeAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NormalizedAddress *NormalizedAddress         `protobuf:"bytes,1,opt,name=NormalizedAddress,proto3" json:"NormalizedAddress,omitempty"`
	Status            *NormalizeAddressMS_Status `protobuf:"bytes,3,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *ResponseNormalizeAddress) Reset() {
	*x = ResponseNormalizeAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_normalize_address_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseNormalizeAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseNormalizeAddress) ProtoMessage() {}

func (x *ResponseNormalizeAddress) ProtoReflect() protoreflect.Message {
	mi := &file_normalize_address_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseNormalizeAddress.ProtoReflect.Descriptor instead.
func (*ResponseNormalizeAddress) Descriptor() ([]byte, []int) {
	return file_normalize_address_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseNormalizeAddress) GetNormalizedAddress() *NormalizedAddress {
	if x != nil {
		return x.NormalizedAddress
	}
	return nil
}

func (x *ResponseNormalizeAddress) GetStatus() *NormalizeAddressMS_Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_normalize_address_proto protoreflect.FileDescriptor

var file_normalize_address_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x6f, 0x72, 0x6d, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x93, 0x01, 0x0a,
	0x17, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06, 0x52, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x06, 0x52, 0x0d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x64, 0x22, 0x51, 0x0a, 0x19, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4d, 0x53, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb0, 0x01, 0x0a, 0x0e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x51, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x51, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x51, 0x63, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x51, 0x63,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x51, 0x63, 0x47, 0x65,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x51, 0x63, 0x47, 0x65, 0x6f, 0x12, 0x18,
	0x0a, 0x07, 0x51, 0x63, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x51, 0x63, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74,
	0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50,
	0x6f, 0x73, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6f,
	0x73, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50,
	0x6f, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x22, 0xe8, 0x02, 0x0a, 0x11, 0x4e, 0x6f, 0x72,
	0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x28,
	0x0a, 0x0f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4e, 0x6f, 0x72, 0x6d,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x4a, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6f,
	0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x69, 0x6d, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x49, 0x73, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x42, 0x79, 0x44, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x11, 0x49, 0x73, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x79, 0x44, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x24, 0x0a, 0x0d, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x42, 0x79, 0x50,
	0x6f, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x49, 0x73, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x42, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x0e, 0x4e, 0x6f, 0x72, 0x6d,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x2e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x52, 0x0e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x22, 0xb4, 0x01, 0x0a, 0x18, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x52, 0x0a, 0x11, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x6f,
	0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x11, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x44, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4d, 0x53, 0x5f, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x40, 0x0a, 0x17, 0x4e, 0x6f,
	0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4d, 0x53,
	0x5f, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x53, 0x5f, 0x41, 0x4e, 0x59, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x53, 0x5f, 0x46, 0x41, 0x4c, 0x53, 0x45, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x49, 0x53, 0x5f, 0x54, 0x52, 0x55, 0x45, 0x10, 0x02, 0x32, 0x8f, 0x01, 0x0a,
	0x1e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x6d, 0x0a, 0x10, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x2a, 0x2e, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a,
	0x2b, 0x2e, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x72, 0x6d,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x00, 0x42, 0x0e,
	0x5a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_normalize_address_proto_rawDescOnce sync.Once
	file_normalize_address_proto_rawDescData = file_normalize_address_proto_rawDesc
)

func file_normalize_address_proto_rawDescGZIP() []byte {
	file_normalize_address_proto_rawDescOnce.Do(func() {
		file_normalize_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_normalize_address_proto_rawDescData)
	})
	return file_normalize_address_proto_rawDescData
}

var file_normalize_address_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_normalize_address_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_normalize_address_proto_goTypes = []interface{}{
	(NormalizeAddressMS_Bool)(0),      // 0: normalize_address.NormalizeAddressMS_Bool
	(*RequestNormalizeAddress)(nil),   // 1: normalize_address.RequestNormalizeAddress
	(*NormalizeAddressMS_Status)(nil), // 2: normalize_address.NormalizeAddressMS_Status
	(*NormalizeCodes)(nil),            // 3: normalize_address.NormalizeCodes
	(*NormalizedAddress)(nil),         // 4: normalize_address.NormalizedAddress
	(*ResponseNormalizeAddress)(nil),  // 5: normalize_address.ResponseNormalizeAddress
}
var file_normalize_address_proto_depIdxs = []int32{
	3, // 0: normalize_address.NormalizedAddress.NormalizeCodes:type_name -> normalize_address.NormalizeCodes
	4, // 1: normalize_address.ResponseNormalizeAddress.NormalizedAddress:type_name -> normalize_address.NormalizedAddress
	2, // 2: normalize_address.ResponseNormalizeAddress.Status:type_name -> normalize_address.NormalizeAddressMS_Status
	1, // 3: normalize_address.NormalizeAddressServicesServer.NormalizeAddress:input_type -> normalize_address.RequestNormalizeAddress
	5, // 4: normalize_address.NormalizeAddressServicesServer.NormalizeAddress:output_type -> normalize_address.ResponseNormalizeAddress
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_normalize_address_proto_init() }
func file_normalize_address_proto_init() {
	if File_normalize_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_normalize_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestNormalizeAddress); i {
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
		file_normalize_address_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NormalizeAddressMS_Status); i {
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
		file_normalize_address_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NormalizeCodes); i {
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
		file_normalize_address_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NormalizedAddress); i {
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
		file_normalize_address_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseNormalizeAddress); i {
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
			RawDescriptor: file_normalize_address_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_normalize_address_proto_goTypes,
		DependencyIndexes: file_normalize_address_proto_depIdxs,
		EnumInfos:         file_normalize_address_proto_enumTypes,
		MessageInfos:      file_normalize_address_proto_msgTypes,
	}.Build()
	File_normalize_address_proto = out.File
	file_normalize_address_proto_rawDesc = nil
	file_normalize_address_proto_goTypes = nil
	file_normalize_address_proto_depIdxs = nil
}
