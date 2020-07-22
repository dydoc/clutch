// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/service/aws/v1/aws.proto

package awsv1

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Config struct {
	Regions              []string `protobuf:"bytes,1,rep,name=regions,proto3" json:"regions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_645f7f145f69443e, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "clutch.config.service.aws.v1.Config")
}

func init() {
	proto.RegisterFile("config/service/aws/v1/aws.proto", fileDescriptor_645f7f145f69443e)
}

var fileDescriptor_645f7f145f69443e = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0x2c, 0x2f, 0xd6, 0x2f,
	0x33, 0x04, 0x51, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x32, 0xc9, 0x39, 0xa5, 0x25, 0xc9,
	0x19, 0x7a, 0x10, 0x75, 0x7a, 0x50, 0x75, 0x7a, 0x20, 0x05, 0x65, 0x86, 0x52, 0xe2, 0x65, 0x89,
	0x39, 0x99, 0x29, 0x89, 0x25, 0xa9, 0xfa, 0x30, 0x06, 0x44, 0x9b, 0x92, 0x0e, 0x17, 0x9b, 0x33,
	0x58, 0x87, 0x90, 0x12, 0x17, 0x7b, 0x51, 0x6a, 0x7a, 0x66, 0x7e, 0x5e, 0xb1, 0x04, 0xa3, 0x02,
	0xb3, 0x06, 0xa7, 0x13, 0xc7, 0x2f, 0x27, 0xd6, 0x49, 0x8c, 0x4c, 0x1c, 0x8c, 0x41, 0x30, 0x09,
	0x27, 0xf6, 0x28, 0xd6, 0xc4, 0xf2, 0xe2, 0x32, 0xc3, 0x24, 0x36, 0xb0, 0x6e, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x41, 0x09, 0xa9, 0xe5, 0x97, 0x00, 0x00, 0x00,
}