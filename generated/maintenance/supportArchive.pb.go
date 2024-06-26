// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: supportArchive.proto

package maintenance

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

type CreateSupportArchiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CollectEtcd defines whether etcd keys/values should be contained in the final archive (true) or nor (false)
	CollectEtcd bool `protobuf:"varint,1,opt,name=collectEtcd,proto3" json:"collectEtcd,omitempty"`
	// CollectOsInformation defines whether information about soft- and hardware should be contained in the final archive (true) or nor (false)
	CollectOsInformation bool `protobuf:"varint,2,opt,name=collectOsInformation,proto3" json:"collectOsInformation,omitempty"`
	// CollectEnvironment defines whether the environment variables should be contained in the final archive (true) or nor (false)
	CollectEnvironment bool `protobuf:"varint,3,opt,name=collectEnvironment,proto3" json:"collectEnvironment,omitempty"`
	// CollectDoguStatusInformation defines whether information about installed dogus should be contained in the final archive (true) or nor (false)
	CollectDoguStatusInformation bool `protobuf:"varint,4,opt,name=collectDoguStatusInformation,proto3" json:"collectDoguStatusInformation,omitempty"`
	// CollectProcessInformation defines whether the Process and service status should be contained in the final archive (true) or nor (false)
	CollectProcessInformation bool `protobuf:"varint,5,opt,name=collectProcessInformation,proto3" json:"collectProcessInformation,omitempty"`
	// CollectVolumeInformation defines whether the volume information should be contained in the final archive (true) or nor (false)
	CollectVolumeInformation bool `protobuf:"varint,6,opt,name=collectVolumeInformation,proto3" json:"collectVolumeInformation,omitempty"`
	// ExcludedEtcdKeys is a list of etcd keys which should not be included in the final archive.
	ExcludedEtcdKeys []string `protobuf:"bytes,7,rep,name=excludedEtcdKeys,proto3" json:"excludedEtcdKeys,omitempty"`
	// FailFast defines whether the archive collection should fail at the first error (true) or not (false).
	FailFast bool `protobuf:"varint,8,opt,name=failFast,proto3" json:"failFast,omitempty"`
	// CollectLogsFromTheseDogus must be a list of dogus. Only the logs of these dogus are collected - all other dogu logs are skipped.
	CollectLogsFromTheseDogus []string `protobuf:"bytes,9,rep,name=collectLogsFromTheseDogus,proto3" json:"collectLogsFromTheseDogus,omitempty"`
	// CollectCesappLog defines whether the cesapp logfile should be collected (true) or not (false).
	CollectCesappLog bool `protobuf:"varint,10,opt,name=collectCesappLog,proto3" json:"collectCesappLog,omitempty"`
	// CollectCesappdLog defines whether the cesappd logfile should be collected (true) or not (false).
	CollectCesappdLog bool `protobuf:"varint,11,opt,name=collectCesappdLog,proto3" json:"collectCesappdLog,omitempty"`
	// CollectSyslog defines whether the syslog logfile should be collected (true) or not (false).
	CollectSyslog bool `protobuf:"varint,12,opt,name=collectSyslog,proto3" json:"collectSyslog,omitempty"`
	// CollectAptLog defines whether the apt logfiles should be collected (true) or not (false).
	CollectAptLog bool `protobuf:"varint,13,opt,name=collectAptLog,proto3" json:"collectAptLog,omitempty"`
	// CollectBackupWatcherLog defines whether the backup-watcher logfile should be collected (true) or not (false).
	CollectBackupWatcherLog bool `protobuf:"varint,14,opt,name=collectBackupWatcherLog,proto3" json:"collectBackupWatcherLog,omitempty"`
}

func (x *CreateSupportArchiveRequest) Reset() {
	*x = CreateSupportArchiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supportArchive_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSupportArchiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSupportArchiveRequest) ProtoMessage() {}

