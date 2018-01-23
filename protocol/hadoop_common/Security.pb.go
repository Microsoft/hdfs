// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Security.proto

package hadoop_common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// *
// Security token identifier
type TokenProto struct {
	Identifier       []byte  `protobuf:"bytes,1,req,name=identifier" json:"identifier,omitempty"`
	Password         []byte  `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
	Kind             *string `protobuf:"bytes,3,req,name=kind" json:"kind,omitempty"`
	Service          *string `protobuf:"bytes,4,req,name=service" json:"service,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TokenProto) Reset()                    { *m = TokenProto{} }
func (m *TokenProto) String() string            { return proto.CompactTextString(m) }
func (*TokenProto) ProtoMessage()               {}
func (*TokenProto) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *TokenProto) GetIdentifier() []byte {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *TokenProto) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *TokenProto) GetKind() string {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return ""
}

func (m *TokenProto) GetService() string {
	if m != nil && m.Service != nil {
		return *m.Service
	}
	return ""
}

type GetDelegationTokenRequestProto struct {
	Renewer          *string `protobuf:"bytes,1,req,name=renewer" json:"renewer,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetDelegationTokenRequestProto) Reset()                    { *m = GetDelegationTokenRequestProto{} }
func (m *GetDelegationTokenRequestProto) String() string            { return proto.CompactTextString(m) }
func (*GetDelegationTokenRequestProto) ProtoMessage()               {}
func (*GetDelegationTokenRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *GetDelegationTokenRequestProto) GetRenewer() string {
	if m != nil && m.Renewer != nil {
		return *m.Renewer
	}
	return ""
}

type GetDelegationTokenResponseProto struct {
	Token            *TokenProto `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *GetDelegationTokenResponseProto) Reset()                    { *m = GetDelegationTokenResponseProto{} }
func (m *GetDelegationTokenResponseProto) String() string            { return proto.CompactTextString(m) }
func (*GetDelegationTokenResponseProto) ProtoMessage()               {}
func (*GetDelegationTokenResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *GetDelegationTokenResponseProto) GetToken() *TokenProto {
	if m != nil {
		return m.Token
	}
	return nil
}

