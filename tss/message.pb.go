// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protob/message.proto

package tss

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//
// Wrapper for TSS messages, often read by the transport layer and not itself sent over the wire
type MessageWrapper struct {
	// Metadata optionally un-marshalled and used by the transport to route this message.
	IsBroadcast bool `protobuf:"varint,1,opt,name=is_broadcast,json=isBroadcast,proto3" json:"is_broadcast,omitempty"`
	// Metadata optionally un-marshalled and used by the transport to route this message.
	IsToOldCommittee bool `protobuf:"varint,2,opt,name=is_to_old_committee,json=isToOldCommittee,proto3" json:"is_to_old_committee,omitempty"`
	// Metadata optionally un-marshalled and used by the transport to route this message.
	IsToOldAndNewCommittees bool `protobuf:"varint,5,opt,name=is_to_old_and_new_committees,json=isToOldAndNewCommittees,proto3" json:"is_to_old_and_new_committees,omitempty"`
	// Metadata optionally un-marshalled and used by the transport to route this message.
	From *MessageWrapper_PartyID `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	// Metadata optionally un-marshalled and used by the transport to route this message.
	To []*MessageWrapper_PartyID `protobuf:"bytes,4,rep,name=to,proto3" json:"to,omitempty"`
	// This field is actually what is sent through the wire and consumed on the other end by UpdateFromBytes.
	// An Any contains an arbitrary serialized message as bytes, along with a URL that
	// acts as a globally unique identifier for and resolves to that message's type.
	Message              *any.Any `protobuf:"bytes,10,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageWrapper) Reset()         { *m = MessageWrapper{} }
func (m *MessageWrapper) String() string { return proto.CompactTextString(m) }
func (*MessageWrapper) ProtoMessage()    {}
func (*MessageWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_5be430ad0e7f3d12, []int{0}
}

func (m *MessageWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWrapper.Unmarshal(m, b)
}
func (m *MessageWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWrapper.Marshal(b, m, deterministic)
}
func (m *MessageWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWrapper.Merge(m, src)
}
func (m *MessageWrapper) XXX_Size() int {
	return xxx_messageInfo_MessageWrapper.Size(m)
}
func (m *MessageWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWrapper proto.InternalMessageInfo

func (m *MessageWrapper) GetIsBroadcast() bool {
	if m != nil {
		return m.IsBroadcast
	}
	return false
}

func (m *MessageWrapper) GetIsToOldCommittee() bool {
	if m != nil {
		return m.IsToOldCommittee
	}
	return false
}

func (m *MessageWrapper) GetIsToOldAndNewCommittees() bool {
	if m != nil {
		return m.IsToOldAndNewCommittees
	}
	return false
}

func (m *MessageWrapper) GetFrom() *MessageWrapper_PartyID {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *MessageWrapper) GetTo() []*MessageWrapper_PartyID {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *MessageWrapper) GetMessage() *any.Any {
	if m != nil {
		return m.Message
	}
	return nil
}

// PartyID represents a participant in the TSS protocol rounds.
// Note: The `id` and `moniker` are provided for convenience to allow you to track participants easier.
// The `id` is intended to be a unique string representation of `key` and `moniker` can be anything (even left blank).
type MessageWrapper_PartyID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Moniker              string   `protobuf:"bytes,2,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Key                  []byte   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageWrapper_PartyID) Reset()         { *m = MessageWrapper_PartyID{} }
func (m *MessageWrapper_PartyID) String() string { return proto.CompactTextString(m) }
func (*MessageWrapper_PartyID) ProtoMessage()    {}
func (*MessageWrapper_PartyID) Descriptor() ([]byte, []int) {
	return fileDescriptor_5be430ad0e7f3d12, []int{0, 0}
}

func (m *MessageWrapper_PartyID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWrapper_PartyID.Unmarshal(m, b)
}
func (m *MessageWrapper_PartyID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWrapper_PartyID.Marshal(b, m, deterministic)
}
func (m *MessageWrapper_PartyID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWrapper_PartyID.Merge(m, src)
}
func (m *MessageWrapper_PartyID) XXX_Size() int {
	return xxx_messageInfo_MessageWrapper_PartyID.Size(m)
}
func (m *MessageWrapper_PartyID) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWrapper_PartyID.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWrapper_PartyID proto.InternalMessageInfo

func (m *MessageWrapper_PartyID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MessageWrapper_PartyID) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *MessageWrapper_PartyID) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageWrapper)(nil), "MessageWrapper")
	proto.RegisterType((*MessageWrapper_PartyID)(nil), "MessageWrapper.PartyID")
}

func init() { proto.RegisterFile("protob/message.proto", fileDescriptor_5be430ad0e7f3d12) }

var fileDescriptor_5be430ad0e7f3d12 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0x87, 0x69, 0xda, 0xbe, 0xfb, 0x76, 0x5a, 0x4a, 0x89, 0x85, 0xc6, 0xe2, 0xa1, 0x7a, 0xb1,
	0x20, 0x66, 0x41, 0xcf, 0x1e, 0x5a, 0xf5, 0xe0, 0xc1, 0x3f, 0x04, 0x41, 0xf0, 0xb2, 0xa4, 0x4d,
	0x5a, 0x42, 0xbb, 0x99, 0x92, 0x44, 0xca, 0x7e, 0x69, 0x3f, 0x83, 0x98, 0xdd, 0xb5, 0x78, 0xf1,
	0x96, 0x99, 0x79, 0x7e, 0x99, 0xe1, 0x81, 0xe1, 0xce, 0x61, 0xc0, 0x45, 0x9a, 0x6b, 0xef, 0xe5,
	0x5a, 0xf3, 0x58, 0x8e, 0x8f, 0xd7, 0x88, 0xeb, 0xad, 0x4e, 0xcb, 0xe1, 0xc7, 0x2a, 0x95, 0xb6,
	0x28, 0x47, 0x67, 0x9f, 0x04, 0xfa, 0x8f, 0x25, 0xfc, 0xe6, 0xe4, 0x6e, 0xa7, 0x1d, 0x3d, 0x85,
	0x9e, 0xf1, 0xd9, 0xc2, 0xa1, 0x54, 0x4b, 0xe9, 0x03, 0x6b, 0x4c, 0x1a, 0xd3, 0xff, 0xa2, 0x6b,
	0xfc, 0xbc, 0x6e, 0xd1, 0x4b, 0x38, 0x32, 0x3e, 0x0b, 0x98, 0xe1, 0x56, 0x65, 0x4b, 0xcc, 0x73,
	0x13, 0x82, 0xd6, 0x8c, 0x44, 0x72, 0x60, 0xfc, 0x2b, 0x3e, 0x6f, 0xd5, 0x6d, 0xdd, 0xa7, 0x37,
	0x70, 0x72, 0xc0, 0xa5, 0x55, 0x99, 0xd5, 0xfb, 0x43, 0xcc, 0xb3, 0x76, 0xcc, 0x8d, 0xaa, 0xdc,
	0xcc, 0xaa, 0x27, 0xbd, 0xff, 0x49, 0x7b, 0x7a, 0x01, 0xad, 0x95, 0xc3, 0x9c, 0x35, 0x27, 0x8d,
	0x69, 0xf7, 0x6a, 0xc4, 0x7f, 0xdf, 0xcb, 0x5f, 0xa4, 0x0b, 0xc5, 0xc3, 0x9d, 0x88, 0x10, 0x3d,
	0x07, 0x12, 0x90, 0xb5, 0x26, 0xcd, 0xbf, 0x50, 0x12, 0x90, 0x72, 0x48, 0x2a, 0x4b, 0x0c, 0xe2,
	0xc7, 0x43, 0x5e, 0x6a, 0xe2, 0xb5, 0x26, 0x3e, 0xb3, 0x85, 0xa8, 0xa1, 0xf1, 0x3d, 0x24, 0x55,
	0x9c, 0xf6, 0x81, 0x18, 0x15, 0xbd, 0x74, 0x04, 0x31, 0x8a, 0x32, 0x48, 0x72, 0xb4, 0x66, 0xa3,
	0x5d, 0x54, 0xd0, 0x11, 0x75, 0x49, 0x07, 0xd0, 0xdc, 0xe8, 0x22, 0x5e, 0xde, 0x13, 0xdf, 0xcf,
	0x79, 0xf2, 0xde, 0xe6, 0x69, 0xf0, 0x7e, 0xf1, 0x2f, 0xae, 0xb9, 0xfe, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0xac, 0xdd, 0x4e, 0x90, 0xb3, 0x01, 0x00, 0x00,
}
