// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.19.4
// source: storage.proto

package storage

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

type FileUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileData      []byte `protobuf:"bytes,1,opt,name=file_data,json=fileData,proto3" json:"file_data,omitempty"`                // 文件数据
	FileName      string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`                // 文件名
	FileType      string `protobuf:"bytes,3,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`                // 文件MIME类型
	BussinessId   uint64 `protobuf:"varint,4,opt,name=bussiness_id,json=bussinessId,proto3" json:"bussiness_id,omitempty"`      // 业务ID
	BussinessName string `protobuf:"bytes,5,opt,name=bussiness_name,json=bussinessName,proto3" json:"bussiness_name,omitempty"` // 业务名称
}

func (x *FileUploadRequest) Reset() {
	*x = FileUploadRequest{}
	mi := &file_storage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadRequest) ProtoMessage() {}

func (x *FileUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadRequest.ProtoReflect.Descriptor instead.
func (*FileUploadRequest) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{0}
}

func (x *FileUploadRequest) GetFileData() []byte {
	if x != nil {
		return x.FileData
	}
	return nil
}

func (x *FileUploadRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileUploadRequest) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *FileUploadRequest) GetBussinessId() uint64 {
	if x != nil {
		return x.BussinessId
	}
	return 0
}

func (x *FileUploadRequest) GetBussinessName() string {
	if x != nil {
		return x.BussinessName
	}
	return ""
}

type FileUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileUrl string `protobuf:"bytes,1,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"` // 文件URL
	FileId  uint64 `protobuf:"varint,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`   // 文件ID
}

func (x *FileUploadResponse) Reset() {
	*x = FileUploadResponse{}
	mi := &file_storage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadResponse) ProtoMessage() {}

func (x *FileUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadResponse.ProtoReflect.Descriptor instead.
func (*FileUploadResponse) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{1}
}

func (x *FileUploadResponse) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *FileUploadResponse) GetFileId() uint64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

type FileDownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId uint64 `protobuf:"varint,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"` // 文件ID
}

func (x *FileDownloadRequest) Reset() {
	*x = FileDownloadRequest{}
	mi := &file_storage_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileDownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDownloadRequest) ProtoMessage() {}

func (x *FileDownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDownloadRequest.ProtoReflect.Descriptor instead.
func (*FileDownloadRequest) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{2}
}

func (x *FileDownloadRequest) GetFileId() uint64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

type FileDownloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileData []byte `protobuf:"bytes,1,opt,name=file_data,json=fileData,proto3" json:"file_data,omitempty"`  // 文件数据
	FileName string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`  // 文件名
	FileType string `protobuf:"bytes,3,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`  // 文件MIME类型
	FileSize uint64 `protobuf:"varint,4,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"` // 文件大小
}

func (x *FileDownloadResponse) Reset() {
	*x = FileDownloadResponse{}
	mi := &file_storage_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileDownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDownloadResponse) ProtoMessage() {}

func (x *FileDownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDownloadResponse.ProtoReflect.Descriptor instead.
func (*FileDownloadResponse) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{3}
}

func (x *FileDownloadResponse) GetFileData() []byte {
	if x != nil {
		return x.FileData
	}
	return nil
}

func (x *FileDownloadResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileDownloadResponse) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *FileDownloadResponse) GetFileSize() uint64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

type FileDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId uint64 `protobuf:"varint,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"` // 文件ID
}

func (x *FileDeleteRequest) Reset() {
	*x = FileDeleteRequest{}
	mi := &file_storage_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDeleteRequest) ProtoMessage() {}

func (x *FileDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDeleteRequest.ProtoReflect.Descriptor instead.
func (*FileDeleteRequest) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{4}
}

func (x *FileDeleteRequest) GetFileId() uint64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

type FileDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FileDeleteResponse) Reset() {
	*x = FileDeleteResponse{}
	mi := &file_storage_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDeleteResponse) ProtoMessage() {}

func (x *FileDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDeleteResponse.ProtoReflect.Descriptor instead.
func (*FileDeleteResponse) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{5}
}

type FileInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId uint64 `protobuf:"varint,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"` // 文件ID
}

func (x *FileInfoRequest) Reset() {
	*x = FileInfoRequest{}
	mi := &file_storage_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfoRequest) ProtoMessage() {}

