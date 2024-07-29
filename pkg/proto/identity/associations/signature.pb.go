// Signing methods for identity associations

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: identity/associations/signature.proto

package associations

import (
	message_contents "github.com/xmtp/xmtpd/pkg/proto/message_contents"
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

// RecoverableEcdsaSignature for EIP-191 and V2 signatures
type RecoverableEcdsaSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 65-bytes [ R || S || V ], with recovery id as the last byte
	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *RecoverableEcdsaSignature) Reset() {
	*x = RecoverableEcdsaSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_associations_signature_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecoverableEcdsaSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecoverableEcdsaSignature) ProtoMessage() {}

func (x *RecoverableEcdsaSignature) ProtoReflect() protoreflect.Message {
	mi := &file_identity_associations_signature_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecoverableEcdsaSignature.ProtoReflect.Descriptor instead.
func (*RecoverableEcdsaSignature) Descriptor() ([]byte, []int) {
	return file_identity_associations_signature_proto_rawDescGZIP(), []int{0}
}

func (x *RecoverableEcdsaSignature) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

// EdDSA signature for 25519
type RecoverableEd25519Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 64 bytes [R(32 bytes) || S(32 bytes)]
	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	// 32 bytes
	PublicKey []byte `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *RecoverableEd25519Signature) Reset() {
	*x = RecoverableEd25519Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_associations_signature_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecoverableEd25519Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecoverableEd25519Signature) ProtoMessage() {}

func (x *RecoverableEd25519Signature) ProtoReflect() protoreflect.Message {
	mi := &file_identity_associations_signature_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecoverableEd25519Signature.ProtoReflect.Descriptor instead.
func (*RecoverableEd25519Signature) Descriptor() ([]byte, []int) {
	return file_identity_associations_signature_proto_rawDescGZIP(), []int{1}
}

func (x *RecoverableEd25519Signature) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *RecoverableEd25519Signature) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

// Smart wallet signature
type Erc1271Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CAIP-10
	// https://github.com/ChainAgnostic/CAIPs/blob/main/CAIPs/caip-10.md
	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// Specify the block number to verify the signature against
	BlockNumber uint64 `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	// The actual signature bytes
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Erc1271Signature) Reset() {
	*x = Erc1271Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_associations_signature_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Erc1271Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Erc1271Signature) ProtoMessage() {}

func (x *Erc1271Signature) ProtoReflect() protoreflect.Message {
	mi := &file_identity_associations_signature_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Erc1271Signature.ProtoReflect.Descriptor instead.
func (*Erc1271Signature) Descriptor() ([]byte, []int) {
	return file_identity_associations_signature_proto_rawDescGZIP(), []int{2}
}

func (x *Erc1271Signature) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Erc1271Signature) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *Erc1271Signature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

// An existing address on xmtpv2 may have already signed a legacy identity key
// of type SignedPublicKey via the 'Create Identity' signature.
// For migration to xmtpv3, the legacy key is permitted to sign on behalf of the
// address to create a matching xmtpv3 installation key.
// This signature type can ONLY be used for CreateXid and AddAssociation
// payloads, and can only be used once in xmtpv3.
type LegacyDelegatedSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DelegatedKey *message_contents.SignedPublicKey `protobuf:"bytes,1,opt,name=delegated_key,json=delegatedKey,proto3" json:"delegated_key,omitempty"`
	Signature    *RecoverableEcdsaSignature        `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *LegacyDelegatedSignature) Reset() {
	*x = LegacyDelegatedSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_associations_signature_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LegacyDelegatedSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LegacyDelegatedSignature) ProtoMessage() {}

func (x *LegacyDelegatedSignature) ProtoReflect() protoreflect.Message {
	mi := &file_identity_associations_signature_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LegacyDelegatedSignature.ProtoReflect.Descriptor instead.
func (*LegacyDelegatedSignature) Descriptor() ([]byte, []int) {
	return file_identity_associations_signature_proto_rawDescGZIP(), []int{3}
}

func (x *LegacyDelegatedSignature) GetDelegatedKey() *message_contents.SignedPublicKey {
	if x != nil {
		return x.DelegatedKey
	}
	return nil
}

func (x *LegacyDelegatedSignature) GetSignature() *RecoverableEcdsaSignature {
	if x != nil {
		return x.Signature
	}
	return nil
}

// A wrapper for all possible signature types
type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Must have two properties:
	//  1. An identifier (address or public key) for the signer must either be
	//     recoverable, or specified as a field.
	//  2. The signer certifies that the signing payload is correct. The payload
	//     must be inferred from the context in which the signature is provided.
	//
	// Types that are assignable to Signature:
	//
	//	*Signature_Erc_191
	//	*Signature_Erc_1271
	//	*Signature_InstallationKey
	//	*Signature_DelegatedErc_191
	Signature isSignature_Signature `protobuf_oneof:"signature"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_associations_signature_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_identity_associations_signature_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_identity_associations_signature_proto_rawDescGZIP(), []int{4}
}

