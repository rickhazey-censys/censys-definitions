// Code generated by protoc-gen-go. DO NOT EDIT.
// source: caa.proto

/*
Package censys_definitions is a generated protocol buffer package.

It is generated from these files:
	caa.proto
	rpc.proto
	protocols.proto
	hoststore.proto
	ct.proto
	search.proto
	anonstore.proto
	pubkey.proto
	certificate.proto
	zlint.proto
	common.proto

It has these top-level messages:
	CAATagValue
	CAARecord
	CAALookup
	MinScanId
	MozillaOneCRLEntry
	Command
	AnonymousStoreStatistics
	StatisticsPair
	StoreStatistics
	ServerStatistics
	PruneStatistics
	CommandReply
	HostQuery
	HostQueryResponse
	AnonymousQuery
	AnonymousQueryResponse
	UserDataRequest
	RootStoreQuery
	RootStoreReply
	WHOISAtom
	LocationAtom
	ProtocolAtom
	AnonymousKey
	Record
	Delta
	CTServerStatus
	CTStatus
	SCT
	AnonymousRecord
	AnonymousDelta
	ExternalCertificate
	RSACryptographicKey
	DSACryptographicKey
	ECCCryptographicKey
	CryptographicKey
	Path
	RootStoreStatus
	CertificateValidation
	MozillaSalesForceStatus
	CertificateRevocation
	CertificateAudit
	Certificate
	LintResult
	ZLint
	Lints
	Metadatum
	UserdataAtom
	ASAtom
*/
package censys_definitions

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CAAResult int32

const (
	CAAResult_CAA_RESULT_RESERVED                CAAResult = 0
	CAAResult_CAA_RESULT_VALIDATION_SUCCESS      CAAResult = 1
	CAAResult_CAA_RESULT_VALIDATION_FAIL         CAAResult = 2
	CAAResult_CAA_RESULT_VALIDATION_SKIPPED      CAAResult = 3
	CAAResult_CAA_RESULT_VALIDATION_NOT_REQUIRED CAAResult = 4
)

var CAAResult_name = map[int32]string{
	0: "CAA_RESULT_RESERVED",
	1: "CAA_RESULT_VALIDATION_SUCCESS",
	2: "CAA_RESULT_VALIDATION_FAIL",
	3: "CAA_RESULT_VALIDATION_SKIPPED",
	4: "CAA_RESULT_VALIDATION_NOT_REQUIRED",
}
var CAAResult_value = map[string]int32{
	"CAA_RESULT_RESERVED":                0,
	"CAA_RESULT_VALIDATION_SUCCESS":      1,
	"CAA_RESULT_VALIDATION_FAIL":         2,
	"CAA_RESULT_VALIDATION_SKIPPED":      3,
	"CAA_RESULT_VALIDATION_NOT_REQUIRED": 4,
}

func (x CAAResult) String() string {
	return proto.EnumName(CAAResult_name, int32(x))
}
func (CAAResult) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CAADomainStatus int32

const (
	CAADomainStatus_CAA_DOMAIN_STATUS_RESERVED                CAADomainStatus = 0
	CAADomainStatus_CAA_DOMAIN_STATUS_VALIDATION_SUCCESS      CAADomainStatus = 1
	CAADomainStatus_CAA_DOMAIN_STATUS_VALIDATION_FAIL         CAADomainStatus = 2
	CAADomainStatus_CAA_DOMAIN_STATUS_VALIDATION_SKIPPED      CAADomainStatus = 3
	CAADomainStatus_CAA_DOMAIN_STATUS_DNS_ERROR               CAADomainStatus = 5
	CAADomainStatus_CAA_DOMAIN_STATUS_DNS_ERROR_SERVFAIL      CAADomainStatus = 6
	CAADomainStatus_CAA_DOMAIN_STATUS_DNS_ERROR_AUTHFAIL      CAADomainStatus = 7
	CAADomainStatus_CAA_DOMAIN_STATUS_DNS_ERROR_NO_RECORD     CAADomainStatus = 8
	CAADomainStatus_CAA_DOMAIN_STATUS_DNS_ERROR_BLACKLIST     CAADomainStatus = 9
	CAADomainStatus_CAA_DOMAIN_STATUS_DNS_ERROR_NO_OUTPUT     CAADomainStatus = 10
	CAADomainStatus_CAA_DOMAIN_STATUS_DNS_ERROR_NO_ANSWER     CAADomainStatus = 11
	CAADomainStatus_CAA_DOMAIN_STATUS_DNS_ERROR_ILLEGAL_INPUT CAADomainStatus = 12
	CAADomainStatus_CAA_DOMAIN_STATUS_DNS_ERROR_TIMEOUT       CAADomainStatus = 13
	CAADomainStatus_CAA_DOMAIN_STATUS_DNS_ERROR_ITER_TIMEOUT  CAADomainStatus = 14
	CAADomainStatus_CAA_DOMAIN_STATUS_DNS_ERROR_TEMPORARY     CAADomainStatus = 15
	CAADomainStatus_CAA_DOMAIN_STATUS_DNS_ERROR_TRUNCATED     CAADomainStatus = 16
)

