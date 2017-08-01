// Code generated by protoc-gen-go.
// source: bgp_nbr_brief_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_vrfs_vrf_sessions_session is a generated protocol buffer package.

It is generated from these files:
	bgp_nbr_brief_bag.proto

It has these top-level messages:
	BgpNbrBriefBag_KEYS
	BgpNbrBriefBag
	IPV4TunnelAddressType
	IPV4MDTAddressType
	RTConstraintAddressType
	IPV6AddressType
	BgpIpv4SrpolicyAddrT
	BgpIpv6SrpolicyAddrT
	BgpL2VpnAddrT
	L2VPNEVPNAddressType
	BgpL2VpnMspwAddrT
	IPV6MVPNAddressType
	IPV4MVPNAddressType
	LS_LSAddressType
	IPv4FlowspecAddressType
	IPv6FlowspecAddressType
	BgpAddrtype
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_vrfs_vrf_sessions_session

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

// BGP Neighbor brief Information
type BgpNbrBriefBag_KEYS struct {
	InstanceName    string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	VrfName         string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	NeighborAddress string `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
}

func (m *BgpNbrBriefBag_KEYS) Reset()                    { *m = BgpNbrBriefBag_KEYS{} }
func (m *BgpNbrBriefBag_KEYS) String() string            { return proto.CompactTextString(m) }
func (*BgpNbrBriefBag_KEYS) ProtoMessage()               {}
func (*BgpNbrBriefBag_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BgpNbrBriefBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpNbrBriefBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpNbrBriefBag_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type BgpNbrBriefBag struct {
	// Speaker this neighbor is allocated to
	SpeakerId uint32 `protobuf:"varint,50,opt,name=speaker_id,json=speakerId" json:"speaker_id,omitempty"`
	// Description
	Description string `protobuf:"bytes,51,opt,name=description" json:"description,omitempty"`
	// Local AS number
	LocalAs uint32 `protobuf:"varint,52,opt,name=local_as,json=localAs" json:"local_as,omitempty"`
	// Remote AS number
	RemoteAs uint32 `protobuf:"varint,53,opt,name=remote_as,json=remoteAs" json:"remote_as,omitempty"`
	// No. of msgs on receive queue
	MessagesQueuedIn uint32 `protobuf:"varint,54,opt,name=messages_queued_in,json=messagesQueuedIn" json:"messages_queued_in,omitempty"`
	// No. of messages on send queue
	MessagesQueuedOut uint32 `protobuf:"varint,55,opt,name=messages_queued_out,json=messagesQueuedOut" json:"messages_queued_out,omitempty"`
	// State of connection
	ConnectionState string `protobuf:"bytes,56,opt,name=connection_state,json=connectionState" json:"connection_state,omitempty"`
	// Local address for the connection
	ConnectionLocalAddress *BgpAddrtype `protobuf:"bytes,57,opt,name=connection_local_address,json=connectionLocalAddress" json:"connection_local_address,omitempty"`
	// Local address configured for the neighbor connection
	IsLocalAddressConfigured bool `protobuf:"varint,58,opt,name=is_local_address_configured,json=isLocalAddressConfigured" json:"is_local_address_configured,omitempty"`
	// Remote address for the connection
	ConnectionRemoteAddress *BgpAddrtype `protobuf:"bytes,59,opt,name=connection_remote_address,json=connectionRemoteAddress" json:"connection_remote_address,omitempty"`
	// Name of the VRF
	VrfName string `protobuf:"bytes,60,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	// Neighbor supports NSR
	NsrEnabled bool `protobuf:"varint,61,opt,name=nsr_enabled,json=nsrEnabled" json:"nsr_enabled,omitempty"`
	// NSR state
	NsrState string `protobuf:"bytes,62,opt,name=nsr_state,json=nsrState" json:"nsr_state,omitempty"`
	// Nbr has postits pending
	PostitPending bool `protobuf:"varint,63,opt,name=postit_pending,json=postitPending" json:"postit_pending,omitempty"`
}

func (m *BgpNbrBriefBag) Reset()                    { *m = BgpNbrBriefBag{} }
func (m *BgpNbrBriefBag) String() string            { return proto.CompactTextString(m) }
func (*BgpNbrBriefBag) ProtoMessage()               {}
func (*BgpNbrBriefBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BgpNbrBriefBag) GetSpeakerId() uint32 {
	if m != nil {
		return m.SpeakerId
	}
	return 0
}

func (m *BgpNbrBriefBag) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *BgpNbrBriefBag) GetLocalAs() uint32 {
	if m != nil {
		return m.LocalAs
	}
	return 0
}

func (m *BgpNbrBriefBag) GetRemoteAs() uint32 {
	if m != nil {
		return m.RemoteAs
	}
	return 0
}

func (m *BgpNbrBriefBag) GetMessagesQueuedIn() uint32 {
	if m != nil {
		return m.MessagesQueuedIn
	}
	return 0
}

func (m *BgpNbrBriefBag) GetMessagesQueuedOut() uint32 {
	if m != nil {
		return m.MessagesQueuedOut
	}
	return 0
}

