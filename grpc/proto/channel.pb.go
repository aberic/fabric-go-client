// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/proto/channel.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ChannelCreate struct {
	ConfigID             string   `protobuf:"bytes,1,opt,name=configID,proto3" json:"configID,omitempty"`
	OrderOrgName         string   `protobuf:"bytes,2,opt,name=orderOrgName,proto3" json:"orderOrgName,omitempty"`
	OrgName              string   `protobuf:"bytes,3,opt,name=orgName,proto3" json:"orgName,omitempty"`
	OrgUser              string   `protobuf:"bytes,4,opt,name=orgUser,proto3" json:"orgUser,omitempty"`
	ChannelID            string   `protobuf:"bytes,5,opt,name=channelID,proto3" json:"channelID,omitempty"`
	ChannelConfigPath    string   `protobuf:"bytes,6,opt,name=channelConfigPath,proto3" json:"channelConfigPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelCreate) Reset()         { *m = ChannelCreate{} }
func (m *ChannelCreate) String() string { return proto.CompactTextString(m) }
func (*ChannelCreate) ProtoMessage()    {}
func (*ChannelCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_056b8b10491c9602, []int{0}
}

func (m *ChannelCreate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelCreate.Unmarshal(m, b)
}
func (m *ChannelCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelCreate.Marshal(b, m, deterministic)
}
func (m *ChannelCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelCreate.Merge(m, src)
}
func (m *ChannelCreate) XXX_Size() int {
	return xxx_messageInfo_ChannelCreate.Size(m)
}
func (m *ChannelCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelCreate.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelCreate proto.InternalMessageInfo

func (m *ChannelCreate) GetConfigID() string {
	if m != nil {
		return m.ConfigID
	}
	return ""
}

func (m *ChannelCreate) GetOrderOrgName() string {
	if m != nil {
		return m.OrderOrgName
	}
	return ""
}

func (m *ChannelCreate) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *ChannelCreate) GetOrgUser() string {
	if m != nil {
		return m.OrgUser
	}
	return ""
}

func (m *ChannelCreate) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *ChannelCreate) GetChannelConfigPath() string {
	if m != nil {
		return m.ChannelConfigPath
	}
	return ""
}

type ChannelJoin struct {
	ConfigID             string   `protobuf:"bytes,1,opt,name=configID,proto3" json:"configID,omitempty"`
	OrderName            string   `protobuf:"bytes,2,opt,name=orderName,proto3" json:"orderName,omitempty"`
	OrgName              string   `protobuf:"bytes,3,opt,name=orgName,proto3" json:"orgName,omitempty"`
	OrgUser              string   `protobuf:"bytes,4,opt,name=orgUser,proto3" json:"orgUser,omitempty"`
	ChannelID            string   `protobuf:"bytes,5,opt,name=channelID,proto3" json:"channelID,omitempty"`
	PeerUrl              string   `protobuf:"bytes,6,opt,name=peerUrl,proto3" json:"peerUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelJoin) Reset()         { *m = ChannelJoin{} }
func (m *ChannelJoin) String() string { return proto.CompactTextString(m) }
func (*ChannelJoin) ProtoMessage()    {}
func (*ChannelJoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_056b8b10491c9602, []int{1}
}

func (m *ChannelJoin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelJoin.Unmarshal(m, b)
}
func (m *ChannelJoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelJoin.Marshal(b, m, deterministic)
}
func (m *ChannelJoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelJoin.Merge(m, src)
}
func (m *ChannelJoin) XXX_Size() int {
	return xxx_messageInfo_ChannelJoin.Size(m)
}
func (m *ChannelJoin) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelJoin.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelJoin proto.InternalMessageInfo

func (m *ChannelJoin) GetConfigID() string {
	if m != nil {
		return m.ConfigID
	}
	return ""
}

func (m *ChannelJoin) GetOrderName() string {
	if m != nil {
		return m.OrderName
	}
	return ""
}

func (m *ChannelJoin) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *ChannelJoin) GetOrgUser() string {
	if m != nil {
		return m.OrgUser
	}
	return ""
}

func (m *ChannelJoin) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *ChannelJoin) GetPeerUrl() string {
	if m != nil {
		return m.PeerUrl
	}
	return ""
}

