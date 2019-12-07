// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eth/v1alpha1/validator.proto

package eth

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DutiesRequest struct {
	Epoch                uint64   `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	PublicKeys           [][]byte `protobuf:"bytes,2,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DutiesRequest) Reset()         { *m = DutiesRequest{} }
func (m *DutiesRequest) String() string { return proto.CompactTextString(m) }
func (*DutiesRequest) ProtoMessage()    {}
func (*DutiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8ef4b6a7db76675, []int{0}
}

func (m *DutiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DutiesRequest.Unmarshal(m, b)
}
func (m *DutiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DutiesRequest.Marshal(b, m, deterministic)
}
func (m *DutiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DutiesRequest.Merge(m, src)
}
func (m *DutiesRequest) XXX_Size() int {
	return xxx_messageInfo_DutiesRequest.Size(m)
}
func (m *DutiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DutiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DutiesRequest proto.InternalMessageInfo

func (m *DutiesRequest) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *DutiesRequest) GetPublicKeys() [][]byte {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

type DutiesResponse struct {
	Duties               []*DutiesResponse_Duty `protobuf:"bytes,1,rep,name=duties,proto3" json:"duties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DutiesResponse) Reset()         { *m = DutiesResponse{} }
func (m *DutiesResponse) String() string { return proto.CompactTextString(m) }
func (*DutiesResponse) ProtoMessage()    {}
func (*DutiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8ef4b6a7db76675, []int{1}
}

func (m *DutiesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DutiesResponse.Unmarshal(m, b)
}
func (m *DutiesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DutiesResponse.Marshal(b, m, deterministic)
}
func (m *DutiesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DutiesResponse.Merge(m, src)
}
func (m *DutiesResponse) XXX_Size() int {
	return xxx_messageInfo_DutiesResponse.Size(m)
}
func (m *DutiesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DutiesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DutiesResponse proto.InternalMessageInfo

func (m *DutiesResponse) GetDuties() []*DutiesResponse_Duty {
	if m != nil {
		return m.Duties
	}
	return nil
}

type DutiesResponse_Duty struct {
	PublicKey            []byte   `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	AttestationSlot      uint64   `protobuf:"varint,2,opt,name=attestation_slot,json=attestationSlot,proto3" json:"attestation_slot,omitempty"`
	AttestationShard     uint64   `protobuf:"varint,3,opt,name=attestation_shard,json=attestationShard,proto3" json:"attestation_shard,omitempty"`
	BlockProposalSlot    uint64   `protobuf:"varint,4,opt,name=block_proposal_slot,json=blockProposalSlot,proto3" json:"block_proposal_slot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DutiesResponse_Duty) Reset()         { *m = DutiesResponse_Duty{} }
func (m *DutiesResponse_Duty) String() string { return proto.CompactTextString(m) }
func (*DutiesResponse_Duty) ProtoMessage()    {}
func (*DutiesResponse_Duty) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8ef4b6a7db76675, []int{1, 0}
}

func (m *DutiesResponse_Duty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DutiesResponse_Duty.Unmarshal(m, b)
}
func (m *DutiesResponse_Duty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DutiesResponse_Duty.Marshal(b, m, deterministic)
}
func (m *DutiesResponse_Duty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DutiesResponse_Duty.Merge(m, src)
}
func (m *DutiesResponse_Duty) XXX_Size() int {
	return xxx_messageInfo_DutiesResponse_Duty.Size(m)
}
func (m *DutiesResponse_Duty) XXX_DiscardUnknown() {
	xxx_messageInfo_DutiesResponse_Duty.DiscardUnknown(m)
}

var xxx_messageInfo_DutiesResponse_Duty proto.InternalMessageInfo

func (m *DutiesResponse_Duty) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *DutiesResponse_Duty) GetAttestationSlot() uint64 {
	if m != nil {
		return m.AttestationSlot
	}
	return 0
}

func (m *DutiesResponse_Duty) GetAttestationShard() uint64 {
	if m != nil {
		return m.AttestationShard
	}
	return 0
}

func (m *DutiesResponse_Duty) GetBlockProposalSlot() uint64 {
	if m != nil {
		return m.BlockProposalSlot
	}
	return 0
}

type BlockRequest struct {
	Slot                 uint64   `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	RandaoReveal         []byte   `protobuf:"bytes,2,opt,name=randao_reveal,json=randaoReveal,proto3" json:"randao_reveal,omitempty"`
	Graffiti             []byte   `protobuf:"bytes,3,opt,name=graffiti,proto3" json:"graffiti,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockRequest) Reset()         { *m = BlockRequest{} }
func (m *BlockRequest) String() string { return proto.CompactTextString(m) }
func (*BlockRequest) ProtoMessage()    {}
func (*BlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8ef4b6a7db76675, []int{2}
}

func (m *BlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockRequest.Unmarshal(m, b)
}
func (m *BlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockRequest.Marshal(b, m, deterministic)
}
func (m *BlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockRequest.Merge(m, src)
}
func (m *BlockRequest) XXX_Size() int {
	return xxx_messageInfo_BlockRequest.Size(m)
}
func (m *BlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlockRequest proto.InternalMessageInfo

func (m *BlockRequest) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *BlockRequest) GetRandaoReveal() []byte {
	if m != nil {
		return m.RandaoReveal
	}
	return nil
}

func (m *BlockRequest) GetGraffiti() []byte {
	if m != nil {
		return m.Graffiti
	}
	return nil
}

type AttestationDataRequest struct {
	ProofOfCustodyBit    []byte   `protobuf:"bytes,1,opt,name=proof_of_custody_bit,json=proofOfCustodyBit,proto3" json:"proof_of_custody_bit,omitempty"`
	Slot                 uint64   `protobuf:"varint,2,opt,name=slot,proto3" json:"slot,omitempty"`
	Shard                uint64   `protobuf:"varint,3,opt,name=shard,proto3" json:"shard,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttestationDataRequest) Reset()         { *m = AttestationDataRequest{} }
func (m *AttestationDataRequest) String() string { return proto.CompactTextString(m) }
func (*AttestationDataRequest) ProtoMessage()    {}
func (*AttestationDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8ef4b6a7db76675, []int{3}
}

func (m *AttestationDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestationDataRequest.Unmarshal(m, b)
}
func (m *AttestationDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestationDataRequest.Marshal(b, m, deterministic)
}
func (m *AttestationDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestationDataRequest.Merge(m, src)
}
func (m *AttestationDataRequest) XXX_Size() int {
	return xxx_messageInfo_AttestationDataRequest.Size(m)
}
func (m *AttestationDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestationDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttestationDataRequest proto.InternalMessageInfo

func (m *AttestationDataRequest) GetProofOfCustodyBit() []byte {
	if m != nil {
		return m.ProofOfCustodyBit
	}
	return nil
}

func (m *AttestationDataRequest) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *AttestationDataRequest) GetShard() uint64 {
	if m != nil {
		return m.Shard
	}
	return 0
}

type Validator struct {
	PublicKey                  []byte   `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	WithdrawalCredentials      []byte   `protobuf:"bytes,2,opt,name=withdrawal_credentials,json=withdrawalCredentials,proto3" json:"withdrawal_credentials,omitempty"`
	EffectiveBalance           uint64   `protobuf:"varint,3,opt,name=effective_balance,json=effectiveBalance,proto3" json:"effective_balance,omitempty"`
	Slashed                    bool     `protobuf:"varint,4,opt,name=slashed,proto3" json:"slashed,omitempty"`
	ActivationEligibilityEpoch uint64   `protobuf:"varint,5,opt,name=activation_eligibility_epoch,json=activationEligibilityEpoch,proto3" json:"activation_eligibility_epoch,omitempty"`
	ActivationEpoch            uint64   `protobuf:"varint,6,opt,name=activation_epoch,json=activationEpoch,proto3" json:"activation_epoch,omitempty"`
	ExitEpoch                  uint64   `protobuf:"varint,7,opt,name=exit_epoch,json=exitEpoch,proto3" json:"exit_epoch,omitempty"`
	WithdrawableEpoch          uint64   `protobuf:"varint,8,opt,name=withdrawable_epoch,json=withdrawableEpoch,proto3" json:"withdrawable_epoch,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8ef4b6a7db76675, []int{4}
}

func (m *Validator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Validator.Unmarshal(m, b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return xxx_messageInfo_Validator.Size(m)
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Validator) GetWithdrawalCredentials() []byte {
	if m != nil {
		return m.WithdrawalCredentials
	}
	return nil
}

func (m *Validator) GetEffectiveBalance() uint64 {
	if m != nil {
		return m.EffectiveBalance
	}
	return 0
}

func (m *Validator) GetSlashed() bool {
	if m != nil {
		return m.Slashed
	}
	return false
}

func (m *Validator) GetActivationEligibilityEpoch() uint64 {
	if m != nil {
		return m.ActivationEligibilityEpoch
	}
	return 0
}

func (m *Validator) GetActivationEpoch() uint64 {
	if m != nil {
		return m.ActivationEpoch
	}
	return 0
}

func (m *Validator) GetExitEpoch() uint64 {
	if m != nil {
		return m.ExitEpoch
	}
	return 0
}

func (m *Validator) GetWithdrawableEpoch() uint64 {
	if m != nil {
		return m.WithdrawableEpoch
	}
	return 0
}

type ValidatorParticipation struct {
	GlobalParticipationRate float32  `protobuf:"fixed32,1,opt,name=global_participation_rate,json=globalParticipationRate,proto3" json:"global_participation_rate,omitempty"`
	VotedEther              uint64   `protobuf:"varint,2,opt,name=voted_ether,json=votedEther,proto3" json:"voted_ether,omitempty"`
	EligibleEther           uint64   `protobuf:"varint,3,opt,name=eligible_ether,json=eligibleEther,proto3" json:"eligible_ether,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ValidatorParticipation) Reset()         { *m = ValidatorParticipation{} }
func (m *ValidatorParticipation) String() string { return proto.CompactTextString(m) }
func (*ValidatorParticipation) ProtoMessage()    {}
func (*ValidatorParticipation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8ef4b6a7db76675, []int{5}
}

func (m *ValidatorParticipation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorParticipation.Unmarshal(m, b)
}
func (m *ValidatorParticipation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorParticipation.Marshal(b, m, deterministic)
}
func (m *ValidatorParticipation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorParticipation.Merge(m, src)
}
func (m *ValidatorParticipation) XXX_Size() int {
	return xxx_messageInfo_ValidatorParticipation.Size(m)
}
func (m *ValidatorParticipation) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorParticipation.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorParticipation proto.InternalMessageInfo

func (m *ValidatorParticipation) GetGlobalParticipationRate() float32 {
	if m != nil {
		return m.GlobalParticipationRate
	}
	return 0
}

func (m *ValidatorParticipation) GetVotedEther() uint64 {
	if m != nil {
		return m.VotedEther
	}
	return 0
}

func (m *ValidatorParticipation) GetEligibleEther() uint64 {
	if m != nil {
		return m.EligibleEther
	}
	return 0
}

func init() {
	proto.RegisterType((*DutiesRequest)(nil), "ethereum.eth.v1alpha1.DutiesRequest")
	proto.RegisterType((*DutiesResponse)(nil), "ethereum.eth.v1alpha1.DutiesResponse")
	proto.RegisterType((*DutiesResponse_Duty)(nil), "ethereum.eth.v1alpha1.DutiesResponse.Duty")
	proto.RegisterType((*BlockRequest)(nil), "ethereum.eth.v1alpha1.BlockRequest")
	proto.RegisterType((*AttestationDataRequest)(nil), "ethereum.eth.v1alpha1.AttestationDataRequest")
	proto.RegisterType((*Validator)(nil), "ethereum.eth.v1alpha1.Validator")
	proto.RegisterType((*ValidatorParticipation)(nil), "ethereum.eth.v1alpha1.ValidatorParticipation")
}

func init() { proto.RegisterFile("eth/v1alpha1/validator.proto", fileDescriptor_f8ef4b6a7db76675) }

var fileDescriptor_f8ef4b6a7db76675 = []byte{
	// 908 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0xd6, 0x6e, 0x7e, 0x9a, 0x9c, 0x6c, 0x02, 0x99, 0x26, 0x61, 0xbb, 0xe4, 0x4f, 0x4e, 0x83,
	0x02, 0x55, 0x6d, 0xb5, 0x08, 0x2e, 0xe0, 0x06, 0xb6, 0x5d, 0x72, 0x81, 0x04, 0x91, 0x2b, 0xf5,
	0x02, 0x55, 0xb2, 0xc6, 0xf6, 0xd9, 0xf5, 0xa8, 0x13, 0x8f, 0xeb, 0x99, 0x4d, 0xbb, 0x48, 0x48,
	0x08, 0xf1, 0x02, 0x08, 0xb8, 0xe3, 0x09, 0x90, 0x78, 0x11, 0x6e, 0x79, 0x04, 0x78, 0x10, 0xe4,
	0x33, 0xf6, 0xda, 0x8b, 0x62, 0x25, 0xdc, 0xd9, 0xdf, 0xf9, 0xe6, 0x7c, 0xe7, 0x1f, 0xf6, 0xd1,
	0x24, 0xde, 0xd5, 0x23, 0x2e, 0xb3, 0x84, 0x3f, 0xf2, 0xae, 0xb8, 0x14, 0x31, 0x37, 0x2a, 0x77,
	0xb3, 0x5c, 0x19, 0xc5, 0x76, 0xd1, 0x24, 0x98, 0xe3, 0xf4, 0xd2, 0x45, 0x93, 0xb8, 0x15, 0x6d,
	0xb0, 0x3f, 0x51, 0x6a, 0x22, 0xd1, 0xe3, 0x99, 0xf0, 0x78, 0x9a, 0x2a, 0xc3, 0x8d, 0x50, 0xa9,
	0xb6, 0x8f, 0x06, 0xef, 0x96, 0x56, 0xfa, 0x0b, 0xa7, 0x63, 0x0f, 0x2f, 0x33, 0x33, 0x2b, 0x8d,
	0x47, 0x0b, 0x7a, 0x21, 0xf2, 0x48, 0xa5, 0x41, 0x28, 0x55, 0xf4, 0xb2, 0x24, 0x1c, 0x2e, 0x10,
	0xb8, 0x31, 0xa8, 0xad, 0x7b, 0x6b, 0x77, 0xbe, 0x80, 0xcd, 0xa7, 0x53, 0x23, 0x50, 0xfb, 0xf8,
	0x6a, 0x8a, 0xda, 0xb0, 0x1d, 0x58, 0xc1, 0x4c, 0x45, 0x49, 0xbf, 0x73, 0xdc, 0x39, 0x5b, 0xf6,
	0xed, 0x0f, 0x3b, 0x82, 0x8d, 0x6c, 0x1a, 0x4a, 0x11, 0x05, 0x2f, 0x71, 0xa6, 0xfb, 0xdd, 0xe3,
	0xa5, 0xb3, 0x9e, 0x0f, 0x16, 0xfa, 0x12, 0x67, 0xda, 0xf9, 0xb1, 0x0b, 0x5b, 0x95, 0x23, 0x9d,
	0xa9, 0x54, 0x23, 0x1b, 0xc2, 0x6a, 0x4c, 0x48, 0xbf, 0x73, 0xbc, 0x74, 0xb6, 0xf1, 0xf8, 0x03,
	0xf7, 0xda, 0xf4, 0xdd, 0xc5, 0x67, 0xc5, 0xef, 0xcc, 0x2f, 0x5f, 0x0e, 0xfe, 0xe8, 0xc0, 0x72,
	0x01, 0xb0, 0x03, 0x80, 0x3a, 0x00, 0x8a, 0xad, 0xe7, 0xaf, 0xcf, 0xf5, 0xd9, 0xfb, 0xf0, 0x76,
	0x23, 0xb7, 0x40, 0x4b, 0x65, 0xfa, 0x5d, 0x4a, 0xe0, 0xad, 0x06, 0xfe, 0x4c, 0x2a, 0xc3, 0x1e,
	0xc0, 0xf6, 0x02, 0x35, 0xe1, 0x79, 0xdc, 0x5f, 0x22, 0x6e, 0xd3, 0xc7, 0xb3, 0x02, 0x67, 0x2e,
	0xdc, 0xa5, 0x6a, 0x06, 0x59, 0xae, 0x32, 0xa5, 0xb9, 0xb4, 0xae, 0x97, 0x89, 0xbe, 0x4d, 0xa6,
	0x8b, 0xd2, 0x52, 0x38, 0x77, 0x22, 0xe8, 0x0d, 0x0b, 0xb0, 0xaa, 0x26, 0x83, 0x65, 0x7a, 0x60,
	0x8b, 0x49, 0xdf, 0xec, 0x04, 0x36, 0x73, 0x9e, 0xc6, 0x5c, 0x05, 0x39, 0x5e, 0x21, 0x97, 0x14,
	0x68, 0xcf, 0xef, 0x59, 0xd0, 0x27, 0x8c, 0x0d, 0x60, 0x6d, 0x92, 0xf3, 0xf1, 0x58, 0x18, 0x41,
	0xc1, 0xf5, 0xfc, 0xf9, 0xbf, 0xa3, 0x61, 0xef, 0xf3, 0x3a, 0xd0, 0xa7, 0xdc, 0xf0, 0x4a, 0xce,
	0x83, 0x9d, 0x2c, 0x57, 0x6a, 0x1c, 0xa8, 0x71, 0x10, 0x4d, 0xb5, 0x51, 0xf1, 0x2c, 0x08, 0x85,
	0x29, 0xeb, 0xb5, 0x4d, 0xb6, 0xaf, 0xc7, 0x4f, 0xac, 0x65, 0x28, 0xea, 0xf8, 0xba, 0x8d, 0xf8,
	0x76, 0x60, 0xa5, 0x59, 0x14, 0xfb, 0xe3, 0xfc, 0xdd, 0x85, 0xf5, 0xe7, 0xd5, 0x3c, 0xdf, 0xd4,
	0x8e, 0x8f, 0x60, 0xef, 0xb5, 0x30, 0x49, 0x9c, 0xf3, 0xd7, 0x5c, 0x06, 0x51, 0x8e, 0x31, 0xa6,
	0x46, 0x70, 0xa9, 0xcb, 0x5c, 0x77, 0x6b, 0xeb, 0x93, 0xda, 0x58, 0xb4, 0x06, 0xc7, 0x63, 0x8c,
	0x8c, 0xb8, 0xc2, 0x20, 0xe4, 0x92, 0xa7, 0x11, 0x56, 0xad, 0x99, 0x1b, 0x86, 0x16, 0x67, 0x7d,
	0xb8, 0xa3, 0x25, 0xd7, 0x09, 0xc6, 0xd4, 0x8e, 0x35, 0xbf, 0xfa, 0x65, 0x9f, 0xc1, 0x3e, 0x2f,
	0xa8, 0xb6, 0xc1, 0x28, 0xc5, 0x44, 0x84, 0x42, 0x0a, 0x33, 0x0b, 0xec, 0x64, 0xaf, 0x90, 0xc7,
	0x41, 0xcd, 0x19, 0xd5, 0x94, 0x11, 0x8d, 0x7b, 0x31, 0x4e, 0x0d, 0x0f, 0xf4, 0x6a, 0xb5, 0x1c,
	0xa7, 0xfa, 0x15, 0x51, 0x0f, 0x00, 0xf0, 0x8d, 0x30, 0x25, 0xe9, 0x0e, 0x91, 0xd6, 0x0b, 0xc4,
	0x9a, 0x1f, 0x02, 0x9b, 0xe7, 0x1a, 0x4a, 0x2c, 0x69, 0x6b, 0x76, 0x7e, 0x9a, 0x16, 0xa2, 0x3b,
	0xbf, 0x75, 0x60, 0x6f, 0x5e, 0xe5, 0x0b, 0x9e, 0x1b, 0x11, 0x89, 0x8c, 0xd4, 0xd8, 0x27, 0x70,
	0x6f, 0x22, 0x55, 0xc8, 0x65, 0x90, 0x35, 0xf1, 0x20, 0xe7, 0x06, 0xa9, 0x03, 0x5d, 0xff, 0x1d,
	0x4b, 0x58, 0x78, 0xe7, 0x73, 0x83, 0xc5, 0xfa, 0x5e, 0x29, 0x83, 0x71, 0x40, 0x1b, 0x58, 0x76,
	0x1b, 0x08, 0x1a, 0x15, 0x08, 0x3b, 0x85, 0x2d, 0x5b, 0xa7, 0x22, 0x44, 0xe2, 0xd8, 0xb2, 0x6f,
	0x56, 0x28, 0xd1, 0x1e, 0xff, 0xb4, 0x0a, 0x77, 0x87, 0x74, 0x64, 0xbe, 0x52, 0x31, 0xd6, 0xe3,
	0xf0, 0x7d, 0x07, 0xd6, 0xcf, 0xd1, 0xd8, 0x4d, 0x66, 0xf7, 0x6f, 0x58, 0x74, 0x9a, 0xd5, 0xc1,
	0xe9, 0xad, 0xce, 0x81, 0xf3, 0xde, 0x0f, 0x7f, 0xfd, 0xf3, 0x73, 0xf7, 0x98, 0x1d, 0x7a, 0xd7,
	0x9f, 0x56, 0xcf, 0x5e, 0x0a, 0xf6, 0x2d, 0xac, 0x9d, 0xa3, 0xa1, 0xe5, 0x63, 0x27, 0x2d, 0xae,
	0x9b, 0xab, 0x39, 0x70, 0xda, 0x48, 0x94, 0x1f, 0x51, 0x9d, 0x53, 0x12, 0x3f, 0x62, 0x07, 0x6d,
	0xe2, 0x74, 0x01, 0x58, 0x0e, 0x3d, 0x7b, 0x05, 0xf0, 0x7f, 0xe8, 0xef, 0xb9, 0xf6, 0xb0, 0xbb,
	0xd5, 0x61, 0x77, 0x47, 0xc5, 0x61, 0xaf, 0x34, 0x9d, 0x1b, 0x34, 0x7f, 0xe9, 0x00, 0x3b, 0x47,
	0xf3, 0x9f, 0x43, 0xc0, 0x1e, 0xb6, 0x48, 0x5f, 0x7f, 0x30, 0x5a, 0x8b, 0xd0, 0xa0, 0x3b, 0x0f,
	0x28, 0xa0, 0x53, 0x76, 0xd2, 0x16, 0x50, 0xe3, 0x6a, 0xb2, 0xef, 0x80, 0x95, 0xa5, 0x68, 0xb8,
	0x60, 0xb7, 0x90, 0x69, 0xad, 0x47, 0x29, 0xef, 0xdc, 0x4a, 0xfe, 0x15, 0x6c, 0x94, 0xf2, 0xa3,
	0x37, 0xc2, 0xb4, 0x4e, 0xe2, 0x73, 0x25, 0xa7, 0xa9, 0xe1, 0xf9, 0xac, 0x60, 0xb5, 0x2a, 0xdf,
	0x27, 0xe5, 0x43, 0x67, 0xbf, 0x4d, 0xb9, 0x58, 0xf3, 0xe1, 0xaf, 0x1d, 0xb8, 0xa7, 0xf2, 0xc9,
	0xf5, 0x4a, 0xc3, 0xad, 0x7a, 0x9b, 0x0b, 0xe7, 0x17, 0x9d, 0x6f, 0x3e, 0x9e, 0x08, 0x93, 0x4c,
	0x43, 0x37, 0x52, 0x97, 0x5e, 0x96, 0xcf, 0xf4, 0x25, 0x37, 0x22, 0x92, 0x3c, 0xd4, 0x5e, 0xe5,
	0x81, 0x67, 0x42, 0x2f, 0x48, 0x7e, 0x8a, 0x26, 0xf9, 0xbd, 0xbb, 0x3b, 0xaa, 0x14, 0x46, 0x0d,
	0x85, 0x3f, 0x6b, 0xfc, 0xc5, 0xc8, 0x24, 0x2f, 0x2a, 0x3c, 0x5c, 0xa5, 0x6c, 0x3e, 0xfc, 0x37,
	0x00, 0x00, 0xff, 0xff, 0x1b, 0x40, 0x73, 0x8f, 0x93, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BeaconNodeValidatorClient is the client API for BeaconNodeValidator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BeaconNodeValidatorClient interface {
	GetDuties(ctx context.Context, in *DutiesRequest, opts ...grpc.CallOption) (*DutiesResponse, error)
	GetBlock(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BeaconBlock, error)
	ProposeBlock(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetAttestationData(ctx context.Context, in *AttestationDataRequest, opts ...grpc.CallOption) (*Attestation, error)
	ProposeAttestation(ctx context.Context, in *Attestation, opts ...grpc.CallOption) (*empty.Empty, error)
	ProposeExit(ctx context.Context, in *VoluntaryExit, opts ...grpc.CallOption) (*empty.Empty, error)
}

type beaconNodeValidatorClient struct {
	cc *grpc.ClientConn
}

func NewBeaconNodeValidatorClient(cc *grpc.ClientConn) BeaconNodeValidatorClient {
	return &beaconNodeValidatorClient{cc}
}

func (c *beaconNodeValidatorClient) GetDuties(ctx context.Context, in *DutiesRequest, opts ...grpc.CallOption) (*DutiesResponse, error) {
	out := new(DutiesResponse)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.BeaconNodeValidator/GetDuties", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconNodeValidatorClient) GetBlock(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BeaconBlock, error) {
	out := new(BeaconBlock)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.BeaconNodeValidator/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconNodeValidatorClient) ProposeBlock(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.BeaconNodeValidator/ProposeBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconNodeValidatorClient) GetAttestationData(ctx context.Context, in *AttestationDataRequest, opts ...grpc.CallOption) (*Attestation, error) {
	out := new(Attestation)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.BeaconNodeValidator/GetAttestationData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconNodeValidatorClient) ProposeAttestation(ctx context.Context, in *Attestation, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.BeaconNodeValidator/ProposeAttestation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconNodeValidatorClient) ProposeExit(ctx context.Context, in *VoluntaryExit, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.BeaconNodeValidator/ProposeExit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BeaconNodeValidatorServer is the server API for BeaconNodeValidator service.
type BeaconNodeValidatorServer interface {
	GetDuties(context.Context, *DutiesRequest) (*DutiesResponse, error)
	GetBlock(context.Context, *BlockRequest) (*BeaconBlock, error)
	ProposeBlock(context.Context, *BlockRequest) (*empty.Empty, error)
	GetAttestationData(context.Context, *AttestationDataRequest) (*Attestation, error)
	ProposeAttestation(context.Context, *Attestation) (*empty.Empty, error)
	ProposeExit(context.Context, *VoluntaryExit) (*empty.Empty, error)
}

func RegisterBeaconNodeValidatorServer(s *grpc.Server, srv BeaconNodeValidatorServer) {
	s.RegisterService(&_BeaconNodeValidator_serviceDesc, srv)
}

func _BeaconNodeValidator_GetDuties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DutiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconNodeValidatorServer).GetDuties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.BeaconNodeValidator/GetDuties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconNodeValidatorServer).GetDuties(ctx, req.(*DutiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconNodeValidator_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconNodeValidatorServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.BeaconNodeValidator/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconNodeValidatorServer).GetBlock(ctx, req.(*BlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconNodeValidator_ProposeBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconNodeValidatorServer).ProposeBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.BeaconNodeValidator/ProposeBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconNodeValidatorServer).ProposeBlock(ctx, req.(*BlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconNodeValidator_GetAttestationData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttestationDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconNodeValidatorServer).GetAttestationData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.BeaconNodeValidator/GetAttestationData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconNodeValidatorServer).GetAttestationData(ctx, req.(*AttestationDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconNodeValidator_ProposeAttestation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Attestation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconNodeValidatorServer).ProposeAttestation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.BeaconNodeValidator/ProposeAttestation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconNodeValidatorServer).ProposeAttestation(ctx, req.(*Attestation))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconNodeValidator_ProposeExit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoluntaryExit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconNodeValidatorServer).ProposeExit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.BeaconNodeValidator/ProposeExit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconNodeValidatorServer).ProposeExit(ctx, req.(*VoluntaryExit))
	}
	return interceptor(ctx, in, info, handler)
}

var _BeaconNodeValidator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethereum.eth.v1alpha1.BeaconNodeValidator",
	HandlerType: (*BeaconNodeValidatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDuties",
			Handler:    _BeaconNodeValidator_GetDuties_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _BeaconNodeValidator_GetBlock_Handler,
		},
		{
			MethodName: "ProposeBlock",
			Handler:    _BeaconNodeValidator_ProposeBlock_Handler,
		},
		{
			MethodName: "GetAttestationData",
			Handler:    _BeaconNodeValidator_GetAttestationData_Handler,
		},
		{
			MethodName: "ProposeAttestation",
			Handler:    _BeaconNodeValidator_ProposeAttestation_Handler,
		},
		{
			MethodName: "ProposeExit",
			Handler:    _BeaconNodeValidator_ProposeExit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eth/v1alpha1/validator.proto",
}
