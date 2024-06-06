// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/store/v1beta1/listening.proto

package types

import (
	fmt "fmt"
	v1 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
	_ "github.com/cosmos/cosmos-proto"
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

// StoreKVPair is a KVStore KVPair used for listening to state changes (Sets and
// Deletes) It optionally includes the StoreKey for the originating KVStore and
// a Boolean flag to distinguish between Sets and Deletes
type StoreKVPair struct {
	StoreKey string `protobuf:"bytes,1,opt,name=store_key,json=storeKey,proto3" json:"store_key,omitempty"`
	Delete   bool   `protobuf:"varint,2,opt,name=delete,proto3" json:"delete,omitempty"`
	Key      []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value    []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *StoreKVPair) Reset()         { *m = StoreKVPair{} }
func (m *StoreKVPair) String() string { return proto.CompactTextString(m) }
func (*StoreKVPair) ProtoMessage()    {}
func (*StoreKVPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6caeb9d7b7c7c10, []int{0}
}
func (m *StoreKVPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoreKVPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoreKVPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoreKVPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreKVPair.Merge(m, src)
}
func (m *StoreKVPair) XXX_Size() int {
	return m.Size()
}
func (m *StoreKVPair) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreKVPair.DiscardUnknown(m)
}

var xxx_messageInfo_StoreKVPair proto.InternalMessageInfo

func (m *StoreKVPair) GetStoreKey() string {
	if m != nil {
		return m.StoreKey
	}
	return ""
}

func (m *StoreKVPair) GetDelete() bool {
	if m != nil {
		return m.Delete
	}
	return false
}

func (m *StoreKVPair) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *StoreKVPair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// BlockMetadata contains all the abci event data of a block
// the file streamer dump them into files together with the state changes.
type BlockMetadata struct {
	ResponseCommit        *v1.CommitResponse        `protobuf:"bytes,6,opt,name=response_commit,json=responseCommit,proto3" json:"response_commit,omitempty"`
	RequestFinalizeBlock  *v1.FinalizeBlockRequest  `protobuf:"bytes,7,opt,name=request_finalize_block,json=requestFinalizeBlock,proto3" json:"request_finalize_block,omitempty"`
	ResponseFinalizeBlock *v1.FinalizeBlockResponse `protobuf:"bytes,8,opt,name=response_finalize_block,json=responseFinalizeBlock,proto3" json:"response_finalize_block,omitempty"`
}

