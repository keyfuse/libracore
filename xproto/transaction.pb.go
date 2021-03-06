// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction.proto

package xproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Type of write operation
type WriteOpType int32

const (
	// The WriteOp is to create/update the field from storage.
	WriteOpType_Write WriteOpType = 0
	// The WriteOp is to delete the field from storage.
	WriteOpType_Delete WriteOpType = 1
)

var WriteOpType_name = map[int32]string{
	0: "Write",
	1: "Delete",
}

var WriteOpType_value = map[string]int32{
	"Write":  0,
	"Delete": 1,
}

func (x WriteOpType) String() string {
	return proto.EnumName(WriteOpType_name, int32(x))
}

func (WriteOpType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
}

type TransactionArgument_ArgType int32

const (
	TransactionArgument_U64       TransactionArgument_ArgType = 0
	TransactionArgument_ADDRESS   TransactionArgument_ArgType = 1
	TransactionArgument_STRING    TransactionArgument_ArgType = 2
	TransactionArgument_BYTEARRAY TransactionArgument_ArgType = 3
)

var TransactionArgument_ArgType_name = map[int32]string{
	0: "U64",
	1: "ADDRESS",
	2: "STRING",
	3: "BYTEARRAY",
}

var TransactionArgument_ArgType_value = map[string]int32{
	"U64":       0,
	"ADDRESS":   1,
	"STRING":    2,
	"BYTEARRAY": 3,
}

func (x TransactionArgument_ArgType) String() string {
	return proto.EnumName(TransactionArgument_ArgType_name, int32(x))
}

func (TransactionArgument_ArgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{3, 0}
}

// A generic structure that describes a transaction that a client submits
type RawTransaction struct {
	// Sender's account address
	SenderAccount []byte `protobuf:"bytes,1,opt,name=sender_account,json=senderAccount,proto3" json:"sender_account,omitempty"`
	// Sequence number of this transaction corresponding to sender's account.
	SequenceNumber uint64 `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*RawTransaction_Program
	//	*RawTransaction_WriteSet
	//	*RawTransaction_Script
	//	*RawTransaction_Module
	Payload isRawTransaction_Payload `protobuf_oneof:"payload"`
	// Maximal total gas specified by wallet to spend for this transaction.
	MaxGasAmount uint64 `protobuf:"varint,5,opt,name=max_gas_amount,json=maxGasAmount,proto3" json:"max_gas_amount,omitempty"`
	// The price to be paid for each unit of gas.
	GasUnitPrice uint64 `protobuf:"varint,6,opt,name=gas_unit_price,json=gasUnitPrice,proto3" json:"gas_unit_price,omitempty"`
	// Expiration time for this transaction.  If storage is queried and
	// the time returned is greater than or equal to this time and this
	// transaction has not been included, you can be certain that it will
	// never be included.
	// If set to 0, there will be no expiration time
	ExpirationTime       uint64   `protobuf:"varint,7,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RawTransaction) Reset()         { *m = RawTransaction{} }
func (m *RawTransaction) String() string { return proto.CompactTextString(m) }
func (*RawTransaction) ProtoMessage()    {}
func (*RawTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
}

func (m *RawTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawTransaction.Unmarshal(m, b)
}
func (m *RawTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawTransaction.Marshal(b, m, deterministic)
}
func (m *RawTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawTransaction.Merge(m, src)
}
func (m *RawTransaction) XXX_Size() int {
	return xxx_messageInfo_RawTransaction.Size(m)
}
func (m *RawTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_RawTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_RawTransaction proto.InternalMessageInfo

func (m *RawTransaction) GetSenderAccount() []byte {
	if m != nil {
		return m.SenderAccount
	}
	return nil
}

func (m *RawTransaction) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

type isRawTransaction_Payload interface {
	isRawTransaction_Payload()
}

type RawTransaction_Program struct {
	Program *Program `protobuf:"bytes,3,opt,name=program,proto3,oneof"`
}

type RawTransaction_WriteSet struct {
	WriteSet *WriteSet `protobuf:"bytes,4,opt,name=write_set,json=writeSet,proto3,oneof"`
}

type RawTransaction_Script struct {
	Script *Script `protobuf:"bytes,8,opt,name=script,proto3,oneof"`
}

type RawTransaction_Module struct {
	Module *Module `protobuf:"bytes,9,opt,name=module,proto3,oneof"`
}

func (*RawTransaction_Program) isRawTransaction_Payload() {}

func (*RawTransaction_WriteSet) isRawTransaction_Payload() {}

func (*RawTransaction_Script) isRawTransaction_Payload() {}

func (*RawTransaction_Module) isRawTransaction_Payload() {}

func (m *RawTransaction) GetPayload() isRawTransaction_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *RawTransaction) GetProgram() *Program {
	if x, ok := m.GetPayload().(*RawTransaction_Program); ok {
		return x.Program
	}
	return nil
}

func (m *RawTransaction) GetWriteSet() *WriteSet {
	if x, ok := m.GetPayload().(*RawTransaction_WriteSet); ok {
		return x.WriteSet
	}
	return nil
}

func (m *RawTransaction) GetScript() *Script {
	if x, ok := m.GetPayload().(*RawTransaction_Script); ok {
		return x.Script
	}
	return nil
}

func (m *RawTransaction) GetModule() *Module {
	if x, ok := m.GetPayload().(*RawTransaction_Module); ok {
		return x.Module
	}
	return nil
}

func (m *RawTransaction) GetMaxGasAmount() uint64 {
	if m != nil {
		return m.MaxGasAmount
	}
	return 0
}

func (m *RawTransaction) GetGasUnitPrice() uint64 {
	if m != nil {
		return m.GasUnitPrice
	}
	return 0
}

func (m *RawTransaction) GetExpirationTime() uint64 {
	if m != nil {
		return m.ExpirationTime
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RawTransaction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RawTransaction_Program)(nil),
		(*RawTransaction_WriteSet)(nil),
		(*RawTransaction_Script)(nil),
		(*RawTransaction_Module)(nil),
	}
}

