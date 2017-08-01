// Code generated by protoc-gen-go.
// source: bgp_global_process_info_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_default_vrf_global_process_info is a generated protocol buffer package.

It is generated from these files:
	bgp_global_process_info_bag.proto

It has these top-level messages:
	BgpGlobalProcessInfoBag_KEYS
	BgpGlobalProcessInfoBag
	ClusterIdBag_
	ColorIdBag_
	BgpGlobalProcessInfoGbl_
	BgpGlobalProcessInfoVrf_
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_default_vrf_global_process_info

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

// BGP process information common to all BGP processes
type BgpGlobalProcessInfoBag_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
}

func (m *BgpGlobalProcessInfoBag_KEYS) Reset()                    { *m = BgpGlobalProcessInfoBag_KEYS{} }
func (m *BgpGlobalProcessInfoBag_KEYS) String() string            { return proto.CompactTextString(m) }
func (*BgpGlobalProcessInfoBag_KEYS) ProtoMessage()               {}
func (*BgpGlobalProcessInfoBag_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BgpGlobalProcessInfoBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

type BgpGlobalProcessInfoBag struct {
	// Name of the VRF
	VrfName string `protobuf:"bytes,50,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	// VRF ID
	Vrfid uint32 `protobuf:"varint,51,opt,name=vrfid" json:"vrfid,omitempty"`
	// Global information
	Global *BgpGlobalProcessInfoGbl_ `protobuf:"bytes,52,opt,name=global" json:"global,omitempty"`
	// VRF information
	Vrf *BgpGlobalProcessInfoVrf_ `protobuf:"bytes,53,opt,name=vrf" json:"vrf,omitempty"`
}

func (m *BgpGlobalProcessInfoBag) Reset()                    { *m = BgpGlobalProcessInfoBag{} }
func (m *BgpGlobalProcessInfoBag) String() string            { return proto.CompactTextString(m) }
func (*BgpGlobalProcessInfoBag) ProtoMessage()               {}
func (*BgpGlobalProcessInfoBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BgpGlobalProcessInfoBag) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpGlobalProcessInfoBag) GetVrfid() uint32 {
	if m != nil {
		return m.Vrfid
	}
	return 0
}

func (m *BgpGlobalProcessInfoBag) GetGlobal() *BgpGlobalProcessInfoGbl_ {
	if m != nil {
		return m.Global
	}
	return nil
}

func (m *BgpGlobalProcessInfoBag) GetVrf() *BgpGlobalProcessInfoVrf_ {
	if m != nil {
		return m.Vrf
	}
	return nil
}

type ClusterIdBag_ struct {
	// Cluster ID
	ClusterIdVal uint32 `protobuf:"varint,1,opt,name=cluster_id_val,json=clusterIdVal" json:"cluster_id_val,omitempty"`
	// Cluster ID type: number or IPv4 address
	ClusterIdType uint32 `protobuf:"varint,2,opt,name=cluster_id_type,json=clusterIdType" json:"cluster_id_type,omitempty"`
}

func (m *ClusterIdBag_) Reset()                    { *m = ClusterIdBag_{} }
func (m *ClusterIdBag_) String() string            { return proto.CompactTextString(m) }
func (*ClusterIdBag_) ProtoMessage()               {}
func (*ClusterIdBag_) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClusterIdBag_) GetClusterIdVal() uint32 {
	if m != nil {
		return m.ClusterIdVal
	}
	return 0
}

func (m *ClusterIdBag_) GetClusterIdType() uint32 {
	if m != nil {
		return m.ClusterIdType
	}
	return 0
}

type ColorIdBag_ struct {
	// Color ID
	ColorIdVal uint32 `protobuf:"varint,1,opt,name=color_id_val,json=colorIdVal" json:"color_id_val,omitempty"`
}

func (m *ColorIdBag_) Reset()                    { *m = ColorIdBag_{} }
func (m *ColorIdBag_) String() string            { return proto.CompactTextString(m) }
func (*ColorIdBag_) ProtoMessage()               {}
func (*ColorIdBag_) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ColorIdBag_) GetColorIdVal() uint32 {
	if m != nil {
		return m.ColorIdVal
	}
	return 0
}

type BgpGlobalProcessInfoGbl_ struct {
	// Standalone or Distributed mode
	InStandaloneMode bool `protobuf:"varint,1,opt,name=in_standalone_mode,json=inStandaloneMode" json:"in_standalone_mode,omitempty"`
	// Local autonomous system number
	LocalAs uint32 `protobuf:"varint,2,opt,name=local_as,json=localAs" json:"local_as,omitempty"`
	// Name of BGP instance
	InstanceName string `protobuf:"bytes,3,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	// No of times BGP has started
	RestartCount uint32 `protobuf:"varint,4,opt,name=restart_count,json=restartCount" json:"restart_count,omitempty"`
	// Update delay timeout time (in seconds)
	UpdateDelay uint32 `protobuf:"varint,5,opt,name=update_delay,json=updateDelay" json:"update_delay,omitempty"`
	// Period (in seconds) of generic scanner runs
	GenericScanPeriod uint32 `protobuf:"varint,6,opt,name=generic_scan_period,json=genericScanPeriod" json:"generic_scan_period,omitempty"`
	// Confederation ID
	ConfederationId uint32 `protobuf:"varint,7,opt,name=confederation_id,json=confederationId" json:"confederation_id,omitempty"`
	// Cluster ID
	ClusterId uint32 `protobuf:"varint,8,opt,name=cluster_id,json=clusterId" json:"cluster_id,omitempty"`
	// Configured cluster ID
	ConfiguredClusterId uint32 `protobuf:"varint,9,opt,name=configured_cluster_id,json=configuredClusterId" json:"configured_cluster_id,omitempty"`
	// Cluster ID specified as an IP address
	IsClusterIdSpecifiedAsIp bool `protobuf:"varint,10,opt,name=is_cluster_id_specified_as_ip,json=isClusterIdSpecifiedAsIp" json:"is_cluster_id_specified_as_ip,omitempty"`
	// All configured Cluster IDs
	ClusterIdList []*ClusterIdBag_ `protobuf:"bytes,11,rep,name=cluster_id_list,json=clusterIdList" json:"cluster_id_list,omitempty"`
	// All configured Color IDs
	ColorIdList []*ColorIdBag_ `protobuf:"bytes,12,rep,name=color_id_list,json=colorIdList" json:"color_id_list,omitempty"`
	// BGP AS Number Format
	AsnFormat uint32 `protobuf:"varint,13,opt,name=asn_format,json=asnFormat" json:"asn_format,omitempty"`
	// Configured segment-routing Global Block start value
	SrgbStartConfigured uint32 `protobuf:"varint,14,opt,name=srgb_start_configured,json=srgbStartConfigured" json:"srgb_start_configured,omitempty"`
	// Configured segment-routing Global Block end value
	SrgbEndConfigured uint32 `protobuf:"varint,15,opt,name=srgb_end_configured,json=srgbEndConfigured" json:"srgb_end_configured,omitempty"`
	// In use segment-routing Global Block start value
	SrgbStart uint32 `protobuf:"varint,16,opt,name=srgb_start,json=srgbStart" json:"srgb_start,omitempty"`
	// In use Segment-routing Global Block end value
	SrgbEnd uint32 `protobuf:"varint,17,opt,name=srgb_end,json=srgbEnd" json:"srgb_end,omitempty"`
	// Graceful Maintenance
	GracefulMaintenance bool `protobuf:"varint,18,opt,name=graceful_maintenance,json=gracefulMaintenance" json:"graceful_maintenance,omitempty"`
	// Graceful Maintenance also for neighbors without GM configuration
	GracefulMaintAllNbrs bool `protobuf:"varint,19,opt,name=graceful_maint_all_nbrs,json=gracefulMaintAllNbrs" json:"graceful_maint_all_nbrs,omitempty"`
	// Retaining routes in RIB when BGP process stops while in Graceful Maintenance
	GracefulMaintRetainRoutes bool `protobuf:"varint,20,opt,name=graceful_maint_retain_routes,json=gracefulMaintRetainRoutes" json:"graceful_maint_retain_routes,omitempty"`
	// Platform RLIMIT max for the process
	ProcessRlimit uint64 `protobuf:"varint,21,opt,name=process_rlimit,json=processRlimit" json:"process_rlimit,omitempty"`
	// Maximum limit user can configure for max-buffer-size command under bmp-server
	BmpMaximumBufferSize uint64 `protobuf:"varint,22,opt,name=bmp_maximum_buffer_size,json=bmpMaximumBufferSize" json:"bmp_maximum_buffer_size,omitempty"`
	// Default value for BMP buffer limit when a value is not configured
	BmpDefaultBufferSize uint64 `protobuf:"varint,23,opt,name=bmp_default_buffer_size,json=bmpDefaultBufferSize" json:"bmp_default_buffer_size,omitempty"`
	// Current value for BMP buffer
	BmpCurrentBufferSize uint64 `protobuf:"varint,24,opt,name=bmp_current_buffer_size,json=bmpCurrentBufferSize" json:"bmp_current_buffer_size,omitempty"`
	// Maximum limit user has configure using max-buffer-size command under bmp-server
	BmpCurMaximumBufferSize uint64 `protobuf:"varint,25,opt,name=bmp_cur_maximum_buffer_size,json=bmpCurMaximumBufferSize" json:"bmp_cur_maximum_buffer_size,omitempty"`
}

func (m *BgpGlobalProcessInfoGbl_) Reset()                    { *m = BgpGlobalProcessInfoGbl_{} }
func (m *BgpGlobalProcessInfoGbl_) String() string            { return proto.CompactTextString(m) }
func (*BgpGlobalProcessInfoGbl_) ProtoMessage()               {}
func (*BgpGlobalProcessInfoGbl_) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BgpGlobalProcessInfoGbl_) GetInStandaloneMode() bool {
	if m != nil {
		return m.InStandaloneMode
	}
	return false
}

