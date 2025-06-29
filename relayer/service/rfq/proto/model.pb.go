// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: service/rfq/model.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/celer-network/intent-rfq-relayer/relayer/service/rfqmm/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderStatus int32

const (
	// an order first enters the pending status when an MM returns a firm quote upon a user's quote request.
	OrderStatus_STATUS_PENDING OrderStatus = 0
	// reached after the user's call to the RFQ contract to deposit funds
	OrderStatus_STATUS_SRC_DEPOSITED OrderStatus = 10
	// reached only if an MM actively calls back to the RFQ server to mark the order as rejected. once marked as rejected,
	// the same order will not appear in the PendingOrders() anymore.
	// note that MMs can choose to not implement this active call and hence this status will never be reached.
	OrderStatus_STATUS_MM_REJECTED OrderStatus = 20
	// reached only if an MM actively calls back to the RFQ server to mark the order as dst executed
	// when they finish submitting the tx on the dst chain to transfer fund to user.
	// note that MMs can choose to not implement this active call and hence this status will never be reached.
	OrderStatus_STATUS_MM_DST_EXECUTED OrderStatus = 30
	// this status marks the observation of the on-chain event DstTransferred
	// this also means that msg2 is on its way but not yet arrived on the src chain
	// note that to the user, when an order reaches this status, it can be considered completed
	OrderStatus_STATUS_DST_TRANSFERRED OrderStatus = 40
	// reached only if an MM actively calls back to the RFQ server to mark the order as src executed
	// when they finish submitting the tx on the src chain to release fund to MM.
	// note that MMs can choose to not implement this active call and hence this status will never be reached.
	OrderStatus_STATUS_MM_SRC_EXECUTED OrderStatus = 50
	// this status marks the observation of the on-chain event RefundInitiated upon msg1 execution
	OrderStatus_STATUS_REFUND_INITIATED OrderStatus = 60
	// this status marks the observation of the on-chain event SrcReleased upon msg2 execution
	OrderStatus_STATUS_SRC_RELEASED OrderStatus = 70
	// this status marks the observation of the on-chain event Refunded upon msg3 execution
	OrderStatus_STATUS_REFUNDED OrderStatus = 80
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0:  "STATUS_PENDING",
		10: "STATUS_SRC_DEPOSITED",
		20: "STATUS_MM_REJECTED",
		30: "STATUS_MM_DST_EXECUTED",
		40: "STATUS_DST_TRANSFERRED",
		50: "STATUS_MM_SRC_EXECUTED",
		60: "STATUS_REFUND_INITIATED",
		70: "STATUS_SRC_RELEASED",
		80: "STATUS_REFUNDED",
	}
	OrderStatus_value = map[string]int32{
		"STATUS_PENDING":          0,
		"STATUS_SRC_DEPOSITED":    10,
		"STATUS_MM_REJECTED":      20,
		"STATUS_MM_DST_EXECUTED":  30,
		"STATUS_DST_TRANSFERRED":  40,
		"STATUS_MM_SRC_EXECUTED":  50,
		"STATUS_REFUND_INITIATED": 60,
		"STATUS_SRC_RELEASED":     70,
		"STATUS_REFUNDED":         80,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_service_rfq_model_proto_enumTypes[0].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_service_rfq_model_proto_enumTypes[0]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_service_rfq_model_proto_rawDescGZIP(), []int{0}
}

