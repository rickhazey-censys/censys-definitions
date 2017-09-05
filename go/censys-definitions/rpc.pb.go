// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package censys_definitions

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CommandReply_CommandStatus int32

const (
	CommandReply_FATAL     CommandReply_CommandStatus = 0
	CommandReply_SUCCESS   CommandReply_CommandStatus = 1
	CommandReply_ERROR     CommandReply_CommandStatus = 2
	CommandReply_NO_RECORD CommandReply_CommandStatus = 3
)

var CommandReply_CommandStatus_name = map[int32]string{
	0: "FATAL",
	1: "SUCCESS",
	2: "ERROR",
	3: "NO_RECORD",
}
var CommandReply_CommandStatus_value = map[string]int32{
	"FATAL":     0,
	"SUCCESS":   1,
	"ERROR":     2,
	"NO_RECORD": 3,
}

func (x CommandReply_CommandStatus) String() string {
	return proto.EnumName(CommandReply_CommandStatus_name, int32(x))
}
func (CommandReply_CommandStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor8, []int{8, 0}
}

type HostQueryResponse_ResponseStatus int32

const (
	HostQueryResponse_RESERVED  HostQueryResponse_ResponseStatus = 0
	HostQueryResponse_SUCCESS   HostQueryResponse_ResponseStatus = 1
	HostQueryResponse_NO_RECORD HostQueryResponse_ResponseStatus = 2
	HostQueryResponse_ERROR     HostQueryResponse_ResponseStatus = 3
)

var HostQueryResponse_ResponseStatus_name = map[int32]string{
	0: "RESERVED",
	1: "SUCCESS",
	2: "NO_RECORD",
	3: "ERROR",
}
var HostQueryResponse_ResponseStatus_value = map[string]int32{
	"RESERVED":  0,
	"SUCCESS":   1,
	"NO_RECORD": 2,
	"ERROR":     3,
}

func (x HostQueryResponse_ResponseStatus) String() string {
	return proto.EnumName(HostQueryResponse_ResponseStatus_name, int32(x))
}
func (HostQueryResponse_ResponseStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor8, []int{10, 0}
}

type AnonymousQueryResponse_ResponseStatus int32

const (
	AnonymousQueryResponse_RESERVED  AnonymousQueryResponse_ResponseStatus = 0
	AnonymousQueryResponse_SUCCESS   AnonymousQueryResponse_ResponseStatus = 1
	AnonymousQueryResponse_NO_RECORD AnonymousQueryResponse_ResponseStatus = 2
	AnonymousQueryResponse_ERROR     AnonymousQueryResponse_ResponseStatus = 3
)

var AnonymousQueryResponse_ResponseStatus_name = map[int32]string{
	0: "RESERVED",
	1: "SUCCESS",
	2: "NO_RECORD",
	3: "ERROR",
}
var AnonymousQueryResponse_ResponseStatus_value = map[string]int32{
	"RESERVED":  0,
	"SUCCESS":   1,
	"NO_RECORD": 2,
	"ERROR":     3,
}

func (x AnonymousQueryResponse_ResponseStatus) String() string {
	return proto.EnumName(AnonymousQueryResponse_ResponseStatus_name, int32(x))
}
func (AnonymousQueryResponse_ResponseStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor8, []int{12, 0}
}

type RootStoreQuery_RootStoreQueryType int32

const (
	RootStoreQuery_RESERVED  RootStoreQuery_RootStoreQueryType = 0
	RootStoreQuery_MOZILLA   RootStoreQuery_RootStoreQueryType = 1
	RootStoreQuery_MICROSOFT RootStoreQuery_RootStoreQueryType = 2
	RootStoreQuery_APPLE     RootStoreQuery_RootStoreQueryType = 3
	RootStoreQuery_ALL       RootStoreQuery_RootStoreQueryType = 4
)

var RootStoreQuery_RootStoreQueryType_name = map[int32]string{
	0: "RESERVED",
	1: "MOZILLA",
	2: "MICROSOFT",
	3: "APPLE",
	4: "ALL",
}
var RootStoreQuery_RootStoreQueryType_value = map[string]int32{
	"RESERVED":  0,
	"MOZILLA":   1,
	"MICROSOFT": 2,
	"APPLE":     3,
	"ALL":       4,
}

func (x RootStoreQuery_RootStoreQueryType) String() string {
	return proto.EnumName(RootStoreQuery_RootStoreQueryType_name, int32(x))
}
func (RootStoreQuery_RootStoreQueryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor8, []int{14, 0}
}

type MinScanId struct {
	Key       *AnonymousKey `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	MinScanId uint32        `protobuf:"varint,2,opt,name=min_scan_id,json=minScanId" json:"min_scan_id,omitempty"`
}

func (m *MinScanId) Reset()                    { *m = MinScanId{} }
func (m *MinScanId) String() string            { return proto.CompactTextString(m) }
func (*MinScanId) ProtoMessage()               {}
func (*MinScanId) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *MinScanId) GetKey() *AnonymousKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MinScanId) GetMinScanId() uint32 {
	if m != nil {
		return m.MinScanId
	}
	return 0
}

type MozillaOneCRLEntry struct {
	Issuer       []byte `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Serial       []byte `protobuf:"bytes,2,opt,name=serial,proto3" json:"serial,omitempty"`
	Id           string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	LastModified uint32 `protobuf:"varint,4,opt,name=last_modified,json=lastModified" json:"last_modified,omitempty"`
}

