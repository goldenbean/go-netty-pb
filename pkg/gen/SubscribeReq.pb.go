// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: src/main/proto/SubscribeReq.proto

package gen

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

type SubscribeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubReqID    int32    `protobuf:"varint,1,opt,name=subReqID,proto3" json:"subReqID,omitempty"`
	UserName    string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	ProductName string   `protobuf:"bytes,3,opt,name=productName,proto3" json:"productName,omitempty"`
	Address     []string `protobuf:"bytes,4,rep,name=address,proto3" json:"address,omitempty"`
}

func (x *SubscribeReq) Reset() {
	*x = SubscribeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_main_proto_SubscribeReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeReq) ProtoMessage() {}

func (x *SubscribeReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_main_proto_SubscribeReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeReq.ProtoReflect.Descriptor instead.
func (*SubscribeReq) Descriptor() ([]byte, []int) {
	return file_src_main_proto_SubscribeReq_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeReq) GetSubReqID() int32 {
	if x != nil {
		return x.SubReqID
	}
	return 0
}

func (x *SubscribeReq) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *SubscribeReq) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *SubscribeReq) GetAddress() []string {
	if x != nil {
		return x.Address
	}
	return nil
}

var File_src_main_proto_SubscribeReq_proto protoreflect.FileDescriptor

var file_src_main_proto_SubscribeReq_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e, 0x65, 0x74, 0x74, 0x79, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x75, 0x62, 0x52, 0x65, 0x71, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73,
	0x75, 0x62, 0x52, 0x65, 0x71, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42,
	0x3a, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x70, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x74, 0x79,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x42,
	0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5a, 0x07, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_src_main_proto_SubscribeReq_proto_rawDescOnce sync.Once
	file_src_main_proto_SubscribeReq_proto_rawDescData = file_src_main_proto_SubscribeReq_proto_rawDesc
)

func file_src_main_proto_SubscribeReq_proto_rawDescGZIP() []byte {
	file_src_main_proto_SubscribeReq_proto_rawDescOnce.Do(func() {
		file_src_main_proto_SubscribeReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_main_proto_SubscribeReq_proto_rawDescData)
	})
	return file_src_main_proto_SubscribeReq_proto_rawDescData
}

var file_src_main_proto_SubscribeReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_src_main_proto_SubscribeReq_proto_goTypes = []interface{}{
	(*SubscribeReq)(nil), // 0: netty.SubscribeReq
}
var file_src_main_proto_SubscribeReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_main_proto_SubscribeReq_proto_init() }
func file_src_main_proto_SubscribeReq_proto_init() {
	if File_src_main_proto_SubscribeReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_main_proto_SubscribeReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeReq); i {
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
			RawDescriptor: file_src_main_proto_SubscribeReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_main_proto_SubscribeReq_proto_goTypes,
		DependencyIndexes: file_src_main_proto_SubscribeReq_proto_depIdxs,
		MessageInfos:      file_src_main_proto_SubscribeReq_proto_msgTypes,
	}.Build()
	File_src_main_proto_SubscribeReq_proto = out.File
	file_src_main_proto_SubscribeReq_proto_rawDesc = nil
	file_src_main_proto_SubscribeReq_proto_goTypes = nil
	file_src_main_proto_SubscribeReq_proto_depIdxs = nil
}
