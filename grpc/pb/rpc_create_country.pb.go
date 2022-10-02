// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: rpc_create_country.proto

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

type CreateCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country *Country `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *CreateCountryRequest) Reset() {
	*x = CreateCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_country_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCountryRequest) ProtoMessage() {}

func (x *CreateCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_country_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCountryRequest.ProtoReflect.Descriptor instead.
func (*CreateCountryRequest) Descriptor() ([]byte, []int) {
	return file_rpc_create_country_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCountryRequest) GetCountry() *Country {
	if x != nil {
		return x.Country
	}
	return nil
}

type CreateCountryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country *Country `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *CreateCountryResponse) Reset() {
	*x = CreateCountryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_country_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCountryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCountryResponse) ProtoMessage() {}

func (x *CreateCountryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_country_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCountryResponse.ProtoReflect.Descriptor instead.
func (*CreateCountryResponse) Descriptor() ([]byte, []int) {
	return file_rpc_create_country_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCountryResponse) GetCountry() *Country {
	if x != nil {
		return x.Country
	}
	return nil
}

var File_rpc_create_country_proto protoreflect.FileDescriptor

var file_rpc_create_country_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x40, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x22, 0x41, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x74, 0x65, 0x73, 0x74, 0x30, 0x31, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_create_country_proto_rawDescOnce sync.Once
	file_rpc_create_country_proto_rawDescData = file_rpc_create_country_proto_rawDesc
)

func file_rpc_create_country_proto_rawDescGZIP() []byte {
	file_rpc_create_country_proto_rawDescOnce.Do(func() {
		file_rpc_create_country_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_create_country_proto_rawDescData)
	})
	return file_rpc_create_country_proto_rawDescData
}

var file_rpc_create_country_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_create_country_proto_goTypes = []interface{}{
	(*CreateCountryRequest)(nil),  // 0: proto.CreateCountryRequest
	(*CreateCountryResponse)(nil), // 1: proto.CreateCountryResponse
	(*Country)(nil),               // 2: proto.Country
}
var file_rpc_create_country_proto_depIdxs = []int32{
	2, // 0: proto.CreateCountryRequest.country:type_name -> proto.Country
	2, // 1: proto.CreateCountryResponse.country:type_name -> proto.Country
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_create_country_proto_init() }
func file_rpc_create_country_proto_init() {
	if File_rpc_create_country_proto != nil {
		return
	}
	file_country_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_create_country_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCountryRequest); i {
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
		file_rpc_create_country_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCountryResponse); i {
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
			RawDescriptor: file_rpc_create_country_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_create_country_proto_goTypes,
		DependencyIndexes: file_rpc_create_country_proto_depIdxs,
		MessageInfos:      file_rpc_create_country_proto_msgTypes,
	}.Build()
	File_rpc_create_country_proto = out.File
	file_rpc_create_country_proto_rawDesc = nil
	file_rpc_create_country_proto_goTypes = nil
	file_rpc_create_country_proto_depIdxs = nil
}