// The code for the transaction to execute
type Program struct {
	Code                 []byte                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Arguments            []*TransactionArgument `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments,omitempty"`
	Modules              [][]byte               `protobuf:"bytes,3,rep,name=modules,proto3" json:"modules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Program) Reset()         { *m = Program{} }
func (m *Program) String() string { return proto.CompactTextString(m) }
func (*Program) ProtoMessage()    {}
func (*Program) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{1}
}

func (m *Program) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Program.Unmarshal(m, b)
}
func (m *Program) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Program.Marshal(b, m, deterministic)
}
func (m *Program) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Program.Merge(m, src)
}
func (m *Program) XXX_Size() int {
	return xxx_messageInfo_Program.Size(m)
}
func (m *Program) XXX_DiscardUnknown() {
	xxx_messageInfo_Program.DiscardUnknown(m)
}

var xxx_messageInfo_Program proto.InternalMessageInfo

func (m *Program) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *Program) GetArguments() []*TransactionArgument {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func (m *Program) GetModules() [][]byte {
	if m != nil {
		return m.Modules
	}
	return nil
}

// The code for the transaction to execute
type Script struct {
	Code                 []byte                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Arguments            []*TransactionArgument `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Script) Reset()         { *m = Script{} }
func (m *Script) String() string { return proto.CompactTextString(m) }
func (*Script) ProtoMessage()    {}
func (*Script) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{2}
}

func (m *Script) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Script.Unmarshal(m, b)
}
func (m *Script) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Script.Marshal(b, m, deterministic)
}
func (m *Script) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Script.Merge(m, src)
}
func (m *Script) XXX_Size() int {
	return xxx_messageInfo_Script.Size(m)
}
func (m *Script) XXX_DiscardUnknown() {
	xxx_messageInfo_Script.DiscardUnknown(m)
}

var xxx_messageInfo_Script proto.InternalMessageInfo

func (m *Script) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *Script) GetArguments() []*TransactionArgument {
	if m != nil {
		return m.Arguments
	}
	return nil
}

// An argument to the transaction if the transaction takes arguments
type TransactionArgument struct {
	Type                 TransactionArgument_ArgType `protobuf:"varint,1,opt,name=type,proto3,enum=types.TransactionArgument_ArgType" json:"type,omitempty"`
	Data                 []byte                      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TransactionArgument) Reset()         { *m = TransactionArgument{} }
func (m *TransactionArgument) String() string { return proto.CompactTextString(m) }
func (*TransactionArgument) ProtoMessage()    {}
func (*TransactionArgument) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{3}
}

func (m *TransactionArgument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionArgument.Unmarshal(m, b)
}
func (m *TransactionArgument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionArgument.Marshal(b, m, deterministic)
}
func (m *TransactionArgument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionArgument.Merge(m, src)
}
func (m *TransactionArgument) XXX_Size() int {
	return xxx_messageInfo_TransactionArgument.Size(m)
}
func (m *TransactionArgument) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionArgument.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionArgument proto.InternalMessageInfo

func (m *TransactionArgument) GetType() TransactionArgument_ArgType {
	if m != nil {
		return m.Type
	}
	return TransactionArgument_U64
}

func (m *TransactionArgument) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// A Move Module to publish
type Module struct {
	Code                 []byte   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Module) Reset()         { *m = Module{} }
func (m *Module) String() string { return proto.CompactTextString(m) }
func (*Module) ProtoMessage()    {}
func (*Module) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{4}
}

func (m *Module) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Module.Unmarshal(m, b)
}
func (m *Module) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Module.Marshal(b, m, deterministic)
}
func (m *Module) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Module.Merge(m, src)
}
func (m *Module) XXX_Size() int {
	return xxx_messageInfo_Module.Size(m)
}
func (m *Module) XXX_DiscardUnknown() {
	xxx_messageInfo_Module.DiscardUnknown(m)
}

var xxx_messageInfo_Module proto.InternalMessageInfo

func (m *Module) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

