// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.8.0
//  VPP:              23.06
// source: plugins/abx.api.json

// Package abx contains generated bindings for API file abx.api.
//
// Contents:
// -  2 structs
// - 10 messages
package abx

import (
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
	ethernet_types "go.pantheon.tech/stonework/plugins/binapi/vpp2306/ethernet_types"
	_ "go.pantheon.tech/stonework/plugins/binapi/vpp2306/interface_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "abx"
	APIVersion = "0.1.0"
	VersionCrc = 0x558e384
)

// AbxInterfaceAttach defines type 'abx_interface_attach'.
type AbxInterfaceAttach struct {
	PolicyID    uint32 `binapi:"u32,name=policy_id" json:"policy_id,omitempty"`
	Priority    uint32 `binapi:"u32,name=priority" json:"priority,omitempty"`
	RxSwIfIndex uint32 `binapi:"u32,name=rx_sw_if_index" json:"rx_sw_if_index,omitempty"`
}

// AbxPolicy defines type 'abx_policy'.
type AbxPolicy struct {
	PolicyID    uint32                    `binapi:"u32,name=policy_id" json:"policy_id,omitempty"`
	ACLIndex    uint32                    `binapi:"u32,name=acl_index" json:"acl_index,omitempty"`
	TxSwIfIndex uint32                    `binapi:"u32,name=tx_sw_if_index" json:"tx_sw_if_index,omitempty"`
	DstMac      ethernet_types.MacAddress `binapi:"mac_address,name=dst_mac" json:"dst_mac,omitempty"`
}

// AbxInterfaceAttachDetach defines message 'abx_interface_attach_detach'.
type AbxInterfaceAttachDetach struct {
	IsAttach uint8              `binapi:"u8,name=is_attach" json:"is_attach,omitempty"`
	Attach   AbxInterfaceAttach `binapi:"abx_interface_attach,name=attach" json:"attach,omitempty"`
}

func (m *AbxInterfaceAttachDetach) Reset()               { *m = AbxInterfaceAttachDetach{} }
func (*AbxInterfaceAttachDetach) GetMessageName() string { return "abx_interface_attach_detach" }
func (*AbxInterfaceAttachDetach) GetCrcString() string   { return "a09d5b0c" }
func (*AbxInterfaceAttachDetach) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AbxInterfaceAttachDetach) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsAttach
	size += 4 // m.Attach.PolicyID
	size += 4 // m.Attach.Priority
	size += 4 // m.Attach.RxSwIfIndex
	return size
}
func (m *AbxInterfaceAttachDetach) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.IsAttach)
	buf.EncodeUint32(m.Attach.PolicyID)
	buf.EncodeUint32(m.Attach.Priority)
	buf.EncodeUint32(m.Attach.RxSwIfIndex)
	return buf.Bytes(), nil
}
func (m *AbxInterfaceAttachDetach) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAttach = buf.DecodeUint8()
	m.Attach.PolicyID = buf.DecodeUint32()
	m.Attach.Priority = buf.DecodeUint32()
	m.Attach.RxSwIfIndex = buf.DecodeUint32()
	return nil
}

// AbxInterfaceAttachDetachReply defines message 'abx_interface_attach_detach_reply'.
type AbxInterfaceAttachDetachReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *AbxInterfaceAttachDetachReply) Reset() { *m = AbxInterfaceAttachDetachReply{} }
func (*AbxInterfaceAttachDetachReply) GetMessageName() string {
	return "abx_interface_attach_detach_reply"
}
func (*AbxInterfaceAttachDetachReply) GetCrcString() string { return "e8d4e804" }
func (*AbxInterfaceAttachDetachReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AbxInterfaceAttachDetachReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *AbxInterfaceAttachDetachReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *AbxInterfaceAttachDetachReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// AbxInterfaceAttachDetails defines message 'abx_interface_attach_details'.
type AbxInterfaceAttachDetails struct {
	Attach AbxInterfaceAttach `binapi:"abx_interface_attach,name=attach" json:"attach,omitempty"`
}

func (m *AbxInterfaceAttachDetails) Reset()               { *m = AbxInterfaceAttachDetails{} }
func (*AbxInterfaceAttachDetails) GetMessageName() string { return "abx_interface_attach_details" }
func (*AbxInterfaceAttachDetails) GetCrcString() string   { return "e7369b44" }
func (*AbxInterfaceAttachDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AbxInterfaceAttachDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Attach.PolicyID
	size += 4 // m.Attach.Priority
	size += 4 // m.Attach.RxSwIfIndex
	return size
}
func (m *AbxInterfaceAttachDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Attach.PolicyID)
	buf.EncodeUint32(m.Attach.Priority)
	buf.EncodeUint32(m.Attach.RxSwIfIndex)
	return buf.Bytes(), nil
}
func (m *AbxInterfaceAttachDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Attach.PolicyID = buf.DecodeUint32()
	m.Attach.Priority = buf.DecodeUint32()
	m.Attach.RxSwIfIndex = buf.DecodeUint32()
	return nil
}

