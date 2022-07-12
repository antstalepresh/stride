// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stakeibc/params.proto

package types

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

// Params defines the parameters for the module.
// next id: 10
type Params struct {
	// define epoch lengths, in stride_epochs
	RewardsInterval        uint64 `protobuf:"varint,1,opt,name=rewards_interval,json=rewardsInterval,proto3" json:"rewards_interval,omitempty"`
	DelegateInterval       uint64 `protobuf:"varint,6,opt,name=delegate_interval,json=delegateInterval,proto3" json:"delegate_interval,omitempty"`
	DepositInterval        uint64 `protobuf:"varint,2,opt,name=deposit_interval,json=depositInterval,proto3" json:"deposit_interval,omitempty"`
	RedemptionRateInterval uint64 `protobuf:"varint,3,opt,name=redemption_rate_interval,json=redemptionRateInterval,proto3" json:"redemption_rate_interval,omitempty"`
	StrideCommission       uint64 `protobuf:"varint,4,opt,name=stride_commission,json=strideCommission,proto3" json:"stride_commission,omitempty"`
	// zone_com_address stores which addresses to
	// send the Stride commission too, as well as what portion
	// of the fee each address is entitled to
	// TODO implement this
	ZoneComAddress                map[string]string `protobuf:"bytes,5,rep,name=zone_com_address,json=zoneComAddress,proto3" json:"zone_com_address,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ReinvestInterval              uint64            `protobuf:"varint,7,opt,name=reinvest_interval,json=reinvestInterval,proto3" json:"reinvest_interval,omitempty"`
	ValidatorRebalancingThreshold uint64            `protobuf:"varint,8,opt,name=validator_rebalancing_threshold,json=validatorRebalancingThreshold,proto3" json:"validator_rebalancing_threshold,omitempty"`
	IcaTimeoutNanos               uint64            `protobuf:"varint,9,opt,name=ica_timeout_nanos,json=icaTimeoutNanos,proto3" json:"ica_timeout_nanos,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f5fe1d2f7ac763, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetRewardsInterval() uint64 {
	if m != nil {
		return m.RewardsInterval
	}
	return 0
}

func (m *Params) GetDelegateInterval() uint64 {
	if m != nil {
		return m.DelegateInterval
	}
	return 0
}

func (m *Params) GetDepositInterval() uint64 {
	if m != nil {
		return m.DepositInterval
	}
	return 0
}

func (m *Params) GetRedemptionRateInterval() uint64 {
	if m != nil {
		return m.RedemptionRateInterval
	}
	return 0
}

func (m *Params) GetStrideCommission() uint64 {
	if m != nil {
		return m.StrideCommission
	}
	return 0
}

func (m *Params) GetZoneComAddress() map[string]string {
	if m != nil {
		return m.ZoneComAddress
	}
	return nil
}

func (m *Params) GetReinvestInterval() uint64 {
	if m != nil {
		return m.ReinvestInterval
	}
	return 0
}

func (m *Params) GetValidatorRebalancingThreshold() uint64 {
	if m != nil {
		return m.ValidatorRebalancingThreshold
	}
	return 0
}

func (m *Params) GetIcaTimeoutNanos() uint64 {
	if m != nil {
		return m.IcaTimeoutNanos
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "Stridelabs.stride.stakeibc.Params")
	proto.RegisterMapType((map[string]string)(nil), "Stridelabs.stride.stakeibc.Params.ZoneComAddressEntry")
}

func init() { proto.RegisterFile("stakeibc/params.proto", fileDescriptor_41f5fe1d2f7ac763) }

