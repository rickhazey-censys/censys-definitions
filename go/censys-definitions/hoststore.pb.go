// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hoststore.proto

package censys_definitions

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WHOISAtom struct {
}

func (m *WHOISAtom) Reset()                    { *m = WHOISAtom{} }
func (m *WHOISAtom) String() string            { return proto.CompactTextString(m) }
func (*WHOISAtom) ProtoMessage()               {}
func (*WHOISAtom) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type LocationAtom struct {
	Continent             string  `protobuf:"bytes,1,opt,name=continent" json:"continent,omitempty"`
	Country               string  `protobuf:"bytes,2,opt,name=country" json:"country,omitempty"`
	CountryCode           string  `protobuf:"bytes,3,opt,name=country_code,json=countryCode" json:"country_code,omitempty"`
	City                  string  `protobuf:"bytes,4,opt,name=city" json:"city,omitempty"`
	PostalCode            string  `protobuf:"bytes,5,opt,name=postal_code,json=postalCode" json:"postal_code,omitempty"`
	Timezone              string  `protobuf:"bytes,6,opt,name=timezone" json:"timezone,omitempty"`
	Province              string  `protobuf:"bytes,7,opt,name=province" json:"province,omitempty"`
	Latitude              float64 `protobuf:"fixed64,8,opt,name=latitude" json:"latitude,omitempty"`
	Longitude             float64 `protobuf:"fixed64,9,opt,name=longitude" json:"longitude,omitempty"`
	RegisteredCountry     string  `protobuf:"bytes,10,opt,name=registered_country,json=registeredCountry" json:"registered_country,omitempty"`
	RegisteredCountryCode string  `protobuf:"bytes,11,opt,name=registered_country_code,json=registeredCountryCode" json:"registered_country_code,omitempty"`
}

func (m *LocationAtom) Reset()                    { *m = LocationAtom{} }
func (m *LocationAtom) String() string            { return proto.CompactTextString(m) }
func (*LocationAtom) ProtoMessage()               {}
func (*LocationAtom) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *LocationAtom) GetContinent() string {
	if m != nil {
		return m.Continent
	}
	return ""
}

func (m *LocationAtom) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *LocationAtom) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *LocationAtom) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *LocationAtom) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

func (m *LocationAtom) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *LocationAtom) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *LocationAtom) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *LocationAtom) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *LocationAtom) GetRegisteredCountry() string {
	if m != nil {
		return m.RegisteredCountry
	}
	return ""
}

func (m *LocationAtom) GetRegisteredCountryCode() string {
	if m != nil {
		return m.RegisteredCountryCode
	}
	return ""
}