func (m *MozillaOneCRLEntry) Reset()                    { *m = MozillaOneCRLEntry{} }
func (m *MozillaOneCRLEntry) String() string            { return proto.CompactTextString(m) }
func (*MozillaOneCRLEntry) ProtoMessage()               {}
func (*MozillaOneCRLEntry) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *MozillaOneCRLEntry) GetIssuer() []byte {
	if m != nil {
		return m.Issuer
	}
	return nil
}

func (m *MozillaOneCRLEntry) GetSerial() []byte {
	if m != nil {
		return m.Serial
	}
	return nil
}

func (m *MozillaOneCRLEntry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MozillaOneCRLEntry) GetLastModified() uint32 {
	if m != nil {
		return m.LastModified
	}
	return 0
}

type Command struct {
	IncrementalDump bool                  `protobuf:"varint,1,opt,name=incremental_dump,json=incrementalDump" json:"incremental_dump,omitempty"`
	MinScanIds      []*MinScanId          `protobuf:"bytes,2,rep,name=min_scan_ids,json=minScanIds" json:"min_scan_ids,omitempty"`
	Filepath        string                `protobuf:"bytes,3,opt,name=filepath" json:"filepath,omitempty"`
	MaxRecords      uint32                `protobuf:"varint,4,opt,name=max_records,json=maxRecords" json:"max_records,omitempty"`
	StartIp         uint32                `protobuf:"varint,5,opt,name=start_ip,json=startIp" json:"start_ip,omitempty"`
	StopIp          uint32                `protobuf:"varint,6,opt,name=stop_ip,json=stopIp" json:"stop_ip,omitempty"`
	OneCrlEntries   []*MozillaOneCRLEntry `protobuf:"bytes,7,rep,name=one_crl_entries,json=oneCrlEntries" json:"one_crl_entries,omitempty"`
	Threads         uint32                `protobuf:"varint,8,opt,name=threads" json:"threads,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *Command) GetIncrementalDump() bool {
	if m != nil {
		return m.IncrementalDump
	}
	return false
}

func (m *Command) GetMinScanIds() []*MinScanId {
	if m != nil {
		return m.MinScanIds
	}
	return nil
}

func (m *Command) GetFilepath() string {
	if m != nil {
		return m.Filepath
	}
	return ""
}

func (m *Command) GetMaxRecords() uint32 {
	if m != nil {
		return m.MaxRecords
	}
	return 0
}

func (m *Command) GetStartIp() uint32 {
	if m != nil {
		return m.StartIp
	}
	return 0
}

func (m *Command) GetStopIp() uint32 {
	if m != nil {
		return m.StopIp
	}
	return 0
}

func (m *Command) GetOneCrlEntries() []*MozillaOneCRLEntry {
	if m != nil {
		return m.OneCrlEntries
	}
	return nil
}

func (m *Command) GetThreads() uint32 {
	if m != nil {
		return m.Threads
	}
	return 0
}

type AnonymousStoreStatistics struct {
	TotalRecords              uint64 `protobuf:"varint,1,opt,name=total_records,json=totalRecords" json:"total_records,omitempty"`
	RecordsAddedLastReset     uint64 `protobuf:"varint,2,opt,name=records_added_last_reset,json=recordsAddedLastReset" json:"records_added_last_reset,omitempty"`
	RecordsUpdatedLastReset   uint64 `protobuf:"varint,3,opt,name=records_updated_last_reset,json=recordsUpdatedLastReset" json:"records_updated_last_reset,omitempty"`
	RecordsUnchangedLastReset uint64 `protobuf:"varint,4,opt,name=records_unchanged_last_reset,json=recordsUnchangedLastReset" json:"records_unchanged_last_reset,omitempty"`
	RecordsReceived           uint64 `protobuf:"varint,5,opt,name=records_received,json=recordsReceived" json:"records_received,omitempty"`
	RecordsInRedisQueue       uint32 `protobuf:"varint,6,opt,name=records_in_redis_queue,json=recordsInRedisQueue" json:"records_in_redis_queue,omitempty"`
	RedisQueueName            string `protobuf:"bytes,7,opt,name=redis_queue_name,json=redisQueueName" json:"redis_queue_name,omitempty"`
	QueueType                 string `protobuf:"bytes,8,opt,name=queue_type,json=queueType" json:"queue_type,omitempty"`
	WorkerThreads             uint32 `protobuf:"varint,9,opt,name=worker_threads,json=workerThreads" json:"worker_threads,omitempty"`
}

func (m *AnonymousStoreStatistics) Reset()                    { *m = AnonymousStoreStatistics{} }
func (m *AnonymousStoreStatistics) String() string            { return proto.CompactTextString(m) }
func (*AnonymousStoreStatistics) ProtoMessage()               {}
func (*AnonymousStoreStatistics) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *AnonymousStoreStatistics) GetTotalRecords() uint64 {
	if m != nil {
		return m.TotalRecords
	}
	return 0
}

func (m *AnonymousStoreStatistics) GetRecordsAddedLastReset() uint64 {
	if m != nil {
		return m.RecordsAddedLastReset
	}
	return 0
}

func (m *AnonymousStoreStatistics) GetRecordsUpdatedLastReset() uint64 {
	if m != nil {
		return m.RecordsUpdatedLastReset
	}
	return 0
}

func (m *AnonymousStoreStatistics) GetRecordsUnchangedLastReset() uint64 {
	if m != nil {
		return m.RecordsUnchangedLastReset
	}
	return 0
}

func (m *AnonymousStoreStatistics) GetRecordsReceived() uint64 {
	if m != nil {
		return m.RecordsReceived
	}
	return 0
}

func (m *AnonymousStoreStatistics) GetRecordsInRedisQueue() uint32 {
	if m != nil {
		return m.RecordsInRedisQueue
	}
	return 0
}

func (m *AnonymousStoreStatistics) GetRedisQueueName() string {
	if m != nil {
		return m.RedisQueueName
	}
	return ""
}

func (m *AnonymousStoreStatistics) GetQueueType() string {
	if m != nil {
		return m.QueueType
	}
	return ""
}

func (m *AnonymousStoreStatistics) GetWorkerThreads() uint32 {
	if m != nil {
		return m.WorkerThreads
	}
	return 0
}

type StatisticsPair struct {
	Port        uint32                    `protobuf:"varint,1,opt,name=port" json:"port,omitempty"`
	Protocol    uint32                    `protobuf:"varint,2,opt,name=protocol" json:"protocol,omitempty"`
	Subprotocol uint32                    `protobuf:"varint,3,opt,name=subprotocol" json:"subprotocol,omitempty"`
	Statistics  *AnonymousStoreStatistics `protobuf:"bytes,4,opt,name=statistics" json:"statistics,omitempty"`
}

func (m *StatisticsPair) Reset()                    { *m = StatisticsPair{} }
func (m *StatisticsPair) String() string            { return proto.CompactTextString(m) }
func (*StatisticsPair) ProtoMessage()               {}
func (*StatisticsPair) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *StatisticsPair) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *StatisticsPair) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *StatisticsPair) GetSubprotocol() uint32 {
	if m != nil {
		return m.Subprotocol
	}
	return 0
}

func (m *StatisticsPair) GetStatistics() *AnonymousStoreStatistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

type StoreStatistics struct {
	Global    *AnonymousStoreStatistics `protobuf:"bytes,1,opt,name=global" json:"global,omitempty"`
	Protocols []*StatisticsPair         `protobuf:"bytes,2,rep,name=protocols" json:"protocols,omitempty"`
}

func (m *StoreStatistics) Reset()                    { *m = StoreStatistics{} }
func (m *StoreStatistics) String() string            { return proto.CompactTextString(m) }
func (*StoreStatistics) ProtoMessage()               {}
func (*StoreStatistics) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{5} }

func (m *StoreStatistics) GetGlobal() *AnonymousStoreStatistics {
	if m != nil {
		return m.Global
	}
	return nil
}

func (m *StoreStatistics) GetProtocols() []*StatisticsPair {
	if m != nil {
		return m.Protocols
	}
	return nil
}

type ServerStatistics struct {
	StoreStatistics          map[string]*StoreStatistics          `protobuf:"bytes,1,rep,name=store_statistics,json=storeStatistics" json:"store_statistics,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	AnonymousStoreStatistics map[string]*AnonymousStoreStatistics `protobuf:"bytes,2,rep,name=anonymous_store_statistics,json=anonymousStoreStatistics" json:"anonymous_store_statistics,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ServerStatistics) Reset()                    { *m = ServerStatistics{} }