func (m *BgpNbrBriefBag) GetConnectionState() string {
	if m != nil {
		return m.ConnectionState
	}
	return ""
}

func (m *BgpNbrBriefBag) GetConnectionLocalAddress() *BgpAddrtype {
	if m != nil {
		return m.ConnectionLocalAddress
	}
	return nil
}

func (m *BgpNbrBriefBag) GetIsLocalAddressConfigured() bool {
	if m != nil {
		return m.IsLocalAddressConfigured
	}
	return false
}

func (m *BgpNbrBriefBag) GetConnectionRemoteAddress() *BgpAddrtype {
	if m != nil {
		return m.ConnectionRemoteAddress
	}
	return nil
}

func (m *BgpNbrBriefBag) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpNbrBriefBag) GetNsrEnabled() bool {
	if m != nil {
		return m.NsrEnabled
	}
	return false
}

func (m *BgpNbrBriefBag) GetNsrState() string {
	if m != nil {
		return m.NsrState
	}
	return ""
}

func (m *BgpNbrBriefBag) GetPostitPending() bool {
	if m != nil {
		return m.PostitPending
	}
	return false
}

// IPV4Tunnel Address type
type IPV4TunnelAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPV4TunnelAddressType) Reset()                    { *m = IPV4TunnelAddressType{} }
func (m *IPV4TunnelAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPV4TunnelAddressType) ProtoMessage()               {}
func (*IPV4TunnelAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IPV4TunnelAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPV4MDT Address type
type IPV4MDTAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPV4MDTAddressType) Reset()                    { *m = IPV4MDTAddressType{} }
func (m *IPV4MDTAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPV4MDTAddressType) ProtoMessage()               {}
func (*IPV4MDTAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IPV4MDTAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPV4 RTConstraint Address type
type RTConstraintAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *RTConstraintAddressType) Reset()                    { *m = RTConstraintAddressType{} }
func (m *RTConstraintAddressType) String() string            { return proto.CompactTextString(m) }
func (*RTConstraintAddressType) ProtoMessage()               {}
func (*RTConstraintAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RTConstraintAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPV6 Address type
type IPV6AddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPV6AddressType) Reset()                    { *m = IPV6AddressType{} }
func (m *IPV6AddressType) String() string            { return proto.CompactTextString(m) }
func (*IPV6AddressType) ProtoMessage()               {}
func (*IPV6AddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IPV6AddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type BgpIpv4SrpolicyAddrT struct {
	Ipv4SrpolicyAddress []byte `protobuf:"bytes,1,opt,name=ipv4_srpolicy_address,json=ipv4SrpolicyAddress,proto3" json:"ipv4_srpolicy_address,omitempty"`
}

func (m *BgpIpv4SrpolicyAddrT) Reset()                    { *m = BgpIpv4SrpolicyAddrT{} }
func (m *BgpIpv4SrpolicyAddrT) String() string            { return proto.CompactTextString(m) }
func (*BgpIpv4SrpolicyAddrT) ProtoMessage()               {}
func (*BgpIpv4SrpolicyAddrT) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *BgpIpv4SrpolicyAddrT) GetIpv4SrpolicyAddress() []byte {
	if m != nil {
		return m.Ipv4SrpolicyAddress
	}
	return nil
}

type BgpIpv6SrpolicyAddrT struct {
	Ipv6SrpolicyAddress []byte `protobuf:"bytes,1,opt,name=ipv6_srpolicy_address,json=ipv6SrpolicyAddress,proto3" json:"ipv6_srpolicy_address,omitempty"`
}

func (m *BgpIpv6SrpolicyAddrT) Reset()                    { *m = BgpIpv6SrpolicyAddrT{} }
func (m *BgpIpv6SrpolicyAddrT) String() string            { return proto.CompactTextString(m) }
func (*BgpIpv6SrpolicyAddrT) ProtoMessage()               {}
func (*BgpIpv6SrpolicyAddrT) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *BgpIpv6SrpolicyAddrT) GetIpv6SrpolicyAddress() []byte {
	if m != nil {
		return m.Ipv6SrpolicyAddress
	}
	return nil
}

type BgpL2VpnAddrT struct {
	L2VpnAddress []byte `protobuf:"bytes,1,opt,name=l2vpn_address,json=l2vpnAddress,proto3" json:"l2vpn_address,omitempty"`
}

func (m *BgpL2VpnAddrT) Reset()                    { *m = BgpL2VpnAddrT{} }
func (m *BgpL2VpnAddrT) String() string            { return proto.CompactTextString(m) }
func (*BgpL2VpnAddrT) ProtoMessage()               {}
func (*BgpL2VpnAddrT) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *BgpL2VpnAddrT) GetL2VpnAddress() []byte {
	if m != nil {
		return m.L2VpnAddress
	}
	return nil
}