func (m *Signature) GetSignature() isSignature_Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (x *Signature) GetErc_191() *RecoverableEcdsaSignature {
	if x, ok := x.GetSignature().(*Signature_Erc_191); ok {
		return x.Erc_191
	}
	return nil
}

func (x *Signature) GetErc_1271() *Erc1271Signature {
	if x, ok := x.GetSignature().(*Signature_Erc_1271); ok {
		return x.Erc_1271
	}
	return nil
}

func (x *Signature) GetInstallationKey() *RecoverableEd25519Signature {
	if x, ok := x.GetSignature().(*Signature_InstallationKey); ok {
		return x.InstallationKey
	}
	return nil
}

func (x *Signature) GetDelegatedErc_191() *LegacyDelegatedSignature {
	if x, ok := x.GetSignature().(*Signature_DelegatedErc_191); ok {
		return x.DelegatedErc_191
	}
	return nil
}

type isSignature_Signature interface {
	isSignature_Signature()
}

type Signature_Erc_191 struct {
	Erc_191 *RecoverableEcdsaSignature `protobuf:"bytes,1,opt,name=erc_191,json=erc191,proto3,oneof"`
}

type Signature_Erc_1271 struct {
	Erc_1271 *Erc1271Signature `protobuf:"bytes,2,opt,name=erc_1271,json=erc1271,proto3,oneof"`
}

type Signature_InstallationKey struct {
	InstallationKey *RecoverableEd25519Signature `protobuf:"bytes,3,opt,name=installation_key,json=installationKey,proto3,oneof"`
}

type Signature_DelegatedErc_191 struct {
	DelegatedErc_191 *LegacyDelegatedSignature `protobuf:"bytes,4,opt,name=delegated_erc_191,json=delegatedErc191,proto3,oneof"`
}

func (*Signature_Erc_191) isSignature_Signature() {}

func (*Signature_Erc_1271) isSignature_Signature() {}

func (*Signature_InstallationKey) isSignature_Signature() {}

func (*Signature_DelegatedErc_191) isSignature_Signature() {}

var File_identity_associations_signature_proto protoreflect.FileDescriptor

var file_identity_associations_signature_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x73, 0x73, 0x6f, 0x63,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x21, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x19, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x63, 0x64, 0x73, 0x61, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0x52, 0x0a, 0x1b, 0x52, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x72, 0x0a,
	0x10, 0x45, 0x72, 0x63, 0x31, 0x32, 0x37, 0x31, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x22, 0xbc, 0x01, 0x0a, 0x18, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x44, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x4b,
	0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x0c, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x53, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x63, 0x64, 0x73, 0x61, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x22, 0xff, 0x02, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x50,
	0x0a, 0x07, 0x65, 0x72, 0x63, 0x5f, 0x31, 0x39, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x63, 0x64, 0x73, 0x61, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x06, 0x65, 0x72, 0x63, 0x31, 0x39, 0x31,
	0x12, 0x49, 0x0a, 0x08, 0x65, 0x72, 0x63, 0x5f, 0x31, 0x32, 0x37, 0x31, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x45, 0x72, 0x63, 0x31, 0x32, 0x37, 0x31, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x48, 0x00, 0x52, 0x07, 0x65, 0x72, 0x63, 0x31, 0x32, 0x37, 0x31, 0x12, 0x64, 0x0a, 0x10, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x64,
	0x32, 0x35, 0x35, 0x31, 0x39, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x48, 0x00,
	0x52, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65,
	0x79, 0x12, 0x62, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65,
	0x72, 0x63, 0x5f, 0x31, 0x39, 0x31, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x78,
	0x6d, 0x74, 0x70, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x73, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79,
	0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x45,
	0x72, 0x63, 0x31, 0x39, 0x31, 0x42, 0x0b, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x42, 0xf1, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x6d, 0x74, 0x70, 0x2f, 0x78, 0x6d, 0x74, 0x70, 0x64, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xa2, 0x02,
	0x03, 0x58, 0x49, 0x41, 0xaa, 0x02, 0x1a, 0x58, 0x6d, 0x74, 0x70, 0x2e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0xca, 0x02, 0x1a, 0x58, 0x6d, 0x74, 0x70, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5c, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xe2, 0x02,
	0x26, 0x58, 0x6d, 0x74, 0x70, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c, 0x41,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x58, 0x6d, 0x74, 0x70, 0x3a, 0x3a,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_identity_associations_signature_proto_rawDescOnce sync.Once
	file_identity_associations_signature_proto_rawDescData = file_identity_associations_signature_proto_rawDesc
)

