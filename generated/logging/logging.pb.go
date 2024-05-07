// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: logging.proto

package logging

import (
	types "github.com/cloudogu/ces-control-api/generated/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogLevel int32

const (
	LogLevel_DEBUG LogLevel = 0
	LogLevel_INFO  LogLevel = 1
	LogLevel_WARN  LogLevel = 2
	LogLevel_ERROR LogLevel = 3
)

// Enum value maps for LogLevel.
var (
	LogLevel_name = map[int32]string{
		0: "DEBUG",
		1: "INFO",
		2: "WARN",
		3: "ERROR",
	}
	LogLevel_value = map[string]int32{
		"DEBUG": 0,
		"INFO":  1,
		"WARN":  2,
		"ERROR": 3,
	}
)

func (x LogLevel) Enum() *LogLevel {
	p := new(LogLevel)
	*p = x
	return p
}

func (x LogLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_logging_proto_enumTypes[0].Descriptor()
}

func (LogLevel) Type() protoreflect.EnumType {
	return &file_logging_proto_enumTypes[0]
}

func (x LogLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogLevel.Descriptor instead.
func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return file_logging_proto_rawDescGZIP(), []int{0}
}

type DoguLogMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// dogu_name contains the name of the dogu for which log lines should be returned in the response. Must not be empty.
	DoguName string `protobuf:"bytes,1,opt,name=dogu_name,json=doguName,proto3" json:"dogu_name,omitempty"`
	// line_count provides a hint with the maximum count of log lines to be returned. The response contains the actual
	// number of lines (which can be less if the request asks for more lines than there really are).
	// A value of 0 means that all lines should be returned (use with care).
	LineCount uint32 `protobuf:"varint,2,opt,name=line_count,json=lineCount,proto3" json:"line_count,omitempty"`
}

func (x *DoguLogMessageRequest) Reset() {
	*x = DoguLogMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logging_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoguLogMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoguLogMessageRequest) ProtoMessage() {}

func (x *DoguLogMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logging_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoguLogMessageRequest.ProtoReflect.Descriptor instead.
func (*DoguLogMessageRequest) Descriptor() ([]byte, []int) {
	return file_logging_proto_rawDescGZIP(), []int{0}
}

func (x *DoguLogMessageRequest) GetDoguName() string {
	if x != nil {
		return x.DoguName
	}
	return ""
}

func (x *DoguLogMessageRequest) GetLineCount() uint32 {
	if x != nil {
		return x.LineCount
	}
	return 0
}

type DoguLogMessageQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// dogu_name contains the name of the dogu for which log lines should be returned in the response. Must not be empty.
	DoguName string `protobuf:"bytes,1,opt,name=dogu_name,json=doguName,proto3" json:"dogu_name,omitempty"`
	// start_date start datetime of the first log message to be retrieved. Optional, if empty is set to 30 days before the end_date.
	StartDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3,oneof" json:"start_date,omitempty"`
	// end_date end datetime of the last log message to be retrieved. Optional, if empty is set to current date.
	EndDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3,oneof" json:"end_date,omitempty"`
	// filter the log for a message, for example the log level. Optional, if empty no filter is applied
	// The filter simply selects all log-lines containing the given string.
	// The filter is case-sensitive.
	Filter *string `protobuf:"bytes,5,opt,name=filter,proto3,oneof" json:"filter,omitempty"`
}

func (x *DoguLogMessageQueryRequest) Reset() {
	*x = DoguLogMessageQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logging_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoguLogMessageQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoguLogMessageQueryRequest) ProtoMessage() {}

func (x *DoguLogMessageQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logging_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoguLogMessageQueryRequest.ProtoReflect.Descriptor instead.
func (*DoguLogMessageQueryRequest) Descriptor() ([]byte, []int) {
	return file_logging_proto_rawDescGZIP(), []int{1}
}

func (x *DoguLogMessageQueryRequest) GetDoguName() string {
	if x != nil {
		return x.DoguName
	}
	return ""
}

func (x *DoguLogMessageQueryRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *DoguLogMessageQueryRequest) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *DoguLogMessageQueryRequest) GetFilter() string {
	if x != nil && x.Filter != nil {
		return *x.Filter
	}
	return ""
}

type DoguLogMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// timestamp contains the timestamp of the log-message
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// message contains the content of logged message
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DoguLogMessage) Reset() {
	*x = DoguLogMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logging_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoguLogMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoguLogMessage) ProtoMessage() {}

func (x *DoguLogMessage) ProtoReflect() protoreflect.Message {
	mi := &file_logging_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoguLogMessage.ProtoReflect.Descriptor instead.
func (*DoguLogMessage) Descriptor() ([]byte, []int) {
	return file_logging_proto_rawDescGZIP(), []int{2}
}

func (x *DoguLogMessage) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *DoguLogMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type LogLevelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// dogu_name contains the name of the dogu for which log level should be set. Must not be empty.
	DoguName string `protobuf:"bytes,1,opt,name=dogu_name,json=doguName,proto3" json:"dogu_name,omitempty"`
	// the log level to be set.
	LogLevel LogLevel `protobuf:"varint,2,opt,name=log_level,json=logLevel,proto3,enum=logging.LogLevel" json:"log_level,omitempty"`
}

