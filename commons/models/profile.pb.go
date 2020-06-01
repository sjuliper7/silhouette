// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.7.1
// source: profile.proto

package models

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID      int64  `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Address     string `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	WorkAt      string `protobuf:"bytes,4,opt,name=WorkAt,proto3" json:"WorkAt,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Gender      string `protobuf:"bytes,6,opt,name=Gender,proto3" json:"Gender,omitempty"`
	IsActive    bool   `protobuf:"varint,7,opt,name=IsActive,proto3" json:"IsActive,omitempty"`
	CreatedAt   string `protobuf:"bytes,8,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt   string `protobuf:"bytes,9,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	Name        string `protobuf:"bytes,10,opt,name=Name,proto3" json:"Name,omitempty"`
	DateOfBirth string `protobuf:"bytes,11,opt,name=DateOfBirth,proto3" json:"DateOfBirth,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{0}
}

func (x *Profile) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Profile) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Profile) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Profile) GetWorkAt() string {
	if x != nil {
		return x.WorkAt
	}
	return ""
}

func (x *Profile) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Profile) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Profile) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Profile) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Profile) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Profile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Profile) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

type UserGetProfileArguments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID      int64  `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Address     string `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	WorkAt      string `protobuf:"bytes,4,opt,name=WorkAt,proto3" json:"WorkAt,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Gender      string `protobuf:"bytes,6,opt,name=Gender,proto3" json:"Gender,omitempty"`
}

func (x *UserGetProfileArguments) Reset() {
	*x = UserGetProfileArguments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGetProfileArguments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGetProfileArguments) ProtoMessage() {}

func (x *UserGetProfileArguments) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGetProfileArguments.ProtoReflect.Descriptor instead.
func (*UserGetProfileArguments) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{1}
}

func (x *UserGetProfileArguments) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UserGetProfileArguments) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UserGetProfileArguments) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UserGetProfileArguments) GetWorkAt() string {
	if x != nil {
		return x.WorkAt
	}
	return ""
}

func (x *UserGetProfileArguments) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserGetProfileArguments) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

var File_profile_proto protoreflect.FileDescriptor

var file_profile_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0xab, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x41, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x41, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72,
	0x74, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x66,
	0x42, 0x69, 0x72, 0x74, 0x68, 0x22, 0xad, 0x01, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x41, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x32, 0x4c, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x12, 0x40, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x1f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x1a, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x22, 0x00, 0x42, 0x02, 0x5a, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profile_proto_rawDescOnce sync.Once
	file_profile_proto_rawDescData = file_profile_proto_rawDesc
)

func file_profile_proto_rawDescGZIP() []byte {
	file_profile_proto_rawDescOnce.Do(func() {
		file_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_profile_proto_rawDescData)
	})
	return file_profile_proto_rawDescData
}

var file_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_profile_proto_goTypes = []interface{}{
	(*Profile)(nil),                 // 0: models.Profile
	(*UserGetProfileArguments)(nil), // 1: models.UserGetProfileArguments
}
var file_profile_proto_depIdxs = []int32{
	1, // 0: models.Profiles.GetProfile:input_type -> models.UserGetProfileArguments
	0, // 1: models.Profiles.GetProfile:output_type -> models.Profile
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_profile_proto_init() }
func file_profile_proto_init() {
	if File_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGetProfileArguments); i {
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
			RawDescriptor: file_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profile_proto_goTypes,
		DependencyIndexes: file_profile_proto_depIdxs,
		MessageInfos:      file_profile_proto_msgTypes,
	}.Build()
	File_profile_proto = out.File
	file_profile_proto_rawDesc = nil
	file_profile_proto_goTypes = nil
	file_profile_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProfilesClient is the client API for Profiles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfilesClient interface {
	GetProfile(ctx context.Context, in *UserGetProfileArguments, opts ...grpc.CallOption) (*Profile, error)
}

type profilesClient struct {
	cc grpc.ClientConnInterface
}

func NewProfilesClient(cc grpc.ClientConnInterface) ProfilesClient {
	return &profilesClient{cc}
}

func (c *profilesClient) GetProfile(ctx context.Context, in *UserGetProfileArguments, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/models.Profiles/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfilesServer is the server API for Profiles service.
type ProfilesServer interface {
	GetProfile(context.Context, *UserGetProfileArguments) (*Profile, error)
}

// UnimplementedProfilesServer can be embedded to have forward compatible implementations.
type UnimplementedProfilesServer struct {
}

func (*UnimplementedProfilesServer) GetProfile(context.Context, *UserGetProfileArguments) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}

func RegisterProfilesServer(s *grpc.Server, srv ProfilesServer) {
	s.RegisterService(&_Profiles_serviceDesc, srv)
}

func _Profiles_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetProfileArguments)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Profiles/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServer).GetProfile(ctx, req.(*UserGetProfileArguments))
	}
	return interceptor(ctx, in, info, handler)
}

var _Profiles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "models.Profiles",
	HandlerType: (*ProfilesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfile",
			Handler:    _Profiles_GetProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}