type ChannelList struct {
	ConfigID             string   `protobuf:"bytes,1,opt,name=configID,proto3" json:"configID,omitempty"`
	OrgName              string   `protobuf:"bytes,2,opt,name=orgName,proto3" json:"orgName,omitempty"`
	OrgUser              string   `protobuf:"bytes,3,opt,name=orgUser,proto3" json:"orgUser,omitempty"`
	PeerName             string   `protobuf:"bytes,4,opt,name=peerName,proto3" json:"peerName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelList) Reset()         { *m = ChannelList{} }
func (m *ChannelList) String() string { return proto.CompactTextString(m) }
func (*ChannelList) ProtoMessage()    {}
func (*ChannelList) Descriptor() ([]byte, []int) {
	return fileDescriptor_056b8b10491c9602, []int{2}
}

func (m *ChannelList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelList.Unmarshal(m, b)
}
func (m *ChannelList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelList.Marshal(b, m, deterministic)
}
func (m *ChannelList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelList.Merge(m, src)
}
func (m *ChannelList) XXX_Size() int {
	return xxx_messageInfo_ChannelList.Size(m)
}
func (m *ChannelList) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelList.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelList proto.InternalMessageInfo

func (m *ChannelList) GetConfigID() string {
	if m != nil {
		return m.ConfigID
	}
	return ""
}

func (m *ChannelList) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *ChannelList) GetOrgUser() string {
	if m != nil {
		return m.OrgUser
	}
	return ""
}

func (m *ChannelList) GetPeerName() string {
	if m != nil {
		return m.PeerName
	}
	return ""
}

func init() {
	proto.RegisterType((*ChannelCreate)(nil), "proto.ChannelCreate")
	proto.RegisterType((*ChannelJoin)(nil), "proto.ChannelJoin")
	proto.RegisterType((*ChannelList)(nil), "proto.ChannelList")
}

func init() { proto.RegisterFile("grpc/proto/channel.proto", fileDescriptor_056b8b10491c9602) }

var fileDescriptor_056b8b10491c9602 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x2f, 0x2a, 0x48,
	0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0xce, 0x48, 0xcc, 0xcb, 0x4b, 0xcd, 0xd1, 0x03,
	0xf3, 0x84, 0x58, 0xc1, 0x94, 0xd2, 0x79, 0x46, 0x2e, 0x5e, 0x67, 0x88, 0x84, 0x73, 0x51, 0x6a,
	0x62, 0x49, 0xaa, 0x90, 0x14, 0x17, 0x47, 0x72, 0x7e, 0x5e, 0x5a, 0x66, 0xba, 0xa7, 0x8b, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0x2f, 0xa4, 0xc4, 0xc5, 0x93, 0x5f, 0x94, 0x92, 0x5a,
	0xe4, 0x5f, 0x94, 0xee, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0x04, 0x96, 0x47, 0x11, 0x13, 0x92, 0xe0,
	0x62, 0xcf, 0x87, 0x4a, 0x33, 0x83, 0xa5, 0x61, 0x5c, 0xa8, 0x4c, 0x68, 0x71, 0x6a, 0x91, 0x04,
	0x0b, 0x5c, 0x06, 0xc4, 0x15, 0x92, 0xe1, 0xe2, 0x84, 0xba, 0xce, 0xd3, 0x45, 0x82, 0x15, 0x2c,
	0x87, 0x10, 0x10, 0xd2, 0xe1, 0x12, 0x84, 0x72, 0x9c, 0xc1, 0x0e, 0x09, 0x48, 0x2c, 0xc9, 0x90,
	0x60, 0x03, 0xab, 0xc2, 0x94, 0x50, 0xda, 0xcc, 0xc8, 0xc5, 0x0d, 0xf5, 0x91, 0x57, 0x7e, 0x66,
	0x1e, 0x5e, 0xff, 0xc8, 0x70, 0x71, 0x82, 0xdd, 0x8e, 0xe4, 0x19, 0x84, 0x00, 0x0d, 0x7c, 0x22,
	0xc1, 0xc5, 0x5e, 0x90, 0x9a, 0x5a, 0x14, 0x5a, 0x94, 0x03, 0x75, 0x3f, 0x8c, 0xab, 0x54, 0x09,
	0x77, 0xb4, 0x4f, 0x66, 0x71, 0x09, 0x5e, 0x47, 0x23, 0x39, 0x8b, 0x09, 0xa7, 0xb3, 0x98, 0x51,
	0x9d, 0x25, 0xc5, 0xc5, 0x01, 0xb2, 0x09, 0xac, 0x09, 0xe2, 0x62, 0x38, 0x3f, 0x89, 0x0d, 0x9c,
	0x12, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xba, 0x6b, 0x12, 0x2c, 0x02, 0x00, 0x00,
}