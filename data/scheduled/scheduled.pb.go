// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scheduled.proto

package scheduled

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_multiversx_mx_chain_core_go_data "github.com/multiversx/mx-chain-core-go/data"
	block "github.com/multiversx/mx-chain-core-go/data/block"
	smartContractResult "github.com/multiversx/mx-chain-core-go/data/smartContractResult"
	transaction "github.com/multiversx/mx-chain-core-go/data/transaction"
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

type GasAndFees struct {
	AccumulatedFees *math_big.Int `protobuf:"bytes,1,opt,name=AccumulatedFees,proto3,casttypewith=math/big.Int;github.com/multiversx/mx-chain-core-go/data.BigIntCaster" json:"AccumulatedFees,omitempty"`
	DeveloperFees   *math_big.Int `protobuf:"bytes,2,opt,name=DeveloperFees,proto3,casttypewith=math/big.Int;github.com/multiversx/mx-chain-core-go/data.BigIntCaster" json:"DeveloperFees,omitempty"`
	GasProvided     uint64        `protobuf:"varint,3,opt,name=GasProvided,proto3" json:"GasProvided,omitempty"`
	GasPenalized    uint64        `protobuf:"varint,4,opt,name=GasPenalized,proto3" json:"GasPenalized,omitempty"`
	GasRefunded     uint64        `protobuf:"varint,5,opt,name=GasRefunded,proto3" json:"GasRefunded,omitempty"`
}

func (m *GasAndFees) Reset()      { *m = GasAndFees{} }
func (*GasAndFees) ProtoMessage() {}
func (*GasAndFees) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80076f37bd30c16, []int{0}
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
	InvalidTransactions []*transaction.Transaction                 `protobuf:"bytes,3,rep,name=invalidTransactions,proto3" json:"invalidTransactions,omitempty"`
	ScheduledMiniBlocks []*block.MiniBlock                         `protobuf:"bytes,4,rep,name=scheduledMiniBlocks,proto3" json:"scheduledMiniBlocks,omitempty"`
	GasAndFees          *GasAndFees                                `protobuf:"bytes,5,opt,name=gasAndFees,proto3" json:"gasAndFees,omitempty"`
}