type RenewDelegationTokenRequestProto struct {
	Token            *TokenProto `protobuf:"bytes,1,req,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *RenewDelegationTokenRequestProto) Reset()         { *m = RenewDelegationTokenRequestProto{} }
func (m *RenewDelegationTokenRequestProto) String() string { return proto.CompactTextString(m) }
func (*RenewDelegationTokenRequestProto) ProtoMessage()    {}
func (*RenewDelegationTokenRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{3}
}

func (m *RenewDelegationTokenRequestProto) GetToken() *TokenProto {
	if m != nil {
		return m.Token
	}
	return nil
}

type RenewDelegationTokenResponseProto struct {
	NewExpiryTime    *uint64 `protobuf:"varint,1,req,name=newExpiryTime" json:"newExpiryTime,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RenewDelegationTokenResponseProto) Reset()         { *m = RenewDelegationTokenResponseProto{} }
func (m *RenewDelegationTokenResponseProto) String() string { return proto.CompactTextString(m) }
func (*RenewDelegationTokenResponseProto) ProtoMessage()    {}
func (*RenewDelegationTokenResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{4}
}

func (m *RenewDelegationTokenResponseProto) GetNewExpiryTime() uint64 {
	if m != nil && m.NewExpiryTime != nil {
		return *m.NewExpiryTime
	}
	return 0
}

type CancelDelegationTokenRequestProto struct {
	Token            *TokenProto `protobuf:"bytes,1,req,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CancelDelegationTokenRequestProto) Reset()         { *m = CancelDelegationTokenRequestProto{} }
func (m *CancelDelegationTokenRequestProto) String() string { return proto.CompactTextString(m) }
func (*CancelDelegationTokenRequestProto) ProtoMessage()    {}
func (*CancelDelegationTokenRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{5}
}

func (m *CancelDelegationTokenRequestProto) GetToken() *TokenProto {
	if m != nil {
		return m.Token
	}
	return nil
}

type CancelDelegationTokenResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CancelDelegationTokenResponseProto) Reset()         { *m = CancelDelegationTokenResponseProto{} }
func (m *CancelDelegationTokenResponseProto) String() string { return proto.CompactTextString(m) }
func (*CancelDelegationTokenResponseProto) ProtoMessage()    {}
func (*CancelDelegationTokenResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{6}
}

func init() {
	proto.RegisterType((*TokenProto)(nil), "hadoop.common.TokenProto")
	proto.RegisterType((*GetDelegationTokenRequestProto)(nil), "hadoop.common.GetDelegationTokenRequestProto")
	proto.RegisterType((*GetDelegationTokenResponseProto)(nil), "hadoop.common.GetDelegationTokenResponseProto")
	proto.RegisterType((*RenewDelegationTokenRequestProto)(nil), "hadoop.common.RenewDelegationTokenRequestProto")
	proto.RegisterType((*RenewDelegationTokenResponseProto)(nil), "hadoop.common.RenewDelegationTokenResponseProto")
	proto.RegisterType((*CancelDelegationTokenRequestProto)(nil), "hadoop.common.CancelDelegationTokenRequestProto")
	proto.RegisterType((*CancelDelegationTokenResponseProto)(nil), "hadoop.common.CancelDelegationTokenResponseProto")
}

func init() { proto.RegisterFile("Security.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xb1, 0x4e, 0xf3, 0x30,
	0x10, 0xc7, 0xe5, 0x7c, 0xfd, 0x54, 0x7a, 0x50, 0x86, 0x4c, 0x86, 0xa1, 0xa4, 0x56, 0x87, 0x4e,
	0x41, 0x62, 0x42, 0x8c, 0x05, 0x84, 0xd8, 0x90, 0xdb, 0x17, 0xb0, 0x92, 0xa3, 0xb5, 0xda, 0xd8,
	0xc6, 0x76, 0x1b, 0xfa, 0x06, 0x3c, 0x06, 0x8f, 0x8a, 0xea, 0x90, 0x92, 0x48, 0x05, 0x31, 0xb0,
	0xf9, 0xfe, 0xe7, 0xfb, 0xe9, 0x77, 0xd2, 0xc1, 0xe9, 0x14, 0xb3, 0xb5, 0x95, 0x7e, 0x9b, 0x1a,
	0xab, 0xbd, 0x8e, 0xfb, 0x0b, 0x91, 0x6b, 0x6d, 0xd2, 0x4c, 0x17, 0x85, 0x56, 0x6c, 0x03, 0x30,
	0xd3, 0x4b, 0x54, 0x4f, 0xa1, 0x39, 0x00, 0x90, 0x39, 0x2a, 0x2f, 0x9f, 0x25, 0x5a, 0x4a, 0x92,
	0x68, 0x7c, 0xc2, 0x1b, 0x49, 0x7c, 0x0e, 0x47, 0x46, 0x38, 0x57, 0x6a, 0x9b, 0xd3, 0x28, 0x74,
	0xf7, 0x75, 0x1c, 0x43, 0x67, 0x29, 0x55, 0x4e, 0xff, 0x25, 0xd1, 0xb8, 0xc7, 0xc3, 0x3b, 0xa6,
	0xd0, 0x75, 0x68, 0x37, 0x32, 0x43, 0xda, 0x09, 0x71, 0x5d, 0xb2, 0x1b, 0x18, 0x3c, 0xa0, 0xbf,
	0xc3, 0x15, 0xce, 0x85, 0x97, 0x5a, 0x05, 0x09, 0x8e, 0x2f, 0x6b, 0x74, 0xbe, 0x72, 0xa1, 0xd0,
	0xb5, 0xa8, 0xb0, 0xfc, 0x14, 0xe9, 0xf1, 0xba, 0x64, 0x1c, 0x2e, 0x0e, 0xcd, 0x3a, 0xa3, 0x95,
	0xc3, 0x6a, 0xf8, 0x12, 0xfe, 0xfb, 0x5d, 0x4a, 0x49, 0x42, 0xc6, 0xc7, 0x57, 0x67, 0x69, 0x6b,
	0xeb, 0xf4, 0x6b, 0x65, 0x5e, 0xfd, 0x63, 0x53, 0x48, 0xf8, 0x0e, 0xff, 0x93, 0x51, 0x03, 0x1a,
	0xfd, 0x0a, 0xfa, 0x08, 0xc3, 0xc3, 0xd0, 0xa6, 0xea, 0x08, 0xfa, 0x0a, 0xcb, 0xfb, 0x57, 0x23,
	0xed, 0x76, 0x26, 0x0b, 0x0c, 0xf4, 0x0e, 0x6f, 0x87, 0x6c, 0x06, 0xc3, 0x5b, 0xa1, 0x32, 0x5c,
	0xfd, 0xa9, 0xe0, 0x08, 0xd8, 0x37, 0xd4, 0x86, 0xe1, 0xe4, 0x1a, 0x12, 0x6d, 0xe7, 0xa9, 0x30,
	0x22, 0x5b, 0x60, 0xcd, 0x74, 0xad, 0xb3, 0x9a, 0xec, 0xcf, 0x2c, 0x8c, 0xb8, 0x37, 0x42, 0xde,
	0x09, 0xf9, 0x08, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x67, 0xa7, 0xeb, 0x7d, 0x02, 0x00, 0x00,
}
