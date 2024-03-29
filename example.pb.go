// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package main

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

type ExampleMsg struct {
	SimpleString string `protobuf:"bytes,1,opt,name=simple_string,json=simpleString,proto3" json:"simple_string,omitempty"`
	// Types that are valid to be assigned to ComplexField:
	//	*ExampleMsg_SomeString
	//	*ExampleMsg_SomeNumber
	//	*ExampleMsg_SomeBool
	ComplexField         isExampleMsg_ComplexField `protobuf_oneof:"complex_field"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ExampleMsg) Reset()         { *m = ExampleMsg{} }
func (m *ExampleMsg) String() string { return proto.CompactTextString(m) }
func (*ExampleMsg) ProtoMessage()    {}
func (*ExampleMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *ExampleMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExampleMsg.Unmarshal(m, b)
}
func (m *ExampleMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExampleMsg.Marshal(b, m, deterministic)
}
func (m *ExampleMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExampleMsg.Merge(m, src)
}
func (m *ExampleMsg) XXX_Size() int {
	return xxx_messageInfo_ExampleMsg.Size(m)
}
func (m *ExampleMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ExampleMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ExampleMsg proto.InternalMessageInfo

func (m *ExampleMsg) GetSimpleString() string {
	if m != nil {
		return m.SimpleString
	}
	return ""
}

type isExampleMsg_ComplexField interface {
	isExampleMsg_ComplexField()
}

type ExampleMsg_SomeString struct {
	SomeString string `protobuf:"bytes,2,opt,name=some_string,json=someString,proto3,oneof"`
}

type ExampleMsg_SomeNumber struct {
	SomeNumber int64 `protobuf:"varint,3,opt,name=some_number,json=someNumber,proto3,oneof"`
}

type ExampleMsg_SomeBool struct {
	SomeBool bool `protobuf:"varint,4,opt,name=some_bool,json=someBool,proto3,oneof"`
}

func (*ExampleMsg_SomeString) isExampleMsg_ComplexField() {}

func (*ExampleMsg_SomeNumber) isExampleMsg_ComplexField() {}

func (*ExampleMsg_SomeBool) isExampleMsg_ComplexField() {}

func (m *ExampleMsg) GetComplexField() isExampleMsg_ComplexField {
	if m != nil {
		return m.ComplexField
	}
	return nil
}

func (m *ExampleMsg) GetSomeString() string {
	if x, ok := m.GetComplexField().(*ExampleMsg_SomeString); ok {
		return x.SomeString
	}
	return ""
}

func (m *ExampleMsg) GetSomeNumber() int64 {
	if x, ok := m.GetComplexField().(*ExampleMsg_SomeNumber); ok {
		return x.SomeNumber
	}
	return 0
}

func (m *ExampleMsg) GetSomeBool() bool {
	if x, ok := m.GetComplexField().(*ExampleMsg_SomeBool); ok {
		return x.SomeBool
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExampleMsg) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExampleMsg_SomeString)(nil),
		(*ExampleMsg_SomeNumber)(nil),
		(*ExampleMsg_SomeBool)(nil),
	}
}

func init() {
	proto.RegisterType((*ExampleMsg)(nil), "ExampleMsg")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x5a, 0xce, 0xc8, 0xc5, 0xe5, 0x0a,
	0x11, 0xf1, 0x2d, 0x4e, 0x17, 0x52, 0xe6, 0xe2, 0x2d, 0xce, 0x04, 0x71, 0xe2, 0x8b, 0x4b, 0x8a,
	0x32, 0xf3, 0xd2, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x78, 0x20, 0x82, 0xc1, 0x60, 0x31,
	0x21, 0x45, 0x2e, 0xee, 0xe2, 0xfc, 0x5c, 0xb8, 0x12, 0x26, 0x90, 0x12, 0x0f, 0x86, 0x20, 0x2e,
	0x90, 0x20, 0x9a, 0x92, 0xbc, 0xd2, 0xdc, 0xa4, 0xd4, 0x22, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x66,
	0x98, 0x12, 0x3f, 0xb0, 0x98, 0x90, 0x2c, 0x17, 0x27, 0x58, 0x49, 0x52, 0x7e, 0x7e, 0x8e, 0x04,
	0x8b, 0x02, 0xa3, 0x06, 0x87, 0x07, 0x43, 0x10, 0x07, 0x48, 0xc8, 0x29, 0x3f, 0x3f, 0xc7, 0x89,
	0x9f, 0x8b, 0x37, 0x39, 0x1f, 0x64, 0x69, 0x45, 0x7c, 0x5a, 0x66, 0x6a, 0x4e, 0x8a, 0x13, 0x5b,
	0x14, 0x4b, 0x6e, 0x62, 0x66, 0x5e, 0x12, 0x1b, 0xd8, 0xe1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf8, 0xb0, 0x9c, 0xfb, 0xc9, 0x00, 0x00, 0x00,
}
