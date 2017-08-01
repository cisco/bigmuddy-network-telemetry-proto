// Code generated by protoc-gen-go.
// source: ospf_sh_topology.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_local_routes_local_route is a generated protocol buffer package.

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
package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_local_routes_local_route

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
	Prefix       string `protobuf:"bytes,2,opt,name=prefix" json:"prefix,omitempty"`
	PrefixLength uint32 `protobuf:"varint,3,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
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
	proto.RegisterType((*OspfShTopology_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.local_routes.local_route.ospf_sh_topology_KEYS")
	proto.RegisterType((*OspfShTopology)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.local_routes.local_route.ospf_sh_topology")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.local_routes.local_route.ospf_sh_time")
	proto.RegisterType((*OspfShRepEl)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.local_routes.local_route.ospf_sh_rep_el")
	proto.RegisterType((*OspfShSrUloopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.local_routes.local_route.ospf_sh_sr_uloop_path")
	proto.RegisterType((*OspfShTopCommon)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.local_routes.local_route.ospf_sh_top_common")
	proto.RegisterType((*OspfShTopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.local_routes.local_route.ospf_sh_top_path")
}

func init() { proto.RegisterFile("ospf_sh_topology.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1051 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xcd, 0x6e, 0xe4, 0xc4,
	0x13, 0x97, 0x37, 0xf9, 0x27, 0x93, 0x9e, 0x64, 0x3f, 0x7a, 0xb3, 0x89, 0xb3, 0xff, 0x85, 0x84,
	0x41, 0x82, 0x08, 0xa1, 0xd1, 0x2a, 0x09, 0xdf, 0x07, 0x14, 0x41, 0x22, 0x46, 0x24, 0xd1, 0xca,
	0x93, 0x45, 0xe2, 0xd4, 0xea, 0xb1, 0xcb, 0x49, 0x4b, 0xb6, 0xdb, 0xea, 0xee, 0xc9, 0xc7, 0x22,
	0x24, 0xde, 0x80, 0x13, 0x37, 0x24, 0x0e, 0x9c, 0x39, 0xf1, 0x02, 0x3c, 0x0f, 0xef, 0x80, 0x84,
	0xba, 0xaa, 0x3d, 0x76, 0x36, 0x7b, 0x26, 0x37, 0xfb, 0x57, 0xd5, 0x55, 0xd5, 0x55, 0xbf, 0xaa,
	0x6a, 0xb6, 0xa6, 0x6d, 0x9d, 0x0b, 0x7b, 0x2e, 0x9c, 0xae, 0x75, 0xa1, 0xcf, 0xae, 0x87, 0xb5,
	0xd1, 0x4e, 0x73, 0x95, 0x2a, 0x9b, 0x6a, 0xa1, 0xb4, 0x15, 0x57, 0x46, 0xa8, 0xfa, 0x62, 0x4f,
	0xa0, 0xa6, 0xae, 0xc1, 0x0c, 0xfd, 0x97, 0xd7, 0x4b, 0xc1, 0x5a, 0xb0, 0xcd, 0xd7, 0x30, 0x83,
	0x5c, 0x4e, 0x0b, 0x27, 0x2e, 0x4c, 0x3e, 0x34, 0x7a, 0xea, 0x40, 0xa8, 0x2a, 0xd7, 0xa6, 0x94,
	0x4e, 0xe9, 0x6a, 0x58, 0xe8, 0x54, 0x16, 0x02, 0x71, 0xdb, 0xfd, 0x19, 0x5c, 0xb2, 0x27, 0xaf,
	0x07, 0x21, 0xbe, 0x3d, 0xf8, 0x7e, 0xcc, 0xdf, 0x61, 0xcb, 0xc1, 0xb4, 0xa8, 0x64, 0x09, 0x71,
	0xb4, 0x15, 0x6d, 0x2f, 0x25, 0xfd, 0x80, 0x9d, 0xc8, 0x12, 0xf8, 0x1a, 0x5b, 0xa8, 0x0d, 0xe4,
	0xea, 0x2a, 0xbe, 0x87, 0xc2, 0xf0, 0xc7, 0xdf, 0x65, 0x2b, 0xf4, 0x25, 0x0a, 0xa8, 0xce, 0xdc,
	0x79, 0x3c, 0xb7, 0x15, 0x6d, 0xaf, 0x24, 0xcb, 0x04, 0x1e, 0x21, 0x36, 0xf8, 0x65, 0x9e, 0x3d,
	0x7c, 0xdd, 0xb3, 0x77, 0x4a, 0xb1, 0x07, 0xbb, 0x3b, 0xe4, 0x14, 0xb1, 0x17, 0x64, 0x7c, 0xc8,
	0x1e, 0x77, 0x55, 0x1a, 0x17, 0xbb, 0xe8, 0xe2, 0x51, 0x47, 0x93, 0xfc, 0xb4, 0x26, 0x4b, 0x70,
	0x46, 0xa5, 0xf1, 0x1e, 0x2a, 0x92, 0xc9, 0x63, 0x84, 0xf8, 0x5b, 0x8c, 0x91, 0x8a, 0xbb, 0xae,
	0x21, 0xfe, 0x08, 0x7d, 0x2e, 0x21, 0x72, 0x7a, 0x5d, 0x03, 0x7f, 0x9f, 0x3d, 0x20, 0x71, 0xaa,
	0xab, 0x0a, 0x52, 0x07, 0x59, 0xfc, 0xf1, 0x56, 0xb4, 0xdd, 0x4b, 0xee, 0x23, 0xfc, 0x55, 0x83,
	0xf2, 0x5f, 0xa3, 0xc6, 0x90, 0x4f, 0x7d, 0xfc, 0xc9, 0x56, 0xb4, 0xdd, 0xdf, 0xf9, 0x71, 0xf8,
	0x9f, 0x15, 0x73, 0xd8, 0xc9, 0xa7, 0x48, 0x75, 0x59, 0xea, 0x2a, 0xdc, 0x63, 0x54, 0xe5, 0x9a,
	0xff, 0x1e, 0x35, 0x17, 0xa9, 0xa5, 0x3b, 0x17, 0x85, 0xb2, 0x2e, 0xfe, 0x74, 0x6b, 0x6e, 0xbb,
	0xbf, 0xf3, 0xc3, 0x1d, 0xc5, 0xe8, 0xe3, 0x48, 0x56, 0xa8, 0x66, 0xd2, 0x9d, 0x1f, 0x29, 0xeb,
	0x06, 0x87, 0x6c, 0x79, 0xa6, 0xa2, 0x88, 0x64, 0x16, 0x52, 0x5d, 0x65, 0xc8, 0xc0, 0x95, 0x24,
	0xfc, 0xf1, 0xb7, 0x19, 0xab, 0x64, 0xa5, 0x83, 0xec, 0x1e, 0xca, 0x3a, 0xc8, 0xe0, 0xe7, 0x88,
	0xdd, 0x6f, 0x0c, 0x19, 0xa8, 0x05, 0x14, 0xfc, 0x03, 0xf6, 0xc8, 0x40, 0x2d, 0x95, 0x11, 0x50,
	0x40, 0x09, 0x95, 0x13, 0x2a, 0x0b, 0xbc, 0x7e, 0x40, 0x82, 0x03, 0xc2, 0x47, 0x19, 0xd2, 0x86,
	0x74, 0x0b, 0x39, 0x81, 0x22, 0x38, 0xe8, 0x13, 0x76, 0xe4, 0x21, 0x64, 0xe2, 0x4d, 0x73, 0xc8,
	0x9f, 0xb9, 0xc0, 0xc4, 0xae, 0x41, 0xcf, 0xa3, 0xc1, 0x3f, 0xf7, 0xda, 0x5e, 0xb3, 0x46, 0x4c,
	0x0b, 0x1d, 0x52, 0xc0, 0xff, 0x88, 0xd8, 0x93, 0x52, 0xa5, 0x46, 0x23, 0xd4, 0xf8, 0xf5, 0xf5,
	0x89, 0xb0, 0x3e, 0xd7, 0x77, 0x50, 0x1f, 0xca, 0x59, 0xf2, 0x78, 0x16, 0x57, 0x42, 0x57, 0x57,
	0xd6, 0xf1, 0x2f, 0xd8, 0xd3, 0x37, 0x86, 0x2b, 0xac, 0x7a, 0x05, 0x21, 0x55, 0xeb, 0x6f, 0x38,
	0x38, 0x56, 0xaf, 0x80, 0x1f, 0xb0, 0xcd, 0xf6, 0xb0, 0x9b, 0x56, 0x15, 0x14, 0x42, 0x55, 0x0e,
	0x4c, 0x2e, 0x53, 0xa0, 0x59, 0x33, 0x87, 0x35, 0x79, 0x36, 0x53, 0x3b, 0x45, 0xad, 0x51, 0xa3,
	0x84, 0xc3, 0xe7, 0x39, 0x5b, 0x6d, 0xcd, 0x58, 0xdf, 0xc7, 0x4e, 0xd8, 0x3a, 0x8f, 0xe7, 0xb1,
	0x35, 0xf9, 0x4c, 0x36, 0x46, 0xd1, 0xb8, 0xce, 0x07, 0x7f, 0x2e, 0x30, 0x7e, 0xbb, 0x43, 0xf8,
	0x80, 0x11, 0x03, 0x85, 0x34, 0x20, 0x1b, 0x46, 0x34, 0x13, 0x62, 0xdf, 0x80, 0x1c, 0x65, 0xfc,
	0xbd, 0xa6, 0x73, 0xda, 0x39, 0x42, 0xb7, 0xa4, 0xa3, 0xa7, 0xcd, 0x24, 0xf1, 0x0c, 0x43, 0x3d,
	0xa3, 0x26, 0xe2, 0x02, 0x8c, 0x55, 0xba, 0x0a, 0x84, 0x20, 0x03, 0x89, 0x9a, 0x7c, 0x47, 0x70,
	0xab, 0xeb, 0x43, 0x6a, 0x74, 0x7d, 0xf4, 0xf3, 0x41, 0x77, 0x5c, 0xe7, 0x8d, 0xee, 0x1e, 0x5b,
	0x23, 0xdd, 0x5c, 0x9b, 0x4b, 0x69, 0x32, 0x91, 0x29, 0xeb, 0x64, 0x95, 0x42, 0xfc, 0x3f, 0x34,
	0xbe, 0x8a, 0xd2, 0x43, 0x12, 0x7e, 0x1d, 0x64, 0xed, 0xe8, 0xb3, 0x7a, 0x6a, 0x52, 0x88, 0x17,
	0x3a, 0x17, 0x1b, 0x23, 0xe4, 0x67, 0x42, 0x88, 0x62, 0x5a, 0x67, 0xd2, 0x5f, 0x50, 0x95, 0x10,
	0x2f, 0xe2, 0xe4, 0xba, 0xbc, 0x8b, 0xa9, 0xa0, 0x4a, 0x08, 0xd7, 0x7f, 0x89, 0x01, 0x9d, 0xfa,
	0x19, 0xf0, 0xdb, 0x6c, 0x72, 0xe5, 0x52, 0x15, 0x14, 0x63, 0xef, 0x6e, 0x63, 0xa4, 0xc2, 0x1f,
	0x4a, 0x55, 0x60, 0x84, 0x1f, 0x32, 0xde, 0x16, 0xb3, 0x36, 0x4a, 0x1b, 0xe5, 0xae, 0xe3, 0x25,
	0x4c, 0xf8, 0xc3, 0xa6, 0x9a, 0x2f, 0x02, 0xde, 0xee, 0x30, 0x39, 0x75, 0x5a, 0xc0, 0x55, 0x5a,
	0x4c, 0x33, 0xc8, 0x62, 0x86, 0xd4, 0xa5, 0x7a, 0xec, 0x4f, 0x9d, 0x3e, 0x08, 0x02, 0xfe, 0x25,
	0x7b, 0x16, 0xac, 0x9b, 0x76, 0xf1, 0x19, 0x38, 0x53, 0xd6, 0x81, 0x81, 0x2c, 0xee, 0xe3, 0xc1,
	0x0d, 0xf2, 0x63, 0x9a, 0x05, 0x98, 0xcc, 0x14, 0xf8, 0x67, 0x6c, 0xa3, 0x63, 0xa0, 0x9a, 0x98,
	0xee, 0xe9, 0x65, 0x8c, 0x72, 0x6d, 0x76, 0xfa, 0x64, 0x62, 0xda, 0xa3, 0x83, 0xbf, 0x6f, 0xee,
	0x69, 0x1a, 0x58, 0xcf, 0xd9, 0x6a, 0x93, 0xb8, 0x1b, 0x8d, 0x4b, 0xc3, 0x94, 0x87, 0x9d, 0xd3,
	0x6d, 0xd7, 0xdd, 0x86, 0xc1, 0x15, 0x5c, 0x39, 0x71, 0xae, 0x6b, 0x21, 0xb3, 0xcc, 0x80, 0xb5,
	0xe1, 0xed, 0x40, 0x09, 0x39, 0x81, 0x2b, 0xf7, 0x8d, 0xae, 0xf7, 0x49, 0x74, 0x8b, 0xc0, 0x73,
	0x9d, 0xe7, 0x40, 0x20, 0xf0, 0x26, 0xa3, 0x5f, 0x51, 0x58, 0xa9, 0x32, 0xec, 0x9f, 0xa5, 0x84,
	0xb6, 0xf0, 0x91, 0x47, 0xf8, 0xe7, 0xec, 0x69, 0x67, 0xe9, 0x29, 0x2b, 0xca, 0x54, 0x5a, 0xe7,
	0x03, 0x97, 0xa9, 0xc3, 0xf6, 0xe9, 0x85, 0xbb, 0xfb, 0x15, 0x34, 0xb2, 0xc7, 0x5e, 0x3c, 0x42,
	0x29, 0xdf, 0x63, 0xeb, 0x37, 0xcf, 0x4e, 0xd3, 0x92, 0x32, 0x80, 0xbd, 0xd4, 0x0b, 0x51, 0xd3,
	0xc1, 0x97, 0x69, 0x59, 0xfb, 0xaf, 0x5b, 0x2f, 0x8e, 0xc5, 0xdb, 0x2f, 0x8e, 0x0d, 0xd6, 0x2b,
	0xac, 0xa4, 0x7d, 0xd1, 0x43, 0xf1, 0x62, 0x61, 0x25, 0xbe, 0x36, 0xd6, 0xd9, 0x62, 0x33, 0x88,
	0x88, 0x3e, 0x0b, 0x92, 0x66, 0xd0, 0x26, 0xeb, 0xa3, 0x80, 0x98, 0x1a, 0xc8, 0xc2, 0x3c, 0x74,
	0x88, 0x08, 0xff, 0x2b, 0x62, 0xff, 0xb7, 0x46, 0xb4, 0x53, 0x51, 0x5e, 0x68, 0x95, 0xf9, 0x49,
	0x40, 0x21, 0xf7, 0xb1, 0x63, 0x7e, 0x8a, 0xee, 0xa0, 0x65, 0x6e, 0xac, 0xbb, 0x24, 0xb6, 0xe6,
	0xb8, 0x09, 0x72, 0xbf, 0x89, 0xd1, 0xa7, 0x6e, 0xb2, 0x80, 0xef, 0xdf, 0xdd, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x24, 0xc4, 0xa2, 0xa7, 0x19, 0x0b, 0x00, 0x00,
}
