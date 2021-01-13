// Copyright (c) 2018-2020 MobileCoin Inc.

// Blockchain-related data types.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: blockchain.proto

package block

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Block ID.
type BlockID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BlockID) Reset() {
	*x = BlockID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockID) ProtoMessage() {}

func (x *BlockID) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockID.ProtoReflect.Descriptor instead.
func (*BlockID) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{0}
}

func (x *BlockID) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Hash of the block's contents.
type BlockContentsHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BlockContentsHash) Reset() {
	*x = BlockContentsHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockContentsHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockContentsHash) ProtoMessage() {}

func (x *BlockContentsHash) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockContentsHash.ProtoReflect.Descriptor instead.
func (*BlockContentsHash) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{1}
}

func (x *BlockContentsHash) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// A block in the blockchain.
type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Block ID.
	Id *BlockID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Block format version.
	Version uint32 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	// Id of the previous block.
	ParentId *BlockID `protobuf:"bytes,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// The index of this block in the blockchain.
	Index uint64 `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	// The cumulative number of TXOs in the blockchain, including this block
	CumulativeTxoCount uint64 `protobuf:"varint,5,opt,name=cumulative_txo_count,json=cumulativeTxoCount,proto3" json:"cumulative_txo_count,omitempty"`
	// Root hash of the membership proofs provided by the untrusted local system for validation.
	// This captures the state of all TxOuts in the ledger that this block was validated against.
	RootElement *TxOutMembershipElement `protobuf:"bytes,6,opt,name=root_element,json=rootElement,proto3" json:"root_element,omitempty"`
	// Hash of the block's contents.
	ContentsHash *BlockContentsHash `protobuf:"bytes,7,opt,name=contents_hash,json=contentsHash,proto3" json:"contents_hash,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{2}
}

func (x *Block) GetId() *BlockID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Block) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Block) GetParentId() *BlockID {
	if x != nil {
		return x.ParentId
	}
	return nil
}

func (x *Block) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Block) GetCumulativeTxoCount() uint64 {
	if x != nil {
		return x.CumulativeTxoCount
	}
	return 0
}

func (x *Block) GetRootElement() *TxOutMembershipElement {
	if x != nil {
		return x.RootElement
	}
	return nil
}

func (x *Block) GetContentsHash() *BlockContentsHash {
	if x != nil {
		return x.ContentsHash
	}
	return nil
}

type BlockContents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key images spent in this block.
	KeyImages []*KeyImage `protobuf:"bytes,1,rep,name=key_images,json=keyImages,proto3" json:"key_images,omitempty"`
	// Outputs created in this block.
	Outputs []*TxOut `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *BlockContents) Reset() {
	*x = BlockContents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockContents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockContents) ProtoMessage() {}

func (x *BlockContents) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockContents.ProtoReflect.Descriptor instead.
func (*BlockContents) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{3}
}

func (x *BlockContents) GetKeyImages() []*KeyImage {
	if x != nil {
		return x.KeyImages
	}
	return nil
}

func (x *BlockContents) GetOutputs() []*TxOut {
	if x != nil {
		return x.Outputs
	}
	return nil
}

type BlockSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The signature of the block.
	Signature *Ed25519Signature `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// The signer that generated the above signature.
	Signer *Ed25519Public `protobuf:"bytes,2,opt,name=signer,proto3" json:"signer,omitempty"`
	// An approximate time in which the block was signed.
	// Represented as seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z.
	SignedAt uint64 `protobuf:"varint,3,opt,name=signed_at,json=signedAt,proto3" json:"signed_at,omitempty"`
}

func (x *BlockSignature) Reset() {
	*x = BlockSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockSignature) ProtoMessage() {}

func (x *BlockSignature) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockSignature.ProtoReflect.Descriptor instead.
func (*BlockSignature) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{4}
}

func (x *BlockSignature) GetSignature() *Ed25519Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *BlockSignature) GetSigner() *Ed25519Public {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *BlockSignature) GetSignedAt() uint64 {
	if x != nil {
		return x.SignedAt
	}
	return 0
}

// Version 1 of an archived block.
type ArchiveBlockV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Block
	Block *Block `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	// Contents of the block.
	BlockContents *BlockContents `protobuf:"bytes,2,opt,name=block_contents,json=blockContents,proto3" json:"block_contents,omitempty"`
	// Block signature, when available.
	Signature *BlockSignature `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ArchiveBlockV1) Reset() {
	*x = ArchiveBlockV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArchiveBlockV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArchiveBlockV1) ProtoMessage() {}

func (x *ArchiveBlockV1) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArchiveBlockV1.ProtoReflect.Descriptor instead.
func (*ArchiveBlockV1) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{5}
}

func (x *ArchiveBlockV1) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *ArchiveBlockV1) GetBlockContents() *BlockContents {
	if x != nil {
		return x.BlockContents
	}
	return nil
}

func (x *ArchiveBlockV1) GetSignature() *BlockSignature {
	if x != nil {
		return x.Signature
	}
	return nil
}

// An archived block.
type ArchiveBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Block:
	//	*ArchiveBlock_V1
	Block isArchiveBlock_Block `protobuf_oneof:"block"`
}

func (x *ArchiveBlock) Reset() {
	*x = ArchiveBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArchiveBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArchiveBlock) ProtoMessage() {}

func (x *ArchiveBlock) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArchiveBlock.ProtoReflect.Descriptor instead.
func (*ArchiveBlock) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{6}
}

func (m *ArchiveBlock) GetBlock() isArchiveBlock_Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (x *ArchiveBlock) GetV1() *ArchiveBlockV1 {
	if x, ok := x.GetBlock().(*ArchiveBlock_V1); ok {
		return x.V1
	}
	return nil
}

type isArchiveBlock_Block interface {
	isArchiveBlock_Block()
}

type ArchiveBlock_V1 struct {
	V1 *ArchiveBlockV1 `protobuf:"bytes,1,opt,name=v1,proto3,oneof"`
}

func (*ArchiveBlock_V1) isArchiveBlock_Block() {}

// A collection of archived blocks.
type ArchiveBlocks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*ArchiveBlock `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *ArchiveBlocks) Reset() {
	*x = ArchiveBlocks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArchiveBlocks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArchiveBlocks) ProtoMessage() {}

func (x *ArchiveBlocks) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArchiveBlocks.ProtoReflect.Descriptor instead.
func (*ArchiveBlocks) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{7}
}

func (x *ArchiveBlocks) GetBlocks() []*ArchiveBlock {
	if x != nil {
		return x.Blocks
	}
	return nil
}

var File_blockchain_proto protoreflect.FileDescriptor

var file_blockchain_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x1a, 0x0e,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d,
	0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x27, 0x0a,
	0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc9, 0x02, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x23, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49,
	0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x30, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x75, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x78, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x54, 0x78, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x0c, 0x72, 0x6f, 0x6f,
	0x74, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x54, 0x78, 0x4f, 0x75, 0x74,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0b, 0x72, 0x6f, 0x6f, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x42,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x48, 0x61, 0x73, 0x68, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x61,
	0x73, 0x68, 0x22, 0x6d, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x4b, 0x65, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x09, 0x6b, 0x65, 0x79,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x54, 0x78, 0x4f, 0x75, 0x74, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x73, 0x22, 0x98, 0x01, 0x0a, 0x0e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x45, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2f,
	0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x45, 0x64, 0x32, 0x35, 0x35, 0x31,
	0x39, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb5, 0x01, 0x0a,
	0x0e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x56, 0x31, 0x12,
	0x27, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x40, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0d, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0x45, 0x0a, 0x0c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2c, 0x0a, 0x02, 0x76, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x56, 0x31, 0x48, 0x00, 0x52, 0x02,
	0x76, 0x31, 0x42, 0x07, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x41, 0x0a, 0x0d, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x30, 0x0a, 0x06,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_proto_rawDescOnce sync.Once
	file_blockchain_proto_rawDescData = file_blockchain_proto_rawDesc
)