func (m *ServerStatistics) String() string            { return proto.CompactTextString(m) }
func (*ServerStatistics) ProtoMessage()               {}
func (*ServerStatistics) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{6} }

func (m *ServerStatistics) GetStoreStatistics() map[string]*StoreStatistics {
	if m != nil {
		return m.StoreStatistics
	}
	return nil
}

func (m *ServerStatistics) GetAnonymousStoreStatistics() map[string]*AnonymousStoreStatistics {
	if m != nil {
		return m.AnonymousStoreStatistics
	}
	return nil
}

type PruneStatistics struct {
	Key           *AnonymousKey `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	RecordsPruned uint64        `protobuf:"varint,2,opt,name=records_pruned,json=recordsPruned" json:"records_pruned,omitempty"`
}

func (m *PruneStatistics) Reset()                    { *m = PruneStatistics{} }
func (m *PruneStatistics) String() string            { return proto.CompactTextString(m) }
func (*PruneStatistics) ProtoMessage()               {}
func (*PruneStatistics) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{7} }

func (m *PruneStatistics) GetKey() *AnonymousKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PruneStatistics) GetRecordsPruned() uint64 {
	if m != nil {
		return m.RecordsPruned
	}
	return 0
}

type CommandReply struct {
	Status          CommandReply_CommandStatus `protobuf:"varint,1,opt,name=status,enum=zsearch.CommandReply_CommandStatus" json:"status,omitempty"`
	Error           string                     `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Stats           *ServerStatistics          `protobuf:"bytes,3,opt,name=stats" json:"stats,omitempty"`
	PruneStatistics []*PruneStatistics         `protobuf:"bytes,4,rep,name=prune_statistics,json=pruneStatistics" json:"prune_statistics,omitempty"`
}

func (m *CommandReply) Reset()                    { *m = CommandReply{} }
func (m *CommandReply) String() string            { return proto.CompactTextString(m) }
func (*CommandReply) ProtoMessage()               {}
func (*CommandReply) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{8} }

func (m *CommandReply) GetStatus() CommandReply_CommandStatus {
	if m != nil {
		return m.Status
	}
	return CommandReply_FATAL
}

func (m *CommandReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CommandReply) GetStats() *ServerStatistics {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *CommandReply) GetPruneStatistics() []*PruneStatistics {
	if m != nil {
		return m.PruneStatistics
	}
	return nil
}

