// Code generated by protoc-gen-go.
// source: bgp_updfilter_vrf_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_vrfs_vrf_update_inbound_filter_vrf is a generated protocol buffer package.

It is generated from these files:
	bgp_updfilter_vrf_bag.proto

It has these top-level messages:
	BgpUpdfilterVrfBag_KEYS
	BgpUpdfilterVrfBag
	BgpTimespec
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_vrfs_vrf_update_inbound_filter_vrf

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

// BGP Update filtering VRF information
type BgpUpdfilterVrfBag_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	VrfName      string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
}

func (m *BgpUpdfilterVrfBag_KEYS) Reset()                    { *m = BgpUpdfilterVrfBag_KEYS{} }
func (m *BgpUpdfilterVrfBag_KEYS) String() string            { return proto.CompactTextString(m) }
func (*BgpUpdfilterVrfBag_KEYS) ProtoMessage()               {}
func (*BgpUpdfilterVrfBag_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BgpUpdfilterVrfBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpUpdfilterVrfBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type BgpUpdfilterVrfBag struct {
	// VRF Name
	UpdateVrfName string `protobuf:"bytes,50,opt,name=update_vrf_name,json=updateVrfName" json:"update_vrf_name,omitempty"`
	// Filtered messages count
	UpdateFilteredMessageCount uint32 `protobuf:"varint,51,opt,name=update_filtered_message_count,json=updateFilteredMessageCount" json:"update_filtered_message_count,omitempty"`
	// Count of neighbors that received filtered messages
	UpdateFilteredNeighborCount uint32 `protobuf:"varint,52,opt,name=update_filtered_neighbor_count,json=updateFilteredNeighborCount" json:"update_filtered_neighbor_count,omitempty"`
	// Last Filtered messages received time: time elapsed since 00:00:00 UTC, January 1, 1970
	LastUpdateFilteredTimestamp *BgpTimespec `protobuf:"bytes,53,opt,name=last_update_filtered_timestamp,json=lastUpdateFilteredTimestamp" json:"last_update_filtered_timestamp,omitempty"`
	// Time since last filtered messages received event (in seconds)
	LastUpdateFilteredAge uint32 `protobuf:"varint,54,opt,name=last_update_filtered_age,json=lastUpdateFilteredAge" json:"last_update_filtered_age,omitempty"`
}

func (m *BgpUpdfilterVrfBag) Reset()                    { *m = BgpUpdfilterVrfBag{} }
func (m *BgpUpdfilterVrfBag) String() string            { return proto.CompactTextString(m) }
func (*BgpUpdfilterVrfBag) ProtoMessage()               {}
func (*BgpUpdfilterVrfBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BgpUpdfilterVrfBag) GetUpdateVrfName() string {
	if m != nil {
		return m.UpdateVrfName
	}
	return ""
}

func (m *BgpUpdfilterVrfBag) GetUpdateFilteredMessageCount() uint32 {
	if m != nil {
		return m.UpdateFilteredMessageCount
	}
	return 0
}

func (m *BgpUpdfilterVrfBag) GetUpdateFilteredNeighborCount() uint32 {
	if m != nil {
		return m.UpdateFilteredNeighborCount
	}
	return 0
}

func (m *BgpUpdfilterVrfBag) GetLastUpdateFilteredTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateFilteredTimestamp
	}
	return nil
}

