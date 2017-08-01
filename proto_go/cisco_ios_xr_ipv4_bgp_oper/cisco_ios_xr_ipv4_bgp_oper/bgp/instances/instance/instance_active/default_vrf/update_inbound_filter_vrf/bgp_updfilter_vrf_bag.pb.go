// Code generated by protoc-gen-go.
// source: bgp_updfilter_vrf_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_default_vrf_update_inbound_filter_vrf is a generated protocol buffer package.

It is generated from these files:
	bgp_updfilter_vrf_bag.proto

It has these top-level messages:
	BgpUpdfilterVrfBag_KEYS
	BgpUpdfilterVrfBag
	BgpTimespec
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_default_vrf_update_inbound_filter_vrf

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
	proto.RegisterType((*BgpUpdfilterVrfBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.update_inbound_filter_vrf.bgp_updfilter_vrf_bag_KEYS")
	proto.RegisterType((*BgpUpdfilterVrfBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.update_inbound_filter_vrf.bgp_updfilter_vrf_bag")
	proto.RegisterType((*BgpTimespec)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.update_inbound_filter_vrf.bgp_timespec")
}

func init() { proto.RegisterFile("bgp_updfilter_vrf_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xa9, 0x82, 0x62, 0xb6, 0x22, 0x04, 0x06, 0x65, 0xc3, 0x51, 0x26, 0xc8, 0x4e, 0x3d,
	0x6c, 0x53, 0xcf, 0x63, 0xe8, 0x41, 0x71, 0x87, 0xfa, 0x07, 0x3c, 0xbd, 0xa4, 0xed, 0xdb, 0x18,
	0x58, 0x93, 0xd0, 0xa4, 0xc3, 0xcf, 0x26, 0xf8, 0xdd, 0xa4, 0xff, 0xe6, 0x36, 0x77, 0xf6, 0x96,
	0x3e, 0x79, 0xfa, 0x7b, 0x9e, 0x97, 0x37, 0x64, 0x10, 0x71, 0x0d, 0x85, 0x4e, 0x52, 0xb1, 0xb2,
	0x98, 0xc3, 0x3a, 0x4f, 0x21, 0x62, 0x3c, 0xd0, 0xb9, 0xb2, 0x8a, 0xf2, 0x58, 0x98, 0x58, 0x81,
	0x50, 0x06, 0x3e, 0x73, 0x10, 0x7a, 0x3d, 0x83, 0xd2, 0xae, 0x34, 0xe6, 0x41, 0xc4, 0x75, 0x20,
	0xa4, 0xb1, 0x4c, 0xc6, 0x68, 0x36, 0xa7, 0xcd, 0x01, 0x58, 0x6c, 0xc5, 0x1a, 0x83, 0x04, 0x53,
	0x56, 0xac, 0x6c, 0x09, 0x0e, 0x0a, 0x9d, 0x30, 0x8b, 0x20, 0x64, 0xa4, 0x0a, 0x99, 0xc0, 0x6f,
	0xe4, 0x68, 0x4e, 0xfa, 0x07, 0x7b, 0xc0, 0xe3, 0xdd, 0xfb, 0x33, 0xbd, 0x24, 0xee, 0x06, 0x2b,
	0x59, 0x86, 0x9e, 0xe3, 0x3b, 0xe3, 0xb3, 0xb0, 0xdb, 0x8a, 0x4b, 0x96, 0xe1, 0xe8, 0xfb, 0x98,
	0xf4, 0x0e, 0x32, 0xe8, 0x15, 0x39, 0x6f, 0x92, 0x4b, 0xa5, 0x02, 0x4c, 0x2a, 0x80, 0x5b, 0xcb,
	0x6f, 0x79, 0x5a, 0x12, 0xe8, 0x9c, 0x5c, 0x34, 0xbe, 0x1a, 0x80, 0x09, 0x64, 0x68, 0x0c, 0xe3,
	0x08, 0xb1, 0x2a, 0xa4, 0xf5, 0xa6, 0xbe, 0x33, 0x76, 0xc3, 0x7e, 0x6d, 0xba, 0x6f, 0x3c, 0x4f,
	0xb5, 0x65, 0x51, 0x3a, 0xe8, 0x82, 0x0c, 0xf7, 0x11, 0x12, 0x05, 0xff, 0x88, 0x54, 0xde, 0x30,
	0x66, 0x15, 0x63, 0xb0, 0xcb, 0x58, 0x36, 0x9e, 0x1a, 0xf2, 0xe5, 0x90, 0xe1, 0x8a, 0x19, 0x0b,
	0xfb, 0x28, 0x2b, 0x32, 0x34, 0x96, 0x65, 0xda, 0xbb, 0xf6, 0x9d, 0x71, 0x67, 0x52, 0x04, 0xff,
	0xb4, 0x9f, 0x12, 0x56, 0xa7, 0x6b, 0x8c, 0xc3, 0x41, 0x59, 0xee, 0x75, 0x67, 0x80, 0x97, 0xb6,
	0x19, 0xbd, 0x25, 0xde, 0xc1, 0xee, 0x8c, 0xa3, 0x77, 0x53, 0xcd, 0xde, 0xfb, 0xfb, 0xfb, 0x9c,
	0xe3, 0xe8, 0x81, 0x74, 0xb7, 0x53, 0xa8, 0x47, 0x4e, 0x0d, 0xc6, 0x4a, 0x26, 0xa6, 0x5a, 0xb7,
	0x1b, 0xb6, 0x9f, 0xd4, 0x27, 0x1d, 0xc9, 0xa4, 0x6a, 0x6f, 0x8f, 0xaa, 0xdb, 0x6d, 0x29, 0x3a,
	0xa9, 0x9e, 0xef, 0xf4, 0x27, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x35, 0x46, 0x08, 0xdd, 0x02, 0x00,
	0x00,
}