type HostQuery struct {
	Ip          uint32 `protobuf:"fixed32,1,opt,name=ip" json:"ip,omitempty"`
	Domain      string `protobuf:"bytes,2,opt,name=domain" json:"domain,omitempty"`
	Port        uint32 `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	Protocol    uint32 `protobuf:"varint,4,opt,name=protocol" json:"protocol,omitempty"`
	Subprotocol uint32 `protobuf:"varint,5,opt,name=subprotocol" json:"subprotocol,omitempty"`
	MaxRecords  uint32 `protobuf:"varint,6,opt,name=max_records,json=maxRecords" json:"max_records,omitempty"`
}

func (m *HostQuery) Reset()                    { *m = HostQuery{} }
func (m *HostQuery) String() string            { return proto.CompactTextString(m) }
func (*HostQuery) ProtoMessage()               {}
func (*HostQuery) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{9} }

func (m *HostQuery) GetIp() uint32 {
	if m != nil {
		return m.Ip
	}
	return 0
}

func (m *HostQuery) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *HostQuery) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *HostQuery) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *HostQuery) GetSubprotocol() uint32 {
	if m != nil {
		return m.Subprotocol
	}
	return 0
}

func (m *HostQuery) GetMaxRecords() uint32 {
	if m != nil {
		return m.MaxRecords
	}
	return 0
}

type HostQueryResponse struct {
	Status      HostQueryResponse_ResponseStatus `protobuf:"varint,1,opt,name=status,enum=zsearch.HostQueryResponse_ResponseStatus" json:"status,omitempty"`
	Ip          uint32                           `protobuf:"fixed32,2,opt,name=ip" json:"ip,omitempty"`
	Domain      string                           `protobuf:"bytes,3,opt,name=domain" json:"domain,omitempty"`
	Port        uint32                           `protobuf:"varint,4,opt,name=port" json:"port,omitempty"`
	Protocol    uint32                           `protobuf:"varint,5,opt,name=protocol" json:"protocol,omitempty"`
	Subprotocol uint32                           `protobuf:"varint,6,opt,name=subprotocol" json:"subprotocol,omitempty"`
	Record      *Record                          `protobuf:"bytes,7,opt,name=record" json:"record,omitempty"`
	Records     []*Record                        `protobuf:"bytes,8,rep,name=records" json:"records,omitempty"`
	Error       string                           `protobuf:"bytes,9,opt,name=error" json:"error,omitempty"`
}

func (m *HostQueryResponse) Reset()                    { *m = HostQueryResponse{} }
func (m *HostQueryResponse) String() string            { return proto.CompactTextString(m) }
func (*HostQueryResponse) ProtoMessage()               {}
func (*HostQueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{10} }

func (m *HostQueryResponse) GetStatus() HostQueryResponse_ResponseStatus {
	if m != nil {
		return m.Status
	}
	return HostQueryResponse_RESERVED
}

func (m *HostQueryResponse) GetIp() uint32 {
	if m != nil {
		return m.Ip
	}
	return 0
}

func (m *HostQueryResponse) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *HostQueryResponse) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *HostQueryResponse) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *HostQueryResponse) GetSubprotocol() uint32 {
	if m != nil {
		return m.Subprotocol
	}
	return 0
}

func (m *HostQueryResponse) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *HostQueryResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *HostQueryResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type AnonymousQuery struct {
	Sha256Fp []byte `protobuf:"bytes,1,opt,name=sha256fp,proto3" json:"sha256fp,omitempty"`
}

func (m *AnonymousQuery) Reset()                    { *m = AnonymousQuery{} }
func (m *AnonymousQuery) String() string            { return proto.CompactTextString(m) }
func (*AnonymousQuery) ProtoMessage()               {}
func (*AnonymousQuery) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{11} }

func (m *AnonymousQuery) GetSha256Fp() []byte {
	if m != nil {
		return m.Sha256Fp
	}
	return nil
}

type AnonymousQueryResponse struct {
	Sha256Fp []byte                                `protobuf:"bytes,1,opt,name=sha256fp,proto3" json:"sha256fp,omitempty"`
	Status   AnonymousQueryResponse_ResponseStatus `protobuf:"varint,2,opt,name=status,enum=zsearch.AnonymousQueryResponse_ResponseStatus" json:"status,omitempty"`
	Record   *AnonymousRecord                      `protobuf:"bytes,3,opt,name=record" json:"record,omitempty"`
	Records  []*AnonymousRecord                    `protobuf:"bytes,4,rep,name=records" json:"records,omitempty"`
	Error    string                                `protobuf:"bytes,59,opt,name=error" json:"error,omitempty"`
}

func (m *AnonymousQueryResponse) Reset()                    { *m = AnonymousQueryResponse{} }
func (m *AnonymousQueryResponse) String() string            { return proto.CompactTextString(m) }
func (*AnonymousQueryResponse) ProtoMessage()               {}
func (*AnonymousQueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{12} }

func (m *AnonymousQueryResponse) GetSha256Fp() []byte {
	if m != nil {
		return m.Sha256Fp
	}
	return nil
}

func (m *AnonymousQueryResponse) GetStatus() AnonymousQueryResponse_ResponseStatus {
	if m != nil {
		return m.Status
	}
	return AnonymousQueryResponse_RESERVED
}

func (m *AnonymousQueryResponse) GetRecord() *AnonymousRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *AnonymousQueryResponse) GetRecords() []*AnonymousRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *AnonymousQueryResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type UserDataRequest struct {
	Ip     uint32        `protobuf:"fixed32,1,opt,name=ip" json:"ip,omitempty"`
	Domain string        `protobuf:"bytes,2,opt,name=domain" json:"domain,omitempty"`
	Data   *UserdataAtom `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *UserDataRequest) Reset()                    { *m = UserDataRequest{} }
func (m *UserDataRequest) String() string            { return proto.CompactTextString(m) }
func (*UserDataRequest) ProtoMessage()               {}
func (*UserDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{13} }