func file_identity_associations_signature_proto_rawDescGZIP() []byte {
	file_identity_associations_signature_proto_rawDescOnce.Do(func() {
		file_identity_associations_signature_proto_rawDescData = protoimpl.X.CompressGZIP(file_identity_associations_signature_proto_rawDescData)
	})
	return file_identity_associations_signature_proto_rawDescData
}

var file_identity_associations_signature_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_identity_associations_signature_proto_goTypes = []any{
	(*RecoverableEcdsaSignature)(nil),        // 0: xmtp.identity.associations.RecoverableEcdsaSignature
	(*RecoverableEd25519Signature)(nil),      // 1: xmtp.identity.associations.RecoverableEd25519Signature
	(*Erc1271Signature)(nil),                 // 2: xmtp.identity.associations.Erc1271Signature
	(*LegacyDelegatedSignature)(nil),         // 3: xmtp.identity.associations.LegacyDelegatedSignature
	(*Signature)(nil),                        // 4: xmtp.identity.associations.Signature
	(*message_contents.SignedPublicKey)(nil), // 5: xmtp.message_contents.SignedPublicKey
}
var file_identity_associations_signature_proto_depIdxs = []int32{
	5, // 0: xmtp.identity.associations.LegacyDelegatedSignature.delegated_key:type_name -> xmtp.message_contents.SignedPublicKey
	0, // 1: xmtp.identity.associations.LegacyDelegatedSignature.signature:type_name -> xmtp.identity.associations.RecoverableEcdsaSignature
	0, // 2: xmtp.identity.associations.Signature.erc_191:type_name -> xmtp.identity.associations.RecoverableEcdsaSignature
	2, // 3: xmtp.identity.associations.Signature.erc_1271:type_name -> xmtp.identity.associations.Erc1271Signature
	1, // 4: xmtp.identity.associations.Signature.installation_key:type_name -> xmtp.identity.associations.RecoverableEd25519Signature
	3, // 5: xmtp.identity.associations.Signature.delegated_erc_191:type_name -> xmtp.identity.associations.LegacyDelegatedSignature
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_identity_associations_signature_proto_init() }
func file_identity_associations_signature_proto_init() {
	if File_identity_associations_signature_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_identity_associations_signature_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RecoverableEcdsaSignature); i {
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
		file_identity_associations_signature_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RecoverableEd25519Signature); i {
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
		file_identity_associations_signature_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Erc1271Signature); i {
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
		file_identity_associations_signature_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*LegacyDelegatedSignature); i {
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
		file_identity_associations_signature_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Signature); i {
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
	file_identity_associations_signature_proto_msgTypes[4].OneofWrappers = []any{
		(*Signature_Erc_191)(nil),
		(*Signature_Erc_1271)(nil),
		(*Signature_InstallationKey)(nil),
		(*Signature_DelegatedErc_191)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_identity_associations_signature_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_identity_associations_signature_proto_goTypes,
		DependencyIndexes: file_identity_associations_signature_proto_depIdxs,
		MessageInfos:      file_identity_associations_signature_proto_msgTypes,
	}.Build()
	File_identity_associations_signature_proto = out.File
	file_identity_associations_signature_proto_rawDesc = nil
	file_identity_associations_signature_proto_goTypes = nil
	file_identity_associations_signature_proto_depIdxs = nil
}
