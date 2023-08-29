// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmonova/client/mutistore_tree.proto

package types

import (
	fmt "fmt"
	github_com_cometbft_cometbft_libs_bytes "github.com/cometbft/cometbft/libs/bytes"
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

// ___________________________[AppHash]______________
// /                                                  \
// _________________[I19]_________________                           ____[I20*]____
// /                                        \	                      /              \
// _______[I15]______                          _______[I16*]______          [GHIJ]           [KLMN]
// /                  \                        /                   \
// __[I8*]__             __[I9]__              __[I10]__             __[I11]__
// /         \          /         \            /         \           /         \
// [I0]       [I1]     [I2]        [I3*]       [I4]      [I5]        [I6]       [I7]
// /   \      /   \     /   \      /    \      /    \    /    \     /    \     /    \
// [0]   [1]  [2]   [3] [4*]  [5*]  [6]   [7]  [8]    [9] [A]    [B] [C]     [D] [E]     [F]
// Left[4], Right[I3], Left[I8], Right[I16], Right[I20]
// [0] - acc (auth) [1] - authz           [2] - bank     [3] - capability [4] - consensus [5] - cosmonova
// [6] - crisis     [7] - distr           [8] - evidence [9] - feegrant   [A] - gov       [B] - group
// [C] - ibc        [D] - icacontroller   [E] - icahost  [F] - mint       [G] - params    [H] - slashing
// [I] - staking    [J] - transfer        [K] - upgrade  [L] - vesting
type MutiStoreTreeFields struct {
	CosmonovaIavlStateHash            github_com_cometbft_cometbft_libs_bytes.HexBytes `protobuf:"bytes,1,opt,name=cosmonova_iavl_state_hash,json=cosmonovaIavlStateHash,proto3,casttype=github.com/cometbft/cometbft/libs/bytes.HexBytes" json:"cosmonova_iavl_state_hash,omitempty"`
	ConsensusStoreMerkleHash          github_com_cometbft_cometbft_libs_bytes.HexBytes `protobuf:"bytes,2,opt,name=consensus_store_merkle_hash,json=consensusStoreMerkleHash,proto3,casttype=github.com/cometbft/cometbft/libs/bytes.HexBytes" json:"consensus_store_merkle_hash,omitempty"`
	CrisisToDistrStoresMerkleHash     github_com_cometbft_cometbft_libs_bytes.HexBytes `protobuf:"bytes,3,opt,name=crisis_to_distr_stores_merkle_hash,json=crisisToDistrStoresMerkleHash,proto3,casttype=github.com/cometbft/cometbft/libs/bytes.HexBytes" json:"crisis_to_distr_stores_merkle_hash,omitempty"`
	AccToCapabilityStoresMerkleHash   github_com_cometbft_cometbft_libs_bytes.HexBytes `protobuf:"bytes,4,opt,name=acc_to_capability_stores_merkle_hash,json=accToCapabilityStoresMerkleHash,proto3,casttype=github.com/cometbft/cometbft/libs/bytes.HexBytes" json:"acc_to_capability_stores_merkle_hash,omitempty"`
	EvidenceToIcahostStoresMerkleHash github_com_cometbft_cometbft_libs_bytes.HexBytes `protobuf:"bytes,5,opt,name=evidence_to_icahost_stores_merkle_hash,json=evidenceToIcahostStoresMerkleHash,proto3,casttype=github.com/cometbft/cometbft/libs/bytes.HexBytes" json:"evidence_to_icahost_stores_merkle_hash,omitempty"`
	MintToVestingStoresMerkleHash     github_com_cometbft_cometbft_libs_bytes.HexBytes `protobuf:"bytes,6,opt,name=mint_to_vesting_stores_merkle_hash,json=mintToVestingStoresMerkleHash,proto3,casttype=github.com/cometbft/cometbft/libs/bytes.HexBytes" json:"mint_to_vesting_stores_merkle_hash,omitempty"`
}

func (m *MutiStoreTreeFields) Reset()         { *m = MutiStoreTreeFields{} }
func (m *MutiStoreTreeFields) String() string { return proto.CompactTextString(m) }
func (*MutiStoreTreeFields) ProtoMessage()    {}
func (*MutiStoreTreeFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_29da94be13d33792, []int{0}
}
func (m *MutiStoreTreeFields) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MutiStoreTreeFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MutiStoreTreeFields.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MutiStoreTreeFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutiStoreTreeFields.Merge(m, src)
}
func (m *MutiStoreTreeFields) XXX_Size() int {
	return m.Size()
}
func (m *MutiStoreTreeFields) XXX_DiscardUnknown() {
	xxx_messageInfo_MutiStoreTreeFields.DiscardUnknown(m)
}