func (m *UserDataRequest) GetIp() uint32 {
	if m != nil {
		return m.Ip
	}
	return 0
}

func (m *UserDataRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *UserDataRequest) GetData() *UserdataAtom {
	if m != nil {
		return m.Data
	}
	return nil
}

type RootStoreQuery struct {
	Type RootStoreQuery_RootStoreQueryType `protobuf:"varint,1,opt,name=type,enum=zsearch.RootStoreQuery_RootStoreQueryType" json:"type,omitempty"`
}

func (m *RootStoreQuery) Reset()                    { *m = RootStoreQuery{} }
func (m *RootStoreQuery) String() string            { return proto.CompactTextString(m) }
func (*RootStoreQuery) ProtoMessage()               {}
func (*RootStoreQuery) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{14} }

func (m *RootStoreQuery) GetType() RootStoreQuery_RootStoreQueryType {
	if m != nil {
		return m.Type
	}
	return RootStoreQuery_RESERVED
}

type RootStoreReply struct {
	Certificates []*Certificate `protobuf:"bytes,1,rep,name=certificates" json:"certificates,omitempty"`
}

func (m *RootStoreReply) Reset()                    { *m = RootStoreReply{} }
func (m *RootStoreReply) String() string            { return proto.CompactTextString(m) }
func (*RootStoreReply) ProtoMessage()               {}
func (*RootStoreReply) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{15} }

func (m *RootStoreReply) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