// AbxInterfaceAttachDump defines message 'abx_interface_attach_dump'.
type AbxInterfaceAttachDump struct{}

func (m *AbxInterfaceAttachDump) Reset()               { *m = AbxInterfaceAttachDump{} }
func (*AbxInterfaceAttachDump) GetMessageName() string { return "abx_interface_attach_dump" }
func (*AbxInterfaceAttachDump) GetCrcString() string   { return "51077d14" }
func (*AbxInterfaceAttachDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AbxInterfaceAttachDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *AbxInterfaceAttachDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *AbxInterfaceAttachDump) Unmarshal(b []byte) error {
	return nil
}

// AbxPluginGetVersion defines message 'abx_plugin_get_version'.
type AbxPluginGetVersion struct{}

func (m *AbxPluginGetVersion) Reset()               { *m = AbxPluginGetVersion{} }
func (*AbxPluginGetVersion) GetMessageName() string { return "abx_plugin_get_version" }
func (*AbxPluginGetVersion) GetCrcString() string   { return "51077d14" }
func (*AbxPluginGetVersion) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AbxPluginGetVersion) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *AbxPluginGetVersion) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *AbxPluginGetVersion) Unmarshal(b []byte) error {
	return nil
}

// AbxPluginGetVersionReply defines message 'abx_plugin_get_version_reply'.
type AbxPluginGetVersionReply struct {
	Major uint32 `binapi:"u32,name=major" json:"major,omitempty"`
	Minor uint32 `binapi:"u32,name=minor" json:"minor,omitempty"`
}

func (m *AbxPluginGetVersionReply) Reset()               { *m = AbxPluginGetVersionReply{} }
func (*AbxPluginGetVersionReply) GetMessageName() string { return "abx_plugin_get_version_reply" }
func (*AbxPluginGetVersionReply) GetCrcString() string   { return "9b32cf86" }
func (*AbxPluginGetVersionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AbxPluginGetVersionReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Major
	size += 4 // m.Minor
	return size
}
func (m *AbxPluginGetVersionReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Major)
	buf.EncodeUint32(m.Minor)
	return buf.Bytes(), nil
}
func (m *AbxPluginGetVersionReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Major = buf.DecodeUint32()
	m.Minor = buf.DecodeUint32()
	return nil
}

// AbxPolicyAddDel defines message 'abx_policy_add_del'.
type AbxPolicyAddDel struct {
	IsAdd  uint8     `binapi:"u8,name=is_add" json:"is_add,omitempty"`
	Policy AbxPolicy `binapi:"abx_policy,name=policy" json:"policy,omitempty"`
}

func (m *AbxPolicyAddDel) Reset()               { *m = AbxPolicyAddDel{} }
func (*AbxPolicyAddDel) GetMessageName() string { return "abx_policy_add_del" }
func (*AbxPolicyAddDel) GetCrcString() string   { return "f5ab75d9" }
func (*AbxPolicyAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AbxPolicyAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1     // m.IsAdd
	size += 4     // m.Policy.PolicyID
	size += 4     // m.Policy.ACLIndex
	size += 4     // m.Policy.TxSwIfIndex
	size += 1 * 6 // m.Policy.DstMac
	return size
}
func (m *AbxPolicyAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.IsAdd)
	buf.EncodeUint32(m.Policy.PolicyID)
	buf.EncodeUint32(m.Policy.ACLIndex)
	buf.EncodeUint32(m.Policy.TxSwIfIndex)
	buf.EncodeBytes(m.Policy.DstMac[:], 6)
	return buf.Bytes(), nil
}
func (m *AbxPolicyAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeUint8()
	m.Policy.PolicyID = buf.DecodeUint32()
	m.Policy.ACLIndex = buf.DecodeUint32()
	m.Policy.TxSwIfIndex = buf.DecodeUint32()
	copy(m.Policy.DstMac[:], buf.DecodeBytes(6))
	return nil
}

