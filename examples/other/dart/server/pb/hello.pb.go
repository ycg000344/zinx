// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package pb

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

type MessageID int32

const (
	MessageID_NaN   MessageID = 0
	MessageID_First MessageID = 1
)

var MessageID_name = map[int32]string{
	0: "NaN",
	1: "First",
}
var MessageID_value = map[string]int32{
	"NaN":   0,
	"First": 1,
}

func (x MessageID) String() string {
	return proto.EnumName(MessageID_name, int32(x))
}
func (MessageID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_hello_33729d7bc163d53e, []int{0}
}

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_33729d7bc163d53e, []int{0}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (dst *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(dst, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Greeter              string   `protobuf:"bytes,1,opt,name=greeter,proto3" json:"greeter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_33729d7bc163d53e, []int{1}
}
func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (dst *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(dst, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetGreeter() string {
	if m != nil {
		return m.Greeter
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "pb.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "pb.HelloResponse")
	proto.RegisterEnum("pb.MessageID", MessageID_name, MessageID_value)
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_hello_33729d7bc163d53e) }

var fileDescriptor_hello_33729d7bc163d53e = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe2, 0xe2, 0xf1,
	0x00, 0x09, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6,
	0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x9a, 0x5c, 0xbc, 0x50, 0x35,
	0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x12, 0x5c, 0xec, 0xe9, 0x45, 0xa9, 0xa9, 0x25, 0xa9,
	0x45, 0x50, 0x75, 0x30, 0xae, 0x96, 0x3c, 0x17, 0xa7, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa,
	0xa7, 0x8b, 0x10, 0x3b, 0x17, 0xb3, 0x5f, 0xa2, 0x9f, 0x00, 0x83, 0x10, 0x27, 0x17, 0xab, 0x5b,
	0x66, 0x51, 0x71, 0x89, 0x00, 0x63, 0x12, 0x1b, 0xd8, 0x6a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf0, 0x48, 0x3c, 0x03, 0x89, 0x00, 0x00, 0x00,
}