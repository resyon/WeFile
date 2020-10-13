// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: service.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EncodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Err   *Error `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *EncodeResp) Reset() {
	*x = EncodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncodeResp) ProtoMessage() {}

func (x *EncodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncodeResp.ProtoReflect.Descriptor instead.
func (*EncodeResp) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *EncodeResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *EncodeResp) GetErr() *Error {
	if x != nil {
		return x.Err
	}
	return nil
}

type DecodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *DecodeReq) Reset() {
	*x = DecodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodeReq) ProtoMessage() {}

func (x *DecodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodeReq.ProtoReflect.Descriptor instead.
func (*DecodeReq) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *DecodeReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DownloadFileMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileID   int64  `protobuf:"varint,1,opt,name=fileID,proto3" json:"fileID,omitempty"`    // 文件 id
	FileName string `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"` // 用户文件名
}

func (x *DownloadFileMeta) Reset() {
	*x = DownloadFileMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFileMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileMeta) ProtoMessage() {}

func (x *DownloadFileMeta) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileMeta.ProtoReflect.Descriptor instead.
func (*DownloadFileMeta) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *DownloadFileMeta) GetFileID() int64 {
	if x != nil {
		return x.FileID
	}
	return 0
}

func (x *DownloadFileMeta) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type DownloadJWTDecodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileMeta *DownloadFileMeta `protobuf:"bytes,1,opt,name=fileMeta,proto3" json:"fileMeta,omitempty"`
	Err      *Error            `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *DownloadJWTDecodeResp) Reset() {
	*x = DownloadJWTDecodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadJWTDecodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadJWTDecodeResp) ProtoMessage() {}

func (x *DownloadJWTDecodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadJWTDecodeResp.ProtoReflect.Descriptor instead.
func (*DownloadJWTDecodeResp) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *DownloadJWTDecodeResp) GetFileMeta() *DownloadFileMeta {
	if x != nil {
		return x.FileMeta
	}
	return nil
}

func (x *DownloadJWTDecodeResp) GetErr() *Error {
	if x != nil {
		return x.Err
	}
	return nil
}

// 用作插入数据库
type UploadFileMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    int64  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Directory string `protobuf:"bytes,2,opt,name=directory,proto3" json:"directory,omitempty"`
	FileName  string `protobuf:"bytes,3,opt,name=fileName,proto3" json:"fileName,omitempty"`
}

func (x *UploadFileMeta) Reset() {
	*x = UploadFileMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileMeta) ProtoMessage() {}

func (x *UploadFileMeta) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileMeta.ProtoReflect.Descriptor instead.
func (*UploadFileMeta) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *UploadFileMeta) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UploadFileMeta) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

func (x *UploadFileMeta) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type UploadJWTDecodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileMeta *UploadFileMeta `protobuf:"bytes,1,opt,name=fileMeta,proto3" json:"fileMeta,omitempty"`
	Err      *Error          `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *UploadJWTDecodeResp) Reset() {
	*x = UploadJWTDecodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadJWTDecodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadJWTDecodeResp) ProtoMessage() {}

func (x *UploadJWTDecodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadJWTDecodeResp.ProtoReflect.Descriptor instead.
func (*UploadJWTDecodeResp) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *UploadJWTDecodeResp) GetFileMeta() *UploadFileMeta {
	if x != nil {
		return x.FileMeta
	}
	return nil
}

func (x *UploadJWTDecodeResp) GetErr() *Error {
	if x != nil {
		return x.Err
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x52, 0x0a,
	0x0a, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x2e, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x03, 0x65, 0x72,
	0x72, 0x22, 0x21, 0x0a, 0x09, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x46, 0x0a, 0x10, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x8c, 0x01, 0x0a,
	0x15, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4a, 0x57, 0x54, 0x44, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x43, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x03, 0x65,
	0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x62, 0x0a, 0x0e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x88, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4a, 0x57, 0x54, 0x44, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x03, 0x65, 0x72,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x32, 0x95, 0x03, 0x0a, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x11, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x4a, 0x57, 0x54, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x1a, 0x21, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x11, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4a, 0x57, 0x54, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20,
	0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x2c, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x4a, 0x57, 0x54, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x5d, 0x0a, 0x0f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4a, 0x57, 0x54, 0x45, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x25, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x21, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x61, 0x0a, 0x0f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4a, 0x57, 0x54, 0x44, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x20, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x2a, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x4a, 0x57, 0x54, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x54, 0x65, 0x6e, 0x73, 0x53, 0x68, 0x69, 0x6e, 0x65, 0x74, 0x2f, 0x57, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_service_proto_goTypes = []interface{}{
	(*Error)(nil),                 // 0: go.micro.service.auth.Error
	(*EncodeResp)(nil),            // 1: go.micro.service.auth.EncodeResp
	(*DecodeReq)(nil),             // 2: go.micro.service.auth.DecodeReq
	(*DownloadFileMeta)(nil),      // 3: go.micro.service.auth.DownloadFileMeta
	(*DownloadJWTDecodeResp)(nil), // 4: go.micro.service.auth.DownloadJWTDecodeResp
	(*UploadFileMeta)(nil),        // 5: go.micro.service.auth.UploadFileMeta
	(*UploadJWTDecodeResp)(nil),   // 6: go.micro.service.auth.UploadJWTDecodeResp
}
var file_service_proto_depIdxs = []int32{
	0, // 0: go.micro.service.auth.EncodeResp.err:type_name -> go.micro.service.auth.Error
	3, // 1: go.micro.service.auth.DownloadJWTDecodeResp.fileMeta:type_name -> go.micro.service.auth.DownloadFileMeta
	0, // 2: go.micro.service.auth.DownloadJWTDecodeResp.err:type_name -> go.micro.service.auth.Error
	5, // 3: go.micro.service.auth.UploadJWTDecodeResp.fileMeta:type_name -> go.micro.service.auth.UploadFileMeta
	0, // 4: go.micro.service.auth.UploadJWTDecodeResp.err:type_name -> go.micro.service.auth.Error
	3, // 5: go.micro.service.auth.Service.DownloadJWTEncode:input_type -> go.micro.service.auth.DownloadFileMeta
	2, // 6: go.micro.service.auth.Service.DownloadJWTDecode:input_type -> go.micro.service.auth.DecodeReq
	5, // 7: go.micro.service.auth.Service.UploadJWTEncode:input_type -> go.micro.service.auth.UploadFileMeta
	2, // 8: go.micro.service.auth.Service.UploadJWTDecode:input_type -> go.micro.service.auth.DecodeReq
	1, // 9: go.micro.service.auth.Service.DownloadJWTEncode:output_type -> go.micro.service.auth.EncodeResp
	4, // 10: go.micro.service.auth.Service.DownloadJWTDecode:output_type -> go.micro.service.auth.DownloadJWTDecodeResp
	1, // 11: go.micro.service.auth.Service.UploadJWTEncode:output_type -> go.micro.service.auth.EncodeResp
	6, // 12: go.micro.service.auth.Service.UploadJWTDecode:output_type -> go.micro.service.auth.UploadJWTDecodeResp
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncodeResp); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodeReq); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFileMeta); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadJWTDecodeResp); i {
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
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileMeta); i {
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
		file_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadJWTDecodeResp); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
