// Code generated by protoc-gen-go.
// source: lldp_neighbor.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ethernet_lldp_oper_lldp_nodes_node_neighbors_devices_device is a generated protocol buffer package.

It is generated from these files:
	lldp_neighbor.proto

It has these top-level messages:
	LldpNeighbor_KEYS
	LldpNeighbor
	LldpNeighborItem
	In6AddrTd
	LldpL3Addr
	LldpAddrEntry
	LldpAddrEntryItem
	LldpUnknownTlvEntry
	LldpUnknownTlvEntryItem
	LldpOrgDefTlvEntry
	LldpOrgDefTlvEntryItem
	LldpNeighborDetail
	LldpNeighborMib
*/
package cisco_ios_xr_ethernet_lldp_oper_lldp_nodes_node_neighbors_devices_device

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

// LLDP neighbor info
type LldpNeighbor_KEYS struct {
	NodeName      string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	DeviceId      string `protobuf:"bytes,2,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	InterfaceName string `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *LldpNeighbor_KEYS) Reset()                    { *m = LldpNeighbor_KEYS{} }
func (m *LldpNeighbor_KEYS) String() string            { return proto.CompactTextString(m) }
func (*LldpNeighbor_KEYS) ProtoMessage()               {}
func (*LldpNeighbor_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LldpNeighbor_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *LldpNeighbor_KEYS) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *LldpNeighbor_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type LldpNeighbor struct {
	// Next neighbor in the list
	LldpNeighbor []*LldpNeighborItem `protobuf:"bytes,50,rep,name=lldp_neighbor,json=lldpNeighbor" json:"lldp_neighbor,omitempty"`
}

func (m *LldpNeighbor) Reset()                    { *m = LldpNeighbor{} }
func (m *LldpNeighbor) String() string            { return proto.CompactTextString(m) }
func (*LldpNeighbor) ProtoMessage()               {}
func (*LldpNeighbor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LldpNeighbor) GetLldpNeighbor() []*LldpNeighborItem {
	if m != nil {
		return m.LldpNeighbor
	}
	return nil
}

type LldpNeighborItem struct {
	// Interface the neighbor entry was received on
	ReceivingInterfaceName string `protobuf:"bytes,1,opt,name=receiving_interface_name,json=receivingInterfaceName" json:"receiving_interface_name,omitempty"`
	// Parent Interface the neighbor entry was received on
	ReceivingParentInterfaceName string `protobuf:"bytes,2,opt,name=receiving_parent_interface_name,json=receivingParentInterfaceName" json:"receiving_parent_interface_name,omitempty"`
	// Device identifier
	DeviceId string `protobuf:"bytes,3,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	// Chassis id
	ChassisId string `protobuf:"bytes,4,opt,name=chassis_id,json=chassisId" json:"chassis_id,omitempty"`
	// Outgoing port identifier
	PortIdDetail string `protobuf:"bytes,5,opt,name=port_id_detail,json=portIdDetail" json:"port_id_detail,omitempty"`
	// Version number
	HeaderVersion uint32 `protobuf:"varint,6,opt,name=header_version,json=headerVersion" json:"header_version,omitempty"`
	// Remaining hold time
	HoldTime uint32 `protobuf:"varint,7,opt,name=hold_time,json=holdTime" json:"hold_time,omitempty"`
	// Enabled Capabilities
	EnabledCapabilities string `protobuf:"bytes,8,opt,name=enabled_capabilities,json=enabledCapabilities" json:"enabled_capabilities,omitempty"`
	// Platform type
	Platform string `protobuf:"bytes,9,opt,name=platform" json:"platform,omitempty"`
	// Detailed neighbor info
	Detail *LldpNeighborDetail `protobuf:"bytes,10,opt,name=detail" json:"detail,omitempty"`
	// MIB nieghbor info
	Mib *LldpNeighborMib `protobuf:"bytes,11,opt,name=mib" json:"mib,omitempty"`
}

func (m *LldpNeighborItem) Reset()                    { *m = LldpNeighborItem{} }
func (m *LldpNeighborItem) String() string            { return proto.CompactTextString(m) }
func (*LldpNeighborItem) ProtoMessage()               {}
func (*LldpNeighborItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LldpNeighborItem) GetReceivingInterfaceName() string {
	if m != nil {
		return m.ReceivingInterfaceName
	}
	return ""
}

func (m *LldpNeighborItem) GetReceivingParentInterfaceName() string {
	if m != nil {
		return m.ReceivingParentInterfaceName
	}
	return ""
}

func (m *LldpNeighborItem) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *LldpNeighborItem) GetChassisId() string {
	if m != nil {
		return m.ChassisId
	}
	return ""
}

func (m *LldpNeighborItem) GetPortIdDetail() string {
	if m != nil {
		return m.PortIdDetail
	}
	return ""
}

func (m *LldpNeighborItem) GetHeaderVersion() uint32 {
	if m != nil {
		return m.HeaderVersion
	}
	return 0
}

func (m *LldpNeighborItem) GetHoldTime() uint32 {
	if m != nil {
		return m.HoldTime
	}
	return 0
}

func (m *LldpNeighborItem) GetEnabledCapabilities() string {
	if m != nil {
		return m.EnabledCapabilities
	}
	return ""
}

func (m *LldpNeighborItem) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *LldpNeighborItem) GetDetail() *LldpNeighborDetail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (m *LldpNeighborItem) GetMib() *LldpNeighborMib {
	if m != nil {
		return m.Mib
	}
	return nil
}

type In6AddrTd struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *In6AddrTd) Reset()                    { *m = In6AddrTd{} }
func (m *In6AddrTd) String() string            { return proto.CompactTextString(m) }
func (*In6AddrTd) ProtoMessage()               {}
func (*In6AddrTd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *In6AddrTd) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type LldpL3Addr struct {
	AddressType string `protobuf:"bytes,1,opt,name=address_type,json=addressType" json:"address_type,omitempty"`
	// IPv4 address
	Ipv4Address string `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address" json:"ipv4_address,omitempty"`
	// IPv6 address
	Ipv6Address *In6AddrTd `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address" json:"ipv6_address,omitempty"`
}

