// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package censys_definitions

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DeltaType int32

const (
	DeltaType_DT_RESERVED  DeltaType = 0
	DeltaType_DT_UPDATE    DeltaType = 1
	DeltaType_DT_DELETE    DeltaType = 2
	DeltaType_DT_NO_CHANGE DeltaType = 3
)

var DeltaType_name = map[int32]string{
	0: "DT_RESERVED",
	1: "DT_UPDATE",
	2: "DT_DELETE",
	3: "DT_NO_CHANGE",
}
var DeltaType_value = map[string]int32{
	"DT_RESERVED":  0,
	"DT_UPDATE":    1,
	"DT_DELETE":    2,
	"DT_NO_CHANGE": 3,
}

func (x DeltaType) String() string {
	return proto.EnumName(DeltaType_name, int32(x))
}
func (DeltaType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type RegionalRegistrar int32

const (
	RegionalRegistrar_RIR_RESERVED RegionalRegistrar = 0
	RegionalRegistrar_RIR_ARIN     RegionalRegistrar = 1
	RegionalRegistrar_RIR_RIPE     RegionalRegistrar = 2
	RegionalRegistrar_RIR_APNIC    RegionalRegistrar = 3
	RegionalRegistrar_RIR_AFRINIC  RegionalRegistrar = 4
	RegionalRegistrar_RIR_LACNIC   RegionalRegistrar = 5
	RegionalRegistrar_RIR_UNKNOWN  RegionalRegistrar = 6
)

var RegionalRegistrar_name = map[int32]string{
	0: "RIR_RESERVED",
	1: "RIR_ARIN",
	2: "RIR_RIPE",
	3: "RIR_APNIC",
	4: "RIR_AFRINIC",
	5: "RIR_LACNIC",
	6: "RIR_UNKNOWN",
}
var RegionalRegistrar_value = map[string]int32{
	"RIR_RESERVED": 0,
	"RIR_ARIN":     1,
	"RIR_RIPE":     2,
	"RIR_APNIC":    3,
	"RIR_AFRINIC":  4,
	"RIR_LACNIC":   5,
	"RIR_UNKNOWN":  6,
}

func (x RegionalRegistrar) String() string {
	return proto.EnumName(RegionalRegistrar_name, int32(x))
}
func (RegionalRegistrar) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type Metadatum struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Metadatum) Reset()                    { *m = Metadatum{} }
func (m *Metadatum) String() string            { return proto.CompactTextString(m) }
func (*Metadatum) ProtoMessage()               {}
func (*Metadatum) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Metadatum) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Metadatum) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type UserdataAtom struct {
	PrivateNotes    string       `protobuf:"bytes,1,opt,name=private_notes,json=privateNotes" json:"private_notes,omitempty"`
	PublicNotes     string       `protobuf:"bytes,2,opt,name=public_notes,json=publicNotes" json:"public_notes,omitempty"`
	PrivateMetadata []*Metadatum `protobuf:"bytes,3,rep,name=private_metadata,json=privateMetadata" json:"private_metadata,omitempty"`
	PublicMetadata  []*Metadatum `protobuf:"bytes,4,rep,name=public_metadata,json=publicMetadata" json:"public_metadata,omitempty"`
	PrivateTags     []string     `protobuf:"bytes,5,rep,name=private_tags,json=privateTags" json:"private_tags,omitempty"`
	PublicTags      []string     `protobuf:"bytes,6,rep,name=public_tags,json=publicTags" json:"public_tags,omitempty"`
}

func (m *UserdataAtom) Reset()                    { *m = UserdataAtom{} }
func (m *UserdataAtom) String() string            { return proto.CompactTextString(m) }
func (*UserdataAtom) ProtoMessage()               {}
func (*UserdataAtom) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *UserdataAtom) GetPrivateNotes() string {
	if m != nil {
		return m.PrivateNotes
	}
	return ""
}

func (m *UserdataAtom) GetPublicNotes() string {
	if m != nil {
		return m.PublicNotes
	}
	return ""
}

func (m *UserdataAtom) GetPrivateMetadata() []*Metadatum {
	if m != nil {
		return m.PrivateMetadata
	}
	return nil
}

func (m *UserdataAtom) GetPublicMetadata() []*Metadatum {
	if m != nil {
		return m.PublicMetadata
	}
	return nil
}

func (m *UserdataAtom) GetPrivateTags() []string {
	if m != nil {
		return m.PrivateTags
	}
	return nil
}

func (m *UserdataAtom) GetPublicTags() []string {
	if m != nil {
		return m.PublicTags
	}
	return nil
}