func file_blockchain_proto_rawDescGZIP() []byte {
	file_blockchain_proto_rawDescOnce.Do(func() {
		file_blockchain_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_proto_rawDescData)
	})
	return file_blockchain_proto_rawDescData
}

var file_blockchain_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_blockchain_proto_goTypes = []interface{}{
	(*BlockID)(nil),                // 0: blockchain.BlockID
	(*BlockContentsHash)(nil),      // 1: blockchain.BlockContentsHash
	(*Block)(nil),                  // 2: blockchain.Block
	(*BlockContents)(nil),          // 3: blockchain.BlockContents
	(*BlockSignature)(nil),         // 4: blockchain.BlockSignature
	(*ArchiveBlockV1)(nil),         // 5: blockchain.ArchiveBlockV1
	(*ArchiveBlock)(nil),           // 6: blockchain.ArchiveBlock
	(*ArchiveBlocks)(nil),          // 7: blockchain.ArchiveBlocks
	(*TxOutMembershipElement)(nil), // 8: external.TxOutMembershipElement
	(*KeyImage)(nil),               // 9: external.KeyImage
	(*TxOut)(nil),                  // 10: external.TxOut
	(*Ed25519Signature)(nil),       // 11: external.Ed25519Signature
	(*Ed25519Public)(nil),          // 12: external.Ed25519Public
}
var file_blockchain_proto_depIdxs = []int32{
	0,  // 0: blockchain.Block.id:type_name -> blockchain.BlockID
	0,  // 1: blockchain.Block.parent_id:type_name -> blockchain.BlockID
	8,  // 2: blockchain.Block.root_element:type_name -> external.TxOutMembershipElement
	1,  // 3: blockchain.Block.contents_hash:type_name -> blockchain.BlockContentsHash
	9,  // 4: blockchain.BlockContents.key_images:type_name -> external.KeyImage
	10, // 5: blockchain.BlockContents.outputs:type_name -> external.TxOut
	11, // 6: blockchain.BlockSignature.signature:type_name -> external.Ed25519Signature
	12, // 7: blockchain.BlockSignature.signer:type_name -> external.Ed25519Public
	2,  // 8: blockchain.ArchiveBlockV1.block:type_name -> blockchain.Block
	3,  // 9: blockchain.ArchiveBlockV1.block_contents:type_name -> blockchain.BlockContents
	4,  // 10: blockchain.ArchiveBlockV1.signature:type_name -> blockchain.BlockSignature
	5,  // 11: blockchain.ArchiveBlock.v1:type_name -> blockchain.ArchiveBlockV1
	6,  // 12: blockchain.ArchiveBlocks.blocks:type_name -> blockchain.ArchiveBlock
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_blockchain_proto_init() }
func file_blockchain_proto_init() {
	if File_blockchain_proto != nil {
		return
	}
	file_external_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_blockchain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockID); i {
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
		file_blockchain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockContentsHash); i {
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
		file_blockchain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_blockchain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockContents); i {
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
		file_blockchain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockSignature); i {
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
		file_blockchain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArchiveBlockV1); i {
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
		file_blockchain_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArchiveBlock); i {
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
		file_blockchain_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArchiveBlocks); i {
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
	file_blockchain_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*ArchiveBlock_V1)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blockchain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_proto_goTypes,
		DependencyIndexes: file_blockchain_proto_depIdxs,
		MessageInfos:      file_blockchain_proto_msgTypes,
	}.Build()
	File_blockchain_proto = out.File
	file_blockchain_proto_rawDesc = nil
	file_blockchain_proto_goTypes = nil
	file_blockchain_proto_depIdxs = nil
}
