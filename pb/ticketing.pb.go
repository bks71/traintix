// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: pb/ticketing.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Ticketing domain model
type Passenger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Passenger) Reset() {
	*x = Passenger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_ticketing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Passenger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Passenger) ProtoMessage() {}

func (x *Passenger) ProtoReflect() protoreflect.Message {
	mi := &file_pb_ticketing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Passenger.ProtoReflect.Descriptor instead.
func (*Passenger) Descriptor() ([]byte, []int) {
	return file_pb_ticketing_proto_rawDescGZIP(), []int{0}
}

func (x *Passenger) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Passenger) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Passenger) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Seat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number  int32  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Section string `protobuf:"bytes,2,opt,name=section,proto3" json:"section,omitempty"`
}

func (x *Seat) Reset() {
	*x = Seat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_ticketing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Seat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Seat) ProtoMessage() {}

func (x *Seat) ProtoReflect() protoreflect.Message {
	mi := &file_pb_ticketing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Seat.ProtoReflect.Descriptor instead.
func (*Seat) Descriptor() ([]byte, []int) {
	return file_pb_ticketing_proto_rawDescGZIP(), []int{1}
}

func (x *Seat) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Seat) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

type Reservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From      string     `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To        string     `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Passenger *Passenger `protobuf:"bytes,3,opt,name=passenger,proto3" json:"passenger,omitempty"`
	Price     float32    `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Seat      *Seat      `protobuf:"bytes,5,opt,name=seat,proto3" json:"seat,omitempty"`
}

func (x *Reservation) Reset() {
	*x = Reservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_ticketing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reservation) ProtoMessage() {}

func (x *Reservation) ProtoReflect() protoreflect.Message {
	mi := &file_pb_ticketing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reservation.ProtoReflect.Descriptor instead.
func (*Reservation) Descriptor() ([]byte, []int) {
	return file_pb_ticketing_proto_rawDescGZIP(), []int{2}
}

func (x *Reservation) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Reservation) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Reservation) GetPassenger() *Passenger {
	if x != nil {
		return x.Passenger
	}
	return nil
}

func (x *Reservation) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Reservation) GetSeat() *Seat {
	if x != nil {
		return x.Seat
	}
	return nil
}

// Request messages
type PurchaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Passenger *Passenger `protobuf:"bytes,1,opt,name=passenger,proto3" json:"passenger,omitempty"`
}

func (x *PurchaseRequest) Reset() {
	*x = PurchaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_ticketing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseRequest) ProtoMessage() {}

func (x *PurchaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_ticketing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseRequest.ProtoReflect.Descriptor instead.
func (*PurchaseRequest) Descriptor() ([]byte, []int) {
	return file_pb_ticketing_proto_rawDescGZIP(), []int{3}
}

func (x *PurchaseRequest) GetPassenger() *Passenger {
	if x != nil {
		return x.Passenger
	}
	return nil
}

type GetReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GetReservationRequest) Reset() {
	*x = GetReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_ticketing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReservationRequest) ProtoMessage() {}

func (x *GetReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_ticketing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReservationRequest.ProtoReflect.Descriptor instead.
func (*GetReservationRequest) Descriptor() ([]byte, []int) {
	return file_pb_ticketing_proto_rawDescGZIP(), []int{4}
}

func (x *GetReservationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetReservationsBySectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section string `protobuf:"bytes,1,opt,name=section,proto3" json:"section,omitempty"`
}

func (x *GetReservationsBySectionRequest) Reset() {
	*x = GetReservationsBySectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_ticketing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReservationsBySectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReservationsBySectionRequest) ProtoMessage() {}

func (x *GetReservationsBySectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_ticketing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReservationsBySectionRequest.ProtoReflect.Descriptor instead.
func (*GetReservationsBySectionRequest) Descriptor() ([]byte, []int) {
	return file_pb_ticketing_proto_rawDescGZIP(), []int{5}
}

func (x *GetReservationsBySectionRequest) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

type CancelReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CancelReservationRequest) Reset() {
	*x = CancelReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_ticketing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelReservationRequest) ProtoMessage() {}

func (x *CancelReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_ticketing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelReservationRequest.ProtoReflect.Descriptor instead.
func (*CancelReservationRequest) Descriptor() ([]byte, []int) {
	return file_pb_ticketing_proto_rawDescGZIP(), []int{6}
}

func (x *CancelReservationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ChangeSeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Seat  *Seat  `protobuf:"bytes,2,opt,name=seat,proto3" json:"seat,omitempty"` //requested seat
}

func (x *ChangeSeatRequest) Reset() {
	*x = ChangeSeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_ticketing_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeSeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeSeatRequest) ProtoMessage() {}

func (x *ChangeSeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_ticketing_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeSeatRequest.ProtoReflect.Descriptor instead.
func (*ChangeSeatRequest) Descriptor() ([]byte, []int) {
	return file_pb_ticketing_proto_rawDescGZIP(), []int{7}
}

func (x *ChangeSeatRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ChangeSeatRequest) GetSeat() *Seat {
	if x != nil {
		return x.Seat
	}
	return nil
}

// Response messages
type TicketingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservation []*Reservation `protobuf:"bytes,1,rep,name=reservation,proto3" json:"reservation,omitempty"`
}

func (x *TicketingResponse) Reset() {
	*x = TicketingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_ticketing_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketingResponse) ProtoMessage() {}

func (x *TicketingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_ticketing_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketingResponse.ProtoReflect.Descriptor instead.
func (*TicketingResponse) Descriptor() ([]byte, []int) {
	return file_pb_ticketing_proto_rawDescGZIP(), []int{8}
}

func (x *TicketingResponse) GetReservation() []*Reservation {
	if x != nil {
		return x.Reservation
	}
	return nil
}

var File_pb_ticketing_proto protoreflect.FileDescriptor

var file_pb_ticketing_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x5d, 0x0a, 0x09, 0x50, 0x61, 0x73, 0x73,
	0x65, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x38, 0x0a, 0x04, 0x53, 0x65, 0x61, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x92, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x2b, 0x0a, 0x09, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61,
	0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x09, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x73, 0x65, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x74,
	0x52, 0x04, 0x73, 0x65, 0x61, 0x74, 0x22, 0x3e, 0x0a, 0x0f, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x09, 0x70, 0x61, 0x73,
	0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x09, 0x70, 0x61, 0x73,
	0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x3b, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x18, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x47, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1c, 0x0a, 0x04, 0x73, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x52, 0x04, 0x73, 0x65, 0x61, 0x74, 0x22, 0x46, 0x0a,
	0x11, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xef, 0x02, 0x0a, 0x09, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x38, 0x0a, 0x08, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x12,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a,
	0x11, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0a, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x53, 0x65, 0x61, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_ticketing_proto_rawDescOnce sync.Once
	file_pb_ticketing_proto_rawDescData = file_pb_ticketing_proto_rawDesc
)

func file_pb_ticketing_proto_rawDescGZIP() []byte {
	file_pb_ticketing_proto_rawDescOnce.Do(func() {
		file_pb_ticketing_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_ticketing_proto_rawDescData)
	})
	return file_pb_ticketing_proto_rawDescData
}

var file_pb_ticketing_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pb_ticketing_proto_goTypes = []interface{}{
	(*Passenger)(nil),                       // 0: pb.Passenger
	(*Seat)(nil),                            // 1: pb.Seat
	(*Reservation)(nil),                     // 2: pb.Reservation
	(*PurchaseRequest)(nil),                 // 3: pb.PurchaseRequest
	(*GetReservationRequest)(nil),           // 4: pb.GetReservationRequest
	(*GetReservationsBySectionRequest)(nil), // 5: pb.GetReservationsBySectionRequest
	(*CancelReservationRequest)(nil),        // 6: pb.CancelReservationRequest
	(*ChangeSeatRequest)(nil),               // 7: pb.ChangeSeatRequest
	(*TicketingResponse)(nil),               // 8: pb.TicketingResponse
}
var file_pb_ticketing_proto_depIdxs = []int32{
	0,  // 0: pb.Reservation.passenger:type_name -> pb.Passenger
	1,  // 1: pb.Reservation.seat:type_name -> pb.Seat
	0,  // 2: pb.PurchaseRequest.passenger:type_name -> pb.Passenger
	1,  // 3: pb.ChangeSeatRequest.seat:type_name -> pb.Seat
	2,  // 4: pb.TicketingResponse.reservation:type_name -> pb.Reservation
	3,  // 5: pb.Ticketing.Purchase:input_type -> pb.PurchaseRequest
	4,  // 6: pb.Ticketing.GetReservation:input_type -> pb.GetReservationRequest
	5,  // 7: pb.Ticketing.GetReservationsBySection:input_type -> pb.GetReservationsBySectionRequest
	6,  // 8: pb.Ticketing.CancelReservation:input_type -> pb.CancelReservationRequest
	7,  // 9: pb.Ticketing.ChangeSeat:input_type -> pb.ChangeSeatRequest
	8,  // 10: pb.Ticketing.Purchase:output_type -> pb.TicketingResponse
	8,  // 11: pb.Ticketing.GetReservation:output_type -> pb.TicketingResponse
	8,  // 12: pb.Ticketing.GetReservationsBySection:output_type -> pb.TicketingResponse
	8,  // 13: pb.Ticketing.CancelReservation:output_type -> pb.TicketingResponse
	8,  // 14: pb.Ticketing.ChangeSeat:output_type -> pb.TicketingResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_pb_ticketing_proto_init() }
func file_pb_ticketing_proto_init() {
	if File_pb_ticketing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_ticketing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Passenger); i {
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
		file_pb_ticketing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Seat); i {
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
		file_pb_ticketing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reservation); i {
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
		file_pb_ticketing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseRequest); i {
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
		file_pb_ticketing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReservationRequest); i {
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
		file_pb_ticketing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReservationsBySectionRequest); i {
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
		file_pb_ticketing_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelReservationRequest); i {
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
		file_pb_ticketing_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeSeatRequest); i {
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
		file_pb_ticketing_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketingResponse); i {
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
			RawDescriptor: file_pb_ticketing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_ticketing_proto_goTypes,
		DependencyIndexes: file_pb_ticketing_proto_depIdxs,
		MessageInfos:      file_pb_ticketing_proto_msgTypes,
	}.Build()
	File_pb_ticketing_proto = out.File
	file_pb_ticketing_proto_rawDesc = nil
	file_pb_ticketing_proto_goTypes = nil
	file_pb_ticketing_proto_depIdxs = nil
}
