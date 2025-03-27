// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: boilerplate.proto

package gapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Boilerplate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Boilerplate) Reset() {
	*x = Boilerplate{}
	mi := &file_boilerplate_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Boilerplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Boilerplate) ProtoMessage() {}

func (x *Boilerplate) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Boilerplate.ProtoReflect.Descriptor instead.
func (*Boilerplate) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{0}
}

func (x *Boilerplate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Boilerplate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Boilerplate) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Boilerplate) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateBoilerplateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBoilerplateRequest) Reset() {
	*x = CreateBoilerplateRequest{}
	mi := &file_boilerplate_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBoilerplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoilerplateRequest) ProtoMessage() {}

func (x *CreateBoilerplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBoilerplateRequest.ProtoReflect.Descriptor instead.
func (*CreateBoilerplateRequest) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBoilerplateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateBoilerplateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBoilerplateResponse) Reset() {
	*x = CreateBoilerplateResponse{}
	mi := &file_boilerplate_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBoilerplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoilerplateResponse) ProtoMessage() {}

func (x *CreateBoilerplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBoilerplateResponse.ProtoReflect.Descriptor instead.
func (*CreateBoilerplateResponse) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBoilerplateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateBoilerplateResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetBoilerplateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBoilerplateRequest) Reset() {
	*x = GetBoilerplateRequest{}
	mi := &file_boilerplate_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBoilerplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBoilerplateRequest) ProtoMessage() {}

func (x *GetBoilerplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBoilerplateRequest.ProtoReflect.Descriptor instead.
func (*GetBoilerplateRequest) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{3}
}

func (x *GetBoilerplateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBoilerplateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBoilerplateResponse) Reset() {
	*x = GetBoilerplateResponse{}
	mi := &file_boilerplate_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBoilerplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBoilerplateResponse) ProtoMessage() {}

func (x *GetBoilerplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBoilerplateResponse.ProtoReflect.Descriptor instead.
func (*GetBoilerplateResponse) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{4}
}

func (x *GetBoilerplateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetBoilerplateResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateBoilerplateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBoilerplateRequest) Reset() {
	*x = UpdateBoilerplateRequest{}
	mi := &file_boilerplate_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBoilerplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBoilerplateRequest) ProtoMessage() {}

func (x *UpdateBoilerplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBoilerplateRequest.ProtoReflect.Descriptor instead.
func (*UpdateBoilerplateRequest) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBoilerplateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBoilerplateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateBoilerplateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Updated       string                 `protobuf:"bytes,1,opt,name=updated,proto3" json:"updated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBoilerplateResponse) Reset() {
	*x = UpdateBoilerplateResponse{}
	mi := &file_boilerplate_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBoilerplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBoilerplateResponse) ProtoMessage() {}

func (x *UpdateBoilerplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBoilerplateResponse.ProtoReflect.Descriptor instead.
func (*UpdateBoilerplateResponse) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBoilerplateResponse) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

type DeleteBoilerplateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteBoilerplateRequest) Reset() {
	*x = DeleteBoilerplateRequest{}
	mi := &file_boilerplate_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBoilerplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBoilerplateRequest) ProtoMessage() {}

func (x *DeleteBoilerplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBoilerplateRequest.ProtoReflect.Descriptor instead.
func (*DeleteBoilerplateRequest) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBoilerplateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteBoilerplateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Deleted       string                 `protobuf:"bytes,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteBoilerplateResponse) Reset() {
	*x = DeleteBoilerplateResponse{}
	mi := &file_boilerplate_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBoilerplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBoilerplateResponse) ProtoMessage() {}

func (x *DeleteBoilerplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBoilerplateResponse.ProtoReflect.Descriptor instead.
func (*DeleteBoilerplateResponse) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteBoilerplateResponse) GetDeleted() string {
	if x != nil {
		return x.Deleted
	}
	return ""
}

type ListBoilerplatesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListBoilerplatesRequest) Reset() {
	*x = ListBoilerplatesRequest{}
	mi := &file_boilerplate_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBoilerplatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBoilerplatesRequest) ProtoMessage() {}

