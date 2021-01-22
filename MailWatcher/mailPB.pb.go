// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.5.1
// source: protoFolder/mailPB.proto

package main

import (
	context "context"
	reflect "reflect"
	sync "sync"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Mails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mails []*MailSimple `protobuf:"bytes,1,rep,name=mails,proto3" json:"mails,omitempty"`
}

func (x *Mails) Reset() {
	*x = Mails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFolder_mailPB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mails) ProtoMessage() {}

func (x *Mails) ProtoReflect() protoreflect.Message {
	mi := &file_protoFolder_mailPB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mails.ProtoReflect.Descriptor instead.
func (*Mails) Descriptor() ([]byte, []int) {
	return file_protoFolder_mailPB_proto_rawDescGZIP(), []int{0}
}

func (x *Mails) GetMails() []*MailSimple {
	if x != nil {
		return x.Mails
	}
	return nil
}

type MailSimple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExchangeID       string `protobuf:"bytes,1,opt,name=exchangeID,proto3" json:"exchangeID,omitempty"`
	ReceivedDateTime string `protobuf:"bytes,2,opt,name=receivedDateTime,proto3" json:"receivedDateTime,omitempty"`
	Subject          string `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	BodyPreview      string `protobuf:"bytes,4,opt,name=bodyPreview,proto3" json:"bodyPreview,omitempty"`
	WebLink          string `protobuf:"bytes,5,opt,name=webLink,proto3" json:"webLink,omitempty"`
	BodyContentType  string `protobuf:"bytes,6,opt,name=body_contentType,json=bodyContentType,proto3" json:"body_contentType,omitempty"`
	BodyContent      string `protobuf:"bytes,7,opt,name=body_content,json=bodyContent,proto3" json:"body_content,omitempty"`
	FromAddress      string `protobuf:"bytes,8,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	ParentFolderId   string `protobuf:"bytes,9,opt,name=parentFolderId,proto3" json:"parentFolderId,omitempty"`
	TaskName         string `protobuf:"bytes,10,opt,name=taskName,proto3" json:"taskName,omitempty"`
}

func (x *MailSimple) Reset() {
	*x = MailSimple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFolder_mailPB_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailSimple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailSimple) ProtoMessage() {}

func (x *MailSimple) ProtoReflect() protoreflect.Message {
	mi := &file_protoFolder_mailPB_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailSimple.ProtoReflect.Descriptor instead.
func (*MailSimple) Descriptor() ([]byte, []int) {
	return file_protoFolder_mailPB_proto_rawDescGZIP(), []int{1}
}

func (x *MailSimple) GetExchangeID() string {
	if x != nil {
		return x.ExchangeID
	}
	return ""
}

func (x *MailSimple) GetReceivedDateTime() string {
	if x != nil {
		return x.ReceivedDateTime
	}
	return ""
}

func (x *MailSimple) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *MailSimple) GetBodyPreview() string {
	if x != nil {
		return x.BodyPreview
	}
	return ""
}

func (x *MailSimple) GetWebLink() string {
	if x != nil {
		return x.WebLink
	}
	return ""
}

func (x *MailSimple) GetBodyContentType() string {
	if x != nil {
		return x.BodyContentType
	}
	return ""
}

func (x *MailSimple) GetBodyContent() string {
	if x != nil {
		return x.BodyContent
	}
	return ""
}

func (x *MailSimple) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *MailSimple) GetParentFolderId() string {
	if x != nil {
		return x.ParentFolderId
	}
	return ""
}

func (x *MailSimple) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

type MailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mails *Mails `protobuf:"bytes,1,opt,name=mails,proto3" json:"mails,omitempty"`
}

func (x *MailsRequest) Reset() {
	*x = MailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFolder_mailPB_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailsRequest) ProtoMessage() {}

func (x *MailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoFolder_mailPB_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailsRequest.ProtoReflect.Descriptor instead.
func (*MailsRequest) Descriptor() ([]byte, []int) {
	return file_protoFolder_mailPB_proto_rawDescGZIP(), []int{2}
}

func (x *MailsRequest) GetMails() *Mails {
	if x != nil {
		return x.Mails
	}
	return nil
}

type MailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MailsResponse) Reset() {
	*x = MailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFolder_mailPB_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailsResponse) ProtoMessage() {}