// AbxPolicyAddDelReply defines message 'abx_policy_add_del_reply'.
type AbxPolicyAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *AbxPolicyAddDelReply) Reset()               { *m = AbxPolicyAddDelReply{} }
func (*AbxPolicyAddDelReply) GetMessageName() string { return "abx_policy_add_del_reply" }
func (*AbxPolicyAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*AbxPolicyAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AbxPolicyAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *AbxPolicyAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *AbxPolicyAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// AbxPolicyDetails defines message 'abx_policy_details'.
type AbxPolicyDetails struct {
	Policy AbxPolicy `binapi:"abx_policy,name=policy" json:"policy,omitempty"`
}

func (m *AbxPolicyDetails) Reset()               { *m = AbxPolicyDetails{} }
func (*AbxPolicyDetails) GetMessageName() string { return "abx_policy_details" }
func (*AbxPolicyDetails) GetCrcString() string   { return "1833567f" }
func (*AbxPolicyDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AbxPolicyDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.Policy.PolicyID
	size += 4     // m.Policy.ACLIndex
	size += 4     // m.Policy.TxSwIfIndex
	size += 1 * 6 // m.Policy.DstMac
	return size
}
func (m *AbxPolicyDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Policy.PolicyID)
	buf.EncodeUint32(m.Policy.ACLIndex)
	buf.EncodeUint32(m.Policy.TxSwIfIndex)
	buf.EncodeBytes(m.Policy.DstMac[:], 6)
	return buf.Bytes(), nil
}
func (m *AbxPolicyDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Policy.PolicyID = buf.DecodeUint32()
	m.Policy.ACLIndex = buf.DecodeUint32()
	m.Policy.TxSwIfIndex = buf.DecodeUint32()
	copy(m.Policy.DstMac[:], buf.DecodeBytes(6))
	return nil
}

// AbxPolicyDump defines message 'abx_policy_dump'.
type AbxPolicyDump struct{}

func (m *AbxPolicyDump) Reset()               { *m = AbxPolicyDump{} }
func (*AbxPolicyDump) GetMessageName() string { return "abx_policy_dump" }
func (*AbxPolicyDump) GetCrcString() string   { return "51077d14" }
func (*AbxPolicyDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AbxPolicyDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *AbxPolicyDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *AbxPolicyDump) Unmarshal(b []byte) error {
	return nil
}

func init() { file_abx_binapi_init() }
func file_abx_binapi_init() {
	api.RegisterMessage((*AbxInterfaceAttachDetach)(nil), "abx_interface_attach_detach_a09d5b0c")
	api.RegisterMessage((*AbxInterfaceAttachDetachReply)(nil), "abx_interface_attach_detach_reply_e8d4e804")
	api.RegisterMessage((*AbxInterfaceAttachDetails)(nil), "abx_interface_attach_details_e7369b44")
	api.RegisterMessage((*AbxInterfaceAttachDump)(nil), "abx_interface_attach_dump_51077d14")
	api.RegisterMessage((*AbxPluginGetVersion)(nil), "abx_plugin_get_version_51077d14")
	api.RegisterMessage((*AbxPluginGetVersionReply)(nil), "abx_plugin_get_version_reply_9b32cf86")
	api.RegisterMessage((*AbxPolicyAddDel)(nil), "abx_policy_add_del_f5ab75d9")
	api.RegisterMessage((*AbxPolicyAddDelReply)(nil), "abx_policy_add_del_reply_e8d4e804")
	api.RegisterMessage((*AbxPolicyDetails)(nil), "abx_policy_details_1833567f")
	api.RegisterMessage((*AbxPolicyDump)(nil), "abx_policy_dump_51077d14")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*AbxInterfaceAttachDetach)(nil),
		(*AbxInterfaceAttachDetachReply)(nil),
		(*AbxInterfaceAttachDetails)(nil),
		(*AbxInterfaceAttachDump)(nil),
		(*AbxPluginGetVersion)(nil),
		(*AbxPluginGetVersionReply)(nil),
		(*AbxPolicyAddDel)(nil),
		(*AbxPolicyAddDelReply)(nil),
		(*AbxPolicyDetails)(nil),
		(*AbxPolicyDump)(nil),
	}
}