func init() {
	proto.RegisterType((*MinScanId)(nil), "zsearch.MinScanId")
	proto.RegisterType((*MozillaOneCRLEntry)(nil), "zsearch.MozillaOneCRLEntry")
	proto.RegisterType((*Command)(nil), "zsearch.Command")
	proto.RegisterType((*AnonymousStoreStatistics)(nil), "zsearch.AnonymousStoreStatistics")
	proto.RegisterType((*StatisticsPair)(nil), "zsearch.StatisticsPair")
	proto.RegisterType((*StoreStatistics)(nil), "zsearch.StoreStatistics")
	proto.RegisterType((*ServerStatistics)(nil), "zsearch.ServerStatistics")
	proto.RegisterType((*PruneStatistics)(nil), "zsearch.PruneStatistics")
	proto.RegisterType((*CommandReply)(nil), "zsearch.CommandReply")
	proto.RegisterType((*HostQuery)(nil), "zsearch.HostQuery")
	proto.RegisterType((*HostQueryResponse)(nil), "zsearch.HostQueryResponse")
	proto.RegisterType((*AnonymousQuery)(nil), "zsearch.AnonymousQuery")
	proto.RegisterType((*AnonymousQueryResponse)(nil), "zsearch.AnonymousQueryResponse")
	proto.RegisterType((*UserDataRequest)(nil), "zsearch.UserDataRequest")
	proto.RegisterType((*RootStoreQuery)(nil), "zsearch.RootStoreQuery")
	proto.RegisterType((*RootStoreReply)(nil), "zsearch.RootStoreReply")
	proto.RegisterEnum("zsearch.CommandReply_CommandStatus", CommandReply_CommandStatus_name, CommandReply_CommandStatus_value)
	proto.RegisterEnum("zsearch.HostQueryResponse_ResponseStatus", HostQueryResponse_ResponseStatus_name, HostQueryResponse_ResponseStatus_value)
	proto.RegisterEnum("zsearch.AnonymousQueryResponse_ResponseStatus", AnonymousQueryResponse_ResponseStatus_name, AnonymousQueryResponse_ResponseStatus_value)
	proto.RegisterEnum("zsearch.RootStoreQuery_RootStoreQueryType", RootStoreQuery_RootStoreQueryType_name, RootStoreQuery_RootStoreQueryType_value)
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 1339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0x5b, 0x8f, 0xdb, 0xc4,
	0x17, 0xff, 0x27, 0xce, 0x26, 0x9b, 0x93, 0x6b, 0xe7, 0xdf, 0x8b, 0x1b, 0x28, 0x2c, 0xae, 0xaa,
	0xee, 0x22, 0x14, 0x50, 0x4a, 0x69, 0xa1, 0x12, 0x28, 0xca, 0xa6, 0x62, 0x21, 0xdb, 0x6c, 0x27,
	0x5b, 0x24, 0x10, 0x92, 0x35, 0xb5, 0x67, 0xbb, 0x56, 0x63, 0x8f, 0x3b, 0x33, 0x2e, 0x4d, 0xdf,
	0x10, 0xdf, 0x82, 0x67, 0xe0, 0x03, 0xf0, 0xc8, 0x0b, 0x6f, 0x7c, 0x2e, 0xe4, 0xf1, 0xd8, 0x8e,
	0x9d, 0xdd, 0xd5, 0xbe, 0xf0, 0x96, 0x73, 0x3f, 0xe7, 0x77, 0x2e, 0xe3, 0x40, 0x93, 0x87, 0xce,
	0x30, 0xe4, 0x4c, 0x32, 0xd4, 0x78, 0x2b, 0x28, 0xe1, 0xce, 0xe9, 0xa0, 0xed, 0x30, 0xdf, 0x67,
	0x41, 0xc2, 0x1e, 0x5c, 0x71, 0x28, 0x97, 0xde, 0x89, 0xe7, 0x10, 0x49, 0x35, 0xab, 0x47, 0x02,
	0x16, 0x08, 0xc9, 0x78, 0xc6, 0x38, 0x65, 0x42, 0xae, 0x31, 0xac, 0x63, 0x68, 0x1e, 0x7a, 0xc1,
	0xc2, 0x21, 0xc1, 0x81, 0x8b, 0xee, 0x82, 0xf1, 0x92, 0xae, 0xcc, 0xca, 0x4e, 0x65, 0xb7, 0x35,
	0xba, 0x36, 0xd4, 0x61, 0x86, 0xe3, 0x80, 0x05, 0x2b, 0x9f, 0x45, 0xe2, 0x5b, 0xba, 0xc2, 0xb1,
	0x06, 0x7a, 0x0f, 0x5a, 0xbe, 0x17, 0xd8, 0xc2, 0x21, 0x81, 0xed, 0xb9, 0x66, 0x75, 0xa7, 0xb2,
	0xdb, 0xc1, 0x4d, 0x3f, 0x75, 0x64, 0xad, 0x00, 0x1d, 0xb2, 0xb7, 0xde, 0x72, 0x49, 0xe6, 0x01,
	0x9d, 0xe0, 0xd9, 0x34, 0x90, 0x7c, 0x85, 0xae, 0x43, 0xdd, 0x13, 0x22, 0xa2, 0x5c, 0x45, 0x68,
	0x63, 0x4d, 0xc5, 0x7c, 0x41, 0xb9, 0x47, 0x96, 0xca, 0x51, 0x1b, 0x6b, 0x0a, 0x75, 0xa1, 0xea,
	0xb9, 0xa6, 0xb1, 0x53, 0xd9, 0x6d, 0xe2, 0xaa, 0xe7, 0xa2, 0xdb, 0xd0, 0x59, 0x12, 0x21, 0x6d,
	0x9f, 0xb9, 0xde, 0x89, 0x47, 0x5d, 0xb3, 0xa6, 0xe2, 0xb6, 0x63, 0xe6, 0xa1, 0xe6, 0x59, 0x7f,
	0x55, 0xa1, 0x31, 0x61, 0xbe, 0x4f, 0x02, 0x17, 0xed, 0x41, 0xdf, 0x0b, 0x1c, 0x4e, 0x7d, 0x1a,
	0x48, 0xb2, 0xb4, 0xdd, 0xc8, 0x0f, 0x55, 0xe8, 0x6d, 0xdc, 0x5b, 0xe3, 0xef, 0x47, 0x7e, 0x88,
	0x3e, 0x85, 0xf6, 0x5a, 0x45, 0xc2, 0xac, 0xee, 0x18, 0xbb, 0xad, 0x11, 0xca, 0x30, 0xc8, 0x40,
	0xc2, 0x90, 0x95, 0x29, 0xd0, 0x00, 0xb6, 0x4f, 0xbc, 0x25, 0x0d, 0x89, 0x3c, 0xd5, 0x79, 0x66,
	0x34, 0x7a, 0x1f, 0x5a, 0x3e, 0x79, 0x63, 0x73, 0xea, 0x30, 0xee, 0x0a, 0x9d, 0x2b, 0xf8, 0xe4,
	0x0d, 0x4e, 0x38, 0xe8, 0x26, 0x6c, 0x0b, 0x49, 0xb8, 0xb4, 0xbd, 0xd0, 0xdc, 0x52, 0xd2, 0x86,
	0xa2, 0x0f, 0x42, 0x74, 0x03, 0x1a, 0x42, 0xb2, 0x30, 0x96, 0xd4, 0x95, 0xa4, 0x1e, 0x93, 0x07,
	0x21, 0x9a, 0x40, 0x8f, 0x05, 0xd4, 0x76, 0xf8, 0xd2, 0xa6, 0x81, 0xe4, 0x1e, 0x15, 0x66, 0x43,
	0x65, 0xfa, 0x4e, 0x9e, 0xe9, 0x06, 0xf0, 0xb8, 0xc3, 0x02, 0x3a, 0xe1, 0xcb, 0x69, 0x62, 0x81,
	0x4c, 0x68, 0xc8, 0x53, 0x4e, 0x89, 0x2b, 0xcc, 0xed, 0x24, 0xae, 0x26, 0xad, 0x7f, 0x0c, 0x30,
	0xb3, 0x6e, 0x2f, 0xe2, 0x31, 0x59, 0x48, 0x22, 0x3d, 0x21, 0x3d, 0x47, 0xc4, 0xf0, 0x4b, 0x16,
	0xe3, 0x98, 0x96, 0x14, 0x43, 0x59, 0xc3, 0x6d, 0xc5, 0x4c, 0x8b, 0x7a, 0x00, 0xa6, 0x16, 0xdb,
	0xc4, 0x75, 0xa9, 0x6b, 0xab, 0x8e, 0x71, 0x2a, 0xa8, 0x54, 0xdd, 0xad, 0xe1, 0x6b, 0x5a, 0x3e,
	0x8e, 0xc5, 0x33, 0x22, 0x24, 0x8e, 0x85, 0xe8, 0x11, 0x0c, 0x52, 0xc3, 0x28, 0x74, 0x89, 0x2c,
	0x9a, 0x1a, 0xca, 0xf4, 0x86, 0xd6, 0x78, 0x96, 0x28, 0xe4, 0xc6, 0x5f, 0xc1, 0xbb, 0x99, 0x71,
	0xe0, 0x9c, 0x92, 0xe0, 0x45, 0xd1, 0xbc, 0xa6, 0xcc, 0x6f, 0xa6, 0xe6, 0xa9, 0x4a, 0xee, 0x60,
	0x0f, 0xfa, 0xa9, 0x03, 0x4e, 0x1d, 0xea, 0xbd, 0xa6, 0xae, 0xea, 0x49, 0x0d, 0xf7, 0x34, 0x1f,
	0x6b, 0x36, 0xba, 0x07, 0xd7, 0x53, 0x55, 0x2f, 0xb0, 0x39, 0x75, 0x3d, 0x61, 0xbf, 0x8a, 0x68,
	0x44, 0x75, 0xab, 0xfe, 0xaf, 0xa5, 0x07, 0x01, 0x8e, 0x65, 0x4f, 0x63, 0x11, 0xda, 0x8d, 0xfd,
	0x67, 0x9a, 0x76, 0x40, 0x7c, 0x6a, 0x36, 0xd4, 0xc0, 0x74, 0x79, 0xa6, 0xf5, 0x84, 0xf8, 0x14,
	0xdd, 0x02, 0x48, 0x74, 0xe4, 0x2a, 0xa4, 0xaa, 0x3f, 0x4d, 0xdc, 0x54, 0x9c, 0xe3, 0x55, 0x48,
	0xd1, 0x1d, 0xe8, 0xfe, 0xc4, 0xf8, 0x4b, 0xca, 0xed, 0xb4, 0x85, 0x4d, 0x15, 0xb5, 0x93, 0x70,
	0x8f, 0x75, 0x23, 0x7f, 0xaf, 0x40, 0x37, 0x6f, 0xdd, 0x11, 0xf1, 0x38, 0x42, 0x50, 0x0b, 0x19,
	0x97, 0xaa, 0x6b, 0x1d, 0xac, 0x7e, 0xc7, 0xf3, 0xab, 0xce, 0x80, 0xc3, 0x96, 0x7a, 0x89, 0x33,
	0x1a, 0xed, 0x40, 0x4b, 0x44, 0xcf, 0x33, 0xb1, 0xa1, 0xc4, 0xeb, 0x2c, 0x34, 0x06, 0x10, 0x59,
	0x0c, 0x85, 0x71, 0x6b, 0xf4, 0xc1, 0xe6, 0xd5, 0x28, 0xcd, 0x11, 0x5e, 0x33, 0xb2, 0x7e, 0xa9,
	0x40, 0xaf, 0x3c, 0x67, 0x9f, 0x43, 0xfd, 0xc5, 0x92, 0x3d, 0x27, 0x4b, 0x7d, 0x88, 0x2e, 0xe1,
	0x52, 0x1b, 0xa0, 0xfb, 0xd0, 0x4c, 0xb3, 0x4b, 0x57, 0xf8, 0x46, 0x66, 0x5d, 0xc4, 0x03, 0xe7,
	0x9a, 0xd6, 0x9f, 0x06, 0xf4, 0x17, 0x94, 0xbf, 0xa6, 0x7c, 0x2d, 0x8d, 0xef, 0xa1, 0xaf, 0x0e,
	0xa5, 0xbd, 0x56, 0x63, 0x45, 0xb9, 0x1c, 0xe6, 0x2e, 0x4b, 0x46, 0xc3, 0x52, 0x62, 0xc9, 0xfa,
	0xf5, 0x44, 0xa9, 0xc2, 0x08, 0x06, 0x24, 0x2d, 0xc5, 0xde, 0x08, 0x92, 0xe4, 0xfd, 0xe0, 0xfc,
	0x20, 0xe7, 0xc1, 0x90, 0x44, 0x33, 0xc9, 0x39, 0xe2, 0xc1, 0x8f, 0x70, 0xf5, 0x2c, 0x0b, 0xd4,
	0xcf, 0xcf, 0x7e, 0x33, 0xb9, 0xef, 0x43, 0xd8, 0x7a, 0x4d, 0x96, 0x11, 0x55, 0x43, 0xd1, 0x1a,
	0x99, 0x6b, 0x18, 0x16, 0x81, 0x4f, 0xd4, 0xbe, 0xa8, 0x3e, 0xac, 0x0c, 0x02, 0xb8, 0x75, 0x61,
	0x62, 0x67, 0x84, 0x79, 0x50, 0x0c, 0x73, 0x89, 0x46, 0xe7, 0xf1, 0x2c, 0x02, 0xbd, 0x23, 0x1e,
	0x05, 0xeb, 0xb8, 0x5e, 0xfa, 0xfd, 0xba, 0x03, 0xdd, 0x74, 0x87, 0xc3, 0xd8, 0x87, 0xab, 0x6f,
	0x53, 0x47, 0x73, 0x95, 0x63, 0xd7, 0xfa, 0xb5, 0x0a, 0x6d, 0xfd, 0x96, 0x60, 0x1a, 0x2e, 0x57,
	0xe8, 0x11, 0xd4, 0xe3, 0x46, 0x45, 0xc9, 0xed, 0xeb, 0x8e, 0x6e, 0x67, 0x31, 0xd6, 0xd5, 0x52,
	0x62, 0xa1, 0x54, 0xb1, 0x36, 0x41, 0x57, 0x61, 0x8b, 0x72, 0xce, 0xb8, 0x8a, 0xd5, 0xc4, 0x09,
	0x81, 0x3e, 0x86, 0xad, 0x58, 0x2e, 0xd4, 0x82, 0xb5, 0x46, 0x37, 0xcf, 0x6d, 0x3b, 0x4e, 0xf4,
	0xd0, 0x04, 0xfa, 0x2a, 0x67, 0xbb, 0xb0, 0x7b, 0x46, 0xa1, 0x4d, 0x25, 0x60, 0x70, 0x2f, 0x2c,
	0x32, 0xac, 0x31, 0x74, 0x0a, 0x49, 0xa2, 0x26, 0x6c, 0x3d, 0x1e, 0x1f, 0x8f, 0x67, 0xfd, 0xff,
	0xa1, 0x16, 0x34, 0x16, 0xcf, 0x26, 0x93, 0xe9, 0x62, 0xd1, 0xaf, 0xc4, 0xfc, 0x29, 0xc6, 0x73,
	0xdc, 0xaf, 0xa2, 0x0e, 0x34, 0x9f, 0xcc, 0x6d, 0x3c, 0x9d, 0xcc, 0xf1, 0x7e, 0xdf, 0xb0, 0xfe,
	0xa8, 0x40, 0xf3, 0x6b, 0x26, 0xe4, 0xd3, 0x88, 0xf2, 0x95, 0x7a, 0xab, 0x93, 0xc7, 0xb5, 0x81,
	0xab, 0x5e, 0x18, 0xbf, 0xe9, 0x2e, 0xf3, 0x89, 0x17, 0xe8, 0x6a, 0x35, 0x95, 0x5d, 0x21, 0xe3,
	0x9c, 0x2b, 0x54, 0xbb, 0xf8, 0x0a, 0x6d, 0x6d, 0x5e, 0xa1, 0xd2, 0x3b, 0x5b, 0x2f, 0xbf, 0xb3,
	0xd6, 0xcf, 0x06, 0x5c, 0xc9, 0x12, 0xc5, 0x54, 0x84, 0x2c, 0x10, 0x14, 0x8d, 0x4b, 0xad, 0xdc,
	0xcb, 0xc0, 0xdb, 0xd0, 0x1d, 0xa6, 0x3f, 0x4a, 0x0d, 0x4d, 0x6a, 0xae, 0x9e, 0x51, 0xb3, 0x71,
	0x66, 0xcd, 0xb5, 0x73, 0x6a, 0xde, 0xba, 0xb8, 0xe6, 0xfa, 0x66, 0xcd, 0x77, 0xa1, 0x9e, 0xd4,
	0xab, 0x1e, 0x91, 0xd6, 0xa8, 0x97, 0x25, 0x9f, 0x14, 0x8d, 0xb5, 0x18, 0xed, 0x41, 0x23, 0x05,
	0x66, 0x5b, 0xcd, 0xc8, 0x86, 0x66, 0x2a, 0xcf, 0xc7, 0xb3, 0xb9, 0x36, 0x9e, 0xd6, 0x14, 0xba,
	0xc5, 0xea, 0x51, 0x1b, 0xb6, 0xf1, 0x74, 0x31, 0xc5, 0xdf, 0x4d, 0xf7, 0xcb, 0xc3, 0x52, 0x98,
	0x90, 0x6a, 0x3e, 0x3b, 0x86, 0xf5, 0x11, 0x74, 0xb3, 0x2d, 0x4c, 0x06, 0x66, 0x00, 0xdb, 0xe2,
	0x94, 0x8c, 0xee, 0x7f, 0x76, 0x12, 0xea, 0xcf, 0xc1, 0x8c, 0xb6, 0xfe, 0xae, 0xc2, 0xf5, 0xa2,
	0x7a, 0xd6, 0xb6, 0x0b, 0xcc, 0xd0, 0xe3, 0xac, 0xa5, 0x55, 0xd5, 0xd2, 0xe1, 0xe6, 0x05, 0xb8,
	0x54, 0x5f, 0x3f, 0xc9, 0xd0, 0x35, 0x4a, 0xe7, 0x2f, 0xf3, 0x53, 0x82, 0x79, 0x94, 0xc3, 0x5c,
	0x5e, 0xc5, 0xb2, 0xc9, 0x26, 0xde, 0x8f, 0xfe, 0x03, 0xbc, 0x5d, 0xe8, 0x3d, 0x13, 0x94, 0xef,
	0x13, 0x49, 0x30, 0x7d, 0x15, 0x51, 0x21, 0x2f, 0xbd, 0xa1, 0x7b, 0x50, 0x73, 0x89, 0x24, 0xba,
	0xf6, 0xfc, 0x8a, 0xc6, 0xfe, 0x62, 0xc1, 0x58, 0x32, 0x1f, 0x2b, 0x15, 0xeb, 0xb7, 0x0a, 0x74,
	0x31, 0x63, 0x52, 0x5d, 0xe9, 0xa4, 0xad, 0x5f, 0x42, 0x4d, 0x7d, 0xb8, 0x24, 0x4b, 0xf5, 0x61,
	0x3e, 0x6d, 0x05, 0xb5, 0x12, 0x19, 0x7f, 0xd9, 0x60, 0x65, 0x67, 0x61, 0x40, 0x9b, 0xb2, 0x4d,
	0x0c, 0x0e, 0xe7, 0x3f, 0x1c, 0xcc, 0x66, 0xe3, 0x04, 0x83, 0xc3, 0x83, 0x09, 0x9e, 0x2f, 0xe6,
	0x8f, 0x8f, 0x13, 0x0c, 0xc6, 0x47, 0x47, 0xb3, 0x69, 0xdf, 0x40, 0x0d, 0x30, 0xc6, 0xb3, 0x59,
	0xbf, 0x66, 0x7d, 0xb3, 0x96, 0x65, 0x72, 0xc7, 0x1f, 0x42, 0x7b, 0xed, 0xcf, 0x52, 0xfa, 0xae,
	0x5f, 0xcd, 0xaf, 0x79, 0x2e, 0xc4, 0x05, 0xcd, 0xe7, 0x75, 0xb5, 0x83, 0xf7, 0xfe, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0xdb, 0x56, 0xe0, 0xca, 0x8f, 0x0d, 0x00, 0x00,
}