type ProtocolAtom struct {
	Metadata []*Metadatum `protobuf:"bytes,1,rep,name=metadata" json:"metadata,omitempty"`
	Tags     []string     `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty"`
	Data     string       `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *ProtocolAtom) Reset()                    { *m = ProtocolAtom{} }
func (m *ProtocolAtom) String() string            { return proto.CompactTextString(m) }
func (*ProtocolAtom) ProtoMessage()               {}
func (*ProtocolAtom) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *ProtocolAtom) GetMetadata() []*Metadatum {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ProtocolAtom) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ProtocolAtom) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type AnonymousKey struct {
	Port        uint32 `protobuf:"varint,1,opt,name=port" json:"port,omitempty"`
	Protocol    uint32 `protobuf:"varint,2,opt,name=protocol" json:"protocol,omitempty"`
	Subprotocol uint32 `protobuf:"varint,3,opt,name=subprotocol" json:"subprotocol,omitempty"`
}

func (m *AnonymousKey) Reset()                    { *m = AnonymousKey{} }
func (m *AnonymousKey) String() string            { return proto.CompactTextString(m) }
func (*AnonymousKey) ProtoMessage()               {}
func (*AnonymousKey) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *AnonymousKey) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *AnonymousKey) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *AnonymousKey) GetSubprotocol() uint32 {
	if m != nil {
		return m.Subprotocol
	}
	return 0
}

type Record struct {
	Ip                uint32 `protobuf:"fixed32,1,opt,name=ip" json:"ip,omitempty"`
	Port              uint32 `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	Protocol          uint32 `protobuf:"varint,3,opt,name=protocol" json:"protocol,omitempty"`
	Subprotocol       uint32 `protobuf:"varint,4,opt,name=subprotocol" json:"subprotocol,omitempty"`
	Domain            string `protobuf:"bytes,5,opt,name=domain" json:"domain,omitempty"`
	Timestamp         uint32 `protobuf:"fixed32,6,opt,name=timestamp" json:"timestamp,omitempty"`
	Scanid            uint32 `protobuf:"varint,7,opt,name=scanid" json:"scanid,omitempty"`
	Sha256Fp          []byte `protobuf:"bytes,8,opt,name=sha256fp,proto3" json:"sha256fp,omitempty"`
	FirstSeenAtScanId uint32 `protobuf:"varint,9,opt,name=first_seen_at_scan_id,json=firstSeenAtScanId" json:"first_seen_at_scan_id,omitempty"`
	LastSeenAtScanId  uint32 `protobuf:"varint,10,opt,name=last_seen_at_scan_id,json=lastSeenAtScanId" json:"last_seen_at_scan_id,omitempty"`
	// Types that are valid to be assigned to DataOneof:
	//	*Record_Atom
	//	*Record_PrivateLocation
	//	*Record_AsAtom
	//	*Record_Whois
	//	*Record_Userdata
	//	*Record_PublicLocation
	//	*Record_AlexaRank
	//	*Record_QuantcastRank
	//	*Record_CiscoUmbrellaRank
	DataOneof isRecord_DataOneof `protobuf_oneof:"data_oneof"`
	Version   uint64             `protobuf:"varint,18,opt,name=version" json:"version,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

type isRecord_DataOneof interface {
	isRecord_DataOneof()
}

type Record_Atom struct {
	Atom *ProtocolAtom `protobuf:"bytes,11,opt,name=atom,oneof"`
}
type Record_PrivateLocation struct {
	PrivateLocation *LocationAtom `protobuf:"bytes,12,opt,name=private_location,json=privateLocation,oneof"`
}
type Record_AsAtom struct {
	AsAtom *ASAtom `protobuf:"bytes,13,opt,name=as_atom,json=asAtom,oneof"`
}
type Record_Whois struct {
	Whois *WHOISAtom `protobuf:"bytes,14,opt,name=whois,oneof"`
}
type Record_Userdata struct {
	Userdata *UserdataAtom `protobuf:"bytes,15,opt,name=userdata,oneof"`
}
type Record_PublicLocation struct {
	PublicLocation *LocationAtom `protobuf:"bytes,17,opt,name=public_location,json=publicLocation,oneof"`
}
type Record_AlexaRank struct {
	AlexaRank uint32 `protobuf:"varint,16,opt,name=alexa_rank,json=alexaRank,oneof"`
}
type Record_QuantcastRank struct {
	QuantcastRank uint32 `protobuf:"varint,19,opt,name=quantcast_rank,json=quantcastRank,oneof"`
}
type Record_CiscoUmbrellaRank struct {
	CiscoUmbrellaRank uint32 `protobuf:"varint,20,opt,name=cisco_umbrella_rank,json=ciscoUmbrellaRank,oneof"`
}

func (*Record_Atom) isRecord_DataOneof()              {}
func (*Record_PrivateLocation) isRecord_DataOneof()   {}
func (*Record_AsAtom) isRecord_DataOneof()            {}
func (*Record_Whois) isRecord_DataOneof()             {}
func (*Record_Userdata) isRecord_DataOneof()          {}
func (*Record_PublicLocation) isRecord_DataOneof()    {}
func (*Record_AlexaRank) isRecord_DataOneof()         {}
func (*Record_QuantcastRank) isRecord_DataOneof()     {}
func (*Record_CiscoUmbrellaRank) isRecord_DataOneof() {}

func (m *Record) GetDataOneof() isRecord_DataOneof {
	if m != nil {
		return m.DataOneof
	}
	return nil
}

func (m *Record) GetIp() uint32 {
	if m != nil {
		return m.Ip
	}
	return 0
}

func (m *Record) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Record) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *Record) GetSubprotocol() uint32 {
	if m != nil {
		return m.Subprotocol
	}
	return 0
}

func (m *Record) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Record) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Record) GetScanid() uint32 {
	if m != nil {
		return m.Scanid
	}
	return 0
}

func (m *Record) GetSha256Fp() []byte {
	if m != nil {
		return m.Sha256Fp
	}
	return nil
}

func (m *Record) GetFirstSeenAtScanId() uint32 {
	if m != nil {
		return m.FirstSeenAtScanId
	}
	return 0
}

func (m *Record) GetLastSeenAtScanId() uint32 {
	if m != nil {
		return m.LastSeenAtScanId
	}
	return 0
}

func (m *Record) GetAtom() *ProtocolAtom {
	if x, ok := m.GetDataOneof().(*Record_Atom); ok {
		return x.Atom
	}
	return nil
}

func (m *Record) GetPrivateLocation() *LocationAtom {
	if x, ok := m.GetDataOneof().(*Record_PrivateLocation); ok {
		return x.PrivateLocation
	}
	return nil
}

func (m *Record) GetAsAtom() *ASAtom {
	if x, ok := m.GetDataOneof().(*Record_AsAtom); ok {
		return x.AsAtom
	}
	return nil
}

func (m *Record) GetWhois() *WHOISAtom {
	if x, ok := m.GetDataOneof().(*Record_Whois); ok {
		return x.Whois
	}
	return nil
}

func (m *Record) GetUserdata() *UserdataAtom {
	if x, ok := m.GetDataOneof().(*Record_Userdata); ok {
		return x.Userdata
	}
	return nil
}

func (m *Record) GetPublicLocation() *LocationAtom {
	if x, ok := m.GetDataOneof().(*Record_PublicLocation); ok {
		return x.PublicLocation
	}
	return nil
}

func (m *Record) GetAlexaRank() uint32 {
	if x, ok := m.GetDataOneof().(*Record_AlexaRank); ok {
		return x.AlexaRank
	}
	return 0
}

func (m *Record) GetQuantcastRank() uint32 {
	if x, ok := m.GetDataOneof().(*Record_QuantcastRank); ok {
		return x.QuantcastRank
	}
	return 0
}

func (m *Record) GetCiscoUmbrellaRank() uint32 {
	if x, ok := m.GetDataOneof().(*Record_CiscoUmbrellaRank); ok {
		return x.CiscoUmbrellaRank
	}
	return 0
}

func (m *Record) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Record) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Record_OneofMarshaler, _Record_OneofUnmarshaler, _Record_OneofSizer, []interface{}{
		(*Record_Atom)(nil),
		(*Record_PrivateLocation)(nil),
		(*Record_AsAtom)(nil),
		(*Record_Whois)(nil),
		(*Record_Userdata)(nil),
		(*Record_PublicLocation)(nil),
		(*Record_AlexaRank)(nil),
		(*Record_QuantcastRank)(nil),
		(*Record_CiscoUmbrellaRank)(nil),
	}
}

func _Record_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Record)
	// data_oneof
	switch x := m.DataOneof.(type) {
	case *Record_Atom:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Atom); err != nil {
			return err
		}
	case *Record_PrivateLocation:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PrivateLocation); err != nil {
			return err
		}
	case *Record_AsAtom:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AsAtom); err != nil {
			return err
		}
	case *Record_Whois:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Whois); err != nil {
			return err
		}
	case *Record_Userdata:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Userdata); err != nil {
			return err
		}
	case *Record_PublicLocation:
		b.EncodeVarint(17<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PublicLocation); err != nil {
			return err
		}
	case *Record_AlexaRank:
		b.EncodeVarint(16<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.AlexaRank))
	case *Record_QuantcastRank:
		b.EncodeVarint(19<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.QuantcastRank))
	case *Record_CiscoUmbrellaRank:
		b.EncodeVarint(20<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.CiscoUmbrellaRank))
	case nil:
	default:
		return fmt.Errorf("Record.DataOneof has unexpected type %T", x)
	}
	return nil
}

func _Record_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Record)
	switch tag {
	case 11: // data_oneof.atom
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProtocolAtom)
		err := b.DecodeMessage(msg)
		m.DataOneof = &Record_Atom{msg}
		return true, err
	case 12: // data_oneof.private_location
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LocationAtom)
		err := b.DecodeMessage(msg)
		m.DataOneof = &Record_PrivateLocation{msg}
		return true, err
	case 13: // data_oneof.as_atom
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ASAtom)
		err := b.DecodeMessage(msg)
		m.DataOneof = &Record_AsAtom{msg}
		return true, err
	case 14: // data_oneof.whois
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WHOISAtom)
		err := b.DecodeMessage(msg)
		m.DataOneof = &Record_Whois{msg}
		return true, err
	case 15: // data_oneof.userdata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UserdataAtom)
		err := b.DecodeMessage(msg)
		m.DataOneof = &Record_Userdata{msg}
		return true, err
	case 17: // data_oneof.public_location
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LocationAtom)
		err := b.DecodeMessage(msg)
		m.DataOneof = &Record_PublicLocation{msg}
		return true, err
	case 16: // data_oneof.alexa_rank
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.DataOneof = &Record_AlexaRank{uint32(x)}
		return true, err
	case 19: // data_oneof.quantcast_rank
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.DataOneof = &Record_QuantcastRank{uint32(x)}
		return true, err
	case 20: // data_oneof.cisco_umbrella_rank
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.DataOneof = &Record_CiscoUmbrellaRank{uint32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Record_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Record)
	// data_oneof
	switch x := m.DataOneof.(type) {
	case *Record_Atom:
		s := proto.Size(x.Atom)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Record_PrivateLocation:
		s := proto.Size(x.PrivateLocation)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Record_AsAtom:
		s := proto.Size(x.AsAtom)
		n += proto.SizeVarint(13<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Record_Whois:
		s := proto.Size(x.Whois)
		n += proto.SizeVarint(14<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Record_Userdata:
		s := proto.Size(x.Userdata)
		n += proto.SizeVarint(15<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Record_PublicLocation:
		s := proto.Size(x.PublicLocation)
		n += proto.SizeVarint(17<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Record_AlexaRank:
		n += proto.SizeVarint(16<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.AlexaRank))
	case *Record_QuantcastRank:
		n += proto.SizeVarint(19<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.QuantcastRank))
	case *Record_CiscoUmbrellaRank:
		n += proto.SizeVarint(20<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.CiscoUmbrellaRank))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Delta struct {
	DeltaType DeltaType `protobuf:"varint,1,opt,name=delta_type,json=deltaType,enum=zsearch.DeltaType" json:"delta_type,omitempty"`
	Ip        uint32    `protobuf:"fixed32,2,opt,name=ip" json:"ip,omitempty"`
	Domain    string    `protobuf:"bytes,3,opt,name=domain" json:"domain,omitempty"`
	Version   uint64    `protobuf:"varint,4,opt,name=version" json:"version,omitempty"`
	Records   []*Record `protobuf:"bytes,5,rep,name=records" json:"records,omitempty"`
}

func (m *Delta) Reset()                    { *m = Delta{} }
func (m *Delta) String() string            { return proto.CompactTextString(m) }
func (*Delta) ProtoMessage()               {}
func (*Delta) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *Delta) GetDeltaType() DeltaType {
	if m != nil {
		return m.DeltaType
	}
	return DeltaType_DT_RESERVED
}

func (m *Delta) GetIp() uint32 {
	if m != nil {
		return m.Ip
	}
	return 0
}

func (m *Delta) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Delta) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Delta) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*WHOISAtom)(nil), "zsearch.WHOISAtom")
	proto.RegisterType((*LocationAtom)(nil), "zsearch.LocationAtom")
	proto.RegisterType((*ProtocolAtom)(nil), "zsearch.ProtocolAtom")
	proto.RegisterType((*AnonymousKey)(nil), "zsearch.AnonymousKey")
	proto.RegisterType((*Record)(nil), "zsearch.Record")
	proto.RegisterType((*Delta)(nil), "zsearch.Delta")
}

func init() { proto.RegisterFile("hoststore.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 777 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdf, 0x93, 0xdb, 0x34,
	0x10, 0xbe, 0xfc, 0xb8, 0xf8, 0xb2, 0x71, 0x92, 0x3b, 0xb5, 0x07, 0x9e, 0x0e, 0x33, 0x0d, 0x79,
	0xe1, 0x28, 0x43, 0xa6, 0x5c, 0x87, 0x3e, 0x93, 0x96, 0x87, 0x74, 0x80, 0x81, 0x51, 0xe8, 0xf0,
	0x68, 0x14, 0x59, 0xb9, 0x68, 0x6a, 0x4b, 0x46, 0x92, 0x0f, 0xd2, 0x3f, 0x85, 0x7f, 0x80, 0x3f,
	0x91, 0x57, 0x46, 0x6b, 0xc5, 0xce, 0xf5, 0xe8, 0xbd, 0x69, 0xf7, 0xfb, 0x3e, 0x4b, 0xbb, 0xfb,
	0xad, 0x61, 0xba, 0xd3, 0xd6, 0x59, 0xa7, 0x8d, 0x58, 0x94, 0x46, 0x3b, 0x4d, 0xa2, 0xf7, 0x56,
	0x30, 0xc3, 0x77, 0x4f, 0x62, 0xae, 0x8b, 0x42, 0xab, 0x3a, 0x3d, 0x1f, 0xc1, 0xf0, 0xb7, 0xd5,
	0xcf, 0x6f, 0xd6, 0x4b, 0xa7, 0x8b, 0xf9, 0xbf, 0x5d, 0x88, 0x7f, 0xd4, 0x9c, 0x39, 0xa9, 0x95,
	0x4f, 0x90, 0xcf, 0x60, 0xc8, 0xb5, 0x72, 0x52, 0x09, 0xe5, 0x92, 0xce, 0xac, 0x73, 0x35, 0xa4,
	0x6d, 0x82, 0x24, 0x10, 0x71, 0x5d, 0x29, 0x67, 0xf6, 0x49, 0x17, 0xb1, 0x43, 0x48, 0x3e, 0x87,
	0x38, 0x1c, 0x53, 0xae, 0x33, 0x91, 0xf4, 0x10, 0x1e, 0x85, 0xdc, 0x6b, 0x9d, 0x09, 0x42, 0xa0,
	0xcf, 0xa5, 0xdb, 0x27, 0x7d, 0x84, 0xf0, 0x4c, 0x9e, 0xc2, 0xa8, 0xd4, 0xd6, 0xb1, 0xbc, 0x56,
	0x9d, 0x22, 0x04, 0x75, 0x0a, 0x45, 0x4f, 0xe0, 0xcc, 0xc9, 0x42, 0xbc, 0xd7, 0x4a, 0x24, 0x03,
	0x44, 0x9b, 0xd8, 0x63, 0xa5, 0xd1, 0xb7, 0x52, 0x71, 0x91, 0x44, 0x35, 0x76, 0x88, 0x3d, 0x96,
	0x33, 0x27, 0x5d, 0x95, 0x89, 0xe4, 0x6c, 0xd6, 0xb9, 0xea, 0xd0, 0x26, 0xf6, 0x35, 0xe6, 0x5a,
	0xdd, 0xd4, 0xe0, 0x10, 0xc1, 0x36, 0x41, 0xbe, 0x06, 0x62, 0xc4, 0x8d, 0xb4, 0x4e, 0x18, 0x91,
	0xa5, 0x87, 0x72, 0x01, 0xbf, 0x7f, 0xd1, 0x22, 0xaf, 0x43, 0xe1, 0x2f, 0xe1, 0xd3, 0xfb, 0xf4,
	0xba, 0x9a, 0x11, 0x6a, 0x2e, 0xef, 0x69, 0x7c, 0x61, 0xf3, 0x2d, 0xc4, 0xbf, 0xf8, 0x79, 0x70,
	0x9d, 0x63, 0xe3, 0x17, 0x70, 0x56, 0x08, 0xc7, 0x32, 0xe6, 0x58, 0xd2, 0x99, 0xf5, 0xae, 0x46,
	0xd7, 0x64, 0x11, 0x06, 0xb8, 0xf8, 0xa9, 0x06, 0xaa, 0x82, 0x36, 0x1c, 0xdf, 0x4d, 0xc7, 0x6e,
	0x6c, 0xd2, 0x9d, 0xf5, 0x7c, 0x37, 0xfd, 0xd9, 0xe7, 0x50, 0x5f, 0x37, 0x1f, 0xcf, 0xf3, 0xdf,
	0x21, 0x5e, 0x2a, 0xad, 0xf6, 0x85, 0xae, 0xec, 0x0f, 0x62, 0xef, 0x39, 0xa5, 0x36, 0xf5, 0x6c,
	0xc7, 0x14, 0xcf, 0xa1, 0x91, 0xf8, 0x16, 0x9c, 0xeb, 0x98, 0x36, 0x31, 0x99, 0xc1, 0xc8, 0x56,
	0x9b, 0x06, 0xee, 0x21, 0x7c, 0x9c, 0x9a, 0xff, 0x3d, 0x80, 0x01, 0x15, 0x5c, 0x9b, 0x8c, 0x4c,
	0xa0, 0x2b, 0x4b, 0xfc, 0x74, 0x44, 0xbb, 0xb2, 0x6c, 0x2e, 0xeb, 0x7e, 0xe4, 0xb2, 0xde, 0xc3,
	0x97, 0xf5, 0xef, 0x5d, 0x46, 0x3e, 0x81, 0x41, 0xa6, 0x0b, 0x26, 0x55, 0xf0, 0x4a, 0x88, 0xfc,
	0x4c, 0xbd, 0x2f, 0xac, 0x63, 0x45, 0x89, 0x46, 0x89, 0x68, 0x9b, 0xf0, 0x2a, 0xcb, 0x99, 0x92,
	0x19, 0xfa, 0x64, 0x4c, 0x43, 0xe4, 0xdf, 0x62, 0x77, 0xec, 0xfa, 0xdb, 0x97, 0xdb, 0x12, 0x5d,
	0x12, 0xd3, 0x26, 0x26, 0xcf, 0xe1, 0x72, 0x2b, 0x8d, 0x75, 0xa9, 0x15, 0x42, 0xa5, 0xcc, 0xa5,
	0x5e, 0x93, 0xca, 0x0c, 0x1d, 0x33, 0xa6, 0x17, 0x08, 0xae, 0x85, 0x50, 0x4b, 0xb7, 0xe6, 0x4c,
	0xbd, 0xc9, 0xc8, 0x02, 0x1e, 0xe7, 0xec, 0x7f, 0x04, 0x80, 0x82, 0x73, 0x8f, 0xdd, 0xe1, 0x7f,
	0x05, 0x7d, 0xe6, 0x74, 0x81, 0x3e, 0x19, 0x5d, 0x5f, 0x36, 0xe3, 0x3e, 0xf6, 0xc5, 0xea, 0x84,
	0x22, 0x89, 0xbc, 0x82, 0xf3, 0xd2, 0xc8, 0x5b, 0xe6, 0x44, 0x9a, 0x87, 0x85, 0x4d, 0xe2, 0x0f,
	0x84, 0xc7, 0x9b, 0xbc, 0x3a, 0xa1, 0xd3, 0x20, 0x38, 0xa4, 0xc9, 0x33, 0x88, 0x98, 0x4d, 0xf1,
	0xce, 0x31, 0x4a, 0xa7, 0x8d, 0x74, 0xb9, 0x0e, 0xa2, 0x01, 0xb3, 0xe8, 0xc7, 0x67, 0x70, 0xfa,
	0xe7, 0x4e, 0x4b, 0x9b, 0x4c, 0x90, 0xd9, 0x9a, 0xb1, 0xf9, 0x79, 0xac, 0x4e, 0x68, 0x4d, 0x21,
	0x2f, 0xe0, 0xac, 0xb2, 0xc2, 0xa0, 0xf7, 0xa6, 0x1f, 0xbc, 0xe9, 0x6d, 0x00, 0x82, 0xa2, 0x21,
	0x92, 0xef, 0x60, 0x5a, 0x56, 0x9b, 0x5c, 0xf2, 0xb6, 0x9e, 0x8b, 0x87, 0xeb, 0x99, 0xd4, 0xfc,
	0xa6, 0x9c, 0xa7, 0x00, 0x2c, 0x17, 0x7f, 0xb1, 0xd4, 0x30, 0xf5, 0x2e, 0x39, 0xf7, 0x5d, 0x5e,
	0x9d, 0xd0, 0x21, 0xe6, 0x28, 0x53, 0xef, 0xc8, 0x17, 0x30, 0xf9, 0xa3, 0x62, 0xca, 0x71, 0x3f,
	0x15, 0x24, 0x3d, 0x0a, 0xa4, 0x71, 0x93, 0x47, 0xe2, 0x73, 0x78, 0xc4, 0xa5, 0xe5, 0x3a, 0xad,
	0x8a, 0x8d, 0x11, 0x79, 0x1e, 0x3e, 0xf9, 0x38, 0xb0, 0x2f, 0x10, 0x7c, 0x1b, 0x30, 0x54, 0x24,
	0x10, 0xdd, 0x0a, 0x63, 0xfd, 0xab, 0xc9, 0xac, 0x73, 0xd5, 0xa7, 0x87, 0xf0, 0x55, 0x0c, 0xe0,
	0xeb, 0x4b, 0xb5, 0x12, 0x7a, 0x3b, 0xff, 0xa7, 0x03, 0xa7, 0xdf, 0x8b, 0xdc, 0x31, 0xf2, 0x0d,
	0x40, 0xe6, 0x0f, 0xa9, 0xdb, 0x97, 0x02, 0x77, 0x64, 0x72, 0xd4, 0x55, 0xe4, 0xfc, 0xba, 0x2f,
	0x05, 0x1d, 0x66, 0x87, 0x63, 0x58, 0xa7, 0x6e, 0xb3, 0x4e, 0xad, 0xf9, 0x7b, 0x77, 0xcc, 0x7f,
	0xf4, 0x98, 0xfe, 0x9d, 0xc7, 0x90, 0x2f, 0x21, 0x32, 0xb8, 0x9a, 0x36, 0x39, 0xc5, 0x9f, 0x4a,
	0x3b, 0xf1, 0x7a, 0x65, 0xe9, 0x01, 0xdf, 0x0c, 0x70, 0xc7, 0x5e, 0xfc, 0x17, 0x00, 0x00, 0xff,
	0xff, 0x9e, 0x04, 0x42, 0xc1, 0x48, 0x06, 0x00, 0x00,
}