func (x *ListBoilerplatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBoilerplatesRequest.ProtoReflect.Descriptor instead.
func (*ListBoilerplatesRequest) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{9}
}

type ListBoilerplatesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Boilerplates  []*Boilerplate         `protobuf:"bytes,1,rep,name=boilerplates,proto3" json:"boilerplates,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListBoilerplatesResponse) Reset() {
	*x = ListBoilerplatesResponse{}
	mi := &file_boilerplate_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBoilerplatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBoilerplatesResponse) ProtoMessage() {}

func (x *ListBoilerplatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBoilerplatesResponse.ProtoReflect.Descriptor instead.
func (*ListBoilerplatesResponse) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{10}
}

func (x *ListBoilerplatesResponse) GetBoilerplates() []*Boilerplate {
	if x != nil {
		return x.Boilerplates
	}
	return nil
}

type SearchBoilerplatesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SearchTerm    string                 `protobuf:"bytes,1,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	PageNumber    int32                  `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize      int32                  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchBoilerplatesRequest) Reset() {
	*x = SearchBoilerplatesRequest{}
	mi := &file_boilerplate_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchBoilerplatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchBoilerplatesRequest) ProtoMessage() {}

func (x *SearchBoilerplatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchBoilerplatesRequest.ProtoReflect.Descriptor instead.
func (*SearchBoilerplatesRequest) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{11}
}

func (x *SearchBoilerplatesRequest) GetSearchTerm() string {
	if x != nil {
		return x.SearchTerm
	}
	return ""
}