func (x *LogLevelRequest) Reset() {
	*x = LogLevelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logging_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogLevelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogLevelRequest) ProtoMessage() {}

func (x *LogLevelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logging_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogLevelRequest.ProtoReflect.Descriptor instead.
func (*LogLevelRequest) Descriptor() ([]byte, []int) {
	return file_logging_proto_rawDescGZIP(), []int{3}
}

func (x *LogLevelRequest) GetDoguName() string {
	if x != nil {
		return x.DoguName
	}
	return ""
}

func (x *LogLevelRequest) GetLogLevel() LogLevel {
	if x != nil {
		return x.LogLevel
	}
	return LogLevel_DEBUG
}

var File_logging_proto protoreflect.FileDescriptor

var file_logging_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x1a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x15, 0x44, 0x6f, 0x67, 0x75, 0x4c, 0x6f,
	0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x67, 0x75, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x67, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xf9, 0x01, 0x0a, 0x1a,
	0x44, 0x6f, 0x67, 0x75, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f,
	0x67, 0x75, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x67, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x64, 0x0a, 0x0e, 0x44, 0x6f, 0x67, 0x75, 0x4c,
	0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5e, 0x0a,
	0x0f, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x67, 0x75, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x67, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a,
	0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x2a, 0x34, 0x0a,
	0x08, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42,
	0x55, 0x47, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x57, 0x41, 0x52, 0x4e, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x03, 0x32, 0xee, 0x01, 0x0a, 0x0f, 0x44, 0x6f, 0x67, 0x75, 0x4c, 0x6f, 0x67, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46, 0x6f,
	0x72, 0x44, 0x6f, 0x67, 0x75, 0x12, 0x1e, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x44, 0x6f, 0x67, 0x75, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x6f, 0x72, 0x44,
	0x6f, 0x67, 0x75, 0x12, 0x23, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x6f,
	0x67, 0x75, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x44, 0x6f, 0x67, 0x75, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x30, 0x01, 0x12, 0x3f, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x6f, 0x67,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6f, 0x67, 0x75, 0x2f, 0x63, 0x65, 0x73, 0x2d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logging_proto_rawDescOnce sync.Once
	file_logging_proto_rawDescData = file_logging_proto_rawDesc
)

func file_logging_proto_rawDescGZIP() []byte {
	file_logging_proto_rawDescOnce.Do(func() {
		file_logging_proto_rawDescData = protoimpl.X.CompressGZIP(file_logging_proto_rawDescData)
	})
	return file_logging_proto_rawDescData
}

var file_logging_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_logging_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_logging_proto_goTypes = []interface{}{
	(LogLevel)(0),                      // 0: logging.LogLevel
	(*DoguLogMessageRequest)(nil),      // 1: logging.DoguLogMessageRequest
	(*DoguLogMessageQueryRequest)(nil), // 2: logging.DoguLogMessageQueryRequest
	(*DoguLogMessage)(nil),             // 3: logging.DoguLogMessage
	(*LogLevelRequest)(nil),            // 4: logging.LogLevelRequest
	(*timestamppb.Timestamp)(nil),      // 5: google.protobuf.Timestamp
	(*types.ChunkedDataResponse)(nil),  // 6: types.ChunkedDataResponse
	(*emptypb.Empty)(nil),              // 7: google.protobuf.Empty
}
var file_logging_proto_depIdxs = []int32{
	5, // 0: logging.DoguLogMessageQueryRequest.start_date:type_name -> google.protobuf.Timestamp
	5, // 1: logging.DoguLogMessageQueryRequest.end_date:type_name -> google.protobuf.Timestamp
	5, // 2: logging.DoguLogMessage.timestamp:type_name -> google.protobuf.Timestamp
	0, // 3: logging.LogLevelRequest.log_level:type_name -> logging.LogLevel
	1, // 4: logging.DoguLogMessages.GetForDogu:input_type -> logging.DoguLogMessageRequest
	2, // 5: logging.DoguLogMessages.QueryForDogu:input_type -> logging.DoguLogMessageQueryRequest
	4, // 6: logging.DoguLogMessages.SetLogLevel:input_type -> logging.LogLevelRequest
	6, // 7: logging.DoguLogMessages.GetForDogu:output_type -> types.ChunkedDataResponse
	3, // 8: logging.DoguLogMessages.QueryForDogu:output_type -> logging.DoguLogMessage
	7, // 9: logging.DoguLogMessages.SetLogLevel:output_type -> google.protobuf.Empty
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_logging_proto_init() }
func file_logging_proto_init() {
	if File_logging_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logging_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoguLogMessageRequest); i {
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
		file_logging_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoguLogMessageQueryRequest); i {
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
		file_logging_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoguLogMessage); i {
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
		file_logging_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogLevelRequest); i {
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
	file_logging_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_logging_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logging_proto_goTypes,
		DependencyIndexes: file_logging_proto_depIdxs,
		EnumInfos:         file_logging_proto_enumTypes,
		MessageInfos:      file_logging_proto_msgTypes,
	}.Build()
	File_logging_proto = out.File
	file_logging_proto_rawDesc = nil
	file_logging_proto_goTypes = nil
	file_logging_proto_depIdxs = nil
}
