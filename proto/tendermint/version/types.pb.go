// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/version/types.proto

package version

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// Consensus captures the consensus rules for processing a block in the
// blockchain, including all blockchain data structures and the rules of the
// application's state transition machine.
type Consensus struct {
	Block uint64 `protobuf:"varint,1,opt,name=block,proto3" json:"block,omitempty"`
	App   uint64 `protobuf:"varint,2,opt,name=app,proto3" json:"app,omitempty"`
}

func (m *Consensus) Reset()         { *m = Consensus{} }
func (m *Consensus) String() string { return proto.CompactTextString(m) }
func (*Consensus) ProtoMessage()    {}
func (*Consensus) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9b42966edc5edad, []int{0}
}
func (m *Consensus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Consensus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Consensus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Consensus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consensus.Merge(m, src)
}
func (m *Consensus) XXX_Size() int {
	return m.Size()
}
func (m *Consensus) XXX_DiscardUnknown() {
	xxx_messageInfo_Consensus.DiscardUnknown(m)
}

var xxx_messageInfo_Consensus proto.InternalMessageInfo

func (m *Consensus) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *Consensus) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

func init() {
	proto.RegisterType((*Consensus)(nil), "tendermint.version.Consensus")
}

func init() { proto.RegisterFile("tendermint/version/types.proto", fileDescriptor_f9b42966edc5edad) }

var fileDescriptor_f9b42966edc5edad = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2f, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0xd3, 0x2f,
	0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x42, 0xc8, 0xeb, 0x41,
	0xe5, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xd2, 0xfa, 0x20, 0x16, 0x44, 0xa5, 0x92, 0x25,
	0x17, 0xa7, 0x73, 0x7e, 0x5e, 0x71, 0x6a, 0x5e, 0x71, 0x69, 0xb1, 0x90, 0x08, 0x17, 0x6b, 0x52,
	0x4e, 0x7e, 0x72, 0xb6, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x84, 0x23, 0x24, 0xc0, 0xc5,
	0x9c, 0x58, 0x50, 0x20, 0xc1, 0x04, 0x16, 0x03, 0x31, 0xad, 0x58, 0x5e, 0x2c, 0x90, 0x67, 0x74,
	0x0a, 0x3e, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c,
	0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xcb, 0xf4, 0xcc, 0x92,
	0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x24, 0x97, 0x22, 0x31, 0x21, 0xee, 0xc0, 0xf4,
	0x45, 0x12, 0x1b, 0x58, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xf0, 0x65, 0xd2, 0xe2,
	0x00, 0x00, 0x00,
}

func (this *Consensus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Consensus)
	if !ok {
		that2, ok := that.(Consensus)
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
	if this.Block != that1.Block {
		return false
	}
	if this.App != that1.App {
		return false
	}
	return true
}
func (m *Consensus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Consensus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Consensus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.App != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.App))
		i--
		dAtA[i] = 0x10
	}
	if m.Block != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Consensus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Block != 0 {
		n += 1 + sovTypes(uint64(m.Block))
	}
	if m.App != 0 {
		n += 1 + sovTypes(uint64(m.App))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Consensus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Consensus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Consensus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field App", wireType)
			}
			m.App = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.App |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
