// Code generated by protoc-gen-go.
// source: protocols.proto
// DO NOT EDIT!

package censys_definitions

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Protocol int32

const (
	Protocol_PROTO_RESERVED  Protocol = 0
	Protocol_PROTO_SYSTEM    Protocol = 1
	Protocol_PROTO_HTTP      Protocol = 2
	Protocol_PROTO_HTTPS     Protocol = 3
	Protocol_PROTO_IMAP      Protocol = 4
	Protocol_PROTO_IMAPS     Protocol = 5
	Protocol_PROTO_SMTP      Protocol = 6
	Protocol_PROTO_SMTPS     Protocol = 7
	Protocol_PROTO_POP3      Protocol = 8
	Protocol_PROTO_POP3S     Protocol = 9
	Protocol_PROTO_MODBUS    Protocol = 10
	Protocol_PROTO_FTP       Protocol = 11
	Protocol_PROTO_SSH       Protocol = 12
	Protocol_PROTO_DNS       Protocol = 13
	Protocol_PROTO_NTP       Protocol = 14
	Protocol_PROTO_TELNET    Protocol = 15
	Protocol_PROTO_UPNP      Protocol = 16
	Protocol_PROTO_CWMP      Protocol = 17
	Protocol_PROTO_HTTP_2    Protocol = 18
	Protocol_PROTO_BACNET    Protocol = 19
	Protocol_PROTO_DNP3      Protocol = 20
	Protocol_PROTO_FOX       Protocol = 21
	Protocol_PROTO_S7        Protocol = 22
	Protocol_PROTO_GLOBAL    Protocol = 23
	Protocol_PROTO_LOOKUP    Protocol = 24
	Protocol_PROTO_HTTP_WWW  Protocol = 26
	Protocol_PROTO_HTTPS_WWW Protocol = 27
)

var Protocol_name = map[int32]string{
	0:  "PROTO_RESERVED",
	1:  "PROTO_SYSTEM",
	2:  "PROTO_HTTP",
	3:  "PROTO_HTTPS",
	4:  "PROTO_IMAP",
	5:  "PROTO_IMAPS",
	6:  "PROTO_SMTP",
	7:  "PROTO_SMTPS",
	8:  "PROTO_POP3",
	9:  "PROTO_POP3S",
	10: "PROTO_MODBUS",
	11: "PROTO_FTP",
	12: "PROTO_SSH",
	13: "PROTO_DNS",
	14: "PROTO_NTP",
	15: "PROTO_TELNET",
	16: "PROTO_UPNP",
	17: "PROTO_CWMP",
	18: "PROTO_HTTP_2",
	19: "PROTO_BACNET",
	20: "PROTO_DNP3",
	21: "PROTO_FOX",
	22: "PROTO_S7",
	23: "PROTO_GLOBAL",
	24: "PROTO_LOOKUP",
	26: "PROTO_HTTP_WWW",
	27: "PROTO_HTTPS_WWW",
}
var Protocol_value = map[string]int32{
	"PROTO_RESERVED":  0,
	"PROTO_SYSTEM":    1,
	"PROTO_HTTP":      2,
	"PROTO_HTTPS":     3,
	"PROTO_IMAP":      4,
	"PROTO_IMAPS":     5,
	"PROTO_SMTP":      6,
	"PROTO_SMTPS":     7,
	"PROTO_POP3":      8,
	"PROTO_POP3S":     9,
	"PROTO_MODBUS":    10,
	"PROTO_FTP":       11,
	"PROTO_SSH":       12,
	"PROTO_DNS":       13,
	"PROTO_NTP":       14,
	"PROTO_TELNET":    15,
	"PROTO_UPNP":      16,
	"PROTO_CWMP":      17,
	"PROTO_HTTP_2":    18,
	"PROTO_BACNET":    19,
	"PROTO_DNP3":      20,
	"PROTO_FOX":       21,
	"PROTO_S7":        22,
	"PROTO_GLOBAL":    23,
	"PROTO_LOOKUP":    24,
	"PROTO_HTTP_WWW":  26,
	"PROTO_HTTPS_WWW": 27,
}