func (m *BgpGlobalProcessInfoGbl_) GetLocalAs() uint32 {
	if m != nil {
		return m.LocalAs
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpGlobalProcessInfoGbl_) GetRestartCount() uint32 {
	if m != nil {
		return m.RestartCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetUpdateDelay() uint32 {
	if m != nil {
		return m.UpdateDelay
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetGenericScanPeriod() uint32 {
	if m != nil {
		return m.GenericScanPeriod
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetConfederationId() uint32 {
	if m != nil {
		return m.ConfederationId
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetClusterId() uint32 {
	if m != nil {
		return m.ClusterId
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetConfiguredClusterId() uint32 {
	if m != nil {
		return m.ConfiguredClusterId
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetIsClusterIdSpecifiedAsIp() bool {
	if m != nil {
		return m.IsClusterIdSpecifiedAsIp
	}
	return false
}

func (m *BgpGlobalProcessInfoGbl_) GetClusterIdList() []*ClusterIdBag_ {
	if m != nil {
		return m.ClusterIdList
	}
	return nil
}

func (m *BgpGlobalProcessInfoGbl_) GetColorIdList() []*ColorIdBag_ {
	if m != nil {
		return m.ColorIdList
	}
	return nil
}

func (m *BgpGlobalProcessInfoGbl_) GetAsnFormat() uint32 {
	if m != nil {
		return m.AsnFormat
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetSrgbStartConfigured() uint32 {
	if m != nil {
		return m.SrgbStartConfigured
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetSrgbEndConfigured() uint32 {
	if m != nil {
		return m.SrgbEndConfigured
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetSrgbStart() uint32 {
	if m != nil {
		return m.SrgbStart
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetSrgbEnd() uint32 {
	if m != nil {
		return m.SrgbEnd
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetGracefulMaintenance() bool {
	if m != nil {
		return m.GracefulMaintenance
	}
	return false
}

func (m *BgpGlobalProcessInfoGbl_) GetGracefulMaintAllNbrs() bool {
	if m != nil {
		return m.GracefulMaintAllNbrs
	}
	return false
}

func (m *BgpGlobalProcessInfoGbl_) GetGracefulMaintRetainRoutes() bool {
	if m != nil {
		return m.GracefulMaintRetainRoutes
	}
	return false
}

func (m *BgpGlobalProcessInfoGbl_) GetProcessRlimit() uint64 {
	if m != nil {
		return m.ProcessRlimit
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetBmpMaximumBufferSize() uint64 {
	if m != nil {
		return m.BmpMaximumBufferSize
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetBmpDefaultBufferSize() uint64 {
	if m != nil {
		return m.BmpDefaultBufferSize
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetBmpCurrentBufferSize() uint64 {
	if m != nil {
		return m.BmpCurrentBufferSize
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetBmpCurMaximumBufferSize() uint64 {
	if m != nil {
		return m.BmpCurMaximumBufferSize
	}
	return 0
}

type BgpGlobalProcessInfoVrf_ struct {
	// VRF state
	VrfIsActive bool `protobuf:"varint,1,opt,name=vrf_is_active,json=vrfIsActive" json:"vrf_is_active,omitempty"`
	// Route Distinguisher
	RouteDistinguisher []byte `protobuf:"bytes,2,opt,name=route_distinguisher,json=routeDistinguisher,proto3" json:"route_distinguisher,omitempty"`
	// Router ID for the VRF
	RouterId string `protobuf:"bytes,3,opt,name=router_id,json=routerId" json:"router_id,omitempty"`
	// Configured router ID
	ConfiguredRouterId string `protobuf:"bytes,4,opt,name=configured_router_id,json=configuredRouterId" json:"configured_router_id,omitempty"`
	// Redistribute iBGP into IGPs enabled
	IsRedistributeIbgpToIgPsEnabled bool `protobuf:"varint,5,opt,name=is_redistribute_ibgp_to_ig_ps_enabled,json=isRedistributeIbgpToIgPsEnabled" json:"is_redistribute_ibgp_to_ig_ps_enabled,omitempty"`
	// Fast external fallover enabled
	IsFastExternalFalloverEnabled bool `protobuf:"varint,6,opt,name=is_fast_external_fallover_enabled,json=isFastExternalFalloverEnabled" json:"is_fast_external_fallover_enabled,omitempty"`
	// Bestpath: Treat missing MED as worst
	IsBestpathMissingMedIsWorstEnabled bool `protobuf:"varint,7,opt,name=is_bestpath_missing_med_is_worst_enabled,json=isBestpathMissingMedIsWorstEnabled" json:"is_bestpath_missing_med_is_worst_enabled,omitempty"`
	// Bestpath: Always compare MED
	IsBestpathAlwaysCompareMedEnabled bool `protobuf:"varint,8,opt,name=is_bestpath_always_compare_med_enabled,json=isBestpathAlwaysCompareMedEnabled" json:"is_bestpath_always_compare_med_enabled,omitempty"`
	// Bestpath: Ignore AS path
	IsBestpathIgnoreAsPathEnabled bool `protobuf:"varint,9,opt,name=is_bestpath_ignore_as_path_enabled,json=isBestpathIgnoreAsPathEnabled" json:"is_bestpath_ignore_as_path_enabled,omitempty"`
	// Bestpath: Relax AS path for mpath
	IsBestpathAsPathMpathRelaxEnabled bool `protobuf:"varint,10,opt,name=is_bestpath_as_path_mpath_relax_enabled,json=isBestpathAsPathMpathRelaxEnabled" json:"is_bestpath_as_path_mpath_relax_enabled,omitempty"`
	// Bestpath: Compare MED from confed peer
	IsBestpathCompareMedFromConfedPeerEnabled bool `protobuf:"varint,11,opt,name=is_bestpath_compare_med_from_confed_peer_enabled,json=isBestpathCompareMedFromConfedPeerEnabled" json:"is_bestpath_compare_med_from_confed_peer_enabled,omitempty"`
	// Bestpath: Compare routerID for eBGP peers
	IsBestpathCompareRouterIdForEbgpPeersEnabled bool `protobuf:"varint,12,opt,name=is_bestpath_compare_router_id_for_ebgp_peers_enabled,json=isBestpathCompareRouterIdForEbgpPeersEnabled" json:"is_bestpath_compare_router_id_for_ebgp_peers_enabled,omitempty"`
	// Bestpath: Ignore AIGP unless both paths have AIGP attribute
	IsBestpathAigpIgnoreEnabled bool `protobuf:"varint,13,opt,name=is_bestpath_aigp_ignore_enabled,json=isBestpathAigpIgnoreEnabled" json:"is_bestpath_aigp_ignore_enabled,omitempty"`
	// Multipath: Ignore everything AS path onwards for mpath
	IsMultipathAsPathIgnoreOnwardsEnabled bool `protobuf:"varint,14,opt,name=is_multipath_as_path_ignore_onwards_enabled,json=isMultipathAsPathIgnoreOnwardsEnabled" json:"is_multipath_as_path_ignore_onwards_enabled,omitempty"`
	// Enforce first AS
	IsEnforceFirstAsEnabled bool `protobuf:"varint,15,opt,name=is_enforce_first_as_enabled,json=isEnforceFirstAsEnabled" json:"is_enforce_first_as_enabled,omitempty"`
	// Default local preference
	DefaultLocalPreference uint32 `protobuf:"varint,16,opt,name=default_local_preference,json=defaultLocalPreference" json:"default_local_preference,omitempty"`
	// Default keepalive timer (seconds)
	KeepAliveTime uint32 `protobuf:"varint,17,opt,name=keep_alive_time,json=keepAliveTime" json:"keep_alive_time,omitempty"`
	// Default hold timer (seconds)
	HoldTime uint32 `protobuf:"varint,18,opt,name=hold_time,json=holdTime" json:"hold_time,omitempty"`
	// Default min acceptable hold time from neighbor(seconds)
	MinAcceptableHoldTime uint32 `protobuf:"varint,19,opt,name=min_acceptable_hold_time,json=minAcceptableHoldTime" json:"min_acceptable_hold_time,omitempty"`
	// Neighbor logging enabled
	IsNeighborLogging bool `protobuf:"varint,20,opt,name=is_neighbor_logging,json=isNeighborLogging" json:"is_neighbor_logging,omitempty"`
	// Default metric configured
	IsDefaultMetricConfigured bool `protobuf:"varint,21,opt,name=is_default_metric_configured,json=isDefaultMetricConfigured" json:"is_default_metric_configured,omitempty"`
	// Default metric
	DefaultMetric uint32 `protobuf:"varint,22,opt,name=default_metric,json=defaultMetric" json:"default_metric,omitempty"`
	// Default route originate configured
	IsDefaultOriginateConfigured bool `protobuf:"varint,23,opt,name=is_default_originate_configured,json=isDefaultOriginateConfigured" json:"is_default_originate_configured,omitempty"`
	// Graceful restart enabled
	IsGracefulRestart bool `protobuf:"varint,24,opt,name=is_graceful_restart,json=isGracefulRestart" json:"is_graceful_restart,omitempty"`
	// Non-stop routing enabled
	IsNsr bool `protobuf:"varint,25,opt,name=is_nsr,json=isNsr" json:"is_nsr,omitempty"`
	// Restart time (in seconds)
	RestartTime uint32 `protobuf:"varint,26,opt,name=restart_time,json=restartTime" json:"restart_time,omitempty"`
	// Stale path timeout time (in seconds)
	StalePathTime uint32 `protobuf:"varint,27,opt,name=stale_path_time,json=stalePathTime" json:"stale_path_time,omitempty"`
	// RIB purge timeout time (in seconds)
	RibPurgeTimeout uint32 `protobuf:"varint,28,opt,name=rib_purge_timeout,json=ribPurgeTimeout" json:"rib_purge_timeout,omitempty"`
}

func (m *BgpGlobalProcessInfoVrf_) Reset()                    { *m = BgpGlobalProcessInfoVrf_{} }
func (m *BgpGlobalProcessInfoVrf_) String() string            { return proto.CompactTextString(m) }
func (*BgpGlobalProcessInfoVrf_) ProtoMessage()               {}
func (*BgpGlobalProcessInfoVrf_) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BgpGlobalProcessInfoVrf_) GetVrfIsActive() bool {
	if m != nil {
		return m.VrfIsActive
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetRouteDistinguisher() []byte {
	if m != nil {
		return m.RouteDistinguisher
	}
	return nil
}

func (m *BgpGlobalProcessInfoVrf_) GetRouterId() string {
	if m != nil {
		return m.RouterId
	}
	return ""
}

func (m *BgpGlobalProcessInfoVrf_) GetConfiguredRouterId() string {
	if m != nil {
		return m.ConfiguredRouterId
	}
	return ""
}

func (m *BgpGlobalProcessInfoVrf_) GetIsRedistributeIbgpToIgPsEnabled() bool {
	if m != nil {
		return m.IsRedistributeIbgpToIgPsEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsFastExternalFalloverEnabled() bool {
	if m != nil {
		return m.IsFastExternalFalloverEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsBestpathMissingMedIsWorstEnabled() bool {
	if m != nil {
		return m.IsBestpathMissingMedIsWorstEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsBestpathAlwaysCompareMedEnabled() bool {
	if m != nil {
		return m.IsBestpathAlwaysCompareMedEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsBestpathIgnoreAsPathEnabled() bool {
	if m != nil {
		return m.IsBestpathIgnoreAsPathEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsBestpathAsPathMpathRelaxEnabled() bool {
	if m != nil {
		return m.IsBestpathAsPathMpathRelaxEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsBestpathCompareMedFromConfedPeerEnabled() bool {
	if m != nil {
		return m.IsBestpathCompareMedFromConfedPeerEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsBestpathCompareRouterIdForEbgpPeersEnabled() bool {
	if m != nil {
		return m.IsBestpathCompareRouterIdForEbgpPeersEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsBestpathAigpIgnoreEnabled() bool {
	if m != nil {
		return m.IsBestpathAigpIgnoreEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsMultipathAsPathIgnoreOnwardsEnabled() bool {
	if m != nil {
		return m.IsMultipathAsPathIgnoreOnwardsEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsEnforceFirstAsEnabled() bool {
	if m != nil {
		return m.IsEnforceFirstAsEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetDefaultLocalPreference() uint32 {
	if m != nil {
		return m.DefaultLocalPreference
	}
	return 0
}

func (m *BgpGlobalProcessInfoVrf_) GetKeepAliveTime() uint32 {
	if m != nil {
		return m.KeepAliveTime
	}
	return 0
}

func (m *BgpGlobalProcessInfoVrf_) GetHoldTime() uint32 {
	if m != nil {
		return m.HoldTime
	}
	return 0
}

func (m *BgpGlobalProcessInfoVrf_) GetMinAcceptableHoldTime() uint32 {
	if m != nil {
		return m.MinAcceptableHoldTime
	}
	return 0
}

func (m *BgpGlobalProcessInfoVrf_) GetIsNeighborLogging() bool {
	if m != nil {
		return m.IsNeighborLogging
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsDefaultMetricConfigured() bool {
	if m != nil {
		return m.IsDefaultMetricConfigured
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetDefaultMetric() uint32 {
	if m != nil {
		return m.DefaultMetric
	}
	return 0
}

func (m *BgpGlobalProcessInfoVrf_) GetIsDefaultOriginateConfigured() bool {
	if m != nil {
		return m.IsDefaultOriginateConfigured
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsGracefulRestart() bool {
	if m != nil {
		return m.IsGracefulRestart
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsNsr() bool {
	if m != nil {
		return m.IsNsr
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetRestartTime() uint32 {
	if m != nil {
		return m.RestartTime
	}
	return 0
}

func (m *BgpGlobalProcessInfoVrf_) GetStalePathTime() uint32 {
	if m != nil {
		return m.StalePathTime
	}
	return 0
}

func (m *BgpGlobalProcessInfoVrf_) GetRibPurgeTimeout() uint32 {
	if m != nil {
		return m.RibPurgeTimeout
	}
	return 0
}

func init() {
	proto.RegisterType((*BgpGlobalProcessInfoBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.global_process_info.bgp_global_process_info_bag_KEYS")
	proto.RegisterType((*BgpGlobalProcessInfoBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.global_process_info.bgp_global_process_info_bag")
	proto.RegisterType((*ClusterIdBag_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.global_process_info.cluster_id_bag_")
	proto.RegisterType((*ColorIdBag_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.global_process_info.color_id_bag_")
	proto.RegisterType((*BgpGlobalProcessInfoGbl_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.global_process_info.bgp_global_process_info_gbl_")
	proto.RegisterType((*BgpGlobalProcessInfoVrf_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.global_process_info.bgp_global_process_info_vrf_")
}

func init() { proto.RegisterFile("bgp_global_process_info_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0x51, 0x73, 0xdc, 0xb6,
	0x11, 0x9e, 0xab, 0x6d, 0xf9, 0x04, 0xe9, 0x24, 0x1b, 0x92, 0x22, 0xba, 0xb2, 0x26, 0xf2, 0xb5,
	0x76, 0x95, 0x36, 0x73, 0x4d, 0x1c, 0x67, 0xda, 0x87, 0xcc, 0x64, 0x2e, 0xb2, 0x94, 0xdc, 0xd4,
	0xa7, 0xa8, 0x94, 0xa6, 0x9d, 0xf6, 0x05, 0x03, 0x92, 0x20, 0xb5, 0x2d, 0x09, 0xb0, 0x00, 0xee,
	0x6c, 0xe5, 0x29, 0xff, 0xa1, 0xd3, 0x87, 0x4e, 0xff, 0x41, 0xff, 0x5a, 0xff, 0x44, 0x07, 0x0b,
	0x82, 0xe4, 0x45, 0x8e, 0xdf, 0xea, 0xe9, 0x8b, 0x8e, 0xdc, 0xfd, 0xbe, 0x0f, 0x8b, 0xc5, 0x72,
	0x17, 0x22, 0x4f, 0x92, 0xa2, 0x66, 0x45, 0xa9, 0x12, 0x5e, 0xb2, 0x5a, 0xab, 0x54, 0x18, 0xc3,
	0x40, 0xe6, 0x8a, 0x25, 0xbc, 0x98, 0xd4, 0x5a, 0x59, 0x45, 0x93, 0x14, 0x4c, 0xaa, 0x18, 0x28,
	0xc3, 0xde, 0x68, 0x06, 0xf5, 0xf2, 0x05, 0x73, 0x24, 0x55, 0x0b, 0x3d, 0x49, 0x8a, 0x7a, 0x02,
	0xd2, 0x58, 0x2e, 0x53, 0x61, 0xda, 0xa7, 0xf6, 0x81, 0xb9, 0x9f, 0x2c, 0xb9, 0x99, 0x64, 0x22,
	0xe7, 0x8b, 0xd2, 0xb2, 0xa5, 0xce, 0x27, 0x6f, 0x59, 0x6d, 0xfc, 0x35, 0x39, 0x7a, 0x47, 0x20,
	0xec, 0x77, 0xa7, 0x7f, 0xba, 0xa4, 0x3f, 0x23, 0xa3, 0x56, 0x57, 0xf2, 0x4a, 0x44, 0x83, 0xa3,
	0xc1, 0xf1, 0x7a, 0xbc, 0x19, 0x8c, 0xe7, 0xbc, 0x12, 0xe3, 0xef, 0xef, 0x90, 0x83, 0x77, 0x28,
	0xd1, 0x47, 0x64, 0xb8, 0xd4, 0xb9, 0xe7, 0x3f, 0x47, 0xfe, 0xfd, 0xa5, 0xce, 0x1d, 0x95, 0xee,
	0x92, 0x7b, 0x4b, 0x9d, 0x43, 0x16, 0x7d, 0x76, 0x34, 0x38, 0x1e, 0xc5, 0xfe, 0x85, 0xfe, 0x73,
	0x40, 0xd6, 0xbc, 0x58, 0xf4, 0xe2, 0x68, 0x70, 0xbc, 0xf1, 0xfc, 0xfb, 0xc1, 0xe4, 0x7f, 0x9f,
	0x90, 0xc9, 0x8f, 0xed, 0xa1, 0x48, 0x4a, 0x16, 0x37, 0x01, 0xd1, 0xbf, 0x0f, 0xc8, 0x9d, 0xa5,
	0xce, 0xa3, 0xcf, 0xff, 0x0f, 0x02, 0x73, 0x59, 0x8d, 0x5d, 0x34, 0x63, 0x46, 0xb6, 0xd3, 0x72,
	0x61, 0xac, 0xd0, 0x0c, 0x32, 0x3c, 0x3e, 0xfa, 0x73, 0xb2, 0xd5, 0x33, 0x2d, 0x79, 0x89, 0x67,
	0x37, 0x8a, 0x37, 0x1b, 0xeb, 0x2c, 0xfb, 0x03, 0x2f, 0xe9, 0xb3, 0x15, 0xa2, 0xbd, 0xa9, 0x45,
	0xf4, 0x13, 0x84, 0x8d, 0x5a, 0xd8, 0xd5, 0x4d, 0x2d, 0xc6, 0x9f, 0x92, 0x51, 0xaa, 0x4a, 0xd5,
	0xc9, 0x1f, 0x91, 0xcd, 0xd6, 0xd0, 0x89, 0x13, 0xb4, 0xa1, 0xf4, 0xf8, 0x3f, 0x84, 0x3c, 0x7e,
	0x57, 0x4a, 0xe9, 0xc7, 0x84, 0x82, 0xf4, 0x49, 0xe0, 0xa5, 0x92, 0x82, 0x55, 0x2a, 0xf3, 0x15,
	0x36, 0x8c, 0x1f, 0x80, 0xbc, 0x6c, 0x1d, 0x73, 0x95, 0x09, 0x57, 0x45, 0xa5, 0x4a, 0x79, 0xc9,
	0xb8, 0x69, 0x42, 0xbc, 0x8f, 0xef, 0x53, 0x73, 0xbb, 0x4a, 0xef, 0xdc, 0xae, 0x52, 0x07, 0xd2,
	0xc2, 0x58, 0xae, 0x2d, 0x4b, 0xd5, 0x42, 0xda, 0xe8, 0xae, 0x4f, 0x47, 0x63, 0x3c, 0x71, 0x36,
	0xfa, 0x84, 0x6c, 0x2e, 0xea, 0x8c, 0x5b, 0xc1, 0x32, 0x51, 0xf2, 0x9b, 0xe8, 0x1e, 0x62, 0x36,
	0xbc, 0xed, 0xa5, 0x33, 0xd1, 0x09, 0xd9, 0x29, 0x84, 0x14, 0x1a, 0x52, 0x66, 0x52, 0x2e, 0x59,
	0x2d, 0x34, 0xa8, 0x2c, 0x5a, 0x43, 0xe4, 0xc3, 0xc6, 0x75, 0x99, 0x72, 0x79, 0x81, 0x0e, 0xfa,
	0x11, 0x79, 0x90, 0x2a, 0x99, 0x8b, 0x4c, 0x68, 0x6e, 0x41, 0x49, 0x06, 0x59, 0x74, 0x1f, 0xc1,
	0xdb, 0x2b, 0xf6, 0x59, 0x46, 0x0f, 0x09, 0xe9, 0x0e, 0x23, 0x1a, 0x22, 0x68, 0xbd, 0x3d, 0x07,
	0xfa, 0x9c, 0xec, 0x39, 0x06, 0x14, 0x0b, 0x2d, 0x32, 0xd6, 0x43, 0xae, 0x23, 0x72, 0xa7, 0x73,
	0x9e, 0xb4, 0x9c, 0x2f, 0xc9, 0x21, 0x98, 0x1e, 0x96, 0x99, 0x5a, 0xa4, 0x90, 0x83, 0xc8, 0x18,
	0x37, 0x0c, 0xea, 0x88, 0x60, 0xba, 0x23, 0x30, 0x2d, 0xe7, 0x32, 0x20, 0xa6, 0x66, 0x56, 0xd3,
	0x7f, 0x0d, 0x56, 0x2a, 0xa4, 0x04, 0x63, 0xa3, 0x8d, 0xa3, 0x3b, 0xc7, 0x1b, 0xcf, 0xcd, 0xfb,
	0x28, 0xfd, 0x1f, 0x54, 0x75, 0xaf, 0x2c, 0x5f, 0x81, 0xb1, 0xf4, 0x1f, 0x83, 0x5e, 0x5d, 0x62,
	0x6c, 0x9b, 0x18, 0xdb, 0xdf, 0xde, 0x4b, 0x6c, 0xfd, 0x0f, 0x22, 0xde, 0x68, 0x4a, 0x1f, 0xe3,
	0x3a, 0x24, 0x84, 0x1b, 0xc9, 0x72, 0xa5, 0x2b, 0x6e, 0xa3, 0x91, 0x3f, 0x49, 0x6e, 0xe4, 0x19,
	0x1a, 0xdc, 0x49, 0x1a, 0x5d, 0x24, 0x2c, 0x94, 0x63, 0x38, 0xb7, 0x68, 0xcb, 0x9f, 0xa4, 0x73,
	0x5e, 0xfa, 0xaa, 0x0c, 0x2e, 0x57, 0x77, 0xc8, 0x11, 0x32, 0xeb, 0x33, 0xb6, 0x7d, 0xdd, 0x39,
	0xd7, 0xa9, 0xcc, 0x7a, 0xf8, 0x43, 0x42, 0xba, 0x35, 0xa2, 0x07, 0x3e, 0x84, 0x56, 0xd8, 0x7d,
	0x4e, 0x41, 0x2e, 0x7a, 0xe8, 0x3f, 0xa7, 0x46, 0x83, 0x7e, 0x4a, 0x76, 0x0b, 0xcd, 0x53, 0x91,
	0x2f, 0x4a, 0x56, 0x71, 0x90, 0x56, 0x48, 0x97, 0x97, 0x88, 0x62, 0xa9, 0xec, 0x04, 0xdf, 0xbc,
	0x73, 0xd1, 0xcf, 0xc9, 0xfe, 0x2a, 0x85, 0xf1, 0xb2, 0x64, 0x32, 0xd1, 0x26, 0xda, 0x41, 0xd6,
	0xee, 0x0a, 0x6b, 0x5a, 0x96, 0xe7, 0x89, 0x36, 0xf4, 0x4b, 0xf2, 0xf8, 0x07, 0x34, 0x2d, 0x2c,
	0x07, 0xc9, 0xb4, 0x5a, 0x58, 0x61, 0xa2, 0x5d, 0xe4, 0x3e, 0x5a, 0xe1, 0xc6, 0x88, 0x88, 0x11,
	0x40, 0x9f, 0x92, 0xad, 0x70, 0x24, 0xba, 0x84, 0x0a, 0x6c, 0xb4, 0x77, 0x34, 0x38, 0xbe, 0x1b,
	0x8f, 0x1a, 0x6b, 0x8c, 0x46, 0x17, 0x5e, 0x52, 0xd5, 0xac, 0xe2, 0x6f, 0xa0, 0x5a, 0x54, 0x2c,
	0x59, 0xe4, 0xb9, 0xd0, 0xcc, 0xc0, 0x77, 0x22, 0xfa, 0x00, 0xf1, 0xbb, 0x49, 0x55, 0xcf, 0xbd,
	0xf7, 0x2b, 0x74, 0x5e, 0xc2, 0x77, 0x22, 0xd0, 0x42, 0x1d, 0xf4, 0x69, 0xfb, 0x2d, 0xed, 0xa5,
	0xf7, 0xde, 0xa6, 0xa5, 0x0b, 0xad, 0x85, 0x5c, 0xa5, 0x45, 0x2d, 0xed, 0xc4, 0x7b, 0x7b, 0xb4,
	0x2f, 0xc8, 0x41, 0x43, 0x7b, 0x6b, 0xa0, 0x8f, 0x90, 0xba, 0xef, 0xa9, 0xb7, 0x62, 0x1d, 0xff,
	0x7b, 0xf4, 0xe3, 0xdd, 0xd6, 0xcd, 0x09, 0x3a, 0x26, 0x23, 0xf7, 0x0b, 0x86, 0xf1, 0xd4, 0xc2,
	0x32, 0x34, 0xda, 0x8d, 0xa5, 0xce, 0x67, 0x66, 0x8a, 0x26, 0xfa, 0x6b, 0xb2, 0x83, 0x99, 0x67,
	0x19, 0x18, 0x0b, 0xb2, 0x58, 0x80, 0xb9, 0x16, 0x1a, 0xdb, 0xed, 0x66, 0x4c, 0xd1, 0xf5, 0xb2,
	0xef, 0xa1, 0x07, 0x64, 0x1d, 0xad, 0xd8, 0x86, 0x7c, 0xd7, 0x1d, 0x7a, 0xc3, 0x2c, 0xa3, 0x9f,
	0x90, 0xdd, 0x5e, 0xbf, 0xea, 0x70, 0x77, 0x11, 0x47, 0x3b, 0x5f, 0x1c, 0x18, 0xe7, 0xe4, 0x29,
	0x18, 0xa6, 0x85, 0x5b, 0x5e, 0x43, 0xe2, 0x22, 0x01, 0xb7, 0x29, 0xab, 0x18, 0x14, 0xac, 0x36,
	0x4c, 0x48, 0x9e, 0x94, 0x22, 0xc3, 0xbe, 0x3c, 0x8c, 0x3f, 0x04, 0x13, 0xf7, 0xb0, 0xb3, 0xa4,
	0xa8, 0xaf, 0xd4, 0xac, 0xb8, 0x30, 0xa7, 0x1e, 0x46, 0xbf, 0x21, 0x4f, 0xc0, 0xb0, 0x9c, 0x1b,
	0xcb, 0xc4, 0x1b, 0x2b, 0xb4, 0xe4, 0x25, 0xcb, 0x79, 0x59, 0xaa, 0xa5, 0xd0, 0xad, 0xd6, 0x1a,
	0x6a, 0x1d, 0x82, 0x39, 0xe3, 0xc6, 0x9e, 0x36, 0xb0, 0xb3, 0x06, 0x15, 0x94, 0xae, 0xc8, 0x31,
	0x18, 0x96, 0x08, 0x63, 0x6b, 0x6e, 0xaf, 0x59, 0x05, 0xc6, 0x80, 0x2c, 0x58, 0x25, 0x32, 0x97,
	0xd1, 0xd7, 0x4a, 0xbb, 0x25, 0x1a, 0xc1, 0xfb, 0x28, 0x38, 0x06, 0xf3, 0x55, 0x03, 0x9f, 0x7b,
	0xf4, 0x5c, 0x64, 0x33, 0xf3, 0x47, 0x07, 0x0d, 0xaa, 0xbf, 0x27, 0xcf, 0xfa, 0xaa, 0xbc, 0x7c,
	0xcd, 0x6f, 0x0c, 0x4b, 0x55, 0x55, 0x73, 0x2d, 0x50, 0x3c, 0x68, 0x0e, 0x51, 0xf3, 0x49, 0xa7,
	0x39, 0x45, 0xec, 0x89, 0x87, 0xce, 0x45, 0x16, 0x24, 0x67, 0x64, 0xdc, 0x97, 0x84, 0x42, 0x2a,
	0x2d, 0x5c, 0xaf, 0xc7, 0xd7, 0x20, 0xb7, 0x1e, 0xf6, 0x1c, 0xe4, 0x66, 0x88, 0x9b, 0x9a, 0x0b,
	0x6e, 0xaf, 0x83, 0x54, 0x4c, 0x7e, 0xb1, 0x12, 0x5d, 0xa3, 0x51, 0xe1, 0x5f, 0x2d, 0x4a, 0xfe,
	0xa6, 0xd5, 0x23, 0xb7, 0xc2, 0x43, 0xa5, 0xb9, 0x7b, 0x8c, 0x1d, 0x32, 0x68, 0xa6, 0xe4, 0x93,
	0xbe, 0x66, 0x7f, 0xab, 0xb9, 0x56, 0x15, 0xf3, 0x23, 0x91, 0xd5, 0xa2, 0x77, 0x40, 0x1b, 0x28,
	0xfe, 0x51, 0x27, 0xde, 0xed, 0xfa, 0x4c, 0xab, 0xea, 0x04, 0x29, 0x17, 0xa2, 0x3b, 0xac, 0xbf,
	0x90, 0x17, 0x6f, 0x5b, 0xa4, 0xad, 0x40, 0xd7, 0x97, 0x99, 0x70, 0x85, 0xe5, 0xd6, 0xe9, 0xaa,
	0x6a, 0x13, 0x17, 0xfa, 0xf8, 0xd6, 0x42, 0xa1, 0x3e, 0xcf, 0x94, 0x3e, 0x4d, 0x8a, 0xda, 0x2d,
	0xd5, 0x96, 0xd8, 0x4b, 0xf2, 0xe1, 0x4a, 0x92, 0xa0, 0xa8, 0x43, 0xd2, 0x83, 0xec, 0x08, 0x65,
	0x0f, 0x7a, 0xc9, 0x81, 0xa2, 0xf6, 0x09, 0x0f, 0x2a, 0x7f, 0x26, 0xbf, 0x02, 0xc3, 0xaa, 0x45,
	0x69, 0x61, 0x25, 0xd7, 0x8d, 0x92, 0x92, 0xaf, 0xb9, 0xce, 0xba, 0x40, 0xb7, 0x50, 0xf1, 0x29,
	0x98, 0x79, 0x60, 0xf8, 0x7c, 0x7b, 0xd1, 0x6f, 0x3d, 0x3a, 0x68, 0x7f, 0x41, 0x0e, 0xc0, 0x51,
	0x73, 0xa5, 0x53, 0xc1, 0x72, 0x70, 0xa5, 0xca, 0x3b, 0xad, 0x6d, 0xd4, 0xda, 0x07, 0x73, 0xea,
	0x11, 0x67, 0x0e, 0x30, 0x6d, 0xd9, 0xbf, 0x25, 0x51, 0xe8, 0x7f, 0xfe, 0xfa, 0x55, 0x6b, 0x91,
	0x0b, 0x2d, 0xdc, 0x40, 0xf0, 0x43, 0xe5, 0x83, 0xc6, 0xff, 0xca, 0xb9, 0x2f, 0x5a, 0xaf, 0xbb,
	0x5a, 0xfe, 0x55, 0x88, 0x9a, 0xf1, 0x12, 0x96, 0x82, 0x59, 0xa8, 0x44, 0x33, 0x68, 0x46, 0xce,
	0x3c, 0x75, 0xd6, 0x2b, 0xa8, 0x84, 0xeb, 0x21, 0xd7, 0xaa, 0xcc, 0x3c, 0x82, 0x22, 0x62, 0xe8,
	0x0c, 0xe8, 0xfc, 0x0d, 0x89, 0x2a, 0x90, 0x8c, 0xa7, 0xa9, 0xa8, 0xad, 0x8b, 0x88, 0x75, 0xd8,
	0x1d, 0xc4, 0xee, 0x55, 0x20, 0xa7, 0xad, 0xfb, 0x9b, 0x40, 0x9c, 0x90, 0x1d, 0x30, 0x4c, 0x0a,
	0x28, 0xae, 0x13, 0xa5, 0x59, 0xa9, 0x8a, 0x02, 0x64, 0xd1, 0x4c, 0x94, 0x87, 0x60, 0xce, 0x1b,
	0xcf, 0x2b, 0xef, 0x70, 0xa3, 0x08, 0x4c, 0xdb, 0xea, 0x2b, 0x61, 0xdd, 0x05, 0xaf, 0x37, 0x67,
	0xf7, 0xfc, 0x28, 0x02, 0xd3, 0xf4, 0xfb, 0x39, 0x22, 0x7a, 0xf3, 0xf6, 0x29, 0xd9, 0x5a, 0x65,
	0xe3, 0x68, 0x19, 0xc5, 0xa3, 0xac, 0x4f, 0xa0, 0xa7, 0x58, 0x2f, 0x01, 0xa9, 0x34, 0x14, 0x20,
	0xdd, 0x7d, 0xb3, 0xb7, 0xd4, 0x3e, 0x2e, 0xf5, 0xb8, 0x5d, 0xea, 0xdb, 0x00, 0x5a, 0xbd, 0x0d,
	0x80, 0x61, 0xed, 0xf0, 0x6c, 0x2e, 0xb1, 0x38, 0x5f, 0x70, 0x7b, 0x5f, 0x37, 0x9e, 0xd8, 0x3b,
	0xe8, 0x1e, 0x59, 0x73, 0xe9, 0x30, 0x1a, 0xe7, 0xc8, 0x30, 0xbe, 0x07, 0xe6, 0xdc, 0x68, 0x77,
	0xdf, 0x0d, 0x97, 0x62, 0x4c, 0xe9, 0x4f, 0xfd, 0x7d, 0xb7, 0xb1, 0x61, 0x22, 0x9f, 0x91, 0x6d,
	0x63, 0x79, 0x29, 0x7c, 0x41, 0x22, 0xea, 0xc0, 0x6f, 0x0c, 0xcd, 0xae, 0xec, 0x10, 0xf7, 0x4b,
	0xf2, 0x50, 0x43, 0xc2, 0xea, 0x85, 0x2e, 0xfc, 0x69, 0xab, 0x85, 0x8d, 0x1e, 0xfb, 0x8b, 0xae,
	0x86, 0xe4, 0xc2, 0xd9, 0xaf, 0xbc, 0x39, 0x59, 0xc3, 0xff, 0x72, 0x3f, 0xfb, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x7a, 0x1e, 0x75, 0x80, 0x0a, 0x0f, 0x00, 0x00,
}
