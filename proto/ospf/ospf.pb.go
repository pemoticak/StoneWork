//
//(simplified, for full configuration tree refer to plugins/ospfplugin/yang/frr/frr-ospfd.tree)
//
//module: frr-ospfd
//augment /frr-rt:routing/frr-rt:control-plane-protocols/frr-rt:control-plane-protocol:
//+--rw ospf
//+--rw ospf
//|  +--rw router-id?              inet:ipv4-address
//+--rw passive-interface* [interface]
//|  +--rw interface    frr-interface:interface-ref
//|  +--rw address?     inet:ipv4-address
//+--rw areas
//+--rw area* [area-id]
//+--rw area-id          ospf-area-id
//+--rw nssa!
//|  +--rw no-summary?   boolean
//+--rw stub!
//+--rw no-summary?   boolean
//
//augment /frr-interface:lib/frr-interface:interface:
//+--rw ospf
//+--rw instance* [id]
//+--rw id                     uint16
//+--rw network?               enumeration
//+--rw area?                  ospf-area-id
//+--rw cost?                  uint16
//+--rw priority?              uint8
//+--rw interface-address* [address]
//+--rw address                inet:ipv4-address
//+--rw area?                  ospf-area-id
//+--rw cost?                  uint16
//+--rw priority?              uint8

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.13.0
// source: ospf/ospf.proto

package ospf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Interface_AreaType int32

const (
	Interface_NORMAL       Interface_AreaType = 0
	Interface_STUB         Interface_AreaType = 1
	Interface_TOTALLY_STUB Interface_AreaType = 2
	Interface_NSSA         Interface_AreaType = 3
	Interface_TOTALLY_NSSA Interface_AreaType = 4
)

// Enum value maps for Interface_AreaType.
var (
	Interface_AreaType_name = map[int32]string{
		0: "NORMAL",
		1: "STUB",
		2: "TOTALLY_STUB",
		3: "NSSA",
		4: "TOTALLY_NSSA",
	}
	Interface_AreaType_value = map[string]int32{
		"NORMAL":       0,
		"STUB":         1,
		"TOTALLY_STUB": 2,
		"NSSA":         3,
		"TOTALLY_NSSA": 4,
	}
)

func (x Interface_AreaType) Enum() *Interface_AreaType {
	p := new(Interface_AreaType)
	*p = x
	return p
}

func (x Interface_AreaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Interface_AreaType) Descriptor() protoreflect.EnumDescriptor {
	return file_ospf_ospf_proto_enumTypes[0].Descriptor()
}

func (Interface_AreaType) Type() protoreflect.EnumType {
	return &file_ospf_ospf_proto_enumTypes[0]
}

func (x Interface_AreaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Interface_AreaType.Descriptor instead.
func (Interface_AreaType) EnumDescriptor() ([]byte, []int) {
	return file_ospf_ospf_proto_rawDescGZIP(), []int{1, 0}
}

// Type of the network behind the interface.
type Interface_Network int32

const (
	Interface_BROADCAST           Interface_Network = 0
	Interface_NON_BROADCAST       Interface_Network = 1
	Interface_POINT_TO_MULTIPOINT Interface_Network = 2
	Interface_POINT_TO_POINT      Interface_Network = 3
)

// Enum value maps for Interface_Network.
var (
	Interface_Network_name = map[int32]string{
		0: "BROADCAST",
		1: "NON_BROADCAST",
		2: "POINT_TO_MULTIPOINT",
		3: "POINT_TO_POINT",
	}
	Interface_Network_value = map[string]int32{
		"BROADCAST":           0,
		"NON_BROADCAST":       1,
		"POINT_TO_MULTIPOINT": 2,
		"POINT_TO_POINT":      3,
	}
)

func (x Interface_Network) Enum() *Interface_Network {
	p := new(Interface_Network)
	*p = x
	return p
}

func (x Interface_Network) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Interface_Network) Descriptor() protoreflect.EnumDescriptor {
	return file_ospf_ospf_proto_enumTypes[1].Descriptor()
}

func (Interface_Network) Type() protoreflect.EnumType {
	return &file_ospf_ospf_proto_enumTypes[1]
}

func (x Interface_Network) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Interface_Network.Descriptor instead.
func (Interface_Network) EnumDescriptor() ([]byte, []int) {
	return file_ospf_ospf_proto_rawDescGZIP(), []int{1, 1}
}