func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}
func (Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type Subprotocol int32

const (
	Subprotocol_SUBPROTO_RESERVED      Subprotocol = 0
	Subprotocol_SUBPROTO_DELETED       Subprotocol = 1
	Subprotocol_SUBPROTO_GENERIC       Subprotocol = 2
	Subprotocol_SUBPROTO_BANNER        Subprotocol = 3
	Subprotocol_SUBPROTO_TLS           Subprotocol = 4
	Subprotocol_SUBPROTO_TLS_1_0       Subprotocol = 5
	Subprotocol_SUBPROTO_TLS_1_1       Subprotocol = 6
	Subprotocol_SUBPROTO_TLS_1_2       Subprotocol = 7
	Subprotocol_SUBPROTO_TLS_1_3       Subprotocol = 8
	Subprotocol_SUBPROTO_HEARTBLEED    Subprotocol = 9
	Subprotocol_SUBPROTO_CIPHERS       Subprotocol = 10
	Subprotocol_SUBPROTO_SSL_2         Subprotocol = 11
	Subprotocol_SUBPROTO_SSL_3         Subprotocol = 12
	Subprotocol_SUBPROTO_GET           Subprotocol = 13
	Subprotocol_SUBPROTO_STARTTLS      Subprotocol = 14
	Subprotocol_SUBPROTO_EXPORT        Subprotocol = 15
	Subprotocol_SUBPROTO_RSA_EXPORT    Subprotocol = 16
	Subprotocol_SUBPROTO_DHE_EXPORT    Subprotocol = 17
	Subprotocol_SUBPROTO_DHE           Subprotocol = 18
	Subprotocol_SUBPROTO_ECDHE         Subprotocol = 19
	Subprotocol_SUBPROTO_SNI           Subprotocol = 20
	Subprotocol_SUBPROTO_NO_SNI        Subprotocol = 21
	Subprotocol_SUBPROTO_QUIC          Subprotocol = 22
	Subprotocol_SUBPROTO_SPDY          Subprotocol = 23
	Subprotocol_SUBPROTO_RSA           Subprotocol = 24
	Subprotocol_SUBPROTO_DSA           Subprotocol = 25
	Subprotocol_SUBPROTO_ECDSA         Subprotocol = 26
	Subprotocol_SUBPROTO_DEVICE_ID     Subprotocol = 27
	Subprotocol_SUBPROTO_OPEN_RESOLVER Subprotocol = 28
	// that we had to later scrap.
	Subprotocol_SUBPROTO_OPEN_PROXY      Subprotocol = 29
	Subprotocol_SUBPROTO_OPEN_RELAY      Subprotocol = 30
	Subprotocol_SUBPROTO_TIME            Subprotocol = 31
	Subprotocol_SUBPROTO_HACKING_TEAM    Subprotocol = 32
	Subprotocol_SUBPROTO_EXTENDED_RANDOM Subprotocol = 33
	Subprotocol_SUBPROTO_DISCOVERY       Subprotocol = 34
	Subprotocol_SUBPROTO_GTLD_A          Subprotocol = 35
	Subprotocol_SUBPROTO_LOOKUP          Subprotocol = 36
	Subprotocol_SUBPROTO_STATUS          Subprotocol = 37
	Subprotocol_SUBPROTO_SZL             Subprotocol = 38
	// System Records
	Subprotocol_SUBPROTO_SYS_PUBLIC_LOCATION     Subprotocol = 192
	Subprotocol_SUBPROTO_SYS_AS                  Subprotocol = 193
	Subprotocol_SUBPROTO_SYS_TAGS                Subprotocol = 194
	Subprotocol_SUBPROTO_SYS_METADATA            Subprotocol = 195
	Subprotocol_SUBPROTO_SYS_WHOIS               Subprotocol = 196
	Subprotocol_SUBPROTO_SYS_USERDATA            Subprotocol = 197
	Subprotocol_SUBPROTO_SYS_BLACKLIST           Subprotocol = 198
	Subprotocol_SUBPROTO_SYS_ALEXA_RANK          Subprotocol = 199
	Subprotocol_SUBPROTO_SYS_RESTRICTED_LOCATION Subprotocol = 200
	Subprotocol_SUBPROTO_SYS_VERSION             Subprotocol = 201
	Subprotocol_SUBPROTO_SYS_QUANTCAST_RANK      Subprotocol = 202
	Subprotocol_SUBPROTO_SYS_CISCO_UMBRELLA_RANK Subprotocol = 203
	Subprotocol_SUBPROTO_SYS_REVERSE_DNS         Subprotocol = 204
	// Non System Global data. This is generally used
	// for things about Alexa domains, such as DNS records, etc.
	Subprotocol_SUBPROTO_SPF   Subprotocol = 220
	Subprotocol_SUBPROTO_DMARC Subprotocol = 221
	Subprotocol_SUBPROTO_DKIM  Subprotocol = 222
	Subprotocol_SUBPROTO_A     Subprotocol = 223
	Subprotocol_SUBPROTO_MX    Subprotocol = 224
	Subprotocol_SUBPROTO_AXFR  Subprotocol = 225
)

var Subprotocol_name = map[int32]string{
	0:   "SUBPROTO_RESERVED",
	1:   "SUBPROTO_DELETED",
	2:   "SUBPROTO_GENERIC",
	3:   "SUBPROTO_BANNER",
	4:   "SUBPROTO_TLS",
	5:   "SUBPROTO_TLS_1_0",
	6:   "SUBPROTO_TLS_1_1",
	7:   "SUBPROTO_TLS_1_2",
	8:   "SUBPROTO_TLS_1_3",
	9:   "SUBPROTO_HEARTBLEED",
	10:  "SUBPROTO_CIPHERS",
	11:  "SUBPROTO_SSL_2",
	12:  "SUBPROTO_SSL_3",
	13:  "SUBPROTO_GET",
	14:  "SUBPROTO_STARTTLS",
	15:  "SUBPROTO_EXPORT",
	16:  "SUBPROTO_RSA_EXPORT",
	17:  "SUBPROTO_DHE_EXPORT",
	18:  "SUBPROTO_DHE",
	19:  "SUBPROTO_ECDHE",
	20:  "SUBPROTO_SNI",
	21:  "SUBPROTO_NO_SNI",
	22:  "SUBPROTO_QUIC",
	23:  "SUBPROTO_SPDY",
	24:  "SUBPROTO_RSA",
	25:  "SUBPROTO_DSA",
	26:  "SUBPROTO_ECDSA",
	27:  "SUBPROTO_DEVICE_ID",
	28:  "SUBPROTO_OPEN_RESOLVER",
	29:  "SUBPROTO_OPEN_PROXY",
	30:  "SUBPROTO_OPEN_RELAY",
	31:  "SUBPROTO_TIME",
	32:  "SUBPROTO_HACKING_TEAM",
	33:  "SUBPROTO_EXTENDED_RANDOM",
	34:  "SUBPROTO_DISCOVERY",
	35:  "SUBPROTO_GTLD_A",
	36:  "SUBPROTO_LOOKUP",
	37:  "SUBPROTO_STATUS",
	38:  "SUBPROTO_SZL",
	192: "SUBPROTO_SYS_PUBLIC_LOCATION",
	193: "SUBPROTO_SYS_AS",
	194: "SUBPROTO_SYS_TAGS",
	195: "SUBPROTO_SYS_METADATA",
	196: "SUBPROTO_SYS_WHOIS",
	197: "SUBPROTO_SYS_USERDATA",
	198: "SUBPROTO_SYS_BLACKLIST",
	199: "SUBPROTO_SYS_ALEXA_RANK",
	200: "SUBPROTO_SYS_RESTRICTED_LOCATION",
	201: "SUBPROTO_SYS_VERSION",
	202: "SUBPROTO_SYS_QUANTCAST_RANK",
	203: "SUBPROTO_SYS_CISCO_UMBRELLA_RANK",
	204: "SUBPROTO_SYS_REVERSE_DNS",
	220: "SUBPROTO_SPF",
	221: "SUBPROTO_DMARC",
	222: "SUBPROTO_DKIM",
	223: "SUBPROTO_A",
	224: "SUBPROTO_MX",
	225: "SUBPROTO_AXFR",
}
var Subprotocol_value = map[string]int32{
	"SUBPROTO_RESERVED":                0,
	"SUBPROTO_DELETED":                 1,
	"SUBPROTO_GENERIC":                 2,
	"SUBPROTO_BANNER":                  3,
	"SUBPROTO_TLS":                     4,
	"SUBPROTO_TLS_1_0":                 5,
	"SUBPROTO_TLS_1_1":                 6,
	"SUBPROTO_TLS_1_2":                 7,
	"SUBPROTO_TLS_1_3":                 8,
	"SUBPROTO_HEARTBLEED":              9,
	"SUBPROTO_CIPHERS":                 10,
	"SUBPROTO_SSL_2":                   11,
	"SUBPROTO_SSL_3":                   12,
	"SUBPROTO_GET":                     13,
	"SUBPROTO_STARTTLS":                14,
	"SUBPROTO_EXPORT":                  15,
	"SUBPROTO_RSA_EXPORT":              16,
	"SUBPROTO_DHE_EXPORT":              17,
	"SUBPROTO_DHE":                     18,
	"SUBPROTO_ECDHE":                   19,
	"SUBPROTO_SNI":                     20,
	"SUBPROTO_NO_SNI":                  21,
	"SUBPROTO_QUIC":                    22,
	"SUBPROTO_SPDY":                    23,
	"SUBPROTO_RSA":                     24,
	"SUBPROTO_DSA":                     25,
	"SUBPROTO_ECDSA":                   26,
	"SUBPROTO_DEVICE_ID":               27,
	"SUBPROTO_OPEN_RESOLVER":           28,
	"SUBPROTO_OPEN_PROXY":              29,
	"SUBPROTO_OPEN_RELAY":              30,
	"SUBPROTO_TIME":                    31,
	"SUBPROTO_HACKING_TEAM":            32,
	"SUBPROTO_EXTENDED_RANDOM":         33,
	"SUBPROTO_DISCOVERY":               34,
	"SUBPROTO_GTLD_A":                  35,
	"SUBPROTO_LOOKUP":                  36,
	"SUBPROTO_STATUS":                  37,
	"SUBPROTO_SZL":                     38,
	"SUBPROTO_SYS_PUBLIC_LOCATION":     192,
	"SUBPROTO_SYS_AS":                  193,
	"SUBPROTO_SYS_TAGS":                194,
	"SUBPROTO_SYS_METADATA":            195,
	"SUBPROTO_SYS_WHOIS":               196,
	"SUBPROTO_SYS_USERDATA":            197,
	"SUBPROTO_SYS_BLACKLIST":           198,
	"SUBPROTO_SYS_ALEXA_RANK":          199,
	"SUBPROTO_SYS_RESTRICTED_LOCATION": 200,
	"SUBPROTO_SYS_VERSION":             201,
	"SUBPROTO_SYS_QUANTCAST_RANK":      202,
	"SUBPROTO_SYS_CISCO_UMBRELLA_RANK": 203,
	"SUBPROTO_SYS_REVERSE_DNS":         204,
	"SUBPROTO_SPF":                     220,
	"SUBPROTO_DMARC":                   221,
	"SUBPROTO_DKIM":                    222,
	"SUBPROTO_A":                       223,
	"SUBPROTO_MX":                      224,
	"SUBPROTO_AXFR":                    225,
}

func (x Subprotocol) String() string {
	return proto.EnumName(Subprotocol_name, int32(x))
}
func (Subprotocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func init() {
	proto.RegisterEnum("zsearch.Protocol", Protocol_name, Protocol_value)
	proto.RegisterEnum("zsearch.Subprotocol", Subprotocol_name, Subprotocol_value)
}

func init() { proto.RegisterFile("protocols.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 829 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x95, 0x4b, 0x73, 0xeb, 0x34,
	0x14, 0xc7, 0x11, 0x17, 0xee, 0x6d, 0x95, 0x26, 0x39, 0x51, 0x9a, 0xf4, 0x91, 0x16, 0x5a, 0xa0,
	0x2c, 0xba, 0x60, 0x68, 0xbb, 0x60, 0x2d, 0xdb, 0x6a, 0xe2, 0x89, 0x6c, 0xab, 0x96, 0xdc, 0x24,
	0x6c, 0x3c, 0xb4, 0xc3, 0x0c, 0x0b, 0x66, 0xca, 0xb4, 0xb0, 0xe1, 0xd3, 0xf1, 0x7e, 0xaf, 0x58,
	0xb0, 0xe2, 0xf9, 0x4d, 0x38, 0x8a, 0x5b, 0xcb, 0xba, 0xde, 0x45, 0xbf, 0x73, 0xa4, 0xf3, 0xd7,
	0xf9, 0x1f, 0x2b, 0xb4, 0xff, 0xd9, 0xfd, 0xdd, 0xe7, 0x77, 0xb7, 0x77, 0x9f, 0x3e, 0xbc, 0xb7,
	0xfe, 0xc5, 0x5e, 0x7c, 0xf9, 0xf0, 0xf1, 0x47, 0xf7, 0xb7, 0x9f, 0x9c, 0xfe, 0xfe, 0x8c, 0x6e,
	0xa8, 0xc7, 0x20, 0x63, 0xb4, 0xa7, 0xf2, 0xcc, 0x64, 0x65, 0x2e, 0xb4, 0xc8, 0xaf, 0x45, 0x04,
	0xaf, 0x30, 0xa0, 0x5b, 0x15, 0xd3, 0x2b, 0x6d, 0x44, 0x02, 0x84, 0xf5, 0x28, 0xad, 0xc8, 0xcc,
	0x18, 0x05, 0xaf, 0xb2, 0x3e, 0xed, 0xb8, 0xb5, 0x86, 0x67, 0x2e, 0x21, 0x4e, 0xb8, 0x82, 0xd7,
	0x5c, 0x82, 0x5d, 0x6b, 0x78, 0xdd, 0x25, 0xe8, 0x04, 0x4f, 0x78, 0xee, 0x12, 0xec, 0x5a, 0xc3,
	0x0b, 0x97, 0xa0, 0x32, 0x75, 0x01, 0x1b, 0x2e, 0xc1, 0xae, 0x35, 0x6c, 0x3a, 0x55, 0x49, 0x16,
	0x05, 0x85, 0x06, 0xca, 0xba, 0x74, 0xb3, 0x22, 0x97, 0x78, 0x64, 0xc7, 0x2d, 0xb5, 0x9e, 0xc1,
	0x96, 0x5b, 0x46, 0xa9, 0x86, 0xae, 0x5b, 0xa6, 0x98, 0xdc, 0x73, 0xa7, 0x19, 0x21, 0x53, 0x61,
	0xa0, 0xef, 0x04, 0x14, 0x2a, 0x55, 0x00, 0x6e, 0x1d, 0x2e, 0x12, 0x05, 0x03, 0xb7, 0xc3, 0xde,
	0xb9, 0x3c, 0x07, 0xe6, 0x48, 0xc0, 0x43, 0x7b, 0xc6, 0xd0, 0xed, 0x89, 0x52, 0xbc, 0xc4, 0x76,
	0x43, 0x61, 0xb6, 0x84, 0x11, 0xdb, 0xc2, 0xc6, 0x57, 0x0a, 0x3f, 0x80, 0xb1, 0xdb, 0x3e, 0x95,
	0x59, 0xc0, 0x25, 0xec, 0x38, 0x22, 0xb3, 0x6c, 0x5e, 0x28, 0xd8, 0x75, 0xf6, 0xac, 0x8b, 0x2e,
	0x16, 0x0b, 0xd8, 0x67, 0x43, 0xda, 0x6f, 0x34, 0x7f, 0x0d, 0x27, 0xa7, 0x7f, 0x50, 0xda, 0xd1,
	0x5f, 0xdc, 0x3c, 0x99, 0xce, 0x46, 0x74, 0xa0, 0x8b, 0xa0, 0x65, 0xed, 0x36, 0x85, 0x1a, 0x47,
	0x42, 0x0a, 0x83, 0x94, 0x78, 0x74, 0x2a, 0x52, 0x91, 0xc7, 0x21, 0x9a, 0x8c, 0x75, 0x6a, 0x1a,
	0xf0, 0x14, 0x31, 0x1a, 0x8d, 0x12, 0x6b, 0x68, 0xa4, 0x46, 0xab, 0x9b, 0x9b, 0x91, 0x94, 0x67,
	0xe5, 0xfb, 0xe8, 0x77, 0x9b, 0x9e, 0xa1, 0xeb, 0x6d, 0x7a, 0x8e, 0xd6, 0xb7, 0xa9, 0x1d, 0x80,
	0x1d, 0x3a, 0xac, 0xe9, 0x4c, 0xf0, 0xdc, 0x04, 0x52, 0xa0, 0xda, 0x4d, 0x2f, 0x3d, 0x8c, 0xd5,
	0x4c, 0xe4, 0x76, 0x18, 0xb0, 0x53, 0x35, 0xd5, 0x5a, 0xe2, 0xc1, 0x9d, 0x16, 0xbb, 0xc0, 0xb1,
	0x68, 0x5e, 0x60, 0x8a, 0xa6, 0x75, 0xbd, 0x56, 0x69, 0x83, 0x85, 0xec, 0xbd, 0x7a, 0xde, 0xf5,
	0xc5, 0x52, 0x65, 0xb9, 0x1d, 0x92, 0xa6, 0xa8, 0x5c, 0xf3, 0xa7, 0x00, 0x78, 0x81, 0x68, 0x26,
	0x9e, 0x02, 0x03, 0xaf, 0x1e, 0x06, 0x70, 0x6c, 0x9a, 0xaa, 0x44, 0x68, 0xd9, 0xd0, 0xcb, 0xd2,
	0x69, 0x8c, 0xa3, 0xd3, 0x2c, 0x9f, 0x56, 0x70, 0xc4, 0x06, 0xb4, 0x5b, 0xc3, 0xab, 0x02, 0x5d,
	0x1a, 0x7b, 0x48, 0xab, 0x68, 0x55, 0x8d, 0x51, 0x53, 0x24, 0x8e, 0x91, 0x27, 0x02, 0xc9, 0xde,
	0xcb, 0x22, 0x90, 0xed, 0xb3, 0x31, 0x65, 0x8d, 0xe1, 0xb8, 0x8e, 0x43, 0x51, 0xc6, 0x11, 0x4c,
	0xd8, 0x3e, 0x1d, 0xd7, 0x3c, 0x53, 0x22, 0xb5, 0x03, 0x95, 0xc9, 0x6b, 0x9c, 0x87, 0x03, 0xef,
	0xde, 0xeb, 0x18, 0xfe, 0x5c, 0xae, 0xe0, 0xb0, 0x1d, 0xc8, 0x85, 0xe4, 0x2b, 0x78, 0xc3, 0x13,
	0x6c, 0xe2, 0x44, 0xc0, 0x9b, 0x6c, 0x8f, 0x8e, 0x9c, 0xd5, 0x3c, 0x9c, 0xc7, 0xe9, 0x14, 0xbf,
	0x4b, 0x9e, 0xc0, 0x11, 0x3b, 0xa0, 0xbb, 0x0d, 0x17, 0x8c, 0x48, 0x23, 0x11, 0x95, 0x39, 0x4f,
	0xa3, 0x2c, 0x81, 0x63, 0x5f, 0x71, 0xac, 0xc3, 0x0c, 0x45, 0xad, 0xe0, 0x2d, 0xaf, 0x79, 0x53,
	0x23, 0xa3, 0x92, 0xc3, 0xdb, 0x1e, 0x7c, 0xfc, 0xc0, 0xde, 0xf1, 0x20, 0x9a, 0x6f, 0xf0, 0x61,
	0x39, 0xf1, 0xdd, 0xf8, 0x50, 0xc2, 0xbb, 0xec, 0x98, 0x1e, 0x38, 0xb2, 0xd2, 0xa5, 0x2a, 0x02,
	0x19, 0x87, 0x78, 0x4c, 0xc8, 0x4d, 0x9c, 0xa5, 0xf0, 0x95, 0xfd, 0x88, 0xfa, 0x5e, 0x0a, 0xd7,
	0xf0, 0x35, 0x41, 0x85, 0x03, 0x8f, 0x1a, 0x3e, 0xd5, 0xf0, 0x0d, 0xc1, 0x9e, 0x8e, 0x3c, 0x9e,
	0x08, 0xc3, 0x23, 0x6e, 0x38, 0x7c, 0x4b, 0xb0, 0x75, 0xcc, 0x8b, 0x2d, 0x66, 0x59, 0xac, 0xe1,
	0xbb, 0xf6, 0xa6, 0x02, 0xbf, 0xeb, 0xf5, 0xa6, 0xef, 0x09, 0x9b, 0x34, 0x4c, 0xb2, 0xb1, 0x40,
	0x62, 0x23, 0x65, 0xac, 0x0d, 0xfc, 0x40, 0xb0, 0x8b, 0x3b, 0xbe, 0x36, 0x29, 0x96, 0xdc, 0xb6,
	0x71, 0x0e, 0x3f, 0x12, 0x76, 0x42, 0x8f, 0xbc, 0x28, 0xda, 0x6b, 0xf0, 0x09, 0xc0, 0xb7, 0xc1,
	0x5d, 0xf0, 0x27, 0x82, 0x2e, 0x6d, 0x7b, 0x69, 0xd8, 0x6b, 0x6d, 0x43, 0x3f, 0x13, 0x76, 0x44,
	0x27, 0x5e, 0xe8, 0xaa, 0xe0, 0xa9, 0x09, 0xb9, 0x36, 0x55, 0x8d, 0x5f, 0xda, 0x35, 0x42, 0xeb,
	0x56, 0x59, 0x24, 0x01, 0x8e, 0x85, 0x7c, 0x94, 0xf2, 0x2b, 0x61, 0x87, 0x0d, 0xbb, 0x2b, 0x29,
	0xb6, 0x8a, 0x58, 0xbf, 0xe1, 0xbf, 0x11, 0x9c, 0x9d, 0x86, 0x31, 0xea, 0x12, 0xfe, 0x24, 0x68,
	0xa0, 0x1b, 0xe4, 0x28, 0xe1, 0x79, 0x08, 0x7f, 0x11, 0x9c, 0x6e, 0x37, 0x63, 0xd1, 0x3c, 0x4e,
	0xe0, 0x6f, 0x82, 0x7f, 0x28, 0xb4, 0x66, 0x1c, 0xfe, 0x21, 0xe8, 0x72, 0xa7, 0x06, 0xc9, 0x12,
	0xfe, 0xf5, 0xb7, 0xf1, 0xe5, 0x65, 0x0e, 0xff, 0x91, 0x9b, 0xe7, 0xeb, 0x27, 0xf5, 0xe2, 0xff,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x9b, 0x88, 0x08, 0x50, 0x07, 0x00, 0x00,
}
