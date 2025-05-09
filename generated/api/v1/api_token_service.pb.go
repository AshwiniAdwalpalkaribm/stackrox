// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: api/v1/api_token_service.proto

package v1

import (
	storage "github.com/stackrox/rox/generated/storage"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type GenerateTokenRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Deprecated: Marked as deprecated in api/v1/api_token_service.proto.
	Role          string                 `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Roles         []string               `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	Expiration    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expiration,proto3" json:"expiration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateTokenRequest) Reset() {
	*x = GenerateTokenRequest{}
	mi := &file_api_v1_api_token_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenRequest) ProtoMessage() {}

func (x *GenerateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_token_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenRequest.ProtoReflect.Descriptor instead.
func (*GenerateTokenRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_api_token_service_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateTokenRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Deprecated: Marked as deprecated in api/v1/api_token_service.proto.
func (x *GenerateTokenRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *GenerateTokenRequest) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *GenerateTokenRequest) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

type GenerateTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Metadata      *storage.TokenMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateTokenResponse) Reset() {
	*x = GenerateTokenResponse{}
	mi := &file_api_v1_api_token_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenResponse) ProtoMessage() {}

func (x *GenerateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_token_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenResponse.ProtoReflect.Descriptor instead.
func (*GenerateTokenResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_api_token_service_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GenerateTokenResponse) GetMetadata() *storage.TokenMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type GetAPITokensRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to RevokedOneof:
	//
	//	*GetAPITokensRequest_Revoked
	RevokedOneof  isGetAPITokensRequest_RevokedOneof `protobuf_oneof:"revoked_oneof"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAPITokensRequest) Reset() {
	*x = GetAPITokensRequest{}
	mi := &file_api_v1_api_token_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAPITokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAPITokensRequest) ProtoMessage() {}

func (x *GetAPITokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_token_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAPITokensRequest.ProtoReflect.Descriptor instead.
func (*GetAPITokensRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_api_token_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAPITokensRequest) GetRevokedOneof() isGetAPITokensRequest_RevokedOneof {
	if x != nil {
		return x.RevokedOneof
	}
	return nil
}

func (x *GetAPITokensRequest) GetRevoked() bool {
	if x != nil {
		if x, ok := x.RevokedOneof.(*GetAPITokensRequest_Revoked); ok {
			return x.Revoked
		}
	}
	return false
}

type isGetAPITokensRequest_RevokedOneof interface {
	isGetAPITokensRequest_RevokedOneof()
}

type GetAPITokensRequest_Revoked struct {
	Revoked bool `protobuf:"varint,1,opt,name=revoked,proto3,oneof"`
}

func (*GetAPITokensRequest_Revoked) isGetAPITokensRequest_RevokedOneof() {}

type GetAPITokensResponse struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Tokens        []*storage.TokenMetadata `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAPITokensResponse) Reset() {
	*x = GetAPITokensResponse{}
	mi := &file_api_v1_api_token_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAPITokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAPITokensResponse) ProtoMessage() {}