func (m *LldpL3Addr) Reset()                    { *m = LldpL3Addr{} }
func (m *LldpL3Addr) String() string            { return proto.CompactTextString(m) }
func (*LldpL3Addr) ProtoMessage()               {}
func (*LldpL3Addr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LldpL3Addr) GetAddressType() string {
	if m != nil {
		return m.AddressType
	}
	return ""
}

func (m *LldpL3Addr) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *LldpL3Addr) GetIpv6Address() *In6AddrTd {
	if m != nil {
		return m.Ipv6Address
	}
	return nil
}

type LldpAddrEntry struct {
	// Next address entry in list
	LldpAddrEntry []*LldpAddrEntryItem `protobuf:"bytes,1,rep,name=lldp_addr_entry,json=lldpAddrEntry" json:"lldp_addr_entry,omitempty"`
}

func (m *LldpAddrEntry) Reset()                    { *m = LldpAddrEntry{} }
func (m *LldpAddrEntry) String() string            { return proto.CompactTextString(m) }
func (*LldpAddrEntry) ProtoMessage()               {}
func (*LldpAddrEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LldpAddrEntry) GetLldpAddrEntry() []*LldpAddrEntryItem {
	if m != nil {
		return m.LldpAddrEntry
	}
	return nil
}

type LldpAddrEntryItem struct {
	// Network layer address
	Address *LldpL3Addr `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// MA sub type
	MaSubtype uint32 `protobuf:"varint,2,opt,name=ma_subtype,json=maSubtype" json:"ma_subtype,omitempty"`
	// Interface num
	IfNum uint32 `protobuf:"varint,3,opt,name=if_num,json=ifNum" json:"if_num,omitempty"`
}

func (m *LldpAddrEntryItem) Reset()                    { *m = LldpAddrEntryItem{} }
func (m *LldpAddrEntryItem) String() string            { return proto.CompactTextString(m) }
func (*LldpAddrEntryItem) ProtoMessage()               {}
func (*LldpAddrEntryItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LldpAddrEntryItem) GetAddress() *LldpL3Addr {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *LldpAddrEntryItem) GetMaSubtype() uint32 {
	if m != nil {
		return m.MaSubtype
	}
	return 0
}

func (m *LldpAddrEntryItem) GetIfNum() uint32 {
	if m != nil {
		return m.IfNum
	}
	return 0
}

type LldpUnknownTlvEntry struct {
	// Next unknown TLV entry in list
	LldpUnknownTlvEntry []*LldpUnknownTlvEntryItem `protobuf:"bytes,1,rep,name=lldp_unknown_tlv_entry,json=lldpUnknownTlvEntry" json:"lldp_unknown_tlv_entry,omitempty"`
}

func (m *LldpUnknownTlvEntry) Reset()                    { *m = LldpUnknownTlvEntry{} }
func (m *LldpUnknownTlvEntry) String() string            { return proto.CompactTextString(m) }
func (*LldpUnknownTlvEntry) ProtoMessage()               {}
func (*LldpUnknownTlvEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LldpUnknownTlvEntry) GetLldpUnknownTlvEntry() []*LldpUnknownTlvEntryItem {
	if m != nil {
		return m.LldpUnknownTlvEntry
	}
	return nil
}

type LldpUnknownTlvEntryItem struct {
	// Unknown TLV type
	TlvType uint32 `protobuf:"varint,1,opt,name=tlv_type,json=tlvType" json:"tlv_type,omitempty"`
	// Unknown TLV payload
	TlvValue []byte `protobuf:"bytes,2,opt,name=tlv_value,json=tlvValue,proto3" json:"tlv_value,omitempty"`
}

func (m *LldpUnknownTlvEntryItem) Reset()                    { *m = LldpUnknownTlvEntryItem{} }
func (m *LldpUnknownTlvEntryItem) String() string            { return proto.CompactTextString(m) }
func (*LldpUnknownTlvEntryItem) ProtoMessage()               {}
func (*LldpUnknownTlvEntryItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *LldpUnknownTlvEntryItem) GetTlvType() uint32 {
	if m != nil {
		return m.TlvType
	}
	return 0
}

func (m *LldpUnknownTlvEntryItem) GetTlvValue() []byte {
	if m != nil {
		return m.TlvValue
	}
	return nil
}

type LldpOrgDefTlvEntry struct {
	// Next Org Def TLV entry in list
	LldpOrgDefTlvEntry []*LldpOrgDefTlvEntryItem `protobuf:"bytes,1,rep,name=lldp_org_def_tlv_entry,json=lldpOrgDefTlvEntry" json:"lldp_org_def_tlv_entry,omitempty"`
}

func (m *LldpOrgDefTlvEntry) Reset()                    { *m = LldpOrgDefTlvEntry{} }
func (m *LldpOrgDefTlvEntry) String() string            { return proto.CompactTextString(m) }
func (*LldpOrgDefTlvEntry) ProtoMessage()               {}
func (*LldpOrgDefTlvEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *LldpOrgDefTlvEntry) GetLldpOrgDefTlvEntry() []*LldpOrgDefTlvEntryItem {
	if m != nil {
		return m.LldpOrgDefTlvEntry
	}
	return nil
}

type LldpOrgDefTlvEntryItem struct {
	// Organizationally Unique Identifier
	Oui uint32 `protobuf:"varint,1,opt,name=oui" json:"oui,omitempty"`
	// Org Def TLV subtype
	TlvSubtype uint32 `protobuf:"varint,2,opt,name=tlv_subtype,json=tlvSubtype" json:"tlv_subtype,omitempty"`
	// lldpRemOrgDefInfoIndex
	TlvInfoIndes uint32 `protobuf:"varint,3,opt,name=tlv_info_indes,json=tlvInfoIndes" json:"tlv_info_indes,omitempty"`
	// Org Def TLV payload
	TlvValue []byte `protobuf:"bytes,4,opt,name=tlv_value,json=tlvValue,proto3" json:"tlv_value,omitempty"`
}

func (m *LldpOrgDefTlvEntryItem) Reset()                    { *m = LldpOrgDefTlvEntryItem{} }
func (m *LldpOrgDefTlvEntryItem) String() string            { return proto.CompactTextString(m) }
func (*LldpOrgDefTlvEntryItem) ProtoMessage()               {}
func (*LldpOrgDefTlvEntryItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *LldpOrgDefTlvEntryItem) GetOui() uint32 {
	if m != nil {
		return m.Oui
	}
	return 0
}

func (m *LldpOrgDefTlvEntryItem) GetTlvSubtype() uint32 {
	if m != nil {
		return m.TlvSubtype
	}
	return 0
}

func (m *LldpOrgDefTlvEntryItem) GetTlvInfoIndes() uint32 {
	if m != nil {
		return m.TlvInfoIndes
	}
	return 0
}

func (m *LldpOrgDefTlvEntryItem) GetTlvValue() []byte {
	if m != nil {
		return m.TlvValue
	}
	return nil
}

type LldpNeighborDetail struct {
	// Port Description
	PortDescription string `protobuf:"bytes,1,opt,name=port_description,json=portDescription" json:"port_description,omitempty"`
	// System Name
	SystemName string `protobuf:"bytes,2,opt,name=system_name,json=systemName" json:"system_name,omitempty"`
	// System Description
	SystemDescription string `protobuf:"bytes,3,opt,name=system_description,json=systemDescription" json:"system_description,omitempty"`
	// Time remaining
	TimeRemaining uint32 `protobuf:"varint,4,opt,name=time_remaining,json=timeRemaining" json:"time_remaining,omitempty"`
	// System Capabilities
	SystemCapabilities string `protobuf:"bytes,5,opt,name=system_capabilities,json=systemCapabilities" json:"system_capabilities,omitempty"`
	// Enabled Capabilities
	EnabledCapabilities string `protobuf:"bytes,6,opt,name=enabled_capabilities,json=enabledCapabilities" json:"enabled_capabilities,omitempty"`
	// Management Addresses
	NetworkAddresses *LldpAddrEntry `protobuf:"bytes,7,opt,name=network_addresses,json=networkAddresses" json:"network_addresses,omitempty"`
	// Auto Negotiation
	AutoNegotiation string `protobuf:"bytes,8,opt,name=auto_negotiation,json=autoNegotiation" json:"auto_negotiation,omitempty"`
	// Physical media capabilities
	PhysicalMediaCapabilities string `protobuf:"bytes,9,opt,name=physical_media_capabilities,json=physicalMediaCapabilities" json:"physical_media_capabilities,omitempty"`
	// Media Attachment Unit type
	MediaAttachmentUnitType uint32 `protobuf:"varint,10,opt,name=media_attachment_unit_type,json=mediaAttachmentUnitType" json:"media_attachment_unit_type,omitempty"`
	// Vlan ID
	PortVlanId uint32 `protobuf:"varint,11,opt,name=port_vlan_id,json=portVlanId" json:"port_vlan_id,omitempty"`
}

func (m *LldpNeighborDetail) Reset()                    { *m = LldpNeighborDetail{} }
func (m *LldpNeighborDetail) String() string            { return proto.CompactTextString(m) }
func (*LldpNeighborDetail) ProtoMessage()               {}
func (*LldpNeighborDetail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *LldpNeighborDetail) GetPortDescription() string {
	if m != nil {
		return m.PortDescription
	}
	return ""
}

func (m *LldpNeighborDetail) GetSystemName() string {
	if m != nil {
		return m.SystemName
	}
	return ""
}

func (m *LldpNeighborDetail) GetSystemDescription() string {
	if m != nil {
		return m.SystemDescription
	}
	return ""
}

func (m *LldpNeighborDetail) GetTimeRemaining() uint32 {
	if m != nil {
		return m.TimeRemaining
	}
	return 0
}

func (m *LldpNeighborDetail) GetSystemCapabilities() string {
	if m != nil {
		return m.SystemCapabilities
	}
	return ""
}

func (m *LldpNeighborDetail) GetEnabledCapabilities() string {
	if m != nil {
		return m.EnabledCapabilities
	}
	return ""
}

func (m *LldpNeighborDetail) GetNetworkAddresses() *LldpAddrEntry {
	if m != nil {
		return m.NetworkAddresses
	}
	return nil
}

func (m *LldpNeighborDetail) GetAutoNegotiation() string {
	if m != nil {
		return m.AutoNegotiation
	}
	return ""
}

func (m *LldpNeighborDetail) GetPhysicalMediaCapabilities() string {
	if m != nil {
		return m.PhysicalMediaCapabilities
	}
	return ""
}

func (m *LldpNeighborDetail) GetMediaAttachmentUnitType() uint32 {
	if m != nil {
		return m.MediaAttachmentUnitType
	}
	return 0
}

func (m *LldpNeighborDetail) GetPortVlanId() uint32 {
	if m != nil {
		return m.PortVlanId
	}
	return 0
}

type LldpNeighborMib struct {
	// TimeFilter
	RemTimeMark uint32 `protobuf:"varint,1,opt,name=rem_time_mark,json=remTimeMark" json:"rem_time_mark,omitempty"`
	// LldpPortNumber
	RemLocalPortNum uint32 `protobuf:"varint,2,opt,name=rem_local_port_num,json=remLocalPortNum" json:"rem_local_port_num,omitempty"`
	// lldpRemIndex
	RemIndex uint32 `protobuf:"varint,3,opt,name=rem_index,json=remIndex" json:"rem_index,omitempty"`
	// Chassis ID sub type
	ChassisIdSubType uint32 `protobuf:"varint,4,opt,name=chassis_id_sub_type,json=chassisIdSubType" json:"chassis_id_sub_type,omitempty"`
	// Chassis ID length
	ChassisIdLen uint32 `protobuf:"varint,5,opt,name=chassis_id_len,json=chassisIdLen" json:"chassis_id_len,omitempty"`
	// Port ID sub type
	PortIdSubType uint32 `protobuf:"varint,6,opt,name=port_id_sub_type,json=portIdSubType" json:"port_id_sub_type,omitempty"`
	// Port ID length
	PortIdLen uint32 `protobuf:"varint,7,opt,name=port_id_len,json=portIdLen" json:"port_id_len,omitempty"`
	// Supported and combined cpabilities
	CombinedCapabilities uint32 `protobuf:"varint,8,opt,name=combined_capabilities,json=combinedCapabilities" json:"combined_capabilities,omitempty"`
	// Unknown TLV list
	UnknownTlvList *LldpUnknownTlvEntry `protobuf:"bytes,9,opt,name=unknown_tlv_list,json=unknownTlvList" json:"unknown_tlv_list,omitempty"`
	// Org Def TLV list
	OrgDefTlvList *LldpOrgDefTlvEntry `protobuf:"bytes,10,opt,name=org_def_tlv_list,json=orgDefTlvList" json:"org_def_tlv_list,omitempty"`
}

func (m *LldpNeighborMib) Reset()                    { *m = LldpNeighborMib{} }
func (m *LldpNeighborMib) String() string            { return proto.CompactTextString(m) }
func (*LldpNeighborMib) ProtoMessage()               {}
func (*LldpNeighborMib) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *LldpNeighborMib) GetRemTimeMark() uint32 {
	if m != nil {
		return m.RemTimeMark
	}
	return 0
}

func (m *LldpNeighborMib) GetRemLocalPortNum() uint32 {
	if m != nil {
		return m.RemLocalPortNum
	}
	return 0
}

func (m *LldpNeighborMib) GetRemIndex() uint32 {
	if m != nil {
		return m.RemIndex
	}
	return 0
}

func (m *LldpNeighborMib) GetChassisIdSubType() uint32 {
	if m != nil {
		return m.ChassisIdSubType
	}
	return 0
}

func (m *LldpNeighborMib) GetChassisIdLen() uint32 {
	if m != nil {
		return m.ChassisIdLen
	}
	return 0
}

func (m *LldpNeighborMib) GetPortIdSubType() uint32 {
	if m != nil {
		return m.PortIdSubType
	}
	return 0
}

func (m *LldpNeighborMib) GetPortIdLen() uint32 {
	if m != nil {
		return m.PortIdLen
	}
	return 0
}

func (m *LldpNeighborMib) GetCombinedCapabilities() uint32 {
	if m != nil {
		return m.CombinedCapabilities
	}
	return 0
}

func (m *LldpNeighborMib) GetUnknownTlvList() *LldpUnknownTlvEntry {
	if m != nil {
		return m.UnknownTlvList
	}
	return nil
}

func (m *LldpNeighborMib) GetOrgDefTlvList() *LldpOrgDefTlvEntry {
	if m != nil {
		return m.OrgDefTlvList
	}
	return nil
}

func init() {
	proto.RegisterType((*LldpNeighbor_KEYS)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.devices.device.lldp_neighbor_KEYS")
	proto.RegisterType((*LldpNeighbor)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.devices.device.lldp_neighbor")
	proto.RegisterType((*LldpNeighborItem)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.devices.device.lldp_neighbor_item")
	proto.RegisterType((*In6AddrTd)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.devices.device.in6_addr_td")
	proto.RegisterType((*LldpL3Addr)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.devices.device.lldp_l3_addr")
	proto.RegisterType((*LldpAddrEntry)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.devices.device.lldp_addr_entry")
	proto.RegisterType((*LldpAddrEntryItem)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.devices.device.lldp_addr_entry_item")
	proto.RegisterType((*LldpUnknownTlvEntry)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.devices.device.lldp_unknown_tlv_entry")
	proto.RegisterType((*LldpUnknownTlvEntryItem)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.devices.device.lldp_unknown_tlv_entry_item")
	proto.RegisterType((*LldpOrgDefTlvEntry)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.devices.device.lldp_org_def_tlv_entry")
	proto.RegisterType((*LldpOrgDefTlvEntryItem)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.devices.device.lldp_org_def_tlv_entry_item")
	proto.RegisterType((*LldpNeighborDetail)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.devices.device.lldp_neighbor_detail")
	proto.RegisterType((*LldpNeighborMib)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.devices.device.lldp_neighbor_mib")
}

func init() { proto.RegisterFile("lldp_neighbor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0x5f, 0x6f, 0x1b, 0x45,
	0x10, 0xd7, 0x35, 0x69, 0x1a, 0xcf, 0xe5, 0x92, 0x74, 0x93, 0x16, 0xb7, 0x01, 0x1a, 0x4c, 0x2b,
	0x82, 0x50, 0x8d, 0x48, 0x51, 0x85, 0x84, 0x84, 0x54, 0xd1, 0x4a, 0x58, 0xb4, 0xa1, 0xba, 0x26,
	0x91, 0x2a, 0x10, 0xcb, 0xda, 0xb7, 0xb6, 0x57, 0xb9, 0xdd, 0x3b, 0xed, 0xed, 0x5d, 0x93, 0x47,
	0x5e, 0x2a, 0xc1, 0x6b, 0x9f, 0xe8, 0xe7, 0x80, 0x57, 0x3e, 0x02, 0xdf, 0x04, 0x21, 0xbe, 0x01,
	0x9a, 0xdd, 0x3d, 0xc7, 0xe7, 0x24, 0x3c, 0xc5, 0x2f, 0xb1, 0xfd, 0x9b, 0xdf, 0xce, 0xec, 0xcc,
	0xce, 0xbf, 0xc0, 0x46, 0x9a, 0x26, 0x39, 0x55, 0x5c, 0x8c, 0xc6, 0xfd, 0x4c, 0x77, 0x73, 0x9d,
	0x99, 0x8c, 0x7c, 0x33, 0x10, 0xc5, 0x20, 0xa3, 0x22, 0x2b, 0xe8, 0xb1, 0xa6, 0xdc, 0x8c, 0xb9,
	0x56, 0xdc, 0x50, 0x4b, 0xcd, 0x72, 0xae, 0xbb, 0xf8, 0xad, 0xab, 0xb2, 0x84, 0x17, 0xf6, 0x6f,
	0xb7, 0x3e, 0x5f, 0x74, 0x13, 0x5e, 0x89, 0x01, 0xaf, 0x3f, 0x3b, 0x25, 0x90, 0x86, 0x01, 0xfa,
	0xed, 0x93, 0x97, 0x2f, 0xc8, 0x16, 0xb4, 0xf0, 0x18, 0x55, 0x4c, 0xf2, 0x76, 0xb0, 0x1d, 0xec,
	0xb4, 0xe2, 0x65, 0x04, 0xf6, 0x98, 0xe4, 0x28, 0x74, 0x87, 0xa9, 0x48, 0xda, 0x57, 0x9c, 0xd0,
	0x01, 0xbd, 0x84, 0xdc, 0x83, 0x55, 0xa1, 0x0c, 0xd7, 0x43, 0x36, 0xf0, 0xc7, 0x17, 0x2c, 0x23,
	0x9a, 0xa0, 0xa8, 0xa3, 0xf3, 0x26, 0x80, 0xa8, 0x61, 0x97, 0xfc, 0x3c, 0x8b, 0xb4, 0x77, 0xb7,
	0x17, 0x76, 0xc2, 0xdd, 0x1f, 0xba, 0x97, 0xe5, 0x6b, 0xb7, 0xe9, 0xa8, 0x30, 0x5c, 0xc6, 0x2b,
	0x88, 0xed, 0x79, 0xa8, 0xf3, 0xef, 0xe2, 0x6c, 0x34, 0x90, 0x44, 0xbe, 0x80, 0xb6, 0xe6, 0x03,
	0x2e, 0x2a, 0xa1, 0x46, 0x74, 0xc6, 0x3b, 0x17, 0x9c, 0x9b, 0x13, 0x79, 0x6f, 0xda, 0x4d, 0xf2,
	0x04, 0xee, 0x9c, 0x9e, 0xcc, 0x99, 0xe6, 0xca, 0xcc, 0x2a, 0x70, 0x01, 0x7c, 0x77, 0x42, 0x7b,
	0x6e, 0x59, 0x4d, 0x35, 0x8d, 0x88, 0x2f, 0xcc, 0x44, 0xfc, 0x3d, 0x80, 0xc1, 0x98, 0x15, 0x85,
	0x28, 0x50, 0xba, 0x68, 0xa5, 0x2d, 0x8f, 0xf4, 0x12, 0x72, 0x17, 0x56, 0xf3, 0x4c, 0x1b, 0x2a,
	0x12, 0x9a, 0x70, 0xc3, 0x44, 0xda, 0xbe, 0x6a, 0x29, 0x2b, 0x88, 0xf6, 0x92, 0xc7, 0x16, 0xc3,
	0x67, 0x1b, 0x73, 0x96, 0x70, 0x4d, 0x2b, 0xae, 0x0b, 0x91, 0xa9, 0xf6, 0xd2, 0x76, 0xb0, 0x13,
	0xc5, 0x91, 0x43, 0x0f, 0x1d, 0x88, 0x17, 0x19, 0x67, 0x69, 0x42, 0x8d, 0x90, 0xbc, 0x7d, 0xcd,
	0x32, 0x96, 0x11, 0xd8, 0x17, 0x92, 0x93, 0xcf, 0x60, 0x93, 0x2b, 0xd6, 0x4f, 0x79, 0x42, 0x07,
	0x2c, 0x67, 0x7d, 0x91, 0x0a, 0x23, 0x78, 0xd1, 0x5e, 0xb6, 0xf6, 0x36, 0xbc, 0xec, 0xeb, 0x29,
	0x11, 0xb9, 0x0d, 0xcb, 0x79, 0xca, 0xcc, 0x30, 0xd3, 0xb2, 0xdd, 0x72, 0x7e, 0xd5, 0xbf, 0x49,
	0x05, 0x4b, 0xfe, 0xc2, 0xb0, 0x1d, 0xec, 0x84, 0xbb, 0x3f, 0xce, 0x2b, 0x11, 0x9c, 0x95, 0xd8,
	0x5b, 0x23, 0x12, 0x16, 0xa4, 0xe8, 0xb7, 0x43, 0x6b, 0xf4, 0xfb, 0x79, 0x19, 0x95, 0xa2, 0x1f,
	0xa3, 0x9d, 0xce, 0x87, 0x10, 0x0a, 0xf5, 0x90, 0xb2, 0x24, 0xd1, 0xd4, 0x24, 0x64, 0x13, 0xae,
	0x56, 0x2c, 0x2d, 0xeb, 0xc4, 0x72, 0x3f, 0x3a, 0x7f, 0x05, 0x60, 0x33, 0x95, 0xa6, 0x0f, 0x2c,
	0x93, 0x7c, 0x00, 0x2b, 0xf8, 0xc9, 0x8b, 0x82, 0x9a, 0x93, 0xbc, 0x66, 0x87, 0x1e, 0xdb, 0x3f,
	0xc9, 0x39, 0x52, 0x44, 0x5e, 0x7d, 0x4e, 0x3d, 0xe6, 0x13, 0x2d, 0x44, 0xec, 0x91, 0x83, 0xc8,
	0xb1, 0xa5, 0x3c, 0x9c, 0x50, 0x16, 0xac, 0xcf, 0x07, 0x97, 0xe7, 0xf3, 0x94, 0x67, 0xd6, 0xf2,
	0x43, 0x6f, 0xb9, 0xf3, 0x36, 0x80, 0x35, 0xab, 0xcf, 0x4a, 0xb9, 0x32, 0xfa, 0x84, 0xbc, 0x3e,
	0x8b, 0xb5, 0x03, 0xdb, 0x03, 0x2e, 0xfb, 0xe9, 0x4f, 0x0d, 0xb8, 0x2e, 0x60, 0x1b, 0x0f, 0x5e,
	0xed, 0x09, 0x62, 0x9d, 0x3f, 0x03, 0xd8, 0x3c, 0x8f, 0x47, 0x72, 0xb8, 0x56, 0x87, 0x2a, 0xb0,
	0xa1, 0x3a, 0xbc, 0xe4, 0x8b, 0xf9, 0xe7, 0x8d, 0x6b, 0x33, 0x58, 0xdc, 0x92, 0xd1, 0xa2, 0xec,
	0xdb, 0x57, 0xbe, 0x62, 0x2b, 0xae, 0x25, 0xd9, 0x0b, 0x07, 0x90, 0x1b, 0xb0, 0x24, 0x86, 0x54,
	0x95, 0xd2, 0x3e, 0x5d, 0x14, 0x5f, 0x15, 0xc3, 0xbd, 0x52, 0x76, 0xfe, 0x08, 0xe0, 0xa6, 0xd5,
	0x57, 0xaa, 0x23, 0x95, 0xbd, 0x52, 0xd4, 0xa4, 0x95, 0x0f, 0xf2, 0xdb, 0x0b, 0x45, 0x3e, 0xd6,
	0xfc, 0x92, 0x5d, 0x3a, 0x63, 0xc7, 0x85, 0xdc, 0x8e, 0xb5, 0x03, 0x27, 0xdb, 0x4f, 0x2b, 0x17,
	0xf8, 0x03, 0xd8, 0xfa, 0x9f, 0x33, 0xe4, 0x16, 0x2c, 0x23, 0x32, 0x49, 0xf8, 0x28, 0xbe, 0x66,
	0xd2, 0xca, 0x26, 0xfb, 0x16, 0xb4, 0x50, 0xe4, 0x4a, 0x07, 0xc3, 0xb4, 0x12, 0x23, 0xf7, 0xd0,
	0x56, 0xcf, 0xef, 0xb5, 0xcf, 0x99, 0x1e, 0xd1, 0x84, 0x0f, 0xa7, 0xc2, 0xf1, 0xdb, 0x85, 0xa2,
	0x39, 0x85, 0xe3, 0x8c, 0x1d, 0x17, 0x0e, 0x3b, 0x76, 0xbe, 0xd3, 0xa3, 0xc7, 0x7c, 0x38, 0x89,
	0xc6, 0x9b, 0xc0, 0x87, 0xe3, 0xfc, 0x33, 0x64, 0x1d, 0x16, 0xb2, 0x52, 0xf8, 0x48, 0xe0, 0x57,
	0x72, 0x07, 0x42, 0xe4, 0x34, 0xd3, 0x05, 0x4c, 0x5a, 0xd5, 0xf9, 0x72, 0x17, 0x56, 0x91, 0x20,
	0xd4, 0x30, 0xa3, 0x42, 0x25, 0xbc, 0xf0, 0x79, 0xb3, 0x62, 0xd2, 0xaa, 0xa7, 0x86, 0x59, 0x0f,
	0xb1, 0x66, 0x30, 0x17, 0x67, 0x82, 0xf9, 0xcf, 0xa2, 0x2f, 0x8e, 0x99, 0xfe, 0x49, 0x3e, 0x86,
	0x75, 0x3b, 0x68, 0x12, 0x5e, 0x0c, 0xb4, 0xc8, 0x0d, 0x0e, 0x11, 0xd7, 0x96, 0xd6, 0x10, 0x7f,
	0x7c, 0x0a, 0xe3, 0x3d, 0x8b, 0x93, 0xc2, 0x70, 0x39, 0x3d, 0x02, 0xc1, 0x41, 0x76, 0xe0, 0xdd,
	0x07, 0xe2, 0x09, 0xd3, 0xda, 0xdc, 0xe4, 0xbb, 0xee, 0x24, 0xd3, 0xfa, 0xee, 0xc1, 0x2a, 0x4e,
	0x24, 0xaa, 0xb9, 0x64, 0x42, 0x09, 0x35, 0xb2, 0xb7, 0x8e, 0xe2, 0x08, 0xd1, 0xb8, 0x06, 0xc9,
	0xa7, 0xb0, 0xe1, 0xb5, 0x36, 0xe6, 0x93, 0x9b, 0x87, 0xde, 0x60, 0x63, 0x3c, 0x5d, 0x34, 0xd1,
	0x96, 0x2e, 0x9e, 0x68, 0xaf, 0x03, 0xb8, 0xae, 0xb8, 0x79, 0x95, 0xe9, 0xa3, 0xba, 0xad, 0xf2,
	0xc2, 0x8e, 0xca, 0x70, 0xf7, 0xe5, 0xdc, 0xda, 0x58, 0xbc, 0xee, 0x6d, 0x3e, 0xaa, 0x4d, 0xe2,
	0x73, 0xb0, 0xd2, 0x64, 0x54, 0xf1, 0x51, 0x66, 0x04, 0xb3, 0x01, 0x74, 0x93, 0x78, 0x0d, 0xf1,
	0xbd, 0x53, 0x98, 0x7c, 0x05, 0x5b, 0xf9, 0xf8, 0xa4, 0x10, 0x03, 0x96, 0x52, 0xc9, 0x13, 0xc1,
	0x9a, 0xde, 0xba, 0xc1, 0x7c, 0xab, 0xa6, 0x3c, 0x43, 0x46, 0xc3, 0xe7, 0x2f, 0xe1, 0xb6, 0x3b,
	0xc6, 0x8c, 0x61, 0x83, 0xb1, 0xc4, 0x2d, 0xa7, 0x54, 0xc2, 0xb8, 0x4a, 0x05, 0xfb, 0x14, 0xef,
	0x58, 0xc6, 0xa3, 0x09, 0xe1, 0x40, 0x09, 0x63, 0x2b, 0x77, 0x1b, 0xec, 0x26, 0x42, 0xab, 0x94,
	0x29, 0x5c, 0x60, 0x42, 0x97, 0xb4, 0x88, 0x1d, 0xa6, 0x4c, 0xf5, 0x92, 0xce, 0xdf, 0x8b, 0x70,
	0xfd, 0xcc, 0xf0, 0x24, 0x1d, 0x88, 0x34, 0x97, 0x76, 0x13, 0xa1, 0x92, 0xe9, 0x23, 0x5f, 0x07,
	0xa1, 0xe6, 0x12, 0xb7, 0x91, 0x67, 0x4c, 0x1f, 0x91, 0x4f, 0x80, 0x20, 0x27, 0xcd, 0xd0, 0x33,
	0x6b, 0x05, 0x5b, 0xa5, 0x2b, 0x8b, 0x35, 0xcd, 0xe5, 0x53, 0x14, 0x3c, 0xcf, 0xb4, 0xd9, 0x2b,
	0x25, 0x66, 0x3d, 0x92, 0xb1, 0x2c, 0x8e, 0x7d, 0x59, 0x2c, 0x6b, 0x2e, 0xb1, 0x24, 0x8e, 0xc9,
	0x7d, 0xd8, 0x38, 0x5d, 0xb2, 0xb0, 0xc0, 0x9c, 0x6f, 0x2e, 0xcd, 0xd6, 0x27, 0xdb, 0xd6, 0x8b,
	0xb2, 0xbf, 0xef, 0xeb, 0x6c, 0x8a, 0x9e, 0x72, 0x65, 0x93, 0x2c, 0x8a, 0x57, 0x26, 0xcc, 0xa7,
	0x5c, 0x91, 0x8f, 0x7c, 0xc5, 0x4c, 0x6b, 0xf4, 0x6b, 0x97, 0x5b, 0xce, 0x6a, 0x75, 0xef, 0x43,
	0x58, 0x13, 0x51, 0x97, 0x5b, 0xbc, 0x5a, 0x8e, 0x83, 0x8a, 0x1e, 0xc0, 0x8d, 0x41, 0x26, 0xfb,
	0x42, 0x9d, 0xb7, 0x7a, 0x45, 0xf1, 0x66, 0x2d, 0x6c, 0xbc, 0xda, 0xaf, 0x01, 0xac, 0x4f, 0x37,
	0xda, 0x54, 0x14, 0xc6, 0xbe, 0x75, 0xb8, 0xfb, 0xd3, 0xbc, 0x67, 0x40, 0xbc, 0x5a, 0x4e, 0x5a,
	0xff, 0x53, 0x51, 0x18, 0xf2, 0x4b, 0x00, 0xeb, 0xd3, 0x6d, 0xce, 0x5e, 0x06, 0xe6, 0x72, 0x99,
	0x33, 0xdd, 0x34, 0x8e, 0xb2, 0xba, 0xf1, 0xe2, 0x5d, 0xfa, 0x4b, 0xf6, 0x7f, 0xac, 0x07, 0xff,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xef, 0x74, 0x0d, 0xba, 0x7a, 0x0d, 0x00, 0x00,
}
