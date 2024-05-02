// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: doguAdministration.proto

package doguAdministration

import (
	types "github.com/cloudogu/ces-control-api/generated/types"
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

// DoguListRequest is empty and is used instead of google.protobuf.Empty because of backward compability
type DoguListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DoguListRequest) Reset() {
	*x = DoguListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doguAdministration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoguListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoguListRequest) ProtoMessage() {}

func (x *DoguListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doguAdministration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoguListRequest.ProtoReflect.Descriptor instead.
func (*DoguListRequest) Descriptor() ([]byte, []int) {
	return file_doguAdministration_proto_rawDescGZIP(), []int{0}
}

type DoguListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of dogus
	Dogus []*Dogu `protobuf:"bytes,1,rep,name=dogus,proto3" json:"dogus,omitempty"`
}

func (x *DoguListResponse) Reset() {
	*x = DoguListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doguAdministration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoguListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoguListResponse) ProtoMessage() {}

func (x *DoguListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doguAdministration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoguListResponse.ProtoReflect.Descriptor instead.
func (*DoguListResponse) Descriptor() ([]byte, []int) {
	return file_doguAdministration_proto_rawDescGZIP(), []int{1}
}

func (x *DoguListResponse) GetDogus() []*Dogu {
	if x != nil {
		return x.Dogus
	}
	return nil
}

type DoguAdministrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoguName string `protobuf:"bytes,1,opt,name=dogu_name,json=doguName,proto3" json:"dogu_name,omitempty"`
}

func (x *DoguAdministrationRequest) Reset() {
	*x = DoguAdministrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doguAdministration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoguAdministrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoguAdministrationRequest) ProtoMessage() {}

func (x *DoguAdministrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doguAdministration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoguAdministrationRequest.ProtoReflect.Descriptor instead.
func (*DoguAdministrationRequest) Descriptor() ([]byte, []int) {
	return file_doguAdministration_proto_rawDescGZIP(), []int{2}
}

func (x *DoguAdministrationRequest) GetDoguName() string {
	if x != nil {
		return x.DoguName
	}
	return ""
}

type Dogu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName string   `protobuf:"bytes,2,opt,name=displayName,proto3" json:"displayName,omitempty"`
	Version     string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Tags        []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Dogu) Reset() {
	*x = Dogu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doguAdministration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dogu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dogu) ProtoMessage() {}

func (x *Dogu) ProtoReflect() protoreflect.Message {
	mi := &file_doguAdministration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dogu.ProtoReflect.Descriptor instead.
func (*Dogu) Descriptor() ([]byte, []int) {
	return file_doguAdministration_proto_rawDescGZIP(), []int{3}
}

func (x *Dogu) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dogu) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Dogu) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Dogu) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Dogu) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// DoguBlueprintId request is empty and is used instead of google.protobuf.Empty because of backward compability
type DoguBlueprinitIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DoguBlueprinitIdRequest) Reset() {
	*x = DoguBlueprinitIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doguAdministration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoguBlueprinitIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoguBlueprinitIdRequest) ProtoMessage() {}

func (x *DoguBlueprinitIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doguAdministration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoguBlueprinitIdRequest.ProtoReflect.Descriptor instead.
func (*DoguBlueprinitIdRequest) Descriptor() ([]byte, []int) {
	return file_doguAdministration_proto_rawDescGZIP(), []int{4}
}

type DoguBlueprintIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlueprintId string `protobuf:"bytes,1,opt,name=blueprint_id,json=blueprintId,proto3" json:"blueprint_id,omitempty"`
}

func (x *DoguBlueprintIdResponse) Reset() {
	*x = DoguBlueprintIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doguAdministration_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoguBlueprintIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoguBlueprintIdResponse) ProtoMessage() {}