func (x *GetAPITokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_token_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAPITokensResponse.ProtoReflect.Descriptor instead.
func (*GetAPITokensResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_api_token_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAPITokensResponse) GetTokens() []*storage.TokenMetadata {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type ListAllowedTokenRolesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoleNames     []string               `protobuf:"bytes,1,rep,name=roleNames,proto3" json:"roleNames,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAllowedTokenRolesResponse) Reset() {
	*x = ListAllowedTokenRolesResponse{}
	mi := &file_api_v1_api_token_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAllowedTokenRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllowedTokenRolesResponse) ProtoMessage() {}

func (x *ListAllowedTokenRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_token_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllowedTokenRolesResponse.ProtoReflect.Descriptor instead.
func (*ListAllowedTokenRolesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_api_token_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListAllowedTokenRolesResponse) GetRoleNames() []string {
	if x != nil {
		return x.RoleNames
	}
	return nil
}

var File_api_v1_api_token_service_proto protoreflect.FileDescriptor

const file_api_v1_api_token_service_proto_rawDesc = "" +
	"\n" +
	"\x1eapi/v1/api_token_service.proto\x12\x02v1\x1a\x13api/v1/common.proto\x1a\x12api/v1/empty.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x17storage/api_token.proto\"\x94\x01\n" +
	"\x14GenerateTokenRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x16\n" +
	"\x04role\x18\x02 \x01(\tB\x02\x18\x01R\x04role\x12\x14\n" +
	"\x05roles\x18\x03 \x03(\tR\x05roles\x12:\n" +
	"\n" +
	"expiration\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"expiration\"a\n" +
	"\x15GenerateTokenResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x122\n" +
	"\bmetadata\x18\x02 \x01(\v2\x16.storage.TokenMetadataR\bmetadata\"B\n" +
	"\x13GetAPITokensRequest\x12\x1a\n" +
	"\arevoked\x18\x01 \x01(\bH\x00R\arevokedB\x0f\n" +
	"\rrevoked_oneof\"F\n" +
	"\x14GetAPITokensResponse\x12.\n" +
	"\x06tokens\x18\x01 \x03(\v2\x16.storage.TokenMetadataR\x06tokens\"=\n" +
	"\x1dListAllowedTokenRolesResponse\x12\x1c\n" +
	"\troleNames\x18\x01 \x03(\tR\troleNames2\xed\x03\n" +
	"\x0fAPITokenService\x12S\n" +
	"\vGetAPIToken\x12\x10.v1.ResourceByID\x1a\x16.storage.TokenMetadata\"\x1a\x82\xd3\xe4\x93\x02\x14\x12\x12/v1/apitokens/{id}\x12X\n" +
	"\fGetAPITokens\x12\x17.v1.GetAPITokensRequest\x1a\x18.v1.GetAPITokensResponse\"\x15\x82\xd3\xe4\x93\x02\x0f\x12\r/v1/apitokens\x12g\n" +
	"\rGenerateToken\x12\x18.v1.GenerateTokenRequest\x1a\x19.v1.GenerateTokenResponse\"!\x82\xd3\xe4\x93\x02\x1b:\x01*\"\x16/v1/apitokens/generate\x12M\n" +
	"\vRevokeToken\x12\x10.v1.ResourceByID\x1a\t.v1.Empty\"!\x82\xd3\xe4\x93\x02\x1b2\x19/v1/apitokens/revoke/{id}\x12s\n" +
	"\x15ListAllowedTokenRoles\x12\t.v1.Empty\x1a!.v1.ListAllowedTokenRolesResponse\",\x82\xd3\xe4\x93\x02&\x12$/v1/apitokens/generate/allowed-rolesB'\n" +
	"\x18io.stackrox.proto.api.v1Z\v./api/v1;v1X\x02b\x06proto3"

var (
	file_api_v1_api_token_service_proto_rawDescOnce sync.Once
	file_api_v1_api_token_service_proto_rawDescData []byte
)

func file_api_v1_api_token_service_proto_rawDescGZIP() []byte {
	file_api_v1_api_token_service_proto_rawDescOnce.Do(func() {
		file_api_v1_api_token_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_api_token_service_proto_rawDesc), len(file_api_v1_api_token_service_proto_rawDesc)))
	})
	return file_api_v1_api_token_service_proto_rawDescData
}

var file_api_v1_api_token_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_api_token_service_proto_goTypes = []any{
	(*GenerateTokenRequest)(nil),          // 0: v1.GenerateTokenRequest
	(*GenerateTokenResponse)(nil),         // 1: v1.GenerateTokenResponse
	(*GetAPITokensRequest)(nil),           // 2: v1.GetAPITokensRequest
	(*GetAPITokensResponse)(nil),          // 3: v1.GetAPITokensResponse
	(*ListAllowedTokenRolesResponse)(nil), // 4: v1.ListAllowedTokenRolesResponse
	(*timestamppb.Timestamp)(nil),         // 5: google.protobuf.Timestamp
	(*storage.TokenMetadata)(nil),         // 6: storage.TokenMetadata
	(*ResourceByID)(nil),                  // 7: v1.ResourceByID
	(*Empty)(nil),                         // 8: v1.Empty
}
var file_api_v1_api_token_service_proto_depIdxs = []int32{
	5, // 0: v1.GenerateTokenRequest.expiration:type_name -> google.protobuf.Timestamp
	6, // 1: v1.GenerateTokenResponse.metadata:type_name -> storage.TokenMetadata
	6, // 2: v1.GetAPITokensResponse.tokens:type_name -> storage.TokenMetadata
	7, // 3: v1.APITokenService.GetAPIToken:input_type -> v1.ResourceByID
	2, // 4: v1.APITokenService.GetAPITokens:input_type -> v1.GetAPITokensRequest
	0, // 5: v1.APITokenService.GenerateToken:input_type -> v1.GenerateTokenRequest
	7, // 6: v1.APITokenService.RevokeToken:input_type -> v1.ResourceByID
	8, // 7: v1.APITokenService.ListAllowedTokenRoles:input_type -> v1.Empty
	6, // 8: v1.APITokenService.GetAPIToken:output_type -> storage.TokenMetadata
	3, // 9: v1.APITokenService.GetAPITokens:output_type -> v1.GetAPITokensResponse
	1, // 10: v1.APITokenService.GenerateToken:output_type -> v1.GenerateTokenResponse
	8, // 11: v1.APITokenService.RevokeToken:output_type -> v1.Empty
	4, // 12: v1.APITokenService.ListAllowedTokenRoles:output_type -> v1.ListAllowedTokenRolesResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_v1_api_token_service_proto_init() }
func file_api_v1_api_token_service_proto_init() {
	if File_api_v1_api_token_service_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	file_api_v1_empty_proto_init()
	file_api_v1_api_token_service_proto_msgTypes[2].OneofWrappers = []any{
		(*GetAPITokensRequest_Revoked)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_api_token_service_proto_rawDesc), len(file_api_v1_api_token_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_api_token_service_proto_goTypes,
		DependencyIndexes: file_api_v1_api_token_service_proto_depIdxs,
		MessageInfos:      file_api_v1_api_token_service_proto_msgTypes,
	}.Build()
	File_api_v1_api_token_service_proto = out.File
	file_api_v1_api_token_service_proto_goTypes = nil
	file_api_v1_api_token_service_proto_depIdxs = nil
}