// A generic structure that represents signed RawTransaction
type SignedTransaction struct {
	// The serialized Protobuf bytes for RawTransaction, for which the signature
	// was signed. Protobuf doesn't guarantee the serialized bytes is canonical
	// across different language implementations, but for our use cases for
	// transaction it is not necessary because the client is the only one to
	// produce this bytes, which is then persisted in storage.
	RawTxnBytes []byte `protobuf:"bytes,1,opt,name=raw_txn_bytes,json=rawTxnBytes,proto3" json:"raw_txn_bytes,omitempty"`
	// public key that corresponds to RawTransaction::sender_account
	SenderPublicKey []byte `protobuf:"bytes,2,opt,name=sender_public_key,json=senderPublicKey,proto3" json:"sender_public_key,omitempty"`
	// signature for the hash
	SenderSignature      []byte   `protobuf:"bytes,3,opt,name=sender_signature,json=senderSignature,proto3" json:"sender_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedTransaction) Reset()         { *m = SignedTransaction{} }
func (m *SignedTransaction) String() string { return proto.CompactTextString(m) }
func (*SignedTransaction) ProtoMessage()    {}
func (*SignedTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{5}
}

func (m *SignedTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedTransaction.Unmarshal(m, b)
}
func (m *SignedTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedTransaction.Marshal(b, m, deterministic)
}
func (m *SignedTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedTransaction.Merge(m, src)
}
func (m *SignedTransaction) XXX_Size() int {
	return xxx_messageInfo_SignedTransaction.Size(m)
}
func (m *SignedTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_SignedTransaction proto.InternalMessageInfo

func (m *SignedTransaction) GetRawTxnBytes() []byte {
	if m != nil {
		return m.RawTxnBytes
	}
	return nil
}

func (m *SignedTransaction) GetSenderPublicKey() []byte {
	if m != nil {
		return m.SenderPublicKey
	}
	return nil
}

func (m *SignedTransaction) GetSenderSignature() []byte {
	if m != nil {
		return m.SenderSignature
	}
	return nil
}

type SignedTransactionWithProof struct {
	// The version of the returned signed transaction.
	Version uint64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// The transaction itself.
	SignedTransaction *SignedTransaction `protobuf:"bytes,2,opt,name=signed_transaction,json=signedTransaction,proto3" json:"signed_transaction,omitempty"`
	// The proof authenticating the signed transaction.
	Proof *SignedTransactionProof `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof,omitempty"`
	// The events yielded by executing the transaction, if requested.
	Events               *EventsList `protobuf:"bytes,4,opt,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SignedTransactionWithProof) Reset()         { *m = SignedTransactionWithProof{} }
func (m *SignedTransactionWithProof) String() string { return proto.CompactTextString(m) }
func (*SignedTransactionWithProof) ProtoMessage()    {}
func (*SignedTransactionWithProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{6}
}

func (m *SignedTransactionWithProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedTransactionWithProof.Unmarshal(m, b)
}
func (m *SignedTransactionWithProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedTransactionWithProof.Marshal(b, m, deterministic)
}
func (m *SignedTransactionWithProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedTransactionWithProof.Merge(m, src)
}
func (m *SignedTransactionWithProof) XXX_Size() int {
	return xxx_messageInfo_SignedTransactionWithProof.Size(m)
}
func (m *SignedTransactionWithProof) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedTransactionWithProof.DiscardUnknown(m)
}

var xxx_messageInfo_SignedTransactionWithProof proto.InternalMessageInfo

func (m *SignedTransactionWithProof) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *SignedTransactionWithProof) GetSignedTransaction() *SignedTransaction {
	if m != nil {
		return m.SignedTransaction
	}
	return nil
}

func (m *SignedTransactionWithProof) GetProof() *SignedTransactionProof {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *SignedTransactionWithProof) GetEvents() *EventsList {
	if m != nil {
		return m.Events
	}
	return nil
}

// A generic structure that represents a block of transactions originated from a
// particular validator instance.
type SignedTransactionsBlock struct {
	// Set of Signed Transactions
	Transactions []*SignedTransaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	// Public key of the validator that created this block
	ValidatorPublicKey []byte `protobuf:"bytes,2,opt,name=validator_public_key,json=validatorPublicKey,proto3" json:"validator_public_key,omitempty"`
	// Signature of the validator that created this block
	ValidatorSignature   []byte   `protobuf:"bytes,3,opt,name=validator_signature,json=validatorSignature,proto3" json:"validator_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedTransactionsBlock) Reset()         { *m = SignedTransactionsBlock{} }
func (m *SignedTransactionsBlock) String() string { return proto.CompactTextString(m) }
func (*SignedTransactionsBlock) ProtoMessage()    {}
func (*SignedTransactionsBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{7}
}

func (m *SignedTransactionsBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedTransactionsBlock.Unmarshal(m, b)
}
func (m *SignedTransactionsBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedTransactionsBlock.Marshal(b, m, deterministic)
}
func (m *SignedTransactionsBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedTransactionsBlock.Merge(m, src)
}
func (m *SignedTransactionsBlock) XXX_Size() int {
	return xxx_messageInfo_SignedTransactionsBlock.Size(m)
}
func (m *SignedTransactionsBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedTransactionsBlock.DiscardUnknown(m)
}

var xxx_messageInfo_SignedTransactionsBlock proto.InternalMessageInfo

func (m *SignedTransactionsBlock) GetTransactions() []*SignedTransaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *SignedTransactionsBlock) GetValidatorPublicKey() []byte {
	if m != nil {
		return m.ValidatorPublicKey
	}
	return nil
}

