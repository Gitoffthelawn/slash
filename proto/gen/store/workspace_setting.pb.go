// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: store/workspace_setting.proto

package store

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

type WorkspaceSettingKey int32

const (
	WorkspaceSettingKey_WORKSPACE_SETTING_KEY_UNSPECIFIED WorkspaceSettingKey = 0
	// Workspace general settings.
	WorkspaceSettingKey_WORKSPACE_SETTING_GENERAL WorkspaceSettingKey = 1
	// Workspace security settings.
	WorkspaceSettingKey_WORKSPACE_SETTING_SECURITY WorkspaceSettingKey = 2
	// Workspace shortcut-related settings.
	WorkspaceSettingKey_WORKSPACE_SETTING_SHORTCUT_RELATED WorkspaceSettingKey = 3
	// Workspace identity provider settings.
	WorkspaceSettingKey_WORKSPACE_SETTING_IDENTITY_PROVIDER WorkspaceSettingKey = 4
	// TODO: remove the following keys.
	// The license key.
	WorkspaceSettingKey_WORKSPACE_SETTING_LICENSE_KEY WorkspaceSettingKey = 10
	// The secret session key used to encrypt session data.
	WorkspaceSettingKey_WORKSPACE_SETTING_SECRET_SESSION WorkspaceSettingKey = 11
	// The custom style.
	WorkspaceSettingKey_WORKSPACE_SETTING_CUSTOM_STYLE WorkspaceSettingKey = 12
	// The default visibility of shortcuts and collections.
	WorkspaceSettingKey_WORKSPACE_SETTING_DEFAULT_VISIBILITY WorkspaceSettingKey = 13
)

// Enum value maps for WorkspaceSettingKey.
var (
	WorkspaceSettingKey_name = map[int32]string{
		0:  "WORKSPACE_SETTING_KEY_UNSPECIFIED",
		1:  "WORKSPACE_SETTING_GENERAL",
		2:  "WORKSPACE_SETTING_SECURITY",
		3:  "WORKSPACE_SETTING_SHORTCUT_RELATED",
		4:  "WORKSPACE_SETTING_IDENTITY_PROVIDER",
		10: "WORKSPACE_SETTING_LICENSE_KEY",
		11: "WORKSPACE_SETTING_SECRET_SESSION",
		12: "WORKSPACE_SETTING_CUSTOM_STYLE",
		13: "WORKSPACE_SETTING_DEFAULT_VISIBILITY",
	}
	WorkspaceSettingKey_value = map[string]int32{
		"WORKSPACE_SETTING_KEY_UNSPECIFIED":    0,
		"WORKSPACE_SETTING_GENERAL":            1,
		"WORKSPACE_SETTING_SECURITY":           2,
		"WORKSPACE_SETTING_SHORTCUT_RELATED":   3,
		"WORKSPACE_SETTING_IDENTITY_PROVIDER":  4,
		"WORKSPACE_SETTING_LICENSE_KEY":        10,
		"WORKSPACE_SETTING_SECRET_SESSION":     11,
		"WORKSPACE_SETTING_CUSTOM_STYLE":       12,
		"WORKSPACE_SETTING_DEFAULT_VISIBILITY": 13,
	}
)

func (x WorkspaceSettingKey) Enum() *WorkspaceSettingKey {
	p := new(WorkspaceSettingKey)
	*p = x
	return p
}

func (x WorkspaceSettingKey) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkspaceSettingKey) Descriptor() protoreflect.EnumDescriptor {
	return file_store_workspace_setting_proto_enumTypes[0].Descriptor()
}

func (WorkspaceSettingKey) Type() protoreflect.EnumType {
	return &file_store_workspace_setting_proto_enumTypes[0]
}

func (x WorkspaceSettingKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkspaceSettingKey.Descriptor instead.
func (WorkspaceSettingKey) EnumDescriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{0}
}

type WorkspaceSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key WorkspaceSettingKey `protobuf:"varint,1,opt,name=key,proto3,enum=slash.store.WorkspaceSettingKey" json:"key,omitempty"`
	Raw string              `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
	// Types that are assignable to Value:
	//
	//	*WorkspaceSetting_General
	//	*WorkspaceSetting_Security
	//	*WorkspaceSetting_ShortcutRelated
	//	*WorkspaceSetting_IdentityProvider
	Value isWorkspaceSetting_Value `protobuf_oneof:"value"`
}

func (x *WorkspaceSetting) Reset() {
	*x = WorkspaceSetting{}
	mi := &file_store_workspace_setting_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkspaceSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceSetting) ProtoMessage() {}

func (x *WorkspaceSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceSetting.ProtoReflect.Descriptor instead.
func (*WorkspaceSetting) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{0}
}

func (x *WorkspaceSetting) GetKey() WorkspaceSettingKey {
	if x != nil {
		return x.Key
	}
	return WorkspaceSettingKey_WORKSPACE_SETTING_KEY_UNSPECIFIED
}

func (x *WorkspaceSetting) GetRaw() string {
	if x != nil {
		return x.Raw
	}
	return ""
}

func (m *WorkspaceSetting) GetValue() isWorkspaceSetting_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *WorkspaceSetting) GetGeneral() *WorkspaceSetting_GeneralSetting {
	if x, ok := x.GetValue().(*WorkspaceSetting_General); ok {
		return x.General
	}
	return nil
}

func (x *WorkspaceSetting) GetSecurity() *WorkspaceSetting_SecuritySetting {
	if x, ok := x.GetValue().(*WorkspaceSetting_Security); ok {
		return x.Security
	}
	return nil
}

func (x *WorkspaceSetting) GetShortcutRelated() *WorkspaceSetting_ShortcutRelatedSetting {
	if x, ok := x.GetValue().(*WorkspaceSetting_ShortcutRelated); ok {
		return x.ShortcutRelated
	}
	return nil
}

func (x *WorkspaceSetting) GetIdentityProvider() *WorkspaceSetting_IdentityProviderSetting {
	if x, ok := x.GetValue().(*WorkspaceSetting_IdentityProvider); ok {
		return x.IdentityProvider
	}
	return nil
}

type isWorkspaceSetting_Value interface {
	isWorkspaceSetting_Value()
}

type WorkspaceSetting_General struct {
	General *WorkspaceSetting_GeneralSetting `protobuf:"bytes,3,opt,name=general,proto3,oneof"`
}

type WorkspaceSetting_Security struct {
	Security *WorkspaceSetting_SecuritySetting `protobuf:"bytes,4,opt,name=security,proto3,oneof"`
}

type WorkspaceSetting_ShortcutRelated struct {
	ShortcutRelated *WorkspaceSetting_ShortcutRelatedSetting `protobuf:"bytes,5,opt,name=shortcut_related,json=shortcutRelated,proto3,oneof"`
}

type WorkspaceSetting_IdentityProvider struct {
	IdentityProvider *WorkspaceSetting_IdentityProviderSetting `protobuf:"bytes,6,opt,name=identity_provider,json=identityProvider,proto3,oneof"`
}

func (*WorkspaceSetting_General) isWorkspaceSetting_Value() {}

func (*WorkspaceSetting_Security) isWorkspaceSetting_Value() {}

func (*WorkspaceSetting_ShortcutRelated) isWorkspaceSetting_Value() {}

func (*WorkspaceSetting_IdentityProvider) isWorkspaceSetting_Value() {}

type WorkspaceSetting_GeneralSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretSession string `protobuf:"bytes,1,opt,name=secret_session,json=secretSession,proto3" json:"secret_session,omitempty"`
	LicenseKey    string `protobuf:"bytes,2,opt,name=license_key,json=licenseKey,proto3" json:"license_key,omitempty"`
	InstanceUrl   string `protobuf:"bytes,3,opt,name=instance_url,json=instanceUrl,proto3" json:"instance_url,omitempty"`
	Branding      []byte `protobuf:"bytes,4,opt,name=branding,proto3" json:"branding,omitempty"`
	CustomStyle   string `protobuf:"bytes,5,opt,name=custom_style,json=customStyle,proto3" json:"custom_style,omitempty"`
}

func (x *WorkspaceSetting_GeneralSetting) Reset() {
	*x = WorkspaceSetting_GeneralSetting{}
	mi := &file_store_workspace_setting_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkspaceSetting_GeneralSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceSetting_GeneralSetting) ProtoMessage() {}

func (x *WorkspaceSetting_GeneralSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceSetting_GeneralSetting.ProtoReflect.Descriptor instead.
func (*WorkspaceSetting_GeneralSetting) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{0, 0}
}

func (x *WorkspaceSetting_GeneralSetting) GetSecretSession() string {
	if x != nil {
		return x.SecretSession
	}
	return ""
}

func (x *WorkspaceSetting_GeneralSetting) GetLicenseKey() string {
	if x != nil {
		return x.LicenseKey
	}
	return ""
}

func (x *WorkspaceSetting_GeneralSetting) GetInstanceUrl() string {
	if x != nil {
		return x.InstanceUrl
	}
	return ""
}

func (x *WorkspaceSetting_GeneralSetting) GetBranding() []byte {
	if x != nil {
		return x.Branding
	}
	return nil
}

func (x *WorkspaceSetting_GeneralSetting) GetCustomStyle() string {
	if x != nil {
		return x.CustomStyle
	}
	return ""
}

type WorkspaceSetting_SecuritySetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisallowUserRegistration bool `protobuf:"varint,1,opt,name=disallow_user_registration,json=disallowUserRegistration,proto3" json:"disallow_user_registration,omitempty"`
	DisallowPasswordAuth     bool `protobuf:"varint,2,opt,name=disallow_password_auth,json=disallowPasswordAuth,proto3" json:"disallow_password_auth,omitempty"`
}

func (x *WorkspaceSetting_SecuritySetting) Reset() {
	*x = WorkspaceSetting_SecuritySetting{}
	mi := &file_store_workspace_setting_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkspaceSetting_SecuritySetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceSetting_SecuritySetting) ProtoMessage() {}

func (x *WorkspaceSetting_SecuritySetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceSetting_SecuritySetting.ProtoReflect.Descriptor instead.
func (*WorkspaceSetting_SecuritySetting) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{0, 1}
}

func (x *WorkspaceSetting_SecuritySetting) GetDisallowUserRegistration() bool {
	if x != nil {
		return x.DisallowUserRegistration
	}
	return false
}

func (x *WorkspaceSetting_SecuritySetting) GetDisallowPasswordAuth() bool {
	if x != nil {
		return x.DisallowPasswordAuth
	}
	return false
}

type WorkspaceSetting_ShortcutRelatedSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultVisibility Visibility `protobuf:"varint,1,opt,name=default_visibility,json=defaultVisibility,proto3,enum=slash.store.Visibility" json:"default_visibility,omitempty"`
}

func (x *WorkspaceSetting_ShortcutRelatedSetting) Reset() {
	*x = WorkspaceSetting_ShortcutRelatedSetting{}
	mi := &file_store_workspace_setting_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkspaceSetting_ShortcutRelatedSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceSetting_ShortcutRelatedSetting) ProtoMessage() {}

func (x *WorkspaceSetting_ShortcutRelatedSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceSetting_ShortcutRelatedSetting.ProtoReflect.Descriptor instead.
func (*WorkspaceSetting_ShortcutRelatedSetting) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{0, 2}
}

func (x *WorkspaceSetting_ShortcutRelatedSetting) GetDefaultVisibility() Visibility {
	if x != nil {
		return x.DefaultVisibility
	}
	return Visibility_VISIBILITY_UNSPECIFIED
}

type WorkspaceSetting_IdentityProviderSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdentityProviders []*IdentityProvider `protobuf:"bytes,1,rep,name=identity_providers,json=identityProviders,proto3" json:"identity_providers,omitempty"`
}

func (x *WorkspaceSetting_IdentityProviderSetting) Reset() {
	*x = WorkspaceSetting_IdentityProviderSetting{}
	mi := &file_store_workspace_setting_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkspaceSetting_IdentityProviderSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceSetting_IdentityProviderSetting) ProtoMessage() {}

func (x *WorkspaceSetting_IdentityProviderSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_workspace_setting_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceSetting_IdentityProviderSetting.ProtoReflect.Descriptor instead.
func (*WorkspaceSetting_IdentityProviderSetting) Descriptor() ([]byte, []int) {
	return file_store_workspace_setting_proto_rawDescGZIP(), []int{0, 3}
}

func (x *WorkspaceSetting_IdentityProviderSetting) GetIdentityProviders() []*IdentityProvider {
	if x != nil {
		return x.IdentityProviders
	}
	return nil
}

var File_store_workspace_setting_proto protoreflect.FileDescriptor

var file_store_workspace_setting_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x12, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x64, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd1, 0x07, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x32, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61,
	0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x48, 0x0a, 0x07,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x07, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x4b, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x61, 0x0a, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x5f,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e,
	0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x63, 0x75, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x12, 0x64, 0x0a, 0x11, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x35, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x10, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0xba, 0x01, 0x0a,
	0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x25, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x1a, 0x85, 0x01, 0x0a, 0x0f, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x3c, 0x0a,
	0x1a, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x18, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x16, 0x64,
	0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x64, 0x69, 0x73,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x41, 0x75, 0x74,
	0x68, 0x1a, 0x60, 0x0a, 0x16, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x46, 0x0a, 0x12, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x11, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x1a, 0x67, 0x0a, 0x17, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x4c,
	0x0a, 0x12, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x6c, 0x61,
	0x73, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x11, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x42, 0x07, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0xe3, 0x02, 0x0a, 0x13, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x0a,
	0x21, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43,
	0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41,
	0x4c, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45,
	0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54,
	0x59, 0x10, 0x02, 0x12, 0x26, 0x0a, 0x22, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45,
	0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x48, 0x4f, 0x52, 0x54, 0x43, 0x55,
	0x54, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x27, 0x0a, 0x23, 0x57,
	0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47,
	0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44,
	0x45, 0x52, 0x10, 0x04, 0x12, 0x21, 0x0a, 0x1d, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43,
	0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53,
	0x45, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x0a, 0x12, 0x24, 0x0a, 0x20, 0x57, 0x4f, 0x52, 0x4b, 0x53,
	0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x43,
	0x52, 0x45, 0x54, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x0b, 0x12, 0x22, 0x0a,
	0x1e, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x53, 0x54, 0x59, 0x4c, 0x45, 0x10,
	0x0c, 0x12, 0x28, 0x0a, 0x24, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53,
	0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x56,
	0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x0d, 0x42, 0xa6, 0x01, 0x0a, 0x0f,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x42,
	0x15, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x6c, 0x66, 0x68, 0x6f, 0x73,
	0x74, 0x65, 0x64, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x58, 0xaa,
	0x02, 0x0b, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0xca, 0x02, 0x0b,
	0x53, 0x6c, 0x61, 0x73, 0x68, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0xe2, 0x02, 0x17, 0x53, 0x6c,
	0x61, 0x73, 0x68, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x3a, 0x3a, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_workspace_setting_proto_rawDescOnce sync.Once
	file_store_workspace_setting_proto_rawDescData = file_store_workspace_setting_proto_rawDesc
)

func file_store_workspace_setting_proto_rawDescGZIP() []byte {
	file_store_workspace_setting_proto_rawDescOnce.Do(func() {
		file_store_workspace_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_workspace_setting_proto_rawDescData)
	})
	return file_store_workspace_setting_proto_rawDescData
}

var file_store_workspace_setting_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_store_workspace_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_store_workspace_setting_proto_goTypes = []any{
	(WorkspaceSettingKey)(0),                         // 0: slash.store.WorkspaceSettingKey
	(*WorkspaceSetting)(nil),                         // 1: slash.store.WorkspaceSetting
	(*WorkspaceSetting_GeneralSetting)(nil),          // 2: slash.store.WorkspaceSetting.GeneralSetting
	(*WorkspaceSetting_SecuritySetting)(nil),         // 3: slash.store.WorkspaceSetting.SecuritySetting
	(*WorkspaceSetting_ShortcutRelatedSetting)(nil),  // 4: slash.store.WorkspaceSetting.ShortcutRelatedSetting
	(*WorkspaceSetting_IdentityProviderSetting)(nil), // 5: slash.store.WorkspaceSetting.IdentityProviderSetting
	(Visibility)(0),                                  // 6: slash.store.Visibility
	(*IdentityProvider)(nil),                         // 7: slash.store.IdentityProvider
}
var file_store_workspace_setting_proto_depIdxs = []int32{
	0, // 0: slash.store.WorkspaceSetting.key:type_name -> slash.store.WorkspaceSettingKey
	2, // 1: slash.store.WorkspaceSetting.general:type_name -> slash.store.WorkspaceSetting.GeneralSetting
	3, // 2: slash.store.WorkspaceSetting.security:type_name -> slash.store.WorkspaceSetting.SecuritySetting
	4, // 3: slash.store.WorkspaceSetting.shortcut_related:type_name -> slash.store.WorkspaceSetting.ShortcutRelatedSetting
	5, // 4: slash.store.WorkspaceSetting.identity_provider:type_name -> slash.store.WorkspaceSetting.IdentityProviderSetting
	6, // 5: slash.store.WorkspaceSetting.ShortcutRelatedSetting.default_visibility:type_name -> slash.store.Visibility
	7, // 6: slash.store.WorkspaceSetting.IdentityProviderSetting.identity_providers:type_name -> slash.store.IdentityProvider
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_store_workspace_setting_proto_init() }
func file_store_workspace_setting_proto_init() {
	if File_store_workspace_setting_proto != nil {
		return
	}
	file_store_common_proto_init()
	file_store_idp_proto_init()
	file_store_workspace_setting_proto_msgTypes[0].OneofWrappers = []any{
		(*WorkspaceSetting_General)(nil),
		(*WorkspaceSetting_Security)(nil),
		(*WorkspaceSetting_ShortcutRelated)(nil),
		(*WorkspaceSetting_IdentityProvider)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_workspace_setting_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_workspace_setting_proto_goTypes,
		DependencyIndexes: file_store_workspace_setting_proto_depIdxs,
		EnumInfos:         file_store_workspace_setting_proto_enumTypes,
		MessageInfos:      file_store_workspace_setting_proto_msgTypes,
	}.Build()
	File_store_workspace_setting_proto = out.File
	file_store_workspace_setting_proto_rawDesc = nil
	file_store_workspace_setting_proto_goTypes = nil
	file_store_workspace_setting_proto_depIdxs = nil
}
