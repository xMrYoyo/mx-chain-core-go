// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scheduled.proto

package scheduled

import (
	bytes "bytes"
	fmt "fmt"
	github_com_ElrondNetwork_elrond_go_core_data "github.com/ElrondNetwork/elrond-go-core/data"
	block "github.com/ElrondNetwork/elrond-go-core/data/block"
	smartContractResult "github.com/ElrondNetwork/elrond-go-core/data/smartContractResult"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	io "io"
	math "math"
	math_big "math/big"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ScheduledMiniBlocks struct {
	MiniBlocks []*block.MiniBlock `protobuf:"bytes,1,rep,name=MiniBlocks,proto3" json:"MiniBlocks,omitempty"`
}

func (m *ScheduledMiniBlocks) Reset()      { *m = ScheduledMiniBlocks{} }
func (*ScheduledMiniBlocks) ProtoMessage() {}
func (*ScheduledMiniBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80076f37bd30c16, []int{0}
}
func (m *ScheduledMiniBlocks) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduledMiniBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ScheduledMiniBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduledMiniBlocks.Merge(m, src)
}
func (m *ScheduledMiniBlocks) XXX_Size() int {
	return m.Size()
}
func (m *ScheduledMiniBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduledMiniBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduledMiniBlocks proto.InternalMessageInfo

func (m *ScheduledMiniBlocks) GetMiniBlocks() []*block.MiniBlock {
	if m != nil {
		return m.MiniBlocks
	}
	return nil
}

type GasAndFees struct {
	AccumulatedFees *math_big.Int `protobuf:"bytes,1,opt,name=AccumulatedFees,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go-core/data.BigIntCaster" json:"AccumulatedFees,omitempty"`
	DeveloperFees   *math_big.Int `protobuf:"bytes,2,opt,name=DeveloperFees,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go-core/data.BigIntCaster" json:"DeveloperFees,omitempty"`
	GasProvided     uint64        `protobuf:"varint,3,opt,name=GasProvided,proto3" json:"GasProvided,omitempty"`
	GasPenalized    uint64        `protobuf:"varint,4,opt,name=GasPenalized,proto3" json:"GasPenalized,omitempty"`
	GasRefunded     uint64        `protobuf:"varint,5,opt,name=GasRefunded,proto3" json:"GasRefunded,omitempty"`
}

func (m *GasAndFees) Reset()      { *m = GasAndFees{} }
func (*GasAndFees) ProtoMessage() {}
func (*GasAndFees) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80076f37bd30c16, []int{1}
}
func (m *GasAndFees) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GasAndFees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *GasAndFees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasAndFees.Merge(m, src)
}
func (m *GasAndFees) XXX_Size() int {
	return m.Size()
}
func (m *GasAndFees) XXX_DiscardUnknown() {
	xxx_messageInfo_GasAndFees.DiscardUnknown(m)
}

var xxx_messageInfo_GasAndFees proto.InternalMessageInfo

func (m *GasAndFees) GetAccumulatedFees() *math_big.Int {
	if m != nil {
		return m.AccumulatedFees
	}
	return nil
}

func (m *GasAndFees) GetDeveloperFees() *math_big.Int {
	if m != nil {
		return m.DeveloperFees
	}
	return nil
}

func (m *GasAndFees) GetGasProvided() uint64 {
	if m != nil {
		return m.GasProvided
	}
	return 0
}

func (m *GasAndFees) GetGasPenalized() uint64 {
	if m != nil {
		return m.GasPenalized
	}
	return 0
}

func (m *GasAndFees) GetGasRefunded() uint64 {
	if m != nil {
		return m.GasRefunded
	}
	return 0
}

