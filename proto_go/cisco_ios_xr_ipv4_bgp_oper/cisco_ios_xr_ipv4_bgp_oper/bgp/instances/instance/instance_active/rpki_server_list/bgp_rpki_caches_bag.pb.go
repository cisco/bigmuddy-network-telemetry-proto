// Code generated by protoc-gen-go.
// source: bgp_rpki_caches_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_rpki_server_list is a generated protocol buffer package.

It is generated from these files:
	bgp_rpki_caches_bag.proto

It has these top-level messages:
	BgpRpkiCachesBag_KEYS
	BgpRpkiCachesBag
	BgpEdmRpkiCacheType_
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_rpki_server_list

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

type BgpRpkiCachesBag_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
}

func (m *BgpRpkiCachesBag_KEYS) Reset()                    { *m = BgpRpkiCachesBag_KEYS{} }
func (m *BgpRpkiCachesBag_KEYS) String() string            { return proto.CompactTextString(m) }
func (*BgpRpkiCachesBag_KEYS) ProtoMessage()               {}
func (*BgpRpkiCachesBag_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BgpRpkiCachesBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

type BgpRpkiCachesBag struct {
	// Array of RPKI servers
	RpkiServers []*BgpEdmRpkiCacheType_ `protobuf:"bytes,50,rep,name=rpki_servers,json=rpkiServers" json:"rpki_servers,omitempty"`
}

func (m *BgpRpkiCachesBag) Reset()                    { *m = BgpRpkiCachesBag{} }
func (m *BgpRpkiCachesBag) String() string            { return proto.CompactTextString(m) }
func (*BgpRpkiCachesBag) ProtoMessage()               {}
func (*BgpRpkiCachesBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BgpRpkiCachesBag) GetRpkiServers() []*BgpEdmRpkiCacheType_ {
	if m != nil {
		return m.RpkiServers
	}
	return nil
}

type BgpEdmRpkiCacheType_ struct {
	// Server Name
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Server Preference
	Preference uint32 `protobuf:"varint,2,opt,name=preference" json:"preference,omitempty"`
	// Server TCP Port number
	Port uint32 `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	// Server Internal State
	State string `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	// Server Internal State timestamp (unix time)
	StateTime uint32 `protobuf:"varint,5,opt,name=state_time,json=stateTime" json:"state_time,omitempty"`
	// Server Shutdown
	Shutdown bool `protobuf:"varint,6,opt,name=shutdown" json:"shutdown,omitempty"`
	// Number of connection retries
	Retries uint32 `protobuf:"varint,7,opt,name=retries" json:"retries,omitempty"`
	// Server close reason
	CloseReason string `protobuf:"bytes,8,opt,name=close_reason,json=closeReason" json:"close_reason,omitempty"`
	// Server close elapsed time
	CloseTime uint32 `protobuf:"varint,9,opt,name=close_time,json=closeTime" json:"close_time,omitempty"`
	// Server close real timestamp (unix time)
	CloseTimeReal uint32 `protobuf:"varint,10,opt,name=close_time_real,json=closeTimeReal" json:"close_time_real,omitempty"`
	// Number of bytes read from the server
	ReadBytes uint32 `protobuf:"varint,11,opt,name=read_bytes,json=readBytes" json:"read_bytes,omitempty"`
	// Number of bytes written to the server
	WriteBytes uint32 `protobuf:"varint,12,opt,name=write_bytes,json=writeBytes" json:"write_bytes,omitempty"`
	// Server transport type
	Transport uint32 `protobuf:"varint,13,opt,name=transport" json:"transport,omitempty"`
	// Server SSH username
	Username string `protobuf:"bytes,14,opt,name=username" json:"username,omitempty"`
	// Server SSH password
	Password string `protobuf:"bytes,15,opt,name=password" json:"password,omitempty"`
	// Server SSH process ID
	Sshpid uint32 `protobuf:"varint,16,opt,name=sshpid" json:"sshpid,omitempty"`
	// Server Protocol state
	ProtoState string `protobuf:"bytes,17,opt,name=proto_state,json=protoState" json:"proto_state,omitempty"`
	// Server Protocol state timestamp (unix time)
	ProtoStateTime uint32 `protobuf:"varint,18,opt,name=proto_state_time,json=protoStateTime" json:"proto_state_time,omitempty"`
	// Server serial number
	Serial uint32 `protobuf:"varint,19,opt,name=serial" json:"serial,omitempty"`
	// Server nonce
	Nonce uint32 `protobuf:"varint,20,opt,name=nonce" json:"nonce,omitempty"`
	// Server refresh time (seconds)
	RefreshTime int32 `protobuf:"zigzag32,21,opt,name=refresh_time,json=refreshTime" json:"refresh_time,omitempty"`
	// Server response time (seconds)
	ResponseTime int32 `protobuf:"zigzag32,22,opt,name=response_time,json=responseTime" json:"response_time,omitempty"`
	// Server purge time (seconds)
	PurgeTime int32 `protobuf:"zigzag32,23,opt,name=purge_time,json=purgeTime" json:"purge_time,omitempty"`
	// Total IPv4 ROAs currently recv'd from server
	Ipv4Roa uint32 `protobuf:"varint,24,opt,name=ipv4_roa,json=ipv4Roa" json:"ipv4_roa,omitempty"`
	// Total IPv4 ROAs announced by the server
	Ipv4RoaAnnounce uint32 `protobuf:"varint,25,opt,name=ipv4_roa_announce,json=ipv4RoaAnnounce" json:"ipv4_roa_announce,omitempty"`
	// Total IPv4 ROAs withdrawn by the server
	Ipv4RoaWithdraw uint32 `protobuf:"varint,26,opt,name=ipv4_roa_withdraw,json=ipv4RoaWithdraw" json:"ipv4_roa_withdraw,omitempty"`
	// Total IPv6 ROAs currently recv'd from server
	Ipv6Roa uint32 `protobuf:"varint,27,opt,name=ipv6_roa,json=ipv6Roa" json:"ipv6_roa,omitempty"`
	// Total IPv6 ROAs announced by the server
	Ipv6RoaAnnounce uint32 `protobuf:"varint,28,opt,name=ipv6_roa_announce,json=ipv6RoaAnnounce" json:"ipv6_roa_announce,omitempty"`
	// Total IPv6 ROAs withdrawn by the server
	Ipv6RoaWithdraw uint32 `protobuf:"varint,29,opt,name=ipv6_roa_withdraw,json=ipv6RoaWithdraw" json:"ipv6_roa_withdraw,omitempty"`
	// Protocol Error Reason
	ProtoError string `protobuf:"bytes,30,opt,name=proto_error,json=protoError" json:"proto_error,omitempty"`
}

func (m *BgpEdmRpkiCacheType_) Reset()                    { *m = BgpEdmRpkiCacheType_{} }
func (m *BgpEdmRpkiCacheType_) String() string            { return proto.CompactTextString(m) }
func (*BgpEdmRpkiCacheType_) ProtoMessage()               {}
func (*BgpEdmRpkiCacheType_) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BgpEdmRpkiCacheType_) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BgpEdmRpkiCacheType_) GetPreference() uint32 {
	if m != nil {
		return m.Preference
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *BgpEdmRpkiCacheType_) GetStateTime() uint32 {
	if m != nil {
		return m.StateTime
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetShutdown() bool {
	if m != nil {
		return m.Shutdown
	}
	return false
}

func (m *BgpEdmRpkiCacheType_) GetRetries() uint32 {
	if m != nil {
		return m.Retries
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetCloseReason() string {
	if m != nil {
		return m.CloseReason
	}
	return ""
}

func (m *BgpEdmRpkiCacheType_) GetCloseTime() uint32 {
	if m != nil {
		return m.CloseTime
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetCloseTimeReal() uint32 {
	if m != nil {
		return m.CloseTimeReal
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetReadBytes() uint32 {
	if m != nil {
		return m.ReadBytes
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetWriteBytes() uint32 {
	if m != nil {
		return m.WriteBytes
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetTransport() uint32 {
	if m != nil {
		return m.Transport
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *BgpEdmRpkiCacheType_) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *BgpEdmRpkiCacheType_) GetSshpid() uint32 {
	if m != nil {
		return m.Sshpid
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetProtoState() string {
	if m != nil {
		return m.ProtoState
	}
	return ""
}

func (m *BgpEdmRpkiCacheType_) GetProtoStateTime() uint32 {
	if m != nil {
		return m.ProtoStateTime
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetSerial() uint32 {
	if m != nil {
		return m.Serial
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetNonce() uint32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetRefreshTime() int32 {
	if m != nil {
		return m.RefreshTime
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetResponseTime() int32 {
	if m != nil {
		return m.ResponseTime
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetPurgeTime() int32 {
	if m != nil {
		return m.PurgeTime
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetIpv4Roa() uint32 {
	if m != nil {
		return m.Ipv4Roa
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetIpv4RoaAnnounce() uint32 {
	if m != nil {
		return m.Ipv4RoaAnnounce
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetIpv4RoaWithdraw() uint32 {
	if m != nil {
		return m.Ipv4RoaWithdraw
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetIpv6Roa() uint32 {
	if m != nil {
		return m.Ipv6Roa
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetIpv6RoaAnnounce() uint32 {
	if m != nil {
		return m.Ipv6RoaAnnounce
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetIpv6RoaWithdraw() uint32 {
	if m != nil {
		return m.Ipv6RoaWithdraw
	}
	return 0
}

func (m *BgpEdmRpkiCacheType_) GetProtoError() string {
	if m != nil {
		return m.ProtoError
	}
	return ""
}

func init() {
	proto.RegisterType((*BgpRpkiCachesBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.rpki_server_list.bgp_rpki_caches_bag_KEYS")
	proto.RegisterType((*BgpRpkiCachesBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.rpki_server_list.bgp_rpki_caches_bag")
	proto.RegisterType((*BgpEdmRpkiCacheType_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.rpki_server_list.bgp_edm_rpki_cache_type_")
}

func init() { proto.RegisterFile("bgp_rpki_caches_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xd6, 0xd2, 0xbf, 0x64, 0x92, 0xf4, 0xc7, 0x2d, 0xc5, 0x2d, 0x6d, 0x09, 0xad, 0x84, 0x22,
	0x0e, 0x39, 0x14, 0x94, 0x2b, 0x02, 0xa9, 0x27, 0x24, 0x0e, 0x5b, 0x24, 0xc4, 0xc9, 0x72, 0x36,
	0xd3, 0xc4, 0x22, 0x59, 0x5b, 0x63, 0xa7, 0xa1, 0x4f, 0xc1, 0x63, 0xf0, 0x32, 0x3c, 0x14, 0xf2,
	0x78, 0xb3, 0x4d, 0x44, 0x8f, 0xdc, 0x66, 0xbe, 0xf9, 0xe6, 0xfb, 0x3c, 0x33, 0xbb, 0x70, 0x32,
	0x1c, 0x3b, 0x45, 0xee, 0x87, 0x51, 0x85, 0x2e, 0x26, 0xe8, 0xd5, 0x50, 0x8f, 0xfb, 0x8e, 0x6c,
	0xb0, 0x22, 0x2f, 0x8c, 0x2f, 0xac, 0x32, 0xd6, 0xab, 0x9f, 0xa4, 0x8c, 0xbb, 0x7f, 0xaf, 0x22,
	0xd9, 0x3a, 0xa4, 0xfe, 0x70, 0xec, 0xfa, 0xa6, 0xf4, 0x41, 0x97, 0x05, 0xfa, 0x3a, 0xaa, 0x03,
	0xa5, 0x8b, 0x60, 0xee, 0xb1, 0xcf, 0xc2, 0x1e, 0xe9, 0x1e, 0x49, 0x4d, 0x8d, 0x0f, 0x97, 0x1f,
	0x40, 0x3e, 0x61, 0xa8, 0x3e, 0xdf, 0x7c, 0xbf, 0x15, 0x57, 0xd0, 0xa9, 0xfb, 0x4b, 0x3d, 0x43,
	0x99, 0x75, 0xb3, 0x5e, 0x33, 0x6f, 0x2f, 0xc1, 0x2f, 0x7a, 0x86, 0x97, 0xbf, 0x33, 0x38, 0x7c,
	0x42, 0x41, 0xfc, 0xca, 0xa0, 0xbd, 0xe2, 0xe6, 0xe5, 0x75, 0x77, 0xa3, 0xd7, 0xba, 0x9e, 0xf6,
	0xff, 0xff, 0x10, 0xb1, 0x4f, 0xe1, 0x68, 0xb6, 0xf2, 0x06, 0x15, 0x1e, 0x1c, 0xaa, 0xbc, 0x15,
	0x91, 0xdb, 0xf4, 0x80, 0xcb, 0x3f, 0x3b, 0x69, 0xd6, 0xa7, 0x98, 0x42, 0xc0, 0xe6, 0xca, 0x88,
	0x1c, 0x8b, 0x0b, 0x00, 0x47, 0x78, 0x87, 0x84, 0x65, 0x81, 0xf2, 0x59, 0x37, 0xeb, 0x75, 0xf2,
	0x15, 0x24, 0xf6, 0x38, 0x4b, 0x41, 0x6e, 0x70, 0x85, 0x63, 0x71, 0x04, 0x5b, 0x3e, 0xe8, 0x80,
	0x72, 0x93, 0x85, 0x52, 0x22, 0xce, 0x01, 0x38, 0x50, 0xc1, 0xcc, 0x50, 0x6e, 0x31, 0xbf, 0xc9,
	0xc8, 0x57, 0x33, 0x43, 0x71, 0x0a, 0x0d, 0x3f, 0x99, 0x87, 0x91, 0x5d, 0x94, 0x72, 0xbb, 0x9b,
	0xf5, 0x1a, 0x79, 0x9d, 0x0b, 0x09, 0x3b, 0x84, 0x81, 0x0c, 0x7a, 0xb9, 0xc3, 0x7d, 0xcb, 0x54,
	0xbc, 0x86, 0x76, 0x31, 0xb5, 0x1e, 0x15, 0xa1, 0xf6, 0xb6, 0x94, 0x0d, 0x76, 0x6c, 0x31, 0x96,
	0x33, 0x14, 0x7d, 0x13, 0x85, 0x7d, 0x9b, 0xc9, 0x97, 0x11, 0xf6, 0x7d, 0x03, 0x7b, 0x8f, 0xe5,
	0x28, 0x33, 0x95, 0xc0, 0x9c, 0x4e, 0xcd, 0xc9, 0x51, 0x4f, 0xa3, 0x0c, 0xa1, 0x1e, 0xa9, 0xe1,
	0x43, 0x40, 0x2f, 0x5b, 0x49, 0x26, 0x22, 0x9f, 0x22, 0x20, 0x5e, 0x41, 0x6b, 0x41, 0x26, 0x60,
	0x55, 0x6f, 0xa7, 0x45, 0x31, 0x94, 0x08, 0x67, 0xd0, 0x0c, 0xa4, 0x4b, 0xcf, 0xdb, 0xea, 0xa4,
	0xf6, 0x1a, 0x88, 0xd3, 0xcf, 0x3d, 0x12, 0xaf, 0x7f, 0x97, 0x67, 0xa8, 0xf3, 0x58, 0x73, 0xda,
	0xfb, 0x85, 0xa5, 0x91, 0xdc, 0x4b, 0xb5, 0x65, 0x2e, 0x8e, 0x61, 0xdb, 0xfb, 0x89, 0x33, 0x23,
	0xb9, 0xcf, 0x92, 0x55, 0x16, 0x9f, 0xc3, 0xff, 0x8b, 0x4a, 0x87, 0x38, 0xe0, 0x36, 0x60, 0xe8,
	0x96, 0xaf, 0xd1, 0x83, 0xfd, 0x15, 0x42, 0xda, 0x8d, 0x60, 0x89, 0xdd, 0x47, 0x16, 0x2f, 0x28,
	0x5a, 0x20, 0x19, 0x3d, 0x95, 0x87, 0x95, 0x05, 0x67, 0xf1, 0xca, 0xa5, 0x8d, 0x1f, 0xc5, 0x11,
	0xc3, 0x29, 0x89, 0x07, 0x21, 0xbc, 0x23, 0xf4, 0x93, 0xa4, 0xf9, 0xbc, 0x9b, 0xf5, 0x0e, 0xf2,
	0x56, 0x85, 0xb1, 0xe0, 0x15, 0x74, 0x08, 0xbd, 0xb3, 0xe5, 0xf2, 0x26, 0xc7, 0xcc, 0x69, 0x2f,
	0x41, 0x26, 0x9d, 0x03, 0xb8, 0x39, 0x8d, 0x2b, 0xc6, 0x0b, 0x66, 0x34, 0x19, 0xe1, 0xf2, 0x09,
	0x34, 0xf8, 0xb7, 0x21, 0xab, 0xa5, 0x4c, 0x9f, 0x44, 0xcc, 0x73, 0xab, 0xc5, 0x5b, 0x38, 0x58,
	0x96, 0x94, 0x2e, 0x4b, 0x3b, 0x8f, 0x6f, 0x3c, 0x61, 0xce, 0x5e, 0xc5, 0xf9, 0x58, 0xc1, 0x6b,
	0xdc, 0x85, 0x09, 0x93, 0x11, 0xe9, 0x85, 0x3c, 0x5d, 0xe3, 0x7e, 0xab, 0xe0, 0xca, 0x72, 0xc0,
	0x96, 0x2f, 0x6b, 0xcb, 0xc1, 0xa3, 0xe5, 0x60, 0xdd, 0xf2, 0xac, 0x96, 0x19, 0xfc, 0x6b, 0x39,
	0x58, 0xb7, 0x3c, 0x5f, 0xe3, 0xd6, 0x96, 0xf5, 0x15, 0x91, 0xc8, 0x92, 0xbc, 0x58, 0xb9, 0xe2,
	0x4d, 0x44, 0x86, 0xdb, 0x1c, 0xbf, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xec, 0xd4, 0xb8, 0xe9,
	0x31, 0x05, 0x00, 0x00,
}