func (x *CreateSupportArchiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supportArchive_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSupportArchiveRequest.ProtoReflect.Descriptor instead.
func (*CreateSupportArchiveRequest) Descriptor() ([]byte, []int) {
	return file_supportArchive_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSupportArchiveRequest) GetCollectEtcd() bool {
	if x != nil {
		return x.CollectEtcd
	}
	return false
}

func (x *CreateSupportArchiveRequest) GetCollectOsInformation() bool {
	if x != nil {
		return x.CollectOsInformation
	}
	return false
}

func (x *CreateSupportArchiveRequest) GetCollectEnvironment() bool {
	if x != nil {
		return x.CollectEnvironment
	}
	return false
}

func (x *CreateSupportArchiveRequest) GetCollectDoguStatusInformation() bool {
	if x != nil {
		return x.CollectDoguStatusInformation
	}
	return false
}

func (x *CreateSupportArchiveRequest) GetCollectProcessInformation() bool {
	if x != nil {
		return x.CollectProcessInformation
	}
	return false
}

func (x *CreateSupportArchiveRequest) GetCollectVolumeInformation() bool {
	if x != nil {
		return x.CollectVolumeInformation
	}
	return false
}

func (x *CreateSupportArchiveRequest) GetExcludedEtcdKeys() []string {
	if x != nil {
		return x.ExcludedEtcdKeys
	}
	return nil
}

func (x *CreateSupportArchiveRequest) GetFailFast() bool {
	if x != nil {
		return x.FailFast
	}
	return false
}

func (x *CreateSupportArchiveRequest) GetCollectLogsFromTheseDogus() []string {
	if x != nil {
		return x.CollectLogsFromTheseDogus
	}
	return nil
}

func (x *CreateSupportArchiveRequest) GetCollectCesappLog() bool {
	if x != nil {
		return x.CollectCesappLog
	}
	return false
}

func (x *CreateSupportArchiveRequest) GetCollectCesappdLog() bool {
	if x != nil {
		return x.CollectCesappdLog
	}
	return false
}

func (x *CreateSupportArchiveRequest) GetCollectSyslog() bool {
	if x != nil {
		return x.CollectSyslog
	}
	return false
}

func (x *CreateSupportArchiveRequest) GetCollectAptLog() bool {
	if x != nil {
		return x.CollectAptLog
	}
	return false
}

func (x *CreateSupportArchiveRequest) GetCollectBackupWatcherLog() bool {
	if x != nil {
		return x.CollectBackupWatcherLog
	}
	return false
}

var File_supportArchive_proto protoreflect.FileDescriptor

var file_supportArchive_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x1a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x05, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x45, 0x74,
	0x63, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x45, 0x74, 0x63, 0x64, 0x12, 0x32, 0x0a, 0x14, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x4f, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x14, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4f, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x1c, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x44, 0x6f, 0x67, 0x75, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x1c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x44, 0x6f, 0x67, 0x75, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a,
	0x19, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x19, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x18, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x78, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x64, 0x45, 0x74, 0x63, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x10, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x45, 0x74, 0x63, 0x64, 0x4b,
	0x65, 0x79, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x46, 0x61, 0x73, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x46, 0x61, 0x73, 0x74, 0x12,
	0x3c, 0x0a, 0x19, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x46, 0x72,
	0x6f, 0x6d, 0x54, 0x68, 0x65, 0x73, 0x65, 0x44, 0x6f, 0x67, 0x75, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x19, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x46,
	0x72, 0x6f, 0x6d, 0x54, 0x68, 0x65, 0x73, 0x65, 0x44, 0x6f, 0x67, 0x75, 0x73, 0x12, 0x2a, 0x0a,
	0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x65, 0x73, 0x61, 0x70, 0x70, 0x4c, 0x6f,
	0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x43, 0x65, 0x73, 0x61, 0x70, 0x70, 0x4c, 0x6f, 0x67, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x43, 0x65, 0x73, 0x61, 0x70, 0x70, 0x64, 0x4c, 0x6f, 0x67, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x65, 0x73,
	0x61, 0x70, 0x70, 0x64, 0x4c, 0x6f, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x53, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x53, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x12, 0x24, 0x0a,
	0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x41, 0x70, 0x74, 0x4c, 0x6f, 0x67, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x41, 0x70, 0x74,
	0x4c, 0x6f, 0x67, 0x12, 0x38, 0x0a, 0x17, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x32, 0x62, 0x0a,
	0x0e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x12,
	0x50, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6f, 0x67, 0x75, 0x2f, 0x63, 0x65, 0x73, 0x2d, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_supportArchive_proto_rawDescOnce sync.Once
	file_supportArchive_proto_rawDescData = file_supportArchive_proto_rawDesc
)

func file_supportArchive_proto_rawDescGZIP() []byte {
	file_supportArchive_proto_rawDescOnce.Do(func() {
		file_supportArchive_proto_rawDescData = protoimpl.X.CompressGZIP(file_supportArchive_proto_rawDescData)
	})
	return file_supportArchive_proto_rawDescData
}

var file_supportArchive_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_supportArchive_proto_goTypes = []interface{}{
	(*CreateSupportArchiveRequest)(nil), // 0: maintenance.CreateSupportArchiveRequest
	(*types.ChunkedDataResponse)(nil),   // 1: types.ChunkedDataResponse
}
var file_supportArchive_proto_depIdxs = []int32{
	0, // 0: maintenance.SupportArchive.Create:input_type -> maintenance.CreateSupportArchiveRequest
	1, // 1: maintenance.SupportArchive.Create:output_type -> types.ChunkedDataResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_supportArchive_proto_init() }
func file_supportArchive_proto_init() {
	if File_supportArchive_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_supportArchive_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSupportArchiveRequest); i {
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
			RawDescriptor: file_supportArchive_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_supportArchive_proto_goTypes,
		DependencyIndexes: file_supportArchive_proto_depIdxs,
		MessageInfos:      file_supportArchive_proto_msgTypes,
	}.Build()
	File_supportArchive_proto = out.File
	file_supportArchive_proto_rawDesc = nil
	file_supportArchive_proto_goTypes = nil
	file_supportArchive_proto_depIdxs = nil
}
