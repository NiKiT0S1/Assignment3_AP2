// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: internal/proto/order/order.proto

package order

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId       int32                  `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	ProductId     int32                  `protobuf:"varint,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	mi := &file_internal_proto_order_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_order_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_internal_proto_order_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderItem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderItem) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderItem) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *OrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int32                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status        string       `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Items         []*OrderItem `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_internal_proto_order_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_order_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_internal_proto_order_order_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type OrderID struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderID) Reset() {
	*x = OrderID{}
	mi := &file_internal_proto_order_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderID) ProtoMessage() {}

func (x *OrderID) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_order_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderID.ProtoReflect.Descriptor instead.
func (*OrderID) Descriptor() ([]byte, []int) {
	return file_internal_proto_order_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderID) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListOrdersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrdersRequest) Reset() {
	*x = ListOrdersRequest{}
	mi := &file_internal_proto_order_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersRequest) ProtoMessage() {}

func (x *ListOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_order_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersRequest.ProtoReflect.Descriptor instead.
func (*ListOrdersRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_order_order_proto_rawDescGZIP(), []int{3}
}

func (x *ListOrdersRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_internal_proto_order_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_order_order_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_internal_proto_order_order_proto_rawDescGZIP(), []int{4}
}

type OrderList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Orders        []*Order               `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderList) Reset() {
	*x = OrderList{}
	mi := &file_internal_proto_order_order_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderList) ProtoMessage() {}

func (x *OrderList) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_order_order_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderList.ProtoReflect.Descriptor instead.
func (*OrderList) Descriptor() ([]byte, []int) {
	return file_internal_proto_order_order_proto_rawDescGZIP(), []int{5}
}

func (x *OrderList) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

var File_internal_proto_order_order_proto protoreflect.FileDescriptor

const file_internal_proto_order_order_proto_rawDesc = "" +
	"\n" +
	" internal/proto/order/order.proto\x12\x05order\"q\n" +
	"\tOrderItem\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x19\n" +
	"\border_id\x18\x02 \x01(\x05R\aorderId\x12\x1d\n" +
	"\n" +
	"product_id\x18\x03 \x01(\x05R\tproductId\x12\x1a\n" +
	"\bquantity\x18\x04 \x01(\x05R\bquantity\"p\n" +
	"\x05Order\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x05R\x06userId\x12\x16\n" +
	"\x06status\x18\x03 \x01(\tR\x06status\x12&\n" +
	"\x05items\x18\x04 \x03(\v2\x10.order.OrderItemR\x05items\"\x19\n" +
	"\aOrderID\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\",\n" +
	"\x11ListOrdersRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x05R\x06userId\"\a\n" +
	"\x05Empty\"1\n" +
	"\tOrderList\x12$\n" +
	"\x06orders\x18\x01 \x03(\v2\f.order.OrderR\x06orders2\xd4\x01\n" +
	"\fOrderService\x12)\n" +
	"\vCreateOrder\x12\f.order.Order\x1a\f.order.Order\x12(\n" +
	"\bGetOrder\x12\x0e.order.OrderID\x1a\f.order.Order\x12/\n" +
	"\x11UpdateOrderStatus\x12\f.order.Order\x1a\f.order.Order\x12>\n" +
	"\x10ListOrdersByUser\x12\x18.order.ListOrdersRequest\x1a\x10.order.OrderListB(Z&orderService/internal/delivery/grpc/pbb\x06proto3"

var (
	file_internal_proto_order_order_proto_rawDescOnce sync.Once
	file_internal_proto_order_order_proto_rawDescData []byte
)

func file_internal_proto_order_order_proto_rawDescGZIP() []byte {
	file_internal_proto_order_order_proto_rawDescOnce.Do(func() {
		file_internal_proto_order_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_proto_order_order_proto_rawDesc), len(file_internal_proto_order_order_proto_rawDesc)))
	})
	return file_internal_proto_order_order_proto_rawDescData
}

var file_internal_proto_order_order_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_proto_order_order_proto_goTypes = []any{
	(*OrderItem)(nil),         // 0: order.OrderItem
	(*Order)(nil),             // 1: order.Order
	(*OrderID)(nil),           // 2: order.OrderID
	(*ListOrdersRequest)(nil), // 3: order.ListOrdersRequest
	(*Empty)(nil),             // 4: order.Empty
	(*OrderList)(nil),         // 5: order.OrderList
}
var file_internal_proto_order_order_proto_depIdxs = []int32{
	0, // 0: order.Order.items:type_name -> order.OrderItem
	1, // 1: order.OrderList.orders:type_name -> order.Order
	1, // 2: order.OrderService.CreateOrder:input_type -> order.Order
	2, // 3: order.OrderService.GetOrder:input_type -> order.OrderID
	1, // 4: order.OrderService.UpdateOrderStatus:input_type -> order.Order
	3, // 5: order.OrderService.ListOrdersByUser:input_type -> order.ListOrdersRequest
	1, // 6: order.OrderService.CreateOrder:output_type -> order.Order
	1, // 7: order.OrderService.GetOrder:output_type -> order.Order
	1, // 8: order.OrderService.UpdateOrderStatus:output_type -> order.Order
	5, // 9: order.OrderService.ListOrdersByUser:output_type -> order.OrderList
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_proto_order_order_proto_init() }
func file_internal_proto_order_order_proto_init() {
	if File_internal_proto_order_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_proto_order_order_proto_rawDesc), len(file_internal_proto_order_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_order_order_proto_goTypes,
		DependencyIndexes: file_internal_proto_order_order_proto_depIdxs,
		MessageInfos:      file_internal_proto_order_order_proto_msgTypes,
	}.Build()
	File_internal_proto_order_order_proto = out.File
	file_internal_proto_order_order_proto_goTypes = nil
	file_internal_proto_order_order_proto_depIdxs = nil
}