func (x *DoguBlueprintIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doguAdministration_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoguBlueprintIdResponse.ProtoReflect.Descriptor instead.
func (*DoguBlueprintIdResponse) Descriptor() ([]byte, []int) {
	return file_doguAdministration_proto_rawDescGZIP(), []int{5}
}

func (x *DoguBlueprintIdResponse) GetBlueprintId() string {
	if x != nil {
		return x.BlueprintId
	}
	return ""
}

var File_doguAdministration_proto protoreflect.FileDescriptor

var file_doguAdministration_proto_rawDesc = []byte{
	0x0a, 0x18, 0x64, 0x6f, 0x67, 0x75, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x6f, 0x67, 0x75,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11,
	0x0a, 0x0f, 0x44, 0x6f, 0x67, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x42, 0x0a, 0x10, 0x44, 0x6f, 0x67, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x64, 0x6f, 0x67, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x6f, 0x67, 0x75, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x6f, 0x67, 0x75, 0x52, 0x05,
	0x64, 0x6f, 0x67, 0x75, 0x73, 0x22, 0x38, 0x0a, 0x19, 0x44, 0x6f, 0x67, 0x75, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x67, 0x75, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x67, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x8c, 0x01, 0x0a, 0x04, 0x44, 0x6f, 0x67, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x19,
	0x0a, 0x17, 0x44, 0x6f, 0x67, 0x75, 0x42, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x69, 0x74,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x17, 0x44, 0x6f, 0x67,
	0x75, 0x42, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6c, 0x75, 0x65,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x32, 0xd1, 0x03, 0x0a, 0x12, 0x44, 0x6f, 0x67, 0x75,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x67, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e,
	0x64, 0x6f, 0x67, 0x75, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x44, 0x6f, 0x67, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x6f, 0x67, 0x75, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x6f, 0x67, 0x75, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x2e, 0x64, 0x6f, 0x67,
	0x75, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x44, 0x6f, 0x67, 0x75, 0x42, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x69, 0x74, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x64, 0x6f, 0x67, 0x75, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x6f, 0x67,
	0x75, 0x42, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x6f, 0x67,
	0x75, 0x12, 0x2d, 0x2e, 0x64, 0x6f, 0x67, 0x75, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x6f, 0x67, 0x75, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x08, 0x53, 0x74, 0x6f, 0x70, 0x44, 0x6f,
	0x67, 0x75, 0x12, 0x2d, 0x2e, 0x64, 0x6f, 0x67, 0x75, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x6f, 0x67, 0x75, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x44, 0x6f, 0x67, 0x75, 0x12, 0x2d, 0x2e, 0x64, 0x6f, 0x67, 0x75, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x6f, 0x67, 0x75,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x42, 0x5a, 0x40, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6f,
	0x67, 0x75, 0x2f, 0x63, 0x65, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x64, 0x6f, 0x67,
	0x75, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doguAdministration_proto_rawDescOnce sync.Once
	file_doguAdministration_proto_rawDescData = file_doguAdministration_proto_rawDesc
)

func file_doguAdministration_proto_rawDescGZIP() []byte {
	file_doguAdministration_proto_rawDescOnce.Do(func() {
		file_doguAdministration_proto_rawDescData = protoimpl.X.CompressGZIP(file_doguAdministration_proto_rawDescData)
	})
	return file_doguAdministration_proto_rawDescData
}

var file_doguAdministration_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_doguAdministration_proto_goTypes = []interface{}{
	(*DoguListRequest)(nil),           // 0: doguAdministration.DoguListRequest
	(*DoguListResponse)(nil),          // 1: doguAdministration.DoguListResponse
	(*DoguAdministrationRequest)(nil), // 2: doguAdministration.DoguAdministrationRequest
	(*Dogu)(nil),                      // 3: doguAdministration.Dogu
	(*DoguBlueprinitIdRequest)(nil),   // 4: doguAdministration.DoguBlueprinitIdRequest
	(*DoguBlueprintIdResponse)(nil),   // 5: doguAdministration.DoguBlueprintIdResponse
	(*types.BasicResponse)(nil),       // 6: types.BasicResponse
}
var file_doguAdministration_proto_depIdxs = []int32{
	3, // 0: doguAdministration.DoguListResponse.dogus:type_name -> doguAdministration.Dogu
	0, // 1: doguAdministration.DoguAdministration.GetDoguList:input_type -> doguAdministration.DoguListRequest
	4, // 2: doguAdministration.DoguAdministration.GetBlueprintId:input_type -> doguAdministration.DoguBlueprinitIdRequest
	2, // 3: doguAdministration.DoguAdministration.StartDogu:input_type -> doguAdministration.DoguAdministrationRequest
	2, // 4: doguAdministration.DoguAdministration.StopDogu:input_type -> doguAdministration.DoguAdministrationRequest
	2, // 5: doguAdministration.DoguAdministration.RestartDogu:input_type -> doguAdministration.DoguAdministrationRequest
	1, // 6: doguAdministration.DoguAdministration.GetDoguList:output_type -> doguAdministration.DoguListResponse
	5, // 7: doguAdministration.DoguAdministration.GetBlueprintId:output_type -> doguAdministration.DoguBlueprintIdResponse
	6, // 8: doguAdministration.DoguAdministration.StartDogu:output_type -> types.BasicResponse
	6, // 9: doguAdministration.DoguAdministration.StopDogu:output_type -> types.BasicResponse
	6, // 10: doguAdministration.DoguAdministration.RestartDogu:output_type -> types.BasicResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_doguAdministration_proto_init() }
func file_doguAdministration_proto_init() {
	if File_doguAdministration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doguAdministration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoguListRequest); i {
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
		file_doguAdministration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoguListResponse); i {
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
		file_doguAdministration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoguAdministrationRequest); i {
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
		file_doguAdministration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dogu); i {
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
		file_doguAdministration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoguBlueprinitIdRequest); i {
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
		file_doguAdministration_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoguBlueprintIdResponse); i {
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
			RawDescriptor: file_doguAdministration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doguAdministration_proto_goTypes,
		DependencyIndexes: file_doguAdministration_proto_depIdxs,
		MessageInfos:      file_doguAdministration_proto_msgTypes,
	}.Build()
	File_doguAdministration_proto = out.File
	file_doguAdministration_proto_rawDesc = nil
	file_doguAdministration_proto_goTypes = nil
	file_doguAdministration_proto_depIdxs = nil
}