func (x *MailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protoFolder_mailPB_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailsResponse.ProtoReflect.Descriptor instead.
func (*MailsResponse) Descriptor() ([]byte, []int) {
	return file_protoFolder_mailPB_proto_rawDescGZIP(), []int{3}
}

func (x *MailsResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_protoFolder_mailPB_proto protoreflect.FileDescriptor

var file_protoFolder_mailPB_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x6d, 0x61,
	0x69, 0x6c, 0x50, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x61, 0x69, 0x6c,
	0x50, 0x42, 0x22, 0x31, 0x0a, 0x05, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x6d,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x69,
	0x6c, 0x50, 0x42, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x05,
	0x6d, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xe3, 0x02, 0x0a, 0x0a, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6f,
	0x64, 0x79, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x62, 0x6f, 0x64, 0x79, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x18, 0x0a, 0x07,
	0x77, 0x65, 0x62, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77,
	0x65, 0x62, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x29, 0x0a, 0x10, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x62, 0x6f, 0x64, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6f, 0x64, 0x79, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x0c, 0x4d,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x6d,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x61, 0x69,
	0x6c, 0x50, 0x42, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x6d, 0x61, 0x69, 0x6c, 0x73,
	0x22, 0x27, 0x0a, 0x0d, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x48, 0x0a, 0x0c, 0x4d, 0x61, 0x69,
	0x6c, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x50, 0x6f, 0x73,
	0x74, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x42, 0x2e,
	0x4d, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d,
	0x61, 0x69, 0x6c, 0x50, 0x42, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x42, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoFolder_mailPB_proto_rawDescOnce sync.Once
	file_protoFolder_mailPB_proto_rawDescData = file_protoFolder_mailPB_proto_rawDesc
)

func file_protoFolder_mailPB_proto_rawDescGZIP() []byte {
	file_protoFolder_mailPB_proto_rawDescOnce.Do(func() {
		file_protoFolder_mailPB_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoFolder_mailPB_proto_rawDescData)
	})
	return file_protoFolder_mailPB_proto_rawDescData
}

var file_protoFolder_mailPB_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protoFolder_mailPB_proto_goTypes = []interface{}{
	(*Mails)(nil),         // 0: mailPB.Mails
	(*MailSimple)(nil),    // 1: mailPB.MailSimple
	(*MailsRequest)(nil),  // 2: mailPB.MailsRequest
	(*MailsResponse)(nil), // 3: mailPB.MailsResponse
}
var file_protoFolder_mailPB_proto_depIdxs = []int32{
	1, // 0: mailPB.Mails.mails:type_name -> mailPB.MailSimple
	0, // 1: mailPB.MailsRequest.mails:type_name -> mailPB.Mails
	2, // 2: mailPB.MailsService.PostMails:input_type -> mailPB.MailsRequest
	3, // 3: mailPB.MailsService.PostMails:output_type -> mailPB.MailsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protoFolder_mailPB_proto_init() }
func file_protoFolder_mailPB_proto_init() {
	if File_protoFolder_mailPB_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoFolder_mailPB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mails); i {
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
		file_protoFolder_mailPB_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailSimple); i {
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
		file_protoFolder_mailPB_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailsRequest); i {
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
		file_protoFolder_mailPB_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailsResponse); i {
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
			RawDescriptor: file_protoFolder_mailPB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protoFolder_mailPB_proto_goTypes,
		DependencyIndexes: file_protoFolder_mailPB_proto_depIdxs,
		MessageInfos:      file_protoFolder_mailPB_proto_msgTypes,
	}.Build()
	File_protoFolder_mailPB_proto = out.File
	file_protoFolder_mailPB_proto_rawDesc = nil
	file_protoFolder_mailPB_proto_goTypes = nil
	file_protoFolder_mailPB_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MailsServiceClient is the client API for MailsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MailsServiceClient interface {
	//Unary
	PostMails(ctx context.Context, in *MailsRequest, opts ...grpc.CallOption) (*MailsResponse, error)
}

type mailsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMailsServiceClient(cc grpc.ClientConnInterface) MailsServiceClient {
	return &mailsServiceClient{cc}
}

func (c *mailsServiceClient) PostMails(ctx context.Context, in *MailsRequest, opts ...grpc.CallOption) (*MailsResponse, error) {
	out := new(MailsResponse)
	err := c.cc.Invoke(ctx, "/mailPB.MailsService/PostMails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailsServiceServer is the server API for MailsService service.
type MailsServiceServer interface {
	//Unary
	PostMails(context.Context, *MailsRequest) (*MailsResponse, error)
}

// UnimplementedMailsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMailsServiceServer struct {
}

func (*UnimplementedMailsServiceServer) PostMails(context.Context, *MailsRequest) (*MailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostMails not implemented")
}

func RegisterMailsServiceServer(s *grpc.Server, srv MailsServiceServer) {
	s.RegisterService(&_MailsService_serviceDesc, srv)
}

func _MailsService_PostMails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailsServiceServer).PostMails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mailPB.MailsService/PostMails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailsServiceServer).PostMails(ctx, req.(*MailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MailsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mailPB.MailsService",
	HandlerType: (*MailsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostMails",
			Handler:    _MailsService_PostMails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoFolder/mailPB.proto",
}