// Definition of the OSPF Router.
// Only one OSPF Router instance can be started per agent.
type Router struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Router ID (required in the IP address format). Must be unique within the OSPF domain.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Router) Reset() {
	*x = Router{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ospf_ospf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Router) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Router) ProtoMessage() {}

func (x *Router) ProtoReflect() protoreflect.Message {
	mi := &file_ospf_ospf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Router.ProtoReflect.Descriptor instead.
func (*Router) Descriptor() ([]byte, []int) {
	return file_ospf_ospf_proto_rawDescGZIP(), []int{0}
}

func (x *Router) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// OSPF enabled interface.
type Interface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Logical name of a VPP interface.
	VppInterface string `protobuf:"bytes,1,opt,name=vpp_interface,json=vppInterface,proto3" json:"vpp_interface,omitempty"`
	// Enable to make this OSPF Passive interface (in that case the attributes below are ignored).
	IsPassive bool `protobuf:"varint,2,opt,name=is_passive,json=isPassive,proto3" json:"is_passive,omitempty"`
	// ID of area behind the interface.
	// The ID should be in the dotted number form (x.x.x.x).
	Area     string             `protobuf:"bytes,3,opt,name=area,proto3" json:"area,omitempty"`
	AreaType Interface_AreaType `protobuf:"varint,4,opt,name=area_type,json=areaType,proto3,enum=ospf.Interface_AreaType" json:"area_type,omitempty"`
	Network  Interface_Network  `protobuf:"varint,5,opt,name=network,proto3,enum=ospf.Interface_Network" json:"network,omitempty"`
	// Interface cost in the range <1,65535>.
	Cost uint32 `protobuf:"varint,6,opt,name=cost,proto3" json:"cost,omitempty"`
	// Router priority in the range <0,255>.
	// Used for the election of a Designed Router.
	// The router with the highest priority will be more eligible to become Designated Router.
	// Setting the value to 0, makes the router ineligible to become Designated Router. The default value is 1.
	Priority uint32 `protobuf:"varint,7,opt,name=priority,proto3" json:"priority,omitempty"`
	// Hello interval in the range <1-65535>.
	HelloInterval uint32 `protobuf:"varint,8,opt,name=hello_interval,json=helloInterval,proto3" json:"hello_interval,omitempty"`
	// Dead interval in the range <1-65535>.
	DeadInterval uint32 `protobuf:"varint,9,opt,name=dead_interval,json=deadInterval,proto3" json:"dead_interval,omitempty"`
	// Retransmit interval in the range <1-65535>.
	RetransmitInterval uint32 `protobuf:"varint,10,opt,name=retransmit_interval,json=retransmitInterval,proto3" json:"retransmit_interval,omitempty"`
	// Transmit delay in the range <1-65535>.
	TransmitDelay uint32 `protobuf:"varint,11,opt,name=transmit_delay,json=transmitDelay,proto3" json:"transmit_delay,omitempty"`
}

func (x *Interface) Reset() {
	*x = Interface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ospf_ospf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interface) ProtoMessage() {}

func (x *Interface) ProtoReflect() protoreflect.Message {
	mi := &file_ospf_ospf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interface.ProtoReflect.Descriptor instead.
func (*Interface) Descriptor() ([]byte, []int) {
	return file_ospf_ospf_proto_rawDescGZIP(), []int{1}
}

func (x *Interface) GetVppInterface() string {
	if x != nil {
		return x.VppInterface
	}
	return ""
}

func (x *Interface) GetIsPassive() bool {
	if x != nil {
		return x.IsPassive
	}
	return false
}

func (x *Interface) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *Interface) GetAreaType() Interface_AreaType {
	if x != nil {
		return x.AreaType
	}
	return Interface_NORMAL
}

func (x *Interface) GetNetwork() Interface_Network {
	if x != nil {
		return x.Network
	}
	return Interface_BROADCAST
}

func (x *Interface) GetCost() uint32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Interface) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *Interface) GetHelloInterval() uint32 {
	if x != nil {
		return x.HelloInterval
	}
	return 0
}

func (x *Interface) GetDeadInterval() uint32 {
	if x != nil {
		return x.DeadInterval
	}
	return 0
}

