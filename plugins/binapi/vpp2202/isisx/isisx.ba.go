// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.6.0
//  VPP:              22.02
// source: /usr/share/vpp/api/plugins/isisx.api.json

// Package isisx contains generated bindings for API file isisx.api.
//
// Contents:
//   1 struct
//   6 messages
//
package isisx

import (
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
	_ "go.pantheon.tech/stonework/plugins/binapi/vpp2202/interface_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "isisx"
	APIVersion = "0.1.0"
	VersionCrc = 0x8b24933
)

// IsisxConnection defines type 'isisx_connection'.
type IsisxConnection struct {
	RxSwIfIndex uint32 `binapi:"u32,name=rx_sw_if_index" json:"rx_sw_if_index,omitempty"`
	TxSwIfIndex uint32 `binapi:"u32,name=tx_sw_if_index" json:"tx_sw_if_index,omitempty"`
}

// IsisxConnectionAddDel defines message 'isisx_connection_add_del'.
// InProgress: the message form may change in the future versions
type IsisxConnectionAddDel struct {
	IsAdd      uint8           `binapi:"u8,name=is_add" json:"is_add,omitempty"`
	Connection IsisxConnection `binapi:"isisx_connection,name=connection" json:"connection,omitempty"`
}

func (m *IsisxConnectionAddDel) Reset()               { *m = IsisxConnectionAddDel{} }
func (*IsisxConnectionAddDel) GetMessageName() string { return "isisx_connection_add_del" }
func (*IsisxConnectionAddDel) GetCrcString() string   { return "2bbf55c3" }
func (*IsisxConnectionAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IsisxConnectionAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsAdd
	size += 4 // m.Connection.RxSwIfIndex
	size += 4 // m.Connection.TxSwIfIndex
	return size
}
func (m *IsisxConnectionAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.IsAdd)
	buf.EncodeUint32(m.Connection.RxSwIfIndex)
	buf.EncodeUint32(m.Connection.TxSwIfIndex)
	return buf.Bytes(), nil
}
func (m *IsisxConnectionAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeUint8()
	m.Connection.RxSwIfIndex = buf.DecodeUint32()
	m.Connection.TxSwIfIndex = buf.DecodeUint32()
	return nil
}

// IsisxConnectionAddDelReply defines message 'isisx_connection_add_del_reply'.
// InProgress: the message form may change in the future versions
type IsisxConnectionAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IsisxConnectionAddDelReply) Reset()               { *m = IsisxConnectionAddDelReply{} }
func (*IsisxConnectionAddDelReply) GetMessageName() string { return "isisx_connection_add_del_reply" }
func (*IsisxConnectionAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*IsisxConnectionAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IsisxConnectionAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IsisxConnectionAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IsisxConnectionAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// IsisxConnectionDetails defines message 'isisx_connection_details'.
// InProgress: the message form may change in the future versions
type IsisxConnectionDetails struct {
	Connection IsisxConnection `binapi:"isisx_connection,name=connection" json:"connection,omitempty"`
}

func (m *IsisxConnectionDetails) Reset()               { *m = IsisxConnectionDetails{} }
func (*IsisxConnectionDetails) GetMessageName() string { return "isisx_connection_details" }
func (*IsisxConnectionDetails) GetCrcString() string   { return "4b667522" }
func (*IsisxConnectionDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IsisxConnectionDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Connection.RxSwIfIndex
	size += 4 // m.Connection.TxSwIfIndex
	return size
}
func (m *IsisxConnectionDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Connection.RxSwIfIndex)
	buf.EncodeUint32(m.Connection.TxSwIfIndex)
	return buf.Bytes(), nil
}
func (m *IsisxConnectionDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Connection.RxSwIfIndex = buf.DecodeUint32()
	m.Connection.TxSwIfIndex = buf.DecodeUint32()
	return nil
}

// IsisxConnectionDump defines message 'isisx_connection_dump'.
// InProgress: the message form may change in the future versions
type IsisxConnectionDump struct{}

func (m *IsisxConnectionDump) Reset()               { *m = IsisxConnectionDump{} }
func (*IsisxConnectionDump) GetMessageName() string { return "isisx_connection_dump" }
func (*IsisxConnectionDump) GetCrcString() string   { return "51077d14" }
func (*IsisxConnectionDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IsisxConnectionDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *IsisxConnectionDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *IsisxConnectionDump) Unmarshal(b []byte) error {
	return nil
}

// IsisxPluginGetVersion defines message 'isisx_plugin_get_version'.
// InProgress: the message form may change in the future versions
type IsisxPluginGetVersion struct{}

func (m *IsisxPluginGetVersion) Reset()               { *m = IsisxPluginGetVersion{} }
func (*IsisxPluginGetVersion) GetMessageName() string { return "isisx_plugin_get_version" }
func (*IsisxPluginGetVersion) GetCrcString() string   { return "51077d14" }
func (*IsisxPluginGetVersion) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IsisxPluginGetVersion) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *IsisxPluginGetVersion) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *IsisxPluginGetVersion) Unmarshal(b []byte) error {
	return nil
}

// IsisxPluginGetVersionReply defines message 'isisx_plugin_get_version_reply'.
// InProgress: the message form may change in the future versions
type IsisxPluginGetVersionReply struct {
	Major uint32 `binapi:"u32,name=major" json:"major,omitempty"`
	Minor uint32 `binapi:"u32,name=minor" json:"minor,omitempty"`
}

func (m *IsisxPluginGetVersionReply) Reset()               { *m = IsisxPluginGetVersionReply{} }
func (*IsisxPluginGetVersionReply) GetMessageName() string { return "isisx_plugin_get_version_reply" }
func (*IsisxPluginGetVersionReply) GetCrcString() string   { return "9b32cf86" }
func (*IsisxPluginGetVersionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IsisxPluginGetVersionReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Major
	size += 4 // m.Minor
	return size
}
func (m *IsisxPluginGetVersionReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Major)
	buf.EncodeUint32(m.Minor)
	return buf.Bytes(), nil
}
func (m *IsisxPluginGetVersionReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Major = buf.DecodeUint32()
	m.Minor = buf.DecodeUint32()
	return nil
}

func init() { file_isisx_binapi_init() }
func file_isisx_binapi_init() {
	api.RegisterMessage((*IsisxConnectionAddDel)(nil), "isisx_connection_add_del_2bbf55c3")
	api.RegisterMessage((*IsisxConnectionAddDelReply)(nil), "isisx_connection_add_del_reply_e8d4e804")
	api.RegisterMessage((*IsisxConnectionDetails)(nil), "isisx_connection_details_4b667522")
	api.RegisterMessage((*IsisxConnectionDump)(nil), "isisx_connection_dump_51077d14")
	api.RegisterMessage((*IsisxPluginGetVersion)(nil), "isisx_plugin_get_version_51077d14")
	api.RegisterMessage((*IsisxPluginGetVersionReply)(nil), "isisx_plugin_get_version_reply_9b32cf86")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*IsisxConnectionAddDel)(nil),
		(*IsisxConnectionAddDelReply)(nil),
		(*IsisxConnectionDetails)(nil),
		(*IsisxConnectionDump)(nil),
		(*IsisxPluginGetVersion)(nil),
		(*IsisxPluginGetVersionReply)(nil),
	}
}