var CAADomainStatus_name = map[int32]string{
	0:  "CAA_DOMAIN_STATUS_RESERVED",
	1:  "CAA_DOMAIN_STATUS_VALIDATION_SUCCESS",
	2:  "CAA_DOMAIN_STATUS_VALIDATION_FAIL",
	3:  "CAA_DOMAIN_STATUS_VALIDATION_SKIPPED",
	5:  "CAA_DOMAIN_STATUS_DNS_ERROR",
	6:  "CAA_DOMAIN_STATUS_DNS_ERROR_SERVFAIL",
	7:  "CAA_DOMAIN_STATUS_DNS_ERROR_AUTHFAIL",
	8:  "CAA_DOMAIN_STATUS_DNS_ERROR_NO_RECORD",
	9:  "CAA_DOMAIN_STATUS_DNS_ERROR_BLACKLIST",
	10: "CAA_DOMAIN_STATUS_DNS_ERROR_NO_OUTPUT",
	11: "CAA_DOMAIN_STATUS_DNS_ERROR_NO_ANSWER",
	12: "CAA_DOMAIN_STATUS_DNS_ERROR_ILLEGAL_INPUT",
	13: "CAA_DOMAIN_STATUS_DNS_ERROR_TIMEOUT",
	14: "CAA_DOMAIN_STATUS_DNS_ERROR_ITER_TIMEOUT",
	15: "CAA_DOMAIN_STATUS_DNS_ERROR_TEMPORARY",
	16: "CAA_DOMAIN_STATUS_DNS_ERROR_TRUNCATED",
}
var CAADomainStatus_value = map[string]int32{
	"CAA_DOMAIN_STATUS_RESERVED":                0,
	"CAA_DOMAIN_STATUS_VALIDATION_SUCCESS":      1,
	"CAA_DOMAIN_STATUS_VALIDATION_FAIL":         2,
	"CAA_DOMAIN_STATUS_VALIDATION_SKIPPED":      3,
	"CAA_DOMAIN_STATUS_DNS_ERROR":               5,
	"CAA_DOMAIN_STATUS_DNS_ERROR_SERVFAIL":      6,
	"CAA_DOMAIN_STATUS_DNS_ERROR_AUTHFAIL":      7,
	"CAA_DOMAIN_STATUS_DNS_ERROR_NO_RECORD":     8,
	"CAA_DOMAIN_STATUS_DNS_ERROR_BLACKLIST":     9,
	"CAA_DOMAIN_STATUS_DNS_ERROR_NO_OUTPUT":     10,
	"CAA_DOMAIN_STATUS_DNS_ERROR_NO_ANSWER":     11,
	"CAA_DOMAIN_STATUS_DNS_ERROR_ILLEGAL_INPUT": 12,
	"CAA_DOMAIN_STATUS_DNS_ERROR_TIMEOUT":       13,
	"CAA_DOMAIN_STATUS_DNS_ERROR_ITER_TIMEOUT":  14,
	"CAA_DOMAIN_STATUS_DNS_ERROR_TEMPORARY":     15,
	"CAA_DOMAIN_STATUS_DNS_ERROR_TRUNCATED":     16,
}

func (x CAADomainStatus) String() string {
	return proto.EnumName(CAADomainStatus_name, int32(x))
}
func (CAADomainStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CAATagValue struct {
	// flag is one byte, this is the cheapest way to store it
	Flag  uint32 `protobuf:"varint,1,opt,name=flag" json:"flag,omitempty"`
	Tag   string `protobuf:"bytes,2,opt,name=tag" json:"tag,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	Ttl   uint32 `protobuf:"varint,4,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *CAATagValue) Reset()                    { *m = CAATagValue{} }
func (m *CAATagValue) String() string            { return proto.CompactTextString(m) }
func (*CAATagValue) ProtoMessage()               {}
func (*CAATagValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CAATagValue) GetFlag() uint32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *CAATagValue) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *CAATagValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *CAATagValue) GetTtl() uint32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type CAARecord struct {
	Domain string          `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	Result CAADomainStatus `protobuf:"varint,2,opt,name=result,enum=zsearch.CAADomainStatus" json:"result,omitempty"`
	Values []*CAATagValue  `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
}

