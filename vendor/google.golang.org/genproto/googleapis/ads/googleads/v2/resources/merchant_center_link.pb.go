// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/merchant_center_link.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A data sharing connection, proposed or in use,
// between a Google Ads Customer and a Merchant Center account.
type MerchantCenterLink struct {
	// The resource name of the merchant center link.
	// Merchant center link resource names have the form:
	//
	// `customers/{customer_id}/merchantCenterLinks/{merchant_center_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the Merchant Center account.
	// This field is readonly.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the Merchant Center account.
	// This field is readonly.
	MerchantCenterAccountName *wrappers.StringValue `protobuf:"bytes,4,opt,name=merchant_center_account_name,json=merchantCenterAccountName,proto3" json:"merchant_center_account_name,omitempty"`
	// The status of the link.
	Status               enums.MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                    `json:"-"`
	XXX_unrecognized     []byte                                                      `json:"-"`
	XXX_sizecache        int32                                                       `json:"-"`
}

func (m *MerchantCenterLink) Reset()         { *m = MerchantCenterLink{} }
func (m *MerchantCenterLink) String() string { return proto.CompactTextString(m) }
func (*MerchantCenterLink) ProtoMessage()    {}
func (*MerchantCenterLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a159a5c045d8944, []int{0}
}

func (m *MerchantCenterLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantCenterLink.Unmarshal(m, b)
}
func (m *MerchantCenterLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantCenterLink.Marshal(b, m, deterministic)
}
func (m *MerchantCenterLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantCenterLink.Merge(m, src)
}
func (m *MerchantCenterLink) XXX_Size() int {
	return xxx_messageInfo_MerchantCenterLink.Size(m)
}
func (m *MerchantCenterLink) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantCenterLink.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantCenterLink proto.InternalMessageInfo

func (m *MerchantCenterLink) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *MerchantCenterLink) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MerchantCenterLink) GetMerchantCenterAccountName() *wrappers.StringValue {
	if m != nil {
		return m.MerchantCenterAccountName
	}
	return nil
}

func (m *MerchantCenterLink) GetStatus() enums.MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.MerchantCenterLinkStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*MerchantCenterLink)(nil), "google.ads.googleads.v2.resources.MerchantCenterLink")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/merchant_center_link.proto", fileDescriptor_7a159a5c045d8944)
}