// L2VPN EVPN Address type
type L2VPNEVPNAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *L2VPNEVPNAddressType) Reset()                    { *m = L2VPNEVPNAddressType{} }
func (m *L2VPNEVPNAddressType) String() string            { return proto.CompactTextString(m) }
func (*L2VPNEVPNAddressType) ProtoMessage()               {}
func (*L2VPNEVPNAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *L2VPNEVPNAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type BgpL2VpnMspwAddrT struct {
	L2VpnAddress []byte `protobuf:"bytes,1,opt,name=l2vpn_address,json=l2vpnAddress,proto3" json:"l2vpn_address,omitempty"`
}

func (m *BgpL2VpnMspwAddrT) Reset()                    { *m = BgpL2VpnMspwAddrT{} }
func (m *BgpL2VpnMspwAddrT) String() string            { return proto.CompactTextString(m) }
func (*BgpL2VpnMspwAddrT) ProtoMessage()               {}
func (*BgpL2VpnMspwAddrT) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *BgpL2VpnMspwAddrT) GetL2VpnAddress() []byte {
	if m != nil {
		return m.L2VpnAddress
	}
	return nil
}

// IPV6 MVPN Address type
type IPV6MVPNAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPV6MVPNAddressType) Reset()                    { *m = IPV6MVPNAddressType{} }
func (m *IPV6MVPNAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPV6MVPNAddressType) ProtoMessage()               {}
func (*IPV6MVPNAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *IPV6MVPNAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPV4 MVPN Address type
type IPV4MVPNAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPV4MVPNAddressType) Reset()                    { *m = IPV4MVPNAddressType{} }
func (m *IPV4MVPNAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPV4MVPNAddressType) ProtoMessage()               {}
func (*IPV4MVPNAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *IPV4MVPNAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// LINKSTATE LINKSTATE Address type
type LS_LSAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *LS_LSAddressType) Reset()                    { *m = LS_LSAddressType{} }
func (m *LS_LSAddressType) String() string            { return proto.CompactTextString(m) }
func (*LS_LSAddressType) ProtoMessage()               {}
func (*LS_LSAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *LS_LSAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPv4 Flowspec Address type
type IPv4FlowspecAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPv4FlowspecAddressType) Reset()                    { *m = IPv4FlowspecAddressType{} }
func (m *IPv4FlowspecAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPv4FlowspecAddressType) ProtoMessage()               {}
func (*IPv4FlowspecAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *IPv4FlowspecAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPv6 Flowspec Address type
type IPv6FlowspecAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPv6FlowspecAddressType) Reset()                    { *m = IPv6FlowspecAddressType{} }
func (m *IPv6FlowspecAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPv6FlowspecAddressType) ProtoMessage()               {}
func (*IPv6FlowspecAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *IPv6FlowspecAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type BgpAddrtype struct {
	Afi string `protobuf:"bytes,1,opt,name=afi" json:"afi,omitempty"`
	// IPv4 Addr
	Ipv4Address string `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address" json:"ipv4_address,omitempty"`
	// IPv4 Mcast Addr
	Ipv4McastAddress string `protobuf:"bytes,3,opt,name=ipv4_mcast_address,json=ipv4McastAddress" json:"ipv4_mcast_address,omitempty"`
	// IPv4 Label Addr
	Ipv4LabelAddress string `protobuf:"bytes,4,opt,name=ipv4_label_address,json=ipv4LabelAddress" json:"ipv4_label_address,omitempty"`
	// IPv4 Tunnel
	Ipv4TunnelAddress *IPV4TunnelAddressType `protobuf:"bytes,5,opt,name=ipv4_tunnel_address,json=ipv4TunnelAddress" json:"ipv4_tunnel_address,omitempty"`
	// IPv4 MDT Addr
	Ipv4MdtAddress *IPV4MDTAddressType `protobuf:"bytes,6,opt,name=ipv4_mdt_address,json=ipv4MdtAddress" json:"ipv4_mdt_address,omitempty"`
	// IPv4 VPN Addr
	Ipv4VpnAddress string `protobuf:"bytes,7,opt,name=ipv4_vpn_address,json=ipv4VpnAddress" json:"ipv4_vpn_address,omitempty"`
	// IPv4 VPN Mcast Addr
	Ipv4VpnaMcastddress string `protobuf:"bytes,8,opt,name=ipv4_vpna_mcastddress,json=ipv4VpnaMcastddress" json:"ipv4_vpna_mcastddress,omitempty"`
	// IPV6 Addr
	Ipv6Address *IPV6AddressType `protobuf:"bytes,9,opt,name=ipv6_address,json=ipv6Address" json:"ipv6_address,omitempty"`
	// IPV6 Mcast Addr
	Ipv6McastAddress *IPV6AddressType `protobuf:"bytes,10,opt,name=ipv6_mcast_address,json=ipv6McastAddress" json:"ipv6_mcast_address,omitempty"`
	// IPv6 Label Addr
	Ipv6LabelAddress *IPV6AddressType `protobuf:"bytes,11,opt,name=ipv6_label_address,json=ipv6LabelAddress" json:"ipv6_label_address,omitempty"`
	// IPv6 VPN Addr
	Ipv6VpnAddress *IPV6AddressType `protobuf:"bytes,12,opt,name=ipv6_vpn_address,json=ipv6VpnAddress" json:"ipv6_vpn_address,omitempty"`
	// IPv6 VPN Mcast Addr
	Ipv6VpnMcastAddress *IPV6AddressType `protobuf:"bytes,13,opt,name=ipv6_vpn_mcast_address,json=ipv6VpnMcastAddress" json:"ipv6_vpn_mcast_address,omitempty"`
	// L2VPN VPLS Addr
	L2VpnvplsAddress *BgpL2VpnAddrT `protobuf:"bytes,14,opt,name=l2_vpnvpls_address,json=l2VpnvplsAddress" json:"l2_vpnvpls_address,omitempty"`
	// RT Constrt Addr
	RtConstraintAddress *RTConstraintAddressType `protobuf:"bytes,15,opt,name=rt_constraint_address,json=rtConstraintAddress" json:"rt_constraint_address,omitempty"`
	// MVPN addr
	Ipv6MvpnAddress *IPV6MVPNAddressType `protobuf:"bytes,16,opt,name=ipv6_mvpn_address,json=ipv6MvpnAddress" json:"ipv6_mvpn_address,omitempty"`
	// MVPN4 addr
	Ipv4MvpnAddress *IPV4MVPNAddressType `protobuf:"bytes,17,opt,name=ipv4_mvpn_address,json=ipv4MvpnAddress" json:"ipv4_mvpn_address,omitempty"`
	// L2VPN EVPN Addr
	L2VpnEvpnAddress *L2VPNEVPNAddressType `protobuf:"bytes,18,opt,name=l2_vpn_evpn_address,json=l2VpnEvpnAddress" json:"l2_vpn_evpn_address,omitempty"`
	// LINKSTATE LINKSTATE Addr
	LsLsAddress *LS_LSAddressType `protobuf:"bytes,19,opt,name=ls_ls_address,json=lsLsAddress" json:"ls_ls_address,omitempty"`
	// L2VPN MSPW Addr
	L2VpnMspwAddress *BgpL2VpnMspwAddrT `protobuf:"bytes,20,opt,name=l2_vpn_mspw_address,json=l2VpnMspwAddress" json:"l2_vpn_mspw_address,omitempty"`
	// IPV4 Flowspec Addr
	Ipv4FlowspecAddress *IPv4FlowspecAddressType `protobuf:"bytes,21,opt,name=ipv4_flowspec_address,json=ipv4FlowspecAddress" json:"ipv4_flowspec_address,omitempty"`
	// IPV6 Flowspec Addr
	Ipv6FlowspecAddress *IPv6FlowspecAddressType `protobuf:"bytes,22,opt,name=ipv6_flowspec_address,json=ipv6FlowspecAddress" json:"ipv6_flowspec_address,omitempty"`
	// IPV4 VPN Flowspec Addr
	Ipv4VpnFlowspecAddress *IPv4FlowspecAddressType `protobuf:"bytes,23,opt,name=ipv4_vpn_flowspec_address,json=ipv4VpnFlowspecAddress" json:"ipv4_vpn_flowspec_address,omitempty"`
	// IPV6 VPN Flowspec Addr
	Ipv6VpnFlowspecAddress *IPv6FlowspecAddressType `protobuf:"bytes,24,opt,name=ipv6_vpn_flowspec_address,json=ipv6VpnFlowspecAddress" json:"ipv6_vpn_flowspec_address,omitempty"`
	// IPV4 Policy Addr
	Ipv4SrPolicyAddress *BgpIpv4SrpolicyAddrT `protobuf:"bytes,25,opt,name=ipv4_sr_policy_address,json=ipv4SrPolicyAddress" json:"ipv4_sr_policy_address,omitempty"`
	// IPV6 Policy Addr
	Ipv6SrPolicyAddress *BgpIpv6SrpolicyAddrT `protobuf:"bytes,26,opt,name=ipv6_sr_policy_address,json=ipv6SrPolicyAddress" json:"ipv6_sr_policy_address,omitempty"`
}

func (m *BgpAddrtype) Reset()                    { *m = BgpAddrtype{} }
func (m *BgpAddrtype) String() string            { return proto.CompactTextString(m) }
func (*BgpAddrtype) ProtoMessage()               {}
func (*BgpAddrtype) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *BgpAddrtype) GetAfi() string {
	if m != nil {
		return m.Afi
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4McastAddress() string {
	if m != nil {
		return m.Ipv4McastAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4LabelAddress() string {
	if m != nil {
		return m.Ipv4LabelAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4TunnelAddress() *IPV4TunnelAddressType {
	if m != nil {
		return m.Ipv4TunnelAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4MdtAddress() *IPV4MDTAddressType {
	if m != nil {
		return m.Ipv4MdtAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4VpnAddress() string {
	if m != nil {
		return m.Ipv4VpnAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4VpnaMcastddress() string {
	if m != nil {
		return m.Ipv4VpnaMcastddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv6Address() *IPV6AddressType {
	if m != nil {
		return m.Ipv6Address
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6McastAddress() *IPV6AddressType {
	if m != nil {
		return m.Ipv6McastAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6LabelAddress() *IPV6AddressType {
	if m != nil {
		return m.Ipv6LabelAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6VpnAddress() *IPV6AddressType {
	if m != nil {
		return m.Ipv6VpnAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6VpnMcastAddress() *IPV6AddressType {
	if m != nil {
		return m.Ipv6VpnMcastAddress
	}
	return nil
}

func (m *BgpAddrtype) GetL2VpnvplsAddress() *BgpL2VpnAddrT {
	if m != nil {
		return m.L2VpnvplsAddress
	}
	return nil
}

func (m *BgpAddrtype) GetRtConstraintAddress() *RTConstraintAddressType {
	if m != nil {
		return m.RtConstraintAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6MvpnAddress() *IPV6MVPNAddressType {
	if m != nil {
		return m.Ipv6MvpnAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4MvpnAddress() *IPV4MVPNAddressType {
	if m != nil {
		return m.Ipv4MvpnAddress
	}
	return nil
}

func (m *BgpAddrtype) GetL2VpnEvpnAddress() *L2VPNEVPNAddressType {
	if m != nil {
		return m.L2VpnEvpnAddress
	}
	return nil
}

func (m *BgpAddrtype) GetLsLsAddress() *LS_LSAddressType {
	if m != nil {
		return m.LsLsAddress
	}
	return nil
}

func (m *BgpAddrtype) GetL2VpnMspwAddress() *BgpL2VpnMspwAddrT {
	if m != nil {
		return m.L2VpnMspwAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4FlowspecAddress() *IPv4FlowspecAddressType {
	if m != nil {
		return m.Ipv4FlowspecAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6FlowspecAddress() *IPv6FlowspecAddressType {
	if m != nil {
		return m.Ipv6FlowspecAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4VpnFlowspecAddress() *IPv4FlowspecAddressType {
	if m != nil {
		return m.Ipv4VpnFlowspecAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6VpnFlowspecAddress() *IPv6FlowspecAddressType {
	if m != nil {
		return m.Ipv6VpnFlowspecAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4SrPolicyAddress() *BgpIpv4SrpolicyAddrT {
	if m != nil {
		return m.Ipv4SrPolicyAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6SrPolicyAddress() *BgpIpv6SrpolicyAddrT {
	if m != nil {
		return m.Ipv6SrPolicyAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpNbrBriefBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.sessions.session.bgp_nbr_brief_bag_KEYS")
	proto.RegisterType((*BgpNbrBriefBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.sessions.session.bgp_nbr_brief_bag")
	proto.RegisterType((*IPV4TunnelAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.sessions.session.IPV4TunnelAddressType")
	proto.RegisterType((*IPV4MDTAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.sessions.session.IPV4MDTAddressType")
	proto.RegisterType((*RTConstraintAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.sessions.session.RTConstraintAddressType")
	proto.RegisterType((*IPV6AddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.sessions.session.IPV6AddressType")
	proto.RegisterType((*BgpIpv4SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.sessions.session.bgp_ipv4_srpolicy_addr_t")
	proto.RegisterType((*BgpIpv6SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.sessions.session.bgp_ipv6_srpolicy_addr_t")
	proto.RegisterType((*BgpL2VpnAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.sessions.session.bgp_l2vpn_addr_t")
	proto.RegisterType((*L2VPNEVPNAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.sessions.session.L2VPNEVPNAddressType")
	proto.RegisterType((*BgpL2VpnMspwAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.sessions.session.bgp_l2vpn_mspw_addr_t")
	proto.RegisterType((*IPV6MVPNAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.sessions.session.IPV6MVPNAddressType")
	proto.RegisterType((*IPV4MVPNAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.sessions.session.IPV4MVPNAddressType")
	proto.RegisterType((*LS_LSAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.sessions.session.LS_LSAddressType")
	proto.RegisterType((*IPv4FlowspecAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.sessions.session.IPv4FlowspecAddressType")
	proto.RegisterType((*IPv6FlowspecAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.sessions.session.IPv6FlowspecAddressType")
	proto.RegisterType((*BgpAddrtype)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.sessions.session.bgp_addrtype")
}

func init() { proto.RegisterFile("bgp_nbr_brief_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x98, 0x5d, 0x6f, 0xe3, 0x44,
	0x17, 0xc7, 0xe5, 0x67, 0x9f, 0xed, 0xcb, 0x49, 0xd2, 0xa6, 0x93, 0x36, 0x75, 0xa9, 0x10, 0xa5,
	0x2b, 0x44, 0x80, 0x12, 0xa4, 0x6c, 0xf0, 0xf2, 0xb2, 0x0b, 0xaa, 0x96, 0x22, 0x55, 0xa4, 0x25,
	0xb8, 0x55, 0x24, 0xc4, 0xc5, 0xc8, 0xb1, 0x27, 0xc1, 0xc2, 0x19, 0x7b, 0x67, 0x9c, 0x94, 0x5e,
	0xf3, 0x01, 0x10, 0x37, 0x48, 0xbc, 0x5c, 0xc0, 0x22, 0x24, 0xc4, 0x0d, 0x9f, 0x80, 0xef, 0x86,
	0x66, 0xc6, 0xef, 0x49, 0x69, 0xb9, 0x48, 0x72, 0x53, 0x39, 0xe7, 0x9c, 0x99, 0xf9, 0xcd, 0x99,
	0xff, 0x19, 0x1f, 0x17, 0x76, 0xfb, 0xc3, 0x00, 0xd3, 0x3e, 0xc3, 0x7d, 0xe6, 0x92, 0x01, 0xee,
	0x5b, 0xc3, 0x66, 0xc0, 0xfc, 0xd0, 0x47, 0x5f, 0xd8, 0x2e, 0xb7, 0x7d, 0xec, 0xfa, 0x1c, 0x7f,
	0xcd, 0xb0, 0x1b, 0x4c, 0xda, 0x58, 0x84, 0xfa, 0x01, 0x61, 0xcd, 0xfe, 0x30, 0x68, 0xba, 0x94,
	0x87, 0x16, 0xb5, 0x09, 0x4f, 0x9e, 0x92, 0x07, 0x6c, 0xd9, 0xa1, 0x3b, 0x21, 0xcd, 0x09, 0x1b,
	0x70, 0xf1, 0xa7, 0xc9, 0x09, 0xe7, 0xae, 0x4f, 0x79, 0xfc, 0x70, 0xf8, 0x8d, 0x06, 0xf5, 0xa9,
	0x85, 0xf1, 0x27, 0x27, 0x9f, 0x5f, 0xa0, 0x07, 0x50, 0x49, 0xe6, 0xa1, 0xd6, 0x88, 0xe8, 0xda,
	0x81, 0xd6, 0x58, 0x37, 0xcb, 0xb1, 0xf1, 0xdc, 0x1a, 0x11, 0xb4, 0x07, 0x6b, 0x13, 0x36, 0x50,
	0xfe, 0xff, 0x49, 0xff, 0xea, 0x84, 0x0d, 0xa4, 0xeb, 0x35, 0xa8, 0x52, 0xe2, 0x0e, 0xbf, 0xec,
	0xfb, 0x0c, 0x5b, 0x8e, 0xc3, 0x08, 0xe7, 0xfa, 0x3d, 0x19, 0xb2, 0x19, 0xdb, 0x8f, 0x95, 0xf9,
	0xf0, 0xf9, 0x0a, 0x6c, 0x4d, 0x51, 0xa0, 0x17, 0x01, 0x78, 0x40, 0xac, 0xaf, 0x08, 0xc3, 0xae,
	0xa3, 0xb7, 0x0e, 0xb4, 0x46, 0xc5, 0x5c, 0x8f, 0x2c, 0xa7, 0x0e, 0x3a, 0x80, 0x92, 0x43, 0xb8,
	0xcd, 0xdc, 0x20, 0x74, 0x7d, 0xaa, 0x3f, 0x94, 0x53, 0x67, 0x4d, 0x02, 0xce, 0xf3, 0x6d, 0xcb,
	0xc3, 0x16, 0xd7, 0xdb, 0x72, 0xf8, 0xaa, 0xfc, 0x7d, 0xcc, 0xd1, 0x3e, 0xac, 0x33, 0x32, 0xf2,
	0x43, 0x22, 0x7c, 0x6f, 0x4b, 0xdf, 0x9a, 0x32, 0x1c, 0x73, 0x74, 0x04, 0x68, 0x44, 0x38, 0xb7,
	0x86, 0x84, 0xe3, 0x67, 0x63, 0x32, 0x26, 0x0e, 0x76, 0xa9, 0x6e, 0xc8, 0xa8, 0x6a, 0xec, 0xf9,
	0x4c, 0x3a, 0x4e, 0x29, 0x6a, 0x42, 0xad, 0x18, 0xed, 0x8f, 0x43, 0xfd, 0x91, 0x0c, 0xdf, 0xca,
	0x87, 0x7f, 0x3a, 0x0e, 0x45, 0x5e, 0x6c, 0x9f, 0x52, 0x62, 0x0b, 0x46, 0xcc, 0x43, 0x2b, 0x24,
	0xfa, 0x3b, 0x2a, 0x2f, 0xa9, 0xfd, 0x42, 0x98, 0xd1, 0x6f, 0x1a, 0xe8, 0x99, 0xd8, 0x68, 0x33,
	0x51, 0x2e, 0xdf, 0x3d, 0xd0, 0x1a, 0xa5, 0x96, 0xdb, 0x9c, 0xa3, 0x3c, 0xc4, 0x04, 0x72, 0xbd,
	0xf0, 0x3a, 0x20, 0x66, 0x3d, 0x45, 0xe9, 0xc8, 0x34, 0x2a, 0x10, 0xf4, 0x04, 0xf6, 0x5d, 0x9e,
	0x87, 0xc3, 0xb6, 0x4f, 0x07, 0xee, 0x70, 0xcc, 0x88, 0xa3, 0xbf, 0x77, 0xa0, 0x35, 0xd6, 0x4c,
	0xdd, 0xe5, 0xd9, 0x41, 0x4f, 0x13, 0x3f, 0xfa, 0x5d, 0x83, 0xbd, 0xcc, 0x26, 0xe3, 0x63, 0x89,
	0x76, 0xf9, 0xfe, 0xa2, 0x77, 0xb9, 0x9b, 0xb2, 0x98, 0x4a, 0x10, 0xd1, 0x36, 0xb3, 0x52, 0x7f,
	0x9c, 0x97, 0xfa, 0x4b, 0x50, 0xa2, 0x9c, 0x61, 0x42, 0xad, 0xbe, 0x47, 0x1c, 0xfd, 0x89, 0xdc,
	0x31, 0x50, 0xce, 0x4e, 0x94, 0x45, 0xc8, 0x4d, 0x04, 0xa8, 0xc3, 0xfe, 0x40, 0x0e, 0x5e, 0xa3,
	0x9c, 0xa9, 0x53, 0x7e, 0x05, 0x36, 0x02, 0x9f, 0x87, 0x6e, 0x88, 0x03, 0x42, 0x1d, 0x97, 0x0e,
	0xf5, 0x0f, 0xe5, 0x04, 0x15, 0x65, 0xed, 0x2a, 0xe3, 0xe1, 0x9b, 0xb0, 0x73, 0xda, 0xed, 0xb5,
	0x2f, 0xc7, 0x94, 0x92, 0x38, 0x8d, 0x97, 0xd7, 0x01, 0x41, 0xdb, 0x70, 0x7f, 0x62, 0x79, 0xe3,
	0xb8, 0x40, 0xd5, 0x8f, 0xc3, 0xd7, 0x01, 0x89, 0xf0, 0xb3, 0x8f, 0x2e, 0x6f, 0x8f, 0x7d, 0x0b,
	0x76, 0xcd, 0xcb, 0xa7, 0x3e, 0xe5, 0x21, 0xb3, 0x5c, 0x1a, 0xde, 0x3e, 0xe0, 0x55, 0xd8, 0x3c,
	0xed, 0xf6, 0x8c, 0xdb, 0x03, 0xcf, 0x41, 0x17, 0xd9, 0x95, 0x07, 0xc6, 0x59, 0xe0, 0x7b, 0xae,
	0x7d, 0x2d, 0x73, 0x8d, 0x43, 0xd4, 0x82, 0x9d, 0x69, 0xbb, 0x38, 0x73, 0x31, 0x43, 0xd9, 0xac,
	0x09, 0xe7, 0x45, 0xe4, 0x8b, 0x6f, 0x8a, 0x74, 0x3e, 0xe3, 0x86, 0xf9, 0x8c, 0x7f, 0x9b, 0xcf,
	0x28, 0xce, 0xf7, 0x08, 0xaa, 0x62, 0x3e, 0xaf, 0x35, 0x09, 0x68, 0x3c, 0xcf, 0x03, 0xa8, 0xa4,
	0xbf, 0xd3, 0xf1, 0x65, 0x69, 0x8c, 0x07, 0x1e, 0xc1, 0x76, 0xa7, 0xd5, 0xeb, 0x9e, 0x9f, 0xf4,
	0xba, 0xe7, 0xb7, 0xa7, 0xe1, 0x31, 0xec, 0xa4, 0xcb, 0x8c, 0x78, 0x70, 0xf5, 0x9f, 0xd6, 0x7a,
	0x03, 0x6a, 0x22, 0xdb, 0x67, 0x77, 0x5a, 0x4a, 0x05, 0xb7, 0xef, 0x16, 0xdc, 0x80, 0x6a, 0xe7,
	0x02, 0x77, 0x2e, 0xee, 0x24, 0x91, 0xd3, 0xee, 0xa4, 0xfd, 0xb1, 0xe7, 0x5f, 0xf1, 0x80, 0xd8,
	0x77, 0x1d, 0x60, 0xdc, 0x7d, 0xc0, 0xdf, 0xfb, 0x50, 0xce, 0x56, 0x22, 0xaa, 0xc2, 0x3d, 0x6b,
	0xe0, 0x46, 0x41, 0xe2, 0x11, 0xbd, 0x0c, 0x65, 0xa9, 0x98, 0x38, 0x59, 0xea, 0x8d, 0x53, 0x12,
	0xb6, 0xb8, 0x4a, 0x8f, 0x00, 0xc9, 0x90, 0x91, 0x6d, 0xf1, 0xb0, 0xf0, 0xde, 0xa9, 0x0a, 0xcf,
	0x99, 0x70, 0x14, 0xa3, 0x3d, 0xab, 0x4f, 0xd2, 0x9b, 0xf5, 0xff, 0x69, 0x74, 0x47, 0x38, 0xe2,
	0xe8, 0xe7, 0x1a, 0x48, 0x51, 0xe2, 0x50, 0xd6, 0x60, 0x12, 0x7f, 0x5f, 0xde, 0x51, 0x6c, 0xae,
	0x77, 0xd4, 0xcc, 0xd2, 0x37, 0xb7, 0xc4, 0x2a, 0x39, 0x33, 0xfa, 0x41, 0x83, 0xaa, 0xca, 0x80,
	0x93, 0xee, 0x7f, 0x45, 0x12, 0xfa, 0x73, 0x27, 0xcc, 0xdf, 0x36, 0xe6, 0x86, 0x4c, 0xb8, 0x93,
	0xa4, 0xbb, 0x11, 0xa1, 0x65, 0x05, 0xbf, 0x2a, 0x93, 0x2d, 0x23, 0x7b, 0x89, 0xe4, 0x93, 0xbb,
	0x61, 0x12, 0x50, 0x4b, 0x9d, 0x65, 0x14, 0xbe, 0x26, 0xc3, 0x6b, 0x51, 0xb8, 0x75, 0x96, 0xba,
	0xd0, 0xb7, 0x9a, 0x94, 0x87, 0x91, 0x4c, 0xbd, 0x2e, 0x77, 0xed, 0xcd, 0x7b, 0xd7, 0xd9, 0x6b,
	0x50, 0x8a, 0x31, 0x36, 0xa0, 0x1f, 0x35, 0xa9, 0x2f, 0xa3, 0xa0, 0x46, 0x58, 0x02, 0x97, 0x48,
	0xbc, 0x91, 0xd3, 0x7e, 0x02, 0x97, 0x17, 0x7f, 0x69, 0x59, 0x70, 0xb9, 0x52, 0xfb, 0x5e, 0xa9,
	0xd8, 0xc8, 0x49, 0xa5, 0xbc, 0x04, 0x34, 0x21, 0x4c, 0x23, 0x23, 0xcc, 0x5f, 0x35, 0xa8, 0x27,
	0x60, 0xf9, 0x63, 0xad, 0x2c, 0x01, 0xaf, 0x16, 0xe1, 0xe5, 0x4e, 0xf6, 0x27, 0x0d, 0x90, 0xd7,
	0x12, 0x84, 0x93, 0xc0, 0xe3, 0x09, 0xdf, 0x86, 0xe4, 0x1b, 0xcd, 0xbd, 0x95, 0xca, 0xbe, 0x4c,
	0xcd, 0xaa, 0xd7, 0xea, 0x29, 0x8e, 0x98, 0xee, 0x0f, 0x0d, 0x76, 0x58, 0x28, 0x3a, 0xc4, 0xa8,
	0xdd, 0x48, 0x00, 0x37, 0x25, 0x60, 0x38, 0x57, 0xc0, 0x1b, 0xfa, 0x1c, 0xb3, 0xc6, 0xc2, 0x29,
	0x07, 0xfa, 0x59, 0x83, 0x2d, 0x55, 0xbf, 0x59, 0x19, 0x56, 0x25, 0x66, 0x30, 0xf7, 0x73, 0x2e,
	0xbc, 0xc2, 0xcd, 0x4d, 0x59, 0xc2, 0x69, 0x5f, 0x10, 0xe3, 0xb5, 0xf3, 0x78, 0x5b, 0x8b, 0xc1,
	0x6b, 0xcf, 0xc2, 0x6b, 0x67, 0xf1, 0x7e, 0xd1, 0xa0, 0xa6, 0x64, 0x88, 0x49, 0x16, 0x10, 0x49,
	0xc0, 0x67, 0x73, 0x05, 0x9c, 0xd5, 0x9b, 0x45, 0x5a, 0x3c, 0xc9, 0x20, 0x7e, 0xa7, 0x41, 0xc5,
	0xe3, 0x38, 0x53, 0x24, 0xb5, 0x05, 0x14, 0x49, 0xb1, 0xe5, 0x32, 0x4b, 0x1e, 0xef, 0xf0, 0xcc,
	0x0d, 0x13, 0xa7, 0x2d, 0xe9, 0x14, 0x05, 0xd9, 0xf6, 0x02, 0xba, 0x8c, 0x99, 0x4d, 0x6a, 0x94,
	0xb7, 0x33, 0x1e, 0x5c, 0x65, 0x6b, 0x58, 0x2e, 0x3d, 0x88, 0xda, 0xbb, 0x84, 0x72, 0x67, 0x01,
	0x35, 0x7c, 0x43, 0x23, 0xaa, 0xba, 0x82, 0x82, 0x23, 0x46, 0x35, 0xa6, 0x51, 0xeb, 0x8b, 0x41,
	0x35, 0x6e, 0x42, 0x2d, 0x3a, 0xd0, 0x5f, 0x1a, 0xec, 0x25, 0xfd, 0xd1, 0x14, 0xee, 0xee, 0x12,
	0x33, 0x5b, 0x8f, 0xfa, 0xad, 0x1b, 0x88, 0x8d, 0xd9, 0xc4, 0xfa, 0x12, 0x13, 0x5c, 0x8f, 0x5e,
	0x8c, 0x45, 0xe2, 0x3f, 0xd5, 0xfb, 0x5b, 0x7c, 0x75, 0xe2, 0xc2, 0x67, 0xe2, 0x9e, 0xc4, 0x1d,
	0xcf, 0xbd, 0xc0, 0x66, 0x7d, 0x0c, 0xc7, 0x5f, 0xbb, 0xdd, 0xec, 0xd7, 0x69, 0x0c, 0x6b, 0xcc,
	0x80, 0x7d, 0x61, 0x71, 0xb0, 0xc6, 0x4c, 0x58, 0xa3, 0x00, 0xdb, 0x5f, 0x91, 0xff, 0xae, 0x7c,
	0xf8, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x1e, 0x8d, 0xca, 0xc9, 0x14, 0x00, 0x00,
}