func (x *FileInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfoRequest.ProtoReflect.Descriptor instead.
func (*FileInfoRequest) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{6}
}

func (x *FileInfoRequest) GetFileId() uint64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

type FileInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId      uint64 `protobuf:"varint,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`                // 文件ID
	FileKey     string `protobuf:"bytes,2,opt,name=file_key,json=fileKey,proto3" json:"file_key,omitempty"`              // 文件key
	FileUrl     string `protobuf:"bytes,3,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`              // 文件URL
	FileName    string `protobuf:"bytes,4,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`           // 文件名
	FileType    string `protobuf:"bytes,5,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`           // 文件MIME类型
	FileSize    uint64 `protobuf:"varint,6,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`          // 文件大小
	BussinessId uint64 `protobuf:"varint,7,opt,name=bussiness_id,json=bussinessId,proto3" json:"bussiness_id,omitempty"` // 业务ID
	CreatedAt   int64  `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`       // 创建时间
	UpdatedAt   int64  `protobuf:"varint,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`       // 更新时间
	DeletedAt   int64  `protobuf:"varint,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`      // 删除时间
}

func (x *FileInfoResponse) Reset() {
	*x = FileInfoResponse{}
	mi := &file_storage_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfoResponse) ProtoMessage() {}

func (x *FileInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfoResponse.ProtoReflect.Descriptor instead.
func (*FileInfoResponse) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{7}
}

func (x *FileInfoResponse) GetFileId() uint64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

func (x *FileInfoResponse) GetFileKey() string {
	if x != nil {
		return x.FileKey
	}
	return ""
}

func (x *FileInfoResponse) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *FileInfoResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileInfoResponse) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *FileInfoResponse) GetFileSize() uint64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *FileInfoResponse) GetBussinessId() uint64 {
	if x != nil {
		return x.BussinessId
	}
	return 0
}

func (x *FileInfoResponse) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *FileInfoResponse) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *FileInfoResponse) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type GetBussinessFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BussinessId   uint64 `protobuf:"varint,1,opt,name=bussiness_id,json=bussinessId,proto3" json:"bussiness_id,omitempty"`      // 业务ID
	BussinessName string `protobuf:"bytes,2,opt,name=bussiness_name,json=bussinessName,proto3" json:"bussiness_name,omitempty"` // 业务名称
}

func (x *GetBussinessFilesRequest) Reset() {
	*x = GetBussinessFilesRequest{}
	mi := &file_storage_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBussinessFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBussinessFilesRequest) ProtoMessage() {}

func (x *GetBussinessFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBussinessFilesRequest.ProtoReflect.Descriptor instead.
func (*GetBussinessFilesRequest) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{8}
}

func (x *GetBussinessFilesRequest) GetBussinessId() uint64 {
	if x != nil {
		return x.BussinessId
	}
	return 0
}

func (x *GetBussinessFilesRequest) GetBussinessName() string {
	if x != nil {
		return x.BussinessName
	}
	return ""
}

type GetBussinessFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*FileInfoResponse `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"` // 文件列表
}

func (x *GetBussinessFilesResponse) Reset() {
	*x = GetBussinessFilesResponse{}
	mi := &file_storage_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBussinessFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBussinessFilesResponse) ProtoMessage() {}

func (x *GetBussinessFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBussinessFilesResponse.ProtoReflect.Descriptor instead.
func (*GetBussinessFilesResponse) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{9}
}

func (x *GetBussinessFilesResponse) GetFiles() []*FileInfoResponse {
	if x != nil {
		return x.Files
	}
	return nil
}

type DeleteBusinessFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BussnessId   uint64 `protobuf:"varint,1,opt,name=bussness_id,json=bussnessId,proto3" json:"bussness_id,omitempty"`
	BusinessName string `protobuf:"bytes,2,opt,name=business_name,json=businessName,proto3" json:"business_name,omitempty"`
}

func (x *DeleteBusinessFilesRequest) Reset() {
	*x = DeleteBusinessFilesRequest{}
	mi := &file_storage_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBusinessFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBusinessFilesRequest) ProtoMessage() {}