type ScheduledSCRs struct {
	RootHash            []byte                                     `protobuf:"bytes,1,opt,name=rootHash,proto3" json:"rootHash,omitempty"`
	Scrs                []*smartContractResult.SmartContractResult `protobuf:"bytes,2,rep,name=scrs,proto3" json:"scrs,omitempty"`
	InvalidTxHashes     [][]byte                                   `protobuf:"bytes,3,rep,name=invalidTxHashes,proto3" json:"invalidTxHashes,omitempty"`
	ScheduledMiniBlocks map[uint32]*ScheduledMiniBlocks            `protobuf:"bytes,4,rep,name=scheduledMiniBlocks,proto3" json:"scheduledMiniBlocks,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	GasAndFees          *GasAndFees                                `protobuf:"bytes,5,opt,name=gasAndFees,proto3" json:"gasAndFees,omitempty"`
}

func (m *ScheduledSCRs) Reset()      { *m = ScheduledSCRs{} }
func (*ScheduledSCRs) ProtoMessage() {}
func (*ScheduledSCRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80076f37bd30c16, []int{2}
}
func (m *ScheduledSCRs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduledSCRs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ScheduledSCRs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduledSCRs.Merge(m, src)
}
func (m *ScheduledSCRs) XXX_Size() int {
	return m.Size()
}
func (m *ScheduledSCRs) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduledSCRs.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduledSCRs proto.InternalMessageInfo

func (m *ScheduledSCRs) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

func (m *ScheduledSCRs) GetScrs() []*smartContractResult.SmartContractResult {
	if m != nil {
		return m.Scrs
	}
	return nil
}

func (m *ScheduledSCRs) GetInvalidTxHashes() [][]byte {
	if m != nil {
		return m.InvalidTxHashes
	}
	return nil
}

func (m *ScheduledSCRs) GetScheduledMiniBlocks() map[uint32]*ScheduledMiniBlocks {
	if m != nil {
		return m.ScheduledMiniBlocks
	}
	return nil
}

func (m *ScheduledSCRs) GetGasAndFees() *GasAndFees {
	if m != nil {
		return m.GasAndFees
	}
	return nil
}

func init() {
	proto.RegisterType((*ScheduledMiniBlocks)(nil), "proto.ScheduledMiniBlocks")
	proto.RegisterType((*GasAndFees)(nil), "proto.GasAndFees")
	proto.RegisterType((*ScheduledSCRs)(nil), "proto.ScheduledSCRs")
	proto.RegisterMapType((map[uint32]*ScheduledMiniBlocks)(nil), "proto.ScheduledSCRs.ScheduledMiniBlocksEntry")
}

func init() { proto.RegisterFile("scheduled.proto", fileDescriptor_f80076f37bd30c16) }

var fileDescriptor_f80076f37bd30c16 = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x31, 0x6f, 0xd3, 0x4c,
	0x18, 0xf6, 0x25, 0xe9, 0xa7, 0x8f, 0x4b, 0xa2, 0x94, 0xeb, 0x62, 0x65, 0x38, 0xa2, 0x4c, 0x59,
	0x62, 0x97, 0xb2, 0x20, 0x90, 0x90, 0x9a, 0xd0, 0x86, 0x22, 0x81, 0xd0, 0x85, 0xa9, 0x0b, 0x3a,
	0xdb, 0x57, 0xc7, 0x8a, 0xe3, 0x8b, 0xee, 0xce, 0x81, 0x32, 0xf1, 0x13, 0xf8, 0x19, 0x88, 0x5f,
	0xc2, 0x98, 0x31, 0x03, 0x12, 0xc4, 0x59, 0x18, 0xfb, 0x07, 0x90, 0x90, 0xcf, 0x89, 0xeb, 0x36,
	0x61, 0xa8, 0xc4, 0x92, 0xdc, 0xfb, 0xdc, 0x73, 0xcf, 0xe3, 0xf7, 0x7d, 0x1f, 0xd8, 0x90, 0xee,
	0x88, 0x79, 0x71, 0xc8, 0x3c, 0x6b, 0x2a, 0xb8, 0xe2, 0x68, 0x4f, 0xff, 0x35, 0xbb, 0x7e, 0xa0,
	0x46, 0xb1, 0x63, 0xb9, 0x7c, 0x62, 0xfb, 0xdc, 0xe7, 0xb6, 0x86, 0x9d, 0xf8, 0x42, 0x57, 0xba,
	0xd0, 0xa7, 0xec, 0x55, 0xf3, 0xbc, 0x40, 0x3f, 0x09, 0x05, 0x8f, 0xbc, 0xd7, 0x4c, 0xbd, 0xe7,
	0x62, 0x6c, 0x33, 0x5d, 0x75, 0x7d, 0xde, 0x75, 0xb9, 0x60, 0xb6, 0x47, 0x15, 0xb5, 0xe5, 0x84,
	0x0a, 0xd5, 0xe7, 0x91, 0x12, 0xd4, 0x55, 0x84, 0xc9, 0x38, 0x54, 0xbb, 0xb0, 0xb5, 0xf6, 0xb3,
	0x3b, 0x69, 0x3b, 0x21, 0x77, 0xc7, 0xd9, 0x6f, 0xf6, 0xbe, 0x3d, 0x80, 0x07, 0xc3, 0x4d, 0x93,
	0xaf, 0x82, 0x28, 0xe8, 0xa5, 0x77, 0x12, 0x1d, 0x42, 0x78, 0x5d, 0x99, 0xa0, 0x55, 0xee, 0x54,
	0x8f, 0xf6, 0xb3, 0x27, 0x56, 0x7e, 0x41, 0x0a, 0x9c, 0xf6, 0xf7, 0x12, 0x84, 0x03, 0x2a, 0x8f,
	0x23, 0xef, 0x94, 0x31, 0x89, 0x14, 0x6c, 0x1c, 0xbb, 0x6e, 0x3c, 0x89, 0x43, 0xaa, 0x98, 0x86,
	0x4c, 0xd0, 0x02, 0x9d, 0x5a, 0xef, 0xe5, 0xd7, 0x1f, 0x0f, 0x4e, 0x27, 0x54, 0x8d, 0x6c, 0x27,
	0xf0, 0xad, 0xb3, 0x48, 0x3d, 0xbd, 0x4b, 0x07, 0x56, 0x2f, 0xf0, 0xcf, 0x22, 0xd5, 0xa7, 0x52,
	0x31, 0x41, 0x6e, 0x5b, 0xa0, 0x29, 0xac, 0x3f, 0x67, 0x33, 0x16, 0xf2, 0x29, 0x13, 0xda, 0xb3,
	0xf4, 0xcf, 0x3d, 0x6f, 0x1a, 0xa0, 0x16, 0xac, 0x0e, 0xa8, 0x7c, 0x23, 0xf8, 0x2c, 0xf0, 0x98,
	0x67, 0x96, 0x5b, 0xa0, 0x53, 0x21, 0x45, 0x08, 0xb5, 0x61, 0x2d, 0x2d, 0x59, 0x44, 0xc3, 0xe0,
	0x23, 0xf3, 0xcc, 0x8a, 0xa6, 0xdc, 0xc0, 0xd6, 0x2a, 0x84, 0x5d, 0xc4, 0x51, 0xaa, 0xb2, 0x97,
	0xab, 0x6c, 0xa0, 0xf6, 0xef, 0x12, 0xac, 0xe7, 0x8b, 0x1a, 0xf6, 0x89, 0x44, 0x4d, 0xf8, 0xbf,
	0xe0, 0x5c, 0xbd, 0xa0, 0x72, 0x94, 0x8d, 0x96, 0xe4, 0x35, 0xb2, 0x60, 0x45, 0xba, 0x22, 0x6d,
	0x3f, 0x5d, 0x5c, 0x73, 0xbd, 0xb8, 0xe1, 0x76, 0x8a, 0x88, 0xe6, 0xa1, 0x0e, 0x6c, 0x04, 0xd1,
	0x8c, 0x86, 0x81, 0xf7, 0xf6, 0x43, 0x2a, 0xc0, 0xa4, 0x59, 0x6e, 0x95, 0x3b, 0x35, 0x72, 0x1b,
	0x46, 0xef, 0xe0, 0x81, 0xdc, 0xce, 0x8b, 0x59, 0xd1, 0x46, 0xdd, 0x8d, 0x51, 0xf1, 0x43, 0xad,
	0x1d, 0xf9, 0x3a, 0x89, 0x94, 0xb8, 0x24, 0xbb, 0x94, 0xd0, 0x43, 0x08, 0xfd, 0x3c, 0x46, 0x7a,
	0x12, 0xd5, 0xa3, 0xfb, 0x6b, 0xdd, 0xeb, 0x7c, 0x91, 0x02, 0xa9, 0xe9, 0x40, 0xf3, 0x6f, 0x1e,
	0x68, 0x1f, 0x96, 0xc7, 0xec, 0x52, 0x0f, 0xa8, 0x4e, 0xd2, 0x23, 0x3a, 0x84, 0x7b, 0x33, 0x1a,
	0xc6, 0x4c, 0x67, 0xa3, 0x30, 0x9c, 0x6d, 0x05, 0x92, 0x11, 0x9f, 0x94, 0x1e, 0x83, 0x5e, 0x7f,
	0xbe, 0xc4, 0xc6, 0x62, 0x89, 0x8d, 0xab, 0x25, 0x06, 0x9f, 0x12, 0x0c, 0xbe, 0x24, 0x18, 0x7c,
	0x4b, 0x30, 0x98, 0x27, 0x18, 0x2c, 0x12, 0x0c, 0x7e, 0x26, 0x18, 0xfc, 0x4a, 0xb0, 0x71, 0x95,
	0x60, 0xf0, 0x79, 0x85, 0x8d, 0xf9, 0x0a, 0x1b, 0x8b, 0x15, 0x36, 0xce, 0xef, 0xe5, 0x5d, 0x3a,
	0xff, 0x69, 0xab, 0x47, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x07, 0xa8, 0x09, 0xd2, 0x58, 0x04,
	0x00, 0x00,
}

func (this *ScheduledMiniBlocks) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ScheduledMiniBlocks)
	if !ok {
		that2, ok := that.(ScheduledMiniBlocks)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.MiniBlocks) != len(that1.MiniBlocks) {
		return false
	}
	for i := range this.MiniBlocks {
		if !this.MiniBlocks[i].Equal(that1.MiniBlocks[i]) {
			return false
		}
	}
	return true
}
func (this *GasAndFees) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GasAndFees)
	if !ok {
		that2, ok := that.(GasAndFees)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_core_data.BigIntCaster{}
		if !__caster.Equal(this.AccumulatedFees, that1.AccumulatedFees) {
			return false
		}
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_core_data.BigIntCaster{}
		if !__caster.Equal(this.DeveloperFees, that1.DeveloperFees) {
			return false
		}
	}
	if this.GasProvided != that1.GasProvided {
		return false
	}
	if this.GasPenalized != that1.GasPenalized {
		return false
	}
	if this.GasRefunded != that1.GasRefunded {
		return false
	}
	return true
}
func (this *ScheduledSCRs) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ScheduledSCRs)
	if !ok {
		that2, ok := that.(ScheduledSCRs)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.RootHash, that1.RootHash) {
		return false
	}
	if len(this.Scrs) != len(that1.Scrs) {
		return false
	}
	for i := range this.Scrs {
		if !this.Scrs[i].Equal(that1.Scrs[i]) {
			return false
		}
	}
	if len(this.InvalidTxHashes) != len(that1.InvalidTxHashes) {
		return false
	}
	for i := range this.InvalidTxHashes {
		if !bytes.Equal(this.InvalidTxHashes[i], that1.InvalidTxHashes[i]) {
			return false
		}
	}
	if len(this.ScheduledMiniBlocks) != len(that1.ScheduledMiniBlocks) {
		return false
	}
	for i := range this.ScheduledMiniBlocks {
		if !this.ScheduledMiniBlocks[i].Equal(that1.ScheduledMiniBlocks[i]) {
			return false
		}
	}
	if !this.GasAndFees.Equal(that1.GasAndFees) {
		return false
	}
	return true
}
func (this *ScheduledMiniBlocks) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&scheduled.ScheduledMiniBlocks{")
	if this.MiniBlocks != nil {
		s = append(s, "MiniBlocks: "+fmt.Sprintf("%#v", this.MiniBlocks)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GasAndFees) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&scheduled.GasAndFees{")
	s = append(s, "AccumulatedFees: "+fmt.Sprintf("%#v", this.AccumulatedFees)+",\n")
	s = append(s, "DeveloperFees: "+fmt.Sprintf("%#v", this.DeveloperFees)+",\n")
	s = append(s, "GasProvided: "+fmt.Sprintf("%#v", this.GasProvided)+",\n")
	s = append(s, "GasPenalized: "+fmt.Sprintf("%#v", this.GasPenalized)+",\n")
	s = append(s, "GasRefunded: "+fmt.Sprintf("%#v", this.GasRefunded)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ScheduledSCRs) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&scheduled.ScheduledSCRs{")
	s = append(s, "RootHash: "+fmt.Sprintf("%#v", this.RootHash)+",\n")
	if this.Scrs != nil {
		s = append(s, "Scrs: "+fmt.Sprintf("%#v", this.Scrs)+",\n")
	}
	s = append(s, "InvalidTxHashes: "+fmt.Sprintf("%#v", this.InvalidTxHashes)+",\n")
	keysForScheduledMiniBlocks := make([]uint32, 0, len(this.ScheduledMiniBlocks))
	for k, _ := range this.ScheduledMiniBlocks {
		keysForScheduledMiniBlocks = append(keysForScheduledMiniBlocks, k)
	}
	github_com_gogo_protobuf_sortkeys.Uint32s(keysForScheduledMiniBlocks)
	mapStringForScheduledMiniBlocks := "map[uint32]*ScheduledMiniBlocks{"
	for _, k := range keysForScheduledMiniBlocks {
		mapStringForScheduledMiniBlocks += fmt.Sprintf("%#v: %#v,", k, this.ScheduledMiniBlocks[k])
	}
	mapStringForScheduledMiniBlocks += "}"
	if this.ScheduledMiniBlocks != nil {
		s = append(s, "ScheduledMiniBlocks: "+mapStringForScheduledMiniBlocks+",\n")
	}
	if this.GasAndFees != nil {
		s = append(s, "GasAndFees: "+fmt.Sprintf("%#v", this.GasAndFees)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringScheduled(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ScheduledMiniBlocks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduledMiniBlocks) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduledMiniBlocks) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MiniBlocks) > 0 {
		for iNdEx := len(m.MiniBlocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MiniBlocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintScheduled(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GasAndFees) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GasAndFees) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GasAndFees) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GasRefunded != 0 {
		i = encodeVarintScheduled(dAtA, i, uint64(m.GasRefunded))
		i--
		dAtA[i] = 0x28
	}
	if m.GasPenalized != 0 {
		i = encodeVarintScheduled(dAtA, i, uint64(m.GasPenalized))
		i--
		dAtA[i] = 0x20
	}
	if m.GasProvided != 0 {
		i = encodeVarintScheduled(dAtA, i, uint64(m.GasProvided))
		i--
		dAtA[i] = 0x18
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_core_data.BigIntCaster{}
		size := __caster.Size(m.DeveloperFees)
		i -= size
		if _, err := __caster.MarshalTo(m.DeveloperFees, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintScheduled(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		__caster := &github_com_ElrondNetwork_elrond_go_core_data.BigIntCaster{}
		size := __caster.Size(m.AccumulatedFees)
		i -= size
		if _, err := __caster.MarshalTo(m.AccumulatedFees, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintScheduled(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ScheduledSCRs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduledSCRs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduledSCRs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GasAndFees != nil {
		{
			size, err := m.GasAndFees.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintScheduled(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ScheduledMiniBlocks) > 0 {
		keysForScheduledMiniBlocks := make([]uint32, 0, len(m.ScheduledMiniBlocks))
		for k := range m.ScheduledMiniBlocks {
			keysForScheduledMiniBlocks = append(keysForScheduledMiniBlocks, uint32(k))
		}
		github_com_gogo_protobuf_sortkeys.Uint32s(keysForScheduledMiniBlocks)
		for iNdEx := len(keysForScheduledMiniBlocks) - 1; iNdEx >= 0; iNdEx-- {
			v := m.ScheduledMiniBlocks[uint32(keysForScheduledMiniBlocks[iNdEx])]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintScheduled(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i = encodeVarintScheduled(dAtA, i, uint64(keysForScheduledMiniBlocks[iNdEx]))
			i--
			dAtA[i] = 0x8
			i = encodeVarintScheduled(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.InvalidTxHashes) > 0 {
		for iNdEx := len(m.InvalidTxHashes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.InvalidTxHashes[iNdEx])
			copy(dAtA[i:], m.InvalidTxHashes[iNdEx])
			i = encodeVarintScheduled(dAtA, i, uint64(len(m.InvalidTxHashes[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Scrs) > 0 {
		for iNdEx := len(m.Scrs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Scrs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintScheduled(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.RootHash) > 0 {
		i -= len(m.RootHash)
		copy(dAtA[i:], m.RootHash)
		i = encodeVarintScheduled(dAtA, i, uint64(len(m.RootHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintScheduled(dAtA []byte, offset int, v uint64) int {
	offset -= sovScheduled(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ScheduledMiniBlocks) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MiniBlocks) > 0 {
		for _, e := range m.MiniBlocks {
			l = e.Size()
			n += 1 + l + sovScheduled(uint64(l))
		}
	}
	return n
}

func (m *GasAndFees) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	{
		__caster := &github_com_ElrondNetwork_elrond_go_core_data.BigIntCaster{}
		l = __caster.Size(m.AccumulatedFees)
		n += 1 + l + sovScheduled(uint64(l))
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_core_data.BigIntCaster{}
		l = __caster.Size(m.DeveloperFees)
		n += 1 + l + sovScheduled(uint64(l))
	}
	if m.GasProvided != 0 {
		n += 1 + sovScheduled(uint64(m.GasProvided))
	}
	if m.GasPenalized != 0 {
		n += 1 + sovScheduled(uint64(m.GasPenalized))
	}
	if m.GasRefunded != 0 {
		n += 1 + sovScheduled(uint64(m.GasRefunded))
	}
	return n
}

func (m *ScheduledSCRs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RootHash)
	if l > 0 {
		n += 1 + l + sovScheduled(uint64(l))
	}
	if len(m.Scrs) > 0 {
		for _, e := range m.Scrs {
			l = e.Size()
			n += 1 + l + sovScheduled(uint64(l))
		}
	}
	if len(m.InvalidTxHashes) > 0 {
		for _, b := range m.InvalidTxHashes {
			l = len(b)
			n += 1 + l + sovScheduled(uint64(l))
		}
	}
	if len(m.ScheduledMiniBlocks) > 0 {
		for k, v := range m.ScheduledMiniBlocks {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovScheduled(uint64(l))
			}
			mapEntrySize := 1 + sovScheduled(uint64(k)) + l
			n += mapEntrySize + 1 + sovScheduled(uint64(mapEntrySize))
		}
	}
	if m.GasAndFees != nil {
		l = m.GasAndFees.Size()
		n += 1 + l + sovScheduled(uint64(l))
	}
	return n
}

func sovScheduled(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozScheduled(x uint64) (n int) {
	return sovScheduled(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ScheduledMiniBlocks) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForMiniBlocks := "[]*MiniBlock{"
	for _, f := range this.MiniBlocks {
		repeatedStringForMiniBlocks += strings.Replace(fmt.Sprintf("%v", f), "MiniBlock", "block.MiniBlock", 1) + ","
	}
	repeatedStringForMiniBlocks += "}"
	s := strings.Join([]string{`&ScheduledMiniBlocks{`,
		`MiniBlocks:` + repeatedStringForMiniBlocks + `,`,
		`}`,
	}, "")
	return s
}
func (this *GasAndFees) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GasAndFees{`,
		`AccumulatedFees:` + fmt.Sprintf("%v", this.AccumulatedFees) + `,`,
		`DeveloperFees:` + fmt.Sprintf("%v", this.DeveloperFees) + `,`,
		`GasProvided:` + fmt.Sprintf("%v", this.GasProvided) + `,`,
		`GasPenalized:` + fmt.Sprintf("%v", this.GasPenalized) + `,`,
		`GasRefunded:` + fmt.Sprintf("%v", this.GasRefunded) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ScheduledSCRs) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForScrs := "[]*SmartContractResult{"
	for _, f := range this.Scrs {
		repeatedStringForScrs += strings.Replace(fmt.Sprintf("%v", f), "SmartContractResult", "smartContractResult.SmartContractResult", 1) + ","
	}
	repeatedStringForScrs += "}"
	keysForScheduledMiniBlocks := make([]uint32, 0, len(this.ScheduledMiniBlocks))
	for k, _ := range this.ScheduledMiniBlocks {
		keysForScheduledMiniBlocks = append(keysForScheduledMiniBlocks, k)
	}
	github_com_gogo_protobuf_sortkeys.Uint32s(keysForScheduledMiniBlocks)
	mapStringForScheduledMiniBlocks := "map[uint32]*ScheduledMiniBlocks{"
	for _, k := range keysForScheduledMiniBlocks {
		mapStringForScheduledMiniBlocks += fmt.Sprintf("%v: %v,", k, this.ScheduledMiniBlocks[k])
	}
	mapStringForScheduledMiniBlocks += "}"
	s := strings.Join([]string{`&ScheduledSCRs{`,
		`RootHash:` + fmt.Sprintf("%v", this.RootHash) + `,`,
		`Scrs:` + repeatedStringForScrs + `,`,
		`InvalidTxHashes:` + fmt.Sprintf("%v", this.InvalidTxHashes) + `,`,
		`ScheduledMiniBlocks:` + mapStringForScheduledMiniBlocks + `,`,
		`GasAndFees:` + strings.Replace(this.GasAndFees.String(), "GasAndFees", "GasAndFees", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringScheduled(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ScheduledMiniBlocks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScheduled
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ScheduledMiniBlocks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduledMiniBlocks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MiniBlocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduled
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthScheduled
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthScheduled
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MiniBlocks = append(m.MiniBlocks, &block.MiniBlock{})
			if err := m.MiniBlocks[len(m.MiniBlocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipScheduled(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthScheduled
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthScheduled
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GasAndFees) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScheduled
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GasAndFees: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GasAndFees: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccumulatedFees", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduled
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthScheduled
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthScheduled
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_core_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.AccumulatedFees = tmp
				}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeveloperFees", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduled
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthScheduled
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthScheduled
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_core_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.DeveloperFees = tmp
				}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasProvided", wireType)
			}
			m.GasProvided = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduled
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasProvided |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPenalized", wireType)
			}
			m.GasPenalized = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduled
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasPenalized |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasRefunded", wireType)
			}
			m.GasRefunded = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduled
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasRefunded |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipScheduled(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthScheduled
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthScheduled
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ScheduledSCRs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScheduled
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ScheduledSCRs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduledSCRs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduled
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthScheduled
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthScheduled
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RootHash = append(m.RootHash[:0], dAtA[iNdEx:postIndex]...)
			if m.RootHash == nil {
				m.RootHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scrs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduled
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthScheduled
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthScheduled
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Scrs = append(m.Scrs, &smartContractResult.SmartContractResult{})
			if err := m.Scrs[len(m.Scrs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InvalidTxHashes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduled
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthScheduled
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthScheduled
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InvalidTxHashes = append(m.InvalidTxHashes, make([]byte, postIndex-iNdEx))
			copy(m.InvalidTxHashes[len(m.InvalidTxHashes)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduledMiniBlocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduled
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthScheduled
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthScheduled
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ScheduledMiniBlocks == nil {
				m.ScheduledMiniBlocks = make(map[uint32]*ScheduledMiniBlocks)
			}
			var mapkey uint32
			var mapvalue *ScheduledMiniBlocks
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowScheduled
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowScheduled
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowScheduled
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthScheduled
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthScheduled
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &ScheduledMiniBlocks{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipScheduled(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthScheduled
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ScheduledMiniBlocks[mapkey] = mapvalue
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasAndFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduled
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthScheduled
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthScheduled
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GasAndFees == nil {
				m.GasAndFees = &GasAndFees{}
			}
			if err := m.GasAndFees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipScheduled(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthScheduled
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthScheduled
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipScheduled(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowScheduled
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowScheduled
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowScheduled
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthScheduled
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupScheduled
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthScheduled
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthScheduled        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowScheduled          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupScheduled = fmt.Errorf("proto: unexpected end of group")
)