func (m *BgpUpdfilterVrfBag) GetLastUpdateFilteredAge() uint32 {
	if m != nil {
		return m.LastUpdateFilteredAge
	}
	return 0
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

func init() {
	proto.RegisterType((*BgpUpdfilterVrfBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.update_inbound_filter_vrf.bgp_updfilter_vrf_bag_KEYS")
	proto.RegisterType((*BgpUpdfilterVrfBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.update_inbound_filter_vrf.bgp_updfilter_vrf_bag")
	proto.RegisterType((*BgpTimespec)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.update_inbound_filter_vrf.bgp_timespec")
}

func init() { proto.RegisterFile("bgp_updfilter_vrf_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x5f, 0x6b, 0xdb, 0x30,
	0x14, 0xc5, 0x71, 0x06, 0xcb, 0xa6, 0xc4, 0x0c, 0x04, 0x01, 0x2f, 0x61, 0xc1, 0x64, 0x30, 0xf2,
	0xe4, 0x87, 0x24, 0xdb, 0x9e, 0x43, 0xd8, 0x1e, 0x36, 0x96, 0x07, 0x6f, 0x2b, 0x14, 0x0a, 0x42,
	0xb6, 0xaf, 0x5d, 0x41, 0x2c, 0x09, 0x49, 0x36, 0xed, 0x37, 0x2b, 0xf4, 0xcb, 0x15, 0xc9, 0x7f,
	0x9a, 0xa4, 0x79, 0xee, 0x8b, 0x2d, 0xdf, 0x73, 0xee, 0xef, 0x1e, 0xa3, 0x8b, 0x66, 0x49, 0x21,
	0x49, 0x25, 0xb3, 0x9c, 0x1d, 0x0c, 0x28, 0x52, 0xab, 0x9c, 0x24, 0xb4, 0x88, 0xa4, 0x12, 0x46,
	0x60, 0x48, 0x99, 0x4e, 0x05, 0x61, 0x42, 0x93, 0x3b, 0x45, 0x98, 0xac, 0x37, 0xc4, 0xda, 0x85,
	0x04, 0x15, 0x25, 0x85, 0x8c, 0x18, 0xd7, 0x86, 0xf2, 0x14, 0x74, 0x7f, 0xea, 0x0f, 0xc4, 0xbe,
	0xb2, 0xe4, 0x3e, 0xaa, 0x55, 0xae, 0xed, 0x23, 0xaa, 0x64, 0x46, 0x0d, 0x10, 0xc6, 0x13, 0x51,
	0xf1, 0x8c, 0x3c, 0x0f, 0x5c, 0xdc, 0xa0, 0xe9, 0xc5, 0x14, 0xe4, 0xf7, 0x8f, 0xeb, 0xbf, 0xf8,
	0x33, 0xf2, 0x7b, 0x28, 0xa7, 0x25, 0x04, 0x5e, 0xe8, 0x2d, 0xdf, 0xc7, 0xe3, 0xae, 0xb8, 0xa7,
	0x25, 0xe0, 0x8f, 0xe8, 0x9d, 0x6d, 0x72, 0xfa, 0xc0, 0xe9, 0xc3, 0x5a, 0xe5, 0x56, 0x5a, 0x3c,
	0xbe, 0x41, 0x93, 0x8b, 0x78, 0xfc, 0x05, 0x7d, 0x68, 0x43, 0xf5, 0xbd, 0x2b, 0xd7, 0xeb, 0x37,
	0xe5, 0xab, 0x86, 0x80, 0xb7, 0xe8, 0x53, 0xeb, 0x6b, 0x00, 0x90, 0x91, 0x12, 0xb4, 0xa6, 0x05,
	0x90, 0x54, 0x54, 0xdc, 0x04, 0xeb, 0xd0, 0x5b, 0xfa, 0xf1, 0xb4, 0x31, 0xfd, 0x6c, 0x3d, 0x7f,
	0x1a, 0xcb, 0xce, 0x3a, 0xf0, 0x0e, 0xcd, 0xcf, 0x11, 0x1c, 0x58, 0x71, 0x9b, 0x08, 0xd5, 0x32,
	0x36, 0x8e, 0x31, 0x3b, 0x65, 0xec, 0x5b, 0x4f, 0x03, 0x79, 0xf0, 0xd0, 0xfc, 0x40, 0xb5, 0x21,
	0xe7, 0x28, 0xc3, 0x4a, 0xd0, 0x86, 0x96, 0x32, 0xf8, 0x1a, 0x7a, 0xcb, 0xd1, 0x4a, 0x47, 0xaf,
	0x72, 0x71, 0x16, 0xd5, 0xcc, 0x96, 0x90, 0xc6, 0x33, 0x1b, 0xed, 0xff, 0x49, 0xfc, 0x7f, 0x5d,
	0x2e, 0xfc, 0x1d, 0x05, 0x17, 0x93, 0xd3, 0x02, 0x82, 0x6f, 0xee, 0xcf, 0x27, 0x2f, 0xdb, 0xb7,
	0x05, 0x2c, 0x7e, 0xa1, 0xf1, 0xf1, 0x14, 0x1c, 0xa0, 0xa1, 0x86, 0x54, 0xf0, 0x4c, 0xbb, 0x3d,
	0xf0, 0xe3, 0xee, 0x13, 0x87, 0x68, 0xc4, 0x29, 0x17, 0x9d, 0x3a, 0x70, 0xea, 0x71, 0x29, 0x79,
	0xeb, 0xb6, 0x7a, 0xfd, 0x14, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xb1, 0x5d, 0xce, 0xf4, 0x02, 0x00,
	0x00,
}
