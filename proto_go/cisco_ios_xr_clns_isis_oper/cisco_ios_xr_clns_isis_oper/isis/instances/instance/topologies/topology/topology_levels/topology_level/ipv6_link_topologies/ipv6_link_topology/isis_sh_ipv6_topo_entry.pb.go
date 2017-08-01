// Code generated by protoc-gen-go.
// source: isis_sh_ipv6_topo_entry.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_clns_isis_oper_isis_instances_instance_topologies_topology_topology_levels_topology_level_ipv6_link_topologies_ipv6_link_topology is a generated protocol buffer package.

It is generated from these files:
	isis_sh_ipv6_topo_entry.proto

It has these top-level messages:
	IsisShIpv6TopoEntry_KEYS
	IsisShIpv6TopoEntry
	IsisNodeIdType
	IsisSnpaType
	IsisPerPriorityCounts
	IsisShRepEl
	IsisShIpv6FrrBackup
	IsisShIpv6Path
	IsisShTopoNeighbor
	IsisShIpv6TopoReachableDetails
	IsisShIpv6TopoReachableStatus
*/
package cisco_ios_xr_clns_isis_oper_isis_instances_instance_topologies_topology_topology_levels_topology_level_ipv6_link_topologies_ipv6_link_topology

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

// IPv6 IS Link Topology Entry
type IsisShIpv6TopoEntry_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	AfName       string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName      string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	TopologyName string `protobuf:"bytes,4,opt,name=topology_name,json=topologyName" json:"topology_name,omitempty"`
	Level        string `protobuf:"bytes,5,opt,name=level" json:"level,omitempty"`
	SystemId     string `protobuf:"bytes,6,opt,name=system_id,json=systemId" json:"system_id,omitempty"`
}

func (m *IsisShIpv6TopoEntry_KEYS) Reset()                    { *m = IsisShIpv6TopoEntry_KEYS{} }
func (m *IsisShIpv6TopoEntry_KEYS) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv6TopoEntry_KEYS) ProtoMessage()               {}
func (*IsisShIpv6TopoEntry_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IsisShIpv6TopoEntry_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShIpv6TopoEntry_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IsisShIpv6TopoEntry_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *IsisShIpv6TopoEntry_KEYS) GetTopologyName() string {
	if m != nil {
		return m.TopologyName
	}
	return ""
}

func (m *IsisShIpv6TopoEntry_KEYS) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *IsisShIpv6TopoEntry_KEYS) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