var xxx_messageInfo_MutiStoreTreeFields proto.InternalMessageInfo

func (m *MutiStoreTreeFields) GetCosmonovaIavlStateHash() github_com_cometbft_cometbft_libs_bytes.HexBytes {
	if m != nil {
		return m.CosmonovaIavlStateHash
	}
	return nil
}

func (m *MutiStoreTreeFields) GetConsensusStoreMerkleHash() github_com_cometbft_cometbft_libs_bytes.HexBytes {
	if m != nil {
		return m.ConsensusStoreMerkleHash
	}
	return nil
}

func (m *MutiStoreTreeFields) GetCrisisToDistrStoresMerkleHash() github_com_cometbft_cometbft_libs_bytes.HexBytes {
	if m != nil {
		return m.CrisisToDistrStoresMerkleHash
	}
	return nil
}

func (m *MutiStoreTreeFields) GetAccToCapabilityStoresMerkleHash() github_com_cometbft_cometbft_libs_bytes.HexBytes {
	if m != nil {
		return m.AccToCapabilityStoresMerkleHash
	}
	return nil
}

func (m *MutiStoreTreeFields) GetEvidenceToIcahostStoresMerkleHash() github_com_cometbft_cometbft_libs_bytes.HexBytes {
	if m != nil {
		return m.EvidenceToIcahostStoresMerkleHash
	}
	return nil
}

func (m *MutiStoreTreeFields) GetMintToVestingStoresMerkleHash() github_com_cometbft_cometbft_libs_bytes.HexBytes {
	if m != nil {
		return m.MintToVestingStoresMerkleHash
	}
	return nil
}

func init() {
	proto.RegisterType((*MutiStoreTreeFields)(nil), "cosmonova.client.MutiStoreTreeFields")
}

func init() {
	proto.RegisterFile("cosmonova/client/mutistore_tree.proto", fileDescriptor_29da94be13d33792)
}

