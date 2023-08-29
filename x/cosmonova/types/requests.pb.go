// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmonova/cosmonova/requests.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type Requests struct {
	QueryData     string     `protobuf:"bytes,1,opt,name=queryData,proto3" json:"queryData,omitempty"`
	ValueList     []*Values  `protobuf:"bytes,2,rep,name=valueList,proto3" json:"valueList,omitempty"`
	Tip           types.Coin `protobuf:"bytes,3,opt,name=tip,proto3" json:"tip"`
	TimeofRequest int64      `protobuf:"varint,4,opt,name=timeofRequest,proto3" json:"timeofRequest,omitempty"`
	Status        string     `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *Requests) Reset()         { *m = Requests{} }
func (m *Requests) String() string { return proto.CompactTextString(m) }
func (*Requests) ProtoMessage()    {}
func (*Requests) Descriptor() ([]byte, []int) {
	return fileDescriptor_d01894a1febd30ab, []int{0}
}
func (m *Requests) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Requests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Requests.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Requests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Requests.Merge(m, src)
}
func (m *Requests) XXX_Size() int {
	return m.Size()
}
func (m *Requests) XXX_DiscardUnknown() {
	xxx_messageInfo_Requests.DiscardUnknown(m)
}

var xxx_messageInfo_Requests proto.InternalMessageInfo

func (m *Requests) GetQueryData() string {
	if m != nil {
		return m.QueryData
	}
	return ""
}

func (m *Requests) GetValueList() []*Values {
	if m != nil {
		return m.ValueList
	}
	return nil
}

func (m *Requests) GetTip() types.Coin {
	if m != nil {
		return m.Tip
	}
	return types.Coin{}
}

func (m *Requests) GetTimeofRequest() int64 {
	if m != nil {
		return m.TimeofRequest
	}
	return 0
}

func (m *Requests) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Requests)(nil), "cosmonova.cosmonova.Requests")
}

func init() {
	proto.RegisterFile("cosmonova/cosmonova/requests.proto", fileDescriptor_d01894a1febd30ab)
}

var fileDescriptor_d01894a1febd30ab = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x18, 0x84, 0x63, 0x52, 0x2a, 0xea, 0x8a, 0xc5, 0x20, 0x14, 0x5a, 0x64, 0xa2, 0x8a, 0x21, 0x93,
	0xa3, 0x14, 0x31, 0xb0, 0x16, 0x46, 0xa6, 0x0c, 0x0c, 0x6c, 0x4e, 0x65, 0xaa, 0x48, 0x34, 0x4e,
	0xf3, 0xff, 0x89, 0xe8, 0x5b, 0xf0, 0x58, 0x1d, 0x3b, 0x30, 0x30, 0x21, 0x94, 0xbc, 0x08, 0x8a,
	0x6b, 0x35, 0x20, 0x75, 0x3b, 0xff, 0x77, 0xd6, 0x7d, 0x3a, 0x3a, 0x99, 0x6b, 0x58, 0xea, 0x4c,
	0x57, 0x32, 0xec, 0x54, 0xa1, 0x56, 0xa5, 0x02, 0x04, 0x91, 0x17, 0x1a, 0x35, 0x3b, 0xdb, 0x3b,
	0x62, 0xaf, 0x46, 0xe7, 0x0b, 0xbd, 0xd0, 0xc6, 0x0f, 0x5b, 0xb5, 0x8b, 0x8e, 0xb8, 0x09, 0x40,
	0x98, 0x48, 0x50, 0x61, 0x15, 0x25, 0x0a, 0x65, 0x14, 0xce, 0x75, 0x9a, 0x59, 0xdf, 0x3f, 0x54,
	0x57, 0xc9, 0xb7, 0x52, 0xd9, 0xb2, 0xc9, 0x27, 0xa1, 0x27, 0xb1, 0xed, 0x67, 0x57, 0x74, 0xb0,
	0x2a, 0x55, 0xb1, 0x7e, 0x94, 0x28, 0x3d, 0xe2, 0x93, 0x60, 0x10, 0x77, 0x07, 0x76, 0x4f, 0x07,
	0xe6, 0xeb, 0x53, 0x0a, 0xe8, 0x1d, 0xf9, 0x6e, 0x30, 0x9c, 0x8e, 0xc5, 0x01, 0x56, 0xf1, 0x6c,
	0x0a, 0xe2, 0x2e, 0xcd, 0x22, 0xea, 0x62, 0x9a, 0x7b, 0xae, 0x4f, 0x82, 0xe1, 0xf4, 0x72, 0x17,
	0x05, 0xd1, 0x52, 0x0b, 0x4b, 0x2d, 0x1e, 0x74, 0x9a, 0xcd, 0x7a, 0x9b, 0xef, 0x6b, 0x27, 0x6e,
	0xb3, 0xec, 0x86, 0x9e, 0x62, 0xba, 0x54, 0xfa, 0xd5, 0xd2, 0x79, 0x3d, 0x9f, 0x04, 0x6e, 0xfc,
	0xff, 0xc8, 0x2e, 0x68, 0x1f, 0x50, 0x62, 0x09, 0xde, 0xb1, 0xc1, 0xb5, 0xaf, 0xd9, 0xdd, 0xa6,
	0xe6, 0x64, 0x5b, 0x73, 0xf2, 0x53, 0x73, 0xf2, 0xd1, 0x70, 0x67, 0xdb, 0x70, 0xe7, 0xab, 0xe1,
	0xce, 0xcb, 0xb8, 0x1b, 0xe2, 0xfd, 0xcf, 0x28, 0xb8, 0xce, 0x15, 0x24, 0x7d, 0x33, 0xca, 0xed,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xf1, 0xbe, 0x32, 0xa7, 0x01, 0x00, 0x00,
}

func (m *Requests) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Requests) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Requests) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintRequests(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TimeofRequest != 0 {
		i = encodeVarintRequests(dAtA, i, uint64(m.TimeofRequest))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.Tip.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintRequests(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ValueList) > 0 {
		for iNdEx := len(m.ValueList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValueList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRequests(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.QueryData) > 0 {
		i -= len(m.QueryData)
		copy(dAtA[i:], m.QueryData)
		i = encodeVarintRequests(dAtA, i, uint64(len(m.QueryData)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRequests(dAtA []byte, offset int, v uint64) int {
	offset -= sovRequests(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Requests) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.QueryData)
	if l > 0 {
		n += 1 + l + sovRequests(uint64(l))
	}
	if len(m.ValueList) > 0 {
		for _, e := range m.ValueList {
			l = e.Size()
			n += 1 + l + sovRequests(uint64(l))
		}
	}
	l = m.Tip.Size()
	n += 1 + l + sovRequests(uint64(l))
	if m.TimeofRequest != 0 {
		n += 1 + sovRequests(uint64(m.TimeofRequest))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovRequests(uint64(l))
	}
	return n
}

func sovRequests(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequests(x uint64) (n int) {
	return sovRequests(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Requests) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequests
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
			return fmt.Errorf("proto: Requests: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Requests: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequests
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
				return ErrInvalidLengthRequests
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValueList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequests
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
				return ErrInvalidLengthRequests
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValueList = append(m.ValueList, &Values{})
			if err := m.ValueList[len(m.ValueList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tip", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequests
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
				return ErrInvalidLengthRequests
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tip.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeofRequest", wireType)
			}
			m.TimeofRequest = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequests
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeofRequest |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequests
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
				return ErrInvalidLengthRequests
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequests(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequests
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
func skipRequests(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequests
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
					return 0, ErrIntOverflowRequests
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
					return 0, ErrIntOverflowRequests
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
				return 0, ErrInvalidLengthRequests
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRequests
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRequests
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRequests        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequests          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRequests = fmt.Errorf("proto: unexpected end of group")
)