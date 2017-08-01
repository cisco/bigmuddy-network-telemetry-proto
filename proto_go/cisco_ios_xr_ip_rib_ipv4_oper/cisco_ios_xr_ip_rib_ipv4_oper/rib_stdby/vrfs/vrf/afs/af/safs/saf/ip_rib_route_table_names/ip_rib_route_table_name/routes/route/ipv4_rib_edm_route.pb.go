// Code generated by protoc-gen-go.
// source: ipv4_rib_edm_route.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_routes_route is a generated protocol buffer package.

It is generated from these files:
	ipv4_rib_edm_route.proto

It has these top-level messages:
	Ipv4RibEdmRoute_KEYS
	Ipv4RibEdmRoute
	Ipv4RibEdmPath
	Ipv4RibEdmPathItem
	RibEdmLocalLabel
*/
package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_routes_route

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

// Information of a rib route head and rib proto route
type Ipv4RibEdmRoute_KEYS struct {
	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName" json:"route_table_name,omitempty"`
	Address        string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	PrefixLength   uint32 `protobuf:"varint,6,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
	NextHopAddress string `protobuf:"bytes,7,opt,name=next_hop_address,json=nextHopAddress" json:"next_hop_address,omitempty"`
	InterfaceName  string `protobuf:"bytes,8,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *Ipv4RibEdmRoute_KEYS) Reset()                    { *m = Ipv4RibEdmRoute_KEYS{} }
func (m *Ipv4RibEdmRoute_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmRoute_KEYS) ProtoMessage()               {}
func (*Ipv4RibEdmRoute_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv4RibEdmRoute_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv4RibEdmRoute_KEYS) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ipv4RibEdmRoute struct {
	// Route prefix
	Prefix string `protobuf:"bytes,50,opt,name=prefix" json:"prefix,omitempty"`
	// Length of prefix
	PrefixLength uint32 `protobuf:"varint,51,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
	// Route version
	RouteVersion uint32 `protobuf:"varint,52,opt,name=route_version,json=routeVersion" json:"route_version,omitempty"`
	// Protocol advertising the route
	ProtocolId uint32 `protobuf:"varint,53,opt,name=protocol_id,json=protocolId" json:"protocol_id,omitempty"`
	//  Name of Protocol
	ProtocolName string `protobuf:"bytes,54,opt,name=protocol_name,json=protocolName" json:"protocol_name,omitempty"`
	// Instance name
	Instance string `protobuf:"bytes,55,opt,name=instance" json:"instance,omitempty"`
	// Client adding the route to RIB
	ClientId uint32 `protobuf:"varint,56,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Route type
	RouteType uint32 `protobuf:"varint,57,opt,name=route_type,json=routeType" json:"route_type,omitempty"`
	// Route priority
	Priority uint32 `protobuf:"varint,58,opt,name=priority" json:"priority,omitempty"`
	// SVD Type of route
	SvdType uint32 `protobuf:"varint,59,opt,name=svd_type,json=svdType" json:"svd_type,omitempty"`
	// Route flags
	Flags uint32 `protobuf:"varint,60,opt,name=flags" json:"flags,omitempty"`
	// Extended Route flags
	ExtendedFlags uint64 `protobuf:"varint,61,opt,name=extended_flags,json=extendedFlags" json:"extended_flags,omitempty"`
	// Opaque proto specific info
	Tag uint32 `protobuf:"varint,62,opt,name=tag" json:"tag,omitempty"`
	// Distance of the route
	Distance uint32 `protobuf:"varint,63,opt,name=distance" json:"distance,omitempty"`
	// Diversion distance of the route
	DiversionDistance uint32 `protobuf:"varint,64,opt,name=diversion_distance,json=diversionDistance" json:"diversion_distance,omitempty"`
	// Route metric
	Metric uint32 `protobuf:"varint,65,opt,name=metric" json:"metric,omitempty"`
	// Number of paths
	PathsCount uint32 `protobuf:"varint,66,opt,name=paths_count,json=pathsCount" json:"paths_count,omitempty"`
	// BGP Attribute ID
	AttributeIdentity uint32 `protobuf:"varint,67,opt,name=attribute_identity,json=attributeIdentity" json:"attribute_identity,omitempty"`
	// BGP Traffic Index
	TrafficIndex uint32 `protobuf:"varint,68,opt,name=traffic_index,json=trafficIndex" json:"traffic_index,omitempty"`
	// Route ip precedence
	RoutePrecedence uint32 `protobuf:"varint,69,opt,name=route_precedence,json=routePrecedence" json:"route_precedence,omitempty"`
	// Route qos group
	QosGroup uint32 `protobuf:"varint,70,opt,name=qos_group,json=qosGroup" json:"qos_group,omitempty"`
	// Flow tag
	FlowTag uint32 `protobuf:"varint,71,opt,name=flow_tag,json=flowTag" json:"flow_tag,omitempty"`
	// Forward Class
	FwdClass uint32 `protobuf:"varint,72,opt,name=fwd_class,json=fwdClass" json:"fwd_class,omitempty"`
	// Number of pic paths in this route
	PicCount uint32 `protobuf:"varint,73,opt,name=pic_count,json=picCount" json:"pic_count,omitempty"`
	// Is the route active or backup
	Active bool `protobuf:"varint,74,opt,name=active" json:"active,omitempty"`
	// Route has a diversion path
	Diversion bool `protobuf:"varint,75,opt,name=diversion" json:"diversion,omitempty"`
	// Diversion route protocol name
	DiversionProtoName string `protobuf:"bytes,76,opt,name=diversion_proto_name,json=diversionProtoName" json:"diversion_proto_name,omitempty"`
	// Age of route (seconds)
	RouteAge uint32 `protobuf:"varint,77,opt,name=route_age,json=routeAge" json:"route_age,omitempty"`
	// Local label of the route
	RouteLabel uint32 `protobuf:"varint,78,opt,name=route_label,json=routeLabel" json:"route_label,omitempty"`
	// Route Version
	Version uint32 `protobuf:"varint,79,opt,name=version" json:"version,omitempty"`
	// Table Version
	TblVersion uint64 `protobuf:"varint,80,opt,name=tbl_version,json=tblVersion" json:"tbl_version,omitempty"`
	// Route modification time(nanoseconds)
	RouteModifyTime uint64 `protobuf:"varint,81,opt,name=route_modify_time,json=routeModifyTime" json:"route_modify_time,omitempty"`
	// Path(s) of the route
	RoutePath *Ipv4RibEdmPath `protobuf:"bytes,82,opt,name=route_path,json=routePath" json:"route_path,omitempty"`
}

func (m *Ipv4RibEdmRoute) Reset()                    { *m = Ipv4RibEdmRoute{} }
func (m *Ipv4RibEdmRoute) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmRoute) ProtoMessage()               {}
func (*Ipv4RibEdmRoute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv4RibEdmRoute) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteVersion() uint32 {
	if m != nil {
		return m.RouteVersion
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetProtocolId() uint32 {
	if m != nil {
		return m.ProtocolId
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteType() uint32 {
	if m != nil {
		return m.RouteType
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetSvdType() uint32 {
	if m != nil {
		return m.SvdType
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetExtendedFlags() uint64 {
	if m != nil {
		return m.ExtendedFlags
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetDiversionDistance() uint32 {
	if m != nil {
		return m.DiversionDistance
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetAttributeIdentity() uint32 {
	if m != nil {
		return m.AttributeIdentity
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTrafficIndex() uint32 {
	if m != nil {
		return m.TrafficIndex
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRoutePrecedence() uint32 {
	if m != nil {
		return m.RoutePrecedence
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetQosGroup() uint32 {
	if m != nil {
		return m.QosGroup
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFlowTag() uint32 {
	if m != nil {
		return m.FlowTag
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFwdClass() uint32 {
	if m != nil {
		return m.FwdClass
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPicCount() uint32 {
	if m != nil {
		return m.PicCount
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Ipv4RibEdmRoute) GetDiversion() bool {
	if m != nil {
		return m.Diversion
	}
	return false
}

func (m *Ipv4RibEdmRoute) GetDiversionProtoName() string {
	if m != nil {
		return m.DiversionProtoName
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetRouteAge() uint32 {
	if m != nil {
		return m.RouteAge
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTblVersion() uint64 {
	if m != nil {
		return m.TblVersion
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteModifyTime() uint64 {
	if m != nil {
		return m.RouteModifyTime
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRoutePath() *Ipv4RibEdmPath {
	if m != nil {
		return m.RoutePath
	}
	return nil
}

// Information of a rib path
type Ipv4RibEdmPath struct {
	// Next path
	Ipv4RibEdmPath []*Ipv4RibEdmPathItem `protobuf:"bytes,1,rep,name=ipv4_rib_edm_path,json=ipv4RibEdmPath" json:"ipv4_rib_edm_path,omitempty"`
}

func (m *Ipv4RibEdmPath) Reset()                    { *m = Ipv4RibEdmPath{} }
func (m *Ipv4RibEdmPath) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmPath) ProtoMessage()               {}
func (*Ipv4RibEdmPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ipv4RibEdmPath) GetIpv4RibEdmPath() []*Ipv4RibEdmPathItem {
	if m != nil {
		return m.Ipv4RibEdmPath
	}
	return nil
}

type Ipv4RibEdmPathItem struct {
	// Nexthop
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// Infosource
	InformationSource string `protobuf:"bytes,2,opt,name=information_source,json=informationSource" json:"information_source,omitempty"`
	// V6 nexthop
	V6Nexthop string `protobuf:"bytes,3,opt,name=v6_nexthop,json=v6Nexthop" json:"v6_nexthop,omitempty"`
	// V6 Infosource
	V6InformationSource string `protobuf:"bytes,4,opt,name=v6_information_source,json=v6InformationSource" json:"v6_information_source,omitempty"`
	// Interface Name
	InterfaceName string `protobuf:"bytes,5,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	// Metrics
	Metric uint32 `protobuf:"varint,6,opt,name=metric" json:"metric,omitempty"`
	// Load Metrics
	LoadMetric uint32 `protobuf:"varint,7,opt,name=load_metric,json=loadMetric" json:"load_metric,omitempty"`
	// Flags extended to 64 bits
	Flags64 uint64 `protobuf:"varint,8,opt,name=flags64" json:"flags64,omitempty"`
	// Flags
	Flags uint32 `protobuf:"varint,9,opt,name=flags" json:"flags,omitempty"`
	// Private Flags
	PrivateFlags uint32 `protobuf:"varint,10,opt,name=private_flags,json=privateFlags" json:"private_flags,omitempty"`
	// Looping path
	Looped bool `protobuf:"varint,11,opt,name=looped" json:"looped,omitempty"`
	// Nexthop tableid
	NextHopTableId uint32 `protobuf:"varint,12,opt,name=next_hop_table_id,json=nextHopTableId" json:"next_hop_table_id,omitempty"`
	// VRF Name of the nh table
	NextHopVrfName string `protobuf:"bytes,13,opt,name=next_hop_vrf_name,json=nextHopVrfName" json:"next_hop_vrf_name,omitempty"`
	// NH table name
	NextHopTableName string `protobuf:"bytes,14,opt,name=next_hop_table_name,json=nextHopTableName" json:"next_hop_table_name,omitempty"`
	// NH afi
	NextHopAfi uint32 `protobuf:"varint,15,opt,name=next_hop_afi,json=nextHopAfi" json:"next_hop_afi,omitempty"`
	// NH safi
	NextHopSafi uint32 `protobuf:"varint,16,opt,name=next_hop_safi,json=nextHopSafi" json:"next_hop_safi,omitempty"`
	// Label associated with this path
	RouteLabel uint32 `protobuf:"varint,17,opt,name=route_label,json=routeLabel" json:"route_label,omitempty"`
	// Tunnel ID associated with this path
	TunnelId uint32 `protobuf:"varint,18,opt,name=tunnel_id,json=tunnelId" json:"tunnel_id,omitempty"`
	// Path id of this path
	Pathid uint32 `protobuf:"varint,19,opt,name=pathid" json:"pathid,omitempty"`
	// Path id of this path's backup
	BackupPathid uint32 `protobuf:"varint,20,opt,name=backup_pathid,json=backupPathid" json:"backup_pathid,omitempty"`
	// Refcnt of backup
	RefCntOfBackup uint32 `protobuf:"varint,21,opt,name=ref_cnt_of_backup,json=refCntOfBackup" json:"ref_cnt_of_backup,omitempty"`
	// Number of extended communities
	NumberOfExtendedCommunities uint32 `protobuf:"varint,22,opt,name=number_of_extended_communities,json=numberOfExtendedCommunities" json:"number_of_extended_communities,omitempty"`
	// MVPN attribute present
	MvpnPresent bool `protobuf:"varint,23,opt,name=mvpn_present,json=mvpnPresent" json:"mvpn_present,omitempty"`
	// Path RT present
	PathrtPresent           bool `protobuf:"varint,24,opt,name=pathrt_present,json=pathrtPresent" json:"pathrt_present,omitempty"`
	VrfimportrtPresent      bool `protobuf:"varint,25,opt,name=vrfimportrt_present,json=vrfimportrtPresent" json:"vrfimportrt_present,omitempty"`
	SourceasrtPresent       bool `protobuf:"varint,26,opt,name=sourceasrt_present,json=sourceasrtPresent" json:"sourceasrt_present,omitempty"`
	SourcerdPresent         bool `protobuf:"varint,27,opt,name=sourcerd_present,json=sourcerdPresent" json:"sourcerd_present,omitempty"`
	SegmentedNexthopPresent bool `protobuf:"varint,28,opt,name=segmented_nexthop_present,json=segmentedNexthopPresent" json:"segmented_nexthop_present,omitempty"`
	// NHID associated with this path
	NextHopId uint32 `protobuf:"varint,29,opt,name=next_hop_id,json=nextHopId" json:"next_hop_id,omitempty"`
	// NHID references
	NextHopIdRefcount uint32 `protobuf:"varint,30,opt,name=next_hop_id_refcount,json=nextHopIdRefcount" json:"next_hop_id_refcount,omitempty"`
	// OSPF area associated with the path
	OspfAreaId string `protobuf:"bytes,31,opt,name=ospf_area_id,json=ospfAreaId" json:"ospf_area_id,omitempty"`
	// Remote backup node address
	RemoteBackupAddr [][]byte `protobuf:"bytes,32,rep,name=remote_backup_addr,json=remoteBackupAddr,proto3" json:"remote_backup_addr,omitempty"`
	// Path has a label stack
	HasLabelstk bool `protobuf:"varint,33,opt,name=has_labelstk,json=hasLabelstk" json:"has_labelstk,omitempty"`
	// Number of labels in stack
	NumLabels uint32 `protobuf:"varint,34,opt,name=num_labels,json=numLabels" json:"num_labels,omitempty"`
	// Outgoing label stack for this path
	Labelstk []uint32 `protobuf:"varint,35,rep,packed,name=labelstk" json:"labelstk,omitempty"`
	// binding Label for this path
	BindingLabel uint32 `protobuf:"varint,36,opt,name=binding_label,json=bindingLabel" json:"binding_label,omitempty"`
	// Fib nhid encap id
	NhidFeid uint64 `protobuf:"varint,37,opt,name=nhid_feid,json=nhidFeid" json:"nhid_feid,omitempty"`
	// Fib mpls encap id
	MplsFeid uint64 `protobuf:"varint,38,opt,name=mpls_feid,json=mplsFeid" json:"mpls_feid,omitempty"`
}

func (m *Ipv4RibEdmPathItem) Reset()                    { *m = Ipv4RibEdmPathItem{} }
func (m *Ipv4RibEdmPathItem) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmPathItem) ProtoMessage()               {}
func (*Ipv4RibEdmPathItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Ipv4RibEdmPathItem) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetInformationSource() string {
	if m != nil {
		return m.InformationSource
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetV6Nexthop() string {
	if m != nil {
		return m.V6Nexthop
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetV6InformationSource() string {
	if m != nil {
		return m.V6InformationSource
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLoadMetric() uint32 {
	if m != nil {
		return m.LoadMetric
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetFlags64() uint64 {
	if m != nil {
		return m.Flags64
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetPrivateFlags() uint32 {
	if m != nil {
		return m.PrivateFlags
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLooped() bool {
	if m != nil {
		return m.Looped
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNextHopTableId() uint32 {
	if m != nil {
		return m.NextHopTableId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopVrfName() string {
	if m != nil {
		return m.NextHopVrfName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetNextHopTableName() string {
	if m != nil {
		return m.NextHopTableName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetNextHopAfi() uint32 {
	if m != nil {
		return m.NextHopAfi
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopSafi() uint32 {
	if m != nil {
		return m.NextHopSafi
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetPathid() uint32 {
	if m != nil {
		return m.Pathid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetBackupPathid() uint32 {
	if m != nil {
		return m.BackupPathid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetRefCntOfBackup() uint32 {
	if m != nil {
		return m.RefCntOfBackup
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNumberOfExtendedCommunities() uint32 {
	if m != nil {
		return m.NumberOfExtendedCommunities
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetMvpnPresent() bool {
	if m != nil {
		return m.MvpnPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetPathrtPresent() bool {
	if m != nil {
		return m.PathrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetVrfimportrtPresent() bool {
	if m != nil {
		return m.VrfimportrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSourceasrtPresent() bool {
	if m != nil {
		return m.SourceasrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSourcerdPresent() bool {
	if m != nil {
		return m.SourcerdPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSegmentedNexthopPresent() bool {
	if m != nil {
		return m.SegmentedNexthopPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNextHopId() uint32 {
	if m != nil {
		return m.NextHopId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopIdRefcount() uint32 {
	if m != nil {
		return m.NextHopIdRefcount
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetOspfAreaId() string {
	if m != nil {
		return m.OspfAreaId
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetRemoteBackupAddr() [][]byte {
	if m != nil {
		return m.RemoteBackupAddr
	}
	return nil
}

func (m *Ipv4RibEdmPathItem) GetHasLabelstk() bool {
	if m != nil {
		return m.HasLabelstk
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNumLabels() uint32 {
	if m != nil {
		return m.NumLabels
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLabelstk() []uint32 {
	if m != nil {
		return m.Labelstk
	}
	return nil
}

func (m *Ipv4RibEdmPathItem) GetBindingLabel() uint32 {
	if m != nil {
		return m.BindingLabel
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNhidFeid() uint64 {
	if m != nil {
		return m.NhidFeid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetMplsFeid() uint64 {
	if m != nil {
		return m.MplsFeid
	}
	return 0
}

// Information of local label for route head
type RibEdmLocalLabel struct {
	// Protocol Name
	ProtocolName string `protobuf:"bytes,1,opt,name=protocol_name,json=protocolName" json:"protocol_name,omitempty"`
	// Client ID
	ClientId uint32 `protobuf:"varint,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Stale
	Stale uint32 `protobuf:"varint,3,opt,name=stale" json:"stale,omitempty"`
	// Mirrored
	Mirrored uint32 `protobuf:"varint,4,opt,name=mirrored" json:"mirrored,omitempty"`
	// Redist only
	RedistOnly uint32 `protobuf:"varint,5,opt,name=redist_only,json=redistOnly" json:"redist_only,omitempty"`
	// Local label
	Label uint32 `protobuf:"varint,6,opt,name=label" json:"label,omitempty"`
	// Distance
	Distance uint32 `protobuf:"varint,7,opt,name=distance" json:"distance,omitempty"`
}

func (m *RibEdmLocalLabel) Reset()                    { *m = RibEdmLocalLabel{} }
func (m *RibEdmLocalLabel) String() string            { return proto.CompactTextString(m) }
func (*RibEdmLocalLabel) ProtoMessage()               {}
func (*RibEdmLocalLabel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RibEdmLocalLabel) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *RibEdmLocalLabel) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *RibEdmLocalLabel) GetStale() uint32 {
	if m != nil {
		return m.Stale
	}
	return 0
}

func (m *RibEdmLocalLabel) GetMirrored() uint32 {
	if m != nil {
		return m.Mirrored
	}
	return 0
}

func (m *RibEdmLocalLabel) GetRedistOnly() uint32 {
	if m != nil {
		return m.RedistOnly
	}
	return 0
}

func (m *RibEdmLocalLabel) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

func (m *RibEdmLocalLabel) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.routes.route.ipv4_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv4RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.routes.route.ipv4_rib_edm_route")
	proto.RegisterType((*Ipv4RibEdmPath)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.routes.route.ipv4_rib_edm_path")
	proto.RegisterType((*Ipv4RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.routes.route.ipv4_rib_edm_path_item")
	proto.RegisterType((*RibEdmLocalLabel)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.routes.route.rib_edm_local_label")
}

func init() { proto.RegisterFile("ipv4_rib_edm_route.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0x4d, 0x72, 0x1b, 0xbb,
	0x11, 0x2e, 0x5a, 0x16, 0x29, 0x82, 0xa2, 0x2c, 0x8e, 0xf4, 0x24, 0xe8, 0xe9, 0x3d, 0x9b, 0x96,
	0xe3, 0x14, 0x95, 0x8a, 0xe8, 0x94, 0xec, 0x28, 0x89, 0xf3, 0x2b, 0xcb, 0xb2, 0xcd, 0x58, 0xb6,
	0x14, 0xda, 0xe5, 0xaa, 0xac, 0x50, 0xe0, 0x00, 0x43, 0xa2, 0x3c, 0x7f, 0x06, 0xc0, 0xb1, 0xb4,
	0xc9, 0x1d, 0xb2, 0xc9, 0x01, 0xb2, 0xc8, 0x3e, 0x57, 0xc8, 0x09, 0xb2, 0xcc, 0x22, 0x87, 0x49,
	0xa1, 0x1b, 0x33, 0xa4, 0x44, 0x65, 0x9d, 0x6c, 0x4c, 0xf7, 0xf7, 0x7d, 0xc0, 0x00, 0xdd, 0x8d,
	0xee, 0x16, 0xa1, 0x2a, 0x2f, 0x9e, 0x31, 0xad, 0x46, 0x4c, 0x8a, 0x84, 0xe9, 0x6c, 0x6a, 0x65,
	0x3f, 0xd7, 0x99, 0xcd, 0x82, 0x3f, 0x85, 0xca, 0x84, 0x19, 0x53, 0x99, 0x61, 0x97, 0x9a, 0xa9,
	0x1c, 0x44, 0xa0, 0xce, 0x72, 0xa9, 0xfb, 0xce, 0x32, 0x56, 0x8c, 0xae, 0xfa, 0x85, 0x8e, 0x8c,
	0xfb, 0xa7, 0xcf, 0x23, 0xd3, 0xe7, 0x51, 0xdf, 0xb8, 0x5f, 0xc3, 0xa3, 0xbe, 0x5f, 0x03, 0x9b,
	0x32, 0xcb, 0x47, 0xb1, 0x64, 0x29, 0x4f, 0xa4, 0xf9, 0x6f, 0x44, 0x1f, 0x00, 0x83, 0x3f, 0x7b,
	0x7f, 0xbd, 0x43, 0xb6, 0x17, 0x0f, 0xc7, 0xde, 0x9e, 0xfe, 0xf1, 0x43, 0xb0, 0x43, 0x56, 0x0a,
	0x1d, 0xc1, 0x22, 0x5a, 0xeb, 0xd6, 0x7a, 0xcd, 0x61, 0xa3, 0xd0, 0xd1, 0x7b, 0x9e, 0xc8, 0x60,
	0x9b, 0x34, 0xb8, 0x67, 0xee, 0x00, 0x53, 0xe7, 0x48, 0xec, 0x90, 0x15, 0x53, 0x32, 0x4b, 0xb8,
	0xc6, 0x78, 0xaa, 0x47, 0xd6, 0x6f, 0x9e, 0x85, 0xde, 0x05, 0xc9, 0x1a, 0xe0, 0x1f, 0x1d, 0x0c,
	0x4a, 0x4a, 0x1a, 0x5c, 0x08, 0x2d, 0x8d, 0xa1, 0xcb, 0xb8, 0x87, 0x37, 0x83, 0x47, 0xa4, 0x9d,
	0x6b, 0x19, 0xa9, 0x4b, 0x16, 0xcb, 0x74, 0x6c, 0x27, 0xb4, 0xde, 0xad, 0xf5, 0xda, 0xc3, 0x55,
	0x04, 0xcf, 0x00, 0x73, 0x1f, 0x4a, 0xe5, 0xa5, 0x65, 0x93, 0x2c, 0x67, 0xe5, 0x3e, 0x0d, 0xfc,
	0x90, 0xc3, 0xdf, 0x64, 0xf9, 0xb1, 0xdf, 0xee, 0x31, 0x59, 0x53, 0xa9, 0x95, 0x3a, 0xe2, 0xa1,
	0x3f, 0xd0, 0x0a, 0xe8, 0xda, 0x15, 0xea, 0xce, 0xb3, 0xf7, 0xcf, 0x26, 0x09, 0x16, 0x9d, 0x14,
	0x6c, 0x91, 0x3a, 0x7e, 0x97, 0x1e, 0xa2, 0x0f, 0xd0, 0x5a, 0x3c, 0xe4, 0xd3, 0x5b, 0x0e, 0xf9,
	0x88, 0xb4, 0xd1, 0x1b, 0x85, 0xd4, 0x46, 0x65, 0x29, 0x7d, 0x86, 0x22, 0x00, 0x3f, 0x21, 0x16,
	0x3c, 0x20, 0x2d, 0x48, 0x93, 0x30, 0x8b, 0x99, 0x12, 0xf4, 0xa7, 0x20, 0x21, 0x25, 0x34, 0x10,
	0xf8, 0x29, 0x2f, 0x80, 0xf3, 0x1f, 0xc1, 0x49, 0x56, 0x4b, 0x10, 0xdc, 0xf9, 0x2d, 0x59, 0x51,
	0xa9, 0xb1, 0x3c, 0x0d, 0x25, 0xfd, 0x19, 0xf0, 0x95, 0x1d, 0xec, 0x92, 0x66, 0x18, 0x2b, 0x99,
	0x5a, 0xb7, 0xff, 0xcf, 0x61, 0xff, 0x15, 0x04, 0x06, 0x22, 0xf8, 0x9e, 0x10, 0x1f, 0xb1, 0xab,
	0x5c, 0xd2, 0x5f, 0x00, 0xdb, 0xc4, 0x58, 0x5d, 0xe5, 0xb0, 0x6f, 0xae, 0x55, 0xa6, 0x95, 0xbd,
	0xa2, 0xcf, 0x71, 0x69, 0x69, 0x43, 0x1e, 0x14, 0x02, 0x17, 0xfe, 0x12, 0xb8, 0x86, 0x29, 0x04,
	0x2c, 0xdb, 0x24, 0xcb, 0x51, 0xcc, 0xc7, 0x86, 0xfe, 0x0a, 0x70, 0x34, 0x5c, 0x28, 0xe4, 0xa5,
	0x95, 0xa9, 0x90, 0x82, 0x21, 0xfd, 0xeb, 0x6e, 0xad, 0x77, 0x77, 0xd8, 0x2e, 0xd1, 0x57, 0x20,
	0x5b, 0x27, 0x4b, 0x96, 0x8f, 0xe9, 0x6f, 0x60, 0xa9, 0xfb, 0xaf, 0x3b, 0x85, 0x50, 0xfe, 0x76,
	0xbf, 0xc5, 0x53, 0x94, 0x76, 0x70, 0x40, 0x02, 0xa1, 0xbc, 0x83, 0x59, 0xa5, 0xfa, 0x1d, 0xa8,
	0x3a, 0x15, 0xf3, 0xb2, 0x94, 0x6f, 0x91, 0x7a, 0x22, 0xad, 0x56, 0x21, 0x3d, 0x06, 0x89, 0xb7,
	0x20, 0x0c, 0xdc, 0x4e, 0x0c, 0x0b, 0xb3, 0x69, 0x6a, 0xe9, 0x0b, 0x1f, 0x06, 0x07, 0x9d, 0x38,
	0xc4, 0x7d, 0x87, 0x5b, 0xab, 0xd5, 0xc8, 0x39, 0x4b, 0x09, 0x99, 0x5a, 0xe7, 0x93, 0x13, 0xfc,
	0x4e, 0xc5, 0x0c, 0x3c, 0xe1, 0xa2, 0x66, 0x35, 0x8f, 0x22, 0x15, 0x32, 0x95, 0x0a, 0x79, 0x49,
	0x5f, 0x62, 0xec, 0x3d, 0x38, 0x70, 0x58, 0xb0, 0x5f, 0x3e, 0x97, 0x5c, 0xcb, 0x50, 0x0a, 0xe9,
	0x4e, 0x7e, 0x0a, 0xba, 0x7b, 0x80, 0x5f, 0x54, 0xb0, 0x0b, 0xe2, 0x97, 0xcc, 0xb0, 0xb1, 0xce,
	0xa6, 0x39, 0x7d, 0x85, 0x3e, 0xf8, 0x92, 0x99, 0xd7, 0xce, 0x76, 0x91, 0x88, 0xe2, 0xec, 0x2b,
	0x73, 0x6e, 0x7b, 0x8d, 0x91, 0x70, 0xf6, 0x47, 0x3e, 0x76, 0xeb, 0xa2, 0xaf, 0x82, 0x85, 0x31,
	0x37, 0x86, 0xbe, 0xc1, 0x75, 0xd1, 0x57, 0x71, 0xe2, 0x6c, 0x47, 0xe6, 0x2a, 0xf4, 0x57, 0x1e,
	0xf8, 0xf0, 0xaa, 0x10, 0x2f, 0xbc, 0x45, 0xea, 0x3c, 0xb4, 0xaa, 0x90, 0xf4, 0xf7, 0xdd, 0x5a,
	0x6f, 0x65, 0xe8, 0xad, 0xe0, 0x3b, 0xd2, 0xac, 0xdc, 0x4a, 0xdf, 0x02, 0x35, 0x03, 0x82, 0x9f,
	0x90, 0xcd, 0x59, 0x38, 0x20, 0x45, 0x31, 0x69, 0xcf, 0x20, 0x29, 0x67, 0xa1, 0xba, 0x70, 0x14,
	0xa4, 0xee, 0x2e, 0xc1, 0x7c, 0x63, 0x7c, 0x2c, 0xe9, 0x3b, 0x3c, 0x04, 0x00, 0xc7, 0x63, 0xe9,
	0xc2, 0x82, 0x64, 0xcc, 0x47, 0x32, 0xa6, 0xef, 0x31, 0x2c, 0x00, 0x9d, 0x39, 0xc4, 0xd5, 0x91,
	0xf2, 0x2c, 0xe7, 0x78, 0xf3, 0x62, 0xf6, 0xb0, 0xec, 0x28, 0xae, 0xde, 0xde, 0x05, 0xa4, 0x1a,
	0xb1, 0xa3, 0xb8, 0x7c, 0x79, 0x3f, 0x22, 0x1d, 0xdc, 0x3b, 0xc9, 0x84, 0x8a, 0xae, 0x98, 0x55,
	0x89, 0xa4, 0x7f, 0x00, 0x19, 0xba, 0xff, 0x1d, 0xe0, 0x1f, 0x55, 0x22, 0x83, 0xbf, 0xd7, 0xca,
	0x77, 0xe2, 0x52, 0x82, 0x0e, 0xbb, 0xb5, 0x5e, 0xeb, 0xf0, 0xcf, 0xb5, 0xfe, 0xff, 0xb6, 0xb4,
	0xf7, 0xaf, 0x55, 0x2c, 0x77, 0x32, 0xff, 0x76, 0x2f, 0xb8, 0x9d, 0xec, 0xfd, 0xab, 0x46, 0x3a,
	0x0b, 0x82, 0xe0, 0x1f, 0xb7, 0xa1, 0xb4, 0xd6, 0x5d, 0xea, 0xb5, 0x0e, 0xff, 0xf2, 0xff, 0x77,
	0x21, 0xa6, 0xac, 0x4c, 0x86, 0x6b, 0x0e, 0x1f, 0xaa, 0xd1, 0xa9, 0x48, 0xe0, 0x6a, 0x7f, 0x6b,
	0x91, 0xad, 0xdb, 0xa5, 0xf3, 0x8d, 0xa5, 0x76, 0xbd, 0xb1, 0x1c, 0x90, 0x40, 0xa5, 0x51, 0xa6,
	0x13, 0x6e, 0x5d, 0x72, 0x9a, 0x6c, 0xaa, 0xc3, 0xb2, 0xb7, 0x75, 0xe6, 0x98, 0x0f, 0x40, 0xb8,
	0xca, 0x58, 0x1c, 0x31, 0xd7, 0x4d, 0x26, 0x59, 0xee, 0x1b, 0x5d, 0xb3, 0x38, 0x7a, 0x8f, 0x40,
	0x70, 0x48, 0xbe, 0x29, 0x8e, 0xd8, 0x2d, 0x1b, 0x62, 0xbf, 0xdb, 0x28, 0x8e, 0x06, 0x0b, 0x5b,
	0x2e, 0xf6, 0xa2, 0xe5, 0x5b, 0x7a, 0xd1, 0x5c, 0x8d, 0xaa, 0xdf, 0xac, 0x51, 0x71, 0xc6, 0x05,
	0xf3, 0x64, 0x03, 0x1f, 0x83, 0x83, 0xde, 0xa1, 0x80, 0x92, 0x06, 0xd4, 0xd5, 0xa3, 0x67, 0xd0,
	0xe4, 0xee, 0x0e, 0x4b, 0x73, 0x56, 0x90, 0x9b, 0xf3, 0x05, 0x19, 0x5a, 0x8b, 0x2a, 0xb8, 0x95,
	0xbe, 0x1e, 0x93, 0xb2, 0x8b, 0x01, 0x88, 0xe5, 0x78, 0x8b, 0xd4, 0xe3, 0x2c, 0xcb, 0xa5, 0xa0,
	0x2d, 0xac, 0x03, 0x68, 0x05, 0xfb, 0xa4, 0x53, 0xb5, 0x60, 0x0c, 0xa7, 0x12, 0x74, 0x15, 0x36,
	0x28, 0x7b, 0x30, 0xb4, 0xfb, 0xc1, 0x75, 0x69, 0x35, 0x6e, 0xb4, 0xaf, 0xb5, 0xeb, 0x4f, 0x7e,
	0xea, 0x38, 0x20, 0x1b, 0x37, 0x76, 0x05, 0xf1, 0x1a, 0x88, 0xd7, 0xe7, 0xf7, 0x05, 0x79, 0x97,
	0xac, 0xce, 0xe6, 0x80, 0x48, 0xd1, 0x7b, 0xe8, 0x93, 0x72, 0x06, 0x88, 0x54, 0xb0, 0x47, 0xda,
	0x95, 0xc2, 0x38, 0xc9, 0x3a, 0x48, 0x5a, 0x5e, 0xf2, 0x81, 0x47, 0xea, 0x66, 0x95, 0xe9, 0x2c,
	0x54, 0x99, 0x5d, 0xd2, 0xb4, 0xd3, 0x34, 0x95, 0xd0, 0xa2, 0x03, 0xac, 0x51, 0x08, 0x0c, 0x04,
	0xcc, 0x08, 0xdc, 0x4e, 0x94, 0xa0, 0x1b, 0x18, 0x2e, 0xb4, 0x9c, 0x77, 0x47, 0x3c, 0xfc, 0x3c,
	0xcd, 0x99, 0xa7, 0x37, 0xd1, 0xbb, 0x08, 0x5e, 0xa0, 0x68, 0x9f, 0x74, 0xb4, 0x8c, 0x58, 0x98,
	0x5a, 0x96, 0x45, 0x0c, 0x29, 0xfa, 0x0d, 0x7a, 0x51, 0xcb, 0xe8, 0x24, 0xb5, 0xe7, 0xd1, 0x0b,
	0x40, 0x83, 0x13, 0x72, 0x3f, 0x9d, 0x26, 0x23, 0xa9, 0x9d, 0xb2, 0x6a, 0xa4, 0x61, 0x96, 0x24,
	0xd3, 0x54, 0x59, 0x25, 0x0d, 0xdd, 0x82, 0x75, 0xbb, 0xa8, 0x3a, 0x8f, 0x4e, 0xbd, 0xe6, 0x64,
	0x26, 0x09, 0x1e, 0x92, 0xd5, 0xa4, 0xc8, 0x5d, 0x69, 0x96, 0x46, 0xa6, 0x96, 0x6e, 0x43, 0x4c,
	0x5b, 0x0e, 0xbb, 0x40, 0xc8, 0x65, 0xa9, 0x3b, 0xb0, 0xb6, 0x95, 0x88, 0x82, 0xa8, 0x8d, 0x68,
	0x29, 0x7b, 0x42, 0x36, 0x0a, 0x1d, 0xa9, 0x24, 0xcf, 0xb4, 0x9d, 0xd3, 0xee, 0x80, 0x36, 0x98,
	0xa3, 0xca, 0x05, 0x07, 0x24, 0xc0, 0x27, 0xc2, 0xcd, 0x9c, 0xfe, 0x5b, 0xd0, 0x77, 0x66, 0x4c,
	0x29, 0xdf, 0x27, 0xeb, 0x08, 0x6a, 0x51, 0x89, 0x77, 0x41, 0x7c, 0xaf, 0xc4, 0x4b, 0xe9, 0x73,
	0xb2, 0x63, 0xe4, 0x38, 0x91, 0xa9, 0x95, 0xa2, 0x7c, 0xb1, 0xd5, 0x9a, 0xef, 0x60, 0xcd, 0x76,
	0x25, 0xf0, 0x0f, 0xb8, 0x5c, 0x7b, 0x9f, 0xb4, 0xaa, 0xfc, 0x50, 0x82, 0x7e, 0x8f, 0x13, 0x90,
	0xcf, 0x8e, 0x81, 0x08, 0x9e, 0x90, 0xcd, 0x39, 0x9e, 0x69, 0x19, 0x61, 0xbb, 0xbc, 0x8f, 0x9d,
	0xbf, 0x12, 0x0e, 0x3d, 0xe1, 0x52, 0x32, 0x33, 0x79, 0xc4, 0xb8, 0x96, 0xdc, 0xed, 0xf8, 0x00,
	0x52, 0x97, 0x38, 0xec, 0x58, 0x4b, 0x3e, 0x10, 0xc1, 0x8f, 0x49, 0xa0, 0x65, 0x92, 0x59, 0xe9,
	0xe3, 0x0d, 0x13, 0x2c, 0xed, 0x76, 0x97, 0x7a, 0xab, 0xc3, 0x75, 0x64, 0x30, 0xe4, 0x6e, 0x86,
	0x75, 0x11, 0x9b, 0x70, 0x83, 0xa9, 0x69, 0xec, 0x67, 0xfa, 0x10, 0x23, 0x36, 0xe1, 0xe6, 0xcc,
	0x43, 0xae, 0x54, 0xa5, 0xd3, 0xc4, 0x4b, 0xe8, 0x9e, 0xbf, 0xc2, 0x34, 0x41, 0x81, 0x1b, 0x9f,
	0xaa, 0xd5, 0x8f, 0xba, 0x4b, 0x2e, 0x79, 0x4b, 0x1b, 0x92, 0x54, 0xa5, 0x42, 0xa5, 0x63, 0x9f,
	0xfc, 0x3f, 0xf0, 0x49, 0x8a, 0x60, 0x95, 0xfe, 0xe9, 0x44, 0x09, 0x16, 0x49, 0x25, 0xe8, 0x63,
	0xa8, 0x2c, 0x2b, 0x0e, 0x78, 0x25, 0x95, 0x70, 0x64, 0x92, 0xc7, 0x06, 0xc9, 0x1f, 0x22, 0xe9,
	0x00, 0x47, 0xee, 0xfd, 0xbb, 0x46, 0x36, 0xca, 0x1a, 0x1d, 0x67, 0x21, 0x8f, 0xf1, 0x2b, 0x8b,
	0x43, 0x6d, 0xed, 0x96, 0xa1, 0xf6, 0xda, 0xe0, 0x7a, 0xe7, 0xc6, 0xe0, 0xba, 0x49, 0x96, 0x8d,
	0xe5, 0x31, 0xfe, 0x09, 0xd2, 0x1e, 0xa2, 0xe1, 0xae, 0x9a, 0x28, 0xad, 0x33, 0x2d, 0x05, 0x14,
	0xe2, 0xf6, 0xb0, 0xb2, 0xe1, 0x95, 0x4b, 0x37, 0x21, 0xb2, 0x2c, 0x8d, 0xaf, 0xa0, 0xf4, 0xba,
	0x57, 0x0e, 0xd0, 0x79, 0x1a, 0x5f, 0xb9, 0x2d, 0xd1, 0x07, 0x58, 0x76, 0xd1, 0xb8, 0x36, 0x7c,
	0x36, 0xae, 0x0f, 0x9f, 0xa3, 0x3a, 0x9c, 0xf7, 0xe9, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x46,
	0x5f, 0x8c, 0xdd, 0xfd, 0x0d, 0x00, 0x00,
}
