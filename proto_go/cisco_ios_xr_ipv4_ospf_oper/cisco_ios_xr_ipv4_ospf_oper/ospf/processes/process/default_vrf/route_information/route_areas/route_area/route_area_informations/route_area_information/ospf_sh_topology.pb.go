// Code generated by protoc-gen-go.
// source: ospf_sh_topology.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_areas_route_area_route_area_informations_route_area_information is a generated protocol buffer package.

It is generated from these files:
	ospf_sh_topology.proto

It has these top-level messages:
	OspfShTopology_KEYS
	OspfShTopology
	OspfShTime
	OspfShRepEl
	OspfShSrUloopPath
	OspfShTopCommon
	OspfShTopPath
*/
package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_areas_route_area_route_area_informations_route_area_information

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

// OSPF Route Information
type OspfShTopology_KEYS struct {
	ProcessName  string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	AreaId       uint32 `protobuf:"varint,2,opt,name=area_id,json=areaId" json:"area_id,omitempty"`
	Prefix       string `protobuf:"bytes,3,opt,name=prefix" json:"prefix,omitempty"`
	PrefixLength uint32 `protobuf:"varint,4,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
}

func (m *OspfShTopology_KEYS) Reset()                    { *m = OspfShTopology_KEYS{} }
func (m *OspfShTopology_KEYS) String() string            { return proto.CompactTextString(m) }
func (*OspfShTopology_KEYS) ProtoMessage()               {}
func (*OspfShTopology_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OspfShTopology_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShTopology_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShTopology_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *OspfShTopology_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type OspfShTopology struct {
	// Prefix
	RoutePrefix string `protobuf:"bytes,50,opt,name=route_prefix,json=routePrefix" json:"route_prefix,omitempty"`
	// Prefix length
	RoutePrefixLength uint32 `protobuf:"varint,51,opt,name=route_prefix_length,json=routePrefixLength" json:"route_prefix_length,omitempty"`
	// Metric
	RouteMetric uint32 `protobuf:"varint,52,opt,name=route_metric,json=routeMetric" json:"route_metric,omitempty"`
	// Route type
	RouteType string `protobuf:"bytes,53,opt,name=route_type,json=routeType" json:"route_type,omitempty"`
	// If true, connected route
	RouteConnected bool `protobuf:"varint,54,opt,name=route_connected,json=routeConnected" json:"route_connected,omitempty"`
	// Route information
	RouteInfo *OspfShTopCommon `protobuf:"bytes,55,opt,name=route_info,json=routeInfo" json:"route_info,omitempty"`
	// List of paths to this route
	RoutePathList []*OspfShTopPath `protobuf:"bytes,56,rep,name=route_path_list,json=routePathList" json:"route_path_list,omitempty"`
}

func (m *OspfShTopology) Reset()                    { *m = OspfShTopology{} }
func (m *OspfShTopology) String() string            { return proto.CompactTextString(m) }
func (*OspfShTopology) ProtoMessage()               {}
func (*OspfShTopology) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OspfShTopology) GetRoutePrefix() string {
	if m != nil {
		return m.RoutePrefix
	}
	return ""
}

func (m *OspfShTopology) GetRoutePrefixLength() uint32 {
	if m != nil {
		return m.RoutePrefixLength
	}
	return 0
}

func (m *OspfShTopology) GetRouteMetric() uint32 {
	if m != nil {
		return m.RouteMetric
	}
	return 0
}

func (m *OspfShTopology) GetRouteType() string {
	if m != nil {
		return m.RouteType
	}
	return ""
}

func (m *OspfShTopology) GetRouteConnected() bool {
	if m != nil {
		return m.RouteConnected
	}
	return false
}

func (m *OspfShTopology) GetRouteInfo() *OspfShTopCommon {
	if m != nil {
		return m.RouteInfo
	}
	return nil
}

func (m *OspfShTopology) GetRoutePathList() []*OspfShTopPath {
	if m != nil {
		return m.RoutePathList
	}
	return nil
}

type OspfShTime struct {
	Second     uint32 `protobuf:"varint,1,opt,name=second" json:"second,omitempty"`
	Nanosecond uint32 `protobuf:"varint,2,opt,name=nanosecond" json:"nanosecond,omitempty"`
}

func (m *OspfShTime) Reset()                    { *m = OspfShTime{} }
func (m *OspfShTime) String() string            { return proto.CompactTextString(m) }
func (*OspfShTime) ProtoMessage()               {}
func (*OspfShTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OspfShTime) GetSecond() uint32 {
	if m != nil {
		return m.Second
	}
	return 0
}

func (m *OspfShTime) GetNanosecond() uint32 {
	if m != nil {
		return m.Nanosecond
	}
	return 0
}

// OSPF Repair Element
type OspfShRepEl struct {
	// Repair Element ID
	RepairElementId string `protobuf:"bytes,1,opt,name=repair_element_id,json=repairElementId" json:"repair_element_id,omitempty"`
	// Repair Label
	RepairLabel uint32 `protobuf:"varint,2,opt,name=repair_label,json=repairLabel" json:"repair_label,omitempty"`
	// Repair Element Type
	RepairElementType uint32 `protobuf:"varint,3,opt,name=repair_element_type,json=repairElementType" json:"repair_element_type,omitempty"`
}

func (m *OspfShRepEl) Reset()                    { *m = OspfShRepEl{} }
func (m *OspfShRepEl) String() string            { return proto.CompactTextString(m) }
func (*OspfShRepEl) ProtoMessage()               {}
func (*OspfShRepEl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OspfShRepEl) GetRepairElementId() string {
	if m != nil {
		return m.RepairElementId
	}
	return ""
}

func (m *OspfShRepEl) GetRepairLabel() uint32 {
	if m != nil {
		return m.RepairLabel
	}
	return 0
}

func (m *OspfShRepEl) GetRepairElementType() uint32 {
	if m != nil {
		return m.RepairElementType
	}
	return 0
}

// OSPF Route SR Uloop Path Information
type OspfShSrUloopPath struct {
	// Microloop Repair List
	MicroloopRepairList []*OspfShRepEl `protobuf:"bytes,1,rep,name=microloop_repair_list,json=microloopRepairList" json:"microloop_repair_list,omitempty"`
	// Microloop Repair List Size
	MicroloopRepairListSize uint32 `protobuf:"varint,2,opt,name=microloop_repair_list_size,json=microloopRepairListSize" json:"microloop_repair_list_size,omitempty"`
	// Microloop Tunnel Interface name
	MicroloopTunnelInterfaceName string `protobuf:"bytes,3,opt,name=microloop_tunnel_interface_name,json=microloopTunnelInterfaceName" json:"microloop_tunnel_interface_name,omitempty"`
	// Strict SPF SID
	MicroloopStrictSpf bool `protobuf:"varint,4,opt,name=microloop_strict_spf,json=microloopStrictSpf" json:"microloop_strict_spf,omitempty"`
}

func (m *OspfShSrUloopPath) Reset()                    { *m = OspfShSrUloopPath{} }
func (m *OspfShSrUloopPath) String() string            { return proto.CompactTextString(m) }
func (*OspfShSrUloopPath) ProtoMessage()               {}
func (*OspfShSrUloopPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OspfShSrUloopPath) GetMicroloopRepairList() []*OspfShRepEl {
	if m != nil {
		return m.MicroloopRepairList
	}
	return nil
}

func (m *OspfShSrUloopPath) GetMicroloopRepairListSize() uint32 {
	if m != nil {
		return m.MicroloopRepairListSize
	}
	return 0
}

func (m *OspfShSrUloopPath) GetMicroloopTunnelInterfaceName() string {
	if m != nil {
		return m.MicroloopTunnelInterfaceName
	}
	return ""
}

func (m *OspfShSrUloopPath) GetMicroloopStrictSpf() bool {
	if m != nil {
		return m.MicroloopStrictSpf
	}
	return false
}

// OSPF Common Route Information
type OspfShTopCommon struct {
	// Area ID
	RouteAreaId uint32 `protobuf:"varint,1,opt,name=route_area_id,json=routeAreaId" json:"route_area_id,omitempty"`
	// TE metric
	RouteTeMetric uint32 `protobuf:"varint,2,opt,name=route_te_metric,json=routeTeMetric" json:"route_te_metric,omitempty"`
	// RIB version
	RouteRibVersion uint32 `protobuf:"varint,3,opt,name=route_rib_version,json=routeRibVersion" json:"route_rib_version,omitempty"`
	// SPF version
	RouteSpfVersion uint64 `protobuf:"varint,4,opt,name=route_spf_version,json=routeSpfVersion" json:"route_spf_version,omitempty"`
	// Forward distance
	RouteForwardDistance uint32 `protobuf:"varint,5,opt,name=route_forward_distance,json=routeForwardDistance" json:"route_forward_distance,omitempty"`
	// Protocol source
	RouteSource uint32 `protobuf:"varint,6,opt,name=route_source,json=routeSource" json:"route_source,omitempty"`
	// Last time updated
	RouteUpdateTime *OspfShTime `protobuf:"bytes,7,opt,name=route_update_time,json=routeUpdateTime" json:"route_update_time,omitempty"`
	// Last time update failed
	RouteFailTime *OspfShTime `protobuf:"bytes,8,opt,name=route_fail_time,json=routeFailTime" json:"route_fail_time,omitempty"`
	// SPF priority
	RouteSpfPriority uint32 `protobuf:"varint,9,opt,name=route_spf_priority,json=routeSpfPriority" json:"route_spf_priority,omitempty"`
	// If true, exclude from TE paths
	RouteAutoExcluded bool `protobuf:"varint,10,opt,name=route_auto_excluded,json=routeAutoExcluded" json:"route_auto_excluded,omitempty"`
	// If true, SRTE registered prefix route
	RouteSrtePrefixRegistered bool `protobuf:"varint,11,opt,name=route_srte_prefix_registered,json=routeSrtePrefixRegistered" json:"route_srte_prefix_registered,omitempty"`
	// SRTE registered neigbhor count on route
	RouteSrteNbrRegistered uint32 `protobuf:"varint,12,opt,name=route_srte_nbr_registered,json=routeSrteNbrRegistered" json:"route_srte_nbr_registered,omitempty"`
}

func (m *OspfShTopCommon) Reset()                    { *m = OspfShTopCommon{} }
func (m *OspfShTopCommon) String() string            { return proto.CompactTextString(m) }
func (*OspfShTopCommon) ProtoMessage()               {}
func (*OspfShTopCommon) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *OspfShTopCommon) GetRouteAreaId() uint32 {
	if m != nil {
		return m.RouteAreaId
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteTeMetric() uint32 {
	if m != nil {
		return m.RouteTeMetric
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteRibVersion() uint32 {
	if m != nil {
		return m.RouteRibVersion
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteSpfVersion() uint64 {
	if m != nil {
		return m.RouteSpfVersion
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteForwardDistance() uint32 {
	if m != nil {
		return m.RouteForwardDistance
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteSource() uint32 {
	if m != nil {
		return m.RouteSource
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteUpdateTime() *OspfShTime {
	if m != nil {
		return m.RouteUpdateTime
	}
	return nil
}

func (m *OspfShTopCommon) GetRouteFailTime() *OspfShTime {
	if m != nil {
		return m.RouteFailTime
	}
	return nil
}

func (m *OspfShTopCommon) GetRouteSpfPriority() uint32 {
	if m != nil {
		return m.RouteSpfPriority
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteAutoExcluded() bool {
	if m != nil {
		return m.RouteAutoExcluded
	}
	return false
}

func (m *OspfShTopCommon) GetRouteSrtePrefixRegistered() bool {
	if m != nil {
		return m.RouteSrtePrefixRegistered
	}
	return false
}

func (m *OspfShTopCommon) GetRouteSrteNbrRegistered() uint32 {
	if m != nil {
		return m.RouteSrteNbrRegistered
	}
	return 0
}

// OSPF Route Path Information
type OspfShTopPath struct {
	// Next hop Interface
	RouteInterfaceName string `protobuf:"bytes,1,opt,name=route_interface_name,json=routeInterfaceName" json:"route_interface_name,omitempty"`
	// Nexthop IP address
	RouteNextHopAddress string `protobuf:"bytes,2,opt,name=route_next_hop_address,json=routeNextHopAddress" json:"route_next_hop_address,omitempty"`
	// IP address of source of route
	RouteSource string `protobuf:"bytes,3,opt,name=route_source,json=routeSource" json:"route_source,omitempty"`
	// LSA ID, see RFC2328
	RouteLsaid string `protobuf:"bytes,4,opt,name=route_lsaid,json=routeLsaid" json:"route_lsaid,omitempty"`
	// Multicast-intact path
	RoutePathIsMcastIntact bool `protobuf:"varint,5,opt,name=route_path_is_mcast_intact,json=routePathIsMcastIntact" json:"route_path_is_mcast_intact,omitempty"`
	// UCMP path
	RoutePathIsUcmpPath bool `protobuf:"varint,6,opt,name=route_path_is_ucmp_path,json=routePathIsUcmpPath" json:"route_path_is_ucmp_path,omitempty"`
	// Metric
	RouteMetric uint32 `protobuf:"varint,7,opt,name=route_metric,json=routeMetric" json:"route_metric,omitempty"`
	// LSA type, see RFC2328 etc.
	LsaType uint32 `protobuf:"varint,8,opt,name=lsa_type,json=lsaType" json:"lsa_type,omitempty"`
	// Area ID
	AreaId uint32 `protobuf:"varint,9,opt,name=area_id,json=areaId" json:"area_id,omitempty"`
	// Area format IP or uint32
	AreaFormat bool `protobuf:"varint,10,opt,name=area_format,json=areaFormat" json:"area_format,omitempty"`
	// SR Microloop avoidance Path Info
	SrMicroloopAvoidancePath *OspfShSrUloopPath `protobuf:"bytes,11,opt,name=sr_microloop_avoidance_path,json=srMicroloopAvoidancePath" json:"sr_microloop_avoidance_path,omitempty"`
}

func (m *OspfShTopPath) Reset()                    { *m = OspfShTopPath{} }
func (m *OspfShTopPath) String() string            { return proto.CompactTextString(m) }
func (*OspfShTopPath) ProtoMessage()               {}
func (*OspfShTopPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *OspfShTopPath) GetRouteInterfaceName() string {
	if m != nil {
		return m.RouteInterfaceName
	}
	return ""
}

func (m *OspfShTopPath) GetRouteNextHopAddress() string {
	if m != nil {
		return m.RouteNextHopAddress
	}
	return ""
}

func (m *OspfShTopPath) GetRouteSource() string {
	if m != nil {
		return m.RouteSource
	}
	return ""
}

func (m *OspfShTopPath) GetRouteLsaid() string {
	if m != nil {
		return m.RouteLsaid
	}
	return ""
}

func (m *OspfShTopPath) GetRoutePathIsMcastIntact() bool {
	if m != nil {
		return m.RoutePathIsMcastIntact
	}
	return false
}

func (m *OspfShTopPath) GetRoutePathIsUcmpPath() bool {
	if m != nil {
		return m.RoutePathIsUcmpPath
	}
	return false
}

func (m *OspfShTopPath) GetRouteMetric() uint32 {
	if m != nil {
		return m.RouteMetric
	}
	return 0
}

func (m *OspfShTopPath) GetLsaType() uint32 {
	if m != nil {
		return m.LsaType
	}
	return 0
}

func (m *OspfShTopPath) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShTopPath) GetAreaFormat() bool {
	if m != nil {
		return m.AreaFormat
	}
	return false
}

func (m *OspfShTopPath) GetSrMicroloopAvoidancePath() *OspfShSrUloopPath {
	if m != nil {
		return m.SrMicroloopAvoidancePath
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShTopology_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.route_area_informations.route_area_information.ospf_sh_topology_KEYS")
	proto.RegisterType((*OspfShTopology)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.route_area_informations.route_area_information.ospf_sh_topology")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.route_area_informations.route_area_information.ospf_sh_time")
	proto.RegisterType((*OspfShRepEl)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.route_area_informations.route_area_information.ospf_sh_rep_el")
	proto.RegisterType((*OspfShSrUloopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.route_area_informations.route_area_information.ospf_sh_sr_uloop_path")
	proto.RegisterType((*OspfShTopCommon)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.route_area_informations.route_area_information.ospf_sh_top_common")
	proto.RegisterType((*OspfShTopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.route_area_informations.route_area_information.ospf_sh_top_path")
}

func init() { proto.RegisterFile("ospf_sh_topology.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1069 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xdd, 0x6e, 0xe4, 0x34,
	0x14, 0x56, 0x68, 0xb7, 0x9d, 0xf1, 0xb4, 0xfb, 0xe3, 0xed, 0xb6, 0xe9, 0xb2, 0xd0, 0x32, 0x48,
	0x30, 0x42, 0x68, 0xb4, 0x6a, 0xcb, 0xff, 0x05, 0xaa, 0xa0, 0x15, 0x23, 0xda, 0x6a, 0x95, 0xe9,
	0x22, 0x71, 0x65, 0x79, 0x12, 0xa7, 0x63, 0x29, 0x89, 0x23, 0xdb, 0x29, 0xd3, 0x7d, 0x89, 0xe5,
	0x0a, 0x5e, 0x00, 0x21, 0x04, 0xe2, 0x45, 0x90, 0xb8, 0xe0, 0x9a, 0x67, 0xe0, 0x82, 0x37, 0x40,
	0x3e, 0xc7, 0x99, 0x64, 0xb6, 0xcb, 0x7d, 0xf7, 0xce, 0xf9, 0xce, 0x8f, 0xcf, 0x39, 0x3e, 0xe7,
	0xb3, 0x43, 0x36, 0x95, 0x29, 0x53, 0x66, 0xa6, 0xcc, 0xaa, 0x52, 0x65, 0xea, 0xe2, 0x6a, 0x58,
	0x6a, 0x65, 0x15, 0xfd, 0x21, 0x88, 0xa5, 0x89, 0x15, 0x93, 0xca, 0xb0, 0x99, 0x66, 0xb2, 0xbc,
	0x3c, 0x60, 0xa0, 0xaa, 0x4a, 0xa1, 0x87, 0x6e, 0xe5, 0x14, 0x63, 0x61, 0x8c, 0x30, 0xf5, 0x6a,
	0x98, 0x88, 0x94, 0x57, 0x99, 0x65, 0x97, 0x3a, 0x1d, 0x6a, 0x55, 0x59, 0xc1, 0x64, 0x91, 0x2a,
	0x9d, 0x73, 0x2b, 0x55, 0xe1, 0x11, 0xae, 0x05, 0x37, 0xad, 0x75, 0x6b, 0xd9, 0xd6, 0x36, 0xff,
	0x83, 0xf7, 0xbf, 0x0f, 0xc8, 0x83, 0x17, 0x63, 0x66, 0x5f, 0x1f, 0x7d, 0x3b, 0xa6, 0x6f, 0x91,
	0x35, 0x1f, 0x08, 0x2b, 0x78, 0x2e, 0xc2, 0x60, 0x37, 0x18, 0x74, 0xa3, 0x9e, 0xc7, 0xce, 0x78,
	0x2e, 0xe8, 0x16, 0x59, 0x45, 0x87, 0x49, 0xf8, 0xda, 0x6e, 0x30, 0x58, 0x8f, 0x56, 0xdc, 0xe7,
	0x28, 0xa1, 0x9b, 0x64, 0xa5, 0xd4, 0x22, 0x95, 0xb3, 0x70, 0x09, 0xac, 0xfc, 0x17, 0x7d, 0x9b,
	0xac, 0xe3, 0x8a, 0x65, 0xa2, 0xb8, 0xb0, 0xd3, 0x70, 0x19, 0xcc, 0xd6, 0x10, 0x3c, 0x01, 0xac,
	0xff, 0xef, 0x32, 0xb9, 0xfb, 0x62, 0x48, 0x2e, 0x1a, 0xcc, 0xc0, 0xfb, 0xdd, 0xc3, 0x68, 0x00,
	0x7b, 0x82, 0xce, 0x87, 0xe4, 0x7e, 0x5b, 0xa5, 0xde, 0x62, 0x1f, 0xb6, 0xb8, 0xd7, 0xd2, 0xc4,
	0x7d, 0x1a, 0x97, 0xb9, 0xb0, 0x5a, 0xc6, 0xe1, 0x01, 0x28, 0xa2, 0xcb, 0x53, 0x80, 0xe8, 0x1b,
	0x84, 0xa0, 0x8a, 0xbd, 0x2a, 0x45, 0xf8, 0x01, 0xec, 0xd9, 0x05, 0xe4, 0xfc, 0xaa, 0x14, 0xf4,
	0x5d, 0x72, 0x07, 0xc5, 0xb1, 0x2a, 0x0a, 0x11, 0x5b, 0x91, 0x84, 0x1f, 0xee, 0x06, 0x83, 0x4e,
	0x74, 0x1b, 0xe0, 0x2f, 0x6a, 0x94, 0xfe, 0x11, 0xd4, 0x8e, 0x5c, 0xed, 0xc3, 0x8f, 0x76, 0x83,
	0x41, 0x6f, 0xef, 0xd7, 0x60, 0x78, 0x33, 0x9b, 0x62, 0xd8, 0xaa, 0x3e, 0x8b, 0x55, 0x9e, 0xab,
	0xc2, 0x67, 0x3d, 0x2a, 0x52, 0x45, 0xff, 0x0a, 0xea, 0xb4, 0x4b, 0x6e, 0xa7, 0x2c, 0x93, 0xc6,
	0x86, 0x1f, 0xef, 0x2e, 0x0d, 0x7a, 0x7b, 0xbf, 0xbc, 0x12, 0x19, 0xb9, 0xa8, 0xa3, 0x75, 0xec,
	0x07, 0x6e, 0xa7, 0x27, 0xd2, 0xd8, 0xfe, 0x31, 0x59, 0x9b, 0xab, 0xc8, 0x5c, 0xb8, 0x06, 0x36,
	0x22, 0x56, 0x45, 0x02, 0x6d, 0xbf, 0x1e, 0xf9, 0x2f, 0xfa, 0x26, 0x21, 0x05, 0x2f, 0x94, 0x97,
	0x61, 0xd3, 0xb7, 0x90, 0xfe, 0xf3, 0x80, 0xdc, 0xae, 0x1d, 0x69, 0x51, 0x32, 0x91, 0xd1, 0xf7,
	0xc8, 0x3d, 0x2d, 0x4a, 0x2e, 0x35, 0x13, 0x99, 0xc8, 0x45, 0x61, 0xdd, 0xb8, 0xe0, 0x30, 0xdd,
	0x41, 0xc1, 0x11, 0xe2, 0xa3, 0x04, 0x5a, 0x12, 0x75, 0x33, 0x3e, 0x11, 0x99, 0xdf, 0xa0, 0x87,
	0xd8, 0x89, 0x83, 0xa0, 0xcb, 0x17, 0xdd, 0x41, 0x6f, 0x2e, 0xf9, 0x2e, 0x6f, 0x3b, 0x74, 0x3d,
	0xda, 0xff, 0x6d, 0xa9, 0x19, 0x70, 0xa3, 0x59, 0x95, 0x29, 0x5f, 0x02, 0xfa, 0x77, 0x40, 0x1e,
	0xe4, 0x32, 0xd6, 0x0a, 0xa0, 0x7a, 0x5f, 0x77, 0x9a, 0x01, 0x9c, 0xe6, 0xcf, 0x37, 0xfe, 0x34,
	0xb1, 0xc2, 0xd1, 0xfd, 0x79, 0x16, 0x11, 0x16, 0x4a, 0x1a, 0x4b, 0x3f, 0x23, 0x0f, 0x5f, 0x9a,
	0x1c, 0x33, 0xf2, 0x99, 0xf0, 0x85, 0xdd, 0x7a, 0x89, 0xe1, 0x58, 0x3e, 0x13, 0xf4, 0x88, 0xec,
	0x34, 0xc6, 0xb6, 0x2a, 0x0a, 0x91, 0x31, 0x59, 0x58, 0xa1, 0x53, 0x1e, 0x0b, 0xa4, 0x43, 0x24,
	0xb6, 0x47, 0x73, 0xb5, 0x73, 0xd0, 0x1a, 0xd5, 0x4a, 0xc0, 0x8f, 0x8f, 0xc9, 0x46, 0xe3, 0xc6,
	0x38, 0x46, 0xb1, 0xcc, 0x94, 0x29, 0xb0, 0x5e, 0x27, 0xa2, 0x73, 0xd9, 0x18, 0x44, 0xe3, 0x32,
	0xed, 0x3f, 0x5f, 0x25, 0xf4, 0xfa, 0xf4, 0xd1, 0x3e, 0x59, 0x6f, 0x97, 0xa2, 0xee, 0x4a, 0xe4,
	0xaa, 0x43, 0xe4, 0xdc, 0x77, 0xea, 0xa9, 0x6c, 0x18, 0x0d, 0xb3, 0x44, 0xd3, 0xf3, 0x9a, 0xd3,
	0x5c, 0x3f, 0x82, 0x9e, 0x96, 0x13, 0x76, 0x29, 0xb4, 0x91, 0xaa, 0xf0, 0xed, 0x83, 0x0e, 0x22,
	0x39, 0xf9, 0x06, 0xe1, 0x46, 0xd7, 0x85, 0x54, 0xeb, 0xba, 0xe8, 0x97, 0xbd, 0xee, 0xb8, 0x4c,
	0x6b, 0xdd, 0x03, 0xb2, 0x89, 0xba, 0xa9, 0xd2, 0xdf, 0x71, 0x9d, 0xb0, 0x44, 0x1a, 0xcb, 0x8b,
	0x58, 0x84, 0xb7, 0xc0, 0xf9, 0x06, 0x48, 0x8f, 0x51, 0xf8, 0xa5, 0x97, 0x35, 0x24, 0x6c, 0x54,
	0xa5, 0x63, 0x11, 0xae, 0xb4, 0x12, 0x1b, 0x03, 0xe4, 0xf8, 0xc6, 0x47, 0x51, 0x95, 0x09, 0x77,
	0x09, 0xca, 0x5c, 0x84, 0xab, 0xc0, 0xa1, 0x3f, 0xdd, 0x7c, 0xc6, 0x91, 0xb9, 0xf0, 0xc5, 0x7a,
	0x0a, 0xe1, 0x9f, 0x3b, 0x7e, 0xf9, 0x73, 0xce, 0xa1, 0x29, 0x97, 0x19, 0x66, 0xd4, 0x79, 0x95,
	0x32, 0xc2, 0xa6, 0x3a, 0xe6, 0x32, 0x83, 0x7c, 0xde, 0x27, 0xb4, 0x69, 0x94, 0x52, 0x4b, 0xa5,
	0xa5, 0xbd, 0x0a, 0xbb, 0x70, 0x98, 0x77, 0xeb, 0x4e, 0x79, 0xe2, 0xf1, 0xe6, 0xa6, 0xe6, 0x95,
	0x55, 0x4c, 0xcc, 0xe2, 0xac, 0x4a, 0x44, 0x12, 0x12, 0x18, 0x0b, 0x3c, 0xeb, 0xc3, 0xca, 0xaa,
	0x23, 0x2f, 0xa0, 0x9f, 0x93, 0x47, 0xde, 0xbb, 0x6e, 0xae, 0x77, 0x2d, 0x2e, 0xa4, 0xb1, 0x42,
	0x8b, 0x24, 0xec, 0x81, 0xe1, 0x36, 0xee, 0xa3, 0xeb, 0x6b, 0x3e, 0x9a, 0x2b, 0xd0, 0x4f, 0xc8,
	0x76, 0xcb, 0x41, 0x31, 0xd1, 0x6d, 0xeb, 0x35, 0x88, 0x72, 0x73, 0x6e, 0x7d, 0x36, 0xd1, 0x8d,
	0x69, 0xff, 0xc7, 0x5b, 0x0b, 0xaf, 0x11, 0xa4, 0xce, 0xc7, 0x64, 0xa3, 0xae, 0xf2, 0x02, 0x29,
	0x20, 0xad, 0x53, 0x7f, 0x57, 0xb6, 0xa9, 0x60, 0xbf, 0x9e, 0x8e, 0x42, 0xcc, 0x2c, 0x9b, 0xaa,
	0x92, 0xf1, 0x24, 0xd1, 0xc2, 0x18, 0x18, 0xd2, 0x6e, 0x84, 0x05, 0x39, 0x13, 0x33, 0xfb, 0x95,
	0x2a, 0x0f, 0x51, 0x74, 0x6d, 0x38, 0x96, 0x5a, 0x8f, 0x1e, 0x3f, 0x1c, 0x3b, 0x04, 0x3f, 0x59,
	0x66, 0xb8, 0x4c, 0x60, 0x36, 0xbb, 0x11, 0xbe, 0x35, 0x4e, 0x1c, 0x42, 0x3f, 0x25, 0x0f, 0x5b,
	0x97, 0xb5, 0x34, 0x2c, 0x8f, 0xb9, 0xb1, 0x2e, 0x70, 0x1e, 0x5b, 0x18, 0xcd, 0x8e, 0xcf, 0xdd,
	0x5d, 0x86, 0x23, 0x73, 0xea, 0xc4, 0x23, 0x90, 0xd2, 0x03, 0xb2, 0xb5, 0x68, 0x5b, 0xc5, 0x39,
	0x56, 0x00, 0xe6, 0xb4, 0xe3, 0xa3, 0x46, 0xc3, 0xa7, 0x71, 0x5e, 0xba, 0xd5, 0xb5, 0x77, 0xd5,
	0xea, 0xf5, 0x77, 0xd5, 0x36, 0xe9, 0x64, 0x86, 0xe3, 0xcd, 0xd5, 0x01, 0xf1, 0x6a, 0x66, 0x38,
	0xbc, 0xa9, 0x5a, 0x6f, 0xca, 0xee, 0xc2, 0x9b, 0x72, 0x87, 0xf4, 0x40, 0x80, 0x0d, 0xe9, 0x9b,
	0x85, 0x38, 0xe8, 0x18, 0x10, 0xfa, 0x4f, 0x40, 0x5e, 0x37, 0x9a, 0x35, 0x8c, 0xcb, 0x2f, 0x95,
	0x4c, 0x1c, 0xcb, 0x60, 0xc8, 0x3d, 0x98, 0xaf, 0xdf, 0x6f, 0xfc, 0x7c, 0x2d, 0xdc, 0xd2, 0x51,
	0x68, 0xf4, 0x69, 0x9d, 0xd1, 0x61, 0x9d, 0x90, 0xab, 0xf3, 0x64, 0x05, 0x7e, 0x2d, 0xf6, 0xff,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0x66, 0x98, 0x46, 0x5a, 0x74, 0x0c, 0x00, 0x00,
}