var fileDescriptor_29da94be13d33792 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x3f, 0xeb, 0xd3, 0x60,
	0x10, 0x80, 0x1b, 0xff, 0x74, 0x78, 0x71, 0x90, 0x2a, 0x52, 0x15, 0xf3, 0xd3, 0xa2, 0xe2, 0x94,
	0x88, 0xfa, 0x09, 0xaa, 0x48, 0x3b, 0x74, 0xb1, 0xc1, 0xc1, 0x25, 0xbc, 0x79, 0x7b, 0x26, 0x87,
	0x49, 0xae, 0xe4, 0x2e, 0xa1, 0x9d, 0x3a, 0x38, 0x38, 0x09, 0x7e, 0x2c, 0xc7, 0x8e, 0x4e, 0x22,
	0xed, 0xb7, 0x70, 0x92, 0x37, 0xb1, 0xa9, 0xfe, 0xda, 0x2d, 0xdb, 0x41, 0x8e, 0xe7, 0x79, 0x5e,
	0xc2, 0xa9, 0x27, 0x86, 0x38, 0xa3, 0x9c, 0x2a, 0xed, 0x9b, 0x14, 0x21, 0x17, 0x3f, 0x2b, 0x05,
	0x59, 0xa8, 0x80, 0x50, 0x0a, 0x00, 0x6f, 0x59, 0x90, 0xd0, 0xe0, 0x66, 0xbb, 0xe6, 0x35, 0x6b,
	0xf7, 0x6e, 0xc7, 0x14, 0x53, 0xfd, 0xd1, 0xb7, 0x53, 0xb3, 0x37, 0xfa, 0xda, 0x57, 0xb7, 0x66,
	0xa5, 0xe0, 0xdc, 0x02, 0x82, 0x02, 0xe0, 0x2d, 0x42, 0xba, 0xe0, 0x01, 0xa9, 0xbb, 0x2d, 0x21,
	0x44, 0x5d, 0xa5, 0x21, 0x8b, 0x16, 0x08, 0x13, 0xcd, 0xc9, 0xd0, 0x79, 0xe8, 0x3c, 0xbb, 0x31,
	0x7e, 0xf5, 0xfb, 0xe7, 0xc5, 0xf3, 0x18, 0x25, 0x29, 0x23, 0xcf, 0x50, 0xe6, 0x1b, 0xca, 0x40,
	0xa2, 0x8f, 0x72, 0x1c, 0x52, 0x8c, 0xd8, 0x8f, 0xd6, 0x02, 0xec, 0x4d, 0x60, 0x35, 0xb6, 0xc3,
	0xbb, 0x3b, 0x2d, 0x76, 0xaa, 0xab, 0x74, 0x6e, 0xa1, 0x13, 0xcd, 0xc9, 0x80, 0xd5, 0x7d, 0x43,
	0x39, 0x43, 0xce, 0x25, 0x87, 0xcd, 0x73, 0x32, 0x28, 0x3e, 0xa5, 0x7f, 0x95, 0x57, 0x3a, 0x28,
	0x87, 0x2d, 0xb8, 0x7e, 0xe5, 0xac, 0xc6, 0xd6, 0xd2, 0x8d, 0x1a, 0x99, 0x02, 0x19, 0x39, 0x14,
	0x0a, 0x17, 0xc8, 0x52, 0x34, 0x6a, 0xfe, 0xcf, 0x7d, 0xb5, 0x83, 0xfb, 0x41, 0xc3, 0x0f, 0xe8,
	0x8d, 0xa5, 0xd7, 0x7e, 0xfe, 0x27, 0xe0, 0xb3, 0xa3, 0x1e, 0x6b, 0x63, 0xac, 0xde, 0xe8, 0xa5,
	0x8e, 0x30, 0x45, 0x59, 0x9f, 0x6b, 0xb8, 0xd6, 0xa1, 0xe1, 0x42, 0x1b, 0x13, 0xd0, 0xeb, 0x96,
	0x7f, 0x52, 0xf1, 0xc5, 0x51, 0x4f, 0xa1, 0xc2, 0x05, 0xe4, 0x06, 0x6c, 0x0a, 0x1a, 0x9d, 0x10,
	0xcb, 0xb9, 0x8e, 0xeb, 0x1d, 0x3a, 0x1e, 0x1d, 0x1c, 0x01, 0x4d, 0x1b, 0xc3, 0x49, 0xc9, 0x46,
	0x8d, 0x32, 0xcc, 0xc5, 0x46, 0x54, 0xc0, 0x82, 0x79, 0x7c, 0x2e, 0xa2, 0xdf, 0xe5, 0x87, 0x58,
	0x7e, 0x40, 0xef, 0x1b, 0xfa, 0xe5, 0x80, 0xf1, 0x8b, 0xef, 0x3b, 0xd7, 0xd9, 0xee, 0x5c, 0xe7,
	0xd7, 0xce, 0x75, 0xbe, 0xed, 0xdd, 0xde, 0x76, 0xef, 0xf6, 0x7e, 0xec, 0xdd, 0xde, 0x87, 0xe1,
	0xf1, 0xf0, 0x56, 0x87, 0xd3, 0x93, 0xf5, 0x12, 0x38, 0xea, 0xd7, 0xa7, 0xf4, 0xf2, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xaf, 0xe4, 0x3f, 0x0a, 0x9b, 0x03, 0x00, 0x00,
}

