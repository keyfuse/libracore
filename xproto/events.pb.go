// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events.proto

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

// An event emitted from a smart contract
type Event struct {
	AccessPath           *AccessPath `protobuf:"bytes,1,opt,name=access_path,json=accessPath,proto3" json:"access_path,omitempty"`
	SequenceNumber       uint64      `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	EventData            []byte      `protobuf:"bytes,3,opt,name=event_data,json=eventData,proto3" json:"event_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetAccessPath() *AccessPath {
	if m != nil {
		return m.AccessPath
	}
	return nil
}

func (m *Event) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *Event) GetEventData() []byte {
	if m != nil {
		return m.EventData
	}
	return nil
}

// An event along with the proof for the event
type EventWithProof struct {
	TransactionVersion   uint64      `protobuf:"varint,1,opt,name=transaction_version,json=transactionVersion,proto3" json:"transaction_version,omitempty"`
	EventIndex           uint64      `protobuf:"varint,2,opt,name=event_index,json=eventIndex,proto3" json:"event_index,omitempty"`
	Event                *Event      `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	Proof                *EventProof `protobuf:"bytes,4,opt,name=proof,proto3" json:"proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EventWithProof) Reset()         { *m = EventWithProof{} }
func (m *EventWithProof) String() string { return proto.CompactTextString(m) }
func (*EventWithProof) ProtoMessage()    {}
func (*EventWithProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{1}
}

func (m *EventWithProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventWithProof.Unmarshal(m, b)
}
func (m *EventWithProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventWithProof.Marshal(b, m, deterministic)
}
func (m *EventWithProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventWithProof.Merge(m, src)
}
func (m *EventWithProof) XXX_Size() int {
	return xxx_messageInfo_EventWithProof.Size(m)
}
func (m *EventWithProof) XXX_DiscardUnknown() {
	xxx_messageInfo_EventWithProof.DiscardUnknown(m)
}

var xxx_messageInfo_EventWithProof proto.InternalMessageInfo

func (m *EventWithProof) GetTransactionVersion() uint64 {
	if m != nil {
		return m.TransactionVersion
	}
	return 0
}

func (m *EventWithProof) GetEventIndex() uint64 {
	if m != nil {
		return m.EventIndex
	}
	return 0
}

func (m *EventWithProof) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *EventWithProof) GetProof() *EventProof {
	if m != nil {
		return m.Proof
	}
	return nil
}

// A list of events.
type EventsList struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventsList) Reset()         { *m = EventsList{} }
func (m *EventsList) String() string { return proto.CompactTextString(m) }
func (*EventsList) ProtoMessage()    {}
func (*EventsList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{2}
}

func (m *EventsList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventsList.Unmarshal(m, b)
}
func (m *EventsList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventsList.Marshal(b, m, deterministic)
}
func (m *EventsList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsList.Merge(m, src)
}
func (m *EventsList) XXX_Size() int {
	return xxx_messageInfo_EventsList.Size(m)
}
func (m *EventsList) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsList.DiscardUnknown(m)
}

var xxx_messageInfo_EventsList proto.InternalMessageInfo

func (m *EventsList) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

