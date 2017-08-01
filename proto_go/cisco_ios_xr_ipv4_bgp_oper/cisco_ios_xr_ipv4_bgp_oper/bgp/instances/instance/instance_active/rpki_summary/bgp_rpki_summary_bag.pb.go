// Code generated by protoc-gen-go.
// source: bgp_rpki_summary_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_rpki_summary is a generated protocol buffer package.

It is generated from these files:
	bgp_rpki_summary_bag.proto

It has these top-level messages:
	BgpRpkiSummaryBag_KEYS
	BgpRpkiSummaryBag
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_rpki_summary

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

type BgpRpkiSummaryBag_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
}

func (m *BgpRpkiSummaryBag_KEYS) Reset()                    { *m = BgpRpkiSummaryBag_KEYS{} }
func (m *BgpRpkiSummaryBag_KEYS) String() string            { return proto.CompactTextString(m) }
func (*BgpRpkiSummaryBag_KEYS) ProtoMessage()               {}
func (*BgpRpkiSummaryBag_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BgpRpkiSummaryBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

type BgpRpkiSummaryBag struct {
	// Number of RPKI Servers configured
	Servers uint32 `protobuf:"varint,50,opt,name=servers" json:"servers,omitempty"`
	// Number of IPv4 ROA Nets
	Ipv4RoaNets uint32 `protobuf:"varint,51,opt,name=ipv4_roa_nets,json=ipv4RoaNets" json:"ipv4_roa_nets,omitempty"`
	// Number of IPv4 ROA Paths
	Ipv4RoaPaths uint32 `protobuf:"varint,52,opt,name=ipv4_roa_paths,json=ipv4RoaPaths" json:"ipv4_roa_paths,omitempty"`
	// Number of IPv6 ROA Nets
	Ipv6RoaNets uint32 `protobuf:"varint,53,opt,name=ipv6_roa_nets,json=ipv6RoaNets" json:"ipv6_roa_nets,omitempty"`
	// Number of IPv6 ROA Paths
	Ipv6RoaPaths uint32 `protobuf:"varint,54,opt,name=ipv6_roa_paths,json=ipv6RoaPaths" json:"ipv6_roa_paths,omitempty"`
	// RPKI Knob disabled
	RpkiDisabled bool `protobuf:"varint,55,opt,name=rpki_disabled,json=rpkiDisabled" json:"rpki_disabled,omitempty"`
	// Use RPKI validity for bestpath calculation
	RpkiUseValidity bool `protobuf:"varint,56,opt,name=rpki_use_validity,json=rpkiUseValidity" json:"rpki_use_validity,omitempty"`
	// Allow invalid paths
	RpkiAllowInvalid bool `protobuf:"varint,57,opt,name=rpki_allow_invalid,json=rpkiAllowInvalid" json:"rpki_allow_invalid,omitempty"`
	// Signal RPKI validity to iBGP peers
	RpkiSignalIbgp bool `protobuf:"varint,58,opt,name=rpki_signal_ibgp,json=rpkiSignalIbgp" json:"rpki_signal_ibgp,omitempty"`
}

func (m *BgpRpkiSummaryBag) Reset()                    { *m = BgpRpkiSummaryBag{} }
func (m *BgpRpkiSummaryBag) String() string            { return proto.CompactTextString(m) }
func (*BgpRpkiSummaryBag) ProtoMessage()               {}
func (*BgpRpkiSummaryBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BgpRpkiSummaryBag) GetServers() uint32 {
	if m != nil {
		return m.Servers
	}
	return 0
}

func (m *BgpRpkiSummaryBag) GetIpv4RoaNets() uint32 {
	if m != nil {
		return m.Ipv4RoaNets
	}
	return 0
}

func (m *BgpRpkiSummaryBag) GetIpv4RoaPaths() uint32 {
	if m != nil {
		return m.Ipv4RoaPaths
	}
	return 0
}

func (m *BgpRpkiSummaryBag) GetIpv6RoaNets() uint32 {
	if m != nil {
		return m.Ipv6RoaNets
	}
	return 0
}

func (m *BgpRpkiSummaryBag) GetIpv6RoaPaths() uint32 {
	if m != nil {
		return m.Ipv6RoaPaths
	}
	return 0
}

func (m *BgpRpkiSummaryBag) GetRpkiDisabled() bool {
	if m != nil {
		return m.RpkiDisabled
	}
	return false
}

func (m *BgpRpkiSummaryBag) GetRpkiUseValidity() bool {
	if m != nil {
		return m.RpkiUseValidity
	}
	return false
}

func (m *BgpRpkiSummaryBag) GetRpkiAllowInvalid() bool {
	if m != nil {
		return m.RpkiAllowInvalid
	}
	return false
}

func (m *BgpRpkiSummaryBag) GetRpkiSignalIbgp() bool {
	if m != nil {
		return m.RpkiSignalIbgp
	}
	return false
}

func init() {
	proto.RegisterType((*BgpRpkiSummaryBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.rpki_summary.bgp_rpki_summary_bag_KEYS")
	proto.RegisterType((*BgpRpkiSummaryBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.rpki_summary.bgp_rpki_summary_bag")
}

func init() { proto.RegisterFile("bgp_rpki_summary_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0x80, 0xe9, 0xef, 0xf0, 0x53, 0xe3, 0x3a, 0x67, 0xf0, 0x10, 0x3d, 0x8d, 0xcd, 0x43, 0x11,
	0xe9, 0xc1, 0xcd, 0xfa, 0xe7, 0xa4, 0xa0, 0x87, 0x21, 0x0c, 0xe9, 0x50, 0xf0, 0x14, 0xd2, 0x2e,
	0xd4, 0x60, 0xdb, 0x84, 0xbc, 0x59, 0x75, 0x1f, 0xda, 0xef, 0x20, 0xc9, 0xba, 0x38, 0x61, 0xb7,
	0xb7, 0xcf, 0xfb, 0xf4, 0xe9, 0x1f, 0x82, 0x4e, 0xb2, 0x42, 0x51, 0xad, 0x3e, 0x04, 0x85, 0x45,
	0x55, 0x31, 0xbd, 0xa4, 0x19, 0x2b, 0x62, 0xa5, 0xa5, 0x91, 0x78, 0x9a, 0x0b, 0xc8, 0x25, 0x15,
	0x12, 0xe8, 0x97, 0xa6, 0x42, 0x35, 0x63, 0x6a, 0x6d, 0xa9, 0xb8, 0x8e, 0xb3, 0x42, 0xc5, 0xa2,
	0x06, 0xc3, 0xea, 0x9c, 0x83, 0x9f, 0xfc, 0x40, 0x59, 0x6e, 0x44, 0xc3, 0xe3, 0xcd, 0xf2, 0xe0,
	0x0e, 0x1d, 0x6f, 0x7b, 0x1a, 0x7d, 0x7a, 0x7c, 0x9b, 0xe1, 0x21, 0x0a, 0xfd, 0xcd, 0x35, 0xab,
	0x38, 0x09, 0xfa, 0x41, 0xb4, 0x97, 0x76, 0xd6, 0x70, 0xca, 0x2a, 0x3e, 0xf8, 0xfe, 0x87, 0x8e,
	0xb6, 0x25, 0x30, 0x41, 0x3b, 0xc0, 0x75, 0xc3, 0x35, 0x90, 0x8b, 0x7e, 0x10, 0x85, 0xe9, 0xfa,
	0x12, 0x0f, 0x50, 0xe8, 0xde, 0x5c, 0x4b, 0x46, 0x6b, 0x6e, 0x80, 0x8c, 0xdc, 0x7e, 0xdf, 0xc2,
	0x54, 0xb2, 0x29, 0x37, 0x80, 0x4f, 0x51, 0xd7, 0x3b, 0x8a, 0x99, 0x77, 0x20, 0x63, 0x27, 0x75,
	0x5a, 0xe9, 0xd9, 0xb2, 0xb6, 0x94, 0xfc, 0x96, 0x2e, 0x7d, 0x29, 0xf9, 0x5b, 0x4a, 0x36, 0x4a,
	0x89, 0x2f, 0x25, 0xbe, 0x34, 0x44, 0xa1, 0xfb, 0x82, 0xb9, 0x00, 0x96, 0x95, 0x7c, 0x4e, 0xae,
	0xfa, 0x41, 0xb4, 0x9b, 0x76, 0x2c, 0x7c, 0x68, 0x19, 0x3e, 0x43, 0x87, 0x4e, 0x5a, 0x00, 0xa7,
	0x0d, 0x2b, 0xc5, 0x5c, 0x98, 0x25, 0xb9, 0x76, 0xe2, 0x81, 0x5d, 0xbc, 0x00, 0x7f, 0x6d, 0x31,
	0x3e, 0x47, 0xd8, 0xb9, 0xac, 0x2c, 0xe5, 0x27, 0x15, 0xb5, 0xf3, 0xc9, 0x8d, 0x93, 0x7b, 0x76,
	0x73, 0x6f, 0x17, 0x93, 0x15, 0xc7, 0x11, 0xea, 0xad, 0x7e, 0xa0, 0x28, 0x6a, 0x56, 0x52, 0x91,
	0x15, 0x8a, 0xdc, 0x3a, 0xb7, 0x6b, 0xf9, 0xcc, 0xe1, 0x49, 0x56, 0xa8, 0xec, 0xbf, 0x3b, 0x08,
	0xa3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xcf, 0x36, 0x1a, 0x26, 0x02, 0x00, 0x00,
}
