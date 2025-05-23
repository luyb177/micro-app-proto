// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.6
// source: api/user/v1/user.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Money         uint32                 `protobuf:"varint,3,opt,name=money,proto3" json:"money,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_api_user_v1_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetMoney() uint32 {
	if x != nil {
		return x.Money
	}
	return 0
}

type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_api_user_v1_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_api_user_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RegisterReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	User          *User                  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	mi := &file_api_user_v1_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_api_user_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RegisterReply) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type DecrMoneyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Money         uint32                 `protobuf:"varint,2,opt,name=money,proto3" json:"money,omitempty"`
	Gid           string                 `protobuf:"bytes,3,opt,name=gid,proto3" json:"gid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DecrMoneyRequest) Reset() {
	*x = DecrMoneyRequest{}
	mi := &file_api_user_v1_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DecrMoneyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecrMoneyRequest) ProtoMessage() {}

func (x *DecrMoneyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecrMoneyRequest.ProtoReflect.Descriptor instead.
func (*DecrMoneyRequest) Descriptor() ([]byte, []int) {
	return file_api_user_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *DecrMoneyRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DecrMoneyRequest) GetMoney() uint32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *DecrMoneyRequest) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

type DecrMoneyReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DecrMoneyReply) Reset() {
	*x = DecrMoneyReply{}
	mi := &file_api_user_v1_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DecrMoneyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecrMoneyReply) ProtoMessage() {}

func (x *DecrMoneyReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecrMoneyReply.ProtoReflect.Descriptor instead.
func (*DecrMoneyReply) Descriptor() ([]byte, []int) {
	return file_api_user_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *DecrMoneyReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type IncrMoneyReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IncrMoneyReply) Reset() {
	*x = IncrMoneyReply{}
	mi := &file_api_user_v1_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncrMoneyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrMoneyReply) ProtoMessage() {}

func (x *IncrMoneyReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrMoneyReply.ProtoReflect.Descriptor instead.
func (*IncrMoneyReply) Descriptor() ([]byte, []int) {
	return file_api_user_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *IncrMoneyReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_api_user_v1_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_api_user_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	User          *User                  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserReply) Reset() {
	*x = GetUserReply{}
	mi := &file_api_user_v1_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReply) ProtoMessage() {}

func (x *GetUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReply.ProtoReflect.Descriptor instead.
func (*GetUserReply) Descriptor() ([]byte, []int) {
	return file_api_user_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetUserReply) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_api_user_v1_user_proto protoreflect.FileDescriptor

const file_api_user_v1_user_proto_rawDesc = "" +
	"\n" +
	"\x16api/user/v1/user.proto\x12\auser.v1\x1a\x1cgoogle/api/annotations.proto\"@\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05money\x18\x03 \x01(\rR\x05money\"%\n" +
	"\x0fRegisterRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"L\n" +
	"\rRegisterReply\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12!\n" +
	"\x04user\x18\x02 \x01(\v2\r.user.v1.UserR\x04user\"S\n" +
	"\x10DecrMoneyRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\rR\x06userId\x12\x14\n" +
	"\x05money\x18\x02 \x01(\rR\x05money\x12\x10\n" +
	"\x03gid\x18\x03 \x01(\tR\x03gid\"*\n" +
	"\x0eDecrMoneyReply\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"*\n" +
	"\x0eIncrMoneyReply\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\" \n" +
	"\x0eGetUserRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\"K\n" +
	"\fGetUserReply\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12!\n" +
	"\x04user\x18\x02 \x01(\v2\r.user.v1.UserR\x04user2\xd0\x02\n" +
	"\x05Users\x12R\n" +
	"\bRegister\x12\x18.user.v1.RegisterRequest\x1a\x16.user.v1.RegisterReply\"\x14\x82\xd3\xe4\x93\x02\x0e:\x01*\"\t/register\x12Q\n" +
	"\tDecrMoney\x12\x19.user.v1.DecrMoneyRequest\x1a\x17.user.v1.DecrMoneyReply\"\x10\x82\xd3\xe4\x93\x02\n" +
	":\x01*\"\x05/decr\x12Q\n" +
	"\tIncrMoney\x12\x19.user.v1.DecrMoneyRequest\x1a\x17.user.v1.IncrMoneyReply\"\x10\x82\xd3\xe4\x93\x02\n" +
	":\x01*\"\x05/incr\x12M\n" +
	"\aGetUser\x12\x17.user.v1.GetUserRequest\x1a\x15.user.v1.GetUserReply\"\x12\x82\xd3\xe4\x93\x02\f\x12\n" +
	"/user/{id}B\x10Z\x0eapi/user/v1;v1b\x06proto3"

var (
	file_api_user_v1_user_proto_rawDescOnce sync.Once
	file_api_user_v1_user_proto_rawDescData []byte
)

func file_api_user_v1_user_proto_rawDescGZIP() []byte {
	file_api_user_v1_user_proto_rawDescOnce.Do(func() {
		file_api_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_user_v1_user_proto_rawDesc), len(file_api_user_v1_user_proto_rawDesc)))
	})
	return file_api_user_v1_user_proto_rawDescData
}

var file_api_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_user_v1_user_proto_goTypes = []any{
	(*User)(nil),             // 0: user.v1.User
	(*RegisterRequest)(nil),  // 1: user.v1.RegisterRequest
	(*RegisterReply)(nil),    // 2: user.v1.RegisterReply
	(*DecrMoneyRequest)(nil), // 3: user.v1.DecrMoneyRequest
	(*DecrMoneyReply)(nil),   // 4: user.v1.DecrMoneyReply
	(*IncrMoneyReply)(nil),   // 5: user.v1.IncrMoneyReply
	(*GetUserRequest)(nil),   // 6: user.v1.GetUserRequest
	(*GetUserReply)(nil),     // 7: user.v1.GetUserReply
}
var file_api_user_v1_user_proto_depIdxs = []int32{
	0, // 0: user.v1.RegisterReply.user:type_name -> user.v1.User
	0, // 1: user.v1.GetUserReply.user:type_name -> user.v1.User
	1, // 2: user.v1.Users.Register:input_type -> user.v1.RegisterRequest
	3, // 3: user.v1.Users.DecrMoney:input_type -> user.v1.DecrMoneyRequest
	3, // 4: user.v1.Users.IncrMoney:input_type -> user.v1.DecrMoneyRequest
	6, // 5: user.v1.Users.GetUser:input_type -> user.v1.GetUserRequest
	2, // 6: user.v1.Users.Register:output_type -> user.v1.RegisterReply
	4, // 7: user.v1.Users.DecrMoney:output_type -> user.v1.DecrMoneyReply
	5, // 8: user.v1.Users.IncrMoney:output_type -> user.v1.IncrMoneyReply
	7, // 9: user.v1.Users.GetUser:output_type -> user.v1.GetUserReply
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_user_v1_user_proto_init() }
func file_api_user_v1_user_proto_init() {
	if File_api_user_v1_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_user_v1_user_proto_rawDesc), len(file_api_user_v1_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_user_v1_user_proto_goTypes,
		DependencyIndexes: file_api_user_v1_user_proto_depIdxs,
		MessageInfos:      file_api_user_v1_user_proto_msgTypes,
	}.Build()
	File_api_user_v1_user_proto = out.File
	file_api_user_v1_user_proto_goTypes = nil
	file_api_user_v1_user_proto_depIdxs = nil
}
