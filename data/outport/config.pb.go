// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

package outport

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
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

type OutportConfig struct {
	IsInImportDBMode bool `protobuf:"varint,1,opt,name=IsInImportDBMode,proto3" json:"isInImportDBMode"`
}

func (m *OutportConfig) Reset()      { *m = OutportConfig{} }
func (*OutportConfig) ProtoMessage() {}
func (*OutportConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}
func (m *OutportConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutportConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *OutportConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutportConfig.Merge(m, src)
}
func (m *OutportConfig) XXX_Size() int {
	return m.Size()
}
func (m *OutportConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OutportConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OutportConfig proto.InternalMessageInfo

func (m *OutportConfig) GetIsInImportDBMode() bool {
	if m != nil {
		return m.IsInImportDBMode
	}
	return false
}

func init() {
	proto.RegisterType((*OutportConfig)(nil), "proto.OutportConfig")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_3eaf2c85e69e9ea4) }

var fileDescriptor_3eaf2c85e69e9ea4 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0xba, 0xe9, 0x99,
	0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9, 0xfa, 0x60, 0xe1,
	0xa4, 0xd2, 0x34, 0x30, 0x0f, 0xcc, 0x01, 0xb3, 0x20, 0xba, 0x94, 0x02, 0xb9, 0x78, 0xfd, 0x4b,
	0x4b, 0x0a, 0xf2, 0x8b, 0x4a, 0x9c, 0xc1, 0x86, 0x09, 0x39, 0x70, 0x09, 0x78, 0x16, 0x7b, 0xe6,
	0x79, 0xe6, 0x82, 0xc4, 0x5c, 0x9c, 0x7c, 0xf3, 0x53, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x9c, 0x44, 0x5e, 0xdd, 0x93, 0x17, 0xc8, 0x44, 0x93, 0x0b, 0xc2, 0x50, 0xed, 0x54, 0x7a, 0xe1,
	0xa1, 0x1c, 0xc3, 0x8d, 0x87, 0x72, 0x0c, 0x1f, 0x1e, 0xca, 0x31, 0x36, 0x3c, 0x92, 0x63, 0x5c,
	0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x6f, 0x3c, 0x92, 0x63,
	0x7c, 0xf0, 0x48, 0x8e, 0xf1, 0xc5, 0x23, 0x39, 0x86, 0x0f, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96,
	0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x6b, 0x24, 0xb7, 0xe7, 0x96,
	0xe6, 0x94, 0x64, 0x96, 0xa5, 0x16, 0x15, 0x57, 0xe8, 0xe7, 0x56, 0xe8, 0x26, 0x67, 0x24, 0x66,
	0xe6, 0xe9, 0x26, 0xe7, 0x17, 0xa5, 0xea, 0xa6, 0xe7, 0xeb, 0xa7, 0x24, 0x96, 0x24, 0xea, 0xe7,
	0x43, 0x9c, 0x6d, 0x0d, 0xa5, 0x93, 0xd8, 0xc0, 0x1e, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xc2, 0x23, 0x6e, 0xa4, 0x16, 0x01, 0x00, 0x00,
}

func (this *OutportConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OutportConfig)
	if !ok {
		that2, ok := that.(OutportConfig)
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
	if this.IsInImportDBMode != that1.IsInImportDBMode {
		return false
	}
	return true
}
func (this *OutportConfig) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&outport.OutportConfig{")
	s = append(s, "IsInImportDBMode: "+fmt.Sprintf("%#v", this.IsInImportDBMode)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringConfig(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *OutportConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutportConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutportConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsInImportDBMode {
		i--
		if m.IsInImportDBMode {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OutportConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsInImportDBMode {
		n += 2
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *OutportConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&OutportConfig{`,
		`IsInImportDBMode:` + fmt.Sprintf("%v", this.IsInImportDBMode) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *OutportConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: OutportConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutportConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsInImportDBMode", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsInImportDBMode = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)
