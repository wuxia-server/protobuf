// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: message/error.proto

package message

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

type ErrorCode int32

const (
	// 错误码从10000开始 预分配 错误码 通用错误码 10000 - 20000 成功为0
	ErrorCode_Success              ErrorCode = 0     // 成功
	ErrorCode_SessionIsNil         ErrorCode = 10000 // session为niL 可能用户未登录
	ErrorCode_OperationUserIdError ErrorCode = 10001 // 该uid不合法或者未注册
	ErrorCode_LoadUserInfoByProxy  ErrorCode = 10002 // 加载用户属性出现问题
	// 增加用户资产错误
	ErrorCode_AddUserItemErr       ErrorCode = 10003 // 增肌道具错误
	ErrorCode_ConsumeItemErr       ErrorCode = 10004 // 消耗道具错误
	ErrorCode_TheUidNotMatch       ErrorCode = 10005 // session中的uid不匹配前端的传递的uid
	ErrorCode_NotEnoughItemDiamond ErrorCode = 10006 // 用户钻石不够
	// 任务错误码从 20000 - 20100
	ErrorCode_UpdateTaskCenterError ErrorCode = 20001 // 更新任务失败
	ErrorCode_NoFoundThisTask       ErrorCode = 20002 // 未找到该任务
	ErrorCode_TaskHasCollect        ErrorCode = 20003 // 任务已领取
	ErrorCode_ResetToMany           ErrorCode = 20004 // 任务已领取
	// 登录注册错误码从 30000开始
	ErrorCode_ParseJwtTokenError      ErrorCode = 30000 // 解析sdk的jwttoken失败
	ErrorCode_UerIdIsZero             ErrorCode = 30001 // 登录传递的uid为0
	ErrorCode_ParseJwtTokenMatchError ErrorCode = 30002 // sdkjwttoken解析的uid 和 传递的uid不一致
	ErrorCode_NoFoundThisUser         ErrorCode = 30003 // 未找到该用户
	ErrorCode_CreateUserError         ErrorCode = 30004 // 创建用户失败
	ErrorCode_OperationDbError        ErrorCode = 30005 // 操作数据库失败
	ErrorCode_UidBoundError           ErrorCode = 30006 // 连接已绑定过UID
	// 商城错误码从40000开始，包括购买和消耗
	ErrorCode_LoadProductError    ErrorCode = 40000
	ErrorCode_ThisProductNotExist ErrorCode = 40001
	// 属性相关错误码
	ErrorCode_NoFoundThisProp      ErrorCode = 50000
	ErrorCode_NoFoundThisItem      ErrorCode = 60000
	ErrorCode_LoadAssetConfigError ErrorCode = 60001
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:     "Success",
		10000: "SessionIsNil",
		10001: "OperationUserIdError",
		10002: "LoadUserInfoByProxy",
		10003: "AddUserItemErr",
		10004: "ConsumeItemErr",
		10005: "TheUidNotMatch",
		10006: "NotEnoughItemDiamond",
		20001: "UpdateTaskCenterError",
		20002: "NoFoundThisTask",
		20003: "TaskHasCollect",
		20004: "ResetToMany",
		30000: "ParseJwtTokenError",
		30001: "UerIdIsZero",
		30002: "ParseJwtTokenMatchError",
		30003: "NoFoundThisUser",
		30004: "CreateUserError",
		30005: "OperationDbError",
		30006: "UidBoundError",
		40000: "LoadProductError",
		40001: "ThisProductNotExist",
		50000: "NoFoundThisProp",
		60000: "NoFoundThisItem",
		60001: "LoadAssetConfigError",
	}
	ErrorCode_value = map[string]int32{
		"Success":                 0,
		"SessionIsNil":            10000,
		"OperationUserIdError":    10001,
		"LoadUserInfoByProxy":     10002,
		"AddUserItemErr":          10003,
		"ConsumeItemErr":          10004,
		"TheUidNotMatch":          10005,
		"NotEnoughItemDiamond":    10006,
		"UpdateTaskCenterError":   20001,
		"NoFoundThisTask":         20002,
		"TaskHasCollect":          20003,
		"ResetToMany":             20004,
		"ParseJwtTokenError":      30000,
		"UerIdIsZero":             30001,
		"ParseJwtTokenMatchError": 30002,
		"NoFoundThisUser":         30003,
		"CreateUserError":         30004,
		"OperationDbError":        30005,
		"UidBoundError":           30006,
		"LoadProductError":        40000,
		"ThisProductNotExist":     40001,
		"NoFoundThisProp":         50000,
		"NoFoundThisItem":         60000,
		"LoadAssetConfigError":    60001,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_message_error_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_message_error_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_message_error_proto_rawDescGZIP(), []int{0}
}