func (m *ScheduledSCRs) Reset()      { *m = ScheduledSCRs{} }
func (*ScheduledSCRs) ProtoMessage() {}
func (*ScheduledSCRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80076f37bd30c16, []int{1}
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

func (m *ScheduledSCRs) GetInvalidTransactions() []*transaction.Transaction {
	if m != nil {
		return m.InvalidTransactions
	}
	return nil
}

func (m *ScheduledSCRs) GetScheduledMiniBlocks() []*block.MiniBlock {
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
	proto.RegisterType((*GasAndFees)(nil), "proto.GasAndFees")
	proto.RegisterType((*ScheduledSCRs)(nil), "proto.ScheduledSCRs")
}

func init() { proto.RegisterFile("scheduled.proto", fileDescriptor_f80076f37bd30c16) }

var fileDescriptor_f80076f37bd30c16 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x31, 0x8f, 0xd3, 0x30,
	0x18, 0x8d, 0xaf, 0x3d, 0x04, 0xee, 0x9d, 0x0e, 0x7c, 0x4b, 0xd5, 0xc1, 0x54, 0x9d, 0xba, 0x34,
	0x11, 0xc7, 0x88, 0x18, 0xae, 0x3d, 0x38, 0x3a, 0x20, 0x21, 0x97, 0x05, 0x36, 0xc7, 0xf1, 0x25,
	0x16, 0x89, 0x7d, 0xb2, 0x9d, 0xea, 0xc4, 0xc4, 0x4f, 0x60, 0xe4, 0x27, 0x20, 0x7e, 0x09, 0x63,
	0xc7, 0x8a, 0x05, 0x9a, 0x2e, 0x8c, 0xf7, 0x13, 0x50, 0x9d, 0x5c, 0x2e, 0x07, 0x5d, 0x2a, 0xb1,
	0x24, 0xfe, 0x9e, 0xdf, 0x7b, 0x9f, 0xe3, 0xef, 0x05, 0x1e, 0x19, 0x96, 0xf0, 0x28, 0x4f, 0x79,
	0xe4, 0x5f, 0x6a, 0x65, 0x15, 0xda, 0x77, 0xaf, 0xde, 0x28, 0x16, 0x36, 0xc9, 0x43, 0x9f, 0xa9,
	0x2c, 0x88, 0x55, 0xac, 0x02, 0x07, 0x87, 0xf9, 0x85, 0xab, 0x5c, 0xe1, 0x56, 0xa5, 0xaa, 0xf7,
	0xae, 0x41, 0xcf, 0xf2, 0xd4, 0x8a, 0x39, 0xd7, 0xe6, 0x2a, 0xc8, 0xae, 0x46, 0x2c, 0xa1, 0x42,
	0x8e, 0x98, 0xd2, 0x7c, 0x14, 0xab, 0x20, 0xa2, 0x96, 0x06, 0x26, 0xa3, 0xda, 0x4e, 0x94, 0xb4,
	0x9a, 0x32, 0x4b, 0xb8, 0xc9, 0x53, 0xbb, 0x0d, 0xab, 0xac, 0x9f, 0xef, 0x62, 0x1d, 0xa6, 0x8a,
	0x7d, 0x28, 0x9f, 0x95, 0x7c, 0xba, 0x8b, 0xdc, 0x6a, 0x2a, 0x0d, 0x65, 0x56, 0x28, 0xd9, 0x5c,
	0x97, 0x56, 0x83, 0x1f, 0x7b, 0x10, 0x9e, 0x53, 0x73, 0x2a, 0xa3, 0x97, 0x9c, 0x1b, 0x64, 0xe0,
	0xd1, 0x29, 0x63, 0x79, 0x96, 0xa7, 0xd4, 0x72, 0x07, 0x75, 0x41, 0x1f, 0x0c, 0x0f, 0xc6, 0xd3,
	0x6f, 0x3f, 0x1f, 0xbf, 0xc8, 0xa8, 0x4d, 0x82, 0x50, 0xc4, 0xfe, 0x54, 0xda, 0x67, 0x3b, 0x9c,
	0xc1, 0x1f, 0x8b, 0x78, 0x2a, 0xed, 0x84, 0x1a, 0xcb, 0x35, 0xf9, 0xbb, 0x03, 0x52, 0xf0, 0xf0,
	0x8c, 0xcf, 0x79, 0xaa, 0x2e, 0xb9, 0x76, 0x2d, 0xf7, 0xfe, 0x77, 0xcb, 0xbb, 0xfe, 0xa8, 0x0f,
	0x3b, 0xe7, 0xd4, 0xbc, 0xd1, 0x6a, 0x2e, 0x22, 0x1e, 0x75, 0x5b, 0x7d, 0x30, 0x6c, 0x93, 0x26,
	0x84, 0x06, 0xf0, 0x60, 0x53, 0x72, 0x49, 0x53, 0xf1, 0x91, 0x47, 0xdd, 0xb6, 0xa3, 0xdc, 0xc1,
	0x2a, 0x17, 0xc2, 0x2f, 0x72, 0xb9, 0x71, 0xd9, 0xaf, 0x5d, 0x6e, 0xa0, 0xc1, 0x97, 0x3d, 0x78,
	0x38, 0xbb, 0xc9, 0xe2, 0x6c, 0x42, 0x0c, 0xea, 0xc1, 0xfb, 0x5a, 0x29, 0xfb, 0x8a, 0x9a, 0xa4,
	0xbc, 0x58, 0x52, 0xd7, 0xc8, 0x87, 0x6d, 0xc3, 0xf4, 0xe6, 0xeb, 0x5b, 0xc3, 0xce, 0x49, 0xaf,
	0x1c, 0x90, 0x3f, 0xfb, 0x37, 0x44, 0xc4, 0xf1, 0xd0, 0x19, 0x3c, 0x16, 0x72, 0x4e, 0x53, 0x11,
	0xbd, 0xbd, 0x1d, 0xab, 0xe9, 0xb6, 0x9c, 0x1c, 0x55, 0xf2, 0xc6, 0x16, 0xd9, 0x46, 0x47, 0x63,
	0x78, 0x5c, 0xff, 0x2e, 0xaf, 0x85, 0x14, 0xe3, 0x4d, 0xce, 0x4c, 0xb7, 0xed, 0x5c, 0x1e, 0x56,
	0x2e, 0xf5, 0x06, 0xd9, 0x46, 0x46, 0x4f, 0x20, 0x8c, 0xeb, 0x0c, 0xb9, 0x8b, 0xe8, 0x9c, 0x3c,
	0xaa, 0xa4, 0xb7, 0xe1, 0x22, 0x0d, 0xd2, 0x78, 0xb2, 0x58, 0x61, 0x6f, 0xb9, 0xc2, 0xde, 0xf5,
	0x0a, 0x83, 0x4f, 0x05, 0x06, 0x5f, 0x0b, 0x0c, 0xbe, 0x17, 0x18, 0x2c, 0x0a, 0x0c, 0x96, 0x05,
	0x06, 0xbf, 0x0a, 0x0c, 0x7e, 0x17, 0xd8, 0xbb, 0x2e, 0x30, 0xf8, 0xbc, 0xc6, 0xde, 0x62, 0x8d,
	0xbd, 0xe5, 0x1a, 0x7b, 0xef, 0x1f, 0xd4, 0x27, 0x08, 0xef, 0xb9, 0x16, 0x4f, 0xff, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x81, 0x81, 0x3b, 0xd5, 0xf1, 0x03, 0x00, 0x00,
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
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		if !__caster.Equal(this.AccumulatedFees, that1.AccumulatedFees) {
			return false
		}
	}
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
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
	if len(this.InvalidTransactions) != len(that1.InvalidTransactions) {
		return false
	}
	for i := range this.InvalidTransactions {
		if !this.InvalidTransactions[i].Equal(that1.InvalidTransactions[i]) {
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
	if this.InvalidTransactions != nil {
		s = append(s, "InvalidTransactions: "+fmt.Sprintf("%#v", this.InvalidTransactions)+",\n")
	}
	if this.ScheduledMiniBlocks != nil {
		s = append(s, "ScheduledMiniBlocks: "+fmt.Sprintf("%#v", this.ScheduledMiniBlocks)+",\n")
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
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
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
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
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
		for iNdEx := len(m.ScheduledMiniBlocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ScheduledMiniBlocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintScheduled(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.InvalidTransactions) > 0 {
		for iNdEx := len(m.InvalidTransactions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InvalidTransactions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintScheduled(dAtA, i, uint64(size))
			}
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
func (m *GasAndFees) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		l = __caster.Size(m.AccumulatedFees)
		n += 1 + l + sovScheduled(uint64(l))
	}
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
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
	if len(m.InvalidTransactions) > 0 {
		for _, e := range m.InvalidTransactions {
			l = e.Size()
			n += 1 + l + sovScheduled(uint64(l))
		}
	}
	if len(m.ScheduledMiniBlocks) > 0 {
		for _, e := range m.ScheduledMiniBlocks {
			l = e.Size()
			n += 1 + l + sovScheduled(uint64(l))
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
	repeatedStringForInvalidTransactions := "[]*Transaction{"
	for _, f := range this.InvalidTransactions {
		repeatedStringForInvalidTransactions += strings.Replace(fmt.Sprintf("%v", f), "Transaction", "transaction.Transaction", 1) + ","
	}
	repeatedStringForInvalidTransactions += "}"
	repeatedStringForScheduledMiniBlocks := "[]*MiniBlock{"
	for _, f := range this.ScheduledMiniBlocks {
		repeatedStringForScheduledMiniBlocks += strings.Replace(fmt.Sprintf("%v", f), "MiniBlock", "block.MiniBlock", 1) + ","
	}
	repeatedStringForScheduledMiniBlocks += "}"
	s := strings.Join([]string{`&ScheduledSCRs{`,
		`RootHash:` + fmt.Sprintf("%v", this.RootHash) + `,`,
		`Scrs:` + repeatedStringForScrs + `,`,
		`InvalidTransactions:` + repeatedStringForInvalidTransactions + `,`,
		`ScheduledMiniBlocks:` + repeatedStringForScheduledMiniBlocks + `,`,
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
				__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
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
				__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
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
				return fmt.Errorf("proto: wrong wireType = %d for field InvalidTransactions", wireType)
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
			m.InvalidTransactions = append(m.InvalidTransactions, &transaction.Transaction{})
			if err := m.InvalidTransactions[len(m.InvalidTransactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
			m.ScheduledMiniBlocks = append(m.ScheduledMiniBlocks, &block.MiniBlock{})
			if err := m.ScheduledMiniBlocks[len(m.ScheduledMiniBlocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