func (m *CAARecord) Reset()                    { *m = CAARecord{} }
func (m *CAARecord) String() string            { return proto.CompactTextString(m) }
func (*CAARecord) ProtoMessage()               {}
func (*CAARecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CAARecord) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *CAARecord) GetResult() CAADomainStatus {
	if m != nil {
		return m.Result
	}
	return CAADomainStatus_CAA_DOMAIN_STATUS_RESERVED
}

func (m *CAARecord) GetValues() []*CAATagValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type CAALookup struct {
	Timestamp int64        `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Records   []*CAARecord `protobuf:"bytes,2,rep,name=records" json:"records,omitempty"`
	Result    CAAResult    `protobuf:"varint,3,opt,name=result,enum=zsearch.CAAResult" json:"result,omitempty"`
}

func (m *CAALookup) Reset()                    { *m = CAALookup{} }
func (m *CAALookup) String() string            { return proto.CompactTextString(m) }
func (*CAALookup) ProtoMessage()               {}
func (*CAALookup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CAALookup) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *CAALookup) GetRecords() []*CAARecord {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *CAALookup) GetResult() CAAResult {
	if m != nil {
		return m.Result
	}
	return CAAResult_CAA_RESULT_RESERVED
}

func init() {
	proto.RegisterType((*CAATagValue)(nil), "zsearch.CAATagValue")
	proto.RegisterType((*CAARecord)(nil), "zsearch.CAARecord")
	proto.RegisterType((*CAALookup)(nil), "zsearch.CAALookup")
	proto.RegisterEnum("zsearch.CAAResult", CAAResult_name, CAAResult_value)
	proto.RegisterEnum("zsearch.CAADomainStatus", CAADomainStatus_name, CAADomainStatus_value)
}

func init() { proto.RegisterFile("caa.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdd, 0x6e, 0xd3, 0x40,
	0x10, 0x85, 0x49, 0x9d, 0x1f, 0x3c, 0xa1, 0xed, 0x6a, 0xa9, 0xc0, 0xe2, 0x37, 0x0d, 0x14, 0xd2,
	0x28, 0x44, 0xa8, 0x3c, 0xc1, 0x62, 0x2f, 0x60, 0xd5, 0xb1, 0xc3, 0xec, 0x3a, 0x08, 0x71, 0x61,
	0x2d, 0x69, 0x08, 0x15, 0x09, 0x8e, 0x62, 0x87, 0x0b, 0x6e, 0x79, 0x1a, 0x6e, 0x79, 0x42, 0xe4,
	0x6d, 0x9c, 0x46, 0x6d, 0xe4, 0xf4, 0xca, 0xb3, 0x3b, 0xdf, 0x9c, 0x39, 0xc7, 0x96, 0x0c, 0xe6,
	0x50, 0xa9, 0xee, 0x6c, 0x1e, 0xa7, 0x31, 0xad, 0xfd, 0x4e, 0x46, 0x6a, 0x3e, 0xfc, 0xde, 0xfc,
	0x02, 0x75, 0x9b, 0x31, 0xa9, 0xc6, 0x03, 0x35, 0x59, 0x8c, 0x28, 0x85, 0xf2, 0xb7, 0x89, 0x1a,
	0x5b, 0xa5, 0x46, 0xa9, 0xb5, 0x8b, 0xba, 0xa6, 0x04, 0x8c, 0x54, 0x8d, 0xad, 0x9d, 0x46, 0xa9,
	0x65, 0x62, 0x56, 0xd2, 0x03, 0xa8, 0xfc, 0xca, 0x70, 0xcb, 0xd0, 0x77, 0x17, 0x07, 0xcd, 0xa5,
	0x13, 0xab, 0xac, 0x47, 0xb3, 0xb2, 0xf9, 0xa7, 0x04, 0xa6, 0xcd, 0x18, 0x8e, 0x86, 0xf1, 0xfc,
	0x8c, 0xde, 0x83, 0xea, 0x59, 0x3c, 0x55, 0xe7, 0x3f, 0xb5, 0xba, 0x89, 0xcb, 0x13, 0x7d, 0x0d,
	0xd5, 0xf9, 0x28, 0x59, 0x4c, 0x52, 0xbd, 0x62, 0xef, 0xc4, 0xea, 0x2e, 0xcd, 0x75, 0x6d, 0xc6,
	0x1c, 0xcd, 0x88, 0x54, 0xa5, 0x8b, 0x04, 0x97, 0x1c, 0xed, 0x40, 0x55, 0xaf, 0x4c, 0x2c, 0xa3,
	0x61, 0xb4, 0xea, 0x27, 0x07, 0xeb, 0x13, 0x79, 0x16, 0x5c, 0x32, 0xb9, 0x0b, 0x2f, 0x8e, 0x7f,
	0x2c, 0x66, 0xf4, 0x11, 0x98, 0xe9, 0xf9, 0x74, 0x94, 0xa4, 0x6a, 0x3a, 0xd3, 0x46, 0x0c, 0xbc,
	0xbc, 0xa0, 0x1d, 0xa8, 0xcd, 0xb5, 0xdb, 0xc4, 0xda, 0xd1, 0xd2, 0x74, 0x5d, 0xfa, 0x22, 0x08,
	0xe6, 0x08, 0x6d, 0xaf, 0x9c, 0x1b, 0xda, 0xf9, 0x15, 0x38, 0xeb, 0xe4, 0x9e, 0xdb, 0xff, 0xf2,
	0x77, 0xa1, 0x13, 0xdc, 0x87, 0xbb, 0x36, 0x63, 0x11, 0x72, 0x11, 0x7a, 0x32, 0x7b, 0x70, 0x1c,
	0x70, 0x87, 0xdc, 0xa2, 0x87, 0xf0, 0x78, 0xad, 0x31, 0x60, 0x9e, 0xeb, 0x30, 0xe9, 0x06, 0x7e,
	0x24, 0x42, 0xdb, 0xe6, 0x42, 0x90, 0x12, 0x7d, 0x02, 0x0f, 0x36, 0x23, 0xef, 0x98, 0xeb, 0x91,
	0x9d, 0x02, 0x89, 0x53, 0xb7, 0xdf, 0xe7, 0x0e, 0x31, 0xe8, 0x0b, 0x68, 0x6e, 0x46, 0xfc, 0x20,
	0x73, 0xf3, 0x31, 0x74, 0x91, 0x3b, 0xa4, 0xdc, 0xfe, 0x5b, 0x81, 0xfd, 0x2b, 0x1f, 0x21, 0x5f,
	0xef, 0x04, 0x3d, 0xe6, 0xfa, 0x91, 0x90, 0x4c, 0x86, 0x62, 0x3d, 0x41, 0x0b, 0x9e, 0x5f, 0xef,
	0x6f, 0x0c, 0x72, 0x04, 0x87, 0x85, 0xe4, 0x32, 0xcf, 0x56, 0xc1, 0x55, 0xac, 0xa7, 0xf0, 0xf0,
	0x3a, 0xe9, 0xf8, 0x22, 0xe2, 0x88, 0x01, 0x92, 0xca, 0x66, 0xa9, 0x15, 0x10, 0x65, 0x19, 0xf4,
	0xd2, 0xea, 0x36, 0x92, 0x85, 0xf2, 0x83, 0x26, 0x6b, 0xf4, 0x18, 0x8e, 0x8a, 0x48, 0x3f, 0x88,
	0x90, 0xdb, 0x01, 0x3a, 0xe4, 0xf6, 0x36, 0xf4, 0xad, 0xc7, 0xec, 0x53, 0xcf, 0x15, 0x92, 0x98,
	0x37, 0x50, 0x0d, 0x42, 0xd9, 0x0f, 0x25, 0x81, 0x1b, 0xa0, 0xcc, 0x17, 0x9f, 0x38, 0x92, 0x3a,
	0x7d, 0x05, 0xc7, 0x45, 0xa8, 0xeb, 0x79, 0xfc, 0x3d, 0xf3, 0x22, 0xd7, 0xcf, 0x94, 0xef, 0xd0,
	0x97, 0xf0, 0xac, 0x08, 0x97, 0x6e, 0x8f, 0x07, 0xa1, 0x24, 0xbb, 0xb4, 0x03, 0xad, 0x42, 0x5d,
	0xc9, 0x2f, 0xe9, 0xbd, 0x6d, 0x86, 0x25, 0xef, 0xf5, 0x03, 0x64, 0xf8, 0x99, 0xec, 0x6f, 0x45,
	0x31, 0xf4, 0x6d, 0x26, 0xb9, 0x43, 0xc8, 0xd7, 0xaa, 0xfe, 0xb3, 0xbd, 0xf9, 0x1f, 0x00, 0x00,
	0xff, 0xff, 0x99, 0x19, 0x0c, 0x70, 0xe6, 0x04, 0x00, 0x00,
}