var File_message_error_proto protoreflect.FileDescriptor

var file_message_error_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0xbb, 0x04, 0x0a,
	0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x73, 0x4e, 0x69, 0x6c, 0x10, 0x90, 0x4e, 0x12, 0x19, 0x0a, 0x14, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x91, 0x4e, 0x12, 0x18, 0x0a, 0x13, 0x4c, 0x6f, 0x61, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x10, 0x92, 0x4e, 0x12,
	0x13, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x45, 0x72,
	0x72, 0x10, 0x93, 0x4e, 0x12, 0x13, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x45, 0x72, 0x72, 0x10, 0x94, 0x4e, 0x12, 0x13, 0x0a, 0x0e, 0x54, 0x68, 0x65,
	0x55, 0x69, 0x64, 0x4e, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x10, 0x95, 0x4e, 0x12, 0x19,
	0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x45, 0x6e, 0x6f, 0x75, 0x67, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x44,
	0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x10, 0x96, 0x4e, 0x12, 0x1b, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0xa1, 0x9c, 0x01, 0x12, 0x15, 0x0a, 0x0f, 0x4e, 0x6f, 0x46, 0x6f, 0x75, 0x6e,
	0x64, 0x54, 0x68, 0x69, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x10, 0xa2, 0x9c, 0x01, 0x12, 0x14, 0x0a,
	0x0e, 0x54, 0x61, 0x73, 0x6b, 0x48, 0x61, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x10,
	0xa3, 0x9c, 0x01, 0x12, 0x11, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x65, 0x74, 0x54, 0x6f, 0x4d, 0x61,
	0x6e, 0x79, 0x10, 0xa4, 0x9c, 0x01, 0x12, 0x18, 0x0a, 0x12, 0x50, 0x61, 0x72, 0x73, 0x65, 0x4a,
	0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0xb0, 0xea, 0x01,
	0x12, 0x11, 0x0a, 0x0b, 0x55, 0x65, 0x72, 0x49, 0x64, 0x49, 0x73, 0x5a, 0x65, 0x72, 0x6f, 0x10,
	0xb1, 0xea, 0x01, 0x12, 0x1d, 0x0a, 0x17, 0x50, 0x61, 0x72, 0x73, 0x65, 0x4a, 0x77, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0xb2,
	0xea, 0x01, 0x12, 0x15, 0x0a, 0x0f, 0x4e, 0x6f, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x68, 0x69,
	0x73, 0x55, 0x73, 0x65, 0x72, 0x10, 0xb3, 0xea, 0x01, 0x12, 0x15, 0x0a, 0x0f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0xb4, 0xea, 0x01,
	0x12, 0x16, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x62, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0xb5, 0xea, 0x01, 0x12, 0x13, 0x0a, 0x0d, 0x55, 0x69, 0x64, 0x42,
	0x6f, 0x75, 0x6e, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0xb6, 0xea, 0x01, 0x12, 0x16, 0x0a,
	0x10, 0x4c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x10, 0xc0, 0xb8, 0x02, 0x12, 0x19, 0x0a, 0x13, 0x54, 0x68, 0x69, 0x73, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0xc1, 0xb8, 0x02,
	0x12, 0x15, 0x0a, 0x0f, 0x4e, 0x6f, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x68, 0x69, 0x73, 0x50,
	0x72, 0x6f, 0x70, 0x10, 0xd0, 0x86, 0x03, 0x12, 0x15, 0x0a, 0x0f, 0x4e, 0x6f, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x54, 0x68, 0x69, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x10, 0xe0, 0xd4, 0x03, 0x12, 0x1a,
	0x0a, 0x14, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0xe1, 0xd4, 0x03, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x75, 0x78, 0x69, 0x61, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_error_proto_rawDescOnce sync.Once
	file_message_error_proto_rawDescData = file_message_error_proto_rawDesc
)

func file_message_error_proto_rawDescGZIP() []byte {
	file_message_error_proto_rawDescOnce.Do(func() {
		file_message_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_error_proto_rawDescData)
	})
	return file_message_error_proto_rawDescData
}

var file_message_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_error_proto_goTypes = []interface{}{
	(ErrorCode)(0), // 0: error.ErrorCode
}
var file_message_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_error_proto_init() }
func file_message_error_proto_init() {
	if File_message_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_error_proto_goTypes,
		DependencyIndexes: file_message_error_proto_depIdxs,
		EnumInfos:         file_message_error_proto_enumTypes,
	}.Build()
	File_message_error_proto = out.File
	file_message_error_proto_rawDesc = nil
	file_message_error_proto_goTypes = nil
	file_message_error_proto_depIdxs = nil
}