// A list of EventList's, each representing all events for a transaction.
type EventsForVersions struct {
	EventsForVersion     []*EventsList `protobuf:"bytes,1,rep,name=events_for_version,json=eventsForVersion,proto3" json:"events_for_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EventsForVersions) Reset()         { *m = EventsForVersions{} }
func (m *EventsForVersions) String() string { return proto.CompactTextString(m) }
func (*EventsForVersions) ProtoMessage()    {}
func (*EventsForVersions) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{3}
}

func (m *EventsForVersions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventsForVersions.Unmarshal(m, b)
}
func (m *EventsForVersions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventsForVersions.Marshal(b, m, deterministic)
}
func (m *EventsForVersions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsForVersions.Merge(m, src)
}
func (m *EventsForVersions) XXX_Size() int {
	return xxx_messageInfo_EventsForVersions.Size(m)
}
func (m *EventsForVersions) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsForVersions.DiscardUnknown(m)
}

var xxx_messageInfo_EventsForVersions proto.InternalMessageInfo

func (m *EventsForVersions) GetEventsForVersion() []*EventsList {
	if m != nil {
		return m.EventsForVersion
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "types.Event")
	proto.RegisterType((*EventWithProof)(nil), "types.EventWithProof")
	proto.RegisterType((*EventsList)(nil), "types.EventsList")
	proto.RegisterType((*EventsForVersions)(nil), "types.EventsForVersions")
}

func init() { proto.RegisterFile("events.proto", fileDescriptor_8f22242cb04491f9) }

var fileDescriptor_8f22242cb04491f9 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x51, 0x4b, 0xc3, 0x30,
	0x10, 0xc7, 0x89, 0x5b, 0x87, 0x5e, 0xc7, 0x74, 0xe7, 0x4b, 0x19, 0x88, 0xa5, 0x08, 0xdb, 0x53,
	0x85, 0xfa, 0x01, 0x44, 0x51, 0x41, 0x10, 0x19, 0x41, 0x14, 0x7c, 0x29, 0x59, 0x97, 0xd1, 0x3e,
	0x98, 0xd4, 0x24, 0x1b, 0xf3, 0xd9, 0x0f, 0xe4, 0x57, 0x94, 0x5e, 0xaa, 0xd4, 0xf9, 0xd4, 0xe6,
	0xff, 0xcb, 0xdd, 0xff, 0x7f, 0x17, 0x18, 0xca, 0x8d, 0x54, 0xce, 0xa6, 0xb5, 0xd1, 0x4e, 0x63,
	0xe0, 0x3e, 0x6a, 0x69, 0x27, 0x63, 0x51, 0x14, 0xd2, 0xda, 0xbc, 0x16, 0xae, 0xf4, 0x64, 0x12,
	0xd6, 0x46, 0xeb, 0x95, 0x3f, 0x24, 0x9f, 0x0c, 0x82, 0xdb, 0xa6, 0x0e, 0x33, 0x08, 0x3b, 0x77,
	0x23, 0x16, 0xb3, 0x59, 0x98, 0x8d, 0x53, 0x6a, 0x93, 0x5e, 0x11, 0x99, 0x0b, 0x57, 0x72, 0x10,
	0xbf, 0xff, 0x38, 0x85, 0x43, 0x2b, 0xdf, 0xd7, 0x52, 0x15, 0x32, 0x57, 0xeb, 0xb7, 0x85, 0x34,
	0xd1, 0x5e, 0xcc, 0x66, 0x7d, 0x3e, 0xfa, 0x91, 0x1f, 0x49, 0xc5, 0x13, 0x00, 0x4a, 0x97, 0x2f,
	0x85, 0x13, 0x51, 0x2f, 0x66, 0xb3, 0x21, 0x3f, 0x20, 0xe5, 0x46, 0x38, 0x91, 0x7c, 0x31, 0x18,
	0x51, 0x8a, 0x97, 0xca, 0x95, 0xf3, 0x26, 0x1e, 0x9e, 0xc3, 0xb1, 0x33, 0x42, 0x59, 0x51, 0xb8,
	0x4a, 0xab, 0x7c, 0x23, 0x8d, 0xad, 0xb4, 0xa2, 0x58, 0x7d, 0x8e, 0x1d, 0xf4, 0xec, 0x09, 0x9e,
	0x42, 0xe8, 0x2d, 0x2a, 0xb5, 0x94, 0xdb, 0x36, 0x87, 0x77, 0xbd, 0x6f, 0x14, 0x4c, 0x20, 0xa0,
	0x13, 0xd9, 0x87, 0xd9, 0xb0, 0x1d, 0x8d, 0x7c, 0xb9, 0x47, 0x38, 0x85, 0x80, 0xb6, 0x13, 0xf5,
	0xff, 0x8c, 0x4f, 0x77, 0x28, 0x17, 0xf7, 0x3c, 0xc9, 0x00, 0x48, 0xb4, 0x0f, 0x95, 0x75, 0x78,
	0x06, 0x03, 0xbf, 0xfc, 0x88, 0xc5, 0xbd, 0x7f, 0xbd, 0x5b, 0x96, 0x3c, 0xc1, 0xd8, 0xd7, 0xdc,
	0x69, 0xd3, 0xa6, 0xb6, 0x78, 0x09, 0xe8, 0x71, 0xbe, 0xd2, 0xa6, 0x33, 0x66, 0x6f, 0xd7, 0x9e,
	0x9c, 0xf8, 0x91, 0xdc, 0xe9, 0x70, 0xbd, 0xff, 0x3a, 0xd8, 0xd2, 0x5b, 0x2e, 0x06, 0xf4, 0xb9,
	0xf8, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x57, 0xbc, 0x59, 0xec, 0x09, 0x02, 0x00, 0x00,
}
