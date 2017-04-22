// Code generated by protoc-gen-go.
// source: certificate.proto
// DO NOT EDIT!

package censys_definitions

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CertificateType int32

const (
	CertificateType_CERTIFICATE_TYPE_RESERVED     CertificateType = 0
	CertificateType_CERTIFICATE_TYPE_UNKNOWN      CertificateType = 1
	CertificateType_CERTIFICATE_TYPE_LEAF         CertificateType = 2
	CertificateType_CERTIFICATE_TYPE_INTERMEDIATE CertificateType = 3
	CertificateType_CERTIFICATE_TYPE_ROOT         CertificateType = 4
)

var CertificateType_name = map[int32]string{
	0: "CERTIFICATE_TYPE_RESERVED",
	1: "CERTIFICATE_TYPE_UNKNOWN",
	2: "CERTIFICATE_TYPE_LEAF",
	3: "CERTIFICATE_TYPE_INTERMEDIATE",
	4: "CERTIFICATE_TYPE_ROOT",
}
var CertificateType_value = map[string]int32{
	"CERTIFICATE_TYPE_RESERVED":     0,
	"CERTIFICATE_TYPE_UNKNOWN":      1,
	"CERTIFICATE_TYPE_LEAF":         2,
	"CERTIFICATE_TYPE_INTERMEDIATE": 3,
	"CERTIFICATE_TYPE_ROOT":         4,
}

