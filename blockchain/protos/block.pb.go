// Code generated by protoc-gen-go. DO NOT EDIT.
// source: block.proto

/*
Package block is a generated protocol buffer package.

It is generated from these files:
	block.proto

It has these top-level messages:
	Block
	BlockHeader
	BlockData
*/
package block

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Block struct {
	Header *BlockHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Data   *BlockData   `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetData() *BlockData {
	if m != nil {
		return m.Data
	}
	return nil
}

type BlockHeader struct {
	Version       uint64 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Timestamp     int64  `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	PrevBlockHash []byte `protobuf:"bytes,3,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`
	Hash          []byte `protobuf:"bytes,4,opt,name=Hash,proto3" json:"Hash,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BlockHeader) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetPrevBlockHash() []byte {
	if m != nil {
		return m.PrevBlockHash
	}
	return nil
}

func (m *BlockHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type BlockData struct {
	Data [][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (m *BlockData) Reset()                    { *m = BlockData{} }
func (m *BlockData) String() string            { return proto.CompactTextString(m) }
func (*BlockData) ProtoMessage()               {}
func (*BlockData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BlockData) GetData() [][]byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Block)(nil), "block.Block")
	proto.RegisterType((*BlockHeader)(nil), "block.BlockHeader")
	proto.RegisterType((*BlockData)(nil), "block.BlockData")
}

func init() { proto.RegisterFile("block.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x3f, 0xcf, 0x82, 0x30,
	0x10, 0x87, 0xd3, 0x97, 0xc2, 0x1b, 0x0e, 0x4d, 0xcc, 0x4d, 0x1d, 0x4c, 0x24, 0x84, 0x81, 0x38,
	0x30, 0xe8, 0x37, 0x30, 0x0e, 0xce, 0xdd, 0x1c, 0x8b, 0x34, 0x81, 0x28, 0x96, 0x94, 0x86, 0xd5,
	0xaf, 0x6e, 0x7a, 0xf8, 0x07, 0xb7, 0xdf, 0xdd, 0xf3, 0xf4, 0x7a, 0x39, 0x48, 0xaa, 0x9b, 0xb9,
	0x5c, 0xcb, 0xde, 0x1a, 0x67, 0x30, 0xa4, 0x22, 0x3b, 0x43, 0x78, 0xf0, 0x01, 0xb7, 0x10, 0x35,
	0x5a, 0xd5, 0xda, 0x0a, 0x96, 0xb2, 0x22, 0xd9, 0x61, 0x39, 0xd9, 0x44, 0x4f, 0x44, 0xe4, 0xcb,
	0xc0, 0x1c, 0x78, 0xad, 0x9c, 0x12, 0x7f, 0x64, 0xae, 0xe6, 0xe6, 0x51, 0x39, 0x25, 0x89, 0x66,
	0x0f, 0x48, 0x66, 0x8f, 0x51, 0xc0, 0xff, 0xa8, 0xed, 0xd0, 0x9a, 0x3b, 0xfd, 0xc0, 0xe5, 0xbb,
	0xc4, 0x35, 0xc4, 0xae, 0xed, 0xf4, 0xe0, 0x54, 0xd7, 0xd3, 0xcc, 0x40, 0x7e, 0x1b, 0x98, 0xc3,
	0xb2, 0xb7, 0x7a, 0x9c, 0x46, 0xa9, 0xa1, 0x11, 0x41, 0xca, 0x8a, 0x85, 0xfc, 0x6d, 0x22, 0x02,
	0x27, 0xc8, 0x09, 0x52, 0xce, 0x36, 0x10, 0x7f, 0x76, 0xf2, 0x02, 0xed, 0xcc, 0xd2, 0xc0, 0x0b,
	0x3e, 0x57, 0x11, 0x9d, 0x62, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x00, 0x4a, 0x70, 0xc6, 0x19,
	0x01, 0x00, 0x00,
}