func (m *BlockMetadata) Reset()         { *m = BlockMetadata{} }
func (m *BlockMetadata) String() string { return proto.CompactTextString(m) }
func (*BlockMetadata) ProtoMessage()    {}
func (*BlockMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6caeb9d7b7c7c10, []int{1}
}
func (m *BlockMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockMetadata.Merge(m, src)
}
func (m *BlockMetadata) XXX_Size() int {
	return m.Size()
}
func (m *BlockMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_BlockMetadata proto.InternalMessageInfo

func (m *BlockMetadata) GetResponseCommit() *v1.CommitResponse {
	if m != nil {
		return m.ResponseCommit
	}
	return nil
}

func (m *BlockMetadata) GetRequestFinalizeBlock() *v1.FinalizeBlockRequest {
	if m != nil {
		return m.RequestFinalizeBlock
	}
	return nil
}

func (m *BlockMetadata) GetResponseFinalizeBlock() *v1.FinalizeBlockResponse {
	if m != nil {
		return m.ResponseFinalizeBlock
	}
	return nil
}

func init() {
	proto.RegisterType((*StoreKVPair)(nil), "cosmos.store.v1beta1.StoreKVPair")
	proto.RegisterType((*BlockMetadata)(nil), "cosmos.store.v1beta1.BlockMetadata")
}

func init() {
	proto.RegisterFile("cosmos/store/v1beta1/listening.proto", fileDescriptor_b6caeb9d7b7c7c10)
}

var fileDescriptor_b6caeb9d7b7c7c10 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0xd6, 0x2d, 0x9e, 0x07, 0x2c, 0x32, 0x65, 0x84, 0x81, 0xa2, 0x68, 0x42, 0xd0,
	0xcb, 0x1c, 0xba, 0x71, 0xe2, 0x38, 0x24, 0x24, 0x32, 0x21, 0xa1, 0x20, 0x71, 0x40, 0x48, 0x91,
	0x93, 0xbe, 0x21, 0xab, 0x69, 0x5c, 0x62, 0x2f, 0x52, 0xb9, 0xf0, 0x15, 0xf8, 0x30, 0x48, 0x7c,
	0x05, 0x8e, 0x13, 0x27, 0x8e, 0xa8, 0xfd, 0x22, 0x28, 0xb6, 0x8b, 0x34, 0x38, 0xec, 0x96, 0xff,
	0x7b, 0xff, 0xff, 0xcf, 0xcf, 0xf1, 0xa3, 0x8f, 0x4a, 0xa5, 0x17, 0x4a, 0x27, 0xda, 0xa8, 0x06,
	0x92, 0x76, 0x5a, 0x80, 0x11, 0xd3, 0xa4, 0x92, 0xda, 0x40, 0x2d, 0xeb, 0x8f, 0x7c, 0xd9, 0x28,
	0xa3, 0xd8, 0xd8, 0xb9, 0xb8, 0x75, 0x71, 0xef, 0x3a, 0x78, 0x58, 0xaa, 0x05, 0x98, 0xe2, 0xdc,
	0x24, 0xa2, 0x28, 0x65, 0xd2, 0x4e, 0x13, 0xb3, 0x5a, 0x82, 0x76, 0x99, 0x83, 0xfb, 0x2e, 0x93,
	0x5b, 0x95, 0x78, 0x80, 0x15, 0x87, 0x5f, 0xe8, 0xee, 0xdb, 0x8e, 0x74, 0xf6, 0xee, 0x8d, 0x90,
	0x0d, 0x7b, 0x40, 0x77, 0x2c, 0x38, 0x9f, 0xc3, 0x2a, 0x44, 0x31, 0x9a, 0xec, 0x64, 0xc4, 0x16,
	0xce, 0x60, 0xc5, 0xf6, 0xe9, 0x68, 0x06, 0x15, 0x18, 0x08, 0xfb, 0x31, 0x9a, 0x90, 0xcc, 0x2b,
	0x16, 0xd0, 0x41, 0x67, 0x1f, 0xc4, 0x68, 0x72, 0x33, 0xeb, 0x3e, 0xd9, 0x98, 0x0e, 0x5b, 0x51,
	0x5d, 0x40, 0x88, 0x6d, 0xcd, 0x89, 0xe7, 0x77, 0x7e, 0x7e, 0x3b, 0xda, 0x73, 0xa7, 0x1f, 0xe9,
	0xd9, 0x3c, 0x7e, 0xca, 0x9f, 0x9d, 0x1c, 0x7e, 0xef, 0xd3, 0x5b, 0xa7, 0x95, 0x2a, 0xe7, 0xaf,
	0xc1, 0x88, 0x99, 0x30, 0x82, 0xbd, 0xa2, 0x7b, 0x0d, 0xe8, 0xa5, 0xaa, 0x35, 0xe4, 0xa5, 0x5a,
	0x2c, 0xa4, 0x09, 0x47, 0x31, 0x9a, 0xec, 0x1e, 0xc7, 0x7c, 0x7b, 0x4b, 0xde, 0xdd, 0x92, 0xb7,
	0x53, 0xfe, 0xc2, 0xf6, 0x33, 0x6f, 0xcf, 0x6e, 0x6f, 0x83, 0xae, 0xce, 0x3e, 0xd0, 0xfd, 0x06,
	0x3e, 0x5d, 0x80, 0x36, 0xf9, 0xb9, 0xac, 0x45, 0x25, 0x3f, 0x43, 0x5e, 0x74, 0x87, 0x85, 0x37,
	0x2c, 0xf1, 0xf1, 0xff, 0xc4, 0x97, 0xde, 0x67, 0x67, 0xca, 0x5c, 0x38, 0x1b, 0x7b, 0xca, 0x95,
	0x26, 0xcb, 0xe9, 0xbd, 0xbf, 0x83, 0xfe, 0x83, 0x27, 0x16, 0xff, 0xe4, 0x5a, 0xbc, 0x9f, 0xfb,
	0xee, 0x96, 0x73, 0xa5, 0x9d, 0x62, 0x82, 0x82, 0x7e, 0x8a, 0x49, 0x3f, 0x18, 0xa4, 0x98, 0x0c,
	0x02, 0x9c, 0x62, 0x82, 0x83, 0x61, 0x8a, 0xc9, 0x30, 0x18, 0x9d, 0x1e, 0xff, 0x58, 0x47, 0xe8,
	0x72, 0x1d, 0xa1, 0xdf, 0xeb, 0x08, 0x7d, 0xdd, 0x44, 0xbd, 0xcb, 0x4d, 0xd4, 0xfb, 0xb5, 0x89,
	0x7a, 0xef, 0x43, 0xf7, 0x93, 0xf5, 0x6c, 0xce, 0xa5, 0xf2, 0xfb, 0x64, 0xf7, 0xa1, 0x18, 0xd9,
	0x57, 0x3f, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x98, 0x9f, 0x12, 0x13, 0x6c, 0x02, 0x00, 0x00,
}

