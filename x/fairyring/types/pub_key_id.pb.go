// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fairyring/fairyring/pub_key_id.proto

package types

import (
	fmt "fmt"
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

type PubKeyID struct {
	Height    uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	PublicKey string `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	IbeID     string `protobuf:"bytes,3,opt,name=ibeID,proto3" json:"ibeID,omitempty"`
	Creator   string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *PubKeyID) Reset()         { *m = PubKeyID{} }
func (m *PubKeyID) String() string { return proto.CompactTextString(m) }
func (*PubKeyID) ProtoMessage()    {}
func (*PubKeyID) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d325226236bbdb8, []int{0}
}
func (m *PubKeyID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PubKeyID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PubKeyID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PubKeyID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubKeyID.Merge(m, src)
}
func (m *PubKeyID) XXX_Size() int {
	return m.Size()
}
func (m *PubKeyID) XXX_DiscardUnknown() {
	xxx_messageInfo_PubKeyID.DiscardUnknown(m)
}

var xxx_messageInfo_PubKeyID proto.InternalMessageInfo

func (m *PubKeyID) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *PubKeyID) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *PubKeyID) GetIbeID() string {
	if m != nil {
		return m.IbeID
	}
	return ""
}

func (m *PubKeyID) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*PubKeyID)(nil), "fairyring.fairyring.PubKeyID")
}

func init() {
	proto.RegisterFile("fairyring/fairyring/pub_key_id.proto", fileDescriptor_0d325226236bbdb8)
}

var fileDescriptor_0d325226236bbdb8 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0x4b, 0xcc, 0x2c,
	0xaa, 0x2c, 0xca, 0xcc, 0x4b, 0xd7, 0x47, 0xb0, 0x0a, 0x4a, 0x93, 0xe2, 0xb3, 0x53, 0x2b, 0xe3,
	0x33, 0x53, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84, 0xe1, 0x72, 0x7a, 0x70, 0x96, 0x52,
	0x01, 0x17, 0x47, 0x40, 0x69, 0x92, 0x77, 0x6a, 0xa5, 0xa7, 0x8b, 0x90, 0x18, 0x17, 0x5b, 0x46,
	0x6a, 0x66, 0x7a, 0x46, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x94, 0x27, 0x24, 0xc3,
	0xc5, 0x59, 0x50, 0x9a, 0x94, 0x93, 0x99, 0xec, 0x9d, 0x5a, 0x29, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1,
	0x19, 0x84, 0x10, 0x10, 0x12, 0xe1, 0x62, 0xcd, 0x4c, 0x4a, 0xf5, 0x74, 0x91, 0x60, 0x06, 0xcb,
	0x40, 0x38, 0x42, 0x12, 0x5c, 0xec, 0xc9, 0x45, 0xa9, 0x89, 0x25, 0xf9, 0x45, 0x12, 0x2c, 0x60,
	0x71, 0x18, 0xd7, 0xc9, 0xf4, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92,
	0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xa4,
	0x11, 0x8e, 0xaf, 0x40, 0xf2, 0x48, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x13, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0x13, 0x89, 0xee, 0xec, 0x00, 0x00, 0x00,
}

func (m *PubKeyID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PubKeyID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PubKeyID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPubKeyId(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.IbeID) > 0 {
		i -= len(m.IbeID)
		copy(dAtA[i:], m.IbeID)
		i = encodeVarintPubKeyId(dAtA, i, uint64(len(m.IbeID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintPubKeyId(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintPubKeyId(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPubKeyId(dAtA []byte, offset int, v uint64) int {
	offset -= sovPubKeyId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PubKeyID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovPubKeyId(uint64(m.Height))
	}
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovPubKeyId(uint64(l))
	}
	l = len(m.IbeID)
	if l > 0 {
		n += 1 + l + sovPubKeyId(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPubKeyId(uint64(l))
	}
	return n
}

func sovPubKeyId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPubKeyId(x uint64) (n int) {
	return sovPubKeyId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PubKeyID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPubKeyId
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
			return fmt.Errorf("proto: PubKeyID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PubKeyID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPubKeyId
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPubKeyId
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPubKeyId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPubKeyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbeID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPubKeyId
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPubKeyId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPubKeyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbeID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPubKeyId
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPubKeyId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPubKeyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPubKeyId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPubKeyId
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
func skipPubKeyId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPubKeyId
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
					return 0, ErrIntOverflowPubKeyId
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
					return 0, ErrIntOverflowPubKeyId
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
				return 0, ErrInvalidLengthPubKeyId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPubKeyId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPubKeyId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPubKeyId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPubKeyId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPubKeyId = fmt.Errorf("proto: unexpected end of group")
)