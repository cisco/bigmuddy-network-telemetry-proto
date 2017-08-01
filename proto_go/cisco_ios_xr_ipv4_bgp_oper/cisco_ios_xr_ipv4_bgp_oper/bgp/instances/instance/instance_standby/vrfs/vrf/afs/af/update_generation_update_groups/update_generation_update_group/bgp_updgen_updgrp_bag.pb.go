// Code generated by protoc-gen-go.
// source: bgp_updgen_updgrp_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_vrfs_vrf_afs_af_update_generation_update_groups_update_generation_update_group is a generated protocol buffer package.

It is generated from these files:
	bgp_updgen_updgrp_bag.proto

It has these top-level messages:
	BgpUpdgenUpdgrpBag_KEYS
	BgpUpdgenUpdgrpBag
	BgpTimespec
	BgpUpdgenStatsBag
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_vrfs_vrf_afs_af_update_generation_update_groups_update_generation_update_group

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

// BGP Update generation Update-group information
type BgpUpdgenUpdgrpBag_KEYS struct {
	InstanceName     string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	VrfName          string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName           string `protobuf:"bytes,3,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	UpdateGroupIndex uint32 `protobuf:"varint,4,opt,name=update_group_index,json=updateGroupIndex" json:"update_group_index,omitempty"`
}

func (m *BgpUpdgenUpdgrpBag_KEYS) Reset()                    { *m = BgpUpdgenUpdgrpBag_KEYS{} }
func (m *BgpUpdgenUpdgrpBag_KEYS) String() string            { return proto.CompactTextString(m) }
func (*BgpUpdgenUpdgrpBag_KEYS) ProtoMessage()               {}
func (*BgpUpdgenUpdgrpBag_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BgpUpdgenUpdgrpBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpUpdgenUpdgrpBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpUpdgenUpdgrpBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpUpdgenUpdgrpBag_KEYS) GetUpdateGroupIndex() uint32 {
	if m != nil {
		return m.UpdateGroupIndex
	}
	return 0
}

type BgpUpdgenUpdgrpBag struct {
	// ProcessID
	ProcessId uint32 `protobuf:"varint,50,opt,name=process_id,json=processId" json:"process_id,omitempty"`
	// VRF Name
	UpdateVrfName string `protobuf:"bytes,51,opt,name=update_vrf_name,json=updateVrfName" json:"update_vrf_name,omitempty"`
	// Address family identifier
	UpdateGroupAfName string `protobuf:"bytes,52,opt,name=update_group_af_name,json=updateGroupAfName" json:"update_group_af_name,omitempty"`
	// Neighbor session address family
	NeighborSessionAfName uint32 `protobuf:"varint,53,opt,name=neighbor_session_af_name,json=neighborSessionAfName" json:"neighbor_session_af_name,omitempty"`
	// Update-group index
	UpdateGroupIndex uint32 `protobuf:"varint,54,opt,name=update_group_index,json=updateGroupIndex" json:"update_group_index,omitempty"`
	// Update-group internal flags2
	UpdateGroupFlags2 uint32 `protobuf:"varint,55,opt,name=update_group_flags2,json=updateGroupFlags2" json:"update_group_flags2,omitempty"`
	// OutQueue Messages
	UpdateOutQueueMessages uint32 `protobuf:"varint,56,opt,name=update_out_queue_messages,json=updateOutQueueMessages" json:"update_out_queue_messages,omitempty"`
	// OutQueue Size
	UpdateOutQueueSize uint32 `protobuf:"varint,57,opt,name=update_out_queue_size,json=updateOutQueueSize" json:"update_out_queue_size,omitempty"`
	// Sub-group count
	UpdateSubGroupCount uint32 `protobuf:"varint,58,opt,name=update_sub_group_count,json=updateSubGroupCount" json:"update_sub_group_count,omitempty"`
	// Throttled sub-group count
	SubGroupThrottledCount uint32 `protobuf:"varint,59,opt,name=sub_group_throttled_count,json=subGroupThrottledCount" json:"sub_group_throttled_count,omitempty"`
	// Refresh sub-group count
	RefreshSubGroupCount uint32 `protobuf:"varint,60,opt,name=refresh_sub_group_count,json=refreshSubGroupCount" json:"refresh_sub_group_count,omitempty"`
	// Throttled refresh sub-group count
	RefreshSubGroupThrottledCount uint32 `protobuf:"varint,61,opt,name=refresh_sub_group_throttled_count,json=refreshSubGroupThrottledCount" json:"refresh_sub_group_throttled_count,omitempty"`
	// Filter-group count
	FilterGroupCount uint32 `protobuf:"varint,62,opt,name=filter_group_count,json=filterGroupCount" json:"filter_group_count,omitempty"`
	// Neighbor count
	NeighborCount uint32 `protobuf:"varint,63,opt,name=neighbor_count,json=neighborCount" json:"neighbor_count,omitempty"`
	// Count of neighbors leaving the update-group
	NeighborLeavingCount uint32 `protobuf:"varint,64,opt,name=neighbor_leaving_count,json=neighborLeavingCount" json:"neighbor_leaving_count,omitempty"`
	// Is update generation recovery pending
	UpdateGenerationRecoveryPending bool `protobuf:"varint,65,opt,name=update_generation_recovery_pending,json=updateGenerationRecoveryPending" json:"update_generation_recovery_pending,omitempty"`
	// Last update timer start time
	LastUpdateTimerStartTimestamp *BgpTimespec `protobuf:"bytes,66,opt,name=last_update_timer_start_timestamp,json=lastUpdateTimerStartTimestamp" json:"last_update_timer_start_timestamp,omitempty"`
	// Last update timer stop time
	LastUpdateTimerStopTimestamp *BgpTimespec `protobuf:"bytes,67,opt,name=last_update_timer_stop_timestamp,json=lastUpdateTimerStopTimestamp" json:"last_update_timer_stop_timestamp,omitempty"`
	// Last update timer expiry time
	LastUpdateTimerExpiryTimestamp *BgpTimespec `protobuf:"bytes,68,opt,name=last_update_timer_expiry_timestamp,json=lastUpdateTimerExpiryTimestamp" json:"last_update_timer_expiry_timestamp,omitempty"`
	// Time since last update timer expiry event
	LastUpdateTimerExpiryAge uint32 `protobuf:"varint,69,opt,name=last_update_timer_expiry_age,json=lastUpdateTimerExpiryAge" json:"last_update_timer_expiry_age,omitempty"`
	// Is update timer running
	IsUpdateTimerRunning bool `protobuf:"varint,70,opt,name=is_update_timer_running,json=isUpdateTimerRunning" json:"is_update_timer_running,omitempty"`
	// Update timer remaining time
	LastUpdateTimerRemainingValue *BgpTimespec `protobuf:"bytes,71,opt,name=last_update_timer_remaining_value,json=lastUpdateTimerRemainingValue" json:"last_update_timer_remaining_value,omitempty"`
	// Last update timer delay value
	LastUpdateTimerDelayValue *BgpTimespec `protobuf:"bytes,72,opt,name=last_update_timer_delay_value,json=lastUpdateTimerDelayValue" json:"last_update_timer_delay_value,omitempty"`
	// Update statistics
	UpdateStatistics *BgpUpdgenStatsBag `protobuf:"bytes,73,opt,name=update_statistics,json=updateStatistics" json:"update_statistics,omitempty"`
	// Count of Perm Pelems seen by updgen
	PermPelemEncountered uint32 `protobuf:"varint,74,opt,name=perm_pelem_encountered,json=permPelemEncountered" json:"perm_pelem_encountered,omitempty"`
	// Count of Perm Pelems allowed by updgen
	PermPelemAllowed uint32 `protobuf:"varint,75,opt,name=perm_pelem_allowed,json=permPelemAllowed" json:"perm_pelem_allowed,omitempty"`
	// Count of Perm Pelems not allowed by updgen
	PermPelemNotAllowed uint32 `protobuf:"varint,76,opt,name=perm_pelem_not_allowed,json=permPelemNotAllowed" json:"perm_pelem_not_allowed,omitempty"`
	// Count of Perm Pelems explicitly wdrn by updgen
	PermPelemExpWdr uint32 `protobuf:"varint,77,opt,name=perm_pelem_exp_wdr,json=permPelemExpWdr" json:"perm_pelem_exp_wdr,omitempty"`
	// Count of Perm Pelems Spurious withdraws by updgen
	PermPelemSpurWdr uint32 `protobuf:"varint,78,opt,name=perm_pelem_spur_wdr,json=permPelemSpurWdr" json:"perm_pelem_spur_wdr,omitempty"`
	// Permanent UG check
	IsPermanent bool `protobuf:"varint,79,opt,name=is_permanent,json=isPermanent" json:"is_permanent,omitempty"`
}

func (m *BgpUpdgenUpdgrpBag) Reset()                    { *m = BgpUpdgenUpdgrpBag{} }
func (m *BgpUpdgenUpdgrpBag) String() string            { return proto.CompactTextString(m) }
func (*BgpUpdgenUpdgrpBag) ProtoMessage()               {}
func (*BgpUpdgenUpdgrpBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BgpUpdgenUpdgrpBag) GetProcessId() uint32 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetUpdateVrfName() string {
	if m != nil {
		return m.UpdateVrfName
	}
	return ""
}

func (m *BgpUpdgenUpdgrpBag) GetUpdateGroupAfName() string {
	if m != nil {
		return m.UpdateGroupAfName
	}
	return ""
}

func (m *BgpUpdgenUpdgrpBag) GetNeighborSessionAfName() uint32 {
	if m != nil {
		return m.NeighborSessionAfName
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetUpdateGroupIndex() uint32 {
	if m != nil {
		return m.UpdateGroupIndex
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetUpdateGroupFlags2() uint32 {
	if m != nil {
		return m.UpdateGroupFlags2
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetUpdateOutQueueMessages() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessages
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetUpdateOutQueueSize() uint32 {
	if m != nil {
		return m.UpdateOutQueueSize
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetUpdateSubGroupCount() uint32 {
	if m != nil {
		return m.UpdateSubGroupCount
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetSubGroupThrottledCount() uint32 {
	if m != nil {
		return m.SubGroupThrottledCount
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetRefreshSubGroupCount() uint32 {
	if m != nil {
		return m.RefreshSubGroupCount
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetRefreshSubGroupThrottledCount() uint32 {
	if m != nil {
		return m.RefreshSubGroupThrottledCount
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetFilterGroupCount() uint32 {
	if m != nil {
		return m.FilterGroupCount
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetNeighborCount() uint32 {
	if m != nil {
		return m.NeighborCount
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetNeighborLeavingCount() uint32 {
	if m != nil {
		return m.NeighborLeavingCount
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetUpdateGenerationRecoveryPending() bool {
	if m != nil {
		return m.UpdateGenerationRecoveryPending
	}
	return false
}

func (m *BgpUpdgenUpdgrpBag) GetLastUpdateTimerStartTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateTimerStartTimestamp
	}
	return nil
}

func (m *BgpUpdgenUpdgrpBag) GetLastUpdateTimerStopTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateTimerStopTimestamp
	}
	return nil
}

func (m *BgpUpdgenUpdgrpBag) GetLastUpdateTimerExpiryTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateTimerExpiryTimestamp
	}
	return nil
}

func (m *BgpUpdgenUpdgrpBag) GetLastUpdateTimerExpiryAge() uint32 {
	if m != nil {
		return m.LastUpdateTimerExpiryAge
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetIsUpdateTimerRunning() bool {
	if m != nil {
		return m.IsUpdateTimerRunning
	}
	return false
}

func (m *BgpUpdgenUpdgrpBag) GetLastUpdateTimerRemainingValue() *BgpTimespec {
	if m != nil {
		return m.LastUpdateTimerRemainingValue
	}
	return nil
}

func (m *BgpUpdgenUpdgrpBag) GetLastUpdateTimerDelayValue() *BgpTimespec {
	if m != nil {
		return m.LastUpdateTimerDelayValue
	}
	return nil
}

func (m *BgpUpdgenUpdgrpBag) GetUpdateStatistics() *BgpUpdgenStatsBag {
	if m != nil {
		return m.UpdateStatistics
	}
	return nil
}

func (m *BgpUpdgenUpdgrpBag) GetPermPelemEncountered() uint32 {
	if m != nil {
		return m.PermPelemEncountered
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetPermPelemAllowed() uint32 {
	if m != nil {
		return m.PermPelemAllowed
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetPermPelemNotAllowed() uint32 {
	if m != nil {
		return m.PermPelemNotAllowed
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetPermPelemExpWdr() uint32 {
	if m != nil {
		return m.PermPelemExpWdr
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetPermPelemSpurWdr() uint32 {
	if m != nil {
		return m.PermPelemSpurWdr
	}
	return 0
}

func (m *BgpUpdgenUpdgrpBag) GetIsPermanent() bool {
	if m != nil {
		return m.IsPermanent
	}
	return false
}

type BgpTimespec struct {
	// Seconds part of time value
	Seconds uint32 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	// Nanoseconds part of time value
	Nanoseconds uint32 `protobuf:"varint,2,opt,name=nanoseconds" json:"nanoseconds,omitempty"`
}

func (m *BgpTimespec) Reset()                    { *m = BgpTimespec{} }
func (m *BgpTimespec) String() string            { return proto.CompactTextString(m) }
func (*BgpTimespec) ProtoMessage()               {}
func (*BgpTimespec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BgpTimespec) GetSeconds() uint32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *BgpTimespec) GetNanoseconds() uint32 {
	if m != nil {
		return m.Nanoseconds
	}
	return 0
}

// BGP Update generation common statistics information
type BgpUpdgenStatsBag struct {
	// OutQueue High Messages
	UpdateOutQueueMessagesHigh uint32 `protobuf:"varint,1,opt,name=update_out_queue_messages_high,json=updateOutQueueMessagesHigh" json:"update_out_queue_messages_high,omitempty"`
	// OutQueue Cumulative Messages
	UpdateOutQueueMessagesCumulative uint32 `protobuf:"varint,2,opt,name=update_out_queue_messages_cumulative,json=updateOutQueueMessagesCumulative" json:"update_out_queue_messages_cumulative,omitempty"`
	// OutQueue Discarded Messages
	UpdateOutQueueMessagesDiscarded uint32 `protobuf:"varint,3,opt,name=update_out_queue_messages_discarded,json=updateOutQueueMessagesDiscarded" json:"update_out_queue_messages_discarded,omitempty"`
	// OutQueue Cleared Messages
	UpdateOutQueueMessagesCleared uint32 `protobuf:"varint,4,opt,name=update_out_queue_messages_cleared,json=updateOutQueueMessagesCleared" json:"update_out_queue_messages_cleared,omitempty"`
	// OutQueue Hi Size
	UpdateOutQueueSizeHigh uint32 `protobuf:"varint,5,opt,name=update_out_queue_size_high,json=updateOutQueueSizeHigh" json:"update_out_queue_size_high,omitempty"`
	// OutQueue Cumulative Size
	UpdateOutQueueSizeCumulative uint64 `protobuf:"varint,6,opt,name=update_out_queue_size_cumulative,json=updateOutQueueSizeCumulative" json:"update_out_queue_size_cumulative,omitempty"`
	// OutQueue Discarded Size
	UpdateOutQueueSizeDiscarded uint64 `protobuf:"varint,7,opt,name=update_out_queue_size_discarded,json=updateOutQueueSizeDiscarded" json:"update_out_queue_size_discarded,omitempty"`
	// OutQueue Cleared Size
	UpdateOutQueueSizeCleared uint64 `protobuf:"varint,8,opt,name=update_out_queue_size_cleared,json=updateOutQueueSizeCleared" json:"update_out_queue_size_cleared,omitempty"`
	// Last Discarded time
	LastUpdateDiscardTimestamp *BgpTimespec `protobuf:"bytes,9,opt,name=last_update_discard_timestamp,json=lastUpdateDiscardTimestamp" json:"last_update_discard_timestamp,omitempty"`
	// Time since last Discard event (in seconds)
	LastUpdateDiscardAge uint32 `protobuf:"varint,10,opt,name=last_update_discard_age,json=lastUpdateDiscardAge" json:"last_update_discard_age,omitempty"`
	// Last Cleared time
	LastUpdateClearedTimestamp *BgpTimespec `protobuf:"bytes,11,opt,name=last_update_cleared_timestamp,json=lastUpdateClearedTimestamp" json:"last_update_cleared_timestamp,omitempty"`
	// Time since last Clear event (in seconds)
	LastUpdateCleardAge uint32 `protobuf:"varint,12,opt,name=last_update_cleard_age,json=lastUpdateCleardAge" json:"last_update_cleard_age,omitempty"`
	// Throttle Count
	UpdateThrottleCount uint32 `protobuf:"varint,13,opt,name=update_throttle_count,json=updateThrottleCount" json:"update_throttle_count,omitempty"`
	// Last Throttled time
	LastUpdateThrottleTimestamp *BgpTimespec `protobuf:"bytes,14,opt,name=last_update_throttle_timestamp,json=lastUpdateThrottleTimestamp" json:"last_update_throttle_timestamp,omitempty"`
	// Time since last Throttle event (in seconds)
	LastUpdateThrottleAge uint32 `protobuf:"varint,15,opt,name=last_update_throttle_age,json=lastUpdateThrottleAge" json:"last_update_throttle_age,omitempty"`
	// Recovery Count
	UpdateRecoveryCount uint32 `protobuf:"varint,16,opt,name=update_recovery_count,json=updateRecoveryCount" json:"update_recovery_count,omitempty"`
	// Last Recovery time
	LastUpdateRecoveryTimestamp *BgpTimespec `protobuf:"bytes,17,opt,name=last_update_recovery_timestamp,json=lastUpdateRecoveryTimestamp" json:"last_update_recovery_timestamp,omitempty"`
	// Time since last Recovery event (in seconds)
	LastUpdateRecoveryAge uint32 `protobuf:"varint,18,opt,name=last_update_recovery_age,json=lastUpdateRecoveryAge" json:"last_update_recovery_age,omitempty"`
	// Memory allocation failure count
	UpdateMemoryAllocationFailCount uint32 `protobuf:"varint,19,opt,name=update_memory_allocation_fail_count,json=updateMemoryAllocationFailCount" json:"update_memory_allocation_fail_count,omitempty"`
	// Memory allocation failure time
	LastUpdateMemoryAllocationFailTimestamp *BgpTimespec `protobuf:"bytes,20,opt,name=last_update_memory_allocation_fail_timestamp,json=lastUpdateMemoryAllocationFailTimestamp" json:"last_update_memory_allocation_fail_timestamp,omitempty"`
	// Time since last memory allocation failure event (in seconds)
	LastUpdateMemoryAllocationFailAge uint32 `protobuf:"varint,21,opt,name=last_update_memory_allocation_fail_age,json=lastUpdateMemoryAllocationFailAge" json:"last_update_memory_allocation_fail_age,omitempty"`
}

func (m *BgpUpdgenStatsBag) Reset()                    { *m = BgpUpdgenStatsBag{} }
func (m *BgpUpdgenStatsBag) String() string            { return proto.CompactTextString(m) }
func (*BgpUpdgenStatsBag) ProtoMessage()               {}
func (*BgpUpdgenStatsBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueMessagesHigh() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessagesHigh
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueMessagesCumulative() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessagesCumulative
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueMessagesDiscarded() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessagesDiscarded
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueMessagesCleared() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessagesCleared
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueSizeHigh() uint32 {
	if m != nil {
		return m.UpdateOutQueueSizeHigh
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueSizeCumulative() uint64 {
	if m != nil {
		return m.UpdateOutQueueSizeCumulative
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueSizeDiscarded() uint64 {
	if m != nil {
		return m.UpdateOutQueueSizeDiscarded
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueSizeCleared() uint64 {
	if m != nil {
		return m.UpdateOutQueueSizeCleared
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetLastUpdateDiscardTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateDiscardTimestamp
	}
	return nil
}

func (m *BgpUpdgenStatsBag) GetLastUpdateDiscardAge() uint32 {
	if m != nil {
		return m.LastUpdateDiscardAge
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetLastUpdateClearedTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateClearedTimestamp
	}
	return nil
}

func (m *BgpUpdgenStatsBag) GetLastUpdateCleardAge() uint32 {
	if m != nil {
		return m.LastUpdateCleardAge
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateThrottleCount() uint32 {
	if m != nil {
		return m.UpdateThrottleCount
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetLastUpdateThrottleTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateThrottleTimestamp
	}
	return nil
}

func (m *BgpUpdgenStatsBag) GetLastUpdateThrottleAge() uint32 {
	if m != nil {
		return m.LastUpdateThrottleAge
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateRecoveryCount() uint32 {
	if m != nil {
		return m.UpdateRecoveryCount
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetLastUpdateRecoveryTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateRecoveryTimestamp
	}
	return nil
}

func (m *BgpUpdgenStatsBag) GetLastUpdateRecoveryAge() uint32 {
	if m != nil {
		return m.LastUpdateRecoveryAge
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateMemoryAllocationFailCount() uint32 {
	if m != nil {
		return m.UpdateMemoryAllocationFailCount
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetLastUpdateMemoryAllocationFailTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateMemoryAllocationFailTimestamp
	}
	return nil
}

func (m *BgpUpdgenStatsBag) GetLastUpdateMemoryAllocationFailAge() uint32 {
	if m != nil {
		return m.LastUpdateMemoryAllocationFailAge
	}
	return 0
}

func init() {
	proto.RegisterType((*BgpUpdgenUpdgrpBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.afs.af.update_generation_update_groups.update_generation_update_group.bgp_updgen_updgrp_bag_KEYS")
	proto.RegisterType((*BgpUpdgenUpdgrpBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.afs.af.update_generation_update_groups.update_generation_update_group.bgp_updgen_updgrp_bag")
	proto.RegisterType((*BgpTimespec)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.afs.af.update_generation_update_groups.update_generation_update_group.bgp_timespec")
	proto.RegisterType((*BgpUpdgenStatsBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.afs.af.update_generation_update_groups.update_generation_update_group.bgp_updgen_stats_bag")
}

func init() { proto.RegisterFile("bgp_updgen_updgrp_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x58, 0x5f, 0x73, 0xd4, 0x54,
	0x14, 0x9f, 0x20, 0x52, 0x38, 0xed, 0x02, 0x0d, 0x2d, 0xa4, 0x85, 0xc2, 0xb6, 0x28, 0x76, 0x46,
	0x5c, 0xc7, 0x16, 0x44, 0x50, 0x91, 0x85, 0xb6, 0xfc, 0x2f, 0xb0, 0x8b, 0x38, 0x3e, 0x65, 0xee,
	0x26, 0x67, 0xd3, 0x3b, 0x93, 0x4d, 0xe2, 0xbd, 0xc9, 0xd2, 0xf2, 0x2d, 0xfc, 0x04, 0x3e, 0xf8,
	0xe0, 0x07, 0x72, 0x7c, 0xf4, 0x45, 0xc7, 0x19, 0xc7, 0xf1, 0x3b, 0xe8, 0xdc, 0x7f, 0x49, 0x76,
	0x93, 0xad, 0x3e, 0xb2, 0x2f, 0xec, 0x72, 0xcf, 0xf9, 0x9d, 0xfb, 0xfb, 0xfd, 0xb2, 0xb9, 0xe7,
	0xdc, 0xc2, 0xf9, 0x5e, 0x90, 0xb8, 0x59, 0xe2, 0x07, 0x18, 0xc9, 0x0f, 0x96, 0xb8, 0x3d, 0x12,
	0xb4, 0x12, 0x16, 0xa7, 0xb1, 0xfd, 0xbd, 0xe5, 0x51, 0xee, 0xc5, 0x2e, 0x8d, 0xb9, 0xbb, 0xcf,
	0x5c, 0x9a, 0x0c, 0xaf, 0xb9, 0x22, 0x3f, 0x4e, 0x90, 0xb5, 0x7a, 0x41, 0xd2, 0xa2, 0x11, 0x4f,
	0x49, 0xe4, 0x21, 0xcf, 0xbf, 0xe5, 0x5f, 0x5c, 0xf1, 0xe1, 0xf7, 0x0e, 0x5a, 0x43, 0xd6, 0xe7,
	0xe2, 0x9f, 0x16, 0xe9, 0xf3, 0x16, 0xe9, 0xb7, 0xb2, 0xc4, 0x27, 0x29, 0xba, 0x01, 0x46, 0xc8,
	0x48, 0x4a, 0x63, 0xb9, 0xaf, 0x5c, 0x61, 0x71, 0x96, 0xf0, 0xff, 0x88, 0xaf, 0xfd, 0x68, 0xc1,
	0x72, 0x2d, 0x67, 0xf7, 0xf1, 0xf6, 0xb7, 0x5d, 0xfb, 0x32, 0x34, 0x72, 0x06, 0x11, 0x19, 0xa0,
	0x63, 0x35, 0xad, 0xf5, 0x13, 0x9d, 0x39, 0xb3, 0xb8, 0x4b, 0x06, 0x68, 0x2f, 0xc1, 0xf1, 0x21,
	0xeb, 0xab, 0xf8, 0x11, 0x19, 0x9f, 0x19, 0xb2, 0xbe, 0x0c, 0x9d, 0x83, 0x19, 0xa2, 0x23, 0xef,
	0xc8, 0xc8, 0x31, 0xa2, 0x02, 0x57, 0xc1, 0x2e, 0xf3, 0x70, 0x69, 0xe4, 0xe3, 0xbe, 0x73, 0xb4,
	0x69, 0xad, 0x37, 0x3a, 0xa7, 0x55, 0xe4, 0xbe, 0x08, 0x3c, 0x14, 0xeb, 0x6b, 0xbf, 0x2e, 0xc0,
	0x62, 0x2d, 0x4b, 0x7b, 0x05, 0x20, 0x61, 0xb1, 0x87, 0x9c, 0xbb, 0xd4, 0x77, 0x36, 0x24, 0xfe,
	0x84, 0x5e, 0x79, 0xe8, 0xdb, 0x57, 0xe0, 0x94, 0xde, 0x26, 0x67, 0xb8, 0x29, 0x79, 0x34, 0xd4,
	0xf2, 0x2b, 0xcd, 0xf3, 0x63, 0x58, 0x18, 0xa1, 0x63, 0x48, 0x5f, 0x93, 0xc9, 0xf3, 0x25, 0x42,
	0x6d, 0x05, 0xb8, 0x01, 0x4e, 0x84, 0x34, 0xd8, 0xeb, 0xc5, 0xcc, 0xe5, 0xc8, 0xb9, 0x30, 0xd6,
	0x80, 0xae, 0x4b, 0x16, 0x8b, 0x26, 0xde, 0x55, 0xe1, 0xf6, 0x61, 0xc2, 0x3f, 0xad, 0x17, 0x6e,
	0xb7, 0xe0, 0xcc, 0x48, 0x76, 0x3f, 0x24, 0x01, 0xdf, 0x70, 0x6e, 0xc8, 0xf4, 0x32, 0xad, 0x1d,
	0x19, 0xb0, 0x6f, 0xc2, 0x92, 0xce, 0x8f, 0xb3, 0xd4, 0xfd, 0x2e, 0xc3, 0x0c, 0xdd, 0x01, 0x72,
	0x4e, 0x02, 0xe4, 0xce, 0x67, 0x12, 0x75, 0x56, 0x25, 0x3c, 0xcb, 0xd2, 0x17, 0x22, 0xfc, 0x54,
	0x47, 0xed, 0x4f, 0x60, 0xb1, 0x02, 0xe5, 0xf4, 0x0d, 0x3a, 0x37, 0x25, 0xcc, 0x1e, 0x85, 0x75,
	0xe9, 0x1b, 0xb4, 0x37, 0x41, 0x17, 0x73, 0x79, 0xd6, 0xd3, 0x0c, 0xbd, 0x38, 0x8b, 0x52, 0xe7,
	0x96, 0xc4, 0x68, 0xee, 0xdd, 0xac, 0x27, 0x39, 0xde, 0x13, 0x21, 0x41, 0xb1, 0xc8, 0x4e, 0xf7,
	0x58, 0x9c, 0xa6, 0x21, 0xfa, 0x1a, 0xf7, 0xb9, 0xa2, 0xc8, 0x35, 0xe2, 0xa5, 0x09, 0x2b, 0xe8,
	0x75, 0x38, 0xc7, 0xb0, 0xcf, 0x90, 0xef, 0x55, 0x36, 0xfc, 0x42, 0x02, 0x17, 0x74, 0x78, 0x74,
	0xc7, 0x07, 0xb0, 0x5a, 0x85, 0x8d, 0xef, 0xfc, 0xa5, 0x2c, 0xb0, 0x32, 0x56, 0x60, 0x8c, 0xc0,
	0x55, 0xb0, 0xfb, 0x34, 0x4c, 0x91, 0x8d, 0xec, 0x7d, 0x5b, 0x3d, 0x3c, 0x15, 0x29, 0xed, 0xfb,
	0x3e, 0x9c, 0xcc, 0x7f, 0x23, 0x2a, 0xf3, 0x2b, 0x99, 0xd9, 0x30, 0xab, 0x2a, 0xed, 0x1a, 0x9c,
	0xcd, 0xd3, 0x42, 0x24, 0x43, 0x1a, 0x05, 0x3a, 0xfd, 0x8e, 0x12, 0x65, 0xa2, 0x4f, 0x54, 0x50,
	0xa1, 0x1e, 0xc3, 0x5a, 0xf5, 0xd5, 0x66, 0xe8, 0xc5, 0x43, 0x64, 0x07, 0x6e, 0x82, 0x91, 0x4f,
	0xa3, 0xc0, 0x69, 0x37, 0xad, 0xf5, 0xe3, 0x9d, 0x4b, 0xfa, 0x87, 0x92, 0x27, 0x76, 0x74, 0xde,
	0x73, 0x95, 0x66, 0xff, 0x65, 0xc1, 0x6a, 0x48, 0x78, 0x6a, 0xce, 0x86, 0x94, 0x0e, 0x90, 0x89,
	0x23, 0x87, 0xa5, 0xf2, 0x3b, 0x4f, 0xc9, 0x20, 0x71, 0xee, 0x36, 0xad, 0xf5, 0xd9, 0x8d, 0x1f,
	0xac, 0xd6, 0x5b, 0x77, 0x8c, 0x89, 0x6d, 0x15, 0xd1, 0x04, 0xbd, 0xce, 0x8a, 0x90, 0xf2, 0xb5,
	0x0c, 0xbf, 0x14, 0x42, 0xba, 0x42, 0xc7, 0x4b, 0x23, 0xc3, 0xfe, 0xd3, 0x82, 0x66, 0x9d, 0xd8,
	0x38, 0x29, 0x69, 0xbd, 0x37, 0x25, 0x5a, 0x2f, 0x54, 0xb4, 0xc6, 0x49, 0x21, 0xf5, 0x6f, 0x0b,
	0xd6, 0xaa, 0x52, 0x71, 0x3f, 0xa1, 0xec, 0xa0, 0x24, 0x76, 0x6b, 0x4a, 0xc4, 0x5e, 0x1c, 0x13,
	0xbb, 0x2d, 0x85, 0x14, 0x72, 0x6f, 0xc3, 0x85, 0x89, 0x6a, 0x49, 0x80, 0xce, 0xb6, 0x7c, 0x9f,
	0x9c, 0xda, 0x2a, 0xed, 0x00, 0xc5, 0xf9, 0x42, 0xf9, 0x28, 0x9a, 0x65, 0x51, 0x24, 0x5e, 0xa4,
	0x1d, 0xf9, 0x22, 0x2d, 0x50, 0x5e, 0x02, 0x76, 0x54, 0x6c, 0xc2, 0xdb, 0xc3, 0x70, 0x40, 0xa8,
	0x88, 0xbb, 0x43, 0x12, 0x66, 0xe8, 0xdc, 0x9f, 0xd2, 0xb7, 0xa7, 0x63, 0x74, 0xbc, 0x12, 0x32,
	0xec, 0xdf, 0x2c, 0x58, 0xa9, 0x8a, 0xf5, 0x31, 0x24, 0x07, 0x5a, 0xe8, 0x83, 0x29, 0x11, 0xba,
	0x34, 0x26, 0x74, 0x4b, 0x68, 0x50, 0x22, 0x7f, 0xb1, 0x60, 0xde, 0x74, 0xb6, 0x94, 0xa4, 0x94,
	0xa7, 0xd4, 0xe3, 0xce, 0x43, 0x29, 0xec, 0xa7, 0xb7, 0x55, 0x98, 0x1e, 0x8e, 0x04, 0x61, 0x2e,
	0x66, 0x23, 0x33, 0x4e, 0x74, 0x73, 0x05, 0xa2, 0xd5, 0x24, 0xc8, 0x06, 0x6e, 0x82, 0x21, 0x0e,
	0x5c, 0x8c, 0x64, 0x9b, 0x41, 0x86, 0xbe, 0xf3, 0x48, 0xb5, 0x1a, 0x11, 0x7d, 0x2e, 0x82, 0xdb,
	0x45, 0x4c, 0x74, 0xbd, 0x12, 0x8a, 0x84, 0x61, 0xfc, 0x1a, 0x7d, 0xe7, 0xb1, 0xea, 0x7a, 0x39,
	0xa2, 0xad, 0xd6, 0xc5, 0x50, 0x50, 0xca, 0x8e, 0xe2, 0x34, 0x47, 0x3c, 0x51, 0x43, 0x41, 0x8e,
	0xd8, 0x8d, 0x53, 0x03, 0xfa, 0x70, 0x64, 0x0b, 0xdc, 0x4f, 0xdc, 0xd7, 0x3e, 0x73, 0x9e, 0x4a,
	0xc0, 0xa9, 0x82, 0xd4, 0x7e, 0xf2, 0x8d, 0xcf, 0xec, 0x8f, 0xe0, 0x4c, 0x29, 0x99, 0x27, 0x19,
	0x93, 0xd9, 0xbb, 0x63, 0x84, 0xba, 0x49, 0xc6, 0x44, 0xfa, 0x2a, 0xcc, 0x51, 0xee, 0x8a, 0x65,
	0x12, 0x61, 0x94, 0x3a, 0xcf, 0xe4, 0xab, 0x3c, 0x4b, 0xf9, 0x73, 0xb3, 0xb4, 0xf6, 0x08, 0xe6,
	0xca, 0x3f, 0x0d, 0xdb, 0x81, 0x19, 0x8e, 0x5e, 0x1c, 0xf9, 0x5c, 0x0e, 0xbc, 0x8d, 0x8e, 0xf9,
	0xaf, 0xdd, 0x84, 0xd9, 0x88, 0x44, 0xb1, 0x89, 0x1e, 0x91, 0xd1, 0xf2, 0xd2, 0xda, 0xcf, 0xf3,
	0xb0, 0x50, 0xf7, 0x38, 0xec, 0xbb, 0x70, 0x71, 0xe2, 0x6c, 0xe6, 0xee, 0xd1, 0x60, 0x4f, 0xef,
	0xb5, 0x5c, 0x3f, 0xa0, 0x3d, 0xa0, 0xc1, 0x9e, 0xbd, 0x0b, 0xef, 0x4d, 0xae, 0xe1, 0x65, 0x83,
	0x2c, 0x24, 0x29, 0x1d, 0xa2, 0xe6, 0xd5, 0xac, 0xaf, 0x74, 0x2f, 0xcf, 0xb3, 0x9f, 0xc0, 0xe5,
	0xc9, 0xf5, 0x7c, 0xca, 0x3d, 0xc2, 0x7c, 0xf4, 0xe5, 0xec, 0xde, 0x30, 0x63, 0xc4, 0x78, 0xb9,
	0x2d, 0x93, 0x26, 0x06, 0xad, 0x43, 0xd8, 0x85, 0x48, 0xc4, 0x2f, 0x4d, 0xcd, 0xf8, 0x2b, 0x13,
	0xa8, 0xa9, 0x24, 0xfb, 0x16, 0x2c, 0xd7, 0x0e, 0xa3, 0xca, 0xa7, 0x77, 0xeb, 0x06, 0x59, 0x31,
	0x91, 0x4a, 0x8f, 0x76, 0xa0, 0x59, 0x8f, 0x2d, 0xf9, 0x73, 0xac, 0x69, 0xad, 0x1f, 0xed, 0x5c,
	0xa8, 0x56, 0x28, 0x79, 0xb3, 0x05, 0x97, 0xea, 0xeb, 0x14, 0xbe, 0xcc, 0xc8, 0x32, 0xe7, 0xab,
	0x65, 0x0a, 0x4f, 0xee, 0xc0, 0xca, 0x04, 0x36, 0xda, 0x8f, 0xe3, 0xb2, 0xc6, 0x52, 0x0d, 0x15,
	0xed, 0xc5, 0xef, 0x63, 0x27, 0xae, 0xde, 0xbe, 0xd4, 0xbf, 0x4f, 0x4c, 0xc9, 0x89, 0xbb, 0x5c,
	0x9c, 0xb8, 0xda, 0xa0, 0xa2, 0x77, 0x5f, 0x87, 0x73, 0x75, 0x22, 0x45, 0xdb, 0x06, 0x75, 0x36,
	0x55, 0xc0, 0xa2, 0x65, 0x8f, 0x9b, 0xa3, 0x5d, 0x2d, 0x99, 0x33, 0x3b, 0x7d, 0xe6, 0xe8, 0x27,
	0x5f, 0x98, 0xb3, 0x09, 0x67, 0x2b, 0x22, 0x95, 0x37, 0x73, 0xea, 0x4c, 0x1d, 0xc3, 0x4a, 0x6b,
	0x36, 0xf2, 0x0b, 0x9d, 0xb9, 0xeb, 0xe8, 0x6b, 0x45, 0xa3, 0x7c, 0x39, 0x33, 0x37, 0x1c, 0x75,
	0xab, 0xf8, 0xc3, 0x82, 0x8b, 0x23, 0xdd, 0xdd, 0x20, 0x0b, 0x3f, 0x4f, 0x4e, 0x89, 0x9f, 0xe7,
	0x4b, 0xed, 0x5d, 0xab, 0x28, 0x0c, 0xbd, 0x01, 0x4e, 0xad, 0x4c, 0x61, 0xe9, 0x29, 0x75, 0x7d,
	0xaf, 0xc2, 0x47, 0x4d, 0xcd, 0xef, 0x5a, 0xca, 0xd4, 0xd3, 0x65, 0x53, 0xcd, 0xfd, 0xaa, 0xde,
	0xd4, 0x1c, 0x59, 0x98, 0x3a, 0x3f, 0x7d, 0xa6, 0x1a, 0x8d, 0x13, 0x4d, 0xcd, 0x65, 0x0a, 0x53,
	0xed, 0x71, 0x53, 0x0d, 0x5c, 0x98, 0x5a, 0x74, 0xa1, 0x01, 0x0e, 0x62, 0x81, 0x08, 0xc3, 0xd8,
	0x53, 0x5c, 0xfa, 0x84, 0x86, 0xda, 0xe2, 0x33, 0xe5, 0x2e, 0xf4, 0x54, 0x66, 0xb6, 0xf3, 0xc4,
	0x1d, 0x42, 0x43, 0x65, 0xf7, 0x3f, 0x16, 0x5c, 0x2d, 0xf3, 0x98, 0x50, 0xb3, 0x30, 0x7f, 0x61,
	0x4a, 0xcc, 0xff, 0xa0, 0x70, 0xaf, 0x4e, 0x7d, 0xf1, 0x20, 0x5e, 0xc0, 0x95, 0xff, 0x61, 0x80,
	0x78, 0x2c, 0x8b, 0xd2, 0xd2, 0xd5, 0xc3, 0x0b, 0xb7, 0x03, 0xec, 0x1d, 0x93, 0x7f, 0xc2, 0xdc,
	0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xba, 0x8d, 0xb0, 0xe1, 0x14, 0x00, 0x00,
}