func (m *StoreKVPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreKVPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoreKVPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintListening(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintListening(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Delete {
		i--
		if m.Delete {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.StoreKey) > 0 {
		i -= len(m.StoreKey)
		copy(dAtA[i:], m.StoreKey)
		i = encodeVarintListening(dAtA, i, uint64(len(m.StoreKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ResponseFinalizeBlock != nil {
		{
			size, err := m.ResponseFinalizeBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintListening(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.RequestFinalizeBlock != nil {
		{
			size, err := m.RequestFinalizeBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintListening(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.ResponseCommit != nil {
		{
			size, err := m.ResponseCommit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintListening(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}

func encodeVarintListening(dAtA []byte, offset int, v uint64) int {
	offset -= sovListening(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoreKVPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StoreKey)
	if l > 0 {
		n += 1 + l + sovListening(uint64(l))
	}
	if m.Delete {
		n += 2
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovListening(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovListening(uint64(l))
	}
	return n
}

func (m *BlockMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ResponseCommit != nil {
		l = m.ResponseCommit.Size()
		n += 1 + l + sovListening(uint64(l))
	}
	if m.RequestFinalizeBlock != nil {
		l = m.RequestFinalizeBlock.Size()
		n += 1 + l + sovListening(uint64(l))
	}
	if m.ResponseFinalizeBlock != nil {
		l = m.ResponseFinalizeBlock.Size()
		n += 1 + l + sovListening(uint64(l))
	}
	return n
}

func sovListening(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozListening(x uint64) (n int) {
	return sovListening(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoreKVPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListening
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
			return fmt.Errorf("proto: StoreKVPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreKVPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
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
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoreKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delete", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
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
			m.Delete = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
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
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
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
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListening(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthListening
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
func (m *BlockMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListening
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
			return fmt.Errorf("proto: BlockMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseCommit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
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
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ResponseCommit == nil {
				m.ResponseCommit = &v1.CommitResponse{}
			}
			if err := m.ResponseCommit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestFinalizeBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
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
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RequestFinalizeBlock == nil {
				m.RequestFinalizeBlock = &v1.FinalizeBlockRequest{}
			}
			if err := m.RequestFinalizeBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseFinalizeBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
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
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ResponseFinalizeBlock == nil {
				m.ResponseFinalizeBlock = &v1.FinalizeBlockResponse{}
			}
			if err := m.ResponseFinalizeBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListening(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthListening
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
func skipListening(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowListening
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
					return 0, ErrIntOverflowListening
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
					return 0, ErrIntOverflowListening
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
				return 0, ErrInvalidLengthListening
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupListening
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthListening
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthListening        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowListening          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupListening = fmt.Errorf("proto: unexpected end of group")
)