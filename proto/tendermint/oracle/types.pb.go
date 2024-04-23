// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/oracle/types.proto

package oracle

import (
	fmt "fmt"
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

type Vote struct {
	Validator []byte `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	OracleId  string `protobuf:"bytes,2,opt,name=oracle_id,json=oracleId,proto3" json:"oracle_id,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Data      string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed9227d272ed5d90, []int{0}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return m.Size()
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetValidator() []byte {
	if m != nil {
		return m.Validator
	}
	return nil
}

func (m *Vote) GetOracleId() string {
	if m != nil {
		return m.OracleId
	}
	return ""
}

func (m *Vote) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Vote) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type GossipVote struct {
	ValidatorIndex  int32   `protobuf:"varint,1,opt,name=validator_index,json=validatorIndex,proto3" json:"validator_index,omitempty"`
	SignType        string  `protobuf:"bytes,2,opt,name=sign_type,json=signType,proto3" json:"sign_type,omitempty"`
	Votes           []*Vote `protobuf:"bytes,3,rep,name=votes,proto3" json:"votes,omitempty"`
	SignedTimestamp int64   `protobuf:"varint,4,opt,name=signed_timestamp,json=signedTimestamp,proto3" json:"signed_timestamp,omitempty"`
	Signature       []byte  `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *GossipVote) Reset()         { *m = GossipVote{} }
func (m *GossipVote) String() string { return proto.CompactTextString(m) }
func (*GossipVote) ProtoMessage()    {}
func (*GossipVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed9227d272ed5d90, []int{1}
}
func (m *GossipVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GossipVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GossipVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GossipVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GossipVote.Merge(m, src)
}
func (m *GossipVote) XXX_Size() int {
	return m.Size()
}
func (m *GossipVote) XXX_DiscardUnknown() {
	xxx_messageInfo_GossipVote.DiscardUnknown(m)
}

var xxx_messageInfo_GossipVote proto.InternalMessageInfo

func (m *GossipVote) GetValidatorIndex() int32 {
	if m != nil {
		return m.ValidatorIndex
	}
	return 0
}

func (m *GossipVote) GetSignType() string {
	if m != nil {
		return m.SignType
	}
	return ""
}

func (m *GossipVote) GetVotes() []*Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *GossipVote) GetSignedTimestamp() int64 {
	if m != nil {
		return m.SignedTimestamp
	}
	return 0
}

func (m *GossipVote) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type CanonicalGossipVote struct {
	ValidatorIndex int32   `protobuf:"varint,1,opt,name=validator_index,json=validatorIndex,proto3" json:"validator_index,omitempty"`
	SignType       string  `protobuf:"bytes,2,opt,name=sign_type,json=signType,proto3" json:"sign_type,omitempty"`
	Votes          []*Vote `protobuf:"bytes,3,rep,name=votes,proto3" json:"votes,omitempty"`
}

func (m *CanonicalGossipVote) Reset()         { *m = CanonicalGossipVote{} }
func (m *CanonicalGossipVote) String() string { return proto.CompactTextString(m) }
func (*CanonicalGossipVote) ProtoMessage()    {}
func (*CanonicalGossipVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed9227d272ed5d90, []int{2}
}
func (m *CanonicalGossipVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CanonicalGossipVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CanonicalGossipVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CanonicalGossipVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CanonicalGossipVote.Merge(m, src)
}
func (m *CanonicalGossipVote) XXX_Size() int {
	return m.Size()
}
func (m *CanonicalGossipVote) XXX_DiscardUnknown() {
	xxx_messageInfo_CanonicalGossipVote.DiscardUnknown(m)
}

var xxx_messageInfo_CanonicalGossipVote proto.InternalMessageInfo

func (m *CanonicalGossipVote) GetValidatorIndex() int32 {
	if m != nil {
		return m.ValidatorIndex
	}
	return 0
}

func (m *CanonicalGossipVote) GetSignType() string {
	if m != nil {
		return m.SignType
	}
	return ""
}

func (m *CanonicalGossipVote) GetVotes() []*Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func init() {
	proto.RegisterType((*Vote)(nil), "tendermint.oracle.Vote")
	proto.RegisterType((*GossipVote)(nil), "tendermint.oracle.GossipVote")
	proto.RegisterType((*CanonicalGossipVote)(nil), "tendermint.oracle.CanonicalGossipVote")
}

func init() { proto.RegisterFile("tendermint/oracle/types.proto", fileDescriptor_ed9227d272ed5d90) }

var fileDescriptor_ed9227d272ed5d90 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x6b, 0x92, 0x22, 0x6a, 0x10, 0x05, 0x33, 0x10, 0x89, 0x12, 0x45, 0x5d, 0x08, 0x03,
	0x89, 0x04, 0x9c, 0x00, 0x06, 0xd4, 0x85, 0x21, 0xaa, 0x18, 0x58, 0x22, 0x37, 0x36, 0xc5, 0x52,
	0x63, 0x47, 0xf1, 0x6b, 0x45, 0xaf, 0xc0, 0xc4, 0xb1, 0x10, 0x53, 0x47, 0x46, 0xd4, 0x5e, 0x04,
	0xd9, 0x16, 0xc9, 0xd0, 0x03, 0xb0, 0x59, 0xdf, 0x7b, 0xbf, 0xfe, 0xff, 0xb7, 0x1e, 0x3e, 0x07,
	0x2e, 0x19, 0xaf, 0x4b, 0x21, 0x21, 0x55, 0x35, 0x2d, 0x66, 0x3c, 0x85, 0x65, 0xc5, 0x75, 0x52,
	0xd5, 0x0a, 0x14, 0x39, 0x6e, 0xc7, 0x89, 0x1b, 0x0f, 0x35, 0xf6, 0x9f, 0x14, 0x70, 0x32, 0xc0,
	0xbd, 0x05, 0x9d, 0x09, 0x46, 0x41, 0xd5, 0x01, 0x8a, 0x50, 0x7c, 0x90, 0xb5, 0x80, 0x9c, 0xe1,
	0x9e, 0xdb, 0xcf, 0x05, 0x0b, 0x76, 0x22, 0x14, 0xf7, 0xb2, 0x3d, 0x07, 0x46, 0xcc, 0x48, 0x41,
	0x94, 0x5c, 0x03, 0x2d, 0xab, 0xc0, 0x8b, 0x50, 0xec, 0x65, 0x2d, 0x20, 0x04, 0xfb, 0x8c, 0x02,
	0x0d, 0x7c, 0xab, 0xb2, 0xef, 0xe1, 0x17, 0xc2, 0xf8, 0x41, 0x69, 0x2d, 0x2a, 0xeb, 0x7d, 0x81,
	0xfb, 0x8d, 0x55, 0x2e, 0x24, 0xe3, 0x6f, 0x36, 0x41, 0x37, 0x3b, 0x6c, 0xf0, 0xc8, 0x50, 0x13,
	0x43, 0x8b, 0xa9, 0xcc, 0x4d, 0xa7, 0xbf, 0x18, 0x06, 0x8c, 0x97, 0x15, 0x27, 0x57, 0xb8, 0xbb,
	0x50, 0xc0, 0x75, 0xe0, 0x45, 0x5e, 0xbc, 0x7f, 0x7d, 0x9a, 0x6c, 0x95, 0x4d, 0x8c, 0x5b, 0xe6,
	0xb6, 0xc8, 0x25, 0x3e, 0x32, 0x52, 0xce, 0xf2, 0x36, 0xbc, 0x6f, 0xc3, 0xf7, 0x1d, 0x1f, 0x37,
	0x15, 0x06, 0xce, 0x96, 0xc2, 0xbc, 0xe6, 0x41, 0xd7, 0xfd, 0x4d, 0x03, 0x86, 0xef, 0x08, 0x9f,
	0xdc, 0x53, 0xa9, 0xa4, 0x28, 0xe8, 0xec, 0x9f, 0x5b, 0xdd, 0x3d, 0x7e, 0xae, 0x43, 0xb4, 0x5a,
	0x87, 0xe8, 0x67, 0x1d, 0xa2, 0x8f, 0x4d, 0xd8, 0x59, 0x6d, 0xc2, 0xce, 0xf7, 0x26, 0xec, 0x3c,
	0xdf, 0x4e, 0x05, 0xbc, 0xce, 0x27, 0x49, 0xa1, 0xca, 0xb4, 0x50, 0x25, 0x87, 0xc9, 0x0b, 0xb4,
	0x0f, 0x7b, 0x1f, 0xe9, 0xd6, 0xf5, 0x4c, 0x76, 0xed, 0xe0, 0xe6, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x58, 0x6b, 0x3b, 0xa5, 0x59, 0x02, 0x00, 0x00,
}

func (m *Vote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x22
	}
	if m.Timestamp != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x18
	}
	if len(m.OracleId) > 0 {
		i -= len(m.OracleId)
		copy(dAtA[i:], m.OracleId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.OracleId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GossipVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GossipVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GossipVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x2a
	}
	if m.SignedTimestamp != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.SignedTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SignType) > 0 {
		i -= len(m.SignType)
		copy(dAtA[i:], m.SignType)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.SignType)))
		i--
		dAtA[i] = 0x12
	}
	if m.ValidatorIndex != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ValidatorIndex))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CanonicalGossipVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CanonicalGossipVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CanonicalGossipVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SignType) > 0 {
		i -= len(m.SignType)
		copy(dAtA[i:], m.SignType)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.SignType)))
		i--
		dAtA[i] = 0x12
	}
	if m.ValidatorIndex != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ValidatorIndex))
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
func (m *Vote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.OracleId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovTypes(uint64(m.Timestamp))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *GossipVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ValidatorIndex != 0 {
		n += 1 + sovTypes(uint64(m.ValidatorIndex))
	}
	l = len(m.SignType)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.SignedTimestamp != 0 {
		n += 1 + sovTypes(uint64(m.SignedTimestamp))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *CanonicalGossipVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ValidatorIndex != 0 {
		n += 1 + sovTypes(uint64(m.ValidatorIndex))
	}
	l = len(m.SignType)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Vote) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Vote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = append(m.Validator[:0], dAtA[iNdEx:postIndex]...)
			if m.Validator == nil {
				m.Validator = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *GossipVote) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GossipVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GossipVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorIndex", wireType)
			}
			m.ValidatorIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidatorIndex |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, &Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedTimestamp", wireType)
			}
			m.SignedTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignedTimestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *CanonicalGossipVote) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CanonicalGossipVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CanonicalGossipVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorIndex", wireType)
			}
			m.ValidatorIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidatorIndex |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, &Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