type IsisShIpv6TopoEntry struct {
	// Source Address
	SourceAddress string `protobuf:"bytes,50,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	// Does the IS participate in the topology?
	IsParticipant bool `protobuf:"varint,51,opt,name=is_participant,json=isParticipant" json:"is_participant,omitempty"`
	// Is the IS overloaded?
	IsOverloaded bool `protobuf:"varint,52,opt,name=is_overloaded,json=isOverloaded" json:"is_overloaded,omitempty"`
	// Is the IS attached?
	IsAttached bool `protobuf:"varint,53,opt,name=is_attached,json=isAttached" json:"is_attached,omitempty"`
	// Is the IS reachable, and, if so, its status within the SPT
	ReachabilityStatus *IsisShIpv6TopoReachableStatus `protobuf:"bytes,54,opt,name=reachability_status,json=reachabilityStatus" json:"reachability_status,omitempty"`
	// Per-priority counts of prefix items advertised by the IS
	AdvertisedPrefixItemCounts *IsisPerPriorityCounts `protobuf:"bytes,55,opt,name=advertised_prefix_item_counts,json=advertisedPrefixItemCounts" json:"advertised_prefix_item_counts,omitempty"`
}

func (m *IsisShIpv6TopoEntry) Reset()                    { *m = IsisShIpv6TopoEntry{} }
func (m *IsisShIpv6TopoEntry) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv6TopoEntry) ProtoMessage()               {}
func (*IsisShIpv6TopoEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IsisShIpv6TopoEntry) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *IsisShIpv6TopoEntry) GetIsParticipant() bool {
	if m != nil {
		return m.IsParticipant
	}
	return false
}

func (m *IsisShIpv6TopoEntry) GetIsOverloaded() bool {
	if m != nil {
		return m.IsOverloaded
	}
	return false
}

func (m *IsisShIpv6TopoEntry) GetIsAttached() bool {
	if m != nil {
		return m.IsAttached
	}
	return false
}

func (m *IsisShIpv6TopoEntry) GetReachabilityStatus() *IsisShIpv6TopoReachableStatus {
	if m != nil {
		return m.ReachabilityStatus
	}
	return nil
}

func (m *IsisShIpv6TopoEntry) GetAdvertisedPrefixItemCounts() *IsisPerPriorityCounts {
	if m != nil {
		return m.AdvertisedPrefixItemCounts
	}
	return nil
}

type IsisNodeIdType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IsisNodeIdType) Reset()                    { *m = IsisNodeIdType{} }
func (m *IsisNodeIdType) String() string            { return proto.CompactTextString(m) }
func (*IsisNodeIdType) ProtoMessage()               {}
func (*IsisNodeIdType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IsisNodeIdType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type IsisSnpaType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IsisSnpaType) Reset()                    { *m = IsisSnpaType{} }
func (m *IsisSnpaType) String() string            { return proto.CompactTextString(m) }
func (*IsisSnpaType) ProtoMessage()               {}
func (*IsisSnpaType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IsisSnpaType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Per-priority counts
type IsisPerPriorityCounts struct {
	// Critical priority
	Critical uint32 `protobuf:"varint,1,opt,name=critical" json:"critical,omitempty"`
	// High priority
	High uint32 `protobuf:"varint,2,opt,name=high" json:"high,omitempty"`
	// Medium priority
	Medium uint32 `protobuf:"varint,3,opt,name=medium" json:"medium,omitempty"`
	// Low priority
	Low uint32 `protobuf:"varint,4,opt,name=low" json:"low,omitempty"`
}

func (m *IsisPerPriorityCounts) Reset()                    { *m = IsisPerPriorityCounts{} }
func (m *IsisPerPriorityCounts) String() string            { return proto.CompactTextString(m) }
func (*IsisPerPriorityCounts) ProtoMessage()               {}
func (*IsisPerPriorityCounts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *IsisPerPriorityCounts) GetCritical() uint32 {
	if m != nil {
		return m.Critical
	}
	return 0
}

func (m *IsisPerPriorityCounts) GetHigh() uint32 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *IsisPerPriorityCounts) GetMedium() uint32 {
	if m != nil {
		return m.Medium
	}
	return 0
}

func (m *IsisPerPriorityCounts) GetLow() uint32 {
	if m != nil {
		return m.Low
	}
	return 0
}

// OSPF Repair Element
type IsisShRepEl struct {
	// RepairElementNodeID
	RepairElementNodeId string `protobuf:"bytes,1,opt,name=repair_element_node_id,json=repairElementNodeId" json:"repair_element_node_id,omitempty"`
	// RepairIPv4Addr
	RepairIpv4Addr string `protobuf:"bytes,2,opt,name=repair_ipv4_addr,json=repairIpv4Addr" json:"repair_ipv4_addr,omitempty"`
	// RepairIPv6Addr
	RepairIpv6Addr string `protobuf:"bytes,3,opt,name=repair_ipv6_addr,json=repairIpv6Addr" json:"repair_ipv6_addr,omitempty"`
	// Repair Label
	RepairLabel uint32 `protobuf:"varint,4,opt,name=repair_label,json=repairLabel" json:"repair_label,omitempty"`
	// Repair Element Type
	RepairElementType uint32 `protobuf:"varint,5,opt,name=repair_element_type,json=repairElementType" json:"repair_element_type,omitempty"`
	// Repair Strict SPF Label
	RepairStrictSpfLabel uint32 `protobuf:"varint,6,opt,name=repair_strict_spf_label,json=repairStrictSpfLabel" json:"repair_strict_spf_label,omitempty"`
}

func (m *IsisShRepEl) Reset()                    { *m = IsisShRepEl{} }
func (m *IsisShRepEl) String() string            { return proto.CompactTextString(m) }
func (*IsisShRepEl) ProtoMessage()               {}
func (*IsisShRepEl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IsisShRepEl) GetRepairElementNodeId() string {
	if m != nil {
		return m.RepairElementNodeId
	}
	return ""
}

func (m *IsisShRepEl) GetRepairIpv4Addr() string {
	if m != nil {
		return m.RepairIpv4Addr
	}
	return ""
}

func (m *IsisShRepEl) GetRepairIpv6Addr() string {
	if m != nil {
		return m.RepairIpv6Addr
	}
	return ""
}

func (m *IsisShRepEl) GetRepairLabel() uint32 {
	if m != nil {
		return m.RepairLabel
	}
	return 0
}

func (m *IsisShRepEl) GetRepairElementType() uint32 {
	if m != nil {
		return m.RepairElementType
	}
	return 0
}

func (m *IsisShRepEl) GetRepairStrictSpfLabel() uint32 {
	if m != nil {
		return m.RepairStrictSpfLabel
	}
	return 0
}

// FRR backup path
type IsisShIpv6FrrBackup struct {
	// Next hop neighbor ID
	NeighborId string `protobuf:"bytes,1,opt,name=neighbor_id,json=neighborId" json:"neighbor_id,omitempty"`
	// Interface to send the packet out of
	EgressInterface string `protobuf:"bytes,2,opt,name=egress_interface,json=egressInterface" json:"egress_interface,omitempty"`
	// Next hop neighbor's forwarding address
	NeighborAddress string `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
	// Tunnel interface to send the packet out of
	TunnelEgressInterface string `protobuf:"bytes,4,opt,name=tunnel_egress_interface,json=tunnelEgressInterface" json:"tunnel_egress_interface,omitempty"`
	// Next hop neighbor's SNPA
	NeighborSnpa *IsisSnpaType `protobuf:"bytes,5,opt,name=neighbor_snpa,json=neighborSnpa" json:"neighbor_snpa,omitempty"`
	// Remote LFA PQ Node's ID
	RemoteLfaSystemId string `protobuf:"bytes,6,opt,name=remote_lfa_system_id,json=remoteLfaSystemId" json:"remote_lfa_system_id,omitempty"`
	// Remote LFA Router ID
	RemoteLfaRouterId string `protobuf:"bytes,7,opt,name=remote_lfa_router_id,json=remoteLfaRouterId" json:"remote_lfa_router_id,omitempty"`
	// Remote LFA PQ Node's ID
	RemoteLfaSystemPid string `protobuf:"bytes,8,opt,name=remote_lfa_system_pid,json=remoteLfaSystemPid" json:"remote_lfa_system_pid,omitempty"`
	// Remote LFA Router ID
	RemoteLfaRouterPid string `protobuf:"bytes,9,opt,name=remote_lfa_router_pid,json=remoteLfaRouterPid" json:"remote_lfa_router_pid,omitempty"`
	// Distance to the network via this backup path
	TotalBackupDistance uint32 `protobuf:"varint,10,opt,name=total_backup_distance,json=totalBackupDistance" json:"total_backup_distance,omitempty"`
	// Segment routing sid value received from first hop
	SegmentRoutingSidValue uint32 `protobuf:"varint,11,opt,name=segment_routing_sid_value,json=segmentRoutingSidValue" json:"segment_routing_sid_value,omitempty"`
	// Number of SIDs in TI-LFA/rLFA
	NumSid uint32 `protobuf:"varint,12,opt,name=num_sid,json=numSid" json:"num_sid,omitempty"`
	// Segment routing sid values for TI-LFA/rLFA
	SegmentRoutingSidValues []uint32 `protobuf:"varint,13,rep,packed,name=segment_routing_sid_values,json=segmentRoutingSidValues" json:"segment_routing_sid_values,omitempty"`
	// Backup Repair List Size
	BackupRepairListSize uint32 `protobuf:"varint,14,opt,name=backup_repair_list_size,json=backupRepairListSize" json:"backup_repair_list_size,omitempty"`
	// Ti LFA computation which provided backup path
	TilfaComputation string `protobuf:"bytes,15,opt,name=tilfa_computation,json=tilfaComputation" json:"tilfa_computation,omitempty"`
	// BAckup Repair List
	BackupRepairList []*IsisShRepEl `protobuf:"bytes,16,rep,name=backup_repair_list,json=backupRepairList" json:"backup_repair_list,omitempty"`
	// PrefixSourceNodeID
	PrefixSourceNodeId string `protobuf:"bytes,17,opt,name=prefix_source_node_id,json=prefixSourceNodeId" json:"prefix_source_node_id,omitempty"`
	// Is the backup path via downstream node?
	IsDownstream bool `protobuf:"varint,18,opt,name=is_downstream,json=isDownstream" json:"is_downstream,omitempty"`
	// Is the backup path line card disjoint with primary?
	IsLcDisjoint bool `protobuf:"varint,19,opt,name=is_lc_disjoint,json=isLcDisjoint" json:"is_lc_disjoint,omitempty"`
	// Is the backup path node protecting?
	IsNodeProtecting bool `protobuf:"varint,20,opt,name=is_node_protecting,json=isNodeProtecting" json:"is_node_protecting,omitempty"`
	// Is the backup path an ECMP to the network?
	IsPrimaryPath bool `protobuf:"varint,21,opt,name=is_primary_path,json=isPrimaryPath" json:"is_primary_path,omitempty"`
	// Is the backup path SRLG disjoint with primary?
	IsSrlgDisjoint bool `protobuf:"varint,22,opt,name=is_srlg_disjoint,json=isSrlgDisjoint" json:"is_srlg_disjoint,omitempty"`
	// Is the backup path via a Remote LFA?
	IsRemoteLfa bool `protobuf:"varint,23,opt,name=is_remote_lfa,json=isRemoteLfa" json:"is_remote_lfa,omitempty"`
	// Is the backup path via a TI-LFA?
	IsEpcfrrLfa bool `protobuf:"varint,24,opt,name=is_epcfrr_lfa,json=isEpcfrrLfa" json:"is_epcfrr_lfa,omitempty"`
	// Is the backup path TI-LFA strict SPF?
	IsStrictSpflfa bool `protobuf:"varint,25,opt,name=is_strict_spflfa,json=isStrictSpflfa" json:"is_strict_spflfa,omitempty"`
	// Is SR TE tunnel requested
	IsTunnelRequested bool `protobuf:"varint,26,opt,name=is_tunnel_requested,json=isTunnelRequested" json:"is_tunnel_requested,omitempty"`
	// Weight configured on the interface
	Weight uint32 `protobuf:"varint,27,opt,name=weight" json:"weight,omitempty"`
}

