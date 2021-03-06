// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dapr/proto/daprinternal/v1/status.proto

package v1

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

// Status represents the response status for HTTP and gRPC app channel.
type Status struct {
	// The status code
	//
	// This field is required.
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Error message
	//
	// This field is optional.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// A list of messages that carry the error details
	//
	// This field is optional.
	Details              []*any.Any `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_f93cffa19bfeff08, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Status) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "dapr.proto.daprinternal.v1.Status")
}

func init() {
	proto.RegisterFile("dapr/proto/daprinternal/v1/status.proto", fileDescriptor_f93cffa19bfeff08)
}

var fileDescriptor_f93cffa19bfeff08 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xb1, 0xce, 0x82, 0x30,
	0x1c, 0xc4, 0xc3, 0xc7, 0x27, 0xc4, 0xba, 0x35, 0x0e, 0x95, 0x89, 0xb8, 0xc8, 0xf4, 0xaf, 0xe8,
	0x13, 0xe8, 0x23, 0xe0, 0xe6, 0x56, 0xa0, 0x54, 0x22, 0xb4, 0x84, 0x16, 0x12, 0xde, 0xde, 0xd0,
	0xa6, 0x89, 0x83, 0x4b, 0x73, 0xfd, 0xdf, 0x2f, 0x77, 0x39, 0x74, 0xaa, 0xd9, 0x30, 0xd2, 0x61,
	0x54, 0x46, 0xd1, 0x55, 0xb6, 0xd2, 0xf0, 0x51, 0xb2, 0x8e, 0xce, 0x39, 0xd5, 0x86, 0x99, 0x49,
	0x83, 0x35, 0x71, 0xb2, 0xba, 0x4e, 0xc3, 0x37, 0x08, 0x73, 0x9e, 0x1c, 0x84, 0x52, 0xa2, 0xe3,
	0x2e, 0xa6, 0x9c, 0x1a, 0xca, 0xe4, 0xe2, 0xd0, 0x63, 0x83, 0xa2, 0x87, 0x8d, 0xc1, 0x18, 0xfd,
	0x57, 0xaa, 0xe6, 0x24, 0x48, 0x83, 0x6c, 0x53, 0x58, 0x8d, 0x09, 0x8a, 0x7b, 0xae, 0x35, 0x13,
	0x9c, 0xfc, 0xa5, 0x41, 0xb6, 0x2d, 0xfc, 0x17, 0x03, 0x8a, 0x6b, 0x6e, 0x58, 0xdb, 0x69, 0x12,
	0xa6, 0x61, 0xb6, 0xbb, 0xec, 0xc1, 0x95, 0x80, 0x2f, 0x81, 0x9b, 0x5c, 0x0a, 0x0f, 0xdd, 0xcf,
	0x4f, 0x10, 0xad, 0x79, 0x4d, 0x25, 0x54, 0xaa, 0xb7, 0x4b, 0xdc, 0x33, 0xbc, 0xc5, 0xef, 0x75,
	0x65, 0x64, 0xcf, 0xd7, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0x71, 0x10, 0x2c, 0x02, 0x01,
	0x00, 0x00,
}
