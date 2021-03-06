// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: error.proto

package Derror

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

type Error int32

const (
	Error_ok            Error = 0
	Error_badRequest    Error = 1    //错误的消息
	Error_noHandler     Error = 2    //没有命令的处理方法
	Error_noAuthority   Error = 3    //权限不足
	Error_net           Error = 4    //网络问题
	Error_server        Error = 5    //服务器错误
	Error_authExpire    Error = 6    //鉴权信息过期
	Error_noAccount     Error = 7    //用户不存在
	Error_noturn        Error = 1000 //不是玩家的操作阶段
	Error_noInHand      Error = 1001 //token不在手牌
	Error_noInDropTop   Error = 1002 //不在弃牌堆顶部
	Error_invaildInvest Error = 1003 //非法投资
	Error_inMatching    Error = 1004 //玩家已在匹配中
)

// Enum value maps for Error.
var (
	Error_name = map[int32]string{
		0:    "ok",
		1:    "badRequest",
		2:    "noHandler",
		3:    "noAuthority",
		4:    "net",
		5:    "server",
		6:    "authExpire",
		7:    "noAccount",
		1000: "noturn",
		1001: "noInHand",
		1002: "noInDropTop",
		1003: "invaildInvest",
		1004: "inMatching",
	}
	Error_value = map[string]int32{
		"ok":            0,
		"badRequest":    1,
		"noHandler":     2,
		"noAuthority":   3,
		"net":           4,
		"server":        5,
		"authExpire":    6,
		"noAccount":     7,
		"noturn":        1000,
		"noInHand":      1001,
		"noInDropTop":   1002,
		"invaildInvest": 1003,
		"inMatching":    1004,
	}
)

func (x Error) Enum() *Error {
	p := new(Error)
	*p = x
	return p
}

func (x Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error) Descriptor() protoreflect.EnumDescriptor {
	return file_error_proto_enumTypes[0].Descriptor()
}

func (Error) Type() protoreflect.EnumType {
	return &file_error_proto_enumTypes[0]
}

func (x Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error.Descriptor instead.
func (Error) EnumDescriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{0}
}

var File_error_proto protoreflect.FileDescriptor

var file_error_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x44, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0xc6, 0x01, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x06, 0x0a, 0x02, 0x6f, 0x6b, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x62, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x6e, 0x6f, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x6e,
	0x6f, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03,
	0x6e, 0x65, 0x74, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x10,
	0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x10,
	0x06, 0x12, 0x0d, 0x0a, 0x09, 0x6e, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x07,
	0x12, 0x0b, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x75, 0x72, 0x6e, 0x10, 0xe8, 0x07, 0x12, 0x0d, 0x0a,
	0x08, 0x6e, 0x6f, 0x49, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x10, 0xe9, 0x07, 0x12, 0x10, 0x0a, 0x0b,
	0x6e, 0x6f, 0x49, 0x6e, 0x44, 0x72, 0x6f, 0x70, 0x54, 0x6f, 0x70, 0x10, 0xea, 0x07, 0x12, 0x12,
	0x0a, 0x0d, 0x69, 0x6e, 0x76, 0x61, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x10,
	0xeb, 0x07, 0x12, 0x0f, 0x0a, 0x0a, 0x69, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67,
	0x10, 0xec, 0x07, 0x42, 0x13, 0x5a, 0x11, 0x4a, 0x46, 0x46, 0x75, 0x6e, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x44, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_error_proto_rawDescOnce sync.Once
	file_error_proto_rawDescData = file_error_proto_rawDesc
)

func file_error_proto_rawDescGZIP() []byte {
	file_error_proto_rawDescOnce.Do(func() {
		file_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_proto_rawDescData)
	})
	return file_error_proto_rawDescData
}

var file_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_error_proto_goTypes = []interface{}{
	(Error)(0), // 0: data.Derror.error
}
var file_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_error_proto_init() }
func file_error_proto_init() {
	if File_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_proto_goTypes,
		DependencyIndexes: file_error_proto_depIdxs,
		EnumInfos:         file_error_proto_enumTypes,
	}.Build()
	File_error_proto = out.File
	file_error_proto_rawDesc = nil
	file_error_proto_goTypes = nil
	file_error_proto_depIdxs = nil
}