var fileDescriptor_7a159a5c045d8944 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0x6a, 0xdb, 0x30,
	0x18, 0xc7, 0xce, 0x16, 0x98, 0xf7, 0xe7, 0xe0, 0xcb, 0xb2, 0x2c, 0x8c, 0x64, 0x23, 0x10, 0x18,
	0xc8, 0xe0, 0x8d, 0x1d, 0xb4, 0xc1, 0x70, 0xc6, 0x08, 0x2d, 0x6d, 0x09, 0x4e, 0xf1, 0xa1, 0xb8,
	0x18, 0xc5, 0x56, 0x5d, 0x11, 0x5b, 0x32, 0x92, 0x9c, 0xbe, 0x40, 0x9f, 0xa4, 0xc7, 0x5e, 0xfb,
	0x16, 0x7d, 0x94, 0x3e, 0x45, 0x89, 0x64, 0x19, 0xda, 0x34, 0xed, 0xed, 0x4b, 0xbe, 0xdf, 0xdf,
	0xcf, 0x72, 0xfe, 0xe4, 0x8c, 0xe5, 0x05, 0xf6, 0x50, 0x26, 0x3c, 0x3d, 0x6e, 0xa6, 0xb5, 0xef,
	0x71, 0x2c, 0x58, 0xcd, 0x53, 0x2c, 0xbc, 0x12, 0xf3, 0xf4, 0x1c, 0x51, 0x99, 0xa4, 0x98, 0x4a,
	0xcc, 0x93, 0x82, 0xd0, 0x15, 0xa8, 0x38, 0x93, 0xcc, 0x1d, 0x69, 0x0a, 0x40, 0x99, 0x00, 0x2d,
	0x1b, 0xac, 0x7d, 0xd0, 0xb2, 0xfb, 0x7f, 0x77, 0x19, 0x60, 0x5a, 0x97, 0x4f, 0x8b, 0x27, 0x42,
	0x22, 0x59, 0x0b, 0xed, 0xd1, 0xff, 0xd2, 0x08, 0xa8, 0x5f, 0xcb, 0xfa, 0xcc, 0xbb, 0xe0, 0xa8,
	0xaa, 0x30, 0x37, 0xfb, 0x81, 0x31, 0xa8, 0x88, 0x87, 0x28, 0x65, 0x12, 0x49, 0xc2, 0x68, 0xb3,
	0xfd, 0x7a, 0x63, 0x3b, 0xee, 0x61, 0xe3, 0xf1, 0x4f, 0x59, 0x1c, 0x10, 0xba, 0x72, 0xbf, 0x39,
	0xef, 0x4d, 0xc4, 0x84, 0xa2, 0x12, 0xf7, 0xac, 0xa1, 0x35, 0x79, 0x13, 0xbe, 0x33, 0x7f, 0x1e,
	0xa1, 0x12, 0xbb, 0xdf, 0x1d, 0x9b, 0x64, 0xbd, 0xce, 0xd0, 0x9a, 0xbc, 0xf5, 0x3f, 0x37, 0xfd,
	0x80, 0x89, 0x01, 0xf6, 0xa8, 0xfc, 0xf5, 0x33, 0x42, 0x45, 0x8d, 0x43, 0x9b, 0x64, 0xee, 0xa9,
	0x33, 0x78, 0xdc, 0x05, 0xa5, 0x29, 0xab, 0xa9, 0xd4, 0x06, 0xaf, 0x94, 0xcc, 0x60, 0x4b, 0x66,
	0x21, 0x39, 0xa1, 0xb9, 0xd6, 0xf9, 0x54, 0x3e, 0x48, 0x1a, 0x68, 0xbe, 0xca, 0x52, 0x38, 0x5d,
	0x7d, 0x95, 0xde, 0xeb, 0xa1, 0x35, 0xf9, 0xe0, 0x1f, 0x83, 0x5d, 0xa7, 0x57, 0x77, 0x05, 0xdb,
	0x9d, 0x17, 0x8a, 0xfe, 0x9f, 0xd6, 0xe5, 0xce, 0x65, 0xd8, 0x78, 0x4c, 0x2f, 0x6d, 0x67, 0x9c,
	0xb2, 0x12, 0xbc, 0xf8, 0x79, 0xa7, 0x1f, 0xb7, 0xb5, 0xe6, 0x9b, 0x6a, 0x73, 0xeb, 0x64, 0xbf,
	0x61, 0xe7, 0xac, 0x40, 0x34, 0x07, 0x8c, 0xe7, 0x5e, 0x8e, 0xa9, 0x2a, 0x6e, 0x5e, 0x42, 0x45,
	0xc4, 0x33, 0x2f, 0xef, 0x77, 0x3b, 0x5d, 0xd9, 0x9d, 0x59, 0x10, 0x5c, 0xdb, 0xa3, 0x99, 0x96,
	0x0c, 0x32, 0x01, 0xf4, 0xb8, 0x99, 0x22, 0x1f, 0x84, 0x06, 0x79, 0x6b, 0x30, 0x71, 0x90, 0x89,
	0xb8, 0xc5, 0xc4, 0x91, 0x1f, 0xb7, 0x98, 0x3b, 0x7b, 0xac, 0x17, 0x10, 0x06, 0x99, 0x80, 0xb0,
	0x45, 0x41, 0x18, 0xf9, 0x10, 0xb6, 0xb8, 0x65, 0x57, 0x85, 0xfd, 0x71, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0x34, 0xdd, 0x5e, 0x7f, 0x25, 0x03, 0x00, 0x00,
}
