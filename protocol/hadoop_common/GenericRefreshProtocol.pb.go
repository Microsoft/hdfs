// Code generated by protoc-gen-go. DO NOT EDIT.
// source: GenericRefreshProtocol.proto

package hadoop_common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// *
//  Refresh request.
type GenericRefreshRequestProto struct {
	Identifier       *string  `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	Args             []string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GenericRefreshRequestProto) Reset()                    { *m = GenericRefreshRequestProto{} }
func (m *GenericRefreshRequestProto) String() string            { return proto.CompactTextString(m) }
func (*GenericRefreshRequestProto) ProtoMessage()               {}
func (*GenericRefreshRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *GenericRefreshRequestProto) GetIdentifier() string {
	if m != nil && m.Identifier != nil {
		return *m.Identifier
	}
	return ""
}

func (m *GenericRefreshRequestProto) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

// *
// A single response from a refresh handler.
type GenericRefreshResponseProto struct {
	ExitStatus       *int32  `protobuf:"varint,1,opt,name=exitStatus" json:"exitStatus,omitempty"`
	UserMessage      *string `protobuf:"bytes,2,opt,name=userMessage" json:"userMessage,omitempty"`
	SenderName       *string `protobuf:"bytes,3,opt,name=senderName" json:"senderName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GenericRefreshResponseProto) Reset()                    { *m = GenericRefreshResponseProto{} }
func (m *GenericRefreshResponseProto) String() string            { return proto.CompactTextString(m) }
func (*GenericRefreshResponseProto) ProtoMessage()               {}
func (*GenericRefreshResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *GenericRefreshResponseProto) GetExitStatus() int32 {
	if m != nil && m.ExitStatus != nil {
		return *m.ExitStatus
	}
	return 0
}

func (m *GenericRefreshResponseProto) GetUserMessage() string {
	if m != nil && m.UserMessage != nil {
		return *m.UserMessage
	}
	return ""
}

func (m *GenericRefreshResponseProto) GetSenderName() string {
	if m != nil && m.SenderName != nil {
		return *m.SenderName
	}
	return ""
}

// *
// Collection of responses from zero or more handlers.
type GenericRefreshResponseCollectionProto struct {
	Responses        []*GenericRefreshResponseProto `protobuf:"bytes,1,rep,name=responses" json:"responses,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *GenericRefreshResponseCollectionProto) Reset()         { *m = GenericRefreshResponseCollectionProto{} }
func (m *GenericRefreshResponseCollectionProto) String() string { return proto.CompactTextString(m) }
func (*GenericRefreshResponseCollectionProto) ProtoMessage()    {}
func (*GenericRefreshResponseCollectionProto) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{2}
}

func (m *GenericRefreshResponseCollectionProto) GetResponses() []*GenericRefreshResponseProto {
	if m != nil {
		return m.Responses
	}
	return nil
}

func init() {
	proto.RegisterType((*GenericRefreshRequestProto)(nil), "hadoop.common.GenericRefreshRequestProto")
	proto.RegisterType((*GenericRefreshResponseProto)(nil), "hadoop.common.GenericRefreshResponseProto")
	proto.RegisterType((*GenericRefreshResponseCollectionProto)(nil), "hadoop.common.GenericRefreshResponseCollectionProto")
}

func init() { proto.RegisterFile("GenericRefreshProtocol.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xd9, 0x56, 0x91, 0x4e, 0xf1, 0xb2, 0xa7, 0xd0, 0xaa, 0x94, 0x80, 0x10, 0x3d, 0xe4,
	0x50, 0x7c, 0x01, 0xeb, 0x41, 0x2f, 0x4a, 0xd9, 0x3e, 0xc1, 0xb2, 0x99, 0x26, 0x2b, 0xc9, 0x6e,
	0xba, 0xb3, 0x11, 0x6f, 0x5e, 0x05, 0x5f, 0xc2, 0x47, 0x95, 0x24, 0x6a, 0xd2, 0x52, 0xff, 0x9c,
	0x76, 0xf8, 0x66, 0xe7, 0x9b, 0xdf, 0x0c, 0x03, 0x27, 0xb7, 0x68, 0xd0, 0x69, 0x25, 0x70, 0xed,
	0x90, 0xb2, 0xa5, 0xb3, 0xde, 0x2a, 0x9b, 0xc7, 0x65, 0x1d, 0xf0, 0xe3, 0x4c, 0x26, 0xd6, 0x96,
	0xb1, 0xb2, 0x45, 0x61, 0x4d, 0xb8, 0x84, 0xc9, 0xf6, 0x77, 0x81, 0x9b, 0x0a, 0xc9, 0x37, 0x55,
	0xfc, 0x0c, 0x40, 0x27, 0x68, 0xbc, 0x5e, 0x6b, 0x74, 0x01, 0x9b, 0xb1, 0x68, 0x24, 0x7a, 0x0a,
	0xe7, 0x70, 0x20, 0x5d, 0x4a, 0xc1, 0x60, 0x36, 0x8c, 0x46, 0xa2, 0x89, 0xc3, 0x17, 0x98, 0xee,
	0x3a, 0x52, 0x69, 0x0d, 0xe1, 0xb7, 0x25, 0x3e, 0x6b, 0xbf, 0xf2, 0xd2, 0x57, 0xd4, 0x58, 0x1e,
	0x8a, 0x9e, 0xc2, 0x67, 0x30, 0xae, 0x08, 0xdd, 0x3d, 0x12, 0xc9, 0x14, 0x83, 0x41, 0xd3, 0xb3,
	0x2f, 0xd5, 0x0e, 0x84, 0x26, 0x41, 0xf7, 0x20, 0x0b, 0x0c, 0x86, 0x2d, 0x54, 0xa7, 0x84, 0x1b,
	0x38, 0xdf, 0x0f, 0x70, 0x63, 0xf3, 0x1c, 0x95, 0xd7, 0xd6, 0xb4, 0x28, 0x77, 0x30, 0x72, 0x9f,
	0xa9, 0x9a, 0x64, 0x18, 0x8d, 0xe7, 0x97, 0xf1, 0xd6, 0x7a, 0xe2, 0x5f, 0x26, 0x11, 0x5d, 0xf1,
	0xfc, 0x8d, 0xc1, 0xe9, 0xfe, 0xad, 0xaf, 0xd0, 0x3d, 0x69, 0x85, 0xfc, 0x11, 0x8e, 0x5c, 0x9b,
	0xe1, 0x17, 0x7f, 0xf4, 0xe8, 0xf6, 0x3f, 0xb9, 0xfa, 0x17, 0xce, 0xce, 0x5c, 0x8b, 0x6b, 0x98,
	0x5a, 0x97, 0xc6, 0xb2, 0x94, 0x2a, 0xc3, 0x2f, 0x07, 0x5d, 0xaa, 0xf6, 0x02, 0x16, 0x3f, 0xdc,
	0x47, 0xf3, 0xd2, 0x2b, 0x63, 0xef, 0x8c, 0x7d, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x01, 0x5c,
	0x07, 0x44, 0x02, 0x00, 0x00,
}