// for MM use
type PendingOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quote            *proto.Quote `protobuf:"bytes,1,opt,name=quote,proto3" json:"quote,omitempty"`
	SrcDepositTxHash string       `protobuf:"bytes,2,opt,name=src_deposit_tx_hash,json=srcDepositTxHash,proto3" json:"src_deposit_tx_hash,omitempty"`
	// indicates whether the user wants native token on the dst chain (only applicable if the dst token is a native wrap)
	DstNative bool `protobuf:"varint,3,opt,name=dst_native,json=dstNative,proto3" json:"dst_native,omitempty"`
	// unix epoch seconds
	ExecMsgCallData []byte      `protobuf:"bytes,4,opt,name=exec_msg_call_data,json=execMsgCallData,proto3" json:"exec_msg_call_data,omitempty"`
	QuoteSig        string      `protobuf:"bytes,5,opt,name=quote_sig,json=quoteSig,proto3" json:"quote_sig,omitempty"`
	Status          OrderStatus `protobuf:"varint,6,opt,name=status,proto3,enum=service.rfq.OrderStatus" json:"status,omitempty"`
}

func (x *PendingOrder) Reset() {
	*x = PendingOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rfq_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PendingOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PendingOrder) ProtoMessage() {}

func (x *PendingOrder) ProtoReflect() protoreflect.Message {
	mi := &file_service_rfq_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PendingOrder.ProtoReflect.Descriptor instead.
func (*PendingOrder) Descriptor() ([]byte, []int) {
	return file_service_rfq_model_proto_rawDescGZIP(), []int{0}
}

func (x *PendingOrder) GetQuote() *proto.Quote {
	if x != nil {
		return x.Quote
	}
	return nil
}

func (x *PendingOrder) GetSrcDepositTxHash() string {
	if x != nil {
		return x.SrcDepositTxHash
	}
	return ""
}

func (x *PendingOrder) GetDstNative() bool {
	if x != nil {
		return x.DstNative
	}
	return false
}

func (x *PendingOrder) GetExecMsgCallData() []byte {
	if x != nil {
		return x.ExecMsgCallData
	}
	return nil
}

func (x *PendingOrder) GetQuoteSig() string {
	if x != nil {
		return x.QuoteSig
	}
	return ""
}

func (x *PendingOrder) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_STATUS_PENDING
}

type UserOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quote *proto.Quote `protobuf:"bytes,1,opt,name=quote,proto3" json:"quote,omitempty"`
	// indicates whether the user wants native token on the dst chain (only applicable if the dst token is a native wrap)
	DstNative bool `protobuf:"varint,2,opt,name=dst_native,json=dstNative,proto3" json:"dst_native,omitempty"`
	// unix epoch seconds
	LastUpdated int64       `protobuf:"varint,3,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	Status      OrderStatus `protobuf:"varint,4,opt,name=status,proto3,enum=service.rfq.OrderStatus" json:"status,omitempty"`
	MmId        string      `protobuf:"bytes,5,opt,name=mm_id,json=mmId,proto3" json:"mm_id,omitempty"`
	MmAddr      string      `protobuf:"bytes,6,opt,name=mm_addr,json=mmAddr,proto3" json:"mm_addr,omitempty"`
	// the tx in which the user deposits the fund on the src chain
	SrcDepositTxHash string `protobuf:"bytes,7,opt,name=src_deposit_tx_hash,json=srcDepositTxHash,proto3" json:"src_deposit_tx_hash,omitempty"`
	// the tx in which the fund is transferred from MMs address to the user on the dst chain
	DstTransferTxHash string `protobuf:"bytes,8,opt,name=dst_transfer_tx_hash,json=dstTransferTxHash,proto3" json:"dst_transfer_tx_hash,omitempty"`
	// the tx in which the fund is released to the MM on the src chain
	SrcReleaseTxHash string `protobuf:"bytes,9,opt,name=src_release_tx_hash,json=srcReleaseTxHash,proto3" json:"src_release_tx_hash,omitempty"`
	// the tx in which the refund is initiated on the dst chain
	DstRefundInitTxHash string `protobuf:"bytes,10,opt,name=dst_refund_init_tx_hash,json=dstRefundInitTxHash,proto3" json:"dst_refund_init_tx_hash,omitempty"`
	// the tx in which the fund is refunded to the user on the src chain
	SrcRefundTxHash string `protobuf:"bytes,11,opt,name=src_refund_tx_hash,json=srcRefundTxHash,proto3" json:"src_refund_tx_hash,omitempty"`
}

func (x *UserOrder) Reset() {
	*x = UserOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rfq_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOrder) ProtoMessage() {}

func (x *UserOrder) ProtoReflect() protoreflect.Message {
	mi := &file_service_rfq_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOrder.ProtoReflect.Descriptor instead.
func (*UserOrder) Descriptor() ([]byte, []int) {
	return file_service_rfq_model_proto_rawDescGZIP(), []int{1}
}

func (x *UserOrder) GetQuote() *proto.Quote {
	if x != nil {
		return x.Quote
	}
	return nil
}

func (x *UserOrder) GetDstNative() bool {
	if x != nil {
		return x.DstNative
	}
	return false
}

func (x *UserOrder) GetLastUpdated() int64 {
	if x != nil {
		return x.LastUpdated
	}
	return 0
}

func (x *UserOrder) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_STATUS_PENDING
}

func (x *UserOrder) GetMmId() string {
	if x != nil {
		return x.MmId
	}
	return ""
}

func (x *UserOrder) GetMmAddr() string {
	if x != nil {
		return x.MmAddr
	}
	return ""
}

func (x *UserOrder) GetSrcDepositTxHash() string {
	if x != nil {
		return x.SrcDepositTxHash
	}
	return ""
}

func (x *UserOrder) GetDstTransferTxHash() string {
	if x != nil {
		return x.DstTransferTxHash
	}
	return ""
}

func (x *UserOrder) GetSrcReleaseTxHash() string {
	if x != nil {
		return x.SrcReleaseTxHash
	}
	return ""
}

func (x *UserOrder) GetDstRefundInitTxHash() string {
	if x != nil {
		return x.DstRefundInitTxHash
	}
	return ""
}

func (x *UserOrder) GetSrcRefundTxHash() string {
	if x != nil {
		return x.SrcRefundTxHash
	}
	return ""
}

type MarketMaker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MarketMaker) Reset() {
	*x = MarketMaker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rfq_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketMaker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketMaker) ProtoMessage() {}

func (x *MarketMaker) ProtoReflect() protoreflect.Message {
	mi := &file_service_rfq_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketMaker.ProtoReflect.Descriptor instead.
func (*MarketMaker) Descriptor() ([]byte, []int) {
	return file_service_rfq_model_proto_rawDescGZIP(), []int{2}
}

func (x *MarketMaker) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MarketMaker) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_service_rfq_model_proto protoreflect.FileDescriptor

var file_service_rfq_model_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x66, 0x71, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x72, 0x66, 0x71, 0x1a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x72, 0x66, 0x71, 0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x84, 0x02, 0x0a, 0x0c, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x2a, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x66, 0x71, 0x6d, 0x6d, 0x2e,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x13,
	0x73, 0x72, 0x63, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x5f, 0x74, 0x78, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x72, 0x63, 0x44, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x64, 0x73, 0x74, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x12, 0x2b, 0x0a, 0x12, 0x65, 0x78,
	0x65, 0x63, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x4d, 0x73, 0x67, 0x43,
	0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x5f, 0x73, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x53, 0x69, 0x67, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72,
	0x66, 0x71, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xcb, 0x03, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x66,
	0x71, 0x6d, 0x6d, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x73, 0x74, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x66, 0x71,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x13, 0x0a, 0x05, 0x6d, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x6d, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6d, 0x41, 0x64,
	0x64, 0x72, 0x12, 0x2d, 0x0a, 0x13, 0x73, 0x72, 0x63, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x5f, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x73, 0x72, 0x63, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x54, 0x78, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x2f, 0x0a, 0x14, 0x64, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x5f, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x64, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x78, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x2d, 0x0a, 0x13, 0x73, 0x72, 0x63, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x5f, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x73, 0x72, 0x63, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x78, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x34, 0x0a, 0x17, 0x64, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x5f,
	0x69, 0x6e, 0x69, 0x74, 0x5f, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x64, 0x73, 0x74, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x69,
	0x74, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2b, 0x0a, 0x12, 0x73, 0x72, 0x63, 0x5f, 0x72,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x72, 0x63, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x54, 0x78,
	0x48, 0x61, 0x73, 0x68, 0x22, 0x31, 0x0a, 0x0b, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x4d, 0x61,
	0x6b, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x2a, 0xf2, 0x01, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x52, 0x43, 0x5f, 0x44, 0x45, 0x50, 0x4f, 0x53, 0x49,
	0x54, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x4d, 0x4d, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x14, 0x12, 0x1a, 0x0a,
	0x16, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4d, 0x4d, 0x5f, 0x44, 0x53, 0x54, 0x5f, 0x45,
	0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x44, 0x10, 0x1e, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x44, 0x53, 0x54, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52,
	0x52, 0x45, 0x44, 0x10, 0x28, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x4d, 0x4d, 0x5f, 0x53, 0x52, 0x43, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x44, 0x10,
	0x32, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x46, 0x55,
	0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x54, 0x45, 0x44, 0x10, 0x3c, 0x12, 0x17,
	0x0a, 0x13, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x52, 0x43, 0x5f, 0x52, 0x45, 0x4c,
	0x45, 0x41, 0x53, 0x45, 0x44, 0x10, 0x46, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x52, 0x45, 0x46, 0x55, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x50, 0x42, 0x45, 0x5a, 0x43,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x65, 0x6c, 0x65, 0x72,
	0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x65, 0x74, 0x69, 0x2d, 0x72, 0x66,
	0x71, 0x2d, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x66, 0x71, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_rfq_model_proto_rawDescOnce sync.Once
	file_service_rfq_model_proto_rawDescData = file_service_rfq_model_proto_rawDesc
)

func file_service_rfq_model_proto_rawDescGZIP() []byte {
	file_service_rfq_model_proto_rawDescOnce.Do(func() {
		file_service_rfq_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_rfq_model_proto_rawDescData)
	})
	return file_service_rfq_model_proto_rawDescData
}

var file_service_rfq_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_service_rfq_model_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_rfq_model_proto_goTypes = []interface{}{
	(OrderStatus)(0),     // 0: service.rfq.OrderStatus
	(*PendingOrder)(nil), // 1: service.rfq.PendingOrder
	(*UserOrder)(nil),    // 2: service.rfq.UserOrder
	(*MarketMaker)(nil),  // 3: service.rfq.MarketMaker
	(*proto.Quote)(nil),  // 4: service.rfqmm.Quote
}
var file_service_rfq_model_proto_depIdxs = []int32{
	4, // 0: service.rfq.PendingOrder.quote:type_name -> service.rfqmm.Quote
	0, // 1: service.rfq.PendingOrder.status:type_name -> service.rfq.OrderStatus
	4, // 2: service.rfq.UserOrder.quote:type_name -> service.rfqmm.Quote
	0, // 3: service.rfq.UserOrder.status:type_name -> service.rfq.OrderStatus
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_service_rfq_model_proto_init() }
func file_service_rfq_model_proto_init() {
	if File_service_rfq_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_rfq_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PendingOrder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_rfq_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOrder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_rfq_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketMaker); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_rfq_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_rfq_model_proto_goTypes,
		DependencyIndexes: file_service_rfq_model_proto_depIdxs,
		EnumInfos:         file_service_rfq_model_proto_enumTypes,
		MessageInfos:      file_service_rfq_model_proto_msgTypes,
	}.Build()
	File_service_rfq_model_proto = out.File
	file_service_rfq_model_proto_rawDesc = nil
	file_service_rfq_model_proto_goTypes = nil
	file_service_rfq_model_proto_depIdxs = nil
}
