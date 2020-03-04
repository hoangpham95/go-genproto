// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/mobile_app_category_constant.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A mobile application category constant.
type MobileAppCategoryConstant struct {
	// The resource name of the mobile app category constant.
	// Mobile app category constant resource names have the form:
	//
	// `mobileAppCategoryConstants/{mobile_app_category_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the mobile app category constant.
	Id *wrappers.Int32Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Mobile app category name.
	Name                 *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MobileAppCategoryConstant) Reset()         { *m = MobileAppCategoryConstant{} }
func (m *MobileAppCategoryConstant) String() string { return proto.CompactTextString(m) }
func (*MobileAppCategoryConstant) ProtoMessage()    {}
func (*MobileAppCategoryConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_14ce0b7d234fc1da, []int{0}
}

func (m *MobileAppCategoryConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileAppCategoryConstant.Unmarshal(m, b)
}
func (m *MobileAppCategoryConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileAppCategoryConstant.Marshal(b, m, deterministic)
}
func (m *MobileAppCategoryConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileAppCategoryConstant.Merge(m, src)
}
func (m *MobileAppCategoryConstant) XXX_Size() int {
	return xxx_messageInfo_MobileAppCategoryConstant.Size(m)
}
func (m *MobileAppCategoryConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileAppCategoryConstant.DiscardUnknown(m)
}

var xxx_messageInfo_MobileAppCategoryConstant proto.InternalMessageInfo

func (m *MobileAppCategoryConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *MobileAppCategoryConstant) GetId() *wrappers.Int32Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MobileAppCategoryConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func init() {
	proto.RegisterType((*MobileAppCategoryConstant)(nil), "google.ads.googleads.v3.resources.MobileAppCategoryConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/mobile_app_category_constant.proto", fileDescriptor_14ce0b7d234fc1da)
}

var fileDescriptor_14ce0b7d234fc1da = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x6a, 0xdb, 0x30,
	0x1c, 0xc6, 0xb1, 0x32, 0x06, 0xf3, 0xb6, 0x8b, 0x4f, 0x49, 0x16, 0x42, 0xb2, 0x11, 0x08, 0x0c,
	0xa4, 0x11, 0x9f, 0xa6, 0x9d, 0x9c, 0x0c, 0xc2, 0x06, 0x1b, 0x21, 0x03, 0x1f, 0x86, 0xc1, 0x28,
	0xb6, 0x2a, 0x0c, 0xb6, 0x24, 0x24, 0x25, 0xa5, 0x94, 0x3e, 0x41, 0x8f, 0x7d, 0x83, 0x1e, 0xfb,
	0x28, 0x7d, 0x94, 0x3c, 0x45, 0x89, 0x2d, 0x2b, 0x85, 0x92, 0xf4, 0xf6, 0xd9, 0xdf, 0x4f, 0xdf,
	0xff, 0xfb, 0x5b, 0xf6, 0x7f, 0x32, 0x21, 0x58, 0x49, 0x11, 0xc9, 0x35, 0x6a, 0xe4, 0x41, 0xed,
	0x42, 0xa4, 0xa8, 0x16, 0x5b, 0x95, 0x51, 0x8d, 0x2a, 0xb1, 0x29, 0x4a, 0x9a, 0x12, 0x29, 0xd3,
	0x8c, 0x18, 0xca, 0x84, 0xba, 0x4a, 0x33, 0xc1, 0xb5, 0x21, 0xdc, 0x40, 0xa9, 0x84, 0x11, 0xc1,
	0xb8, 0x39, 0x0a, 0x49, 0xae, 0xa1, 0x4b, 0x81, 0xbb, 0x10, 0xba, 0x94, 0x7e, 0xaf, 0x1d, 0x24,
	0x0b, 0x97, 0xdd, 0x9c, 0xee, 0x0f, 0xad, 0x55, 0x3f, 0x6d, 0xb6, 0x17, 0xe8, 0x52, 0x11, 0x29,
	0xa9, 0xd2, 0xd6, 0x1f, 0x3c, 0x3b, 0x4a, 0x38, 0x17, 0x86, 0x98, 0x42, 0x70, 0xeb, 0x7e, 0xbe,
	0x03, 0x7e, 0xef, 0x4f, 0x5d, 0x31, 0x92, 0x72, 0x61, 0x0b, 0x2e, 0x6c, 0xbf, 0xe0, 0x8b, 0xff,
	0xb1, 0x9d, 0x96, 0x72, 0x52, 0xd1, 0xae, 0x37, 0xf2, 0xa6, 0xef, 0xd6, 0x1f, 0xda, 0x97, 0x7f,
	0x49, 0x45, 0x83, 0xaf, 0x3e, 0x28, 0xf2, 0x2e, 0x18, 0x79, 0xd3, 0xf7, 0xb3, 0x4f, 0x76, 0x01,
	0xd8, 0xb6, 0x81, 0xbf, 0xb8, 0x09, 0x67, 0x31, 0x29, 0xb7, 0x74, 0x0d, 0x8a, 0x3c, 0xf8, 0xe6,
	0xbf, 0xa9, 0x83, 0x3a, 0x35, 0x3e, 0x78, 0x81, 0xff, 0x33, 0xaa, 0xe0, 0xac, 0xe1, 0x6b, 0x12,
	0xab, 0x7d, 0x24, 0xfc, 0xd9, 0xf1, 0xbb, 0x58, 0x25, 0x0b, 0x0d, 0x33, 0x51, 0xa1, 0xd3, 0xe5,
	0xbf, 0x57, 0xa7, 0x2c, 0x8d, 0xae, 0xcf, 0x5d, 0xcb, 0xcd, 0xfc, 0x16, 0xf8, 0x93, 0x4c, 0x54,
	0xf0, 0xd5, 0x8b, 0x99, 0x0f, 0x4f, 0xce, 0x5f, 0x1d, 0x56, 0x5a, 0x79, 0xff, 0x7f, 0xdb, 0x10,
	0x26, 0x4a, 0xc2, 0x19, 0x14, 0x8a, 0x21, 0x46, 0x79, 0xbd, 0x30, 0x3a, 0xee, 0x72, 0xe6, 0x17,
	0xfa, 0xe1, 0xd4, 0x3d, 0xe8, 0x2c, 0xa3, 0xe8, 0x01, 0x8c, 0x97, 0x4d, 0x64, 0x94, 0x6b, 0xd8,
	0xc8, 0x83, 0x8a, 0x43, 0xb8, 0x6e, 0xc9, 0xc7, 0x96, 0x49, 0xa2, 0x5c, 0x27, 0x8e, 0x49, 0xe2,
	0x30, 0x71, 0xcc, 0x1e, 0x4c, 0x1a, 0x03, 0xe3, 0x28, 0xd7, 0x18, 0x3b, 0x0a, 0xe3, 0x38, 0xc4,
	0xd8, 0x71, 0x9b, 0xb7, 0x75, 0xd9, 0xf0, 0x29, 0x00, 0x00, 0xff, 0xff, 0x84, 0x0f, 0xa0, 0x59,
	0xee, 0x02, 0x00, 0x00,
}