func (m *MutiStoreTreeFields) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MutiStoreTreeFields) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MutiStoreTreeFields) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MintToVestingStoresMerkleHash) > 0 {
		i -= len(m.MintToVestingStoresMerkleHash)
		copy(dAtA[i:], m.MintToVestingStoresMerkleHash)
		i = encodeVarintMutistoreTree(dAtA, i, uint64(len(m.MintToVestingStoresMerkleHash)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.EvidenceToIcahostStoresMerkleHash) > 0 {
		i -= len(m.EvidenceToIcahostStoresMerkleHash)
		copy(dAtA[i:], m.EvidenceToIcahostStoresMerkleHash)
		i = encodeVarintMutistoreTree(dAtA, i, uint64(len(m.EvidenceToIcahostStoresMerkleHash)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AccToCapabilityStoresMerkleHash) > 0 {
		i -= len(m.AccToCapabilityStoresMerkleHash)
		copy(dAtA[i:], m.AccToCapabilityStoresMerkleHash)
		i = encodeVarintMutistoreTree(dAtA, i, uint64(len(m.AccToCapabilityStoresMerkleHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CrisisToDistrStoresMerkleHash) > 0 {
		i -= len(m.CrisisToDistrStoresMerkleHash)
		copy(dAtA[i:], m.CrisisToDistrStoresMerkleHash)
		i = encodeVarintMutistoreTree(dAtA, i, uint64(len(m.CrisisToDistrStoresMerkleHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConsensusStoreMerkleHash) > 0 {
		i -= len(m.ConsensusStoreMerkleHash)
		copy(dAtA[i:], m.ConsensusStoreMerkleHash)
		i = encodeVarintMutistoreTree(dAtA, i, uint64(len(m.ConsensusStoreMerkleHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CosmonovaIavlStateHash) > 0 {
		i -= len(m.CosmonovaIavlStateHash)
		copy(dAtA[i:], m.CosmonovaIavlStateHash)
		i = encodeVarintMutistoreTree(dAtA, i, uint64(len(m.CosmonovaIavlStateHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMutistoreTree(dAtA []byte, offset int, v uint64) int {
	offset -= sovMutistoreTree(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MutiStoreTreeFields) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CosmonovaIavlStateHash)
	if l > 0 {
		n += 1 + l + sovMutistoreTree(uint64(l))
	}
	l = len(m.ConsensusStoreMerkleHash)
	if l > 0 {
		n += 1 + l + sovMutistoreTree(uint64(l))
	}
	l = len(m.CrisisToDistrStoresMerkleHash)
	if l > 0 {
		n += 1 + l + sovMutistoreTree(uint64(l))
	}
	l = len(m.AccToCapabilityStoresMerkleHash)
	if l > 0 {
		n += 1 + l + sovMutistoreTree(uint64(l))
	}
	l = len(m.EvidenceToIcahostStoresMerkleHash)
	if l > 0 {
		n += 1 + l + sovMutistoreTree(uint64(l))
	}
	l = len(m.MintToVestingStoresMerkleHash)
	if l > 0 {
		n += 1 + l + sovMutistoreTree(uint64(l))
	}
	return n
}

func sovMutistoreTree(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMutistoreTree(x uint64) (n int) {
	return sovMutistoreTree(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MutiStoreTreeFields) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMutistoreTree
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
			return fmt.Errorf("proto: MutiStoreTreeFields: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MutiStoreTreeFields: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CosmonovaIavlStateHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMutistoreTree
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
				return ErrInvalidLengthMutistoreTree
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMutistoreTree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CosmonovaIavlStateHash = append(m.CosmonovaIavlStateHash[:0], dAtA[iNdEx:postIndex]...)
			if m.CosmonovaIavlStateHash == nil {
				m.CosmonovaIavlStateHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusStoreMerkleHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMutistoreTree
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
				return ErrInvalidLengthMutistoreTree
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMutistoreTree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsensusStoreMerkleHash = append(m.ConsensusStoreMerkleHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ConsensusStoreMerkleHash == nil {
				m.ConsensusStoreMerkleHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrisisToDistrStoresMerkleHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMutistoreTree
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
				return ErrInvalidLengthMutistoreTree
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMutistoreTree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CrisisToDistrStoresMerkleHash = append(m.CrisisToDistrStoresMerkleHash[:0], dAtA[iNdEx:postIndex]...)
			if m.CrisisToDistrStoresMerkleHash == nil {
				m.CrisisToDistrStoresMerkleHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccToCapabilityStoresMerkleHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMutistoreTree
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
				return ErrInvalidLengthMutistoreTree
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMutistoreTree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccToCapabilityStoresMerkleHash = append(m.AccToCapabilityStoresMerkleHash[:0], dAtA[iNdEx:postIndex]...)
			if m.AccToCapabilityStoresMerkleHash == nil {
				m.AccToCapabilityStoresMerkleHash = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EvidenceToIcahostStoresMerkleHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMutistoreTree
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
				return ErrInvalidLengthMutistoreTree
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMutistoreTree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EvidenceToIcahostStoresMerkleHash = append(m.EvidenceToIcahostStoresMerkleHash[:0], dAtA[iNdEx:postIndex]...)
			if m.EvidenceToIcahostStoresMerkleHash == nil {
				m.EvidenceToIcahostStoresMerkleHash = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintToVestingStoresMerkleHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMutistoreTree
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
				return ErrInvalidLengthMutistoreTree
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMutistoreTree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintToVestingStoresMerkleHash = append(m.MintToVestingStoresMerkleHash[:0], dAtA[iNdEx:postIndex]...)
			if m.MintToVestingStoresMerkleHash == nil {
				m.MintToVestingStoresMerkleHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMutistoreTree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMutistoreTree
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
func skipMutistoreTree(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMutistoreTree
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
					return 0, ErrIntOverflowMutistoreTree
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
					return 0, ErrIntOverflowMutistoreTree
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
				return 0, ErrInvalidLengthMutistoreTree
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMutistoreTree
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMutistoreTree
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMutistoreTree        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMutistoreTree          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMutistoreTree = fmt.Errorf("proto: unexpected end of group")
)