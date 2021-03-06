// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mempool_status.proto

package xproto

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

type MempoolAddTransactionStatusCode int32

const (
	// Transaction was sent to Mempool
	MempoolAddTransactionStatusCode_Valid MempoolAddTransactionStatusCode = 0
	// The sender does not have enough balance for the transaction.
	MempoolAddTransactionStatusCode_InsufficientBalance MempoolAddTransactionStatusCode = 1
	// Sequence number is old, etc.
	MempoolAddTransactionStatusCode_InvalidSeqNumber MempoolAddTransactionStatusCode = 2
	// Mempool is full (reached max global capacity)
	MempoolAddTransactionStatusCode_MempoolIsFull MempoolAddTransactionStatusCode = 3
	// Account reached max capacity per account
	MempoolAddTransactionStatusCode_TooManyTransactions MempoolAddTransactionStatusCode = 4
	// Invalid update. Only gas price increase is allowed
	MempoolAddTransactionStatusCode_InvalidUpdate MempoolAddTransactionStatusCode = 5
)

var MempoolAddTransactionStatusCode_name = map[int32]string{
	0: "Valid",
	1: "InsufficientBalance",
	2: "InvalidSeqNumber",
	3: "MempoolIsFull",
	4: "TooManyTransactions",
	5: "InvalidUpdate",
}

var MempoolAddTransactionStatusCode_value = map[string]int32{
	"Valid":               0,
	"InsufficientBalance": 1,
	"InvalidSeqNumber":    2,
	"MempoolIsFull":       3,
	"TooManyTransactions": 4,
	"InvalidUpdate":       5,
}

func (x MempoolAddTransactionStatusCode) String() string {
	return proto.EnumName(MempoolAddTransactionStatusCode_name, int32(x))
}

func (MempoolAddTransactionStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cad4a86f8a5465be, []int{0}
}

type MempoolAddTransactionStatus struct {
	Code                 MempoolAddTransactionStatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=mempool.MempoolAddTransactionStatusCode" json:"code,omitempty"`
	Message              string                          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *MempoolAddTransactionStatus) Reset()         { *m = MempoolAddTransactionStatus{} }
func (m *MempoolAddTransactionStatus) String() string { return proto.CompactTextString(m) }
func (*MempoolAddTransactionStatus) ProtoMessage()    {}
func (*MempoolAddTransactionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_cad4a86f8a5465be, []int{0}
}

func (m *MempoolAddTransactionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MempoolAddTransactionStatus.Unmarshal(m, b)
}
func (m *MempoolAddTransactionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MempoolAddTransactionStatus.Marshal(b, m, deterministic)
}
func (m *MempoolAddTransactionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MempoolAddTransactionStatus.Merge(m, src)
}
func (m *MempoolAddTransactionStatus) XXX_Size() int {
	return xxx_messageInfo_MempoolAddTransactionStatus.Size(m)
}
func (m *MempoolAddTransactionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_MempoolAddTransactionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_MempoolAddTransactionStatus proto.InternalMessageInfo

func (m *MempoolAddTransactionStatus) GetCode() MempoolAddTransactionStatusCode {
	if m != nil {
		return m.Code
	}
	return MempoolAddTransactionStatusCode_Valid
}

func (m *MempoolAddTransactionStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("mempool.MempoolAddTransactionStatusCode", MempoolAddTransactionStatusCode_name, MempoolAddTransactionStatusCode_value)
	proto.RegisterType((*MempoolAddTransactionStatus)(nil), "mempool.MempoolAddTransactionStatus")
}

func init() { proto.RegisterFile("mempool_status.proto", fileDescriptor_cad4a86f8a5465be) }

var fileDescriptor_cad4a86f8a5465be = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x4d, 0xcd, 0x2d,
	0xc8, 0xcf, 0xcf, 0x89, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x87, 0x8a, 0x2a, 0x95, 0x72, 0x49, 0xfb, 0x42, 0x98, 0x8e, 0x29, 0x29, 0x21, 0x45,
	0x89, 0x79, 0xc5, 0x89, 0xc9, 0x25, 0x99, 0xf9, 0x79, 0xc1, 0x60, 0xd5, 0x42, 0x36, 0x5c, 0x2c,
	0xc9, 0xf9, 0x29, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0x1a, 0x7a, 0x50, 0x6d, 0x7a,
	0x78, 0xf4, 0x38, 0xe7, 0xa7, 0xa4, 0x06, 0x81, 0x75, 0x09, 0x49, 0x70, 0xb1, 0xe7, 0xa6, 0x16,
	0x17, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x5a, 0xb3, 0x18,
	0xb9, 0xe4, 0x09, 0x98, 0x21, 0xc4, 0xc9, 0xc5, 0x1a, 0x96, 0x98, 0x93, 0x99, 0x22, 0xc0, 0x20,
	0x24, 0xce, 0x25, 0xec, 0x99, 0x57, 0x5c, 0x9a, 0x96, 0x96, 0x99, 0x9c, 0x99, 0x9a, 0x57, 0xe2,
	0x94, 0x98, 0x93, 0x98, 0x97, 0x9c, 0x2a, 0xc0, 0x28, 0x24, 0xc2, 0x25, 0xe0, 0x99, 0x57, 0x06,
	0x52, 0x15, 0x9c, 0x5a, 0xe8, 0x57, 0x9a, 0x9b, 0x94, 0x5a, 0x24, 0xc0, 0x24, 0x24, 0xc8, 0xc5,
	0x0b, 0x35, 0xdc, 0xb3, 0xd8, 0xad, 0x34, 0x27, 0x47, 0x80, 0x19, 0x64, 0x42, 0x48, 0x7e, 0xbe,
	0x6f, 0x62, 0x5e, 0x25, 0x92, 0x65, 0xc5, 0x02, 0x2c, 0x20, 0xb5, 0x50, 0x13, 0x42, 0x0b, 0x52,
	0x12, 0x4b, 0x52, 0x05, 0x58, 0x9d, 0x38, 0xa2, 0xd8, 0x2a, 0xc0, 0xc1, 0x94, 0xc4, 0x06, 0xa6,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x5e, 0xf2, 0xe6, 0x45, 0x01, 0x00, 0x00,
}
