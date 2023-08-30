// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: dnspb/dns.proto

package dnspb

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

type DNSMessageType int32

const (
	DNSMessageType_NOP               DNSMessageType = 0 // aka FINGERPRINT
	DNSMessageType_TOTP              DNSMessageType = 1
	DNSMessageType_INIT              DNSMessageType = 2
	DNSMessageType_POLL              DNSMessageType = 3
	DNSMessageType_CLOSE             DNSMessageType = 4
	DNSMessageType_MANIFEST          DNSMessageType = 6
	DNSMessageType_DATA_TO_IMPLANT   DNSMessageType = 7
	DNSMessageType_DATA_FROM_IMPLANT DNSMessageType = 8
	DNSMessageType_CLEAR             DNSMessageType = 9
)

// Enum value maps for DNSMessageType.
var (
	DNSMessageType_name = map[int32]string{
		0: "NOP",
		1: "TOTP",
		2: "INIT",
		3: "POLL",
		4: "CLOSE",
		6: "MANIFEST",
		7: "DATA_TO_IMPLANT",
		8: "DATA_FROM_IMPLANT",
		9: "CLEAR",
	}
	DNSMessageType_value = map[string]int32{
		"NOP":               0,
		"TOTP":              1,
		"INIT":              2,
		"POLL":              3,
		"CLOSE":             4,
		"MANIFEST":          6,
		"DATA_TO_IMPLANT":   7,
		"DATA_FROM_IMPLANT": 8,
		"CLEAR":             9,
	}
)

func (x DNSMessageType) Enum() *DNSMessageType {
	p := new(DNSMessageType)
	*p = x
	return p
}

func (x DNSMessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DNSMessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_dnspb_dns_proto_enumTypes[0].Descriptor()
}

func (DNSMessageType) Type() protoreflect.EnumType {
	return &file_dnspb_dns_proto_enumTypes[0]
}

func (x DNSMessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DNSMessageType.Descriptor instead.
func (DNSMessageType) EnumDescriptor() ([]byte, []int) {
	return file_dnspb_dns_proto_rawDescGZIP(), []int{0}
}

// NOTE: DNS is very space sensitive so certain fields are re-purposed
// depending on the DNSMessageType as noted below:
//
// [Type TOTP]: ID field is used for the TOTP code
type DNSMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  DNSMessageType `protobuf:"varint,1,opt,name=Type,proto3,enum=dnspb.DNSMessageType" json:"Type,omitempty"`
	ID    uint32         `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`       // 8 bit message id + 24 bit dns session ID
	Start uint32         `protobuf:"varint,3,opt,name=Start,proto3" json:"Start,omitempty"` // Bytes start at
	Stop  uint32         `protobuf:"varint,4,opt,name=Stop,proto3" json:"Stop,omitempty"`   // Bytes stop at
	Size  uint32         `protobuf:"varint,5,opt,name=Size,proto3" json:"Size,omitempty"`   // Total size
	Data  []byte         `protobuf:"bytes,6,opt,name=Data,proto3" json:"Data,omitempty"`    // Actual data
}

func (x *DNSMessage) Reset() {
	*x = DNSMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dnspb_dns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSMessage) ProtoMessage() {}

func (x *DNSMessage) ProtoReflect() protoreflect.Message {
	mi := &file_dnspb_dns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSMessage.ProtoReflect.Descriptor instead.
func (*DNSMessage) Descriptor() ([]byte, []int) {
	return file_dnspb_dns_proto_rawDescGZIP(), []int{0}
}

func (x *DNSMessage) GetType() DNSMessageType {
	if x != nil {
		return x.Type
	}
	return DNSMessageType_NOP
}

func (x *DNSMessage) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DNSMessage) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *DNSMessage) GetStop() uint32 {
	if x != nil {
		return x.Stop
	}
	return 0
}

func (x *DNSMessage) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *DNSMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_dnspb_dns_proto protoreflect.FileDescriptor

var file_dnspb_dns_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x6e, 0x73, 0x70, 0x62, 0x2f, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x64, 0x6e, 0x73, 0x70, 0x62, 0x22, 0x99, 0x01, 0x0a, 0x0a, 0x44, 0x4e, 0x53,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x64, 0x6e, 0x73, 0x70, 0x62, 0x2e, 0x44, 0x4e,
	0x53, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x2a, 0x87, 0x01, 0x0a, 0x0e, 0x44, 0x4e, 0x53, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x4f, 0x50, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x54, 0x4f, 0x54, 0x50, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e,
	0x49, 0x54, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x4c, 0x4c, 0x10, 0x03, 0x12, 0x09,
	0x0a, 0x05, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x41, 0x4e,
	0x49, 0x46, 0x45, 0x53, 0x54, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x54, 0x4f, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x41, 0x4e, 0x54, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11,
	0x44, 0x41, 0x54, 0x41, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x41, 0x4e,
	0x54, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x45, 0x41, 0x52, 0x10, 0x09, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x73,
	0x68, 0x6f, 0x70, 0x66, 0x6f, 0x78, 0x2f, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x6e, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dnspb_dns_proto_rawDescOnce sync.Once
	file_dnspb_dns_proto_rawDescData = file_dnspb_dns_proto_rawDesc
)

func file_dnspb_dns_proto_rawDescGZIP() []byte {
	file_dnspb_dns_proto_rawDescOnce.Do(func() {
		file_dnspb_dns_proto_rawDescData = protoimpl.X.CompressGZIP(file_dnspb_dns_proto_rawDescData)
	})
	return file_dnspb_dns_proto_rawDescData
}

var file_dnspb_dns_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dnspb_dns_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_dnspb_dns_proto_goTypes = []interface{}{
	(DNSMessageType)(0), // 0: dnspb.DNSMessageType
	(*DNSMessage)(nil),  // 1: dnspb.DNSMessage
}
var file_dnspb_dns_proto_depIdxs = []int32{
	0, // 0: dnspb.DNSMessage.Type:type_name -> dnspb.DNSMessageType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dnspb_dns_proto_init() }
func file_dnspb_dns_proto_init() {
	if File_dnspb_dns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dnspb_dns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSMessage); i {
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
			RawDescriptor: file_dnspb_dns_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dnspb_dns_proto_goTypes,
		DependencyIndexes: file_dnspb_dns_proto_depIdxs,
		EnumInfos:         file_dnspb_dns_proto_enumTypes,
		MessageInfos:      file_dnspb_dns_proto_msgTypes,
	}.Build()
	File_dnspb_dns_proto = out.File
	file_dnspb_dns_proto_rawDesc = nil
	file_dnspb_dns_proto_goTypes = nil
	file_dnspb_dns_proto_depIdxs = nil
}
