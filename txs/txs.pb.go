// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: txs.proto

package txs // import "github.com/hyperledger/burrow/txs"

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import crypto "github.com/hyperledger/burrow/crypto"

import github_com_hyperledger_burrow_crypto "github.com/hyperledger/burrow/crypto"
import github_com_hyperledger_burrow_txs_payload "github.com/hyperledger/burrow/txs/payload"
import github_com_hyperledger_burrow_binary "github.com/hyperledger/burrow/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// An envelope contains both the signable Tx and the signatures for each input (in signatories)
type Envelope struct {
	Signatories []Signatory `protobuf:"bytes,1,rep,name=Signatories" json:"Signatories"`
	// Canonical bytes of the Tx ready to be signed
	Tx                   *Tx      `protobuf:"bytes,2,opt,name=Tx,proto3,customtype=Tx" json:"Tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Envelope) Reset()      { *m = Envelope{} }
func (*Envelope) ProtoMessage() {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_txs_d68f38b98c76a1e7, []int{0}
}
func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(dst, src)
}
func (m *Envelope) XXX_Size() int {
	return m.Size()
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func (m *Envelope) GetSignatories() []Signatory {
	if m != nil {
		return m.Signatories
	}
	return nil
}

func (*Envelope) XXX_MessageName() string {
	return "txs.Envelope"
}

// Signatory contains signature and one or both of Address and PublicKey to identify the signer
type Signatory struct {
	Address              *github_com_hyperledger_burrow_crypto.Address  `protobuf:"bytes,1,opt,name=Address,proto3,customtype=github.com/hyperledger/burrow/crypto.Address" json:"Address,omitempty"`
	PublicKey            *crypto.PublicKey                              `protobuf:"bytes,2,opt,name=PublicKey" json:"PublicKey,omitempty"`
	Signature            github_com_hyperledger_burrow_crypto.Signature `protobuf:"bytes,3,opt,name=Signature,proto3,customtype=github.com/hyperledger/burrow/crypto.Signature" json:"Signature"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *Signatory) Reset()         { *m = Signatory{} }
func (m *Signatory) String() string { return proto.CompactTextString(m) }
func (*Signatory) ProtoMessage()    {}
func (*Signatory) Descriptor() ([]byte, []int) {
	return fileDescriptor_txs_d68f38b98c76a1e7, []int{1}
}
func (m *Signatory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Signatory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Signatory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Signatory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signatory.Merge(dst, src)
}
func (m *Signatory) XXX_Size() int {
	return m.Size()
}
func (m *Signatory) XXX_DiscardUnknown() {
	xxx_messageInfo_Signatory.DiscardUnknown(m)
}

var xxx_messageInfo_Signatory proto.InternalMessageInfo

func (m *Signatory) GetPublicKey() *crypto.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (*Signatory) XXX_MessageName() string {
	return "txs.Signatory"
}

