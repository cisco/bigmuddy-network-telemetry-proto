// Code generated by protoc-gen-go.
// source: bgp_nbr_adv_cnt_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_vrfs_vrf_afs_af_advertised_path_counts_advertised_path_count is a generated protocol buffer package.

It is generated from these files:
	bgp_nbr_adv_cnt_bag.proto

It has these top-level messages:
	BgpNbrAdvCntBag_KEYS
	BgpNbrAdvCntBag
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_vrfs_vrf_afs_af_advertised_path_counts_advertised_path_count

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

// BGP Neighbor adv cnt bag
type BgpNbrAdvCntBag_KEYS struct {
	InstanceName    string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	VrfName         string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName          string `protobuf:"bytes,3,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	NeighborAddress string `protobuf:"bytes,4,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
}

func (m *BgpNbrAdvCntBag_KEYS) Reset()                    { *m = BgpNbrAdvCntBag_KEYS{} }
func (m *BgpNbrAdvCntBag_KEYS) String() string            { return proto.CompactTextString(m) }
func (*BgpNbrAdvCntBag_KEYS) ProtoMessage()               {}
func (*BgpNbrAdvCntBag_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BgpNbrAdvCntBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpNbrAdvCntBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpNbrAdvCntBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpNbrAdvCntBag_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type BgpNbrAdvCntBag struct {
	// Number of prefixes advertsied to neighbor
	MaxPrefixAdvertisedcount uint32 `protobuf:"varint,50,opt,name=max_prefix_advertisedcount,json=maxPrefixAdvertisedcount" json:"max_prefix_advertisedcount,omitempty"`
}

func (m *BgpNbrAdvCntBag) Reset()                    { *m = BgpNbrAdvCntBag{} }
func (m *BgpNbrAdvCntBag) String() string            { return proto.CompactTextString(m) }
func (*BgpNbrAdvCntBag) ProtoMessage()               {}
func (*BgpNbrAdvCntBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BgpNbrAdvCntBag) GetMaxPrefixAdvertisedcount() uint32 {
	if m != nil {
		return m.MaxPrefixAdvertisedcount
	}
	return 0
}

func init() {
	proto.RegisterType((*BgpNbrAdvCntBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.afs.af.advertised_path_counts.advertised_path_count.bgp_nbr_adv_cnt_bag_KEYS")
	proto.RegisterType((*BgpNbrAdvCntBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.afs.af.advertised_path_counts.advertised_path_count.bgp_nbr_adv_cnt_bag")
}

func init() { proto.RegisterFile("bgp_nbr_adv_cnt_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x15, 0x40, 0x2d, 0x58, 0x54, 0xa0, 0x30, 0x90, 0x32, 0x55, 0x65, 0x29, 0x8b, 0x07,
	0x60, 0x64, 0xe9, 0xc0, 0x84, 0x84, 0x50, 0x3b, 0x31, 0x9d, 0xce, 0xb1, 0x9d, 0x7a, 0x88, 0x6d,
	0xf9, 0x8c, 0x15, 0x26, 0xde, 0x84, 0x67, 0x45, 0x71, 0x69, 0x10, 0x52, 0x16, 0xfb, 0xf4, 0x7d,
	0x3f, 0xe9, 0xfe, 0xb0, 0xb9, 0x68, 0x3c, 0x58, 0x11, 0x00, 0x65, 0x82, 0xda, 0x46, 0x10, 0xd8,
	0x70, 0x1f, 0x5c, 0x74, 0xe5, 0x57, 0x6d, 0xa8, 0x76, 0x60, 0x1c, 0x41, 0x17, 0xc0, 0xf8, 0xf4,
	0x08, 0x7d, 0xd8, 0x79, 0x15, 0xb8, 0x68, 0x3c, 0x37, 0x96, 0x22, 0xda, 0x5a, 0xd1, 0x50, 0x0d,
	0x05, 0xf4, 0x9f, 0x14, 0x9f, 0x3c, 0x05, 0x4d, 0xfd, 0xc3, 0x51, 0x13, 0x47, 0xcd, 0x51, 0x26,
	0x15, 0xa2, 0x21, 0x25, 0xc1, 0x63, 0xdc, 0x41, 0xed, 0x3e, 0x6c, 0xa4, 0x71, 0xbc, 0xfc, 0x2e,
	0x58, 0x35, 0x32, 0x1e, 0xbc, 0x3c, 0xbf, 0x6f, 0xcb, 0x5b, 0x36, 0x1b, 0xba, 0x59, 0x6c, 0x55,
	0x55, 0x2c, 0x8a, 0xd5, 0xd9, 0xe6, 0xfc, 0x00, 0x5f, 0xb1, 0x55, 0xe5, 0x9c, 0x9d, 0xa6, 0xa0,
	0xf7, 0xfe, 0x28, 0xfb, 0x69, 0x0a, 0x3a, 0xab, 0x6b, 0x36, 0xc5, 0x5f, 0x73, 0x9c, 0xcd, 0x04,
	0xf7, 0xe2, 0x8e, 0x5d, 0x5a, 0x65, 0x9a, 0x9d, 0x70, 0x7d, 0x57, 0x19, 0x14, 0x51, 0x75, 0x92,
	0x13, 0x17, 0x07, 0xbe, 0xde, 0xe3, 0xe5, 0x96, 0x5d, 0x8d, 0xcc, 0x57, 0x3e, 0xb1, 0x9b, 0x16,
	0x3b, 0xf0, 0x41, 0x69, 0xd3, 0xc1, 0xdf, 0x6e, 0x79, 0xab, 0xea, 0x7e, 0x51, 0xac, 0x66, 0x9b,
	0xaa, 0xc5, 0xee, 0x2d, 0x07, 0xd6, 0xff, 0xbd, 0x98, 0xe4, 0xeb, 0x3f, 0xfc, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xfe, 0x9e, 0x65, 0x1a, 0x9a, 0x01, 0x00, 0x00,
}