type ASAtom struct {
	Asn          uint32            `protobuf:"varint,1,opt,name=asn" json:"asn,omitempty"`
	Description  string            `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Path         []uint32          `protobuf:"varint,3,rep,packed,name=path" json:"path,omitempty"`
	Rir          RegionalRegistrar `protobuf:"varint,4,opt,name=rir,enum=zsearch.RegionalRegistrar" json:"rir,omitempty"`
	BgpPrefix    string            `protobuf:"bytes,5,opt,name=bgp_prefix,json=bgpPrefix" json:"bgp_prefix,omitempty"`
	Name         string            `protobuf:"bytes,6,opt,name=name" json:"name,omitempty"`
	CountryCode  string            `protobuf:"bytes,7,opt,name=country_code,json=countryCode" json:"country_code,omitempty"`
	Organization string            `protobuf:"bytes,8,opt,name=organization" json:"organization,omitempty"`
}

func (m *ASAtom) Reset()                    { *m = ASAtom{} }
func (m *ASAtom) String() string            { return proto.CompactTextString(m) }
func (*ASAtom) ProtoMessage()               {}
func (*ASAtom) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ASAtom) GetAsn() uint32 {
	if m != nil {
		return m.Asn
	}
	return 0
}

func (m *ASAtom) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ASAtom) GetPath() []uint32 {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *ASAtom) GetRir() RegionalRegistrar {
	if m != nil {
		return m.Rir
	}
	return RegionalRegistrar_RIR_RESERVED
}

func (m *ASAtom) GetBgpPrefix() string {
	if m != nil {
		return m.BgpPrefix
	}
	return ""
}

func (m *ASAtom) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ASAtom) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *ASAtom) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

func init() {
	proto.RegisterType((*Metadatum)(nil), "zsearch.Metadatum")
	proto.RegisterType((*UserdataAtom)(nil), "zsearch.UserdataAtom")
	proto.RegisterType((*ASAtom)(nil), "zsearch.ASAtom")
	proto.RegisterEnum("zsearch.DeltaType", DeltaType_name, DeltaType_value)
	proto.RegisterEnum("zsearch.RegionalRegistrar", RegionalRegistrar_name, RegionalRegistrar_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x5f, 0x8f, 0xd2, 0x4c,
	0x14, 0xc6, 0xdf, 0x52, 0x60, 0xb7, 0x87, 0xb2, 0x3b, 0xef, 0xc4, 0x98, 0xc6, 0xc4, 0x88, 0x78,
	0x43, 0x36, 0x86, 0x8b, 0xdd, 0x4b, 0xe3, 0x45, 0xa5, 0x55, 0x89, 0xd8, 0x25, 0xb3, 0x45, 0x2f,
	0x9b, 0xa1, 0x8c, 0xdd, 0x46, 0xda, 0x69, 0xa6, 0xc3, 0x46, 0xf6, 0xd2, 0x0f, 0xe1, 0xa7, 0xf5,
	0xc2, 0xcc, 0x1f, 0x08, 0x89, 0xf1, 0x8a, 0x79, 0x7e, 0xf3, 0x9c, 0x73, 0x9e, 0x33, 0xa4, 0xe0,
	0xe7, 0xbc, 0xaa, 0x78, 0x3d, 0x6d, 0x04, 0x97, 0x1c, 0x9f, 0x3d, 0xb6, 0x8c, 0x8a, 0xfc, 0x7e,
	0x7c, 0x03, 0xde, 0x67, 0x26, 0xe9, 0x86, 0xca, 0x5d, 0x85, 0x11, 0xb8, 0xdf, 0xd9, 0x3e, 0x70,
	0x46, 0xce, 0xc4, 0x23, 0xea, 0x88, 0x9f, 0x40, 0xef, 0x81, 0x6e, 0x77, 0x2c, 0xe8, 0x68, 0x66,
	0xc4, 0xf8, 0x57, 0x07, 0xfc, 0x55, 0xcb, 0xc4, 0x86, 0x4a, 0x1a, 0x4a, 0x5e, 0xe1, 0x57, 0x30,
	0x6c, 0x44, 0xf9, 0x40, 0x25, 0xcb, 0x6a, 0x2e, 0x59, 0x6b, 0x5b, 0xf8, 0x16, 0x26, 0x8a, 0xe1,
	0x97, 0xe0, 0x37, 0xbb, 0xf5, 0xb6, 0xcc, 0xad, 0xc7, 0xb4, 0x1c, 0x18, 0x66, 0x2c, 0x6f, 0x01,
	0x1d, 0xfa, 0x54, 0x26, 0x15, 0x0d, 0xdc, 0x91, 0x3b, 0x19, 0x5c, 0xe3, 0xa9, 0x4d, 0x3c, 0x3d,
	0xc6, 0x25, 0x97, 0xd6, 0x6b, 0x09, 0xc5, 0x6f, 0xe0, 0xd2, 0x4e, 0x38, 0x56, 0x77, 0xff, 0x59,
	0x7d, 0x61, 0xac, 0xc7, 0x62, 0x15, 0xcf, 0xce, 0x96, 0xb4, 0x68, 0x83, 0xde, 0xc8, 0xd5, 0xf1,
	0x0c, 0x4b, 0x69, 0xd1, 0xe2, 0x17, 0x60, 0xd3, 0x1a, 0x47, 0x5f, 0x3b, 0xc0, 0x20, 0x65, 0x18,
	0xff, 0x76, 0xa0, 0x1f, 0xde, 0xe9, 0x27, 0x41, 0xe0, 0xd2, 0xb6, 0xd6, 0x0f, 0x31, 0x24, 0xea,
	0x88, 0x47, 0x30, 0xd8, 0xb0, 0x36, 0x17, 0x65, 0x23, 0x4b, 0x5e, 0x1f, 0xd6, 0x3f, 0x41, 0xf8,
	0x29, 0x74, 0x1b, 0x2a, 0xef, 0xf5, 0xca, 0xc3, 0x77, 0x1d, 0xe4, 0x10, 0xad, 0xf1, 0x6b, 0x70,
	0x45, 0x29, 0x82, 0xee, 0xc8, 0x99, 0x5c, 0x5c, 0x3f, 0x3b, 0xee, 0x42, 0x58, 0x51, 0xf2, 0x9a,
	0x6e, 0xd5, 0x6f, 0x2b, 0x05, 0x15, 0x44, 0xd9, 0xf0, 0x73, 0x80, 0x75, 0xd1, 0x64, 0x8d, 0x60,
	0xdf, 0xca, 0x1f, 0x41, 0x4f, 0x8f, 0xf1, 0xd6, 0x45, 0xb3, 0xd4, 0x00, 0x63, 0xe8, 0xd6, 0xb4,
	0x62, 0x41, 0x5f, 0x5f, 0xe8, 0xb3, 0xda, 0x3d, 0xe7, 0xbb, 0x5a, 0x8a, 0x7d, 0x96, 0xf3, 0x0d,
	0x0b, 0xce, 0x4c, 0x36, 0xcb, 0x66, 0x7c, 0xc3, 0xf0, 0x18, 0x7c, 0x2e, 0x0a, 0x5a, 0x97, 0x8f,
	0x54, 0xc7, 0x3f, 0x37, 0xff, 0xf0, 0x29, 0xbb, 0x5a, 0x80, 0x17, 0xb1, 0xad, 0xa4, 0xe9, 0xbe,
	0x61, 0xf8, 0x12, 0x06, 0x51, 0x9a, 0x91, 0xf8, 0x2e, 0x26, 0x5f, 0xe2, 0x08, 0xfd, 0x87, 0x87,
	0xe0, 0x45, 0x69, 0xb6, 0x5a, 0x46, 0x61, 0x1a, 0x23, 0xc7, 0xca, 0x28, 0x5e, 0xc4, 0x69, 0x8c,
	0x3a, 0x18, 0x81, 0x1f, 0xa5, 0x59, 0x72, 0x9b, 0xcd, 0x3e, 0x86, 0xc9, 0x87, 0x18, 0xb9, 0x57,
	0x3f, 0x1d, 0xf8, 0xff, 0xaf, 0x15, 0x95, 0x8f, 0xcc, 0xc9, 0x69, 0x5f, 0x1f, 0xce, 0x15, 0x09,
	0xc9, 0x3c, 0x41, 0xce, 0x41, 0x91, 0xf9, 0x52, 0x75, 0x1d, 0x82, 0xa7, 0xef, 0x96, 0xc9, 0x7c,
	0x86, 0x5c, 0x95, 0x49, 0xcb, 0xf7, 0x64, 0xae, 0x40, 0x17, 0x5f, 0x00, 0x28, 0xb0, 0x08, 0x67,
	0x4a, 0xf7, 0x0e, 0x86, 0x55, 0xf2, 0x29, 0xb9, 0xfd, 0x9a, 0xa0, 0xfe, 0xba, 0xaf, 0xbf, 0x97,
	0x9b, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x88, 0x4b, 0xe2, 0x3f, 0x03, 0x00, 0x00,
}
