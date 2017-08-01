// Code generated by protoc-gen-go.
// source: bgp_upderr_vrf_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_vrfs_vrf_update_inbound_error_vrf is a generated protocol buffer package.

It is generated from these files:
	bgp_upderr_vrf_bag.proto

It has these top-level messages:
	BgpUpderrVrfBag_KEYS
	BgpUpderrVrfBag
	BgpTimespec
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_vrfs_vrf_update_inbound_error_vrf

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

// BGP Update error-handling VRF information
type BgpUpderrVrfBag_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	VrfName      string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
}

func (m *BgpUpderrVrfBag_KEYS) Reset()                    { *m = BgpUpderrVrfBag_KEYS{} }
func (m *BgpUpderrVrfBag_KEYS) String() string            { return proto.CompactTextString(m) }
func (*BgpUpderrVrfBag_KEYS) ProtoMessage()               {}
func (*BgpUpderrVrfBag_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BgpUpderrVrfBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpUpderrVrfBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type BgpUpderrVrfBag struct {
	// VRF Name
	UpdateVrfName string `protobuf:"bytes,50,opt,name=update_vrf_name,json=updateVrfName" json:"update_vrf_name,omitempty"`
	// Malformed messages count
	UpdateMalformedMessageCount uint32 `protobuf:"varint,51,opt,name=update_malformed_message_count,json=updateMalformedMessageCount" json:"update_malformed_message_count,omitempty"`
	// Count of neighbors that received malformed messages
	UpdateMalformedNeighborCount uint32 `protobuf:"varint,52,opt,name=update_malformed_neighbor_count,json=updateMalformedNeighborCount" json:"update_malformed_neighbor_count,omitempty"`
	// Last malformed messages received time: time elapsed since 00:00:00 UTC, January 1, 1970
	LastUpdateMalformedTimestamp *BgpTimespec `protobuf:"bytes,53,opt,name=last_update_malformed_timestamp,json=lastUpdateMalformedTimestamp" json:"last_update_malformed_timestamp,omitempty"`
	// Time since last malformed messages received event (in seconds)
	LastUpdateMalformedAge uint32 `protobuf:"varint,54,opt,name=last_update_malformed_age,json=lastUpdateMalformedAge" json:"last_update_malformed_age,omitempty"`
}

func (m *BgpUpderrVrfBag) Reset()                    { *m = BgpUpderrVrfBag{} }
func (m *BgpUpderrVrfBag) String() string            { return proto.CompactTextString(m) }
func (*BgpUpderrVrfBag) ProtoMessage()               {}
func (*BgpUpderrVrfBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BgpUpderrVrfBag) GetUpdateVrfName() string {
	if m != nil {
		return m.UpdateVrfName
	}
	return ""
}

func (m *BgpUpderrVrfBag) GetUpdateMalformedMessageCount() uint32 {
	if m != nil {
		return m.UpdateMalformedMessageCount
	}
	return 0
}

func (m *BgpUpderrVrfBag) GetUpdateMalformedNeighborCount() uint32 {
	if m != nil {
		return m.UpdateMalformedNeighborCount
	}
	return 0
}

func (m *BgpUpderrVrfBag) GetLastUpdateMalformedTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateMalformedTimestamp
	}
	return nil
}