// BroadcastTx or Transaction receipt
type Receipt struct {
	// Transaction type
	TxType github_com_hyperledger_burrow_txs_payload.Type `protobuf:"varint,1,opt,name=TxType,proto3,casttype=github.com/hyperledger/burrow/txs/payload.Type" json:"TxType,omitempty"`
	// The hash of the transaction that caused this event to be generated
	TxHash github_com_hyperledger_burrow_binary.HexBytes `protobuf:"bytes,2,opt,name=TxHash,proto3,customtype=github.com/hyperledger/burrow/binary.HexBytes" json:"TxHash"`
	// Whether the transaction creates a contract
	CreatesContract bool `protobuf:"varint,3,opt,name=CreatesContract,proto3" json:"CreatesContract,omitempty"`
	// The address of the contract being called
	ContractAddress      github_com_hyperledger_burrow_crypto.Address `protobuf:"bytes,4,opt,name=ContractAddress,proto3,customtype=github.com/hyperledger/burrow/crypto.Address" json:"ContractAddress"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *Receipt) Reset()         { *m = Receipt{} }
func (m *Receipt) String() string { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()    {}
func (*Receipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_txs_d68f38b98c76a1e7, []int{2}
}
func (m *Receipt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Receipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Receipt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Receipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receipt.Merge(dst, src)
}
func (m *Receipt) XXX_Size() int {
	return m.Size()
}
func (m *Receipt) XXX_DiscardUnknown() {
	xxx_messageInfo_Receipt.DiscardUnknown(m)
}

var xxx_messageInfo_Receipt proto.InternalMessageInfo

func (m *Receipt) GetTxType() github_com_hyperledger_burrow_txs_payload.Type {
	if m != nil {
		return m.TxType
	}
	return 0
}

func (m *Receipt) GetCreatesContract() bool {
	if m != nil {
		return m.CreatesContract
	}
	return false
}

func (*Receipt) XXX_MessageName() string {
	return "txs.Receipt"
}
func init() {
	proto.RegisterType((*Envelope)(nil), "txs.Envelope")
	golang_proto.RegisterType((*Envelope)(nil), "txs.Envelope")
	proto.RegisterType((*Signatory)(nil), "txs.Signatory")
	golang_proto.RegisterType((*Signatory)(nil), "txs.Signatory")
	proto.RegisterType((*Receipt)(nil), "txs.Receipt")
	golang_proto.RegisterType((*Receipt)(nil), "txs.Receipt")
}
func (m *Envelope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Envelope) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Signatories) > 0 {
		for _, msg := range m.Signatories {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTxs(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Tx != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTxs(dAtA, i, uint64(m.Tx.Size()))
		n1, err := m.Tx.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Signatory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Signatory) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Address != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTxs(dAtA, i, uint64(m.Address.Size()))
		n2, err := m.Address.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.PublicKey != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTxs(dAtA, i, uint64(m.PublicKey.Size()))
		n3, err := m.PublicKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintTxs(dAtA, i, uint64(m.Signature.Size()))
	n4, err := m.Signature.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Receipt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Receipt) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TxType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTxs(dAtA, i, uint64(m.TxType))
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintTxs(dAtA, i, uint64(m.TxHash.Size()))
	n5, err := m.TxHash.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if m.CreatesContract {
		dAtA[i] = 0x18
		i++
		if m.CreatesContract {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintTxs(dAtA, i, uint64(m.ContractAddress.Size()))
	n6, err := m.ContractAddress.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTxs(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Envelope) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Signatories) > 0 {
		for _, e := range m.Signatories {
			l = e.Size()
			n += 1 + l + sovTxs(uint64(l))
		}
	}
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovTxs(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Signatory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Address != nil {
		l = m.Address.Size()
		n += 1 + l + sovTxs(uint64(l))
	}
	if m.PublicKey != nil {
		l = m.PublicKey.Size()
		n += 1 + l + sovTxs(uint64(l))
	}
	l = m.Signature.Size()
	n += 1 + l + sovTxs(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Receipt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TxType != 0 {
		n += 1 + sovTxs(uint64(m.TxType))
	}
	l = m.TxHash.Size()
	n += 1 + l + sovTxs(uint64(l))
	if m.CreatesContract {
		n += 2
	}
	l = m.ContractAddress.Size()
	n += 1 + l + sovTxs(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTxs(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTxs(x uint64) (n int) {
	return sovTxs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Envelope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Envelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Envelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatories", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatories = append(m.Signatories, Signatory{})
			if err := m.Signatories[len(m.Signatories)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v Tx
			m.Tx = &v
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTxs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Signatory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Signatory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Signatory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_hyperledger_burrow_crypto.Address
			m.Address = &v
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PublicKey == nil {
				m.PublicKey = &crypto.PublicKey{}
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Signature.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTxs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Receipt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Receipt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Receipt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxType", wireType)
			}
			m.TxType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxType |= (github_com_hyperledger_burrow_txs_payload.Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatesContract", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CreatesContract = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ContractAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTxs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTxs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTxs
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
					return 0, ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTxs
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTxs
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTxs
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTxs(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTxs = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTxs   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("txs.proto", fileDescriptor_txs_d68f38b98c76a1e7) }
func init() { golang_proto.RegisterFile("txs.proto", fileDescriptor_txs_d68f38b98c76a1e7) }

var fileDescriptor_txs_d68f38b98c76a1e7 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0xeb, 0xec, 0x6a, 0xdb, 0xf5, 0x16, 0x2a, 0x7c, 0x40, 0xab, 0x1e, 0x92, 0xb2, 0xa7,
	0x3d, 0xd0, 0x04, 0x2d, 0x50, 0x24, 0x6e, 0xa4, 0x42, 0xaa, 0x8a, 0x90, 0x90, 0xc9, 0x89, 0x03,
	0x22, 0x7f, 0x86, 0x6c, 0xa4, 0x10, 0x47, 0xb6, 0x03, 0xf6, 0x9b, 0x70, 0xe4, 0x09, 0x78, 0x06,
	0x8e, 0x7b, 0xe4, 0xbc, 0x87, 0x08, 0x6d, 0xdf, 0x82, 0x13, 0x8a, 0x71, 0xb6, 0x55, 0x0f, 0x94,
	0xde, 0x3c, 0xf6, 0x7c, 0xbf, 0xf9, 0x66, 0xc6, 0x78, 0x2c, 0x95, 0xf0, 0x6b, 0xce, 0x24, 0x23,
	0x03, 0xa9, 0xc4, 0xe1, 0x71, 0x5e, 0xc8, 0x65, 0x93, 0xf8, 0x29, 0xfb, 0x14, 0xe4, 0x2c, 0x67,
	0x81, 0x79, 0x4b, 0x9a, 0x8f, 0x26, 0x32, 0x81, 0x39, 0xfd, 0xd5, 0x1c, 0xee, 0xa7, 0x5c, 0xd7,
	0xd2, 0x46, 0xb3, 0x0f, 0x78, 0xef, 0x65, 0xf5, 0x19, 0x4a, 0x56, 0x03, 0x39, 0xc1, 0x93, 0xb7,
	0x45, 0x5e, 0xc5, 0x92, 0xf1, 0x02, 0xc4, 0x14, 0x1d, 0x0d, 0xe6, 0x93, 0xc5, 0x5d, 0xbf, 0x2b,
	0xd7, 0xdf, 0xeb, 0x70, 0xb8, 0x6a, 0xbd, 0x1d, 0x7a, 0x35, 0x91, 0xdc, 0xc7, 0x4e, 0xa4, 0xa6,
	0xce, 0x11, 0x9a, 0xef, 0x87, 0xa3, 0x75, 0xeb, 0x39, 0x91, 0xa2, 0x4e, 0xa4, 0x9e, 0x0f, 0xbf,
	0x7e, 0xf3, 0x76, 0x66, 0x2d, 0xc2, 0xe3, 0xad, 0x9c, 0x9c, 0xe3, 0xdd, 0x17, 0x59, 0xc6, 0x41,
	0x74, 0xfc, 0x4e, 0xf0, 0x68, 0xdd, 0x7a, 0x0f, 0xaf, 0x74, 0xb0, 0xd4, 0x35, 0xf0, 0x12, 0xb2,
	0x1c, 0x78, 0x90, 0x34, 0x9c, 0xb3, 0x2f, 0x81, 0x35, 0x6c, 0x75, 0xb4, 0x07, 0x90, 0x00, 0x8f,
	0xdf, 0x34, 0x49, 0x59, 0xa4, 0xaf, 0x40, 0x9b, 0xf2, 0x93, 0xc5, 0x3d, 0xdf, 0x26, 0x6f, 0x1f,
	0xe8, 0x65, 0x0e, 0x89, 0x7a, 0x27, 0x0d, 0x87, 0xe9, 0xc0, 0x94, 0x3f, 0xe9, 0xda, 0x59, 0xb7,
	0x9e, 0xff, 0x5f, 0x16, 0xb6, 0x6a, 0x7a, 0x09, 0x9a, 0x7d, 0x77, 0xf0, 0x2e, 0x85, 0x14, 0x8a,
	0x5a, 0x92, 0x73, 0x3c, 0x8a, 0x54, 0xa4, 0x6b, 0x30, 0xdd, 0xdd, 0x09, 0x17, 0xbf, 0x6f, 0x44,
	0x4b, 0x25, 0x82, 0x3a, 0xd6, 0x25, 0x8b, 0x33, 0xbf, 0x53, 0x52, 0x4b, 0x20, 0xaf, 0x3b, 0xd6,
	0x59, 0x2c, 0x96, 0x76, 0xb4, 0x4f, 0xad, 0xd5, 0xe3, 0x7f, 0xf3, 0x92, 0xa2, 0x8a, 0xb9, 0xf6,
	0xcf, 0x40, 0x85, 0x5a, 0x82, 0xa0, 0x16, 0x42, 0xe6, 0xf8, 0xe0, 0x94, 0x43, 0x2c, 0x41, 0x9c,
	0xb2, 0x4a, 0xf2, 0x38, 0x95, 0x66, 0x04, 0x7b, 0xf4, 0xfa, 0x35, 0x79, 0x8f, 0x0f, 0xfa, 0x73,
	0xbf, 0xab, 0xa1, 0x71, 0xf0, 0xc4, 0x3a, 0xb8, 0xdd, 0xbe, 0xae, 0xc3, 0xc2, 0x67, 0xab, 0x8d,
	0x8b, 0x7e, 0x6e, 0x5c, 0xf4, 0x6b, 0xe3, 0xa2, 0x1f, 0x17, 0x2e, 0x5a, 0x5d, 0xb8, 0xe8, 0xdd,
	0x83, 0x1b, 0xc7, 0x94, 0x8c, 0xcc, 0x9f, 0x7d, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x2f,
	0x9b, 0x0b, 0x02, 0x03, 0x00, 0x00,
}
