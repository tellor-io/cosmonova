// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmonova/cosmonova/data_request.proto

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

type DataRequest struct {
	QueryData    string     `protobuf:"bytes,1,opt,name=queryData,proto3" json:"queryData,omitempty"`
	Value        string     `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	LastReported uint64     `protobuf:"varint,3,opt,name=lastReported,proto3" json:"lastReported,omitempty"`
	TipAmount    types.Coin `protobuf:"bytes,4,opt,name=tipAmount,proto3" json:"tipAmount"`
	Status       string     `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *DataRequest) Reset()         { *m = DataRequest{} }
func (m *DataRequest) String() string { return proto.CompactTextString(m) }
func (*DataRequest) ProtoMessage()    {}
func (*DataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_faf99196d94e15f1, []int{0}
}
func (m *DataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRequest.Merge(m, src)
}
func (m *DataRequest) XXX_Size() int {
	return m.Size()
}
func (m *DataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataRequest proto.InternalMessageInfo

func (m *DataRequest) GetQueryData() string {
	if m != nil {
		return m.QueryData
	}
	return ""
}

func (m *DataRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *DataRequest) GetLastReported() uint64 {
	if m != nil {
		return m.LastReported
	}
	return 0
}

func (m *DataRequest) GetTipAmount() types.Coin {
	if m != nil {
		return m.TipAmount
	}
	return types.Coin{}
}

func (m *DataRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*DataRequest)(nil), "cosmonova.cosmonova.DataRequest")
}

func init() {
	proto.RegisterFile("cosmonova/cosmonova/data_request.proto", fileDescriptor_faf99196d94e15f1)
}

var fileDescriptor_faf99196d94e15f1 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0x68, 0x2b, 0xc5, 0x65, 0x32, 0x15, 0x0a, 0x05, 0x99, 0xa8, 0x03, 0xca, 0x94,
	0xa8, 0x20, 0x46, 0x06, 0x0a, 0x4f, 0x90, 0x91, 0x05, 0x5d, 0x5a, 0xab, 0x8a, 0xd4, 0xe6, 0xd2,
	0xf8, 0x12, 0xd1, 0xb7, 0xe0, 0x81, 0x78, 0x80, 0x8e, 0x1d, 0x99, 0x10, 0x4a, 0x5e, 0x04, 0xc5,
	0xae, 0x1a, 0xd8, 0xfe, 0xfb, 0xbf, 0xb3, 0xee, 0xf7, 0xcf, 0x6f, 0xe7, 0xa8, 0xd7, 0x98, 0x61,
	0x05, 0x51, 0xa7, 0x16, 0x40, 0xf0, 0x56, 0xa8, 0x4d, 0xa9, 0x34, 0x85, 0x79, 0x81, 0x84, 0xe2,
	0xfc, 0x48, 0xc3, 0xa3, 0x1a, 0x8f, 0x96, 0xb8, 0x44, 0xc3, 0xa3, 0x56, 0xd9, 0xd5, 0xb1, 0x34,
	0x0b, 0x3a, 0x4a, 0x40, 0xab, 0xa8, 0x9a, 0x26, 0x8a, 0x60, 0x1a, 0xcd, 0x31, 0xcd, 0x2c, 0x9f,
	0x7c, 0x32, 0x3e, 0x7c, 0x01, 0x82, 0xd8, 0x1e, 0x10, 0xd7, 0xdc, 0xdd, 0x94, 0xaa, 0xd8, 0xb6,
	0x9e, 0xc7, 0x7c, 0x16, 0xb8, 0x71, 0x67, 0x88, 0x11, 0xef, 0x57, 0xb0, 0x2a, 0x95, 0x77, 0x62,
	0x88, 0x1d, 0xc4, 0x84, 0x9f, 0xad, 0x40, 0x53, 0xac, 0x72, 0x2c, 0x48, 0x2d, 0xbc, 0x53, 0x9f,
	0x05, 0xbd, 0xf8, 0x9f, 0x27, 0x1e, 0xb9, 0x4b, 0x69, 0xfe, 0xb4, 0xc6, 0x32, 0x23, 0xaf, 0xe7,
	0xb3, 0x60, 0x78, 0x77, 0x69, 0xc3, 0xeb, 0xb0, 0xcd, 0x16, 0x1e, 0xb2, 0x85, 0xcf, 0x98, 0x66,
	0xb3, 0xde, 0xee, 0xfb, 0xc6, 0x89, 0xbb, 0x17, 0xe2, 0x82, 0x0f, 0x34, 0x01, 0x95, 0xda, 0xeb,
	0x9b, 0xcb, 0x87, 0x69, 0xf6, 0xb0, 0xab, 0x25, 0xdb, 0xd7, 0x92, 0xfd, 0xd4, 0x92, 0x7d, 0x34,
	0xd2, 0xd9, 0x37, 0xd2, 0xf9, 0x6a, 0xa4, 0xf3, 0x7a, 0xd5, 0x35, 0xf8, 0xfe, 0xa7, 0x4d, 0xda,
	0xe6, 0x4a, 0x27, 0x03, 0xf3, 0xf9, 0xfb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0xa1, 0xe6,
	0x1c, 0x71, 0x01, 0x00, 0x00,
}

func (m *DataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintDataRequest(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size, err := m.TipAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDataRequest(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.LastReported != 0 {
		i = encodeVarintDataRequest(dAtA, i, uint64(m.LastReported))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintDataRequest(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.QueryData) > 0 {
		i -= len(m.QueryData)
		copy(dAtA[i:], m.QueryData)
		i = encodeVarintDataRequest(dAtA, i, uint64(len(m.QueryData)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDataRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovDataRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.QueryData)
	if l > 0 {
		n += 1 + l + sovDataRequest(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovDataRequest(uint64(l))
	}
	if m.LastReported != 0 {
		n += 1 + sovDataRequest(uint64(m.LastReported))
	}
	l = m.TipAmount.Size()
	n += 1 + l + sovDataRequest(uint64(l))
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovDataRequest(uint64(l))
	}
	return n
}

func sovDataRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDataRequest(x uint64) (n int) {
	return sovDataRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataRequest
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
			return fmt.Errorf("proto: DataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataRequest
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
				return ErrInvalidLengthDataRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataRequest
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
				return ErrInvalidLengthDataRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastReported", wireType)
			}
			m.LastReported = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastReported |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TipAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataRequest
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
				return ErrInvalidLengthDataRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDataRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TipAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataRequest
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
				return ErrInvalidLengthDataRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDataRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDataRequest
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
func skipDataRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDataRequest
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
					return 0, ErrIntOverflowDataRequest
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
					return 0, ErrIntOverflowDataRequest
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
				return 0, ErrInvalidLengthDataRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDataRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDataRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDataRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDataRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDataRequest = fmt.Errorf("proto: unexpected end of group")
)