// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/currency_code_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Enum describing possible currency code errors.
type CurrencyCodeErrorEnum_CurrencyCodeError int32

const (
	// Enum unspecified.
	CurrencyCodeErrorEnum_UNSPECIFIED CurrencyCodeErrorEnum_CurrencyCodeError = 0
	// The received error code is not known in this version.
	CurrencyCodeErrorEnum_UNKNOWN CurrencyCodeErrorEnum_CurrencyCodeError = 1
	// The currency code is not supported.
	CurrencyCodeErrorEnum_UNSUPPORTED CurrencyCodeErrorEnum_CurrencyCodeError = 2
)

var CurrencyCodeErrorEnum_CurrencyCodeError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "UNSUPPORTED",
}

var CurrencyCodeErrorEnum_CurrencyCodeError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"UNSUPPORTED": 2,
}

func (x CurrencyCodeErrorEnum_CurrencyCodeError) String() string {
	return proto.EnumName(CurrencyCodeErrorEnum_CurrencyCodeError_name, int32(x))
}

func (CurrencyCodeErrorEnum_CurrencyCodeError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d99fa3ea74d7fca0, []int{0, 0}
}

// Container for enum describing possible currency code errors.
type CurrencyCodeErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrencyCodeErrorEnum) Reset()         { *m = CurrencyCodeErrorEnum{} }
func (m *CurrencyCodeErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CurrencyCodeErrorEnum) ProtoMessage()    {}
func (*CurrencyCodeErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_d99fa3ea74d7fca0, []int{0}
}

func (m *CurrencyCodeErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrencyCodeErrorEnum.Unmarshal(m, b)
}
func (m *CurrencyCodeErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrencyCodeErrorEnum.Marshal(b, m, deterministic)
}
func (m *CurrencyCodeErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrencyCodeErrorEnum.Merge(m, src)
}
func (m *CurrencyCodeErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CurrencyCodeErrorEnum.Size(m)
}
func (m *CurrencyCodeErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrencyCodeErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CurrencyCodeErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.CurrencyCodeErrorEnum_CurrencyCodeError", CurrencyCodeErrorEnum_CurrencyCodeError_name, CurrencyCodeErrorEnum_CurrencyCodeError_value)
	proto.RegisterType((*CurrencyCodeErrorEnum)(nil), "google.ads.googleads.v2.errors.CurrencyCodeErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/currency_code_error.proto", fileDescriptor_d99fa3ea74d7fca0)
}

var fileDescriptor_d99fa3ea74d7fca0 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x23, 0xfd, 0xd4, 0xa2,
	0xa2, 0xfc, 0xa2, 0x62, 0xfd, 0xe4, 0xd2, 0xa2, 0xa2, 0xd4, 0xbc, 0xe4, 0xca, 0xf8, 0xe4, 0xfc,
	0x94, 0xd4, 0x78, 0xb0, 0xa0, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x1c, 0x44, 0xb9, 0x5e,
	0x62, 0x4a, 0xb1, 0x1e, 0x5c, 0xa7, 0x5e, 0x99, 0x91, 0x1e, 0x44, 0xa7, 0x94, 0x0c, 0xcc, 0xe4,
	0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x6e,
	0xa5, 0x68, 0x2e, 0x51, 0x67, 0xa8, 0xd1, 0xce, 0xf9, 0x29, 0xa9, 0xae, 0x20, 0x3d, 0xae, 0x79,
	0xa5, 0xb9, 0x4a, 0x4e, 0x5c, 0x82, 0x18, 0x12, 0x42, 0xfc, 0x5c, 0xdc, 0xa1, 0x7e, 0xc1, 0x01,
	0xae, 0xce, 0x9e, 0x6e, 0x9e, 0xae, 0x2e, 0x02, 0x0c, 0x42, 0xdc, 0x5c, 0xec, 0xa1, 0x7e, 0xde,
	0x7e, 0xfe, 0xe1, 0x7e, 0x02, 0x8c, 0x50, 0xd9, 0xd0, 0x80, 0x00, 0xff, 0xa0, 0x10, 0x57, 0x17,
	0x01, 0x26, 0xa7, 0x8f, 0x8c, 0x5c, 0x4a, 0xc9, 0xf9, 0xb9, 0x7a, 0xf8, 0x5d, 0xe8, 0x24, 0x86,
	0x61, 0x51, 0x00, 0xc8, 0x6d, 0x01, 0x8c, 0x51, 0x2e, 0x50, 0x9d, 0xe9, 0xf9, 0x39, 0x89, 0x79,
	0xe9, 0x7a, 0xf9, 0x45, 0xe9, 0xfa, 0xe9, 0xa9, 0x79, 0x60, 0x97, 0xc3, 0x42, 0xa9, 0x20, 0xb3,
	0x18, 0x57, 0xa0, 0x59, 0x43, 0xa8, 0x45, 0x4c, 0xcc, 0xee, 0x8e, 0x8e, 0xab, 0x98, 0xe4, 0xdc,
	0x21, 0x86, 0x39, 0xa6, 0x14, 0xeb, 0x41, 0x98, 0x20, 0x56, 0x98, 0x91, 0x1e, 0xd8, 0xca, 0xe2,
	0x53, 0x30, 0x05, 0x31, 0x8e, 0x29, 0xc5, 0x31, 0x70, 0x05, 0x31, 0x61, 0x46, 0x31, 0x10, 0x05,
	0xaf, 0x98, 0x94, 0x20, 0xa2, 0x56, 0x56, 0x8e, 0x29, 0xc5, 0x56, 0x56, 0x70, 0x25, 0x56, 0x56,
	0x61, 0x46, 0x56, 0x56, 0x10, 0x45, 0x49, 0x6c, 0x60, 0xd7, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x5f, 0x5e, 0xbb, 0x68, 0xd1, 0x01, 0x00, 0x00,
}