func (x *SearchBoilerplatesRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *SearchBoilerplatesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type SearchBoilerplatesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Boilerplates  []*Boilerplate         `protobuf:"bytes,1,rep,name=boilerplates,proto3" json:"boilerplates,omitempty"`
	TotalResults  int32                  `protobuf:"varint,2,opt,name=total_results,json=totalResults,proto3" json:"total_results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchBoilerplatesResponse) Reset() {
	*x = SearchBoilerplatesResponse{}
	mi := &file_boilerplate_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchBoilerplatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchBoilerplatesResponse) ProtoMessage() {}

func (x *SearchBoilerplatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchBoilerplatesResponse.ProtoReflect.Descriptor instead.
func (*SearchBoilerplatesResponse) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{12}
}

func (x *SearchBoilerplatesResponse) GetBoilerplates() []*Boilerplate {
	if x != nil {
		return x.Boilerplates
	}
	return nil
}

func (x *SearchBoilerplatesResponse) GetTotalResults() int32 {
	if x != nil {
		return x.TotalResults
	}
	return 0
}

var File_boilerplate_proto protoreflect.FileDescriptor

var file_boilerplate_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x0b, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2e, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x19, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x69, 0x6c,
	0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x69,
	0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x69,
	0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x2a, 0x0a, 0x18, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x19, 0x0a,
	0x17, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x58, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x6f, 0x69,
	0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x22, 0x7a, 0x0a, 0x19, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x69, 0x6c,
	0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x7f,
	0x0a, 0x1a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c,
	0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x62, 0x6f,
	0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32,
	0xe3, 0x04, 0x0a, 0x12, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x62, 0x6f,
	0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x62,
	0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x62, 0x6f, 0x69,
	0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x25,
	0x2e, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x24, 0x2e, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65,
	0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x62,
	0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x69, 0x74, 0x6f, 0x73, 0x2d, 0x65, 0x73, 0x2f,
	0x6d, 0x61, 0x72, 0x6b, 0x69, 0x74, 0x6f, 0x73, 0x2d, 0x73, 0x76, 0x63, 0x2d, 0x62, 0x6f, 0x69,
	0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x67, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_boilerplate_proto_rawDescOnce sync.Once
	file_boilerplate_proto_rawDescData []byte
)

func file_boilerplate_proto_rawDescGZIP() []byte {
	file_boilerplate_proto_rawDescOnce.Do(func() {
		file_boilerplate_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_boilerplate_proto_rawDesc), len(file_boilerplate_proto_rawDesc)))
	})
	return file_boilerplate_proto_rawDescData
}

var file_boilerplate_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_boilerplate_proto_goTypes = []any{
	(*Boilerplate)(nil),                // 0: boilerplate.Boilerplate
	(*CreateBoilerplateRequest)(nil),   // 1: boilerplate.CreateBoilerplateRequest
	(*CreateBoilerplateResponse)(nil),  // 2: boilerplate.CreateBoilerplateResponse
	(*GetBoilerplateRequest)(nil),      // 3: boilerplate.GetBoilerplateRequest
	(*GetBoilerplateResponse)(nil),     // 4: boilerplate.GetBoilerplateResponse
	(*UpdateBoilerplateRequest)(nil),   // 5: boilerplate.UpdateBoilerplateRequest
	(*UpdateBoilerplateResponse)(nil),  // 6: boilerplate.UpdateBoilerplateResponse
	(*DeleteBoilerplateRequest)(nil),   // 7: boilerplate.DeleteBoilerplateRequest
	(*DeleteBoilerplateResponse)(nil),  // 8: boilerplate.DeleteBoilerplateResponse
	(*ListBoilerplatesRequest)(nil),    // 9: boilerplate.ListBoilerplatesRequest
	(*ListBoilerplatesResponse)(nil),   // 10: boilerplate.ListBoilerplatesResponse
	(*SearchBoilerplatesRequest)(nil),  // 11: boilerplate.SearchBoilerplatesRequest
	(*SearchBoilerplatesResponse)(nil), // 12: boilerplate.SearchBoilerplatesResponse
	(*timestamppb.Timestamp)(nil),      // 13: google.protobuf.Timestamp
}
var file_boilerplate_proto_depIdxs = []int32{
	13, // 0: boilerplate.Boilerplate.created_at:type_name -> google.protobuf.Timestamp
	13, // 1: boilerplate.Boilerplate.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: boilerplate.ListBoilerplatesResponse.boilerplates:type_name -> boilerplate.Boilerplate
	0,  // 3: boilerplate.SearchBoilerplatesResponse.boilerplates:type_name -> boilerplate.Boilerplate
	1,  // 4: boilerplate.BoilerplateService.CreateBoilerplate:input_type -> boilerplate.CreateBoilerplateRequest
	3,  // 5: boilerplate.BoilerplateService.GetBoilerplate:input_type -> boilerplate.GetBoilerplateRequest
	5,  // 6: boilerplate.BoilerplateService.UpdateBoilerplate:input_type -> boilerplate.UpdateBoilerplateRequest
	7,  // 7: boilerplate.BoilerplateService.DeleteBoilerplate:input_type -> boilerplate.DeleteBoilerplateRequest
	9,  // 8: boilerplate.BoilerplateService.ListBoilerplates:input_type -> boilerplate.ListBoilerplatesRequest
	11, // 9: boilerplate.BoilerplateService.SearchBoilerplates:input_type -> boilerplate.SearchBoilerplatesRequest
	2,  // 10: boilerplate.BoilerplateService.CreateBoilerplate:output_type -> boilerplate.CreateBoilerplateResponse
	4,  // 11: boilerplate.BoilerplateService.GetBoilerplate:output_type -> boilerplate.GetBoilerplateResponse
	6,  // 12: boilerplate.BoilerplateService.UpdateBoilerplate:output_type -> boilerplate.UpdateBoilerplateResponse
	8,  // 13: boilerplate.BoilerplateService.DeleteBoilerplate:output_type -> boilerplate.DeleteBoilerplateResponse
	10, // 14: boilerplate.BoilerplateService.ListBoilerplates:output_type -> boilerplate.ListBoilerplatesResponse
	12, // 15: boilerplate.BoilerplateService.SearchBoilerplates:output_type -> boilerplate.SearchBoilerplatesResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_boilerplate_proto_init() }
func file_boilerplate_proto_init() {
	if File_boilerplate_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_boilerplate_proto_rawDesc), len(file_boilerplate_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_boilerplate_proto_goTypes,
		DependencyIndexes: file_boilerplate_proto_depIdxs,
		MessageInfos:      file_boilerplate_proto_msgTypes,
	}.Build()
	File_boilerplate_proto = out.File
	file_boilerplate_proto_goTypes = nil
	file_boilerplate_proto_depIdxs = nil
}
