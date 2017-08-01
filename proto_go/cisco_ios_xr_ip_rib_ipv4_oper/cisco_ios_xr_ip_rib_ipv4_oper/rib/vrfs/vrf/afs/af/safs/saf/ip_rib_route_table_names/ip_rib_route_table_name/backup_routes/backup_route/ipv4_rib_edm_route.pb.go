// Code generated by protoc-gen-go.
// source: ipv4_rib_edm_route.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_backup_routes_backup_route is a generated protocol buffer package.

It is generated from these files:
	ipv4_rib_edm_route.proto

It has these top-level messages:
	Ipv4RibEdmRoute_KEYS
	Ipv4RibEdmRoute
	Ipv4RibEdmPath
	Ipv4RibEdmPathItem
	RibEdmLocalLabel
*/
package cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_backup_routes_backup_route

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
	Protoid        uint32 `protobuf:"varint,7,opt,name=protoid" json:"protoid,omitempty"`
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

func (m *Ipv4RibEdmRoute_KEYS) GetProtoid() uint32 {
	if m != nil {
		return m.Protoid
	}
	return 0
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
	proto.RegisterType((*Ipv4RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.backup_routes.backup_route.ipv4_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv4RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.backup_routes.backup_route.ipv4_rib_edm_route")
	proto.RegisterType((*Ipv4RibEdmPath)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.backup_routes.backup_route.ipv4_rib_edm_path")
	proto.RegisterType((*Ipv4RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.backup_routes.backup_route.ipv4_rib_edm_path_item")
	proto.RegisterType((*RibEdmLocalLabel)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.backup_routes.backup_route.rib_edm_local_label")
}

func init() { proto.RegisterFile("ipv4_rib_edm_route.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0xcd, 0x72, 0x1b, 0xb9,
	0x11, 0x2e, 0x5a, 0x36, 0x25, 0x82, 0xa2, 0x2c, 0x8e, 0x64, 0x09, 0xb2, 0xfc, 0x43, 0xcb, 0x71,
	0x8a, 0x4a, 0x45, 0x74, 0x4a, 0x76, 0x94, 0xc4, 0xf9, 0x95, 0x65, 0xd9, 0x66, 0x2c, 0x5b, 0x0a,
	0xed, 0x72, 0x55, 0x4e, 0x28, 0x70, 0x80, 0x21, 0x51, 0x9e, 0x3f, 0x03, 0xe0, 0x58, 0x7a, 0x81,
	0xbc, 0x43, 0x2e, 0x7b, 0xdc, 0x77, 0xd8, 0x57, 0xd8, 0x27, 0xd8, 0xc3, 0x1e, 0xf6, 0xbe, 0x2f,
	0xb1, 0x85, 0x6e, 0xcc, 0x90, 0xfa, 0xd9, 0xbb, 0x2f, 0x92, 0xfa, 0xfb, 0x3e, 0xf4, 0x00, 0xdd,
	0x8d, 0x6e, 0x88, 0x50, 0x95, 0x17, 0x4f, 0x99, 0x56, 0x43, 0x26, 0x45, 0xc2, 0x74, 0x36, 0xb1,
	0xb2, 0x97, 0xeb, 0xcc, 0x66, 0xc1, 0xff, 0x6a, 0xa1, 0x32, 0x61, 0xc6, 0x54, 0x66, 0xd8, 0xa9,
	0x66, 0x2a, 0x07, 0x15, 0xc8, 0xb3, 0x5c, 0xea, 0x9e, 0x56, 0xc3, 0x5e, 0xa1, 0x23, 0xe3, 0x7e,
	0xf4, 0x78, 0x64, 0x7a, 0x3c, 0xea, 0x19, 0xf7, 0xdb, 0xf0, 0xa8, 0xe7, 0xd5, 0xe0, 0x8f, 0x59,
	0x3e, 0x8c, 0x25, 0x4b, 0x79, 0x22, 0xcd, 0xaf, 0x11, 0xbd, 0x21, 0x0f, 0x3f, 0x4d, 0x72, 0xc4,
	0xcd, 0x39, 0x6b, 0xeb, 0xe7, 0x1a, 0x59, 0xbf, 0xbc, 0x4b, 0xf6, 0xe6, 0xf0, 0xbf, 0xef, 0x83,
	0x0d, 0xb2, 0x50, 0xe8, 0x08, 0x5c, 0xd0, 0x5a, 0xa7, 0xd6, 0x6d, 0x0c, 0xe6, 0x0b, 0x1d, 0xbd,
	0xe3, 0x89, 0x0c, 0xd6, 0xc9, 0x3c, 0xf7, 0xcc, 0x35, 0x60, 0xea, 0x1c, 0x89, 0x0d, 0xb2, 0x60,
	0x4a, 0x66, 0x0e, 0xd7, 0x18, 0x4f, 0x75, 0xc9, 0xf2, 0xc5, 0x9d, 0xd1, 0xeb, 0x20, 0x59, 0x02,
	0xfc, 0x83, 0x83, 0x41, 0x49, 0xc9, 0x3c, 0x17, 0x42, 0x4b, 0x63, 0xe8, 0x0d, 0xf4, 0xe1, 0xcd,
	0xe0, 0x21, 0x69, 0xe5, 0x5a, 0x46, 0xea, 0x94, 0xc5, 0x32, 0x1d, 0xd9, 0x31, 0xad, 0x77, 0x6a,
	0xdd, 0xd6, 0x60, 0x11, 0xc1, 0x23, 0xc0, 0xdc, 0x72, 0x88, 0xb2, 0x12, 0x74, 0x1e, 0xe8, 0xd2,
	0xdc, 0xfa, 0xa1, 0x41, 0x82, 0xcb, 0xa7, 0x0d, 0xd6, 0x48, 0x1d, 0x1d, 0xd0, 0x5d, 0x3c, 0x0c,
	0x5a, 0x97, 0xbf, 0xf6, 0xe4, 0x8a, 0xaf, 0x3d, 0x24, 0x2d, 0x3c, 0x56, 0x21, 0xb5, 0x51, 0x59,
	0x4a, 0x9f, 0xa2, 0x08, 0xc0, 0x8f, 0x88, 0x05, 0xf7, 0x49, 0x13, 0xf6, 0x10, 0x66, 0x31, 0x53,
	0x82, 0xfe, 0x11, 0x24, 0xa4, 0x84, 0xfa, 0x02, 0x3f, 0xe5, 0x05, 0x10, 0x99, 0x3d, 0xd8, 0xc9,
	0x62, 0x09, 0x42, 0x5c, 0x6e, 0x93, 0x05, 0x95, 0x1a, 0xcb, 0xd3, 0x50, 0xd2, 0x3f, 0x01, 0x5f,
	0xd9, 0xc1, 0x26, 0x69, 0x84, 0xb1, 0x92, 0xa9, 0x75, 0xfe, 0xff, 0x0c, 0xfe, 0x17, 0x10, 0xe8,
	0x8b, 0xe0, 0x2e, 0x21, 0x3e, 0xf4, 0x67, 0xb9, 0xa4, 0x7f, 0x01, 0xb6, 0x81, 0x41, 0x3f, 0xcb,
	0xc1, 0x6f, 0xae, 0x55, 0xa6, 0x95, 0x3d, 0xa3, 0xcf, 0x70, 0x69, 0x69, 0x43, 0x42, 0x0b, 0x81,
	0x0b, 0xff, 0x8a, 0xd1, 0x34, 0x85, 0x80, 0x65, 0xab, 0xe4, 0x46, 0x14, 0xf3, 0x91, 0xa1, 0x7f,
	0x03, 0x1c, 0x8d, 0xe0, 0x11, 0x59, 0x92, 0xa7, 0x56, 0xa6, 0x42, 0x0a, 0x86, 0xf4, 0xdf, 0x3b,
	0xb5, 0xee, 0xf5, 0x41, 0xab, 0x44, 0x5f, 0x82, 0x6c, 0x99, 0xcc, 0x59, 0x3e, 0xa2, 0xff, 0x80,
	0xa5, 0xee, 0x4f, 0xb7, 0x0b, 0xa1, 0xfc, 0xe9, 0xfe, 0x89, 0xbb, 0x28, 0xed, 0x60, 0x87, 0x04,
	0x42, 0xf9, 0x00, 0xb3, 0x4a, 0xf5, 0x2f, 0x50, 0xb5, 0x2b, 0xe6, 0x45, 0x29, 0x5f, 0x23, 0xf5,
	0x44, 0x5a, 0xad, 0x42, 0xba, 0x0f, 0x12, 0x6f, 0x41, 0x1a, 0xb8, 0x1d, 0x1b, 0x16, 0x66, 0x93,
	0xd4, 0xd2, 0xe7, 0x3e, 0x0d, 0x0e, 0x3a, 0x70, 0x88, 0xfb, 0x0e, 0xb7, 0x56, 0xab, 0xa1, 0x0b,
	0x96, 0x12, 0x32, 0xb5, 0x2e, 0x26, 0x07, 0xf8, 0x9d, 0x8a, 0xe9, 0x7b, 0xc2, 0x65, 0xcd, 0x6a,
	0x1e, 0x45, 0x2a, 0x64, 0x2a, 0x15, 0xf2, 0x94, 0xbe, 0xc0, 0xdc, 0x7b, 0xb0, 0xef, 0xb0, 0x60,
	0xbb, 0xac, 0xfb, 0x5c, 0xcb, 0x50, 0x0a, 0xe9, 0x76, 0x7e, 0x08, 0xba, 0x9b, 0x80, 0x9f, 0x54,
	0xb0, 0x4b, 0xe2, 0xe7, 0xcc, 0xb0, 0x91, 0xce, 0x26, 0x39, 0x7d, 0x89, 0x31, 0xf8, 0x9c, 0x99,
	0x57, 0xce, 0x76, 0x99, 0x88, 0xe2, 0xec, 0x0b, 0x73, 0x61, 0x7b, 0x85, 0x99, 0x70, 0xf6, 0x07,
	0x3e, 0x72, 0xeb, 0xa2, 0x2f, 0x82, 0x85, 0x31, 0x37, 0x86, 0xbe, 0xc6, 0x75, 0xd1, 0x17, 0x71,
	0xe0, 0x6c, 0x47, 0xe6, 0x2a, 0xf4, 0x47, 0xee, 0xfb, 0xf4, 0xaa, 0x10, 0x0f, 0xbc, 0x46, 0xea,
	0x3c, 0xb4, 0xaa, 0x90, 0xf4, 0xdf, 0x9d, 0x5a, 0x77, 0x61, 0xe0, 0xad, 0xe0, 0x0e, 0x69, 0x54,
	0x61, 0xa5, 0x6f, 0x80, 0x9a, 0x02, 0xc1, 0x1f, 0xc8, 0xea, 0x34, 0x1d, 0x50, 0xa2, 0x58, 0xb4,
	0x47, 0x50, 0x94, 0xd3, 0x54, 0x9d, 0x38, 0x0a, 0x4a, 0x77, 0x93, 0x60, 0xbd, 0x31, 0x3e, 0x92,
	0xf4, 0x2d, 0x6e, 0x02, 0x80, 0xfd, 0x91, 0x74, 0x69, 0x41, 0x32, 0xe6, 0x43, 0x19, 0xd3, 0x77,
	0x98, 0x16, 0x80, 0x8e, 0x1c, 0xe2, 0x6e, 0x74, 0xb9, 0x97, 0x63, 0x3c, 0x79, 0x31, 0xbd, 0x58,
	0x76, 0x18, 0x57, 0x77, 0xef, 0x04, 0x4a, 0x8d, 0xd8, 0x61, 0x5c, 0xde, 0xbc, 0xdf, 0x91, 0x36,
	0xfa, 0x4e, 0x32, 0xa1, 0xa2, 0x33, 0x66, 0x55, 0x22, 0xe9, 0x7f, 0x40, 0x86, 0xe1, 0x7f, 0x0b,
	0xf8, 0x07, 0x95, 0xc8, 0xe0, 0xbb, 0x5a, 0x79, 0x4f, 0x5c, 0x49, 0xd0, 0x41, 0xa7, 0xd6, 0x6d,
	0xee, 0xfe, 0xbf, 0xd6, 0xfb, 0x3a, 0x7a, 0x75, 0xef, 0x5c, 0xe7, 0x72, 0x3b, 0xf4, 0x77, 0xf8,
	0x84, 0xdb, 0xf1, 0xd6, 0x4f, 0x35, 0xd2, 0xbe, 0x24, 0x08, 0xbe, 0xbf, 0x0a, 0xa5, 0xb5, 0xce,
	0x5c, 0xb7, 0xb9, 0xfb, 0xcd, 0xd7, 0x7b, 0x30, 0xa6, 0xac, 0x4c, 0x06, 0x4b, 0x0e, 0x1f, 0xa8,
	0xe1, 0xa1, 0x48, 0xe0, 0x88, 0xdf, 0x36, 0xc9, 0xda, 0xd5, 0xd2, 0xd9, 0x89, 0x51, 0x3b, 0x3f,
	0x31, 0x76, 0x48, 0xa0, 0xd2, 0x28, 0xd3, 0x09, 0xb7, 0xae, 0x58, 0x4d, 0x36, 0xd1, 0x61, 0x39,
	0xb4, 0xda, 0x33, 0xcc, 0x7b, 0x20, 0x5c, 0xa7, 0x2c, 0xf6, 0x58, 0x2a, 0x4f, 0xed, 0x38, 0xcb,
	0xfd, 0x04, 0x6b, 0x14, 0x7b, 0xef, 0x10, 0x08, 0x76, 0xc9, 0xad, 0x62, 0x8f, 0x5d, 0xe1, 0x10,
	0x07, 0xd9, 0x4a, 0xb1, 0xd7, 0xbf, 0xe4, 0xf2, 0x11, 0x59, 0x52, 0xa9, 0x95, 0x3a, 0xe2, 0xa1,
	0x9f, 0x7a, 0x38, 0xd4, 0x5a, 0x15, 0x0a, 0x37, 0x64, 0xda, 0xb3, 0xea, 0x17, 0x7b, 0x56, 0x9c,
	0x71, 0xc1, 0x3c, 0x89, 0x13, 0x8d, 0x38, 0xe8, 0x2d, 0x0a, 0x28, 0x99, 0x87, 0x3e, 0xbb, 0xf7,
	0x94, 0x2e, 0x40, 0x5d, 0x97, 0xe6, 0xb4, 0x41, 0x37, 0x66, 0x1b, 0x34, 0x8c, 0x1a, 0x55, 0x70,
	0x2b, 0x7d, 0x7f, 0x26, 0xe5, 0x54, 0x03, 0x10, 0xdb, 0xf3, 0x1a, 0xa9, 0xc7, 0x59, 0x96, 0x4b,
	0x41, 0x9b, 0xd8, 0x17, 0xd0, 0x0a, 0xb6, 0x49, 0xdb, 0x05, 0x87, 0x8d, 0xb3, 0xdc, 0x27, 0x57,
	0x09, 0xba, 0x08, 0x0e, 0x96, 0x1c, 0xf1, 0x3a, 0xcb, 0x61, 0x8e, 0xf7, 0xcf, 0x4b, 0xab, 0x77,
	0x44, 0x0b, 0x07, 0xbe, 0x97, 0x7e, 0xf4, 0xcf, 0x89, 0x1d, 0xb2, 0x72, 0xc1, 0x2b, 0x88, 0x97,
	0x40, 0xbc, 0x3c, 0xeb, 0x17, 0xe4, 0x1d, 0xb2, 0x58, 0xc9, 0x79, 0xa4, 0xe8, 0x4d, 0x8c, 0x89,
	0xd7, 0xed, 0x47, 0x2a, 0xd8, 0x22, 0xad, 0x4a, 0x61, 0x9c, 0x64, 0x19, 0x24, 0x4d, 0x2f, 0x79,
	0xcf, 0x23, 0x75, 0xb1, 0xeb, 0xb4, 0x2f, 0x75, 0x9d, 0x4d, 0xd2, 0xb0, 0x93, 0x34, 0x95, 0x30,
	0xb2, 0x03, 0xec, 0x59, 0x08, 0xf4, 0x05, 0xbc, 0x19, 0xb8, 0x1d, 0x2b, 0x41, 0x57, 0x30, 0x5d,
	0x68, 0xb9, 0xe8, 0xfa, 0xda, 0xf6, 0xf4, 0x2a, 0x46, 0x17, 0xc1, 0x13, 0x14, 0x6d, 0x93, 0xb6,
	0x96, 0x11, 0x0b, 0x53, 0xcb, 0xb2, 0x88, 0x21, 0x45, 0x6f, 0x61, 0x14, 0xb5, 0x8c, 0x0e, 0x52,
	0x7b, 0x1c, 0x3d, 0x07, 0x34, 0x38, 0x20, 0xf7, 0xd2, 0x49, 0x32, 0x94, 0xda, 0x29, 0xab, 0xc1,
	0x1a, 0x66, 0x49, 0x32, 0x49, 0x95, 0x55, 0xd2, 0xd0, 0x35, 0x58, 0xb7, 0x89, 0xaa, 0xe3, 0xe8,
	0xd0, 0x6b, 0x0e, 0xa6, 0x92, 0xe0, 0x01, 0x59, 0x4c, 0x8a, 0xdc, 0xb5, 0x6a, 0x69, 0x64, 0x6a,
	0xe9, 0x3a, 0xe4, 0xb4, 0xe9, 0xb0, 0x13, 0x84, 0x5c, 0x95, 0xba, 0x0d, 0x6b, 0x5b, 0x89, 0x28,
	0x88, 0x5a, 0x88, 0x96, 0xb2, 0xc7, 0x64, 0xa5, 0xd0, 0x91, 0x4a, 0xf2, 0x4c, 0xdb, 0x19, 0xed,
	0x06, 0x68, 0x83, 0x19, 0xaa, 0x5c, 0xb0, 0x43, 0x02, 0xbc, 0x22, 0xdc, 0xcc, 0xe8, 0x6f, 0x83,
	0xbe, 0x3d, 0x65, 0x4a, 0xf9, 0x36, 0x59, 0x46, 0x50, 0x8b, 0x4a, 0xbc, 0x09, 0xe2, 0x9b, 0x25,
	0x5e, 0x4a, 0x9f, 0x91, 0x0d, 0x23, 0x47, 0x89, 0x4c, 0xad, 0x14, 0xe5, 0x8d, 0xad, 0xd6, 0xdc,
	0x81, 0x35, 0xeb, 0x95, 0xc0, 0x5f, 0xe0, 0x72, 0xed, 0x3d, 0xd2, 0xac, 0xea, 0x43, 0x09, 0x7a,
	0x17, 0x5f, 0x44, 0xbe, 0x3a, 0xfa, 0x22, 0x78, 0x4c, 0x56, 0x67, 0x78, 0xa6, 0x65, 0x84, 0xe3,
	0xf3, 0x1e, 0xbe, 0x04, 0x2a, 0xe1, 0xc0, 0x13, 0xae, 0x24, 0x33, 0x93, 0x47, 0x8c, 0x6b, 0xc9,
	0x9d, 0xc7, 0xfb, 0x50, 0xba, 0xc4, 0x61, 0xfb, 0x5a, 0xf2, 0xbe, 0x08, 0x7e, 0x4f, 0x02, 0x2d,
	0x93, 0xcc, 0x4a, 0x9f, 0x6f, 0xe6, 0x3a, 0x14, 0xed, 0x74, 0xe6, 0xba, 0x8b, 0x83, 0x65, 0x64,
	0x30, 0xe5, 0xfb, 0x42, 0x68, 0x97, 0xb1, 0x31, 0x37, 0x58, 0x9a, 0xc6, 0x7e, 0xa2, 0x0f, 0x30,
	0x63, 0x63, 0x6e, 0x8e, 0x3c, 0xe4, 0x5a, 0x55, 0x3a, 0x49, 0xbc, 0x84, 0x6e, 0xf9, 0x23, 0x4c,
	0x12, 0x14, 0xb8, 0xe7, 0x54, 0xb5, 0xfa, 0x61, 0x67, 0xce, 0x15, 0x6f, 0x69, 0x43, 0x91, 0xaa,
	0x54, 0xa8, 0x74, 0xe4, 0x8b, 0xff, 0x37, 0xbe, 0x48, 0x11, 0xac, 0xca, 0x3f, 0x1d, 0x2b, 0xc1,
	0x22, 0xa9, 0x04, 0x7d, 0x04, 0x9d, 0x65, 0xc1, 0x01, 0x2f, 0xa5, 0x12, 0x8e, 0x4c, 0xf2, 0xd8,
	0x20, 0xf9, 0x5b, 0x24, 0x1d, 0xe0, 0xc8, 0xad, 0x1f, 0x6b, 0x64, 0xa5, 0xec, 0xd1, 0x71, 0x16,
	0xf2, 0x18, 0xbf, 0x72, 0xf9, 0x91, 0x5b, 0xbb, 0xe2, 0x91, 0x7b, 0xee, 0x21, 0x7b, 0xed, 0xc2,
	0x43, 0x76, 0x95, 0xdc, 0x30, 0x96, 0xc7, 0xf8, 0xbf, 0x45, 0x6b, 0x80, 0x86, 0x3b, 0x6a, 0xa2,
	0xb4, 0xce, 0xb4, 0x14, 0xd0, 0x88, 0x5b, 0x83, 0xca, 0x86, 0x5b, 0x2e, 0xdd, 0x8b, 0x91, 0x65,
	0x69, 0x7c, 0x06, 0xad, 0xd7, 0xdd, 0x72, 0x80, 0x8e, 0xd3, 0xf8, 0xcc, 0xb9, 0xc4, 0x18, 0x60,
	0xdb, 0x45, 0xe3, 0xdc, 0x63, 0x74, 0xfe, 0xfc, 0x63, 0x74, 0x58, 0x87, 0xfd, 0x3e, 0xf9, 0x25,
	0x00, 0x00, 0xff, 0xff, 0xab, 0xaa, 0x04, 0xe5, 0xdf, 0x0d, 0x00, 0x00,
}