func (m *SignedTransactionsBlock) GetValidatorSignature() []byte {
	if m != nil {
		return m.ValidatorSignature
	}
	return nil
}

// Set of WriteOps to save to storage.
type WriteSet struct {
	// Set of WriteOp for storage update.
	WriteSet             []*WriteOp `protobuf:"bytes,1,rep,name=write_set,json=writeSet,proto3" json:"write_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *WriteSet) Reset()         { *m = WriteSet{} }
func (m *WriteSet) String() string { return proto.CompactTextString(m) }
func (*WriteSet) ProtoMessage()    {}
func (*WriteSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{8}
}

func (m *WriteSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteSet.Unmarshal(m, b)
}
func (m *WriteSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteSet.Marshal(b, m, deterministic)
}
func (m *WriteSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteSet.Merge(m, src)
}
func (m *WriteSet) XXX_Size() int {
	return xxx_messageInfo_WriteSet.Size(m)
}
func (m *WriteSet) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteSet.DiscardUnknown(m)
}

var xxx_messageInfo_WriteSet proto.InternalMessageInfo

func (m *WriteSet) GetWriteSet() []*WriteOp {
	if m != nil {
		return m.WriteSet
	}
	return nil
}

// Write Operation on underlying storage.
type WriteOp struct {
	// AccessPath of the write set.
	AccessPath *AccessPath `protobuf:"bytes,1,opt,name=access_path,json=accessPath,proto3" json:"access_path,omitempty"`
	// The value of the write op. Empty if `type` is Delete.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// WriteOp type.
	Type                 WriteOpType `protobuf:"varint,3,opt,name=type,proto3,enum=types.WriteOpType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *WriteOp) Reset()         { *m = WriteOp{} }
func (m *WriteOp) String() string { return proto.CompactTextString(m) }
func (*WriteOp) ProtoMessage()    {}
func (*WriteOp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{9}
}

func (m *WriteOp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteOp.Unmarshal(m, b)
}
func (m *WriteOp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteOp.Marshal(b, m, deterministic)
}
func (m *WriteOp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteOp.Merge(m, src)
}
func (m *WriteOp) XXX_Size() int {
	return xxx_messageInfo_WriteOp.Size(m)
}
func (m *WriteOp) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteOp.DiscardUnknown(m)
}

var xxx_messageInfo_WriteOp proto.InternalMessageInfo

func (m *WriteOp) GetAccessPath() *AccessPath {
	if m != nil {
		return m.AccessPath
	}
	return nil
}

func (m *WriteOp) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *WriteOp) GetType() WriteOpType {
	if m != nil {
		return m.Type
	}
	return WriteOpType_Write
}

