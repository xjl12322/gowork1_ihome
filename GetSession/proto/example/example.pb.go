// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package go_micro_srv_GetSession

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Sessionid            string   `protobuf:"bytes,1,opt,name=Sessionid,proto3" json:"Sessionid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSessionid() string {
	if m != nil {
		return m.Sessionid
	}
	return ""
}

type Response struct {
	Errno  string `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	Errmsg string `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	//返回用户名
	UserName             string   `protobuf:"bytes,3,opt,name=User_name,json=UserName,proto3" json:"User_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.GetSession.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.GetSession.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.GetSession.Response")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x3f, 0x4b, 0xc4, 0x40,
	0x10, 0xc5, 0x8d, 0xc1, 0xfc, 0x19, 0x10, 0x64, 0x10, 0x0d, 0xc6, 0x22, 0x6e, 0xa3, 0xd5, 0x16,
	0xfa, 0x19, 0x82, 0x95, 0x16, 0x09, 0x69, 0x95, 0x55, 0x87, 0x10, 0x70, 0x77, 0xe3, 0x4c, 0x14,
	0xef, 0xdb, 0x1f, 0x49, 0x16, 0x52, 0xdd, 0x75, 0xf3, 0xde, 0xbc, 0x81, 0xdf, 0x1b, 0x38, 0xa7,
	0x7f, 0x63, 0xc7, 0x6f, 0xd2, 0x23, 0xfb, 0xc9, 0xe3, 0x75, 0xef, 0xb5, 0x1d, 0x3e, 0xd9, 0x6b,
	0xe1, 0x3f, 0xfd, 0x4c, 0x53, 0x4b, 0x22, 0x83, 0x77, 0xaa, 0x84, 0xf4, 0x85, 0x44, 0x4c, 0x4f,
	0x78, 0x01, 0xb1, 0x98, 0x5d, 0x11, 0x55, 0xd1, 0x43, 0xde, 0xcc, 0xa3, 0xba, 0x87, 0xb4, 0xa1,
	0x9f, 0x5f, 0x92, 0x09, 0x6f, 0x21, 0x0f, 0x27, 0xc3, 0x57, 0x88, 0x6c, 0x86, 0xea, 0x20, 0x6b,
	0x48, 0x46, 0xef, 0x84, 0xf0, 0x12, 0xce, 0x6a, 0x66, 0xe7, 0x43, 0x6a, 0x15, 0x78, 0x05, 0x49,
	0xcd, 0x6c, 0xa5, 0x2f, 0x4e, 0x17, 0x3b, 0x28, 0x2c, 0x21, 0xef, 0x84, 0xf8, 0xdd, 0x19, 0x4b,
	0x45, 0xbc, 0xac, 0xb2, 0xd9, 0x78, 0x35, 0x96, 0x1e, 0xdf, 0x20, 0xad, 0xd7, 0x1a, 0xd8, 0x02,
	0x6c, 0xd4, 0x58, 0xe9, 0x03, 0x7d, 0x74, 0xe0, 0xbd, 0xb9, 0x3b, 0x92, 0x58, 0x41, 0xd5, 0xc9,
	0x47, 0xb2, 0x3c, 0xe7, 0x69, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xbc, 0x35, 0xba, 0x2d, 0x01,
	0x00, 0x00,
}