func (x *Interface) GetRetransmitInterval() uint32 {
	if x != nil {
		return x.RetransmitInterval
	}
	return 0
}

func (x *Interface) GetTransmitDelay() uint32 {
	if x != nil {
		return x.TransmitDelay
	}
	return 0
}

var File_ospf_ospf_proto protoreflect.FileDescriptor

var file_ospf_ospf_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x73, 0x70, 0x66, 0x2f, 0x6f, 0x73, 0x70, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6f, 0x73, 0x70, 0x66, 0x22, 0x18, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xcb, 0x04, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x76, 0x70, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x70, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x69,
	0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x61, 0x73, 0x73,
	0x69, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x12, 0x35, 0x0a, 0x09, 0x61, 0x72, 0x65, 0x61, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6f, 0x73, 0x70,
	0x66, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x41, 0x72, 0x65, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31,
	0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x17, 0x2e, 0x6f, 0x73, 0x70, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x25, 0x0a, 0x0e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x61, 0x64,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x64, 0x65, 0x61, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2f, 0x0a,
	0x13, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x72, 0x65, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x25,
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74,
	0x44, 0x65, 0x6c, 0x61, 0x79, 0x22, 0x4e, 0x0a, 0x08, 0x41, 0x72, 0x65, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x53, 0x54, 0x55, 0x42, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x4f, 0x54, 0x41, 0x4c,
	0x4c, 0x59, 0x5f, 0x53, 0x54, 0x55, 0x42, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x53, 0x53,
	0x41, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x4e,
	0x53, 0x53, 0x41, 0x10, 0x04, 0x22, 0x58, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x0d, 0x0a, 0x09, 0x42, 0x52, 0x4f, 0x41, 0x44, 0x43, 0x41, 0x53, 0x54, 0x10, 0x00, 0x12,
	0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x4e, 0x5f, 0x42, 0x52, 0x4f, 0x41, 0x44, 0x43, 0x41, 0x53, 0x54,
	0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x54, 0x4f, 0x5f, 0x4d,
	0x55, 0x4c, 0x54, 0x49, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x50,
	0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x54, 0x4f, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x03, 0x42,
	0x27, 0x5a, 0x25, 0x70, 0x61, 0x6e, 0x74, 0x68, 0x65, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x63, 0x68,
	0x2f, 0x63, 0x6e, 0x66, 0x2d, 0x66, 0x72, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f,
	0x73, 0x70, 0x66, 0x3b, 0x6f, 0x73, 0x70, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ospf_ospf_proto_rawDescOnce sync.Once
	file_ospf_ospf_proto_rawDescData = file_ospf_ospf_proto_rawDesc
)

func file_ospf_ospf_proto_rawDescGZIP() []byte {
	file_ospf_ospf_proto_rawDescOnce.Do(func() {
		file_ospf_ospf_proto_rawDescData = protoimpl.X.CompressGZIP(file_ospf_ospf_proto_rawDescData)
	})
	return file_ospf_ospf_proto_rawDescData
}

var file_ospf_ospf_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ospf_ospf_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ospf_ospf_proto_goTypes = []interface{}{
	(Interface_AreaType)(0), // 0: ospf.Interface.AreaType
	(Interface_Network)(0),  // 1: ospf.Interface.Network
	(*Router)(nil),          // 2: ospf.Router
	(*Interface)(nil),       // 3: ospf.Interface
}
var file_ospf_ospf_proto_depIdxs = []int32{
	0, // 0: ospf.Interface.area_type:type_name -> ospf.Interface.AreaType
	1, // 1: ospf.Interface.network:type_name -> ospf.Interface.Network
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ospf_ospf_proto_init() }
func file_ospf_ospf_proto_init() {
	if File_ospf_ospf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ospf_ospf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Router); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ospf_ospf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interface); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ospf_ospf_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ospf_ospf_proto_goTypes,
		DependencyIndexes: file_ospf_ospf_proto_depIdxs,
		EnumInfos:         file_ospf_ospf_proto_enumTypes,
		MessageInfos:      file_ospf_ospf_proto_msgTypes,
	}.Build()
	File_ospf_ospf_proto = out.File
	file_ospf_ospf_proto_rawDesc = nil
	file_ospf_ospf_proto_goTypes = nil
	file_ospf_ospf_proto_depIdxs = nil
}