// Account state as a whole.
// After execution, updates to accounts are passed in this form to storage for
// persistence.
type AccountState struct {
	// Account address
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Account state blob
	Blob                 []byte   `protobuf:"bytes,2,opt,name=blob,proto3" json:"blob,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountState) Reset()         { *m = AccountState{} }
func (m *AccountState) String() string { return proto.CompactTextString(m) }
func (*AccountState) ProtoMessage()    {}
func (*AccountState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{10}
}

func (m *AccountState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountState.Unmarshal(m, b)
}
func (m *AccountState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountState.Marshal(b, m, deterministic)
}
func (m *AccountState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountState.Merge(m, src)
}
func (m *AccountState) XXX_Size() int {
	return xxx_messageInfo_AccountState.Size(m)
}
func (m *AccountState) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountState.DiscardUnknown(m)
}

var xxx_messageInfo_AccountState proto.InternalMessageInfo

func (m *AccountState) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *AccountState) GetBlob() []byte {
	if m != nil {
		return m.Blob
	}
	return nil
}

// Transaction struct to commit to storage
type TransactionToCommit struct {
	// The signed transaction which was executed
	SignedTxn *SignedTransaction `protobuf:"bytes,1,opt,name=signed_txn,json=signedTxn,proto3" json:"signed_txn,omitempty"`
	// State db updates
	AccountStates []*AccountState `protobuf:"bytes,2,rep,name=account_states,json=accountStates,proto3" json:"account_states,omitempty"`
	// Events yielded by the transaction.
	Events []*Event `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	// The amount of gas used.
	GasUsed              uint64   `protobuf:"varint,4,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionToCommit) Reset()         { *m = TransactionToCommit{} }
func (m *TransactionToCommit) String() string { return proto.CompactTextString(m) }
func (*TransactionToCommit) ProtoMessage()    {}
func (*TransactionToCommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{11}
}

func (m *TransactionToCommit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionToCommit.Unmarshal(m, b)
}
func (m *TransactionToCommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionToCommit.Marshal(b, m, deterministic)
}
func (m *TransactionToCommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionToCommit.Merge(m, src)
}
func (m *TransactionToCommit) XXX_Size() int {
	return xxx_messageInfo_TransactionToCommit.Size(m)
}
func (m *TransactionToCommit) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionToCommit.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionToCommit proto.InternalMessageInfo

func (m *TransactionToCommit) GetSignedTxn() *SignedTransaction {
	if m != nil {
		return m.SignedTxn
	}
	return nil
}

func (m *TransactionToCommit) GetAccountStates() []*AccountState {
	if m != nil {
		return m.AccountStates
	}
	return nil
}

func (m *TransactionToCommit) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *TransactionToCommit) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

// A list of consecutive transactions with proof. This is mainly used for state
// synchronization when a validator would request a list of transactions from a
// peer, verify the proof, execute the transactions and persist them. Note that
// the transactions are supposed to belong to the same epoch E, otherwise
// verification will fail.
type TransactionListWithProof struct {
	// The list of transactions.
	Transactions []*SignedTransaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	// The list of corresponding TransactionInfo objects.
	Infos []*TransactionInfo `protobuf:"bytes,2,rep,name=infos,proto3" json:"infos,omitempty"`
	// The list of corresponding Event objects (only present if fetch_events was set to true in req)
	EventsForVersions *EventsForVersions `protobuf:"bytes,3,opt,name=events_for_versions,json=eventsForVersions,proto3" json:"events_for_versions,omitempty"`
	// If the list is not empty, the version of the first transaction.
	FirstTransactionVersion *wrappers.UInt64Value `protobuf:"bytes,4,opt,name=first_transaction_version,json=firstTransactionVersion,proto3" json:"first_transaction_version,omitempty"`
	// The proofs of the first and last transaction in this chunk. When this is
	// used for state synchronization, the validator who requests the transactions
	// will provide a version in the request and the proofs will be relative to
	// the given version. When this is returned in GetTransactionsResponse, the
	// proofs will be relative to the ledger info returned in
	// UpdateToLatestLedgerResponse.
	ProofOfFirstTransaction *AccumulatorProof `protobuf:"bytes,5,opt,name=proof_of_first_transaction,json=proofOfFirstTransaction,proto3" json:"proof_of_first_transaction,omitempty"`
	ProofOfLastTransaction  *AccumulatorProof `protobuf:"bytes,6,opt,name=proof_of_last_transaction,json=proofOfLastTransaction,proto3" json:"proof_of_last_transaction,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}          `json:"-"`
	XXX_unrecognized        []byte            `json:"-"`
	XXX_sizecache           int32             `json:"-"`
}

func (m *TransactionListWithProof) Reset()         { *m = TransactionListWithProof{} }
func (m *TransactionListWithProof) String() string { return proto.CompactTextString(m) }
func (*TransactionListWithProof) ProtoMessage()    {}
func (*TransactionListWithProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{12}
}

func (m *TransactionListWithProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionListWithProof.Unmarshal(m, b)
}
func (m *TransactionListWithProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionListWithProof.Marshal(b, m, deterministic)
}
func (m *TransactionListWithProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionListWithProof.Merge(m, src)
}
func (m *TransactionListWithProof) XXX_Size() int {
	return xxx_messageInfo_TransactionListWithProof.Size(m)
}
func (m *TransactionListWithProof) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionListWithProof.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionListWithProof proto.InternalMessageInfo

func (m *TransactionListWithProof) GetTransactions() []*SignedTransaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *TransactionListWithProof) GetInfos() []*TransactionInfo {
	if m != nil {
		return m.Infos
	}
	return nil
}

func (m *TransactionListWithProof) GetEventsForVersions() *EventsForVersions {
	if m != nil {
		return m.EventsForVersions
	}
	return nil
}

func (m *TransactionListWithProof) GetFirstTransactionVersion() *wrappers.UInt64Value {
	if m != nil {
		return m.FirstTransactionVersion
	}
	return nil
}

func (m *TransactionListWithProof) GetProofOfFirstTransaction() *AccumulatorProof {
	if m != nil {
		return m.ProofOfFirstTransaction
	}
	return nil
}

func (m *TransactionListWithProof) GetProofOfLastTransaction() *AccumulatorProof {
	if m != nil {
		return m.ProofOfLastTransaction
	}
	return nil
}

func init() {
	proto.RegisterEnum("types.WriteOpType", WriteOpType_name, WriteOpType_value)
	proto.RegisterEnum("types.TransactionArgument_ArgType", TransactionArgument_ArgType_name, TransactionArgument_ArgType_value)
	proto.RegisterType((*RawTransaction)(nil), "types.RawTransaction")
	proto.RegisterType((*Program)(nil), "types.Program")
	proto.RegisterType((*Script)(nil), "types.Script")
	proto.RegisterType((*TransactionArgument)(nil), "types.TransactionArgument")
	proto.RegisterType((*Module)(nil), "types.Module")
	proto.RegisterType((*SignedTransaction)(nil), "types.SignedTransaction")
	proto.RegisterType((*SignedTransactionWithProof)(nil), "types.SignedTransactionWithProof")
	proto.RegisterType((*SignedTransactionsBlock)(nil), "types.SignedTransactionsBlock")
	proto.RegisterType((*WriteSet)(nil), "types.WriteSet")
	proto.RegisterType((*WriteOp)(nil), "types.WriteOp")
	proto.RegisterType((*AccountState)(nil), "types.AccountState")
	proto.RegisterType((*TransactionToCommit)(nil), "types.TransactionToCommit")
	proto.RegisterType((*TransactionListWithProof)(nil), "types.TransactionListWithProof")
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_2cc4e03d2c28c490) }