func (x CertificateType) String() string {
	return proto.EnumName(CertificateType_name, int32(x))
}
func (CertificateType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type CertificateSource int32

const (
	CertificateSource_CERTIFICATE_SOURCE_RESERVED           CertificateSource = 0
	CertificateSource_CERTIFICATE_SOURCE_UNKNOWN            CertificateSource = 1
	CertificateSource_CERTIFICATE_SOURCE_SCAN               CertificateSource = 2
	CertificateSource_CERTIFICATE_SOURCE_CT                 CertificateSource = 3
	CertificateSource_CERTIFICATE_SOURCE_MOZILLA_SALESFORCE CertificateSource = 4
	CertificateSource_CERTIFICATE_SOURCE_RESEARCH           CertificateSource = 5
	CertificateSource_CERTIFICATE_SOURCE_RAPID7             CertificateSource = 6
	CertificateSource_CERTIFICATE_SOURCE_HUBBLE             CertificateSource = 7
	CertificateSource_CERTIFICATE_SOURCE_CT_CHAIN           CertificateSource = 8
)

var CertificateSource_name = map[int32]string{
	0: "CERTIFICATE_SOURCE_RESERVED",
	1: "CERTIFICATE_SOURCE_UNKNOWN",
	2: "CERTIFICATE_SOURCE_SCAN",
	3: "CERTIFICATE_SOURCE_CT",
	4: "CERTIFICATE_SOURCE_MOZILLA_SALESFORCE",
	5: "CERTIFICATE_SOURCE_RESEARCH",
	6: "CERTIFICATE_SOURCE_RAPID7",
	7: "CERTIFICATE_SOURCE_HUBBLE",
	8: "CERTIFICATE_SOURCE_CT_CHAIN",
}
var CertificateSource_value = map[string]int32{
	"CERTIFICATE_SOURCE_RESERVED":           0,
	"CERTIFICATE_SOURCE_UNKNOWN":            1,
	"CERTIFICATE_SOURCE_SCAN":               2,
	"CERTIFICATE_SOURCE_CT":                 3,
	"CERTIFICATE_SOURCE_MOZILLA_SALESFORCE": 4,
	"CERTIFICATE_SOURCE_RESEARCH":           5,
	"CERTIFICATE_SOURCE_RAPID7":             6,
	"CERTIFICATE_SOURCE_HUBBLE":             7,
	"CERTIFICATE_SOURCE_CT_CHAIN":           8,
}

func (x CertificateSource) String() string {
	return proto.EnumName(CertificateSource_name, int32(x))
}
func (CertificateSource) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type CertificateParseStatus int32

const (
	CertificateParseStatus_CERTIFICATE_PARSE_STATUS_RESERVED   CertificateParseStatus = 0
	CertificateParseStatus_CERTIFICATE_PARSE_STATUS_UNKNOWN    CertificateParseStatus = 1
	CertificateParseStatus_CERTIFICATE_PARSE_STATUS_SUCCESS    CertificateParseStatus = 2
	CertificateParseStatus_CERTIFICATE_PARSE_STATUS_FAIL       CertificateParseStatus = 3
	CertificateParseStatus_CERTIFICATE_PARSE_STATUS_NOT_PARSED CertificateParseStatus = 4
)

var CertificateParseStatus_name = map[int32]string{
	0: "CERTIFICATE_PARSE_STATUS_RESERVED",
	1: "CERTIFICATE_PARSE_STATUS_UNKNOWN",
	2: "CERTIFICATE_PARSE_STATUS_SUCCESS",
	3: "CERTIFICATE_PARSE_STATUS_FAIL",
	4: "CERTIFICATE_PARSE_STATUS_NOT_PARSED",
}
var CertificateParseStatus_value = map[string]int32{
	"CERTIFICATE_PARSE_STATUS_RESERVED":   0,
	"CERTIFICATE_PARSE_STATUS_UNKNOWN":    1,
	"CERTIFICATE_PARSE_STATUS_SUCCESS":    2,
	"CERTIFICATE_PARSE_STATUS_FAIL":       3,
	"CERTIFICATE_PARSE_STATUS_NOT_PARSED": 4,
}

func (x CertificateParseStatus) String() string {
	return proto.EnumName(CertificateParseStatus_name, int32(x))
}
func (CertificateParseStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type CertificateRevocationReason int32

const (
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_RESERVED CertificateRevocationReason = 0
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_UNKNOWN  CertificateRevocationReason = 1
	// RFC 5280
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_UNSPECIFIED            CertificateRevocationReason = 2
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_KEY_COMPROMISE         CertificateRevocationReason = 3
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_CA_COMPROMISE          CertificateRevocationReason = 4
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_AFFILIATION_CHANGED    CertificateRevocationReason = 5
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_SUPERSEDED             CertificateRevocationReason = 6
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_CESSATION_OF_OPERATION CertificateRevocationReason = 7
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_CERTIFICATE_HOLD       CertificateRevocationReason = 8
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_REMOVE_FROM_CRL        CertificateRevocationReason = 9
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_PRIVILEGE_WITHDRAWN    CertificateRevocationReason = 10
	CertificateRevocationReason_CERTIFICATE_REVOCATION_REASON_AA_COMPROMISE          CertificateRevocationReason = 11
)

var CertificateRevocationReason_name = map[int32]string{
	0:  "CERTIFICATE_REVOCATION_REASON_RESERVED",
	1:  "CERTIFICATE_REVOCATION_REASON_UNKNOWN",
	2:  "CERTIFICATE_REVOCATION_REASON_UNSPECIFIED",
	3:  "CERTIFICATE_REVOCATION_REASON_KEY_COMPROMISE",
	4:  "CERTIFICATE_REVOCATION_REASON_CA_COMPROMISE",
	5:  "CERTIFICATE_REVOCATION_REASON_AFFILIATION_CHANGED",
	6:  "CERTIFICATE_REVOCATION_REASON_SUPERSEDED",
	7:  "CERTIFICATE_REVOCATION_REASON_CESSATION_OF_OPERATION",
	8:  "CERTIFICATE_REVOCATION_REASON_CERTIFICATE_HOLD",
	9:  "CERTIFICATE_REVOCATION_REASON_REMOVE_FROM_CRL",
	10: "CERTIFICATE_REVOCATION_REASON_PRIVILEGE_WITHDRAWN",
	11: "CERTIFICATE_REVOCATION_REASON_AA_COMPROMISE",
}
var CertificateRevocationReason_value = map[string]int32{
	"CERTIFICATE_REVOCATION_REASON_RESERVED":               0,
	"CERTIFICATE_REVOCATION_REASON_UNKNOWN":                1,
	"CERTIFICATE_REVOCATION_REASON_UNSPECIFIED":            2,
	"CERTIFICATE_REVOCATION_REASON_KEY_COMPROMISE":         3,
	"CERTIFICATE_REVOCATION_REASON_CA_COMPROMISE":          4,
	"CERTIFICATE_REVOCATION_REASON_AFFILIATION_CHANGED":    5,
	"CERTIFICATE_REVOCATION_REASON_SUPERSEDED":             6,
	"CERTIFICATE_REVOCATION_REASON_CESSATION_OF_OPERATION": 7,
	"CERTIFICATE_REVOCATION_REASON_CERTIFICATE_HOLD":       8,
	"CERTIFICATE_REVOCATION_REASON_REMOVE_FROM_CRL":        9,
	"CERTIFICATE_REVOCATION_REASON_PRIVILEGE_WITHDRAWN":    10,
	"CERTIFICATE_REVOCATION_REASON_AA_COMPROMISE":          11,
}

func (x CertificateRevocationReason) String() string {
	return proto.EnumName(CertificateRevocationReason_name, int32(x))
}
func (CertificateRevocationReason) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type RootStoreStatus struct {
	// (trusted path || whitelisted) && other certificate validation (e.g., expiration)
	Valid bool `protobuf:"varint,1,opt,name=valid" json:"valid,omitempty"`
	// was certificate ever valid in the past?
	WasValid bool `protobuf:"varint,2,opt,name=was_valid,json=wasValid" json:"was_valid,omitempty"`
	// does there exist a path from the certificate to the root store
	TrustedPath    bool `protobuf:"varint,3,opt,name=trusted_path,json=trustedPath" json:"trusted_path,omitempty"`
	WasTrustedPath bool `protobuf:"varint,4,opt,name=was_trusted_path,json=wasTrustedPath" json:"was_trusted_path,omitempty"`
	// is the certificate explicitly blacklisted by the browser (e.g., OneCRL)
	Blacklisted bool            `protobuf:"varint,5,opt,name=blacklisted" json:"blacklisted,omitempty"`
	Whitelisted bool            `protobuf:"varint,6,opt,name=whitelisted" json:"whitelisted,omitempty"`
	Type        CertificateType `protobuf:"varint,7,opt,name=type,enum=zsearch.CertificateType" json:"type,omitempty"`
	// contains sha256fp of path to root store
	Path [][]byte `protobuf:"bytes,8,rep,name=path,proto3" json:"path,omitempty"`
}

func (m *RootStoreStatus) Reset()                    { *m = RootStoreStatus{} }
func (m *RootStoreStatus) String() string            { return proto.CompactTextString(m) }
func (*RootStoreStatus) ProtoMessage()               {}
func (*RootStoreStatus) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type CertificateValidation struct {
	Nss       *RootStoreStatus `protobuf:"bytes,1,opt,name=nss" json:"nss,omitempty"`
	Microsoft *RootStoreStatus `protobuf:"bytes,2,opt,name=microsoft" json:"microsoft,omitempty"`
	Apple     *RootStoreStatus `protobuf:"bytes,3,opt,name=apple" json:"apple,omitempty"`
	Java      *RootStoreStatus `protobuf:"bytes,4,opt,name=java" json:"java,omitempty"`
	Android   *RootStoreStatus `protobuf:"bytes,5,opt,name=android" json:"android,omitempty"`
	// also track for Google CT servers so we know what to push
	GoogleCtPrimary *RootStoreStatus `protobuf:"bytes,10,opt,name=google_ct_primary,json=googleCtPrimary" json:"google_ct_primary,omitempty"`
}

func (m *CertificateValidation) Reset()                    { *m = CertificateValidation{} }
func (m *CertificateValidation) String() string            { return proto.CompactTextString(m) }
func (*CertificateValidation) ProtoMessage()               {}
func (*CertificateValidation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *CertificateValidation) GetNss() *RootStoreStatus {
	if m != nil {
		return m.Nss
	}
	return nil
}

func (m *CertificateValidation) GetMicrosoft() *RootStoreStatus {
	if m != nil {
		return m.Microsoft
	}
	return nil
}

func (m *CertificateValidation) GetApple() *RootStoreStatus {
	if m != nil {
		return m.Apple
	}
	return nil
}

func (m *CertificateValidation) GetJava() *RootStoreStatus {
	if m != nil {
		return m.Java
	}
	return nil
}

func (m *CertificateValidation) GetAndroid() *RootStoreStatus {
	if m != nil {
		return m.Android
	}
	return nil
}

func (m *CertificateValidation) GetGoogleCtPrimary() *RootStoreStatus {
	if m != nil {
		return m.GoogleCtPrimary
	}
	return nil
}

type MozillaSalesForceStatus struct {
	CurrentIn                      bool   `protobuf:"varint,1,opt,name=current_in,json=currentIn" json:"current_in,omitempty"`
	WasIn                          bool   `protobuf:"varint,2,opt,name=was_in,json=wasIn" json:"was_in,omitempty"`
	OwnerName                      string `protobuf:"bytes,3,opt,name=owner_name,json=ownerName" json:"owner_name,omitempty"`
	ParentName                     string `protobuf:"bytes,4,opt,name=parent_name,json=parentName" json:"parent_name,omitempty"`
	CertificateName                string `protobuf:"bytes,5,opt,name=certificate_name,json=certificateName" json:"certificate_name,omitempty"`
	CertificatePolicy              string `protobuf:"bytes,6,opt,name=certificate_policy,json=certificatePolicy" json:"certificate_policy,omitempty"`
	CertificationPracticeStatement string `protobuf:"bytes,7,opt,name=certification_practice_statement,json=certificationPracticeStatement" json:"certification_practice_statement,omitempty"`
	CpSameAsParent                 bool   `protobuf:"varint,8,opt,name=cp_same_as_parent,json=cpSameAsParent" json:"cp_same_as_parent,omitempty"`
	AuditSameAsParent              bool   `protobuf:"varint,9,opt,name=audit_same_as_parent,json=auditSameAsParent" json:"audit_same_as_parent,omitempty"`
	StandardAudit                  string `protobuf:"bytes,10,opt,name=standard_audit,json=standardAudit" json:"standard_audit,omitempty"`
	BrAudit                        string `protobuf:"bytes,11,opt,name=br_audit,json=brAudit" json:"br_audit,omitempty"`
	Auditor                        string `protobuf:"bytes,12,opt,name=auditor" json:"auditor,omitempty"`
}

func (m *MozillaSalesForceStatus) Reset()                    { *m = MozillaSalesForceStatus{} }
func (m *MozillaSalesForceStatus) String() string            { return proto.CompactTextString(m) }
func (*MozillaSalesForceStatus) ProtoMessage()               {}
func (*MozillaSalesForceStatus) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type CertificateRevocation struct {
	Revoked bool                        `protobuf:"varint,1,opt,name=revoked" json:"revoked,omitempty"`
	Reason  CertificateRevocationReason `protobuf:"varint,2,opt,name=reason,enum=zsearch.CertificateRevocationReason" json:"reason,omitempty"`
}

func (m *CertificateRevocation) Reset()                    { *m = CertificateRevocation{} }
func (m *CertificateRevocation) String() string            { return proto.CompactTextString(m) }
func (*CertificateRevocation) ProtoMessage()               {}
func (*CertificateRevocation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type CertificateAudit struct {
	Mozilla *MozillaSalesForceStatus `protobuf:"bytes,1,opt,name=mozilla" json:"mozilla,omitempty"`
}

func (m *CertificateAudit) Reset()                    { *m = CertificateAudit{} }
func (m *CertificateAudit) String() string            { return proto.CompactTextString(m) }
func (*CertificateAudit) ProtoMessage()               {}
func (*CertificateAudit) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *CertificateAudit) GetMozilla() *MozillaSalesForceStatus {
	if m != nil {
		return m.Mozilla
	}
	return nil
}

// Database Records
type Certificate struct {
	Sha1Fp       []byte                 `protobuf:"bytes,1,opt,name=sha1fp,proto3" json:"sha1fp,omitempty"`
	Sha256Fp     []byte                 `protobuf:"bytes,2,opt,name=sha256fp,proto3" json:"sha256fp,omitempty"`
	Raw          []byte                 `protobuf:"bytes,3,opt,name=raw,proto3" json:"raw,omitempty"`
	Parsed       string                 `protobuf:"bytes,4,opt,name=parsed" json:"parsed,omitempty"`
	ParseStatus  CertificateParseStatus `protobuf:"varint,44,opt,name=parse_status,json=parseStatus,enum=zsearch.CertificateParseStatus" json:"parse_status,omitempty"`
	ParseVersion uint32                 `protobuf:"varint,39,opt,name=parse_version,json=parseVersion" json:"parse_version,omitempty"`
	ParseError   string                 `protobuf:"bytes,47,opt,name=parse_error,json=parseError" json:"parse_error,omitempty"`
	Parents      [][]byte               `protobuf:"bytes,5,rep,name=parents,proto3" json:"parents,omitempty"`
	// the chain that we received with the certificate and pass
	// to the certificate processing daemon.
	PresentedChain       [][]byte               `protobuf:"bytes,45,rep,name=presented_chain,json=presentedChain,proto3" json:"presented_chain,omitempty"`
	Source               CertificateSource      `protobuf:"varint,28,opt,name=source,enum=zsearch.CertificateSource" json:"source,omitempty"`
	SeenInScan           bool                   `protobuf:"varint,29,opt,name=seen_in_scan,json=seenInScan" json:"seen_in_scan,omitempty"`
	PostProcessed        bool                   `protobuf:"varint,26,opt,name=post_processed,json=postProcessed" json:"post_processed,omitempty"`
	PostProcessTimestamp uint32                 `protobuf:"varint,37,opt,name=post_process_timestamp,json=postProcessTimestamp" json:"post_process_timestamp,omitempty"`
	Validation           *CertificateValidation `protobuf:"bytes,35,opt,name=validation" json:"validation,omitempty"`
	Ct                   *CTStatus              `protobuf:"bytes,30,opt,name=ct" json:"ct,omitempty"`
	Zlint                *ZLint                 `protobuf:"bytes,38,opt,name=zlint" json:"zlint,omitempty"`
	Revocation           *CertificateRevocation `protobuf:"bytes,43,opt,name=revocation" json:"revocation,omitempty"`
	Audit                *CertificateAudit      `protobuf:"bytes,46,opt,name=audit" json:"audit,omitempty"`
	// store wheter record is precert so that CT synchronization
	// daemon can quickly decide what to do without reparsing the
	// raw certificate
	IsPrecert bool `protobuf:"varint,32,opt,name=is_precert,json=isPrecert" json:"is_precert,omitempty"`
	// store time range that certificate is valid so that cert
	// dump job can mark certificates as expired
	NotValidAfter         uint32                   `protobuf:"varint,41,opt,name=not_valid_after,json=notValidAfter" json:"not_valid_after,omitempty"`
	NotValidBefore        uint32                   `protobuf:"varint,42,opt,name=not_valid_before,json=notValidBefore" json:"not_valid_before,omitempty"`
	InNss                 bool                     `protobuf:"varint,6,opt,name=in_nss,json=inNss" json:"in_nss,omitempty"`
	InMicrosoft           bool                     `protobuf:"varint,7,opt,name=in_microsoft,json=inMicrosoft" json:"in_microsoft,omitempty"`
	InApple               bool                     `protobuf:"varint,8,opt,name=in_apple,json=inApple" json:"in_apple,omitempty"`
	ValidationTimestamp   uint32                   `protobuf:"varint,10,opt,name=validation_timestamp,json=validationTimestamp" json:"validation_timestamp,omitempty"`
	ValidNss              bool                     `protobuf:"varint,11,opt,name=valid_nss,json=validNss" json:"valid_nss,omitempty"`
	ValidMicrosoft        bool                     `protobuf:"varint,12,opt,name=valid_microsoft,json=validMicrosoft" json:"valid_microsoft,omitempty"`
	ValidApple            bool                     `protobuf:"varint,13,opt,name=valid_apple,json=validApple" json:"valid_apple,omitempty"`
	WasValidNss           bool                     `protobuf:"varint,14,opt,name=was_valid_nss,json=wasValidNss" json:"was_valid_nss,omitempty"`
	WasValidMicrosoft     bool                     `protobuf:"varint,15,opt,name=was_valid_microsoft,json=wasValidMicrosoft" json:"was_valid_microsoft,omitempty"`
	WasValidApple         bool                     `protobuf:"varint,16,opt,name=was_valid_apple,json=wasValidApple" json:"was_valid_apple,omitempty"`
	WasInNss              bool                     `protobuf:"varint,17,opt,name=was_in_nss,json=wasInNss" json:"was_in_nss,omitempty"`
	WasInMicrosoft        bool                     `protobuf:"varint,18,opt,name=was_in_microsoft,json=wasInMicrosoft" json:"was_in_microsoft,omitempty"`
	WasInApple            bool                     `protobuf:"varint,19,opt,name=was_in_apple,json=wasInApple" json:"was_in_apple,omitempty"`
	CurrentValidNss       bool                     `protobuf:"varint,20,opt,name=current_valid_nss,json=currentValidNss" json:"current_valid_nss,omitempty"`
	CurrentValidMicrosoft bool                     `protobuf:"varint,21,opt,name=current_valid_microsoft,json=currentValidMicrosoft" json:"current_valid_microsoft,omitempty"`
	CurrentValidApple     bool                     `protobuf:"varint,22,opt,name=current_valid_apple,json=currentValidApple" json:"current_valid_apple,omitempty"`
	CurrentInNss          bool                     `protobuf:"varint,23,opt,name=current_in_nss,json=currentInNss" json:"current_in_nss,omitempty"`
	CurrentInMicrosoft    bool                     `protobuf:"varint,24,opt,name=current_in_microsoft,json=currentInMicrosoft" json:"current_in_microsoft,omitempty"`
	CurrentInApple        bool                     `protobuf:"varint,25,opt,name=current_in_apple,json=currentInApple" json:"current_in_apple,omitempty"`
	NssAudit              *MozillaSalesForceStatus `protobuf:"bytes,31,opt,name=nss_audit,json=nssAudit" json:"nss_audit,omitempty"`
	ShouldPostProcess     bool                     `protobuf:"varint,27,opt,name=should_post_process,json=shouldPostProcess" json:"should_post_process,omitempty"`
	DoNotPostProcess      bool                     `protobuf:"varint,36,opt,name=do_not_post_process,json=doNotPostProcess" json:"do_not_post_process,omitempty"`
}

func (m *Certificate) Reset()                    { *m = Certificate{} }
func (m *Certificate) String() string            { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()               {}
func (*Certificate) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *Certificate) GetValidation() *CertificateValidation {
	if m != nil {
		return m.Validation
	}
	return nil
}

func (m *Certificate) GetCt() *CTStatus {
	if m != nil {
		return m.Ct
	}
	return nil
}

func (m *Certificate) GetZlint() *ZLint {
	if m != nil {
		return m.Zlint
	}
	return nil
}

func (m *Certificate) GetRevocation() *CertificateRevocation {
	if m != nil {
		return m.Revocation
	}
	return nil
}

func (m *Certificate) GetAudit() *CertificateAudit {
	if m != nil {
		return m.Audit
	}
	return nil
}

func (m *Certificate) GetNssAudit() *MozillaSalesForceStatus {
	if m != nil {
		return m.NssAudit
	}
	return nil
}

func init() {
	proto.RegisterType((*RootStoreStatus)(nil), "zsearch.RootStoreStatus")
	proto.RegisterType((*CertificateValidation)(nil), "zsearch.CertificateValidation")
	proto.RegisterType((*MozillaSalesForceStatus)(nil), "zsearch.MozillaSalesForceStatus")
	proto.RegisterType((*CertificateRevocation)(nil), "zsearch.CertificateRevocation")
	proto.RegisterType((*CertificateAudit)(nil), "zsearch.CertificateAudit")
	proto.RegisterType((*Certificate)(nil), "zsearch.Certificate")
	proto.RegisterEnum("zsearch.CertificateType", CertificateType_name, CertificateType_value)
	proto.RegisterEnum("zsearch.CertificateSource", CertificateSource_name, CertificateSource_value)
	proto.RegisterEnum("zsearch.CertificateParseStatus", CertificateParseStatus_name, CertificateParseStatus_value)
	proto.RegisterEnum("zsearch.CertificateRevocationReason", CertificateRevocationReason_name, CertificateRevocationReason_value)
}

func init() { proto.RegisterFile("certificate.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1790 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x58, 0xdd, 0x52, 0xdb, 0x48,
	0x16, 0x5e, 0x7e, 0x0c, 0xf6, 0xb1, 0xb1, 0x45, 0x07, 0x12, 0x85, 0x4c, 0x12, 0x86, 0x90, 0x3f,
	0x92, 0x90, 0x0d, 0x9b, 0x99, 0xdd, 0x9a, 0xda, 0xda, 0x2a, 0x61, 0xe4, 0x41, 0x35, 0xc6, 0x76,
	0xb5, 0x0c, 0x53, 0x33, 0x37, 0x2a, 0x21, 0xc4, 0xa0, 0x8d, 0x2d, 0xb9, 0x24, 0x41, 0x2a, 0x79,
	0x91, 0x7d, 0x80, 0xbd, 0xd8, 0x47, 0xd8, 0x07, 0xd9, 0x8b, 0x7d, 0x88, 0xbd, 0xdd, 0x07, 0xd8,
	0xd3, 0xa7, 0x65, 0xb5, 0x0c, 0xc6, 0xec, 0x15, 0xea, 0xef, 0x7c, 0xa7, 0xfd, 0xf5, 0xe9, 0xf3,
	0xd3, 0x05, 0xac, 0x7a, 0x7e, 0x9c, 0x06, 0xe7, 0x81, 0xe7, 0xa6, 0xfe, 0xee, 0x28, 0x8e, 0xd2,
	0x88, 0x2d, 0x7f, 0x4d, 0x7c, 0x37, 0xf6, 0x2e, 0x36, 0x6a, 0x5e, 0x34, 0x1c, 0x46, 0xa1, 0x84,
	0x37, 0xaa, 0x5f, 0x07, 0x41, 0x98, 0x66, 0x8b, 0xb2, 0x97, 0x7d, 0x6d, 0xfd, 0x6d, 0x1e, 0x1a,
	0x3c, 0x8a, 0x52, 0x3b, 0x8d, 0x62, 0xdf, 0x4e, 0xdd, 0xf4, 0x32, 0x61, 0x6b, 0x50, 0xba, 0x72,
	0x07, 0xc1, 0x99, 0x3e, 0xb7, 0x39, 0xf7, 0xaa, 0xcc, 0xe5, 0x82, 0x3d, 0x82, 0xca, 0x67, 0x37,
	0x71, 0xa4, 0x65, 0x9e, 0x2c, 0x65, 0x04, 0x4e, 0xc8, 0xf8, 0x2d, 0xd4, 0xd2, 0xf8, 0x32, 0x49,
	0xfd, 0x33, 0x67, 0xe4, 0xa6, 0x17, 0xfa, 0x02, 0xd9, 0xab, 0x19, 0xd6, 0x43, 0x88, 0xbd, 0x02,
	0x4d, 0xf8, 0x4f, 0xd0, 0x16, 0x89, 0x56, 0x47, 0xbc, 0x5f, 0x60, 0x6e, 0x42, 0xf5, 0x74, 0xe0,
	0x7a, 0x9f, 0x06, 0x81, 0x80, 0xf4, 0x92, 0xdc, 0xab, 0x00, 0x09, 0xc6, 0xe7, 0x8b, 0x20, 0xf5,
	0x33, 0xc6, 0x92, 0x64, 0x14, 0x20, 0xf6, 0x16, 0x16, 0xd3, 0x2f, 0x23, 0x5f, 0x5f, 0x46, 0x53,
	0x7d, 0x4f, 0xdf, 0xcd, 0x82, 0xb2, 0xdb, 0x54, 0xf1, 0xea, 0xa3, 0x9d, 0x13, 0x8b, 0x31, 0x58,
	0x24, 0x3d, 0xe5, 0xcd, 0x85, 0x57, 0x35, 0x4e, 0xdf, 0x5b, 0xff, 0x9a, 0x87, 0xf5, 0x02, 0x9b,
	0xce, 0xe9, 0xa6, 0x41, 0x14, 0xb2, 0x1d, 0x58, 0x08, 0x93, 0x84, 0xa2, 0x53, 0x2d, 0x6c, 0x7d,
	0x2d, 0x8c, 0x5c, 0x90, 0xd8, 0xf7, 0x50, 0x19, 0x06, 0x5e, 0x1c, 0x25, 0xd1, 0x79, 0x4a, 0x51,
	0x9b, 0xe5, 0xa1, 0xa8, 0x6c, 0x17, 0x4a, 0xee, 0x68, 0x34, 0xf0, 0x29, 0x92, 0xb3, 0x7c, 0x24,
	0x4d, 0x9c, 0xf7, 0xaf, 0xee, 0x95, 0x4b, 0x11, 0x9d, 0x45, 0x27, 0x16, 0xdb, 0x83, 0x65, 0x37,
	0x3c, 0x8b, 0xa3, 0x40, 0x46, 0x77, 0x96, 0xc3, 0x98, 0xc8, 0x0e, 0x60, 0xf5, 0xb7, 0x28, 0xfa,
	0x6d, 0xe0, 0x3b, 0x5e, 0xea, 0x8c, 0xe2, 0x60, 0xe8, 0xc6, 0x5f, 0x74, 0xb8, 0xc3, 0xbb, 0x21,
	0x5d, 0x9a, 0x69, 0x4f, 0x3a, 0x6c, 0xfd, 0x67, 0x01, 0x1e, 0x1c, 0x45, 0x5f, 0x83, 0xc1, 0xc0,
	0xb5, 0xdd, 0x81, 0x9f, 0xb4, 0xa2, 0xd8, 0x1b, 0xe7, 0xdd, 0x63, 0x00, 0xef, 0x32, 0x8e, 0xfd,
	0x30, 0x75, 0x82, 0x30, 0x4b, 0xbe, 0x4a, 0x86, 0x58, 0x21, 0x5b, 0x87, 0x25, 0x91, 0x40, 0x68,
	0x92, 0xd9, 0x57, 0xc2, 0x15, 0xc2, 0xe8, 0x15, 0x7d, 0x0e, 0xfd, 0xd8, 0x09, 0xdd, 0xa1, 0x0c,
	0x57, 0x85, 0x57, 0x08, 0xe9, 0x20, 0xc0, 0x9e, 0x42, 0x75, 0xe4, 0xd2, 0x9e, 0x64, 0x5f, 0x24,
	0x3b, 0x48, 0x88, 0x08, 0xaf, 0x41, 0x2b, 0x14, 0x91, 0x64, 0x95, 0x88, 0xd5, 0x28, 0xe0, 0x44,
	0x7d, 0x07, 0xac, 0x48, 0x1d, 0x45, 0x83, 0xc0, 0xfb, 0x42, 0xd9, 0x57, 0xe1, 0xc5, 0x4a, 0xec,
	0x91, 0x81, 0x1d, 0xc2, 0xa6, 0x02, 0x31, 0x71, 0x30, 0x6a, 0xae, 0x97, 0x06, 0x9e, 0xef, 0x24,
	0x78, 0x60, 0x7f, 0x88, 0x0a, 0x28, 0x3f, 0x2b, 0xfc, 0xc9, 0x04, 0xaf, 0x97, 0xd1, 0xec, 0x31,
	0x0b, 0x35, 0xae, 0x7a, 0x23, 0x27, 0x41, 0x0d, 0x0e, 0x46, 0x40, 0x8a, 0xc7, 0x64, 0xa5, 0xe2,
	0xf1, 0x46, 0x36, 0xe2, 0x46, 0xd2, 0x23, 0x94, 0xbd, 0x87, 0x35, 0xf7, 0xf2, 0x2c, 0x48, 0xaf,
	0xb3, 0x2b, 0xc4, 0x5e, 0x25, 0xdb, 0x84, 0xc3, 0x73, 0xa8, 0xa3, 0x9c, 0xf0, 0xcc, 0x8d, 0xcf,
	0x1c, 0xb2, 0xd2, 0xa5, 0x56, 0xf8, 0xca, 0x18, 0x35, 0x04, 0xc8, 0x1e, 0x42, 0xf9, 0x34, 0xce,
	0x08, 0x55, 0x22, 0x2c, 0x9f, 0xc6, 0xd2, 0xa4, 0x63, 0x36, 0x89, 0x8f, 0x28, 0xd6, 0x6b, 0xd2,
	0x92, 0x2d, 0xb7, 0xa2, 0x89, 0x12, 0xe2, 0xfe, 0x55, 0x24, 0x4f, 0x28, 0x5c, 0x62, 0x5c, 0x7d,
	0xf2, 0xc7, 0x4d, 0x66, 0xbc, 0x64, 0x7f, 0x86, 0xa5, 0xd8, 0x77, 0x93, 0x48, 0xde, 0x72, 0x7d,
	0x6f, 0x7b, 0x5a, 0xe9, 0xaa, 0x9d, 0x38, 0x71, 0x79, 0xe6, 0xb3, 0xd5, 0x01, 0xad, 0x40, 0x93,
	0xf2, 0x7e, 0x80, 0xe5, 0xa1, 0xcc, 0xb8, 0xac, 0x64, 0x37, 0xf3, 0x2d, 0x6f, 0xc9, 0x44, 0x3e,
	0x76, 0xd8, 0xfa, 0x7b, 0x1d, 0xaa, 0x85, 0x0d, 0xd9, 0x7d, 0x58, 0x4a, 0x2e, 0xdc, 0x0f, 0xe7,
	0x23, 0xda, 0xaa, 0xc6, 0xb3, 0x15, 0xdb, 0x80, 0x32, 0x7e, 0xed, 0x7d, 0xf7, 0x3d, 0x5a, 0xe6,
	0xc9, 0x92, 0xaf, 0x99, 0x06, 0x0b, 0xb1, 0xfb, 0x99, 0x32, 0xb3, 0xc6, 0xc5, 0xa7, 0xd8, 0x05,
	0x6f, 0x25, 0xc1, 0xc3, 0xcb, 0x74, 0xcc, 0x56, 0x6c, 0x1f, 0x6a, 0xf4, 0x45, 0xf9, 0x71, 0x99,
	0xe8, 0x6f, 0x29, 0x02, 0x4f, 0xa7, 0x45, 0xa0, 0x27, 0x78, 0x99, 0xda, 0xea, 0x48, 0x2d, 0xd8,
	0x33, 0x58, 0x91, 0x7b, 0x5c, 0xf9, 0x71, 0x82, 0x01, 0xd2, 0x5f, 0xe2, 0x26, 0x2b, 0x5c, 0x6e,
	0x7c, 0x22, 0xb1, 0xac, 0x28, 0x90, 0xe4, 0xc7, 0x31, 0xde, 0xda, 0xfb, 0xbc, 0x28, 0x12, 0xdf,
	0x14, 0x88, 0xb8, 0x1f, 0x99, 0x37, 0x09, 0xd6, 0x82, 0xe8, 0x89, 0xe3, 0x25, 0x7b, 0x09, 0x8d,
	0x51, 0xec, 0x27, 0xf8, 0x8d, 0x4d, 0xdc, 0xbb, 0x70, 0xb1, 0x1c, 0xdf, 0x11, 0xa3, 0x9e, 0xc3,
	0x4d, 0x81, 0x62, 0x8f, 0x59, 0x4a, 0xa2, 0x4b, 0x8c, 0xa9, 0xfe, 0x0d, 0x1d, 0x63, 0x63, 0xda,
	0x31, 0x6c, 0x62, 0xf0, 0x8c, 0x89, 0x7d, 0xbd, 0x96, 0xf8, 0x7e, 0x88, 0x35, 0xee, 0x24, 0x9e,
	0x1b, 0xea, 0x8f, 0x29, 0x37, 0x40, 0x60, 0x56, 0x68, 0x23, 0x22, 0xb2, 0x75, 0x14, 0x25, 0xa2,
	0x01, 0x45, 0x9e, 0x9f, 0x88, 0x10, 0x6e, 0x10, 0x67, 0x45, 0xa0, 0xbd, 0x31, 0xc8, 0x3e, 0xc2,
	0xfd, 0x22, 0xcd, 0x49, 0x83, 0xa1, 0x8f, 0x41, 0x1d, 0x8e, 0xf4, 0xe7, 0x14, 0x8e, 0xb5, 0x02,
	0xbd, 0x3f, 0xb6, 0xb1, 0xbf, 0x00, 0x5c, 0xe5, 0x6d, 0x5e, 0x7f, 0x46, 0xc9, 0xf2, 0x64, 0x9a,
	0x6c, 0x35, 0x0c, 0x78, 0xc1, 0x03, 0xa7, 0xe0, 0xbc, 0x97, 0xea, 0x4f, 0xc8, 0x6f, 0x55, 0xf9,
	0xf5, 0xb3, 0x7b, 0x42, 0x23, 0xdb, 0x86, 0x12, 0x0d, 0x62, 0xfd, 0x05, 0xb1, 0xea, 0x39, 0xeb,
	0xd7, 0x36, 0xa2, 0x5c, 0x1a, 0x85, 0x90, 0x38, 0x4f, 0x71, 0xfd, 0xcd, 0xed, 0x42, 0x0a, 0x85,
	0x50, 0xf0, 0xc0, 0x26, 0x50, 0x92, 0x95, 0xba, 0x4b, 0xae, 0x0f, 0xa7, 0xb9, 0x52, 0x71, 0x70,
	0xc9, 0x13, 0x4d, 0x34, 0xc0, 0x56, 0x11, 0xfb, 0xa2, 0x11, 0xe9, 0x9b, 0xb2, 0xf5, 0x06, 0x49,
	0x4f, 0x02, 0xec, 0x05, 0x34, 0xc2, 0x28, 0x95, 0xb3, 0xdf, 0x71, 0xcf, 0x53, 0x3f, 0xd6, 0x5f,
	0x53, 0x1c, 0x57, 0x10, 0xa6, 0x60, 0x18, 0x02, 0x14, 0x33, 0x5e, 0xf1, 0x4e, 0xfd, 0x73, 0x1c,
	0x04, 0xfa, 0x0e, 0x11, 0xeb, 0x63, 0xe2, 0x3e, 0xa1, 0xa2, 0x99, 0xe3, 0x25, 0x8b, 0x31, 0x2a,
	0x87, 0x77, 0x29, 0x08, 0x3b, 0x38, 0x2e, 0x9f, 0x43, 0x0d, 0x61, 0x35, 0x31, 0x45, 0x7b, 0x2c,
	0xef, 0xcf, 0xeb, 0x73, 0xbc, 0x1a, 0x84, 0x47, 0xf9, 0x74, 0x7c, 0x0c, 0x65, 0xa4, 0xc9, 0x01,
	0x59, 0xce, 0x29, 0xcb, 0x41, 0x68, 0xd0, 0x30, 0xfc, 0x00, 0x6b, 0xea, 0x56, 0x0a, 0x77, 0x0f,
	0x24, 0xe5, 0x9e, 0xb2, 0xa9, 0xab, 0x7f, 0x0a, 0x15, 0xa9, 0x5a, 0x48, 0xaa, 0xe6, 0x5b, 0x96,
	0x09, 0x14, 0xca, 0xde, 0x40, 0x43, 0x12, 0x94, 0xb8, 0x5a, 0x4e, 0xab, 0x93, 0x49, 0xe9, 0x7b,
	0x06, 0xd5, 0x2c, 0x56, 0x24, 0x71, 0x25, 0x27, 0xca, 0x6c, 0x91, 0x2a, 0xb7, 0x60, 0x25, 0x7f,
	0x50, 0xd1, 0xcf, 0xd6, 0xb3, 0x67, 0x4c, 0xf6, 0xa8, 0x12, 0xbf, 0xba, 0x07, 0xf7, 0x14, 0x47,
	0xfd, 0x72, 0x23, 0xdf, 0x70, 0x75, 0xcc, 0x56, 0x3f, 0xbe, 0x03, 0x0d, 0xe5, 0x23, 0x05, 0x68,
	0x39, 0x7f, 0x65, 0xcc, 0x97, 0x1a, 0x36, 0x01, 0xe4, 0x4c, 0x25, 0x01, 0xab, 0xea, 0xdc, 0x34,
	0x5b, 0x85, 0x82, 0xb7, 0xf2, 0xd9, 0x36, 0x71, 0x2b, 0x4c, 0x1d, 0x9c, 0x78, 0xea, 0xb7, 0xb7,
	0xa1, 0x96, 0xb1, 0xe5, 0x0f, 0xdf, 0x53, 0x27, 0x27, 0xa6, 0xfc, 0xd5, 0x1d, 0x1c, 0x67, 0xd9,
	0xa0, 0x57, 0xa7, 0x5f, 0xa3, 0xd3, 0x37, 0x32, 0x43, 0x1e, 0x81, 0x1f, 0xe0, 0xc1, 0x24, 0x57,
	0xc9, 0x58, 0xcf, 0x37, 0x5f, 0x2f, 0x7a, 0x29, 0x35, 0x18, 0xbd, 0x49, 0x5f, 0x29, 0xea, 0xbe,
	0x8a, 0x5e, 0xd1, 0x4f, 0x6a, 0xdb, 0x86, 0xba, 0x7a, 0x84, 0x90, 0xb0, 0x07, 0x24, 0xac, 0x96,
	0x3f, 0x44, 0x84, 0xaa, 0x8f, 0xb0, 0x56, 0x60, 0x29, 0x49, 0x7a, 0xbe, 0x35, 0xcb, 0xf9, 0x4a,
	0x0f, 0xc6, 0xb2, 0xe0, 0x25, 0xc5, 0x3c, 0x54, 0xb1, 0xcc, 0x3d, 0xa4, 0x12, 0x03, 0x2a, 0xf8,
	0xf3, 0xd9, 0xc8, 0x7d, 0xfa, 0xff, 0x4d, 0x2e, 0x79, 0x79, 0xe8, 0x26, 0x47, 0x1f, 0x06, 0x20,
	0xb9, 0x88, 0x2e, 0x07, 0xf8, 0xdc, 0x2e, 0x74, 0x43, 0xfd, 0x91, 0x0a, 0x80, 0x34, 0xf7, 0x54,
	0x37, 0xc4, 0xe2, 0xb9, 0x77, 0x16, 0x39, 0xa2, 0x8c, 0x27, 0x7c, 0xb6, 0x73, 0x1f, 0xed, 0x2c,
	0xea, 0x44, 0x69, 0xc1, 0x65, 0xe7, 0x1f, 0x73, 0xd0, 0xb8, 0xf6, 0xb0, 0xc6, 0x12, 0x7d, 0xd8,
	0x34, 0x79, 0xdf, 0x6a, 0x59, 0x4d, 0xa3, 0x6f, 0x3a, 0xfd, 0x5f, 0x7a, 0xa6, 0xc3, 0x4d, 0xdb,
	0xe4, 0x27, 0xe6, 0x81, 0xf6, 0x3b, 0xf6, 0x0d, 0xe8, 0x37, 0xcc, 0xc7, 0x9d, 0x9f, 0x3a, 0xdd,
	0x9f, 0x3b, 0xda, 0x1c, 0x3e, 0x36, 0xd6, 0x6f, 0x58, 0xdb, 0xa6, 0xd1, 0xd2, 0xe6, 0xb1, 0xc7,
	0x3e, 0xbe, 0x61, 0xb2, 0x3a, 0x7d, 0x93, 0x1f, 0x99, 0x07, 0x16, 0x22, 0xda, 0xc2, 0x54, 0x6f,
	0xde, 0xed, 0xf6, 0xb5, 0xc5, 0x9d, 0x7f, 0xce, 0xc3, 0xea, 0x8d, 0xf1, 0x83, 0xc5, 0xff, 0xa8,
	0xe8, 0x60, 0x77, 0x8f, 0x79, 0x73, 0x42, 0xed, 0x13, 0xd8, 0x98, 0x42, 0x50, 0x7a, 0x1f, 0xc1,
	0x83, 0x29, 0x76, 0xbb, 0x69, 0x74, 0x50, 0xf1, 0x35, 0x39, 0x99, 0xb1, 0xd9, 0x47, 0xa5, 0xaf,
	0xe1, 0xf9, 0x14, 0xd3, 0x51, 0xf7, 0x57, 0xab, 0xdd, 0x36, 0x1c, 0xdb, 0x68, 0x9b, 0x76, 0xab,
	0x8b, 0x90, 0xb6, 0x38, 0x43, 0xa3, 0xc1, 0x9b, 0x87, 0x5a, 0xe9, 0x7a, 0xc0, 0xc7, 0x04, 0xa3,
	0x67, 0x1d, 0xfc, 0x51, 0x5b, 0xba, 0xc5, 0x7c, 0x78, 0xbc, 0xbf, 0xdf, 0x36, 0xb5, 0xe5, 0x5b,
	0xb6, 0x6f, 0xf6, 0x9d, 0xe6, 0xa1, 0x61, 0x75, 0xb4, 0xf2, 0xce, 0xbf, 0xe7, 0xe0, 0xfe, 0xf4,
	0xf7, 0x07, 0x36, 0xed, 0x6f, 0x8b, 0xbe, 0x3d, 0x83, 0xdb, 0xb8, 0x43, 0xdf, 0xe8, 0x1f, 0xdb,
	0xc5, 0x20, 0x6e, 0xc3, 0xe6, 0xad, 0x34, 0x15, 0xca, 0x59, 0x2c, 0xfb, 0xb8, 0xd9, 0x34, 0x6d,
	0xfb, 0x66, 0x16, 0x4c, 0xb0, 0x5a, 0x86, 0xd5, 0xc6, 0xd8, 0xbe, 0x84, 0x67, 0xb7, 0x52, 0x3a,
	0xdd, 0xbe, 0x04, 0x0e, 0x30, 0x27, 0xfe, 0xbb, 0x88, 0x67, 0xbf, 0xfd, 0x6d, 0x89, 0xdd, 0xea,
	0x45, 0x71, 0x23, 0x6e, 0x9e, 0x74, 0xf1, 0xc3, 0xea, 0x76, 0xf0, 0xd3, 0xb0, 0xe9, 0x4f, 0x7e,
	0xc6, 0x6b, 0x17, 0x7a, 0x93, 0xab, 0x0e, 0xfa, 0x0e, 0x5e, 0xdf, 0x45, 0xb5, 0x7b, 0x66, 0x13,
	0xed, 0xb8, 0xf3, 0x3c, 0xfb, 0x3d, 0xbc, 0x9d, 0x4d, 0xff, 0xc9, 0xfc, 0xc5, 0x69, 0x76, 0x8f,
	0x7a, 0xbc, 0x7b, 0x64, 0xd9, 0xa2, 0x0c, 0xde, 0xc3, 0x9b, 0xd9, 0x1e, 0x4d, 0xa3, 0xe8, 0xb0,
	0xc8, 0xbe, 0x83, 0x0f, 0xb3, 0x1d, 0x8c, 0x56, 0xcb, 0x6a, 0x5b, 0x12, 0xc2, 0xbc, 0xe8, 0xfc,
	0x88, 0xca, 0x4a, 0xd8, 0xd5, 0x5e, 0xcd, 0x76, 0xb3, 0x8f, 0x7b, 0xa6, 0x88, 0x35, 0xb2, 0x97,
	0xd8, 0x9f, 0xe0, 0xe3, 0x1d, 0xaa, 0xf0, 0x86, 0x25, 0xd0, 0x6d, 0x39, 0x5d, 0xf4, 0xa4, 0x05,
	0xa6, 0xe8, 0x1e, 0xec, 0xde, 0xe5, 0xa9, 0xac, 0x87, 0xdd, 0xf6, 0x81, 0x56, 0xc6, 0x66, 0xf6,
	0xee, 0xae, 0xbb, 0x3b, 0xea, 0x9e, 0x98, 0x4e, 0x0b, 0x83, 0xe0, 0x34, 0x79, 0x5b, 0xab, 0xdc,
	0x1d, 0x85, 0x1e, 0xb7, 0x4e, 0xac, 0xb6, 0xf9, 0xa3, 0xe9, 0xfc, 0x6c, 0xf5, 0x0f, 0x0f, 0xb8,
	0x81, 0xd7, 0x09, 0x77, 0x47, 0xdb, 0x98, 0x88, 0x76, 0xf5, 0x74, 0x89, 0xfe, 0x01, 0xf3, 0x87,
	0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xda, 0xa8, 0xad, 0x18, 0xc3, 0x11, 0x00, 0x00,
}
