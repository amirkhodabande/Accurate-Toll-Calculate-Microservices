// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/invoicer/invoicer.proto

package invoicer

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

type GetInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OBUID int64 `protobuf:"varint,1,opt,name=OBUID,proto3" json:"OBUID,omitempty"`
}

func (x *GetInvoiceRequest) Reset() {
	*x = GetInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_invoicer_invoicer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvoiceRequest) ProtoMessage() {}

func (x *GetInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_invoicer_invoicer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvoiceRequest.ProtoReflect.Descriptor instead.
func (*GetInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_proto_invoicer_invoicer_proto_rawDescGZIP(), []int{0}
}

func (x *GetInvoiceRequest) GetOBUID() int64 {
	if x != nil {
		return x.OBUID
	}
	return 0
}

type GetInvoiceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OBUID    int64   `protobuf:"varint,1,opt,name=OBUID,proto3" json:"OBUID,omitempty"`
	Distance float64 `protobuf:"fixed64,2,opt,name=Distance,proto3" json:"Distance,omitempty"`
	Amount   float64 `protobuf:"fixed64,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *GetInvoiceReply) Reset() {
	*x = GetInvoiceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_invoicer_invoicer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvoiceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvoiceReply) ProtoMessage() {}

func (x *GetInvoiceReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_invoicer_invoicer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvoiceReply.ProtoReflect.Descriptor instead.
func (*GetInvoiceReply) Descriptor() ([]byte, []int) {
	return file_proto_invoicer_invoicer_proto_rawDescGZIP(), []int{1}
}

func (x *GetInvoiceReply) GetOBUID() int64 {
	if x != nil {
		return x.OBUID
	}
	return 0
}

func (x *GetInvoiceReply) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *GetInvoiceReply) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_proto_invoicer_invoicer_proto protoreflect.FileDescriptor

var file_proto_invoicer_invoicer_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72,
	0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4f,
	0x42, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4f, 0x42, 0x55, 0x49,
	0x44, 0x22, 0x5b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x42, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x4f, 0x42, 0x55, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x44, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x4c,
	0x0a, 0x08, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x42, 0x0a, 0x18,
	0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x01, 0x5a, 0x13, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_invoicer_invoicer_proto_rawDescOnce sync.Once
	file_proto_invoicer_invoicer_proto_rawDescData = file_proto_invoicer_invoicer_proto_rawDesc
)

func file_proto_invoicer_invoicer_proto_rawDescGZIP() []byte {
	file_proto_invoicer_invoicer_proto_rawDescOnce.Do(func() {
		file_proto_invoicer_invoicer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_invoicer_invoicer_proto_rawDescData)
	})
	return file_proto_invoicer_invoicer_proto_rawDescData
}

var file_proto_invoicer_invoicer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_invoicer_invoicer_proto_goTypes = []interface{}{
	(*GetInvoiceRequest)(nil), // 0: proto.GetInvoiceRequest
	(*GetInvoiceReply)(nil),   // 1: proto.GetInvoiceReply
}
var file_proto_invoicer_invoicer_proto_depIdxs = []int32{
	0, // 0: proto.Invoicer.GetInvoice:input_type -> proto.GetInvoiceRequest
	1, // 1: proto.Invoicer.GetInvoice:output_type -> proto.GetInvoiceReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_invoicer_invoicer_proto_init() }
func file_proto_invoicer_invoicer_proto_init() {
	if File_proto_invoicer_invoicer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_invoicer_invoicer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvoiceRequest); i {
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
		file_proto_invoicer_invoicer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvoiceReply); i {
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
			RawDescriptor: file_proto_invoicer_invoicer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_invoicer_invoicer_proto_goTypes,
		DependencyIndexes: file_proto_invoicer_invoicer_proto_depIdxs,
		MessageInfos:      file_proto_invoicer_invoicer_proto_msgTypes,
	}.Build()
	File_proto_invoicer_invoicer_proto = out.File
	file_proto_invoicer_invoicer_proto_rawDesc = nil
	file_proto_invoicer_invoicer_proto_goTypes = nil
	file_proto_invoicer_invoicer_proto_depIdxs = nil
}
