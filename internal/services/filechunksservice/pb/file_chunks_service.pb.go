// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.23.3
// source: file_chunks_service.proto

package pb

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

type FileChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FilePath       string `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	ChunkIndex     int32  `protobuf:"varint,3,opt,name=chunk_index,json=chunkIndex,proto3" json:"chunk_index,omitempty"`
	NumTotalChunks int32  `protobuf:"varint,4,opt,name=num_total_chunks,json=numTotalChunks,proto3" json:"num_total_chunks,omitempty"`
	Content        string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *FileChunk) Reset() {
	*x = FileChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_chunks_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunk) ProtoMessage() {}

func (x *FileChunk) ProtoReflect() protoreflect.Message {
	mi := &file_file_chunks_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunk.ProtoReflect.Descriptor instead.
func (*FileChunk) Descriptor() ([]byte, []int) {
	return file_file_chunks_service_proto_rawDescGZIP(), []int{0}
}

func (x *FileChunk) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FileChunk) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *FileChunk) GetChunkIndex() int32 {
	if x != nil {
		return x.ChunkIndex
	}
	return 0
}

func (x *FileChunk) GetNumTotalChunks() int32 {
	if x != nil {
		return x.NumTotalChunks
	}
	return 0
}

func (x *FileChunk) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SaveFileChunksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileChunks []*FileChunk `protobuf:"bytes,1,rep,name=file_chunks,json=fileChunks,proto3" json:"file_chunks,omitempty"`
}

func (x *SaveFileChunksRequest) Reset() {
	*x = SaveFileChunksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_chunks_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveFileChunksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveFileChunksRequest) ProtoMessage() {}

func (x *SaveFileChunksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_chunks_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveFileChunksRequest.ProtoReflect.Descriptor instead.
func (*SaveFileChunksRequest) Descriptor() ([]byte, []int) {
	return file_file_chunks_service_proto_rawDescGZIP(), []int{1}
}

func (x *SaveFileChunksRequest) GetFileChunks() []*FileChunk {
	if x != nil {
		return x.FileChunks
	}
	return nil
}

type FileChunkSaveStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath         string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	IsLastSavedChunk bool   `protobuf:"varint,2,opt,name=is_last_saved_chunk,json=isLastSavedChunk,proto3" json:"is_last_saved_chunk,omitempty"`
}

func (x *FileChunkSaveStatus) Reset() {
	*x = FileChunkSaveStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_chunks_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileChunkSaveStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunkSaveStatus) ProtoMessage() {}

func (x *FileChunkSaveStatus) ProtoReflect() protoreflect.Message {
	mi := &file_file_chunks_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunkSaveStatus.ProtoReflect.Descriptor instead.
func (*FileChunkSaveStatus) Descriptor() ([]byte, []int) {
	return file_file_chunks_service_proto_rawDescGZIP(), []int{2}
}

func (x *FileChunkSaveStatus) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *FileChunkSaveStatus) GetIsLastSavedChunk() bool {
	if x != nil {
		return x.IsLastSavedChunk
	}
	return false
}

type SaveFileChunksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileChunkStatuses []*FileChunkSaveStatus `protobuf:"bytes,1,rep,name=file_chunk_statuses,json=fileChunkStatuses,proto3" json:"file_chunk_statuses,omitempty"`
}

func (x *SaveFileChunksResponse) Reset() {
	*x = SaveFileChunksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_chunks_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveFileChunksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveFileChunksResponse) ProtoMessage() {}

func (x *SaveFileChunksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_chunks_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveFileChunksResponse.ProtoReflect.Descriptor instead.
func (*SaveFileChunksResponse) Descriptor() ([]byte, []int) {
	return file_file_chunks_service_proto_rawDescGZIP(), []int{3}
}

func (x *SaveFileChunksResponse) GetFileChunkStatuses() []*FileChunkSaveStatus {
	if x != nil {
		return x.FileChunkStatuses
	}
	return nil
}

var File_file_chunks_service_proto protoreflect.FileDescriptor

var file_file_chunks_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x09,
	0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x28, 0x0a, 0x10, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0x44, 0x0a, 0x15, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a,
	0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x0a,
	0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x22, 0x61, 0x0a, 0x13, 0x46, 0x69,
	0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2d,
	0x0a, 0x13, 0x69, 0x73, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x64, 0x5f,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x4c,
	0x61, 0x73, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x5e, 0x0a,
	0x16, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x13, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x11, 0x66, 0x69, 0x6c, 0x65,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x32, 0x56, 0x0a,
	0x11, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x73, 0x12, 0x16, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_chunks_service_proto_rawDescOnce sync.Once
	file_file_chunks_service_proto_rawDescData = file_file_chunks_service_proto_rawDesc
)

func file_file_chunks_service_proto_rawDescGZIP() []byte {
	file_file_chunks_service_proto_rawDescOnce.Do(func() {
		file_file_chunks_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_chunks_service_proto_rawDescData)
	})
	return file_file_chunks_service_proto_rawDescData
}

var file_file_chunks_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_file_chunks_service_proto_goTypes = []interface{}{
	(*FileChunk)(nil),              // 0: FileChunk
	(*SaveFileChunksRequest)(nil),  // 1: SaveFileChunksRequest
	(*FileChunkSaveStatus)(nil),    // 2: FileChunkSaveStatus
	(*SaveFileChunksResponse)(nil), // 3: SaveFileChunksResponse
}
var file_file_chunks_service_proto_depIdxs = []int32{
	0, // 0: SaveFileChunksRequest.file_chunks:type_name -> FileChunk
	2, // 1: SaveFileChunksResponse.file_chunk_statuses:type_name -> FileChunkSaveStatus
	1, // 2: FileChunksService.SaveFileChunks:input_type -> SaveFileChunksRequest
	3, // 3: FileChunksService.SaveFileChunks:output_type -> SaveFileChunksResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_file_chunks_service_proto_init() }
func file_file_chunks_service_proto_init() {
	if File_file_chunks_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_chunks_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileChunk); i {
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
		file_file_chunks_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveFileChunksRequest); i {
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
		file_file_chunks_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileChunkSaveStatus); i {
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
		file_file_chunks_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveFileChunksResponse); i {
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
			RawDescriptor: file_file_chunks_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_chunks_service_proto_goTypes,
		DependencyIndexes: file_file_chunks_service_proto_depIdxs,
		MessageInfos:      file_file_chunks_service_proto_msgTypes,
	}.Build()
	File_file_chunks_service_proto = out.File
	file_file_chunks_service_proto_rawDesc = nil
	file_file_chunks_service_proto_goTypes = nil
	file_file_chunks_service_proto_depIdxs = nil
}
