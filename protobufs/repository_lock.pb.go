// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.23.3
// source: protobufs/repository_lock.proto

package protobufs

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

type ReportSaveRawSuccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId            string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FilePath          string `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	ContentChunkIndex int32  `protobuf:"varint,3,opt,name=content_chunk_index,json=contentChunkIndex,proto3" json:"content_chunk_index,omitempty"`
	NumTotalChunks    int32  `protobuf:"varint,4,opt,name=num_total_chunks,json=numTotalChunks,proto3" json:"num_total_chunks,omitempty"`
}

func (x *ReportSaveRawSuccessRequest) Reset() {
	*x = ReportSaveRawSuccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufs_repository_lock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportSaveRawSuccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportSaveRawSuccessRequest) ProtoMessage() {}

func (x *ReportSaveRawSuccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobufs_repository_lock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportSaveRawSuccessRequest.ProtoReflect.Descriptor instead.
func (*ReportSaveRawSuccessRequest) Descriptor() ([]byte, []int) {
	return file_protobufs_repository_lock_proto_rawDescGZIP(), []int{0}
}

func (x *ReportSaveRawSuccessRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReportSaveRawSuccessRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *ReportSaveRawSuccessRequest) GetContentChunkIndex() int32 {
	if x != nil {
		return x.ContentChunkIndex
	}
	return 0
}

func (x *ReportSaveRawSuccessRequest) GetNumTotalChunks() int32 {
	if x != nil {
		return x.NumTotalChunks
	}
	return 0
}

type ReportSaveRawSuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath       string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	AllChunksSaved bool   `protobuf:"varint,2,opt,name=all_chunks_saved,json=allChunksSaved,proto3" json:"all_chunks_saved,omitempty"`
}

func (x *ReportSaveRawSuccessResponse) Reset() {
	*x = ReportSaveRawSuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufs_repository_lock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportSaveRawSuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportSaveRawSuccessResponse) ProtoMessage() {}

func (x *ReportSaveRawSuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobufs_repository_lock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportSaveRawSuccessResponse.ProtoReflect.Descriptor instead.
func (*ReportSaveRawSuccessResponse) Descriptor() ([]byte, []int) {
	return file_protobufs_repository_lock_proto_rawDescGZIP(), []int{1}
}

func (x *ReportSaveRawSuccessResponse) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *ReportSaveRawSuccessResponse) GetAllChunksSaved() bool {
	if x != nil {
		return x.AllChunksSaved
	}
	return false
}

var File_protobufs_repository_lock_proto protoreflect.FileDescriptor

var file_protobufs_repository_lock_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x22, 0xad, 0x01, 0x0a,
	0x1b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x61, 0x76, 0x65, 0x52, 0x61, 0x77, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x11, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6e, 0x75,
	0x6d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x22, 0x65, 0x0a, 0x1c,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x61, 0x76, 0x65, 0x52, 0x61, 0x77, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x6c, 0x6c,
	0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x53, 0x61,
	0x76, 0x65, 0x64, 0x32, 0x7d, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x6b, 0x0a, 0x14, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x53,
	0x61, 0x76, 0x65, 0x52, 0x61, 0x77, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x26, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x53, 0x61, 0x76, 0x65, 0x52, 0x61, 0x77, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x61, 0x76, 0x65, 0x52, 0x61, 0x77, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobufs_repository_lock_proto_rawDescOnce sync.Once
	file_protobufs_repository_lock_proto_rawDescData = file_protobufs_repository_lock_proto_rawDesc
)

func file_protobufs_repository_lock_proto_rawDescGZIP() []byte {
	file_protobufs_repository_lock_proto_rawDescOnce.Do(func() {
		file_protobufs_repository_lock_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobufs_repository_lock_proto_rawDescData)
	})
	return file_protobufs_repository_lock_proto_rawDescData
}

var file_protobufs_repository_lock_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobufs_repository_lock_proto_goTypes = []interface{}{
	(*ReportSaveRawSuccessRequest)(nil),  // 0: protobufs.ReportSaveRawSuccessRequest
	(*ReportSaveRawSuccessResponse)(nil), // 1: protobufs.ReportSaveRawSuccessResponse
}
var file_protobufs_repository_lock_proto_depIdxs = []int32{
	0, // 0: protobufs.RepositoryLock.HandleSaveRawSuccess:input_type -> protobufs.ReportSaveRawSuccessRequest
	1, // 1: protobufs.RepositoryLock.HandleSaveRawSuccess:output_type -> protobufs.ReportSaveRawSuccessResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobufs_repository_lock_proto_init() }
func file_protobufs_repository_lock_proto_init() {
	if File_protobufs_repository_lock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobufs_repository_lock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportSaveRawSuccessRequest); i {
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
		file_protobufs_repository_lock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportSaveRawSuccessResponse); i {
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
			RawDescriptor: file_protobufs_repository_lock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobufs_repository_lock_proto_goTypes,
		DependencyIndexes: file_protobufs_repository_lock_proto_depIdxs,
		MessageInfos:      file_protobufs_repository_lock_proto_msgTypes,
	}.Build()
	File_protobufs_repository_lock_proto = out.File
	file_protobufs_repository_lock_proto_rawDesc = nil
	file_protobufs_repository_lock_proto_goTypes = nil
	file_protobufs_repository_lock_proto_depIdxs = nil
}