func (m *BgpUpderrVrfBag) GetLastUpdateMalformedAge() uint32 {
	if m != nil {
		return m.LastUpdateMalformedAge
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
	proto.RegisterType((*BgpUpderrVrfBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.update_inbound_error_vrf.bgp_upderr_vrf_bag_KEYS")
	proto.RegisterType((*BgpUpderrVrfBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.update_inbound_error_vrf.bgp_upderr_vrf_bag")
	proto.RegisterType((*BgpTimespec)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.update_inbound_error_vrf.bgp_timespec")
}

func init() { proto.RegisterFile("bgp_upderr_vrf_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x4f, 0x4b, 0xeb, 0x40,
	0x14, 0xc5, 0x49, 0x1f, 0xbc, 0xbe, 0x37, 0x6d, 0x10, 0x66, 0xa1, 0x29, 0x16, 0x1b, 0x2a, 0x48,
	0x57, 0x59, 0xb4, 0x55, 0x70, 0x29, 0xa5, 0x1b, 0xa5, 0x5d, 0xc4, 0x3f, 0xd0, 0xd5, 0x30, 0x49,
	0x6e, 0x62, 0xa0, 0x99, 0x19, 0x66, 0x92, 0xa0, 0x1f, 0xcd, 0x95, 0x5f, 0x4d, 0x66, 0xf2, 0x87,
	0x6a, 0xba, 0x76, 0x93, 0x4c, 0xee, 0x39, 0xf7, 0x77, 0x4f, 0x98, 0x8b, 0x9c, 0x20, 0x11, 0xa4,
	0x10, 0x11, 0x48, 0x49, 0x4a, 0x19, 0x93, 0x80, 0x26, 0x9e, 0x90, 0x3c, 0xe7, 0x38, 0x0a, 0x53,
	0x15, 0x72, 0x92, 0x72, 0x45, 0xde, 0x24, 0x49, 0x45, 0xb9, 0x24, 0xda, 0xcb, 0x05, 0x48, 0x2f,
	0x48, 0x84, 0x97, 0x32, 0x95, 0x53, 0x16, 0x82, 0x6a, 0x4f, 0xed, 0x81, 0xe8, 0x57, 0x14, 0xbc,
	0x7b, 0xa5, 0x8c, 0x95, 0x7e, 0x78, 0x85, 0x88, 0x68, 0x0e, 0x24, 0x65, 0x01, 0x2f, 0x58, 0x44,
	0x40, 0x4a, 0x6e, 0xe6, 0x4d, 0x77, 0xe8, 0xac, 0x9b, 0x80, 0x3c, 0xac, 0x77, 0x8f, 0xf8, 0x12,
	0xd9, 0x2d, 0x90, 0xd1, 0x0c, 0x1c, 0xcb, 0xb5, 0x66, 0xff, 0xfd, 0x61, 0x53, 0xdc, 0xd2, 0x0c,
	0xf0, 0x08, 0xfd, 0xd3, 0x4d, 0x46, 0xef, 0x19, 0xbd, 0x5f, 0xca, 0x58, 0x4b, 0xd3, 0xcf, 0x3f,
	0x08, 0x77, 0xd9, 0xf8, 0x0a, 0x9d, 0xd4, 0x69, 0xda, 0xc6, 0xb9, 0x69, 0xb4, 0xab, 0xf2, 0x4b,
	0xd5, 0x8e, 0x57, 0xe8, 0xa2, 0xf6, 0x65, 0x74, 0x1f, 0x73, 0x99, 0x41, 0x44, 0x32, 0x50, 0x8a,
	0x26, 0x40, 0x42, 0x5e, 0xb0, 0xdc, 0x59, 0xb8, 0xd6, 0xcc, 0xf6, 0xcf, 0x2b, 0xd7, 0xa6, 0x31,
	0x6d, 0x2a, 0xcf, 0x4a, 0x5b, 0xf0, 0x1a, 0x4d, 0x3a, 0x10, 0x06, 0x69, 0xf2, 0x1a, 0x70, 0x59,
	0x53, 0x96, 0x86, 0x32, 0xfe, 0x41, 0xd9, 0xd6, 0xa6, 0x0a, 0xf3, 0x61, 0xa1, 0xc9, 0x9e, 0xaa,
	0x9c, 0x74, 0x60, 0x79, 0x9a, 0x81, 0xca, 0x69, 0x26, 0x9c, 0x6b, 0xd7, 0x9a, 0x0d, 0xe6, 0xd2,
	0xfb, 0x8d, 0x6b, 0xd3, 0xa4, 0x6a, 0xb4, 0x80, 0xd0, 0x1f, 0xeb, 0x68, 0xcf, 0xdf, 0xf3, 0x3f,
	0x35, 0xb9, 0xf0, 0x2d, 0x1a, 0x1d, 0x8f, 0x4e, 0x13, 0x70, 0x6e, 0xcc, 0xcf, 0x9f, 0x1e, 0x01,
	0xdc, 0x25, 0x30, 0xbd, 0x47, 0xc3, 0xc3, 0x41, 0xd8, 0x41, 0x7d, 0x05, 0x21, 0x67, 0x91, 0x32,
	0xbb, 0x60, 0xfb, 0xcd, 0x27, 0x76, 0xd1, 0x80, 0x51, 0xc6, 0x1b, 0xb5, 0x67, 0xd4, 0xc3, 0x52,
	0xf0, 0xd7, 0x6c, 0xf5, 0xe2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x03, 0x83, 0xec, 0xf1, 0x02,
	0x00, 0x00,
}
