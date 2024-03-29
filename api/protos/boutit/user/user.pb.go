// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: api/protos/boutit/user/user.proto

package user

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

type SignupUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailOrPhone string `protobuf:"bytes,1,opt,name=email_or_phone,json=emailOrPhone,proto3" json:"email_or_phone,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password     string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Username     string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *SignupUserRequest) Reset() {
	*x = SignupUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_boutit_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupUserRequest) ProtoMessage() {}

func (x *SignupUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_boutit_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupUserRequest.ProtoReflect.Descriptor instead.
func (*SignupUserRequest) Descriptor() ([]byte, []int) {
	return file_api_protos_boutit_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *SignupUserRequest) GetEmailOrPhone() string {
	if x != nil {
		return x.EmailOrPhone
	}
	return ""
}

func (x *SignupUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SignupUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignupUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type SignupUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *SignupUserResponse) Reset() {
	*x = SignupUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_boutit_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupUserResponse) ProtoMessage() {}

func (x *SignupUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_boutit_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupUserResponse.ProtoReflect.Descriptor instead.
func (*SignupUserResponse) Descriptor() ([]byte, []int) {
	return file_api_protos_boutit_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *SignupUserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type LoginUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailPhoneOrUsername string `protobuf:"bytes,1,opt,name=email_phone_or_username,json=emailPhoneOrUsername,proto3" json:"email_phone_or_username,omitempty"`
	Password             string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginUserRequest) Reset() {
	*x = LoginUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_boutit_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserRequest) ProtoMessage() {}

func (x *LoginUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_boutit_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserRequest.ProtoReflect.Descriptor instead.
func (*LoginUserRequest) Descriptor() ([]byte, []int) {
	return file_api_protos_boutit_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *LoginUserRequest) GetEmailPhoneOrUsername() string {
	if x != nil {
		return x.EmailPhoneOrUsername
	}
	return ""
}

func (x *LoginUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoginUserResponse) Reset() {
	*x = LoginUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_boutit_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserResponse) ProtoMessage() {}

func (x *LoginUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_boutit_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserResponse.ProtoReflect.Descriptor instead.
func (*LoginUserResponse) Descriptor() ([]byte, []int) {
	return file_api_protos_boutit_user_user_proto_rawDescGZIP(), []int{3}
}

var File_api_protos_boutit_user_user_proto protoreflect.FileDescriptor

var file_api_protos_boutit_user_user_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6f, 0x75,
	0x74, 0x69, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x62, 0x6f, 0x75, 0x74, 0x69, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x22, 0x85, 0x01, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x6f, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x4f, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x12,
	0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x10, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x35, 0x0a, 0x17, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6f,
	0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x14, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4f, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x6f, 0x75, 0x74, 0x69, 0x74, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6f, 0x75,
	0x74, 0x69, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protos_boutit_user_user_proto_rawDescOnce sync.Once
	file_api_protos_boutit_user_user_proto_rawDescData = file_api_protos_boutit_user_user_proto_rawDesc
)

func file_api_protos_boutit_user_user_proto_rawDescGZIP() []byte {
	file_api_protos_boutit_user_user_proto_rawDescOnce.Do(func() {
		file_api_protos_boutit_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protos_boutit_user_user_proto_rawDescData)
	})
	return file_api_protos_boutit_user_user_proto_rawDescData
}

var file_api_protos_boutit_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_protos_boutit_user_user_proto_goTypes = []interface{}{
	(*SignupUserRequest)(nil),  // 0: boutit.user.api.SignupUserRequest
	(*SignupUserResponse)(nil), // 1: boutit.user.api.SignupUserResponse
	(*LoginUserRequest)(nil),   // 2: boutit.user.api.LoginUserRequest
	(*LoginUserResponse)(nil),  // 3: boutit.user.api.LoginUserResponse
}
var file_api_protos_boutit_user_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_protos_boutit_user_user_proto_init() }
func file_api_protos_boutit_user_user_proto_init() {
	if File_api_protos_boutit_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_protos_boutit_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupUserRequest); i {
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
		file_api_protos_boutit_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupUserResponse); i {
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
		file_api_protos_boutit_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUserRequest); i {
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
		file_api_protos_boutit_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUserResponse); i {
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
			RawDescriptor: file_api_protos_boutit_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_protos_boutit_user_user_proto_goTypes,
		DependencyIndexes: file_api_protos_boutit_user_user_proto_depIdxs,
		MessageInfos:      file_api_protos_boutit_user_user_proto_msgTypes,
	}.Build()
	File_api_protos_boutit_user_user_proto = out.File
	file_api_protos_boutit_user_user_proto_rawDesc = nil
	file_api_protos_boutit_user_user_proto_goTypes = nil
	file_api_protos_boutit_user_user_proto_depIdxs = nil
}