var fileDescriptor_2cc4e03d2c28c490 = []byte{
	// 1090 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x8e, 0xe3, 0xbf, 0xf8, 0xd8, 0x71, 0x6c, 0xa6, 0x48, 0x94, 0xa0, 0x1b, 0x02, 0x21, 0x5b,
	0xd3, 0x6c, 0x70, 0x06, 0xb7, 0x48, 0x87, 0xa2, 0x37, 0xce, 0x92, 0x26, 0xc1, 0xba, 0x36, 0xa0,
	0x9d, 0x74, 0xdd, 0x8d, 0x40, 0xdb, 0xb4, 0x23, 0x54, 0x12, 0x55, 0x92, 0x8a, 0x6d, 0xec, 0x1d,
	0xfa, 0x06, 0x7b, 0x91, 0xbd, 0xc2, 0x6e, 0x77, 0xbf, 0x57, 0x19, 0x48, 0x51, 0x96, 0x9c, 0x9f,
	0xde, 0x6c, 0x57, 0x16, 0x3f, 0x7e, 0xe7, 0xe8, 0xfc, 0x7d, 0x47, 0x86, 0xa6, 0xe4, 0x24, 0x10,
	0x64, 0x20, 0x5d, 0x16, 0xb4, 0x42, 0xce, 0x24, 0x43, 0x45, 0x39, 0x0b, 0xa9, 0xd8, 0x6e, 0x92,
	0xc1, 0x80, 0x0a, 0xe1, 0x84, 0x44, 0x5e, 0xc7, 0x37, 0xdb, 0x35, 0x7a, 0x43, 0x03, 0x29, 0xcc,
	0xa9, 0x1a, 0x72, 0xc6, 0x46, 0xe6, 0xb0, 0x91, 0xf1, 0xe3, 0xb8, 0xc1, 0x88, 0x19, 0xfc, 0xeb,
	0x31, 0x63, 0x63, 0x8f, 0x1e, 0xe8, 0x53, 0x3f, 0x1a, 0x1d, 0x4c, 0x38, 0x09, 0x43, 0xca, 0x8d,
	0x13, 0xfb, 0x73, 0x1e, 0xea, 0x98, 0x4c, 0x7a, 0xa9, 0x35, 0xfa, 0x06, 0xea, 0x82, 0x06, 0x43,
	0xca, 0x1d, 0x32, 0x18, 0xb0, 0x28, 0x90, 0x56, 0x6e, 0x27, 0xb7, 0x57, 0xc3, 0xab, 0x31, 0xda,
	0x89, 0x41, 0xf4, 0x04, 0xd6, 0x04, 0xfd, 0x14, 0xd1, 0x60, 0x40, 0x9d, 0x20, 0xf2, 0xfb, 0x94,
	0x5b, 0xcb, 0x3b, 0xb9, 0xbd, 0x02, 0xae, 0x27, 0xf0, 0x5b, 0x8d, 0xa2, 0x7d, 0x28, 0x87, 0x9c,
	0x8d, 0x39, 0xf1, 0xad, 0xfc, 0x4e, 0x6e, 0xaf, 0xda, 0xae, 0xb7, 0x74, 0x86, 0xad, 0x8b, 0x18,
	0x3d, 0x5b, 0xc2, 0x09, 0x01, 0xb5, 0xa0, 0x32, 0xe1, 0xae, 0xa4, 0x8e, 0xa0, 0xd2, 0x2a, 0x68,
	0xf6, 0x9a, 0x61, 0xbf, 0x57, 0x78, 0x97, 0xca, 0xb3, 0x25, 0xbc, 0x32, 0x31, 0xcf, 0xe8, 0x09,
	0x94, 0xc4, 0x80, 0xbb, 0xa1, 0xb4, 0x56, 0x34, 0x79, 0xd5, 0x90, 0xbb, 0x1a, 0x3c, 0x5b, 0xc2,
	0xe6, 0x5a, 0x11, 0x7d, 0x36, 0x8c, 0x3c, 0x6a, 0x55, 0x16, 0x88, 0xbf, 0x68, 0x50, 0x11, 0xe3,
	0x6b, 0xb4, 0x0b, 0x75, 0x9f, 0x4c, 0x9d, 0x31, 0x11, 0x0e, 0xf1, 0x75, 0xf6, 0x45, 0x9d, 0x55,
	0xcd, 0x27, 0xd3, 0x53, 0x22, 0x3a, 0x1a, 0x53, 0x2c, 0xc5, 0x88, 0x02, 0x57, 0x3a, 0x21, 0x77,
	0x07, 0xd4, 0x2a, 0xc5, 0xac, 0x31, 0x11, 0x97, 0x81, 0x2b, 0x2f, 0x14, 0xa6, 0x4a, 0x44, 0xa7,
	0xa1, 0xcb, 0x89, 0xee, 0x8a, 0x74, 0x7d, 0x6a, 0x95, 0xe3, 0x12, 0xa5, 0x70, 0xcf, 0xf5, 0xe9,
	0x51, 0x05, 0xca, 0x21, 0x99, 0x79, 0x8c, 0x0c, 0xed, 0x4f, 0x50, 0x36, 0x75, 0x41, 0x08, 0x0a,
	0x03, 0x36, 0xa4, 0xa6, 0xfc, 0xfa, 0x19, 0xfd, 0x08, 0x15, 0xc2, 0xc7, 0x91, 0xaf, 0xe6, 0xc0,
	0x5a, 0xde, 0xc9, 0xef, 0x55, 0xdb, 0xdb, 0x26, 0x95, 0x4c, 0x0f, 0x3b, 0x86, 0x82, 0x53, 0x32,
	0xb2, 0xa0, 0x1c, 0xa7, 0x28, 0xac, 0xfc, 0x4e, 0x7e, 0xaf, 0x86, 0x93, 0xa3, 0x7d, 0x05, 0xa5,
	0xb8, 0x5e, 0xff, 0xef, 0x1b, 0xed, 0x3f, 0x72, 0xb0, 0x7e, 0x0f, 0x05, 0x1d, 0x42, 0x41, 0xd9,
	0xeb, 0xb7, 0xd4, 0xdb, 0xf6, 0xc3, 0xce, 0x5a, 0x1d, 0x3e, 0xee, 0xcd, 0x42, 0x8a, 0x35, 0x5f,
	0x45, 0x37, 0x24, 0x92, 0xe8, 0x31, 0xab, 0x61, 0xfd, 0x6c, 0xbf, 0x84, 0xb2, 0x21, 0xa1, 0x32,
	0xe4, 0x2f, 0x0f, 0x9f, 0x37, 0x96, 0x50, 0x15, 0xca, 0x9d, 0xe3, 0x63, 0x7c, 0xd2, 0xed, 0x36,
	0x72, 0x08, 0xa0, 0xd4, 0xed, 0xe1, 0xf3, 0xb7, 0xa7, 0x8d, 0x65, 0xb4, 0x0a, 0x95, 0xa3, 0x0f,
	0xbd, 0x93, 0x0e, 0xc6, 0x9d, 0x0f, 0x8d, 0xbc, 0xfd, 0x18, 0x4a, 0x71, 0xfb, 0xef, 0xcb, 0xdb,
	0xfe, 0x9c, 0x83, 0x66, 0xd7, 0x1d, 0x07, 0x74, 0x98, 0x15, 0x87, 0x0d, 0xab, 0x9c, 0x4c, 0x1c,
	0x39, 0x0d, 0x9c, 0xfe, 0x4c, 0x52, 0x61, 0x4c, 0xaa, 0x9c, 0x4c, 0x7a, 0xd3, 0xe0, 0x48, 0x41,
	0x68, 0x1f, 0x9a, 0x46, 0x40, 0x61, 0xd4, 0xf7, 0xdc, 0x81, 0xf3, 0x91, 0xce, 0x4c, 0xd0, 0x6b,
	0xf1, 0xc5, 0x85, 0xc6, 0x7f, 0xa6, 0x33, 0xf4, 0x14, 0x1a, 0x86, 0x2b, 0xdc, 0x71, 0x40, 0x64,
	0xc4, 0xa9, 0x56, 0xc9, 0x9c, 0xda, 0x4d, 0x60, 0xfb, 0x9f, 0x1c, 0x6c, 0xdf, 0x09, 0xe8, 0xbd,
	0x2b, 0xaf, 0x2f, 0xd4, 0x1e, 0x50, 0xfd, 0xbd, 0xa1, 0x5c, 0xb8, 0x2c, 0xd0, 0x31, 0x15, 0x70,
	0x72, 0x44, 0xa7, 0x80, 0x84, 0xb6, 0x73, 0x32, 0x4b, 0x42, 0x07, 0x54, 0x6d, 0x5b, 0x89, 0x60,
	0x6e, 0x3b, 0xc6, 0x4d, 0x71, 0x27, 0xf9, 0x67, 0x50, 0xd4, 0x3b, 0xc7, 0xe8, 0xf8, 0xab, 0x87,
	0x6c, 0x75, 0x40, 0x38, 0xe6, 0xa2, 0xa7, 0x50, 0x8a, 0xd7, 0x96, 0xd1, 0x73, 0xd3, 0x58, 0x9d,
	0x68, 0xf0, 0x8d, 0x2b, 0x24, 0x36, 0x04, 0xfb, 0xcf, 0x1c, 0x6c, 0xde, 0x71, 0x26, 0x8e, 0x3c,
	0x36, 0xf8, 0x88, 0x5e, 0x41, 0x2d, 0x13, 0xbd, 0xaa, 0x7b, 0xfe, 0x8b, 0xe1, 0x2f, 0xb0, 0xd1,
	0x0f, 0xf0, 0xe8, 0x86, 0x78, 0xee, 0x90, 0x48, 0x76, 0x4f, 0x57, 0xd0, 0xfc, 0x2e, 0x6d, 0xcc,
	0x01, 0xac, 0xa7, 0x16, 0xb7, 0x7b, 0x93, 0x1a, 0xa4, 0xed, 0x79, 0x01, 0x2b, 0xc9, 0x8a, 0x42,
	0xdf, 0x65, 0xd7, 0x58, 0x1c, 0x69, 0x3d, 0xbb, 0xc6, 0xde, 0x85, 0xe9, 0x0e, 0xb3, 0x7f, 0x87,
	0xb2, 0x01, 0x51, 0x1b, 0xaa, 0x99, 0xad, 0xaf, 0xfb, 0x98, 0x16, 0xac, 0xa3, 0x6f, 0x2e, 0x88,
	0xbc, 0xc6, 0x40, 0xe6, 0xcf, 0xe8, 0x11, 0x14, 0x6f, 0x88, 0x17, 0x51, 0x93, 0x4b, 0x7c, 0x40,
	0xdf, 0x1a, 0x8d, 0xe5, 0xb5, 0xc6, 0xd0, 0xe2, 0xcb, 0x53, 0x4d, 0xd9, 0xaf, 0xa0, 0x66, 0x16,
	0x7a, 0x57, 0x12, 0x49, 0xd5, 0x14, 0x91, 0xe1, 0x90, 0x53, 0x91, 0x4c, 0x76, 0x72, 0x54, 0x1a,
	0xe9, 0x7b, 0xac, 0x9f, 0xa8, 0x4f, 0x3d, 0xdb, 0x7f, 0x2d, 0x2a, 0xbc, 0xc7, 0x7e, 0x62, 0xbe,
	0xef, 0x4a, 0xf4, 0x02, 0x20, 0x99, 0xb8, 0x69, 0x60, 0xd2, 0x78, 0xb8, 0x55, 0x15, 0x33, 0x69,
	0xd3, 0x00, 0xbd, 0x84, 0xba, 0xf9, 0xe8, 0x38, 0x42, 0xc5, 0x93, 0x6c, 0x9c, 0xf5, 0xb4, 0x06,
	0xf3, 0x58, 0xf1, 0x2a, 0xc9, 0x9c, 0x04, 0xda, 0x9d, 0x0f, 0x5a, 0x5e, 0xdb, 0xd4, 0xb2, 0x83,
	0x96, 0xcc, 0x18, 0xda, 0x82, 0x15, 0xbd, 0xb9, 0x05, 0x1d, 0xea, 0x81, 0x2c, 0xe0, 0xb2, 0xda,
	0xd9, 0x82, 0x0e, 0xed, 0xbf, 0xf3, 0x60, 0x65, 0xe2, 0x52, 0xa3, 0x99, 0xca, 0xeb, 0xbf, 0xcd,
	0xdf, 0xf7, 0x50, 0x54, 0x1f, 0xe5, 0x24, 0x9d, 0x8d, 0xbb, 0x3b, 0xef, 0x3c, 0x18, 0x31, 0x1c,
	0x93, 0xd0, 0x19, 0xac, 0xc7, 0xd1, 0x3a, 0x23, 0xc6, 0x1d, 0x23, 0x63, 0x61, 0x54, 0x67, 0x2d,
	0xe8, 0xe7, 0x35, 0xe3, 0x57, 0xe6, 0x1e, 0x37, 0xe9, 0x6d, 0x08, 0xfd, 0x0a, 0x5b, 0x23, 0x97,
	0x0b, 0x99, 0x55, 0x7e, 0xe2, 0xd0, 0xe8, 0xf1, 0x71, 0x2b, 0xfe, 0x8b, 0xd0, 0x4a, 0xfe, 0x22,
	0xb4, 0x2e, 0xcf, 0x03, 0x79, 0xf8, 0xfc, 0x4a, 0xcd, 0x11, 0xde, 0xd4, 0xe6, 0x99, 0x38, 0x8d,
	0x6b, 0xd4, 0x83, 0x6d, 0xad, 0x6f, 0x87, 0x8d, 0x9c, 0x3b, 0xaf, 0xd0, 0xdf, 0xcc, 0x6a, 0x7b,
	0x33, 0xed, 0x5a, 0xe4, 0x47, 0x9e, 0x16, 0x98, 0x5e, 0x0d, 0x9b, 0xda, 0xf4, 0xdd, 0xe8, 0xf5,
	0x2d, 0xe7, 0x08, 0xc3, 0xd6, 0xdc, 0xab, 0x47, 0x6e, 0x39, 0x2d, 0x7d, 0xd9, 0xe9, 0x86, 0x71,
	0xfa, 0x86, 0x2c, 0xf8, 0xdc, 0xdf, 0x85, 0x6a, 0x66, 0xee, 0x51, 0x05, 0x8a, 0xfa, 0xd8, 0x58,
	0x52, 0xdf, 0x86, 0x63, 0xea, 0x51, 0x49, 0x1b, 0xb9, 0xa3, 0x95, 0xdf, 0x4a, 0xd3, 0xb8, 0x02,
	0x25, 0xfd, 0xf3, 0xec, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x78, 0x2d, 0xf4, 0x9b, 0x09,
	0x00, 0x00,
}