var fileDescriptor_41f5fe1d2f7ac763 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe3, 0x26, 0x0d, 0x64, 0x91, 0xc0, 0x31, 0x05, 0x59, 0x91, 0x70, 0x2b, 0x4e, 0x2d,
	0x15, 0xb6, 0x04, 0x12, 0xaa, 0x7a, 0x2b, 0x08, 0x04, 0x12, 0x42, 0xc8, 0xf4, 0xd4, 0xcb, 0x32,
	0xf6, 0x8e, 0x9c, 0x55, 0xed, 0x5d, 0x6b, 0x77, 0x13, 0x48, 0x9f, 0x82, 0x23, 0x47, 0x9e, 0x84,
	0x33, 0xc7, 0x1e, 0x39, 0xa2, 0xe4, 0x45, 0x90, 0x77, 0x13, 0xbb, 0x91, 0xe0, 0x36, 0xfe, 0xe7,
	0x9b, 0xf1, 0xcc, 0xbf, 0x43, 0x1e, 0x68, 0x03, 0x97, 0xc8, 0xb3, 0x3c, 0xa9, 0x41, 0x41, 0xa5,
	0xe3, 0x5a, 0x49, 0x23, 0x83, 0xc9, 0x27, 0xa3, 0x38, 0xc3, 0x12, 0x32, 0x1d, 0x6b, 0x1b, 0xc6,
	0x1b, 0x70, 0xb2, 0x57, 0xc8, 0x42, 0x5a, 0x2c, 0x69, 0x22, 0x57, 0xf1, 0xf8, 0xe7, 0x80, 0x0c,
	0x3f, 0xda, 0x16, 0xc1, 0x11, 0xf1, 0x15, 0x7e, 0x01, 0xc5, 0x34, 0xe5, 0xc2, 0xa0, 0x9a, 0x43,
	0x19, 0x7a, 0x07, 0xde, 0xe1, 0x20, 0xbd, 0xb7, 0xd6, 0xdf, 0xad, 0xe5, 0xe0, 0x98, 0x8c, 0x19,
	0x96, 0x58, 0x80, 0xc1, 0x8e, 0x1d, 0x5a, 0xd6, 0xdf, 0x24, 0x5a, 0xf8, 0x88, 0xf8, 0x0c, 0x6b,
	0xa9, 0xb9, 0xe9, 0xd8, 0x1d, 0xd7, 0x77, 0xad, 0xb7, 0xe8, 0x09, 0x09, 0x15, 0x32, 0xac, 0x6a,
	0xc3, 0xa5, 0xa0, 0x6a, 0xab, 0x7d, 0xdf, 0x96, 0x3c, 0xec, 0xf2, 0xe9, 0xcd, 0x9f, 0x1c, 0x93,
	0xb1, 0x5b, 0x98, 0xe6, 0xb2, 0xaa, 0xb8, 0xd6, 0x5c, 0x8a, 0x70, 0xe0, 0x26, 0x72, 0x89, 0x57,
	0xad, 0x1e, 0x7c, 0x26, 0xfe, 0x95, 0x14, 0x16, 0xa5, 0xc0, 0x98, 0x42, 0xad, 0xc3, 0xdd, 0x83,
	0xfe, 0xe1, 0x9d, 0x67, 0x2f, 0xe2, 0xff, 0x3b, 0x18, 0x3b, 0x9f, 0xe2, 0x0b, 0x29, 0x9a, 0x66,
	0x67, 0xae, 0xf0, 0xb5, 0x30, 0x6a, 0x91, 0xde, 0xbd, 0xda, 0x12, 0x9b, 0x71, 0x14, 0x72, 0x31,
	0x47, 0x7d, 0x63, 0xe9, 0x5b, 0x6e, 0x9c, 0x4d, 0xa2, 0x9d, 0xfd, 0x0d, 0xd9, 0x9f, 0x43, 0xc9,
	0x19, 0x18, 0xa9, 0xa8, 0xc2, 0x0c, 0x4a, 0x10, 0x39, 0x17, 0x05, 0x35, 0x53, 0x85, 0x7a, 0x2a,
	0x4b, 0x16, 0xde, 0xb6, 0xa5, 0x8f, 0x5a, 0x2c, 0xed, 0xa8, 0xf3, 0x0d, 0x14, 0x3c, 0x21, 0x63,
	0x9e, 0x03, 0x35, 0xbc, 0x42, 0x39, 0x33, 0x54, 0x80, 0x90, 0x3a, 0x1c, 0x39, 0xa7, 0x79, 0x0e,
	0xe7, 0x4e, 0xff, 0xd0, 0xc8, 0x93, 0x33, 0x72, 0xff, 0x1f, 0x7b, 0x04, 0x3e, 0xe9, 0x5f, 0xe2,
	0xc2, 0x3e, 0xfb, 0x28, 0x6d, 0xc2, 0x60, 0x8f, 0xec, 0xce, 0xa1, 0x9c, 0xa1, 0x7d, 0xb2, 0x51,
	0xea, 0x3e, 0x4e, 0x77, 0x4e, 0xbc, 0xd3, 0xc1, 0xf7, 0x1f, 0xfb, 0xbd, 0x97, 0x6f, 0x7f, 0x2d,
	0x23, 0xef, 0x7a, 0x19, 0x79, 0x7f, 0x96, 0x91, 0xf7, 0x6d, 0x15, 0xf5, 0xae, 0x57, 0x51, 0xef,
	0xf7, 0x2a, 0xea, 0x5d, 0xc4, 0x05, 0x37, 0xd3, 0x59, 0x16, 0xe7, 0xb2, 0x4a, 0x9c, 0xab, 0x4f,
	0xdf, 0x43, 0xa6, 0x13, 0x67, 0x6b, 0xf2, 0x35, 0x69, 0x6f, 0xd8, 0x2c, 0x6a, 0xd4, 0xd9, 0xd0,
	0x5e, 0xe4, 0xf3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x97, 0x1b, 0x06, 0x90, 0xdc, 0x02, 0x00,
	0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IcaTimeoutNanos != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.IcaTimeoutNanos))
		i--
		dAtA[i] = 0x48
	}
	if m.ValidatorRebalancingThreshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ValidatorRebalancingThreshold))
		i--
		dAtA[i] = 0x40
	}
	if m.ReinvestInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ReinvestInterval))
		i--
		dAtA[i] = 0x38
	}
	if m.DelegateInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DelegateInterval))
		i--
		dAtA[i] = 0x30
	}
	if len(m.ZoneComAddress) > 0 {
		for k := range m.ZoneComAddress {
			v := m.ZoneComAddress[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintParams(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintParams(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintParams(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.StrideCommission != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.StrideCommission))
		i--
		dAtA[i] = 0x20
	}
	if m.RedemptionRateInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RedemptionRateInterval))
		i--
		dAtA[i] = 0x18
	}
	if m.DepositInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DepositInterval))
		i--
		dAtA[i] = 0x10
	}
	if m.RewardsInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RewardsInterval))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RewardsInterval != 0 {
		n += 1 + sovParams(uint64(m.RewardsInterval))
	}
	if m.DepositInterval != 0 {
		n += 1 + sovParams(uint64(m.DepositInterval))
	}
	if m.RedemptionRateInterval != 0 {
		n += 1 + sovParams(uint64(m.RedemptionRateInterval))
	}
	if m.StrideCommission != 0 {
		n += 1 + sovParams(uint64(m.StrideCommission))
	}
	if len(m.ZoneComAddress) > 0 {
		for k, v := range m.ZoneComAddress {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovParams(uint64(len(k))) + 1 + len(v) + sovParams(uint64(len(v)))
			n += mapEntrySize + 1 + sovParams(uint64(mapEntrySize))
		}
	}
	if m.DelegateInterval != 0 {
		n += 1 + sovParams(uint64(m.DelegateInterval))
	}
	if m.ReinvestInterval != 0 {
		n += 1 + sovParams(uint64(m.ReinvestInterval))
	}
	if m.ValidatorRebalancingThreshold != 0 {
		n += 1 + sovParams(uint64(m.ValidatorRebalancingThreshold))
	}
	if m.IcaTimeoutNanos != 0 {
		n += 1 + sovParams(uint64(m.IcaTimeoutNanos))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardsInterval", wireType)
			}
			m.RewardsInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RewardsInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositInterval", wireType)
			}
			m.DepositInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DepositInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedemptionRateInterval", wireType)
			}
			m.RedemptionRateInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RedemptionRateInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StrideCommission", wireType)
			}
			m.StrideCommission = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StrideCommission |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZoneComAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ZoneComAddress == nil {
				m.ZoneComAddress = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowParams
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthParams
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthParams
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthParams
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthParams
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipParams(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthParams
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ZoneComAddress[mapkey] = mapvalue
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegateInterval", wireType)
			}
			m.DelegateInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DelegateInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReinvestInterval", wireType)
			}
			m.ReinvestInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReinvestInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorRebalancingThreshold", wireType)
			}
			m.ValidatorRebalancingThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidatorRebalancingThreshold |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IcaTimeoutNanos", wireType)
			}
			m.IcaTimeoutNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IcaTimeoutNanos |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