func (m *IsisShIpv6FrrBackup) Reset()                    { *m = IsisShIpv6FrrBackup{} }
func (m *IsisShIpv6FrrBackup) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv6FrrBackup) ProtoMessage()               {}
func (*IsisShIpv6FrrBackup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *IsisShIpv6FrrBackup) GetNeighborId() string {
	if m != nil {
		return m.NeighborId
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetEgressInterface() string {
	if m != nil {
		return m.EgressInterface
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetTunnelEgressInterface() string {
	if m != nil {
		return m.TunnelEgressInterface
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetNeighborSnpa() *IsisSnpaType {
	if m != nil {
		return m.NeighborSnpa
	}
	return nil
}

func (m *IsisShIpv6FrrBackup) GetRemoteLfaSystemId() string {
	if m != nil {
		return m.RemoteLfaSystemId
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetRemoteLfaRouterId() string {
	if m != nil {
		return m.RemoteLfaRouterId
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetRemoteLfaSystemPid() string {
	if m != nil {
		return m.RemoteLfaSystemPid
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetRemoteLfaRouterPid() string {
	if m != nil {
		return m.RemoteLfaRouterPid
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetTotalBackupDistance() uint32 {
	if m != nil {
		return m.TotalBackupDistance
	}
	return 0
}

func (m *IsisShIpv6FrrBackup) GetSegmentRoutingSidValue() uint32 {
	if m != nil {
		return m.SegmentRoutingSidValue
	}
	return 0
}

func (m *IsisShIpv6FrrBackup) GetNumSid() uint32 {
	if m != nil {
		return m.NumSid
	}
	return 0
}

func (m *IsisShIpv6FrrBackup) GetSegmentRoutingSidValues() []uint32 {
	if m != nil {
		return m.SegmentRoutingSidValues
	}
	return nil
}

func (m *IsisShIpv6FrrBackup) GetBackupRepairListSize() uint32 {
	if m != nil {
		return m.BackupRepairListSize
	}
	return 0
}

func (m *IsisShIpv6FrrBackup) GetTilfaComputation() string {
	if m != nil {
		return m.TilfaComputation
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetBackupRepairList() []*IsisShRepEl {
	if m != nil {
		return m.BackupRepairList
	}
	return nil
}

func (m *IsisShIpv6FrrBackup) GetPrefixSourceNodeId() string {
	if m != nil {
		return m.PrefixSourceNodeId
	}
	return ""
}

func (m *IsisShIpv6FrrBackup) GetIsDownstream() bool {
	if m != nil {
		return m.IsDownstream
	}
	return false
}

func (m *IsisShIpv6FrrBackup) GetIsLcDisjoint() bool {
	if m != nil {
		return m.IsLcDisjoint
	}
	return false
}

func (m *IsisShIpv6FrrBackup) GetIsNodeProtecting() bool {
	if m != nil {
		return m.IsNodeProtecting
	}
	return false
}

func (m *IsisShIpv6FrrBackup) GetIsPrimaryPath() bool {
	if m != nil {
		return m.IsPrimaryPath
	}
	return false
}

func (m *IsisShIpv6FrrBackup) GetIsSrlgDisjoint() bool {
	if m != nil {
		return m.IsSrlgDisjoint
	}
	return false
}

func (m *IsisShIpv6FrrBackup) GetIsRemoteLfa() bool {
	if m != nil {
		return m.IsRemoteLfa
	}
	return false
}

func (m *IsisShIpv6FrrBackup) GetIsEpcfrrLfa() bool {
	if m != nil {
		return m.IsEpcfrrLfa
	}
	return false
}

func (m *IsisShIpv6FrrBackup) GetIsStrictSpflfa() bool {
	if m != nil {
		return m.IsStrictSpflfa
	}
	return false
}

func (m *IsisShIpv6FrrBackup) GetIsTunnelRequested() bool {
	if m != nil {
		return m.IsTunnelRequested
	}
	return false
}

func (m *IsisShIpv6FrrBackup) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

// IPv6 path to a destination
type IsisShIpv6Path struct {
	// Next hop neighbor ID
	NeighborId string `protobuf:"bytes,1,opt,name=neighbor_id,json=neighborId" json:"neighbor_id,omitempty"`
	// Interface to send the packet out of
	EgressInterface string `protobuf:"bytes,2,opt,name=egress_interface,json=egressInterface" json:"egress_interface,omitempty"`
	// Next hop neighbor's forwarding address
	NeighborAddress string `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
	// Next hop neighbor's SNPA
	NeighborSnpa *IsisSnpaType `protobuf:"bytes,4,opt,name=neighbor_snpa,json=neighborSnpa" json:"neighbor_snpa,omitempty"`
	// Tag associated with the path
	Tag uint32 `protobuf:"varint,5,opt,name=tag" json:"tag,omitempty"`
	// FRR backup for this path
	FrrBackup *IsisShIpv6FrrBackup `protobuf:"bytes,6,opt,name=frr_backup,json=frrBackup" json:"frr_backup,omitempty"`
	// Uloop Explicit List
	UloopExplicitList []*IsisShRepEl `protobuf:"bytes,7,rep,name=uloop_explicit_list,json=uloopExplicitList" json:"uloop_explicit_list,omitempty"`
	// Explicit path tunnel interface
	TunnelInterface string `protobuf:"bytes,8,opt,name=tunnel_interface,json=tunnelInterface" json:"tunnel_interface,omitempty"`
	// Segment routing sid value received from first hop
	SegmentRoutingSidValue uint32 `protobuf:"varint,9,opt,name=segment_routing_sid_value,json=segmentRoutingSidValue" json:"segment_routing_sid_value,omitempty"`
	// Weight configured on the interface
	Weight uint32 `protobuf:"varint,10,opt,name=weight" json:"weight,omitempty"`
	// Is path via a TE tunnel
	IsTeTunnelInterface bool `protobuf:"varint,11,opt,name=is_te_tunnel_interface,json=isTeTunnelInterface" json:"is_te_tunnel_interface,omitempty"`
	// Is path via an SR-exclude TE tunnel
	IsSrExcludeTunnelInterface bool `protobuf:"varint,12,opt,name=is_sr_exclude_tunnel_interface,json=isSrExcludeTunnelInterface" json:"is_sr_exclude_tunnel_interface,omitempty"`
}

func (m *IsisShIpv6Path) Reset()                    { *m = IsisShIpv6Path{} }
func (m *IsisShIpv6Path) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv6Path) ProtoMessage()               {}
func (*IsisShIpv6Path) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *IsisShIpv6Path) GetNeighborId() string {
	if m != nil {
		return m.NeighborId
	}
	return ""
}

func (m *IsisShIpv6Path) GetEgressInterface() string {
	if m != nil {
		return m.EgressInterface
	}
	return ""
}

func (m *IsisShIpv6Path) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *IsisShIpv6Path) GetNeighborSnpa() *IsisSnpaType {
	if m != nil {
		return m.NeighborSnpa
	}
	return nil
}

func (m *IsisShIpv6Path) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *IsisShIpv6Path) GetFrrBackup() *IsisShIpv6FrrBackup {
	if m != nil {
		return m.FrrBackup
	}
	return nil
}

func (m *IsisShIpv6Path) GetUloopExplicitList() []*IsisShRepEl {
	if m != nil {
		return m.UloopExplicitList
	}
	return nil
}

func (m *IsisShIpv6Path) GetTunnelInterface() string {
	if m != nil {
		return m.TunnelInterface
	}
	return ""
}

func (m *IsisShIpv6Path) GetSegmentRoutingSidValue() uint32 {
	if m != nil {
		return m.SegmentRoutingSidValue
	}
	return 0
}

func (m *IsisShIpv6Path) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *IsisShIpv6Path) GetIsTeTunnelInterface() bool {
	if m != nil {
		return m.IsTeTunnelInterface
	}
	return false
}

func (m *IsisShIpv6Path) GetIsSrExcludeTunnelInterface() bool {
	if m != nil {
		return m.IsSrExcludeTunnelInterface
	}
	return false
}

// SPT Neighbor
type IsisShTopoNeighbor struct {
	// Neighbor ID
	NeighborId string `protobuf:"bytes,1,opt,name=neighbor_id,json=neighborId" json:"neighbor_id,omitempty"`
	// Pseudonode between system and its neighbor
	IntermediatePseudonode *IsisNodeIdType `protobuf:"bytes,2,opt,name=intermediate_pseudonode,json=intermediatePseudonode" json:"intermediate_pseudonode,omitempty"`
}

func (m *IsisShTopoNeighbor) Reset()                    { *m = IsisShTopoNeighbor{} }
func (m *IsisShTopoNeighbor) String() string            { return proto.CompactTextString(m) }
func (*IsisShTopoNeighbor) ProtoMessage()               {}
func (*IsisShTopoNeighbor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *IsisShTopoNeighbor) GetNeighborId() string {
	if m != nil {
		return m.NeighborId
	}
	return ""
}

func (m *IsisShTopoNeighbor) GetIntermediatePseudonode() *IsisNodeIdType {
	if m != nil {
		return m.IntermediatePseudonode
	}
	return nil
}

// Status of a reachable IPv6 IS
type IsisShIpv6TopoReachableDetails struct {
	// Distance to the IS
	RootDistance uint32 `protobuf:"varint,1,opt,name=root_distance,json=rootDistance" json:"root_distance,omitempty"`
	// Distance to the IS
	MulticastRootDistance uint32 `protobuf:"varint,2,opt,name=multicast_root_distance,json=multicastRootDistance" json:"multicast_root_distance,omitempty"`
	// First hops towards the IS
	Paths []*IsisShIpv6Path `protobuf:"bytes,3,rep,name=paths" json:"paths,omitempty"`
	// Multicast intact first hops towards the IS
	MulticastPaths []*IsisShIpv6Path `protobuf:"bytes,4,rep,name=multicast_paths,json=multicastPaths" json:"multicast_paths,omitempty"`
	// Parents of the IS within the SPT
	Parents []*IsisShTopoNeighbor `protobuf:"bytes,5,rep,name=parents" json:"parents,omitempty"`
	// Children of the IS within the SPT
	Children []*IsisShTopoNeighbor `protobuf:"bytes,6,rep,name=children" json:"children,omitempty"`
}

func (m *IsisShIpv6TopoReachableDetails) Reset()                    { *m = IsisShIpv6TopoReachableDetails{} }
func (m *IsisShIpv6TopoReachableDetails) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv6TopoReachableDetails) ProtoMessage()               {}
func (*IsisShIpv6TopoReachableDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *IsisShIpv6TopoReachableDetails) GetRootDistance() uint32 {
	if m != nil {
		return m.RootDistance
	}
	return 0
}

func (m *IsisShIpv6TopoReachableDetails) GetMulticastRootDistance() uint32 {
	if m != nil {
		return m.MulticastRootDistance
	}
	return 0
}

func (m *IsisShIpv6TopoReachableDetails) GetPaths() []*IsisShIpv6Path {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *IsisShIpv6TopoReachableDetails) GetMulticastPaths() []*IsisShIpv6Path {
	if m != nil {
		return m.MulticastPaths
	}
	return nil
}

func (m *IsisShIpv6TopoReachableDetails) GetParents() []*IsisShTopoNeighbor {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *IsisShIpv6TopoReachableDetails) GetChildren() []*IsisShTopoNeighbor {
	if m != nil {
		return m.Children
	}
	return nil
}

// Reachability status of an IPv6 IS
type IsisShIpv6TopoReachableStatus struct {
	ReachableStatus string `protobuf:"bytes,1,opt,name=reachable_status,json=reachableStatus" json:"reachable_status,omitempty"`
	// Status of the IS within the SPT
	ReachableDetails *IsisShIpv6TopoReachableDetails `protobuf:"bytes,2,opt,name=reachable_details,json=reachableDetails" json:"reachable_details,omitempty"`
}

func (m *IsisShIpv6TopoReachableStatus) Reset()                    { *m = IsisShIpv6TopoReachableStatus{} }
func (m *IsisShIpv6TopoReachableStatus) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv6TopoReachableStatus) ProtoMessage()               {}
func (*IsisShIpv6TopoReachableStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *IsisShIpv6TopoReachableStatus) GetReachableStatus() string {
	if m != nil {
		return m.ReachableStatus
	}
	return ""
}

func (m *IsisShIpv6TopoReachableStatus) GetReachableDetails() *IsisShIpv6TopoReachableDetails {
	if m != nil {
		return m.ReachableDetails
	}
	return nil
}

func init() {
	proto.RegisterType((*IsisShIpv6TopoEntry_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_sh_ipv6_topo_entry_KEYS")
	proto.RegisterType((*IsisShIpv6TopoEntry)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_sh_ipv6_topo_entry")
	proto.RegisterType((*IsisNodeIdType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_node_id_type")
	proto.RegisterType((*IsisSnpaType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_snpa_type")
	proto.RegisterType((*IsisPerPriorityCounts)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_per_priority_counts")
	proto.RegisterType((*IsisShRepEl)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_sh_rep_el")
	proto.RegisterType((*IsisShIpv6FrrBackup)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_sh_ipv6_frr_backup")
	proto.RegisterType((*IsisShIpv6Path)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_sh_ipv6_path")
	proto.RegisterType((*IsisShTopoNeighbor)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_sh_topo_neighbor")
	proto.RegisterType((*IsisShIpv6TopoReachableDetails)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_sh_ipv6_topo_reachable_details")
	proto.RegisterType((*IsisShIpv6TopoReachableStatus)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv6_link_topologies.ipv6_link_topology.isis_sh_ipv6_topo_reachable_status")
}

func init() { proto.RegisterFile("isis_sh_ipv6_topo_entry.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x58, 0xbd, 0x6f, 0x1b, 0x47,
	0x16, 0xc7, 0xea, 0x5b, 0x43, 0x52, 0x22, 0x47, 0x1f, 0x5c, 0xcb, 0xe7, 0x3b, 0x1d, 0x7d, 0x67,
	0xc8, 0xb8, 0x03, 0x0f, 0x27, 0xdb, 0x0a, 0x82, 0x54, 0xb6, 0xa5, 0x42, 0x88, 0xe1, 0x08, 0x4b,
	0x21, 0x40, 0xaa, 0xc1, 0x68, 0x77, 0x48, 0x4e, 0xbc, 0xdc, 0xdd, 0xcc, 0xcc, 0xca, 0x92, 0xdb,
	0x34, 0xa9, 0xd2, 0x04, 0x48, 0x19, 0x20, 0x08, 0x82, 0xa4, 0x8e, 0x9b, 0x94, 0x41, 0x10, 0x20,
	0x45, 0x10, 0xf8, 0x0f, 0x48, 0x15, 0x20, 0x65, 0xea, 0xd4, 0xc1, 0xbc, 0x99, 0xd9, 0xa5, 0x97,
	0x96, 0x5c, 0x86, 0xea, 0x76, 0xdf, 0xef, 0xbd, 0x79, 0x1f, 0x7c, 0xef, 0xf7, 0x66, 0x89, 0x6e,
	0x70, 0xc9, 0x25, 0x91, 0x43, 0xc2, 0xb3, 0xd3, 0x3d, 0xa2, 0xd2, 0x2c, 0x25, 0x2c, 0x51, 0xe2,
	0xbc, 0x9b, 0x89, 0x54, 0xa5, 0xf8, 0x63, 0x2f, 0xe4, 0x32, 0x4c, 0x09, 0x4f, 0x25, 0x39, 0x13,
	0x24, 0x8c, 0x13, 0x49, 0xc0, 0x22, 0xcd, 0x98, 0xe8, 0xea, 0xa7, 0x2e, 0x4f, 0xa4, 0xa2, 0x49,
	0xc8, 0xca, 0xa7, 0xae, 0x3e, 0x26, 0x4e, 0x07, 0x9c, 0x49, 0xf7, 0x78, 0x5e, 0x3c, 0x90, 0x98,
	0x9d, 0xb2, 0x58, 0x56, 0xde, 0xbb, 0xe0, 0x3e, 0xe6, 0xc9, 0x13, 0x32, 0x66, 0x3c, 0x21, 0x3c,
	0xef, 0xbc, 0xf0, 0xd0, 0xdf, 0x2e, 0x08, 0x99, 0xbc, 0x7d, 0xf0, 0x5e, 0x0f, 0xdf, 0x44, 0x0d,
	0x17, 0x07, 0x49, 0xe8, 0x88, 0xf9, 0xde, 0xb6, 0xb7, 0xb3, 0x1c, 0xd4, 0x9d, 0xf0, 0x31, 0x1d,
	0x31, 0xdc, 0x46, 0x8b, 0xb4, 0x6f, 0xe0, 0x19, 0x80, 0x17, 0x68, 0x1f, 0x80, 0x6b, 0x68, 0x49,
	0x3a, 0x64, 0x16, 0x90, 0x45, 0x69, 0xa1, 0x9b, 0xa8, 0x51, 0xc4, 0x0c, 0xf8, 0x9c, 0x39, 0xd8,
	0x09, 0x41, 0x69, 0x1d, 0xcd, 0x43, 0x3e, 0xfe, 0x3c, 0x80, 0xe6, 0x05, 0x5f, 0x47, 0xcb, 0xf2,
	0x5c, 0x2a, 0x36, 0x22, 0x3c, 0xf2, 0x17, 0x00, 0x59, 0x32, 0x82, 0xc3, 0xa8, 0xf3, 0xe1, 0x3c,
	0x6a, 0x5f, 0x90, 0x11, 0xfe, 0x37, 0x5a, 0x91, 0x69, 0x2e, 0x42, 0x46, 0x68, 0x14, 0x09, 0x26,
	0xa5, 0xbf, 0x0b, 0xd6, 0x0d, 0x23, 0xbd, 0x6f, 0x84, 0x5a, 0x8d, 0x4b, 0x92, 0x51, 0xa1, 0x78,
	0xc8, 0x33, 0x9a, 0x28, 0xff, 0xce, 0xb6, 0xb7, 0xb3, 0x14, 0x34, 0xb8, 0x3c, 0x2a, 0x85, 0x50,
	0x1a, 0x49, 0xd2, 0x53, 0x26, 0xe2, 0x94, 0x46, 0x2c, 0xf2, 0xef, 0x82, 0x56, 0x9d, 0xcb, 0x77,
	0x0a, 0x19, 0xfe, 0x07, 0xaa, 0x71, 0x49, 0xa8, 0x52, 0x34, 0x1c, 0xb2, 0xc8, 0xbf, 0x07, 0x2a,
	0x88, 0xcb, 0xfb, 0x56, 0x82, 0x7f, 0xf3, 0xd0, 0x9a, 0x60, 0x34, 0x1c, 0xd2, 0x13, 0x1e, 0x73,
	0x75, 0x4e, 0xa4, 0xa2, 0x2a, 0x97, 0xfe, 0xde, 0xb6, 0xb7, 0x53, 0xdb, 0xfd, 0xc6, 0xeb, 0x4e,
	0x57, 0xc7, 0x74, 0x27, 0x6b, 0x6b, 0xa3, 0x8f, 0x99, 0x0d, 0x3d, 0xc0, 0xe3, 0xf9, 0xf4, 0x40,
	0x86, 0x7f, 0xf7, 0xd0, 0x0d, 0x1a, 0x9d, 0x32, 0xa1, 0xb8, 0x64, 0x11, 0xc9, 0x04, 0xeb, 0xf3,
	0x33, 0xc2, 0xf5, 0x8f, 0x18, 0xa6, 0x79, 0xa2, 0xa4, 0xff, 0x06, 0x24, 0xfc, 0xf5, 0x74, 0x26,
	0x9c, 0x31, 0x41, 0x32, 0xc1, 0x53, 0xa1, 0x7f, 0x21, 0x13, 0x70, 0xb0, 0x55, 0xe6, 0x73, 0x04,
	0xe9, 0x1c, 0x2a, 0x36, 0x7a, 0x08, 0x58, 0xe7, 0x36, 0x6a, 0x81, 0x5d, 0x92, 0x46, 0x8c, 0xf0,
	0x88, 0xa8, 0xf3, 0x0c, 0xba, 0xf9, 0x94, 0xc6, 0xb9, 0x9b, 0x21, 0xf3, 0xd2, 0xb9, 0xa5, 0xbb,
	0x4d, 0xd7, 0x34, 0xc9, 0xe8, 0x65, 0x7a, 0x0a, 0xf9, 0x17, 0x85, 0x82, 0xb7, 0xd0, 0x52, 0x28,
	0xb8, 0xe2, 0x21, 0x8d, 0xc1, 0xa8, 0x11, 0x14, 0xef, 0x18, 0xa3, 0xb9, 0x21, 0x1f, 0x0c, 0x61,
	0x32, 0x1b, 0x01, 0x3c, 0xe3, 0x4d, 0xb4, 0x30, 0x62, 0x11, 0xcf, 0x47, 0x30, 0x95, 0x8d, 0xc0,
	0xbe, 0xe1, 0x26, 0x9a, 0x8d, 0xd3, 0xa7, 0x30, 0x8a, 0x8d, 0x40, 0x3f, 0x76, 0xbe, 0x98, 0x71,
	0xe1, 0x0d, 0x89, 0x60, 0x19, 0x61, 0x31, 0xbe, 0x83, 0x36, 0x05, 0xcb, 0x28, 0x17, 0x84, 0xc5,
	0x6c, 0xc4, 0x12, 0xe5, 0xb2, 0xb4, 0xf1, 0xae, 0x19, 0xf4, 0xc0, 0x80, 0x8f, 0xd3, 0x88, 0x1d,
	0x46, 0x78, 0x07, 0x35, 0xad, 0x11, 0xcf, 0x4e, 0xef, 0xc2, 0xfc, 0x59, 0xae, 0x58, 0x31, 0xf2,
	0xc3, 0xec, 0xf4, 0xae, 0x1e, 0xc0, 0x97, 0x35, 0xf7, 0x8c, 0xe6, 0x6c, 0x45, 0x73, 0x0f, 0x34,
	0xff, 0x89, 0xea, 0x56, 0x33, 0xa6, 0x27, 0x2c, 0xb6, 0x61, 0xd7, 0x8c, 0xec, 0x91, 0x16, 0xe1,
	0x2e, 0x5a, 0xab, 0xc4, 0xaa, 0x2b, 0x0c, 0x74, 0xd2, 0x08, 0x5a, 0x2f, 0x05, 0x7a, 0xac, 0x4b,
	0x7f, 0x0f, 0xb5, 0xad, 0xbe, 0x54, 0x82, 0x87, 0x8a, 0xc8, 0xac, 0x6f, 0x4f, 0x5f, 0x00, 0x9b,
	0x75, 0x03, 0xf7, 0x00, 0xed, 0x65, 0x7d, 0x70, 0xd3, 0xf9, 0xa8, 0x5e, 0x21, 0x9d, 0xbe, 0x10,
	0xe4, 0x84, 0x86, 0x4f, 0xf2, 0x4c, 0x33, 0x40, 0xc2, 0xf8, 0x60, 0x78, 0x92, 0x8a, 0xb2, 0x46,
	0xc8, 0x89, 0x0e, 0x23, 0x7c, 0x1b, 0x35, 0xd9, 0x40, 0x13, 0x0f, 0xe1, 0x89, 0x62, 0xa2, 0x4f,
	0x43, 0x47, 0xa3, 0xab, 0x46, 0x7e, 0xe8, 0xc4, 0x5a, 0xb5, 0x38, 0xcb, 0x51, 0x98, 0xa9, 0xcd,
	0xaa, 0x93, 0x3b, 0x12, 0xdb, 0x43, 0x6d, 0x95, 0x27, 0x09, 0x8b, 0xc9, 0xc4, 0xe1, 0x86, 0x69,
	0x37, 0x0c, 0x7c, 0x50, 0x71, 0xf1, 0x83, 0x87, 0x1a, 0x85, 0x0f, 0xdd, 0x93, 0x50, 0xac, 0xda,
	0xee, 0x67, 0x53, 0xca, 0x44, 0x6e, 0x6a, 0x82, 0xba, 0x8b, 0xba, 0x97, 0x64, 0x14, 0xff, 0x0f,
	0xad, 0x0b, 0x36, 0x4a, 0x15, 0x23, 0x71, 0x9f, 0x92, 0xea, 0xba, 0x68, 0x19, 0xec, 0x51, 0x9f,
	0xf6, 0xec, 0xde, 0xa8, 0x18, 0x88, 0x34, 0x57, 0x0c, 0x7e, 0xaf, 0xc5, 0x8a, 0x41, 0x00, 0xc8,
	0x61, 0x84, 0xff, 0x8f, 0x36, 0x26, 0x3d, 0x64, 0x3c, 0xf2, 0x97, 0xc0, 0x02, 0x57, 0x5c, 0x1c,
	0xf1, 0xaa, 0x89, 0xf5, 0xa1, 0x4d, 0x96, 0x2b, 0x26, 0xc6, 0x89, 0x36, 0xd9, 0x45, 0x1b, 0x2a,
	0x55, 0x34, 0xb6, 0xdd, 0x44, 0x22, 0x6e, 0x0a, 0xeb, 0x23, 0x68, 0xc7, 0x35, 0x00, 0x1f, 0x00,
	0xb6, 0x6f, 0x21, 0xfc, 0x26, 0xba, 0x26, 0xd9, 0x00, 0xba, 0x5d, 0xfb, 0xe0, 0xc9, 0x80, 0x48,
	0x1e, 0x11, 0xc3, 0x29, 0x35, 0xb0, 0xdb, 0xb4, 0x0a, 0x81, 0xc1, 0x7b, 0x3c, 0x7a, 0x57, 0xa3,
	0x7a, 0x93, 0x27, 0xf9, 0x48, 0xab, 0xfb, 0x75, 0xc3, 0x0c, 0x49, 0x3e, 0xea, 0xf1, 0x08, 0xbf,
	0x85, 0xb6, 0x2e, 0x3c, 0x53, 0xfa, 0x8d, 0xed, 0xd9, 0x9d, 0x46, 0xd0, 0x7e, 0xf5, 0xa1, 0x52,
	0x4f, 0x95, 0x0d, 0xdf, 0xcd, 0x2b, 0x97, 0x8a, 0x48, 0xfe, 0x8c, 0xf9, 0x2b, 0x66, 0xaa, 0x0c,
	0x1c, 0x98, 0xc9, 0xe5, 0x52, 0xf5, 0xf8, 0x33, 0x86, 0xff, 0x83, 0x5a, 0x8a, 0xeb, 0x4a, 0x85,
	0xe9, 0x28, 0xcb, 0x15, 0x55, 0x3c, 0x4d, 0xfc, 0x55, 0x28, 0x55, 0x13, 0x80, 0x87, 0xa5, 0x1c,
	0xff, 0xec, 0x21, 0x3c, 0xe9, 0xc4, 0x6f, 0x6e, 0xcf, 0x4e, 0x71, 0xf3, 0x3a, 0x4e, 0x0d, 0x9a,
	0xd5, 0x02, 0xe8, 0x5e, 0xb1, 0x4b, 0xd2, 0x5e, 0x59, 0x1c, 0xc9, 0xb6, 0x4c, 0xaf, 0x18, 0xb0,
	0x07, 0x98, 0xe5, 0x58, 0x73, 0x21, 0x89, 0xd2, 0xa7, 0x89, 0x54, 0x82, 0xd1, 0x91, 0x8f, 0xdd,
	0x85, 0x64, 0xbf, 0x90, 0xe1, 0x7f, 0xc1, 0xe5, 0x26, 0x0e, 0x75, 0x27, 0xbd, 0x9f, 0xf2, 0x44,
	0xf9, 0x6b, 0x4e, 0xeb, 0x51, 0xb8, 0x6f, 0x65, 0xf8, 0xbf, 0x08, 0xbb, 0xed, 0xa5, 0x6f, 0xae,
	0x2c, 0xd4, 0x3f, 0xa8, 0xbf, 0x0e, 0x9a, 0x4d, 0x2e, 0xb5, 0xc3, 0xa3, 0x42, 0x8e, 0x6f, 0xa1,
	0x55, 0xbd, 0x98, 0x04, 0x1f, 0x51, 0x71, 0x4e, 0x32, 0xaa, 0x86, 0xfe, 0x46, 0x71, 0x63, 0x32,
	0xd2, 0x23, 0xaa, 0x86, 0x9a, 0xda, 0x75, 0xd6, 0x22, 0x1e, 0x94, 0xde, 0x37, 0x41, 0x71, 0x85,
	0xcb, 0x9e, 0x88, 0x07, 0x85, 0xff, 0x0e, 0xa4, 0x52, 0x0e, 0x8b, 0xdf, 0x06, 0xb5, 0x1a, 0x97,
	0x81, 0x9b, 0x11, 0xab, 0xc3, 0xb2, 0x50, 0x93, 0xad, 0xd6, 0xf1, 0x9d, 0xce, 0x01, 0xc8, 0xb4,
	0x8e, 0xf5, 0x58, 0x70, 0xb9, 0x56, 0xbb, 0x56, 0x78, 0x74, 0x24, 0x1e, 0xf7, 0xa9, 0xde, 0x14,
	0x5c, 0x12, 0x4b, 0x99, 0x82, 0x7d, 0x90, 0x33, 0xa9, 0x58, 0xe4, 0x6f, 0x81, 0x72, 0x8b, 0xcb,
	0x63, 0x40, 0x02, 0x07, 0xe8, 0x15, 0xfa, 0x54, 0x13, 0x8e, 0xf2, 0xaf, 0x9b, 0x41, 0x31, 0x6f,
	0x9d, 0x3f, 0x16, 0xed, 0xea, 0x77, 0xab, 0x40, 0x97, 0xe3, 0xaf, 0x5a, 0x02, 0x93, 0x64, 0x3e,
	0x77, 0x15, 0xc9, 0xbc, 0x89, 0x66, 0x15, 0x1d, 0xd8, 0xad, 0xad, 0x1f, 0xf1, 0x8f, 0x1e, 0x42,
	0xe5, 0x8e, 0x05, 0x56, 0xaf, 0xed, 0x7e, 0x35, 0xdd, 0x97, 0xe5, 0x32, 0xde, 0x60, 0xb9, 0x2f,
	0x84, 0x21, 0x6d, 0xfc, 0xc2, 0x43, 0x6b, 0x79, 0x9c, 0xa6, 0x19, 0x61, 0x67, 0x59, 0xcc, 0x43,
	0xae, 0x0c, 0x71, 0x2d, 0x5e, 0x0d, 0xe2, 0x6a, 0x41, 0xec, 0x07, 0x36, 0x74, 0x60, 0xae, 0xdb,
	0xa8, 0x69, 0xc7, 0xa8, 0x6c, 0x65, 0xb3, 0x13, 0x57, 0x8d, 0xbc, 0x6c, 0xe5, 0x4b, 0x37, 0xd5,
	0xf2, 0xa5, 0x9b, 0xaa, 0x9c, 0x3f, 0x34, 0x3e, 0x7f, 0xfa, 0x76, 0xaa, 0xe7, 0x98, 0x91, 0x89,
	0x18, 0x6a, 0x30, 0xca, 0x6b, 0x5c, 0x1e, 0xb3, 0xe3, 0x4a, 0x1c, 0x0f, 0xd0, 0xdf, 0x81, 0x98,
	0x08, 0x3b, 0x0b, 0xe3, 0x3c, 0x7a, 0x85, 0x71, 0x1d, 0x8c, 0xb7, 0x34, 0x4d, 0x1d, 0x18, 0x9d,
	0xca, 0x19, 0x9d, 0x4f, 0x67, 0xd0, 0x86, 0x2b, 0x0e, 0x7c, 0x17, 0xb9, 0x16, 0x7e, 0xfd, 0xf0,
	0xff, 0xe2, 0xa1, 0x36, 0xb8, 0xd2, 0xd7, 0x70, 0xaa, 0x18, 0xc9, 0x24, 0xcb, 0xa3, 0x54, 0xd3,
	0x2f, 0x90, 0x40, 0x6d, 0xf7, 0xf3, 0xe9, 0xec, 0x83, 0xf1, 0xcf, 0x9b, 0x60, 0x73, 0x3c, 0x85,
	0xa3, 0x22, 0x83, 0xce, 0x27, 0x8b, 0xe8, 0xe6, 0x65, 0x5f, 0x8d, 0x11, 0x53, 0x94, 0xc7, 0x52,
	0xaf, 0x2f, 0x91, 0xa6, 0xaa, 0xbc, 0xe2, 0x98, 0x2f, 0x99, 0xba, 0x16, 0x16, 0x77, 0x9b, 0x3d,
	0xd4, 0x1e, 0xe5, 0xb1, 0xfe, 0xb2, 0x91, 0xba, 0x67, 0xc6, 0xd5, 0xcd, 0x07, 0xce, 0x46, 0x01,
	0x07, 0xe3, 0x76, 0xdf, 0x7a, 0x68, 0x5e, 0x33, 0xb1, 0xa6, 0xca, 0xd9, 0xe9, 0x2d, 0xe8, 0xf8,
	0xd2, 0x08, 0x4c, 0xc0, 0xf8, 0x27, 0x0f, 0xad, 0x96, 0x39, 0x9b, 0x24, 0xe6, 0xae, 0x4c, 0x12,
	0x2b, 0x45, 0xe8, 0x47, 0x90, 0xcd, 0x77, 0x1e, 0x5a, 0xcc, 0xa8, 0x60, 0xfa, 0x93, 0x7f, 0x1e,
	0xb2, 0xf8, 0x72, 0x6a, 0xb3, 0x78, 0x69, 0x8c, 0x03, 0x17, 0x36, 0xfe, 0xde, 0x43, 0x4b, 0xe1,
	0x90, 0xc7, 0x91, 0x60, 0x89, 0xbf, 0x70, 0xa5, 0x72, 0x28, 0xe2, 0xee, 0x3c, 0x9f, 0x41, 0x9d,
	0xd7, 0xff, 0x97, 0xa3, 0xc9, 0xbc, 0x2a, 0xb3, 0x04, 0xb6, 0x5a, 0xc8, 0xed, 0x5f, 0x3c, 0xbf,
	0x7a, 0xa8, 0x35, 0x31, 0xd5, 0x96, 0xbf, 0x9e, 0x5f, 0xa9, 0xff, 0xb1, 0x6c, 0xec, 0x41, 0x99,
	0xfa, 0xbe, 0x91, 0x9c, 0x2c, 0xc0, 0xff, 0xb8, 0x77, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xb9,
	0x54, 0x29, 0x81, 0xe8, 0x15, 0x00, 0x00,
}