func (x *DeleteBusinessFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBusinessFilesRequest.ProtoReflect.Descriptor instead.
func (*DeleteBusinessFilesRequest) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteBusinessFilesRequest) GetBussnessId() uint64 {
	if x != nil {
		return x.BussnessId
	}
	return 0
}

func (x *DeleteBusinessFilesRequest) GetBusinessName() string {
	if x != nil {
		return x.BusinessName
	}
	return ""
}

var File_storage_proto protoreflect.FileDescriptor

var file_storage_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x11, 0x46, 0x69, 0x6c,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x75, 0x73,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x75, 0x73, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x62, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x48, 0x0a, 0x12, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x13, 0x46, 0x69, 0x6c,
	0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x14, 0x46, 0x69,
	0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x2c, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x0a, 0x0f, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0xb8, 0x02, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x64, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x62, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x62, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x75, 0x73, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x42, 0x75,
	0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x62, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x73, 0x6e, 0x65, 0x73, 0x73, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x73, 0x6e, 0x65,
	0x73, 0x73, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xda, 0x03, 0x0a, 0x07, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x1a, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0c,
	0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x46, 0x69, 0x6c,
	0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3f, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a,
	0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_proto_rawDescOnce sync.Once
	file_storage_proto_rawDescData = file_storage_proto_rawDesc
)

func file_storage_proto_rawDescGZIP() []byte {
	file_storage_proto_rawDescOnce.Do(func() {
		file_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_proto_rawDescData)
	})
	return file_storage_proto_rawDescData
}

var file_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_storage_proto_goTypes = []any{
	(*FileUploadRequest)(nil),          // 0: storage.FileUploadRequest
	(*FileUploadResponse)(nil),         // 1: storage.FileUploadResponse
	(*FileDownloadRequest)(nil),        // 2: storage.FileDownloadRequest
	(*FileDownloadResponse)(nil),       // 3: storage.FileDownloadResponse
	(*FileDeleteRequest)(nil),          // 4: storage.FileDeleteRequest
	(*FileDeleteResponse)(nil),         // 5: storage.FileDeleteResponse
	(*FileInfoRequest)(nil),            // 6: storage.FileInfoRequest
	(*FileInfoResponse)(nil),           // 7: storage.FileInfoResponse
	(*GetBussinessFilesRequest)(nil),   // 8: storage.GetBussinessFilesRequest
	(*GetBussinessFilesResponse)(nil),  // 9: storage.GetBussinessFilesResponse
	(*DeleteBusinessFilesRequest)(nil), // 10: storage.DeleteBusinessFilesRequest
}
var file_storage_proto_depIdxs = []int32{
	7,  // 0: storage.GetBussinessFilesResponse.files:type_name -> storage.FileInfoResponse
	0,  // 1: storage.Storage.FileUpload:input_type -> storage.FileUploadRequest
	2,  // 2: storage.Storage.FileDownload:input_type -> storage.FileDownloadRequest
	4,  // 3: storage.Storage.FileDelete:input_type -> storage.FileDeleteRequest
	6,  // 4: storage.Storage.FileInfo:input_type -> storage.FileInfoRequest
	8,  // 5: storage.Storage.GetBussinessFiles:input_type -> storage.GetBussinessFilesRequest
	10, // 6: storage.Storage.DeleteBusinessFiles:input_type -> storage.DeleteBusinessFilesRequest
	1,  // 7: storage.Storage.FileUpload:output_type -> storage.FileUploadResponse
	3,  // 8: storage.Storage.FileDownload:output_type -> storage.FileDownloadResponse
	5,  // 9: storage.Storage.FileDelete:output_type -> storage.FileDeleteResponse
	7,  // 10: storage.Storage.FileInfo:output_type -> storage.FileInfoResponse
	9,  // 11: storage.Storage.GetBussinessFiles:output_type -> storage.GetBussinessFilesResponse
	5,  // 12: storage.Storage.DeleteBusinessFiles:output_type -> storage.FileDeleteResponse
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_storage_proto_init() }
func file_storage_proto_init() {
	if File_storage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storage_proto_goTypes,
		DependencyIndexes: file_storage_proto_depIdxs,
		MessageInfos:      file_storage_proto_msgTypes,
	}.Build()
	File_storage_proto = out.File
	file_storage_proto_rawDesc = nil
	file_storage_proto_goTypes = nil
	file_storage_proto_depIdxs = nil
}
